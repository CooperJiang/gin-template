#!/bin/bash

# 创建新模块脚手架脚本
# 使用方法: ./scripts/create_module.sh module_name

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 检查参数
if [ $# -eq 0 ]; then
    echo -e "${RED}错误: 请提供模块名称${NC}"
    echo "使用方法: $0 <module_name>"
    echo "示例: $0 product"
    exit 1
fi

MODULE_NAME=$1
MODULE_NAME_LOWER=$(echo "$MODULE_NAME" | tr '[:upper:]' '[:lower:]')
MODULE_NAME_UPPER=$(echo "$MODULE_NAME" | tr '[:lower:]' '[:upper:]')
MODULE_NAME_TITLE=$(echo "$MODULE_NAME" | sed 's/\b\w/\U&/g')

echo -e "${BLUE}创建模块: ${MODULE_NAME_TITLE}${NC}"

# 创建目录结构
echo -e "${YELLOW}创建目录结构...${NC}"
mkdir -p "internal/controllers/${MODULE_NAME_LOWER}"
mkdir -p "internal/services/${MODULE_NAME_LOWER}"
mkdir -p "internal/repositories/${MODULE_NAME_LOWER}"
mkdir -p "internal/dto/request/${MODULE_NAME_LOWER}"
mkdir -p "internal/dto/response/${MODULE_NAME_LOWER}"

# 创建模型文件
echo -e "${YELLOW}创建模型文件...${NC}"
cat > "internal/models/${MODULE_NAME_LOWER}.go" << EOF
package models

import (
	"time"
	"gorm.io/gorm"
)

// ${MODULE_NAME_TITLE} ${MODULE_NAME_LOWER}模型
type ${MODULE_NAME_TITLE} struct {
	BaseModel
	Name        string \`gorm:"size:100;not null" json:"name"\`
	Description string \`gorm:"size:500" json:"description"\`
	Status      int    \`gorm:"default:1" json:"status"\`
}

// TableName 指定表名
func (${MODULE_NAME_TITLE}) TableName() string {
	return "${MODULE_NAME_LOWER}s"
}

// BeforeCreate 创建前钩子
func (m *${MODULE_NAME_TITLE}) BeforeCreate(tx *gorm.DB) error {
	// 在这里添加创建前的逻辑
	return nil
}

// BeforeUpdate 更新前钩子
func (m *${MODULE_NAME_TITLE}) BeforeUpdate(tx *gorm.DB) error {
	// 在这里添加更新前的逻辑
	return nil
}
EOF

# 创建请求DTO
echo -e "${YELLOW}创建请求DTO...${NC}"
cat > "internal/dto/request/${MODULE_NAME_LOWER}/create.go" << EOF
package ${MODULE_NAME_LOWER}

// Create${MODULE_NAME_TITLE}Request 创建${MODULE_NAME_LOWER}请求
type Create${MODULE_NAME_TITLE}Request struct {
	Name        string \`json:"name" binding:"required,min=1,max=100"\`
	Description string \`json:"description" binding:"max=500"\`
}

// GetValidationMessages 获取验证错误信息
func (r *Create${MODULE_NAME_TITLE}Request) GetValidationMessages() map[string]string {
	return map[string]string{
		"Name.required": "名称不能为空",
		"Name.min":      "名称长度不能小于1个字符",
		"Name.max":      "名称长度不能超过100个字符",
		"Description.max": "描述长度不能超过500个字符",
	}
}
EOF

cat > "internal/dto/request/${MODULE_NAME_LOWER}/update.go" << EOF
package ${MODULE_NAME_LOWER}

// Update${MODULE_NAME_TITLE}Request 更新${MODULE_NAME_LOWER}请求
type Update${MODULE_NAME_TITLE}Request struct {
	Name        *string \`json:"name" binding:"omitempty,min=1,max=100"\`
	Description *string \`json:"description" binding:"omitempty,max=500"\`
	Status      *int    \`json:"status" binding:"omitempty,oneof=0 1"\`
}

// GetValidationMessages 获取验证错误信息
func (r *Update${MODULE_NAME_TITLE}Request) GetValidationMessages() map[string]string {
	return map[string]string{
		"Name.min":        "名称长度不能小于1个字符",
		"Name.max":        "名称长度不能超过100个字符",
		"Description.max": "描述长度不能超过500个字符",
		"Status.oneof":    "状态值必须是0或1",
	}
}
EOF

cat > "internal/dto/request/${MODULE_NAME_LOWER}/query.go" << EOF
package ${MODULE_NAME_LOWER}

import "template/pkg/common"

// Query${MODULE_NAME_TITLE}Request 查询${MODULE_NAME_LOWER}请求
type Query${MODULE_NAME_TITLE}Request struct {
	common.PaginationRequest
	Name   string \`form:"name" binding:"omitempty,max=100"\`
	Status *int   \`form:"status" binding:"omitempty,oneof=0 1"\`
}

// GetValidationMessages 获取验证错误信息
func (r *Query${MODULE_NAME_TITLE}Request) GetValidationMessages() map[string]string {
	return map[string]string{
		"Name.max":     "名称长度不能超过100个字符",
		"Status.oneof": "状态值必须是0或1",
	}
}
EOF

# 创建响应DTO
echo -e "${YELLOW}创建响应DTO...${NC}"
cat > "internal/dto/response/${MODULE_NAME_LOWER}/${MODULE_NAME_LOWER}_info.go" << EOF
package ${MODULE_NAME_LOWER}

import (
	"time"
	"template/internal/models"
)

// ${MODULE_NAME_TITLE}Response ${MODULE_NAME_LOWER}响应
type ${MODULE_NAME_TITLE}Response struct {
	ID          uint      \`json:"id"\`
	Name        string    \`json:"name"\`
	Description string    \`json:"description"\`
	Status      int       \`json:"status"\`
	CreatedAt   time.Time \`json:"created_at"\`
	UpdatedAt   time.Time \`json:"updated_at"\`
}

// From${MODULE_NAME_TITLE} 从模型转换为响应DTO
func From${MODULE_NAME_TITLE}(${MODULE_NAME_LOWER} *models.${MODULE_NAME_TITLE}) *${MODULE_NAME_TITLE}Response {
	return &${MODULE_NAME_TITLE}Response{
		ID:          ${MODULE_NAME_LOWER}.ID,
		Name:        ${MODULE_NAME_LOWER}.Name,
		Description: ${MODULE_NAME_LOWER}.Description,
		Status:      ${MODULE_NAME_LOWER}.Status,
		CreatedAt:   ${MODULE_NAME_LOWER}.CreatedAt,
		UpdatedAt:   ${MODULE_NAME_LOWER}.UpdatedAt,
	}
}

// From${MODULE_NAME_TITLE}List 从模型列表转换为响应DTO列表
func From${MODULE_NAME_TITLE}List(${MODULE_NAME_LOWER}s []*models.${MODULE_NAME_TITLE}) []*${MODULE_NAME_TITLE}Response {
	result := make([]*${MODULE_NAME_TITLE}Response, len(${MODULE_NAME_LOWER}s))
	for i, ${MODULE_NAME_LOWER} := range ${MODULE_NAME_LOWER}s {
		result[i] = From${MODULE_NAME_TITLE}(${MODULE_NAME_LOWER})
	}
	return result
}
EOF

# 创建Repository
echo -e "${YELLOW}创建Repository...${NC}"
cat > "internal/repositories/${MODULE_NAME_LOWER}/${MODULE_NAME_LOWER}_repository.go" << EOF
package ${MODULE_NAME_LOWER}

import (
	"context"
	"template/internal/models"
	"template/pkg/common"
	"template/pkg/errors"
	
	"gorm.io/gorm"
)

// ${MODULE_NAME_TITLE}Repository ${MODULE_NAME_LOWER}数据访问接口
type ${MODULE_NAME_TITLE}Repository interface {
	Create(ctx context.Context, ${MODULE_NAME_LOWER} *models.${MODULE_NAME_TITLE}) error
	GetByID(ctx context.Context, id uint) (*models.${MODULE_NAME_TITLE}, error)
	Update(ctx context.Context, id uint, updates map[string]interface{}) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, req *common.PaginationRequest, filters map[string]interface{}) ([]*models.${MODULE_NAME_TITLE}, int64, error)
	ExistsByName(ctx context.Context, name string, excludeID ...uint) (bool, error)
}

// ${MODULE_NAME_LOWER}Repository ${MODULE_NAME_LOWER}数据访问实现
type ${MODULE_NAME_LOWER}Repository struct {
	db *gorm.DB
}

// New${MODULE_NAME_TITLE}Repository 创建${MODULE_NAME_LOWER}数据访问实例
func New${MODULE_NAME_TITLE}Repository(db *gorm.DB) ${MODULE_NAME_TITLE}Repository {
	return &${MODULE_NAME_LOWER}Repository{
		db: db,
	}
}

// Create 创建${MODULE_NAME_LOWER}
func (r *${MODULE_NAME_LOWER}Repository) Create(ctx context.Context, ${MODULE_NAME_LOWER} *models.${MODULE_NAME_TITLE}) error {
	if err := r.db.WithContext(ctx).Create(${MODULE_NAME_LOWER}).Error; err != nil {
		return errors.Wrap(err, errors.CodeDBQueryFailed)
	}
	return nil
}

// GetByID 根据ID获取${MODULE_NAME_LOWER}
func (r *${MODULE_NAME_LOWER}Repository) GetByID(ctx context.Context, id uint) (*models.${MODULE_NAME_TITLE}, error) {
	var ${MODULE_NAME_LOWER} models.${MODULE_NAME_TITLE}
	err := r.db.WithContext(ctx).First(&${MODULE_NAME_LOWER}, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(errors.CodeNotFound, "${MODULE_NAME_TITLE}不存在")
		}
		return nil, errors.Wrap(err, errors.CodeDBQueryFailed)
	}
	return &${MODULE_NAME_LOWER}, nil
}

// Update 更新${MODULE_NAME_LOWER}
func (r *${MODULE_NAME_LOWER}Repository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	result := r.db.WithContext(ctx).Model(&models.${MODULE_NAME_TITLE}{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return errors.Wrap(result.Error, errors.CodeDBQueryFailed)
	}
	if result.RowsAffected == 0 {
		return errors.New(errors.CodeNotFound, "${MODULE_NAME_TITLE}不存在")
	}
	return nil
}

// Delete 删除${MODULE_NAME_LOWER}
func (r *${MODULE_NAME_LOWER}Repository) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&models.${MODULE_NAME_TITLE}{}, id)
	if result.Error != nil {
		return errors.Wrap(result.Error, errors.CodeDBQueryFailed)
	}
	if result.RowsAffected == 0 {
		return errors.New(errors.CodeNotFound, "${MODULE_NAME_TITLE}不存在")
	}
	return nil
}

// List 获取${MODULE_NAME_LOWER}列表
func (r *${MODULE_NAME_LOWER}Repository) List(ctx context.Context, req *common.PaginationRequest, filters map[string]interface{}) ([]*models.${MODULE_NAME_TITLE}, int64, error) {
	var ${MODULE_NAME_LOWER}s []*models.${MODULE_NAME_TITLE}
	var total int64
	
	query := r.db.WithContext(ctx).Model(&models.${MODULE_NAME_TITLE}{})
	
	// 应用过滤条件
	for key, value := range filters {
		switch key {
		case "name":
			if name, ok := value.(string); ok && name != "" {
				query = query.Where("name LIKE ?", "%"+name+"%")
			}
		case "status":
			if status, ok := value.(int); ok {
				query = query.Where("status = ?", status)
			}
		}
	}
	
	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, errors.Wrap(err, errors.CodeDBQueryFailed)
	}
	
	// 分页查询
	offset := (req.Page - 1) * req.PageSize
	if err := query.Offset(offset).Limit(req.PageSize).Order("created_at DESC").Find(&${MODULE_NAME_LOWER}s).Error; err != nil {
		return nil, 0, errors.Wrap(err, errors.CodeDBQueryFailed)
	}
	
	return ${MODULE_NAME_LOWER}s, total, nil
}

// ExistsByName 检查名称是否存在
func (r *${MODULE_NAME_LOWER}Repository) ExistsByName(ctx context.Context, name string, excludeID ...uint) (bool, error) {
	var count int64
	query := r.db.WithContext(ctx).Model(&models.${MODULE_NAME_TITLE}{}).Where("name = ?", name)
	
	if len(excludeID) > 0 {
		query = query.Where("id != ?", excludeID[0])
	}
	
	err := query.Count(&count).Error
	if err != nil {
		return false, errors.Wrap(err, errors.CodeDBQueryFailed)
	}
	
	return count > 0, nil
}
EOF

# 创建Service
echo -e "${YELLOW}创建Service...${NC}"
cat > "internal/services/${MODULE_NAME_LOWER}/${MODULE_NAME_LOWER}_service.go" << EOF
package ${MODULE_NAME_LOWER}

import (
	"context"
	"template/internal/models"
	"template/internal/repositories/${MODULE_NAME_LOWER}"
	requestDto "template/internal/dto/request/${MODULE_NAME_LOWER}"
	responseDto "template/internal/dto/response/${MODULE_NAME_LOWER}"
	"template/pkg/common"
	"template/pkg/errors"
)

// ${MODULE_NAME_TITLE}Service ${MODULE_NAME_LOWER}服务接口
type ${MODULE_NAME_TITLE}Service interface {
	Create(ctx context.Context, req *requestDto.Create${MODULE_NAME_TITLE}Request) (*responseDto.${MODULE_NAME_TITLE}Response, error)
	GetByID(ctx context.Context, id uint) (*responseDto.${MODULE_NAME_TITLE}Response, error)
	Update(ctx context.Context, id uint, req *requestDto.Update${MODULE_NAME_TITLE}Request) (*responseDto.${MODULE_NAME_TITLE}Response, error)
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, req *requestDto.Query${MODULE_NAME_TITLE}Request) (*common.PaginationResponse, error)
}

// ${MODULE_NAME_LOWER}Service ${MODULE_NAME_LOWER}服务实现
type ${MODULE_NAME_LOWER}Service struct {
	${MODULE_NAME_LOWER}Repo ${MODULE_NAME_LOWER}.${MODULE_NAME_TITLE}Repository
}

// New${MODULE_NAME_TITLE}Service 创建${MODULE_NAME_LOWER}服务实例
func New${MODULE_NAME_TITLE}Service(${MODULE_NAME_LOWER}Repo ${MODULE_NAME_LOWER}.${MODULE_NAME_TITLE}Repository) ${MODULE_NAME_TITLE}Service {
	return &${MODULE_NAME_LOWER}Service{
		${MODULE_NAME_LOWER}Repo: ${MODULE_NAME_LOWER}Repo,
	}
}

// Create 创建${MODULE_NAME_LOWER}
func (s *${MODULE_NAME_LOWER}Service) Create(ctx context.Context, req *requestDto.Create${MODULE_NAME_TITLE}Request) (*responseDto.${MODULE_NAME_TITLE}Response, error) {
	// 检查名称是否已存在
	exists, err := s.${MODULE_NAME_LOWER}Repo.ExistsByName(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New(errors.CodeConflict, "名称已存在")
	}
	
	// 创建${MODULE_NAME_LOWER}
	${MODULE_NAME_LOWER} := &models.${MODULE_NAME_TITLE}{
		Name:        req.Name,
		Description: req.Description,
		Status:      1, // 默认启用
	}
	
	if err := s.${MODULE_NAME_LOWER}Repo.Create(ctx, ${MODULE_NAME_LOWER}); err != nil {
		return nil, err
	}
	
	return responseDto.From${MODULE_NAME_TITLE}(${MODULE_NAME_LOWER}), nil
}

// GetByID 根据ID获取${MODULE_NAME_LOWER}
func (s *${MODULE_NAME_LOWER}Service) GetByID(ctx context.Context, id uint) (*responseDto.${MODULE_NAME_TITLE}Response, error) {
	${MODULE_NAME_LOWER}, err := s.${MODULE_NAME_LOWER}Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	return responseDto.From${MODULE_NAME_TITLE}(${MODULE_NAME_LOWER}), nil
}

// Update 更新${MODULE_NAME_LOWER}
func (s *${MODULE_NAME_LOWER}Service) Update(ctx context.Context, id uint, req *requestDto.Update${MODULE_NAME_TITLE}Request) (*responseDto.${MODULE_NAME_TITLE}Response, error) {
	// 检查${MODULE_NAME_LOWER}是否存在
	existing, err := s.${MODULE_NAME_LOWER}Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	// 构建更新字段
	updates := make(map[string]interface{})
	
	if req.Name != nil {
		// 检查新名称是否已存在
		exists, err := s.${MODULE_NAME_LOWER}Repo.ExistsByName(ctx, *req.Name, id)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, errors.New(errors.CodeConflict, "名称已存在")
		}
		updates["name"] = *req.Name
	}
	
	if req.Description != nil {
		updates["description"] = *req.Description
	}
	
	if req.Status != nil {
		updates["status"] = *req.Status
	}
	
	// 执行更新
	if len(updates) > 0 {
		if err := s.${MODULE_NAME_LOWER}Repo.Update(ctx, id, updates); err != nil {
			return nil, err
		}
	}
	
	// 返回更新后的数据
	updated, err := s.${MODULE_NAME_LOWER}Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	return responseDto.From${MODULE_NAME_TITLE}(updated), nil
}

// Delete 删除${MODULE_NAME_LOWER}
func (s *${MODULE_NAME_LOWER}Service) Delete(ctx context.Context, id uint) error {
	// 检查${MODULE_NAME_LOWER}是否存在
	_, err := s.${MODULE_NAME_LOWER}Repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	
	// 执行删除
	return s.${MODULE_NAME_LOWER}Repo.Delete(ctx, id)
}

// List 获取${MODULE_NAME_LOWER}列表
func (s *${MODULE_NAME_LOWER}Service) List(ctx context.Context, req *requestDto.Query${MODULE_NAME_TITLE}Request) (*common.PaginationResponse, error) {
	// 构建过滤条件
	filters := make(map[string]interface{})
	
	if req.Name != "" {
		filters["name"] = req.Name
	}
	
	if req.Status != nil {
		filters["status"] = *req.Status
	}
	
	// 查询数据
	${MODULE_NAME_LOWER}s, total, err := s.${MODULE_NAME_LOWER}Repo.List(ctx, &req.PaginationRequest, filters)
	if err != nil {
		return nil, err
	}
	
	// 转换为响应DTO
	items := responseDto.From${MODULE_NAME_TITLE}List(${MODULE_NAME_LOWER}s)
	
	return &common.PaginationResponse{
		Items: items,
		Pagination: common.PaginationInfo{
			Page:       req.Page,
			PageSize:   req.PageSize,
			Total:      total,
			TotalPages: (total + int64(req.PageSize) - 1) / int64(req.PageSize),
		},
	}, nil
}
EOF

# 创建Controller
echo -e "${YELLOW}创建Controller...${NC}"
cat > "internal/controllers/${MODULE_NAME_LOWER}/${MODULE_NAME_LOWER}_controller.go" << EOF
package ${MODULE_NAME_LOWER}

import (
	"strconv"
	"template/internal/services/${MODULE_NAME_LOWER}"
	requestDto "template/internal/dto/request/${MODULE_NAME_LOWER}"
	"template/pkg/common"
	"template/pkg/errors"
	
	"github.com/gin-gonic/gin"
)

// ${MODULE_NAME_TITLE}Controller ${MODULE_NAME_LOWER}控制器
type ${MODULE_NAME_TITLE}Controller struct {
	${MODULE_NAME_LOWER}Service ${MODULE_NAME_LOWER}.${MODULE_NAME_TITLE}Service
}

// New${MODULE_NAME_TITLE}Controller 创建${MODULE_NAME_LOWER}控制器实例
func New${MODULE_NAME_TITLE}Controller(${MODULE_NAME_LOWER}Service ${MODULE_NAME_LOWER}.${MODULE_NAME_TITLE}Service) *${MODULE_NAME_TITLE}Controller {
	return &${MODULE_NAME_TITLE}Controller{
		${MODULE_NAME_LOWER}Service: ${MODULE_NAME_LOWER}Service,
	}
}

// Create 创建${MODULE_NAME_LOWER}
// @Summary 创建${MODULE_NAME_LOWER}
// @Description 创建新的${MODULE_NAME_LOWER}
// @Tags ${MODULE_NAME_TITLE}
// @Accept json
// @Produce json
// @Param request body requestDto.Create${MODULE_NAME_TITLE}Request true "创建${MODULE_NAME_LOWER}请求"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response
// @Router /api/v1/${MODULE_NAME_LOWER}s [post]
func (c *${MODULE_NAME_TITLE}Controller) Create(ctx *gin.Context) {
	req, err := common.ValidateRequest[requestDto.Create${MODULE_NAME_TITLE}Request](ctx)
	if err != nil {
		errors.HandleError(ctx, err)
		return
	}
	
	result, err := c.${MODULE_NAME_LOWER}Service.Create(ctx.Request.Context(), req)
	if err != nil {
		errors.HandleError(ctx, err)
		return
	}
	
	common.Success(ctx, result, "创建成功")
}

// GetByID 根据ID获取${MODULE_NAME_LOWER}
// @Summary 获取${MODULE_NAME_LOWER}详情
// @Description 根据ID获取${MODULE_NAME_LOWER}详情
// @Tags ${MODULE_NAME_TITLE}
// @Accept json
// @Produce json
// @Param id path int true "${MODULE_NAME_TITLE} ID"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response
// @Failure 404 {object} common.Response
// @Router /api/v1/${MODULE_NAME_LOWER}s/{id} [get]
func (c *${MODULE_NAME_TITLE}Controller) GetByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		errors.HandleError(ctx, errors.New(errors.CodeInvalidParameter, "无效的ID"))
		return
	}
	
	result, err := c.${MODULE_NAME_LOWER}Service.GetByID(ctx.Request.Context(), uint(id))
	if err != nil {
		errors.HandleError(ctx, err)
		return
	}
	
	common.Success(ctx, result, "获取成功")
}

// Update 更新${MODULE_NAME_LOWER}
// @Summary 更新${MODULE_NAME_LOWER}
// @Description 更新${MODULE_NAME_LOWER}信息
// @Tags ${MODULE_NAME_TITLE}
// @Accept json
// @Produce json
// @Param id path int true "${MODULE_NAME_TITLE} ID"
// @Param request body requestDto.Update${MODULE_NAME_TITLE}Request true "更新${MODULE_NAME_LOWER}请求"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response
// @Failure 404 {object} common.Response
// @Router /api/v1/${MODULE_NAME_LOWER}s/{id} [put]
func (c *${MODULE_NAME_TITLE}Controller) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		errors.HandleError(ctx, errors.New(errors.CodeInvalidParameter, "无效的ID"))
		return
	}
	
	req, err := common.ValidateRequest[requestDto.Update${MODULE_NAME_TITLE}Request](ctx)
	if err != nil {
		errors.HandleError(ctx, err)
		return
	}
	
	result, err := c.${MODULE_NAME_LOWER}Service.Update(ctx.Request.Context(), uint(id), req)
	if err != nil {
		errors.HandleError(ctx, err)
		return
	}
	
	common.Success(ctx, result, "更新成功")
}

// Delete 删除${MODULE_NAME_LOWER}
// @Summary 删除${MODULE_NAME_LOWER}
// @Description 删除${MODULE_NAME_LOWER}
// @Tags ${MODULE_NAME_TITLE}
// @Accept json
// @Produce json
// @Param id path int true "${MODULE_NAME_TITLE} ID"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response
// @Failure 404 {object} common.Response
// @Router /api/v1/${MODULE_NAME_LOWER}s/{id} [delete]
func (c *${MODULE_NAME_TITLE}Controller) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		errors.HandleError(ctx, errors.New(errors.CodeInvalidParameter, "无效的ID"))
		return
	}
	
	err = c.${MODULE_NAME_LOWER}Service.Delete(ctx.Request.Context(), uint(id))
	if err != nil {
		errors.HandleError(ctx, err)
		return
	}
	
	common.Success(ctx, nil, "删除成功")
}

// List 获取${MODULE_NAME_LOWER}列表
// @Summary 获取${MODULE_NAME_LOWER}列表
// @Description 分页获取${MODULE_NAME_LOWER}列表
// @Tags ${MODULE_NAME_TITLE}
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Param name query string false "名称搜索"
// @Param status query int false "状态筛选"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response
// @Router /api/v1/${MODULE_NAME_LOWER}s [get]
func (c *${MODULE_NAME_TITLE}Controller) List(ctx *gin.Context) {
	req, err := common.ValidateRequest[requestDto.Query${MODULE_NAME_TITLE}Request](ctx)
	if err != nil {
		errors.HandleError(ctx, err)
		return
	}
	
	// 设置默认分页参数
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}
	
	result, err := c.${MODULE_NAME_LOWER}Service.List(ctx.Request.Context(), req)
	if err != nil {
		errors.HandleError(ctx, err)
		return
	}
	
	common.Success(ctx, result, "获取成功")
}
EOF

# 创建路由文件
echo -e "${YELLOW}创建路由文件...${NC}"
cat > "internal/routes/api/v1/${MODULE_NAME_LOWER}.go" << EOF
package v1

import (
	"template/internal/controllers/${MODULE_NAME_LOWER}"
	"template/internal/repositories/${MODULE_NAME_LOWER}"
	"template/internal/services/${MODULE_NAME_LOWER}"
	"template/pkg/database"
	
	"github.com/gin-gonic/gin"
)

// Register${MODULE_NAME_TITLE}Routes 注册${MODULE_NAME_LOWER}路由
func Register${MODULE_NAME_TITLE}Routes(r *gin.RouterGroup) {
	// 初始化依赖
	${MODULE_NAME_LOWER}Repo := ${MODULE_NAME_LOWER}.New${MODULE_NAME_TITLE}Repository(database.GetDB())
	${MODULE_NAME_LOWER}Service := ${MODULE_NAME_LOWER}.New${MODULE_NAME_TITLE}Service(${MODULE_NAME_LOWER}Repo)
	${MODULE_NAME_LOWER}Controller := ${MODULE_NAME_LOWER}.New${MODULE_NAME_TITLE}Controller(${MODULE_NAME_LOWER}Service)
	
	// 注册路由
	${MODULE_NAME_LOWER}Group := r.Group("/${MODULE_NAME_LOWER}s")
	{
		${MODULE_NAME_LOWER}Group.POST("", ${MODULE_NAME_LOWER}Controller.Create)
		${MODULE_NAME_LOWER}Group.GET("/:id", ${MODULE_NAME_LOWER}Controller.GetByID)
		${MODULE_NAME_LOWER}Group.PUT("/:id", ${MODULE_NAME_LOWER}Controller.Update)
		${MODULE_NAME_LOWER}Group.DELETE("/:id", ${MODULE_NAME_LOWER}Controller.Delete)
		${MODULE_NAME_LOWER}Group.GET("", ${MODULE_NAME_LOWER}Controller.List)
	}
}
EOF

# 创建迁移文件
echo -e "${YELLOW}创建数据库迁移文件...${NC}"
mkdir -p internal/migrations
timestamp=$(date +%Y%m%d%H%M%S)
cat > "internal/migrations/${timestamp}_create_${MODULE_NAME_LOWER}s_table.sql" << EOF
-- +migrate Up
CREATE TABLE ${MODULE_NAME_LOWER}s (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(500) DEFAULT '',
    status TINYINT DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    
    INDEX idx_name (name),
    INDEX idx_status (status),
    INDEX idx_deleted_at (deleted_at)
);

-- +migrate Down
DROP TABLE IF EXISTS ${MODULE_NAME_LOWER}s;
EOF

# 创建测试文件
echo -e "${YELLOW}创建测试文件...${NC}"
mkdir -p "tests/unit/services"
cat > "tests/unit/services/${MODULE_NAME_LOWER}_service_test.go" << EOF
package services

import (
	"context"
	"testing"
	"template/internal/models"
	requestDto "template/internal/dto/request/${MODULE_NAME_LOWER}"
	"template/internal/services/${MODULE_NAME_LOWER}"
	"template/pkg/errors"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock${MODULE_NAME_TITLE}Repository 模拟${MODULE_NAME_LOWER}仓库
type Mock${MODULE_NAME_TITLE}Repository struct {
	mock.Mock
}

func (m *Mock${MODULE_NAME_TITLE}Repository) Create(ctx context.Context, ${MODULE_NAME_LOWER} *models.${MODULE_NAME_TITLE}) error {
	args := m.Called(ctx, ${MODULE_NAME_LOWER})
	return args.Error(0)
}

func (m *Mock${MODULE_NAME_TITLE}Repository) GetByID(ctx context.Context, id uint) (*models.${MODULE_NAME_TITLE}, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.${MODULE_NAME_TITLE}), args.Error(1)
}

func (m *Mock${MODULE_NAME_TITLE}Repository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	args := m.Called(ctx, id, updates)
	return args.Error(0)
}

func (m *Mock${MODULE_NAME_TITLE}Repository) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *Mock${MODULE_NAME_TITLE}Repository) List(ctx context.Context, req *common.PaginationRequest, filters map[string]interface{}) ([]*models.${MODULE_NAME_TITLE}, int64, error) {
	args := m.Called(ctx, req, filters)
	return args.Get(0).([]*models.${MODULE_NAME_TITLE}), args.Get(1).(int64), args.Error(2)
}

func (m *Mock${MODULE_NAME_TITLE}Repository) ExistsByName(ctx context.Context, name string, excludeID ...uint) (bool, error) {
	args := m.Called(ctx, name, excludeID)
	return args.Bool(0), args.Error(1)
}

func Test${MODULE_NAME_TITLE}Service_Create(t *testing.T) {
	mockRepo := new(Mock${MODULE_NAME_TITLE}Repository)
	service := ${MODULE_NAME_LOWER}.New${MODULE_NAME_TITLE}Service(mockRepo)
	
	tests := []struct {
		name    string
		req     *requestDto.Create${MODULE_NAME_TITLE}Request
		setup   func()
		wantErr bool
		errCode errors.ErrorCode
	}{
		{
			name: "正常创建",
			req: &requestDto.Create${MODULE_NAME_TITLE}Request{
				Name:        "测试${MODULE_NAME_LOWER}",
				Description: "测试描述",
			},
			setup: func() {
				mockRepo.On("ExistsByName", mock.Anything, "测试${MODULE_NAME_LOWER}", mock.Anything).Return(false, nil)
				mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.${MODULE_NAME_TITLE}")).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "名称已存在",
			req: &requestDto.Create${MODULE_NAME_TITLE}Request{
				Name:        "已存在的${MODULE_NAME_LOWER}",
				Description: "测试描述",
			},
			setup: func() {
				mockRepo.On("ExistsByName", mock.Anything, "已存在的${MODULE_NAME_LOWER}", mock.Anything).Return(true, nil)
			},
			wantErr: true,
			errCode: errors.CodeConflict,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.ExpectedCalls = nil
			tt.setup()
			
			result, err := service.Create(context.Background(), tt.req)
			
			if tt.wantErr {
				assert.Error(t, err)
				if tt.errCode != 0 {
					assert.True(t, errors.Is(err, tt.errCode))
				}
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.req.Name, result.Name)
			}
			
			mockRepo.AssertExpectations(t)
		})
	}
}
EOF

echo -e "${GREEN}模块创建完成！${NC}"
echo ""
echo -e "${BLUE}已创建的文件:${NC}"
echo "  - internal/models/${MODULE_NAME_LOWER}.go"
echo "  - internal/dto/request/${MODULE_NAME_LOWER}/create.go"
echo "  - internal/dto/request/${MODULE_NAME_LOWER}/update.go"
echo "  - internal/dto/request/${MODULE_NAME_LOWER}/query.go"
echo "  - internal/dto/response/${MODULE_NAME_LOWER}/${MODULE_NAME_LOWER}_info.go"
echo "  - internal/repositories/${MODULE_NAME_LOWER}/${MODULE_NAME_LOWER}_repository.go"
echo "  - internal/services/${MODULE_NAME_LOWER}/${MODULE_NAME_LOWER}_service.go"
echo "  - internal/controllers/${MODULE_NAME_LOWER}/${MODULE_NAME_LOWER}_controller.go"
echo "  - internal/routes/api/v1/${MODULE_NAME_LOWER}.go"
echo "  - internal/migrations/${timestamp}_create_${MODULE_NAME_LOWER}s_table.sql"
echo "  - tests/unit/services/${MODULE_NAME_LOWER}_service_test.go"
echo ""
echo -e "${YELLOW}下一步操作:${NC}"
echo "1. 在 internal/routes/api/v1/routes.go 中注册新路由:"
echo "   Register${MODULE_NAME_TITLE}Routes(v1)"
echo ""
echo "2. 运行数据库迁移:"
echo "   make migrate"
echo ""
echo "3. 运行测试:"
echo "   make test"
echo ""
echo -e "${GREEN}模块 ${MODULE_NAME_TITLE} 创建完成！${NC}"
EOF 