#!/bin/bash
echo "🚀 在VSCode终端中启动三端开发环境"
echo ""
echo "请按以下步骤操作:"
echo "1. 使用 Cmd+Shift+\` 创建新终端 (重复3次)"
echo "2. 在各终端中分别运行:"
echo "   📱 管理端: ./scripts/start-admin.sh"
echo "   🌐 用户端: ./scripts/start-web.sh"
echo "   ⚙️  后端:   ./scripts/start-backend.sh"
echo ""
echo "或者使用以下命令在当前终端并行启动:"
echo "   make fullstack-dev"
echo ""
read -p "按回车键继续，或者 Ctrl+C 取消..." dummy
echo ""
echo "🎯 现在可以手动执行上述步骤，或运行其他启动方式"

