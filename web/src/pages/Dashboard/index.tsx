import { useState, useEffect } from 'react'
import { useAuth } from '@/auth/useAuth'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { LayoutDashboard, Users, Activity, TrendingUp } from 'lucide-react'

interface StatItem {
  title: string
  value: string
  icon: typeof Users
  color: string
}

export default function Dashboard() {
  const { user } = useAuth()
  const [loading, setLoading] = useState(true)
  const [stats, setStats] = useState<StatItem[]>([])

  useEffect(() => {
    // TODO: 替换为真实 API 调用，例如 apiClient.get('/dashboard/stats')
    const timer = setTimeout(() => {
      setStats([
        { title: '总用户', value: '128', icon: Users, color: 'text-blue-600 bg-blue-100 dark:text-blue-400 dark:bg-blue-900/30' },
        { title: '活跃用户', value: '42', icon: Activity, color: 'text-green-600 bg-green-100 dark:text-green-400 dark:bg-green-900/30' },
        { title: '今日访问', value: '1,024', icon: TrendingUp, color: 'text-purple-600 bg-purple-100 dark:text-purple-400 dark:bg-purple-900/30' },
        { title: '系统状态', value: '正常', icon: LayoutDashboard, color: 'text-orange-600 bg-orange-100 dark:text-orange-400 dark:bg-orange-900/30' },
      ])
      setLoading(false)
    }, 600)
    return () => clearTimeout(timer)
  }, [])

  return (
    <div>
      <div className="mb-8">
        <h1 className="text-2xl font-bold text-gray-900 dark:text-gray-100">
          欢迎回来，{user?.username || '用户'}
        </h1>
        <p className="text-gray-500 dark:text-gray-400 mt-1">以下是您的系统概览</p>
      </div>

      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        {loading
          ? Array.from({ length: 4 }).map((_, i) => (
              <Card key={i}>
                <CardContent className="p-6">
                  <div className="flex items-center justify-between animate-pulse">
                    <div>
                      <div className="h-4 w-16 bg-gray-200 dark:bg-gray-700 rounded mb-3" />
                      <div className="h-7 w-12 bg-gray-200 dark:bg-gray-700 rounded" />
                    </div>
                    <div className="w-12 h-12 rounded-lg bg-gray-200 dark:bg-gray-700" />
                  </div>
                </CardContent>
              </Card>
            ))
          : stats.map((stat) => {
              const Icon = stat.icon
              return (
                <Card key={stat.title}>
                  <CardContent className="p-6">
                    <div className="flex items-center justify-between">
                      <div>
                        <p className="text-sm text-gray-500 dark:text-gray-400">{stat.title}</p>
                        <p className="text-2xl font-bold text-gray-900 dark:text-gray-100 mt-1">
                          {stat.value}
                        </p>
                      </div>
                      <div
                        className={`w-12 h-12 rounded-lg flex items-center justify-center ${stat.color}`}
                      >
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
            <div className="space-y-4 text-sm text-gray-600 dark:text-gray-400">
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
                <span className="text-gray-500 dark:text-gray-400">前端框架</span>
                <span className="text-gray-900 dark:text-gray-100 font-medium">React 19</span>
              </div>
              <div className="flex justify-between">
                <span className="text-gray-500 dark:text-gray-400">构建工具</span>
                <span className="text-gray-900 dark:text-gray-100 font-medium">Vite 6</span>
              </div>
              <div className="flex justify-between">
                <span className="text-gray-500 dark:text-gray-400">UI 框架</span>
                <span className="text-gray-900 dark:text-gray-100 font-medium">
                  Tailwind CSS + shadcn/ui
                </span>
              </div>
              <div className="flex justify-between">
                <span className="text-gray-500 dark:text-gray-400">后端框架</span>
                <span className="text-gray-900 dark:text-gray-100 font-medium">Gin</span>
              </div>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
