package response

import (
	"template/internal/dto"
	"template/internal/models"
	"time"
)

// LoginResponse 登录响应
type LoginResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	User      UserInfo  `json:"user"`
}

// UserInfo 用户信息
type UserInfo struct {
	dto.BaseResponse
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar,omitempty"`
	Status   int    `json:"status"`
}

// FromModel 从用户模型转换
func (u *UserInfo) FromModel(user models.User) *UserInfo {
	return &UserInfo{
		BaseResponse: dto.BaseResponse{
			ID:        user.ID.String(),
			CreatedAt: time.Time(user.CreatedAt),
			UpdatedAt: time.Time(user.UpdatedAt),
		},
		Username: user.Username,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Status:   user.Status,
	}
}

// UserListResponse 用户列表响应
type UserListResponse struct {
	Users      []UserInfo `json:"users"`
	Total      int64      `json:"total"`
	Page       int        `json:"page"`
	PageSize   int        `json:"page_size"`
	TotalPages int        `json:"total_pages"`
}

// ProfileResponse 用户资料响应
type ProfileResponse struct {
	UserInfo
}

// RegisterResponse 注册响应
type RegisterResponse struct {
	Message string   `json:"message"`
	User    UserInfo `json:"user"`
}
