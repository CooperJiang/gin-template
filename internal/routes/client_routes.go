package routes

import (
	"io"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"
	"email-manage/internal/static"

	"github.com/gin-gonic/gin"
)

// RegisterClientRoutes 注册前端静态文件路由
func RegisterClientRoutes(r *gin.Engine) {
	registerWebRoutes(r)
}

// registerWebRoutes 注册用户端路由
func registerWebRoutes(r *gin.Engine) {
	webFS := static.GetWebDistFS()

	// 用户端静态文件路由 - 处理根目录下的静态文件
	r.GET("/favicon.ico", func(c *gin.Context) {
		file, err := webFS.Open("favicon.ico")
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to read favicon")
			return
		}

		c.Data(http.StatusOK, "image/x-icon", content)
	})

	// 用户端assets路由
	r.GET("/assets/*filepath", func(c *gin.Context) {
		filePath := strings.TrimPrefix(c.Param("filepath"), "/")
		assetPath := filepath.Join("assets", filePath)

		file, err := webFS.Open(assetPath)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to read asset")
			return
		}

		// 设置正确的Content-Type
		contentType := getContentType(filePath)
		c.Data(http.StatusOK, contentType, content)
	})

	// SPA路由 - 使用NoRoute作为最后的fallback
	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path

		// 如果是 API 或 debug 路径，跳过
		if strings.HasPrefix(path, "/api/") ||
			strings.HasPrefix(path, "/debug/") {
			c.Next()
			return
		}

		// SPA路由支持
		serveIndexHTML(c, webFS, "web")
	})
}

// serveIndexHTML 服务index.html文件
func serveIndexHTML(c *gin.Context, fs fs.FS, moduleName string) {
	file, err := fs.Open("index.html")
	if err != nil {
		c.String(http.StatusNotFound, moduleName+" index.html not found")
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to read "+moduleName+" index.html")
		return
	}

	c.Data(http.StatusOK, "text/html; charset=utf-8", content)
}

// getContentType 根据文件扩展名返回正确的Content-Type
func getContentType(filePath string) string {
	ext := filepath.Ext(filePath)
	switch ext {
	case ".js":
		return "application/javascript; charset=utf-8"
	case ".css":
		return "text/css; charset=utf-8"
	case ".html":
		return "text/html; charset=utf-8"
	case ".json":
		return "application/json; charset=utf-8"
	case ".png":
		return "image/png"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".gif":
		return "image/gif"
	case ".svg":
		return "image/svg+xml"
	case ".ico":
		return "image/x-icon"
	default:
		return "application/octet-stream"
	}
}
