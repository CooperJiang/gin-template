package mail_provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"email-manage/pkg/logger"
)

type ClientConfig struct {
	BaseURL string
}

type client struct {
	cfg ClientConfig
}

func NewClient(cfg ClientConfig) Client {
	return &client{cfg: cfg}
}

func (c *client) DetectPermission(clientID, refreshToken string) (*DetectPermissionResult, error) {
	if strings.TrimSpace(clientID) == "" || strings.TrimSpace(refreshToken) == "" {
		return nil, fmt.Errorf("client_id or refresh_token is empty")
	}

	base := strings.TrimRight(c.cfg.BaseURL, "/")
	if base == "" {
		base = "https://app.wyx66.com"
	}

	payload := map[string]string{
		"client_id":     strings.TrimSpace(clientID),
		"refresh_token": strings.TrimSpace(refreshToken),
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, base+"/detect-permission", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("detect-permission http status: %s", resp.Status)
	}

	var res struct {
		Success    bool   `json:"success"`
		TokenType  string `json:"token_type"`
		UseLocalIP bool   `json:"use_local_ip"`
		Scope      string `json:"scope"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	if !res.Success {
		return nil, fmt.Errorf("detect-permission failed")
	}

	return &DetectPermissionResult{
		TokenType:  res.TokenType,
		Scope:      res.Scope,
		UseLocalIP: res.UseLocalIP,
		CheckedAt:  time.Now(),
	}, nil
}

func (c *client) FetchMailbox(account AccountCredentials) (*FetchMailboxResult, error) {
	logger.Info("[provider.FetchMailbox] 开始 email=%s clientID=%s folder=%s tokenType=%s",
		account.Email, account.ClientID, account.Folder, account.TokenType)

	if strings.TrimSpace(account.ClientID) == "" || strings.TrimSpace(account.RefreshToken) == "" {
		logger.Error("[provider.FetchMailbox] 凭据不完整 email=%s hasClientID=%v hasRefreshToken=%v",
			account.Email, strings.TrimSpace(account.ClientID) != "", strings.TrimSpace(account.RefreshToken) != "")
		return nil, fmt.Errorf("mail provider credentials are incomplete")
	}

	base := strings.TrimRight(c.cfg.BaseURL, "/")
	if base == "" {
		base = "https://app.wyx66.com"
	}

	payload := map[string]string{
		"email_address": account.Email,
		"client_id":     account.ClientID,
		"refresh_token": account.RefreshToken,
		"folder":        account.Folder,
		"token_type":    account.TokenType,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	url := base + "/api/emails/refresh"
	logger.Info("[provider.FetchMailbox] POST %s email=%s", url, account.Email)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/plain, */*")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error("[provider.FetchMailbox] HTTP 请求失败 email=%s err=%v", account.Email, err)
		return nil, err
	}
	defer resp.Body.Close()

	logger.Info("[provider.FetchMailbox] HTTP 响应 status=%d email=%s", resp.StatusCode, account.Email)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("refresh mailbox http status: %s", resp.Status)
	}

	var res struct {
		Success bool `json:"success"`
		Data    []struct {
			ID           string `json:"id"`
			Subject      string `json:"subject"`
			FromAddress  string `json:"from_address"`
			FromName     string `json:"from_name"`
			ReceivedTime string `json:"received_time"`
			BodyPreview  string `json:"body_preview"`
			Body         string `json:"body"`
			IsRead       bool   `json:"is_read"`
		} `json:"data"`
		Message string `json:"message"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		logger.Error("[provider.FetchMailbox] 解析响应失败 email=%s err=%v", account.Email, err)
		return nil, err
	}

	logger.Info("[provider.FetchMailbox] 响应解析 success=%v message=%s 邮件数=%d email=%s",
		res.Success, res.Message, len(res.Data), account.Email)

	if !res.Success {
		logger.Error("[provider.FetchMailbox] 接口返回失败 email=%s message=%s", account.Email, res.Message)
		return nil, fmt.Errorf("refresh mailbox failed: %s", res.Message)
	}

	messages := make([]MailMessage, 0, len(res.Data))
	for i, m := range res.Data {
		t, _ := time.Parse(time.RFC3339, m.ReceivedTime)
		msg := MailMessage{
			ID:           m.ID,
			Subject:      m.Subject,
			FromAddress:  m.FromAddress,
			FromName:     m.FromName,
			ReceivedTime: t,
			BodyPreview:  m.BodyPreview,
			Body:         m.Body,
			IsRead:       m.IsRead,
		}
		messages = append(messages, msg)
		logger.Info("[provider.FetchMailbox] 原始邮件[%d] id=%s from=%s subject=%s time=%s bodyPreview前80字=%s",
			i, m.ID, m.FromAddress, m.Subject, m.ReceivedTime, truncate(m.BodyPreview, 80))
	}

	sort.Slice(messages, func(i, j int) bool {
		return messages[i].ReceivedTime.After(messages[j].ReceivedTime)
	})

	logger.Info("[provider.FetchMailbox] 排序后第一封(最新) subject=%s time=%s", safeFirst(messages), safeFirstTime(messages))

	return &FetchMailboxResult{Messages: messages}, nil
}

func truncate(s string, maxLen int) string {
	s = strings.ReplaceAll(s, "\r\n", " ")
	s = strings.ReplaceAll(s, "\n", " ")
	if len(s) > maxLen {
		return s[:maxLen] + "..."
	}
	return s
}

func safeFirst(msgs []MailMessage) string {
	if len(msgs) == 0 {
		return "(empty)"
	}
	return msgs[0].Subject
}

func safeFirstTime(msgs []MailMessage) string {
	if len(msgs) == 0 {
		return "(empty)"
	}
	return msgs[0].ReceivedTime.Format(time.RFC3339)
}
