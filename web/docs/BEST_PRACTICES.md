# Web 微前端最佳实践

本文档提供基于 qiankun 微前端架构的最佳实践指南，帮助你构建可维护、高性能的应用。

## 目录

- [架构决策](#架构决策)
- [状态管理](#状态管理)
- [路由设计](#路由设计)
- [样式隔离](#样式隔离)
- [性能优化](#性能优化)
- [错误处理](#错误处理)
- [测试策略](#测试策略)
- [部署指南](#部署指南)

---

## 架构决策

### 何时使用子应用 vs 路由？

#### 使用子应用的场景 ✅

1. **业务域独立**
   - 用户管理系统、订单系统、内容管理等
   - 团队独立开发和维护
   - 发布周期不同

2. **技术栈差异**
   - 需要使用不同的 React 版本
   - 需要集成 Vue/Angular 等其他框架
   - 需要独立的依赖管理

3. **权限控制**
   - 不同用户角色访问不同子应用
   - 子应用需要独立的认证流程

4. **性能考虑**
   - 功能模块体积大（> 500KB）
   - 可以按需加载

#### 使用路由的场景 ✅

1. **功能紧密相关**
   - 同一业务流程的不同页面
   - 共享大量组件和状态
   - 需要频繁跳转

2. **性能敏感**
   - 页面切换需要极快响应
   - 避免子应用加载开销

3. **开发团队统一**
   - 同一团队维护
   - 统一的发布周期

#### 示例对比

```
❌ 不推荐：为每个页面创建子应用
├── packages/user-list/      # 过度拆分
├── packages/user-detail/
└── packages/user-edit/

✅ 推荐：按业务域拆分
├── packages/user/           # 用户管理子应用
│   └── pages/
│       ├── UserList
│       ├── UserDetail
│       └── UserEdit
└── packages/order/          # 订单管理子应用
```

---

## 状态管理

### 主应用状态

主应用负责管理**全局共享状态**：

```typescript
// packages/main/src/store/global.ts
import { create } from 'zustand'

interface GlobalState {
  theme: 'light' | 'dark'
  locale: string
  setTheme: (theme: 'light' | 'dark') => void
  setLocale: (locale: string) => void
}

export const useGlobalStore = create<GlobalState>((set) => ({
  theme: 'light',
  locale: 'zh-CN',
  setTheme: (theme) => set({ theme }),
  setLocale: (locale) => set({ locale }),
}))
```

### 子应用状态

子应用维护**局部业务状态**：

```typescript
// packages/user/src/store/user.ts
import { create } from 'zustand'

interface UserState {
  users: User[]
  loading: boolean
  fetchUsers: () => Promise<void>
}

export const useUserStore = create<UserState>((set) => ({
  users: [],
  loading: false,
  fetchUsers: async () => {
    set({ loading: true })
    const users = await userService.list()
    set({ users, loading: false })
  },
}))
```

### 跨应用通信

#### 1. 使用认证状态（推荐）

```typescript
// 所有应用都可以使用
import { useAuth } from '@app/auth'

const { user, isAuthenticated, login, logout } = useAuth()
```

#### 2. 使用 CustomEvent

```typescript
// 发送事件
window.dispatchEvent(
  new CustomEvent('app:user-updated', {
    detail: { userId: 123 },
  })
)

// 监听事件
useEffect(() => {
  const handler = (e: CustomEvent) => {
    console.log('User updated:', e.detail.userId)
  }
  window.addEventListener('app:user-updated', handler)
  return () => window.removeEventListener('app:user-updated', handler)
}, [])
```

#### 3. 使用 localStorage（慎用）

```typescript
import { SecureStorage } from '@app/shared'

// 设置
SecureStorage.setItem('shared-data', data)

// 监听变化
window.addEventListener('storage', (e) => {
  if (e.key === 'shared-data') {
    console.log('Data changed:', e.newValue)
  }
})
```

### 状态管理建议

| 状态类型 | 推荐方案 | 示例 |
|---------|---------|------|
| 用户认证 | `@app/auth` | 登录状态、用户信息 |
| 全局配置 | 主应用 Zustand | 主题、语言、布局 |
| 业务数据 | 子应用 Zustand/React Query | 用户列表、订单数据 |
| 表单状态 | React Hook Form | 表单输入、验证 |
| UI 状态 | useState | 模态框、展开状态 |

---

## 路由设计

### 路由层级规划

```
/                         主应用首页
├── /user                 用户管理子应用
│   ├── /user/list        用户列表
│   ├── /user/:id         用户详情
│   └── /user/:id/edit    编辑用户
├── /order                订单管理子应用
│   ├── /order/list
│   └── /order/:id
└── /analytics            数据分析子应用
```

### 子应用路由配置

```typescript
// packages/user/src/router/index.tsx
import { createBrowserRouter } from 'react-router-dom'

export const router = createBrowserRouter(
  [
    {
      path: '/',
      element: <Layout />,
      children: [
        { index: true, element: <Navigate to="/list" /> },
        { path: 'list', element: <UserList /> },
        { path: ':id', element: <UserDetail /> },
        { path: ':id/edit', element: <UserEdit /> },
      ],
    },
  ],
  { basename: '/user' } // 重要：匹配 activeRule
)
```

### 路由跳转

#### 子应用内跳转

```typescript
import { useNavigate } from 'react-router-dom'

const navigate = useNavigate()
navigate('/list') // 相对路径
navigate('/user/123') // 绝对路径
```

#### 跨子应用跳转

```typescript
// 使用原生 API
window.location.href = '/order/123'

// 或使用主应用的导航方法
window.dispatchEvent(
  new CustomEvent('app:navigate', {
    detail: { path: '/order/123' },
  })
)
```

---

## 样式隔离

### 主应用样式

主应用使用 `important: true`：

```javascript
// packages/main/tailwind.config.js
export default {
  important: true, // 防止被子应用覆盖
  content: ['./src/**/*.{ts,tsx}'],
}
```

### 子应用样式

子应用禁用 `preflight`：

```javascript
// packages/app/tailwind.config.js
export default {
  corePlugins: { preflight: false }, // 避免重置样式冲突
  content: ['./src/**/*.{ts,tsx}'],
}
```

### CSS 模块化

推荐使用 CSS Modules 或 CSS-in-JS：

```typescript
// ✅ 推荐：CSS Modules
import styles from './Button.module.css'

export function Button() {
  return <button className={styles.button}>Click</button>
}

// ✅ 推荐：Tailwind（已配置隔离）
export function Button() {
  return <button className="bg-primary-500 text-white">Click</button>
}

// ❌ 避免：全局样式
import './Button.css' // 可能污染其他应用
```

### qiankun 样式隔离

```typescript
// packages/main/src/micro-apps.ts
start({
  sandbox: {
    experimentalStyleIsolation: true, // 已启用
  },
})
```

---

## 性能优化

### 1. 代码分割

```typescript
// 路由懒加载
const UserList = lazy(() => import('./pages/UserList'))
const UserDetail = lazy(() => import('./pages/UserDetail'))

// 组件懒加载
const HeavyChart = lazy(() => import('./components/HeavyChart'))

<Suspense fallback={<Loading />}>
  <HeavyChart />
</Suspense>
```

### 2. 预加载策略

```typescript
// packages/main/src/micro-apps.ts
start({
  prefetch: 'all', // 预加载所有子应用（默认关闭）
})

// 或选择性预加载
start({
  prefetch: (apps) => apps.filter((app) => app.name !== 'heavy-app'),
})
```

### 3. 构建优化

```typescript
// vite.config.ts
export default {
  build: {
    rollupOptions: {
      output: {
        manualChunks: {
          'react-vendor': ['react', 'react-dom'],
          'router': ['react-router-dom'],
        },
      },
    },
  },
}
```

### 4. 图片优化

```typescript
// 使用 WebP 格式
<img src="/images/hero.webp" alt="Hero" loading="lazy" />

// 响应式图片
<picture>
  <source srcset="/images/hero-mobile.webp" media="(max-width: 768px)" />
  <img src="/images/hero.webp" alt="Hero" />
</picture>
```

### 5. API 请求优化

```typescript
// 使用 React Query 缓存
import { useQuery } from '@tanstack/react-query'

const { data, isLoading } = useQuery({
  queryKey: ['users', page],
  queryFn: () => userService.list(page),
  staleTime: 5 * 60 * 1000, // 5 分钟内不重新请求
})
```

---

## 错误处理

### React Error Boundary

```typescript
// shared/components/ErrorBoundary.tsx
import { Component, ReactNode } from 'react'

interface Props {
  children: ReactNode
  fallback?: ReactNode
}

interface State {
  hasError: boolean
  error?: Error
}

export class ErrorBoundary extends Component<Props, State> {
  state: State = { hasError: false }

  static getDerivedStateFromError(error: Error): State {
    return { hasError: true, error }
  }

  componentDidCatch(error: Error, errorInfo: any) {
    console.error('Error caught:', error, errorInfo)
    // 上报错误到监控系统
  }

  render() {
    if (this.state.hasError) {
      return this.props.fallback || <div>出错了</div>
    }
    return this.props.children
  }
}

// 使用
<ErrorBoundary>
  <App />
</ErrorBoundary>
```

### 子应用加载错误

qiankun 已配置友好的错误提示：

```typescript
// packages/main/src/micro-apps.ts
loadMicroApp({
  name: 'app',
  entry: '//localhost:3001',
  container: '#container',
  props: {},
}).catch((error) => {
  console.error('子应用加载失败:', error)
  // 显示友好提示
})
```

### API 错误处理

```typescript
// HttpClient 已自动处理常见错误
try {
  await http.get('/users')
} catch (error) {
  // 401, 403, 500 等已自动处理
  // 只需处理业务逻辑错误
}
```

---

## 测试策略

### 单元测试

```typescript
// packages/app/src/components/Button.test.tsx
import { render, screen } from '@testing-library/react'
import { Button } from './Button'

describe('Button', () => {
  it('renders correctly', () => {
    render(<Button>Click me</Button>)
    expect(screen.getByText('Click me')).toBeInTheDocument()
  })

  it('handles click events', () => {
    const onClick = vi.fn()
    render(<Button onClick={onClick}>Click</Button>)
    screen.getByText('Click').click()
    expect(onClick).toHaveBeenCalled()
  })
})
```

### 集成测试

```typescript
// packages/app/src/pages/UserList.test.tsx
import { render, screen, waitFor } from '@testing-library/react'
import { UserList } from './UserList'
import { userService } from '@/services/user.service'

vi.mock('@/services/user.service')

describe('UserList', () => {
  it('loads and displays users', async () => {
    vi.mocked(userService.list).mockResolvedValue({
      items: [{ id: 1, name: 'Alice' }],
      total: 1,
    })

    render(<UserList />)

    await waitFor(() => {
      expect(screen.getByText('Alice')).toBeInTheDocument()
    })
  })
})
```

### E2E 测试

```typescript
// tests/e2e/user-flow.spec.ts (Playwright)
import { test, expect } from '@playwright/test'

test('user can login and view user list', async ({ page }) => {
  await page.goto('http://localhost:3000')
  await page.fill('[name=email]', 'test@example.com')
  await page.fill('[name=password]', 'password')
  await page.click('button[type=submit]')

  await expect(page).toHaveURL('/user/list')
  await expect(page.locator('.user-item')).toHaveCount(10)
})
```

---

## 部署指南

### 构建顺序

```bash
# 1. 构建共享包
pnpm --filter @app/shared build

# 2. 构建核心包
pnpm --filter @app/core build

# 3. 构建子应用
pnpm --filter @app/app build
pnpm --filter @app/user build

# 4. 构建主应用
pnpm --filter @app/main build
```

### 静态资源路径

确保子应用的 `base` 配置正确：

```typescript
// packages/app/vite.config.ts
export default defineConfig(({ mode }) => ({
  base: mode === 'production' ? '/subapps/app/' : '/',
}))
```

### Nginx 配置

```nginx
server {
  listen 80;
  server_name example.com;

  # 主应用
  location / {
    root /var/www/main;
    try_files $uri $uri/ /index.html;
  }

  # 子应用
  location /subapps/ {
    root /var/www;
    try_files $uri $uri/ =404;
  }

  # API 代理
  location /api/ {
    proxy_pass http://localhost:9000;
  }
}
```

### Docker 部署

```dockerfile
# Dockerfile
FROM node:20-alpine AS builder

WORKDIR /app
COPY package.json pnpm-lock.yaml ./
RUN npm install -g pnpm && pnpm install

COPY . .
RUN pnpm build

FROM nginx:alpine
COPY --from=builder /app/packages/main/dist /usr/share/nginx/html
COPY --from=builder /app/packages/app/dist /usr/share/nginx/html/subapps/app
COPY nginx.conf /etc/nginx/conf.d/default.conf
```

---

## 常见问题

### Q: 子应用路由刷新 404？

确保子应用的 router 配置了正确的 basename：

```typescript
createBrowserRouter(routes, { basename: '/user' })
```

### Q: 样式冲突？

1. 主应用使用 `important: true`
2. 子应用禁用 `preflight`
3. 使用 CSS Modules 或 Tailwind

### Q: 子应用加载慢？

1. 启用 `prefetch`
2. 优化构建产物（code splitting）
3. 使用 CDN

### Q: 认证状态不同步？

使用 `@app/auth` 提供的 `useAuth` hook，它会自动监听 storage 事件。

---

## 相关资源

- [qiankun 官方文档](https://qiankun.umijs.org/)
- [HTTP Client 使用指南](./HTTP_CLIENT.md)
- [故障排除指南](../TROUBLESHOOTING.md)
- [项目 README](../README.md)
