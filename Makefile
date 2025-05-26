# 项目配置
APP_NAME := template
VERSION := $(shell git describe --tags --always --dirty)
BUILD_TIME := $(shell date +%Y-%m-%d_%H:%M:%S)
GO_VERSION := $(shell go version | awk '{print $$3}')

# 构建配置
BINARY_NAME := $(APP_NAME)
BINARY_PATH := ./bin/$(BINARY_NAME)
MAIN_PATH := ./cmd/main.go

# Docker配置
DOCKER_IMAGE := $(APP_NAME):$(VERSION)
DOCKER_REGISTRY := your-registry.com

# 颜色输出
RED := \033[31m
GREEN := \033[32m
YELLOW := \033[33m
BLUE := \033[34m
RESET := \033[0m

.PHONY: help
help: ## 显示帮助信息
	@echo "$(BLUE)$(APP_NAME) 开发工具$(RESET)"
	@echo ""
	@echo "$(GREEN)可用命令:$(RESET)"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  $(YELLOW)%-15s$(RESET) %s\n", $$1, $$2}' $(MAKEFILE_LIST)

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

.PHONY: fmt
fmt: ## 格式化代码
	@echo "$(GREEN)格式化代码...$(RESET)"
	@go fmt ./...
	@goimports -w .

.PHONY: lint
lint: ## 代码检查
	@echo "$(GREEN)运行代码检查...$(RESET)"
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "$(RED)golangci-lint 未安装$(RESET)"; \
		echo "$(YELLOW)安装命令: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest$(RESET)"; \
	fi

.PHONY: security
security: ## 安全检查
	@echo "$(GREEN)运行安全检查...$(RESET)"
	@if command -v gosec > /dev/null; then \
		gosec ./...; \
	else \
		echo "$(RED)gosec 未安装$(RESET)"; \
		echo "$(YELLOW)安装命令: go install github.com/securego/gosec/v2/cmd/gosec@latest$(RESET)"; \
	fi

.PHONY: deps
deps: ## 检查依赖
	@echo "$(GREEN)检查依赖...$(RESET)"
	@go mod tidy
	@go mod verify
	@if command -v govulncheck > /dev/null; then \
		govulncheck ./...; \
	else \
		echo "$(YELLOW)govulncheck 未安装，跳过漏洞检查$(RESET)"; \
		echo "$(YELLOW)安装命令: go install golang.org/x/vuln/cmd/govulncheck@latest$(RESET)"; \
	fi

.PHONY: migrate
migrate: ## 执行数据库迁移
	@echo "$(GREEN)执行数据库迁移...$(RESET)"
	@go run $(MAIN_PATH) migrate

.PHONY: migrate-down
migrate-down: ## 回滚数据库迁移
	@echo "$(YELLOW)回滚数据库迁移...$(RESET)"
	@go run $(MAIN_PATH) migrate-down

.PHONY: migration
migration: ## 创建新的迁移文件 (使用: make migration name=add_user_avatar)
	@if [ -z "$(name)" ]; then \
		echo "$(RED)请提供迁移名称: make migration name=your_migration_name$(RESET)"; \
		exit 1; \
	fi
	@echo "$(GREEN)创建迁移文件: $(name)$(RESET)"
	@mkdir -p internal/migrations
	@timestamp=$$(date +%Y%m%d%H%M%S); \
	filename="internal/migrations/$${timestamp}_$(name).sql"; \
	echo "-- +migrate Up" > $$filename; \
	echo "-- 在这里添加你的SQL语句" >> $$filename; \
	echo "" >> $$filename; \
	echo "-- +migrate Down" >> $$filename; \
	echo "-- 在这里添加回滚SQL语句" >> $$filename; \
	echo "$(GREEN)迁移文件已创建: $$filename$(RESET)"

.PHONY: db-reset
db-reset: ## 重置数据库
	@echo "$(RED)重置数据库...$(RESET)"
	@read -p "确定要重置数据库吗？这将删除所有数据 [y/N]: " confirm; \
	if [ "$$confirm" = "y" ] || [ "$$confirm" = "Y" ]; then \
		rm -f app.db; \
		echo "$(GREEN)数据库已重置$(RESET)"; \
	else \
		echo "$(YELLOW)操作已取消$(RESET)"; \
	fi

.PHONY: new-module
new-module: ## 创建新模块 (使用: make new-module name=product)
	@if [ -z "$(name)" ]; then \
		echo "$(RED)请提供模块名称: make new-module name=your_module_name$(RESET)"; \
		exit 1; \
	fi
	@echo "$(GREEN)创建新模块: $(name)$(RESET)"
	@./scripts/create_module.sh $(name)

.PHONY: docs
docs: ## 生成API文档
	@echo "$(GREEN)生成API文档...$(RESET)"
	@if command -v swag > /dev/null; then \
		swag init -g cmd/main.go -o docs/swagger; \
		echo "$(GREEN)API文档已生成: docs/swagger/$(RESET)"; \
	else \
		echo "$(RED)swag 未安装$(RESET)"; \
		echo "$(YELLOW)安装命令: go install github.com/swaggo/swag/cmd/swag@latest$(RESET)"; \
	fi

.PHONY: docker-build
docker-build: ## 构建Docker镜像
	@echo "$(GREEN)构建Docker镜像...$(RESET)"
	@docker build -t $(DOCKER_IMAGE) .
	@echo "$(GREEN)Docker镜像构建完成: $(DOCKER_IMAGE)$(RESET)"

.PHONY: docker-run
docker-run: ## 运行Docker容器
	@echo "$(GREEN)运行Docker容器...$(RESET)"
	@docker run -p 8080:8080 --name $(APP_NAME) $(DOCKER_IMAGE)

.PHONY: docker-push
docker-push: docker-build ## 推送Docker镜像到仓库
	@echo "$(GREEN)推送Docker镜像...$(RESET)"
	@docker tag $(DOCKER_IMAGE) $(DOCKER_REGISTRY)/$(DOCKER_IMAGE)
	@docker push $(DOCKER_REGISTRY)/$(DOCKER_IMAGE)

.PHONY: docker-clean
docker-clean: ## 清理Docker资源
	@echo "$(YELLOW)清理Docker资源...$(RESET)"
	@docker stop $(APP_NAME) 2>/dev/null || true
	@docker rm $(APP_NAME) 2>/dev/null || true
	@docker rmi $(DOCKER_IMAGE) 2>/dev/null || true

.PHONY: install-tools
install-tools: ## 安装开发工具
	@echo "$(GREEN)安装开发工具...$(RESET)"
	@go install github.com/air-verse/air@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/securego/gosec/v2/cmd/gosec@latest
	@go install golang.org/x/vuln/cmd/govulncheck@latest
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go install golang.org/x/tools/cmd/goimports@latest
	@echo "$(GREEN)开发工具安装完成$(RESET)"

.PHONY: setup
setup: install-tools ## 初始化开发环境
	@echo "$(GREEN)初始化开发环境...$(RESET)"
	@go mod tidy
	@cp config.example.yaml config.yaml
	@echo "$(GREEN)开发环境初始化完成$(RESET)"
	@echo "$(YELLOW)请编辑 config.yaml 文件配置你的环境$(RESET)"

.PHONY: run
run: build ## 构建并运行应用程序
	@echo "$(GREEN)运行应用程序...$(RESET)"
	@$(BINARY_PATH)

.PHONY: deploy-dev
deploy-dev: ## 部署到开发环境
	@echo "$(GREEN)部署到开发环境...$(RESET)"
	@./scripts/deploy.sh dev

.PHONY: deploy-prod
deploy-prod: ## 部署到生产环境
	@echo "$(GREEN)部署到生产环境...$(RESET)"
	@./scripts/deploy.sh prod

.PHONY: logs
logs: ## 查看应用日志
	@echo "$(GREEN)查看应用日志...$(RESET)"
	@tail -f logs/app.log

.PHONY: status
status: ## 显示项目状态
	@echo "$(BLUE)项目状态$(RESET)"
	@echo "应用名称: $(APP_NAME)"
	@echo "版本: $(VERSION)"
	@echo "Go版本: $(GO_VERSION)"
	@echo "构建时间: $(BUILD_TIME)"
	@echo ""
	@echo "$(GREEN)Git状态:$(RESET)"
	@git status --short

# ========================================
# 新增的工具命令
# ========================================

.PHONY: perf-check
perf-check: ## 运行性能检查
	@echo "$(GREEN)运行性能检查...$(RESET)"
	@./scripts/performance_check.sh

.PHONY: perf-check-url
perf-check-url: ## 运行性能检查 (自定义URL: make perf-check-url url=http://localhost:9000)
	@if [ -z "$(url)" ]; then \
		echo "$(RED)请提供URL: make perf-check-url url=http://localhost:9000$(RESET)"; \
		exit 1; \
	fi
	@echo "$(GREEN)运行性能检查 (URL: $(url))...$(RESET)"
	@./scripts/performance_check.sh -u $(url)

.PHONY: db-backup
db-backup: ## 备份数据库
	@echo "$(GREEN)备份数据库...$(RESET)"
	@./scripts/db_manager.sh backup

.PHONY: db-restore
db-restore: ## 恢复数据库
	@echo "$(GREEN)恢复数据库...$(RESET)"
	@./scripts/db_manager.sh restore

.PHONY: db-list
db-list: ## 列出数据库备份
	@echo "$(GREEN)列出数据库备份...$(RESET)"
	@./scripts/db_manager.sh list

.PHONY: db-cleanup
db-cleanup: ## 清理旧备份 (使用: make db-cleanup days=7)
	@echo "$(GREEN)清理旧备份...$(RESET)"
	@if [ -n "$(days)" ]; then \
		./scripts/db_manager.sh cleanup $(days); \
	else \
		./scripts/db_manager.sh cleanup; \
	fi

.PHONY: db-info
db-info: ## 显示数据库信息
	@echo "$(GREEN)显示数据库信息...$(RESET)"
	@./scripts/db_manager.sh info

.PHONY: db-optimize
db-optimize: ## 优化数据库
	@echo "$(GREEN)优化数据库...$(RESET)"
	@./scripts/db_manager.sh optimize

.PHONY: db-export
db-export: ## 导出数据库数据
	@echo "$(GREEN)导出数据库数据...$(RESET)"
	@./scripts/db_manager.sh export

.PHONY: db-sql
db-sql: ## 执行SQL查询 (使用: make db-sql file=query.sql 或直接 make db-sql 进入交互模式)
	@echo "$(GREEN)执行SQL查询...$(RESET)"
	@if [ -n "$(file)" ]; then \
		./scripts/db_manager.sh sql $(file); \
	else \
		./scripts/db_manager.sh sql; \
	fi

.PHONY: benchmark
benchmark: ## 运行性能基准测试
	@echo "$(GREEN)运行性能基准测试...$(RESET)"
	@go test -bench=. -benchmem ./...

.PHONY: profile
profile: ## 生成性能分析报告
	@echo "$(GREEN)生成性能分析报告...$(RESET)"
	@go test -cpuprofile=cpu.prof -memprofile=mem.prof -bench=. ./...
	@echo "$(YELLOW)CPU分析: go tool pprof cpu.prof$(RESET)"
	@echo "$(YELLOW)内存分析: go tool pprof mem.prof$(RESET)"

.PHONY: check-deps
check-deps: ## 检查依赖更新
	@echo "$(GREEN)检查依赖更新...$(RESET)"
	@go list -u -m all

.PHONY: update-deps
update-deps: ## 更新依赖
	@echo "$(GREEN)更新依赖...$(RESET)"
	@go get -u ./...
	@go mod tidy

.PHONY: generate
generate: ## 运行代码生成
	@echo "$(GREEN)运行代码生成...$(RESET)"
	@go generate ./...

.PHONY: mock
mock: ## 生成Mock文件
	@echo "$(GREEN)生成Mock文件...$(RESET)"
	@if command -v mockgen > /dev/null; then \
		find . -name "*.go" -exec grep -l "//go:generate mockgen" {} \; | xargs -I {} go generate {}; \
		echo "$(GREEN)Mock文件生成完成$(RESET)"; \
	else \
		echo "$(RED)mockgen 未安装$(RESET)"; \
		echo "$(YELLOW)安装命令: go install github.com/golang/mock/mockgen@latest$(RESET)"; \
	fi

.PHONY: swagger
swagger: ## 生成Swagger文档
	@echo "$(GREEN)生成Swagger文档...$(RESET)"
	@if command -v swag > /dev/null; then \
		swag init -g cmd/main.go -o docs/swagger --parseDependency --parseInternal; \
		echo "$(GREEN)Swagger文档已生成: docs/swagger/$(RESET)"; \
		echo "$(YELLOW)访问地址: http://localhost:8080/swagger/index.html$(RESET)"; \
	else \
		echo "$(RED)swag 未安装$(RESET)"; \
		echo "$(YELLOW)安装命令: go install github.com/swaggo/swag/cmd/swag@latest$(RESET)"; \
	fi

.PHONY: api-test
api-test: ## 运行API测试
	@echo "$(GREEN)运行API测试...$(RESET)"
	@if [ -f "tests/api_test.sh" ]; then \
		./tests/api_test.sh; \
	else \
		echo "$(YELLOW)API测试脚本不存在: tests/api_test.sh$(RESET)"; \
	fi

.PHONY: check-config
check-config: ## 检查配置文件
	@echo "$(GREEN)检查配置文件...$(RESET)"
	@if [ -f "config.yaml" ]; then \
		echo "$(GREEN)✓ 配置文件存在$(RESET)"; \
		if command -v yq > /dev/null; then \
			yq eval '.' config.yaml > /dev/null && echo "$(GREEN)✓ 配置文件格式正确$(RESET)" || echo "$(RED)✗ 配置文件格式错误$(RESET)"; \
		else \
			echo "$(YELLOW)⚠ yq 未安装，无法验证YAML格式$(RESET)"; \
		fi; \
	else \
		echo "$(RED)✗ 配置文件不存在$(RESET)"; \
		echo "$(YELLOW)请运行: cp config.example.yaml config.yaml$(RESET)"; \
	fi

.PHONY: clean-all
clean-all: clean docker-clean ## 清理所有构建文件和Docker资源
	@echo "$(YELLOW)清理所有文件...$(RESET)"
	@rm -rf coverage.out coverage.html
	@rm -rf *.prof
	@rm -rf docs/swagger
	@rm -rf exports/
	@echo "$(GREEN)清理完成$(RESET)"

.PHONY: quick-start
quick-start: setup migrate dev ## 快速启动项目 (初始化+迁移+开发服务器)
	@echo "$(GREEN)项目快速启动完成!$(RESET)"

.PHONY: full-check
full-check: fmt lint security test test-coverage ## 完整代码检查 (格式化+检查+安全+测试+覆盖率)
	@echo "$(GREEN)完整代码检查完成!$(RESET)"

.PHONY: all
all: clean fmt lint test build ## 执行完整的构建流程

# 默认目标
.DEFAULT_GOAL := help 