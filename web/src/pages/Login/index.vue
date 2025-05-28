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

      <!-- 错误提示 -->
      <div v-if="error" class="mt-4 p-3 bg-red-100 border border-red-400 text-red-700 rounded">
        {{ error }}
      </div>

      <!-- 成功提示 -->
      <div
        v-if="success"
        class="mt-4 p-3 bg-green-100 border border-green-400 text-green-700 rounded"
      >
        {{ success }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuth } from '../../hooks/user/useAuth'
import type { LoginRequest } from '../../types'

const router = useRouter()
const { login, loading: authLoading } = useAuth()

const form = ref<LoginRequest & { remember: boolean }>({
  account: '',
  password: '',
  remember: false,
})

const error = ref('')
const success = ref('')

// 记住我功能相关
const REMEMBER_KEY = 'login_remember'
const REMEMBER_EXPIRE_DAYS = 30

// 保存记住我信息
const saveRememberInfo = () => {
  if (form.value.remember) {
    const rememberData = {
      account: form.value.account,
      password: form.value.password,
      expire: Date.now() + (REMEMBER_EXPIRE_DAYS * 24 * 60 * 60 * 1000) // 30天后过期
    }
    localStorage.setItem(REMEMBER_KEY, JSON.stringify(rememberData))
  } else {
    localStorage.removeItem(REMEMBER_KEY)
  }
}

// 恢复记住我信息
const loadRememberInfo = () => {
  try {
    const rememberData = localStorage.getItem(REMEMBER_KEY)
    if (rememberData) {
      const data = JSON.parse(rememberData)
      // 检查是否过期
      if (data.expire && Date.now() < data.expire) {
        form.value.account = data.account || ''
        form.value.password = data.password || ''
        form.value.remember = true
      } else {
        // 过期则删除
        localStorage.removeItem(REMEMBER_KEY)
      }
    }
  } catch (error) {
    console.error('恢复记住我信息失败:', error)
    localStorage.removeItem(REMEMBER_KEY)
  }
}

const handleLogin = async () => {
  try {
    error.value = ''
    success.value = ''

    const loginData = {
      account: form.value.account,
      password: form.value.password
    }

    console.log('发送登录请求:', loginData)

    const response = await login(loginData)

    // 保存记住我信息
    saveRememberInfo()

    success.value = '登录成功！'

    // 登录成功后跳转到dashboard
    setTimeout(() => {
      router.push('/dashboard')
    }, 1000)

  } catch (err: any) {
    console.error('登录错误详情:', err)
    error.value = err.message || '登录失败，请检查账户名和密码'
  }
}

// 页面加载时恢复记住我信息
onMounted(() => {
  loadRememberInfo()
})
</script>
