package upload

import (
	"crypto/md5"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"path/filepath"
	"strings"
	"template/pkg/constants"

	"github.com/google/uuid"
)

// GenerateStoredFilename 生成存储文件名
func GenerateStoredFilename(originalFilename string) string {
	ext := filepath.Ext(originalFilename)
	return fmt.Sprintf("%s%s", uuid.New().String(), ext)
}

// GetMimeTypeFromExtension 根据扩展名获取MIME类型
func GetMimeTypeFromExtension(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	if mimeType, exists := constants.ExtensionMimeMap[ext]; exists {
		return mimeType
	}
	// 使用系统默认MIME类型检测
	return mime.TypeByExtension(ext)
}

// ValidateFileExtension 验证文件扩展名
func ValidateFileExtension(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	_, exists := constants.ExtensionMimeMap[ext]
	return exists
}

// CalculateFileMD5 计算文件MD5哈希
func CalculateFileMD5(file multipart.File) (string, error) {
	hasher := md5.New()

	// 重置文件指针到开头
	file.Seek(0, 0)

	// 计算MD5
	_, err := io.Copy(hasher, file)
	if err != nil {
		return "", err
	}

	// 重置文件指针到开头
	file.Seek(0, 0)

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}

// CalculateChunkMD5 计算分片MD5哈希
func CalculateChunkMD5(chunk io.Reader) (string, error) {
	hasher := md5.New()
	_, err := io.Copy(hasher, chunk)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}

// ValidateFilename 验证文件名
func ValidateFilename(filename string) error {
	if len(filename) == 0 {
		return fmt.Errorf("文件名不能为空")
	}

	if len(filename) > constants.MaxFilenameLength {
		return fmt.Errorf("文件名过长，最大长度为 %d 字符", constants.MaxFilenameLength)
	}

	// 检查是否包含非法字符
	invalidChars := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|"}
	for _, char := range invalidChars {
		if strings.Contains(filename, char) {
			return fmt.Errorf("文件名包含非法字符: %s", char)
		}
	}

	return nil
}

// GetFileExtension 获取文件扩展名
func GetFileExtension(filename string) string {
	return strings.ToLower(filepath.Ext(filename))
}

// CalculateChunkTotal 计算总分片数
func CalculateChunkTotal(fileSize, chunkSize int64) int {
	chunks := fileSize / chunkSize
	if fileSize%chunkSize != 0 {
		chunks++
	}
	return int(chunks)
}

// GenerateChunkPath 生成分片文件路径
func GenerateChunkPath(tempDir, fileID string, chunkIndex int) string {
	return filepath.Join(tempDir, fileID, fmt.Sprintf("chunk_%d", chunkIndex))
}
