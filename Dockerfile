# 构建阶段：使用官方Go镜像
FROM golang:1.20-alpine AS builder

# 设置工作目录
WORKDIR /app

# 安装编译依赖
RUN apk add --no-cache gcc musl-dev

# 复制go.mod和go.sum文件
COPY go.mod go.sum ./

# 下载所有依赖
RUN go mod download

# 复制源代码
COPY . .

# 编译应用
RUN CGO_ENABLED=1 GOOS=linux go build -a -o app ./cmd

# 运行阶段：使用轻量级alpine镜像
FROM alpine:3.17

# 安装运行时依赖
RUN apk add --no-cache ca-certificates tzdata

# 设置时区为上海
ENV TZ=Asia/Shanghai

# 创建非root用户
RUN adduser -D -H -h /app appuser

# 创建工作目录
WORKDIR /app

# 从构建阶段复制编译好的二进制文件
COPY --from=builder /app/app /app/
COPY --from=builder /app/config.yaml /app/

# 设置目录权限
RUN chown -R appuser:appuser /app && \
    chmod +x /app/app

# 切换到非root用户
USER appuser

# 暴露应用端口
EXPOSE 9000

# 运行应用
CMD ["./app"] 