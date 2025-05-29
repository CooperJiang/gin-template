package upload

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"template/internal/dto/response"
	"template/internal/models"
	uploadRepo "template/internal/repositories/upload"
	"template/pkg/common"
	"template/pkg/constants"
	"template/pkg/database"
	"template/pkg/logger"
	"template/pkg/upload"
)

var uploadService *UploadService
var config *upload.Config

// UploadService 上传服务
type UploadService struct {
	uploadRepo *uploadRepo.UploadRepository
	storage    upload.Storage
	config     *upload.Config
}

// InitUploadService 初始化上传服务
func InitUploadService() {
	config = upload.NewDefaultConfig()
	db := database.GetDB()
	uploadService = &UploadService{
		uploadRepo: uploadRepo.NewUploadRepository(db),
		storage:    upload.NewLocalStorage(config.UploadDir),
		config:     config,
	}
}

// GetUploadService 获取上传服务实例
func GetUploadService() *UploadService {
	return uploadService
}

// SimpleUpload 简单文件上传
func SimpleUpload(file *multipart.FileHeader, userID string) (*response.SimpleUploadResponse, error) {
	if uploadService == nil {
		InitUploadService()
	}
	return uploadService.SimpleUpload(file, userID)
}

// InitChunkUpload 初始化分片上传
func InitChunkUpload(filename string, fileSize int64, md5Hash string, chunkSize int64, userID string) (*response.ChunkUploadInitResponse, error) {
	if uploadService == nil {
		InitUploadService()
	}
	return uploadService.InitChunkUpload(filename, fileSize, md5Hash, chunkSize, userID)
}

// UploadChunk 上传分片
func UploadChunk(fileID string, chunkIndex int, md5Hash string, chunk *multipart.FileHeader) (*response.ChunkUploadResponse, error) {
	if uploadService == nil {
		InitUploadService()
	}
	return uploadService.UploadChunk(fileID, chunkIndex, md5Hash, chunk)
}

// MergeChunks 合并分片
func MergeChunks(fileID string) (*response.ChunkMergeResponse, error) {
	if uploadService == nil {
		InitUploadService()
	}
	return uploadService.MergeChunks(fileID)
}

// GetUploadProgress 获取上传进度
func GetUploadProgress(fileID string) (*response.UploadProgressResponse, error) {
	if uploadService == nil {
		InitUploadService()
	}
	return uploadService.GetUploadProgress(fileID)
}

// SimpleUpload 简单文件上传
func (s *UploadService) SimpleUpload(file *multipart.FileHeader, userID string) (*response.SimpleUploadResponse, error) {
	// 1. 验证文件
	if err := s.validateFile(file); err != nil {
		return nil, err
	}

	// 2. 打开文件
	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("无法打开文件: %v", err)
	}
	defer src.Close()

	// 3. 计算文件MD5
	md5Hash, err := upload.CalculateFileMD5(src)
	if err != nil {
		return nil, fmt.Errorf("计算文件MD5失败: %v", err)
	}

	// 4. 检查文件是否已存在（秒传功能）
	if existingFile, err := s.uploadRepo.GetUploadFileByMD5(md5Hash); err == nil {
		logger.Info("文件已存在，执行秒传", "md5", md5Hash, "fileID", existingFile.ID.String())
		return &response.SimpleUploadResponse{
			FileID:     existingFile.ID.String(),
			Filename:   existingFile.Filename,
			StoredName: existingFile.StoredName,
			FileSize:   existingFile.FileSize,
			MimeType:   existingFile.MimeType,
			Extension:  existingFile.Extension,
			MD5Hash:    existingFile.MD5Hash,
			FilePath:   fmt.Sprintf("/files/preview/%s", existingFile.ID.String()),
			UploadedAt: *existingFile.UploadedAt,
		}, nil
	}

	// 5. 生成存储文件名
	storedName := upload.GenerateStoredFilename(file.Filename)

	// 6. 保存文件
	filePath, err := s.storage.SaveFile(storedName, src)
	if err != nil {
		return nil, fmt.Errorf("保存文件失败: %v", err)
	}

	// 7. 创建数据库记录
	now := time.Now()
	uploadFile := &models.UploadFile{
		BaseModel:     models.BaseModel{ID: common.NewUUID()},
		Filename:      file.Filename,
		StoredName:    storedName,
		FilePath:      filePath,
		FileSize:      file.Size,
		MimeType:      upload.GetMimeTypeFromExtension(file.Filename),
		Extension:     upload.GetFileExtension(file.Filename),
		MD5Hash:       md5Hash,
		UploadStatus:  constants.UploadStatusCompleted,
		ChunkTotal:    1,
		ChunkUploaded: 1,
		UserID:        userID,
		UploadedAt:    &now,
		IsPublic:      true,
	}

	if err := s.uploadRepo.CreateUploadFile(uploadFile); err != nil {
		// 如果数据库保存失败，删除已上传的文件
		s.storage.DeleteFile(filePath)
		return nil, fmt.Errorf("保存文件记录失败: %v", err)
	}

	return &response.SimpleUploadResponse{
		FileID:     uploadFile.ID.String(),
		Filename:   uploadFile.Filename,
		StoredName: uploadFile.StoredName,
		FileSize:   uploadFile.FileSize,
		MimeType:   uploadFile.MimeType,
		Extension:  uploadFile.Extension,
		MD5Hash:    uploadFile.MD5Hash,
		FilePath:   fmt.Sprintf("/files/preview/%s", uploadFile.ID.String()),
		UploadedAt: *uploadFile.UploadedAt,
	}, nil
}

// InitChunkUpload 初始化分片上传
func (s *UploadService) InitChunkUpload(filename string, fileSize int64, md5Hash string, chunkSize int64, userID string) (*response.ChunkUploadInitResponse, error) {
	// 1. 验证参数
	if err := upload.ValidateFilename(filename); err != nil {
		return nil, err
	}

	if !s.config.ValidateFileSize(fileSize) {
		return nil, fmt.Errorf("文件大小超出限制，最大允许 %d 字节", s.config.MaxFileSize)
	}

	mimeType := upload.GetMimeTypeFromExtension(filename)
	if !s.config.ValidateMimeType(mimeType) {
		return nil, fmt.Errorf("不支持的文件类型: %s", mimeType)
	}

	// 2. 检查文件是否已存在（秒传功能）
	if existingFile, err := s.uploadRepo.GetUploadFileByMD5(md5Hash); err == nil {
		logger.Info("文件已存在，执行秒传", "md5", md5Hash, "fileID", existingFile.ID.String())
		return &response.ChunkUploadInitResponse{
			FileID:     existingFile.ID.String(),
			ChunkSize:  chunkSize,
			ChunkTotal: 1,
		}, nil
	}

	// 3. 计算分片数量
	chunkTotal := upload.CalculateChunkTotal(fileSize, chunkSize)

	// 4. 创建文件记录
	fileID := common.NewUUID()
	uploadFile := &models.UploadFile{
		BaseModel:     models.BaseModel{ID: fileID},
		Filename:      filename,
		StoredName:    upload.GenerateStoredFilename(filename),
		FileSize:      fileSize,
		MimeType:      mimeType,
		Extension:     upload.GetFileExtension(filename),
		MD5Hash:       md5Hash,
		UploadStatus:  constants.UploadStatusUploading,
		ChunkTotal:    chunkTotal,
		ChunkUploaded: 0,
		UserID:        userID,
		IsPublic:      true,
	}

	if err := s.uploadRepo.CreateUploadFile(uploadFile); err != nil {
		return nil, fmt.Errorf("创建文件记录失败: %v", err)
	}

	// 5. 创建分片记录
	for i := 0; i < chunkTotal; i++ {
		chunkPath := upload.GenerateChunkPath(s.config.TempDir, fileID.String(), i)
		chunk := &models.ChunkInfo{
			BaseModel:  models.BaseModel{ID: common.NewUUID()},
			FileID:     fileID.String(),
			ChunkIndex: i,
			ChunkPath:  chunkPath,
			IsUploaded: false,
		}

		if err := s.uploadRepo.CreateChunkInfo(chunk); err != nil {
			return nil, fmt.Errorf("创建分片记录失败: %v", err)
		}
	}

	logger.Info("分片上传初始化成功", "fileID", fileID.String(), "chunkTotal", chunkTotal)

	return &response.ChunkUploadInitResponse{
		FileID:     fileID.String(),
		ChunkSize:  chunkSize,
		ChunkTotal: chunkTotal,
	}, nil
}

// UploadChunk 上传分片
func (s *UploadService) UploadChunk(fileID string, chunkIndex int, md5Hash string, chunk *multipart.FileHeader) (*response.ChunkUploadResponse, error) {
	// 1. 验证文件记录是否存在
	uploadFile, err := s.uploadRepo.GetUploadFileByID(fileID)
	if err != nil {
		return nil, fmt.Errorf("文件记录不存在: %v", err)
	}

	// 2. 检查分片是否已存在
	if s.uploadRepo.CheckChunkExists(fileID, chunkIndex) {
		// 分片已存在，直接返回成功
		count, _ := s.uploadRepo.GetUploadedChunksCount(fileID)
		return &response.ChunkUploadResponse{
			FileID:        fileID,
			ChunkIndex:    chunkIndex,
			ChunkUploaded: int(count),
			ChunkTotal:    uploadFile.ChunkTotal,
			IsCompleted:   int(count) == uploadFile.ChunkTotal,
		}, nil
	}

	// 3. 获取分片信息
	chunkInfo, err := s.uploadRepo.GetChunkInfo(fileID, chunkIndex)
	if err != nil {
		return nil, fmt.Errorf("分片信息不存在: %v", err)
	}

	// 4. 打开分片文件
	src, err := chunk.Open()
	if err != nil {
		return nil, fmt.Errorf("无法打开分片文件: %v", err)
	}
	defer src.Close()

	// 5. 验证分片MD5（可选）
	if md5Hash != "" {
		calculatedMD5, err := upload.CalculateChunkMD5(src)
		if err != nil {
			return nil, fmt.Errorf("计算分片MD5失败: %v", err)
		}
		if calculatedMD5 != md5Hash {
			return nil, fmt.Errorf("分片MD5校验失败")
		}
		src.Seek(0, 0) // 重置文件指针
	}

	// 6. 保存分片
	if err := s.storage.SaveChunk(chunkInfo.ChunkPath, src); err != nil {
		return nil, fmt.Errorf("保存分片失败: %v", err)
	}

	// 7. 更新分片状态
	chunkInfo.ChunkSize = chunk.Size
	chunkInfo.MD5Hash = md5Hash
	chunkInfo.IsUploaded = true

	if err := s.uploadRepo.UpdateChunkStatus(fileID, chunkIndex, true); err != nil {
		return nil, fmt.Errorf("更新分片状态失败: %v", err)
	}

	// 8. 更新上传进度
	count, err := s.uploadRepo.GetUploadedChunksCount(fileID)
	if err != nil {
		return nil, fmt.Errorf("获取上传进度失败: %v", err)
	}

	if err := s.uploadRepo.UpdateUploadProgress(fileID, int(count)); err != nil {
		return nil, fmt.Errorf("更新上传进度失败: %v", err)
	}

	logger.Info("分片上传成功", "fileID", fileID, "chunkIndex", chunkIndex, "progress", fmt.Sprintf("%d/%d", count, uploadFile.ChunkTotal))

	return &response.ChunkUploadResponse{
		FileID:        fileID,
		ChunkIndex:    chunkIndex,
		ChunkUploaded: int(count),
		ChunkTotal:    uploadFile.ChunkTotal,
		IsCompleted:   int(count) == uploadFile.ChunkTotal,
	}, nil
}

// MergeChunks 合并分片
func (s *UploadService) MergeChunks(fileID string) (*response.ChunkMergeResponse, error) {
	// 1. 获取文件信息
	uploadFile, err := s.uploadRepo.GetUploadFileByID(fileID)
	if err != nil {
		return nil, fmt.Errorf("文件记录不存在: %v", err)
	}

	// 2. 检查所有分片是否已上传
	count, err := s.uploadRepo.GetUploadedChunksCount(fileID)
	if err != nil {
		return nil, fmt.Errorf("获取上传进度失败: %v", err)
	}

	if int(count) != uploadFile.ChunkTotal {
		return nil, fmt.Errorf("分片尚未完全上传，进度: %d/%d", count, uploadFile.ChunkTotal)
	}

	// 3. 获取所有分片信息
	chunks, err := s.uploadRepo.GetFileChunks(fileID)
	if err != nil {
		return nil, fmt.Errorf("获取分片信息失败: %v", err)
	}

	// 4. 准备分片路径列表
	var chunkPaths []string
	for _, chunk := range chunks {
		chunkPaths = append(chunkPaths, chunk.ChunkPath)
	}

	// 5. 合并分片
	targetPath := filepath.Join(s.config.UploadDir, uploadFile.StoredName)
	if err := s.storage.MergeChunks(chunkPaths, targetPath); err != nil {
		return nil, fmt.Errorf("合并分片失败: %v", err)
	}

	// 6. 更新文件记录
	now := time.Now()
	uploadFile.FilePath = targetPath
	uploadFile.UploadStatus = constants.UploadStatusCompleted
	uploadFile.UploadedAt = &now

	if err := s.uploadRepo.UpdateUploadFile(uploadFile); err != nil {
		return nil, fmt.Errorf("更新文件记录失败: %v", err)
	}

	// 7. 清理分片记录
	if err := s.uploadRepo.DeleteFileChunks(fileID); err != nil {
		logger.Error("清理分片记录失败", "fileID", fileID, "error", err)
	}

	logger.Info("分片合并成功", "fileID", fileID, "filename", uploadFile.Filename)

	return &response.ChunkMergeResponse{
		FileID:     uploadFile.ID.String(),
		Filename:   uploadFile.Filename,
		StoredName: uploadFile.StoredName,
		FileSize:   uploadFile.FileSize,
		MimeType:   uploadFile.MimeType,
		Extension:  uploadFile.Extension,
		MD5Hash:    uploadFile.MD5Hash,
		FilePath:   fmt.Sprintf("/files/preview/%s", uploadFile.ID.String()),
		UploadedAt: *uploadFile.UploadedAt,
	}, nil
}

// GetUploadProgress 获取上传进度
func (s *UploadService) GetUploadProgress(fileID string) (*response.UploadProgressResponse, error) {
	uploadFile, err := s.uploadRepo.GetUploadFileByID(fileID)
	if err != nil {
		return nil, fmt.Errorf("文件记录不存在: %v", err)
	}

	progress := float64(uploadFile.ChunkUploaded) / float64(uploadFile.ChunkTotal) * 100

	return &response.UploadProgressResponse{
		FileID:        uploadFile.ID.String(),
		Filename:      uploadFile.Filename,
		FileSize:      uploadFile.FileSize,
		ChunkTotal:    uploadFile.ChunkTotal,
		ChunkUploaded: uploadFile.ChunkUploaded,
		Progress:      progress,
		Status:        uploadFile.UploadStatus,
	}, nil
}

// validateFile 验证文件
func (s *UploadService) validateFile(file *multipart.FileHeader) error {
	// 验证文件名
	if err := upload.ValidateFilename(file.Filename); err != nil {
		return err
	}

	// 验证文件大小
	if !s.config.ValidateFileSize(file.Size) {
		return fmt.Errorf("文件大小超出限制，最大允许 %d 字节", s.config.MaxFileSize)
	}

	// 验证文件类型
	mimeType := upload.GetMimeTypeFromExtension(file.Filename)
	if !s.config.ValidateMimeType(mimeType) {
		return fmt.Errorf("不支持的文件类型: %s", mimeType)
	}

	// 验证文件扩展名
	if !upload.ValidateFileExtension(file.Filename) {
		return fmt.Errorf("不支持的文件扩展名: %s", upload.GetFileExtension(file.Filename))
	}

	return nil
}
