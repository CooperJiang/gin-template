#!/bin/bash

# 颜色定义
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
BLUE='\033[0;34m'
NC='\033[0m' # 无颜色

# 默认配置
APP_NAME="app"
VERSION="1.0.0"
REMOTE_USER="root"
REMOTE_HOST=""
REMOTE_PORT="22"
REMOTE_DIR="/opt/myapp"
LOCAL_PACKAGE="./build/linux_amd64/${APP_NAME}"
SSH_KEY=""
KEEP_RELEASES=3
ENV_MODE="prod"  # 新增：环境模式，默认为生产环境
CONFIG_FILE_PATH="./config.prod.yaml"  # 新增：指定配置文件路径

# 配置文件
CONFIG_FILE="deploy.conf"

# 读取配置文件（如果存在）
if [ -f "$CONFIG_FILE" ]; then
    source "$CONFIG_FILE"
fi

# 显示帮助信息
show_help() {
    echo -e "${BLUE}自动部署脚本${NC}"
    echo
    echo -e "用法: $0 [选项] [命令]"
    echo
    echo "命令:"
    echo "  deploy      部署应用到远程服务器"
    echo "  rollback    回滚到上一个版本"
    echo "  restart     重启远程服务器上的应用"
    echo "  status      检查远程服务器上的应用状态"
    echo "  logs        查看远程服务器上的应用日志"
    echo "  setup       设置远程服务器环境"
    echo "  config      配置部署参数"
    echo "  envconfig   配置环境特定参数（生产/测试）"  # 新增：环境配置命令
    echo "  help        显示帮助信息"
    echo
    echo "选项:"
    echo "  -h, --host      指定远程主机地址"
    echo "  -u, --user      指定远程用户名"
    echo "  -p, --port      指定SSH端口"
    echo "  -d, --dir       指定远程部署目录"
    echo "  -n, --name      指定应用名称"
    echo "  -k, --key       指定SSH密钥路径"
    echo "  -v, --version   指定版本号"
    echo "  -e, --env       指定部署环境 (prod/test/dev)"  # 新增：环境选项
    echo "  -c, --config    指定配置文件路径"  # 新增：指定配置文件
    echo
    echo "示例:"
    echo "  $0 config                       # 配置部署参数"
    echo "  $0 envconfig                    # 配置环境参数"
    echo "  $0 -h 192.168.1.100 deploy      # 部署到指定服务器"
    echo "  $0 -e test deploy               # 部署到测试环境"
    echo "  $0 -c config.prod.yaml deploy   # 使用生产配置部署"
    echo "  $0 rollback                     # 回滚到上一个版本"
    echo
    echo "提示: 首次使用建议先运行 '$0 config' 配置部署参数"
}

# 保存配置到文件
save_config() {
    cat > "$CONFIG_FILE" << EOF
# 部署配置文件
APP_NAME="$APP_NAME"
VERSION="$VERSION"
REMOTE_USER="$REMOTE_USER"
REMOTE_HOST="$REMOTE_HOST"
REMOTE_PORT="$REMOTE_PORT"
REMOTE_DIR="$REMOTE_DIR"
SSH_KEY="$SSH_KEY"
KEEP_RELEASES=$KEEP_RELEASES
ENV_MODE="$ENV_MODE"
CONFIG_FILE_PATH="$CONFIG_FILE_PATH"
EOF
    echo -e "${GREEN}配置已保存到 $CONFIG_FILE${NC}"
}

# SSH/SCP命令构建
build_ssh_cmd() {
    local CMD="ssh"
    
    if [ -n "$SSH_KEY" ]; then
        CMD="$CMD -i $SSH_KEY"
    fi
    
    CMD="$CMD -p $REMOTE_PORT ${REMOTE_USER}@${REMOTE_HOST}"
    echo "$CMD"
}

build_scp_cmd() {
    local CMD="scp"
    
    if [ -n "$SSH_KEY" ]; then
        CMD="$CMD -i $SSH_KEY"
    fi
    
    CMD="$CMD -P $REMOTE_PORT"
    echo "$CMD"
}

# 检查必要参数
check_required_params() {
    local MISSING=""
    
    if [ -z "$REMOTE_HOST" ]; then
        MISSING="$MISSING\n  - 远程主机地址 (-h, --host)"
    fi
    
    if [ -z "$REMOTE_USER" ]; then
        MISSING="$MISSING\n  - 远程用户名 (-u, --user)"
    fi
    
    if [ -n "$MISSING" ]; then
        echo -e "${RED}错误: 缺少必要参数:${NC}$MISSING"
        echo -e "${YELLOW}提示: 使用 '$0 config' 配置部署参数${NC}"
        return 1
    fi
    
    return 0
}

# 检查本地打包
check_local_package() {
    if [ ! -f "./build.sh" ]; then
        echo -e "${RED}错误: 未找到构建脚本 './build.sh'${NC}"
        return 1
    fi
    
    echo -e "${BLUE}准备应用包...${NC}"
    
    # 检查build目录是否存在
    local SHOULD_BUILD=0
    if [ ! -d "./build/linux_amd64" ]; then
        echo -e "${YELLOW}未找到Linux平台的构建包，将进行构建...${NC}"
        SHOULD_BUILD=1
    else
        # 二进制文件存在，询问是否重新构建
        LOCAL_PACKAGE="./build/linux_amd64/${APP_NAME}"
        if [ -f "$LOCAL_PACKAGE" ]; then
            # 获取文件修改时间
            local FILE_TIME=$(date -r "$LOCAL_PACKAGE" "+%Y-%m-%d %H:%M:%S")
            echo -e "${YELLOW}发现已存在的构建文件:${NC}"
            echo -e "  路径: ${LOCAL_PACKAGE}"
            echo -e "  修改时间: ${FILE_TIME}"
            read -p "是否重新构建应用? (y/n) [n]: " rebuild_choice
            if [ "$rebuild_choice" = "y" ] || [ "$rebuild_choice" = "Y" ]; then
                SHOULD_BUILD=1
                echo -e "${YELLOW}将重新构建应用...${NC}"
            else
                echo -e "${GREEN}使用现有构建文件${NC}"
            fi
        else
            echo -e "${YELLOW}构建目录存在但未找到二进制文件，将进行构建...${NC}"
            SHOULD_BUILD=1
        fi
    fi
    
    # 如果需要构建，则执行构建脚本
    if [ $SHOULD_BUILD -eq 1 ]; then
        ./build.sh --linux
        
        if [ ! -d "./build/linux_amd64" ]; then
            echo -e "${RED}错误: 构建Linux平台包失败${NC}"
            return 1
        fi
    fi
    
    # 检查二进制文件
    LOCAL_PACKAGE="./build/linux_amd64/${APP_NAME}"
    if [ ! -f "$LOCAL_PACKAGE" ]; then
        echo -e "${RED}错误: 未找到Linux平台的构建包: $LOCAL_PACKAGE${NC}"
        return 1
    fi
    
    echo -e "${GREEN}应用二进制文件已准备好: $LOCAL_PACKAGE${NC}"
    
    # 创建部署包目录
    DEPLOY_TEMP_DIR="./build/deploy_package"
    rm -rf "$DEPLOY_TEMP_DIR"
    mkdir -p "$DEPLOY_TEMP_DIR"
    
    # 复制二进制文件
    cp "$LOCAL_PACKAGE" "$DEPLOY_TEMP_DIR/"
    
    # 检查配置文件路径
    local CONFIG_SRC="config.yaml"  # 默认配置文件
    if [ -n "$CONFIG_FILE_PATH" ] && [ -f "$CONFIG_FILE_PATH" ]; then
        CONFIG_SRC="$CONFIG_FILE_PATH"
        echo -e "${BLUE}使用环境特定配置: $CONFIG_SRC${NC}"
    else
        echo -e "${YELLOW}使用默认配置文件: $CONFIG_SRC${NC}"
    fi
    
    # 复制配置文件到部署包
    if [ -f "$CONFIG_SRC" ]; then
        cp "$CONFIG_SRC" "$DEPLOY_TEMP_DIR/config.yaml"
        echo -e "${GREEN}配置文件已添加到部署包: $CONFIG_SRC${NC}"
    else
        echo -e "${RED}警告: 配置文件不存在: $CONFIG_SRC${NC}"
    fi
    
    # 创建启动脚本
    cat > "$DEPLOY_TEMP_DIR/run.sh" << EOF
#!/bin/bash

# 应用启动脚本
APP_NAME="${APP_NAME}"
CONFIG_FILE="./config.yaml"
LOG_FILE="./app.log"
PID_FILE="./${APP_NAME}.pid"

# 启动应用
start() {
    if [ -f "\$PID_FILE" ]; then
        PID=\$(cat "\$PID_FILE")
        if ps -p "\$PID" > /dev/null; then
            echo "应用已在运行，PID: \$PID"
            return 1
        else
            echo "发现过期的PID文件，将删除"
            rm -f "\$PID_FILE"
        fi
    fi
    
    echo "正在启动应用..."
    nohup ./\$APP_NAME > "\$LOG_FILE" 2>&1 &
    PID=\$!
    echo \$PID > "\$PID_FILE"
    echo "应用已启动，PID: \$PID"
}

# 停止应用
stop() {
    if [ ! -f "\$PID_FILE" ]; then
        echo "应用未在运行"
        return 0
    fi
    
    PID=\$(cat "\$PID_FILE")
    if ! ps -p "\$PID" > /dev/null; then
        echo "应用未在运行 (PID文件存在但进程不存在)"
        rm -f "\$PID_FILE"
        return 0
    fi
    
    echo "正在停止应用，PID: \$PID"
    kill "\$PID"
    sleep 2
    
    if ps -p "\$PID" > /dev/null; then
        echo "强制终止应用，PID: \$PID"
        kill -9 "\$PID"
        sleep 1
    fi
    
    rm -f "\$PID_FILE"
    echo "应用已停止"
}

# 查看状态
status() {
    if [ ! -f "\$PID_FILE" ]; then
        echo "应用未在运行"
        return 1
    fi
    
    PID=\$(cat "\$PID_FILE")
    if ! ps -p "\$PID" > /dev/null; then
        echo "应用未在运行 (PID文件存在但进程不存在)"
        rm -f "\$PID_FILE"
        return 1
    fi
    
    echo "应用正在运行，PID: \$PID"
    echo "日志文件: \$LOG_FILE"
    return 0
}

# 查看日志
logs() {
    LINES=\${1:-50}
    if [ -f "\$LOG_FILE" ]; then
        tail -n \$LINES "\$LOG_FILE"
    else
        echo "日志文件不存在: \$LOG_FILE"
        return 1
    fi
}

# 处理命令
case "\$1" in
    start)
        start
        ;;
    stop)
        stop
        ;;
    restart)
        stop
        start
        ;;
    status)
        status
        ;;
    logs)
        logs \$2
        ;;
    *)
        echo "用法: \$0 {start|stop|restart|status|logs}"
        exit 1
        ;;
esac

exit 0
EOF
    
    # 设置执行权限
    chmod +x "$DEPLOY_TEMP_DIR/run.sh"
    echo -e "${GREEN}启动脚本已添加到部署包${NC}"
    
    # 打包文件
    DEPLOY_PACKAGE="./build/${APP_NAME}_${ENV_MODE}_package.tar.gz"
    tar -czf "$DEPLOY_PACKAGE" -C "$DEPLOY_TEMP_DIR" .
    
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}部署包已创建: $DEPLOY_PACKAGE${NC}"
        
        # 清理无用的构建文件，保留必要的部署包
        echo -e "${BLUE}清理无用的构建文件...${NC}"
        
        # 保留部署包和必要的二进制文件
        find ./build -type f ! -name "${APP_NAME}_${ENV_MODE}_package.tar.gz" ! -path "*/linux_amd64/${APP_NAME}" -delete 2>/dev/null
        
        # 清理空目录
        find ./build -type d -empty -delete 2>/dev/null
        
        # 确保linux_amd64目录被保留
        mkdir -p "./build/linux_amd64"
        
        # 清理临时目录
        rm -rf "$DEPLOY_TEMP_DIR"
        
        echo -e "${GREEN}无用文件清理完成${NC}"
        return 0
    else
        echo -e "${RED}创建部署包失败${NC}"
        return 1
    fi
}

# 配置参数
configure() {
    echo -e "${BLUE}配置部署参数${NC}"
    echo
    
    # 读取应用名称
    read -p "应用名称 [$APP_NAME]: " input
    if [ -n "$input" ]; then APP_NAME="$input"; fi
    
    # 读取远程主机
    read -p "远程主机地址 [$REMOTE_HOST]: " input
    if [ -n "$input" ]; then REMOTE_HOST="$input"; fi
    
    # 读取远程用户
    read -p "远程用户名 [$REMOTE_USER]: " input
    if [ -n "$input" ]; then REMOTE_USER="$input"; fi
    
    # 读取SSH端口
    read -p "SSH端口 [$REMOTE_PORT]: " input
    if [ -n "$input" ]; then REMOTE_PORT="$input"; fi
    
    # 读取远程目录
    read -p "远程部署目录 [$REMOTE_DIR]: " input
    if [ -n "$input" ]; then REMOTE_DIR="$input"; fi
    
    # 读取SSH密钥
    read -p "SSH密钥路径 (留空使用密码) [$SSH_KEY]: " input
    if [ -n "$input" ]; then SSH_KEY="$input"; fi
    
    # 读取保留版本数
    read -p "保留最近版本数量 [$KEEP_RELEASES]: " input
    if [ -n "$input" ]; then KEEP_RELEASES="$input"; fi
    
    # 保存配置
    save_config
    
    echo -e "${GREEN}配置完成!${NC}"
    echo -e "使用 '$0 deploy' 部署应用，或使用 '$0 envconfig' 配置环境特定参数"
}

# 配置环境特定参数
configure_env() {
    echo -e "${BLUE}配置环境特定参数${NC}"
    echo
    
    # 读取环境模式
    echo "部署环境选项:"
    echo "  1) 生产环境 (prod)"
    echo "  2) 测试环境 (test)"
    echo "  3) 开发环境 (dev)"
    read -p "选择部署环境 [1-3] (默认: 1): " env_choice
    
    case $env_choice in
        2)
            ENV_MODE="test"
            ;;
        3)
            ENV_MODE="dev"
            ;;
        *)
            ENV_MODE="prod"
            ;;
    esac
    
    # 检查环境特定配置文件
    local env_config="config.${ENV_MODE}.yaml"
    if [ -f "$env_config" ]; then
        echo -e "${GREEN}发现环境配置文件: $env_config${NC}"
        CONFIG_FILE_PATH="$env_config"
    else
        echo -e "${YELLOW}未找到环境配置文件: $env_config${NC}"
        echo -e "您可以创建环境特定的配置文件，或者选择使用本地配置文件上传"
        read -p "使用的配置文件路径 [config.yaml]: " config_path
        if [ -n "$config_path" ]; then
            if [ -f "$config_path" ]; then
                CONFIG_FILE_PATH="$config_path"
            else
                echo -e "${RED}错误: 配置文件不存在: $config_path${NC}"
                CONFIG_FILE_PATH=""
            fi
        else
            CONFIG_FILE_PATH="config.yaml"
        fi
    fi
    
    # 保存配置
    save_config
    
    echo -e "${GREEN}环境配置完成!${NC}"
    echo -e "环境: ${ENV_MODE}"
    echo -e "配置文件: ${CONFIG_FILE_PATH}"
    echo
    echo -e "使用 '$0 deploy' 部署应用"
}

# 设置远程服务器环境
setup_remote() {
    echo -e "${YELLOW}正在设置远程服务器环境...${NC}"
    
    if ! check_required_params; then
        return 1
    fi
    
    local SSH_CMD=$(build_ssh_cmd)
    
    # 创建目录结构
    echo -e "${BLUE}创建目录结构...${NC}"
    if ! $SSH_CMD "mkdir -p $REMOTE_DIR/{releases,shared,shared/config,tmp}"; then
        echo -e "${RED}创建目录结构失败${NC}"
        return 1
    fi
    
    # 创建服务管理脚本
    echo -e "${BLUE}创建服务管理脚本...${NC}"
    cat > ./tmp_service.sh << EOF
#!/bin/bash

# 应用服务管理脚本
APP_NAME="$APP_NAME"
APP_DIR="$REMOTE_DIR"
PID_FILE="\$APP_DIR/shared/\$APP_NAME.pid"
LOG_FILE="\$APP_DIR/shared/\$APP_NAME.log"
CONFIG_DIR="\$APP_DIR/shared/config"
CURRENT_LINK="\$APP_DIR/current"
ENV_FILE="\$APP_DIR/shared/.env"

# 确保日志目录存在
mkdir -p "\$(dirname \$LOG_FILE)"

# 获取环境变量
if [ -f "\$ENV_FILE" ]; then
    source "\$ENV_FILE"
fi

# 启动应用
start() {
    if [ -f "\$PID_FILE" ] && ps -p \$(cat "\$PID_FILE") > /dev/null; then
        echo "应用已在运行，PID: \$(cat "\$PID_FILE")"
        return 1
    fi
    
    if [ ! -L "\$CURRENT_LINK" ]; then
        echo "错误: 未找到当前版本链接"
        return 1
    fi
    
    cd "\$CURRENT_LINK"
    nohup ./\$APP_NAME > "\$LOG_FILE" 2>&1 &
    echo \$! > "\$PID_FILE"
    echo "应用已启动，PID: \$(cat "\$PID_FILE")"
}

# 停止应用
stop() {
    if [ ! -f "\$PID_FILE" ]; then
        echo "应用未在运行"
        return 0
    fi
    
    local PID=\$(cat "\$PID_FILE")
    if ! ps -p "\$PID" > /dev/null; then
        echo "应用未在运行 (PID文件存在但进程不存在)"
        rm -f "\$PID_FILE"
        return 0
    fi
    
    echo "正在停止应用，PID: \$PID"
    kill "\$PID"
    sleep 2
    
    if ps -p "\$PID" > /dev/null; then
        echo "强制终止应用，PID: \$PID"
        kill -9 "\$PID"
        sleep 1
    fi
    
    rm -f "\$PID_FILE"
    echo "应用已停止"
}

# 查看状态
status() {
    if [ ! -f "\$PID_FILE" ]; then
        echo "应用未在运行"
        return 1
    fi
    
    local PID=\$(cat "\$PID_FILE")
    if ! ps -p "\$PID" > /dev/null; then
        echo "应用未在运行 (PID文件存在但进程不存在)"
        rm -f "\$PID_FILE"
        return 1
    fi
    
    echo "应用正在运行，PID: \$PID"
    echo "日志文件: \$LOG_FILE"
    echo "当前版本: \$(readlink \$CURRENT_LINK 2>/dev/null || echo "未知")"
    echo "当前环境: \$(grep -o 'APP_ENV=.*' \$ENV_FILE 2>/dev/null | cut -d= -f2 || echo "未知")"
    return 0
}

# 查看日志
logs() {
    local LINES=\${1:-50}
    if [ -f "\$LOG_FILE" ]; then
        tail -n \$LINES "\$LOG_FILE"
    else
        echo "日志文件不存在: \$LOG_FILE"
        return 1
    fi
}

# 显示帮助
help() {
    echo "用法: \$0 {start|stop|restart|status|logs}"
    echo
    echo "命令:"
    echo "  start      启动应用"
    echo "  stop       停止应用"
    echo "  restart    重启应用"
    echo "  status     查看应用状态"
    echo "  logs [n]   查看最后n行日志 (默认50行)"
    echo "  help       显示帮助信息"
}

# 处理命令
case "\$1" in
    start)
        start
        ;;
    stop)
        stop
        ;;
    restart)
        stop
        start
        ;;
    status)
        status
        ;;
    logs)
        logs \$2
        ;;
    *)
        help
        exit 1
        ;;
esac

exit 0
EOF

    # 上传服务管理脚本
    local SCP_CMD=$(build_scp_cmd)
    if ! $SCP_CMD ./tmp_service.sh ${REMOTE_USER}@${REMOTE_HOST}:${REMOTE_DIR}/service.sh; then
        echo -e "${RED}服务管理脚本上传失败${NC}"
        rm -f ./tmp_service.sh
        return 1
    fi
    
    rm -f ./tmp_service.sh
    
    # 设置权限
    if ! $SSH_CMD "chmod +x $REMOTE_DIR/service.sh"; then
        echo -e "${RED}设置服务脚本权限失败${NC}"
        return 1
    fi
    
    echo -e "${GREEN}远程服务器环境设置完成!${NC}"
    echo -e "使用 '$0 deploy' 部署应用"
    return 0
}

# 准备环境变量文件
prepare_env_file() {
    echo -e "${BLUE}准备环境变量文件...${NC}"
    
    cat > ./tmp_env << EOF
# 环境变量配置 - 由部署脚本自动生成
APP_ENV=${ENV_MODE}
APP_VERSION=${VERSION}
APP_DEPLOY_TIME="$(date +"%Y-%m-%d %H:%M:%S")"
EOF
    
    local SCP_CMD=$(build_scp_cmd)
    
    # 确保目录存在
    local SSH_CMD=$(build_ssh_cmd)
    $SSH_CMD "mkdir -p ${REMOTE_DIR}/shared"
    
    # 上传文件并检查结果
    if $SCP_CMD ./tmp_env ${REMOTE_USER}@${REMOTE_HOST}:${REMOTE_DIR}/shared/.env; then
        echo -e "${GREEN}环境变量文件已上传${NC}"
    else
        echo -e "${RED}环境变量文件上传失败${NC}"
        rm -f ./tmp_env
        return 1
    fi
    
    rm -f ./tmp_env
    return 0
}

# 检查并处理端口占用
check_port_usage() {
    echo -e "${BLUE}检查端口占用情况...${NC}"
    
    local PORT_TO_CHECK=$(get_app_port)
    if [ -z "$PORT_TO_CHECK" ]; then
        echo -e "${YELLOW}无法从配置文件中获取端口号，将使用默认端口9000${NC}"
        PORT_TO_CHECK=9000
    fi
    
    local SSH_CMD=$(build_ssh_cmd)
    
    # 检查远程服务器上指定端口是否被占用
    echo -e "${BLUE}检查端口 ${PORT_TO_CHECK} 是否被占用...${NC}"
    local PORT_USAGE=$($SSH_CMD "lsof -i:${PORT_TO_CHECK} | grep LISTEN || echo '未占用'")
    
    if [ "$PORT_USAGE" != "未占用" ]; then
        echo -e "${YELLOW}端口 ${PORT_TO_CHECK} 已被占用${NC}"
        
        # 查找占用端口的进程PID
        local PID_LIST=$($SSH_CMD "lsof -i:${PORT_TO_CHECK} -t")
        
        if [ -n "$PID_LIST" ]; then
            # 检查是否是自己的应用进程
            local IS_OWN_APP=0
            local PROCESS_DETAILS=$($SSH_CMD "for pid in $PID_LIST; do ps -p \$pid -o pid,ppid,user,comm,args; done")
            echo -e "占用端口的进程信息:\n$PROCESS_DETAILS"
            
            # 检查各种可能识别为自己应用的特征
            if [[ "$PROCESS_DETAILS" == *"$APP_NAME"* ]] || \
               [[ "$PROCESS_DETAILS" == *"${REMOTE_DIR}"* ]] || \
               $SSH_CMD "for pid in $PID_LIST; do ls -l /proc/\$pid/cwd 2>/dev/null | grep -q '${REMOTE_DIR}'; if [ \$? -eq 0 ]; then exit 0; fi; done"; then
                IS_OWN_APP=1
                echo -e "${GREEN}检测到端口被自己的旧实例占用，将自动处理...${NC}"
            fi
            
            if [ $IS_OWN_APP -eq 1 ]; then
                # 自动处理自己的旧实例
                echo -e "${BLUE}尝试停止旧应用实例...${NC}"
                
                # 首先尝试使用服务脚本停止
                if $SSH_CMD "[ -f \"${REMOTE_DIR}/service.sh\" ]" 2>/dev/null; then
                    echo -e "${YELLOW}使用服务脚本停止旧应用...${NC}"
                    $SSH_CMD "${REMOTE_DIR}/service.sh stop"
                    sleep 2
                    
                    # 检查端口是否已释放
                    PORT_USAGE=$($SSH_CMD "lsof -i:${PORT_TO_CHECK} | grep LISTEN || echo '未占用'")
                    if [ "$PORT_USAGE" == "未占用" ]; then
                        echo -e "${GREEN}端口 ${PORT_TO_CHECK} 已成功释放${NC}"
                        return 0
                    fi
                fi
                
                # 如果服务脚本未能停止，直接终止进程
                echo -e "${YELLOW}服务脚本未能停止旧应用，尝试直接终止进程...${NC}"
                
                for PID in $PID_LIST; do
                    echo -e "${BLUE}终止进程 ${PID}...${NC}"
                    $SSH_CMD "kill -15 ${PID} 2>/dev/null"
                done
                
                sleep 2
                
                # 检查是否还有进程未终止
                local REMAINING_PIDS=$($SSH_CMD "lsof -i:${PORT_TO_CHECK} -t")
                if [ -n "$REMAINING_PIDS" ]; then
                    echo -e "${YELLOW}某些进程未响应SIGTERM，尝试强制终止...${NC}"
                    for PID in $REMAINING_PIDS; do
                        $SSH_CMD "kill -9 ${PID} 2>/dev/null"
                    done
                    sleep 1
                fi
                
                # 最终检查
                PORT_USAGE=$($SSH_CMD "lsof -i:${PORT_TO_CHECK} | grep LISTEN || echo '未占用'")
                if [ "$PORT_USAGE" == "未占用" ]; then
                    echo -e "${GREEN}端口 ${PORT_TO_CHECK} 已成功释放${NC}"
                    return 0
                else
                    echo -e "${RED}警告: 无法释放端口 ${PORT_TO_CHECK}，尝试最后的方法...${NC}"
                    
                    # 尝试使用fuser命令
                    $SSH_CMD "fuser -k ${PORT_TO_CHECK}/tcp 2>/dev/null || true"
                    sleep 1
                    
                    PORT_USAGE=$($SSH_CMD "lsof -i:${PORT_TO_CHECK} | grep LISTEN || echo '未占用'")
                    if [ "$PORT_USAGE" == "未占用" ]; then
                        echo -e "${GREEN}端口 ${PORT_TO_CHECK} 已成功释放${NC}"
                        return 0
                    else
                        echo -e "${RED}错误: 无法释放端口，部署可能会失败${NC}"
                        read -p "是否继续部署? (y/n) [n]: " continue_deploy
                        if [ "$continue_deploy" != "y" ] && [ "$continue_deploy" != "Y" ]; then
                            return 1
                        fi
                    fi
                fi
            else
                # 端口被其他应用占用
                echo -e "${YELLOW}端口被其他应用占用，请手动处理...${NC}"
                read -p "是否尝试释放占用的端口? (y/n) [n]: " release_port
                if [ "$release_port" = "y" ] || [ "$release_port" = "Y" ]; then
                    echo -e "${BLUE}尝试释放端口 ${PORT_TO_CHECK}...${NC}"
                    
                    # 显示进程详情供用户确认
                    echo -e "占用端口的进程列表:"
                    for PID in $PID_LIST; do
                        local PROCESS_INFO=$($SSH_CMD "ps -p ${PID} -o pid,ppid,user,comm,args || echo '未知进程'")
                        echo -e "$PROCESS_INFO"
                        read -p "是否终止此进程? (y/n) [n]: " kill_process
                        if [ "$kill_process" = "y" ] || [ "$kill_process" = "Y" ]; then
                            echo -e "${YELLOW}终止进程 ${PID}...${NC}"
                            $SSH_CMD "kill -15 ${PID} 2>/dev/null"
                            sleep 1
                            if $SSH_CMD "ps -p ${PID} > /dev/null 2>&1"; then
                                echo -e "${YELLOW}进程未响应，尝试强制终止...${NC}"
                                $SSH_CMD "kill -9 ${PID} 2>/dev/null"
                            fi
                        fi
                    done
                    
                    # 检查端口是否已释放
                    PORT_USAGE=$($SSH_CMD "lsof -i:${PORT_TO_CHECK} | grep LISTEN || echo '未占用'")
                    if [ "$PORT_USAGE" == "未占用" ]; then
                        echo -e "${GREEN}端口 ${PORT_TO_CHECK} 已成功释放${NC}"
                        return 0
                    else
                        echo -e "${RED}警告: 端口 ${PORT_TO_CHECK} 仍被占用${NC}"
                        read -p "是否继续部署? (y/n) [n]: " continue_deploy
                        if [ "$continue_deploy" != "y" ] && [ "$continue_deploy" != "Y" ]; then
                            return 1
                        fi
                    fi
                else
                    echo -e "${RED}警告: 端口 ${PORT_TO_CHECK} 仍被占用，可能导致部署后应用无法启动${NC}"
                    read -p "是否继续部署? (y/n) [n]: " continue_deploy
                    if [ "$continue_deploy" != "y" ] && [ "$continue_deploy" != "Y" ]; then
                        return 1
                    fi
                fi
            fi
        else
            echo -e "${RED}错误: 无法获取占用端口的进程ID${NC}"
            read -p "是否继续部署? (y/n) [n]: " continue_deploy
            if [ "$continue_deploy" != "y" ] && [ "$continue_deploy" != "Y" ]; then
                return 1
            fi
        fi
    else
        echo -e "${GREEN}端口 ${PORT_TO_CHECK} 未被占用${NC}"
    fi
    
    return 0
}

# 从配置文件获取应用端口
get_app_port() {
    local CONFIG_FILE="${CONFIG_FILE_PATH:-config.yaml}"
    # 尝试使用grep命令从配置文件中提取端口号
    if [ -f "$CONFIG_FILE" ]; then
        local PORT=$(grep -E "port:.*[0-9]+" "$CONFIG_FILE" | grep -oE "[0-9]+" | head -1)
        echo "$PORT"
    else
        echo ""
    fi
}

# 部署应用
deploy_app() {
    echo -e "${YELLOW}正在部署应用...${NC}"
    
    if ! check_required_params; then
        return 1
    fi
    
    # 准备部署包
    if ! check_local_package; then
        return 1
    fi
    
    local SSH_CMD=$(build_ssh_cmd)
    local SCP_CMD=$(build_scp_cmd)
    
    # 检查远程目录是否存在
    if ! $SSH_CMD "[ -d \"${REMOTE_DIR}\" ]" 2>/dev/null; then
        echo -e "${RED}错误: 远程目录 ${REMOTE_DIR} 不存在${NC}"
        echo -e "${YELLOW}提示: 请先运行 '$0 setup' 初始化服务器环境${NC}"
        return 1
    fi
    
    # 检查服务脚本是否存在
    if ! $SSH_CMD "[ -f \"${REMOTE_DIR}/service.sh\" ]" 2>/dev/null; then
        echo -e "${RED}错误: 服务管理脚本不存在${NC}"
        echo -e "${YELLOW}提示: 请先运行 '$0 setup' 初始化服务器环境${NC}"
        return 1
    fi
    
    # 检查端口占用情况
    if ! check_port_usage; then
        echo -e "${RED}取消部署: 端口占用问题未解决${NC}"
        return 1
    fi
    
    local TIMESTAMP=$(date +%Y%m%d%H%M%S)
    local RELEASE_DIR="${REMOTE_DIR}/releases/${TIMESTAMP}"
    
    # 创建远程目录
    echo -e "${BLUE}创建发布目录...${NC}"
    if ! $SSH_CMD "mkdir -p ${RELEASE_DIR}"; then
        echo -e "${RED}创建发布目录失败${NC}"
        return 1
    fi
    
    # 准备环境变量文件
    if ! prepare_env_file; then
        echo -e "${RED}环境变量文件准备失败，但将继续部署${NC}"
    fi
    
    # 上传部署包
    local DEPLOY_PACKAGE="./build/${APP_NAME}_${ENV_MODE}_package.tar.gz"
    local REMOTE_PACKAGE="${REMOTE_DIR}/tmp/package.tar.gz"
    
    echo -e "${BLUE}上传部署包...${NC}"
    echo -e "本地文件: ${DEPLOY_PACKAGE}"
    echo -e "目标位置: ${REMOTE_USER}@${REMOTE_HOST}:${REMOTE_PACKAGE}"
    
    # 确保远程临时目录存在
    $SSH_CMD "mkdir -p ${REMOTE_DIR}/tmp"
    
    if ! $SCP_CMD "$DEPLOY_PACKAGE" "${REMOTE_USER}@${REMOTE_HOST}:${REMOTE_PACKAGE}"; then
        echo -e "${RED}部署包上传失败${NC}"
        return 1
    fi
    
    # 解压部署包到发布目录
    echo -e "${BLUE}解压部署包到发布目录...${NC}"
    if ! $SSH_CMD "tar -xzf ${REMOTE_PACKAGE} -C ${RELEASE_DIR}"; then
        echo -e "${RED}解压部署包失败${NC}"
        return 1
    fi
    echo -e "${GREEN}部署包已解压到: ${RELEASE_DIR}${NC}"
    
    # 设置权限
    echo -e "${BLUE}设置文件权限...${NC}"
    $SSH_CMD "chmod +x ${RELEASE_DIR}/${APP_NAME} ${RELEASE_DIR}/run.sh"
    
    # 确保配置目录存在
    $SSH_CMD "mkdir -p ${REMOTE_DIR}/shared/config"
    
    # 复制配置文件到共享目录
    echo -e "${BLUE}保存配置文件到共享目录...${NC}"
    if $SSH_CMD "cp ${RELEASE_DIR}/config.yaml ${REMOTE_DIR}/shared/config/config.yaml"; then
        echo -e "${GREEN}配置文件已复制到共享目录${NC}"
    else
        echo -e "${RED}配置文件复制失败${NC}"
    fi
    
    # 创建配置软链接
    echo -e "${BLUE}创建配置软链接...${NC}"
    $SSH_CMD "ln -sf ${REMOTE_DIR}/shared/config/config.yaml ${RELEASE_DIR}/config.yaml"
    
    # 停止当前应用
    echo -e "${BLUE}停止当前应用...${NC}"
    $SSH_CMD "${REMOTE_DIR}/service.sh stop" || echo -e "${YELLOW}应用可能未在运行${NC}"
    
    # 更新当前版本链接
    echo -e "${BLUE}更新当前版本链接...${NC}"
    if ! $SSH_CMD "ln -sfn ${RELEASE_DIR} ${REMOTE_DIR}/current"; then
        echo -e "${RED}更新当前版本链接失败${NC}"
        return 1
    fi
    
    # 启动应用
    echo -e "${BLUE}启动应用...${NC}"
    if ! $SSH_CMD "${REMOTE_DIR}/service.sh start"; then
        echo -e "${RED}应用启动失败，请检查日志${NC}"
        # 继续流程，以便查看状态
    fi
    
    # 清理旧版本
    echo -e "${BLUE}清理旧版本...${NC}"
    $SSH_CMD "cd ${REMOTE_DIR}/releases && ls -t | tail -n +$((KEEP_RELEASES+1)) | xargs rm -rf 2>/dev/null || true"
    
    # 清理临时文件
    echo -e "${BLUE}清理临时文件...${NC}"
    $SSH_CMD "rm -f ${REMOTE_PACKAGE}"
    
    # 查看状态
    echo -e "${BLUE}应用状态:${NC}"
    $SSH_CMD "${REMOTE_DIR}/service.sh status" || echo -e "${RED}获取应用状态失败${NC}"
    
    echo -e "${GREEN}部署完成!${NC}"
    echo -e "应用已部署到: ${REMOTE_HOST}:${REMOTE_DIR}"
    echo -e "环境: ${ENV_MODE}"
    echo -e "当前版本: ${TIMESTAMP}"
}

# 回滚到上一个版本
rollback() {
    echo -e "${YELLOW}正在回滚到上一个版本...${NC}"
    
    if ! check_required_params; then
        return 1
    fi
    
    local SSH_CMD=$(build_ssh_cmd)
    
    # 获取当前版本和上一个版本
    local CURRENT_VERSION=$($SSH_CMD "readlink ${REMOTE_DIR}/current | xargs basename")
    local PREV_VERSION=$($SSH_CMD "ls -t ${REMOTE_DIR}/releases | sed -n '2p'")
    
    if [ -z "$PREV_VERSION" ]; then
        echo -e "${RED}错误: 没有找到可回滚的版本${NC}"
        return 1
    fi
    
    echo -e "当前版本: ${CURRENT_VERSION}"
    echo -e "回滚到版本: ${PREV_VERSION}"
    
    # 停止当前应用
    echo -e "${BLUE}停止当前应用...${NC}"
    $SSH_CMD "${REMOTE_DIR}/service.sh stop"
    
    # 更新当前版本链接
    echo -e "${BLUE}更新当前版本链接到上一个版本...${NC}"
    $SSH_CMD "ln -sfn ${REMOTE_DIR}/releases/${PREV_VERSION} ${REMOTE_DIR}/current"
    
    # 启动应用
    echo -e "${BLUE}启动应用...${NC}"
    $SSH_CMD "${REMOTE_DIR}/service.sh start"
    
    # 查看状态
    echo -e "${BLUE}应用状态:${NC}"
    $SSH_CMD "${REMOTE_DIR}/service.sh status"
    
    echo -e "${GREEN}回滚完成!${NC}"
}

# 重启远程应用
restart_app() {
    echo -e "${YELLOW}正在重启应用...${NC}"
    
    if ! check_required_params; then
        return 1
    fi
    
    local SSH_CMD=$(build_ssh_cmd)
    $SSH_CMD "${REMOTE_DIR}/service.sh restart"
}

# 查看应用状态
check_status() {
    if ! check_required_params; then
        return 1
    fi
    
    local SSH_CMD=$(build_ssh_cmd)
    echo -e "${YELLOW}应用状态:${NC}"
    $SSH_CMD "${REMOTE_DIR}/service.sh status"
}

# 查看应用日志
view_logs() {
    if ! check_required_params; then
        return 1
    fi
    
    local SSH_CMD=$(build_ssh_cmd)
    local LINES=${1:-50}
    
    echo -e "${YELLOW}应用日志 (最后 ${LINES} 行):${NC}"
    $SSH_CMD "${REMOTE_DIR}/service.sh logs ${LINES}"
}

# 处理命令行参数
process_args() {
    local COMMAND=""
    local LOGS_LINES=50
    
    while [ $# -gt 0 ]; do
        case "$1" in
            -h|--host)
                REMOTE_HOST="$2"
                shift 2
                ;;
            -u|--user)
                REMOTE_USER="$2"
                shift 2
                ;;
            -p|--port)
                REMOTE_PORT="$2"
                shift 2
                ;;
            -d|--dir)
                REMOTE_DIR="$2"
                shift 2
                ;;
            -n|--name)
                APP_NAME="$2"
                shift 2
                ;;
            -k|--key)
                SSH_KEY="$2"
                shift 2
                ;;
            -v|--version)
                VERSION="$2"
                shift 2
                ;;
            -e|--env)
                ENV_MODE="$2"
                # 自动设置对应的配置文件
                if [ -f "config.${ENV_MODE}.yaml" ]; then
                    CONFIG_FILE_PATH="config.${ENV_MODE}.yaml"
                fi
                shift 2
                ;;
            -c|--config)
                CONFIG_FILE_PATH="$2"
                shift 2
                ;;
            deploy|rollback|restart|status|logs|setup|config|envconfig|help)
                COMMAND="$1"
                shift
                if [ "$COMMAND" = "logs" ] && [[ "$1" =~ ^[0-9]+$ ]]; then
                    LOGS_LINES="$1"
                    shift
                fi
                ;;
            *)
                echo -e "${RED}错误: 未知选项或命令 $1${NC}"
                show_help
                exit 1
                ;;
        esac
    done
    
    # 执行命令
    case "$COMMAND" in
        deploy)
            deploy_app
            ;;
        rollback)
            rollback
            ;;
        restart)
            restart_app
            ;;
        status)
            check_status
            ;;
        logs)
            view_logs "$LOGS_LINES"
            ;;
        setup)
            setup_remote
            ;;
        config)
            configure
            ;;
        envconfig)
            configure_env
            ;;
        help|"")
            show_help
            ;;
    esac
}

# 主函数
main() {
    if [ $# -eq 0 ]; then
        show_help
        exit 0
    fi
    
    process_args "$@"
}

# 执行主函数
main "$@" 