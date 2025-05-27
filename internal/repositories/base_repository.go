package repositories

import (
	"context"
	"template/pkg/common"
	"template/pkg/errors"

	"gorm.io/gorm"
)

// BaseRepository 基础仓库接口
type BaseRepository[T any] interface {
	Create(ctx context.Context, entity *T) error
	GetByID(ctx context.Context, id string) (*T, error)
	Update(ctx context.Context, id string, updates map[string]interface{}) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, req *common.PaginationRequest, filters map[string]interface{}) ([]*T, int64, error)
	Count(ctx context.Context, filters map[string]interface{}) (int64, error)
	Exists(ctx context.Context, id string) (bool, error)
	BatchCreate(ctx context.Context, entities []*T) error
	BatchDelete(ctx context.Context, ids []string) error
}

// baseRepository 基础仓库实现
type baseRepository[T any] struct {
	db *gorm.DB
}

// NewBaseRepository 创建基础仓库实例
func NewBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{
		db: db,
	}
}

// Create 创建实体
func (r *baseRepository[T]) Create(ctx context.Context, entity *T) error {
	if err := r.db.WithContext(ctx).Create(entity).Error; err != nil {
		return errors.Wrap(err, errors.CodeQueryFailed)
	}
	return nil
}

// GetByID 根据ID获取实体
func (r *baseRepository[T]) GetByID(ctx context.Context, id string) (*T, error) {
	// 验证UUID格式
	if !common.ValidateUUID(id) {
		return nil, errors.New(errors.CodeInvalidParameter, "无效的ID格式")
	}

	var entity T
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&entity).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New(errors.CodeNotFound, "记录不存在")
		}
		return nil, errors.Wrap(err, errors.CodeQueryFailed)
	}
	return &entity, nil
}

// Update 更新实体
func (r *baseRepository[T]) Update(ctx context.Context, id string, updates map[string]interface{}) error {
	// 验证UUID格式
	if !common.ValidateUUID(id) {
		return errors.New(errors.CodeInvalidParameter, "无效的ID格式")
	}

	var entity T
	result := r.db.WithContext(ctx).Model(&entity).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return errors.Wrap(result.Error, errors.CodeQueryFailed)
	}
	if result.RowsAffected == 0 {
		return errors.New(errors.CodeNotFound, "记录不存在")
	}
	return nil
}

// Delete 删除实体
func (r *baseRepository[T]) Delete(ctx context.Context, id string) error {
	// 验证UUID格式
	if !common.ValidateUUID(id) {
		return errors.New(errors.CodeInvalidParameter, "无效的ID格式")
	}

	var entity T
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity)
	if result.Error != nil {
		return errors.Wrap(result.Error, errors.CodeQueryFailed)
	}
	if result.RowsAffected == 0 {
		return errors.New(errors.CodeNotFound, "记录不存在")
	}
	return nil
}

// List 获取实体列表
func (r *baseRepository[T]) List(ctx context.Context, req *common.PaginationRequest, filters map[string]interface{}) ([]*T, int64, error) {
	var entities []*T
	var total int64

	query := r.db.WithContext(ctx).Model(new(T))

	// 应用过滤条件
	query = r.applyFilters(query, filters)

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, errors.Wrap(err, errors.CodeQueryFailed)
	}

	// 分页查询
	offset := req.GetOffset()
	pageSize := req.GetPageSize()

	if err := query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&entities).Error; err != nil {
		return nil, 0, errors.Wrap(err, errors.CodeQueryFailed)
	}

	return entities, total, nil
}

// Count 统计实体数量
func (r *baseRepository[T]) Count(ctx context.Context, filters map[string]interface{}) (int64, error) {
	var count int64
	query := r.db.WithContext(ctx).Model(new(T))

	// 应用过滤条件
	query = r.applyFilters(query, filters)

	if err := query.Count(&count).Error; err != nil {
		return 0, errors.Wrap(err, errors.CodeQueryFailed)
	}

	return count, nil
}

// Exists 检查实体是否存在
func (r *baseRepository[T]) Exists(ctx context.Context, id string) (bool, error) {
	// 验证UUID格式
	if !common.ValidateUUID(id) {
		return false, errors.New(errors.CodeInvalidParameter, "无效的ID格式")
	}

	var count int64
	err := r.db.WithContext(ctx).Model(new(T)).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, errors.Wrap(err, errors.CodeQueryFailed)
	}
	return count > 0, nil
}

// BatchCreate 批量创建实体
func (r *baseRepository[T]) BatchCreate(ctx context.Context, entities []*T) error {
	if len(entities) == 0 {
		return nil
	}

	if err := r.db.WithContext(ctx).CreateInBatches(entities, 100).Error; err != nil {
		return errors.Wrap(err, errors.CodeQueryFailed)
	}
	return nil
}

// BatchDelete 批量删除实体
func (r *baseRepository[T]) BatchDelete(ctx context.Context, ids []string) error {
	if len(ids) == 0 {
		return nil
	}

	// 验证所有UUID格式
	for _, id := range ids {
		if !common.ValidateUUID(id) {
			return errors.New(errors.CodeInvalidParameter, "包含无效的ID格式")
		}
	}

	var entity T
	result := r.db.WithContext(ctx).Where("id IN ?", ids).Delete(&entity)
	if result.Error != nil {
		return errors.Wrap(result.Error, errors.CodeQueryFailed)
	}

	return nil
}

// applyFilters 应用过滤条件
func (r *baseRepository[T]) applyFilters(query *gorm.DB, filters map[string]interface{}) *gorm.DB {
	for key, value := range filters {
		switch key {
		case "status":
			if status, ok := value.(int); ok {
				query = query.Where("status = ?", status)
			}
		case "keyword":
			if keyword, ok := value.(string); ok && keyword != "" {
				// 这里可以根据具体模型调整搜索字段
				query = query.Where("name LIKE ? OR description LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
			}
		case "created_after":
			if date, ok := value.(string); ok && date != "" {
				query = query.Where("created_at >= ?", date)
			}
		case "created_before":
			if date, ok := value.(string); ok && date != "" {
				query = query.Where("created_at <= ?", date)
			}
		}
	}
	return query
}

// Transaction 事务操作
func (r *baseRepository[T]) Transaction(ctx context.Context, fn func(*gorm.DB) error) error {
	return r.db.WithContext(ctx).Transaction(fn)
}

// GetDB 获取数据库连接
func (r *baseRepository[T]) GetDB() *gorm.DB {
	return r.db
}
