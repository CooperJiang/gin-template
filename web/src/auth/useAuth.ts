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
import type {
  LoginRequest,
  RegisterRequest,
  ResetPasswordRequest,
  ChangePasswordRequest,
  UpdateProfileRequest,
} from '@/types/auth'

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

  const changePassword = useCallback(async (data: ChangePasswordRequest) => {
    try {
      setLoading(true)
      setError('')
      return await authApi.changePassword(data)
    } catch (err: unknown) {
      const msg = err instanceof Error ? err.message : '修改密码失败'
      setError(msg)
      throw err
    } finally {
      setLoading(false)
    }
  }, [])

  const updateProfile = useCallback(async (data: UpdateProfileRequest) => {
    try {
      setLoading(true)
      setError('')
      const updated = await authApi.updateProfile(data)
      setUser(updated)
      return updated
    } catch (err: unknown) {
      const msg = err instanceof Error ? err.message : '更新资料失败'
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

  const logout = useCallback(() => {
    clearAuth()
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
    changePassword,
    updateProfile,
    getUserInfo,
    logout,
  }
}
