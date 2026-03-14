#!/bin/bash

# 前端打包脚本
# 将 web 应用打包后同步到 internal/static/web

set -euo pipefail

# 颜色输出
RED='\033[31m'
GREEN='\033[32m'
YELLOW='\033[33m'
BLUE='\033[34m'
RESET='\033[0m'

# 项目路径
PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
WEB_DIR="${PROJECT_ROOT}/web"
STATIC_DIR="${PROJECT_ROOT}/internal/static"
DIST_DIR="${WEB_DIR}/dist"
TARGET_WEB_DIR="${STATIC_DIR}/web"

echo -e "${BLUE}开始前端项目打包...${RESET}"

# 检查web目录是否存在
if [ ! -d "$WEB_DIR" ]; then
    echo -e "${RED}错误: web目录不存在: $WEB_DIR${RESET}"
    exit 1
fi

# 检查package.json是否存在
if [ ! -f "$WEB_DIR/package.json" ]; then
    echo -e "${RED}错误: package.json不存在: $WEB_DIR/package.json${RESET}"
    exit 1
fi

# 检查pnpm是否可用
if ! command -v pnpm &> /dev/null; then
    echo -e "${RED}错误: pnpm未安装，请先安装: npm install -g pnpm${RESET}"
    exit 1
fi

# 进入web目录
cd "$WEB_DIR"

# 检查是否安装了node_modules
if [ ! -d "node_modules" ]; then
    echo -e "${YELLOW}未找到node_modules，正在安装依赖...${RESET}"
    pnpm install
    echo -e "${GREEN}依赖安装完成${RESET}"
fi

# 清理之前的构建
echo -e "${YELLOW}清理之前的构建文件...${RESET}"
rm -rf "$DIST_DIR"

# 构建前端项目
echo -e "${BLUE}构建前端项目...${RESET}"
pnpm build

# 检查构建结果
if [ ! -d "$DIST_DIR" ]; then
    echo -e "${RED}错误: 构建失败，dist目录不存在${RESET}"
    exit 1
fi

# 创建静态目录（如果不存在）
mkdir -p "$STATIC_DIR"

# 删除旧的静态文件
if [ -d "$TARGET_WEB_DIR" ]; then
    echo -e "${YELLOW}删除旧的静态文件...${RESET}"
    rm -rf "$TARGET_WEB_DIR"
fi

# 复制dist到 static/web
echo -e "${BLUE}复制构建文件到静态目录...${RESET}"
cp -r "$DIST_DIR" "$TARGET_WEB_DIR"

# 显示构建结果
echo -e "${GREEN}前端打包完成!${RESET}"
echo -e "${BLUE}构建信息:${RESET}"
echo -e "  源目录: $DIST_DIR"
echo -e "  目标目录: $TARGET_WEB_DIR"
echo -e "  文件数量: $(find "$TARGET_WEB_DIR" -type f | wc -l | tr -d ' ')"
echo -e "  总大小: $(du -sh "$TARGET_WEB_DIR" | cut -f1)"

# 检查关键文件
if [ -f "$TARGET_WEB_DIR/index.html" ]; then
    echo -e "${GREEN}✓ index.html 存在${RESET}"
else
    echo -e "${RED}✗ index.html 不存在${RESET}"
    exit 1
fi

if [ -d "$TARGET_WEB_DIR/assets" ]; then
    echo -e "${GREEN}✓ assets目录存在${RESET}"
    echo -e "  CSS文件: $(find "$TARGET_WEB_DIR/assets" -name "*.css" | wc -l | tr -d ' ')"
    echo -e "  JS文件: $(find "$TARGET_WEB_DIR/assets" -name "*.js" | wc -l | tr -d ' ')"
else
    echo -e "${YELLOW}⚠ assets目录不存在${RESET}"
fi

echo -e "${GREEN}打包流程完成!${RESET}"
echo -e "${BLUE}现在可以运行 'make build' 来构建包含静态文件的Go二进制文件${RESET}"
