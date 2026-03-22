package mail_provider

import "time"

type DetectPermissionResult struct {
	TokenType  string    `json:"token_type"`
	Scope      string    `json:"scope"`
	UseLocalIP bool      `json:"use_local_ip"`
	CheckedAt  time.Time `json:"checked_at"`
}

type MailMessage struct {
	ID           string    `json:"id"`
	Subject      string    `json:"subject"`
	FromAddress  string    `json:"from_address"`
	FromName     string    `json:"from_name"`
	ReceivedTime time.Time `json:"received_time"`
	BodyPreview  string    `json:"body_preview"`
	Body         string    `json:"body,omitempty"`
	IsRead       bool      `json:"is_read"`
}

type FetchMailboxResult struct {
	Messages []MailMessage `json:"messages"`
}

type AccountCredentials struct {
	Email        string
	Password     string
	ClientID     string
	RefreshToken string
	TokenType    string
	Folder       string
}

type Client interface {
	DetectPermission(clientID, refreshToken string) (*DetectPermissionResult, error)
	FetchMailbox(account AccountCredentials) (*FetchMailboxResult, error)
}
