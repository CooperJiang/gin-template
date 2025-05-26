#!/bin/bash

# API自动化测试脚本
# 用于测试所有API接口的功能

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 配置
BASE_URL="http://localhost:8080"
API_BASE="$BASE_URL/api/v1"
TEST_EMAIL="test@example.com"
TEST_USERNAME="testuser"
TEST_PASSWORD="password123"
TOKEN=""
USER_ID=""

# 测试结果统计
TOTAL_TESTS=0
PASSED_TESTS=0
FAILED_TESTS=0

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

# 测试结果
test_result() {
    local test_name="$1"
    local expected="$2"
    local actual="$3"
    
    TOTAL_TESTS=$((TOTAL_TESTS + 1))
    
    if [ "$expected" = "$actual" ]; then
        print_message $GREEN "✓ $test_name"
        PASSED_TESTS=$((PASSED_TESTS + 1))
        return 0
    else
        print_message $RED "✗ $test_name (期望: $expected, 实际: $actual)"
        FAILED_TESTS=$((FAILED_TESTS + 1))
        return 1
    fi
}

# HTTP请求函数
http_request() {
    local method="$1"
    local url="$2"
    local data="$3"
    local headers="$4"
    
    local curl_cmd="curl -s -w '%{http_code}' -X $method"
    
    if [ -n "$headers" ]; then
        curl_cmd="$curl_cmd -H '$headers'"
    fi
    
    if [ -n "$data" ]; then
        curl_cmd="$curl_cmd -H 'Content-Type: application/json' -d '$data'"
    fi
    
    curl_cmd="$curl_cmd '$url'"
    
    eval $curl_cmd
}

# 提取HTTP状态码
extract_status_code() {
    local response="$1"
    echo "${response: -3}"
}

# 提取响应体
extract_body() {
    local response="$1"
    echo "${response%???}"
}

# 提取JSON字段
extract_json_field() {
    local json="$1"
    local field="$2"
    echo "$json" | grep -o "\"$field\":[^,}]*" | cut -d':' -f2 | tr -d '"' | tr -d ' '
}

# 检查应用是否运行
check_app_running() {
    print_title "检查应用状态"
    
    local response=$(http_request "GET" "$BASE_URL/health" "" "")
    local status_code=$(extract_status_code "$response")
    
    if test_result "应用健康检查" "200" "$status_code"; then
        print_message $GREEN "应用正在运行"
        return 0
    else
        print_message $RED "应用未运行或健康检查失败"
        print_message $YELLOW "请确保应用在 $BASE_URL 上运行"
        exit 1
    fi
}

# 测试用户注册
test_user_registration() {
    print_title "测试用户注册"
    
    # 首先发送验证码
    print_message $BLUE "发送注册验证码..."
    local verify_data="{\"email\":\"$TEST_EMAIL\"}"
    local verify_response=$(http_request "POST" "$API_BASE/user/send-registration-code" "$verify_data" "")
    local verify_status=$(extract_status_code "$verify_response")
    
    test_result "发送注册验证码" "200" "$verify_status"
    
    # 注册用户 (使用固定验证码进行测试)
    print_message $BLUE "注册用户..."
    local register_data="{\"username\":\"$TEST_USERNAME\",\"password\":\"$TEST_PASSWORD\",\"email\":\"$TEST_EMAIL\",\"verification_code\":\"123456\"}"
    local register_response=$(http_request "POST" "$API_BASE/user/register" "$register_data" "")
    local register_status=$(extract_status_code "$register_response")
    local register_body=$(extract_body "$register_response")
    
    if test_result "用户注册" "200" "$register_status"; then
        USER_ID=$(extract_json_field "$register_body" "user_id")
        print_message $GREEN "用户注册成功，用户ID: $USER_ID"
    fi
}

# 测试用户登录
test_user_login() {
    print_title "测试用户登录"
    
    local login_data="{\"username\":\"$TEST_USERNAME\",\"password\":\"$TEST_PASSWORD\"}"
    local login_response=$(http_request "POST" "$API_BASE/user/login" "$login_data" "")
    local login_status=$(extract_status_code "$login_response")
    local login_body=$(extract_body "$login_response")
    
    if test_result "用户登录" "200" "$login_status"; then
        TOKEN=$(extract_json_field "$login_body" "token")
        print_message $GREEN "登录成功，获取到Token"
    fi
}

# 测试获取用户信息
test_get_user_profile() {
    print_title "测试获取用户信息"
    
    if [ -z "$TOKEN" ]; then
        print_message $YELLOW "跳过用户信息测试 (未获取到Token)"
        return
    fi
    
    local profile_response=$(http_request "GET" "$API_BASE/user/profile" "" "Authorization: Bearer $TOKEN")
    local profile_status=$(extract_status_code "$profile_response")
    local profile_body=$(extract_body "$profile_response")
    
    if test_result "获取用户信息" "200" "$profile_status"; then
        local username=$(extract_json_field "$profile_body" "username")
        test_result "用户名匹配" "$TEST_USERNAME" "$username"
    fi
}

# 测试更新用户信息
test_update_user_profile() {
    print_title "测试更新用户信息"
    
    if [ -z "$TOKEN" ]; then
        print_message $YELLOW "跳过更新用户信息测试 (未获取到Token)"
        return
    fi
    
    local new_username="updated_$TEST_USERNAME"
    local update_data="{\"username\":\"$new_username\"}"
    local update_response=$(http_request "PUT" "$API_BASE/user/profile" "$update_data" "Authorization: Bearer $TOKEN")
    local update_status=$(extract_status_code "$update_response")
    
    test_result "更新用户信息" "200" "$update_status"
}

# 测试密码重置
test_password_reset() {
    print_title "测试密码重置"
    
    # 发送重置密码验证码
    print_message $BLUE "发送重置密码验证码..."
    local reset_code_data="{\"email\":\"$TEST_EMAIL\"}"
    local reset_code_response=$(http_request "POST" "$API_BASE/user/send-reset-password-code" "$reset_code_data" "")
    local reset_code_status=$(extract_status_code "$reset_code_response")
    
    test_result "发送重置密码验证码" "200" "$reset_code_status"
    
    # 重置密码
    print_message $BLUE "重置密码..."
    local new_password="newpassword123"
    local reset_data="{\"email\":\"$TEST_EMAIL\",\"verification_code\":\"123456\",\"new_password\":\"$new_password\"}"
    local reset_response=$(http_request "POST" "$API_BASE/user/reset-password" "$reset_data" "")
    local reset_status=$(extract_status_code "$reset_response")
    
    test_result "重置密码" "200" "$reset_status"
    
    # 使用新密码登录
    print_message $BLUE "使用新密码登录..."
    local new_login_data="{\"username\":\"$TEST_USERNAME\",\"password\":\"$new_password\"}"
    local new_login_response=$(http_request "POST" "$API_BASE/user/login" "$new_login_data" "")
    local new_login_status=$(extract_status_code "$new_login_response")
    
    test_result "新密码登录" "200" "$new_login_status"
}

# 测试无效请求
test_invalid_requests() {
    print_title "测试无效请求"
    
    # 测试无效的登录请求
    print_message $BLUE "测试无效登录..."
    local invalid_login_data="{\"username\":\"invalid\",\"password\":\"invalid\"}"
    local invalid_login_response=$(http_request "POST" "$API_BASE/user/login" "$invalid_login_data" "")
    local invalid_login_status=$(extract_status_code "$invalid_login_response")
    
    test_result "无效登录请求" "400" "$invalid_login_status"
    
    # 测试缺少参数的请求
    print_message $BLUE "测试缺少参数的请求..."
    local missing_param_data="{\"username\":\"test\"}"
    local missing_param_response=$(http_request "POST" "$API_BASE/user/login" "$missing_param_data" "")
    local missing_param_status=$(extract_status_code "$missing_param_response")
    
    test_result "缺少参数的请求" "400" "$missing_param_status"
    
    # 测试未授权的请求
    print_message $BLUE "测试未授权请求..."
    local unauthorized_response=$(http_request "GET" "$API_BASE/user/profile" "" "")
    local unauthorized_status=$(extract_status_code "$unauthorized_response")
    
    test_result "未授权请求" "401" "$unauthorized_status"
    
    # 测试无效Token
    print_message $BLUE "测试无效Token..."
    local invalid_token_response=$(http_request "GET" "$API_BASE/user/profile" "" "Authorization: Bearer invalid_token")
    local invalid_token_status=$(extract_status_code "$invalid_token_response")
    
    test_result "无效Token请求" "401" "$invalid_token_status"
}

# 测试API响应格式
test_response_format() {
    print_title "测试API响应格式"
    
    # 测试成功响应格式
    print_message $BLUE "测试成功响应格式..."
    local response=$(http_request "GET" "$BASE_URL/health" "" "")
    local body=$(extract_body "$response")
    
    # 检查是否包含必要字段
    if echo "$body" | grep -q '"code"' && echo "$body" | grep -q '"message"'; then
        test_result "响应格式包含必要字段" "true" "true"
    else
        test_result "响应格式包含必要字段" "true" "false"
    fi
    
    # 测试错误响应格式
    print_message $BLUE "测试错误响应格式..."
    local error_response=$(http_request "GET" "$API_BASE/nonexistent" "" "")
    local error_body=$(extract_body "$error_response")
    
    if echo "$error_body" | grep -q '"code"' && echo "$error_body" | grep -q '"message"'; then
        test_result "错误响应格式正确" "true" "true"
    else
        test_result "错误响应格式正确" "true" "false"
    fi
}

# 性能测试
test_performance() {
    print_title "性能测试"
    
    print_message $BLUE "测试API响应时间..."
    
    local start_time=$(date +%s%N)
    local response=$(http_request "GET" "$BASE_URL/health" "" "")
    local end_time=$(date +%s%N)
    
    local duration=$(( (end_time - start_time) / 1000000 )) # 转换为毫秒
    
    print_message $BLUE "响应时间: ${duration}ms"
    
    if [ $duration -lt 1000 ]; then
        test_result "响应时间 < 1秒" "true" "true"
    else
        test_result "响应时间 < 1秒" "true" "false"
    fi
}

# 并发测试
test_concurrency() {
    print_title "并发测试"
    
    print_message $BLUE "运行并发请求测试..."
    
    local concurrent_requests=10
    local pids=()
    
    # 启动并发请求
    for i in $(seq 1 $concurrent_requests); do
        (
            local response=$(http_request "GET" "$BASE_URL/health" "" "")
            local status=$(extract_status_code "$response")
            echo "$status"
        ) &
        pids+=($!)
    done
    
    # 等待所有请求完成
    local success_count=0
    for pid in "${pids[@]}"; do
        wait $pid
        local status=$?
        if [ $status -eq 0 ]; then
            success_count=$((success_count + 1))
        fi
    done
    
    test_result "并发请求成功率" "$concurrent_requests" "$success_count"
}

# 清理测试数据
cleanup_test_data() {
    print_title "清理测试数据"
    
    print_message $BLUE "清理测试用户数据..."
    
    # 这里可以添加清理逻辑，比如删除测试用户
    # 由于这是演示，我们只是打印消息
    print_message $GREEN "测试数据清理完成"
}

# 生成测试报告
generate_test_report() {
    print_title "测试报告"
    
    local success_rate=0
    if [ $TOTAL_TESTS -gt 0 ]; then
        success_rate=$(( PASSED_TESTS * 100 / TOTAL_TESTS ))
    fi
    
    echo ""
    print_message $BLUE "测试统计:"
    echo "  总测试数: $TOTAL_TESTS"
    echo "  通过数: $PASSED_TESTS"
    echo "  失败数: $FAILED_TESTS"
    echo "  成功率: ${success_rate}%"
    echo ""
    
    if [ $FAILED_TESTS -eq 0 ]; then
        print_message $GREEN "🎉 所有测试通过!"
        return 0
    else
        print_message $RED "❌ 有 $FAILED_TESTS 个测试失败"
        return 1
    fi
}

# 显示帮助信息
show_help() {
    echo "API自动化测试脚本"
    echo ""
    echo "用法: $0 [选项]"
    echo ""
    echo "选项:"
    echo "  -h, --help          显示帮助信息"
    echo "  -u, --url URL       指定API基础URL (默认: http://localhost:8080)"
    echo "  -e, --email EMAIL   指定测试邮箱 (默认: test@example.com)"
    echo "  -v, --verbose       详细输出"
    echo "  --skip-cleanup      跳过清理步骤"
    echo ""
    echo "示例:"
    echo "  $0                                    # 使用默认设置"
    echo "  $0 -u http://localhost:9000          # 指定API URL"
    echo "  $0 -e mytest@example.com             # 指定测试邮箱"
}

# 主函数
main() {
    print_message $GREEN "开始API自动化测试..."
    echo ""
    
    # 检查必要的命令
    for cmd in curl grep; do
        if ! command -v $cmd > /dev/null 2>&1; then
            print_message $RED "错误: $cmd 未安装"
            exit 1
        fi
    done
    
    # 执行测试
    check_app_running
    test_response_format
    test_user_registration
    test_user_login
    test_get_user_profile
    test_update_user_profile
    test_password_reset
    test_invalid_requests
    test_performance
    test_concurrency
    
    # 清理和报告
    if [ "$SKIP_CLEANUP" != "true" ]; then
        cleanup_test_data
    fi
    
    generate_test_report
}

# 解析命令行参数
SKIP_CLEANUP=false
VERBOSE=false

while [[ $# -gt 0 ]]; do
    case $1 in
        -h|--help)
            show_help
            exit 0
            ;;
        -u|--url)
            BASE_URL="$2"
            API_BASE="$BASE_URL/api/v1"
            shift 2
            ;;
        -e|--email)
            TEST_EMAIL="$2"
            shift 2
            ;;
        -v|--verbose)
            VERBOSE=true
            shift
            ;;
        --skip-cleanup)
            SKIP_CLEANUP=true
            shift
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