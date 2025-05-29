package response

import "time"

// SimpleUploadResponse 简单上传响应
type SimpleUploadResponse struct {
	FileID     string    `json:"fileID"`     // 文件ID
	Filename   string    `json:"filename"`   // 原始文件名
	StoredName string    `json:"storedName"` // 存储文件名
	FileSize   int64     `json:"fileSize"`   // 文件大小
	MimeType   string    `json:"mimeType"`   // 文件类型
	Extension  string    `json:"extension"`  // 文件扩展名
	MD5Hash    string    `json:"md5Hash"`    // 文件MD5哈希
	FilePath   string    `json:"filePath"`   // 文件访问路径
	UploadedAt time.Time `json:"uploadedAt"` // 上传时间
}

// ChunkUploadInitResponse 分片上传初始化响应
type ChunkUploadInitResponse struct {
	FileID      string `json:"fileID"`      // 文件ID
	ChunkSize   int64  `json:"chunkSize"`   // 分片大小
	ChunkTotal  int    `json:"chunkTotal"`  // 总分片数
	UploadToken string `json:"uploadToken"` // 上传令牌(可选)
}

// ChunkUploadResponse 分片上传响应
type ChunkUploadResponse struct {
	FileID        string `json:"fileID"`        // 文件ID
	ChunkIndex    int    `json:"chunkIndex"`    // 分片索引
	ChunkUploaded int    `json:"chunkUploaded"` // 已上传分片数
	ChunkTotal    int    `json:"chunkTotal"`    // 总分片数
	IsCompleted   bool   `json:"isCompleted"`   // 是否完成
}

// ChunkMergeResponse 分片合并响应
type ChunkMergeResponse struct {
	FileID     string    `json:"fileID"`     // 文件ID
	Filename   string    `json:"filename"`   // 原始文件名
	StoredName string    `json:"storedName"` // 存储文件名
	FileSize   int64     `json:"fileSize"`   // 文件大小
	MimeType   string    `json:"mimeType"`   // 文件类型
	Extension  string    `json:"extension"`  // 文件扩展名
	MD5Hash    string    `json:"md5Hash"`    // 文件MD5哈希
	FilePath   string    `json:"filePath"`   // 文件访问路径
	UploadedAt time.Time `json:"uploadedAt"` // 上传完成时间
}

// UploadConfigResponse 上传配置响应
type UploadConfigResponse struct {
	MaxFileSize      int64           `json:"maxFileSize"`      // 最大文件大小
	AllowedMimeTypes map[string]bool `json:"allowedMimeTypes"` // 允许的MIME类型
	ChunkSize        int64           `json:"chunkSize"`        // 分片大小
}

// UploadProgressResponse 上传进度响应
type UploadProgressResponse struct {
	FileID        string  `json:"fileID"`        // 文件ID
	Filename      string  `json:"filename"`      // 文件名
	FileSize      int64   `json:"fileSize"`      // 文件大小
	ChunkTotal    int     `json:"chunkTotal"`    // 总分片数
	ChunkUploaded int     `json:"chunkUploaded"` // 已上传分片数
	Progress      float64 `json:"progress"`      // 上传进度百分比
	Status        int     `json:"status"`        // 上传状态
}
