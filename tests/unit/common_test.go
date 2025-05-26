package unit

import (
	"template/pkg/common"
	"template/pkg/constants"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPaginationRequest 测试分页请求
func TestPaginationRequest(t *testing.T) {
	tests := []struct {
		name     string
		req      common.PaginationRequest
		wantPage int
		wantSize int
	}{
		{
			name:     "默认值",
			req:      common.PaginationRequest{},
			wantPage: constants.DefaultPage,
			wantSize: constants.DefaultPageSize,
		},
		{
			name:     "正常值",
			req:      common.PaginationRequest{Page: 2, PageSize: 10},
			wantPage: 2,
			wantSize: 10,
		},
		{
			name:     "超过最大值",
			req:      common.PaginationRequest{Page: 1, PageSize: 200},
			wantPage: 1,
			wantSize: constants.MaxPageSize,
		},
		{
			name:     "负数值",
			req:      common.PaginationRequest{Page: -1, PageSize: -10},
			wantPage: constants.DefaultPage,
			wantSize: constants.DefaultPageSize,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantPage, tt.req.GetPage())
			assert.Equal(t, tt.wantSize, tt.req.GetPageSize())
		})
	}
}

// TestPaginationOffset 测试分页偏移量计算
func TestPaginationOffset(t *testing.T) {
	tests := []struct {
		name       string
		page       int
		pageSize   int
		wantOffset int
	}{
		{
			name:       "第一页",
			page:       1,
			pageSize:   10,
			wantOffset: 0,
		},
		{
			name:       "第二页",
			page:       2,
			pageSize:   10,
			wantOffset: 10,
		},
		{
			name:       "第三页",
			page:       3,
			pageSize:   20,
			wantOffset: 40,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := common.PaginationRequest{
				Page:     tt.page,
				PageSize: tt.pageSize,
			}
			assert.Equal(t, tt.wantOffset, req.GetOffset())
		})
	}
}

// TestNewPaginationResponse 测试分页响应创建
func TestNewPaginationResponse(t *testing.T) {
	items := []string{"item1", "item2", "item3"}
	req := &common.PaginationRequest{Page: 2, PageSize: 10}
	total := int64(25)

	resp := common.NewPaginationResponse(items, req, total)

	assert.Equal(t, items, resp.Items)
	assert.Equal(t, 2, resp.Pagination.Page)
	assert.Equal(t, 10, resp.Pagination.PageSize)
	assert.Equal(t, int64(25), resp.Pagination.Total)
	assert.Equal(t, int64(3), resp.Pagination.TotalPages)
}

// TestPaginationInfo 测试分页信息
func TestPaginationInfo(t *testing.T) {
	info := common.PaginationInfo{
		Page:       2,
		PageSize:   10,
		Total:      25,
		TotalPages: 3,
	}

	// 测试是否有下一页
	assert.True(t, info.HasNextPage())

	// 测试是否有上一页
	assert.True(t, info.HasPrevPage())

	// 测试获取下一页
	assert.Equal(t, 3, info.GetNextPage())

	// 测试获取上一页
	assert.Equal(t, 1, info.GetPrevPage())

	// 测试最后一页
	lastPageInfo := common.PaginationInfo{
		Page:       3,
		PageSize:   10,
		Total:      25,
		TotalPages: 3,
	}
	assert.False(t, lastPageInfo.HasNextPage())
	assert.Equal(t, 3, lastPageInfo.GetNextPage()) // 应该返回当前页

	// 测试第一页
	firstPageInfo := common.PaginationInfo{
		Page:       1,
		PageSize:   10,
		Total:      25,
		TotalPages: 3,
	}
	assert.False(t, firstPageInfo.HasPrevPage())
	assert.Equal(t, 1, firstPageInfo.GetPrevPage()) // 应该返回当前页
}
