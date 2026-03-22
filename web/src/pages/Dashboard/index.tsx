import { useState, useEffect } from 'react'
import { useAuth } from '@/auth/useAuth'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Mail, MailCheck, MailX, Loader2 } from 'lucide-react'
import { mailAccountsApi } from '@/api/mailAccounts'
import type { MailAccountStats, DailyUsageItem } from '@/types/mailAccount'
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
} from 'recharts'

export default function Dashboard() {
  const { user } = useAuth()
  const [statsLoading, setStatsLoading] = useState(true)
  const [usageLoading, setUsageLoading] = useState(true)
  const [stats, setStats] = useState<MailAccountStats | null>(null)
  const [dailyUsage, setDailyUsage] = useState<DailyUsageItem[]>([])

  useEffect(() => {
    const loadStats = async () => {
      try {
        const data = await mailAccountsApi.stats()
        setStats(data)
      } catch {
        // handled by apiClient
      } finally {
        setStatsLoading(false)
      }
    }

    const loadUsage = async () => {
      try {
        const data = await mailAccountsApi.dailyUsage()
        setDailyUsage(data.items ?? [])
      } catch {
        // handled by apiClient
      } finally {
        setUsageLoading(false)
      }
    }

    void loadStats()
    void loadUsage()
  }, [])

  const statCards = [
    {
      title: '总邮箱数',
      value: stats?.total ?? 0,
      icon: Mail,
      color:
        'text-blue-600 bg-blue-100 dark:text-blue-400 dark:bg-blue-900/30',
      valueColor: 'text-blue-700 dark:text-blue-300',
    },
    {
      title: '未使用',
      value: stats?.unused ?? 0,
      icon: MailCheck,
      color:
        'text-emerald-600 bg-emerald-100 dark:text-emerald-400 dark:bg-emerald-900/30',
      valueColor: 'text-emerald-700 dark:text-emerald-300',
    },
    {
      title: '已使用',
      value: stats?.used ?? 0,
      icon: MailX,
      color:
        'text-orange-600 bg-orange-100 dark:text-orange-400 dark:bg-orange-900/30',
      valueColor: 'text-orange-700 dark:text-orange-300',
    },
  ]

  return (
    <div>
      <div className="mb-6">
        <h1 className="text-xl font-bold text-gray-900 dark:text-gray-100">
          欢迎回来，{user?.username || '用户'}
        </h1>
        <p className="text-gray-500 dark:text-gray-400 mt-1 text-sm">
          以下是您的邮箱资源概览
        </p>
      </div>

      {/* Stats cards */}
      <div className="grid grid-cols-1 sm:grid-cols-3 gap-4 mb-6">
        {statsLoading
          ? Array.from({ length: 3 }).map((_, i) => (
              <Card key={i}>
                <CardContent className="p-5">
                  <div className="flex items-center justify-between animate-pulse">
                    <div>
                      <div className="h-4 w-16 bg-gray-200 dark:bg-gray-700 rounded mb-3" />
                      <div className="h-7 w-12 bg-gray-200 dark:bg-gray-700 rounded" />
                    </div>
                    <div className="w-11 h-11 rounded-lg bg-gray-200 dark:bg-gray-700" />
                  </div>
                </CardContent>
              </Card>
            ))
          : statCards.map((card) => {
              const Icon = card.icon
              return (
                <Card key={card.title}>
                  <CardContent className="p-5">
                    <div className="flex items-center justify-between">
                      <div>
                        <p className="text-sm text-gray-500 dark:text-gray-400">
                          {card.title}
                        </p>
                        <p
                          className={`text-2xl font-bold mt-1 ${card.valueColor}`}
                        >
                          {card.value.toLocaleString()}
                        </p>
                      </div>
                      <div
                        className={`w-11 h-11 rounded-lg flex items-center justify-center ${card.color}`}
                      >
                        <Icon className="w-5 h-5" />
                      </div>
                    </div>
                  </CardContent>
                </Card>
              )
            })}
      </div>

      {/* Daily usage chart */}
      <Card>
        <CardHeader className="pb-2">
          <CardTitle className="text-base">每日使用量（近 30 天）</CardTitle>
        </CardHeader>
        <CardContent className="pt-2">
          {usageLoading ? (
            <div className="flex items-center justify-center h-64 text-gray-400">
              <Loader2 className="h-5 w-5 animate-spin mr-2" />
              加载中...
            </div>
          ) : dailyUsage.length === 0 ? (
            <div className="flex items-center justify-center h-64 text-gray-400 text-sm">
              暂无使用量数据
            </div>
          ) : (
            <ResponsiveContainer width="100%" height={280}>
              <BarChart
                data={dailyUsage}
                margin={{ top: 8, right: 8, left: -16, bottom: 0 }}
              >
                <CartesianGrid
                  strokeDasharray="3 3"
                  vertical={false}
                  stroke="currentColor"
                  className="text-gray-200 dark:text-gray-700"
                />
                <XAxis
                  dataKey="date"
                  tick={{ fontSize: 11 }}
                  tickFormatter={(v: string) => v.slice(5)}
                  stroke="currentColor"
                  className="text-gray-400"
                />
                <YAxis
                  allowDecimals={false}
                  tick={{ fontSize: 11 }}
                  stroke="currentColor"
                  className="text-gray-400"
                />
                <Tooltip
                  contentStyle={{
                    backgroundColor: 'var(--color-gray-800, #1f2937)',
                    border: 'none',
                    borderRadius: '8px',
                    color: '#fff',
                    fontSize: '13px',
                  }}
                  labelFormatter={(label) => `日期: ${String(label)}`}
                  formatter={(value) => [`${String(value)} 个`, '使用量']}
                />
                <Bar
                  dataKey="count"
                  fill="#3b82f6"
                  radius={[4, 4, 0, 0]}
                  maxBarSize={32}
                />
              </BarChart>
            </ResponsiveContainer>
          )}
        </CardContent>
      </Card>
    </div>
  )
}
