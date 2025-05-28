# 🚀 Gin 全栈项目模板

一个功能完整、开箱即用的全栈Web应用模板，包含Go后端（基于Gin框架）和Vue3前端，集成了企业级开发所需的各种功能和最佳实践。

## 🎯 项目特色

### 🔥 全栈解决方案
- **后端**: Go + Gin + GORM + JWT + Redis/MySQL
- **前端**: Vue3 + TypeScript + TailwindCSS + Pinia + Vue Router
- **开发工具**: 热重载、自动生成、一键部署
- **企业级**: 分层架构、统一规范、完整测试

### ⚡ 极速开发
- **5分钟启动**: 一键初始化完整开发环境
- **模块生成**: 前后端CRUD代码自动生成
- **热重载**: 前后端代码变更自动刷新
- **类型安全**: 前后端类型定义完全同步

## 🏗️ 项目结构

```
gin-template/
├── 📁 后端 (Go + Gin)
│   ├── cmd/                   # 应用程序入口
│   ├── internal/              # 内部包
│   │   ├── controllers/       # 控制器层
│   │   ├── services/          # 业务逻辑层
│   │   ├── repositories/      # 数据访问层
│   │   ├── models/            # 数据模型
│   │   └── dto/               # 数据传输对象
│   ├── pkg/                   # 可重用库
│   ├── scripts/               # 脚本工具
│   └── Makefile              # 后端构建工具
│
├── 📁 前端 (Vue3 + TypeScript)
│   ├── web/
│   │   ├── src/
│   │   │   ├── api/           # API接口
│   │   │   ├── components/    # Vue组件
│   │   │   ├── stores/        # Pinia状态管理
│   │   │   ├── views/         # 页面组件
│   │   │   ├── router/        # 路由配置
│   │   │   └── types/         # TypeScript类型
│   │   ├── package.json
│   │   └── Makefile          # 前端构建工具
│
└── 📁 文档和配置
    ├── docs/                  # 项目文档
    ├── docker-compose.yml     # 容器编排
    └── README.md             # 项目说明
```

## ⚡ 快速开始

### 🚀 后端启动

```bash
# 1. 克隆项目
git clone <your-repo-url>
cd gin-template

# 2. 一键初始化后端环境
make setup

# 3. 启动后端开发服务器（热重载）
make dev
```

后端服务运行在: http://localhost:8080

### 🎨 前端启动

```bash
# 1. 进入前端目录
cd web

# 2. 安装依赖
npm install

# 3. 启动前端开发服务器
npm run dev
# 或者使用便捷脚本
./start.sh
```

前端应用运行在: http://localhost:3000

### 🎯 5分钟创建你的第一个模块

```bash
# 后端：自动生成完整的CRUD模块
make new-module name=product

# 运行数据库迁移
make migrate

# 前端：TODO - 前端模块生成器（开发中）
# cd web && npm run generate:module product

# 重启服务
make dev          # 后端
cd web && npm run dev # 前端
```

## 📚 文档

| 文档 | 描述 |
|------|------|
| 📋 **[项目开发规范](docs/PROJECT_STANDARDS.md)** | 完整的项目开发规范和最佳实践 |
| 🤖 **[AI开发指南](docs/AI_DEVELOPMENT_GUIDE.md)** | 专门为AI助手提供的开发指南 |
| 🎨 **[前端文档](web/README.md)** | Vue3前端项目详细文档 |

## 🛠️ 技术栈

### 后端技术栈
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

### 前端技术栈
| 类别 | 技术 | 说明 |
|------|------|------|
| **框架** | Vue 3 | 渐进式JavaScript框架 |
| **语言** | TypeScript | 类型安全的JavaScript |
| **构建工具** | Vite | 下一代前端构建工具 |
| **状态管理** | Pinia | Vue官方状态管理库 |
| **路由** | Vue Router | Vue官方路由管理器 |
| **UI框架** | TailwindCSS | 实用优先的CSS框架 |
| **HTTP客户端** | Axios | 基于Promise的HTTP客户端 |
| **工具库** | VueUse | Vue组合式API工具集 |
| **图标** | Heroicons | 精美的SVG图标库 |

## ⚡ 核心特性

### 🏗️ 开发效率
- **一键模块生成**: `make new-module name=product` 自动生成完整CRUD
- **热重载开发**: 前后端代码变更自动重启/刷新
- **完整脚手架**: Controller、Service、Repository、DTO、测试一键生成
- **智能工具链**: 集成代码检查、测试、构建、部署等完整工具链

### 🛡️ 企业级特性
- **分层架构**: Controller → Service → Repository → Model
- **UUID主键**: 所有模型使用UUID，确保数据安全
- **统一错误处理**: 标准化错误码和响应格式
- **参数验证**: 自动参数验证和错误提示
- **缓存支持**: Redis/内存缓存自动降级
- **数据库支持**: MySQL/SQLite自动降级
- **用户认证**: JWT + 角色权限控制

### 🔧 开发工具
- **代码质量**: 集成 golangci-lint、gosec、ESLint、Prettier
- **自动化测试**: 单元测试和集成测试模板
- **API测试**: 完整的API自动化测试脚本
- **性能分析**: 内置性能检查和优化工具
- **Docker支持**: 多阶段构建和容器化部署

### 🎨 前端特性
- **现代化UI**: 基于TailwindCSS的响应式设计
- **类型安全**: 完整的TypeScript类型定义
- **状态管理**: Pinia状态管理，支持持久化
- **路由守卫**: 基于JWT的路由权限控制
- **通知系统**: 全局通知组件，支持多种类型
- **API集成**: 与后端API完美集成，统一错误处理

## 🚀 常用命令

### 后端命令
```bash
make dev          # 启动开发服务器（热重载）
make build        # 构建应用
make test         # 运行所有测试
make lint         # 代码检查
make migrate      # 运行数据库迁移
make new-module name=product # 生成新模块
```

### 前端命令
```bash
cd web
npm run dev       # 启动前端开发服务器
npm run build     # 构建生产版本
npm run lint      # 代码检查
npm run format    # 格式化代码
npm run type-check # TypeScript类型检查
```

### Docker命令
```bash
make docker-build # 构建Docker镜像
make docker-run   # 运行Docker容器
docker-compose up # 启动完整服务栈
```

## 🎯 项目亮点

### 🚀 开箱即用
- **零配置启动**: 克隆即用，无需复杂配置
- **完整示例**: 包含用户管理、认证等完整功能示例
- **最佳实践**: 集成行业最佳实践和设计模式

### 🔧 开发体验
- **智能提示**: 完整的TypeScript类型定义
- **错误处理**: 统一的错误处理和用户友好的错误提示
- **调试友好**: 详细的日志和调试信息

### 🛡️ 生产就绪
- **性能优化**: 内置缓存、连接池、资源优化
- **安全加固**: JWT认证、参数验证、SQL注入防护
- **监控支持**: 健康检查、指标收集、日志聚合

## 📝 许可证

MIT License - 详见 [LICENSE](LICENSE) 文件

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📞 支持

如有问题，请提交 Issue 或联系维护者。