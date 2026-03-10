# HTTP Client 使用指南

本项目使用基于 Axios 的 `HttpClient` 进行所有 HTTP 请求。它提供了统一的错误处理、请求/响应拦截和类型安全的 API。

## 快速开始

### 创建客户端实例

```typescript
import { createHttpClient } from '@app/core'

const http = createHttpClient({
  baseURL: '/api',
  timeout: 10000, // 可选，默认 10 秒
})
```

### 基础用法

```typescript
// GET 请求
const users = await http.get<User[]>('/users')

// POST 请求
const newUser = await http.post<User>('/users', {
  name: 'Alice',
  email: 'alice@example.com',
})

// PUT 请求
const updatedUser = await http.put<User>(`/users/${id}`, userData)

// DELETE 请求
await http.delete(`/users/${id}`)
```

## 类型定义

### BaseResponse

所有 API 响应都遵循统一的结构：

```typescript
interface BaseResponse<T> {
  code: number
  message: string
  data: T
}
```

### PaginationResponse

分页数据使用：

```typescript
interface PaginationResponse<T> {
  items: T[]
  total: number
  page: number
  pageSize: number
  totalPages: number
}
```

示例：

```typescript
import type { PaginationResponse } from '@app/shared/types'

const response = await http.get<PaginationResponse<User>>('/users', {
  params: {
    page: 1,
    pageSize: 20,
  },
})

console.log(response.items) // User[]
console.log(response.total) // 总记录数
```

## 错误处理

### 自动错误处理

HttpClient 会自动处理常见错误：

- **网络错误**：显示 "网络连接失败" 消息
- **超时**：显示 "请求超时" 消息
- **401 未授权**：自动跳转到登录页
- **403 禁止访问**：显示权限不足消息
- **404 未找到**：显示资源不存在消息
- **500 服务器错误**：显示服务器错误消息

### 手动错误处理

如果需要自定义错误处理：

```typescript
try {
  const data = await http.get<User>('/users/123')
  // 处理成功响应
} catch (error) {
  if (error.response?.status === 404) {
    // 自定义 404 处理
    console.log('用户不存在')
  } else {
    // 其他错误由全局处理器处理
    throw error
  }
}
```

### 禁用自动错误提示

如果不希望显示错误消息：

```typescript
const http = createHttpClient({
  baseURL: '/api',
  // 添加自定义响应拦截器
  interceptors: {
    response: {
      onRejected: (error) => {
        // 静默处理错误
        return Promise.reject(error)
      },
    },
  },
})
```

## 请求拦截

### 添加认证令牌

默认情况下，HttpClient 会自动从 `SecureStorage` 读取 token 并添加到请求头：

```typescript
// 这是自动完成的，无需手动操作
headers: {
  'Authorization': 'Bearer <token>'
}
```

### 自定义请求头

```typescript
const http = createHttpClient({
  baseURL: '/api',
  headers: {
    'X-Custom-Header': 'value',
  },
})
```

或者在单个请求中：

```typescript
await http.get('/users', {
  headers: {
    'X-Request-ID': generateId(),
  },
})
```

## 响应拦截

### 自动解包响应数据

HttpClient 会自动解包 `BaseResponse` 的 `data` 字段：

```typescript
// 后端返回：{ code: 200, message: "成功", data: { id: 1, name: "Alice" } }
const user = await http.get<User>('/users/1')
// user = { id: 1, name: "Alice" } （已自动解包）
```

### 访问完整响应

如果需要访问完整响应（包括 headers、status 等）：

```typescript
import axios from 'axios'

const response = await axios.get('/api/users/1')
console.log(response.status) // 200
console.log(response.headers)
console.log(response.data) // { code: 200, message: "成功", data: {...} }
```

## 高级用法

### 取消请求

```typescript
import axios from 'axios'

const controller = new AbortController()

const request = http.get('/users', {
  signal: controller.signal,
})

// 稍后取消
controller.abort()
```

### 上传文件

```typescript
const formData = new FormData()
formData.append('file', file)

await http.post('/upload', formData, {
  headers: {
    'Content-Type': 'multipart/form-data',
  },
  onUploadProgress: (progressEvent) => {
    const percentCompleted = Math.round(
      (progressEvent.loaded * 100) / progressEvent.total
    )
    console.log(`上传进度: ${percentCompleted}%`)
  },
})
```

### 下载文件

```typescript
const blob = await http.get<Blob>('/files/download/123', {
  responseType: 'blob',
})

const url = window.URL.createObjectURL(blob)
const a = document.createElement('a')
a.href = url
a.download = 'file.pdf'
a.click()
window.URL.revokeObjectURL(url)
```

### 并发请求

```typescript
const [users, posts, comments] = await Promise.all([
  http.get<User[]>('/users'),
  http.get<Post[]>('/posts'),
  http.get<Comment[]>('/comments'),
])
```

### 请求重试

```typescript
import axios from 'axios'
import axiosRetry from 'axios-retry'

const http = createHttpClient({ baseURL: '/api' })

// 配置重试逻辑
axiosRetry(http as any, {
  retries: 3,
  retryDelay: axiosRetry.exponentialDelay,
  retryCondition: (error) => {
    return error.response?.status === 503 || error.code === 'ECONNABORTED'
  },
})
```

## 实战示例

### 用户列表页面

```typescript
import { useState, useEffect } from 'react'
import { createHttpClient } from '@app/core'
import type { PaginationResponse } from '@app/shared/types'

interface User {
  id: number
  name: string
  email: string
}

export function UserList() {
  const [users, setUsers] = useState<User[]>([])
  const [loading, setLoading] = useState(false)
  const [total, setTotal] = useState(0)
  const [page, setPage] = useState(1)

  const http = createHttpClient({ baseURL: '/api' })

  useEffect(() => {
    const fetchUsers = async () => {
      setLoading(true)
      try {
        const response = await http.get<PaginationResponse<User>>('/users', {
          params: { page, pageSize: 20 },
        })
        setUsers(response.items)
        setTotal(response.total)
      } catch (error) {
        // 错误已自动处理
      } finally {
        setLoading(false)
      }
    }

    fetchUsers()
  }, [page])

  return (
    <div>
      {loading && <p>加载中...</p>}
      <ul>
        {users.map((user) => (
          <li key={user.id}>{user.name}</li>
        ))}
      </ul>
      <button onClick={() => setPage(page + 1)}>下一页</button>
    </div>
  )
}
```

### 表单提交

```typescript
import { useState } from 'react'
import { createHttpClient } from '@app/core'
import { message } from '@app/core'

export function UserForm() {
  const [formData, setFormData] = useState({ name: '', email: '' })
  const [submitting, setSubmitting] = useState(false)

  const http = createHttpClient({ baseURL: '/api' })

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setSubmitting(true)

    try {
      await http.post('/users', formData)
      message.success('用户创建成功')
      setFormData({ name: '', email: '' })
    } catch (error) {
      // 错误已由全局拦截器处理
    } finally {
      setSubmitting(false)
    }
  }

  return (
    <form onSubmit={handleSubmit}>
      <input
        value={formData.name}
        onChange={(e) => setFormData({ ...formData, name: e.target.value })}
        placeholder="姓名"
      />
      <input
        value={formData.email}
        onChange={(e) => setFormData({ ...formData, email: e.target.value })}
        placeholder="邮箱"
      />
      <button type="submit" disabled={submitting}>
        {submitting ? '提交中...' : '提交'}
      </button>
    </form>
  )
}
```

## 最佳实践

### 1. 在服务层创建 HttpClient

```typescript
// services/api.ts
import { createHttpClient } from '@app/core'

export const api = createHttpClient({
  baseURL: '/api/v1',
  timeout: 15000,
})

// services/user.service.ts
import { api } from './api'
import type { User, PaginationResponse } from '@app/shared/types'

export const userService = {
  list: (page: number, pageSize: number) =>
    api.get<PaginationResponse<User>>('/users', {
      params: { page, pageSize },
    }),

  get: (id: number) => api.get<User>(`/users/${id}`),

  create: (data: Partial<User>) => api.post<User>('/users', data),

  update: (id: number, data: Partial<User>) =>
    api.put<User>(`/users/${id}`, data),

  delete: (id: number) => api.delete(`/users/${id}`),
}
```

### 2. 在组件中使用服务

```typescript
import { userService } from '@/services/user.service'

const users = await userService.list(1, 20)
```

### 3. 使用 React Query 管理缓存

```typescript
import { useQuery, useMutation } from '@tanstack/react-query'
import { userService } from '@/services/user.service'

export function useUsers(page: number) {
  return useQuery({
    queryKey: ['users', page],
    queryFn: () => userService.list(page, 20),
  })
}

export function useCreateUser() {
  return useMutation({
    mutationFn: userService.create,
    onSuccess: () => {
      message.success('创建成功')
    },
  })
}
```

## 常见问题

### Q: 如何修改超时时间？

```typescript
const http = createHttpClient({
  baseURL: '/api',
  timeout: 30000, // 30 秒
})
```

### Q: 如何在请求中携带 Cookie？

```typescript
const http = createHttpClient({
  baseURL: '/api',
  withCredentials: true,
})
```

### Q: 如何调试请求？

在开发环境中，所有请求和响应都会打印到控制台。生产环境中自动禁用。

### Q: 如何处理 CORS 错误？

CORS 错误通常由后端配置决定。确保后端正确配置了 CORS 中间件。开发环境中，Vite 的代理已配置：

```typescript
// vite.config.ts 中已配置
proxy: {
  '/api': {
    target: 'http://localhost:9000',
    changeOrigin: true,
  },
}
```

## 相关资源

- [Axios 官方文档](https://axios-http.com/)
- [TypeScript 类型定义](../shared/src/types/base.ts)
- [认证系统文档](./AUTHENTICATION.md)
