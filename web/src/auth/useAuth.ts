import { useState, useCallback, useSyncExternalStore } from 'react'
import {
  subscribe,
  getSnapshot,
  getToken,
  getUser,
  getIsAuthenticated,
  setToken,
  setUser,
  clearAuth,
} from './store'
import { authApi } from './api'
import type { LoginRequest, RegisterRequest, ResetPasswordRequest } from '@/types/auth'

export function useAuth() {
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  useSyncExternalStore(subscribe, getSnapshot)

  const token = getToken()
  const user = getUser()
  const isAuthenticated = getIsAuthenticated()

  const login = useCallback(async (credentials: LoginRequest) => {
    try {
      setLoading(true)
      setError('')
      const response = await authApi.login(credentials)
      setToken(response.token)
      setUser(response.user)
      return response
    } catch (err: unknown) {
      const msg = err instanceof Error ? err.message : '登录失败'
      setError(msg)
      throw err
    } finally {
      setLoading(false)
    }
  }, [])

  const register = useCallback(async (userData: RegisterRequest) => {
    try {
      setLoading(true)
      setError('')
      return await authApi.register(userData)
    } catch (err: unknown) {
      const msg = err instanceof Error ? err.message : '注册失败'
      setError(msg)
      throw err
    } finally {
      setLoading(false)
    }
  }, [])

  const sendRegistrationCode = useCallback(async (email: string) => {
    try {
      setLoading(true)
      setError('')
      return await authApi.sendRegistrationCode({ email })
    } catch (err: unknown) {
      const msg = err instanceof Error ? err.message : '发送验证码失败'
      setError(msg)
      throw err
    } finally {
      setLoading(false)
    }
  }, [])

  const sendResetPasswordCode = useCallback(async (email: string) => {
    try {
      setLoading(true)
      setError('')
      return await authApi.sendResetPasswordCode({ email })
    } catch (err: unknown) {
      const msg = err instanceof Error ? err.message : '发送验证码失败'
      setError(msg)
      throw err
    } finally {
      setLoading(false)
    }
  }, [])

  const resetPassword = useCallback(async (data: ResetPasswordRequest) => {
    try {
      setLoading(true)
      setError('')
      return await authApi.resetPassword(data)
    } catch (err: unknown) {
      const msg = err instanceof Error ? err.message : '重置密码失败'
      setError(msg)
      throw err
    } finally {
      setLoading(false)
    }
  }, [])

  const getUserInfo = useCallback(async () => {
    try {
      setLoading(true)
      const response = await authApi.getUserInfo()
      setUser(response)
      return response
    } catch (err: unknown) {
      const msg = err instanceof Error ? err.message : '获取用户信息失败'
      setError(msg)
      throw err
    } finally {
      setLoading(false)
    }
  }, [])

  const logout = useCallback((shouldRedirect = true) => {
    clearAuth()
    if (shouldRedirect) {
      window.location.href = '/login'
    }
  }, [])

  return {
    loading,
    error,
    user,
    token,
    isAuthenticated,
    login,
    register,
    sendRegistrationCode,
    sendResetPasswordCode,
    resetPassword,
    getUserInfo,
    logout,
  }
}
