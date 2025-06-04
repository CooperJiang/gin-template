# 项目配置
APP_NAME := template
VERSION := $(shell git describe --tags --always --dirty)
BUILD_TIME := $(shell date +%Y-%m-%d_%H:%M:%S)
GO_VERSION := $(shell go version | awk '{print $$3}')

# 构建配置
BINARY_NAME := $(APP_NAME)
BINARY_PATH := ./bin/$(BINARY_NAME)
MAIN_PATH := ./cmd/main.go

# 前端配置
ADMIN_DIR := ./admin
WEB_DIR := ./web
STATIC_DIR := ./internal/static
ADMIN_DIST_DIR := $(ADMIN_DIR)/dist
WEB_DIST_DIR := $(WEB_DIR)/dist
TARGET_ADMIN_DIR := $(STATIC_DIR)/admin
TARGET_WEB_DIR := $(STATIC_DIR)/web

# Docker配置
DOCKER_IMAGE := $(APP_NAME):$(VERSION)
DOCKER_REGISTRY := your-registry.com

# 颜色输出
RED := \033[31m
GREEN := \033[32m
YELLOW := \033[33m
BLUE := \033[34m
RESET := \033[0m

# ========================================
# 管理端前端命令
# ========================================

.PHONY: admin-install
admin-install: ## 安装管理端前端依赖
	@echo "$(GREEN)安装管理端前端依赖...$(RESET)"
	@if [ ! -d "$(ADMIN_DIR)" ]; then \
		echo "$(RED)错误: admin目录不存在$(RESET)"; \
		exit 1; \
	fi
	@cd $(ADMIN_DIR) && npm install
	@echo "$(GREEN)管理端前端依赖安装完成$(RESET)"

.PHONY: admin-dev
admin-dev: ## 启动管理端前端开发服务器
	@echo "$(GREEN)启动管理端前端开发服务器...$(RESET)"
	@if [ ! -d "$(ADMIN_DIR)" ]; then \
		echo "$(RED)错误: admin目录不存在$(RESET)"; \
		exit 1; \
	fi
	@cd $(ADMIN_DIR) && npm run dev

.PHONY: admin-build
admin-build: ## 构建管理端前端项目并部署到static/admin目录
	@echo "$(GREEN)构建管理端前端项目...$(RESET)"
	@./scripts/build_admin.sh

.PHONY: admin-lint
admin-lint: ## 检查管理端前端代码
	@echo "$(GREEN)检查管理端前端代码...$(RESET)"
	@if [ ! -d "$(ADMIN_DIR)" ]; then \
		echo "$(RED)错误: admin目录不存在$(RESET)"; \
		exit 1; \
	fi
	@cd $(ADMIN_DIR) && npm run lint

.PHONY: admin-format
admin-format: ## 格式化管理端前端代码
	@echo "$(GREEN)格式化管理端前端代码...$(RESET)"
	@if [ ! -d "$(ADMIN_DIR)" ]; then \
		echo "$(RED)错误: admin目录不存在$(RESET)"; \
		exit 1; \
	fi
	@cd $(ADMIN_DIR) && npm run format

.PHONY: admin-type-check
admin-type-check: ## 管理端前端类型检查
	@echo "$(GREEN)运行管理端前端类型检查...$(RESET)"
	@if [ ! -d "$(ADMIN_DIR)" ]; then \
		echo "$(RED)错误: admin目录不存在$(RESET)"; \
		exit 1; \
	fi
	@cd $(ADMIN_DIR) && npm run type-check

.PHONY: admin-clean
admin-clean: ## 清理管理端前端构建文件
	@echo "$(YELLOW)清理管理端前端构建文件...$(RESET)"
	@rm -rf $(ADMIN_DIST_DIR)
	@rm -rf $(TARGET_ADMIN_DIR)
	@if [ -d "$(ADMIN_DIR)" ]; then \
		cd $(ADMIN_DIR) && npm run clean; \
	fi
	@echo "$(GREEN)管理端前端文件清理完成$(RESET)"

# ========================================
# 用户端前端命令
# ========================================

.PHONY: web-install
web-install: ## 安装用户端前端依赖
	@echo "$(GREEN)安装用户端前端依赖...$(RESET)"
	@if [ ! -d "$(WEB_DIR)" ]; then \
		echo "$(RED)错误: web目录不存在$(RESET)"; \
		exit 1; \
	fi
	@cd $(WEB_DIR) && npm install
	@echo "$(GREEN)用户端前端依赖安装完成$(RESET)"

.PHONY: web-dev
web-dev: ## 启动用户端前端开发服务器
	@echo "$(GREEN)启动用户端前端开发服务器...$(RESET)"
	@if [ ! -d "$(WEB_DIR)" ]; then \
		echo "$(RED)错误: web目录不存在$(RESET)"; \
		exit 1; \
	fi
	@cd $(WEB_DIR) && npm run dev

.PHONY: web-build
web-build: ## 构建用户端前端项目并部署到static/web目录
	@echo "$(GREEN)构建用户端前端项目...$(RESET)"
	@./scripts/build_web.sh

.PHONY: web-lint
web-lint: ## 检查用户端前端代码
	@echo "$(GREEN)检查用户端前端代码...$(RESET)"
	@if [ ! -d "$(WEB_DIR)" ]; then \
		echo "$(RED)错误: web目录不存在$(RESET)"; \
		exit 1; \
	fi
	@cd $(WEB_DIR) && npm run lint

.PHONY: web-format
web-format: ## 格式化用户端前端代码
	@echo "$(GREEN)格式化用户端前端代码...$(RESET)"
	@if [ ! -d "$(WEB_DIR)" ]; then \
		echo "$(RED)错误: web目录不存在$(RESET)"; \
		exit 1; \
	fi
	@cd $(WEB_DIR) && npm run format

.PHONY: web-clean
web-clean: ## 清理用户端前端构建文件
	@echo "$(YELLOW)清理用户端前端构建文件...$(RESET)"
	@rm -rf $(WEB_DIST_DIR)
	@rm -rf $(TARGET_WEB_DIR)
	@if [ -d "$(WEB_DIR)" ]; then \
		cd $(WEB_DIR) && npm run clean 2>/dev/null || true; \
	fi
	@echo "$(GREEN)用户端前端文件清理完成$(RESET)"

# ========================================
# 前端组合命令
# ========================================

.PHONY: frontend-install
frontend-install: admin-install web-install ## 安装所有前端依赖

.PHONY: frontend-build
frontend-build: admin-build web-build ## 构建所有前端项目

.PHONY: frontend-clean
frontend-clean: admin-clean web-clean ## 清理所有前端构建文件

.PHONY: frontend-lint
frontend-lint: admin-lint web-lint ## 检查所有前端代码

.PHONY: frontend-format
frontend-format: admin-format web-format ## 格式化所有前端代码

# ========================================
# 全栈构建命令
# ========================================

.PHONY: fullstack-build
fullstack-build: frontend-build build ## 构建完整的全栈应用 (管理端+用户端+后端)
	@echo "$(GREEN)全栈应用构建完成!$(RESET)"
	@echo "$(BLUE)二进制文件: $(BINARY_PATH)$(RESET)"
	@echo "$(BLUE)管理端: /admin 路径访问$(RESET)"
	@echo "$(BLUE)用户端: / 根路径访问$(RESET)"
	@echo "$(BLUE)静态文件已嵌入到二进制文件中$(RESET)"

.PHONY: fullstack-dev
fullstack-dev: ## 启动全栈开发环境 (并行启动前后端)
	@echo "$(GREEN)启动全栈开发环境...$(RESET)"
	@echo "$(YELLOW)后端将在 :9000 端口启动$(RESET)"
	@echo "$(YELLOW)管理端将在 :3000 端口启动$(RESET)"
	@echo "$(YELLOW)用户端将在 :4000 端口启动$(RESET)"
	@echo "$(BLUE)按 Ctrl+C 停止所有服务$(RESET)"
	@trap 'kill 0' INT; \
	make admin-dev & \
	make web-dev & \
	make dev & \
	wait

.PHONY: fullstack-clean
fullstack-clean: clean frontend-clean ## 清理所有构建文件 (前端+后端)
	@echo "$(GREEN)全栈清理完成$(RESET)"

# ========================================
# 后端构建命令
# ========================================

.PHONY: dev
dev: ## 启动开发服务器（热重载）
	@echo "$(GREEN)启动开发服务器...$(RESET)"
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "$(RED)air 未安装，请运行: go install github.com/air-verse/air@latest$(RESET)"; \
		echo "$(YELLOW)使用普通模式启动...$(RESET)"; \
		go run $(MAIN_PATH); \
	fi

.PHONY: build
build: clean ## 构建应用程序
	@echo "$(GREEN)构建应用程序...$(RESET)"
	@mkdir -p bin
	@go build -ldflags="-X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME)" -o $(BINARY_PATH) $(MAIN_PATH)
	@echo "$(GREEN)构建完成: $(BINARY_PATH)$(RESET)"

.PHONY: build-linux
build-linux: clean ## 构建Linux版本
	@echo "$(GREEN)构建Linux版本...$(RESET)"
	@mkdir -p bin
	@GOOS=linux GOARCH=amd64 go build -ldflags="-X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME)" -o $(BINARY_PATH)-linux $(MAIN_PATH)
	@echo "$(GREEN)构建完成: $(BINARY_PATH)-linux$(RESET)"

.PHONY: clean
clean: ## 清理构建文件
	@echo "$(YELLOW)清理构建文件...$(RESET)"
	@rm -rf bin/
	@rm -rf dist/
	@go clean

# ========================================
# 测试命令
# ========================================

.PHONY: test
test: ## 运行所有测试
	@echo "$(GREEN)运行测试...$(RESET)"
	@go test -v ./...

.PHONY: test-unit
test-unit: ## 运行单元测试
	@echo "$(GREEN)运行单元测试...$(RESET)"
	@go test -v ./tests/unit/...

.PHONY: test-integration
test-integration: ## 运行集成测试
	@echo "$(GREEN)运行集成测试...$(RESET)"
	@go test -v ./tests/integration/...

.PHONY: test-coverage
test-coverage: ## 运行测试并生成覆盖率报告
	@echo "$(GREEN)生成测试覆盖率报告...$(RESET)"
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)覆盖率报告已生成: coverage.html$(RESET)"

# ========================================
# 代码质量命令
# ========================================

.PHONY: fmt
fmt: ## 格式化代码
	@echo "$(GREEN)格式化代码...$(RESET)"
	@go fmt ./...
	@goimports -w .

.PHONY: lint
lint: ## 检查代码
	@echo "$(GREEN)检查代码...$(RESET)"
	@golangci-lint run

# ========================================
# 兼容性命令 (保持向后兼容)
# ========================================

# 保留旧的命令名以确保向后兼容
.PHONY: web-check
web-check: admin-type-check admin-lint ## 完整前端代码检查 (兼容性保留)

# ========================================
# 帮助命令
# ========================================

.PHONY: help
help: ## 显示帮助信息
	@echo "$(BLUE)$(APP_NAME) 三端开发工具$(RESET)"
	@echo ""
	@echo "$(GREEN)📊 项目架构:$(RESET)"
	@echo "  🎯 后端 (Go)     - 端口 9000"
	@echo "  🎨 管理端 (Vue3) - 端口 3000 (开发) / /admin (生产)"  
	@echo "  👥 用户端 (Vue3) - 端口 4000 (开发) / / (生产)"
	@echo ""
	@echo "$(GREEN)🔧 可用命令:$(RESET)"
	@echo ""
	@echo "$(YELLOW)📱 管理端命令:$(RESET)"
	@awk 'BEGIN {FS = ":.*?## "} /^admin-[a-zA-Z_-]+:.*?## / {printf "  $(YELLOW)%-20s$(RESET) %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo ""
	@echo "$(YELLOW)🌐 用户端命令:$(RESET)"
	@awk 'BEGIN {FS = ":.*?## "} /^web-[a-zA-Z_-]+:.*?## / {printf "  $(YELLOW)%-20s$(RESET) %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo ""
	@echo "$(YELLOW)🔗 前端组合命令:$(RESET)"
	@awk 'BEGIN {FS = ":.*?## "} /^frontend-[a-zA-Z_-]+:.*?## / {printf "  $(YELLOW)%-20s$(RESET) %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo ""
	@echo "$(YELLOW)🚀 全栈命令:$(RESET)"
	@awk 'BEGIN {FS = ":.*?## "} /^fullstack-[a-zA-Z_-]+:.*?## / {printf "  $(YELLOW)%-20s$(RESET) %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo ""
	@echo "$(YELLOW)⚙️ 后端命令:$(RESET)"
	@awk 'BEGIN {FS = ":.*?## "} /^(dev|build|build-linux|clean):.*?## / {printf "  $(YELLOW)%-20s$(RESET) %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo ""
	@echo "$(YELLOW)🧪 测试命令:$(RESET)"
	@awk 'BEGIN {FS = ":.*?## "} /^test[a-zA-Z_-]*:.*?## / {printf "  $(YELLOW)%-20s$(RESET) %s\n", $$1, $$2}' $(MAKEFILE_LIST) 