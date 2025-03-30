package middleware

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// isStaticAsset 判断是否是静态资源请求
func isStaticAsset(filePath string) bool {
	ext := strings.ToLower(filepath.Ext(filePath))
	staticExts := []string{
		".js", ".css", ".png", ".jpg", ".jpeg", ".gif",
		".svg", ".ico", ".woff", ".woff2", ".ttf", ".eot",
		".pdf", ".mp4", ".webp", ".map",
	}

	for _, staticExt := range staticExts {
		if ext == staticExt {
			return true
		}
	}
	return false
}

// getCacheableExtensions 获取可缓存的文件扩展名及其缓存时间（单位：秒）
func getCacheableExtensions() map[string]int {
	return map[string]int{
		".js":    2592000, // 30天: 30*24*60*60
		".css":   2592000, // 30天
		".png":   2592000, // 30天
		".jpg":   2592000, // 30天
		".jpeg":  2592000, // 30天
		".gif":   2592000, // 30天
		".svg":   2592000, // 30天
		".ico":   2592000, // 30天
		".woff":  2592000, // 30天
		".woff2": 2592000, // 30天
		".ttf":   2592000, // 30天
		".eot":   2592000, // 30天
		".webp":  2592000, // 30天
		".map":   86400,   // 1天: 24*60*60
		".pdf":   604800,  // 7天: 7*24*60*60
		".mp4":   604800,  // 7天
	}
}

// calculateETag 计算文件的ETag
func calculateETag(file fs.File) (string, error) {
	// 获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		return "", err
	}

	// 使用文件大小和修改时间生成ETag
	// 这种方法不需要读取文件内容，更高效
	modTime := fileInfo.ModTime().Unix()
	size := fileInfo.Size()
	etag := fmt.Sprintf("%x-%x", modTime, size)

	return etag, nil
}

// StaticFileHandler 创建一个处理静态文件的中间件
func StaticFileHandler(subFS fs.FS) gin.HandlerFunc {
	fileServer := http.FileServer(http.FS(subFS))
	cacheableExts := getCacheableExtensions()

	return func(c *gin.Context) {
		// 获取请求路径
		reqPath := strings.TrimPrefix(c.Request.URL.Path, "/")
		if reqPath == "" {
			reqPath = "index.html"
		}

		// 检查文件是否存在
		file, err := subFS.Open(reqPath)
		if err != nil {
			// 如果是静态资源请求但文件不存在，返回 404
			if isStaticAsset(reqPath) {
				c.Status(http.StatusNotFound)
				return
			}

			// 非静态资源的请求都返回 index.html（用于 SPA 前端路由）
			c.Request.URL.Path = "/"
			fileServer.ServeHTTP(c.Writer, c.Request)
			return
		}
		defer file.Close()

		// 获取文件信息
		fileInfo, err := file.Stat()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		// 设置 MIME 类型
		ext := strings.ToLower(filepath.Ext(reqPath))
		switch ext {
		case ".js":
			c.Writer.Header().Set("Content-Type", "application/javascript")
		case ".css":
			c.Writer.Header().Set("Content-Type", "text/css")
		case ".html":
			c.Writer.Header().Set("Content-Type", "text/html")
		case ".svg":
			c.Writer.Header().Set("Content-Type", "image/svg+xml")
		case ".woff":
			c.Writer.Header().Set("Content-Type", "font/woff")
		case ".woff2":
			c.Writer.Header().Set("Content-Type", "font/woff2")
		case ".ttf":
			c.Writer.Header().Set("Content-Type", "font/ttf")
		}

		// 计算并设置 ETag
		etag, err := calculateETag(file)
		if err == nil {
			c.Writer.Header().Set("ETag", fmt.Sprintf(`"%s"`, etag))

			// 检查 If-None-Match 头，如果匹配则返回 304
			if match := c.GetHeader("If-None-Match"); match != "" {
				if strings.Contains(match, etag) {
					c.Writer.WriteHeader(http.StatusNotModified)
					return
				}
			}
		}

		// 设置 Last-Modified 头
		modTime := fileInfo.ModTime()
		c.Writer.Header().Set("Last-Modified", modTime.UTC().Format(http.TimeFormat))

		// 检查 If-Modified-Since 头，如果文件未修改则返回 304
		if ifModSince := c.GetHeader("If-Modified-Since"); ifModSince != "" {
			ifModSinceTime, err := time.Parse(http.TimeFormat, ifModSince)
			if err == nil && modTime.Before(ifModSinceTime.Add(time.Second)) {
				c.Writer.WriteHeader(http.StatusNotModified)
				return
			}
		}

		// 检查是否需要缓存该文件类型
		if maxAge, ok := cacheableExts[ext]; ok {
			c.Writer.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", maxAge))
			c.Writer.Header().Set("Expires", time.Now().Add(time.Duration(maxAge)*time.Second).Format(time.RFC1123))
		} else {
			// 对于不需特别缓存的资源（如index.html等），设置不缓存或短时间缓存
			c.Writer.Header().Set("Cache-Control", "no-cache, must-revalidate")
			c.Writer.Header().Set("Pragma", "no-cache")
		}

		// 由于我们已经手动处理了304逻辑，这里需要重新打开文件
		newFile, err := subFS.Open(reqPath)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		defer newFile.Close()

		// 尝试将文件转换为ReadSeeker
		readSeeker, ok := newFile.(io.ReadSeeker)
		if ok {
			// 直接写入响应
			http.ServeContent(c.Writer, c.Request, reqPath, modTime, readSeeker)
		} else {
			// 如果不能转换为ReadSeeker，则需要读取整个文件内容
			content, err := io.ReadAll(newFile)
			if err != nil {
				c.Status(http.StatusInternalServerError)
				return
			}

			// 设置Content-Length
			c.Writer.Header().Set("Content-Length", fmt.Sprintf("%d", len(content)))

			// 写入响应
			c.Writer.WriteHeader(http.StatusOK)
			c.Writer.Write(content)
		}
	}
}
