#!/bin/bash

# 部署功能脚本 - 完整版
# 整合原deploy.sh的所有功能

# 颜色定义
RED='\033[31m'
GREEN='\033[32m'
YELLOW='\033[33m'
BLUE='\033[34m'
RESET='\033[0m'

# 获取应用端口
get_app_port() {
    local CONFIG_FILE="${1:-config.yaml}"
    # 尝试使用grep命令从配置文件中提取端口号
    if [ -f "$CONFIG_FILE" ]; then
        local PORT=$(grep -E "port:.*[0-9]+" "$CONFIG_FILE" | grep -oE "[0-9]+" | head -1)
        echo "$PORT"
    else
        echo ""
    fi
}

# 检测SSH密钥
detect_ssh_key() {
    local AUTO_KEY_PATH=""
    
    if [ -f "$HOME/.ssh/id_rsa" ]; then
        AUTO_KEY_PATH="$HOME/.ssh/id_rsa"
    elif [ -f "$HOME/.ssh/id_ed25519" ]; then
        AUTO_KEY_PATH="$HOME/.ssh/id_ed25519"
    elif [ -f "$HOME/.ssh/id_ecdsa" ]; then
        AUTO_KEY_PATH="$HOME/.ssh/id_ecdsa"
    else
        return 1
    fi
    
    echo "$AUTO_KEY_PATH"
    return 0
}

# SSH/SCP命令构建
build_ssh_cmd() {
    local SSH_KEY="$1"
    local REMOTE_PORT="$2"
    local REMOTE_USER="$3"
    local REMOTE_HOST="$4"
    
    local CMD="ssh"
    
    if [ -n "$SSH_KEY" ]; then
        CMD="$CMD -i $SSH_KEY"
    fi
    
    CMD="$CMD -p $REMOTE_PORT ${REMOTE_USER}@${REMOTE_HOST}"
    echo "$CMD"
}

build_scp_cmd() {
    local SSH_KEY="$1"
    local REMOTE_PORT="$2"
    
    local CMD="scp"
    
    if [ -n "$SSH_KEY" ]; then
        CMD="$CMD -i $SSH_KEY"
    fi
    
    CMD="$CMD -P $REMOTE_PORT"
    echo "$CMD"
}

# 检查必要参数
check_required_params() {
    local REMOTE_HOST="$1"
    local REMOTE_USER="$2"
    local MISSING=""
    
    if [ -z "$REMOTE_HOST" ]; then
        MISSING="$MISSING\n  - 远程主机地址"
    fi
    
    if [ -z "$REMOTE_USER" ]; then
        MISSING="$MISSING\n  - 远程用户名"
    fi
    
    if [ -n "$MISSING" ]; then
        echo -e "${RED}错误: 缺少必要参数:${RESET}$MISSING"
        echo -e "${YELLOW}提示: 请先运行 'make deploy-config' 配置部署参数${RESET}"
        return 1
    fi
    
    return 0
}

# 检查并处理端口占用
check_port_usage() {
    local SSH_KEY="$1"
    local REMOTE_PORT="$2" 
    local REMOTE_USER="$3"
    local REMOTE_HOST="$4"
    local REMOTE_DIR="$5"
    local APP_NAME="$6"
    local CONFIG_FILE="${7:-config.yaml}"
    
    echo -e "${BLUE}检查端口占用情况...${RESET}"
    
    local PORT_TO_CHECK=$(get_app_port "$CONFIG_FILE")
    if [ -z "$PORT_TO_CHECK" ]; then
        echo -e "${YELLOW}无法从配置文件中获取端口号，将使用默认端口9000${RESET}"
        PORT_TO_CHECK=9000
    fi
    
    local SSH_CMD=$(build_ssh_cmd "$SSH_KEY" "$REMOTE_PORT" "$REMOTE_USER" "$REMOTE_HOST")
    
    # 检查远程服务器上指定端口是否被占用
    echo -e "${BLUE}检查端口 ${PORT_TO_CHECK} 是否被占用...${RESET}"
    local PORT_USAGE=$($SSH_CMD "lsof -i:${PORT_TO_CHECK} | grep LISTEN || echo '未占用'")
    
    if [ "$PORT_USAGE" != "未占用" ]; then
        echo -e "${YELLOW}端口 ${PORT_TO_CHECK} 已被占用${RESET}"
        
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
                echo -e "${GREEN}检测到端口被自己的旧实例占用，将自动处理...${RESET}"
            fi
            
            if [ $IS_OWN_APP -eq 1 ]; then
                # 自动处理自己的旧实例
                echo -e "${BLUE}尝试停止旧应用实例...${RESET}"
                
                # 首先尝试使用服务脚本停止
                if $SSH_CMD "[ -f \"${REMOTE_DIR}/service.sh\" ]" 2>/dev/null; then
                    echo -e "${YELLOW}使用服务脚本停止旧应用...${RESET}"
                    $SSH_CMD "${REMOTE_DIR}/service.sh stop"
                    sleep 2
                    
                    # 检查端口是否已释放
                    PORT_USAGE=$($SSH_CMD "lsof -i:${PORT_TO_CHECK} | grep LISTEN || echo '未占用'")
                    if [ "$PORT_USAGE" == "未占用" ]; then
                        echo -e "${GREEN}端口 ${PORT_TO_CHECK} 已成功释放${RESET}"
                        return 0
                    fi
                fi
                
                # 如果服务脚本未能停止，直接终止进程
                echo -e "${YELLOW}服务脚本未能停止旧应用，尝试直接终止进程...${RESET}"
                
                for PID in $PID_LIST; do
                    echo -e "${BLUE}终止进程 ${PID}...${RESET}"
                    $SSH_CMD "kill -15 ${PID} 2>/dev/null"
                done
                
                sleep 2
                
                # 检查是否还有进程未终止
                local REMAINING_PIDS=$($SSH_CMD "lsof -i:${PORT_TO_CHECK} -t")
                if [ -n "$REMAINING_PIDS" ]; then
                    echo -e "${YELLOW}某些进程未响应SIGTERM，尝试强制终止...${RESET}"
                    for PID in $REMAINING_PIDS; do
                        $SSH_CMD "kill -9 ${PID} 2>/dev/null"
                    done
                    sleep 1
                fi
                
                # 最终检查
                PORT_USAGE=$($SSH_CMD "lsof -i:${PORT_TO_CHECK} | grep LISTEN || echo '未占用'")
                if [ "$PORT_USAGE" == "未占用" ]; then
                    echo -e "${GREEN}端口 ${PORT_TO_CHECK} 已成功释放${RESET}"
                    return 0
                else
                    echo -e "${RED}警告: 无法释放端口 ${PORT_TO_CHECK}${RESET}"
                    read -p "是否继续部署? (y/n) [n]: " continue_deploy
                    if [ "$continue_deploy" != "y" ] && [ "$continue_deploy" != "Y" ]; then
                        return 1
                    fi
                fi
            else
                # 端口被其他应用占用
                echo -e "${YELLOW}端口被其他应用占用，请手动处理...${RESET}"
                read -p "是否尝试释放占用的端口? (y/n) [n]: " release_port
                if [ "$release_port" = "y" ] || [ "$release_port" = "Y" ]; then
                    echo -e "${BLUE}尝试释放端口 ${PORT_TO_CHECK}...${RESET}"
                    
                    # 显示进程详情供用户确认
                    echo -e "占用端口的进程列表:"
                    for PID in $PID_LIST; do
                        local PROCESS_INFO=$($SSH_CMD "ps -p ${PID} -o pid,ppid,user,comm,args || echo '未知进程'")
                        echo -e "$PROCESS_INFO"
                        read -p "是否终止此进程? (y/n) [n]: " kill_process
                        if [ "$kill_process" = "y" ] || [ "$kill_process" = "Y" ]; then
                            echo -e "${YELLOW}终止进程 ${PID}...${RESET}"
                            $SSH_CMD "kill -15 ${PID} 2>/dev/null"
                            sleep 1
                            if $SSH_CMD "ps -p ${PID} > /dev/null 2>&1"; then
                                echo -e "${YELLOW}进程未响应，尝试强制终止...${RESET}"
                                $SSH_CMD "kill -9 ${PID} 2>/dev/null"
                            fi
                        fi
                    done
                    
                    # 检查端口是否已释放
                    PORT_USAGE=$($SSH_CMD "lsof -i:${PORT_TO_CHECK} | grep LISTEN || echo '未占用'")
                    if [ "$PORT_USAGE" == "未占用" ]; then
                        echo -e "${GREEN}端口 ${PORT_TO_CHECK} 已成功释放${RESET}"
                        return 0
                    else
                        echo -e "${RED}警告: 端口 ${PORT_TO_CHECK} 仍被占用${RESET}"
                        read -p "是否继续部署? (y/n) [n]: " continue_deploy
                        if [ "$continue_deploy" != "y" ] && [ "$continue_deploy" != "Y" ]; then
                            return 1
                        fi
                    fi
                else
                    echo -e "${RED}警告: 端口 ${PORT_TO_CHECK} 仍被占用，可能导致部署后应用无法启动${RESET}"
                    read -p "是否继续部署? (y/n) [n]: " continue_deploy
                    if [ "$continue_deploy" != "y" ] && [ "$continue_deploy" != "Y" ]; then
                        return 1
                    fi
                fi
            fi
        else
            echo -e "${RED}错误: 无法获取占用端口的进程ID${RESET}"
            read -p "是否继续部署? (y/n) [n]: " continue_deploy
            if [ "$continue_deploy" != "y" ] && [ "$continue_deploy" != "Y" ]; then
                return 1
            fi
        fi
    else
        echo -e "${GREEN}端口 ${PORT_TO_CHECK} 未被占用${RESET}"
    fi
    
    return 0
}

# 智能构建检查
check_local_package() {
    local APP_NAME="$1"
    local ENV_MODE="${2:-prod}"
    
    if [ ! -f "./build.sh" ]; then
        echo -e "${RED}错误: 未找到构建脚本 './build.sh'${RESET}"
        return 1
    fi
    
    echo -e "${BLUE}准备应用包...${RESET}"
    
    # 检查release目录是否存在
    local SHOULD_BUILD=0
    if [ ! -d "./release/linux_amd64" ]; then
        echo -e "${YELLOW}未找到Linux平台的构建包，将进行构建...${RESET}"
        SHOULD_BUILD=1
    else
        # 二进制文件存在，询问是否重新构建
        local LOCAL_PACKAGE="./release/linux_amd64/${APP_NAME}"
        if [ -f "$LOCAL_PACKAGE" ]; then
            # 获取文件修改时间
            local FILE_TIME=$(date -r "$LOCAL_PACKAGE" "+%Y-%m-%d %H:%M:%S")
            echo -e "${YELLOW}发现已存在的构建文件:${RESET}"
            echo -e "  路径: ${LOCAL_PACKAGE}"
            echo -e "  修改时间: ${FILE_TIME}"
            read -p "是否重新构建应用? (y/n) [n]: " rebuild_choice
            if [ "$rebuild_choice" = "y" ] || [ "$rebuild_choice" = "Y" ]; then
                SHOULD_BUILD=1
                echo -e "${YELLOW}将重新构建应用...${RESET}"
            else
                echo -e "${GREEN}使用现有构建文件${RESET}"
            fi
        else
            echo -e "${YELLOW}构建目录存在但未找到二进制文件，将进行构建...${RESET}"
            SHOULD_BUILD=1
        fi
    fi
    
    # 如果需要构建，则执行构建脚本
    if [ $SHOULD_BUILD -eq 1 ]; then
        ./build.sh --linux
        
        if [ ! -d "./release/linux_amd64" ]; then
            echo -e "${RED}错误: 构建Linux平台包失败${RESET}"
            return 1
        fi
    fi
    
    # 检查二进制文件
    local LOCAL_PACKAGE="./release/linux_amd64/${APP_NAME}"
    if [ ! -f "$LOCAL_PACKAGE" ]; then
        echo -e "${RED}错误: 未找到Linux平台的构建包: $LOCAL_PACKAGE${RESET}"
        return 1
    fi
    
    echo -e "${GREEN}应用二进制文件已准备好: $LOCAL_PACKAGE${RESET}"
    return 0
}

# 创建部署包
create_deploy_package() {
    local APP_NAME="$1"
    local ENV_MODE="${2:-prod}"
    local CONFIG_FILE_PATH="${3:-config.yaml}"
    
    echo -e "${BLUE}创建部署包...${RESET}" >&2
    
    # 创建部署包目录
    local DEPLOY_TEMP_DIR="./release/deploy_package"
    rm -rf "$DEPLOY_TEMP_DIR"
    mkdir -p "$DEPLOY_TEMP_DIR"
    
    # 复制二进制文件
    local LOCAL_PACKAGE="./release/linux_amd64/${APP_NAME}"
    cp "$LOCAL_PACKAGE" "$DEPLOY_TEMP_DIR/"
    
    # 检查配置文件路径
    local CONFIG_SRC="config.yaml"  # 默认配置文件
    if [ -n "$CONFIG_FILE_PATH" ] && [ -f "$CONFIG_FILE_PATH" ]; then
        CONFIG_SRC="$CONFIG_FILE_PATH"
        echo -e "${BLUE}使用环境特定配置: $CONFIG_SRC${RESET}" >&2
    else
        echo -e "${YELLOW}使用默认配置文件: $CONFIG_SRC${RESET}" >&2
    fi
    
    # 复制配置文件到部署包
    if [ -f "$CONFIG_SRC" ]; then
        cp "$CONFIG_SRC" "$DEPLOY_TEMP_DIR/config.yaml"
        echo -e "${GREEN}配置文件已添加到部署包: $CONFIG_SRC${RESET}" >&2
    else
        echo -e "${RED}警告: 配置文件不存在: $CONFIG_SRC${RESET}" >&2
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
    echo -e "${GREEN}启动脚本已添加到部署包${RESET}" >&2
    
    # 打包文件
    local DEPLOY_PACKAGE="./release/${APP_NAME}_${ENV_MODE}_package.tar.gz"
    tar -czf "$DEPLOY_PACKAGE" -C "$DEPLOY_TEMP_DIR" .
    
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}部署包已创建: $DEPLOY_PACKAGE${RESET}" >&2
        
        # 清理临时目录
        rm -rf "$DEPLOY_TEMP_DIR"
        
        echo "$DEPLOY_PACKAGE"
        return 0
    else
        echo -e "${RED}创建部署包失败${RESET}" >&2
        return 1
    fi
}

# 生成服务管理脚本
generate_service_script() {
    local app_name="$1"
    local deploy_dir="$2"
    local script_path="$3"
    
    cat > "$script_path" << 'EOF'
#!/bin/bash
APP_NAME="REPLACE_APP_NAME"
APP_DIR="REPLACE_DEPLOY_DIR"
PID_FILE="$APP_DIR/shared/$APP_NAME.pid"
LOG_FILE="$APP_DIR/shared/$APP_NAME.log"
CURRENT_LINK="$APP_DIR/current"
ENV_FILE="$APP_DIR/shared/.env"

# 确保日志目录存在
mkdir -p "$(dirname $LOG_FILE)"

# 获取环境变量
if [ -f "$ENV_FILE" ]; then
    source "$ENV_FILE"
fi

start() {
    if [ -f "$PID_FILE" ] && ps -p $(cat "$PID_FILE") > /dev/null 2>&1; then
        echo "应用已在运行，PID: $(cat "$PID_FILE")"
        return 1
    fi
    
    if [ ! -L "$CURRENT_LINK" ]; then
        echo "错误: 未找到当前版本链接"
        return 1
    fi
    
    cd "$CURRENT_LINK"
    nohup ./$APP_NAME > "$LOG_FILE" 2>&1 &
    echo $! > "$PID_FILE"
    echo "应用已启动，PID: $(cat "$PID_FILE")"
}

stop() {
    if [ ! -f "$PID_FILE" ]; then
        echo "应用未在运行"
        return 0
    fi
    
    local PID=$(cat "$PID_FILE")
    if ! ps -p "$PID" > /dev/null 2>&1; then
        echo "应用未在运行 (PID文件存在但进程不存在)"
        rm -f "$PID_FILE"
        return 0
    fi
    
    echo "正在停止应用，PID: $PID"
    kill "$PID"
    sleep 2
    
    if ps -p "$PID" > /dev/null 2>&1; then
        echo "强制终止应用，PID: $PID"
        kill -9 "$PID"
        sleep 1
    fi
    
    rm -f "$PID_FILE"
    echo "应用已停止"
}

status() {
    if [ ! -f "$PID_FILE" ]; then
        echo "应用未在运行"
        return 1
    fi
    
    local PID=$(cat "$PID_FILE")
    if ! ps -p "$PID" > /dev/null 2>&1; then
        echo "应用未在运行 (PID文件存在但进程不存在)"
        rm -f "$PID_FILE"
        return 1
    fi
    
    echo "应用正在运行，PID: $PID"
    echo "日志文件: $LOG_FILE"
    echo "当前版本: $(readlink $CURRENT_LINK 2>/dev/null || echo "未知")"
    echo "当前环境: $(grep -o 'APP_ENV=.*' $ENV_FILE 2>/dev/null | cut -d= -f2 || echo "未知")"
    return 0
}

logs() {
    local LINES=${1:-50}
    if [ -f "$LOG_FILE" ]; then
        tail -n $LINES "$LOG_FILE"
    else
        echo "日志文件不存在: $LOG_FILE"
        return 1
    fi
}

case "$1" in
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
        logs $2
        ;;
    *)
        echo "用法: $0 {start|stop|restart|status|logs [lines]}"
        exit 1
        ;;
esac

exit 0
EOF

    # 替换占位符
    sed -i.bak "s/REPLACE_APP_NAME/$app_name/g" "$script_path"
    sed -i.bak "s#REPLACE_DEPLOY_DIR#$deploy_dir#g" "$script_path"
    rm -f "$script_path.bak"
}

# 准备环境变量文件
prepare_env_file() {
    local SSH_KEY="$1"
    local REMOTE_PORT="$2"
    local REMOTE_USER="$3"
    local REMOTE_HOST="$4"
    local REMOTE_DIR="$5"
    local ENV_MODE="$6"
    local VERSION="$7"
    
    echo -e "${BLUE}准备环境变量文件...${RESET}"
    
    cat > ./tmp_env << EOF
# 环境变量配置 - 由部署脚本自动生成
APP_ENV=${ENV_MODE}
APP_VERSION=${VERSION}
APP_DEPLOY_TIME="$(date +"%Y-%m-%d %H:%M:%S")"
EOF
    
    local SCP_CMD=$(build_scp_cmd "$SSH_KEY" "$REMOTE_PORT")
    local SSH_CMD=$(build_ssh_cmd "$SSH_KEY" "$REMOTE_PORT" "$REMOTE_USER" "$REMOTE_HOST")
    
    # 确保目录存在
    $SSH_CMD "mkdir -p ${REMOTE_DIR}/shared"
    
    # 上传文件并检查结果
    if $SCP_CMD ./tmp_env ${REMOTE_USER}@${REMOTE_HOST}:${REMOTE_DIR}/shared/.env; then
        echo -e "${GREEN}环境变量文件已上传${RESET}"
        rm -f ./tmp_env
        return 0
    else
        echo -e "${RED}环境变量文件上传失败${RESET}"
        rm -f ./tmp_env
        return 1
    fi
}

# 配置部署参数 (增强版)
config_deploy() {
    local app_name="$1"
    local config_file="deploy.conf"
    
    echo -e "${BLUE}配置部署参数${RESET}"
    echo ""
    
    read -p "应用名称 [$app_name]: " input_app_name
    app_name="${input_app_name:-$app_name}"
    
    read -p "远程主机地址: " deploy_host
    if [ -z "$deploy_host" ]; then
        echo -e "${RED}错误: 远程主机地址不能为空${RESET}"
        return 1
    fi
    
    read -p "远程用户名 [root]: " deploy_user
    deploy_user="${deploy_user:-root}"
    
    read -p "SSH端口 [22]: " deploy_port
    deploy_port="${deploy_port:-22}"
    
    read -p "远程部署目录 [/opt/myapp]: " deploy_dir
    deploy_dir="${deploy_dir:-/opt/myapp}"
    
    # SSH密钥检测和选择
    echo -e "${BLUE}检测SSH密钥...${RESET}"
    auto_key=$(detect_ssh_key 2>/dev/null)
    deploy_key=""
    
    if [ $? -eq 0 ] && [ -n "$auto_key" ]; then
        echo -e "${GREEN}找到SSH密钥: $auto_key${RESET}"
        echo -e "${BLUE}是否使用检测到的SSH密钥?${RESET}"
        echo "  路径: $auto_key"
        read -p "使用此密钥? (y/n) [y]: " use_auto_key
        
        if [ "$use_auto_key" != "n" ] && [ "$use_auto_key" != "N" ]; then
            deploy_key="$auto_key"
        else
            read -p "自定义SSH密钥路径 (留空使用密码): " custom_key
            if [ -n "$custom_key" ] && [ ! -f "$custom_key" ]; then
                echo -e "${RED}错误: SSH密钥文件不存在: $custom_key${RESET}"
                return 1
            fi
            deploy_key="$custom_key"
        fi
    else
        echo -e "${YELLOW}未找到常见的SSH密钥文件${RESET}"
        echo -e "${YELLOW}常见位置: ~/.ssh/id_rsa, ~/.ssh/id_ed25519, ~/.ssh/id_ecdsa${RESET}"
        echo ""
        echo -e "${BLUE}建议生成SSH密钥:${RESET}"
        echo "  ssh-keygen -t rsa -b 4096 -f ~/.ssh/id_rsa"
        echo "  ssh-copy-id -i ~/.ssh/id_rsa.pub $deploy_user@$deploy_host"
        echo ""
        read -p "是否继续配置 (将使用密码认证)? (y/n) [y]: " continue_config
        if [ "$continue_config" = "n" ] || [ "$continue_config" = "N" ]; then
            echo -e "${YELLOW}配置已取消，请先生成SSH密钥${RESET}"
            return 1
        fi
    fi
    
    read -p "部署环境 (prod/test/dev) [prod]: " deploy_env
    deploy_env="${deploy_env:-prod}"
    
    read -p "保留版本数量 [5]: " keep_releases
    keep_releases="${keep_releases:-5}"
    
    # 保存配置
    cat > "$config_file" << EOF
# 部署配置文件 - 由 deploy-config 生成
APP_NAME := $app_name
DEPLOY_HOST := $deploy_host
DEPLOY_USER := $deploy_user
DEPLOY_PORT := $deploy_port
DEPLOY_DIR := $deploy_dir
DEPLOY_KEY := $deploy_key
DEPLOY_ENV := $deploy_env
KEEP_RELEASES := $keep_releases
EOF
    
    echo -e "${GREEN}配置已保存到 $config_file${RESET}"
    echo -e "${YELLOW}下一步: 运行 'make deploy-setup' 初始化服务器环境${RESET}"
    return 0
}

# 环境配置
config_env() {
    local config_file="deploy.conf"
    
    echo -e "${BLUE}配置环境特定参数${RESET}"
    echo ""
    
    # 读取环境模式
    echo "部署环境选项:"
    echo "  1) 生产环境 (prod)"
    echo "  2) 测试环境 (test)"  
    echo "  3) 开发环境 (dev)"
    read -p "选择部署环境 [1-3] (默认: 1): " env_choice
    
    local ENV_MODE="prod"
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
    local CONFIG_FILE_PATH=""
    if [ -f "$env_config" ]; then
        echo -e "${GREEN}发现环境配置文件: $env_config${RESET}"
        CONFIG_FILE_PATH="$env_config"
    else
        echo -e "${YELLOW}未找到环境配置文件: $env_config${RESET}"
        echo -e "您可以创建环境特定的配置文件，或者选择使用本地配置文件上传"
        read -p "使用的配置文件路径 [config.yaml]: " config_path
        if [ -n "$config_path" ]; then
            if [ -f "$config_path" ]; then
                CONFIG_FILE_PATH="$config_path"
            else
                echo -e "${RED}错误: 配置文件不存在: $config_path${RESET}"
                CONFIG_FILE_PATH=""
            fi
        else
            CONFIG_FILE_PATH="config.yaml"
        fi
    fi
    
    # 如果存在deploy.conf，则更新，否则创建
    if [ -f "$config_file" ]; then
        # 更新现有配置
        if grep -q "DEPLOY_ENV" "$config_file"; then
            sed -i.bak "s/DEPLOY_ENV := .*/DEPLOY_ENV := $ENV_MODE/" "$config_file"
        else
            echo "DEPLOY_ENV := $ENV_MODE" >> "$config_file"
        fi
        
        if [ -n "$CONFIG_FILE_PATH" ]; then
            if grep -q "CONFIG_FILE_PATH" "$config_file"; then
                sed -i.bak "s#CONFIG_FILE_PATH := .*#CONFIG_FILE_PATH := $CONFIG_FILE_PATH#" "$config_file"
            else
                echo "CONFIG_FILE_PATH := $CONFIG_FILE_PATH" >> "$config_file"
            fi
        fi
        rm -f "$config_file.bak"
    else
        echo -e "${RED}错误: 配置文件不存在，请先运行 'make deploy-config'${RESET}"
        return 1
    fi
    
    echo -e "${GREEN}环境配置完成!${RESET}"
    echo -e "环境: ${ENV_MODE}"
    echo -e "配置文件: ${CONFIG_FILE_PATH:-默认}"
    echo ""
    echo -e "使用 'make deploy' 部署应用"
    return 0
}

# 主函数
case "$1" in
    config)
        config_deploy "$2"
        ;;
    config_env)
        config_env
        ;;
    generate_service)
        generate_service_script "$2" "$3" "$4"
        ;;
    detect_key)
        detect_ssh_key
        ;;
    check_package)
        check_local_package "$2" "$3"
        ;;
    create_package)
        create_deploy_package "$2" "$3" "$4"
        ;;
    check_port)
        check_port_usage "$2" "$3" "$4" "$5" "$6" "$7" "$8"
        ;;
    prepare_env)
        prepare_env_file "$2" "$3" "$4" "$5" "$6" "$7" "$8"
        ;;
    check_params)
        check_required_params "$2" "$3"
        ;;
    build_ssh)
        build_ssh_cmd "$2" "$3" "$4" "$5"
        ;;
    build_scp)
        build_scp_cmd "$2" "$3"
        ;;
    *)
        echo "用法: $0 {config|config_env|generate_service|detect_key|check_package|create_package|check_port|prepare_env|check_params|build_ssh|build_scp} [参数...]"
        exit 1
        ;;
esac 