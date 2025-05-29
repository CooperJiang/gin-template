package request

import "time"

// CreateShareLinkRequest 创建分享链接请求
type CreateShareLinkRequest struct {
	FileID         string     `json:"fileID" binding:"required"`                              // 文件ID
	ShareName      string     `json:"shareName" binding:"required"`                           // 分享名称
	ShareType      string     `json:"shareType" binding:"required,oneof=temporary permanent"` // 分享类型
	ExpiresAt      *time.Time `json:"expiresAt"`                                              // 过期时间
	AccessPassword string     `json:"accessPassword"`                                         // 访问密码
	DownloadLimit  int        `json:"downloadLimit"`                                          // 下载次数限制
	AllowPreview   *bool      `json:"allowPreview"`                                           // 是否允许预览
	AllowDownload  *bool      `json:"allowDownload"`                                          // 是否允许下载
}

// UpdateFileSettingsRequest 更新文件设置请求
type UpdateFileSettingsRequest struct {
	IsPublic *bool `json:"isPublic"` // 是否公开
}

// CreateTempAccessRequest 创建临时访问请求
type CreateTempAccessRequest struct {
	FileID     string    `json:"fileID" binding:"required"`    // 文件ID
	Purpose    string    `json:"purpose" binding:"required"`   // 用途说明
	ExpiresAt  time.Time `json:"expiresAt" binding:"required"` // 过期时间
	UsageLimit int       `json:"usageLimit"`                   // 使用次数限制
	AllowedIPs []string  `json:"allowedIPs"`                   // 允许的IP地址
}

// GrantFilePermissionRequest 授权文件权限请求
type GrantFilePermissionRequest struct {
	FileID     string     `json:"fileID" binding:"required"`                                // 文件ID
	UserID     string     `json:"userID" binding:"required"`                                // 被授权用户ID
	Permission string     `json:"permission" binding:"required,oneof=read download manage"` // 权限类型
	ExpiresAt  *time.Time `json:"expiresAt"`                                                // 过期时间
}
