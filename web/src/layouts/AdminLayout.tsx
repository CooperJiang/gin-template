import { useState } from 'react'
import { Link, Outlet, useLocation } from 'react-router-dom'
import { useAuth } from '@/auth/useAuth'
import { getStoredTheme, setTheme, type Theme } from '@/lib/theme'
import {
  LayoutDashboard,
  User,
  Settings,
  LogOut,
  Menu,
  X,
  ChevronDown,
  Sun,
  Moon,
  Monitor,
} from 'lucide-react'

const navItems = [
  { path: '/', label: '仪表盘', icon: LayoutDashboard },
  { path: '/profile', label: '个人资料', icon: User },
  { path: '/settings', label: '系统设置', icon: Settings },
]

const themeOptions: { value: Theme; icon: typeof Sun; label: string }[] = [
  { value: 'light', icon: Sun, label: '浅色' },
  { value: 'dark', icon: Moon, label: '深色' },
  { value: 'system', icon: Monitor, label: '跟随系统' },
]

export default function AdminLayout() {
  const { user, logout } = useAuth()
  const location = useLocation()
  const [sidebarOpen, setSidebarOpen] = useState(false)
  const [userMenuOpen, setUserMenuOpen] = useState(false)
  const [currentTheme, setCurrentTheme] = useState<Theme>(getStoredTheme)

  const isActive = (path: string) => {
    if (path === '/') return location.pathname === '/'
    return location.pathname.startsWith(path)
  }

  const cycleTheme = () => {
    const order: Theme[] = ['light', 'dark', 'system']
    const next = order[(order.indexOf(currentTheme) + 1) % order.length]
    setTheme(next)
    setCurrentTheme(next)
  }

  const ThemeIcon = themeOptions.find((o) => o.value === currentTheme)?.icon ?? Sun
  const themeLabel = themeOptions.find((o) => o.value === currentTheme)?.label ?? '浅色'

  return (
    <div className="min-h-screen bg-nb-bg transition-colors">
      {/* Mobile sidebar overlay */}
      {sidebarOpen && (
        <div
          className="fixed inset-0 z-40 bg-black/40 lg:hidden"
          onClick={() => setSidebarOpen(false)}
        />
      )}

      {/* Header — full width, pinned top */}
      <header className="sticky top-0 z-50 h-16 bg-nb-surface border-b-nb border-nb-border flex items-center px-6 transition-colors">
        {/* Logo area — same width as sidebar */}
        <div className="flex items-center w-64 shrink-0">
          <button
            className="lg:hidden p-2 mr-3 rounded-[var(--nb-radius-sm)] border-nb border-nb-border hover:bg-nb-bg-soft shadow-nb-sm nb-interactive"
            onClick={() => setSidebarOpen(!sidebarOpen)}
          >
            {sidebarOpen ? <X className="w-5 h-5 text-nb-text" /> : <Menu className="w-5 h-5 text-nb-text" />}
          </button>
          <Link to="/" className="flex items-center gap-2.5">
            <div className="w-8 h-8 rounded-[var(--nb-radius-sm)] flex items-center justify-center bg-nb-primary border-nb border-nb-border shadow-nb-sm">
              <svg
                width="16"
                height="16"
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
            <span className="font-bold text-nb-text text-sm">Gin Template</span>
          </Link>
        </div>

        <div className="flex-1" />

        {/* Theme toggle */}
        <button
          onClick={cycleTheme}
          title={themeLabel}
          className="p-2 rounded-[var(--nb-radius-sm)] border-nb border-nb-border text-nb-text hover:bg-nb-bg-soft shadow-nb-sm nb-interactive mr-3"
        >
          <ThemeIcon className="w-5 h-5" />
        </button>

        {/* User menu */}
        <div className="relative">
          <button
            onClick={() => setUserMenuOpen(!userMenuOpen)}
            className="flex items-center gap-2 px-2 py-2 rounded-[var(--nb-radius-sm)] border-nb border-nb-border hover:bg-nb-bg-soft shadow-nb-sm nb-interactive"
          >
            <div className="w-5 h-5 rounded-[var(--nb-radius-sm)] bg-nb-primary border-nb border-nb-border flex items-center justify-center">
              <span className="text-xs font-bold text-[var(--nb-primary-text)]">
                {user?.username?.charAt(0).toUpperCase() || 'U'}
              </span>
            </div>
            <span className="text-sm font-bold text-nb-text hidden sm:block">
              {user?.username || '用户'}
            </span>
            <ChevronDown className="w-4 h-4 text-nb-text-muted" />
          </button>

          {userMenuOpen && (
            <>
              <div className="fixed inset-0 z-40" onClick={() => setUserMenuOpen(false)} />
              <div className="absolute right-0 mt-2 w-48 bg-nb-surface rounded-[var(--nb-radius)] border-nb border-nb-border shadow-nb py-1 z-50">
                <div className="px-4 py-2 border-b-nb border-nb-border">
                  <p className="text-sm font-bold text-nb-text">
                    {user?.username}
                  </p>
                  <p className="text-xs text-nb-text-muted truncate">
                    {user?.email}
                  </p>
                </div>
                <Link
                  to="/profile"
                  onClick={() => setUserMenuOpen(false)}
                  className="flex items-center gap-2 px-4 py-2 text-sm font-medium text-nb-text hover:bg-nb-bg-soft"
                >
                  <User className="w-4 h-4" />
                  个人资料
                </Link>
                <Link
                  to="/settings"
                  onClick={() => setUserMenuOpen(false)}
                  className="flex items-center gap-2 px-4 py-2 text-sm font-medium text-nb-text hover:bg-nb-bg-soft"
                >
                  <Settings className="w-4 h-4" />
                  系统设置
                </Link>
                <div className="border-t-nb border-nb-border mt-1">
                  <button
                    onClick={() => {
                      setUserMenuOpen(false)
                      logout()
                    }}
                    className="flex items-center gap-2 px-4 py-2 text-sm font-bold text-nb-red hover:bg-red-50 dark:hover:bg-red-900/20 w-full text-left"
                  >
                    <LogOut className="w-4 h-4" />
                    退出登录
                  </button>
                </div>
              </div>
            </>
          )}
        </div>
      </header>

      <div className="flex">
        {/* Sidebar — below header */}
        <aside
          className={`fixed top-16 left-0 z-30 h-[calc(100vh-4rem)] w-64 bg-nb-surface border-r-nb border-nb-border transform transition-transform duration-200 ease-in-out lg:translate-x-0 ${
            sidebarOpen ? 'translate-x-0' : '-translate-x-full'
          }`}
        >
          <nav className="p-4 space-y-2">
            {navItems.map((item) => {
              const Icon = item.icon
              const active = isActive(item.path)
              return (
                <Link
                  key={item.path}
                  to={item.path}
                  onClick={() => setSidebarOpen(false)}
                  className={`flex items-center gap-3 px-3 py-2.5 rounded-[var(--nb-radius)] text-sm font-bold transition-all border-nb ${
                    active
                      ? 'bg-nb-primary text-[var(--nb-primary-text)] border-nb-border shadow-nb-sm'
                      : 'text-nb-text-secondary border-transparent hover:bg-nb-bg-soft hover:border-nb-border hover:shadow-nb-sm'
                  }`}
                >
                  <Icon className="w-5 h-5" />
                  {item.label}
                </Link>
              )
            })}
          </nav>
        </aside>

        {/* Main content */}
        <main className="flex-1 lg:ml-64 p-6">
          <Outlet />
        </main>
      </div>
    </div>
  )
}
