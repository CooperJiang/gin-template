import { BrowserRouter, Routes, Route, Navigate, Outlet } from 'react-router-dom'
import { getIsAuthenticated } from '@/auth/store'
import AdminLayout from '@/layouts/AdminLayout'
import Login from '@/auth/pages/Login'
import Register from '@/auth/pages/Register'
import ForgotPassword from '@/auth/pages/ForgotPassword'
import Dashboard from '@/pages/Dashboard'
import Profile from '@/pages/Profile'
import Settings from '@/pages/Settings'
import NotFound from '@/pages/NotFound'

function RequireAuth() {
  if (!getIsAuthenticated()) {
    const redirect = window.location.pathname + window.location.search
    return <Navigate to={`/login?redirect=${encodeURIComponent(redirect)}`} replace />
  }
  return <Outlet />
}

function GuestOnly() {
  if (getIsAuthenticated()) {
    return <Navigate to="/" replace />
  }
  return <Outlet />
}

export default function AppRouter() {
  return (
    <BrowserRouter>
      <Routes>
        {/* Guest-only routes */}
        <Route element={<GuestOnly />}>
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route path="/forgot-password" element={<ForgotPassword />} />
        </Route>

        {/* Authenticated routes */}
        <Route element={<RequireAuth />}>
          <Route element={<AdminLayout />}>
            <Route index element={<Dashboard />} />
            <Route path="/profile" element={<Profile />} />
            <Route path="/settings" element={<Settings />} />
            <Route path="*" element={<NotFound />} />
          </Route>
        </Route>
      </Routes>
    </BrowserRouter>
  )
}
