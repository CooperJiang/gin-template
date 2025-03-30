# 项目部署指南

本文档详细介绍如何将Go Web项目部署到Linux服务器上，包括环境准备、SSH密钥配置、自动化部署等全部步骤。

## 目录

- [部署前准备](#部署前准备)
- [SSH密钥配置](#ssh密钥配置)
- [环境配置](#环境配置)
- [部署脚本说明](#部署脚本说明)
- [部署流程](#部署流程)
- [常见问题排查](#常见问题排查)

## 部署前准备

### 1. 本地环境要求

- Go开发环境 (1.16+)
- Git
- Bash终端环境
- SSH客户端

### 2. 远程服务器要求

- Linux服务器 (推荐Ubuntu/CentOS)
- SSH服务开启
- 足够的磁盘空间 (至少200MB)
- 建议4GB以上内存

### 3. 本地项目结构确认

确认项目中已包含以下文件：
- `build.sh`: 项目构建脚本
- `deploy.sh`: 部署脚本
- `config.example.yaml`: 配置文件模板

## SSH密钥配置

SSH密钥认证允许你无需每次输入密码即可连接服务器，这对自动化部署至关重要。

### 1. 检查现有SSH密钥

```bash
ls -la ~/.ssh/
```

常见的密钥文件包括：
- `id_rsa` (私钥) 和 `id_rsa.pub` (公钥)
- `id_ed25519` 和 `id_ed25519.pub`

### 2. 创建新的SSH密钥 (如果没有)

```bash
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
```

按提示操作，建议使用默认保存路径 (~/.ssh/id_rsa)。

### 3. 获取SSH密钥的完整路径

```bash
realpath ~/.ssh/id_rsa  # 查看私钥路径
```

记住这个路径，配置部署脚本时会用到。

### 4. 将SSH公钥上传到服务器

最简单的方法是使用`ssh-copy-id`：

```bash
ssh-copy-id 用户名@服务器IP  # 例如: ssh-copy-id root@192.168.1.100
```

如果`ssh-copy-id`命令不可用，可以手动上传：

```bash
# 查看公钥内容
cat ~/.ssh/id_rsa.pub

# 登录服务器
ssh 用户名@服务器IP

# 在服务器上执行
mkdir -p ~/.ssh
chmod 700 ~/.ssh
echo "你的公钥内容" >> ~/.ssh/authorized_keys
chmod 600 ~/.ssh/authorized_keys
exit
```

### 5. 测试SSH密钥连接

```bash
ssh 用户名@服务器IP
```

如果不需要输入密码即可登录，说明SSH密钥配置成功。

## 环境配置

项目支持多环境部署，每个环境可以使用不同的配置文件。

### 1. 配置文件命名约定

- `config.yaml` - 本地开发配置（默认）
- `config.prod.yaml` - 生产环境配置
- `config.example.yaml` - 示例配置（提交到Git）

### 2. 创建环境配置文件

为每个目标环境创建特定的配置文件：

```bash
# 创建生产环境配置（从示例配置复制）
cp config.example.yaml config.prod.yaml

# 编辑配置，修改环境特定参数
vi config.prod.yaml
```

修改关键配置项：
- 数据库连接信息（主机、端口、用户名、密码）
- Redis连接信息
- JWT密钥（生产环境需使用强密钥）
- 日志路径和级别
- 应用运行模式（生产环境使用"release"模式）

### 3. 配置文件安全

- 将特定环境的配置文件添加到`.gitignore`以避免提交敏感信息
- 只提交`config.example.yaml`模板文件到代码库
- 考虑使用环境变量或加密存储敏感信息

### 4. 部署时指定环境配置

部署时可以通过以下方式指定使用的配置：

```bash
# 使用环境参数指定环境
./deploy.sh -e prod deploy    # 使用config.prod.yaml

# 或明确指定配置文件路径
./deploy.sh -c /path/to/config.yaml deploy
```

更多详细信息请参考：[环境配置使用指南](ENV_CONFIG_README.md)

## 部署脚本说明

项目包含三个主要脚本:

### 1. `build.sh` - 项目构建脚本

用于将项目编译为不同平台的可执行文件。

**主要参数:**

```bash
./build.sh --help              # 显示帮助信息
./build.sh --local             # 构建当前平台版本
./build.sh --linux             # 构建Linux版本(服务器部署用)
./build.sh --all               # 构建所有支持平台版本
./build.sh -n myapp --linux    # 指定应用名称为"myapp"并构建Linux版本
./build.sh -o dist --linux     # 指定输出目录为"dist"并构建Linux版本
```

### 2. `run.sh` - 本地管理脚本

用于在本地开发环境管理应用。

**主要命令:**

```bash
./run.sh help      # 显示帮助信息
./run.sh build     # 构建应用
./run.sh start     # 启动应用
./run.sh stop      # 停止应用
./run.sh restart   # 重启应用
./run.sh status    # 查看应用状态
./run.sh log       # 查看应用日志
```

### 3. `deploy.sh` - 远程部署脚本

用于将应用部署到远程服务器并管理应用生命周期。

**主要命令和参数:**

```bash
./deploy.sh help                          # 显示帮助信息
./deploy.sh config                        # 配置部署参数
./deploy.sh envconfig                     # 配置环境特定参数
./deploy.sh setup                         # 设置远程服务器环境
./deploy.sh deploy                        # 部署应用到服务器
./deploy.sh rollback                      # 回滚到上一个版本
./deploy.sh restart                       # 重启服务器上的应用
./deploy.sh status                        # 查看服务器上的应用状态
./deploy.sh logs                          # 查看服务器上的应用日志
./deploy.sh -h 192.168.1.100 deploy       # 部署到指定服务器
./deploy.sh -u admin -h 192.168.1.100 deploy  # 指定用户名和服务器
./deploy.sh -n myapp deploy               # 指定应用名称部署
./deploy.sh -d /opt/webapp deploy         # 指定部署目录
./deploy.sh -k ~/.ssh/id_rsa deploy       # 指定SSH密钥路径
./deploy.sh -e prod deploy                # 指定环境（使用对应配置）
./deploy.sh -c config.prod.yaml deploy    # 指定配置文件
```

## 部署流程

下面是完整的部署流程，从本地代码到服务器运行：

### 1. 首次部署准备

> ⚠️ **重要提示：首次部署必须先执行`setup`命令初始化服务器环境！** ⚠️

#### 步骤1: 配置部署参数

```bash
./deploy.sh config
```

按提示输入信息：
- 应用名称 (默认: app)
- 远程主机地址 (服务器IP)
- 远程用户名 (如root)
- SSH端口 (默认: 22)
- 远程部署目录 (如 /opt/myapp)
- SSH密钥路径 (如 ~/.ssh/id_rsa)
- 保留版本数量 (默认: 3)

配置信息会保存到`deploy.conf`文件中。

#### 步骤2: 配置环境特定参数

```bash
./deploy.sh envconfig
```

按提示选择：
- 部署环境 (prod/test/dev)
- 使用的配置文件

#### 步骤3: 设置远程服务器环境

```bash
./deploy.sh setup
```

**此步骤是首次部署的必要操作！** 如果跳过此步骤，部署将会失败。

此命令会:
- 在服务器上创建必要的目录结构
- 上传服务管理脚本
- 设置适当的文件权限

### 2. 部署应用

```bash
./deploy.sh deploy
```

部署过程包括:
- 检查并构建Linux版本应用
- 检查目标端口是否被占用（智能处理已有实例）
- 创建带时间戳的发布目录
- 上传应用和指定的环境配置文件
- 创建环境变量文件
- 停止当前运行的版本(如果有)
- 更新当前版本的软链接
- 启动新版本
- 清理旧版本

**一键部署到指定环境：**
```bash
# 部署到生产环境
./deploy.sh -e prod deploy

# 部署到测试环境
./deploy.sh -e test deploy
```

### 3. 管理部署的应用

```bash
./deploy.sh status                  # 查看应用状态
./deploy.sh restart                 # 重启应用
./deploy.sh logs                    # 查看应用日志
./deploy.sh logs 100                # 查看最后100行日志
```

### 4. 版本回滚 (如需)

如果新版本有问题，可以一键回滚到上一个版本:

```bash
./deploy.sh rollback
```

此操作会:
- 自动停止当前版本
- 将当前版本链接切换到上一个部署版本
- 启动上一个版本

## 远程服务器目录结构

部署脚本会在服务器上创建以下目录结构:

```
/opt/myapp/                 # 应用根目录(可配置)
├── current -> ./releases/20231030120000/  # 指向当前版本的软链接
├── releases/              # 存放所有部署版本
│   ├── 20231030120000/    # 基于时间戳的版本目录
│   │   ├── app           # 应用二进制文件
│   │   └── config.yaml -> ../../shared/config/config.yaml  # 配置文件软链接
│   └── 20231029120000/    # 旧版本目录
├── shared/                # 共享文件
│   ├── config/            # 配置文件目录
│   │   └── config.yaml    # 当前环境配置文件
│   ├── .env               # 环境变量文件
│   ├── app.pid           # 应用PID文件
│   └── app.log           # 应用日志文件
├── service.sh            # 服务管理脚本
└── tmp/                  # 临时文件
```

## 智能端口占用处理

从`v1.1.0`版本开始，部署脚本支持智能端口占用检测和处理：

- 自动从配置文件中识别应用使用的端口
- 在部署前检查目标端口是否被占用
- 自动识别端口被自身旧实例占用的情况并处理
- 对于非自身应用占用的端口，提供交互式处理选项

此功能可以有效避免部署后应用因端口被占用而无法启动的问题。

## 常见问题排查

### 1. SSH密钥问题

**问题**: 每次部署都要求输入密码

**解决方法**:
- 检查SSH密钥路径配置是否正确
- 确认公钥已上传到服务器
- 检查服务器上`.ssh/authorized_keys`文件权限 (应为600)
- 使用`ssh -v 用户名@服务器IP`查看详细连接信息

### 2. 权限问题

**问题**: 应用部署后无法启动

**解决方法**:
- 检查应用二进制文件权限 (应为755): `chmod 755 app`
- 检查日志目录权限: `chmod 755 shared`
- 查看应用日志了解具体错误

### 3. 网络连接问题

**问题**: 无法连接到服务器

**解决方法**:
- 确认服务器IP和SSH端口正确
- 检查服务器防火墙设置
- 确认服务器SSH服务正在运行

### 4. 配置文件问题

**问题**: 应用启动后立即崩溃

**解决方法**:
- 检查配置文件格式是否正确（有效的YAML）
- 确认配置文件中的连接信息在目标环境中有效
- 检查日志文件了解具体错误
- 使用`-c`参数指定正确的配置文件路径

### 5. 端口占用问题

**问题**: 应用启动失败，提示端口已被占用

**解决方法**:
- 在部署前，脚本会自动检查端口占用情况
- 如果是应用自身的旧实例，会自动尝试停止
- 对于其他应用占用的端口，可以根据提示选择处理方式
- 也可以手动停止占用端口的进程: `lsof -i:<端口> | grep LISTEN`

## 部署脚本高级功能

### 1. 多环境部署

创建环境特定的配置文件:

```bash
cp config.example.yaml config.prod.yaml  # 创建生产环境配置
```

部署到生产环境:

```bash
./deploy.sh -h prod-server.com -u deploy -e prod deploy
```

### 2. 自动化部署

可以结合CI/CD系统(如Jenkins、GitHub Actions)使用部署脚本:

```yaml
# 在GitHub Actions工作流中
jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Deploy to production
        run: |
          echo "$SSH_KEY" > deploy_key.pem
          chmod 600 deploy_key.pem
          ./deploy.sh -h $SERVER_IP -u $SERVER_USER -k ./deploy_key.pem -e prod deploy
        env:
          SSH_KEY: ${{ secrets.SSH_PRIVATE_KEY }}
          SERVER_IP: ${{ secrets.SERVER_IP }}
          SERVER_USER: ${{ secrets.SERVER_USER }}
```

### 3. 自定义部署脚本

可以根据项目需求修改`deploy.sh`脚本，例如:
- 添加数据库迁移步骤
- 添加部署前后的钩子功能
- 集成与消息通知系统 (如Slack、钉钉)

## 小结

通过本指南中的部署脚本和流程，你可以实现:
- 自动化构建和部署Go应用
- 支持版本管理和快速回滚
- 无密码SSH认证部署
- 多环境配置管理
- 智能端口占用检测与处理
- 远程应用状态监控

遵循这些步骤可以大大简化应用部署和管理流程，减少手动操作错误，提高开发和运维效率。

**重要提示：首次使用务必执行完整的初始化流程，尤其是`setup`命令！** 