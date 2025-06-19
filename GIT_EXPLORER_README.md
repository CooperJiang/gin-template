# Git 版本探索工具 🔍

这里提供了两个脚本来帮助您方便地浏览和切换Git历史版本。

## 🚀 脚本介绍

### 1. 交互式版本浏览器 (`git_explorer.sh`)
完整的交互式Git提交浏览工具，提供丰富的功能。

### 2. 快速切换工具 (`quick_checkout.sh`)
简化版本的快速切换脚本，支持关键词搜索。

## 📖 使用方法

### 交互式版本浏览器

```bash
# 启动交互式浏览器
./scripts/git_explorer.sh
```

**功能列表：**
- `数字` - 切换到对应编号的提交
- `b` - 基于当前提交创建新分支
- `l` - 显示更多提交（默认20个）
- `s` - 显示当前提交的详细信息
- `d` - 显示当前提交的文件变更
- `f` - 查看指定文件在当前提交的内容
- `r` - 返回原始分支
- `h` - 显示帮助
- `q` - 退出

### 快速切换工具

```bash
# 显示最近30个提交，选择切换
./scripts/quick_checkout.sh

# 搜索包含特定关键词的提交
./scripts/quick_checkout.sh "uuid"
./scripts/quick_checkout.sh "feat"
./scripts/quick_checkout.sh "fix"
```

## 🎯 使用场景示例

### 场景1：查找UUID相关的提交
```bash
./scripts/quick_checkout.sh "uuid"
```

### 场景2：浏览所有提交并测试代码
```bash
./scripts/git_explorer.sh
# 然后输入数字选择提交，输入 's' 查看详情，输入 'b' 创建分支保存
```

### 场景3：查找特定功能的提交
```bash
./scripts/quick_checkout.sh "admin"
./scripts/quick_checkout.sh "component"
```

## 🔧 实用技巧

### 1. 安全地探索历史版本
- 脚本会记住您的原始分支
- 使用 `r` 命令可以随时返回原始分支
- 退出时会提示您当前的位置

### 2. 保存有用的版本
```bash
# 在交互式浏览器中，找到想要的版本后
> b  # 创建分支
> my-useful-version  # 输入分支名
```

### 3. 快速定位提交
```bash
# 按关键词搜索
./scripts/quick_checkout.sh "关键词"

# 如果只有一个匹配结果，会自动切换
# 如果有多个匹配，会显示列表供选择
```

### 4. 查看代码变更
```bash
# 在交互式浏览器中
> d  # 查看文件变更
> f  # 查看特定文件内容
> s  # 查看详细提交信息
```

## 📋 常用命令组合

### 探索并保存版本
1. `./scripts/git_explorer.sh` - 启动浏览器
2. 输入数字切换到感兴趣的提交
3. 输入 `s` 查看详细信息
4. 输入 `d` 查看文件变更
5. 如果满意，输入 `b` 创建分支保存

### 快速查找特定功能
1. `./scripts/quick_checkout.sh "功能关键词"`
2. 选择合适的提交编号
3. 选择是否创建分支保存

## ⚠️ 重要提示

1. **分支安全**：脚本会自动记住您的起始分支，可以安全使用
2. **未提交的更改**：使用前请确保当前工作区是干净的（已提交所有更改）
3. **创建分支**：找到有用的版本时，建议创建分支保存
4. **返回原位**：使用 `git checkout 原分支名` 或脚本中的 `r` 命令返回

## 🎉 示例输出

```
=== Git 提交浏览器 ===
当前分支: master

使用说明:
  数字 - 切换到对应的提交
  b    - 基于当前提交创建新分支
  l    - 显示更多提交 (默认显示20个)
  ...

最近 20 次提交:

  0. 0cd5b21 2025-06-04 JLY fix: docs (HEAD -> uuid-id-version, origin/v1, origin/master, origin/HEAD, master)
  1. 486d6fd 2025-06-04 JLY feat: v1.0
  2. bced4a8 2025-06-04 JLY feat: add make
  3. 596c307 2025-06-04 JLY fix: ico
  ...

请选择操作 (输入数字选择提交，或输入命令字母):
> 