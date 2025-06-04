package routes

import (
	"io"
	"net/http"
	"path/filepath"
	"strings"
	"template/internal/static"

	"github.com/gin-gonic/gin"
)

// RegisterClientRoutes 注册前端静态文件路由
func RegisterClientRoutes(r *gin.Engine) {
	// 获取嵌入的静态文件系统
	adminFS := static.GetAdminDistFS()
	webFS := static.GetWebDistFS()

	// 管理端所有路由（包括静态资源和SPA路由）
	r.GET("/admin/*filepath", func(c *gin.Context) {
		filePath := strings.TrimPrefix(c.Param("filepath"), "/")

		// 如果是根路径或空路径，返回index.html
		if filePath == "" || filePath == "/" {
			file, err := adminFS.Open("index.html")
			if err != nil {
				c.String(http.StatusNotFound, "Admin index.html not found")
				return
			}
			defer file.Close()

			content, err := io.ReadAll(file)
			if err != nil {
				c.String(http.StatusInternalServerError, "Failed to read admin index.html")
				return
			}

			c.Data(http.StatusOK, "text/html; charset=utf-8", content)
			return
		}

		// 尝试打开文件
		file, err := adminFS.Open(filePath)
		if err != nil {
			// 如果文件不存在，返回admin首页（SPA路由支持）
			file, err := adminFS.Open("index.html")
			if err != nil {
				c.String(http.StatusNotFound, "Admin page not found")
				return
			}
			defer file.Close()

			content, err := io.ReadAll(file)
			if err != nil {
				c.String(http.StatusInternalServerError, "Failed to read admin index.html")
				return
			}

			c.Data(http.StatusOK, "text/html; charset=utf-8", content)
			return
		}
		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to read file")
			return
		}

		// 设置正确的Content-Type
		contentType := getContentType(filePath)
		c.Data(http.StatusOK, contentType, content)
	})

	// 用户端路由 - 根路径，但排除 API 和 admin 路径
	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path

		// 如果是 API、debug 或 admin 路径，跳过
		if strings.HasPrefix(path, "/api/") ||
			strings.HasPrefix(path, "/debug/") ||
			strings.HasPrefix(path, "/admin") {
			c.Next()
			return
		}

		// 处理用户端静态资源
		if strings.HasPrefix(path, "/assets/") {
			filePath := strings.TrimPrefix(path, "/assets/")
			assetPath := filepath.Join("assets", filePath)

			file, err := webFS.Open(assetPath)
			if err != nil {
				// 如果资源不存在，返回用户端首页（SPA路由支持）
				file, err := webFS.Open("index.html")
				if err != nil {
					c.String(http.StatusNotFound, "Page not found")
					return
				}
				defer file.Close()

				content, err := io.ReadAll(file)
				if err != nil {
					c.String(http.StatusInternalServerError, "Failed to read index.html")
					return
				}

				c.Data(http.StatusOK, "text/html; charset=utf-8", content)
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
			return
		}

		// 处理其他静态文件（如favicon.ico, vite.svg等）
		if strings.Contains(path, ".") && !strings.Contains(path, "/") {
			filename := strings.TrimPrefix(path, "/")

			file, err := webFS.Open(filename)
			if err != nil {
				// 如果文件不存在，返回用户端首页（SPA路由支持）
				file, err := webFS.Open("index.html")
				if err != nil {
					c.String(http.StatusNotFound, "Page not found")
					return
				}
				defer file.Close()

				content, err := io.ReadAll(file)
				if err != nil {
					c.String(http.StatusInternalServerError, "Failed to read index.html")
					return
				}

				c.Data(http.StatusOK, "text/html; charset=utf-8", content)
				return
			}
			defer file.Close()

			content, err := io.ReadAll(file)
			if err != nil {
				c.String(http.StatusInternalServerError, "Failed to read file")
				return
			}

			// 设置正确的Content-Type
			contentType := getContentType(filename)
			c.Data(http.StatusOK, contentType, content)
			return
		}

		// 用户端默认页面（SPA路由支持）
		file, err := webFS.Open("index.html")
		if err != nil {
			c.String(http.StatusNotFound, "Page not found")
			return
		}
		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to read index.html")
			return
		}

		c.Data(http.StatusOK, "text/html; charset=utf-8", content)
	})
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
