package request

import "template/pkg/common"

// IDRequest ID请求参数
type IDRequest struct {
	ID uint `uri:"id" binding:"required,min=1" json:"id"`
}

// GetValidationMessages 获取验证错误信息
func (r *IDRequest) GetValidationMessages() map[string]string {
	return map[string]string{
		"ID.required": "ID不能为空",
		"ID.min":      "ID必须大于0",
	}
}

// PaginationRequest 分页请求
type PaginationRequest struct {
	common.PaginationRequest
}

// SearchRequest 搜索请求
type SearchRequest struct {
	PaginationRequest
	Keyword string `form:"keyword" binding:"omitempty,max=100" json:"keyword"`
	Status  *int   `form:"status" binding:"omitempty,oneof=0 1" json:"status"`
}

// GetValidationMessages 获取验证错误信息
func (r *SearchRequest) GetValidationMessages() map[string]string {
	return map[string]string{
		"Keyword.max":  "关键词长度不能超过100个字符",
		"Status.oneof": "状态值必须是0或1",
	}
}

// SortRequest 排序请求
type SortRequest struct {
	SortBy    string `form:"sort_by" binding:"omitempty,oneof=id created_at updated_at name" json:"sort_by"`
	SortOrder string `form:"sort_order" binding:"omitempty,oneof=asc desc" json:"sort_order"`
}

// GetValidationMessages 获取验证错误信息
func (r *SortRequest) GetValidationMessages() map[string]string {
	return map[string]string{
		"SortBy.oneof":    "排序字段必须是id、created_at、updated_at或name",
		"SortOrder.oneof": "排序方向必须是asc或desc",
	}
}

// GetSortBy 获取排序字段，如果为空则返回默认值
func (r *SortRequest) GetSortBy() string {
	if r.SortBy == "" {
		return "created_at"
	}
	return r.SortBy
}

// GetSortOrder 获取排序方向，如果为空则返回默认值
func (r *SortRequest) GetSortOrder() string {
	if r.SortOrder == "" {
		return "desc"
	}
	return r.SortOrder
}

// BatchRequest 批量操作请求
type BatchRequest struct {
	IDs []uint `json:"ids" binding:"required,min=1,dive,min=1"`
}

// GetValidationMessages 获取验证错误信息
func (r *BatchRequest) GetValidationMessages() map[string]string {
	return map[string]string{
		"IDs.required": "ID列表不能为空",
		"IDs.min":      "至少需要选择一个项目",
		"IDs.dive":     "ID必须大于0",
	}
}
