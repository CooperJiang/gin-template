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