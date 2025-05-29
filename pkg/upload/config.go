package upload

import (
	"template/pkg/constants"
)

// Config 上传配置
type Config struct {
	MaxFileSize      int64           `yaml:"max_file_size"`      // 最大文件大小
	AllowedMimeTypes map[string]bool `yaml:"allowed_mime_types"` // 允许的MIME类型
	ChunkSize        int64           `yaml:"chunk_size"`         // 分片大小
	UploadDir        string          `yaml:"upload_dir"`         // 上传目录
	TempDir          string          `yaml:"temp_dir"`           // 临时目录
}

// NewDefaultConfig 创建默认配置
func NewDefaultConfig() *Config {
	return &Config{
		MaxFileSize:      constants.MaxFileSize,
		AllowedMimeTypes: constants.AllowedMimeTypes,
		ChunkSize:        constants.ChunkSize,
		UploadDir:        constants.DefaultUploadDir,
		TempDir:          constants.TempChunkDir,
	}
}

// ValidateMimeType 验证MIME类型
func (c *Config) ValidateMimeType(mimeType string) bool {
	if c.AllowedMimeTypes == nil {
		return false
	}
	return c.AllowedMimeTypes[mimeType]
}

// ValidateFileSize 验证文件大小
func (c *Config) ValidateFileSize(size int64) bool {
	return size <= c.MaxFileSize && size > 0
}
