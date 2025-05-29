package models

import (
	"time"
)

// UploadFile 文件上传记录模型
type UploadFile struct {
	BaseModel
	Filename      string     `gorm:"not null" json:"filename"`          // 原始文件名
	StoredName    string     `gorm:"not null;unique" json:"storedName"` // 存储的文件名(UUID)
	FilePath      string     `gorm:"not null" json:"filePath"`          // 文件存储路径
	FileSize      int64      `gorm:"not null" json:"fileSize"`          // 文件大小(字节)
	MimeType      string     `gorm:"not null" json:"mimeType"`          // 文件MIME类型
	Extension     string     `gorm:"not null" json:"extension"`         // 文件扩展名
	MD5Hash       string     `gorm:"not null;index" json:"md5Hash"`     // 文件MD5哈希
	UploadStatus  int        `gorm:"default:1" json:"uploadStatus"`     // 上传状态: 1-进行中 2-完成 3-失败
	ChunkTotal    int        `gorm:"default:1" json:"chunkTotal"`       // 总分片数
	ChunkUploaded int        `gorm:"default:0" json:"chunkUploaded"`    // 已上传分片数
	UserID        string     `gorm:"index" json:"userID"`               // 上传用户ID
	UploadedAt    *time.Time `json:"uploadedAt"`                        // 上传完成时间
	// 文件访问控制
	IsPublic      bool `gorm:"default:false" json:"isPublic"`  // 是否公开（只有true/false）
	DownloadCount int  `gorm:"default:0" json:"downloadCount"` // 文件总下载次数（统计用）
}

// ChunkInfo 分片上传信息模型
type ChunkInfo struct {
	BaseModel
	FileID     string `gorm:"not null;index" json:"fileID"`    // 文件ID
	ChunkIndex int    `gorm:"not null" json:"chunkIndex"`      // 分片索引
	ChunkSize  int64  `gorm:"not null" json:"chunkSize"`       // 分片大小
	ChunkPath  string `gorm:"not null" json:"chunkPath"`       // 分片存储路径
	MD5Hash    string `gorm:"not null" json:"md5Hash"`         // 分片MD5哈希
	IsUploaded bool   `gorm:"default:false" json:"isUploaded"` // 是否已上传
}

// FileShare 文件分享模型（分享功能独立于文件的公开/私有状态）
type FileShare struct {
	BaseModel
	FileID         string     `gorm:"not null;index" json:"fileID"`         // 文件ID
	ShareToken     string     `gorm:"unique;not null" json:"shareToken"`    // 分享token
	ShareName      string     `gorm:"not null" json:"shareName"`            // 分享名称/描述
	ShareType      string     `gorm:"default:'temporary'" json:"shareType"` // 分享类型: temporary(临时), permanent(永久)
	ExpiresAt      *time.Time `json:"expiresAt"`                            // 过期时间
	AccessPassword string     `json:"accessPassword,omitempty"`             // 访问密码（可选）
	DownloadLimit  int        `gorm:"default:0" json:"downloadLimit"`       // 此分享的下载次数限制（0表示无限制）
	DownloadCount  int        `gorm:"default:0" json:"downloadCount"`       // 此分享已下载次数
	AllowPreview   bool       `gorm:"default:true" json:"allowPreview"`     // 是否允许在线预览
	AllowDownload  bool       `gorm:"default:true" json:"allowDownload"`    // 是否允许下载
	CreatedBy      string     `gorm:"not null" json:"createdBy"`            // 创建者ID
	IsActive       bool       `gorm:"default:true" json:"isActive"`         // 是否激活
	ViewCount      int        `gorm:"default:0" json:"viewCount"`           // 分享链接访问次数
}

// FilePermission 文件访问权限模型（用于私有文件的权限管理）
type FilePermission struct {
	BaseModel
	FileID     string     `gorm:"not null;index" json:"fileID"`     // 文件ID
	UserID     string     `gorm:"not null;index" json:"userID"`     // 被授权用户ID
	Permission string     `gorm:"default:'read'" json:"permission"` // 权限类型: read(查看), download(下载), manage(管理)
	GrantedBy  string     `gorm:"not null" json:"grantedBy"`        // 授权者ID
	ExpiresAt  *time.Time `json:"expiresAt"`                        // 权限过期时间（可选）
	IsActive   bool       `gorm:"default:true" json:"isActive"`     // 是否激活
}

// TemporaryAccess 临时访问令牌模型（用于系统生成的临时链接）
type TemporaryAccess struct {
	BaseModel
	FileID      string    `gorm:"not null;index" json:"fileID"`       // 文件ID
	AccessToken string    `gorm:"unique;not null" json:"accessToken"` // 临时访问token
	Purpose     string    `gorm:"not null" json:"purpose"`            // 用途说明：email_attachment, api_response, system_share等
	ExpiresAt   time.Time `gorm:"not null" json:"expiresAt"`          // 过期时间（必须有）
	UsageLimit  int       `gorm:"default:1" json:"usageLimit"`        // 使用次数限制
	UsageCount  int       `gorm:"default:0" json:"usageCount"`        // 已使用次数
	AllowedIPs  string    `json:"allowedIPs"`                         // 允许的IP地址（可选，JSON数组格式）
	CreatedBy   string    `json:"createdBy"`                          // 创建者（系统或用户ID）
	IsActive    bool      `gorm:"default:true" json:"isActive"`       // 是否激活
}
