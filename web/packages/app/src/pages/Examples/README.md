# 示例页面说明

本目录包含完整的示例页面，展示微前端应用的最佳实践。

## 页面列表

### 1. Dashboard（仪表板）

**文件**: `Dashboard.tsx`

**展示内容**:
- 卡片布局
- 数据统计展示
- 并发 API 请求
- 加载状态处理

**使用的组件**:
- `Card`, `Loading`, `Button` from `@app/core`
- `createHttpClient` from `@app/core`

**路由**: `/examples/dashboard`

---

### 2. UserList（用户列表）

**文件**: `UserList.tsx`

**展示内容**:
- 表格展示
- 分页功能
- 搜索过滤
- 加载和空状态

**使用的组件**:
- `Card`, `Button`, `Input`, `Loading` from `@app/core`
- `PaginationResponse` from `@app/shared/types`

**路由**: `/examples/user/list`

---

### 3. UserForm（用户表单）

**文件**: `UserForm.tsx`

**展示内容**:
- 表单输入
- 实时验证
- 错误处理
- 提交状态

**使用的组件**:
- `Card`, `Button`, `Input` from `@app/core`
- `message` from `@app/core`

**路由**: `/examples/user/new`

---

## 如何使用

### 1. 添加路由

在 `src/router/index.tsx` 中添加示例路由：

```typescript
import { Dashboard, UserList, UserForm } from '@/pages/Examples'

{
  path: 'examples',
  children: [
    { path: 'dashboard', element: <Dashboard /> },
    { path: 'user/list', element: <UserList /> },
    { path: 'user/new', element: <UserForm /> },
  ],
}
```

### 2. 添加导航链接

```typescript
<nav>
  <Link to="/examples/dashboard">仪表板</Link>
  <Link to="/examples/user/list">用户列表</Link>
  <Link to="/examples/user/new">添加用户</Link>
</nav>
```

### 3. Mock 数据（开发环境）

如果后端 API 未就绪，可以使用 MSW (Mock Service Worker) 模拟数据：

```typescript
// src/mocks/handlers.ts
import { http, HttpResponse } from 'msw'

export const handlers = [
  http.get('/api/v1/users', () => {
    return HttpResponse.json({
      code: 200,
      message: 'success',
      data: {
        items: [
          {
            id: 1,
            name: 'Alice',
            email: 'alice@example.com',
            role: 'admin',
            createdAt: '2024-01-01T00:00:00Z',
          },
        ],
        total: 1,
        page: 1,
        pageSize: 10,
        totalPages: 1,
      },
    })
  }),
]
```

## 最佳实践亮点

### ✅ 类型安全

所有 API 响应都有完整的 TypeScript 类型定义：

```typescript
interface User {
  id: number
  name: string
  email: string
  role: string
  createdAt: string
}

const users = await http.get<PaginationResponse<User>>('/users')
```

### ✅ 错误处理

使用 HttpClient 的全局错误处理 + 局部处理：

```typescript
try {
  await http.post('/users', data)
  message.success('创建成功')
} catch (error) {
  // 全局处理器已处理网络错误、401、500 等
  // 这里只需处理业务逻辑错误
}
```

### ✅ 状态管理

- 使用 `useState` 管理本地 UI 状态
- 使用 `useEffect` 处理副作用
- 推荐使用 React Query 管理服务器状态（见文档）

### ✅ 组件复用

使用 `@app/core` 的共享组件：

```typescript
import { Button, Input, Card, Loading } from '@app/core'
```

### ✅ 样式一致性

使用 Tailwind CSS 的主题变量：

```typescript
className="text-app-text bg-surface border-app-border"
```

## 扩展建议

### 添加 React Query

```bash
pnpm add @tanstack/react-query
```

```typescript
import { useQuery } from '@tanstack/react-query'

const { data, isLoading } = useQuery({
  queryKey: ['users', page],
  queryFn: () => userService.list(page),
})
```

### 添加表单库

```bash
pnpm add react-hook-form zod @hookform/resolvers
```

```typescript
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { z } from 'zod'

const schema = z.object({
  name: z.string().min(2),
  email: z.string().email(),
})

const { register, handleSubmit } = useForm({
  resolver: zodResolver(schema),
})
```

## 相关文档

- [最佳实践指南](../../../docs/BEST_PRACTICES.md)
- [HTTP Client 使用指南](../../../docs/HTTP_CLIENT.md)
- [组件库文档](../../../core/src/components/README.md)
