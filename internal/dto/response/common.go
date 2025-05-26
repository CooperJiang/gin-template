package response

import (
	"template/internal/dto"
	"time"
)

// BaseResponse 基础响应DTO
type BaseResponse struct {
	dto.BaseResponse
}

// StatusResponse 状态响应
type StatusResponse struct {
	Status     int    `json:"status"`
	StatusText string `json:"status_text"`
}

// CountResponse 计数响应
type CountResponse struct {
	Count int64 `json:"count"`
}

// MessageResponse 消息响应
type MessageResponse struct {
	Message string `json:"message"`
}

// SuccessResponse 成功响应
type SuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// IDResponse ID响应
type IDResponse struct {
	ID uint `json:"id"`
}

// ListResponse 列表响应基础结构
type ListResponse[T any] struct {
	Items []T   `json:"items"`
	Total int64 `json:"total"`
}

// TimeResponse 时间响应
type TimeResponse struct {
	Timestamp time.Time `json:"timestamp"`
}

// VersionResponse 版本响应
type VersionResponse struct {
	Version   string    `json:"version"`
	BuildTime time.Time `json:"build_time"`
	GoVersion string    `json:"go_version"`
}

// HealthResponse 健康检查响应
type HealthResponse struct {
	Status    string            `json:"status"`
	Timestamp time.Time         `json:"timestamp"`
	Services  map[string]string `json:"services"`
	Version   string            `json:"version"`
}

// StatisticsResponse 统计响应
type StatisticsResponse struct {
	TotalUsers    int64 `json:"total_users"`
	ActiveUsers   int64 `json:"active_users"`
	TotalRequests int64 `json:"total_requests"`
	Uptime        int64 `json:"uptime"` // 秒
}

// FileResponse 文件响应
type FileResponse struct {
	FileName string `json:"file_name"`
	FileSize int64  `json:"file_size"`
	FileURL  string `json:"file_url"`
	MimeType string `json:"mime_type"`
}

// BatchResponse 批量操作响应
type BatchResponse struct {
	SuccessCount int    `json:"success_count"`
	FailedCount  int    `json:"failed_count"`
	TotalCount   int    `json:"total_count"`
	Message      string `json:"message"`
}
