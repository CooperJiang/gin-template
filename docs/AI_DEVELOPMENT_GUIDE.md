# AI开发指南 - Gin Web项目模板

> 本文档专门为AI助手（如Cursor、Claude等）提供完整的项目架构和开发规范信息，确保AI能够准确理解项目结构并协助开发。

## 📋 项目概述

这是一个基于Gin框架的企业级Go Web项目模板，采用分层架构设计，支持快速开发和模块化扩展。

### 核心特性
- **分层架构**：Controller → Service → Repository → Model
- **UUID主键**：所有模型使用UUID作为主键，确保数据安全
- **自动降级**：数据库(MySQL→SQLite)、缓存(Redis→内存)自动降级
- **完整工具链**：40+个Makefile命令覆盖开发全流程
- **模块生成**：一键生成CRUD模块脚手架
- **统一规范**：错误处理、参数验证、API设计统一标准

## 🏗️ 项目架构

### 分层架构图
```
┌─────────────────┐
│   Controller    │ ← HTTP路由处理、参数验证、响应格式化
├─────────────────┤
│    Service      │ ← 业务逻辑处理、事务管理
├─────────────────┤
│   Repository    │ ← 数据访问层、CRUD操作
├─────────────────┤
│     Model       │ ← 数据模型定义
└─────────────────┘
```

### 数据流向
```
Request → Middleware → Controller → Service → Repository → Database
                                      ↓
Response ← Controller ← Service ← Repository ← Database
```

## 📁 目录结构规范

```
gin-template/
├── cmd/                          # 应用程序入口
│   └── main.go                   # 主程序入口
├── internal/                     # 私有应用代码
│   ├── controllers/              # 控制器层
│   │   ├── user/                 # 用户相关控制器
│   │   │   └── user_controller.go
│   │   └── base_controller.go    # 基础控制器
│   ├── services/                 # 服务层（业务逻辑）
│   │   ├── user/                 # 用户相关服务
│   │   │   └── user_service.go
│   │   └── base_service.go       # 基础服务
│   ├── repositories/             # 数据访问层
│   │   ├── user/                 # 用户相关仓库
│   │   │   └── user_repository.go
│   │   └── base_repository.go    # 基础仓库（泛型实现）
│   ├── models/                   # 数据模型
│   │   ├── base_model.go         # 基础模型（UUID主键）
│   │   └── user.go               # 用户模型
│   ├── dto/                      # 数据传输对象
│   │   ├── base_dto.go           # 基础DTO
│   │   ├── request/              # 请求DTO
│   │   │   └── user.go
│   │   └── response/             # 响应DTO
│   │       └── user.go
│   ├── middleware/               # 中间件
│   │   ├── auth.go               # 认证中间件
│   │   ├── cors.go               # 跨域中间件
│   │   └── logger.go             # 日志中间件
│   ├── routes/                   # 路由定义
│   │   ├── api.go                # API路由
│   │   └── routes.go             # 路由注册
│   ├── migrations/               # 数据库迁移文件
│   │   └── *.sql                 # SQL迁移文件
│   ├── cron/                     # 定时任务
│   └── static/                   # 静态文件
├── pkg/                          # 公共库代码
│   ├── common/                   # 通用工具
│   │   ├── uuid.go               # UUID工具类
│   │   ├── pagination.go         # 分页工具
│   │   ├── validate.go           # 验证工具
│   │   ├── response.go           # 响应工具
│   │   └── time.go               # 时间工具
│   ├── constants/                # 常量定义
│   │   ├── common.go             # 通用常量
│   │   └── user_status.go        # 用户状态常量
│   ├── config/                   # 配置管理
│   │   └── config.go
│   ├── database/                 # 数据库连接
│   │   └── database.go
│   ├── cache/                    # 缓存管理
│   │   └── cache.go
│   ├── logger/                   # 日志管理
│   │   └── logger.go
│   └── errors/                   # 错误处理
│       └── errors.go
├── tests/                        # 测试文件
│   ├── unit/                     # 单元测试
│   └── integration/              # 集成测试
├── scripts/                      # 脚本文件
│   ├── create_module.sh          # 模块生成脚本
│   ├── db_manager.sh             # 数据库管理脚本
│   └── performance_check.sh      # 性能检查脚本
├── docs/                         # 文档
│   ├── PROJECT_STANDARDS.md      # 项目开发规范
│   ├── QUICK_START.md            # 快速开始指南
│   ├── API_DOCS.md               # API文档
│   └── AI_DEVELOPMENT_GUIDE.md   # AI开发指南（本文档）
├── config.yaml                   # 配置文件
├── Makefile                      # 构建工具
├── Dockerfile                    # Docker配置
├── docker-compose.yml            # Docker Compose配置
├── go.mod                        # Go模块文件
└── README.md                     # 项目说明
```

## 🆔 UUID主键规范

### 重要原则
**所有数据模型必须使用UUID作为主键，禁止使用自增ID**

### UUID实现
```go
// 基础模型定义
type BaseModel struct {
    ID        common.UUID     `gorm:"type:char(36);primaryKey" json:"id"`
    CreatedAt common.JSONTime `json:"created_at"`
    UpdatedAt common.JSONTime `json:"updated_at"`
    DeletedAt gorm.DeletedAt  `gorm:"index" json:"-"`
}

// 自动生成UUID
func (m *BaseModel) BeforeCreate(tx *gorm.DB) error {
    if m.ID.IsZero() {
        m.ID = common.NewUUID()
    }
    return nil
}
```

### 数据库字段类型
- **MySQL**: `CHAR(36)`
- **PostgreSQL**: `UUID`
- **SQLite**: `TEXT`

### UUID优势
1. **数据安全**：用户无法通过ID推测数据量
2. **唯一性保证**：UUID v4几乎不可能重复
3. **分布式友好**：支持多实例部署
4. **防止枚举攻击**：ID不可预测

## 🏛️ 分层架构详解

### 1. Controller层（控制器层）
**职责**：HTTP请求处理、参数验证、响应格式化

```go
// 标准Controller结构
func CreateUser(c *gin.Context) {
    // 1. 参数验证
    req, err := common.ValidateRequest[request.CreateUserRequest](c)
    if err != nil {
        errors.HandleError(c, err)
        return
    }
    
    // 2. 调用Service层
    user, err := userService.CreateUser(req.Username, req.Email, req.Password)
    if err != nil {
        errors.HandleError(c, err)
        return
    }
    
    // 3. 构造响应
    resp := response.UserInfo{}.FromModel(*user)
    errors.ResponseSuccess(c, resp, "创建成功")
}
```

**规范**：
- 只处理HTTP相关逻辑
- 不包含业务逻辑
- 统一错误处理
- 统一响应格式

### 2. Service层（服务层）
**职责**：业务逻辑处理、事务管理、数据组装

```go
// 标准Service结构
func (s *userService) CreateUser(username, email, password string) (*models.User, error) {
    // 1. 业务逻辑验证
    if exists, _ := s.userRepo.ExistsByEmail(email); exists {
        return nil, errors.New(errors.CodeUserExists, "邮箱已存在")
    }
    
    // 2. 数据处理
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    
    // 3. 调用Repository层
    user := &models.User{
        Username: username,
        Email:    email,
        Password: string(hashedPassword),
    }
    
    if err := s.userRepo.Create(context.Background(), user); err != nil {
        return nil, err
    }
    
    return user, nil
}
```

**规范**：
- 包含所有业务逻辑
- 管理事务
- 调用Repository层
- 不直接操作数据库

### 3. Repository层（数据访问层）
**职责**：数据库CRUD操作、查询封装

```go
// 使用泛型基础Repository
type UserRepository interface {
    repositories.BaseRepository[models.User]
    GetByEmail(ctx context.Context, email string) (*models.User, error)
    ExistsByEmail(ctx context.Context, email string) (bool, error)
}

// 实现自定义方法
func (r *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
    var user models.User
    err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, errors.New(errors.CodeNotFound, "用户不存在")
        }
        return nil, errors.Wrap(err, errors.CodeQueryFailed)
    }
    return &user, nil
}
```

**规范**：
- 继承BaseRepository获得通用CRUD
- 只包含数据访问逻辑
- 统一错误处理
- 支持事务操作

### 4. Model层（数据模型层）
**职责**：数据结构定义、数据库映射

```go
// 标准Model结构
type User struct {
    BaseModel                                    // 嵌入基础模型（UUID主键）
    Username string `gorm:"size:50;not null;uniqueIndex" json:"username"`
    Password string `gorm:"size:100;not null" json:"-"`
    Email    string `gorm:"size:100;uniqueIndex" json:"email"`
    Status   int    `gorm:"default:1" json:"status"`
    Role     int    `gorm:"default:3" json:"role"`
}

// 表名定义
func (User) TableName() string {
    return "user"
}

// 创建前钩子
func (u *User) BeforeCreate(tx *gorm.DB) error {
    // 调用基础模型的BeforeCreate（生成UUID）
    if err := u.BaseModel.BeforeCreate(tx); err != nil {
        return err
    }
    
    // 设置默认值
    if u.Status == 0 {
        u.Status = constants.UserStatusNormal
    }
    return nil
}
```

**规范**：
- 必须嵌入BaseModel
- 使用UUID主键
- 定义表名
- 实现必要的钩子函数

## 📝 DTO规范

### Request DTO
```go
type CreateUserRequest struct {
    dto.BaseRequest
    Username string `json:"username" binding:"required,min=2,max=20"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6,max=20"`
}

// 验证消息
func (r *CreateUserRequest) GetValidationMessages() map[string]string {
    return map[string]string{
        "Username.required": "用户名不能为空",
        "Username.min":      "用户名长度不能小于2个字符",
        "Email.required":    "邮箱不能为空",
        "Email.email":       "请输入有效的邮箱地址",
        "Password.required": "密码不能为空",
        "Password.min":      "密码长度不能小于6个字符",
    }
}
```

### Response DTO
```go
type UserInfo struct {
    dto.BaseResponse                    // ID为string类型（UUID）
    Username string `json:"username"`
    Email    string `json:"email"`
    Status   int    `json:"status"`
}

// 从模型转换
func (u *UserInfo) FromModel(user models.User) *UserInfo {
    return &UserInfo{
        BaseResponse: dto.BaseResponse{
            ID:        user.ID.String(),    // UUID转字符串
            CreatedAt: time.Time(user.CreatedAt),
            UpdatedAt: time.Time(user.UpdatedAt),
        },
        Username: user.Username,
        Email:    user.Email,
        Status:   user.Status,
    }
}
```

## 🔧 开发工具链

### Makefile命令分类

#### 开发相关
- `make dev` - 启动热重载开发服务器
- `make build` - 构建应用程序
- `make run` - 构建并运行
- `make clean` - 清理构建文件

#### 测试相关
- `make test` - 运行所有测试
- `make test-coverage` - 生成测试覆盖率报告
- `make benchmark` - 运行性能基准测试

#### 代码质量
- `make fmt` - 格式化代码
- `make lint` - 代码检查
- `make security` - 安全检查
- `make full-check` - 完整代码检查

#### 数据库相关
- `make migrate` - 执行数据库迁移
- `make migration name=xxx` - 创建新迁移文件
- `make db-reset` - 重置数据库

#### 模块生成
- `make new-module name=product` - 创建新模块

#### 快速启动
- `make quick-start` - 初始化+迁移+开发服务器
- `make setup` - 初始化开发环境

## 🚀 模块生成规范

### 使用方法
```bash
make new-module name=product
```

### 生成的文件结构
```
internal/
├── models/product.go              # 数据模型
├── repositories/product/          # 数据访问层
│   └── product_repository.go
├── services/product/              # 业务逻辑层
│   └── product_service.go
├── controllers/product/           # 控制器层
│   └── product_controller.go
├── dto/request/product.go         # 请求DTO
├── dto/response/product.go        # 响应DTO
└── migrations/xxx_create_product_table.sql  # 数据库迁移
```

### 生成的代码特点
- 自动使用UUID主键
- 完整的CRUD操作
- 统一的错误处理
- 标准的分层架构
- 参数验证和响应格式

## 📊 数据库规范

### 表设计规范
1. **主键**：必须使用UUID，字段名为`id`
2. **时间字段**：`created_at`、`updated_at`、`deleted_at`
3. **软删除**：使用`deleted_at`字段
4. **索引**：为常用查询字段添加索引
5. **字符集**：UTF-8

### 迁移文件规范
```sql
-- +migrate Up
-- 创建产品表
CREATE TABLE product (
    id CHAR(36) PRIMARY KEY,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    status INT DEFAULT 1,
    KEY idx_deleted_at (deleted_at),
    KEY idx_status (status)
);

-- +migrate Down
DROP TABLE IF EXISTS product;
```

## 🔒 错误处理规范

### 错误码定义
```go
const (
    CodeSuccess           ErrorCode = 0   // 成功
    CodeInvalidParameter  ErrorCode = 100 // 无效参数
    CodeUnauthorized      ErrorCode = 101 // 未授权
    CodeNotFound          ErrorCode = 102 // 资源不存在
    CodeUserExists        ErrorCode = 103 // 用户已存在
    CodeQueryFailed       ErrorCode = 200 // 查询失败
)
```

### 错误处理方式
```go
// 创建错误
err := errors.New(errors.CodeUserExists, "用户已存在")

// 包装错误
err := errors.Wrap(dbErr, errors.CodeQueryFailed)

// 处理错误
errors.HandleError(c, err)

// 成功响应
errors.ResponseSuccess(c, data, "操作成功")
```

## 🔐 认证授权规范

### JWT Token结构
```go
type Claims struct {
    UserID   string `json:"user_id"`    // UUID字符串
    Username string `json:"username"`
    Role     int    `json:"role"`
    jwt.StandardClaims
}
```

### 中间件使用
```go
// 需要认证的路由
auth := r.Group("/api/v1")
auth.Use(middleware.AuthMiddleware())
{
    auth.GET("/profile", user.GetProfile)
    auth.PUT("/profile", user.UpdateProfile)
}
```

## 📝 API设计规范

### RESTful API规范
- `GET /api/v1/users` - 获取用户列表
- `GET /api/v1/users/:id` - 获取单个用户（ID为UUID）
- `POST /api/v1/users` - 创建用户
- `PUT /api/v1/users/:id` - 更新用户
- `DELETE /api/v1/users/:id` - 删除用户

### 统一响应格式
```json
{
    "code": 0,
    "message": "success",
    "data": {
        "id": "550e8400-e29b-41d4-a716-446655440000",
        "username": "john_doe",
        "email": "john@example.com",
        "created_at": "2024-01-01T00:00:00Z"
    }
}
```

## 🧪 测试规范

### 单元测试
```go
func TestUserService_CreateUser(t *testing.T) {
    // 使用testify框架
    assert := assert.New(t)
    
    // Mock Repository
    mockRepo := &mocks.UserRepository{}
    service := NewUserService(mockRepo)
    
    // 测试用例
    user, err := service.CreateUser("test", "test@example.com", "password")
    
    assert.NoError(err)
    assert.NotNil(user)
    assert.True(common.ValidateUUID(user.ID.String()))
}
```

### 集成测试
```go
func TestUserAPI_CreateUser(t *testing.T) {
    // 设置测试环境
    router := setupTestRouter()
    
    // 构造请求
    body := `{"username":"test","email":"test@example.com","password":"password"}`
    req := httptest.NewRequest("POST", "/api/v1/users", strings.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    
    // 执行请求
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    
    // 验证响应
    assert.Equal(t, 200, w.Code)
}
```

## 🔄 开发流程

### 1. 创建新模块
```bash
# 生成模块脚手架
make new-module name=product

# 执行数据库迁移
make migrate

# 启动开发服务器
make dev
```

### 2. 开发步骤
1. **定义Model**：在生成的模型基础上调整字段
2. **实现Repository**：添加自定义查询方法
3. **实现Service**：编写业务逻辑
4. **实现Controller**：处理HTTP请求
5. **定义DTO**：请求和响应对象
6. **编写测试**：单元测试和集成测试

### 3. 代码检查
```bash
# 完整检查
make full-check

# 单独检查
make fmt lint security test
```

## 🚨 重要注意事项

### 必须遵守的规则
1. **UUID主键**：所有模型必须使用UUID，禁止自增ID
2. **分层架构**：严格按照Controller→Service→Repository→Model分层
3. **错误处理**：使用统一的错误处理机制
4. **参数验证**：所有输入参数必须验证
5. **事务管理**：在Service层管理事务
6. **测试覆盖**：新功能必须有测试覆盖

### 禁止的操作
1. ❌ 在Controller中写业务逻辑
2. ❌ 在Service中直接操作HTTP请求/响应
3. ❌ 使用自增ID作为主键
4. ❌ 跳过参数验证
5. ❌ 在Repository中写业务逻辑
6. ❌ 硬编码配置信息

### 推荐的做法
1. ✅ 使用BaseModel嵌入获得UUID主键
2. ✅ 使用BaseRepository获得通用CRUD
3. ✅ 实现GetValidationMessages方法
4. ✅ 使用FromModel方法转换响应
5. ✅ 编写单元测试和集成测试
6. ✅ 使用make命令进行开发

## 📚 相关文档

- [项目开发规范](PROJECT_STANDARDS.md) - 详细的开发规范
- [快速开始指南](QUICK_START.md) - 5分钟上手指南
- [API文档](API_DOCS.md) - API接口文档

## 🤖 AI助手使用建议

当AI助手协助开发时，请：

1. **严格遵循UUID主键规范**
2. **按照分层架构组织代码**
3. **使用模块生成脚手架**
4. **参考现有代码风格**
5. **运行代码检查命令**
6. **编写相应的测试**

这份文档确保AI助手能够准确理解项目架构，协助开发出符合规范的高质量代码。 