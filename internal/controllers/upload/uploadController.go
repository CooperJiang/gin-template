package upload

import (
	"strconv"

	"email-manage/internal/dto/request"
	"email-manage/internal/middleware"
	uploadService "email-manage/internal/services/upload"
	"email-manage/pkg/common"
	"email-manage/pkg/upload"

	"github.com/gin-gonic/gin"
)

// Controller 上传控制器
type Controller struct {
	service uploadService.Service
}

var defaultController *Controller

// NewController 创建上传控制器
func NewController(service uploadService.Service) *Controller {
	if service == nil {
		service = uploadService.GetUploadService()
	}
	return &Controller{service: service}
}

func getDefaultController() *Controller {
	if defaultController == nil {
		defaultController = NewController(nil)
	}
	return defaultController
}

// SimpleUpload 简单文件上传
func (ctl *Controller) SimpleUpload(c *gin.Context) {
	user, err := middleware.GetUserFromContext(c)
	if err != nil {
		common.Unauthorized(c, err.Error())
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		common.BadRequest(c, "未找到上传文件")
		return
	}

	var req request.SimpleUploadRequest
	if err := c.ShouldBind(&req); err != nil {
		common.BadRequest(c, "参数错误")
		return
	}

	result, err := ctl.service.SimpleUpload(file, user.UserID)
	if err != nil {
		common.ServerError(c, err.Error())
		return
	}

	common.Success(c, result, "文件上传成功")
}

// InitChunkUpload 初始化分片上传
func (ctl *Controller) InitChunkUpload(c *gin.Context) {
	user, err := middleware.GetUserFromContext(c)
	if err != nil {
		common.Unauthorized(c, err.Error())
		return
	}

	req, err := common.ValidateRequest[request.ChunkUploadInitRequest](c)
	if err != nil {
		common.BadRequest(c, err.Error())
		return
	}

	result, err := ctl.service.InitChunkUpload(
		req.Filename,
		req.FileSize,
		req.MD5Hash,
		req.ChunkSize,
		user.UserID,
	)
	if err != nil {
		common.ServerError(c, err.Error())
		return
	}

	common.Success(c, result, "分片上传初始化成功")
}

// UploadChunk 上传分片
func (ctl *Controller) UploadChunk(c *gin.Context) {
	if _, err := middleware.GetUserFromContext(c); err != nil {
		common.Unauthorized(c, err.Error())
		return
	}

	fileID := c.PostForm("fileID")
	chunkIndexStr := c.PostForm("chunkIndex")
	md5Hash := c.PostForm("md5Hash")

	if fileID == "" || chunkIndexStr == "" {
		common.BadRequest(c, "缺少必要参数")
		return
	}

	chunkIndex, err := strconv.Atoi(chunkIndexStr)
	if err != nil {
		common.BadRequest(c, "分片索引格式错误")
		return
	}

	chunk, err := c.FormFile("chunk")
	if err != nil {
		common.BadRequest(c, "未找到分片文件")
		return
	}

	result, err := ctl.service.UploadChunk(fileID, chunkIndex, md5Hash, chunk)
	if err != nil {
		common.ServerError(c, err.Error())
		return
	}

	common.Success(c, result, "分片上传成功")
}

// MergeChunks 合并分片
func (ctl *Controller) MergeChunks(c *gin.Context) {
	if _, err := middleware.GetUserFromContext(c); err != nil {
		common.Unauthorized(c, err.Error())
		return
	}

	req, err := common.ValidateRequest[request.ChunkMergeRequest](c)
	if err != nil {
		common.BadRequest(c, err.Error())
		return
	}

	result, err := ctl.service.MergeChunks(req.FileID)
	if err != nil {
		common.ServerError(c, err.Error())
		return
	}

	common.Success(c, result, "分片合并成功")
}

// GetUploadProgress 获取上传进度
func (ctl *Controller) GetUploadProgress(c *gin.Context) {
	if _, err := middleware.GetUserFromContext(c); err != nil {
		common.Unauthorized(c, err.Error())
		return
	}

	fileID := c.Param("fileID")
	if fileID == "" {
		common.BadRequest(c, "文件ID不能为空")
		return
	}

	result, err := ctl.service.GetUploadProgress(fileID)
	if err != nil {
		common.ServerError(c, err.Error())
		return
	}

	common.Success(c, result, "获取上传进度成功")
}

// GetUploadConfig 获取上传配置
func (ctl *Controller) GetUploadConfig(c *gin.Context) {
	cfg := upload.NewDefaultConfig()
	result := map[string]interface{}{
		"maxFileSize":      cfg.MaxFileSize,
		"allowedMimeTypes": cfg.AllowedMimeTypes,
		"chunkSize":        cfg.ChunkSize,
	}

	common.Success(c, result, "获取上传配置成功")
}

// -------- 兼容旧路由调用 --------

func SimpleUpload(c *gin.Context) {
	getDefaultController().SimpleUpload(c)
}

func InitChunkUpload(c *gin.Context) {
	getDefaultController().InitChunkUpload(c)
}

func UploadChunk(c *gin.Context) {
	getDefaultController().UploadChunk(c)
}

func MergeChunks(c *gin.Context) {
	getDefaultController().MergeChunks(c)
}

func GetUploadProgress(c *gin.Context) {
	getDefaultController().GetUploadProgress(c)
}

func GetUploadConfig(c *gin.Context) {
	getDefaultController().GetUploadConfig(c)
}
