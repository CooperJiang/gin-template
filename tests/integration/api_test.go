package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"template/internal/app"
	"template/internal/middleware"
	"template/internal/models"
	"template/internal/routes"
	"template/pkg/common"
	"template/pkg/config"
	"template/pkg/database"
	apiErrors "template/pkg/errors"
	"template/pkg/logger"
	"template/pkg/utils"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	testRouter   *gin.Engine
	testAccount  string
	testEmail    string
	testPassword = "integration_test_password_123"
)

type apiResponse struct {
	Code int             `json:"code"`
	Data json.RawMessage `json:"data"`
}

type loginData struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	User         struct {
		Username string `json:"username"`
	} `json:"user"`
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	setProjectRoot()

	logger.Init()
	config.InitConfig()
	database.InitDB()

	if err := seedIntegrationUser(); err != nil {
		panic(fmt.Sprintf("seed integration user failed: %v", err))
	}

	testRouter = buildTestRouter()

	code := m.Run()

	cleanupIntegrationUser()
	_ = database.Close()
	os.Exit(code)
}

func setProjectRoot() {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return
	}

	root := filepath.Clean(filepath.Join(filepath.Dir(file), "../.."))
	_ = os.Chdir(root)
}

func buildTestRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(apiErrors.ErrorHandler())
	r.Use(middleware.CORSMiddleware())
	routes.RegisterRoutes(r, app.NewDependencies())
	return r
}

func seedIntegrationUser() error {
	now := time.Now().UnixNano()
	testAccount = fmt.Sprintf("itest_%d", now)
	testEmail = fmt.Sprintf("%s@example.com", testAccount)

	hashedPassword, err := utils.HashPassword(testPassword)
	if err != nil {
		return err
	}

	db := database.GetDB()
	return db.Create(&models.User{
		Username: testAccount,
		Email:    testEmail,
		Password: hashedPassword,
		Status:   common.UserStatusNormal,
		Role:     common.UserRoleUser,
	}).Error
}

func cleanupIntegrationUser() {
	db := database.GetDB()
	if db == nil {
		return
	}
	_ = db.Where("username = ?", testAccount).Delete(&models.User{}).Error
}

func performJSONRequest(method, path string, payload interface{}, headers map[string]string) *httptest.ResponseRecorder {
	var body *bytes.Reader
	if payload == nil {
		body = bytes.NewReader(nil)
	} else {
		encoded, _ := json.Marshal(payload)
		body = bytes.NewReader(encoded)
	}

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, body)
	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	testRouter.ServeHTTP(rec, req)
	return rec
}

func mustLoginToken(t *testing.T) string {
	t.Helper()

	data := mustLoginData(t)
	return data.Token
}

func mustLoginData(t *testing.T) loginData {
	t.Helper()

	rec := performJSONRequest(http.MethodPost, "/api/v1/user/login", map[string]string{
		"account":  testAccount,
		"password": testPassword,
	}, nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("login status = %d, body = %s", rec.Code, rec.Body.String())
	}

	var resp apiResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("unmarshal login response failed: %v", err)
	}
	if resp.Code != 200 {
		t.Fatalf("login business code = %d, body = %s", resp.Code, rec.Body.String())
	}

	var data loginData
	if err := json.Unmarshal(resp.Data, &data); err != nil {
		t.Fatalf("unmarshal login data failed: %v", err)
	}
	if data.Token == "" {
		t.Fatalf("empty token in login response: %s", rec.Body.String())
	}
	if data.RefreshToken == "" {
		t.Fatalf("empty refresh token in login response: %s", rec.Body.String())
	}

	return data
}

func TestLogin(t *testing.T) {
	rec := performJSONRequest(http.MethodPost, "/api/v1/user/login", map[string]string{
		"account":  testAccount,
		"password": testPassword,
	}, nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", rec.Code, rec.Body.String())
	}

	var resp apiResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("unmarshal response failed: %v", err)
	}
	if resp.Code != 200 {
		t.Fatalf("business code = %d, body = %s", resp.Code, rec.Body.String())
	}

	var data loginData
	if err := json.Unmarshal(resp.Data, &data); err != nil {
		t.Fatalf("unmarshal data failed: %v", err)
	}
	if data.User.Username != testAccount {
		t.Fatalf("username mismatch, got %q, want %q", data.User.Username, testAccount)
	}
}

func TestGetUserInfoUnauthorized(t *testing.T) {
	rec := performJSONRequest(http.MethodGet, "/api/v1/user/info", nil, nil)
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, body = %s", rec.Code, rec.Body.String())
	}
}

func TestGetUserInfoAuthorized(t *testing.T) {
	token := mustLoginToken(t)

	rec := performJSONRequest(http.MethodGet, "/api/v1/user/info", nil, map[string]string{
		"Authorization": "Bearer " + token,
	})

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", rec.Code, rec.Body.String())
	}

	var raw map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &raw); err != nil {
		t.Fatalf("unmarshal response failed: %v", err)
	}

	if int(raw["code"].(float64)) != 200 {
		t.Fatalf("business code = %v, body = %s", raw["code"], rec.Body.String())
	}

	data, ok := raw["data"].(map[string]interface{})
	if !ok {
		t.Fatalf("invalid data payload: %s", rec.Body.String())
	}

	if data["username"] != testAccount {
		t.Fatalf("username mismatch, got %v, want %s", data["username"], testAccount)
	}
}

func TestRefreshToken(t *testing.T) {
	login := mustLoginData(t)

	rec := performJSONRequest(http.MethodPost, "/api/v1/user/refresh-token", map[string]string{
		"refreshToken": login.RefreshToken,
	}, nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", rec.Code, rec.Body.String())
	}

	var resp apiResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("unmarshal response failed: %v", err)
	}
	if resp.Code != 200 {
		t.Fatalf("business code = %d, body = %s", resp.Code, rec.Body.String())
	}

	var data struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
	if err := json.Unmarshal(resp.Data, &data); err != nil {
		t.Fatalf("unmarshal data failed: %v", err)
	}
	if data.Token == "" || data.RefreshToken == "" {
		t.Fatalf("invalid token pair response: %s", rec.Body.String())
	}
}

func TestLogoutRevokesAccessToken(t *testing.T) {
	login := mustLoginData(t)

	logoutRec := performJSONRequest(http.MethodPost, "/api/v1/user/logout", nil, map[string]string{
		"Authorization": "Bearer " + login.Token,
	})

	if logoutRec.Code != http.StatusOK {
		t.Fatalf("logout status = %d, body = %s", logoutRec.Code, logoutRec.Body.String())
	}

	infoRec := performJSONRequest(http.MethodGet, "/api/v1/user/info", nil, map[string]string{
		"Authorization": "Bearer " + login.Token,
	})

	if infoRec.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, body = %s", infoRec.Code, infoRec.Body.String())
	}

	// logout 后 refresh token 也应失效（由 token_version 控制）
	refreshRec := performJSONRequest(http.MethodPost, "/api/v1/user/refresh-token", map[string]string{
		"refreshToken": login.RefreshToken,
	}, nil)

	if refreshRec.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, body = %s", refreshRec.Code, refreshRec.Body.String())
	}
}

func TestUpdateProfileRequiresEmailCode(t *testing.T) {
	token := mustLoginToken(t)
	newEmail := fmt.Sprintf("new_%d@example.com", time.Now().UnixNano())

	rec := performJSONRequest(http.MethodPut, "/api/v1/user/profile", map[string]string{
		"email": newEmail,
	}, map[string]string{
		"Authorization": "Bearer " + token,
	})

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, body = %s", rec.Code, rec.Body.String())
	}
}

func TestSendRegistrationCodeRateLimitedByCooldown(t *testing.T) {
	emailAddr := fmt.Sprintf("rate_%d@example.com", time.Now().UnixNano())

	first := performJSONRequest(http.MethodPost, "/api/v1/user/send-registration-code", map[string]string{
		"email": emailAddr,
	}, nil)
	if first.Code == http.StatusTooManyRequests {
		t.Fatalf("first request should not be rate limited, body = %s", first.Body.String())
	}

	second := performJSONRequest(http.MethodPost, "/api/v1/user/send-registration-code", map[string]string{
		"email": emailAddr,
	}, nil)
	if second.Code != http.StatusTooManyRequests {
		t.Fatalf("status = %d, body = %s", second.Code, second.Body.String())
	}
}
