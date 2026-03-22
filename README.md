# 🚀 Gin三端全栈模板

一个现代化的Go+Vue3全栈项目模板，采用**真正的三端分离架构**：后端API + 管理端前端 + 用户端前端。

## 📊 项目架构

```
email-manage/
├── 🎯 后端 API (Go + Gin)              - 端口 9000
├── 🎨 管理端前端 (Vue3 + TypeScript)    - 端口 3000 (开发) / /admin (生产)
└── 👥 用户端前端 (Vue3 + TypeScript)    - 端口 4000 (开发) / / (生产)
```

### 🏗️ 技术栈

**后端 (Go)**
- 🔸 **Gin** - 高性能Web框架
- 🔸 **GORM** - ORM数据库操作
- 🔸 **JWT** - 用户认证
- 🔸 **Swagger** - API文档
- 🔸 **SQLite/MySQL/PostgreSQL** - 数据库支持

**前端 (Vue3)**
- 🔸 **Vue 3** - 渐进式框架
- 🔸 **TypeScript** - 类型安全
- 🔸 **Vite** - 极速构建工具
- 🔸 **TailwindCSS** - 原子化CSS
- 🔸 **Pinia** - 状态管理
- 🔸 **Vue Router** - 路由管理
- 🔸 **Axios** - HTTP客户端

## 🚀 快速开始

### 1️⃣ 克隆项目
```bash
git clone <repository-url>
cd email-manage
```

### 2️⃣ 启动后端服务
```bash
# 安装Go依赖并启动后端
go mod tidy
make dev

# 后端将在 http://localhost:9000 启动
```

### 3️⃣ 启动管理端前端
```bash
# 新开终端，安装并启动管理端
make admin-install
make admin-dev

# 管理端将在 http://localhost:3000 启动
```

### 4️⃣ 启动用户端前端
```bash
# 新开终端，安装并启动用户端
make web-install  
make web-dev

# 用户端将在 http://localhost:4000 启动
```

### 5️⃣ 一键启动全栈开发环境
```bash
# 并行启动所有服务
make fullstack-dev

# 将同时启动：
# - 后端 API: http://localhost:9000
# - 管理端: http://localhost:3000 
# - 用户端: http://localhost:4000
```

## 📁 目录结构

```
email-manage/
├── 📂 cmd/                    # 应用入口
├── 📂 internal/               # 内部代码
│   ├── 📂 api/               # API控制器
│   ├── 📂 middleware/        # 中间件
│   ├── 📂 models/            # 数据模型
│   ├── 📂 routes/            # 路由配置
│   ├── 📂 services/          # 业务逻辑
│   └── 📂 static/            # 静态文件嵌入
│       ├── 📂 admin/         # 管理端构建文件
│       └── 📂 web/           # 用户端构建文件
├── 📂 admin/                  # 🎨 管理端前端
│   ├── 📂 src/
│   │   ├── 📂 components/    # Vue组件
│   │   ├── 📂 views/         # 页面组件
│   │   ├── 📂 stores/        # Pinia状态管理
│   │   ├── 📂 router/        # 路由配置
│   │   └── 📂 utils/         # 工具函数
│   ├── 📄 package.json
│   ├── 📄 vite.config.ts
│   └── 📄 tailwind.config.js
├── 📂 web/                    # 👥 用户端前端
│   ├── 📂 src/
│   │   ├── 📂 components/    # Vue组件
│   │   ├── 📂 views/         # 页面组件
│   │   ├── 📂 stores/        # Pinia状态管理
│   │   ├── 📂 router/        # 路由配置
│   │   └── 📂 utils/         # 工具函数
│   ├── 📄 package.json
│   ├── 📄 vite.config.ts
│   └── 📄 tailwind.config.js
├── 📂 scripts/               # 构建脚本
│   ├── 📄 build_admin.sh     # 管理端构建
│   └── 📄 build_web.sh       # 用户端构建
├── 📄 Makefile               # 构建工具
├── 📄 go.mod
└── 📄 README.md
```

## 🔧 开发命令

### 📱 管理端命令
```bash
make admin-install     # 安装管理端依赖
make admin-dev        # 启动管理端开发服务器
make admin-build      # 构建管理端并部署到静态目录
make admin-lint       # 检查管理端代码
make admin-format     # 格式化管理端代码
make admin-clean      # 清理管理端构建文件
```

### 🌐 用户端命令
```bash
make web-install      # 安装用户端依赖
make web-dev         # 启动用户端开发服务器
make web-build       # 构建用户端并部署到静态目录
make web-lint        # 检查用户端代码
make web-format      # 格式化用户端代码
make web-clean       # 清理用户端构建文件
```

### 🔗 前端组合命令
```bash
make frontend-install # 安装所有前端依赖
make frontend-build   # 构建所有前端项目
make frontend-clean   # 清理所有前端构建文件
make frontend-lint    # 检查所有前端代码
make frontend-format  # 格式化所有前端代码
```

### 🚀 全栈命令
```bash
make fullstack-build  # 构建完整全栈应用 (管理端+用户端+后端)
make fullstack-dev    # 启动全栈开发环境 (并行启动前后端)
make fullstack-clean  # 清理所有构建文件
```

### ⚙️ 后端命令
```bash
make dev             # 启动后端开发服务器
make build           # 构建后端应用
make test            # 运行测试
make clean           # 清理构建文件
```

## 🎯 访问地址

### 开发环境
- 🎯 **后端API**: `http://localhost:9000`
- 🎨 **管理端**: `http://localhost:3000`
- 👥 **用户端**: `http://localhost:4000`

### 生产环境
- 🎯 **后端API**: `http://your-domain:9000/api/v1`
- 🎨 **管理端**: `http://your-domain:9000/admin`
- 👥 **用户端**: `http://your-domain:9000/`

## 🏗️ 构建部署

### 开发构建
```bash
# 单独构建
make admin-build      # 仅构建管理端
make web-build        # 仅构建用户端
make build            # 仅构建后端

# 全栈构建
make fullstack-build  # 构建所有组件
```

### 🚀 一键部署到服务器
```bash
# 首次部署 (3步搞定)
make deploy-config    # 1. 配置服务器信息
make deploy-setup     # 2. 初始化服务器环境
make build-deploy     # 3. 构建并部署

# 日常更新 (一个命令)
make build-deploy     # 构建并部署

# 更多部署命令
make deploy-status    # 查看应用状态
make deploy-logs      # 查看应用日志  
make deploy-rollback  # 回滚到上一版本
```

**特性:**
- ✅ **2-3秒快速切换** - 几乎零停机部署
- ✅ **版本管理** - 保留最近5个版本，支持一键回滚
- ✅ **智能检测** - 自动处理端口占用和进程管理
- ✅ **多环境支持** - prod/test/dev环境配置
- ✅ **版本自动递增** - Admin后台版本号自动更新

📖 详细部署文档: [DEPLOY.md](DEPLOY.md) | ⚡ 快速参考: [QUICK_DEPLOY.md](QUICK_DEPLOY.md)

### 生产部署
```bash
# 1. 构建完整应用
make fullstack-build

# 2. 运行二进制文件
./bin/email-manage

# 应用将在端口9000启动，包含：
# - API服务
# - 管理端 (通过 /admin 访问)
# - 用户端 (通过 / 访问)
```

## 🔐 认证系统

### 管理端认证
- 登录地址：`/admin/login`
- 管理员用户管理
- 权限控制系统
- JWT令牌认证

### 用户端认证
- 登录地址：`/login`
- 用户注册/登录
- 密码重置功能
- 个人资料管理

## 🎨 前端特性

### 管理端特性
- 🔸 **现代化管理界面** - 基于Vue3 + TypeScript
- 🔸 **响应式设计** - 适配各种屏幕尺寸
- 🔸 **权限管理** - 基于角色的访问控制
- 🔸 **数据表格** - 完整的CRUD操作
- 🔸 **图表统计** - 数据可视化支持

### 用户端特性
- 🔸 **用户友好界面** - 简洁优雅的设计
- 🔸 **完整认证流程** - 登录/注册/找回密码
- 🔸 **个人中心** - 用户信息管理
- 🔸 **响应式布局** - 移动端适配
- 🔸 **PWA支持** - 可安装为应用

## 🌟 项目特色

### 💡 真正的三端分离
- 管理端和用户端**完全独立**的前端项目
- 不同的技术栈配置和依赖管理
- 独立的构建和部署流程

### 🚀 现代化技术栈
- Go 1.21+ 后端服务
- Vue 3 + TypeScript 前端
- Vite + TailwindCSS 开发体验
- 完整的类型安全保障

### 🔧 开发友好
- 丰富的Makefile命令
- 热重载开发环境
- 完整的代码规范和检查工具
- 详细的错误处理和日志记录

### 📦 部署简单
- 单一二进制文件部署
- 静态文件嵌入技术
- Docker支持
- 云原生友好

## 🤝 贡献指南

1. Fork 本项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

## 📄 许可证

本项目基于 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情

## 💬 联系我们

如有问题或建议，请提交 Issue 或联系项目维护者。

---

⭐ 如果这个项目对你有帮助，请给它一个星标！