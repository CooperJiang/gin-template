import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { authApi } from '../../api/auth'
import { useLocalStorage } from '../common'
import type {
  LoginRequest,
  RegisterRequest,
  SendCodeRequest,
  ResetPasswordRequest,
  User
} from '../../types'

export function useAuth() {
  const router = useRouter()
  const loading = ref(false)
  const error = ref('')

  // 使用localStorage存储用户信息和token
  const [token, setToken, removeToken] = useLocalStorage<string | null>('auth_token', null)
  const [user, setUser, removeUser] = useLocalStorage<User | null>('auth_user', null)

  // 计算属性
  const isAuthenticated = computed(() => !!token.value && !!user.value)

  // 登录
  const login = async (credentials: LoginRequest) => {
    try {
      loading.value = true
      error.value = ''

      const response = await authApi.login(credentials)

      setToken(response.token)
      setUser(response.user)

      return response
    } catch (err: any) {
      error.value = err.message || '登录失败'
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
    } catch (err: any) {
      error.value = err.message || '注册失败'
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
    } catch (err: any) {
      error.value = err.message || '发送验证码失败'
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
    } catch (err: any) {
      error.value = err.message || '发送验证码失败'
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
    } catch (err: any) {
      error.value = err.message || '重置密码失败'
      throw err
    } finally {
      loading.value = false
    }
  }

  // 登出
  const logout = () => {
    removeToken()
    removeUser()
    router.push('/login')
  }

  // 获取用户信息
  const getUserInfo = async () => {
    try {
      loading.value = true
      const userInfo = await authApi.getUserInfo()
      setUser(userInfo)
      return userInfo
    } catch (err: any) {
      error.value = err.message || '获取用户信息失败'
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
    getUserInfo
  }
}
