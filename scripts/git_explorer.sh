#!/bin/bash

# Git 提交浏览器脚本
# 帮助用户浏览历史提交并切换到指定版本

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# 检查是否在Git仓库中
if ! git rev-parse --git-dir > /dev/null 2>&1; then
    echo -e "${RED}错误: 当前目录不是Git仓库${NC}"
    exit 1
fi

# 获取当前分支
current_branch=$(git branch --show-current)
original_branch=$current_branch

echo -e "${BLUE}=== Git 提交浏览器 ===${NC}"
echo -e "${CYAN}当前分支: ${original_branch}${NC}"
echo ""

# 显示帮助信息
show_help() {
    echo -e "${YELLOW}使用说明:${NC}"
    echo "  数字 - 切换到对应的提交"
    echo "  b    - 基于当前提交创建新分支"
    echo "  l    - 显示更多提交 (默认显示20个)"
    echo "  s    - 显示当前提交的详细信息"
    echo "  d    - 显示当前提交的文件变更"
    echo "  f    - 查看指定文件在当前提交的内容"
    echo "  r    - 返回原始分支"
    echo "  q    - 退出"
    echo "  h    - 显示帮助"
    echo ""
}

# 显示提交列表
show_commits() {
    local count=${1:-20}
    echo -e "${BLUE}最近 ${count} 次提交:${NC}"
    echo ""
    
    # 使用格式化输出显示提交
    git log --oneline --graph --decorate -${count} --pretty=format:"%C(yellow)%h%C(reset) %C(blue)%ad%C(reset) %C(green)%an%C(reset) %s %C(red)%d%C(reset)" --date=short | nl -v0 -w3 -s'. '
    echo ""
    echo ""
}

# 显示当前状态
show_current_status() {
    local current_commit=$(git rev-parse --short HEAD)
    local current_date=$(git show -s --format=%ad --date=short HEAD)
    local current_author=$(git show -s --format=%an HEAD)
    local current_message=$(git show -s --format=%s HEAD)
    
    echo -e "${GREEN}当前位置:${NC}"
    echo -e "  提交: ${YELLOW}${current_commit}${NC}"
    echo -e "  日期: ${BLUE}${current_date}${NC}"
    echo -e "  作者: ${PURPLE}${current_author}${NC}"
    echo -e "  信息: ${current_message}"
    echo ""
}

# 切换到指定提交
checkout_commit() {
    local index=$1
    local commit_hash=$(git log --oneline -$((index + 1)) | sed -n "$((index + 1))p" | cut -d' ' -f1)
    
    if [ -z "$commit_hash" ]; then
        echo -e "${RED}错误: 无效的提交索引${NC}"
        return 1
    fi
    
    echo -e "${YELLOW}切换到提交: ${commit_hash}${NC}"
    if git checkout $commit_hash 2>/dev/null; then
        echo -e "${GREEN}✓ 切换成功${NC}"
        show_current_status
    else
        echo -e "${RED}✗ 切换失败${NC}"
    fi
}

# 创建新分支
create_branch() {
    echo -n "请输入新分支名称: "
    read branch_name
    
    if [ -z "$branch_name" ]; then
        echo -e "${RED}分支名称不能为空${NC}"
        return 1
    fi
    
    if git checkout -b "$branch_name" 2>/dev/null; then
        echo -e "${GREEN}✓ 创建并切换到新分支: ${branch_name}${NC}"
    else
        echo -e "${RED}✗ 创建分支失败 (可能已存在)${NC}"
    fi
}

# 显示提交详情
show_commit_details() {
    echo -e "${BLUE}=== 当前提交详细信息 ===${NC}"
    git show --stat --pretty=format:"%C(yellow)提交: %H%C(reset)%n%C(blue)日期: %ad%C(reset)%n%C(green)作者: %an <%ae>%C(reset)%n%C(purple)信息: %s%C(reset)%n" --date=full
    echo ""
}

# 显示文件变更
show_file_changes() {
    echo -e "${BLUE}=== 当前提交的文件变更 ===${NC}"
    git show --name-status --pretty=format:""
    echo ""
    echo -e "${YELLOW}变更统计:${NC}"
    git show --stat --pretty=format:""
    echo ""
}

# 查看文件内容
view_file() {
    echo -n "请输入文件路径: "
    read file_path
    
    if [ -z "$file_path" ]; then
        echo -e "${RED}文件路径不能为空${NC}"
        return 1
    fi
    
    if git show HEAD:$file_path 2>/dev/null; then
        echo ""
    else
        echo -e "${RED}错误: 文件不存在或无法访问${NC}"
    fi
}

# 返回原始分支
return_to_original() {
    if [ "$original_branch" != "HEAD" ] && [ -n "$original_branch" ]; then
        echo -e "${YELLOW}返回到原始分支: ${original_branch}${NC}"
        if git checkout "$original_branch" 2>/dev/null; then
            echo -e "${GREEN}✓ 返回成功${NC}"
        else
            echo -e "${RED}✗ 返回失败${NC}"
        fi
    else
        echo -e "${YELLOW}无法确定原始分支，请手动切换${NC}"
    fi
}

# 主循环
main() {
    local commit_count=20
    
    show_help
    show_commits $commit_count
    show_current_status
    
    while true; do
        echo -e "${CYAN}请选择操作 (输入数字选择提交，或输入命令字母):${NC}"
        read -p "> " choice
        
        case $choice in
            [0-9]*)
                if [[ $choice =~ ^[0-9]+$ ]]; then
                    checkout_commit $choice
                else
                    echo -e "${RED}请输入有效的数字${NC}"
                fi
                ;;
            b|B)
                create_branch
                ;;
            l|L)
                echo -n "显示多少个提交? (默认20): "
                read new_count
                commit_count=${new_count:-20}
                show_commits $commit_count
                ;;
            s|S)
                show_commit_details
                ;;
            d|D)
                show_file_changes
                ;;
            f|F)
                view_file
                ;;
            r|R)
                return_to_original
                show_current_status
                ;;
            h|H)
                show_help
                ;;
            q|Q)
                echo -e "${YELLOW}退出Git浏览器${NC}"
                if [ "$(git branch --show-current)" != "$original_branch" ] && [ -n "$original_branch" ]; then
                    echo -e "${CYAN}提示: 您当前不在原始分支上，需要的话请使用 'git checkout $original_branch' 返回${NC}"
                fi
                break
                ;;
            *)
                echo -e "${RED}无效选择，请重试 (输入 h 查看帮助)${NC}"
                ;;
        esac
        echo ""
    done
}

# 运行主程序
main 