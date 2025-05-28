import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { authApi } from '@/api/auth'
import type { User, LoginRequest } from '@/types'

export const useAuthStore = defineStore('auth', () => {
  // 状态
  const token = ref<string | null>(localStorage.getItem('auth_token'))
  const user = ref<User | null>(JSON.parse(localStorage.getItem('auth_user') || 'null'))

  // 计算属性
  const isLoggedIn = computed(() => !!token.value && !!user.value)

  // 登录
  const login = async (credentials: LoginRequest) => {
    try {
      const response = await authApi.login(credentials)

      // 存储token和用户信息
      token.value = response.token
      user.value = response.user

      // 持久化存储
      localStorage.setItem('auth_token', response.token)
      localStorage.setItem('auth_user', JSON.stringify(response.user))

      return response
    } catch (error) {
      throw error
    }
  }

  // 获取当前用户信息
  const getCurrentUser = async () => {
    try {
      const response = await authApi.getUserInfo()
      user.value = response

      // 更新本地存储
      localStorage.setItem('auth_user', JSON.stringify(response))

      return response
    } catch (error) {
      // 如果获取用户信息失败，清除认证状态
      await logout()
      throw error
    }
  }

  // 初始化认证状态
  const initAuth = async () => {
    // 如果没有token，直接返回
    if (!token.value) {
      return
    }

    // 如果有用户信息，不需要重新获取
    if (user.value) {
      return
    }

    // 获取用户信息
    try {
      await getCurrentUser()
    } catch (error) {
      console.error('Failed to initialize auth:', error)
      // 初始化失败，清除认证状态
      await logout()
    }
  }

  // 登出
  const logout = async () => {
    token.value = null
    user.value = null

    // 清除本地存储
    localStorage.removeItem('auth_token')
    localStorage.removeItem('auth_user')
  }

  return {
    // 状态
    token,
    user,

    // 计算属性
    isLoggedIn,

    // 方法
    login,
    getCurrentUser,
    initAuth,
    logout
  }
})

