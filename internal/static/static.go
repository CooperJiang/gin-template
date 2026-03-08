package static

import (
	"embed"
	"io/fs"
	"template/pkg/config"
)

// WebDistDir 用户端静态文件
// 如果web目录不存在，embed会创建空的文件系统
//
//go:embed web
var WebDistDir embed.FS

// GetWebDistFS 返回嵌入的用户端静态文件系统
func GetWebDistFS() fs.FS {
	// 检查配置，如果web模块被禁用，返回空文件系统
	cfg := config.GetConfig()
	if cfg != nil && !cfg.Frontend.Web.Enabled {
		return embed.FS{}
	}

	webFS, err := fs.Sub(WebDistDir, "web")
	if err != nil {
		// 如果web目录不存在，返回空的文件系统
		return embed.FS{}
	}
	return webFS
}

// IsWebEnabled 检查用户端是否启用
func IsWebEnabled() bool {
	cfg := config.GetConfig()
	return cfg != nil && cfg.Frontend.Web.Enabled
}
