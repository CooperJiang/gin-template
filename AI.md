# AI 开发规范文档

这是一个基于 Go + Gin + Vue3 + TypeScript 的全栈模板项目，旨在为 AI 开发者提供标准化的开发规范和最佳实践。

## 项目概述

### 架构设计
- **后端**: Go + Gin 框架，采用分层架构设计
- **前端**: Vue3 + TypeScript + Vite + TailwindCSS
- **部署**: 前端构建后嵌入到 Go 二进制文件中，实现单文件部署
- **数据库**: SQLite (开发) / PostgreSQL (生产)
- **缓存**: Redis
- **认证**: JWT

### 构建流程
1. 前端项目在 `web/` 目录下独立开发
2. 构建时将前端 `dist/` 文件夹复制到 `internal/static/dist/`
3. 通过 Go embed 将静态文件嵌入到二进制文件中
4. 最终生成单个二进制文件，包含完整的前后端功能

---

## 后端开发规范

### 📁 项目结构

```
├── cmd/                    # 应用程序入口
│   └── main.go            # 主程序启动文件
├── internal/              # 内部模块 (不对外暴露)
│   ├── controllers/       # 控制器层 - 处理 HTTP 请求
│   ├── services/          # 业务逻辑层 - 核心业务处理
│   ├── repositories/      # 数据访问层 - 数据库操作
│   ├── models/           # 数据模型 - 数据库表结构
│   ├── dto/              # 数据传输对象 - 请求/响应结构
│   ├── routes/           # 路由定义 - API 路由配置
│   ├── middleware/       # 中间件 - 认证、日志、CORS 等
│   ├── migrations/       # 数据库迁移文件
│   └── static/           # 嵌入的静态文件 (前端构建产物)
├── pkg/                   # 可复用的公共包
│   ├── common/           # 通用工具和助手函数
│   ├── config/           # 配置管理
│   ├── database/         # 数据库连接和管理
│   ├── logger/           # 日志系统
│   ├── cache/            # Redis 缓存
│   ├── email/            # 邮件服务
│   ├── errors/           # 错误处理和统一返回
│   ├── utils/            # 工具函数
│   └── constants/        # 全局常量定义
└── tests/                # 测试文件
    ├── unit/             # 单元测试
    └── integration/      # 集成测试
```

### 🎯 命名规范

**统一采用驼峰命名法 (camelCase)**

- **文件名**: `userController.go`, `emailService.go`
- **包名**: `usercontroller`, `emailservice` (小写)
- **函数名**: `GetUserInfo()`, `SendEmail()`
- **变量名**: `userId`, `emailAddress`
- **常量名**: `DefaultPageSize`, `MaxRetryCount`
- **结构体**: `UserInfo`, `LoginRequest`

### 🏗️ MVC 架构分层

#### Controller 层 (控制器)
**位置**: `internal/controllers/`
**职责**: 处理 HTTP 请求，参数校验，调用 Service 层

```go
// internal/controllers/user/userController.go
func (ctrl *UserController) GetUserInfo(c *gin.Context) {
    // 1. 参数校验
    var req request.GetUserInfoRequest
    if err := common.ValidateRequest[request.GetUserInfoRequest](c); err != nil {
        return
    }
    
    // 2. 调用 Service 层
    user, err := ctrl.userService.GetUserInfo(req.UserID)
    if err != nil {
        errors.ResponseError(c, err)
        return
    }
    
    // 3. 返回结果
    errors.ResponseSuccess(c, user, "获取用户信息成功")
}
```

#### Service 层 (业务逻辑)
**位置**: `internal/services/`
**职责**: 核心业务逻辑处理，调用 Repository 层

```go
// internal/services/user/userService.go
func (s *UserService) GetUserInfo(userID string) (*models.User, error) {
    // 业务逻辑处理
    user, err := s.userRepo.GetByID(userID)
    if err != nil {
        return nil, err
    }
    
    // 业务规则验证
    if user.Status != constants.UserStatusActive {
        return nil, errors.New("用户已被禁用")
    }
    
    return user, nil
}
```

#### Repository 层 (数据访问)
**位置**: `internal/repositories/`
**职责**: 数据库操作，CRUD 操作

```go
// internal/repositories/user/userRepository.go
func (r *UserRepository) GetByID(userID string) (*models.User, error) {
    var user models.User
    err := r.db.Where("id = ?", userID).First(&user).Error
    return &user, err
}
```

### 📝 DTO 和参数校验

#### DTO 定义
**位置**: `internal/dto/request/` 和 `internal/dto/response/`

```go
// internal/dto/request/userRequest.go
type LoginRequest struct {
    Account  string `json:"account" binding:"required" validate:"min=3,max=50"`
    Password string `json:"password" binding:"required" validate:"min=6,max=50"`
}

type RegisterRequest struct {
    Username string `json:"username" binding:"required" validate:"min=3,max=30"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required" validate:"min=6,max=50"`
    Code     string `json:"code" binding:"required" validate:"len=6"`
}
```

#### 参数校验使用
```go
// 在 Controller 中使用
var req request.LoginRequest
if err := common.ValidateRequest[request.LoginRequest](c); err != nil {
    return // 自动返回错误信息
}
```

### 🗄️ 数据库层

#### 数据库连接
**位置**: `pkg/database/`
**配置**: `config.yaml` 中的 database 配置节

```go
// 数据库初始化
db := database.InitDB()

// 自动迁移 (在 main.go 中)
database.AutoMigrate(db)
```

#### Model 定义
**位置**: `internal/models/`

```go
// internal/models/user.go
type User struct {
    BaseModel
    Username string    `gorm:"uniqueIndex;not null" json:"username"`
    Email    string    `gorm:"uniqueIndex;not null" json:"email"`
    Password string    `gorm:"not null" json:"-"`
    Status   int       `gorm:"default:1" json:"status"`
    Role     int       `gorm:"default:2" json:"role"`
}
```

#### 自动迁移
**在 `internal/models/` 中的任何 model 文件放置后**:
1. 在 `pkg/database/migrate.go` 中添加到迁移列表
2. 程序启动时自动完成建表迁移

### 📨 统一返回方法

**使用 `errors.ResponseSuccess` 和 `errors.ResponseError`**

```go
// 成功返回
errors.ResponseSuccess(c, data, "操作成功")

// 错误返回
errors.ResponseError(c, err)

// 自定义错误
errors.ResponseErrorWithCode(c, 400, "参数错误")
```

### 📚 常量和公共方法

#### 常量定义
**位置**: `pkg/constants/`

```go
// pkg/constants/user.go
const (
    UserStatusActive   = 1
    UserStatusInactive = 2
    UserRoleAdmin      = 1
    UserRoleUser       = 2
)
```

#### 公共方法
**位置**: `pkg/common/`

- `ValidateRequest[T]()` - 参数校验
- `GetUserFromContext()` - 从上下文获取用户
- `GenerateID()` - 生成唯一ID
- `HashPassword()` - 密码加密

### 📋 日志系统

**位置**: `pkg/logger/`
**使用方式**:

```go
import "your-project/pkg/logger"

// 记录日志
logger.Info("用户登录", "user_id", userID)
logger.Error("数据库连接失败", "error", err)
logger.Debug("调试信息", "data", data)
```

### 🚀 Redis 缓存

**位置**: `pkg/cache/`
**使用方式**:

```go
import "your-project/pkg/cache"

// 基本操作
cache.Set("key", value, time.Hour)
value, err := cache.Get("key")
cache.Delete("key")

// 在 Service 中使用
cacheKey := fmt.Sprintf("user:%s", userID)
if user, exists := cache.GetUser(cacheKey); exists {
    return user, nil
}
```

### 🗃️ 数据库管理

#### 连接配置
**文件**: `config.yaml`
```yaml
database:
  driver: "sqlite"
  dsn: "app.db"
  max_open_conns: 100
  max_idle_conns: 10
```

#### 自动迁移流程
1. 在 `internal/models/` 创建模型文件
2. 在 `pkg/database/migrate.go` 注册模型
3. 程序启动时自动建表

### 📧 Email 服务

**位置**: `pkg/email/`
**配置**: `config.yaml` 中的 email 配置节

```go
// 发送邮件
emailService := email.NewEmailService()
err := emailService.SendVerificationCode(email, code)
err := emailService.SendResetPassword(email, resetLink)
```

### 🛠️ Common 库功能

**位置**: `pkg/common/`
**包含功能**:

- `validator.go` - 参数校验助手
- `jwt.go` - JWT 令牌处理
- `context.go` - 上下文处理
- `password.go` - 密码加密/验证
- `response.go` - 响应处理助手

### 🛣️ 路由模块规范

#### 路由文件组织
**位置**: `internal/routes/`
**规范**: 一个模块一个文件

```go
// internal/routes/userRoutes.go
func RegisterUserRoutes(r *gin.Engine, userController *controllers.UserController) {
    api := r.Group("/api/v1")
    {
        user := api.Group("/user")
        {
            user.POST("/login", userController.Login)
            user.POST("/register", userController.Register)
            
            // 需要认证的路由
            authUser := user.Use(middleware.AuthMiddleware())
            {
                authUser.GET("/info", userController.GetUserInfo)
                authUser.PUT("/profile", userController.UpdateProfile)
            }
        }
    }
}
```

#### 路由注册
**在 `main.go` 中**:
```go
// 注册路由
routes.RegisterUserRoutes(r, userController)
routes.RegisterAdminRoutes(r, adminController)
```

### 🔒 中间件和 JWT

#### 可用中间件
**位置**: `internal/middleware/`

- `authMiddleware.go` - JWT 认证中间件
- `corsMiddleware.go` - CORS 跨域中间件
- `loggerMiddleware.go` - 请求日志中间件
- `rateLimitMiddleware.go` - 限流中间件

#### JWT 使用
```go
// 生成 JWT
token, err := common.GenerateJWT(userID, role)

// 验证 JWT (在中间件中自动处理)
// 获取当前用户
user, err := common.GetUserFromContext(c)
```

### 🎬 项目启动

**启动文件**: `cmd/main.go`
**功能**:
- 配置加载
- 数据库连接
- 路由注册
- 中间件配置
- 服务启动

### 🔧 Makefile 脚本功能

```bash
# 开发相关
make dev                 # 启动开发服务器 (热重载)
make build               # 构建后端程序
make clean               # 清理构建文件

# 前端相关
make web-dev             # 启动前端开发服务器
make web-build           # 构建前端并嵌入到后端
make web-lint            # 前端代码检查

# 全栈相关
make fullstack-build     # 构建完整全栈应用
make fullstack-dev       # 并行启动前后端开发环境
make fullstack-clean     # 清理所有构建文件

# 测试相关
make test                # 运行所有测试
make test-coverage       # 生成测试覆盖率报告

# 代码质量
make fmt                 # 格式化代码
make lint                # 代码检查
make security            # 安全检查

# 部署相关
make docker-build        # 构建 Docker 镜像
make deploy              # 部署到服务器
```

### ⚠️ 开发约束

**禁止随意创建新文件或目录**，必须遵循以下规范：

1. **新增 Model**: 放在 `internal/models/`，并在 `migrate.go` 中注册
2. **新增 API**: 按模块在对应的 controller/service/repository 中添加
3. **新增常量**: 放在 `pkg/constants/` 对应模块文件中
4. **新增工具函数**: 放在 `pkg/utils/` 或 `pkg/common/` 中
5. **新增中间件**: 放在 `internal/middleware/` 中
6. **新增路由**: 在 `internal/routes/` 对应模块文件中添加

---

## 前端开发规范

### 📁 项目结构

```
web/
├── src/
│   ├── api/              # API 接口封装层
│   │   ├── auth/         # 认证相关接口
│   │   ├── user/         # 用户相关接口
│   │   └── index.ts      # 统一导出
│   ├── components/       # 全局组件
│   │   ├── Button/       # 按钮组件
│   │   │   ├── index.vue # 组件实现
│   │   │   ├── index.ts  # 组件导出
│   │   │   └── types.ts  # 类型定义
│   │   └── index.ts      # 全局注册
│   ├── composables/      # 组合式函数
│   │   ├── useMessage.ts # 消息提示
│   │   └── useLoading.ts # 加载状态
│   ├── constants/        # 常量定义
│   │   ├── index.ts      # 通用常量
│   │   └── api.ts        # API 相关常量
│   ├── hooks/            # Vue Hooks
│   │   ├── common/       # 通用 hooks
│   │   ├── user/         # 用户相关 hooks
│   │   └── index.ts      # 统一导出
│   ├── layouts/          # 布局组件
│   │   ├── AdminLayout.vue
│   │   └── components/   # 布局子组件
│   ├── pages/            # 页面组件
│   │   ├── Login/        # 登录页面
│   │   │   ├── index.vue # 页面主文件
│   │   │   └── components/ # 页面私有组件
│   │   └── Dashboard/    # 仪表板
│   ├── router/           # 路由配置
│   │   ├── index.ts      # 路由主文件
│   │   └── modules/      # 路由模块
│   ├── stores/           # 状态管理
│   │   ├── user/         # 用户状态
│   │   │   ├── index.ts  # 状态定义
│   │   │   └── types.ts  # 类型定义
│   │   └── index.ts      # 统一导出
│   ├── styles/           # 样式文件
│   ├── types/            # TypeScript 类型定义
│   ├── utils/            # 工具函数
│   └── main.ts           # 应用入口
├── package.json          # 依赖配置
├── tsconfig.json         # TypeScript 配置
├── tailwind.config.js    # TailwindCSS 配置
└── vite.config.ts        # Vite 配置
```

### 🎯 命名规范

**统一采用驼峰命名法 (camelCase)**

- **文件名**: `UserProfile.vue`, `userService.ts`
- **组件名**: `UserProfile`, `DataTable`
- **函数名**: `getUserInfo()`, `handleSubmit()`
- **变量名**: `userInfo`, `isLoading`
- **常量名**: `API_BASE_URL`, `DEFAULT_PAGE_SIZE`

### 📄 Pages 页面规范

#### 页面组织原则
1. **一个页面一个文件夹**，主文件命名为 `index.vue`
2. **多级嵌套路由** 对应多层级目录结构
3. **复杂页面** 在同目录创建 `components/` 文件夹存放私有组件

```
pages/
├── Login/
│   └── index.vue
├── Dashboard/
│   ├── index.vue
│   └── components/
│       ├── StatsCard.vue
│       └── ChartWidget.vue
├── System/
│   ├── User/
│   │   ├── index.vue          # /system/user
│   │   ├── List/
│   │   │   └── index.vue      # /system/user/list
│   │   └── Detail/
│   │       └── index.vue      # /system/user/detail
│   └── Role/
│       └── index.vue          # /system/role
```

#### 页面文件模板
```vue
<template>
  <div class="page-container">
    <!-- 页面内容 -->
  </div>
</template>

<script setup lang="ts">
import { defineOptions } from 'vue'

// 必须定义组件名称
defineOptions({
  name: 'LoginPage', // 或其他有意义的名称
})

// 页面逻辑
</script>

<style scoped>
/* 仅在需要自定义样式时添加 */
</style>
```

### 🏗️ Layout 布局组件

#### 布局系统说明
**位置**: `src/layouts/`

- `AdminLayout.vue` - 管理后台布局 (侧边栏 + 顶栏)
- `AuthLayout.vue` - 认证页面布局 (登录/注册)
- `components/` - 布局相关的子组件

#### 布局组件职责
1. **页面框架结构** - 定义整体页面布局
2. **导航管理** - 侧边栏、面包屑、用户菜单
3. **权限控制** - 根据用户角色显示不同内容
4. **响应式设计** - 移动端适配

### 🛣️ Router 路由配置

#### 路由文件结构
```typescript
// router/index.ts - 主路由配置
const routes = [
  {
    path: '/login',
    name: 'login',
    component: () => import('@/pages/Login/index.vue'),
    meta: {
      title: '登录',           // 页面标题
      requiresAuth: false,     // 是否需要认证
      layout: 'auth',          // 使用的布局
      icon: 'login',           // 菜单图标
      hidden: true,            // 是否在菜单中隐藏
      breadcrumb: '登录',      // 面包屑显示名称
      roles: ['admin', 'user'], // 允许访问的角色
    }
  }
]
```

#### Meta 参数说明
- `title` - 页面标题，用于浏览器标签页
- `requiresAuth` - 是否需要登录认证
- `layout` - 使用的布局组件 (admin/auth)
- `icon` - 菜单中显示的图标
- `hidden` - 是否在侧边栏菜单中隐藏
- `breadcrumb` - 面包屑导航显示的名称
- `roles` - 允许访问的用户角色数组

#### 大型项目路由拆分
```typescript
// router/modules/user.ts
export const userRoutes = [
  // 用户相关路由
]

// router/modules/system.ts  
export const systemRoutes = [
  // 系统管理路由
]

// router/index.ts
import { userRoutes } from './modules/user'
import { systemRoutes } from './modules/system'
```

### 🗃️ Stores 状态管理规范

#### 状态管理结构
**规范**: 一个模块一个文件夹

```
stores/
├── user/
│   ├── index.ts          # 用户状态管理
│   ├── types.ts          # 类型定义
│   └── actions.ts        # 异步操作 (可选)
├── app/
│   ├── index.ts          # 应用全局状态
│   └── types.ts
└── index.ts              # 统一导出
```

#### Store 模板
```typescript
// stores/user/types.ts
export interface UserState {
  userInfo: User | null
  token: string | null
  permissions: string[]
}

// stores/user/index.ts
import { defineStore } from 'pinia'
import type { UserState } from './types'

export const useUserStore = defineStore('user', {
  state: (): UserState => ({
    userInfo: null,
    token: null,
    permissions: []
  }),
  
  getters: {
    isLoggedIn: (state) => !!state.token,
    userName: (state) => state.userInfo?.name || ''
  },
  
  actions: {
    setToken(token: string) {
      this.token = token
    },
    
    async login(credentials: LoginRequest) {
      // 异步操作
    }
  }
})
```

### 🧩 Components 全局组件规范

#### 组件文件结构
**每个组件必须包含三个文件**:

```
components/
├── Button/
│   ├── index.vue         # 组件实现
│   ├── index.ts          # 组件导出
│   └── types.ts          # 类型定义
└── DataTable/
    ├── index.vue
    ├── index.ts
    └── types.ts
```

#### 组件模板
```typescript
// components/Button/types.ts
export interface ButtonProps {
  type?: 'primary' | 'secondary' | 'danger'
  size?: 'small' | 'medium' | 'large'
  loading?: boolean
  disabled?: boolean
}

// components/Button/index.vue
<template>
  <button 
    :class="buttonClass" 
    :disabled="disabled || loading"
    @click="handleClick"
  >
    <slot />
  </button>
</template>

<script setup lang="ts">
import type { ButtonProps } from './types'

defineOptions({
  name: 'GlobalButton', // 注意：全局组件需要 Global 前缀
})

withDefaults(defineProps<ButtonProps>(), {
  type: 'primary',
  size: 'medium',
  loading: false,
  disabled: false,
})
</script>

// components/Button/index.ts
export { default } from './index.vue'
export type * from './types'
```

#### 全局注册
```typescript
// components/index.ts
import Button from './Button'
import DataTable from './DataTable'

export default {
  GlobalButton: Button,        // 必须添加 Global 前缀
  GlobalDataTable: DataTable,
}

// main.ts 中注册
import globalComponents from '@/components'

Object.entries(globalComponents).forEach(([name, component]) => {
  app.component(name, component)
})
```

#### 使用方式
```vue
<template>
  <!-- 直接使用，无需导入 -->
  <GlobalButton type="primary" @click="handleClick">
    确认
  </GlobalButton>
  
  <GlobalDataTable :columns="columns" :data="tableData" />
</template>
```

### 📊 Constants 常量定义

**位置**: `src/constants/`
**原则**: 所有可定义为常量的值都应该定义，避免在页面中硬编码

```typescript
// constants/index.ts
export const PAGE_SIZE = 20
export const MAX_UPLOAD_SIZE = 5 * 1024 * 1024 // 5MB

// constants/api.ts
export const API_ENDPOINTS = {
  LOGIN: '/user/login',
  REGISTER: '/user/register',
  USER_INFO: '/user/info',
} as const

// constants/user.ts
export const USER_STATUS = {
  ACTIVE: 1,
  INACTIVE: 2,
  BANNED: 3,
} as const

export const USER_ROLES = {
  ADMIN: 1,
  USER: 2,
} as const
```

### 🛠️ Utils 工具函数

**位置**: `src/utils/`

```typescript
// utils/format.ts
export const formatDate = (date: Date | string) => {
  // 日期格式化
}

export const formatFileSize = (bytes: number) => {
  // 文件大小格式化
}

// utils/validation.ts
export const isEmail = (email: string) => {
  // 邮箱验证
}

export const isPhone = (phone: string) => {
  // 手机号验证
}
```

### 🎣 Hooks 和 Composables

#### Hooks 规范
**位置**: `src/hooks/`
**组织**: 按模块分文件夹

```typescript
// hooks/user/useAuth.ts
export function useAuth() {
  const userStore = useUserStore()
  
  const login = async (credentials: LoginRequest) => {
    // 登录逻辑
  }
  
  const logout = () => {
    // 登出逻辑
  }
  
  return {
    login,
    logout,
    isAuthenticated: computed(() => userStore.isLoggedIn)
  }
}

// hooks/common/useSecureStorage.ts
export function useSecureStorage<T>(key: string, defaultValue: T) {
  // 安全存储逻辑
  return [ref(value), setValue, removeValue]
}
```

#### Composables 规范  
**位置**: `src/composables/`

```typescript
// composables/useMessage.ts
export function useMessage() {
  const success = (message: string, duration = 3000) => {
    // 成功提示
  }
  
  const error = (message: string, duration = 5000) => {
    // 错误提示
  }
  
  return { success, error, warning, info }
}
```

### 🌐 API 接口封装

#### API 文件组织
**位置**: `src/api/`
**规范**: 按模块分文件夹，统一导出

```typescript
// api/user/index.ts
import { request } from '@/utils/request'
import type { LoginRequest, User } from '@/types'

export const userApi = {
  login: (data: LoginRequest) => 
    request.post<LoginResponse>('/user/login', data),
    
  getUserInfo: () => 
    request.get<User>('/user/info'),
    
  updateProfile: (data: UpdateProfileRequest) =>
    request.put<User>('/user/profile', data),
}

// api/index.ts
export * from './user'
export * from './admin'
```

### 🎨 样式规范

#### TailwindCSS 优先
```vue
<template>
  <!-- 优先使用 TailwindCSS -->
  <div class="flex items-center justify-between p-4 bg-white rounded-lg shadow">
    <h1 class="text-xl font-bold text-gray-900">标题</h1>
    <button class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">
      按钮
    </button>
  </div>
</template>

<script setup lang="ts">
// 如果不需要自定义样式，不要添加 style 标签
</script>
```

#### 自定义样式规范
```vue
<template>
  <div class="custom-component">
    <!-- 复杂布局才使用自定义样式 -->
  </div>
</template>

<style scoped>
/* 仅在 TailwindCSS 无法满足需求时使用 */
.custom-component {
  /* 自定义样式 */
}
</style>
```

### ⚠️ 前端开发约束

**严格遵循以下规范，禁止随意创建文件**:

1. **新增页面**: 必须在 `pages/` 对应模块下创建
2. **新增组件**: 全局组件放 `components/`，页面组件放页面的 `components/`
3. **新增 API**: 按模块在 `api/` 下创建对应文件夹
4. **新增状态**: 在 `stores/` 下按模块创建
5. **新增常量**: 放在 `constants/` 对应文件中
6. **新增工具**: 放在 `utils/` 或 `hooks/` 中
7. **新增类型**: 放在 `types/` 中，按模块组织

### 📦 构建和部署

#### 构建命令
```bash
# 开发环境
npm run dev               # 启动开发服务器

# 构建相关  
npm run build             # 生产构建 (带完整检查)
npm run build-only        # 快速构建 (跳过检查)

# 代码质量
npm run lint              # ESLint 检查并修复
npm run lint:check        # 仅检查不修复
npm run type-check        # TypeScript 类型检查
```

#### 部署流程
1. 前端项目在 `web/` 目录独立开发
2. 运行 `make web-build` 构建前端并复制到 `internal/static/`
3. 运行 `make build` 构建包含前端的 Go 二进制文件
4. 部署单个二进制文件即可

---

## 开发最佳实践

### 🔍 代码质量
1. **类型安全**: 充分利用 TypeScript，避免 `any` 类型
2. **错误处理**: 统一的错误处理机制，友好的错误提示
3. **代码复用**: 提取公共逻辑到 hooks、utils 或 common 中
4. **性能优化**: 合理使用组件懒加载、缓存等技术

### 🧪 测试规范
1. **单元测试**: 核心业务逻辑必须有单元测试
2. **集成测试**: API 接口和数据库操作需要集成测试
3. **E2E 测试**: 关键用户流程需要端到端测试

### 📝 文档规范
1. **代码注释**: 复杂逻辑必须添加注释说明
2. **API 文档**: 所有 API 接口需要文档说明
3. **组件文档**: 公共组件需要使用说明和示例

### 🔧 工具配置
1. **EditorConfig**: 统一编辑器配置
2. **Prettier**: 代码格式化
3. **ESLint**: 代码质量检查
4. **Git Hooks**: 提交前自动检查

---

## 总结

这个全栈模板项目提供了完整的开发规范和最佳实践，旨在让 AI 开发者能够快速上手并保持代码的一致性和质量。

**核心原则**:
- 📏 **统一规范**: 严格遵循命名和文件组织规范
- 🏗️ **模块化**: 清晰的分层架构和模块划分
- 🔒 **类型安全**: 充分利用 TypeScript 类型系统
- 🚀 **高效开发**: 丰富的工具链和自动化脚本
- 📦 **简化部署**: 单文件部署，降低运维复杂度

**请严格按照此规范进行开发，确保项目的可维护性和团队协作效率。** 