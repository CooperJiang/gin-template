import { useState, useEffect } from 'react'
import { Card, CardSection, Loading, Button } from '@app/core'
import { createHttpClient } from '@app/core'

interface DashboardStats {
  totalUsers: number
  activeUsers: number
  totalRevenue: number
  growthRate: number
}

interface RecentActivity {
  id: number
  type: 'login' | 'signup' | 'purchase'
  user: string
  timestamp: string
}

/**
 * 仪表板页面示例
 * 展示：卡片布局、数据展示、状态管理
 */
export function Dashboard() {
  const [stats, setStats] = useState<DashboardStats | null>(null)
  const [activities, setActivities] = useState<RecentActivity[]>([])
  const [loading, setLoading] = useState(true)

  const http = createHttpClient({ baseURL: '/api/v1' })

  useEffect(() => {
    fetchDashboardData()
  }, [])

  const fetchDashboardData = async () => {
    setLoading(true)
    try {
      // 并发请求
      const [statsData, activitiesData] = await Promise.all([
        http.get<DashboardStats>('/dashboard/stats'),
        http.get<RecentActivity[]>('/dashboard/activities'),
      ])
      setStats(statsData)
      setActivities(activitiesData)
    } catch (error) {
      // 错误已处理
    } finally {
      setLoading(false)
    }
  }

  if (loading) {
    return (
      <div className="flex items-center justify-center h-96">
        <Loading text="加载仪表板..." />
      </div>
    )
  }

  return (
    <div className="max-w-7xl mx-auto p-6 space-y-6">
      <div className="flex items-center justify-between">
        <h1 className="text-2xl font-bold text-app-heading">仪表板</h1>
        <Button onClick={fetchDashboardData}>刷新</Button>
      </div>

      {/* 统计卡片 */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
        <StatCard
          title="总用户数"
          value={stats?.totalUsers || 0}
          icon="👥"
          trend="+12%"
        />
        <StatCard
          title="活跃用户"
          value={stats?.activeUsers || 0}
          icon="✨"
          trend="+8%"
        />
        <StatCard
          title="总收入"
          value={`¥${stats?.totalRevenue || 0}`}
          icon="💰"
          trend="+23%"
        />
        <StatCard
          title="增长率"
          value={`${stats?.growthRate || 0}%`}
          icon="📈"
          trend="+5%"
        />
      </div>

      {/* 最近活动 */}
      <Card title="最近活动">
        <div className="space-y-3">
          {activities.map((activity) => (
            <div
              key={activity.id}
              className="flex items-center justify-between p-3 rounded-lg hover:bg-app-bg-soft transition-colors"
            >
              <div className="flex items-center gap-3">
                <div className="w-8 h-8 rounded-full bg-primary-100 flex items-center justify-center">
                  {getActivityIcon(activity.type)}
                </div>
                <div>
                  <p className="text-sm font-medium text-app-text">
                    {activity.user}
                  </p>
                  <p className="text-xs text-app-text-muted">
                    {getActivityText(activity.type)}
                  </p>
                </div>
              </div>
              <div className="text-xs text-app-text-muted">
                {formatTimestamp(activity.timestamp)}
              </div>
            </div>
          ))}
        </div>
      </Card>

      {/* 快速操作 */}
      <Card title="快速操作">
        <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
          <QuickActionButton icon="➕" label="添加用户" />
          <QuickActionButton icon="📊" label="查看报表" />
          <QuickActionButton icon="⚙️" label="系统设置" />
          <QuickActionButton icon="📧" label="发送通知" />
        </div>
      </Card>
    </div>
  )
}

// 辅助组件
interface StatCardProps {
  title: string
  value: string | number
  icon: string
  trend?: string
}

function StatCard({ title, value, icon, trend }: StatCardProps) {
  return (
    <Card hoverable>
      <div className="flex items-center justify-between">
        <div>
          <p className="text-sm text-app-text-muted">{title}</p>
          <p className="text-2xl font-bold text-app-heading mt-1">{value}</p>
          {trend && (
            <p className="text-xs text-success mt-1">
              {trend} vs 上月
            </p>
          )}
        </div>
        <div className="text-4xl">{icon}</div>
      </div>
    </Card>
  )
}

interface QuickActionButtonProps {
  icon: string
  label: string
}

function QuickActionButton({ icon, label }: QuickActionButtonProps) {
  return (
    <button className="flex flex-col items-center gap-2 p-4 rounded-lg border border-app-border hover:bg-app-bg-soft hover:border-primary-500 transition-colors">
      <span className="text-2xl">{icon}</span>
      <span className="text-sm text-app-text">{label}</span>
    </button>
  )
}

// 辅助函数
function getActivityIcon(type: string) {
  const icons = {
    login: '🔐',
    signup: '✅',
    purchase: '🛒',
  }
  return icons[type as keyof typeof icons] || '📋'
}

function getActivityText(type: string) {
  const texts = {
    login: '登录系统',
    signup: '注册账号',
    purchase: '完成购买',
  }
  return texts[type as keyof typeof texts] || '执行操作'
}

function formatTimestamp(timestamp: string) {
  const date = new Date(timestamp)
  const now = new Date()
  const diff = now.getTime() - date.getTime()
  const minutes = Math.floor(diff / 60000)

  if (minutes < 1) return '刚刚'
  if (minutes < 60) return `${minutes}分钟前`
  if (minutes < 1440) return `${Math.floor(minutes / 60)}小时前`
  return date.toLocaleDateString()
}
