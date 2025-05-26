package constants

import "time"

// 应用常量
const (
	AppName    = "template"
	AppVersion = "1.0.0"
)

// 缓存键前缀
const (
	CacheKeyPrefix = "template:"
)

// 缓存键模板
const (
	CacheKeyUserInfo     = "template:user:info:%d"     // 用户信息
	CacheKeyUserToken    = "template:user:token:%s"    // 用户令牌
	CacheKeyEmailCode    = "template:email:code:%s"    // 邮箱验证码
	CacheKeyLoginAttempt = "template:login:attempt:%s" // 登录尝试次数
)

// 缓存时间常量
const (
	CacheTimeShort  = 5 * time.Minute // 短期缓存：5分钟
	CacheTimeMedium = 1 * time.Hour   // 中期缓存：1小时
	CacheTimeLong   = 24 * time.Hour  // 长期缓存：24小时
)

// 分页常量
const (
	DefaultPage     = 1   // 默认页码
	DefaultPageSize = 20  // 默认每页数量
	MaxPageSize     = 100 // 最大每页数量
)

// 文件上传常量
const (
	MaxUploadSize = 10 * 1024 * 1024 // 最大上传文件大小：10MB
)

// 允许的文件类型
var AllowedFileTypes = []string{
	"image/jpeg",
	"image/png",
	"image/gif",
	"application/pdf",
	"application/msword",
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
}

// JWT常量
const (
	JWTIssuer   = "template-app"
	JWTAudience = "template-users"
)

// 请求头常量
const (
	HeaderAuthorization = "Authorization"
	HeaderContentType   = "Content-Type"
	HeaderUserAgent     = "User-Agent"
	HeaderRequestID     = "X-Request-ID"
)

// 上下文键
const (
	ContextKeyUserID    = "user_id"
	ContextKeyUsername  = "username"
	ContextKeyRequestID = "request_id"
)
