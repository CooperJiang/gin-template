import { lazy, Suspense } from 'react'
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom'
import { useAuth } from '@app/auth'
import { AppLayout } from '@/components/Layout'

const Login = lazy(() => import('@app/auth/pages/Login'))
const Register = lazy(() => import('@app/auth/pages/Register'))
const ForgotPassword = lazy(() => import('@app/auth/pages/ForgotPassword'))

/** 仅未登录用户可访问 */
function GuestOnly({ children }: { children: React.ReactNode }) {
  const { isAuthenticated } = useAuth()
  if (isAuthenticated) return <Navigate to="/app" replace />
  return <>{children}</>
}

const Loading = () => (
  <div className="min-h-screen flex items-center justify-center">
    <div className="animate-spin h-8 w-8 border-4 border-gray-200 border-t-blue-600 rounded-full" />
  </div>
)

export default function AppRouter() {
  return (
    <BrowserRouter>
      <Suspense fallback={<Loading />}>
        <Routes>
          <Route path="/login" element={<GuestOnly><Login /></GuestOnly>} />
          <Route path="/register" element={<GuestOnly><Register /></GuestOnly>} />
          <Route path="/forgot-password" element={<GuestOnly><ForgotPassword /></GuestOnly>} />
          <Route path="/" element={<Navigate to="/app" replace />} />
          <Route path="/*" element={<AppLayout />} />
        </Routes>
      </Suspense>
    </BrowserRouter>
  )
}
