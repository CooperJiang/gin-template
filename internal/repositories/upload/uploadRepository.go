package upload

import (
	"template/internal/models"
	"template/pkg/constants"

	"gorm.io/gorm"
)

type UploadRepository struct {
	db *gorm.DB
}

func NewUploadRepository(db *gorm.DB) *UploadRepository {
	return &UploadRepository{db: db}
}

// CreateUploadFile 创建文件上传记录
func (r *UploadRepository) CreateUploadFile(file *models.UploadFile) error {
	return r.db.Create(file).Error
}

// GetUploadFileByID 根据ID获取文件信息
func (r *UploadRepository) GetUploadFileByID(fileID string) (*models.UploadFile, error) {
	var file models.UploadFile
	err := r.db.Where("id = ?", fileID).First(&file).Error
	return &file, err
}

// GetUploadFileByMD5 根据MD5获取文件信息
func (r *UploadRepository) GetUploadFileByMD5(md5Hash string) (*models.UploadFile, error) {
	var file models.UploadFile
	err := r.db.Where("md5_hash = ? AND upload_status = ?", md5Hash, constants.UploadStatusCompleted).First(&file).Error
	return &file, err
}

// UpdateUploadFile 更新文件上传记录
func (r *UploadRepository) UpdateUploadFile(file *models.UploadFile) error {
	return r.db.Save(file).Error
}

// UpdateUploadProgress 更新上传进度
func (r *UploadRepository) UpdateUploadProgress(fileID string, chunkUploaded int) error {
	return r.db.Model(&models.UploadFile{}).
		Where("id = ?", fileID).
		Update("chunk_uploaded", chunkUploaded).Error
}

// UpdateUploadStatus 更新上传状态
func (r *UploadRepository) UpdateUploadStatus(fileID string, status int) error {
	return r.db.Model(&models.UploadFile{}).
		Where("id = ?", fileID).
		Update("upload_status", status).Error
}

// DeleteUploadFile 删除文件记录
func (r *UploadRepository) DeleteUploadFile(fileID string) error {
	return r.db.Where("id = ?", fileID).Delete(&models.UploadFile{}).Error
}

// GetUserFiles 获取用户文件列表
func (r *UploadRepository) GetUserFiles(userID string, page, pageSize int) ([]models.UploadFile, int64, error) {
	var files []models.UploadFile
	var total int64

	// 计算总数
	if err := r.db.Model(&models.UploadFile{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err := r.db.Where("user_id = ?", userID).
		Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&files).Error

	return files, total, err
}

// CreateChunkInfo 创建分片信息
func (r *UploadRepository) CreateChunkInfo(chunk *models.ChunkInfo) error {
	return r.db.Create(chunk).Error
}

// GetChunkInfo 获取分片信息
func (r *UploadRepository) GetChunkInfo(fileID string, chunkIndex int) (*models.ChunkInfo, error) {
	var chunk models.ChunkInfo
	err := r.db.Where("file_id = ? AND chunk_index = ?", fileID, chunkIndex).First(&chunk).Error
	return &chunk, err
}

// GetFileChunks 获取文件的所有分片信息
func (r *UploadRepository) GetFileChunks(fileID string) ([]models.ChunkInfo, error) {
	var chunks []models.ChunkInfo
	err := r.db.Where("file_id = ?", fileID).Order("chunk_index").Find(&chunks).Error
	return chunks, err
}

// UpdateChunkStatus 更新分片状态
func (r *UploadRepository) UpdateChunkStatus(fileID string, chunkIndex int, isUploaded bool) error {
	return r.db.Model(&models.ChunkInfo{}).
		Where("file_id = ? AND chunk_index = ?", fileID, chunkIndex).
		Update("is_uploaded", isUploaded).Error
}

// GetUploadedChunksCount 获取已上传分片数量
func (r *UploadRepository) GetUploadedChunksCount(fileID string) (int64, error) {
	var count int64
	err := r.db.Model(&models.ChunkInfo{}).
		Where("file_id = ? AND is_uploaded = ?", fileID, true).
		Count(&count).Error
	return count, err
}

// DeleteFileChunks 删除文件的所有分片记录
func (r *UploadRepository) DeleteFileChunks(fileID string) error {
	return r.db.Where("file_id = ?", fileID).Delete(&models.ChunkInfo{}).Error
}

// CheckChunkExists 检查分片是否已存在
func (r *UploadRepository) CheckChunkExists(fileID string, chunkIndex int) bool {
	var count int64
	r.db.Model(&models.ChunkInfo{}).
		Where("file_id = ? AND chunk_index = ? AND is_uploaded = ?", fileID, chunkIndex, true).
		Count(&count)
	return count > 0
}
