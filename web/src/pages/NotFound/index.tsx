import { Link } from 'react-router-dom'
import { Button } from '@/components/ui/button'

export default function NotFound() {
  return (
    <div className="min-h-[60vh] flex items-center justify-center">
      <div className="text-center">
        <h1 className="text-6xl font-bold text-gray-200">404</h1>
        <h2 className="mt-4 text-xl font-semibold text-gray-900">页面未找到</h2>
        <p className="mt-2 text-gray-500">您访问的页面不存在或已被移除</p>
        <div className="mt-6">
          <Button asChild>
            <Link to="/">返回首页</Link>
          </Button>
        </div>
      </div>
    </div>
  )
}
