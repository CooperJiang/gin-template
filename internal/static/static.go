package static

import (
	"embed"
	"io/fs"
)

// WebDistDir 用户端静态文件
// 如果web目录不存在，embed会创建空的文件系统
//
//go:embed web
var WebDistDir embed.FS

// GetWebDistFS 返回嵌入的用户端静态文件系统
func GetWebDistFS() fs.FS {
	webFS, err := fs.Sub(WebDistDir, "web")
	if err != nil {
		// 如果web目录不存在，返回空的文件系统
		return embed.FS{}
	}
	return webFS
}
