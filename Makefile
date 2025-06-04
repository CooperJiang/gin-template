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

# 设置默认目标为help
.DEFAULT_GOAL := help

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

.PHONY: build
build: frontend-build ## 🔨 构建生产部署版本 (Linux + 部署包)
	@echo "$(GREEN)构建生产部署版本...$(RESET)"
	@# 清理所有旧的构建文件
	@rm -rf release/ 
	@# 创建release目录结构
	@mkdir -p release/linux_amd64
	@# 构建Linux版本（包含前端静态文件）
	@GOOS=linux GOARCH=amd64 go build -ldflags="-X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME)" -o release/linux_amd64/$(APP_NAME) $(MAIN_PATH)
	@# 配置文件处理
	@CONFIG_SRC="config.yaml"; \
	if [ -n "$(CONFIG_FILE_PATH)" ] && [ -f "$(CONFIG_FILE_PATH)" ]; then \
		CONFIG_SRC="$(CONFIG_FILE_PATH)"; \
	elif [ -f "config.$(DEPLOY_ENV).yaml" ]; then \
		CONFIG_SRC="config.$(DEPLOY_ENV).yaml"; \
	fi; \
	cp "$$CONFIG_SRC" release/linux_amd64/config.yaml
	@# 创建启动脚本
	@echo "#!/bin/bash" > release/linux_amd64/run.sh
	@echo "./$(APP_NAME)" >> release/linux_amd64/run.sh
	@chmod +x release/linux_amd64/run.sh release/linux_amd64/$(APP_NAME)
	@# 创建部署包
	@echo "$(BLUE)创建部署包...$(RESET)"
	@cd release && tar -czf $(APP_NAME)_$(DEPLOY_ENV)_package.tar.gz -C linux_amd64 .
	@# 清理中间目录，只保留部署包
	@rm -rf release/linux_amd64
	@echo "$(GREEN)生产部署版本构建完成!$(RESET)"
	@echo "$(BLUE)部署包: ./release/$(APP_NAME)_$(DEPLOY_ENV)_package.tar.gz$(RESET)"
	@echo "$(YELLOW)💡 可直接用于生产部署$(RESET)"

.PHONY: build-release
build-release: frontend-build ## 🚀 构建生产部署版本 (与 build 相同)
	@echo "$(GREEN)构建生产部署版本...$(RESET)"
	@# 清理所有旧的构建文件
	@rm -rf release/
	@# 创建release目录结构
	@mkdir -p release/linux_amd64
	@# 构建Linux版本（包含前端静态文件）
	@GOOS=linux GOARCH=amd64 go build -ldflags="-X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME)" -o release/linux_amd64/$(APP_NAME) $(MAIN_PATH)
	@# 配置文件处理
	@CONFIG_SRC="config.yaml"; \
	if [ -n "$(CONFIG_FILE_PATH)" ] && [ -f "$(CONFIG_FILE_PATH)" ]; then \
		CONFIG_SRC="$(CONFIG_FILE_PATH)"; \
	elif [ -f "config.$(DEPLOY_ENV).yaml" ]; then \
		CONFIG_SRC="config.$(DEPLOY_ENV).yaml"; \
	fi; \
	cp "$$CONFIG_SRC" release/linux_amd64/config.yaml
	@# 创建启动脚本
	@echo "#!/bin/bash" > release/linux_amd64/run.sh
	@echo "./$(APP_NAME)" >> release/linux_amd64/run.sh
	@chmod +x release/linux_amd64/run.sh release/linux_amd64/$(APP_NAME)
	@# 创建部署包
	@echo "$(BLUE)创建部署包...$(RESET)"
	@cd release && tar -czf $(APP_NAME)_$(DEPLOY_ENV)_package.tar.gz -C linux_amd64 .
	@# 清理中间目录，只保留部署包
	@rm -rf release/linux_amd64
	@echo "$(GREEN)生产部署版本构建完成!$(RESET)"
	@echo "$(BLUE)部署包: ./release/$(APP_NAME)_$(DEPLOY_ENV)_package.tar.gz$(RESET)"
	@echo "$(YELLOW)💡 可直接用于生产部署$(RESET)"

.PHONY: start
start: ## 🚀 一键启动三端开发环境 (系统终端模式)
	@echo "$(GREEN)🚀 启动三端开发环境 (系统终端)...$(RESET)"
	@echo "$(BLUE)正在打开三个终端窗口...$(RESET)"
	@# 停止可能已存在的服务
	@pkill -f "go run cmd/main.go" || true
	@pkill -f "npm run dev" || true
	@sleep 1
	@# 获取项目目录路径
	@PROJECT_DIR=$$(pwd); \
	echo "$(BLUE)🔧 启动后端 (Go) - 端口 9000$(RESET)"; \
	osascript -e "tell application \"Terminal\" to do script \"cd $$PROJECT_DIR && echo '🔧 启动后端服务器...' && make dev\""; \
	sleep 1; \
	echo "$(BLUE)📱 启动管理端 (Vue3) - 端口 3000$(RESET)"; \
	osascript -e "tell application \"Terminal\" to do script \"cd $$PROJECT_DIR && echo '📱 启动管理端...' && make admin-dev\""; \
	sleep 1; \
	echo "$(BLUE)🌐 启动用户端 (Vue3) - 端口 4000$(RESET)"; \
	osascript -e "tell application \"Terminal\" to do script \"cd $$PROJECT_DIR && echo '🌐 启动用户端...' && make web-dev\""; \
	echo ""
	@echo "$(GREEN)✅ 三个终端窗口已打开！$(RESET)"
	@echo ""
	@echo "$(GREEN)📱 服务访问地址:$(RESET)"
	@echo "  🎯 后端API:  http://localhost:9000"
	@echo "  📱 管理端:   http://localhost:3000"
	@echo "  🌐 用户端:   http://localhost:4000"
	@echo ""
	@echo "$(YELLOW)💡 停止服务: 在各个终端窗口中按 Ctrl+C，或运行 make stop$(RESET)"

.PHONY: start-bg
start-bg: ## 🚀 一键启动三端开发环境 (后台模式)
	@echo "$(GREEN)🚀 启动三端开发环境 (后台模式)...$(RESET)"
	@echo "$(BLUE)正在启动所有服务...$(RESET)"
	@# 停止可能已存在的服务
	@pkill -f "go run cmd/main.go" || true
	@pkill -f "npm run dev" || true
	@sleep 1
	@# 启动后端 (后台)
	@echo "$(BLUE)🔧 启动后端 (Go) - 端口 9000$(RESET)"
	@mkdir -p logs pids
	@nohup go run cmd/main.go > logs/backend.log 2>&1 & echo $$! > pids/backend.pid
	@sleep 2
	@# 启动管理端 (后台)
	@echo "$(BLUE)📱 启动管理端 (Vue3) - 端口 3000$(RESET)"
	@cd admin && nohup npm run dev > ../logs/admin.log 2>&1 & echo $$! > ../pids/admin.pid; cd ..
	@sleep 2
	@# 启动用户端 (后台)
	@echo "$(BLUE)🌐 启动用户端 (Vue3) - 端口 4000$(RESET)"
	@cd web && nohup npm run dev > ../logs/web.log 2>&1 & echo $$! > ../pids/web.pid; cd ..
	@sleep 2
	@echo ""
	@echo "$(GREEN)✅ 三端开发环境已启动！$(RESET)"
	@echo ""
	@echo "$(GREEN)📱 服务访问地址:$(RESET)"
	@echo "  🎯 后端API:  http://localhost:9000"
	@echo "  📱 管理端:   http://localhost:3000"
	@echo "  🌐 用户端:   http://localhost:4000"
	@echo ""
	@echo "$(YELLOW)💡 日志文件:$(RESET)"
	@echo "  后端: logs/backend.log"
	@echo "  管理端: logs/admin.log" 
	@echo "  用户端: logs/web.log"
	@echo ""
	@echo "$(YELLOW)💡 停止服务: make stop$(RESET)"

.PHONY: stop
stop: ## 🛑 停止所有开发服务
	@echo "$(YELLOW)🛑 停止所有开发服务...$(RESET)"
	@# 使用PID文件停止服务
	@if [ -f pids/backend.pid ]; then kill $$(cat pids/backend.pid) 2>/dev/null || true; rm -f pids/backend.pid; fi
	@if [ -f pids/admin.pid ]; then kill $$(cat pids/admin.pid) 2>/dev/null || true; rm -f pids/admin.pid; fi
	@if [ -f pids/web.pid ]; then kill $$(cat pids/web.pid) 2>/dev/null || true; rm -f pids/web.pid; fi
	@# 备用清理
	@pkill -f "go run cmd/main.go" || true
	@pkill -f "npm run dev" || true
	@sleep 1
	@echo "$(GREEN)✅ 所有服务已停止$(RESET)"

.PHONY: status
status: ## 📊 查看服务状态
	@echo "$(BLUE)📊 服务状态:$(RESET)"
	@echo ""
	@# 检查端口占用
	@echo "$(YELLOW)端口状态:$(RESET)"
	@lsof -i :9000 2>/dev/null | head -2 || echo "  端口 9000: 未占用"
	@lsof -i :3000 2>/dev/null | head -2 || echo "  端口 3000: 未占用"
	@lsof -i :4000 2>/dev/null | head -2 || echo "  端口 4000: 未占用"
	@echo ""
	@# 检查PID文件
	@echo "$(YELLOW)PID文件:$(RESET)"
	@if [ -f pids/backend.pid ]; then echo "  后端PID: $$(cat pids/backend.pid)"; else echo "  后端PID: 不存在"; fi
	@if [ -f pids/admin.pid ]; then echo "  管理端PID: $$(cat pids/admin.pid)"; else echo "  管理端PID: 不存在"; fi
	@if [ -f pids/web.pid ]; then echo "  用户端PID: $$(cat pids/web.pid)"; else echo "  用户端PID: 不存在"; fi

.PHONY: logs
logs: ## 📋 查看服务日志
	@echo "$(BLUE)📋 服务日志:$(RESET)"
	@echo ""
	@echo "$(YELLOW)=== 后端日志 (最后10行) ===$(RESET)"
	@tail -n 10 logs/backend.log 2>/dev/null || echo "后端日志文件不存在"
	@echo ""
	@echo "$(YELLOW)=== 管理端日志 (最后10行) ===$(RESET)"
	@tail -n 10 logs/admin.log 2>/dev/null || echo "管理端日志文件不存在"
	@echo ""
	@echo "$(YELLOW)=== 用户端日志 (最后10行) ===$(RESET)"
	@tail -n 10 logs/web.log 2>/dev/null || echo "用户端日志文件不存在"

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

.PHONY: build-backend
build-backend: clean ## 仅构建后端应用程序
	@echo "$(GREEN)构建后端应用程序...$(RESET)"
	@mkdir -p bin
	@go build -ldflags="-X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME)" -o $(BINARY_PATH) $(MAIN_PATH)
	@echo "$(GREEN)后端构建完成: $(BINARY_PATH)$(RESET)"

.PHONY: build-linux
build-linux: clean ## 构建Linux版本
	@mkdir -p bin
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME)" -o $(BINARY_PATH)-linux $(MAIN_PATH)
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
# 部署命令
# ========================================

# 部署配置
DEPLOY_CONFIG := deploy.conf
DEPLOY_HOST := 
DEPLOY_USER := root
DEPLOY_PORT := 22
DEPLOY_DIR := /opt/myapp
DEPLOY_KEY := 
DEPLOY_ENV := prod
KEEP_RELEASES := 5
CONFIG_FILE_PATH := 

-include $(DEPLOY_CONFIG)

.PHONY: deploy-config
deploy-config: ## 配置部署参数
	@echo "$(BLUE)配置部署参数$(RESET)"
	@mkdir -p scripts
	@if [ ! -f scripts/deploy_functions.sh ]; then \
		echo "$(RED)错误: scripts/deploy_functions.sh 不存在$(RESET)"; \
		exit 1; \
	fi
	@chmod +x scripts/deploy_functions.sh
	@./scripts/deploy_functions.sh config $(APP_NAME)

.PHONY: deploy-env-config
deploy-env-config: ## 配置环境特定参数
	@echo "$(BLUE)配置环境特定参数$(RESET)"
	@if [ ! -f scripts/deploy_functions.sh ]; then \
		echo "$(RED)错误: scripts/deploy_functions.sh 不存在$(RESET)"; \
		exit 1; \
	fi
	@chmod +x scripts/deploy_functions.sh
	@./scripts/deploy_functions.sh config_env

.PHONY: deploy-check
deploy-check: ## 检查部署配置
	@echo "$(BLUE)当前部署配置:$(RESET)"
	@echo "  应用名称: $(APP_NAME)"
	@echo "  远程主机: $(DEPLOY_HOST)"
	@echo "  远程用户: $(DEPLOY_USER)"
	@echo "  SSH端口: $(DEPLOY_PORT)"
	@echo "  部署目录: $(DEPLOY_DIR)"
	@echo "  SSH密钥: $(DEPLOY_KEY)"
	@echo "  部署环境: $(DEPLOY_ENV)"
	@echo "  保留版本: $(KEEP_RELEASES)"
	@if [ -n "$(CONFIG_FILE_PATH)" ]; then \
		echo "  配置文件: $(CONFIG_FILE_PATH)"; \
	fi
	@if [ -z "$(DEPLOY_HOST)" ]; then \
		echo "$(RED)错误: 未配置远程主机地址$(RESET)"; \
		echo "$(YELLOW)请运行 'make deploy-config' 配置部署参数$(RESET)"; \
		exit 1; \
	fi

.PHONY: deploy-setup
deploy-setup: ## 初始化服务器环境
	@if [ -z "$(DEPLOY_HOST)" ]; then echo "$(RED)请先运行 make deploy-config$(RESET)"; exit 1; fi
	@echo "$(YELLOW)设置服务器环境...$(RESET)"
	@chmod +x scripts/deploy_functions.sh
	@SSH_CMD=$$(./scripts/deploy_functions.sh build_ssh "$(DEPLOY_KEY)" "$(DEPLOY_PORT)" "$(DEPLOY_USER)" "$(DEPLOY_HOST)"); \
	$$SSH_CMD "mkdir -p $(DEPLOY_DIR)/{releases,shared,shared/config,tmp}"
	@echo "$(BLUE)生成服务脚本...$(RESET)"
	@./scripts/deploy_functions.sh generate_service $(APP_NAME) $(DEPLOY_DIR) /tmp/service_script.sh
	@SCP_CMD=$$(./scripts/deploy_functions.sh build_scp "$(DEPLOY_KEY)" "$(DEPLOY_PORT)"); \
	$$SCP_CMD /tmp/service_script.sh $(DEPLOY_USER)@$(DEPLOY_HOST):$(DEPLOY_DIR)/service.sh
	@SSH_CMD=$$(./scripts/deploy_functions.sh build_ssh "$(DEPLOY_KEY)" "$(DEPLOY_PORT)" "$(DEPLOY_USER)" "$(DEPLOY_HOST)"); \
	$$SSH_CMD "chmod +x $(DEPLOY_DIR)/service.sh"
	@rm -f /tmp/service_script.sh
	@echo "$(GREEN)服务器环境设置完成$(RESET)"

.PHONY: deploy
deploy: deploy-check ## 🚀 快速切换部署（2-3秒停机时间）
	@echo "$(BLUE)开始应用部署...$(RESET)"
	@chmod +x scripts/deploy_functions.sh
	@# 检查必要参数
	@if ! ./scripts/deploy_functions.sh check_params "$(DEPLOY_HOST)" "$(DEPLOY_USER)"; then \
		exit 1; \
	fi
	@# 检查部署包是否存在
	@EXPECTED_PACKAGE="./release/$(APP_NAME)_$(DEPLOY_ENV)_package.tar.gz"; \
	if [ ! -f "$$EXPECTED_PACKAGE" ]; then \
		echo "$(RED)错误: 部署包不存在: $$EXPECTED_PACKAGE$(RESET)"; \
		echo "$(YELLOW)提示: 请先运行 'make build' 或 'make build-release'$(RESET)"; \
		exit 1; \
	fi; \
	echo "$(GREEN)找到部署包: $$EXPECTED_PACKAGE$(RESET)"; \
	echo "$$EXPECTED_PACKAGE" > .tmp_package_path
	@# 检查远程环境
	@echo "$(BLUE)检查远程环境...$(RESET)"
	@SSH_CMD=$$(./scripts/deploy_functions.sh build_ssh "$(DEPLOY_KEY)" "$(DEPLOY_PORT)" "$(DEPLOY_USER)" "$(DEPLOY_HOST)"); \
	if ! $$SSH_CMD "[ -d \"$(DEPLOY_DIR)\" ]" 2>/dev/null; then \
		echo "$(RED)错误: 远程目录 $(DEPLOY_DIR) 不存在$(RESET)"; \
		echo "$(YELLOW)提示: 请先运行 'make deploy-setup' 初始化服务器环境$(RESET)"; \
		exit 1; \
	fi; \
	if ! $$SSH_CMD "[ -f \"$(DEPLOY_DIR)/service.sh\" ]" 2>/dev/null; then \
		echo "$(RED)错误: 服务管理脚本不存在$(RESET)"; \
		echo "$(YELLOW)提示: 请先运行 'make deploy-setup' 初始化服务器环境$(RESET)"; \
		exit 1; \
	fi
	@# 检查端口占用
	@echo "$(BLUE)检查端口占用...$(RESET)"
	@CONFIG_SRC="config.yaml"; \
	if [ -n "$(CONFIG_FILE_PATH)" ] && [ -f "$(CONFIG_FILE_PATH)" ]; then \
		CONFIG_SRC="$(CONFIG_FILE_PATH)"; \
	elif [ -f "config.$(DEPLOY_ENV).yaml" ]; then \
		CONFIG_SRC="config.$(DEPLOY_ENV).yaml"; \
	fi; \
	if ! ./scripts/deploy_functions.sh check_port "$(DEPLOY_KEY)" "$(DEPLOY_PORT)" "$(DEPLOY_USER)" "$(DEPLOY_HOST)" "$(DEPLOY_DIR)" "$(APP_NAME)" "$$CONFIG_SRC"; then \
		echo "$(RED)端口检查失败，部署终止$(RESET)"; \
		exit 1; \
	fi
	@# 准备环境变量
	@echo "$(BLUE)准备环境变量...$(RESET)"
	@./scripts/deploy_functions.sh prepare_env "$(DEPLOY_KEY)" "$(DEPLOY_PORT)" "$(DEPLOY_USER)" "$(DEPLOY_HOST)" "$(DEPLOY_DIR)" "$(DEPLOY_ENV)" "$(VERSION)" || true
	@# 执行部署流程
	@echo "$(BLUE)开始应用部署流程...$(RESET)"
	@PACKAGE_PATH=$$(cat .tmp_package_path); \
	TIMESTAMP=$$(date +%Y%m%d%H%M%S); \
	RELEASE_DIR="$(DEPLOY_DIR)/releases/$$TIMESTAMP"; \
	SSH_CMD=$$(./scripts/deploy_functions.sh build_ssh "$(DEPLOY_KEY)" "$(DEPLOY_PORT)" "$(DEPLOY_USER)" "$(DEPLOY_HOST)"); \
	SCP_CMD=$$(./scripts/deploy_functions.sh build_scp "$(DEPLOY_KEY)" "$(DEPLOY_PORT)"); \
	\
	echo "$(BLUE)步骤1: 创建版本目录...$(RESET)"; \
	$$SSH_CMD "mkdir -p $$RELEASE_DIR"; \
	\
	echo "$(BLUE)步骤2: 上传应用包...$(RESET)"; \
	$$SCP_CMD "$$PACKAGE_PATH" "$(DEPLOY_USER)@$(DEPLOY_HOST):$(DEPLOY_DIR)/tmp/package.tar.gz"; \
	\
	echo "$(BLUE)步骤3: 解压应用包...$(RESET)"; \
	$$SSH_CMD "cd $$RELEASE_DIR && tar -xzf $(DEPLOY_DIR)/tmp/package.tar.gz 2>/dev/null"; \
	$$SSH_CMD "chmod +x $$RELEASE_DIR/$(APP_NAME) $$RELEASE_DIR/run.sh"; \
	\
	echo "$(BLUE)步骤4: 配置应用...$(RESET)"; \
	$$SSH_CMD "mkdir -p $(DEPLOY_DIR)/shared/config && cp $$RELEASE_DIR/config.yaml $(DEPLOY_DIR)/shared/config/"; \
	$$SSH_CMD "ln -sf $(DEPLOY_DIR)/shared/config/config.yaml $$RELEASE_DIR/config.yaml"; \
	\
	echo "$(BLUE)步骤5: 切换应用版本...$(RESET)"; \
	echo "$(YELLOW)⚡ 正在切换应用版本（预计停机2-3秒）...$(RESET)"; \
	$$SSH_CMD "$(DEPLOY_DIR)/service.sh stop || true; ln -sfn $$RELEASE_DIR $(DEPLOY_DIR)/current; $(DEPLOY_DIR)/service.sh start"; \
	sleep 2; \
	\
	echo "$(BLUE)步骤6: 验证应用状态...$(RESET)"; \
	if $$SSH_CMD "$(DEPLOY_DIR)/service.sh status" >/dev/null 2>&1; then \
		echo "$(GREEN)✅ 应用部署成功！$(RESET)"; \
	else \
		echo "$(RED)❌ 应用启动异常，请检查日志$(RESET)"; \
	fi; \
	\
	echo "$(BLUE)步骤7: 清理旧版本...$(RESET)"; \
	$$SSH_CMD "cd $(DEPLOY_DIR)/releases && ls -t | tail -n +$$(($(KEEP_RELEASES)+1)) | xargs rm -rf 2>/dev/null || true"; \
	$$SSH_CMD "rm -f $(DEPLOY_DIR)/tmp/package.tar.gz"; \
	\
	echo "$(GREEN)应用部署完成！$(RESET)"; \
	$$SSH_CMD "$(DEPLOY_DIR)/service.sh status" || echo "$(YELLOW)状态检查失败，请手动检查$(RESET)"; \
	echo "$(BLUE)部署信息:$(RESET)"; \
	echo "  服务器: $(DEPLOY_HOST)"; \
	echo "  目录: $(DEPLOY_DIR)"; \
	echo "  环境: $(DEPLOY_ENV)"; \
	echo "  版本: $$TIMESTAMP"; \
	rm -f .tmp_package_path

.PHONY: deploy-rollback
deploy-rollback: ## 回滚到上一个版本
	@if [ -z "$(DEPLOY_HOST)" ]; then echo "$(RED)请先运行 make deploy-config$(RESET)"; exit 1; fi
	@echo "$(YELLOW)正在回滚到上一个版本...$(RESET)"
	@SSH_CMD=$$(./scripts/deploy_functions.sh build_ssh "$(DEPLOY_KEY)" "$(DEPLOY_PORT)" "$(DEPLOY_USER)" "$(DEPLOY_HOST)"); \
	CURRENT_VERSION=$$( $$SSH_CMD "readlink $(DEPLOY_DIR)/current | xargs basename" ); \
	PREV_VERSION=$$( $$SSH_CMD "ls -t $(DEPLOY_DIR)/releases | sed -n '2p'" ); \
	if [ -z "$$PREV_VERSION" ]; then \
		echo "$(RED)错误: 没有找到可回滚的版本$(RESET)"; \
		exit 1; \
	fi; \
	echo "当前版本: $$CURRENT_VERSION"; \
	echo "回滚到版本: $$PREV_VERSION"; \
	echo "$(BLUE)停止当前应用...$(RESET)"; \
	$$SSH_CMD "$(DEPLOY_DIR)/service.sh stop"; \
	echo "$(BLUE)更新当前版本链接...$(RESET)"; \
	$$SSH_CMD "ln -sfn $(DEPLOY_DIR)/releases/$$PREV_VERSION $(DEPLOY_DIR)/current"; \
	echo "$(BLUE)启动应用...$(RESET)"; \
	$$SSH_CMD "$(DEPLOY_DIR)/service.sh start"; \
	echo "$(BLUE)应用状态:$(RESET)"; \
	$$SSH_CMD "$(DEPLOY_DIR)/service.sh status"; \
	echo "$(GREEN)回滚完成!$(RESET)"

.PHONY: deploy-restart
deploy-restart: ## 重启远程应用
	@if [ -z "$(DEPLOY_HOST)" ]; then echo "$(RED)请先运行 make deploy-config$(RESET)"; exit 1; fi
	@echo "$(YELLOW)正在重启应用...$(RESET)"
	@SSH_CMD=$$(./scripts/deploy_functions.sh build_ssh "$(DEPLOY_KEY)" "$(DEPLOY_PORT)" "$(DEPLOY_USER)" "$(DEPLOY_HOST)"); \
	$$SSH_CMD "$(DEPLOY_DIR)/service.sh restart"

.PHONY: deploy-status
deploy-status: ## 查看应用状态
	@if [ -z "$(DEPLOY_HOST)" ]; then echo "$(RED)请先运行 make deploy-config$(RESET)"; exit 1; fi
	@SSH_CMD=$$(./scripts/deploy_functions.sh build_ssh "$(DEPLOY_KEY)" "$(DEPLOY_PORT)" "$(DEPLOY_USER)" "$(DEPLOY_HOST)"); \
	if ! $$SSH_CMD "[ -f \"$(DEPLOY_DIR)/service.sh\" ]" 2>/dev/null; then \
		echo "$(RED)错误: 服务管理脚本不存在$(RESET)"; \
		echo "$(YELLOW)提示: 请先运行 'make deploy-setup' 初始化服务器环境$(RESET)"; \
		exit 1; \
	fi; \
	echo "$(YELLOW)应用状态:$(RESET)"; \
	$$SSH_CMD "$(DEPLOY_DIR)/service.sh status"

.PHONY: deploy-logs
deploy-logs: ## 查看应用日志
	@if [ -z "$(DEPLOY_HOST)" ]; then echo "$(RED)请先运行 make deploy-config$(RESET)"; exit 1; fi
	@SSH_CMD=$$(./scripts/deploy_functions.sh build_ssh "$(DEPLOY_KEY)" "$(DEPLOY_PORT)" "$(DEPLOY_USER)" "$(DEPLOY_HOST)"); \
	echo "$(YELLOW)应用日志 (最后50行):$(RESET)"; \
	$$SSH_CMD "$(DEPLOY_DIR)/service.sh logs $${LINES:-50}"

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
	@echo "$(GREEN)🚀 快速开始:$(RESET)"
	@awk 'BEGIN {FS = ":.*?## "} /^start:.*?## / {printf "  $(GREEN)%-20s$(RESET) %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@awk 'BEGIN {FS = ":.*?## "} /^start-bg:.*?## / {printf "  $(GREEN)%-20s$(RESET) %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo ""
	@echo "$(GREEN)🔧 可用命令:$(RESET)"
	@echo ""
	@echo "$(YELLOW)🚀 启动命令:$(RESET)"
	@awk 'BEGIN {FS = ":.*?## "} /^(start|start-bg|stop|status|logs):.*?## / {printf "  $(YELLOW)%-20s$(RESET) %s\n", $$1, $$2}' $(MAKEFILE_LIST)
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
	@awk 'BEGIN {FS = ":.*?## "} /^(dev|build-backend|build-linux|clean):.*?## / {printf "  $(YELLOW)%-20s$(RESET) %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo ""
	@echo "$(YELLOW)🔨 构建命令:$(RESET)"
	@awk 'BEGIN {FS = ":.*?## "} /^build:.*?## / {printf "  $(YELLOW)%-20s$(RESET) %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@awk 'BEGIN {FS = ":.*?## "} /^build-release:.*?## / {printf "  $(YELLOW)%-20s$(RESET) %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo ""
	@echo "$(YELLOW)🧪 测试命令:$(RESET)"
	@awk 'BEGIN {FS = ":.*?## "} /^test[a-zA-Z_-]*:.*?## / {printf "  $(YELLOW)%-20s$(RESET) %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo ""
	@echo "$(YELLOW)🚢 部署命令:$(RESET)"
	@awk 'BEGIN {FS = ":.*?## "} /^deploy[a-zA-Z_-]*:.*?## / {printf "  $(YELLOW)%-20s$(RESET) %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@awk 'BEGIN {FS = ":.*?## "} /^build-deploy:.*?## / {printf "  $(YELLOW)%-20s$(RESET) %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo ""
	@echo "$(GREEN)📋 部署流程:$(RESET)"
	@echo "  $(YELLOW)1. make deploy-config$(RESET)       # 配置部署参数"
	@echo "  $(YELLOW)2. make deploy-env-config$(RESET)   # 配置环境参数 (可选)"
	@echo "  $(YELLOW)3. make deploy-setup$(RESET)        # 初始化服务器环境"
	@echo "  $(YELLOW)4. make build-deploy$(RESET)        # 一键构建并部署"
	@echo "  $(YELLOW)或分步操作:$(RESET)"
	@echo "  $(YELLOW)4a. make build$(RESET)              # 构建生产版本 (推荐)"
	@echo "  $(YELLOW)4b. make deploy$(RESET)             # 部署应用"
	@echo ""
	@echo "$(GREEN)🔄 部署管理:$(RESET)"
	@echo "  $(YELLOW)make deploy-status$(RESET)          # 查看应用状态"
	@echo "  $(YELLOW)make deploy-logs$(RESET)            # 查看应用日志"
	@echo "  $(YELLOW)make deploy-restart$(RESET)         # 重启应用"
	@echo "  $(YELLOW)make deploy-rollback$(RESET)        # 回滚到上一版本"
	@echo ""
	@echo "$(GREEN)💡 常用示例:$(RESET)"
	@echo "  $(YELLOW)make start$(RESET)              # 🚀 一键启动三端开发环境 (系统终端)"
	@echo "  $(YELLOW)make start-bg$(RESET)           # 🚀 一键启动三端开发环境 (后台模式)"
	@echo "  $(YELLOW)make stop$(RESET)               # 🛑 停止所有开发服务"
	@echo "  $(YELLOW)make status$(RESET)             # 📊 查看服务状态"
	@echo "  $(YELLOW)make logs$(RESET)               # 📋 查看服务日志"
	@echo "  $(YELLOW)make build$(RESET)              # 🔨 构建生产部署版本 (主要命令)"
	@echo "  $(YELLOW)make deploy-config$(RESET)      # 配置部署参数"
	@echo "  $(YELLOW)make build-deploy$(RESET)       # 🔨 一键构建并部署 (推荐)"
	@echo "  $(YELLOW)make deploy$(RESET)             # 🚀 部署应用 (使用现有包)"
	@echo "  $(YELLOW)make deploy-status$(RESET)      # 查看远程应用状态"
	@echo "  $(YELLOW)make deploy-rollback$(RESET)    # 回滚到上一版本"

.PHONY: build-deploy
build-deploy: build-release deploy ## 🔨 一键构建并快速部署 (完整流程)
	@echo "$(GREEN)🎉 构建并快速部署完成!$(RESET)"
