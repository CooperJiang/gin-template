package static

import (
	"embed"
	"io/fs"
	"template/pkg/config"
)

// AdminDistDir 管理端静态文件
// 如果admin目录不存在，embed会创建空的文件系统
//
//go:embed admin
var AdminDistDir embed.FS

// WebDistDir 用户端静态文件
// 如果web目录不存在，embed会创建空的文件系统
//
//go:embed web
var WebDistDir embed.FS

// GetAdminDistFS 返回嵌入的管理端静态文件系统
func GetAdminDistFS() fs.FS {
	// 检查配置，如果admin模块被禁用，返回空文件系统
	cfg := config.GetConfig()
	if cfg != nil && !cfg.Frontend.Admin.Enabled {
		return embed.FS{}
	}

	adminFS, err := fs.Sub(AdminDistDir, "admin")
	if err != nil {
		// 如果admin目录不存在，返回空的文件系统
		return embed.FS{}
	}
	return adminFS
}

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

// GetDistFS 返回嵌入的静态文件系统 (兼容性保留，默认返回管理端)
// @Deprecated 建议使用 GetAdminDistFS() 或 GetWebDistFS()
func GetDistFS() fs.FS {
	return GetAdminDistFS()
}

// IsAdminEnabled 检查管理端是否启用
func IsAdminEnabled() bool {
	cfg := config.GetConfig()
	return cfg != nil && cfg.Frontend.Admin.Enabled
}

// IsWebEnabled 检查用户端是否启用
func IsWebEnabled() bool {
	cfg := config.GetConfig()
	return cfg != nil && cfg.Frontend.Web.Enabled
}
