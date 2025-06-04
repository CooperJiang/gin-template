# Admin 版本管理功能

## 功能概述

为 admin 后台管理系统实现了自动版本管理功能，每次构建时自动递增版本号并显示构建时间。

## 版本规则

- 版本格式：`major.minor.patch`（如：1.0.1, 1.0.2, ..., 1.1.0）
- 每次构建时 patch 版本自动 +1
- 当 patch 达到 10 时，minor 版本 +1，patch 重置为 0
- 例如：1.0.9 → 1.1.0

## 实现文件

### 1. 版本管理脚本
- `scripts/version-manager.js` - 核心版本管理逻辑
- 自动更新 `package.json` 中的版本号
- 生成 `src/config/version.json` 配置文件

### 2. 类型定义
- `src/types/version.ts` - 版本信息的 TypeScript 类型定义
- `src/types/json.d.ts` - JSON 模块的类型声明

### 3. 版本信息组合式函数
- `src/composables/useVersion.ts` - 获取和管理版本信息的 Composable

### 4. 界面显示
- `src/layouts/AdminLayout.vue` - 在页面底部显示版本信息和构建时间

### 5. 构建配置
- `vite.config.ts` - 配置版本信息注入到构建中
- `package.json` - 添加版本更新脚本

## 使用方法

### 手动更新版本
```bash
cd admin
npm run version-update
```

### 构建时自动更新
```bash
cd admin
npm run build  # 会自动调用 version-update
```

### 完整项目构建
```bash
make build  # 会自动更新 admin 版本
```

## 显示效果

在管理后台页面底部会显示：
- 版本号：v1.1.2
- 构建时间：2025/06/04 18:08:41
- 加载状态指示器（绿色圆点表示已加载）

## 文件结构

```
admin/
├── scripts/
│   └── version-manager.js          # 版本管理脚本
├── src/
│   ├── config/
│   │   └── version.json           # 自动生成的版本配置
│   ├── types/
│   │   ├── version.ts             # 版本类型定义
│   │   └── json.d.ts              # JSON 模块声明
│   ├── composables/
│   │   └── useVersion.ts          # 版本信息 Composable
│   └── layouts/
│       └── AdminLayout.vue        # 显示版本信息的布局
├── package.json                   # 包含版本号和构建脚本
└── vite.config.ts                # 构建配置
```

## 版本信息数据结构

```typescript
interface VersionInfo {
  version: string;        // 版本号，如 "1.1.2"
  buildTime: string;      // ISO 时间戳
  buildTimestamp: number; // Unix 时间戳
  buildDate: string;      // 格式化的中文日期时间
}
```

## 部署说明

每次运行 `make build` 或 `make build-deploy` 时：
1. 自动更新 admin 版本号
2. 生成新的构建时间
3. 将版本信息嵌入到构建产物中
4. 部署后在管理后台底部可见版本信息

这样您就可以通过查看管理后台底部的版本号和构建时间来确认是否部署了最新代码。 