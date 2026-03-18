import { useState } from 'react'
import { useAuth } from '@/auth/useAuth'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { message } from '@/lib/message'
import { INPUT_CLASS, VALIDATION } from '@/lib/constants'
import { User, Mail, Calendar, Pencil, X } from 'lucide-react'

export default function Profile() {
  const { user, getUserInfo, updateProfile, loading } = useAuth()
  const [refreshing, setRefreshing] = useState(false)
  const [editing, setEditing] = useState(false)
  const [form, setForm] = useState({ username: '', bio: '' })

  const handleRefresh = async () => {
    try {
      setRefreshing(true)
      await getUserInfo()
      message.success('资料已刷新')
    } catch {
      // error handled by useAuth
    } finally {
      setRefreshing(false)
    }
  }

  const startEdit = () => {
    setForm({
      username: user?.username || '',
      bio: user?.bio || '',
    })
    setEditing(true)
  }

  const cancelEdit = () => {
    setEditing(false)
  }

  const handleSave = async (e: React.FormEvent) => {
    e.preventDefault()
    const trimmed = form.username.trim()
    if (
      trimmed.length < VALIDATION.USERNAME_MIN ||
      trimmed.length > VALIDATION.USERNAME_MAX
    ) {
      message.error(`用户名长度为 ${VALIDATION.USERNAME_MIN}-${VALIDATION.USERNAME_MAX} 个字符`)
      return
    }
    try {
      await updateProfile({ username: trimmed })
      message.success('资料已更新')
      setEditing(false)
    } catch {
      // error handled by useAuth
    }
  }

  return (
    <div>
      <div className="mb-8">
        <h1 className="text-2xl font-bold text-gray-900 dark:text-gray-100">个人资料</h1>
        <p className="text-gray-500 dark:text-gray-400 mt-1">查看和管理您的账户信息</p>
      </div>

      <div className="max-w-2xl">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between">
            <CardTitle>基本信息</CardTitle>
            <div className="flex gap-2">
              {!editing && (
                <>
                  <Button variant="outline" size="sm" onClick={startEdit}>
                    <Pencil className="w-4 h-4 mr-1" />
                    编辑
                  </Button>
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={handleRefresh}
                    disabled={loading || refreshing}
                  >
                    {refreshing ? '刷新中...' : '刷新'}
                  </Button>
                </>
              )}
              {editing && (
                <Button variant="outline" size="sm" onClick={cancelEdit}>
                  <X className="w-4 h-4 mr-1" />
                  取消
                </Button>
              )}
            </div>
          </CardHeader>
          <CardContent>
            <div className="flex items-center gap-6 mb-8">
              <div className="w-20 h-20 rounded-full bg-blue-100 dark:bg-blue-900/40 flex items-center justify-center">
                <span className="text-3xl font-bold text-blue-700 dark:text-blue-400">
                  {user?.username?.charAt(0).toUpperCase() || 'U'}
                </span>
              </div>
              <div>
                <h3 className="text-lg font-semibold text-gray-900 dark:text-gray-100">
                  {user?.username || '—'}
                </h3>
                <p className="text-sm text-gray-500 dark:text-gray-400">
                  {user?.bio || '暂无简介'}
                </p>
              </div>
            </div>

            {editing ? (
              <form onSubmit={handleSave} className="space-y-4">
                <div>
                  <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1.5">
                    用户名
                  </label>
                  <input
                    type="text"
                    required
                    disabled={loading}
                    value={form.username}
                    onChange={(e) => setForm({ ...form, username: e.target.value })}
                    placeholder={`${VALIDATION.USERNAME_MIN}-${VALIDATION.USERNAME_MAX} 个字符`}
                    className={INPUT_CLASS}
                  />
                </div>
                <div className="pt-2">
                  <Button type="submit" disabled={loading}>
                    {loading ? '保存中...' : '保存'}
                  </Button>
                </div>
              </form>
            ) : (
              <div className="space-y-4">
                <div className="flex items-center gap-3 p-4 bg-gray-50 dark:bg-gray-800 rounded-lg">
                  <User className="w-5 h-5 text-gray-400" />
                  <div>
                    <p className="text-xs text-gray-500 dark:text-gray-400">用户名</p>
                    <p className="text-sm font-medium text-gray-900 dark:text-gray-100">
                      {user?.username || '—'}
                    </p>
                  </div>
                </div>

                <div className="flex items-center gap-3 p-4 bg-gray-50 dark:bg-gray-800 rounded-lg">
                  <Mail className="w-5 h-5 text-gray-400" />
                  <div>
                    <p className="text-xs text-gray-500 dark:text-gray-400">邮箱</p>
                    <p className="text-sm font-medium text-gray-900 dark:text-gray-100">
                      {user?.email || '—'}
                    </p>
                  </div>
                </div>

                <div className="flex items-center gap-3 p-4 bg-gray-50 dark:bg-gray-800 rounded-lg">
                  <Calendar className="w-5 h-5 text-gray-400" />
                  <div>
                    <p className="text-xs text-gray-500 dark:text-gray-400">注册时间</p>
                    <p className="text-sm font-medium text-gray-900 dark:text-gray-100">
                      {user?.created_at
                        ? new Date(user.created_at).toLocaleDateString('zh-CN')
                        : '—'}
                    </p>
                  </div>
                </div>
              </div>
            )}
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
