package request

// SimpleUploadRequest 简单上传请求
type SimpleUploadRequest struct {
	// 文件通过 multipart/form-data 上传，这里不需要定义文件字段
	// 但可以添加其他参数
	Description string `form:"description" json:"description"` // 文件描述
}

// ChunkUploadInitRequest 分片上传初始化请求
type ChunkUploadInitRequest struct {
	Filename  string `json:"filename" binding:"required" validate:"min=1,max=255"` // 文件名
	FileSize  int64  `json:"fileSize" binding:"required" validate:"min=1"`         // 文件总大小
	MD5Hash   string `json:"md5Hash" binding:"required" validate:"len=32"`         // 文件MD5哈希
	ChunkSize int64  `json:"chunkSize" binding:"required" validate:"min=1024"`     // 分片大小
}

// ChunkUploadRequest 分片上传请求
type ChunkUploadRequest struct {
	FileID     string `form:"fileID" json:"fileID" binding:"required"`         // 文件ID
	ChunkIndex int    `form:"chunkIndex" json:"chunkIndex" binding:"required"` // 分片索引
	MD5Hash    string `form:"md5Hash" json:"md5Hash" binding:"required"`       // 分片MD5哈希
	// 分片文件通过 multipart/form-data 上传
}

// ChunkMergeRequest 分片合并请求
type ChunkMergeRequest struct {
	FileID string `json:"fileID" binding:"required"` // 文件ID
}

// UploadConfigRequest 上传配置请求
type UploadConfigRequest struct {
	MaxFileSize      int64           `json:"maxFileSize"`      // 最大文件大小
	AllowedMimeTypes map[string]bool `json:"allowedMimeTypes"` // 允许的MIME类型
	ChunkSize        int64           `json:"chunkSize"`        // 分片大小
}
