import { Link } from 'react-router-dom'

export default function About() {
  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-4xl mx-auto py-12 px-4 sm:px-6 lg:px-8">
        <div className="text-center mb-12">
          <h1 className="text-4xl font-bold text-gray-900 mb-4">关于 __TITLE__</h1>
          <p className="text-xl text-gray-600">了解更多关于这个子应用模板的信息</p>
        </div>

        <div className="bg-white rounded-lg shadow-md p-8 mb-8">
          <h2 className="text-2xl font-bold text-gray-900 mb-6">项目介绍</h2>
          <p className="text-gray-600 mb-4">这是一个基于现代前端技术栈构建的子应用模板，旨在为开发者提供一个快速启动项目的基础框架。</p>
          <p className="text-gray-600 mb-4">该模板采用了最新的技术和最佳实践，确保代码的可维护性、可扩展性和性能。</p>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-8 mb-8">
          <div className="bg-white rounded-lg shadow-md p-6">
            <h3 className="text-xl font-bold text-gray-900 mb-4">技术栈</h3>
            <ul className="space-y-2 text-gray-600">
              <li>• React 19 - 声明式 UI 框架</li>
              <li>• TypeScript - 类型安全的 JavaScript</li>
              <li>• TailwindCSS - 实用优先的 CSS 框架</li>
              <li>• React Router - 路由管理</li>
              <li>• qiankun - 微前端框架</li>
              <li>• Vite - 快速的构建工具</li>
            </ul>
          </div>
          <div className="bg-white rounded-lg shadow-md p-6">
            <h3 className="text-xl font-bold text-gray-900 mb-4">特性</h3>
            <ul className="space-y-2 text-gray-600">
              <li>• 响应式设计</li>
              <li>• 用户认证系统</li>
              <li>• 安全的数据存储</li>
              <li>• 微前端架构</li>
              <li>• 框架无关的核心层</li>
              <li>• 开发工具集成</li>
            </ul>
          </div>
        </div>

        <div className="text-center">
          <Link to="/" className="inline-flex items-center px-6 py-3 border border-transparent text-base font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700">返回首页</Link>
        </div>
      </div>
    </div>
  )
}
