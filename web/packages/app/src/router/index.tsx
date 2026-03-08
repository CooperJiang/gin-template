import { lazy, Suspense } from 'react'
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom'
import { qiankunWindow } from 'vite-plugin-qiankun/dist/helper'
import { useAuth } from '@app/auth'

const Home = lazy(() => import('@/pages/Home'))
const Profile = lazy(() => import('@/pages/Profile'))
const Settings = lazy(() => import('@/pages/Settings'))
const About = lazy(() => import('@/pages/About'))
const Docs = lazy(() => import('@/pages/Docs'))
const NotFound = lazy(() => import('@/pages/NotFound'))

// 独立运行时加载认证页面
const Login = lazy(() => import('@app/auth/pages/Login'))
const Register = lazy(() => import('@app/auth/pages/Register'))
const ForgotPassword = lazy(() => import('@app/auth/pages/ForgotPassword'))

const isQiankun = qiankunWindow.__POWERED_BY_QIANKUN__
const base = isQiankun ? '/app' : '/'

function RequireAuth({ children }: { children: React.ReactNode }) {
  const { isAuthenticated } = useAuth()
  if (!isAuthenticated) {
    window.location.href = '/login?redirect=' + encodeURIComponent(window.location.pathname)
    return null
  }
  return <>{children}</>
}

/** 独立运行时，已登录用户不可访问认证页面 */
function GuestOnly({ children }: { children: React.ReactNode }) {
  const { isAuthenticated } = useAuth()
  if (isAuthenticated) return <Navigate to="/" replace />
  return <>{children}</>
}

const Loading = () => (
  <div className="min-h-screen flex items-center justify-center">
    <div className="animate-spin h-8 w-8 border-4 border-blue-600 border-t-transparent rounded-full" />
  </div>
)

export default function AppRouter() {
  return (
    <BrowserRouter basename={base}>
      <Suspense fallback={<Loading />}>
        <Routes>
          {/* 独立运行时提供认证页面 */}
          {!isQiankun && (
            <>
              <Route path="/login" element={<GuestOnly><Login /></GuestOnly>} />
              <Route path="/register" element={<GuestOnly><Register /></GuestOnly>} />
              <Route path="/forgot-password" element={<GuestOnly><ForgotPassword /></GuestOnly>} />
            </>
          )}
          <Route path="/" element={<Home />} />
          <Route path="/profile" element={<RequireAuth><Profile /></RequireAuth>} />
          <Route path="/settings" element={<RequireAuth><Settings /></RequireAuth>} />
          <Route path="/about" element={<About />} />
          <Route path="/docs" element={<Docs />} />
          <Route path="*" element={<NotFound />} />
        </Routes>
      </Suspense>
    </BrowserRouter>
  )
}
