package common

import "template/pkg/constants"

// PaginationRequest 分页请求
type PaginationRequest struct {
	Page     int `form:"page" binding:"omitempty,min=1" json:"page"`           // 页码，从1开始
	PageSize int `form:"page_size" binding:"omitempty,min=1" json:"page_size"` // 每页数量
}

// GetPage 获取页码，如果为0则返回默认值
func (p *PaginationRequest) GetPage() int {
	if p.Page <= 0 {
		return constants.DefaultPage
	}
	return p.Page
}

// GetPageSize 获取每页数量，如果为0则返回默认值，如果超过最大值则返回最大值
func (p *PaginationRequest) GetPageSize() int {
	if p.PageSize <= 0 {
		return constants.DefaultPageSize
	}
	if p.PageSize > constants.MaxPageSize {
		return constants.MaxPageSize
	}
	return p.PageSize
}

// GetOffset 获取偏移量
func (p *PaginationRequest) GetOffset() int {
	return (p.GetPage() - 1) * p.GetPageSize()
}

// PaginationInfo 分页信息
type PaginationInfo struct {
	Page       int   `json:"page"`        // 当前页码
	PageSize   int   `json:"page_size"`   // 每页数量
	Total      int64 `json:"total"`       // 总记录数
	TotalPages int64 `json:"total_pages"` // 总页数
}

// PaginationResponse 分页响应
type PaginationResponse struct {
	Items      interface{}    `json:"items"`      // 数据列表
	Pagination PaginationInfo `json:"pagination"` // 分页信息
}

// NewPaginationResponse 创建分页响应
func NewPaginationResponse(items interface{}, req *PaginationRequest, total int64) *PaginationResponse {
	page := req.GetPage()
	pageSize := req.GetPageSize()
	totalPages := (total + int64(pageSize) - 1) / int64(pageSize)

	return &PaginationResponse{
		Items: items,
		Pagination: PaginationInfo{
			Page:       page,
			PageSize:   pageSize,
			Total:      total,
			TotalPages: totalPages,
		},
	}
}

// HasNextPage 是否有下一页
func (p *PaginationInfo) HasNextPage() bool {
	return int64(p.Page) < p.TotalPages
}

// HasPrevPage 是否有上一页
func (p *PaginationInfo) HasPrevPage() bool {
	return p.Page > 1
}

// GetNextPage 获取下一页页码
func (p *PaginationInfo) GetNextPage() int {
	if p.HasNextPage() {
		return p.Page + 1
	}
	return p.Page
}

// GetPrevPage 获取上一页页码
func (p *PaginationInfo) GetPrevPage() int {
	if p.HasPrevPage() {
		return p.Page - 1
	}
	return p.Page
}
