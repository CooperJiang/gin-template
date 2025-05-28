#!/bin/bash

echo "🚀 启动 Gin Template 前端项目..."
echo ""

# 检查 Node.js 是否安装
if ! command -v node &> /dev/null; then
    echo "❌ Node.js 未安装，请先安装 Node.js"
    exit 1
fi

# 检查 npm 是否安装
if ! command -v npm &> /dev/null; then
    echo "❌ npm 未安装，请先安装 npm"
    exit 1
fi

# 检查是否已安装依赖
if [ ! -d "node_modules" ]; then
    echo "📦 安装依赖..."
    npm install
fi

echo "🎨 启动前端开发服务器..."
echo "📍 访问地址: http://localhost:3000"
echo "🔄 代码变更将自动刷新"
echo ""
echo "按 Ctrl+C 停止服务器"
echo ""

npm run dev 