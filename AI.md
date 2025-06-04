# AI 开发规范文档

> 专为AI助手（Cursor、Claude等）提供的开发规范，确保严格遵循项目架构和代码规范。

## 📋 项目概述

**技术栈**: Go + Gin + Vue3 + TypeScript + SQLite/PostgreSQL + Redis  
**架构**: 前后端分离开发，构建时合并为单个二进制文件  
**部署**: 前端打包后嵌入Go程序，最终生成单文件部署

---

## ⚠️ 核心约束 (必须严格遵守)

### 🚫 绝对禁止的操作
1. **禁止随意创建新文件或目录** - 必须按照既定目录结构
2. **禁止使用自增ID** - 所有数据模型必须使用UUID主键
3. **禁止跳过分层架构** - 必须按Controller→Service→Repository→Model分层
4. **禁止在Controller中写业务逻辑** - Controller只处理HTTP请求
5. **禁止跳过参数验证** - 所有输入参数必须验证
6. **禁止硬编码配置信息** - 必须使用配置文件或常量

### ✅ 必须遵守的规则
1. **UUID主键** - 所有Model必须嵌入BaseModel使用UUID
2. **统一命名** - 全项目采用驼峰命名法(camelCase)
3. **分层调用** - 严格按分层架构调用，不可跨层
4. **统一响应** - 使用errors.ResponseSuccess/ResponseError
5. **参数校验** - 使用common.ValidateRequest统一校验
6. **错误处理** - 使用统一错误处理机制

---

## 🔧 后端开发规范

### 📁 目录结构规则

#### 1. 目录组织规范
```
cmd/main.go                    # ✅ 程序入口，唯一启动文件
internal/
├── controllers/模块名/        # ✅ HTTP请求处理层
├── services/模块名/          # ✅ 业务逻辑层  
├── repositories/模块名/      # ✅ 数据访问层
├── models/                   # ✅ 数据模型定义
├── dto/request/             # ✅ 请求参数对象
├── dto/response/            # ✅ 响应数据对象
├── routes/                  # ✅ 路由定义
├── middleware/              # ✅ 中间件
└── migrations/              # ✅ 数据库迁移文件
pkg/
├── common/                  # ✅ 通用工具函数
├── config/                  # ✅ 配置管理
├── database/                # ✅ 数据库连接
├── logger/                  # ✅ 日志系统
├── cache/                   # ✅ Redis缓存
├── email/                   # ✅ 邮件服务
├── errors/                  # ✅ 错误处理
├── utils/                   # ✅ 工具函数
└── constants/               # ✅ 常量定义
```

#### 2. 文件创建规则
- **新增Model**: 只能在`internal/models/`创建，必须嵌入BaseModel
- **新增API**: 按模块在`controllers/服务名/`、`services/服务名/`、`repositories/服务名/`创建
- **新增常量**: 只能在`pkg/constants/`对应模块文件中添加
- **新增工具**: 只能在`pkg/utils/`或`pkg/common/`中添加
- **新增中间件**: 只能在`internal/middleware/`中添加
- **新增路由**: 只能在`internal/routes/`对应模块文件中添加

### 🏛️ 分层架构规则

#### 3. Controller层规则
- **职责**: 仅处理HTTP请求、参数校验、响应格式化
- **禁止**: 不得包含任何业务逻辑
- **调用**: 只能调用Service层
- **响应**: 必须使用`errors.ResponseSuccess(c, data, "消息")`或`errors.ResponseError(c, err)`

```go
func (ctrl *UserController) CreateUser(c *gin.Context) {
    // ✅ 正确：参数校验
    req, err := common.ValidateRequest[request.CreateUserRequest](c)
    if err != nil {
        errors.HandleError(c, err)
        return
    }
    
    // ✅ 正确：调用Service层
    user, err := ctrl.userService.CreateUser(req.Username, req.Email, req.Password)
    if err != nil {
        errors.ResponseError(c, err)
        return
    }
    
    // ✅ 正确：统一响应格式
    errors.ResponseSuccess(c, user, "创建成功")
}
```

#### 4. Service层规则
- **职责**: 业务逻辑处理、事务管理、数据组装
- **调用**: 只能调用Repository层和其他Service
- **事务**: 在此层管理数据库事务
- **验证**: 包含业务规则验证

```go
func (s *userService) CreateUser(username, email, password string) (*models.User, error) {
    // ✅ 正确：业务逻辑验证
    if exists, _ := s.userRepo.ExistsByEmail(email); exists {
        return nil, errors.New(errors.CodeUserExists, "邮箱已存在")
    }
    
    // ✅ 正确：数据处理
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    
    // ✅ 正确：调用Repository层
    user := &models.User{
        Username: username,
        Email:    email,
        Password: string(hashedPassword),
    }
    
    return s.userRepo.Create(context.Background(), user)
}
```

#### 5. Repository层规则
- **职责**: 仅数据库CRUD操作
- **继承**: 必须嵌入BaseRepository获得通用CRUD
- **查询**: 只包含数据访问逻辑，不含业务逻辑
- **错误**: 统一错误处理和转换

```go
type UserRepository interface {
    repositories.BaseRepository[models.User]  // ✅ 必须继承
    GetByEmail(ctx context.Context, email string) (*models.User, error)
    ExistsByEmail(ctx context.Context, email string) (bool, error)
}
```

### 🆔 数据模型规则

#### 6. Model定义规则
- **主键**: 必须使用UUID，禁止自增ID
- **嵌入**: 必须嵌入BaseModel
- **表名**: 必须定义TableName()方法
- **钩子**: 实现BeforeCreate等必要钩子

```go
type User struct {
    BaseModel                                    // ✅ 必须嵌入
    Username string `gorm:"size:50;not null;uniqueIndex" json:"username"`
    Password string `gorm:"size:100;not null" json:"-"`
    Email    string `gorm:"size:100;uniqueIndex" json:"email"`
    Status   int    `gorm:"default:1" json:"status"`
}

func (User) TableName() string {           // ✅ 必须定义
    return "user"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {  // ✅ 必须实现
    return u.BaseModel.BeforeCreate(tx)
}
```

### 📝 数据传输规则

#### 7. DTO定义规则
- **位置**: Request放`internal/dto/request/`，Response放`internal/dto/response/`
- **校验**: Request必须包含binding和validate标签
- **消息**: Request必须实现GetValidationMessages()方法
- **转换**: Response必须实现FromModel()方法

```go
// ✅ Request DTO
type CreateUserRequest struct {
    dto.BaseRequest
    Username string `json:"username" binding:"required,min=2,max=20"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6,max=20"`
}

func (r *CreateUserRequest) GetValidationMessages() map[string]string {
    return map[string]string{
        "Username.required": "用户名不能为空",
        "Email.email":       "请输入有效的邮箱地址",
    }
}

// ✅ Response DTO  
type UserInfo struct {
    dto.BaseResponse
    Username string `json:"username"`
    Email    string `json:"email"`
}

func (u *UserInfo) FromModel(user models.User) *UserInfo {
    return &UserInfo{
        BaseResponse: dto.BaseResponse{
            ID:        user.ID.String(),    // ✅ UUID转字符串
            CreatedAt: time.Time(user.CreatedAt),
            UpdatedAt: time.Time(user.UpdatedAt),
        },
        Username: user.Username,
        Email:    user.Email,
    }
}
```

### 🛣️ 路由规则

#### 8. 路由定义规则
- **组织**: 一个模块一个路由文件
- **分组**: 使用gin.Group分组管理
- **中间件**: 在路由级别应用中间件
- **注册**: 在main.go中统一注册

```go
// ✅ 路由文件: internal/routes/userRoutes.go
func RegisterUserRoutes(r *gin.Engine, userController *controllers.UserController) {
    api := r.Group("/api/v1")
    {
        user := api.Group("/user")
        {
            user.POST("/login", userController.Login)
            user.POST("/register", userController.Register)
            
            authUser := user.Use(middleware.AuthMiddleware())  // ✅ 中间件应用
            {
                authUser.GET("/info", userController.GetUserInfo)
            }
        }
    }
}
```

### 🔧 开发工具规则

#### 9. 工具使用规则
- **模块生成**: 使用`make new-module name=模块名`创建新模块
- **数据库**: 使用`make migrate`执行迁移
- **开发**: 使用`make dev`启动开发服务器
- **测试**: 使用`make test`运行测试
- **检查**: 使用`make full-check`完整检查

---

## 🎨 前端开发规范

### 📁 目录结构规则

#### 10. 前端目录组织
```
web/src/
├── api/模块名/              # ✅ API接口封装
├── components/组件名/       # ✅ 全局组件(已注册,无需引入)
├── composables/            # ✅ 组合式函数
├── constants/              # ✅ 常量定义
├── hooks/分类/             # ✅ Vue Hooks
├── layouts/                # ✅ 布局组件
├── pages/页面名/           # ✅ 页面组件
├── router/                 # ✅ 路由配置
├── stores/模块名/          # ✅ 状态管理
├── styles/                 # ✅ 样式文件
├── types/                  # ✅ 类型定义
└── utils/                  # ✅ 工具函数
```

#### 11. 文件创建规则
- **新增页面**: 必须在`pages/`对应模块下创建，主文件名为`index.vue`
- **新增组件**: 全局组件放`components/`，页面组件放页面的`components/`
- **新增API**: 按模块在`api/`下创建对应文件夹
- **新增状态**: 在`stores/`下按模块创建
- **新增常量**: 放在`constants/`对应文件中
- **新增类型**: 放在`types/`中，按模块组织

### 🧩 组件开发规则

#### 12. 全局组件规则
- **结构**: 每个组件必须包含`index.vue`、`index.ts`、`types.ts`最少三个文件 复杂的可以拆分为多个
- **命名**: 组件名必须有`Global`前缀，如`GlobalButton`
- **注册**: 已全局注册，页面中直接使用无需引入
- **优先**: 开发功能前先检查是否有现成全局组件

```vue
<!-- ✅ 全局组件定义 -->
<script setup lang="ts">
import type { ButtonProps } from './types'

defineOptions({
  name: 'GlobalButton',  // ✅ 必须有Global前缀
})

withDefaults(defineProps<ButtonProps>(), {
  type: 'primary',
  size: 'medium',
})
</script>
```

#### 13. 页面组件规则
- **组织**: 一个页面一个文件夹，主文件命名为`index.vue`
- **嵌套**: 多级路由对应多层级目录结构
- **私有组件**: 复杂页面在同目录创建`components/`文件夹
- **命名**: 组件必须定义`name`属性

```vue
<!-- ✅ 页面组件模板 -->
<template>
  <div class="page-container">
    <!-- 页面内容 -->
  </div>
</template>

<script setup lang="ts">
defineOptions({
  name: 'LoginPage',  // ✅ 必须定义组件名
})
</script>
```

### 🛣️ 路由规则

#### 14. 路由配置规则
- **Meta信息**: 每个路由必须包含完整meta信息
- **权限**: 使用`requiresAuth`和`roles`控制访问权限
- **布局**: 通过`layout`指定使用的布局组件
- **拆分**: 大型项目按模块拆分路由文件

```javascript
// ✅ 路由配置模板
{
  path: '/users',
  name: 'users',
  component: () => import('@/pages/System/User/index.vue'),
  meta: {
    title: '用户管理',        // ✅ 必须：页面标题
    requiresAuth: true,      // ✅ 必须：是否需要认证
    layout: 'admin',         // ✅ 必须：使用的布局
    icon: 'users',           // ✅ 必须：菜单图标
    hidden: false,           // ✅ 必须：是否隐藏
    roles: ['admin'],        // ✅ 必须：允许的角色
  }
}
```

### 🗃️ 状态管理规则

#### 15. Store定义规则
- **组织**: 一个模块一个文件夹，包含`index.ts`和`types.ts`
- **类型**: 必须定义完整的State接口
- **命名**: store名称使用`useXxxStore`格式
- **结构**: 包含state、getters、actions三部分

```typescript
// ✅ Store定义模板
export const useUserStore = defineStore('user', {
  state: (): UserState => ({
    userInfo: null,
    token: null,
  }),
  
  getters: {
    isLoggedIn: (state) => !!state.token,
  },
  
  actions: {
    async login(credentials: LoginRequest) {
      // 异步操作
    }
  }
})
```

### 🌐 API接口规则

#### 16. API封装规则
- **组织**: 按模块分文件夹，统一导出
- **命名**: API方法使用动词+名词格式
- **类型**: 所有请求和响应必须有类型定义
- **错误**: 统一错误处理机制

```typescript
// ✅ API封装模板
export const userApi = {
  login: (data: LoginRequest) => 
    request.post<LoginResponse>('/user/login', data),
    
  getUserInfo: () => 
    request.get<User>('/user/info'),
    
  updateProfile: (data: UpdateProfileRequest) =>
    request.put<User>('/user/profile', data),
}
```

### 🎨 样式规则

#### 17. 样式开发规则
- **优先级**: 优先使用TailwindCSS类名
- **自定义**: 仅在TailwindCSS无法满足时使用自定义样式
- **作用域**: 自定义样式必须使用`scoped`
- **组织**: 复杂样式抽取到独立文件

```vue
<!-- ✅ 样式使用模板 -->
<template>
  <!-- 优先使用 TailwindCSS -->
  <div class="flex items-center justify-between p-4 bg-white rounded-lg shadow">
    <h1 class="text-xl font-bold text-gray-900">标题</h1>
  </div>
</template>

<style scoped>
/* 仅在必要时使用自定义样式 */
.custom-component {
  /* 自定义样式 */
}
</style>
```

---

## 📋 开发流程规则

### 🚀 标准开发步骤
1. **分析需求** - 确定涉及的模块和功能
2. **检查组件** - 优先使用现有全局组件和工具
3. **后端开发** - 按Model→Repository→Service→Controller顺序
4. **前端开发** - 按API→Store→Page→Component顺序  
5. **测试验证** - 运行`make test`和前端测试
6. **代码检查** - 运行`make full-check`确保质量

### 🔍 质量检查规则
- **类型安全**: 充分利用TypeScript，避免`any`类型
- **错误处理**: 统一错误处理机制，友好错误提示
- **代码复用**: 提取公共逻辑到hooks、utils或common中
- **性能优化**: 合理使用组件懒加载、缓存等技术

---

## ⚡ 快捷开发命令

### 后端命令
```bash
make new-module name=product  # 创建新模块脚手架
make dev                      # 启动开发服务器
make build                    # 构建应用程序
make test                     # 运行测试
make full-check              # 完整代码检查
make migrate                 # 执行数据库迁移
```

### 前端命令
```bash
make web-dev                 # 启动前端开发服务器
make web-build               # 构建前端
make web-lint                # 前端代码检查
```

### 全栈命令
```bash
make fullstack-build         # 构建完整应用
make fullstack-dev           # 并行启动前后端开发
make fullstack-clean         # 清理所有构建文件
```

---

## 📋 检查清单

开发完成后请确认：
- [ ] 严格按照目录结构组织文件
- [ ] 后端使用分层架构，前端使用模块化结构
- [ ] 所有Model使用UUID主键
- [ ] 统一使用驼峰命名法
- [ ] API有完整的类型定义
- [ ] 路由有完整的meta信息
- [ ] 优先使用现有组件和工具
- [ ] 通过所有代码质量检查
- [ ] 有适当的错误处理
- [ ] 代码有必要的注释说明

**严格遵循以上规范，确保代码质量和项目一致性！** 