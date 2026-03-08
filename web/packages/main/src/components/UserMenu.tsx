import { useState, useRef, useEffect, useCallback } from 'react'
import { useAuth } from '@app/auth'
import type { User } from '@app/shared'

function getInitial(user: User | null): string {
  if (!user) return '?'
  return (user.username || user.email || '?').charAt(0).toUpperCase()
}

export function UserMenu() {
  const { user, logout } = useAuth()
  const [open, setOpen] = useState(false)
  const ref = useRef<HTMLDivElement>(null)

  const [isMobile, setIsMobile] = useState(() => window.innerWidth < 640)

  const close = useCallback(() => setOpen(false), [])

  useEffect(() => {
    if (!open) return
    const handler = (e: MouseEvent) => {
      if (ref.current && !ref.current.contains(e.target as Node)) close()
    }
    document.addEventListener('mousedown', handler)
    return () => document.removeEventListener('mousedown', handler)
  }, [open, close])

  useEffect(() => {
    const onResize = () => setIsMobile(window.innerWidth < 640)
    window.addEventListener('resize', onResize)
    return () => window.removeEventListener('resize', onResize)
  }, [])

  return (
    <div className="relative" ref={ref}>
      <button
        onClick={() => setOpen(!open)}
        className="flex items-center gap-2 py-1 px-1 pr-2 rounded-full border-none cursor-pointer bg-transparent hover:bg-black/[.04] transition-colors"
      >
        <div
          className="w-7 h-7 rounded-full flex items-center justify-center text-white text-xs font-semibold"
          style={{ background: 'linear-gradient(135deg, #3b82f6, #6366f1)' }}
        >
          {getInitial(user)}
        </div>
        {!isMobile && (
          <span className="text-[13px] font-medium text-gray-700 max-w-[80px] truncate">
            {user?.username || '用户'}
          </span>
        )}
        <svg
          width="12" height="12" viewBox="0 0 12 12" fill="none"
          className="text-gray-400 transition-transform"
          style={{ transform: open ? 'rotate(180deg)' : 'rotate(0deg)' }}
        >
          <path d="M2.5 4.5L6 8L9.5 4.5" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round" />
        </svg>
      </button>

      {open && (
        <div
          className="absolute right-0 mt-1.5 w-56 bg-white rounded-xl overflow-hidden z-50"
          style={{
            boxShadow: '0 4px 24px rgba(0,0,0,.08), 0 1px 2px rgba(0,0,0,.06)',
            border: '1px solid rgba(0,0,0,.06)',
            animation: 'menuIn .15s ease-out',
          }}
        >
          <style>{`@keyframes menuIn { from { opacity:0; transform:translateY(-4px) scale(.98); } to { opacity:1; transform:translateY(0) scale(1); } }`}</style>

          <div className="px-3.5 py-3 border-b border-gray-100" style={{ background: 'linear-gradient(135deg, #fafbfc, #f3f4f6)' }}>
            <div className="flex items-center gap-2.5">
              <div
                className="w-9 h-9 rounded-full flex items-center justify-center text-white text-sm font-semibold shrink-0"
                style={{ background: 'linear-gradient(135deg, #3b82f6, #6366f1)' }}
              >
                {getInitial(user)}
              </div>
              <div className="min-w-0">
                <div className="text-[13px] font-semibold text-gray-900 truncate">{user?.username}</div>
                <div className="text-[11px] text-gray-500 truncate">{user?.email}</div>
              </div>
            </div>
          </div>

          <div className="p-1">
            <a
              href="/app/profile"
              className="flex items-center gap-2.5 px-2.5 py-2 rounded-lg text-[13px] text-gray-700 no-underline hover:bg-gray-50 transition-colors"
            >
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="#9ca3af" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round">
                <path d="M20 21v-2a4 4 0 00-4-4H8a4 4 0 00-4 4v2" />
                <circle cx="12" cy="7" r="4" />
              </svg>
              个人资料
            </a>
            <a
              href="/app/settings"
              className="flex items-center gap-2.5 px-2.5 py-2 rounded-lg text-[13px] text-gray-700 no-underline hover:bg-gray-50 transition-colors"
            >
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="#9ca3af" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round">
                <circle cx="12" cy="12" r="3" />
                <path d="M19.4 15a1.65 1.65 0 00.33 1.82l.06.06a2 2 0 010 2.83 2 2 0 01-2.83 0l-.06-.06a1.65 1.65 0 00-1.82-.33 1.65 1.65 0 00-1 1.51V21a2 2 0 01-4 0v-.09A1.65 1.65 0 009 19.4a1.65 1.65 0 00-1.82.33l-.06.06a2 2 0 01-2.83-2.83l.06-.06A1.65 1.65 0 004.68 15a1.65 1.65 0 00-1.51-1H3a2 2 0 010-4h.09A1.65 1.65 0 004.6 9a1.65 1.65 0 00-.33-1.82l-.06-.06a2 2 0 012.83-2.83l.06.06A1.65 1.65 0 009 4.68a1.65 1.65 0 001-1.51V3a2 2 0 014 0v.09a1.65 1.65 0 001 1.51 1.65 1.65 0 001.82-.33l.06-.06a2 2 0 012.83 2.83l-.06.06A1.65 1.65 0 0019.4 9a1.65 1.65 0 001.51 1H21a2 2 0 010 4h-.09a1.65 1.65 0 00-1.51 1z" />
              </svg>
              设置
            </a>
          </div>

          <div className="p-1 border-t border-gray-100">
            <button
              onClick={() => logout()}
              className="flex items-center gap-2.5 w-full px-2.5 py-2 rounded-lg text-[13px] text-red-600 border-none bg-transparent cursor-pointer hover:bg-red-50 transition-colors text-left"
            >
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round">
                <path d="M9 21H5a2 2 0 01-2-2V5a2 2 0 012-2h4" />
                <polyline points="16 17 21 12 16 7" />
                <line x1="21" y1="12" x2="9" y2="12" />
              </svg>
              退出登录
            </button>
          </div>
        </div>
      )}
    </div>
  )
}
