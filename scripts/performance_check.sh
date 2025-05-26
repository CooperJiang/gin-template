#!/bin/bash

# 性能检查和优化脚本
# 用于分析应用性能并提供优化建议

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 配置
APP_URL="http://localhost:8080"
APP_PID=""
REPORT_FILE="performance_report_$(date +%Y%m%d_%H%M%S).txt"

# 打印带颜色的消息
print_message() {
    local color=$1
    local message=$2
    echo -e "${color}${message}${NC}"
}

# 打印标题
print_title() {
    echo ""
    print_message $BLUE "=================================="
    print_message $BLUE "$1"
    print_message $BLUE "=================================="
    echo ""
}

# 检查应用是否运行
check_app_running() {
    print_title "检查应用状态"
    
    # 检查端口是否被占用
    if netstat -tlnp 2>/dev/null | grep -q ":8080 "; then
        APP_PID=$(netstat -tlnp 2>/dev/null | grep ":8080 " | awk '{print $7}' | cut -d'/' -f1)
        print_message $GREEN "✓ 应用正在运行 (PID: $APP_PID)"
        return 0
    else
        print_message $RED "✗ 应用未运行在端口 8080"
        return 1
    fi
}

# 检查应用健康状态
check_app_health() {
    print_title "检查应用健康状态"
    
    # 检查健康端点
    if curl -s -f "$APP_URL/health" > /dev/null 2>&1; then
        print_message $GREEN "✓ 健康检查通过"
    else
        print_message $YELLOW "⚠ 健康检查失败或端点不存在"
    fi
    
    # 检查API响应
    response_time=$(curl -o /dev/null -s -w "%{time_total}" "$APP_URL/api/v1/health" 2>/dev/null || echo "0")
    if (( $(echo "$response_time > 0" | bc -l) )); then
        print_message $GREEN "✓ API响应时间: ${response_time}s"
        if (( $(echo "$response_time > 1" | bc -l) )); then
            print_message $YELLOW "⚠ API响应时间较慢，建议优化"
        fi
    else
        print_message $YELLOW "⚠ 无法获取API响应时间"
    fi
}

# 系统资源检查
check_system_resources() {
    print_title "系统资源检查"
    
    # CPU使用率
    cpu_usage=$(top -bn1 | grep "Cpu(s)" | awk '{print $2}' | cut -d'%' -f1)
    print_message $BLUE "CPU使用率: ${cpu_usage}%"
    if (( $(echo "$cpu_usage > 80" | bc -l) )); then
        print_message $RED "✗ CPU使用率过高"
    elif (( $(echo "$cpu_usage > 60" | bc -l) )); then
        print_message $YELLOW "⚠ CPU使用率较高"
    else
        print_message $GREEN "✓ CPU使用率正常"
    fi
    
    # 内存使用率
    memory_info=$(free | grep Mem)
    total_mem=$(echo $memory_info | awk '{print $2}')
    used_mem=$(echo $memory_info | awk '{print $3}')
    memory_usage=$(echo "scale=2; $used_mem * 100 / $total_mem" | bc)
    
    print_message $BLUE "内存使用率: ${memory_usage}%"
    if (( $(echo "$memory_usage > 90" | bc -l) )); then
        print_message $RED "✗ 内存使用率过高"
    elif (( $(echo "$memory_usage > 75" | bc -l) )); then
        print_message $YELLOW "⚠ 内存使用率较高"
    else
        print_message $GREEN "✓ 内存使用率正常"
    fi
    
    # 磁盘使用率
    disk_usage=$(df -h / | awk 'NR==2 {print $5}' | cut -d'%' -f1)
    print_message $BLUE "磁盘使用率: ${disk_usage}%"
    if (( disk_usage > 90 )); then
        print_message $RED "✗ 磁盘使用率过高"
    elif (( disk_usage > 75 )); then
        print_message $YELLOW "⚠ 磁盘使用率较高"
    else
        print_message $GREEN "✓ 磁盘使用率正常"
    fi
}

# 应用进程检查
check_app_process() {
    if [ -z "$APP_PID" ]; then
        return
    fi
    
    print_title "应用进程分析"
    
    # 进程内存使用
    if ps -p $APP_PID > /dev/null 2>&1; then
        process_info=$(ps -p $APP_PID -o pid,ppid,pcpu,pmem,vsz,rss,comm --no-headers)
        cpu_percent=$(echo $process_info | awk '{print $3}')
        mem_percent=$(echo $process_info | awk '{print $4}')
        vsz=$(echo $process_info | awk '{print $5}')
        rss=$(echo $process_info | awk '{print $6}')
        
        print_message $BLUE "进程CPU使用率: ${cpu_percent}%"
        print_message $BLUE "进程内存使用率: ${mem_percent}%"
        print_message $BLUE "虚拟内存: $(echo "scale=2; $vsz / 1024" | bc) MB"
        print_message $BLUE "物理内存: $(echo "scale=2; $rss / 1024" | bc) MB"
        
        # 检查是否有内存泄漏的迹象
        if (( $(echo "$mem_percent > 10" | bc -l) )); then
            print_message $YELLOW "⚠ 进程内存使用率较高，请关注是否存在内存泄漏"
        fi
    else
        print_message $RED "✗ 无法获取进程信息"
    fi
}

# 数据库连接检查
check_database() {
    print_title "数据库连接检查"
    
    # 检查SQLite数据库文件
    if [ -f "app.db" ]; then
        db_size=$(du -h app.db | cut -f1)
        print_message $GREEN "✓ SQLite数据库文件存在 (大小: $db_size)"
        
        # 检查数据库文件权限
        db_perms=$(stat -c "%a" app.db)
        if [ "$db_perms" = "644" ] || [ "$db_perms" = "600" ]; then
            print_message $GREEN "✓ 数据库文件权限正常"
        else
            print_message $YELLOW "⚠ 数据库文件权限: $db_perms (建议设置为600或644)"
        fi
    else
        print_message $YELLOW "⚠ SQLite数据库文件不存在"
    fi
    
    # 检查MySQL连接（如果配置了）
    if command -v mysql > /dev/null 2>&1; then
        if mysql -h localhost -u root -e "SELECT 1;" > /dev/null 2>&1; then
            print_message $GREEN "✓ MySQL连接正常"
        else
            print_message $YELLOW "⚠ MySQL连接失败或未配置"
        fi
    fi
}

# 网络性能测试
test_network_performance() {
    print_title "网络性能测试"
    
    if ! command -v curl > /dev/null 2>&1; then
        print_message $YELLOW "⚠ curl未安装，跳过网络性能测试"
        return
    fi
    
    # 测试API响应时间
    print_message $BLUE "测试API响应时间..."
    
    total_time=0
    success_count=0
    
    for i in {1..10}; do
        response_time=$(curl -o /dev/null -s -w "%{time_total}" "$APP_URL/api/v1/health" 2>/dev/null || echo "0")
        if (( $(echo "$response_time > 0" | bc -l) )); then
            total_time=$(echo "$total_time + $response_time" | bc)
            success_count=$((success_count + 1))
        fi
        echo -n "."
    done
    echo ""
    
    if [ $success_count -gt 0 ]; then
        avg_time=$(echo "scale=3; $total_time / $success_count" | bc)
        print_message $GREEN "✓ 平均响应时间: ${avg_time}s (成功率: $success_count/10)"
        
        if (( $(echo "$avg_time > 1" | bc -l) )); then
            print_message $YELLOW "⚠ 响应时间较慢，建议优化"
        fi
    else
        print_message $RED "✗ 所有请求都失败了"
    fi
}

# 日志文件检查
check_logs() {
    print_title "日志文件检查"
    
    # 检查常见日志文件
    log_files=("logs/app.log" "app.log" "/var/log/template/app.log")
    
    for log_file in "${log_files[@]}"; do
        if [ -f "$log_file" ]; then
            log_size=$(du -h "$log_file" | cut -f1)
            print_message $GREEN "✓ 日志文件: $log_file (大小: $log_size)"
            
            # 检查最近的错误
            error_count=$(tail -n 1000 "$log_file" | grep -i "error\|fatal\|panic" | wc -l)
            if [ $error_count -gt 0 ]; then
                print_message $YELLOW "⚠ 最近1000行日志中发现 $error_count 个错误"
                print_message $BLUE "最近的错误:"
                tail -n 1000 "$log_file" | grep -i "error\|fatal\|panic" | tail -n 3
            else
                print_message $GREEN "✓ 最近日志中无错误"
            fi
            break
        fi
    done
}

# 配置文件检查
check_config() {
    print_title "配置文件检查"
    
    if [ -f "config.yaml" ]; then
        print_message $GREEN "✓ 配置文件存在"
        
        # 检查配置文件权限
        config_perms=$(stat -c "%a" config.yaml)
        if [ "$config_perms" = "600" ] || [ "$config_perms" = "644" ]; then
            print_message $GREEN "✓ 配置文件权限正常"
        else
            print_message $YELLOW "⚠ 配置文件权限: $config_perms (建议设置为600)"
        fi
        
        # 检查敏感信息
        if grep -q "password.*password\|secret.*secret" config.yaml; then
            print_message $RED "✗ 配置文件中可能包含默认密码，请检查安全性"
        else
            print_message $GREEN "✓ 配置文件安全检查通过"
        fi
    else
        print_message $YELLOW "⚠ 配置文件不存在"
    fi
}

# 生成优化建议
generate_recommendations() {
    print_title "性能优化建议"
    
    print_message $BLUE "基于检查结果，以下是优化建议："
    echo ""
    
    # CPU优化建议
    cpu_usage=$(top -bn1 | grep "Cpu(s)" | awk '{print $2}' | cut -d'%' -f1)
    if (( $(echo "$cpu_usage > 60" | bc -l) )); then
        print_message $YELLOW "🔧 CPU优化:"
        echo "   - 考虑增加GOMAXPROCS设置"
        echo "   - 优化算法复杂度"
        echo "   - 使用性能分析工具: go tool pprof"
        echo ""
    fi
    
    # 内存优化建议
    memory_info=$(free | grep Mem)
    total_mem=$(echo $memory_info | awk '{print $2}')
    used_mem=$(echo $memory_info | awk '{print $3}')
    memory_usage=$(echo "scale=2; $used_mem * 100 / $total_mem" | bc)
    
    if (( $(echo "$memory_usage > 75" | bc -l) )); then
        print_message $YELLOW "🔧 内存优化:"
        echo "   - 检查是否存在内存泄漏"
        echo "   - 优化数据结构使用"
        echo "   - 调整垃圾回收参数: GOGC"
        echo "   - 考虑使用对象池"
        echo ""
    fi
    
    # 数据库优化建议
    if [ -f "app.db" ]; then
        db_size_bytes=$(stat -c%s app.db)
        if [ $db_size_bytes -gt 104857600 ]; then  # 100MB
            print_message $YELLOW "🔧 数据库优化:"
            echo "   - 数据库文件较大，考虑数据清理"
            echo "   - 添加适当的索引"
            echo "   - 考虑数据分页查询"
            echo "   - 定期执行VACUUM操作"
            echo ""
        fi
    fi
    
    # 网络优化建议
    print_message $YELLOW "🔧 网络优化:"
    echo "   - 启用HTTP/2"
    echo "   - 使用CDN加速静态资源"
    echo "   - 启用gzip压缩"
    echo "   - 配置适当的缓存策略"
    echo ""
    
    # 监控建议
    print_message $YELLOW "🔧 监控建议:"
    echo "   - 设置应用性能监控(APM)"
    echo "   - 配置日志聚合和分析"
    echo "   - 设置关键指标告警"
    echo "   - 定期执行性能测试"
    echo ""
}

# 生成报告
generate_report() {
    print_title "生成性能报告"
    
    {
        echo "性能检查报告"
        echo "生成时间: $(date)"
        echo "========================================"
        echo ""
        
        echo "系统信息:"
        echo "- 操作系统: $(uname -s)"
        echo "- 内核版本: $(uname -r)"
        echo "- CPU核心数: $(nproc)"
        echo "- 总内存: $(free -h | grep Mem | awk '{print $2}')"
        echo ""
        
        if [ -n "$APP_PID" ]; then
            echo "应用信息:"
            echo "- 进程ID: $APP_PID"
            echo "- 启动时间: $(ps -o lstart= -p $APP_PID 2>/dev/null || echo '未知')"
            echo ""
        fi
        
        echo "检查结果已保存到控制台输出"
        echo "建议定期运行此脚本以监控应用性能"
        
    } > "$REPORT_FILE"
    
    print_message $GREEN "✓ 报告已保存到: $REPORT_FILE"
}

# 主函数
main() {
    print_message $GREEN "开始性能检查..."
    echo ""
    
    # 检查必要的命令
    for cmd in netstat curl bc; do
        if ! command -v $cmd > /dev/null 2>&1; then
            print_message $YELLOW "⚠ $cmd 未安装，某些检查可能无法执行"
        fi
    done
    
    # 执行各项检查
    if check_app_running; then
        check_app_health
        check_app_process
    fi
    
    check_system_resources
    check_database
    test_network_performance
    check_logs
    check_config
    generate_recommendations
    generate_report
    
    print_title "性能检查完成"
    print_message $GREEN "所有检查已完成，请查看上述结果和建议"
}

# 显示帮助信息
show_help() {
    echo "性能检查脚本"
    echo ""
    echo "用法: $0 [选项]"
    echo ""
    echo "选项:"
    echo "  -h, --help     显示帮助信息"
    echo "  -u, --url URL  指定应用URL (默认: http://localhost:8080)"
    echo ""
    echo "示例:"
    echo "  $0                           # 使用默认设置"
    echo "  $0 -u http://localhost:9000  # 指定应用URL"
}

# 解析命令行参数
while [[ $# -gt 0 ]]; do
    case $1 in
        -h|--help)
            show_help
            exit 0
            ;;
        -u|--url)
            APP_URL="$2"
            shift 2
            ;;
        *)
            echo "未知选项: $1"
            show_help
            exit 1
            ;;
    esac
done

# 运行主函数
main 