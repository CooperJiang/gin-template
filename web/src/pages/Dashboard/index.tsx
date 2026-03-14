import { useAuth } from '@/auth/useAuth'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { LayoutDashboard, Users, Activity, TrendingUp } from 'lucide-react'

const stats = [
  { title: '总用户', value: '—', icon: Users, color: 'text-blue-600 bg-blue-100' },
  { title: '活跃用户', value: '—', icon: Activity, color: 'text-green-600 bg-green-100' },
  { title: '今日访问', value: '—', icon: TrendingUp, color: 'text-purple-600 bg-purple-100' },
  { title: '系统状态', value: '正常', icon: LayoutDashboard, color: 'text-orange-600 bg-orange-100' },
]

export default function Dashboard() {
  const { user } = useAuth()

  return (
    <div>
      <div className="mb-8">
        <h1 className="text-2xl font-bold text-gray-900">
          欢迎回来，{user?.username || '用户'}
        </h1>
        <p className="text-gray-500 mt-1">以下是您的系统概览</p>
      </div>

      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        {stats.map((stat) => {
          const Icon = stat.icon
          return (
            <Card key={stat.title}>
              <CardContent className="p-6">
                <div className="flex items-center justify-between">
                  <div>
                    <p className="text-sm text-gray-500">{stat.title}</p>
                    <p className="text-2xl font-bold text-gray-900 mt-1">{stat.value}</p>
                  </div>
                  <div className={`w-12 h-12 rounded-lg flex items-center justify-center ${stat.color}`}>
                    <Icon className="w-6 h-6" />
                  </div>
                </div>
              </CardContent>
            </Card>
          )
        })}
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <Card>
          <CardHeader>
            <CardTitle>快速开始</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="space-y-4 text-sm text-gray-600">
              <p>这是一个基于 Gin + React 的全栈应用脚手架。</p>
              <ul className="space-y-2 list-disc list-inside">
                <li>后端：Go + Gin + GORM</li>
                <li>前端：React + TypeScript + Tailwind CSS</li>
                <li>认证：JWT Token + 安全加密存储</li>
                <li>部署：一键构建 + SSH 远程部署</li>
              </ul>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>系统信息</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="space-y-3 text-sm">
              <div className="flex justify-between">
                <span className="text-gray-500">前端框架</span>
                <span className="text-gray-900 font-medium">React 19</span>
              </div>
              <div className="flex justify-between">
                <span className="text-gray-500">构建工具</span>
                <span className="text-gray-900 font-medium">Vite 6</span>
              </div>
              <div className="flex justify-between">
                <span className="text-gray-500">UI 框架</span>
                <span className="text-gray-900 font-medium">Tailwind CSS + shadcn/ui</span>
              </div>
              <div className="flex justify-between">
                <span className="text-gray-500">后端框架</span>
                <span className="text-gray-900 font-medium">Gin</span>
              </div>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
