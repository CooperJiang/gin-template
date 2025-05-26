# 🚀 快速开始指南

## 📋 目录

- [环境准备](#环境准备)
- [5分钟快速上手](#5分钟快速上手)
- [创建第一个模块](#创建第一个模块)
- [常用命令](#常用命令)
- [开发最佳实践](#开发最佳实践)
- [高级功能](#高级功能)
- [部署指南](#部署指南)
- [常见问题](#常见问题)

## 🛠️ 环境准备

### 系统要求
- Go 1.21+
- Git
- Make (可选，但推荐)

### 可选依赖
- MySQL 8.0+ (不配置则自动使用SQLite)
- Redis 6.0+ (不配置则自动使用内存缓存)
- Docker & Docker Compose (用于容器化部署)

## ⚡ 5分钟快速上手

### 第一步：克隆项目
```bash
git clone <your-repo-url>
cd gin-template
```

### 第二步：一键初始化
```bash
# 安装开发工具并初始化配置
make setup
```

这个命令会：
- 安装所有必需的开发工具
- 创建配置文件 `config.yaml`
- 下载Go依赖

### 第三步：配置数据库（可选）
编辑 `config.yaml` 文件：

```yaml
# 使用MySQL（可选）
database:
  host: "localhost"
  port: 3306
  username: "root"
  password: "password"
  name: "gin_template"

# 使用Redis（可选）
redis:
  host: "localhost"
  port: 6379
  password: ""
  db: 0
```

> 💡 **提示**: 如果不配置MySQL和Redis，系统会自动使用SQLite和内存缓存

### 第四步：启动开发服务器
```bash
# 启动热重载开发服务器
make dev
```

🎉 **恭喜！** 你的应用现在运行在 http://localhost:8080

### 第五步：测试API
```bash
# 用户注册
curl -X POST http://localhost:8080/api/v1/user/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123",
    "email": "test@example.com"
  }'
```

## 🏗️ 创建第一个模块

### 使用脚手架生成模块
```bash
# 生成完整的CRUD模块
make new-module name=product
```

这个命令会自动生成：
- ✅ Model: `internal/models/product.go`
- ✅ Repository: `internal/repositories/product_repository.go`
- ✅ Service: `internal/services/product_service.go`
- ✅ Controller: `internal/controllers/product_controller.go`
- ✅ DTO: `internal/dto/request/product.go` 和 `internal/dto/response/product.go`
- ✅ 路由: 自动注册到路由系统
- ✅ 迁移文件: `migrations/xxx_create_products_table.sql`
- ✅ 测试文件: 相应的测试模板

### 运行数据库迁移
```bash
make migrate
```

### 重启开发服务器
```bash
make dev
```

### 测试新模块
```bash
# 创建产品
curl -X POST http://localhost:8080/api/v1/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "iPhone 15",
    "price": 999.99,
    "description": "Latest iPhone"
  }'

# 获取产品列表
curl http://localhost:8080/api/v1/products
```

## 🔧 常用命令

### 开发相关
```bash
make dev          # 启动开发服务器（热重载）
make build        # 构建应用
make run          # 构建并运行
make clean        # 清理构建文件
```

### 代码质量
```bash
make fmt          # 格式化代码
make lint         # 代码检查
make security     # 安全检查
make full-check   # 完整检查（格式化+检查+安全+测试）
```

### 测试相关
```bash
make test         # 运行所有测试
make test-unit    # 运行单元测试
make test-integration # 运行集成测试
make test-coverage    # 生成测试覆盖率报告
make api-test     # 运行API测试
```

### 数据库相关
```bash
make migrate      # 运行数据库迁移
make migrate-down # 回滚数据库迁移
make migration name=add_user_avatar # 创建新迁移
make db-reset     # 重置数据库
make db-backup    # 备份数据库
make db-restore   # 恢复数据库
```

### 模块生成
```bash
make new-module name=order    # 生成订单模块
make new-module name=category # 生成分类模块
```

### Docker相关
```bash
make docker-build # 构建Docker镜像
make docker-run   # 运行Docker容器
```

## 💡 开发最佳实践

### 1. 项目结构
遵循四层架构模式：
```
Controller → Service → Repository → Model
```

### 2. 错误处理
使用统一的错误处理机制：
```go
// 创建业务错误
err := errors.New(errors.CodeEmailExists, "邮箱已存在")

// 包装系统错误
err = errors.Wrap(dbErr, errors.CodeQueryFailed)

// 在控制器中处理
if err != nil {
    errors.HandleError(c, err)
    return
}

// 成功响应
errors.ResponseSuccess(c, data, "操作成功")
```

### 3. 参数验证
使用DTO进行参数验证：
```go
type CreateUserRequest struct {
    Username string `json:"username" binding:"required,min=3,max=50"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}
```

### 4. 数据库操作
使用Repository模式：
```go
// 在Service中调用Repository
user, err := s.userRepo.GetByID(ctx, id)
if err != nil {
    return nil, errors.Wrap(err, errors.CodeQueryFailed)
}
```

### 5. 缓存使用
遵循缓存命名规范：
```go
cacheKey := fmt.Sprintf("%s%d", constants.CacheKeyUser, userID)
```

## 🚀 高级功能

### 1. 性能监控
```bash
# 性能检查
make perf-check

# 性能分析
make profile
```

### 2. API文档
```bash
# 生成Swagger文档
make swagger

# 访问文档: http://localhost:8080/swagger/index.html
```

### 3. 数据库管理
```bash
# 数据库信息
make db-info

# 数据库优化
make db-optimize

# 执行SQL
make db-sql file=query.sql
```

## 🐳 部署指南

### 开发环境部署
```bash
# 使用Docker Compose启动完整环境
docker-compose up -d

# 查看日志
docker-compose logs -f app
```

### 生产环境部署
```bash
# 构建生产镜像
make docker-build

# 部署到生产环境
make deploy-prod
```

### 环境变量配置
```bash
# 设置环境变量
export APP_APP_PORT=8080
export APP_DB_HOST=localhost
export APP_REDIS_HOST=localhost
```

## ❓ 常见问题

### Q: 如何切换数据库？
A: 编辑 `config.yaml` 文件中的数据库配置。如果不配置MySQL，系统会自动使用SQLite。

### Q: 如何添加新的API接口？
A: 使用 `make new-module name=模块名` 生成完整模块，或手动在对应的Controller中添加方法。

### Q: 如何自定义错误码？
A: 在 `pkg/errors/codes.go` 中添加新的错误码定义。

### Q: 如何添加中间件？
A: 在 `internal/middleware/` 目录下创建中间件文件，然后在路由中使用。

### Q: 如何配置日志级别？
A: 在 `config.yaml` 中设置 `logger.level` 字段。

### Q: 如何进行单元测试？
A: 运行 `make test-unit` 或 `go test ./tests/unit/...`

### Q: 如何查看API文档？
A: 运行 `make swagger` 生成文档，然后访问 http://localhost:8080/swagger/index.html

### Q: 如何备份数据库？
A: 运行 `make db-backup` 进行备份，`make db-restore` 进行恢复。

### Q: 如何更新依赖？
A: 运行 `make update-deps` 更新所有依赖。

### Q: 如何查看应用日志？
A: 运行 `make logs` 查看实时日志。

## 📞 获取帮助

- 查看所有可用命令: `make help`
- 查看项目规范: [PROJECT_STANDARDS.md](PROJECT_STANDARDS.md)

## 🎯 下一步

1. 阅读 [项目开发规范](PROJECT_STANDARDS.md)
2. 创建你的第一个业务模块
3. 编写单元测试
4. 部署到测试环境

**祝你开发愉快！** 🚀 