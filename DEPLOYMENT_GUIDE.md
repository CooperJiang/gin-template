# 🚀 完整部署指南

> 前后端一体化部署的生产级配置指南

## 📋 目录

- [部署架构](#部署架构)
- [部署前检查](#部署前检查)
- [安全配置](#安全配置)
- [环境变量配置](#环境变量配置)
- [前后端一体化构建](#前后端一体化构建)
- [Docker部署](#docker部署)
- [服务器部署](#服务器部署)
- [故障排除](#故障排除)

---

## 部署架构

### 前后端一体化方案

```
┌─────────────────────────────────────────┐
│     单一 Go 二进制文件（template）        │
├─────────────────────────────────────────┤
│  Go Embed (编译时嵌入)                    │
│  ├─ internal/static/web/                │
│  │   ├─ index.html (主应用)              │
│  │   ├─ assets/*.js, *.css              │
│  │   └─ subapps/                        │
│  │       ├─ app/                        │
│  │       │   ├─ index.html              │
│  │       │   └─ assets/                 │
│  │       └─ [其他子应用]/                │
│  └─ Gin 路由处理                         │
│      ├─ /api/v1/*  → 后端 API           │
│      ├─ /subapps/:app/*  → 子应用资源   │
│      └─ /*  → SPA fallback (主应用)     │
└─────────────────────────────────────────┘
```

**优点**：
- ✅ 单一二进制文件，部署简单
- ✅ 无需单独的静态文件服务器
- ✅ 前后端版本一致性保证
- ✅ 零停机部署支持

---

## 部署前检查

### ✅ 完整检查清单

```bash
# 1. 代码质量检查
cd web
pnpm lint          # 前端代码检查
pnpm type-check    # TypeScript 类型检查
pnpm format:check  # 代码格式检查

cd ..
go vet ./...       # Go 代码检查
go test ./...      # 运行测试

# 2. 前端构建验证
make web-build
ls -la internal/static/web/index.html  # 应该存在
ls -la internal/static/web/assets/     # 应该有 JS 和 CSS 文件
ls -la internal/static/web/subapps/app/  # 子应用应该存在

# 3. 后端编译验证
make build
ls -lh bin/template  # 应该 > 20MB（包含前端资源）

# 4. 配置文件检查
grep "\${" config.prod.yaml  # 应该有环境变量占位符
```

### 🔍 关键文件验证

| 文件/目录 | 检查项 | 预期结果 |
|----------|--------|----------|
| `internal/static/web/index.html` | 存在性 | 必须存在 |
| `internal/static/web/assets/*.js` | 文件大小 | > 100KB |
| `bin/template` | 二进制大小 | > 20MB（包含前端）|
| `config.prod.yaml` | 敏感信息 | 无明文密码 |
| `.env.production` | 敏感信息 | 已配置，不提交Git |

---

## 安全配置

### 🔐 必须配置的环境变量

#### 1. 生成强密钥

```bash
# 生成 JWT 密钥（至少32字符）
openssl rand -base64 48

# 生成数据库密码
openssl rand -base64 24

# 生成管理员密码
openssl rand -base64 16
```

#### 2. 创建 .env.production

```bash
# 复制示例文件
cp .env.production.example .env.production

# 编辑并填入实际值
vi .env.production
```

**.env.production 最小配置**:

```bash
# === 必须配置 ===
APP_DEFAULT_ROOT_PASS=<强密码>
DB_USERNAME=<数据库用户>
DB_PASSWORD=<数据库密码>
JWT_SECRET_KEY=<至少32字符的随机字符串>

# === 根据实际情况配置 ===
DB_HOST=localhost
DB_NAME=template
APP_APP_PORT=7500
APP_APP_MODE=release
```

#### 3. 文件权限设置

```bash
# .env.production 仅所有者可读
chmod 600 .env.production

# 配置文件权限
chmod 644 config.prod.yaml

# 二进制文件可执行
chmod 755 bin/template
```

### 🚫 安全检查

```bash
# 确保敏感文件不在 Git 中
git status --ignored | grep -E "\.env\.production|\.key|\.pem"

# 检查配置文件没有明文密码
grep -r "password.*:" config.prod.yaml | grep -v "\${" && echo "❌ 发现明文密码！"

# 检查 JWT 密钥不是默认值
grep "CHANGE-TO" config.prod.yaml && echo "❌ 发现默认密钥！"
```

---

## 环境变量配置

### 加载顺序（优先级从高到低）

```
1. 环境变量（APP_*）
2. .env.production（如果使用 godotenv）
3. config.prod.yaml
4. 代码默认值
```

### 完整环境变量列表

| 环境变量 | 对应配置 | 默认值 | 必填 |
|---------|---------|--------|------|
| `APP_DEFAULT_ROOT_PASS` | app.defaultRootPass | - | ✅ |
| `JWT_SECRET_KEY` | jwt.secret_key | - | ✅ |
| `DB_HOST` | database.host | localhost | ✅ |
| `DB_USERNAME` | database.username | - | ✅ |
| `DB_PASSWORD` | database.password | - | ✅ |
| `DB_NAME` | database.name | template | ❌ |
| `REDIS_HOST` | redis.host | "" (禁用) | ❌ |
| `REDIS_PASSWORD` | redis.password | "" | ❌ |
| `APP_APP_PORT` | app.port | 7500 | ❌ |
| `APP_APP_MODE` | app.mode | release | ❌ |
| `CORS_ENABLED` | cors.enabled | false | ❌ |

### 使用方式

**方式1：导出环境变量（推荐）**

```bash
export APP_DEFAULT_ROOT_PASS="your_secure_password"
export JWT_SECRET_KEY=$(openssl rand -base64 48)
export DB_USERNAME="prod_user"
export DB_PASSWORD="prod_password"

./bin/template --config=config.prod.yaml
```

**方式2：一次性设置**

```bash
APP_DEFAULT_ROOT_PASS="password" \
JWT_SECRET_KEY="..." \
DB_USERNAME="user" \
DB_PASSWORD="pass" \
./bin/template --config=config.prod.yaml
```

**方式3：使用 systemd（推荐生产环境）**

```ini
# /etc/systemd/system/template.service
[Service]
EnvironmentFile=/opt/myapp/shared/config/.env.production
ExecStart=/opt/myapp/current/template --config=/opt/myapp/shared/config/config.prod.yaml
```

---

## 前后端一体化构建

### 构建流程

```bash
# 完整构建流程
make build-deploy

# 等价于：
#  1. make web-build    (构建前端 → internal/static/web)
#  2. go build          (编译 Go + 嵌入前端)
#  3. tar.gz 打包       (生成部署包)
```

### 详细步骤

#### 1. 前端构建

```bash
cd web

# 构建主应用
cd packages/main
pnpm build  # 输出到 dist/

# 构建所有子应用
cd ../app
pnpm build  # 输出到 dist/

# 复制到后端
cd ../..
./scripts/build_web.sh
```

**build_web.sh 做的事情**:

1. 解析 `micro-app-registry.ts` 找到所有子应用
2. 构建主应用 → `internal/static/web/`
3. 构建每个子应用 → `internal/static/web/subapps/<app>/`
4. 验证关键文件存在

#### 2. 后端编译

```bash
# 开发环境（包含调试信息）
go build -o bin/template cmd/main.go

# 生产环境（优化 + 压缩）
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 \
go build -ldflags="-s -w" -o bin/template cmd/main.go
```

**编译标志说明**:
- `-ldflags="-s -w"`: 移除调试信息，减小体积
- `CGO_ENABLED=1`: 启用CGO（SQLite需要）
- `GOOS=linux`: 目标操作系统
- `GOARCH=amd64`: 目标架构

#### 3. 验证嵌入

```bash
# 检查二进制文件大小（应该 > 20MB）
ls -lh bin/template

# 验证前端资源已嵌入
strings bin/template | grep "index.html"  # 应该有输出

# 运行测试
./bin/template --config=config.prod.yaml &
sleep 2
curl http://localhost:7500/  # 应该返回 HTML
curl http://localhost:7500/api/v1/health  # 应该返回 JSON
killall template
```

---

## Docker部署

### 完整 Dockerfile（包含前端构建）

```dockerfile
# ===== 阶段1：构建前端 =====
FROM node:18-alpine AS frontend-builder

WORKDIR /app/web
COPY web/package.json web/pnpm-lock.yaml ./
COPY web/pnpm-workspace.yaml ./
COPY web/packages/ ./packages/
COPY web/shared/ ./shared/
COPY web/core/ ./core/
COPY web/auth/ ./auth/
COPY web/template/ ./template/
COPY web/scripts/ ./scripts/

# 安装依赖并构建
RUN npm install -g pnpm@latest && \
    pnpm install --frozen-lockfile && \
    pnpm build

# ===== 阶段2：编译后端 =====
FROM golang:1.24-alpine AS backend-builder

WORKDIR /app

# 安装构建依赖
RUN apk add --no-cache gcc musl-dev

# 复制 Go 模块文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 复制前端构建产物
COPY --from=frontend-builder /app/web/packages/main/dist ./internal/static/web/
COPY --from=frontend-builder /app/web/packages/app/dist ./internal/static/web/subapps/app/

# 验证前端资源
RUN test -f internal/static/web/index.html || (echo "❌ 前端构建失败" && exit 1)

# 编译 Go 应用
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-s -w" -o template ./cmd

# ===== 阶段3：运行时镜像 =====
FROM alpine:3.17

# 安装运行时依赖
RUN apk add --no-cache ca-certificates tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

# 创建非root用户
RUN adduser -D -H -h /app appuser

WORKDIR /app

# 复制二进制文件
COPY --from=backend-builder /app/template ./
COPY --from=backend-builder /app/config.prod.yaml ./config.yaml

# 设置权限
RUN chown -R appuser:appuser /app && \
    chmod 755 ./template && \
    chmod 600 ./config.yaml

USER appuser

EXPOSE 7500

# 健康检查
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:7500/api/v1/health || exit 1

CMD ["./template", "--config=config.yaml"]
```

### docker-compose.yml（生产）

```yaml
version: '3.8'

services:
  app:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "7500:7500"
    environment:
      - APP_DEFAULT_ROOT_PASS=${APP_DEFAULT_ROOT_PASS}
      - JWT_SECRET_KEY=${JWT_SECRET_KEY}
      - DB_HOST=mysql
      - DB_USERNAME=${DB_USERNAME}
      - DB_PASSWORD=${DB_PASSWORD}
      - DB_NAME=${DB_NAME}
      - REDIS_HOST=redis
      - REDIS_PASSWORD=${REDIS_PASSWORD}
    depends_on:
      - mysql
      - redis
    networks:
      - app-network
    restart: unless-stopped
    volumes:
      - ./uploads:/app/uploads
      - ./logs:/app/logs

  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: ${MYSQL_ROOT_PASSWORD}
      MYSQL_DATABASE: ${DB_NAME}
      MYSQL_USER: ${DB_USERNAME}
      MYSQL_PASSWORD: ${DB_PASSWORD}
    volumes:
      - mysql-data:/var/lib/mysql
    networks:
      - app-network
    restart: unless-stopped

  redis:
    image: redis:7-alpine
    command: redis-server --requirepass ${REDIS_PASSWORD}
    volumes:
      - redis-data:/data
    networks:
      - app-network
    restart: unless-stopped

networks:
  app-network:
    driver: bridge

volumes:
  mysql-data:
  redis-data:
```

### 使用 Docker 部署

```bash
# 1. 创建 .env 文件
cat > .env <<EOF
APP_DEFAULT_ROOT_PASS=$(openssl rand -base64 16)
JWT_SECRET_KEY=$(openssl rand -base64 48)
DB_USERNAME=template_user
DB_PASSWORD=$(openssl rand -base64 24)
DB_NAME=template
MYSQL_ROOT_PASSWORD=$(openssl rand -base64 24)
REDIS_PASSWORD=$(openssl rand -base64 16)
EOF

# 2. 构建并启动
docker-compose up -d

# 3. 查看日志
docker-compose logs -f app

# 4. 健康检查
curl http://localhost:7500/api/v1/health
curl http://localhost:7500/
```

---

## 服务器部署

### 使用 Makefile 自动化部署

```bash
# 1. 配置服务器信息
make deploy-config

# 2. 初始化服务器环境（首次）
make deploy-setup

# 3. 构建并部署
make build-deploy

# 4. 验证部署
make deploy-status

# 5. 如有问题，一键回滚
make deploy-rollback
```

### 手动部署步骤

#### 1. 准备服务器

```bash
# 安装依赖（CentOS/RHEL）
sudo yum install -y gcc

# 安装依赖（Ubuntu/Debian）
sudo apt-get update
sudo apt-get install -y build-essential

# 创建应用目录
sudo mkdir -p /opt/myapp
sudo chown $USER:$USER /opt/myapp
```

#### 2. 上传部署包

```bash
# 本地构建
make build-deploy

# 上传到服务器
scp release/template_prod_package.tar.gz user@server:/opt/myapp/
```

#### 3. 解压并配置

```bash
# 在服务器上
cd /opt/myapp
tar -xzf template_prod_package.tar.gz

# 创建配置文件
cat > .env.production <<EOF
APP_DEFAULT_ROOT_PASS=your_password
JWT_SECRET_KEY=your_jwt_secret
DB_HOST=localhost
DB_USERNAME=db_user
DB_PASSWORD=db_password
DB_NAME=template
EOF

chmod 600 .env.production
```

#### 4. 配置 systemd 服务

```bash
sudo tee /etc/systemd/system/template.service > /dev/null <<EOF
[Unit]
Description=Template Application
After=network.target mysql.service

[Service]
Type=simple
User=$USER
WorkingDirectory=/opt/myapp
EnvironmentFile=/opt/myapp/.env.production
ExecStart=/opt/myapp/template --config=/opt/myapp/config.yaml
Restart=on-failure
RestartSec=5s

[Install]
WantedBy=multi-user.target
EOF

# 重新加载 systemd
sudo systemctl daemon-reload

# 启动服务
sudo systemctl start template

# 设置开机自启
sudo systemctl enable template

# 查看状态
sudo systemctl status template
```

#### 5. 配置 Nginx 反向代理

```nginx
server {
    listen 80;
    server_name yourdomain.com;

    # 前端资源（主应用和子应用）
    location / {
        proxy_pass http://localhost:7500;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    # API 请求
    location /api/ {
        proxy_pass http://localhost:7500;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    # 静态资源缓存
    location ~* \.(js|css|png|jpg|jpeg|gif|ico|svg|woff|woff2)$ {
        proxy_pass http://localhost:7500;
        expires 1y;
        add_header Cache-Control "public, immutable";
    }

    # Gzip 压缩
    gzip on;
    gzip_types text/plain text/css application/json application/javascript text/xml application/xml application/xml+rss text/javascript;
}
```

---

## 故障排除

### 常见问题

#### 1. 前端资源404

**症状**: 访问 `/` 返回空白页，控制台显示 JS/CSS 404

**原因**: 前端没有正确构建或嵌入

**解决**:

```bash
# 验证前端已构建
ls -la internal/static/web/index.html

# 重新构建
make web-build

# 验证嵌入
strings bin/template | grep "index.html"

# 重新编译
make build
```

#### 2. API 返回 HTML 而不是 JSON

**症状**: API 请求返回 HTML（SPA fallback）

**原因**: API 路由没有正确注册，被 NoRoute 捕获

**解决**:

```go
// 检查路由注册（internal/routes/routes.go）
// 确保 API 路由在 NoRoute 之前注册
```

#### 3. 子应用加载失败

**症状**: 主应用可以访问，但子应用404

**原因**: 子应用没有构建或注册

**解决**:

```bash
# 验证子应用已构建
ls -la internal/static/web/subapps/app/

# 检查注册表
cat web/packages/main/src/micro-app-registry.ts

# 重新构建
make web-build
```

#### 4. 数据库连接失败

**症状**: 应用启动失败，日志显示数据库连接错误

**解决**:

```bash
# 检查环境变量
printenv | grep DB_

# 检查数据库是否运行
mysql -h localhost -u $DB_USERNAME -p$DB_PASSWORD

# 检查防火墙
sudo firewall-cmd --list-ports
```

#### 5. JWT 验证失败

**症状**: 登录后立即提示未授权

**原因**: JWT 密钥配置错误或不一致

**解决**:

```bash
# 确保 JWT_SECRET_KEY 已设置
echo $JWT_SECRET_KEY

# 检查密钥长度（至少32字符）
echo $JWT_SECRET_KEY | wc -c

# 重新生成密钥
export JWT_SECRET_KEY=$(openssl rand -base64 48)
```

---

## 监控和日志

### 日志配置

```yaml
# config.prod.yaml
log:
  level: "info"              # debug/info/warn/error
  format: "json"             # json/text
  output: "logs/app.log"     # 日志文件路径
  max_size: 100              # MB
  max_backups: 10            # 保留文件数
  max_age: 30                # 保留天数
```

### 查看日志

```bash
# systemd 日志
sudo journalctl -u template -f

# 应用日志
tail -f /opt/myapp/logs/app.log

# 错误日志
grep ERROR /opt/myapp/logs/app.log | tail -20
```

### 性能监控

```bash
# CPU 和内存使用
top -p $(pgrep template)

# 连接数
netstat -an | grep :7500 | wc -l

# 请求日志分析
awk '{print $6}' access.log | sort | uniq -c | sort -rn | head -10
```

---

## 备份和恢复

### 数据库备份

```bash
# 备份
mysqldump -u $DB_USERNAME -p$DB_PASSWORD $DB_NAME > backup_$(date +%Y%m%d).sql

# 恢复
mysql -u $DB_USERNAME -p$DB_PASSWORD $DB_NAME < backup_20240220.sql
```

### 应用备份

```bash
# 备份配置和数据
tar -czf backup_$(date +%Y%m%d).tar.gz \
    /opt/myapp/*.yaml \
    /opt/myapp/.env.production \
    /opt/myapp/uploads \
    /opt/myapp/*.db
```

---

## 总结

### ✅ 部署检查清单

- [ ] 前端已构建且验证通过
- [ ] 后端已编译且包含前端资源（> 20MB）
- [ ] 所有敏感信息已移到环境变量
- [ ] `.env.production` 已配置且权限正确（600）
- [ ] 数据库已创建并可连接
- [ ] JWT 密钥已生成（至少32字符）
- [ ] 防火墙已配置（开放端口7500）
- [ ] Nginx 反向代理已配置（可选）
- [ ] systemd 服务已配置并启动
- [ ] 健康检查通过
- [ ] 日志正常输出
- [ ] 备份计划已设置

### 📞 获取帮助

- 查看 [TROUBLESHOOTING.md](./TROUBLESHOOTING.md)
- 查看 [web/QUICK_START.md](./web/QUICK_START.md)
- 检查应用日志：`sudo journalctl -u template -f`

---

**记住**：生产环境安全第一！永远不要在配置文件中硬编码敏感信息。
