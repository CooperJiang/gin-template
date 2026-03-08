import { Link } from 'react-router-dom'
import { useAuth } from '@app/auth'

export default function Home() {
  const { isAuthenticated } = useAuth()

  return (
    <div className="bg-gray-50">
      <main className="max-w-7xl mx-auto py-12 px-4 sm:px-6 lg:px-8">
        <div className="text-center">
          <h1 className="text-4xl font-bold text-gray-900 mb-8">欢迎来到 __TITLE__</h1>
          <p className="text-xl text-gray-600 mb-12">基于 React + TypeScript + TailwindCSS 的子应用模板</p>

          <div className="grid grid-cols-1 md:grid-cols-3 gap-8 mt-12">
            <div className="bg-white p-6 rounded-lg shadow-md">
              <h3 className="text-lg font-semibold text-gray-900 mb-4">现代技术栈</h3>
              <p className="text-gray-600">使用 React 19、TypeScript 和 TailwindCSS 构建</p>
            </div>
            <div className="bg-white p-6 rounded-lg shadow-md">
              <h3 className="text-lg font-semibold text-gray-900 mb-4">安全认证</h3>
              <p className="text-gray-600">内置用户认证系统，支持登录、注册、密码重置等功能</p>
            </div>
            <div className="bg-white p-6 rounded-lg shadow-md">
              <h3 className="text-lg font-semibold text-gray-900 mb-4">微前端架构</h3>
              <p className="text-gray-600">基于 qiankun 的微前端架构，支持多应用独立开发部署</p>
            </div>
          </div>

          <div className="mt-12 flex items-center justify-center gap-4">
            <Link to="/about" className="inline-flex items-center px-6 py-3 border border-transparent text-base font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700">了解更多</Link>
            {isAuthenticated && (
              <Link to="/profile" className="inline-flex items-center px-6 py-3 border border-gray-300 text-base font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50">个人资料</Link>
            )}
          </div>
        </div>
      </main>
    </div>
  )
}
