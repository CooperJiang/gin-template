package upload

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// Storage 文件存储接口
type Storage interface {
	// SaveFile 保存文件
	SaveFile(filename string, file multipart.File) (string, error)
	// SaveChunk 保存分片
	SaveChunk(chunkPath string, chunk io.Reader) error
	// MergeChunks 合并分片
	MergeChunks(chunkPaths []string, targetPath string) error
	// DeleteFile 删除文件
	DeleteFile(filePath string) error
	// FileExists 检查文件是否存在
	FileExists(filePath string) bool
	// GetFileSize 获取文件大小
	GetFileSize(filePath string) (int64, error)
}

// LocalStorage 本地存储实现
type LocalStorage struct {
	baseDir string
}

// NewLocalStorage 创建本地存储实例
func NewLocalStorage(baseDir string) *LocalStorage {
	return &LocalStorage{
		baseDir: baseDir,
	}
}

// SaveFile 保存文件到本地
func (ls *LocalStorage) SaveFile(filename string, file multipart.File) (string, error) {
	// 确保目录存在
	if err := os.MkdirAll(ls.baseDir, 0755); err != nil {
		return "", err
	}

	// 创建目标文件
	filePath := filepath.Join(ls.baseDir, filename)
	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// 复制文件内容
	_, err = io.Copy(dst, file)
	if err != nil {
		os.Remove(filePath) // 删除失败的文件
		return "", err
	}

	return filePath, nil
}

// SaveChunk 保存分片到本地
func (ls *LocalStorage) SaveChunk(chunkPath string, chunk io.Reader) error {
	// 确保目录存在
	dir := filepath.Dir(chunkPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// 创建分片文件
	dst, err := os.Create(chunkPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// 复制分片内容
	_, err = io.Copy(dst, chunk)
	return err
}

// MergeChunks 合并分片
func (ls *LocalStorage) MergeChunks(chunkPaths []string, targetPath string) error {
	// 确保目标目录存在
	dir := filepath.Dir(targetPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// 创建目标文件
	target, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer target.Close()

	// 依次读取并合并分片
	for _, chunkPath := range chunkPaths {
		chunk, err := os.Open(chunkPath)
		if err != nil {
			return err
		}

		_, err = io.Copy(target, chunk)
		chunk.Close()

		if err != nil {
			return err
		}

		// 删除已合并的分片
		os.Remove(chunkPath)
	}

	return nil
}

// DeleteFile 删除本地文件
func (ls *LocalStorage) DeleteFile(filePath string) error {
	return os.Remove(filePath)
}

// FileExists 检查本地文件是否存在
func (ls *LocalStorage) FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

// GetFileSize 获取本地文件大小
func (ls *LocalStorage) GetFileSize(filePath string) (int64, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}
