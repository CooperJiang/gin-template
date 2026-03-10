#!/bin/bash

# 用户端前端打包脚本
# 将 web 主应用与已注册子应用打包后同步到 internal/static/web

set -euo pipefail

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
WEB_DIR="${PROJECT_ROOT}/web"
STATIC_DIR="${PROJECT_ROOT}/internal/static"
MAIN_DIST_DIR="${WEB_DIR}/packages/main/dist"
TARGET_WEB_DIR="${STATIC_DIR}/web"
REGISTRY_FILE="${WEB_DIR}/packages/main/src/micro-app-registry.ts"

declare -a SUB_APP_NAMES=()

echo -e "${BLUE}开始用户端前端项目打包...${RESET}"
echo -e "${BLUE}备份功能: $([ "$ENABLE_BACKUP" = "true" ] && echo "启用" || echo "禁用")${RESET}"

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

# 检查注册表文件
if [ ! -f "$REGISTRY_FILE" ]; then
    echo -e "${RED}错误: 微前端注册表不存在: $REGISTRY_FILE${RESET}"
    exit 1
fi

# 检查pnpm是否可用
if ! command -v pnpm &> /dev/null; then
    echo -e "${RED}错误: pnpm未安装，请先安装: npm install -g pnpm${RESET}"
    exit 1
fi

# 从主应用注册表提取子应用列表
read_sub_apps_from_registry() {
    local registry_apps=()

    mapfile -t registry_apps < <(
        grep -E "name:\s*'[^']+'" "$REGISTRY_FILE" \
            | sed -E "s/.*name:\s*'([^']+)'.*/\1/" \
            | awk '!seen[$0]++'
    )

    if [ ${#registry_apps[@]} -eq 0 ]; then
        echo -e "${YELLOW}警告: 注册表中未解析到子应用，将仅构建主应用${RESET}"
        return
    fi

    for app in "${registry_apps[@]}"; do
        if [ -d "$WEB_DIR/packages/$app" ]; then
            SUB_APP_NAMES+=("$app")
        else
            echo -e "${YELLOW}警告: 跳过非本地子应用 '$app'（未找到目录 web/packages/$app）${RESET}"
        fi
    done
}

# 进入web目录
cd "$WEB_DIR"

# 检查是否安装了node_modules
if [ ! -d "node_modules" ]; then
    echo -e "${YELLOW}未找到node_modules，正在安装依赖...${RESET}"
    pnpm install
    echo -e "${GREEN}依赖安装完成${RESET}"
fi

read_sub_apps_from_registry
if [ ${#SUB_APP_NAMES[@]} -gt 0 ]; then
    echo -e "${BLUE}已纳入构建的子应用: ${SUB_APP_NAMES[*]}${RESET}"
else
    echo -e "${YELLOW}当前没有可构建的本地子应用，仅构建主应用${RESET}"
fi

# 清理之前的构建

echo -e "${YELLOW}清理之前的构建文件...${RESET}"
rm -rf "$MAIN_DIST_DIR"
for app in "${SUB_APP_NAMES[@]}"; do
    rm -rf "$WEB_DIR/packages/$app/dist"
done

# 构建用户端项目

echo -e "${BLUE}构建用户端项目...${RESET}"
pnpm build

# 检查主应用构建结果
if [ ! -d "$MAIN_DIST_DIR" ]; then
    echo -e "${RED}错误: 主应用构建失败，dist目录不存在${RESET}"
    exit 1
fi

# 检查子应用构建结果
for app in "${SUB_APP_NAMES[@]}"; do
    app_dist_dir="$WEB_DIR/packages/$app/dist"
    if [ ! -d "$app_dist_dir" ]; then
        echo -e "${RED}错误: 子应用 '$app' 构建失败，dist目录不存在: $app_dist_dir${RESET}"
        exit 1
    fi
done

# 创建静态目录（如果不存在）
mkdir -p "$STATIC_DIR"

# 备份旧的web目录（如果存在且启用备份）
if [ -d "$TARGET_WEB_DIR" ] && [ "$ENABLE_BACKUP" = "true" ]; then
    echo -e "${YELLOW}备份旧的静态文件...${RESET}"
    BACKUP_DIR="${STATIC_DIR}/web_backup_$(date +%Y%m%d_%H%M%S)"
    mv "$TARGET_WEB_DIR" "$BACKUP_DIR"
    echo -e "${GREEN}旧文件已备份到: $BACKUP_DIR${RESET}"
elif [ -d "$TARGET_WEB_DIR" ]; then
    echo -e "${YELLOW}删除旧的静态文件 (备份功能已禁用)...${RESET}"
    rm -rf "$TARGET_WEB_DIR"
fi

# 复制主应用dist到 static/web

echo -e "${BLUE}复制主应用构建文件到静态目录...${RESET}"
cp -r "$MAIN_DIST_DIR" "$TARGET_WEB_DIR"

# 复制子应用dist到 static/web/subapps/<name>
if [ ${#SUB_APP_NAMES[@]} -gt 0 ]; then
    echo -e "${BLUE}复制子应用构建文件到静态目录...${RESET}"
    mkdir -p "$TARGET_WEB_DIR/subapps"

    for app in "${SUB_APP_NAMES[@]}"; do
        app_dist_dir="$WEB_DIR/packages/$app/dist"
        target_sub_app_dir="$TARGET_WEB_DIR/subapps/$app"
        mkdir -p "$target_sub_app_dir"
        cp -r "$app_dist_dir"/* "$target_sub_app_dir/"
    done
fi

# 显示构建结果

echo -e "${GREEN}用户端前端打包完成!${RESET}"
echo -e "${BLUE}构建信息:${RESET}"
echo -e "  主应用源目录: $MAIN_DIST_DIR"
echo -e "  目标目录: $TARGET_WEB_DIR"
echo -e "  文件数量: $(find "$TARGET_WEB_DIR" -type f | wc -l | tr -d ' ')"
echo -e "  总大小: $(du -sh "$TARGET_WEB_DIR" | cut -f1)"

# 显示主要文件
echo -e "${BLUE}主要文件:${RESET}"
ls -la "$TARGET_WEB_DIR"

# 检查关键文件
if [ -f "$TARGET_WEB_DIR/index.html" ]; then
    echo -e "${GREEN}✓ 主应用 index.html 存在${RESET}"
else
    echo -e "${RED}✗ 主应用 index.html 不存在${RESET}"
    exit 1
fi

# 检查assets目录
if [ -d "$TARGET_WEB_DIR/assets" ]; then
    echo -e "${GREEN}✓ 主应用 assets目录存在${RESET}"
    echo -e "  CSS文件: $(find "$TARGET_WEB_DIR/assets" -name "*.css" | wc -l | tr -d ' ')"
    echo -e "  JS文件: $(find "$TARGET_WEB_DIR/assets" -name "*.js" | wc -l | tr -d ' ')"
else
    echo -e "${YELLOW}⚠ 主应用 assets目录不存在${RESET}"
fi

# 检查子应用关键文件
if [ ${#SUB_APP_NAMES[@]} -gt 0 ]; then
    for app in "${SUB_APP_NAMES[@]}"; do
        if [ -f "$TARGET_WEB_DIR/subapps/$app/index.html" ]; then
            echo -e "${GREEN}✓ 子应用 '$app' index.html 存在${RESET}"
        else
            echo -e "${RED}✗ 子应用 '$app' index.html 不存在${RESET}"
            exit 1
        fi
    done
fi

echo -e "${GREEN}用户端打包流程完成!${RESET}"
echo -e "${BLUE}提示: 主应用通过 / 访问，子应用通过 /subapps/<app-name>/ 加载${RESET}"
echo -e "${BLUE}现在可以运行 'make build' 来构建包含静态文件的Go二进制文件${RESET}"
