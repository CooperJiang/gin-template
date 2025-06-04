# 🎨 Gin Template 前端项目

基于 Vue 3 + TypeScript + TailwindCSS 构建的现代化前端应用，与 Gin 后端完美集成。

## 🚀 快速开始

### 环境要求
- Node.js >= 18.0.0
- npm >= 9.0.0

### 安装和启动

```bash
# 安装依赖
npm install

# 启动开发服务器
npm run dev

# 或使用便捷脚本
./start.sh
```

访问地址: http://localhost:3000

## 🛠️ 技术栈

| 技术 | 版本 | 说明 |
|------|------|------|
| **Vue 3** | ^3.5.13 | 渐进式JavaScript框架 |
| **TypeScript** | ^5.6.3 | 类型安全的JavaScript |
| **Vite** | ^6.0.1 | 下一代前端构建工具 |
| **Pinia** | ^2.2.6 | Vue官方状态管理库 |
| **Vue Router** | ^4.5.0 | Vue官方路由管理器 |
| **TailwindCSS** | ^4.0.0 | 实用优先的CSS框架 |
| **Axios** | ^1.7.9 | HTTP客户端 |
| **VueUse** | ^11.3.0 | Vue组合式API工具集 |
| **Heroicons** | ^2.2.0 | 精美的SVG图标库 |

## 📁 项目结构

```
web/
├── public/                    # 静态资源
├── src/
│   ├── api/                   # API接口层
│   │   ├── auth.ts           # 认证相关API
│   │   └── user.ts           # 用户相关API
│   ├── assets/               # 资源文件
│   │   ├── main.css          # 主样式文件
│   │   └── base.css          # 基础样式
│   ├── components/           # 组件
│   │   └── common/           # 通用组件
│   │       ├── AppNotification.vue  # 通知组件
│   │       └── AppLoading.vue       # 加载组件
│   ├── router/               # 路由配置
│   │   └── index.ts          # 路由定义
│   ├── stores/               # Pinia状态管理
│   │   ├── auth.ts           # 认证状态
│   │   ├── notification.ts   # 通知状态
│   │   └── user.ts           # 用户状态
│   ├── types/                # TypeScript类型定义
│   │   └── index.ts          # 全局类型
│   ├── utils/                # 工具函数
│   │   └── request.ts        # HTTP请求封装
│   ├── views/                # 页面组件
│   │   ├── LoginView.vue     # 登录页面
│   │   ├── RegisterView.vue  # 注册页面
│   │   ├── DashboardView.vue # 仪表板
│   │   └── ...               # 其他页面
│   ├── App.vue               # 根组件
│   └── main.ts               # 应用入口
├── index.html                # HTML模板
├── package.json              # 项目配置
├── vite.config.ts            # Vite配置
├── tailwind.config.js        # TailwindCSS配置
├── postcss.config.js         # PostCSS配置
└── tsconfig.json             # TypeScript配置
```

## 🎯 核心功能

### 🔐 认证系统
- JWT Token 认证
- 自动token刷新
- 路由权限控制
- 登录状态持久化

### 📊 状态管理
- Pinia 状态管理
- 模块化store设计
- 类型安全的状态操作
- 持久化支持

### 🌐 HTTP请求
- Axios 请求封装
- 自动错误处理
- 请求/响应拦截器
- 统一API接口

### 🎨 UI组件
- TailwindCSS 样式系统
- 响应式设计
- 通用组件库
- 主题定制

### 🔔 通知系统
- 全局通知组件
- 多种通知类型
- 自动消失
- 可自定义样式

## 🚀 开发命令

```bash
# 开发
npm run dev              # 启动开发服务器
npm run dev:host         # 启动开发服务器（允许外部访问）

# 构建
npm run build            # 构建生产版本
npm run preview          # 预览生产构建

# 代码质量
npm run lint             # ESLint检查
npm run lint:fix         # 自动修复ESLint错误
npm run format           # Prettier格式化
npm run type-check       # TypeScript类型检查

# 测试
npm run test             # 运行测试
npm run test:coverage    # 运行测试并生成覆盖率报告
```

## 🔧 配置说明

### 环境变量

创建 `.env` 文件配置环境变量：

```bash
# API基础URL
VITE_API_BASE_URL=http://localhost:8080/api/v1

# 应用标题
VITE_APP_TITLE=Gin Template

# 是否启用调试模式
VITE_DEBUG=true
```

### TailwindCSS配置

项目使用 TailwindCSS v4，配置文件 `tailwind.config.js`：

```javascript
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: {
          // 自定义主色调
        },
      },
    },
  },
  plugins: [],
}
```

### TypeScript配置

严格的TypeScript配置，确保类型安全：

```json
{
  "compilerOptions": {
    "strict": true,
    "noImplicitAny": true,
    "strictNullChecks": true,
    // ...其他配置
  }
}
```

## 📝 开发规范

### 代码风格
- 使用 ESLint + Prettier 统一代码风格
- 遵循 Vue 3 Composition API 最佳实践
- TypeScript 严格模式
- 组件命名使用 PascalCase

### 文件命名
- 组件文件：`PascalCase.vue`
- 工具文件：`camelCase.ts`
- 页面文件：`PascalCaseView.vue`
- 类型文件：`camelCase.ts`

### 提交规范
```bash
feat: 新功能
fix: 修复bug
docs: 文档更新
style: 代码格式调整
refactor: 代码重构
test: 测试相关
chore: 构建工具或辅助工具的变动
```

## 🔗 API集成

### 请求封装

```typescript
// utils/request.ts
import axios from 'axios'

const request = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
  timeout: 10000,
})

// 请求拦截器
request.interceptors.request.use(config => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

// 响应拦截器
request.interceptors.response.use(
  response => response.data.data,
  error => {
    // 统一错误处理
    return Promise.reject(error)
  }
)
```

### API接口定义

```typescript
// api/auth.ts
export const authApi = {
  login(data: LoginRequest): Promise<LoginResponse> {
    return ApiClient.post('/auth/login', data)
  },
  
  getCurrentUser(): Promise<User> {
    return ApiClient.get('/auth/profile')
  },
  
  // ...其他接口
}
```

## 🎨 组件开发

### 组件模板

```vue
<template>
  <div class="component-wrapper">
    <!-- 组件内容 -->
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

// Props定义
interface Props {
  title: string
  visible?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  visible: false
})

// Emits定义
interface Emits {
  close: []
  confirm: [value: string]
}

const emit = defineEmits<Emits>()

// 响应式数据
const loading = ref(false)

// 计算属性
const isVisible = computed(() => props.visible && !loading.value)

// 方法
const handleClose = () => {
  emit('close')
}
</script>

<style scoped>
.component-wrapper {
  @apply p-4 bg-white rounded-lg shadow;
}
</style>
```

## 🚀 部署

### 构建生产版本

```bash
npm run build
```

构建产物在 `dist/` 目录下。

### 部署到静态服务器

```bash
# 使用 nginx
cp -r dist/* /var/www/html/

# 使用 Apache
cp -r dist/* /var/www/html/

# 使用 CDN
# 上传 dist/ 目录到 CDN
```

### Docker部署

```dockerfile
FROM nginx:alpine
COPY dist/ /usr/share/nginx/html/
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
```

## 🐛 常见问题

### 1. 开发服务器启动失败
```bash
# 清除缓存
rm -rf node_modules package-lock.json
npm install
```

### 2. TypeScript类型错误
```bash
# 重新生成类型
npm run type-check
```

### 3. 样式不生效
```bash
# 检查TailwindCSS配置
npm run build
```

## 📚 学习资源

- [Vue 3 官方文档](https://vuejs.org/)
- [TypeScript 官方文档](https://www.typescriptlang.org/)
- [TailwindCSS 官方文档](https://tailwindcss.com/)
- [Pinia 官方文档](https://pinia.vuejs.org/)
- [Vite 官方文档](https://vitejs.dev/)

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## �� 许可证

MIT License
