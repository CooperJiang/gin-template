package upload

import (
	"strconv"

	"template/internal/dto/request"
	"template/internal/middleware"
	uploadService "template/internal/services/upload"
	"template/pkg/common"
	"template/pkg/upload"

	"github.com/gin-gonic/gin"
)

// SimpleUpload 简单文件上传
// @Summary 简单文件上传
// @Description 上传单个文件，适用于小文件
// @Tags 文件上传
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "上传的文件"
// @Param description formData string false "文件描述"
// @Success 200 {object} response.SimpleUploadResponse
// @Failure 400 {object} errors.ErrorResponse
// @Router /upload/simple [post]
func SimpleUpload(c *gin.Context) {
	// 1. 获取当前用户
	user, err := middleware.GetUserFromContext(c)
	if err != nil {
		common.Unauthorized(c, err.Error())
		return
	}

	// 2. 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		common.BadRequest(c, "未找到上传文件")
		return
	}

	// 3. 验证请求参数（可选）
	var req request.SimpleUploadRequest
	if err := c.ShouldBind(&req); err != nil {
		common.BadRequest(c, "参数错误")
		return
	}

	// 4. 调用服务层处理上传
	result, err := uploadService.SimpleUpload(file, user.UserID)
	if err != nil {
		common.ServerError(c, err.Error())
		return
	}

	// 5. 返回成功响应
	common.Success(c, result, "文件上传成功")
}

// InitChunkUpload 初始化分片上传
// @Summary 初始化分片上传
// @Description 初始化大文件分片上传，返回上传参数
// @Tags 文件上传
// @Accept json
// @Produce json
// @Param request body request.ChunkUploadInitRequest true "初始化参数"
// @Success 200 {object} response.ChunkUploadInitResponse
// @Failure 400 {object} errors.ErrorResponse
// @Router /upload/chunk/init [post]
func InitChunkUpload(c *gin.Context) {
	// 1. 获取当前用户
	user, err := middleware.GetUserFromContext(c)
	if err != nil {
		common.Unauthorized(c, err.Error())
		return
	}

	// 2. 参数校验
	req, err := common.ValidateRequest[request.ChunkUploadInitRequest](c)
	if err != nil {
		common.BadRequest(c, err.Error())
		return
	}

	// 3. 调用服务层初始化分片上传
	result, err := uploadService.InitChunkUpload(
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

	// 4. 返回成功响应
	common.Success(c, result, "分片上传初始化成功")
}

// UploadChunk 上传分片
// @Summary 上传分片
// @Description 上传单个分片文件
// @Tags 文件上传
// @Accept multipart/form-data
// @Produce json
// @Param fileID formData string true "文件ID"
// @Param chunkIndex formData int true "分片索引"
// @Param md5Hash formData string true "分片MD5哈希"
// @Param chunk formData file true "分片文件"
// @Success 200 {object} response.ChunkUploadResponse
// @Failure 400 {object} errors.ErrorResponse
// @Router /upload/chunk [post]
func UploadChunk(c *gin.Context) {
	// 1. 验证用户登录
	_, err := middleware.GetUserFromContext(c)
	if err != nil {
		common.Unauthorized(c, err.Error())
		return
	}

	// 2. 获取表单参数
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

	// 3. 获取分片文件
	chunk, err := c.FormFile("chunk")
	if err != nil {
		common.BadRequest(c, "未找到分片文件")
		return
	}

	// 4. 调用服务层处理分片上传
	result, err := uploadService.UploadChunk(fileID, chunkIndex, md5Hash, chunk)
	if err != nil {
		common.ServerError(c, err.Error())
		return
	}

	// 5. 返回成功响应
	common.Success(c, result, "分片上传成功")
}

// MergeChunks 合并分片
// @Summary 合并分片
// @Description 合并所有分片为完整文件
// @Tags 文件上传
// @Accept json
// @Produce json
// @Param request body request.ChunkMergeRequest true "合并参数"
// @Success 200 {object} response.ChunkMergeResponse
// @Failure 400 {object} errors.ErrorResponse
// @Router /upload/chunk/merge [post]
func MergeChunks(c *gin.Context) {
	// 1. 验证用户登录
	_, err := middleware.GetUserFromContext(c)
	if err != nil {
		common.Unauthorized(c, err.Error())
		return
	}

	// 2. 参数校验
	req, err := common.ValidateRequest[request.ChunkMergeRequest](c)
	if err != nil {
		common.BadRequest(c, err.Error())
		return
	}

	// 3. 调用服务层合并分片
	result, err := uploadService.MergeChunks(req.FileID)
	if err != nil {
		common.ServerError(c, err.Error())
		return
	}

	// 4. 返回成功响应
	common.Success(c, result, "分片合并成功")
}

// GetUploadProgress 获取上传进度
// @Summary 获取上传进度
// @Description 获取文件上传进度信息
// @Tags 文件上传
// @Accept json
// @Produce json
// @Param fileID path string true "文件ID"
// @Success 200 {object} response.UploadProgressResponse
// @Failure 400 {object} errors.ErrorResponse
// @Router /upload/progress/{fileID} [get]
func GetUploadProgress(c *gin.Context) {
	// 1. 验证用户登录
	_, err := middleware.GetUserFromContext(c)
	if err != nil {
		common.Unauthorized(c, err.Error())
		return
	}

	// 2. 获取路径参数
	fileID := c.Param("fileID")
	if fileID == "" {
		common.BadRequest(c, "文件ID不能为空")
		return
	}

	// 3. 调用服务层获取进度
	result, err := uploadService.GetUploadProgress(fileID)
	if err != nil {
		common.ServerError(c, err.Error())
		return
	}

	// 4. 返回成功响应
	common.Success(c, result, "获取上传进度成功")
}

// GetUploadConfig 获取上传配置
// @Summary 获取上传配置
// @Description 获取前端上传所需的配置信息
// @Tags 文件上传
// @Accept json
// @Produce json
// @Success 200 {object} response.UploadConfigResponse
// @Router /upload/config [get]
func GetUploadConfig(c *gin.Context) {
	// 获取当前配置
	config := upload.NewDefaultConfig()

	// 返回配置信息
	result := map[string]interface{}{
		"maxFileSize":      config.MaxFileSize,
		"allowedMimeTypes": config.AllowedMimeTypes,
		"chunkSize":        config.ChunkSize,
	}

	common.Success(c, result, "获取上传配置成功")
}
