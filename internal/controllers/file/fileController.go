package file

import (
	"os"
	"strconv"
	"template/internal/models"
	"template/pkg/common"

	"github.com/gin-gonic/gin"
)

// DownloadFile 下载文件
func DownloadFile(c *gin.Context) {
	fileInfo, exists := c.Get("file_info")
	if !exists {
		common.ServerError(c, "无法获取文件信息")
		return
	}

	file := fileInfo.(*models.UploadFile)

	// 检查文件是否存在
	if _, err := os.Stat(file.FilePath); os.IsNotExist(err) {
		common.NotFound(c, "文件不存在")
		return
	}

	// 设置响应头
	c.Header("Content-Type", file.MimeType)
	c.Header("Content-Disposition", "attachment; filename=\""+file.Filename+"\"")
	c.Header("Content-Length", strconv.FormatInt(file.FileSize, 10))

	// 返回文件内容
	c.File(file.FilePath)
}

// PreviewFile 预览文件（在线查看）
func PreviewFile(c *gin.Context) {
	fileInfo, exists := c.Get("file_info")
	if !exists {
		common.ServerError(c, "无法获取文件信息")
		return
	}

	file := fileInfo.(*models.UploadFile)

	// 检查文件是否存在
	if _, err := os.Stat(file.FilePath); os.IsNotExist(err) {
		common.NotFound(c, "文件不存在")
		return
	}

	// 设置预览响应头
	c.Header("Content-Type", file.MimeType)
	c.Header("Content-Disposition", "inline; filename=\""+file.Filename+"\"")

	// 返回文件内容
	c.File(file.FilePath)
}
