package pagination

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Pagination struct {
	Page       int   `json:"page" form:"page"`
	PageSize   int   `json:"page_size" form:"page_size"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

type PaginationResult struct {
	Data       interface{} `json:"data"`
	Pagination *Pagination `json:"pagination"`
}

func GetPagination(c *gin.Context) *Pagination {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	return &Pagination{
		Page:     page,
		PageSize: pageSize,
	}
}

func (p *Pagination) GetOffset() int {
	return (p.Page - 1) * p.PageSize
}

func (p *Pagination) GetLimit() int {
	return p.PageSize
}

func (p *Pagination) Paginate() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.GetOffset()).Limit(p.GetLimit())
	}
}

func (p *Pagination) SetTotal(total int64) {
	p.Total = total
	p.TotalPages = int(math.Ceil(float64(total) / float64(p.PageSize)))
}

func (p *Pagination) HasMore() bool {
	return p.Page < p.TotalPages
}

func (p *Pagination) HasPrev() bool {
	return p.Page > 1
}

func PaginateQuery(db *gorm.DB, pagination *Pagination, query func(*gorm.DB) *gorm.DB, result interface{}) error {
	var total int64
	if err := query(db).Count(&total).Error; err != nil {
		return err
	}
	pagination.SetTotal(total)

	if total == 0 {
		return nil
	}

	if err := query(db).Scopes(pagination.Paginate()).Find(result).Error; err != nil {
		return err
	}

	return nil
}

func NewPaginationResult(data interface{}, pagination *Pagination) *PaginationResult {
	return &PaginationResult{
		Data:       data,
		Pagination: pagination,
	}
}

type CursorPagination struct {
	Cursor   string `json:"cursor" form:"cursor"`
	PageSize int    `json:"page_size" form:"page_size"`
	HasMore  bool   `json:"has_more"`
}

func GetCursorPagination(c *gin.Context) *CursorPagination {
	cursor := c.Query("cursor")
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	return &CursorPagination{
		Cursor:   cursor,
		PageSize: pageSize,
	}
}

type CursorResult struct {
	Data       interface{}       `json:"data"`
	Pagination *CursorPagination `json:"pagination"`
}

func NewCursorResult(data interface{}, pagination *CursorPagination) *CursorResult {
	return &CursorResult{
		Data:       data,
		Pagination: pagination,
	}
}
