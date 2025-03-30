#!/bin/bash

# 颜色定义
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
BLUE='\033[0;34m'
NC='\033[0m' # 无颜色

# 版本信息
VERSION="1.0.0"
BUILD_TIME=$(date "+%Y-%m-%d %H:%M:%S")
GIT_HASH=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# 默认值
OUTPUT_DIR="./build"
BINARY_NAME="app"
MAIN_PATH="./cmd/main.go"

# 打印帮助信息
show_help() {
    echo -e "${BLUE}Go项目打包脚本${NC} - v0.1"
    echo
    echo -e "用法: $0 [选项]"
    echo
    echo "选项:"
    echo "  -h, --help          显示帮助信息"
    echo "  -o, --output        指定输出目录 (默认: ./build)"
    echo "  -n, --name          指定二进制文件名 (默认: app)"
    echo "  --local             为当前平台打包"
    echo "  --all               为所有主要平台打包"
    echo "  --windows           为Windows平台打包 (amd64)"
    echo "  --linux             为Linux平台打包 (amd64)"
    echo "  --darwin            为macOS平台打包 (amd64 和 arm64)"
    echo "  --arm               为ARM架构打包 (linux/arm64)"
    echo
    echo "示例:"
    echo "  $0 --local          # 为当前平台编译"
    echo "  $0 --linux          # 为Linux平台编译"
    echo "  $0 --all            # 为所有主要平台编译"
    echo "  $0 -o dist --windows # 将Windows二进制文件输出到dist目录"
}

# 打印构建信息
print_build_info() {
    echo -e "${BLUE}构建信息:${NC}"
    echo -e "  版本: ${VERSION}"
    echo -e "  Git提交: ${GIT_HASH}"
    echo -e "  构建时间: ${BUILD_TIME}"
    echo -e "  输出目录: ${OUTPUT_DIR}"
    echo -e "  二进制名称: ${BINARY_NAME}"
    echo
}

# 执行构建
do_build() {
    local os=$1
    local arch=$2
    local ext=""
    
    # Windows平台添加.exe后缀
    if [ "$os" = "windows" ]; then
        ext=".exe"
    fi
    
    # 创建输出目录
    mkdir -p "${OUTPUT_DIR}/${os}_${arch}"
    
    # 设置构建信息
    local ldflags="-X 'main.version=${VERSION}' -X 'main.buildTime=${BUILD_TIME}' -X 'main.gitHash=${GIT_HASH}'"
    
    echo -e "${YELLOW}正在构建 ${os}/${arch}...${NC}"
    
    # 执行构建
    GOOS=$os GOARCH=$arch go build -ldflags "${ldflags}" -o "${OUTPUT_DIR}/${os}_${arch}/${BINARY_NAME}${ext}" ${MAIN_PATH}
    
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}构建成功:${NC} ${OUTPUT_DIR}/${os}_${arch}/${BINARY_NAME}${ext}"
    else
        echo -e "${RED}构建失败: ${os}/${arch}${NC}"
        return 1
    fi
    
    # 复制配置文件
    cp config.yaml "${OUTPUT_DIR}/${os}_${arch}/" 2>/dev/null || echo -e "${YELLOW}警告: 配置文件未复制${NC}"
    
    # 创建运行脚本
    if [ "$os" = "windows" ]; then
        echo -e "@echo off\r\n${BINARY_NAME}.exe" > "${OUTPUT_DIR}/${os}_${arch}/run.bat"
    else
        echo -e "#!/bin/bash\n./${BINARY_NAME}" > "${OUTPUT_DIR}/${os}_${arch}/run.sh"
        chmod +x "${OUTPUT_DIR}/${os}_${arch}/run.sh"
    fi
    
    return 0
}

# 打包为压缩文件
create_package() {
    local os=$1
    local arch=$2
    local source_dir="${OUTPUT_DIR}/${os}_${arch}"
    local target_file="${OUTPUT_DIR}/${BINARY_NAME}_${VERSION}_${os}_${arch}"
    
    echo -e "${YELLOW}正在打包 ${os}/${arch}...${NC}"
    
    if [ "$os" = "windows" ]; then
        # Windows使用ZIP
        if command -v zip >/dev/null 2>&1; then
            zip -j "${target_file}.zip" "${source_dir}"/* >/dev/null
            echo -e "${GREEN}打包成功:${NC} ${target_file}.zip"
        else
            echo -e "${YELLOW}警告: 未找到zip命令，跳过打包${NC}"
        fi
    else
        # Linux/macOS使用tar.gz
        tar -czf "${target_file}.tar.gz" -C "${OUTPUT_DIR}" "${os}_${arch}" >/dev/null 2>&1
        echo -e "${GREEN}打包成功:${NC} ${target_file}.tar.gz"
    fi
}

# 清理之前的构建
clean_build() {
    if [ -d "$OUTPUT_DIR" ]; then
        echo -e "${YELLOW}清理之前的构建...${NC}"
        rm -rf "$OUTPUT_DIR"
    fi
    mkdir -p "$OUTPUT_DIR"
}

# 打包本地平台
build_local() {
    local os=$(go env GOOS)
    local arch=$(go env GOARCH)
    
    echo -e "${BLUE}打包当前平台 (${os}/${arch})${NC}"
    do_build "$os" "$arch"
    create_package "$os" "$arch"
}

# 打包Windows平台
build_windows() {
    echo -e "${BLUE}打包Windows平台${NC}"
    do_build "windows" "amd64" && create_package "windows" "amd64"
}

# 打包Linux平台
build_linux() {
    echo -e "${BLUE}打包Linux平台${NC}"
    do_build "linux" "amd64" && create_package "linux" "amd64"
}

# 打包macOS平台
build_darwin() {
    echo -e "${BLUE}打包macOS平台${NC}"
    do_build "darwin" "amd64" && create_package "darwin" "amd64"
    do_build "darwin" "arm64" && create_package "darwin" "arm64"
}

# 打包ARM平台
build_arm() {
    echo -e "${BLUE}打包ARM平台${NC}"
    do_build "linux" "arm64" && create_package "linux" "arm64"
}

# 打包所有平台
build_all() {
    echo -e "${BLUE}打包所有主要平台${NC}"
    build_windows
    build_linux
    build_darwin
    build_arm
}

# 处理参数
if [ $# -eq 0 ]; then
    show_help
    exit 0
fi

# 解析命令行参数
while [ $# -gt 0 ]; do
    case "$1" in
        -h|--help)
            show_help
            exit 0
            ;;
        -o|--output)
            OUTPUT_DIR="$2"
            shift 2
            ;;
        -n|--name)
            BINARY_NAME="$2"
            shift 2
            ;;
        --local)
            clean_build
            print_build_info
            build_local
            shift
            ;;
        --windows)
            clean_build
            print_build_info
            build_windows
            shift
            ;;
        --linux)
            clean_build
            print_build_info
            build_linux
            shift
            ;;
        --darwin)
            clean_build
            print_build_info
            build_darwin
            shift
            ;;
        --arm)
            clean_build
            print_build_info
            build_arm
            shift
            ;;
        --all)
            clean_build
            print_build_info
            build_all
            shift
            ;;
        *)
            echo -e "${RED}错误: 未知选项 $1${NC}"
            show_help
            exit 1
            ;;
    esac
done

echo -e "${GREEN}所有操作完成!${NC}"
exit 0 