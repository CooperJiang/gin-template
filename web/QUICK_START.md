# ⚡ 快速上手指南

> 5分钟快速了解常用命令和开发流程

## 📦 首次安装

```bash
cd web
pnpm install
```

首次安装会自动创建 `.env.local` 文件。

---

## 🚀 常用命令速查

### 开发命令

```bash
# 启动主应用 + 默认子应用（最常用）
pnpm dev

# 启动所有应用（包括 theme-editor）
pnpm dev:full

# 仅启动主应用
pnpm dev:main

# 仅启动某个子应用
pnpm dev:app        # app 子应用
pnpm dev:admin      # admin 子应用（如果已创建）
```

### 构建命令

```bash
# 构建所有应用
pnpm build

# 构建单个应用
pnpm build:main
pnpm build:app
```

### 代码质量

```bash
# 类型检查
pnpm type-check

# 代码检查
pnpm lint

# 自动修复
pnpm lint:fix

# 格式化代码
pnpm format

# 完整检查（类型 + lint + 格式）
pnpm check
```

### 子应用管理

```bash
# 创建新子应用
pnpm create:app <name> --port <port> --title "<标题>" --install

# 示例
pnpm create:app admin --port 3002 --title "管理后台" --install
pnpm create:app dashboard --install  # 端口自动分配
```

### 清理

```bash
# 清理所有构建产物
pnpm clean
```

---

## 🎯 典型工作流

### 场景1：启动项目开发

```bash
# 1. 进入 web 目录
cd web

# 2. 安装依赖（首次）
pnpm install

# 3. 启动开发服务器
pnpm dev

# 访问：
# - 主应用：http://localhost:3000
# - app子应用：http://localhost:3001
```

### 场景2：创建新子应用

```bash
# 1. 创建子应用
pnpm create:app admin --port 3002 --title "管理后台" --install

# 2. 开发新子应用
pnpm dev:admin

# 3. 或启动所有应用
pnpm dev
```

### 场景3：修复代码规范问题

```bash
# 1. 检查问题
pnpm lint

# 2. 自动修复
pnpm lint:fix

# 3. 格式化代码
pnpm format
```

### 场景4：提交代码

```bash
# Git hooks 会自动检查
git add .
git commit -m "feat: add new feature"
# ✓ 自动运行 eslint 和 prettier
# ✓ 有问题会阻止提交
```

---

## 📂 目录结构速览

```
web/
├── packages/          # 微前端应用
│   ├── main/         # qiankun 宿主应用（端口 3000）
│   ├── app/          # 默认子应用（端口 3001）
│   └── [其他子应用]/
├── shared/           # 共享类型和工具
├── core/             # 核心服务和组件
├── auth/             # 认证模块
├── template/         # 子应用模板
├── scripts/          # CLI 脚本
└── docs/             # 文档
```

---

## 🛠️ 开发技巧

### 使用共享组件

```typescript
import { Button, Input, Card, Loading, Modal } from '@app/core'

<Button variant="primary" loading={submitting}>
  提交
</Button>
```

### 使用 HTTP Client

```typescript
import { createHttpClient } from '@app/core'

const http = createHttpClient({ baseURL: '/api/v1' })

// 自动错误处理
const users = await http.get<User[]>('/users')
```

### 使用认证

```typescript
import { useAuth } from '@app/auth'

const { user, isAuthenticated, login, logout } = useAuth()

if (!isAuthenticated) {
  return <Login />
}
```

### 跨应用通信

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

---

## 🐛 常见问题速查

### 端口被占用

```bash
# 查看占用端口的进程
lsof -ti:3000

# 杀死进程
kill -9 $(lsof -ti:3000)
```

### 依赖安装失败

```bash
# 清除缓存重新安装
rm -rf node_modules pnpm-lock.yaml
pnpm install
```

### 类型错误

```bash
# 重新检查类型
pnpm type-check

# 如果是 shared 包，先构建
pnpm --filter @app/shared build
```

### 样式不生效

```bash
# 检查 Tailwind 配置
# 主应用：important: true
# 子应用：corePlugins: { preflight: false }
```

### Git 提交被阻止

```bash
# 先修复代码问题
pnpm lint:fix
pnpm format

# 再提交
git commit -m "message"
```

---

## 📖 参考文档

| 文档 | 说明 |
|------|------|
| [README.md](./README.md) | 项目总览和详细说明 |
| [BEST_PRACTICES.md](./docs/BEST_PRACTICES.md) | 最佳实践指南 |
| [HTTP_CLIENT.md](./docs/HTTP_CLIENT.md) | HTTP Client 使用指南 |
| [TROUBLESHOOTING.md](./TROUBLESHOOTING.md) | 故障排除指南 |
| [示例代码](./packages/app/src/pages/Examples/) | 真实世界示例 |

---

## 🎓 学习路径

### 新手（第1天）

1. ✅ 运行 `pnpm install && pnpm dev`
2. ✅ 访问 http://localhost:3000
3. ✅ 查看 `packages/app/src/pages/Examples/` 示例代码

### 进阶（第2-3天）

4. ✅ 阅读 [BEST_PRACTICES.md](./docs/BEST_PRACTICES.md)
5. ✅ 创建第一个子应用：`pnpm create:app test --install`
6. ✅ 使用共享组件开发页面

### 高级（第4-7天）

7. ✅ 阅读 [HTTP_CLIENT.md](./docs/HTTP_CLIENT.md)
8. ✅ 实现完整的 CRUD 功能
9. ✅ 配置 CI/CD 和部署

---

## 💡 提示

- **Git hooks**: 每次提交都会自动检查代码，确保代码质量
- **热更新**: 主应用支持 HMR，子应用需手动刷新（qiankun 限制）
- **端口分配**: 新子应用会自动分配可用端口
- **环境变量**: 修改 `.env.local` 后需重启开发服务器

---

需要帮助？查看 [故障排除文档](./TROUBLESHOOTING.md) 或联系团队。
