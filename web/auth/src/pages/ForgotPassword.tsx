import { useState, useRef, useEffect } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { useAuth } from '../useAuth'
import { useMessage } from '../useMessage'
import type { ResetPasswordRequest } from '@app/shared'

export default function ForgotPassword() {
  const navigate = useNavigate()
  const { resetPassword, sendResetPasswordCode, loading } = useAuth()
  const { success } = useMessage()

  const [form, setForm] = useState({ email: '', code: '', newPassword: '', confirmPassword: '' })
  const [showPwd, setShowPwd] = useState(false)
  const [error, setError] = useState('')
  const [countdown, setCountdown] = useState(0)
  const timerRef = useRef<ReturnType<typeof setInterval>>(undefined)

  useEffect(() => () => { if (timerRef.current) clearInterval(timerRef.current) }, [])

  const canSubmit = form.email.trim().length > 0 && form.code.trim().length > 0 && form.newPassword.length >= 6 && form.confirmPassword.length > 0

  const handleSendCode = async () => {
    if (!form.email.trim()) { setError('请输入邮箱'); return }
    if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(form.email.trim())) { setError('请输入有效的邮箱地址'); return }
    try {
      setError('')
      await sendResetPasswordCode(form.email.trim())
      success('验证码已发送到您的邮箱')
      setCountdown(60)
      timerRef.current = setInterval(() => {
        setCountdown(prev => { if (prev <= 1) { clearInterval(timerRef.current); return 0 } return prev - 1 })
      }, 1000)
    } catch (err: unknown) { setError(err instanceof Error ? err.message : '发送验证码失败') }
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setError('')
    if (form.newPassword.length < 6 || form.newPassword.length > 20) { setError('密码长度为 6-20 个字符'); return }
    if (form.newPassword !== form.confirmPassword) { setError('两次输入的密码不一致'); return }
    try {
      const data: ResetPasswordRequest = { email: form.email.trim(), code: form.code.trim(), newPassword: form.newPassword }
      await resetPassword(data)
      success('密码重置成功，请登录')
      setTimeout(() => navigate('/login'), 800)
    } catch (err: unknown) { setError(err instanceof Error ? err.message : '重置密码失败，请重试') }
  }

  const inputCls = "block w-full rounded-lg border border-gray-200 px-3.5 py-2.5 text-sm text-gray-900 placeholder:text-gray-400 focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition disabled:bg-gray-50 disabled:opacity-60"

  return (
    <div className="min-h-screen flex">
      {/* Left dark panel */}
      <div className="hidden lg:flex lg:w-[45%] relative overflow-hidden items-center justify-center" style={{ background: 'linear-gradient(to bottom right, #0f172a, #1e293b)' }}>
        <div className="absolute top-1/4 left-1/4 w-72 h-72 rounded-full border border-white/[0.04]" />
        <div className="absolute bottom-1/4 right-1/4 w-96 h-96 rounded-full border border-white/[0.03]" />
        <div className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[500px] h-[500px] rounded-full border border-white/[0.02]" />

        <div className="relative z-10 px-12 xl:px-20 max-w-lg">
          <div className="flex items-center gap-3 mb-16">
            <div className="w-10 h-10 rounded-xl flex items-center justify-center" style={{ background: 'linear-gradient(135deg, #3b82f6, #6366f1)' }}>
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="white" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M12 2L2 7l10 5 10-5-10-5z" /><path d="M2 17l10 5 10-5" /><path d="M2 12l10 5 10-5" /></svg>
            </div>
            <span className="text-white text-xl" style={{ fontFamily: "'Outfit', sans-serif", fontWeight: 700 }}>应用平台</span>
          </div>

          <h1 className="text-4xl font-bold text-white leading-tight tracking-tight">重置密码</h1>
          <p className="mt-4 text-slate-400 text-base leading-relaxed">通过邮箱验证码安全地重置您的密码，整个过程只需几步。</p>

          <div className="mt-14 space-y-0">
            {[
              { step: '01', t: '输入注册邮箱', d: '我们将发送验证码到您的邮箱' },
              { step: '02', t: '填写验证码', d: '输入邮箱中收到的 6 位数字验证码' },
              { step: '03', t: '设置新密码', d: '输入并确认您的新密码即可完成' },
            ].map((item, i) => (
              <div key={item.step} className="flex items-start gap-4">
                <div className="relative flex flex-col items-center">
                  <div className="w-10 h-10 rounded-lg bg-blue-500/10 border border-blue-500/20 flex items-center justify-center shrink-0">
                    <span className="text-xs font-bold text-blue-400">{item.step}</span>
                  </div>
                  {i < 2 && <div className="w-px h-8 bg-white/10 mt-1" />}
                </div>
                <div className="pt-2">
                  <div className="text-sm font-medium text-white">{item.t}</div>
                  <div className="text-xs text-slate-500 mt-0.5">{item.d}</div>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>

      {/* Right form panel */}
      <div className="flex-1 flex items-center justify-center px-6 sm:px-12 py-12 bg-white">
        <div className="w-full max-w-sm">
          {/* Mobile-only logo */}
          <div className="lg:hidden flex items-center gap-2.5 mb-12">
            <div className="w-9 h-9 rounded-lg flex items-center justify-center" style={{ background: 'linear-gradient(135deg, #3b82f6, #6366f1)' }}>
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="white" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M12 2L2 7l10 5 10-5-10-5z" /><path d="M2 17l10 5 10-5" /><path d="M2 12l10 5 10-5" /></svg>
            </div>
            <span className="text-lg text-gray-900" style={{ fontFamily: "'Outfit', sans-serif", fontWeight: 700 }}>应用平台</span>
          </div>

          <h2 className="text-2xl font-bold text-gray-900 tracking-tight">重置密码</h2>
          <p className="text-gray-500 text-sm mt-1.5 mb-7">通过邮箱验证码重置您的密码</p>

          {error && (
            <div className="mb-5 px-4 py-3 rounded-lg bg-red-50 border border-red-200 text-sm text-red-600 flex items-center gap-2">
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><circle cx="12" cy="12" r="10" /><line x1="15" y1="9" x2="9" y2="15" /><line x1="9" y1="9" x2="15" y2="15" /></svg>
              {error}
            </div>
          )}

          <form onSubmit={handleSubmit}>
            <div className="mb-4">
              <label htmlFor="email" className="block text-sm font-medium text-gray-700 mb-1.5">邮箱</label>
              <input id="email" type="email" autoComplete="email" autoFocus required disabled={loading}
                value={form.email} onChange={e => setForm({ ...form, email: e.target.value })}
                placeholder="请输入注册邮箱" className={inputCls} />
            </div>

            <div className="mb-4">
              <label htmlFor="code" className="block text-sm font-medium text-gray-700 mb-1.5">验证码</label>
              <div className="flex gap-2">
                <input id="code" type="text" required disabled={loading} maxLength={6}
                  value={form.code} onChange={e => setForm({ ...form, code: e.target.value.replace(/\D/g, '') })}
                  placeholder="6 位数字" className={inputCls + ' flex-1'} />
                <button type="button" onClick={handleSendCode} disabled={loading || countdown > 0}
                  className="shrink-0 px-4 py-2.5 text-sm font-medium text-blue-600 bg-blue-50 border border-blue-200 rounded-lg hover:bg-blue-100 disabled:opacity-50 disabled:cursor-not-allowed whitespace-nowrap transition">
                  {countdown > 0 ? `${countdown}s` : '发送验证码'}
                </button>
              </div>
            </div>

            <div className="mb-4">
              <label htmlFor="newPassword" className="block text-sm font-medium text-gray-700 mb-1.5">新密码</label>
              <div className="relative">
                <input id="newPassword" type={showPwd ? 'text' : 'password'} autoComplete="new-password" required disabled={loading}
                  value={form.newPassword} onChange={e => setForm({ ...form, newPassword: e.target.value })}
                  placeholder="6-20 个字符" className={inputCls + ' pr-10'} />
                <button type="button" tabIndex={-1} onClick={() => setShowPwd(!showPwd)}
                  className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-600">
                  {showPwd ? (
                    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94" /><path d="M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19" /><line x1="1" y1="1" x2="23" y2="23" /></svg>
                  ) : (
                    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z" /><circle cx="12" cy="12" r="3" /></svg>
                  )}
                </button>
              </div>
            </div>

            <div className="mb-6">
              <label htmlFor="confirmPassword" className="block text-sm font-medium text-gray-700 mb-1.5">确认密码</label>
              <input id="confirmPassword" type="password" autoComplete="new-password" required disabled={loading}
                value={form.confirmPassword} onChange={e => setForm({ ...form, confirmPassword: e.target.value })}
                placeholder="请再次输入新密码" className={inputCls} />
            </div>

            <button type="submit" disabled={loading || !canSubmit}
              className="w-full rounded-lg px-4 py-2.5 text-sm font-semibold text-white shadow-sm hover:opacity-90 focus:outline-none focus:ring-2 focus:ring-blue-500/40 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed transition"
              style={{ background: 'linear-gradient(135deg, #3b82f6, #6366f1)' }}>
              {loading ? (
                <span className="inline-flex items-center gap-2">
                  <svg width="16" height="16" viewBox="0 0 24 24" className="animate-spin"><circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4" fill="none" /><path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z" /></svg>
                  重置中...
                </span>
              ) : '重置密码'}
            </button>
          </form>

          <p className="mt-7 text-center text-sm text-gray-500">
            想起密码了？<Link to="/login" className="ml-1 font-semibold text-blue-600 hover:text-blue-700">返回登录</Link>
          </p>
        </div>
      </div>
    </div>
  )
}
