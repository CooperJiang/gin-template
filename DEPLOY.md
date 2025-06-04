# 部署指南

本项目已将部署功能集成到 Makefile 中，提供了完整的一键部署解决方案。

## 🚀 快速开始

### 1. 配置部署参数
```bash
make deploy-config
```

**自动功能**：
- ✅ **自动检测SSH密钥** - 自动查找 `~/.ssh/id_rsa`、`~/.ssh/id_ed25519`、`~/.ssh/id_ecdsa`
- ✅ **智能路径获取** - 使用 `realpath` 获取完整路径
- ✅ **友好提示** - 如果没有SSH密钥，提供生成和配置指导

按提示输入：
- 应用名称 (默认: template)
- 远程主机地址 (必填)
- 远程用户名 (默认: root)
- SSH端口 (默认: 22)
- 远程部署目录 (默认: /opt/myapp)
- SSH密钥路径 (自动检测或手动输入)
- 部署环境 (prod/test/dev，默认: prod)

配置会保存到 `deploy.conf` 文件中。

### 2. 初始化服务器环境
```bash
make deploy-setup
```
此命令会：
- 创建必要的目录结构
- 上传服务管理脚本 (`service.sh`)
- 设置正确的权限

### 3. 部署应用

#### 方式一：一键打包并部署
```bash
make build-deploy
```

#### 方式二：分步部署
```bash
make deploy-package  # 创建部署包
make deploy          # 部署应用
```

## 📦 部署包说明

部署包存储在 `release/` 目录中，包含：
- **编译好的二进制文件** - 包含嵌入式静态文件
- **环境配置文件** - 根据 `DEPLOY_ENV` 自动选择：
  - `prod` → `config.prod.yaml`
  - `test` → `config.test.yaml`
  - `dev` → `config.dev.yaml`
  - 默认 → `config.yaml`

## 🔧 管理命令

### 创建部署包
```bash
make deploy-package  # 创建部署包到release目录
```

### 仅部署 (不打包)
```bash
make deploy
```
注意：需要先运行 `make deploy-package` 创建部署包

### 查看部署配置
```bash
make deploy-check  # 显示当前配置
```

### 查看应用状态
```bash
make deploy-status
```

### 查看应用日志
```bash
make deploy-logs
# 或指定行数
LINES=100 make deploy-logs
```

### 重启应用
```bash
make deploy-restart
```

### 回滚到上一版本
```bash
make deploy-rollback
```

## 🛠️ 服务管理

远程服务器上的 `service.sh` 脚本支持以下命令：
```bash
# 在服务器上执行
/path/to/deploy/service.sh start    # 启动应用
/path/to/deploy/service.sh stop     # 停止应用
/path/to/deploy/service.sh restart  # 重启应用
/path/to/deploy/service.sh status   # 查看状态
/path/to/deploy/service.sh logs 50  # 查看日志
```

## 🔐 SSH密钥配置

### 如果没有SSH密钥：
```bash
# 生成SSH密钥
ssh-keygen -t rsa -b 4096 -f ~/.ssh/id_rsa

# 复制公钥到服务器
ssh-copy-id -i ~/.ssh/id_rsa.pub root@your-server.com
```

### 支持的密钥类型：
- RSA: `~/.ssh/id_rsa`
- Ed25519: `~/.ssh/id_ed25519`
- ECDSA: `~/.ssh/id_ecdsa`

## 📋 完整部署流程

### 方式一：一键部署
```bash
# 1. 配置部署参数 (首次)
make deploy-config

# 2. 初始化服务器环境 (首次)
make deploy-setup

# 3. 一键打包并部署
make build-deploy

# 4. 查看状态
make deploy-status

# 5. 查看日志
make deploy-logs
```

### 方式二：分步部署
```bash
# 1. 配置部署参数 (首次)
make deploy-config

# 2. 初始化服务器环境 (首次)
make deploy-setup

# 3. 创建部署包
make deploy-package

# 4. 部署应用
make deploy

# 5. 查看状态
make deploy-status

# 6. 查看日志
make deploy-logs
```

## 🎯 生产环境优化

项目已包含生产环境配置 `config.prod.yaml`：
- 发布模式 (`mode: release`)
- 禁用调试日志
- 优化的CORS设置
- 生产级安全配置

## 🚨 故障排除

### 常见问题：

1. **SSH连接失败**
   - 检查SSH密钥权限：`chmod 600 ~/.ssh/id_rsa`
   - 验证服务器连接：`ssh -i ~/.ssh/id_rsa root@your-server.com`

2. **service.sh不存在**
   - 运行 `make deploy-setup` 重新初始化

3. **应用启动失败**
   - 检查日志：`make deploy-logs`
   - 检查配置文件是否正确

4. **端口冲突**
   - 修改 `config.prod.yaml` 中的端口设置
   - 重新部署：`make deploy`

## 💡 高级用法

### 多环境部署
```bash
# 部署到测试环境
DEPLOY_ENV=test make deploy

# 部署到开发环境  
DEPLOY_ENV=dev make deploy
```

### 自定义配置
编辑 `deploy.conf` 文件直接修改配置，或重新运行 `make deploy-config`。

---

🎉 现在你可以享受一键部署的便利！ 