import { useState } from 'react'
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { useAuth } from '@/auth/useAuth'
import { message } from '@/lib/message'
import { VALIDATION, INPUT_CLASS } from '@/lib/constants'
import { Eye, EyeOff } from 'lucide-react'

export default function Settings() {
  const { changePassword, loading } = useAuth()
  const [form, setForm] = useState({ oldPassword: '', newPassword: '', confirmPassword: '' })
  const [showOld, setShowOld] = useState(false)
  const [showNew, setShowNew] = useState(false)

  const canSubmit =
    form.oldPassword.length >= VALIDATION.PASSWORD_MIN &&
    form.newPassword.length >= VALIDATION.PASSWORD_MIN &&
    form.confirmPassword.length > 0

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    if (
      form.newPassword.length < VALIDATION.PASSWORD_MIN ||
      form.newPassword.length > VALIDATION.PASSWORD_MAX
    ) {
      message.error(`新密码长度为 ${VALIDATION.PASSWORD_MIN}-${VALIDATION.PASSWORD_MAX} 个字符`)
      return
    }
    if (form.newPassword !== form.confirmPassword) {
      message.error('两次输入的密码不一致')
      return
    }
    if (form.newPassword === form.oldPassword) {
      message.error('新密码不能与旧密码相同')
      return
    }
    try {
      await changePassword({
        oldPassword: form.oldPassword,
        newPassword: form.newPassword,
      })
      message.success('密码修改成功')
      setForm({ oldPassword: '', newPassword: '', confirmPassword: '' })
    } catch {
      // error handled by useAuth
    }
  }

  return (
    <div>
      <div className="mb-8">
        <h1 className="text-2xl font-bold text-nb-text">系统设置</h1>
        <p className="text-nb-text-secondary mt-1">管理应用程序设置和偏好</p>
      </div>

      <div className="max-w-2xl space-y-6">
        <Card>
          <CardHeader>
            <CardTitle>修改密码</CardTitle>
            <CardDescription>更新您的账户密码，建议定期修改</CardDescription>
          </CardHeader>
          <CardContent>
            <form onSubmit={handleSubmit} className="space-y-4">
              <div>
                <label className="block text-sm font-bold text-nb-text mb-1.5">
                  当前密码
                </label>
                <div className="relative">
                  <input
                    type={showOld ? 'text' : 'password'}
                    autoComplete="current-password"
                    required
                    disabled={loading}
                    value={form.oldPassword}
                    onChange={(e) => setForm({ ...form, oldPassword: e.target.value })}
                    placeholder="请输入当前密码"
                    className={INPUT_CLASS + ' pr-10'}
                  />
                  <button
                    type="button"
                    tabIndex={-1}
                    onClick={() => setShowOld(!showOld)}
                    className="absolute right-3 top-1/2 -translate-y-1/2 text-nb-text-muted hover:text-nb-text"
                  >
                    {showOld ? <EyeOff className="w-4 h-4" /> : <Eye className="w-4 h-4" />}
                  </button>
                </div>
              </div>

              <div>
                <label className="block text-sm font-bold text-nb-text mb-1.5">
                  新密码
                </label>
                <div className="relative">
                  <input
                    type={showNew ? 'text' : 'password'}
                    autoComplete="new-password"
                    required
                    disabled={loading}
                    value={form.newPassword}
                    onChange={(e) => setForm({ ...form, newPassword: e.target.value })}
                    placeholder={`${VALIDATION.PASSWORD_MIN}-${VALIDATION.PASSWORD_MAX} 个字符`}
                    className={INPUT_CLASS + ' pr-10'}
                  />
                  <button
                    type="button"
                    tabIndex={-1}
                    onClick={() => setShowNew(!showNew)}
                    className="absolute right-3 top-1/2 -translate-y-1/2 text-nb-text-muted hover:text-nb-text"
                  >
                    {showNew ? <EyeOff className="w-4 h-4" /> : <Eye className="w-4 h-4" />}
                  </button>
                </div>
              </div>

              <div>
                <label className="block text-sm font-bold text-nb-text mb-1.5">
                  确认新密码
                </label>
                <input
                  type="password"
                  autoComplete="new-password"
                  required
                  disabled={loading}
                  value={form.confirmPassword}
                  onChange={(e) => setForm({ ...form, confirmPassword: e.target.value })}
                  placeholder="请再次输入新密码"
                  className={INPUT_CLASS}
                />
              </div>

              <div className="pt-2">
                <Button type="submit" disabled={loading || !canSubmit}>
                  {loading ? '提交中...' : '修改密码'}
                </Button>
              </div>
            </form>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
