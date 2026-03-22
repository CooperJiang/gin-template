# 🚀 部署指南

本项目提供了基于 Makefile 的一键部署解决方案，简单高效。

## 📋 快速部署 (3步搞定)

### 1️⃣ 配置部署参数 (首次)
```bash
make deploy-config
```
**输入以下信息:**
- 应用名称 (默认: email-manage)
- 服务器地址 (必填)
- SSH用户名 (默认: root)
- SSH端口 (默认: 22)
- 部署目录 (默认: /opt/myapp)
- SSH密钥路径 (自动检测)
- 部署环境 (prod/test/dev，默认: prod)

### 2️⃣ 初始化服务器 (首次)
```bash
make deploy-setup
```

### 3️⃣ 一键部署
```bash
make build-deploy
```

**🎉 完成！访问您的应用即可！**

---

## 🛠️ 常用命令

### 构建命令
```bash
make build          # 构建项目 (包含前端版本自动更新)
```

### 部署命令
```bash
make build-deploy    # 🚀 一键构建并部署
make deploy          # 仅部署 (需先运行 make build)
```

### 管理命令
```bash
make deploy-status   # 查看应用状态
make deploy-logs     # 查看应用日志
make deploy-restart  # 重启应用
make deploy-rollback # 回滚到上一版本
```

### 查看配置
```bash
make deploy-check    # 查看当前部署配置
```

---

## 🔧 版本管理特性

### Admin 后台版本自动更新
- ✅ 每次构建自动递增版本号 (1.0.1 → 1.0.2 → ... → 1.1.0)
- ✅ 显示构建时间，方便确认部署是否成功
- ✅ 在管理后台底部显示版本信息

### 快速部署特性
- ✅ **2-3秒切换时间** - 几乎零停机部署
- ✅ **版本管理** - 保留最近5个版本，支持一键回滚
- ✅ **智能检测** - 自动处理端口占用和进程管理
- ✅ **多环境支持** - prod/test/dev环境配置

---

## 📂 部署包说明

运行 `make build` 后会生成：
- `./release/email-manage_prod_package.tar.gz` - 完整部署包

**包含内容:**
- 编译好的 Go 二进制文件 (嵌入前端静态文件)
- 环境配置文件 (config.prod.yaml)
- 启动脚本 (run.sh)

---

## 🌍 多环境部署

```bash
# 部署到测试环境
DEPLOY_ENV=test make build-deploy

# 部署到开发环境
DEPLOY_ENV=dev make build-deploy
```

---

## 🔐 SSH配置

### 如果没有SSH密钥
```bash
# 生成密钥
ssh-keygen -t rsa -b 4096 -f ~/.ssh/id_rsa

# 复制到服务器
ssh-copy-id root@your-server.com
```

### 支持的密钥类型
- `~/.ssh/id_rsa` (RSA)
- `~/.ssh/id_ed25519` (Ed25519)  
- `~/.ssh/id_ecdsa` (ECDSA)

---

## 🚨 故障排除

### 常见问题
```bash
# SSH连接失败
chmod 600 ~/.ssh/id_rsa
ssh root@your-server.com  # 测试连接

# 应用启动失败
make deploy-logs          # 查看日志
make deploy-status        # 查看状态

# 重新初始化服务器
make deploy-setup         # 重新设置环境
```

### 端口冲突
如果遇到端口占用，系统会自动处理旧进程。

---

## 💡 最佳实践

### 生产环境部署流程
```bash
# 1. 首次部署
make deploy-config        # 配置参数
make deploy-setup         # 初始化服务器
make build-deploy         # 部署应用

# 2. 日常更新
make build-deploy         # 一键更新部署

# 3. 紧急回滚
make deploy-rollback      # 回滚到上一版本
```

### 检查部署结果
- 访问管理后台底部查看版本号是否更新
- 运行 `make deploy-status` 确认应用运行状态
- 运行 `make deploy-logs` 查看启动日志

---

**🎯 一个命令搞定部署：`make build-deploy`** 