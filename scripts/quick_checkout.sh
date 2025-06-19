#!/bin/bash

# 快速Git提交切换脚本
# 使用方法: ./scripts/quick_checkout.sh [搜索关键词]

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m'

# 显示提交列表
show_commits() {
    local search_term="$1"
    echo -e "${BLUE}=== 提交历史 ===${NC}"
    
    if [ -n "$search_term" ]; then
        echo -e "${CYAN}搜索: '$search_term'${NC}"
        git log --oneline --grep="$search_term" -i --pretty=format:"%C(yellow)%h%C(reset) %C(blue)%ad%C(reset) %s" --date=short -20 | nl -v1 -w3 -s'. '
    else
        git log --oneline --pretty=format:"%C(yellow)%h%C(reset) %C(blue)%ad%C(reset) %s" --date=short -30 | nl -v1 -w3 -s'. '
    fi
    echo ""
}

# 快速切换
quick_switch() {
    local search_term="$1"
    
    if [ -n "$search_term" ]; then
        # 搜索包含关键词的提交
        local commits=($(git log --oneline --grep="$search_term" -i --pretty=format:"%h" -10))
        
        if [ ${#commits[@]} -eq 0 ]; then
            echo -e "${RED}没有找到包含 '$search_term' 的提交${NC}"
            return 1
        elif [ ${#commits[@]} -eq 1 ]; then
            # 只有一个匹配的提交，直接切换
            echo -e "${YELLOW}找到匹配的提交，正在切换...${NC}"
            git checkout ${commits[0]}
            return 0
        else
            # 多个匹配的提交，显示列表
            echo -e "${YELLOW}找到 ${#commits[@]} 个匹配的提交:${NC}"
            show_commits "$search_term"
        fi
    else
        # 没有搜索词，显示最近的提交
        show_commits
    fi
    
    echo ""
    echo -e "${CYAN}请输入要切换到的提交编号:${NC}"
    read -p "> " choice
    
    if [[ $choice =~ ^[0-9]+$ ]]; then
        local commit_hash
        if [ -n "$search_term" ]; then
            commit_hash=$(git log --oneline --grep="$search_term" -i --pretty=format:"%h" -20 | sed -n "${choice}p")
        else
            commit_hash=$(git log --oneline --pretty=format:"%h" -30 | sed -n "${choice}p")
        fi
        
        if [ -n "$commit_hash" ]; then
            echo -e "${YELLOW}切换到提交: $commit_hash${NC}"
            git checkout $commit_hash
            
            # 显示当前提交信息
            echo -e "${GREEN}当前提交:${NC}"
            git show --stat --oneline | head -1
            echo ""
            
            # 询问是否创建分支
            echo -e "${CYAN}是否基于此提交创建新分支? (y/n):${NC}"
            read -p "> " create_branch
            
            if [[ $create_branch =~ ^[Yy]$ ]]; then
                echo -n "请输入分支名称: "
                read branch_name
                if [ -n "$branch_name" ]; then
                    git checkout -b "$branch_name"
                    echo -e "${GREEN}✓ 创建分支: $branch_name${NC}"
                fi
            fi
        else
            echo -e "${RED}无效的选择${NC}"
        fi
    else
        echo -e "${RED}请输入有效的数字${NC}"
    fi
}

# 主函数
main() {
    echo -e "${BLUE}=== Git 快速切换工具 ===${NC}"
    echo -e "${CYAN}当前分支: $(git branch --show-current)${NC}"
    echo ""
    
    # 如果有参数，作为搜索词
    local search_term="$1"
    
    if [ -n "$search_term" ]; then
        echo -e "${YELLOW}搜索包含 '$search_term' 的提交...${NC}"
    fi
    
    quick_switch "$search_term"
}

# 运行主程序
main "$@" 