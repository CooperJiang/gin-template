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
mkdir -p "internal/dto/request"
mkdir -p "internal/dto/response"

# 创建模型文件
echo -e "${YELLOW}创建模型文件...${NC}"
cat > "internal/models/${MODULE_NAME_LOWER}.go" << EOF
package models

import (
	"email-manage/pkg/common"
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
	return "${MODULE_NAME_LOWER}"
}

// BeforeCreate 创建前钩子
func (m *${MODULE_NAME_TITLE}) BeforeCreate(tx *gorm.DB) error {
	// 调用基础模型的BeforeCreate（生成UUID）
	if err := m.BaseModel.BeforeCreate(tx); err != nil {
		return err
	}
	
	// 设置默认值
	if m.Status == 0 {
		m.Status = common.StatusActive
	}
	return nil
}

// BeforeUpdate 更新前钩子
func (m *${MODULE_NAME_TITLE}) BeforeUpdate(tx *gorm.DB) error {
	// 调用基础模型的BeforeUpdate
	return m.BaseModel.BeforeUpdate(tx)
}

// IsActive 是否激活状态
func (m *${MODULE_NAME_TITLE}) IsActive() bool {
	return m.Status == common.StatusActive
}
EOF

# 创建请求DTO
echo -e "${YELLOW}创建请求DTO...${NC}"
cat > "internal/dto/request/${MODULE_NAME_LOWER}.go" << EOF
package request

import "email-manage/internal/dto"

// Create${MODULE_NAME_TITLE}Request 创建${MODULE_NAME_LOWER}请求
type Create${MODULE_NAME_TITLE}Request struct {
	dto.BaseRequest
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

// Update${MODULE_NAME_TITLE}Request 更新${MODULE_NAME_LOWER}请求
type Update${MODULE_NAME_TITLE}Request struct {
	dto.BaseRequest
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

// Query${MODULE_NAME_TITLE}Request 查询${MODULE_NAME_LOWER}请求
type Query${MODULE_NAME_TITLE}Request struct {
	dto.BaseRequest
	Name   string \`form:"name" binding:"omitempty,max=100"\`
	Status *int   \`form:"status" binding:"omitempty,oneof=0 1"\`
	Page   int    \`form:"page" binding:"omitempty,min=1"\`
	Size   int    \`form:"size" binding:"omitempty,min=1,max=100"\`
}

// GetValidationMessages 获取验证错误信息
func (r *Query${MODULE_NAME_TITLE}Request) GetValidationMessages() map[string]string {
	return map[string]string{
		"Name.max":     "名称长度不能超过100个字符",
		"Status.oneof": "状态值必须是0或1",
		"Page.min":     "页码必须大于0",
		"Size.min":     "每页数量必须大于0",
		"Size.max":     "每页数量不能超过100",
	}
}
EOF

# 创建响应DTO
echo -e "${YELLOW}创建响应DTO...${NC}"
cat > "internal/dto/response/${MODULE_NAME_LOWER}.go" << EOF
package response

import (
	"email-manage/internal/dto"
	"email-manage/internal/models"
	"time"
)

// ${MODULE_NAME_TITLE}Info ${MODULE_NAME_LOWER}信息
type ${MODULE_NAME_TITLE}Info struct {
	dto.BaseResponse
	Name        string \`json:"name"\`
	Description string \`json:"description"\`
	Status      int    \`json:"status"\`
}

// FromModel 从模型转换
func (r *${MODULE_NAME_TITLE}Info) FromModel(${MODULE_NAME_LOWER} models.${MODULE_NAME_TITLE}) *${MODULE_NAME_TITLE}Info {
	return &${MODULE_NAME_TITLE}Info{
		BaseResponse: dto.BaseResponse{
			ID:        ${MODULE_NAME_LOWER}.ID.String(),
			CreatedAt: time.Time(${MODULE_NAME_LOWER}.CreatedAt),
			UpdatedAt: time.Time(${MODULE_NAME_LOWER}.UpdatedAt),
		},
		Name:        ${MODULE_NAME_LOWER}.Name,
		Description: ${MODULE_NAME_LOWER}.Description,
		Status:      ${MODULE_NAME_LOWER}.Status,
	}
}

// ${MODULE_NAME_TITLE}ListResponse ${MODULE_NAME_LOWER}列表响应
type ${MODULE_NAME_TITLE}ListResponse struct {
	Items      []${MODULE_NAME_TITLE}Info \`json:"items"\`
	Total      int64                      \`json:"total"\`
	Page       int                        \`json:"page"\`
	PageSize   int                        \`json:"page_size"\`
	TotalPages int                        \`json:"total_pages"\`
}

// Create${MODULE_NAME_TITLE}Response 创建${MODULE_NAME_LOWER}响应
type Create${MODULE_NAME_TITLE}Response struct {
	Message string                \`json:"message"\`
	Data    ${MODULE_NAME_TITLE}Info \`json:"data"\`
}
EOF

# 创建Repository
echo -e "${YELLOW}创建Repository...${NC}"
cat > "internal/repositories/${MODULE_NAME_LOWER}/${MODULE_NAME_LOWER}_repository.go" << EOF
package ${MODULE_NAME_LOWER}

import (
	"context"
	"email-manage/internal/models"
	"email-manage/internal/repositories"
	"email-manage/pkg/errors"
	
	"gorm.io/gorm"
)

// ${MODULE_NAME_TITLE}Repository ${MODULE_NAME_LOWER}数据访问接口
type ${MODULE_NAME_TITLE}Repository interface {
	repositories.BaseRepository[models.${MODULE_NAME_TITLE}]
	GetByName(ctx context.Context, name string) (*models.${MODULE_NAME_TITLE}, error)
	ExistsByName(ctx context.Context, name string, excludeID ...string) (bool, error)
}

// ${MODULE_NAME_LOWER}Repository ${MODULE_NAME_LOWER}数据访问实现
type ${MODULE_NAME_LOWER}Repository struct {
	repositories.BaseRepository[models.${MODULE_NAME_TITLE}]
	db *gorm.DB
}

// New${MODULE_NAME_TITLE}Repository 创建${MODULE_NAME_LOWER}数据访问实例
func New${MODULE_NAME_TITLE}Repository(db *gorm.DB) ${MODULE_NAME_TITLE}Repository {
	return &${MODULE_NAME_LOWER}Repository{
		BaseRepository: repositories.NewBaseRepository[models.${MODULE_NAME_TITLE}](db),
		db:             db,
	}
}

// GetByName 根据名称获取${MODULE_NAME_LOWER}
func (r *${MODULE_NAME_LOWER}Repository) GetByName(ctx context.Context, name string) (*models.${MODULE_NAME_TITLE}, error) {
	var ${MODULE_NAME_LOWER} models.${MODULE_NAME_TITLE}
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&${MODULE_NAME_LOWER}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New(errors.CodeNotFound, "${MODULE_NAME_TITLE}不存在")
		}
		return nil, errors.Wrap(err, errors.CodeQueryFailed)
	}
	return &${MODULE_NAME_LOWER}, nil
}

// ExistsByName 检查名称是否存在
func (r *${MODULE_NAME_LOWER}Repository) ExistsByName(ctx context.Context, name string, excludeID ...string) (bool, error) {
	var count int64
	query := r.db.WithContext(ctx).Model(&models.${MODULE_NAME_TITLE}{}).Where("name = ?", name)
	
	if len(excludeID) > 0 {
		query = query.Where("id != ?", excludeID[0])
	}
	
	err := query.Count(&count).Error
	if err != nil {
		return false, errors.Wrap(err, errors.CodeQueryFailed)
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
	"email-manage/internal/models"
	"email-manage/internal/repositories/${MODULE_NAME_LOWER}"
	"email-manage/internal/dto/request"
	"email-manage/internal/dto/response"
	"email-manage/pkg/common"
	"email-manage/pkg/errors"
)

// ${MODULE_NAME_TITLE}Service ${MODULE_NAME_LOWER}服务接口
type ${MODULE_NAME_TITLE}Service interface {
	Create(ctx context.Context, req *request.Create${MODULE_NAME_TITLE}Request) (*response.${MODULE_NAME_TITLE}Info, error)
	GetByID(ctx context.Context, id string) (*response.${MODULE_NAME_TITLE}Info, error)
	Update(ctx context.Context, id string, req *request.Update${MODULE_NAME_TITLE}Request) (*response.${MODULE_NAME_TITLE}Info, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, req *request.Query${MODULE_NAME_TITLE}Request) (*response.${MODULE_NAME_TITLE}ListResponse, error)
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
func (s *${MODULE_NAME_LOWER}Service) Create(ctx context.Context, req *request.Create${MODULE_NAME_TITLE}Request) (*response.${MODULE_NAME_TITLE}Info, error) {
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
		Status:      common.StatusActive, // 默认启用
	}
	
	if err := s.${MODULE_NAME_LOWER}Repo.Create(ctx, ${MODULE_NAME_LOWER}); err != nil {
		return nil, err
	}
	
	return (&response.${MODULE_NAME_TITLE}Info{}).FromModel(*${MODULE_NAME_LOWER}), nil
}

// GetByID 根据ID获取${MODULE_NAME_LOWER}
func (s *${MODULE_NAME_LOWER}Service) GetByID(ctx context.Context, id string) (*response.${MODULE_NAME_TITLE}Info, error) {
	${MODULE_NAME_LOWER}, err := s.${MODULE_NAME_LOWER}Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	return (&response.${MODULE_NAME_TITLE}Info{}).FromModel(*${MODULE_NAME_LOWER}), nil
}

// Update 更新${MODULE_NAME_LOWER}
func (s *${MODULE_NAME_LOWER}Service) Update(ctx context.Context, id string, req *request.Update${MODULE_NAME_TITLE}Request) (*response.${MODULE_NAME_TITLE}Info, error) {
	// 检查${MODULE_NAME_LOWER}是否存在
	existing, err := s.${MODULE_NAME_LOWER}Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	// 构建更新数据
	updates := make(map[string]interface{})
	
	if req.Name != nil {
		// 检查名称是否已被其他记录使用
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
	
	// 获取更新后的数据
	updated, err := s.${MODULE_NAME_LOWER}Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	return (&response.${MODULE_NAME_TITLE}Info{}).FromModel(*updated), nil
}

// Delete 删除${MODULE_NAME_LOWER}
func (s *${MODULE_NAME_LOWER}Service) Delete(ctx context.Context, id string) error {
	// 检查${MODULE_NAME_LOWER}是否存在
	_, err := s.${MODULE_NAME_LOWER}Repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	
	return s.${MODULE_NAME_LOWER}Repo.Delete(ctx, id)
}

// List 获取${MODULE_NAME_LOWER}列表
func (s *${MODULE_NAME_LOWER}Service) List(ctx context.Context, req *request.Query${MODULE_NAME_TITLE}Request) (*response.${MODULE_NAME_TITLE}ListResponse, error) {
	// 设置默认分页参数
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 20
	}
	
	// 构建分页请求
	paginationReq := &common.PaginationRequest{
		Page:     req.Page,
		PageSize: req.Size,
	}
	
	// 构建过滤条件
	filters := make(map[string]interface{})
	if req.Name != "" {
		filters["keyword"] = req.Name
	}
	if req.Status != nil {
		filters["status"] = *req.Status
	}
	
	// 查询数据
	${MODULE_NAME_LOWER}s, total, err := s.${MODULE_NAME_LOWER}Repo.List(ctx, paginationReq, filters)
	if err != nil {
		return nil, err
	}
	
	// 转换为响应DTO
	items := make([]response.${MODULE_NAME_TITLE}Info, len(${MODULE_NAME_LOWER}s))
	for i, ${MODULE_NAME_LOWER} := range ${MODULE_NAME_LOWER}s {
		items[i] = *(&response.${MODULE_NAME_TITLE}Info{}).FromModel(*${MODULE_NAME_LOWER})
	}
	
	totalPages := int((total + int64(req.Size) - 1) / int64(req.Size))
	
	return &response.${MODULE_NAME_TITLE}ListResponse{
		Items:      items,
		Total:      total,
		Page:       req.Page,
		PageSize:   req.Size,
		TotalPages: totalPages,
	}, nil
}
EOF

# 创建Controller
echo -e "${YELLOW}创建Controller...${NC}"
cat > "internal/controllers/${MODULE_NAME_LOWER}/${MODULE_NAME_LOWER}_controller.go" << EOF
package ${MODULE_NAME_LOWER}

import (
	"email-manage/internal/services/${MODULE_NAME_LOWER}"
	"email-manage/internal/dto/request"
	"email-manage/pkg/common"
	"email-manage/pkg/errors"
	
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
// @Param request body request.Create${MODULE_NAME_TITLE}Request true "创建${MODULE_NAME_LOWER}请求"
// @Success 200 {object} errors.Response
// @Failure 400 {object} errors.Response
// @Router /api/v1/${MODULE_NAME_LOWER}s [post]
func (c *${MODULE_NAME_TITLE}Controller) Create(ctx *gin.Context) {
	req, err := common.ValidateRequest[request.Create${MODULE_NAME_TITLE}Request](ctx)
	if err != nil {
		errors.HandleError(ctx, err)
		return
	}
	
	result, err := c.${MODULE_NAME_LOWER}Service.Create(ctx.Request.Context(), req)
	if err != nil {
		errors.HandleError(ctx, err)
		return
	}
	
	errors.ResponseSuccess(ctx, result, "创建成功")
}

// GetByID 根据ID获取${MODULE_NAME_LOWER}
// @Summary 获取${MODULE_NAME_LOWER}详情
// @Description 根据ID获取${MODULE_NAME_LOWER}详情
// @Tags ${MODULE_NAME_TITLE}
// @Accept json
// @Produce json
// @Param id path string true "${MODULE_NAME_TITLE} ID (UUID)"
// @Success 200 {object} errors.Response
// @Failure 400 {object} errors.Response
// @Failure 404 {object} errors.Response
// @Router /api/v1/${MODULE_NAME_LOWER}s/{id} [get]
func (c *${MODULE_NAME_TITLE}Controller) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if !common.ValidateUUID(id) {
		errors.HandleError(ctx, errors.New(errors.CodeInvalidParameter, "无效的ID格式"))
		return
	}
	
	result, err := c.${MODULE_NAME_LOWER}Service.GetByID(ctx.Request.Context(), id)
	if err != nil {
		errors.HandleError(ctx, err)
		return
	}
	
	errors.ResponseSuccess(ctx, result, "获取成功")
}

// Update 更新${MODULE_NAME_LOWER}
// @Summary 更新${MODULE_NAME_LOWER}
// @Description 更新${MODULE_NAME_LOWER}信息
// @Tags ${MODULE_NAME_TITLE}
// @Accept json
// @Produce json
// @Param id path string true "${MODULE_NAME_TITLE} ID (UUID)"
// @Param request body request.Update${MODULE_NAME_TITLE}Request true "更新${MODULE_NAME_LOWER}请求"
// @Success 200 {object} errors.Response
// @Failure 400 {object} errors.Response
// @Failure 404 {object} errors.Response
// @Router /api/v1/${MODULE_NAME_LOWER}s/{id} [put]
func (c *${MODULE_NAME_TITLE}Controller) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	if !common.ValidateUUID(id) {
		errors.HandleError(ctx, errors.New(errors.CodeInvalidParameter, "无效的ID格式"))
		return
	}
	
	req, err := common.ValidateRequest[request.Update${MODULE_NAME_TITLE}Request](ctx)
	if err != nil {
		errors.HandleError(ctx, err)
		return
	}
	
	result, err := c.${MODULE_NAME_LOWER}Service.Update(ctx.Request.Context(), id, req)
	if err != nil {
		errors.HandleError(ctx, err)
		return
	}
	
	errors.ResponseSuccess(ctx, result, "更新成功")
}

// Delete 删除${MODULE_NAME_LOWER}
// @Summary 删除${MODULE_NAME_LOWER}
// @Description 删除${MODULE_NAME_LOWER}
// @Tags ${MODULE_NAME_TITLE}
// @Accept json
// @Produce json
// @Param id path string true "${MODULE_NAME_TITLE} ID (UUID)"
// @Success 200 {object} errors.Response
// @Failure 400 {object} errors.Response
// @Failure 404 {object} errors.Response
// @Router /api/v1/${MODULE_NAME_LOWER}s/{id} [delete]
func (c *${MODULE_NAME_TITLE}Controller) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if !common.ValidateUUID(id) {
		errors.HandleError(ctx, errors.New(errors.CodeInvalidParameter, "无效的ID格式"))
		return
	}
	
	err := c.${MODULE_NAME_LOWER}Service.Delete(ctx.Request.Context(), id)
	if err != nil {
		errors.HandleError(ctx, err)
		return
	}
	
	errors.ResponseSuccess(ctx, nil, "删除成功")
}

// List 获取${MODULE_NAME_LOWER}列表
// @Summary 获取${MODULE_NAME_LOWER}列表
// @Description 分页获取${MODULE_NAME_LOWER}列表
// @Tags ${MODULE_NAME_TITLE}
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param size query int false "每页数量" default(20)
// @Param name query string false "名称搜索"
// @Param status query int false "状态筛选"
// @Success 200 {object} errors.Response
// @Failure 400 {object} errors.Response
// @Router /api/v1/${MODULE_NAME_LOWER}s [get]
func (c *${MODULE_NAME_TITLE}Controller) List(ctx *gin.Context) {
	req, err := common.ValidateRequest[request.Query${MODULE_NAME_TITLE}Request](ctx)
	if err != nil {
		errors.HandleError(ctx, err)
		return
	}
	
	result, err := c.${MODULE_NAME_LOWER}Service.List(ctx.Request.Context(), req)
	if err != nil {
		errors.HandleError(ctx, err)
		return
	}
	
	errors.ResponseSuccess(ctx, result, "获取成功")
}
EOF

# 创建路由文件
echo -e "${YELLOW}创建路由文件...${NC}"
mkdir -p "internal/routes/api/v1"
cat > "internal/routes/api/v1/${MODULE_NAME_LOWER}.go" << EOF
package v1

import (
	"email-manage/internal/controllers/${MODULE_NAME_LOWER}"
	"email-manage/internal/repositories/${MODULE_NAME_LOWER}"
	"email-manage/internal/services/${MODULE_NAME_LOWER}"
	"email-manage/pkg/database"
	
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
cat > "internal/migrations/${timestamp}_create_${MODULE_NAME_LOWER}_table.sql" << EOF
-- +migrate Up
-- 创建${MODULE_NAME_LOWER}表
CREATE TABLE ${MODULE_NAME_LOWER} (
    id CHAR(36) PRIMARY KEY,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME NULL,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(500) DEFAULT '',
    status INT DEFAULT 1,
    
    KEY idx_name (name),
    KEY idx_status (status),
    KEY idx_deleted_at (deleted_at)
);

-- +migrate Down
DROP TABLE IF EXISTS ${MODULE_NAME_LOWER};
EOF

# 添加常量定义
echo -e "${YELLOW}更新常量定义...${NC}"
if [ ! -f "pkg/constants/common.go" ]; then
    mkdir -p pkg/constants
    cat > "pkg/constants/common.go" << EOF
package constants

// 通用状态常量
const (
	StatusInactive = 0 // 未激活
	StatusActive   = 1 // 激活
)
EOF
else
    # 检查是否已存在状态常量
    if ! grep -q "StatusActive" pkg/constants/common.go; then
        echo "" >> pkg/constants/common.go
        echo "// 通用状态常量" >> pkg/constants/common.go
        echo "const (" >> pkg/constants/common.go
        echo "	StatusInactive = 0 // 未激活" >> pkg/constants/common.go
        echo "	StatusActive   = 1 // 激活" >> pkg/constants/common.go
        echo ")" >> pkg/constants/common.go
    fi
fi

# 创建测试文件
echo -e "${YELLOW}创建测试文件...${NC}"
mkdir -p "tests/unit/services"
cat > "tests/unit/services/${MODULE_NAME_LOWER}_service_test.go" << EOF
package services

import (
	"context"
	"testing"
	"email-manage/internal/models"
	"email-manage/internal/dto/request"
	"email-manage/internal/services/${MODULE_NAME_LOWER}"
	"email-manage/pkg/common"
	"email-manage/pkg/errors"
	
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

func (m *Mock${MODULE_NAME_TITLE}Repository) GetByID(ctx context.Context, id string) (*models.${MODULE_NAME_TITLE}, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.${MODULE_NAME_TITLE}), args.Error(1)
}

func (m *Mock${MODULE_NAME_TITLE}Repository) Update(ctx context.Context, id string, updates map[string]interface{}) error {
	args := m.Called(ctx, id, updates)
	return args.Error(0)
}

func (m *Mock${MODULE_NAME_TITLE}Repository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *Mock${MODULE_NAME_TITLE}Repository) List(ctx context.Context, req *common.PaginationRequest, filters map[string]interface{}) ([]*models.${MODULE_NAME_TITLE}, int64, error) {
	args := m.Called(ctx, req, filters)
	return args.Get(0).([]*models.${MODULE_NAME_TITLE}), args.Get(1).(int64), args.Error(2)
}

func (m *Mock${MODULE_NAME_TITLE}Repository) Count(ctx context.Context, filters map[string]interface{}) (int64, error) {
	args := m.Called(ctx, filters)
	return args.Get(0).(int64), args.Error(1)
}

func (m *Mock${MODULE_NAME_TITLE}Repository) Exists(ctx context.Context, id string) (bool, error) {
	args := m.Called(ctx, id)
	return args.Bool(0), args.Error(1)
}

func (m *Mock${MODULE_NAME_TITLE}Repository) BatchCreate(ctx context.Context, entities []*models.${MODULE_NAME_TITLE}) error {
	args := m.Called(ctx, entities)
	return args.Error(0)
}

func (m *Mock${MODULE_NAME_TITLE}Repository) BatchDelete(ctx context.Context, ids []string) error {
	args := m.Called(ctx, ids)
	return args.Error(0)
}

func (m *Mock${MODULE_NAME_TITLE}Repository) GetByName(ctx context.Context, name string) (*models.${MODULE_NAME_TITLE}, error) {
	args := m.Called(ctx, name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.${MODULE_NAME_TITLE}), args.Error(1)
}

func (m *Mock${MODULE_NAME_TITLE}Repository) ExistsByName(ctx context.Context, name string, excludeID ...string) (bool, error) {
	args := m.Called(ctx, name, excludeID)
	return args.Bool(0), args.Error(1)
}

func Test${MODULE_NAME_TITLE}Service_Create(t *testing.T) {
	mockRepo := new(Mock${MODULE_NAME_TITLE}Repository)
	service := ${MODULE_NAME_LOWER}.New${MODULE_NAME_TITLE}Service(mockRepo)
	
	tests := []struct {
		name    string
		req     *request.Create${MODULE_NAME_TITLE}Request
		setup   func()
		wantErr bool
		errCode errors.ErrorCode
	}{
		{
			name: "正常创建",
			req: &request.Create${MODULE_NAME_TITLE}Request{
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
			req: &request.Create${MODULE_NAME_TITLE}Request{
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
					customErr, ok := err.(*errors.CustomError)
					assert.True(t, ok)
					assert.Equal(t, tt.errCode, customErr.Code)
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
echo "  - internal/dto/request/${MODULE_NAME_LOWER}.go"
echo "  - internal/dto/response/${MODULE_NAME_LOWER}.go"
echo "  - internal/repositories/${MODULE_NAME_LOWER}/${MODULE_NAME_LOWER}_repository.go"
echo "  - internal/services/${MODULE_NAME_LOWER}/${MODULE_NAME_LOWER}_service.go"
echo "  - internal/controllers/${MODULE_NAME_LOWER}/${MODULE_NAME_LOWER}_controller.go"
echo "  - internal/routes/api/v1/${MODULE_NAME_LOWER}.go"
echo "  - internal/migrations/${timestamp}_create_${MODULE_NAME_LOWER}_table.sql"
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
echo "4. 启动开发服务器:"
echo "   make dev"
echo ""
echo -e "${GREEN}🎉 模块 ${MODULE_NAME_TITLE} 创建完成！${NC}"
echo -e "${BLUE}📝 特性说明:${NC}"
echo "  ✅ 使用UUID主键，确保数据安全"
echo "  ✅ 完整的CRUD操作"
echo "  ✅ 统一的错误处理"
echo "  ✅ 参数验证和响应格式"
echo "  ✅ 单元测试模板"
echo "  ✅ Swagger API文档注释"
EOF 