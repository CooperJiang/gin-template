# ⚡ 快速部署参考

## 🚀 首次部署
```bash
make deploy-config   # 配置服务器信息
make deploy-setup    # 初始化服务器
make build-deploy    # 构建并部署
```

## 🔄 日常更新
```bash
make build-deploy    # 一键构建并部署
```

## 🛠️ 管理命令
```bash
make deploy-status   # 查看状态
make deploy-logs     # 查看日志
make deploy-restart  # 重启应用
make deploy-rollback # 回滚版本
```

## 📊 版本查看
- 访问管理后台底部查看版本号
- 每次部署版本自动递增

---
**核心命令：`make build-deploy`** 🎯 