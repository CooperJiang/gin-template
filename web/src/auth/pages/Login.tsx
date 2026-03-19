import { useState } from 'react'
import { Link, useSearchParams, useNavigate } from 'react-router-dom'
import { useAuth } from '../useAuth'
import { message } from '@/lib/message'
import { VALIDATION, INPUT_CLASS } from '@/lib/constants'
import { Eye, EyeOff } from 'lucide-react'
import type { LoginRequest } from '@/types/auth'

export default function Login() {
  const [searchParams] = useSearchParams()
  const redirect = searchParams.get('redirect') || '/'
  const navigate = useNavigate()
  const { login, loading } = useAuth()

  const [form, setForm] = useState({ account: '', password: '' })
  const [showPwd, setShowPwd] = useState(false)
  const [error, setError] = useState('')

  const canSubmit = form.account.trim().length > 0 && form.password.length >= VALIDATION.PASSWORD_MIN

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setError('')
    try {
      const data: LoginRequest = { account: form.account.trim(), password: form.password }
      await login(data)
      message.success('登录成功')
      setTimeout(() => navigate(redirect, { replace: true }), 300)
    } catch (err: unknown) {
      setError(err instanceof Error ? err.message : '登录失败，请重试')
    }
  }

  return (
    <div className="min-h-screen flex bg-nb-bg">
      {/* Left panel — neubrutalism style */}
      <div className="hidden lg:flex lg:w-[46%] relative overflow-hidden items-center bg-nb-primary border-r-nb border-nb-border">
        {/* Decorative shapes */}
        <div className="absolute top-[15%] right-[-40px] w-64 h-64 rounded-full border-nb border-nb-border bg-nb-surface opacity-30" />
        <div className="absolute top-[50%] right-[20%] w-32 h-32 border-nb border-nb-border bg-[var(--nb-accent-pink)] opacity-40 rotate-12" />
        <div className="absolute bottom-[20%] left-[-30px] w-48 h-48 rounded-full border-nb border-nb-border bg-[var(--nb-accent-blue)] opacity-30" />

        <div className="relative z-10 px-12 xl:px-16 pb-14 pt-14 w-full">
          <div className="flex items-center gap-3 mb-20">
            <div className="flex items-center justify-center rounded-[var(--nb-radius)] w-10 h-10 bg-nb-surface border-nb border-nb-border shadow-nb-sm">
              <svg
                width="20"
                height="20"
                viewBox="0 0 24 24"
                fill="none"
                stroke="var(--nb-text)"
                strokeWidth="2.5"
                strokeLinecap="round"
                strokeLinejoin="round"
              >
                <path d="M12 2L2 7l10 5 10-5-10-5z" />
                <path d="M2 17l10 5 10-5" />
                <path d="M2 12l10 5 10-5" />
              </svg>
            </div>
            <span className="text-[var(--nb-primary-text)] text-lg font-bold tracking-tight">
              Gin Template
            </span>
          </div>

          <h1 className="text-4xl font-bold text-[var(--nb-primary-text)] leading-tight tracking-tight">
            欢迎回来
          </h1>
          <p className="mt-4 text-base leading-relaxed text-[var(--nb-primary-text)] opacity-70">
            登录您的账户，访问所有功能和服务。
            <br />
            基于 React + Go 的现代化全栈应用。
          </p>

          <div className="mt-16 grid grid-cols-3 gap-6 border-t-nb border-[var(--nb-primary-text)] pt-8">
            <div>
              <div className="text-2xl font-bold text-[var(--nb-primary-text)]">99.9%</div>
              <div className="text-xs mt-1.5 text-[var(--nb-primary-text)] opacity-60">服务可用性</div>
            </div>
            <div>
              <div className="text-2xl font-bold text-[var(--nb-primary-text)]">&lt;50ms</div>
              <div className="text-xs mt-1.5 text-[var(--nb-primary-text)] opacity-60">响应延迟</div>
            </div>
            <div>
              <div className="text-2xl font-bold text-[var(--nb-primary-text)]">10K+</div>
              <div className="text-xs mt-1.5 text-[var(--nb-primary-text)] opacity-60">活跃用户</div>
            </div>
          </div>
        </div>
      </div>

      {/* Right form panel */}
      <div className="flex-1 flex items-center justify-center px-6 sm:px-12 bg-nb-bg">
        <div className="w-full max-w-[360px]">
          <div className="lg:hidden flex items-center gap-2.5 mb-12">
            <div className="flex items-center justify-center rounded-[var(--nb-radius)] w-9 h-9 bg-nb-primary border-nb border-nb-border shadow-nb-sm">
              <svg
                width="18"
                height="18"
                viewBox="0 0 24 24"
                fill="none"
                stroke="var(--nb-primary-text)"
                strokeWidth="2.5"
                strokeLinecap="round"
                strokeLinejoin="round"
              >
                <path d="M12 2L2 7l10 5 10-5-10-5z" />
                <path d="M2 17l10 5 10-5" />
                <path d="M2 12l10 5 10-5" />
              </svg>
            </div>
            <span className="text-nb-text text-lg font-bold tracking-tight">
              Gin Template
            </span>
          </div>

          <h2 className="text-2xl font-bold text-nb-text tracking-tight">登录</h2>
          <p className="text-nb-text-secondary text-sm mt-1.5 mb-8">请输入您的账户信息</p>

          {error && (
            <div className="mb-6 px-4 py-3 rounded-[var(--nb-radius)] bg-[var(--nb-accent-red)] border-nb border-nb-border shadow-nb-sm text-sm font-bold text-white flex items-center gap-2.5">
              <svg
                width="16"
                height="16"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                strokeWidth="2.5"
                strokeLinecap="round"
                strokeLinejoin="round"
                className="shrink-0"
              >
                <circle cx="12" cy="12" r="10" />
                <line x1="15" y1="9" x2="9" y2="15" />
                <line x1="9" y1="9" x2="15" y2="15" />
              </svg>
              {error}
            </div>
          )}

          <form onSubmit={handleSubmit}>
            <div className="mb-5">
              <label
                htmlFor="account"
                className="block text-sm font-bold text-nb-text mb-1.5"
              >
                账户
              </label>
              <input
                id="account"
                type="text"
                autoComplete="username"
                autoFocus
                required
                disabled={loading}
                value={form.account}
                onChange={(e) => setForm({ ...form, account: e.target.value })}
                placeholder="用户名或邮箱"
                className={INPUT_CLASS}
              />
            </div>

            <div className="mb-6">
              <div className="flex items-center justify-between mb-1.5">
                <label
                  htmlFor="password"
                  className="text-sm font-bold text-nb-text"
                >
                  密码
                </label>
                <Link
                  to="/forgot-password"
                  className="text-xs font-bold text-nb-text-secondary hover:text-nb-text underline underline-offset-2"
                >
                  忘记密码？
                </Link>
              </div>
              <div className="relative">
                <input
                  id="password"
                  type={showPwd ? 'text' : 'password'}
                  autoComplete="current-password"
                  required
                  disabled={loading}
                  value={form.password}
                  onChange={(e) => setForm({ ...form, password: e.target.value })}
                  placeholder="请输入密码"
                  className={INPUT_CLASS + ' pr-10'}
                />
                <button
                  type="button"
                  tabIndex={-1}
                  onClick={() => setShowPwd(!showPwd)}
                  className="absolute right-3 top-1/2 -translate-y-1/2 text-nb-text-muted hover:text-nb-text"
                >
                  {showPwd ? (
                    <EyeOff className="w-4 h-4" />
                  ) : (
                    <Eye className="w-4 h-4" />
                  )}
                </button>
              </div>
            </div>

            <button
              type="submit"
              disabled={loading || !canSubmit}
              className="w-full rounded-[var(--nb-radius)] px-4 py-2.5 text-sm font-bold text-[var(--nb-primary-text)] bg-nb-primary border-nb border-nb-border shadow-nb nb-interactive focus:outline-none focus:ring-2 focus:ring-nb-primary disabled:opacity-50 disabled:cursor-not-allowed disabled:transform-none disabled:shadow-nb"
            >
              {loading ? (
                <span className="inline-flex items-center gap-2">
                  <svg width="16" height="16" viewBox="0 0 24 24" className="animate-spin">
                    <circle
                      className="opacity-25"
                      cx="12"
                      cy="12"
                      r="10"
                      stroke="currentColor"
                      strokeWidth="4"
                      fill="none"
                    />
                    <path
                      className="opacity-75"
                      fill="currentColor"
                      d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"
                    />
                  </svg>
                  登录中...
                </span>
              ) : (
                '登录'
              )}
            </button>
          </form>

          <p className="mt-8 text-center text-sm text-nb-text-secondary">
            还没有账户？
            <Link to="/register" className="ml-1 font-bold text-nb-text underline underline-offset-2 hover:text-nb-primary">
              立即注册
            </Link>
          </p>
        </div>
      </div>
    </div>
  )
}
