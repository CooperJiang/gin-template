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
        { title: '总用户', value: '128', icon: Users, color: 'bg-nb-blue' },
        { title: '活跃用户', value: '42', icon: Activity, color: 'bg-nb-green' },
        { title: '今日访问', value: '1,024', icon: TrendingUp, color: 'bg-nb-purple' },
        { title: '系统状态', value: '正常', icon: LayoutDashboard, color: 'bg-nb-orange' },
      ])
      setLoading(false)
    }, 600)
    return () => clearTimeout(timer)
  }, [])

  return (
    <div>
      <div className="mb-8">
        <h1 className="text-2xl font-bold text-nb-text">
          欢迎回来，{user?.username || '用户'}
        </h1>
        <p className="text-nb-text-secondary mt-1">以下是您的系统概览</p>
      </div>

      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        {loading
          ? Array.from({ length: 4 }).map((_, i) => (
              <Card key={i}>
                <CardContent className="p-6">
                  <div className="flex items-center justify-between animate-pulse">
                    <div>
                      <div className="h-4 w-16 bg-nb-bg-soft border-nb border-nb-border rounded-[var(--nb-radius-sm)] mb-3" />
                      <div className="h-7 w-12 bg-nb-bg-soft border-nb border-nb-border rounded-[var(--nb-radius-sm)]" />
                    </div>
                    <div className="w-12 h-12 bg-nb-bg-soft border-nb border-nb-border rounded-[var(--nb-radius)]" />
                  </div>
                </CardContent>
              </Card>
            ))
          : stats.map((stat) => {
              const Icon = stat.icon
              return (
                <Card key={stat.title} className="nb-interactive">
                  <CardContent className="p-6">
                    <div className="flex items-center justify-between">
                      <div>
                        <p className="text-sm font-medium text-nb-text-secondary">{stat.title}</p>
                        <p className="text-2xl font-bold text-nb-text mt-1">
                          {stat.value}
                        </p>
                      </div>
                      <div
                        className={`w-12 h-12 rounded-[var(--nb-radius)] flex items-center justify-center border-nb border-nb-border shadow-nb-sm ${stat.color}`}
                      >
                        <Icon className="w-6 h-6 text-[var(--nb-primary-text)]" />
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
            <div className="space-y-4 text-sm text-nb-text-secondary">
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
              <div className="flex justify-between p-2 rounded-[var(--nb-radius-sm)] hover:bg-nb-bg-soft transition-colors">
                <span className="text-nb-text-secondary">前端框架</span>
                <span className="text-nb-text font-bold">React 19</span>
              </div>
              <div className="flex justify-between p-2 rounded-[var(--nb-radius-sm)] hover:bg-nb-bg-soft transition-colors">
                <span className="text-nb-text-secondary">构建工具</span>
                <span className="text-nb-text font-bold">Vite 6</span>
              </div>
              <div className="flex justify-between p-2 rounded-[var(--nb-radius-sm)] hover:bg-nb-bg-soft transition-colors">
                <span className="text-nb-text-secondary">UI 风格</span>
                <span className="text-nb-text font-bold">
                  Neubrutalism
                </span>
              </div>
              <div className="flex justify-between p-2 rounded-[var(--nb-radius-sm)] hover:bg-nb-bg-soft transition-colors">
                <span className="text-nb-text-secondary">后端框架</span>
                <span className="text-nb-text font-bold">Gin</span>
              </div>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
