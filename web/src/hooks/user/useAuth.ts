import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { authApi } from '../../api/auth'
import { useAuthStorage } from '../common'
import { STORAGE_KEYS } from '../../utils/storage'
import type {
  LoginRequest,
  RegisterRequest,
  ResetPasswordRequest,
  User,
} from '../../types'

export function useAuth() {
  const router = useRouter()
  const loading = ref(false)
  const error = ref('')

  // 使用安全存储存储用户信息和token (7天过期)
  const [token, setToken, removeToken] = useAuthStorage<string | null>(STORAGE_KEYS.AUTH_TOKEN, null)
  const [user, setUser, removeUser] = useAuthStorage<User | null>(STORAGE_KEYS.AUTH_USER, null)

  // 计算属性
  const isAuthenticated = computed(() => !!token.value && !!user.value)

  // 登录
  const login = async (credentials: LoginRequest) => {
    try {
      loading.value = true
      error.value = ''

      const response = await authApi.login(credentials)

      // 从响应中提取数据
      const responseData = (response as any).data || response

      setToken(responseData.token)
      setUser(responseData.user)

      return response
    } catch (err: unknown) {
      const errorMessage = err instanceof Error ? err.message : '登录失败'
      error.value = errorMessage
      throw err
    } finally {
      loading.value = false
    }
  }

  // 注册
  const register = async (userData: RegisterRequest) => {
    try {
      loading.value = true
      error.value = ''

      const response = await authApi.register(userData)
      return response
    } catch (err: unknown) {
      const errorMessage = err instanceof Error ? err.message : '注册失败'
      error.value = errorMessage
      throw err
    } finally {
      loading.value = false
    }
  }

  // 发送注册验证码
  const sendRegistrationCode = async (email: string) => {
    try {
      loading.value = true
      error.value = ''

      const response = await authApi.sendRegistrationCode({ email })
      return response
    } catch (err: unknown) {
      const errorMessage = err instanceof Error ? err.message : '发送验证码失败'
      error.value = errorMessage
      throw err
    } finally {
      loading.value = false
    }
  }

  // 发送重置密码验证码
  const sendResetPasswordCode = async (email: string) => {
    try {
      loading.value = true
      error.value = ''

      const response = await authApi.sendResetPasswordCode({ email })
      return response
    } catch (err: unknown) {
      const errorMessage = err instanceof Error ? err.message : '发送验证码失败'
      error.value = errorMessage
      throw err
    } finally {
      loading.value = false
    }
  }

  // 重置密码
  const resetPassword = async (data: ResetPasswordRequest) => {
    try {
      loading.value = true
      error.value = ''

      const response = await authApi.resetPassword(data)
      return response
    } catch (err: unknown) {
      const errorMessage = err instanceof Error ? err.message : '重置密码失败'
      error.value = errorMessage
      throw err
    } finally {
      loading.value = false
    }
  }

  // 登出
  const logout = (shouldRedirect: boolean = true) => {
    removeToken()
    removeUser()
    if (shouldRedirect) {
      router.push('/login')
    }
  }

  // 获取用户信息
  const getUserInfo = async () => {
    try {
      loading.value = true
      const response = await authApi.getUserInfo()
      // 从响应中提取数据
      const userInfo = (response as any).data || response
      setUser(userInfo)
      return userInfo
    } catch (err: unknown) {
      const errorMessage = err instanceof Error ? err.message : '获取用户信息失败'
      error.value = errorMessage
      throw err
    } finally {
      loading.value = false
    }
  }

  return {
    // 状态
    loading,
    error,
    user,
    token,
    isAuthenticated,

    // 方法
    login,
    register,
    sendRegistrationCode,
    sendResetPasswordCode,
    resetPassword,
    logout,
    getUserInfo,
  }
}
