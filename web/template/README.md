# __TITLE__

> qiankun 微前端子应用

## 📦 安装

```bash
pnpm install
```

## 🚀 开发

```bash
# 单独启动此子应用
pnpm dev

# 或在根目录启动所有应用
cd ../..
pnpm dev
```

访问地址：http://localhost:__PORT__

## 🏗️ 构建

```bash
pnpm build
```

## 📝 脚本命令

```bash
pnpm dev          # 启动开发服务器
pnpm build        # 构建生产版本
pnpm preview      # 预览生产构建
pnpm type-check   # TypeScript 类型检查
pnpm lint         # ESLint 代码检查
pnpm lint:fix     # 自动修复 ESLint 问题
pnpm format       # Prettier 代码格式化
pnpm clean        # 清理构建产物
```

## 🔧 技术栈

- **React** 19.1.0 - UI 框架
- **TypeScript** 5.8.0 - 类型安全
- **Vite** 6.2.4 - 构建工具
- **qiankun** 2.10.16 - 微前端框架
- **React Router** 7.6.1 - 路由管理
- **TailwindCSS** 3.4.17 - 样式框架

## 📚 相关文档

- [最佳实践指南](../../docs/BEST_PRACTICES.md)
- [HTTP Client 使用指南](../../docs/HTTP_CLIENT.md)
- [示例页面](../app/src/pages/Examples/README.md)

## 🎨 开发建议

### 使用共享组件

```typescript
import { Button, Input, Card, Loading, Modal } from '@app/core'

<Button variant="primary" onClick={handleClick}>
  提交
</Button>
```

### 使用 HTTP Client

```typescript
import { createHttpClient } from '@app/core'

const http = createHttpClient({ baseURL: '/api/v1' })
const data = await http.get('/users')
```

### 使用认证

```typescript
import { useAuth } from '@app/auth'

const { user, isAuthenticated, login, logout } = useAuth()
```

## ⚙️ 配置

### 路由配置

编辑 `src/router/index.tsx` 添加新路由：

```typescript
{
  path: 'new-page',
  element: <NewPage />,
}
```

### API 代理

开发环境的 API 代理配置在根目录的 `.env.local` 中：

```env
VITE_API_PROXY_TARGET=http://localhost:9000
```

## 📖 项目结构

```
src/
├── main.tsx          # 应用入口
├── App.tsx           # 根组件
├── router/           # 路由配置
│   └── index.tsx
├── pages/            # 页面组件
│   ├── Home/
│   ├── About/
│   └── Settings/
├── api/              # API 服务
│   └── request.ts
├── hooks/            # 自定义 Hooks
│   └── useMessage.ts
└── styles.css        # 全局样式
```

## 🐛 故障排除

### 端口被占用

如果端口 __PORT__ 已被占用，可以修改 `vite.config.ts` 中的端口号。

### 样式冲突

确保使用了 qiankun 的样式隔离配置，主应用使用 `important: true`，子应用禁用 `preflight`。

### 路由 404

确保子应用的路由 basename 与 qiankun 的 activeRule 匹配：

```typescript
// router/index.tsx
basename: '/__NAME__'
```

## 📞 支持

如有问题，请查看 [故障排除文档](../../TROUBLESHOOTING.md) 或联系团队。
