package routes

import (
	"io"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"
	"template/internal/static"
	"template/pkg/config"

	"github.com/gin-gonic/gin"
)

// RegisterClientRoutes 注册前端静态文件路由
func RegisterClientRoutes(r *gin.Engine) {
	cfg := config.GetConfig()

	// 注册用户端路由(如果启用)
	if cfg.Frontend.Web.Enabled {
		registerWebRoutes(r)
	}

	// 如果所有前端模块都禁用，且启用了备用页面，则注册备用路由
	if !cfg.Frontend.Web.Enabled && cfg.Frontend.Fallback.Enabled {
		registerFallbackRoutes(r, cfg.Frontend.Fallback.Message)
	}
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

	// 子应用静态文件路由（动态支持任意子应用）
	r.GET("/subapps/:app", func(c *gin.Context) {
		serveSubAppResource(c, webFS)
	})
	r.GET("/subapps/:app/*filepath", func(c *gin.Context) {
		serveSubAppResource(c, webFS)
	})

	// 用户端SPA路由 - 使用NoRoute作为最后的fallback
	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path

		// 如果是 API、debug 或子应用路径，跳过
		if strings.HasPrefix(path, "/api/") ||
			strings.HasPrefix(path, "/debug/") ||
			strings.HasPrefix(path, "/subapps/") {
			c.Next()
			return
		}

		// 用户端默认页面（SPA路由支持）
		serveIndexHTML(c, webFS, "web")
	})
}

func serveSubAppResource(c *gin.Context, webFS fs.FS) {
	appName := strings.TrimSpace(c.Param("app"))
	if appName == "" || strings.Contains(appName, "..") || strings.Contains(appName, "/") {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	filePath := strings.TrimPrefix(c.Param("filepath"), "/")
	subAppBase := filepath.Join("subapps", appName)

	// 子应用根路径 → 返回该子应用 index.html
	if filePath == "" || filePath == "index.html" {
		subAppFS, err := fs.Sub(webFS, subAppBase)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		serveIndexHTML(c, subAppFS, appName)
		return
	}

	cleanPath := filepath.Clean(filePath)
	if cleanPath == "." || strings.HasPrefix(cleanPath, "..") {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	assetPath := filepath.Join(subAppBase, cleanPath)
	file, err := webFS.Open(assetPath)
	if err != nil {
		// SPA 深层路由回退到该子应用 index.html
		if !strings.Contains(filepath.Base(cleanPath), ".") {
			subAppFS, subErr := fs.Sub(webFS, subAppBase)
			if subErr == nil {
				serveIndexHTML(c, subAppFS, appName)
				return
			}
		}
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to read sub-app asset")
		return
	}

	c.Data(http.StatusOK, getContentType(cleanPath), content)
}

// registerFallbackRoutes 注册备用路由（当所有前端模块都禁用时）
func registerFallbackRoutes(r *gin.Engine, message string) {
	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path

		// 如果是 API 或 debug 路径，跳过
		if strings.HasPrefix(path, "/api/") ||
			strings.HasPrefix(path, "/debug/") {
			c.Next()
			return
		}

		// 返回美化的维护页面
		html := `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>服务维护中</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            display: flex;
            justify-content: center;
            align-items: center;
            position: relative;
            overflow: hidden;
        }
        
        /* 背景装饰 */
        .bg-decoration {
            position: absolute;
            width: 100%;
            height: 100%;
            opacity: 0.1;
            pointer-events: none;
        }
        
        .bg-circle {
            position: absolute;
            border-radius: 50%;
            background: rgba(255, 255, 255, 0.1);
            animation: float 6s ease-in-out infinite;
        }
        
        .bg-circle:nth-child(1) {
            width: 300px;
            height: 300px;
            top: -50px;
            left: -50px;
            animation-delay: 0s;
        }
        
        .bg-circle:nth-child(2) {
            width: 200px;
            height: 200px;
            top: 50%;
            right: -100px;
            animation-delay: 2s;
        }
        
        .bg-circle:nth-child(3) {
            width: 150px;
            height: 150px;
            bottom: -75px;
            left: 50%;
            animation-delay: 4s;
        }
        
        @keyframes float {
            0%, 100% { transform: translateY(0px) rotate(0deg); }
            33% { transform: translateY(-20px) rotate(120deg); }
            66% { transform: translateY(10px) rotate(240deg); }
        }
        
        /* 主容器 */
        .container {
            background: rgba(255, 255, 255, 0.95);
            backdrop-filter: blur(10px);
            border-radius: 20px;
            padding: 3rem 2rem;
            text-align: center;
            box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
            max-width: 480px;
            width: 90%;
            position: relative;
            animation: slideIn 0.8s ease-out;
        }
        
        @keyframes slideIn {
            from {
                opacity: 0;
                transform: translateY(30px);
            }
            to {
                opacity: 1;
                transform: translateY(0);
            }
        }
        
        /* 图标容器 */
        .icon-container {
            width: 100px;
            height: 100px;
            margin: 0 auto 2rem;
            background: linear-gradient(135deg, #ff6b6b, #feca57);
            border-radius: 50%;
            display: flex;
            align-items: center;
            justify-content: center;
            animation: pulse 2s ease-in-out infinite alternate;
            position: relative;
        }
        
        .icon-container::before {
            content: '';
            position: absolute;
            width: 120px;
            height: 120px;
            border-radius: 50%;
            background: linear-gradient(135deg, #ff6b6b, #feca57);
            opacity: 0.3;
            animation: ripple 2s ease-out infinite;
        }
        
        @keyframes pulse {
            0% { transform: scale(1); }
            100% { transform: scale(1.05); }
        }
        
        @keyframes ripple {
            0% {
                transform: scale(0.8);
                opacity: 0.6;
            }
            100% {
                transform: scale(1.2);
                opacity: 0;
            }
        }
        
        .icon {
            font-size: 3rem;
            position: relative;
            z-index: 1;
        }
        
        /* 标题样式 */
        h1 {
            color: #2c3e50;
            font-size: 2.5rem;
            font-weight: 700;
            margin-bottom: 1rem;
            background: linear-gradient(135deg, #667eea, #764ba2);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
            background-clip: text;
        }
        
        /* 消息样式 */
        .message {
            color: #5a6c7d;
            font-size: 1.1rem;
            line-height: 1.6;
            margin-bottom: 2rem;
        }
        
        /* 状态指示器 */
        .status-indicator {
            display: flex;
            align-items: center;
            justify-content: center;
            gap: 0.5rem;
            margin-bottom: 2rem;
        }
        
        .status-dot {
            width: 8px;
            height: 8px;
            border-radius: 50%;
            background: #ff6b6b;
            animation: blink 1.5s ease-in-out infinite;
        }
        
        .status-dot:nth-child(2) { animation-delay: 0.5s; }
        .status-dot:nth-child(3) { animation-delay: 1s; }
        
        @keyframes blink {
            0%, 50% { opacity: 0.3; }
            25% { opacity: 1; }
        }
        
        .status-text {
            color: #7f8c8d;
            font-size: 0.9rem;
            font-weight: 500;
        }
        
        /* 底部信息 */
        .footer {
            margin-top: 2rem;
            padding-top: 1.5rem;
            border-top: 1px solid rgba(0, 0, 0, 0.1);
            color: #95a5a6;
            font-size: 0.85rem;
        }
        
        /* 响应式设计 */
        @media (max-width: 480px) {
            .container {
                padding: 2rem 1.5rem;
                margin: 1rem;
            }
            
            h1 {
                font-size: 2rem;
            }
            
            .icon-container {
                width: 80px;
                height: 80px;
                margin-bottom: 1.5rem;
            }
            
            .icon {
                font-size: 2.5rem;
            }
            
            .message {
                font-size: 1rem;
            }
        }
        
        /* 深色模式支持 */
        @media (prefers-color-scheme: dark) {
            .container {
                background: rgba(45, 55, 72, 0.95);
                color: #e2e8f0;
            }
            
            h1 {
                color: #e2e8f0;
            }
            
            .message {
                color: #a0aec0;
            }
            
            .status-text {
                color: #718096;
            }
            
            .footer {
                color: #718096;
                border-top-color: rgba(255, 255, 255, 0.1);
            }
        }
    </style>
</head>
<body>
    <!-- 背景装饰 -->
    <div class="bg-decoration">
        <div class="bg-circle"></div>
        <div class="bg-circle"></div>
        <div class="bg-circle"></div>
    </div>
    
    <div class="container">
        <div class="icon-container">
            <div class="icon">🔧</div>
        </div>
        
        <h1>系统维护中</h1>
        
        <div class="message">
            ` + message + `
        </div>
        
        <div class="status-indicator">
            <div class="status-dot"></div>
            <div class="status-dot"></div>
            <div class="status-dot"></div>
            <span class="status-text">正在进行系统升级</span>
        </div>
        
        <div class="footer">
            <p>我们正在努力提供更好的服务体验</p>
            <p>预计很快就会恢复正常，感谢您的耐心等待</p>
        </div>
    </div>
</body>
</html>`
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
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
