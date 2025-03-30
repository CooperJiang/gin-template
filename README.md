# 通用Go Web项目模板

这是一个基于Go语言的Web应用程序模板，集成了常用的功能和最佳实践。

## 功能特性

- 基于Gin框架的轻量级Web服务
- 完整的用户认证系统（注册、登录、密码重置）
- JWT身份验证
- 基于角色的权限控制（普通用户、管理员、超级管理员）
- MySQL数据库集成（使用GORM）
- 自动降级到SQLite数据库（未配置MySQL时）
- Redis缓存支持（可选，自动降级为内存缓存）
- 前端资源嵌入（使用Go embed打包静态文件）
- 邮件发送功能
- 统一的错误处理和响应格式
- 请求参数验证
- 定时任务支持
- 跨域(CORS)支持
- 统一的日志记录
- 多级健康检查系统
- 标准化的错误处理系统（错误码、堆栈追踪、请求ID关联）

## 项目结构

```
template/
├── cmd/                   # 命令行工具和入口点
│   └── main.go            # 应用程序入口点
├── internal/              # 内部包，不对外暴露
│   ├── controllers/       # 控制器层，处理HTTP请求
│   ├── middleware/        # HTTP中间件
│   ├── models/            # 数据模型定义
│   ├── routes/            # 路由定义
│   ├── services/          # 业务逻辑层
│   ├── cron/              # 定时任务
│   └── static/            # 静态资源
│       └── dist/          # 前端构建后的静态文件
│           ├── js/        # JavaScript文件
│           └── index.html # 主HTML文件
├── pkg/                   # 可重用的库代码
│   ├── cache/             # 缓存相关功能
│   ├── common/            # 常用工具和通用代码
│   ├── config/            # 配置相关
│   ├── database/          # 数据库连接和操作
│   ├── email/             # 邮件功能
│   ├── logger/            # 日志功能
│   └── utils/             # 通用工具函数
├── config.yaml            # 配置文件
├── go.mod                 # Go模块定义
└── README.md              # 项目说明文档
```

## 数据库配置

项目支持两种数据库：

1. **MySQL数据库**（默认）
   - 在`config.yaml`中配置MySQL连接参数
   - 完整配置示例：
   ```yaml
   database:
     host: localhost
     port: 3306
     username: root
     password: yourpassword
     name: template_db
     charset: utf8mb4
   ```

2. **SQLite数据库**（自动降级）
   - 当MySQL配置不完整时，系统自动使用SQLite数据库
   - 无需额外配置，系统会创建`app.db`文件作为SQLite数据库

## 环境变量配置

项目支持通过环境变量覆盖配置文件中的设置，格式为：`APP_[模块]_[配置项]`

环境变量配置优先级：**环境变量 > 配置文件 > 默认值**

### 示例：

```bash
# 应用基本配置
export APP_APP_NAME="我的应用"
export APP_APP_PORT=8080
export APP_APP_MODE=release

# 数据库配置
export APP_DB_HOST=db.example.com
export APP_DB_PORT=3306
export APP_DB_USERNAME=dbuser
export APP_DB_PASSWORD=secret
export APP_DB_NAME=myapp

# Redis配置
export APP_REDIS_HOST=redis.example.com
export APP_REDIS_PORT=6379

# JWT配置
export APP_JWT_SECRET_KEY=my-secret-key
export APP_JWT_EXPIRES_IN=48
```

### Docker环境配置示例：

```yaml
# docker-compose.yml 示例
version: '3'
services:
  app:
    image: your-app-image
    environment:
      - APP_APP_PORT=8080
      - APP_DB_HOST=mysql
      - APP_DB_USERNAME=root
      - APP_DB_PASSWORD=password
      - APP_DB_NAME=app_db
      - APP_REDIS_HOST=redis
```

## 前端集成

本项目支持将前端构建文件嵌入到Go二进制文件中：

1. 构建前端项目并将生成的文件放入`internal/static/dist/`目录
   ```bash
   # 例如，使用Vue.js构建前端
   cd frontend
   npm run build
   cp -r dist/* ../internal/static/dist/
   ```

2. 在`internal/static/static.go`中，已通过Go embed指令将这些文件嵌入：
   ```go
   //go:embed dist/*.html
   var DistDir embed.FS
   ```

3. 嵌入后的前端文件会随Go二进制一起打包，无需额外部署静态文件

## 健康检查系统

项目内置了完整的健康检查系统，提供多种粒度的健康状态监控：

### 健康检查接口

- `GET /health` - 简单健康检查，返回基础服务状态
- `GET /health/basic` - 基础健康检查，检查关键服务组件状态
- `GET /health/complete` - 完整健康检查，检查所有系统组件状态

### 健康检查响应示例

```json
{
  "status": "up",
  "version": "1.0.0",
  "check_time": "2023-01-01T12:00:00Z",
  "components": [
    {
      "name": "database",
      "status": "up",
      "check_time": "2023-01-01T12:00:00Z",
      "details": {
        "ping_time_ms": 5,
        "max_open_conns": 10,
        "open_conns": 2,
        "in_use": 1
      }
    },
    {
      "name": "cache",
      "status": "up",
      "check_time": "2023-01-01T12:00:00Z",
      "details": {
        "set_time_ms": 3,
        "get_time_ms": 2,
        "type": "redis"
      }
    }
  ],
  "system_info": {
    "go_version": "go1.20",
    "goroutines": 15,
    "cpu_num": 8,
    "memory_alloc": 8562400,
    "os": "linux",
    "arch": "amd64"
  },
  "uptime": "24h0m0s",
  "start_time": "2023-01-01T00:00:00Z"
}
```

### 状态说明

- `up`: 组件正常运行
- `down`: 组件不可用
- `degraded`: 组件性能下降但仍可用

### 自定义健康检查

可通过实现`health.Checker`接口并注册到健康检查系统来添加自定义检查项：

```go
// 创建自定义检查器
type MyServiceChecker struct{}

func (c *MyServiceChecker) Name() string {
	return "my_service"
}

func (c *MyServiceChecker) Check() (health.Status, map[string]interface{}) {
	// 执行检查逻辑
	return health.StatusUp, map[string]interface{}{
		"custom_metric": 100,
	}
}

func (c *MyServiceChecker) Type() health.CheckType {
	return health.CheckTypeBasic
}

// 注册自定义检查器
func init() {
	health.RegisterChecker(&MyServiceChecker{})
}
```

## 标准化错误处理系统

项目实现了完整的标准化错误处理系统，确保统一的错误响应格式和错误处理流程：

### 特性

- 统一的错误码体系（包括通用错误、用户相关错误、数据库错误、第三方服务错误等）
- 错误堆栈跟踪和请求ID关联，方便排查问题
- 安全的错误信息处理，防止敏感信息泄露
- 请求验证错误的统一处理
- 统一的API响应格式

### 错误响应格式

```json
{
  "code": 100,
  "message": "无效的请求参数",
  "request_id": "550e8400-e29b-41d4-a716-446655440000",
  "timestamp": "2023-01-01T12:00:00Z"
}
```

### 错误码分类

- **通用错误码**: 1-999
  - 包括未知错误、内部错误、参数错误、权限错误等

- **用户相关错误码**: 1000-1999
  - 包括用户不存在、密码错误、账号禁用等

- **数据库相关错误码**: 2000-2999
  - 包括数据库连接失败、查询失败、记录不存在等

- **第三方服务错误码**: 3000-3999
  - 包括外部API错误、Redis错误、邮件服务错误等

### 使用示例

在控制器中抛出和处理错误：

```go
import "template/pkg/errors"

func SomeHandler(c *gin.Context) {
    // 验证请求
    req, err := common.ValidateRequest[SomeDTO](c)
    if err != nil {
        errors.HandleError(c, err)
        return
    }

    // 业务处理
    if somethingIsWrong {
        err := errors.New(errors.CodeInvalidParameter, "详细错误信息")
        errors.HandleError(c, err)
        return
    }

    // 成功响应
    errors.ResponseSuccess(c, data, "操作成功")
}
```

### 自定义错误

可以通过以下方式创建定制化错误：

```go
// 创建基本错误
err := errors.New(errors.CodeDBNoRecord, "找不到用户记录")

// 添加请求ID
err = err.WithRequestID("request-123")

// 添加元数据
err = err.WithMetadata("user_id", 123)

// 包装错误
originalErr := someDatabase.Query()
wrappedErr := errors.Wrap(originalErr, errors.CodeDBQueryFailed, "查询数据库失败")
```

### 验证错误处理

系统集成了请求验证功能，可以自动生成友好的验证错误消息：

```go
// 定义请求DTO类型
type LoginRequest struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required,min=6"`
}

// 实现GetValidationMessages方法（可选）
func (r *LoginRequest) GetValidationMessages() map[string]string {
    return map[string]string{
        "Username.required": "用户名不能为空",
        "Password.required": "密码不能为空",
        "Password.min":      "密码长度不能少于6位",
    }
}

// 在控制器中使用
req, err := common.ValidateRequest[LoginRequest](c)
if err != nil {
    errors.HandleError(c, err)
    return
}
```

## 快速开始

1. 克隆项目模板
```bash
git clone https://github.com/yourusername/template.git
cd template
```

2. 修改配置文件 `config.yaml` 适配你的环境
   - 如不配置MySQL，将自动使用SQLite数据库

3. 将前端构建文件放入`internal/static/dist/`目录

4. 构建并运行应用
```bash
go build -o app cmd/main.go
./app
```

## Docker部署

项目支持Docker容器化部署，提供了生产环境和开发环境的配置。

### 开发环境

使用Docker Compose快速启动完整的开发环境，包括应用、MySQL和Redis：

```bash
# 启动所有服务
docker-compose up -d

# 查看日志
docker-compose logs -f

# 停止所有服务
docker-compose down
```

开发环境特性：
- 代码热重载（使用Air工具）
- 源代码与容器同步
- 预配置的MySQL和Redis服务
- 持久化的数据卷

### 生产环境

使用多阶段构建优化的Dockerfile部署到生产环境：

```bash
# 构建生产镜像
docker build -t your-app-name:latest .

# 运行容器
docker run -d --name your-app -p 9000:9000 \
  -e APP_APP_PORT=9000 \
  -e APP_DB_HOST=your-db-host \
  -e APP_DB_USERNAME=your-username \
  -e APP_DB_PASSWORD=your-password \
  -e APP_DB_NAME=your-db-name \
  your-app-name:latest
```

生产环境特性：
- 多阶段构建，镜像体积小
- 非root用户运行，增强安全性
- 环境变量配置，便于在不同环境部署
- 基于Alpine的轻量级基础镜像

### 使用GitHub Actions自动构建Docker镜像

```yaml
name: Docker Build and Push

on:
  push:
    branches: [ main ]
    tags: [ 'v*' ]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v3
        
      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v2
        
      - name: Login to Docker Hub
        uses: docker/login-action@v2
        with:
          username: ${{ secrets.DOCKERHUB_USERNAME }}
          password: ${{ secrets.DOCKERHUB_TOKEN }}
          
      - name: Build and push
        uses: docker/build-push-action@v4
        with:
          context: .
          push: true
          tags: yourusername/yourapp:latest
```

## 默认基础API文档

### 用户相关API

- `POST /api/user/register` - 用户注册
- `POST /api/user/login` - 用户登录
- `POST /api/user/send-registration-code` - 发送注册验证码
- `POST /api/user/send-reset-password-code` - 发送重置密码验证码
- `POST /api/user/reset-password` - 重置密码

## 自定义和扩展

1. 添加新的控制器:
   - 在 `internal/controllers/` 目录下创建新文件
   - 在 `internal/routes/` 中注册路由

2. 添加新的模型:
   - 在 `internal/models/` 目录下创建新模型
   - 在 `pkg/database/database.go` 的 `autoMigrate()` 函数中注册模型

3. 添加新的定时任务:
   - 在 `internal/cron/` 目录下创建新的定时任务文件
   - 在 `internal/cron/cron.go` 的 `registerTasks()` 函数中注册任务

4. 修改前端嵌入配置:
   - 根据需要在 `internal/static/static.go` 中调整 embed 指令
   - 可添加更多静态资源类型，如CSS、图片等

## 许可证

MIT 