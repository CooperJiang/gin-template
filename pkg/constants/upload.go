package constants

// 上传状态常量
const (
	UploadStatusUploading = 1 // 上传中
	UploadStatusCompleted = 2 // 已完成
	UploadStatusFailed    = 3 // 上传失败
)

// 文件存储相关常量
const (
	DefaultUploadDir = "./uploads"     // 默认上传目录
	TempChunkDir     = "./uploads/tmp" // 临时分片目录
	ChunkSize        = 1024 * 1024 * 2 // 默认分片大小 2MB
)

// 文件限制常量
const (
	MaxFileSize       = 1024 * 1024 * 1024 * 5 // 默认最大文件大小 5GB
	MaxFilenameLength = 255                    // 最大文件名长度
)

// 支持的文件类型
var AllowedMimeTypes = map[string]bool{
	"image/jpeg":                   true,
	"image/png":                    true,
	"image/gif":                    true,
	"image/webp":                   true,
	"video/mp4":                    true,
	"video/avi":                    true,
	"video/mov":                    true,
	"application/pdf":              true,
	"text/plain":                   true,
	"application/zip":              true,
	"application/x-rar-compressed": true,
}

// 文件扩展名映射
var ExtensionMimeMap = map[string]string{
	".jpg":  "image/jpeg",
	".jpeg": "image/jpeg",
	".png":  "image/png",
	".gif":  "image/gif",
	".webp": "image/webp",
	".mp4":  "video/mp4",
	".avi":  "video/avi",
	".mov":  "video/mov",
	".pdf":  "application/pdf",
	".txt":  "text/plain",
	".zip":  "application/zip",
	".rar":  "application/x-rar-compressed",
}
