# Gin Template Web Frontend Makefile

.PHONY: help install dev build preview clean lint format type-check test

# 默认目标
help:
	@echo "Gin Template Web Frontend - 可用命令:"
	@echo ""
	@echo "  install      - 安装依赖"
	@echo "  dev          - 启动开发服务器"
	@echo "  build        - 构建生产版本"
	@echo "  preview      - 预览生产版本"
	@echo "  clean        - 清理构建文件"
	@echo "  lint         - 代码检查和修复"
	@echo "  lint-check   - 仅检查代码规范"
	@echo "  format       - 格式化代码"
	@echo "  format-check - 检查代码格式"
	@echo "  type-check   - TypeScript 类型检查"
	@echo "  test         - 运行测试"
	@echo "  setup        - 初始化项目（安装依赖）"
	@echo "  check-all    - 运行所有检查"
	@echo ""

# 安装依赖
install:
	@echo "📦 安装依赖..."
	npm install

# 启动开发服务器
dev:
	@echo "🚀 启动开发服务器..."
	npm run dev

# 构建生产版本
build:
	@echo "🏗️  构建生产版本..."
	npm run build

# 预览生产版本
preview:
	@echo "👀 预览生产版本..."
	npm run preview

# 清理构建文件
clean:
	@echo "🧹 清理构建文件..."
	npm run clean

# 代码检查和修复
lint:
	@echo "🔍 代码检查和修复..."
	npm run lint

# 仅检查代码规范
lint-check:
	@echo "🔍 检查代码规范..."
	npm run lint:check

# 格式化代码
format:
	@echo "💅 格式化代码..."
	npm run format

# 检查代码格式
format-check:
	@echo "💅 检查代码格式..."
	npm run format:check

# TypeScript 类型检查
type-check:
	@echo "🔧 TypeScript 类型检查..."
	npm run type-check

# 运行测试
test:
	@echo "🧪 运行测试..."
	npm run test

# 初始化项目
setup: install
	@echo "✅ 项目初始化完成！"
	@echo ""
	@echo "下一步："
	@echo "  make dev    - 启动开发服务器"
	@echo "  make build  - 构建生产版本"

# 运行所有检查
check-all: lint-check format-check type-check
	@echo "✅ 所有检查完成！"

# 快速开发流程
quick-dev: install dev

# 生产构建流程
production: check-all build
	@echo "🎉 生产版本构建完成！" 