package routes

import (
	"net/http"
	"strings"
	"template/internal/middleware"
	"template/internal/static"

	"github.com/gin-gonic/gin"
)

// RegisterClientRoutes 注册前端静态文件路由
func RegisterClientRoutes(r *gin.Engine) {
	// 获取嵌入的静态文件系统
	distFS := static.GetDistFS()
	httpFS := http.FS(distFS)

	// 注册静态文件路由组，但排除 API 和 debug 路径
	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		// 如果是 API 或 debug 路径，跳过
		if strings.HasPrefix(path, "/api/") || strings.HasPrefix(path, "/debug/") {
			c.Next()
			return
		}
		// 使用静态文件处理中间件
		if middleware.StaticFileHandler(distFS)(c); c.Writer.Written() {
			return
		}
		// 如果静态文件不存在，返回 index.html
		c.FileFromFS("dist/index.html", httpFS)
	})
}
