#!/bin/bash

# 管理端前端打包脚本
# 用于打包Vue管理端项目并将生成的dist文件夹放置到Go项目的static/admin目录

set -e  # 遇到错误立即退出

# ===========================================
# 配置变量 - 可根据需要修改
# ===========================================
ENABLE_BACKUP=false  # 是否启用备份功能 (true/false)

# 颜色输出
RED='\033[31m'
GREEN='\033[32m'
YELLOW='\033[33m'
BLUE='\033[34m'
RESET='\033[0m'

# 项目路径
PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
ADMIN_DIR="${PROJECT_ROOT}/admin"
STATIC_DIR="${PROJECT_ROOT}/internal/static"
DIST_DIR="${ADMIN_DIR}/dist"
TARGET_ADMIN_DIR="${STATIC_DIR}/admin"

echo -e "${BLUE}开始管理端前端项目打包...${RESET}"
echo -e "${BLUE}备份功能: $([ "$ENABLE_BACKUP" = "true" ] && echo "启用" || echo "禁用")${RESET}"

# 检查admin目录是否存在
if [ ! -d "$ADMIN_DIR" ]; then
    echo -e "${RED}错误: admin目录不存在: $ADMIN_DIR${RESET}"
    exit 1
fi

# 检查package.json是否存在
if [ ! -f "$ADMIN_DIR/package.json" ]; then
    echo -e "${RED}错误: package.json不存在: $ADMIN_DIR/package.json${RESET}"
    exit 1
fi

# 进入admin目录
cd "$ADMIN_DIR"

# 检查是否安装了node_modules
if [ ! -d "node_modules" ]; then
    echo -e "${YELLOW}未找到node_modules，正在安装依赖...${RESET}"
    
    # 检查npm是否可用
    if ! command -v npm &> /dev/null; then
        echo -e "${RED}错误: npm未安装，请先安装Node.js和npm${RESET}"
        exit 1
    fi
    
    npm install
    echo -e "${GREEN}依赖安装完成${RESET}"
fi

# 清理之前的构建
echo -e "${YELLOW}清理之前的构建文件...${RESET}"
rm -rf "$DIST_DIR"

# 运行类型检查
# echo -e "${BLUE}运行类型检查...${RESET}"
# npm run type-check

# 构建管理端项目
echo -e "${BLUE}更新版本信息...${RESET}"
npm run version-update

echo -e "${BLUE}构建管理端项目...${RESET}"
npm run build-only

# 检查构建是否成功
if [ ! -d "$DIST_DIR" ]; then
    echo -e "${RED}错误: 构建失败，dist目录不存在${RESET}"
    exit 1
fi

# 创建静态目录（如果不存在）
mkdir -p "$STATIC_DIR"

# 备份旧的admin目录（如果存在且启用备份）
if [ -d "$TARGET_ADMIN_DIR" ] && [ "$ENABLE_BACKUP" = "true" ]; then
    echo -e "${YELLOW}备份旧的静态文件...${RESET}"
    BACKUP_DIR="${STATIC_DIR}/admin_backup_$(date +%Y%m%d_%H%M%S)"
    mv "$TARGET_ADMIN_DIR" "$BACKUP_DIR"
    echo -e "${GREEN}旧文件已备份到: $BACKUP_DIR${RESET}"
elif [ -d "$TARGET_ADMIN_DIR" ]; then
    echo -e "${YELLOW}删除旧的静态文件 (备份功能已禁用)...${RESET}"
    rm -rf "$TARGET_ADMIN_DIR"
fi

# 复制新的dist目录到static/admin目录
echo -e "${BLUE}复制构建文件到静态目录...${RESET}"
cp -r "$DIST_DIR" "$TARGET_ADMIN_DIR"

# 显示构建结果
echo -e "${GREEN}管理端前端打包完成!${RESET}"
echo -e "${BLUE}构建信息:${RESET}"
echo -e "  源目录: $DIST_DIR"
echo -e "  目标目录: $TARGET_ADMIN_DIR"
echo -e "  文件数量: $(find "$TARGET_ADMIN_DIR" -type f | wc -l)"
echo -e "  总大小: $(du -sh "$TARGET_ADMIN_DIR" | cut -f1)"

# 显示主要文件
echo -e "${BLUE}主要文件:${RESET}"
ls -la "$TARGET_ADMIN_DIR"

# 检查关键文件
if [ -f "$TARGET_ADMIN_DIR/index.html" ]; then
    echo -e "${GREEN}✓ index.html 存在${RESET}"
else
    echo -e "${RED}✗ index.html 不存在${RESET}"
fi

# 检查assets目录
if [ -d "$TARGET_ADMIN_DIR/assets" ]; then
    echo -e "${GREEN}✓ assets目录存在${RESET}"
    echo -e "  CSS文件: $(find "$TARGET_ADMIN_DIR/assets" -name "*.css" | wc -l)"
    echo -e "  JS文件: $(find "$TARGET_ADMIN_DIR/assets" -name "*.js" | wc -l)"
else
    echo -e "${YELLOW}⚠ assets目录不存在${RESET}"
fi

echo -e "${GREEN}管理端打包流程完成!${RESET}"
echo -e "${BLUE}提示: 管理端将通过 /admin 路径访问${RESET}"
echo -e "${BLUE}现在可以运行 'make build' 来构建包含静态文件的Go二进制文件${RESET}" 