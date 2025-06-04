# 用户端前端

基于 Vue 3 + TypeScript + TailwindCSS 的现代化用户端前端应用。

## 🚀 特性

- **现代技术栈**: Vue 3 + TypeScript + TailwindCSS + Vite
- **用户认证系统**: 完整的登录、注册、密码重置流程
- **安全存储**: 加密的本地存储，支持过期管理
- **响应式设计**: 适配各种设备尺寸
- **代码规范**: ESLint + Prettier 确保代码质量
- **开发体验**: 热重载、TypeScript 类型检查

## 📦 技术栈

| 技术 | 版本 | 用途 |
|------|------|------|
| Vue | ^3.5.13 | 前端框架 |
| TypeScript | ~5.8.0 | 类型安全 |
| TailwindCSS | ^3.4.17 | CSS 框架 |
| Vue Router | ^4.5.0 | 路由管理 |
| Pinia | ^3.0.1 | 状态管理 |
| Axios | ^1.9.0 | HTTP 请求 |
| Vite | ^6.2.4 | 构建工具 |

## 🏗️ 项目结构

```
frontend/
├── src/
│   ├── api/                 # API 接口
│   │   └── auth/           # 认证相关 API
│   ├── components/         # 通用组件
│   ├── composables/        # 组合式函数
│   ├── hooks/              # 自定义 hooks
│   │   ├── common/         # 通用 hooks
│   │   └── user/           # 用户相关 hooks
│   ├── pages/              # 页面组件
│   │   ├── About/          # 关于页面
│   │   ├── ForgotPassword/ # 忘记密码
│   │   ├── Home/           # 首页
│   │   ├── Login/          # 登录页面
│   │   ├── NotFound/       # 404 页面
│   │   ├── Profile/        # 个人资料
│   │   ├── Register/       # 注册页面
│   │   └── Settings/       # 设置页面
│   ├── router/             # 路由配置
│   ├── styles/             # 样式文件
│   ├── types/              # TypeScript 类型定义
│   ├── utils/              # 工具函数
│   ├── App.vue             # 主应用组件
│   └── main.ts             # 应用入口
├── public/                 # 静态资源
├── package.json            # 项目配置
├── vite.config.ts          # Vite 配置
├── tailwind.config.js      # TailwindCSS 配置
├── tsconfig.json           # TypeScript 配置
└── README.md               # 项目文档
```

## 🚦 快速开始

### 安装依赖

```bash
npm install
```

### 开发环境

```bash
npm run dev
```

应用将在 http://localhost:4000 启动

### 构建生产版本

```bash
npm run build
```

### 代码检查

```bash
# 运行 ESLint 检查
npm run lint

# 自动修复代码风格问题
npm run lint:fix

# 代码格式化
npm run format
```

## 🔐 认证流程

### 登录

1. 访问 `/login` 页面
2. 输入用户名/邮箱和密码
3. 系统验证后跳转到首页
4. token 自动保存到安全存储

### 注册

1. 访问 `/register` 页面
2. 填写用户信息
3. 发送邮箱验证码
4. 验证通过后完成注册

### 密码重置

1. 在登录页点击"忘记密码"
2. 输入邮箱地址
3. 接收验证码
4. 设置新密码

## 🛡️ 安全特性

- **Token 认证**: 使用 JWT token 进行身份验证
- **安全存储**: 本地数据加密存储，支持过期时间
- **路由守卫**: 自动重定向未认证用户
- **请求拦截**: 自动添加认证头和错误处理

## 📄 页面说明

| 路由 | 页面 | 权限 | 说明 |
|------|------|------|------|
| `/` | 首页 | 公开 | 欢迎页面，展示产品特性 |
| `/login` | 登录 | 游客 | 用户登录，已登录用户会被重定向 |
| `/register` | 注册 | 游客 | 用户注册 |
| `/forgot-password` | 忘记密码 | 游客 | 密码重置 |
| `/profile` | 个人资料 | 需认证 | 查看和编辑个人信息 |
| `/settings` | 设置 | 需认证 | 账户设置和偏好 |
| `/about` | 关于 | 公开 | 项目介绍 |

## 🔧 配置说明

### API 配置

修改 `src/utils/request.ts` 中的 `baseURL` 配置后台 API 地址：

```typescript
const request: AxiosInstance = axios.create({
  baseURL: 'http://localhost:9000/api/v1', // 修改为实际 API 地址
  timeout: 10000,
})
```

### 端口配置

修改 `vite.config.ts` 更改开发服务器端口：

```typescript
export default defineConfig({
  server: {
    port: 4000, // 修改端口号
  },
})
```

## 🎨 主题定制

项目使用 TailwindCSS，可以在 `tailwind.config.js` 中自定义主题：

```javascript
module.exports = {
  theme: {
    extend: {
      colors: {
        primary: {
          // 自定义主色调
        },
      },
    },
  },
}
```

## 📱 响应式支持

- **移动端优先**: 使用 TailwindCSS 的响应式前缀
- **断点设置**: `sm: 640px`, `md: 768px`, `lg: 1024px`, `xl: 1280px`
- **自适应布局**: 所有页面都支持移动端显示

## 🐛 问题解决

### 常见问题

1. **模块找不到错误**
   - 检查 `tsconfig.json` 路径映射配置
   - 确保所有依赖已正确安装

2. **API 请求失败**
   - 检查后台服务是否启动
   - 确认 API 地址配置正确
   - 查看浏览器控制台的网络请求

3. **登录后页面空白**
   - 检查路由配置
   - 确认用户数据格式正确

## 🤝 贡献

1. Fork 项目
2. 创建特性分支
3. 提交更改
4. 推送到分支
5. 创建 Pull Request

## �� 许可证

MIT License
