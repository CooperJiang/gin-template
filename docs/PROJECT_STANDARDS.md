# 项目开发规范

## 📋 目录

- [目录结构规范](#目录结构规范)
- [分层架构规范](#分层架构规范)
- [代码规范](#代码规范)
- [API设计规范](#api设计规范)
- [数据库规范](#数据库规范)
- [缓存使用规范](#缓存使用规范)
- [错误处理规范](#错误处理规范)
- [常量定义规范](#常量定义规范)
- [参数验证规范](#参数验证规范)
- [日志记录规范](#日志记录规范)
- [测试规范](#测试规范)
- [部署规范](#部署规范)

## 📁 目录结构规范

### 项目根目录结构

```
gin-template/
├── cmd/                   # 应用程序入口
│   └── main.go           # 主程序入口
├── internal/             # 内部包，不对外暴露
│   ├── controllers/      # 控制器层
│   ├── services/         # 业务逻辑层
│   ├── repositories/     # 数据访问层
│   ├── models/           # 数据模型
│   ├── dto/              # 数据传输对象
│   ├── middleware/       # 中间件
│   ├── routes/           # 路由定义
│   ├── cron/             # 定时任务
│   └── static/           # 静态资源
├── pkg/                  # 可重用的库代码
│   ├── cache/            # 缓存功能
│   ├── common/           # 通用工具
│   ├── config/           # 配置管理
│   ├── constants/        # 常量定义
│   ├── database/         # 数据库连接
│   ├── email/            # 邮件功能
│   ├── errors/           # 错误处理
│   ├── logger/           # 日志功能
│   └── utils/            # 工具函数
├── tests/                # 测试文件
├── scripts/              # 脚本文件
├── migrations/           # 数据库迁移文件
├── docs/                 # 文档
├── config.yaml           # 配置文件
├── Makefile             # 构建工具
├── Dockerfile           # Docker配置
└── docker-compose.yml   # Docker Compose配置
```

### 目录命名规范

1. **目录名使用小写字母和下划线**
   - ✅ `user_controller`
   - ❌ `UserController` 或 `userController`

2. **包名与目录名保持一致**
   ```go
   // internal/controllers/user_controller.go
   package controllers
   ```

3. **文件名使用小写字母和下划线**
   - ✅ `user_controller.go`
   - ✅ `user_service.go`
   - ❌ `UserController.go`

### 模块目录结构

每个业务模块应按以下结构组织：

```
internal/
├── controllers/
│   └── user_controller.go      # 用户控制器
├── services/
│   └── user_service.go         # 用户服务
├── repositories/
│   └── user_repository.go      # 用户数据访问
├── models/
│   └── user.go                 # 用户模型
└── dto/
    ├── request/
    │   └── user.go             # 用户请求DTO
    └── response/
        └── user.go             # 用户响应DTO
```

## 🏗️ 分层架构规范

### 四层架构模式

项目采用 **Controller → Service → Repository → Model** 四层架构：

```
┌─────────────────┐
│   Controller    │  ← HTTP请求处理、参数验证、响应格式化
├─────────────────┤
│    Service      │  ← 业务逻辑、事务管理、业务规则
├─────────────────┤
│   Repository    │  ← 数据访问、查询封装、缓存处理
├─────────────────┤
│     Model       │  ← 数据模型、数据库映射
└─────────────────┘
```

### 各层职责

#### 1. Controller层（控制器层）
**职责**：
- 处理HTTP请求和响应
- 参数验证和格式化
- 调用Service层处理业务逻辑
- 统一错误处理和响应格式

**规范**：
```go
// internal/controllers/user_controller.go
package controllers

import (
    "template/internal/dto/request"
    "template/internal/dto/response"
    "template/internal/services"
    "template/pkg/common"
    "template/pkg/errors"
    
    "github.com/gin-gonic/gin"
)

type UserController struct {
    userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
    return &UserController{
        userService: userService,
    }
}

// CreateUser 创建用户
func (uc *UserController) CreateUser(c *gin.Context) {
    // 1. 参数验证
    req, err := common.ValidateRequest[request.CreateUserRequest](c)
    if err != nil {
        errors.HandleError(c, err)
        return
    }
    
    // 2. 调用服务层
    user, err := uc.userService.CreateUser(c.Request.Context(), req)
    if err != nil {
        errors.HandleError(c, err)
        return
    }
    
    // 3. 返回响应
    resp := response.UserResponse{}.FromModel(user)
    errors.ResponseSuccess(c, resp, "用户创建成功")
}
```

#### 2. Service层（业务逻辑层）
**职责**：
- 实现业务逻辑和业务规则
- 事务管理
- 调用Repository层进行数据操作
- 业务数据验证

**规范**：
```go
// internal/services/user_service.go
package services

import (
    "context"
    "template/internal/dto/request"
    "template/internal/models"
    "template/internal/repositories"
    "template/pkg/errors"
)

type UserService interface {
    CreateUser(ctx context.Context, req *request.CreateUserRequest) (*models.User, error)
    GetUserByID(ctx context.Context, id uint) (*models.User, error)
    UpdateUser(ctx context.Context, id uint, req *request.UpdateUserRequest) error
    DeleteUser(ctx context.Context, id uint) error
}

type userService struct {
    userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
    return &userService{
        userRepo: userRepo,
    }
}

func (s *userService) CreateUser(ctx context.Context, req *request.CreateUserRequest) (*models.User, error) {
    // 1. 业务验证
    exists, err := s.userRepo.ExistsByEmail(ctx, req.Email)
    if err != nil {
        return nil, errors.Wrap(err, errors.CodeQueryFailed)
    }
    if exists {
        return nil, errors.New(errors.CodeEmailExists, "邮箱已存在")
    }
    
    // 2. 创建模型
    user := &models.User{
        Username: req.Username,
        Email:    req.Email,
        Password: req.Password, // 应该加密
    }
    
    // 3. 保存数据
    if err := s.userRepo.Create(ctx, user); err != nil {
        return nil, errors.Wrap(err, errors.CodeQueryFailed)
    }
    
    return user, nil
}
```

#### 3. Repository层（数据访问层）
**职责**：
- 数据库操作封装
- 查询逻辑实现
- 缓存处理
- 数据转换

**规范**：
```go
// internal/repositories/user_repository.go
package repositories

import (
    "context"
    "template/internal/models"
    "template/pkg/common"
    
    "gorm.io/gorm"
)

type UserRepository interface {
    BaseRepository[models.User]
    ExistsByEmail(ctx context.Context, email string) (bool, error)
    GetByEmail(ctx context.Context, email string) (*models.User, error)
    GetByUsername(ctx context.Context, username string) (*models.User, error)
}

type userRepository struct {
    BaseRepository[models.User]
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{
        BaseRepository: NewBaseRepository[models.User](db),
        db:             db,
    }
}

func (r *userRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
    var count int64
    err := r.db.WithContext(ctx).Model(&models.User{}).
        Where("email = ?", email).Count(&count).Error
    return count > 0, err
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
    var user models.User
    err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}
```

#### 4. Model层（数据模型层）
**职责**：
- 定义数据结构
- 数据库表映射
- 模型关联关系

**规范**：
```go
// internal/models/user.go
package models

import (
    "time"
    "template/pkg/constants"
)

type User struct {
    BaseModel
    Username string             `json:"username" gorm:"uniqueIndex;size:50;not null"`
    Email    string             `json:"email" gorm:"uniqueIndex;size:100;not null"`
    Password string             `json:"-" gorm:"size:255;not null"`
    Status   constants.UserStatus `json:"status" gorm:"default:1"`
    LastLoginAt *time.Time      `json:"last_login_at"`
}

// TableName 指定表名
func (User) TableName() string {
    return "users"
}

// BeforeCreate 创建前钩子
func (u *User) BeforeCreate(tx *gorm.DB) error {
    // 密码加密等逻辑
    return nil
}
```

### 依赖注入规范

使用构造函数注入模式：

```go
// cmd/main.go
func main() {
    // 初始化数据库
    db := database.GetDB()
    
    // 初始化Repository
    userRepo := repositories.NewUserRepository(db)
    
    // 初始化Service
    userService := services.NewUserService(userRepo)
    
    // 初始化Controller
    userController := controllers.NewUserController(userService)
    
    // 注册路由
    routes.RegisterUserRoutes(router, userController)
}
```

## 📝 代码规范

### 命名规范

#### 1. 变量命名
```go
// ✅ 正确
var userName string
var userID uint
var isActive bool

// ❌ 错误
var user_name string
var userid uint
var is_active bool
```

#### 2. 函数命名
```go
// ✅ 正确
func GetUserByID(id uint) (*User, error)
func CreateUser(user *User) error
func IsEmailExists(email string) bool

// ❌ 错误
func getUserById(id uint) (*User, error)
func create_user(user *User) error
```

#### 3. 常量命名
```go
// ✅ 正确
const (
    DefaultPageSize = 20
    MaxPageSize     = 100
    UserStatusActive = 1
)

// ❌ 错误
const (
    default_page_size = 20
    maxPageSize       = 100
)
```

### 注释规范

#### 1. 包注释
```go
// Package controllers 提供HTTP请求处理功能
// 包含用户、订单等业务模块的控制器实现
package controllers
```

#### 2. 函数注释
```go
// CreateUser 创建新用户
// 参数:
//   - ctx: 上下文
//   - req: 创建用户请求
// 返回:
//   - *models.User: 创建的用户信息
//   - error: 错误信息
func CreateUser(ctx context.Context, req *request.CreateUserRequest) (*models.User, error) {
    // 实现逻辑
}
```

#### 3. 结构体注释
```go
// User 用户模型
// 包含用户的基本信息和状态
type User struct {
    ID       uint   `json:"id"`       // 用户ID
    Username string `json:"username"` // 用户名
    Email    string `json:"email"`    // 邮箱地址
    Status   int    `json:"status"`   // 用户状态：1-正常，2-禁用
}
```

## 🌐 API设计规范

### RESTful API规范

#### 1. URL设计
```
GET    /api/v1/users          # 获取用户列表
GET    /api/v1/users/{id}     # 获取指定用户
POST   /api/v1/users          # 创建用户
PUT    /api/v1/users/{id}     # 更新用户
DELETE /api/v1/users/{id}     # 删除用户
```

#### 2. HTTP状态码
```
200 OK          # 请求成功
201 Created     # 创建成功
400 Bad Request # 请求参数错误
401 Unauthorized # 未授权
403 Forbidden   # 禁止访问
404 Not Found   # 资源不存在
500 Internal Server Error # 服务器错误
```

### 统一响应格式

#### 成功响应
```json
{
  "code": 200,
  "message": "操作成功",
  "data": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com"
  },
  "request_id": "550e8400-e29b-41d4-a716-446655440000",
  "timestamp": "2024-01-01T12:00:00Z"
}
```

#### 错误响应
```json
{
  "code": 1001,
  "message": "用户名已存在",
  "request_id": "550e8400-e29b-41d4-a716-446655440000",
  "timestamp": "2024-01-01T12:00:00Z"
}
```

#### 分页响应
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "items": [...],
    "pagination": {
      "page": 1,
      "page_size": 20,
      "total": 100,
      "total_pages": 5
    }
  },
  "request_id": "550e8400-e29b-41d4-a716-446655440000",
  "timestamp": "2024-01-01T12:00:00Z"
}
```

## 🗄️ 数据库规范

### 表设计规范

#### 1. 表命名
- 使用复数形式：`users`, `orders`, `products`
- 使用下划线分隔：`user_profiles`, `order_items`
- 避免使用保留字

#### 2. 字段命名
```sql
-- ✅ 正确
CREATE TABLE users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- ❌ 错误
CREATE TABLE user (
    ID BIGINT PRIMARY KEY AUTO_INCREMENT,
    UserName VARCHAR(50) NOT NULL,
    Email VARCHAR(100) NOT NULL UNIQUE
);
```

#### 3. 基础字段
每个表都应包含以下基础字段：
```sql
id BIGINT PRIMARY KEY AUTO_INCREMENT,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
deleted_at TIMESTAMP NULL  -- 软删除字段
```

### 索引规范

#### 1. 主键索引
```sql
-- 每个表必须有主键
PRIMARY KEY (id)
```

#### 2. 唯一索引
```sql
-- 唯一字段添加唯一索引
UNIQUE KEY uk_users_email (email),
UNIQUE KEY uk_users_username (username)
```

#### 3. 普通索引
```sql
-- 经常查询的字段添加索引
KEY idx_users_status (status),
KEY idx_users_created_at (created_at)
```

#### 4. 复合索引
```sql
-- 多字段查询添加复合索引
KEY idx_users_status_created (status, created_at)
```

### 迁移文件规范

#### 1. 文件命名
```
migrations/
├── 20240101120000_create_users_table.sql
├── 20240101120001_create_orders_table.sql
└── 20240101120002_add_status_to_users.sql
```

#### 2. 迁移文件内容
```sql
-- 20240101120000_create_users_table.sql
-- +migrate Up
CREATE TABLE users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(255) NOT NULL,
    status TINYINT DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    
    UNIQUE KEY uk_users_email (email),
    UNIQUE KEY uk_users_username (username),
    KEY idx_users_status (status)
);

-- +migrate Down
DROP TABLE IF EXISTS users;
```

## 🔧 常量定义规范

### 常量组织结构

```
pkg/constants/
├── common.go          # 通用常量
├── user_status.go     # 用户状态常量
├── order_status.go    # 订单状态常量
└── error_codes.go     # 错误码常量
```

### 常量定义规范

#### 1. 通用常量
```go
// pkg/constants/common.go
package constants

// 分页相关常量
const (
    DefaultPage     = 1   // 默认页码
    DefaultPageSize = 20  // 默认页大小
    MaxPageSize     = 100 // 最大页大小
)

// 时间格式常量
const (
    DateFormat     = "2006-01-02"
    TimeFormat     = "2006-01-02 15:04:05"
    DateTimeFormat = "2006-01-02T15:04:05Z07:00"
)
```

#### 2. 枚举类型常量
```go
// pkg/constants/user_status.go
package constants

// UserStatus 用户状态类型
type UserStatus int

// 用户状态枚举
const (
    UserStatusInactive UserStatus = iota // 0 - 未激活
    UserStatusActive                     // 1 - 正常
    UserStatusDisabled                   // 2 - 禁用
    UserStatusDeleted                    // 3 - 已删除
)

// String 返回状态的字符串表示
func (s UserStatus) String() string {
    switch s {
    case UserStatusInactive:
        return "未激活"
    case UserStatusActive:
        return "正常"
    case UserStatusDisabled:
        return "禁用"
    case UserStatusDeleted:
        return "已删除"
    default:
        return "未知"
    }
}

// IsValid 检查状态是否有效
func (s UserStatus) IsValid() bool {
    return s >= UserStatusInactive && s <= UserStatusDeleted
}
```

#### 3. 配置常量
```go
// pkg/constants/config.go
package constants

// 缓存相关常量
const (
    CacheKeyPrefix = "gin-template:"
    
    // 用户相关缓存键
    CacheKeyUser       = CacheKeyPrefix + "user:"
    CacheKeyUserToken  = CacheKeyPrefix + "user:token:"
    CacheKeyUserEmail  = CacheKeyPrefix + "user:email:"
    
    // 验证码相关缓存键
    CacheKeyVerifyCode = CacheKeyPrefix + "verify:"
)

// 缓存过期时间
const (
    CacheUserTTL       = 3600  // 用户信息缓存1小时
    CacheTokenTTL      = 86400 // Token缓存24小时
    CacheVerifyCodeTTL = 300   // 验证码缓存5分钟
)
```

## ✅ 参数验证规范

### DTO定义规范

#### 1. 请求DTO
```go
// internal/dto/request/user.go
package request

import "template/pkg/constants"

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
    Username string `json:"username" binding:"required,min=3,max=50" example:"testuser"`
    Email    string `json:"email" binding:"required,email,max=100" example:"test@example.com"`
    Password string `json:"password" binding:"required,min=6,max=50" example:"password123"`
}

// GetValidationMessages 获取验证错误信息
func (r *CreateUserRequest) GetValidationMessages() map[string]string {
    return map[string]string{
        "Username.required": "用户名不能为空",
        "Username.min":      "用户名长度不能少于3个字符",
        "Username.max":      "用户名长度不能超过50个字符",
        "Email.required":    "邮箱不能为空",
        "Email.email":       "邮箱格式不正确",
        "Email.max":         "邮箱长度不能超过100个字符",
        "Password.required": "密码不能为空",
        "Password.min":      "密码长度不能少于6个字符",
        "Password.max":      "密码长度不能超过50个字符",
    }
}

// UpdateUserRequest 更新用户请求
type UpdateUserRequest struct {
    Username *string               `json:"username,omitempty" binding:"omitempty,min=3,max=50"`
    Status   *constants.UserStatus `json:"status,omitempty" binding:"omitempty,oneof=0 1 2 3"`
}

// GetValidationMessages 获取验证错误信息
func (r *UpdateUserRequest) GetValidationMessages() map[string]string {
    return map[string]string{
        "Username.min":    "用户名长度不能少于3个字符",
        "Username.max":    "用户名长度不能超过50个字符",
        "Status.oneof":    "状态值必须是0、1、2或3",
    }
}
```

#### 2. 响应DTO
```go
// internal/dto/response/user.go
package response

import (
    "time"
    "template/internal/models"
    "template/pkg/constants"
)

// UserResponse 用户响应
type UserResponse struct {
    ID          uint                  `json:"id"`
    Username    string                `json:"username"`
    Email       string                `json:"email"`
    Status      constants.UserStatus  `json:"status"`
    StatusText  string                `json:"status_text"`
    CreatedAt   time.Time             `json:"created_at"`
    UpdatedAt   time.Time             `json:"updated_at"`
}

// FromModel 从模型转换
func (r UserResponse) FromModel(user *models.User) UserResponse {
    return UserResponse{
        ID:         user.ID,
        Username:   user.Username,
        Email:      user.Email,
        Status:     user.Status,
        StatusText: user.Status.String(),
        CreatedAt:  user.CreatedAt,
        UpdatedAt:  user.UpdatedAt,
    }
}

// UserListResponse 用户列表响应
type UserListResponse struct {
    Items      []UserResponse           `json:"items"`
    Pagination common.PaginationInfo    `json:"pagination"`
}
```

### 验证标签规范

#### 1. 基础验证
```go
type ExampleRequest struct {
    // 必填字段
    Name string `binding:"required"`
    
    // 长度限制
    Title string `binding:"min=1,max=100"`
    
    // 数值范围
    Age int `binding:"min=0,max=150"`
    
    // 邮箱格式
    Email string `binding:"email"`
    
    // URL格式
    Website string `binding:"url"`
    
    // 枚举值
    Status int `binding:"oneof=0 1 2"`
}
```

#### 2. 自定义验证
```go
// 注册自定义验证器
func init() {
    if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        v.RegisterValidation("username", validateUsername)
    }
}

// 用户名验证器
func validateUsername(fl validator.FieldLevel) bool {
    username := fl.Field().String()
    // 只允许字母、数字和下划线
    matched, _ := regexp.MatchString("^[a-zA-Z0-9_]+$", username)
    return matched
}

// 使用自定义验证
type UserRequest struct {
    Username string `binding:"required,username,min=3,max=20"`
}
```

### 验证错误处理

```go
// pkg/common/validator.go
package common

import (
    "reflect"
    "strings"
    "template/pkg/errors"
    
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
)

// ValidateRequest 验证请求参数
func ValidateRequest[T any](c *gin.Context) (*T, error) {
    var req T
    
    // 绑定参数
    if err := c.ShouldBindJSON(&req); err != nil {
        return nil, handleValidationError(err, req)
    }
    
    return &req, nil
}

// handleValidationError 处理验证错误
func handleValidationError(err error, req interface{}) error {
    if validationErrors, ok := err.(validator.ValidationErrors); ok {
        // 获取自定义错误消息
        messages := getValidationMessages(req)
        
        for _, fieldError := range validationErrors {
            field := fieldError.Field()
            tag := fieldError.Tag()
            key := field + "." + tag
            
            if message, exists := messages[key]; exists {
                return errors.New(errors.CodeValidationFailed, message)
            }
        }
    }
    
    return errors.New(errors.CodeValidationFailed, "参数验证失败")
}
```

## 📊 缓存使用规范

### 缓存键命名规范

```go
// pkg/constants/cache.go
package constants

// 缓存键前缀
const (
    CacheKeyPrefix = "gin-template:"
    
    // 用户相关缓存键
    CacheKeyUser       = CacheKeyPrefix + "user:"
    CacheKeyUserToken  = CacheKeyPrefix + "user:token:"
    CacheKeyUserEmail  = CacheKeyPrefix + "user:email:"
    
    // 验证码相关缓存键
    CacheKeyVerifyCode = CacheKeyPrefix + "verify:"
)

// 缓存过期时间
const (
    CacheUserTTL       = 3600  // 用户信息缓存1小时
    CacheTokenTTL      = 86400 // Token缓存24小时
    CacheVerifyCodeTTL = 300   // 验证码缓存5分钟
)
```

### 缓存使用模式

```go
// internal/services/user_service.go
func (s *userService) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
    // 1. 尝试从缓存获取
    cacheKey := fmt.Sprintf("%s%d", constants.CacheKeyUser, id)
    if cached, err := s.cache.Get(ctx, cacheKey); err == nil {
        var user models.User
        if err := json.Unmarshal([]byte(cached), &user); err == nil {
            return &user, nil
        }
    }
    
    // 2. 从数据库获取
    user, err := s.userRepo.GetByID(ctx, id)
    if err != nil {
        return nil, err
    }
    
    // 3. 写入缓存
    if data, err := json.Marshal(user); err == nil {
        s.cache.Set(ctx, cacheKey, string(data), constants.CacheUserTTL)
    }
    
    return user, nil
}
```

## 🚨 错误处理规范

### 错误码定义

```go
// pkg/errors/codes.go
package errors

// ErrorCode 错误码类型
type ErrorCode int

// 错误码定义
const (
    // 通用错误码（1-999）
    CodeUnknown          ErrorCode = 1
    CodeInternal         ErrorCode = 2
    CodeInvalidParameter ErrorCode = 100
    CodeUnauthorized     ErrorCode = 101
    CodeForbidden        ErrorCode = 102
    CodeNotFound         ErrorCode = 103
    CodeValidationFailed ErrorCode = 108
    
    // 用户相关错误码（1000-1999）
    CodeUserNotFound     ErrorCode = 1000
    CodeWrongPassword    ErrorCode = 1001
    CodeUserExists       ErrorCode = 1003
    CodeEmailExists      ErrorCode = 1007
    
    // 数据库相关错误码（2000-2999）
    CodeQueryFailed      ErrorCode = 2001
    CodeDBNoRecord       ErrorCode = 2002
)
```

### 错误处理模式

```go
// 在Service层创建业务错误
func (s *userService) CreateUser(ctx context.Context, req *request.CreateUserRequest) (*models.User, error) {
    // 检查邮箱是否存在
    exists, err := s.userRepo.ExistsByEmail(ctx, req.Email)
    if err != nil {
        // 包装数据库错误
        return nil, errors.Wrap(err, errors.CodeQueryFailed)
    }
    if exists {
        // 创建业务错误
        return nil, errors.New(errors.CodeEmailExists, "邮箱已存在")
    }
    
    // ... 其他逻辑
}

// 在Controller层处理错误
func (uc *UserController) CreateUser(c *gin.Context) {
    user, err := uc.userService.CreateUser(c.Request.Context(), req)
    if err != nil {
        // 统一错误处理
        errors.HandleError(c, err)
        return
    }
    
    // 成功响应
    errors.ResponseSuccess(c, response.UserResponse{}.FromModel(user), "用户创建成功")
}
```

## 📝 日志记录规范

### 日志级别使用

```go
// 调试信息
logger.Debug("用户查询参数: %+v", req)

// 一般信息
logger.Info("用户创建成功: ID=%d, Username=%s", user.ID, user.Username)

// 警告信息
logger.Warn("用户登录失败: Username=%s, IP=%s", username, ip)

// 错误信息
logger.Error("数据库连接失败: %v", err)

// 致命错误
logger.Fatal("应用启动失败: %v", err)
```

### 结构化日志

```go
// 使用结构化字段
logger.WithFields(logger.Fields{
    "user_id":   userID,
    "action":    "login",
    "ip":        clientIP,
    "user_agent": userAgent,
}).Info("用户登录")

// 错误日志包含堆栈信息
logger.WithFields(logger.Fields{
    "error":     err.Error(),
    "user_id":   userID,
    "request_id": requestID,
}).Error("用户操作失败")
```

## 🧪 测试规范

### 测试文件组织

```
tests/
├── unit/                    # 单元测试
│   ├── services/
│   │   └── user_service_test.go
│   ├── repositories/
│   │   └── user_repository_test.go
│   └── common_test.go
├── integration/             # 集成测试
│   ├── api_test.go
│   └── database_test.go
└── api_test.sh             # API自动化测试脚本
```

### 单元测试规范

```go
// tests/unit/services/user_service_test.go
package services

import (
    "context"
    "testing"
    "template/internal/dto/request"
    "template/internal/models"
    "template/internal/services"
    "template/pkg/errors"
    
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

// MockUserRepository 模拟用户仓库
type MockUserRepository struct {
    mock.Mock
}

func (m *MockUserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
    args := m.Called(ctx, email)
    return args.Bool(0), args.Error(1)
}

func TestUserService_CreateUser(t *testing.T) {
    tests := []struct {
        name    string
        req     *request.CreateUserRequest
        setup   func(*MockUserRepository)
        wantErr bool
        errCode errors.ErrorCode
    }{
        {
            name: "成功创建用户",
            req: &request.CreateUserRequest{
                Username: "testuser",
                Email:    "test@example.com",
                Password: "password123",
            },
            setup: func(repo *MockUserRepository) {
                repo.On("ExistsByEmail", mock.Anything, "test@example.com").Return(false, nil)
                repo.On("Create", mock.Anything, mock.AnythingOfType("*models.User")).Return(nil)
            },
            wantErr: false,
        },
        {
            name: "邮箱已存在",
            req: &request.CreateUserRequest{
                Username: "testuser",
                Email:    "test@example.com",
                Password: "password123",
            },
            setup: func(repo *MockUserRepository) {
                repo.On("ExistsByEmail", mock.Anything, "test@example.com").Return(true, nil)
            },
            wantErr: true,
            errCode: errors.CodeEmailExists,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // 准备
            mockRepo := new(MockUserRepository)
            tt.setup(mockRepo)
            
            service := services.NewUserService(mockRepo)
            
            // 执行
            user, err := service.CreateUser(context.Background(), tt.req)
            
            // 验证
            if tt.wantErr {
                assert.Error(t, err)
                if tt.errCode != 0 {
                    if appErr, ok := err.(*errors.Error); ok {
                        assert.Equal(t, tt.errCode, appErr.Code)
                    }
                }
                assert.Nil(t, user)
            } else {
                assert.NoError(t, err)
                assert.NotNil(t, user)
                assert.Equal(t, tt.req.Username, user.Username)
            }
            
            // 验证Mock调用
            mockRepo.AssertExpectations(t)
        })
    }
}
```

### 集成测试规范

```go
// tests/integration/api_test.go
func (suite *APITestSuite) TestCreateUser() {
    // 准备测试数据
    reqData := map[string]string{
        "username": "testuser",
        "email":    "test@example.com",
        "password": "password123",
    }
    
    // 发送请求
    resp := suite.makeRequest("POST", "/api/v1/users", reqData, nil)
    
    // 验证响应
    suite.assertResponse(resp, 200, 200)
    
    // 验证数据库
    var user models.User
    err := suite.db.Where("email = ?", "test@example.com").First(&user).Error
    assert.NoError(suite.T(), err)
    assert.Equal(suite.T(), "testuser", user.Username)
}
```

## 🚀 部署规范

### 环境配置

#### 1. 开发环境
```yaml
# config.dev.yaml
app:
  mode: debug
  port: 8080

database:
  host: localhost
  port: 3306
  name: gin_template_dev

redis:
  host: localhost
  port: 6379
  db: 0
```

#### 2. 测试环境
```yaml
# config.test.yaml
app:
  mode: release
  port: 8080

database:
  host: test-db.example.com
  port: 3306
  name: gin_template_test

redis:
  host: test-redis.example.com
  port: 6379
  db: 0
```

#### 3. 生产环境
```yaml
# config.prod.yaml
app:
  mode: release
  port: 8080

database:
  host: prod-db.example.com
  port: 3306
  name: gin_template_prod

redis:
  host: prod-redis.example.com
  port: 6379
  db: 0
```

### Docker配置

#### 1. 开发环境Docker Compose
```yaml
# docker-compose.dev.yml
version: '3.8'
services:
  app:
    build:
      context: .
      dockerfile: Dockerfile.dev
    ports:
      - "8080:8080"
    volumes:
      - .:/app
    environment:
      - APP_APP_MODE=debug
    depends_on:
      - mysql
      - redis

  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: password
      MYSQL_DATABASE: gin_template_dev
    ports:
      - "3306:3306"

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
```

#### 2. 生产环境Dockerfile
```dockerfile
# 多阶段构建
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/config.yaml .

EXPOSE 8080
CMD ["./main"]
```

---

## 📋 总结

本规范文档涵盖了项目开发的各个方面，包括：

1. **目录结构规范** - 统一的项目组织方式
2. **分层架构规范** - 清晰的四层架构模式
3. **代码规范** - 命名、注释等编码规范
4. **API设计规范** - RESTful API和响应格式规范
5. **数据库规范** - 表设计、索引、迁移规范
6. **常量定义规范** - 统一的常量管理方式
7. **参数验证规范** - DTO设计和验证规则
8. **缓存使用规范** - 缓存键命名和使用模式
9. **错误处理规范** - 统一的错误处理机制
10. **日志记录规范** - 结构化日志记录方式
11. **测试规范** - 单元测试和集成测试规范
12. **部署规范** - 多环境部署配置

遵循这些规范可以确保项目的：
- **一致性** - 团队成员编写的代码风格统一
- **可维护性** - 清晰的架构和规范便于维护
- **可扩展性** - 模块化设计便于功能扩展
- **可测试性** - 完善的测试规范保证代码质量
- **可部署性** - 标准化的部署流程提高效率

**请所有开发人员严格遵循本规范进行开发！**

## 📞 联系方式

如有问题或建议，请联系项目维护者或提交Issue。

**最后更新时间**：2024-01-01
**文档版本**：v1.0.0 