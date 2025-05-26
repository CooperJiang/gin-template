package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"template/internal/routes"
	"template/pkg/config"
	"template/pkg/database"
	"template/pkg/logger"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// APITestSuite API集成测试套件
type APITestSuite struct {
	suite.Suite
	router *gin.Engine
	token  string
}

// SetupSuite 测试套件初始化
func (suite *APITestSuite) SetupSuite() {
	// 设置测试模式
	gin.SetMode(gin.TestMode)

	// 初始化配置
	config.Init("../../config.yaml")

	// 初始化日志
	logger.Init()

	// 初始化数据库
	database.Init()

	// 初始化路由
	suite.router = routes.SetupRoutes()
}

// TearDownSuite 测试套件清理
func (suite *APITestSuite) TearDownSuite() {
	// 清理测试数据
	db := database.GetDB()
	if db != nil {
		// 这里可以添加清理逻辑
		// db.Exec("DELETE FROM users WHERE email LIKE '%test%'")
	}
}

// TestHealthCheck 测试健康检查
func (suite *APITestSuite) TestHealthCheck() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 200, int(response["code"].(float64)))
}

// TestAPIHealthCheck 测试API健康检查
func (suite *APITestSuite) TestAPIHealthCheck() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/health", nil)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 200, int(response["code"].(float64)))
}

// TestUserRegistration 测试用户注册
func (suite *APITestSuite) TestUserRegistration() {
	// 先发送验证码
	verifyData := map[string]string{
		"email": "test@example.com",
	}
	verifyBody, _ := json.Marshal(verifyData)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/user/send-registration-code", bytes.NewBuffer(verifyBody))
	req.Header.Set("Content-Type", "application/json")
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)

	// 注册用户
	registerData := map[string]string{
		"username":          "testuser",
		"password":          "password123",
		"email":             "test@example.com",
		"verification_code": "123456", // 测试环境使用固定验证码
	}
	registerBody, _ := json.Marshal(registerData)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/v1/user/register", bytes.NewBuffer(registerBody))
	req.Header.Set("Content-Type", "application/json")
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 200, int(response["code"].(float64)))
}

// TestUserLogin 测试用户登录
func (suite *APITestSuite) TestUserLogin() {
	loginData := map[string]string{
		"username": "testuser",
		"password": "password123",
	}
	loginBody, _ := json.Marshal(loginData)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/user/login", bytes.NewBuffer(loginBody))
	req.Header.Set("Content-Type", "application/json")
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 200, int(response["code"].(float64)))

	// 保存token用于后续测试
	if data, ok := response["data"].(map[string]interface{}); ok {
		if token, ok := data["token"].(string); ok {
			suite.token = token
		}
	}
}

// TestGetUserProfile 测试获取用户信息
func (suite *APITestSuite) TestGetUserProfile() {
	if suite.token == "" {
		suite.T().Skip("需要先登录获取token")
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/user/profile", nil)
	req.Header.Set("Authorization", "Bearer "+suite.token)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 200, int(response["code"].(float64)))

	// 验证返回的用户信息
	if data, ok := response["data"].(map[string]interface{}); ok {
		assert.Equal(suite.T(), "testuser", data["username"])
		assert.Equal(suite.T(), "test@example.com", data["email"])
	}
}

// TestUpdateUserProfile 测试更新用户信息
func (suite *APITestSuite) TestUpdateUserProfile() {
	if suite.token == "" {
		suite.T().Skip("需要先登录获取token")
	}

	updateData := map[string]string{
		"username": "updated_testuser",
	}
	updateBody, _ := json.Marshal(updateData)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/v1/user/profile", bytes.NewBuffer(updateBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+suite.token)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 200, int(response["code"].(float64)))
}

// TestInvalidRequests 测试无效请求
func (suite *APITestSuite) TestInvalidRequests() {
	// 测试无效的登录请求
	invalidLoginData := map[string]string{
		"username": "invalid",
		"password": "invalid",
	}
	invalidLoginBody, _ := json.Marshal(invalidLoginData)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/user/login", bytes.NewBuffer(invalidLoginBody))
	req.Header.Set("Content-Type", "application/json")
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 400, w.Code)

	// 测试未授权的请求
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/v1/user/profile", nil)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 401, w.Code)

	// 测试无效Token
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/v1/user/profile", nil)
	req.Header.Set("Authorization", "Bearer invalid_token")
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 401, w.Code)
}

// TestPasswordReset 测试密码重置
func (suite *APITestSuite) TestPasswordReset() {
	// 发送重置密码验证码
	resetCodeData := map[string]string{
		"email": "test@example.com",
	}
	resetCodeBody, _ := json.Marshal(resetCodeData)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/user/send-reset-password-code", bytes.NewBuffer(resetCodeBody))
	req.Header.Set("Content-Type", "application/json")
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)

	// 重置密码
	resetData := map[string]string{
		"email":             "test@example.com",
		"verification_code": "123456",
		"new_password":      "newpassword123",
	}
	resetBody, _ := json.Marshal(resetData)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/v1/user/reset-password", bytes.NewBuffer(resetBody))
	req.Header.Set("Content-Type", "application/json")
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)

	// 使用新密码登录
	newLoginData := map[string]string{
		"username": "updated_testuser",
		"password": "newpassword123",
	}
	newLoginBody, _ := json.Marshal(newLoginData)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/v1/user/login", bytes.NewBuffer(newLoginBody))
	req.Header.Set("Content-Type", "application/json")
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)
}

// makeRequest 辅助函数：发送HTTP请求
func (suite *APITestSuite) makeRequest(method, url string, body []byte, headers map[string]string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	var req *http.Request

	if body != nil {
		req, _ = http.NewRequest(method, url, bytes.NewBuffer(body))
	} else {
		req, _ = http.NewRequest(method, url, nil)
	}

	// 设置默认Content-Type
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// 设置自定义headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	suite.router.ServeHTTP(w, req)
	return w
}

// assertResponse 辅助函数：验证响应
func (suite *APITestSuite) assertResponse(w *httptest.ResponseRecorder, expectedCode int, expectedAPICode int) map[string]interface{} {
	assert.Equal(suite.T(), expectedCode, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)

	if expectedAPICode > 0 {
		assert.Equal(suite.T(), expectedAPICode, int(response["code"].(float64)))
	}

	return response
}

// TestAPISuite 运行API测试套件
func TestAPISuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}
