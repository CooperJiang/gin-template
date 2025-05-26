#!/bin/bash

# 数据库管理脚本
# 用于数据库备份、恢复、清理等操作

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 配置
DB_FILE="app.db"
BACKUP_DIR="backups"
CONFIG_FILE="config.yaml"

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

# 检查SQLite是否安装
check_sqlite() {
    if ! command -v sqlite3 > /dev/null 2>&1; then
        print_message $RED "✗ SQLite3 未安装"
        print_message $YELLOW "请安装SQLite3: "
        print_message $YELLOW "  Ubuntu/Debian: sudo apt-get install sqlite3"
        print_message $YELLOW "  CentOS/RHEL: sudo yum install sqlite"
        print_message $YELLOW "  macOS: brew install sqlite"
        exit 1
    fi
}

# 检查数据库文件是否存在
check_db_exists() {
    if [ ! -f "$DB_FILE" ]; then
        print_message $RED "✗ 数据库文件 $DB_FILE 不存在"
        return 1
    fi
    return 0
}

# 创建备份目录
create_backup_dir() {
    if [ ! -d "$BACKUP_DIR" ]; then
        mkdir -p "$BACKUP_DIR"
        print_message $GREEN "✓ 创建备份目录: $BACKUP_DIR"
    fi
}

# 备份数据库
backup_database() {
    print_title "备份数据库"
    
    check_sqlite
    
    if ! check_db_exists; then
        return 1
    fi
    
    create_backup_dir
    
    local timestamp=$(date +%Y%m%d_%H%M%S)
    local backup_file="$BACKUP_DIR/app_backup_$timestamp.db"
    local sql_backup_file="$BACKUP_DIR/app_backup_$timestamp.sql"
    
    # 创建二进制备份
    print_message $BLUE "创建二进制备份..."
    cp "$DB_FILE" "$backup_file"
    
    # 创建SQL备份
    print_message $BLUE "创建SQL备份..."
    sqlite3 "$DB_FILE" .dump > "$sql_backup_file"
    
    # 压缩备份文件
    print_message $BLUE "压缩备份文件..."
    tar -czf "$BACKUP_DIR/app_backup_$timestamp.tar.gz" -C "$BACKUP_DIR" \
        "app_backup_$timestamp.db" "app_backup_$timestamp.sql"
    
    # 删除临时文件
    rm "$backup_file" "$sql_backup_file"
    
    local backup_size=$(du -h "$BACKUP_DIR/app_backup_$timestamp.tar.gz" | cut -f1)
    print_message $GREEN "✓ 备份完成: $BACKUP_DIR/app_backup_$timestamp.tar.gz ($backup_size)"
    
    # 显示备份信息
    print_message $BLUE "备份信息:"
    echo "  - 时间: $(date)"
    echo "  - 文件: $BACKUP_DIR/app_backup_$timestamp.tar.gz"
    echo "  - 大小: $backup_size"
    
    # 记录备份日志
    echo "$(date): 备份创建 - $BACKUP_DIR/app_backup_$timestamp.tar.gz" >> "$BACKUP_DIR/backup.log"
}

# 列出备份文件
list_backups() {
    print_title "备份文件列表"
    
    if [ ! -d "$BACKUP_DIR" ]; then
        print_message $YELLOW "⚠ 备份目录不存在"
        return
    fi
    
    local backup_files=($(ls -t "$BACKUP_DIR"/*.tar.gz 2>/dev/null || true))
    
    if [ ${#backup_files[@]} -eq 0 ]; then
        print_message $YELLOW "⚠ 没有找到备份文件"
        return
    fi
    
    print_message $GREEN "找到 ${#backup_files[@]} 个备份文件:"
    echo ""
    
    for i in "${!backup_files[@]}"; do
        local file="${backup_files[$i]}"
        local filename=$(basename "$file")
        local size=$(du -h "$file" | cut -f1)
        local date=$(stat -c %y "$file" | cut -d' ' -f1,2 | cut -d'.' -f1)
        
        printf "%2d. %-40s %8s  %s\n" $((i+1)) "$filename" "$size" "$date"
    done
}

# 恢复数据库
restore_database() {
    print_title "恢复数据库"
    
    check_sqlite
    
    # 列出可用的备份
    list_backups
    
    local backup_files=($(ls -t "$BACKUP_DIR"/*.tar.gz 2>/dev/null || true))
    
    if [ ${#backup_files[@]} -eq 0 ]; then
        print_message $RED "✗ 没有可用的备份文件"
        return 1
    fi
    
    echo ""
    read -p "请选择要恢复的备份 (1-${#backup_files[@]}): " choice
    
    if ! [[ "$choice" =~ ^[0-9]+$ ]] || [ "$choice" -lt 1 ] || [ "$choice" -gt ${#backup_files[@]} ]; then
        print_message $RED "✗ 无效的选择"
        return 1
    fi
    
    local selected_backup="${backup_files[$((choice-1))]}"
    local backup_name=$(basename "$selected_backup")
    
    print_message $YELLOW "⚠ 警告: 这将覆盖当前数据库!"
    read -p "确定要恢复备份 '$backup_name' 吗? (y/N): " confirm
    
    if [ "$confirm" != "y" ] && [ "$confirm" != "Y" ]; then
        print_message $YELLOW "操作已取消"
        return
    fi
    
    # 备份当前数据库
    if check_db_exists; then
        local current_backup="$BACKUP_DIR/current_backup_$(date +%Y%m%d_%H%M%S).db"
        cp "$DB_FILE" "$current_backup"
        print_message $GREEN "✓ 当前数据库已备份到: $current_backup"
    fi
    
    # 解压并恢复
    print_message $BLUE "解压备份文件..."
    local temp_dir=$(mktemp -d)
    tar -xzf "$selected_backup" -C "$temp_dir"
    
    # 查找数据库文件
    local db_backup=$(find "$temp_dir" -name "*.db" | head -1)
    if [ -n "$db_backup" ]; then
        cp "$db_backup" "$DB_FILE"
        print_message $GREEN "✓ 数据库恢复完成"
    else
        print_message $RED "✗ 备份文件中未找到数据库文件"
        rm -rf "$temp_dir"
        return 1
    fi
    
    # 清理临时文件
    rm -rf "$temp_dir"
    
    # 记录恢复日志
    echo "$(date): 数据库恢复 - 从 $backup_name" >> "$BACKUP_DIR/backup.log"
    
    print_message $GREEN "✓ 数据库恢复完成"
}

# 清理旧备份
cleanup_backups() {
    print_title "清理旧备份"
    
    if [ ! -d "$BACKUP_DIR" ]; then
        print_message $YELLOW "⚠ 备份目录不存在"
        return
    fi
    
    local days=${1:-30}  # 默认保留30天
    
    print_message $BLUE "清理 $days 天前的备份文件..."
    
    local deleted_count=0
    while IFS= read -r -d '' file; do
        rm "$file"
        print_message $YELLOW "删除: $(basename "$file")"
        ((deleted_count++))
    done < <(find "$BACKUP_DIR" -name "*.tar.gz" -mtime +$days -print0 2>/dev/null)
    
    if [ $deleted_count -eq 0 ]; then
        print_message $GREEN "✓ 没有需要清理的备份文件"
    else
        print_message $GREEN "✓ 已删除 $deleted_count 个旧备份文件"
        echo "$(date): 清理备份 - 删除了 $deleted_count 个文件" >> "$BACKUP_DIR/backup.log"
    fi
}

# 数据库信息
show_db_info() {
    print_title "数据库信息"
    
    check_sqlite
    
    if ! check_db_exists; then
        return 1
    fi
    
    # 基本信息
    local db_size=$(du -h "$DB_FILE" | cut -f1)
    local db_modified=$(stat -c %y "$DB_FILE" | cut -d'.' -f1)
    
    print_message $BLUE "基本信息:"
    echo "  - 文件: $DB_FILE"
    echo "  - 大小: $db_size"
    echo "  - 修改时间: $db_modified"
    echo ""
    
    # 表信息
    print_message $BLUE "表信息:"
    sqlite3 "$DB_FILE" "SELECT name FROM sqlite_master WHERE type='table';" | while read table; do
        if [ -n "$table" ]; then
            local count=$(sqlite3 "$DB_FILE" "SELECT COUNT(*) FROM $table;")
            printf "  - %-20s %s 条记录\n" "$table:" "$count"
        fi
    done
    echo ""
    
    # 索引信息
    print_message $BLUE "索引信息:"
    local index_count=$(sqlite3 "$DB_FILE" "SELECT COUNT(*) FROM sqlite_master WHERE type='index' AND name NOT LIKE 'sqlite_%';")
    echo "  - 用户索引数量: $index_count"
    
    if [ $index_count -gt 0 ]; then
        sqlite3 "$DB_FILE" "SELECT name, tbl_name FROM sqlite_master WHERE type='index' AND name NOT LIKE 'sqlite_%';" | while IFS='|' read index_name table_name; do
            printf "    - %-20s (表: %s)\n" "$index_name" "$table_name"
        done
    fi
    echo ""
    
    # 数据库完整性检查
    print_message $BLUE "完整性检查:"
    local integrity_result=$(sqlite3 "$DB_FILE" "PRAGMA integrity_check;")
    if [ "$integrity_result" = "ok" ]; then
        print_message $GREEN "  ✓ 数据库完整性正常"
    else
        print_message $RED "  ✗ 数据库完整性检查失败:"
        echo "    $integrity_result"
    fi
}

# 优化数据库
optimize_database() {
    print_title "优化数据库"
    
    check_sqlite
    
    if ! check_db_exists; then
        return 1
    fi
    
    # 备份数据库
    print_message $BLUE "创建优化前备份..."
    local backup_file="$BACKUP_DIR/pre_optimize_$(date +%Y%m%d_%H%M%S).db"
    create_backup_dir
    cp "$DB_FILE" "$backup_file"
    
    local original_size=$(du -h "$DB_FILE" | cut -f1)
    print_message $BLUE "原始大小: $original_size"
    
    # 执行VACUUM
    print_message $BLUE "执行 VACUUM 操作..."
    sqlite3 "$DB_FILE" "VACUUM;"
    
    # 重建索引
    print_message $BLUE "重建索引..."
    sqlite3 "$DB_FILE" "REINDEX;"
    
    # 分析统计信息
    print_message $BLUE "更新统计信息..."
    sqlite3 "$DB_FILE" "ANALYZE;"
    
    local optimized_size=$(du -h "$DB_FILE" | cut -f1)
    print_message $GREEN "✓ 优化完成"
    print_message $BLUE "优化后大小: $optimized_size"
    
    # 记录优化日志
    echo "$(date): 数据库优化 - 从 $original_size 到 $optimized_size" >> "$BACKUP_DIR/backup.log"
}

# 执行SQL查询
execute_sql() {
    print_title "执行SQL查询"
    
    check_sqlite
    
    if ! check_db_exists; then
        return 1
    fi
    
    local sql_file="$1"
    
    if [ -n "$sql_file" ] && [ -f "$sql_file" ]; then
        print_message $BLUE "执行SQL文件: $sql_file"
        sqlite3 "$DB_FILE" < "$sql_file"
        print_message $GREEN "✓ SQL文件执行完成"
    else
        print_message $BLUE "进入交互式SQL模式 (输入 .quit 退出):"
        sqlite3 "$DB_FILE"
    fi
}

# 导出数据
export_data() {
    print_title "导出数据"
    
    check_sqlite
    
    if ! check_db_exists; then
        return 1
    fi
    
    local export_dir="exports"
    local timestamp=$(date +%Y%m%d_%H%M%S)
    
    mkdir -p "$export_dir"
    
    # 导出为SQL
    local sql_file="$export_dir/export_$timestamp.sql"
    sqlite3 "$DB_FILE" .dump > "$sql_file"
    print_message $GREEN "✓ SQL导出: $sql_file"
    
    # 导出为CSV
    print_message $BLUE "导出表为CSV格式..."
    sqlite3 "$DB_FILE" "SELECT name FROM sqlite_master WHERE type='table';" | while read table; do
        if [ -n "$table" ]; then
            local csv_file="$export_dir/${table}_$timestamp.csv"
            sqlite3 -header -csv "$DB_FILE" "SELECT * FROM $table;" > "$csv_file"
            print_message $GREEN "✓ CSV导出: $csv_file"
        fi
    done
    
    # 创建压缩包
    local archive_file="$export_dir/export_$timestamp.tar.gz"
    tar -czf "$archive_file" -C "$export_dir" \
        "export_$timestamp.sql" *_$timestamp.csv
    
    # 清理临时文件
    rm "$export_dir/export_$timestamp.sql" "$export_dir"/*_$timestamp.csv
    
    local archive_size=$(du -h "$archive_file" | cut -f1)
    print_message $GREEN "✓ 数据导出完成: $archive_file ($archive_size)"
}

# 重置数据库
reset_database() {
    print_title "重置数据库"
    
    print_message $RED "⚠ 警告: 这将删除所有数据!"
    read -p "确定要重置数据库吗? (y/N): " confirm
    
    if [ "$confirm" != "y" ] && [ "$confirm" != "Y" ]; then
        print_message $YELLOW "操作已取消"
        return
    fi
    
    # 备份当前数据库
    if check_db_exists; then
        create_backup_dir
        local backup_file="$BACKUP_DIR/pre_reset_$(date +%Y%m%d_%H%M%S).db"
        cp "$DB_FILE" "$backup_file"
        print_message $GREEN "✓ 当前数据库已备份到: $backup_file"
    fi
    
    # 删除数据库文件
    rm -f "$DB_FILE"
    print_message $GREEN "✓ 数据库已重置"
    
    # 记录重置日志
    create_backup_dir
    echo "$(date): 数据库重置" >> "$BACKUP_DIR/backup.log"
}

# 显示帮助信息
show_help() {
    echo "数据库管理脚本"
    echo ""
    echo "用法: $0 <命令> [选项]"
    echo ""
    echo "命令:"
    echo "  backup              备份数据库"
    echo "  restore             恢复数据库"
    echo "  list                列出备份文件"
    echo "  cleanup [days]      清理旧备份 (默认30天)"
    echo "  info                显示数据库信息"
    echo "  optimize            优化数据库"
    echo "  sql [file]          执行SQL查询或文件"
    echo "  export              导出数据"
    echo "  reset               重置数据库"
    echo ""
    echo "选项:"
    echo "  -h, --help          显示帮助信息"
    echo "  -d, --db FILE       指定数据库文件 (默认: app.db)"
    echo "  -b, --backup-dir DIR 指定备份目录 (默认: backups)"
    echo ""
    echo "示例:"
    echo "  $0 backup                    # 备份数据库"
    echo "  $0 restore                   # 恢复数据库"
    echo "  $0 cleanup 7                 # 清理7天前的备份"
    echo "  $0 sql queries.sql           # 执行SQL文件"
    echo "  $0 -d custom.db info         # 查看自定义数据库信息"
}

# 主函数
main() {
    local command="$1"
    shift || true
    
    case "$command" in
        backup)
            backup_database
            ;;
        restore)
            restore_database
            ;;
        list)
            list_backups
            ;;
        cleanup)
            cleanup_backups "$1"
            ;;
        info)
            show_db_info
            ;;
        optimize)
            optimize_database
            ;;
        sql)
            execute_sql "$1"
            ;;
        export)
            export_data
            ;;
        reset)
            reset_database
            ;;
        -h|--help|help)
            show_help
            ;;
        "")
            print_message $RED "错误: 请指定命令"
            echo ""
            show_help
            exit 1
            ;;
        *)
            print_message $RED "错误: 未知命令 '$command'"
            echo ""
            show_help
            exit 1
            ;;
    esac
}

# 解析全局选项
while [[ $# -gt 0 ]]; do
    case $1 in
        -d|--db)
            DB_FILE="$2"
            shift 2
            ;;
        -b|--backup-dir)
            BACKUP_DIR="$2"
            shift 2
            ;;
        -h|--help)
            show_help
            exit 0
            ;;
        *)
            break
            ;;
    esac
done

# 运行主函数
main "$@" 