#!/bin/bash

# 颜色定义
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
BLUE='\033[0;34m'
NC='\033[0m' # 无颜色

# 默认配置
APP_NAME="app"
APP_BIN="./app"
CONFIG_FILE="config.yaml"
PID_FILE="./.app.pid"
LOG_FILE="./logs/app.log"

# 创建日志目录
mkdir -p logs

# 显示帮助信息
show_help() {
    echo -e "${BLUE}应用程序管理脚本${NC}"
    echo
    echo -e "用法: $0 [选项]"
    echo
    echo "选项:"
    echo "  start       启动应用程序"
    echo "  stop        停止应用程序"
    echo "  restart     重启应用程序"
    echo "  status      查看应用程序状态"
    echo "  log         查看应用程序日志"
    echo "  build       构建应用程序(本地平台)"
    echo "  help        显示此帮助信息"
    echo
}

# 检查应用是否在运行
is_running() {
    if [ -f "$PID_FILE" ]; then
        PID=$(cat "$PID_FILE")
        if ps -p "$PID" > /dev/null; then
            return 0  # 正在运行
        fi
    fi
    return 1  # 未运行
}

# 启动应用
start_app() {
    echo -e "${YELLOW}正在启动 ${APP_NAME}...${NC}"
    
    if is_running; then
        echo -e "${RED}应用已经在运行，PID: $(cat $PID_FILE)${NC}"
        return 1
    fi
    
    # 检查二进制文件是否存在
    if [ ! -f "$APP_BIN" ]; then
        echo -e "${RED}错误: 找不到应用程序 $APP_BIN${NC}"
        echo -e "${YELLOW}尝试先构建应用程序: $0 build${NC}"
        return 1
    fi
    
    # 检查配置文件是否存在
    if [ ! -f "$CONFIG_FILE" ]; then
        echo -e "${RED}错误: 找不到配置文件 $CONFIG_FILE${NC}"
        return 1
    fi
    
    nohup "$APP_BIN" > "$LOG_FILE" 2>&1 &
    PID=$!
    echo $PID > "$PID_FILE"
    
    sleep 1
    if is_running; then
        echo -e "${GREEN}应用已启动，PID: $PID${NC}"
        echo -e "使用 '$0 log' 查看日志"
    else
        echo -e "${RED}应用启动失败!${NC}"
        echo -e "检查日志文件: $LOG_FILE"
        return 1
    fi
}

# 停止应用
stop_app() {
    echo -e "${YELLOW}正在停止 ${APP_NAME}...${NC}"
    
    if is_running; then
        PID=$(cat "$PID_FILE")
        kill "$PID"
        sleep 2
        
        if is_running; then
            echo -e "${RED}无法正常停止应用，尝试强制终止...${NC}"
            kill -9 "$PID"
            sleep 1
        fi
        
        if ! is_running; then
            rm -f "$PID_FILE"
            echo -e "${GREEN}应用已停止${NC}"
        else
            echo -e "${RED}无法停止应用!${NC}"
            return 1
        fi
    else
        echo -e "${YELLOW}应用未在运行${NC}"
    fi
}

# 重启应用
restart_app() {
    stop_app
    sleep 1
    start_app
}

# 查看应用状态
show_status() {
    if is_running; then
        PID=$(cat "$PID_FILE")
        echo -e "${GREEN}应用正在运行${NC}"
        echo -e "PID: $PID"
        
        # 显示运行时间
        if [[ "$OSTYPE" == "darwin"* ]]; then
            # macOS
            STARTED=$(ps -p "$PID" -o lstart= 2>/dev/null)
            if [ -n "$STARTED" ]; then
                echo -e "启动时间: $STARTED"
            fi
        else
            # Linux
            if [ -d "/proc/$PID" ]; then
                STARTED=$(stat -c %x "/proc/$PID" 2>/dev/null | cut -d' ' -f1,2)
                if [ -n "$STARTED" ]; then
                    echo -e "启动时间: $STARTED"
                fi
            fi
        fi
        
        # 检查端口是否在监听
        if command -v netstat &>/dev/null; then
            PORT=$(grep -o "port:.*" "$CONFIG_FILE" | awk '{print $2}' | tr -d ' ')
            if [ -n "$PORT" ]; then
                echo -e "\n${YELLOW}正在检查端口状态...${NC}"
                if netstat -an | grep -q ":$PORT.*LISTEN"; then
                    echo -e "${GREEN}端口 $PORT 正在监听${NC}"
                else
                    echo -e "${RED}端口 $PORT 未在监听${NC}"
                fi
            fi
        fi
    else
        echo -e "${YELLOW}应用未在运行${NC}"
    fi
}

# 查看日志
view_log() {
    if [ -f "$LOG_FILE" ]; then
        echo -e "${BLUE}最新日志内容:${NC}\n"
        tail -n 50 "$LOG_FILE"
        echo -e "\n${YELLOW}使用 'tail -f $LOG_FILE' 实时查看日志${NC}"
    else
        echo -e "${RED}日志文件不存在: $LOG_FILE${NC}"
    fi
}

# 构建应用
build_app() {
    echo -e "${YELLOW}正在构建应用...${NC}"
    if [ -f "./build.sh" ]; then
        ./build.sh --local
        if [ -d "./build" ]; then
            local os=$(go env GOOS)
            local arch=$(go env GOARCH)
            cp "./build/${os}_${arch}/${APP_NAME}" "./${APP_NAME}"
            echo -e "${GREEN}应用构建完成，已复制到当前目录${NC}"
        fi
    else
        go build -o "$APP_NAME" ./cmd/main.go
        if [ $? -eq 0 ]; then
            echo -e "${GREEN}应用构建成功: $APP_NAME${NC}"
        else
            echo -e "${RED}应用构建失败!${NC}"
            return 1
        fi
    fi
}

# 命令处理
case "$1" in
    start)
        start_app
        ;;
    stop)
        stop_app
        ;;
    restart)
        restart_app
        ;;
    status)
        show_status
        ;;
    log)
        view_log
        ;;
    build)
        build_app
        ;;
    help|--help|-h)
        show_help
        ;;
    *)
        show_help
        exit 1
        ;;
esac

exit 0 