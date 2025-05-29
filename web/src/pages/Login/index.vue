<template>
  <div
    class="w-screen min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8"
  >
    <div class="max-w-md w-full space-y-8">
      <div>
        <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">登录您的账户</h2>
        <p class="mt-2 text-center text-sm text-gray-600">
          或者
          <router-link to="/register" class="font-medium text-blue-600 hover:text-blue-500">
            创建新账户
          </router-link>
        </p>
      </div>

      <form class="mt-8 space-y-6" @submit.prevent="handleLogin">
        <div class="space-y-4">
          <div>
            <label for="account" class="block text-sm font-medium text-gray-700">账户名/邮箱</label>
            <input
              id="account"
              v-model="form.account"
              name="account"
              type="text"
              autocomplete="username"
              required
              class="mt-1 appearance-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
              placeholder="请输入用户名或邮箱地址"
            />
          </div>
          <div>
            <label for="password" class="block text-sm font-medium text-gray-700">密码</label>
            <input
              id="password"
              v-model="form.password"
              name="password"
              type="password"
              autocomplete="current-password"
              required
              class="mt-1 appearance-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
              placeholder="请输入密码"
            />
          </div>
        </div>

        <div class="flex items-center justify-between">
          <div class="flex items-center">
            <input
              id="remember-me"
              v-model="form.remember"
              name="remember-me"
              type="checkbox"
              class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
            />
            <label for="remember-me" class="ml-2 block text-sm text-gray-900"> 记住我 </label>
          </div>

          <div class="text-sm">
            <router-link
              to="/forgot-password"
              class="font-medium text-blue-600 hover:text-blue-500"
            >
              忘记密码？
            </router-link>
          </div>
        </div>

        <div>
          <button
            type="submit"
            :disabled="authLoading"
            class="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            <span v-if="authLoading" class="absolute left-0 inset-y-0 flex items-center pl-3">
              <div
                class="w-5 h-5 border-2 border-white border-t-transparent rounded-full animate-spin"
              ></div>
            </span>
            {{ authLoading ? '登录中...' : '登录' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, defineOptions } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuth } from '../../hooks/user/useAuth'
import { useMessage } from '../../composables/useMessage'
import SecureStorage, { STORAGE_KEYS, EXPIRY_TIME } from '../../utils/storage'
import type { LoginRequest } from '../../types'

const router = useRouter()
const route = useRoute()
const { login, loading: authLoading } = useAuth()
const { success, error } = useMessage()

const form = ref<LoginRequest & { remember: boolean }>({
  account: '',
  password: '',
  remember: false,
})

// 保存记住我信息
const saveRememberInfo = () => {
  if (form.value.remember) {
    const rememberData = {
      account: form.value.account,
      password: form.value.password
    }
    // 使用安全存储，30天过期
    SecureStorage.setItem(STORAGE_KEYS.REMEMBER_LOGIN, rememberData, {
      encrypt: true,
      expiry: 30 * EXPIRY_TIME.DAY
    })
  } else {
    SecureStorage.removeItem(STORAGE_KEYS.REMEMBER_LOGIN)
  }
}

// 恢复记住我信息
const loadRememberInfo = () => {
  try {
    const rememberData = SecureStorage.getItem<{account: string; password: string}>(STORAGE_KEYS.REMEMBER_LOGIN)
    if (rememberData) {
      form.value.account = rememberData.account || ''
      form.value.password = rememberData.password || ''
      form.value.remember = true
    }
  } catch (error) {
    console.error('恢复记住我信息失败:', error)
    SecureStorage.removeItem(STORAGE_KEYS.REMEMBER_LOGIN)
  }
}

const handleLogin = async () => {
  if (!validateForm()) return

  authLoading.value = true
  try {
    const loginData = {
      account: form.value.account,
      password: form.value.password,
    }

    await login(loginData)

    // 保存记住我信息
    saveRememberInfo()

    success('登录成功！正在跳转...', 2000)

    // 登录成功后确保认证状态完全更新再跳转
    await new Promise(resolve => setTimeout(resolve, 100))

    // 登录成功后跳转到dashboard
    const redirectTo = (route.query.redirect as string) || '/dashboard'
    await router.push(redirectTo)
  } catch (err: unknown) {
    console.error('登录错误详情:', err)
    const errorMessage = err instanceof Error ? err.message : '登录失败，请检查账户名和密码'
    error(errorMessage)
  } finally {
    authLoading.value = false
  }
}

// 页面加载时恢复记住我信息
onMounted(() => {
  loadRememberInfo()
})

// 登录处理
const validateForm = (): boolean => {
  if (!form.value.account.trim()) {
    error('请输入账户名或邮箱')
    return false
  }
  if (!form.value.password.trim()) {
    error('请输入密码')
    return false
  }
  return true
}

defineOptions({
  name: 'LoginPage',
})
</script>
