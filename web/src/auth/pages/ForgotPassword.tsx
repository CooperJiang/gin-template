import { useState, useRef, useEffect } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { useAuth } from '../useAuth'
import { message } from '@/lib/message'
import { VALIDATION, EMAIL_REGEX, INPUT_CLASS } from '@/lib/constants'
import { Eye, EyeOff } from 'lucide-react'
import type { ResetPasswordRequest } from '@/types/auth'

export default function ForgotPassword() {
  const navigate = useNavigate()
  const { resetPassword, sendResetPasswordCode, loading } = useAuth()

  const [form, setForm] = useState({
    email: '',
    code: '',
    newPassword: '',
    confirmPassword: '',
  })
  const [showPwd, setShowPwd] = useState(false)
  const [error, setError] = useState('')
  const [countdown, setCountdown] = useState(0)
  const timerRef = useRef<ReturnType<typeof setInterval>>(undefined)

  useEffect(
    () => () => {
      if (timerRef.current) clearInterval(timerRef.current)
    },
    [],
  )

  const canSubmit =
    form.email.trim().length > 0 &&
    form.code.trim().length > 0 &&
    form.newPassword.length >= VALIDATION.PASSWORD_MIN &&
    form.confirmPassword.length > 0

  const handleSendCode = async () => {
    if (!form.email.trim()) {
      setError('请输入邮箱')
      return
    }
    if (!EMAIL_REGEX.test(form.email.trim())) {
      setError('请输入有效的邮箱地址')
      return
    }
    try {
      setError('')
      await sendResetPasswordCode(form.email.trim())
      message.success('验证码已发送到您的邮箱')
      setCountdown(VALIDATION.CODE_COUNTDOWN)
      timerRef.current = setInterval(() => {
        setCountdown((prev) => {
          if (prev <= 1) {
            clearInterval(timerRef.current)
            return 0
          }
          return prev - 1
        })
      }, 1000)
    } catch (err: unknown) {
      setError(err instanceof Error ? err.message : '发送验证码失败')
    }
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setError('')
    if (
      form.newPassword.length < VALIDATION.PASSWORD_MIN ||
      form.newPassword.length > VALIDATION.PASSWORD_MAX
    ) {
      setError(`密码长度为 ${VALIDATION.PASSWORD_MIN}-${VALIDATION.PASSWORD_MAX} 个字符`)
      return
    }
    if (form.newPassword !== form.confirmPassword) {
      setError('两次输入的密码不一致')
      return
    }
    try {
      const data: ResetPasswordRequest = {
        email: form.email.trim(),
        code: form.code.trim(),
        newPassword: form.newPassword,
      }
      await resetPassword(data)
      message.success('密码重置成功，请登录')
      setTimeout(() => navigate('/login'), 800)
    } catch (err: unknown) {
      setError(err instanceof Error ? err.message : '重置密码失败，请重试')
    }
  }

  return (
    <div className="min-h-screen flex bg-nb-bg">
      {/* Left panel — neubrutalism style */}
      <div className="hidden lg:flex lg:w-[45%] relative overflow-hidden items-center justify-center bg-[var(--nb-accent-blue)] border-r-nb border-nb-border">
        {/* Decorative shapes */}
        <div className="absolute top-[20%] left-[15%] w-48 h-48 rounded-full border-nb border-nb-border bg-nb-primary opacity-50" />
        <div className="absolute bottom-[10%] right-[5%] w-36 h-36 border-nb border-nb-border bg-[var(--nb-accent-orange)] opacity-40 rotate-45" />
        <div className="absolute top-[60%] left-[-20px] w-28 h-28 border-nb border-nb-border bg-nb-surface opacity-30" />

        <div className="relative z-10 px-12 xl:px-20 max-w-lg">
          <div className="flex items-center gap-3 mb-16">
            <div className="w-10 h-10 rounded-[var(--nb-radius)] flex items-center justify-center bg-nb-surface border-nb border-nb-border shadow-nb-sm">
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
            <span className="text-[var(--nb-primary-text)] text-xl font-bold">
              Gin Template
            </span>
          </div>

          <h1 className="text-4xl font-bold text-[var(--nb-primary-text)] leading-tight tracking-tight">重置密码</h1>
          <p className="mt-4 text-[var(--nb-primary-text)] opacity-70 text-base leading-relaxed">
            通过邮箱验证码安全地重置您的密码，整个过程只需几步。
          </p>

          <div className="mt-14 space-y-0">
            {[
              { step: '01', t: '输入注册邮箱', d: '我们将发送验证码到您的邮箱' },
              { step: '02', t: '填写验证码', d: '输入邮箱中收到的 6 位数字验证码' },
              { step: '03', t: '设置新密码', d: '输入并确认您的新密码即可完成' },
            ].map((item, i) => (
              <div key={item.step} className="flex items-start gap-4">
                <div className="relative flex flex-col items-center">
                  <div className="w-10 h-10 rounded-[var(--nb-radius-sm)] bg-nb-surface border-nb border-nb-border shadow-nb-sm flex items-center justify-center shrink-0">
                    <span className="text-xs font-bold text-nb-text">{item.step}</span>
                  </div>
                  {i < 2 && <div className="w-[2.5px] h-8 bg-nb-border mt-1" />}
                </div>
                <div className="pt-2">
                  <div className="text-sm font-bold text-[var(--nb-primary-text)]">{item.t}</div>
                  <div className="text-xs text-[var(--nb-primary-text)] opacity-60 mt-0.5">{item.d}</div>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>

      {/* Right form panel */}
      <div className="flex-1 flex items-center justify-center px-6 sm:px-12 py-12 bg-nb-bg">
        <div className="w-full max-w-sm">
          <div className="lg:hidden flex items-center gap-2.5 mb-12">
            <div className="w-9 h-9 rounded-[var(--nb-radius)] flex items-center justify-center bg-nb-primary border-nb border-nb-border shadow-nb-sm">
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
            <span className="text-lg text-nb-text font-bold">
              Gin Template
            </span>
          </div>

          <h2 className="text-2xl font-bold text-nb-text tracking-tight">重置密码</h2>
          <p className="text-nb-text-secondary text-sm mt-1.5 mb-7">通过邮箱验证码重置您的密码</p>

          {error && (
            <div className="mb-5 px-4 py-3 rounded-[var(--nb-radius)] bg-[var(--nb-accent-red)] border-nb border-nb-border shadow-nb-sm text-sm font-bold text-white flex items-center gap-2">
              <svg
                width="16"
                height="16"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                strokeWidth="2.5"
                strokeLinecap="round"
                strokeLinejoin="round"
              >
                <circle cx="12" cy="12" r="10" />
                <line x1="15" y1="9" x2="9" y2="15" />
                <line x1="9" y1="9" x2="15" y2="15" />
              </svg>
              {error}
            </div>
          )}

          <form onSubmit={handleSubmit}>
            <div className="mb-4">
              <label
                htmlFor="email"
                className="block text-sm font-bold text-nb-text mb-1.5"
              >
                邮箱
              </label>
              <input
                id="email"
                type="email"
                autoComplete="email"
                autoFocus
                required
                disabled={loading}
                value={form.email}
                onChange={(e) => setForm({ ...form, email: e.target.value })}
                placeholder="请输入注册邮箱"
                className={INPUT_CLASS}
              />
            </div>

            <div className="mb-4">
              <label
                htmlFor="code"
                className="block text-sm font-bold text-nb-text mb-1.5"
              >
                验证码
              </label>
              <div className="flex gap-2">
                <input
                  id="code"
                  type="text"
                  required
                  disabled={loading}
                  maxLength={VALIDATION.CODE_LENGTH}
                  value={form.code}
                  onChange={(e) =>
                    setForm({ ...form, code: e.target.value.replace(/\D/g, '') })
                  }
                  placeholder={`${VALIDATION.CODE_LENGTH} 位数字`}
                  className={INPUT_CLASS + ' flex-1'}
                />
                <button
                  type="button"
                  onClick={handleSendCode}
                  disabled={loading || countdown > 0}
                  className="shrink-0 px-4 py-2.5 text-sm font-bold text-[var(--nb-primary-text)] bg-nb-primary border-nb border-nb-border rounded-[var(--nb-radius)] shadow-nb-sm nb-interactive disabled:opacity-50 disabled:cursor-not-allowed disabled:transform-none whitespace-nowrap"
                >
                  {countdown > 0 ? `${countdown}s` : '发送验证码'}
                </button>
              </div>
            </div>

            <div className="mb-4">
              <label
                htmlFor="newPassword"
                className="block text-sm font-bold text-nb-text mb-1.5"
              >
                新密码
              </label>
              <div className="relative">
                <input
                  id="newPassword"
                  type={showPwd ? 'text' : 'password'}
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

            <div className="mb-6">
              <label
                htmlFor="confirmPassword"
                className="block text-sm font-bold text-nb-text mb-1.5"
              >
                确认密码
              </label>
              <input
                id="confirmPassword"
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
                  重置中...
                </span>
              ) : (
                '重置密码'
              )}
            </button>
          </form>

          <p className="mt-7 text-center text-sm text-nb-text-secondary">
            想起密码了？
            <Link to="/login" className="ml-1 font-bold text-nb-text underline underline-offset-2 hover:text-nb-primary">
              返回登录
            </Link>
          </p>
        </div>
      </div>
    </div>
  )
}
