# 🚀 Gin Web 项目模板

一个功能完整、开箱即用的Go Web应用模板，基于Gin框架构建，集成了企业级开发所需的各种功能和最佳实践。

## ⚡ 快速开始

```bash
# 1. 克隆项目
git clone <your-repo-url>
cd gin-template

# 2. 一键初始化开发环境
make setup

# 3. 启动开发服务器（热重载）
make dev
```

🎉 **就这么简单！** 你的应用现在运行在 http://localhost:8080

### 🎯 5分钟创建你的第一个模块

```bash
# 自动生成完整的CRUD模块（包含Controller、Service、Repository、DTO、测试）
make new-module name=product

# 运行数据库迁移
make migrate

# 重启开发服务器
make dev
```

## 📚 文档

| 文档 | 描述 |
|------|------|
| 📋 **[项目开发规范](docs/PROJECT_STANDARDS.md)** | 完整的项目开发规范和最佳实践 |

## ⚡ 核心特性

### 🏗️ 开发效率
- **一键模块生成**: `make new-module name=product` 自动生成完整CRUD
- **热重载开发**: `make dev` 代码变更自动重启
- **完整脚手架**: Controller、Service、Repository、DTO、测试一键生成
- **智能工具链**: 集成代码检查、测试、构建、部署等完整工具链

### 🛡️ 企业级特性
- **分层架构**: Controller → Service → Repository → Model
- **统一错误处理**: 标准化错误码和响应格式
- **参数验证**: 自动参数验证和错误提示
- **缓存支持**: Redis/内存缓存自动降级
- **数据库支持**: MySQL/SQLite自动降级
- **用户认证**: JWT + 角色权限控制

### 🔧 开发工具
- **代码质量**: 集成 golangci-lint、gosec 等工具
- **自动化测试**: 单元测试和集成测试模板
- **API测试**: 完整的API自动化测试脚本
- **性能分析**: 内置性能检查和优化工具
- **Docker支持**: 多阶段构建和容器化部署

## 🛠️ 技术栈

| 类别 | 技术 | 说明 |
|------|------|------|
| **Web框架** | Gin | 高性能HTTP Web框架 |
| **数据库** | MySQL/SQLite | 支持自动降级 |
| **缓存** | Redis/内存 | 支持自动降级 |
| **ORM** | GORM | 功能强大的ORM库 |
| **认证** | JWT | JSON Web Token |
| **日志** | Logrus | 结构化日志 |
| **配置** | Viper | 配置管理 |
| **测试** | Testify | 测试框架 |
| **部署** | Docker | 容器化部署 |

## 📁 项目结构

```
gin-template/
├── cmd/                   # 应用程序入口
│   └── main.go
├── internal/              # 内部包，不对外暴露
│   ├── controllers/       # 控制器层
│   ├── services/          # 业务逻辑层
│   ├── repositories/      # 数据访问层
│   ├── models/            # 数据模型
│   ├── dto/               # 数据传输对象
│   │   ├── request/       # 请求DTO
│   │   └── response/      # 响应DTO
│   ├── middleware/        # 中间件
│   ├── routes/            # 路由定义
│   ├── cron/              # 定时任务
│   └── static/            # 静态资源
├── pkg/                   # 可重用的库代码
│   ├── cache/             # 缓存功能
│   ├── common/            # 通用工具
│   ├── config/            # 配置管理
│   ├── constants/         # 常量定义
│   ├── database/          # 数据库连接
│   ├── email/             # 邮件功能
│   ├── errors/            # 错误处理
│   ├── logger/            # 日志功能
│   └── utils/             # 工具函数
├── tests/                 # 测试文件
│   ├── unit/              # 单元测试
│   ├── integration/       # 集成测试
│   └── api_test.sh        # API测试脚本
├── scripts/               # 脚本文件
│   ├── create_module.sh   # 模块生成脚本
│   ├── build.sh           # 构建脚本
│   └── deploy.sh          # 部署脚本
├── migrations/            # 数据库迁移文件
├── docs/                  # 文档
├── config.yaml            # 配置文件
├── Makefile              # 构建工具
├── Dockerfile            # Docker配置
├── docker-compose.yml    # Docker Compose配置
└── README.md             # 项目说明
```

## 🚀 常用命令

### 开发相关
```bash
make dev          # 启动开发服务器（热重载）
make build        # 构建应用
make run          # 运行应用
make clean        # 清理构建文件
```

### 测试相关
```bash
make test         # 运行所有测试
make test-unit    # 运行单元测试
make test-integration # 运行集成测试
make test-coverage    # 生成测试覆盖率报告
```

### 代码质量
```bash
make fmt          # 格式化代码
make lint         # 代码检查
make security     # 安全检查
```

### 数据库相关
```bash
make migrate      # 运行数据库迁移
make migrate-down # 回滚数据库迁移
make migration name=create_users # 创建新的迁移文件
make db-reset     # 重置数据库
```

### 模块生成
```bash
make new-module name=product # 生成新模块
```

### Docker相关
```bash
make docker-build # 构建Docker镜像
make docker-run   # 运行Docker容器
```

## ⚙️ 配置说明

### 基础配置
编辑 `config.yaml` 文件：

```yaml
app:
  name: "gin-template"
  port: 8080
  mode: "debug"  # debug, release, test

database:
  # MySQL配置（可选，不配置则使用SQLite）
  host: "localhost"
  port: 3306
  username: "root"
  password: "password"
  name: "gin_template"

redis:
  # Redis配置（可选，不配置则使用内存缓存）
  host: "localhost"
  port: 6379
  password: ""
  db: 0

jwt:
  secret_key: "your-secret-key"
  expires_in: 24  # 小时

email:
  smtp_host: "smtp.gmail.com"
  smtp_port: 587
  username: "your-email@gmail.com"
  password: "your-app-password"
```

### 环境变量
支持通过环境变量覆盖配置，格式：`APP_[模块]_[配置项]`

```bash
export APP_APP_PORT=8080
export APP_DB_HOST=localhost
export APP_REDIS_HOST=localhost
```

## 🔧 开发指南

### 创建新模块
使用脚手架快速创建完整的CRUD模块：

```bash
make new-module name=product
```

这将自动生成：
- Model: `internal/models/product.go`
- Repository: `internal/repositories/product_repository.go`
- Service: `internal/services/product_service.go`
- Controller: `internal/controllers/product_controller.go`
- DTO: `internal/dto/request/product.go` 和 `internal/dto/response/product.go`
- 路由: 自动注册到路由系统
- 迁移文件: `migrations/xxx_create_products_table.sql`
- 测试文件: 相应的测试模板

### API响应格式
所有API响应都遵循统一格式：

```json
{
  "code": 200,
  "message": "操作成功",
  "data": {},
  "request_id": "uuid",
  "timestamp": "2024-01-01T12:00:00Z"
}
```

### 错误处理
使用统一的错误处理机制：

```go
// 创建错误
err := errors.New(errors.CodeInvalidParameter, "参数无效")

// 包装错误
err = errors.Wrap(originalErr, errors.CodeDBQueryFailed)

// 在控制器中处理错误
if err != nil {
    errors.HandleError(c, err)
    return
}

// 成功响应
errors.ResponseSuccess(c, data, "操作成功")
```

## 🧪 测试

### 运行测试
```bash
# 运行所有测试
make test

# 运行API测试
./tests/api_test.sh

# 查看测试覆盖率
make test-coverage
```

### 测试结构
- `tests/unit/` - 单元测试
- `tests/integration/` - 集成测试
- `tests/api_test.sh` - API自动化测试脚本

## 🐳 Docker部署

### 开发环境
```bash
# 启动完整开发环境（包含MySQL和Redis）
docker-compose up -d

# 查看日志
docker-compose logs -f app
```

### 生产环境
```bash
# 构建镜像
make docker-build

# 运行容器
docker run -d --name gin-template -p 8080:8080 gin-template:latest
```

## 📝 API文档

### 用户认证API

#### 用户注册
```bash
POST /api/v1/user/register
Content-Type: application/json

{
  "username": "testuser",
  "password": "password123",
  "email": "test@example.com",
  "verification_code": "123456"
}
```

#### 用户登录
```bash
POST /api/v1/user/login
Content-Type: application/json

{
  "username": "testuser",
  "password": "password123"
}
```

#### 获取用户信息
```bash
GET /api/v1/user/profile
Authorization: Bearer <token>
```

### 健康检查
```bash
GET /health
GET /api/v1/health
```

## 🤝 贡献指南

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开 Pull Request

## 📄 许可证

MIT License