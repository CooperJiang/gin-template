import { useEffect, useState, useCallback } from 'react'
import { Link, useLocation } from 'react-router-dom'
import { useSyncExternalStore } from 'react'
import { useAuth } from '@app/auth'
import { UserMenu } from './UserMenu'
import { Loading, ErrorFallback } from './Loading'
import { buttonVariants } from './ui/button'
import { getMicroAppMenuLinks } from '@/micro-app-registry'
import {
  subscribeMicroState,
  getMicroStateSnapshot,
  getMicroLoading,
  getMicroError,
  retryMicroApp,
} from '@/micro-apps'

const SM = 640

export function AppLayout() {
  const { isAuthenticated } = useAuth()
  const location = useLocation()
  const [mobileOpen, setMobileOpen] = useState(false)
  const [scrolled, setScrolled] = useState(false)
  const [isMobile, setIsMobile] = useState(() => window.innerWidth < SM)

  useEffect(() => {
    const onResize = () => setIsMobile(window.innerWidth < SM)
    window.addEventListener('resize', onResize)
    return () => window.removeEventListener('resize', onResize)
  }, [])

  useEffect(() => {
    const onScroll = () => setScrolled(window.scrollY > 2)
    window.addEventListener('scroll', onScroll, { passive: true })
    return () => window.removeEventListener('scroll', onScroll)
  }, [])

  useEffect(() => { setMobileOpen(false) }, [location.pathname, isMobile])

  useSyncExternalStore(subscribeMicroState, getMicroStateSnapshot)
  const loading = getMicroLoading()
  const error = getMicroError()
  const microMenus = getMicroAppMenuLinks()
  const navLinks = location.pathname.startsWith('/app')
    ? [...microMenus, { href: '/app/docs', label: '文档' }, { href: '/app/about', label: '关于' }]
    : microMenus

  const isActive = useCallback((href: string) => {
    if (href === '/app') return location.pathname === '/app' || location.pathname === '/app/'
    return location.pathname.startsWith(href)
  }, [location.pathname])

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Header */}
      <header
        className={`sticky top-0 z-50 transition-shadow duration-300 border-b border-black/[.06] ${
          scrolled ? 'bg-white/[.92] shadow-sm' : 'bg-white/[.78]'
        }`}
        style={{ backdropFilter: 'blur(16px) saturate(1.8)', WebkitBackdropFilter: 'blur(16px) saturate(1.8)' }}
      >
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex items-center justify-between h-14">
            {/* Left: Logo + Nav */}
            <div className="flex items-center gap-10">
              <a href="/app" className="flex items-center gap-2.5 shrink-0 no-underline">
                <div
                  className="w-8 h-8 rounded-lg flex items-center justify-center"
                  style={{ background: 'linear-gradient(135deg, #3b82f6, #6366f1)' }}
                >
                  <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="white" strokeWidth="2.2" strokeLinecap="round" strokeLinejoin="round">
                    <path d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                  </svg>
                </div>
                <span className="text-gray-900" style={{ fontFamily: "'Outfit', sans-serif", fontWeight: 700, fontSize: 17, letterSpacing: '-0.02em' }}>
                  应用平台
                </span>
              </a>

              {/* Desktop Nav */}
              {!isMobile && (
                <nav className="flex items-center gap-1">
                  {navLinks.map((link) => {
                    const active = isActive(link.href)
                    return (
                      <a
                        key={link.href}
                        href={link.href}
                        className={`px-3 py-1.5 text-[13px] font-medium rounded-md no-underline transition-colors ${
                          active ? 'text-gray-900 bg-black/[.04]' : 'text-gray-500 hover:text-gray-700 hover:bg-black/[.03]'
                        }`}
                      >
                        {link.label}
                      </a>
                    )
                  })}
                </nav>
              )}
            </div>

            {/* Right */}
            <div className="flex items-center gap-2">
              {isAuthenticated ? (
                <UserMenu />
              ) : (
                <div className="flex items-center gap-2">
                  <Link
                    to="/login"
                    className={buttonVariants({ variant: 'ghost', size: 'sm', className: 'no-underline' })}
                  >
                    登录
                  </Link>
                  <Link
                    to="/register"
                    className={buttonVariants({ variant: 'default', size: 'sm', className: 'no-underline' })}
                  >
                    注册
                  </Link>
                </div>
              )}

              {isMobile && (
                <button
                  onClick={() => setMobileOpen(!mobileOpen)}
                  className={buttonVariants({
                    variant: 'ghost',
                    size: 'icon',
                    className: 'ml-1 border-none bg-transparent cursor-pointer',
                  })}
                >
                  <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                    {mobileOpen
                      ? <path d="M18 6L6 18M6 6l12 12" />
                      : <><line x1="4" y1="7" x2="20" y2="7" /><line x1="4" y1="12" x2="20" y2="12" /><line x1="4" y1="17" x2="20" y2="17" /></>
                    }
                  </svg>
                </button>
              )}
            </div>
          </div>
        </div>

        {/* Mobile menu */}
        {isMobile && mobileOpen && (
          <div className="border-t border-gray-100 bg-white/95" style={{ backdropFilter: 'blur(16px)' }}>
            <div className="px-4 py-2">
              {navLinks.map((link) => {
                const active = isActive(link.href)
                return (
                  <a
                    key={link.href}
                    href={link.href}
                    onClick={() => setMobileOpen(false)}
                    className={`block px-3 py-2.5 rounded-lg text-sm font-medium no-underline ${
                      active ? 'text-blue-600 bg-blue-50/60' : 'text-gray-600 hover:text-gray-900 hover:bg-gray-50'
                    }`}
                  >
                    {link.label}
                  </a>
                )
              })}
            </div>
          </div>
        )}
      </header>

      {/* Content */}
      <div className="min-h-[calc(100vh-57px)]">
        {loading && <Loading text="正在加载应用..." />}
        {error && <ErrorFallback error={error} onRetry={retryMicroApp} />}
        <main id="subapp-container" style={{ display: loading || error ? 'none' : undefined }} />
      </div>
    </div>
  )
}
