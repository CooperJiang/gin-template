# 环境配置使用指南

本文档详细说明如何管理不同环境（开发、测试、生产）的配置文件，确保安全和灵活的部署。

## 配置文件命名约定

项目使用以下命名约定来区分不同环境的配置文件：

- `config.yaml` - 本地开发配置（默认）
- `config.dev.yaml` - 开发环境配置
- `config.test.yaml` - 测试环境配置 
- `config.prod.yaml` - 生产环境配置
- `config.example.yaml` - 示例配置（包含所有可配置项和说明）

**重要：** 除了`config.example.yaml`以外，所有环境配置文件都应该添加到`.gitignore`中，避免将敏感信息提交到代码库。

## 配置文件模板使用

项目提供了一个配置文件模板`config.example.yaml`，它包含所有配置选项并附有说明。

### 创建新环境配置

1. 复制示例配置文件作为模板：
   ```bash
   cp config.example.yaml config.prod.yaml
   ```

2. 编辑新配置文件，根据目标环境修改相关配置项：
   ```bash
   vi config.prod.yaml
   ```

3. 重点关注以下关键配置项：
   - 数据库连接参数
   - Redis连接参数
   - JWT密钥
   - 日志路径和级别
   - 应用运行模式

## 部署流程中的配置选择

部署脚本`deploy.sh`提供了多种方式来选择要使用的配置文件：

### 1. 使用环境参数

```bash
./deploy.sh -e prod deploy       # 使用config.prod.yaml
./deploy.sh -e test deploy       # 使用config.test.yaml
./deploy.sh -e dev deploy        # 使用config.dev.yaml
```

### 2. 明确指定配置文件

```bash
./deploy.sh -c /path/to/special-config.yaml deploy
```

### 3. 交互式配置环境

```bash
./deploy.sh envconfig            # 交互式选择环境和配置
./deploy.sh deploy               # 使用之前配置的环境和配置文件
```

## 环境变量和配置

部署到远程服务器时，系统会生成`.env`文件，包含以下环境变量：

```
APP_ENV=prod            # 当前环境（prod/test/dev）
APP_VERSION=1.0.0       # 应用版本号
APP_DEPLOY_TIME=...     # 部署时间戳
```

应用程序可以通过以下方式读取这些环境变量：

```go
os.Getenv("APP_ENV")
```

## 配置文件安全处理

### 敏感信息处理

以下敏感信息应谨慎处理：

1. 数据库凭据
2. API密钥和JWT密钥
3. 邮件服务凭据
4. 第三方服务账号信息

### 配置加密（可选）

对于高度敏感的生产环境配置，可以考虑：

1. 使用环境变量而非配置文件存储敏感信息
2. 使用密钥管理服务（如AWS KMS、HashiCorp Vault）
3. 在部署过程中解密敏感配置

## 最佳实践

1. **永不提交实际配置文件到代码库**
   - 只提交`config.example.yaml`作为模板

2. **使用环境特定的默认值**
   - 开发环境：详细日志、禁用缓存
   - 测试环境：模拟外部服务
   - 生产环境：优化性能参数

3. **配置变更管理**
   - 记录所有配置变更
   - 在部署前审查配置变更
   - 使用版本控制系统管理配置模板

4. **配置验证**
   - 在应用启动时验证配置的完整性和正确性
   - 对缺失或无效的配置提供清晰的错误信息

## 环境差异配置指南

不同环境通常需要差异化的配置设置：

### 开发环境 (`config.dev.yaml`)
```yaml
app:
  mode: "debug"
  enable_pprof: true
log:
  level: "debug"
```

### 测试环境 (`config.test.yaml`)
```yaml
app:
  mode: "debug"  # 便于调试
database:
  # 使用测试数据库
  name: "template_test" 
```

### 生产环境 (`config.prod.yaml`)
```yaml
app:
  mode: "release"  # 生产模式
  enable_pprof: false
  enable_rate_limit: true
database:
  max_open_conns: 50  # 更高的连接池配置
log:
  level: "info"  # 减少日志量
  path: "/var/log/template"  # 使用系统日志目录
```

## 故障排查

如果配置相关问题导致部署失败：

1. 检查`deploy.conf`中的`CONFIG_FILE_PATH`值
2. 确认指定的配置文件存在且格式正确
3. 使用`-c`参数明确指定配置文件路径
4. 查看服务器日志输出以确认配置加载错误

通过遵循这些指南，您可以安全有效地管理不同环境的配置，确保应用程序在各种环境中正常运行，同时保护敏感信息安全。 