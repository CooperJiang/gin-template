package middleware

import (
	"template/internal/models"
	"template/internal/repositories/upload"
	"template/pkg/common"
	"template/pkg/database"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// FileAuthMiddleware 文件权限验证中间件
func FileAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fileID := c.Param("fileId")
		if fileID == "" {
			fileID = c.Query("fileId")
		}
		
		if fileID == "" {
			common.BadRequest(c, "文件ID不能为空")
			c.Abort()
			return
		}

		db := database.GetDB()
		uploadRepo := upload.NewUploadRepository(db)

		file, err := uploadRepo.GetUploadFileByID(fileID)
		if err != nil {
			common.NotFound(c, "文件不存在")
			c.Abort()
			return
		}

		// 检查文件是否公开
		if file.IsPublic {
			go incrementFileDownloadCount(fileID)
			c.Set("file_info", file)
			c.Set("access_type", "public")
			c.Next()
			return
		}

		// 私有文件需要登录验证
		user, err := GetUserFromContext(c)
		if err != nil {
			common.Unauthorized(c, "此文件为私有文件，需要登录才能访问")
			c.Abort()
			return
		}

		if hasPrivateFilePermission(user.UserID, file) {
			go incrementFileDownloadCount(fileID)
			c.Set("file_info", file)
			c.Set("current_user", user)
			c.Set("access_type", "private")
			c.Next()
			return
		}

		common.Forbidden(c, "没有权限访问此文件")
		c.Abort()
	}
}

func hasPrivateFilePermission(userID string, file *models.UploadFile) bool {
	if file.UserID == userID {
		return true
	}

	db := database.GetDB()
	var permission models.FilePermission
	err := db.Where("file_id = ? AND user_id = ? AND is_active = ?", file.ID, userID, true).First(&permission).Error
	if err == nil {
		if permission.ExpiresAt == nil || permission.ExpiresAt.After(time.Now()) {
			return true
		}
	}
	
	return false
}

func incrementFileDownloadCount(fileID string) {
	db := database.GetDB()
	db.Model(&models.UploadFile{}).Where("id = ?", fileID).UpdateColumn("download_count", gorm.Expr("download_count + 1"))
}
