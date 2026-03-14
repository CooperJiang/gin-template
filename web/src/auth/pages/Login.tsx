import { useState } from 'react'
import { Link, useSearchParams } from 'react-router-dom'
import { useAuth } from '../useAuth'
import { message } from '@/lib/message'
import type { LoginRequest } from '@/types/auth'

export default function Login() {
  const [searchParams] = useSearchParams()
  const redirect = searchParams.get('redirect') || '/'
  const { login, loading } = useAuth()

  const [form, setForm] = useState({ account: '', password: '' })
  const [showPwd, setShowPwd] = useState(false)
  const [error, setError] = useState('')

  const canSubmit = form.account.trim().length > 0 && form.password.length >= 6

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setError('')
    try {
      const data: LoginRequest = { account: form.account.trim(), password: form.password }
      await login(data)
      message.success('登录成功')
      setTimeout(() => {
        window.location.href = redirect
      }, 300)
    } catch (err: unknown) {
      setError(err instanceof Error ? err.message : '登录失败，请重试')
    }
  }

  return (
    <div className="min-h-screen flex">
      {/* Left dark panel */}
      <div
        className="hidden lg:flex lg:w-[46%] relative overflow-hidden items-end"
        style={{ background: 'linear-gradient(165deg, #0f172a 0%, #1e293b 100%)' }}
      >
        <div
          className="absolute top-[18%] right-[-60px]"
          style={{
            width: 280,
            height: 280,
            borderRadius: '50%',
            border: '1px solid rgba(255,255,255,0.04)',
          }}
        />
        <div
          className="absolute top-[10%] right-[-120px]"
          style={{
            width: 440,
            height: 440,
            borderRadius: '50%',
            border: '1px solid rgba(255,255,255,0.025)',
          }}
        />
        <div
          className="absolute bottom-[-80px] left-[-40px]"
          style={{
            width: 320,
            height: 320,
            borderRadius: '50%',
            border: '1px solid rgba(255,255,255,0.03)',
          }}
        />

        <div className="relative z-10 px-12 xl:px-16 pb-14 pt-14 w-full">
          <div className="flex items-center gap-3 mb-20">
            <div
              className="flex items-center justify-center rounded-xl"
              style={{
                width: 40,
                height: 40,
                background: 'linear-gradient(135deg, #3b82f6, #6366f1)',
              }}
            >
              <svg
                width="20"
                height="20"
                viewBox="0 0 24 24"
                fill="none"
                stroke="white"
                strokeWidth="2"
                strokeLinecap="round"
                strokeLinejoin="round"
              >
                <path d="M12 2L2 7l10 5 10-5-10-5z" />
                <path d="M2 17l10 5 10-5" />
                <path d="M2 12l10 5 10-5" />
              </svg>
            </div>
            <span
              className="text-white text-lg tracking-tight"
              style={{ fontFamily: "'Outfit', sans-serif", fontWeight: 700 }}
            >
              Gin Template
            </span>
          </div>

          <h1
            className="text-4xl text-white leading-tight tracking-tight"
            style={{ fontWeight: 700 }}
          >
            欢迎回来
          </h1>
          <p className="mt-4 text-base leading-relaxed" style={{ color: '#94a3b8' }}>
            登录您的账户，访问所有功能和服务。
            <br />
            基于 React + Go 的现代化全栈应用。
          </p>

          <div
            className="mt-16 grid grid-cols-3 gap-6 border-t pt-8"
            style={{ borderColor: 'rgba(255,255,255,0.06)' }}
          >
            <div>
              <div className="text-2xl text-white" style={{ fontWeight: 700 }}>
                99.9%
              </div>
              <div className="text-xs mt-1.5" style={{ color: '#64748b' }}>
                服务可用性
              </div>
            </div>
            <div>
              <div className="text-2xl text-white" style={{ fontWeight: 700 }}>
                &lt;50ms
              </div>
              <div className="text-xs mt-1.5" style={{ color: '#64748b' }}>
                响应延迟
              </div>
            </div>
            <div>
              <div className="text-2xl text-white" style={{ fontWeight: 700 }}>
                10K+
              </div>
              <div className="text-xs mt-1.5" style={{ color: '#64748b' }}>
                活跃用户
              </div>
            </div>
          </div>
        </div>
      </div>

      {/* Right form panel */}
      <div className="flex-1 flex items-center justify-center px-6 sm:px-12 bg-white">
        <div className="w-full max-w-[360px]">
          <div className="lg:hidden flex items-center gap-2.5 mb-12">
            <div
              className="flex items-center justify-center rounded-lg"
              style={{
                width: 36,
                height: 36,
                background: 'linear-gradient(135deg, #3b82f6, #6366f1)',
              }}
            >
              <svg
                width="18"
                height="18"
                viewBox="0 0 24 24"
                fill="none"
                stroke="white"
                strokeWidth="2"
                strokeLinecap="round"
                strokeLinejoin="round"
              >
                <path d="M12 2L2 7l10 5 10-5-10-5z" />
                <path d="M2 17l10 5 10-5" />
                <path d="M2 12l10 5 10-5" />
              </svg>
            </div>
            <span
              className="text-gray-900 text-lg tracking-tight"
              style={{ fontFamily: "'Outfit', sans-serif", fontWeight: 700 }}
            >
              Gin Template
            </span>
          </div>

          <h2 className="text-2xl text-gray-900 tracking-tight" style={{ fontWeight: 700 }}>
            登录
          </h2>
          <p className="text-gray-500 text-sm mt-1.5 mb-8">请输入您的账户信息</p>

          {error && (
            <div className="mb-6 px-4 py-3 rounded-lg bg-red-50 border border-red-200 text-sm text-red-600 flex items-center gap-2.5">
              <svg
                width="16"
                height="16"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                strokeWidth="2"
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
              <label htmlFor="account" className="block text-sm font-medium text-gray-700 mb-1.5">
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
                className="block w-full rounded-lg border border-gray-200 px-3.5 py-2.5 text-sm text-gray-900 placeholder:text-gray-400 focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition disabled:bg-gray-50 disabled:opacity-60"
              />
            </div>

            <div className="mb-6">
              <div className="flex items-center justify-between mb-1.5">
                <label htmlFor="password" className="text-sm font-medium text-gray-700">
                  密码
                </label>
                <Link
                  to="/forgot-password"
                  className="text-xs text-blue-600 hover:text-blue-700 font-medium"
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
                  className="block w-full rounded-lg border border-gray-200 px-3.5 py-2.5 pr-10 text-sm text-gray-900 placeholder:text-gray-400 focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition disabled:bg-gray-50 disabled:opacity-60"
                />
                <button
                  type="button"
                  tabIndex={-1}
                  onClick={() => setShowPwd(!showPwd)}
                  className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-600"
                >
                  {showPwd ? (
                    <svg
                      width="16"
                      height="16"
                      viewBox="0 0 24 24"
                      fill="none"
                      stroke="currentColor"
                      strokeWidth="2"
                      strokeLinecap="round"
                      strokeLinejoin="round"
                    >
                      <path d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94" />
                      <path d="M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19" />
                      <line x1="1" y1="1" x2="23" y2="23" />
                    </svg>
                  ) : (
                    <svg
                      width="16"
                      height="16"
                      viewBox="0 0 24 24"
                      fill="none"
                      stroke="currentColor"
                      strokeWidth="2"
                      strokeLinecap="round"
                      strokeLinejoin="round"
                    >
                      <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z" />
                      <circle cx="12" cy="12" r="3" />
                    </svg>
                  )}
                </button>
              </div>
            </div>

            <button
              type="submit"
              disabled={loading || !canSubmit}
              className="w-full rounded-lg px-4 py-2.5 text-sm font-semibold text-white focus:outline-none focus:ring-2 focus:ring-blue-500/40 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed transition hover:brightness-110"
              style={{ background: 'linear-gradient(135deg, #3b82f6, #6366f1)' }}
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

          <p className="mt-8 text-center text-sm text-gray-500">
            还没有账户？
            <Link to="/register" className="ml-1 font-semibold text-blue-600 hover:text-blue-700">
              立即注册
            </Link>
          </p>
        </div>
      </div>
    </div>
  )
}
