<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full space-y-8">
      <div>
        <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">
          登录您的账户
        </h2>
        <p class="mt-2 text-center text-sm text-gray-600">
          或者
          <router-link to="/register" class="font-medium text-blue-600 hover:text-blue-500">
            创建新账户
          </router-link>
        </p>
      </div>

      <!-- 错误提示 -->
      <div v-if="error" class="bg-red-50 border border-red-200 rounded-md p-4">
        <div class="flex">
          <div class="flex-shrink-0">
            <ExclamationCircleIcon class="h-5 w-5 text-red-400" />
          </div>
          <div class="ml-3">
            <p class="text-red-800">{{ error }}</p>
          </div>
        </div>
      </div>

      <form class="mt-8 space-y-6" @submit.prevent="handleSubmit">
        <div class="rounded-md shadow-sm -space-y-px">
          <div>
            <label for="account" class="sr-only">账户</label>
            <input
              id="account"
              v-model="form.account"
              name="account"
              type="text"
              required
              :disabled="loading"
              class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 focus:z-10 sm:text-sm disabled:opacity-50 disabled:cursor-not-allowed"
              placeholder="用户名或邮箱"
            />
          </div>
          <div>
            <label for="password" class="sr-only">密码</label>
            <input
              id="password"
              v-model="form.password"
              name="password"
              type="password"
              required
              :disabled="loading"
              class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 focus:z-10 sm:text-sm disabled:opacity-50 disabled:cursor-not-allowed"
              placeholder="密码"
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
            <label for="remember-me" class="ml-2 block text-sm text-gray-900">
              记住我
            </label>
          </div>

          <div class="text-sm">
            <router-link to="/forgot-password" class="font-medium text-blue-600 hover:text-blue-500">
              忘记密码？
            </router-link>
          </div>
        </div>

        <div>
          <button
            type="submit"
            :disabled="loading || !isFormValid"
            class="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
          >
            <span v-if="loading" class="absolute left-0 inset-y-0 flex items-center pl-3">
              <!-- Loading spinner -->
              <svg class="animate-spin h-5 w-5 text-blue-300" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
            </span>
            {{ loading ? '登录中...' : '登录' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onBeforeUnmount } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { ExclamationCircleIcon } from '@heroicons/vue/24/outline'
import { useAuth } from '@/hooks/user/useAuth'
import { useMessage } from '@/composables/useMessage'
import type { LoginRequest } from '@/types'

const router = useRouter()
const route = useRoute()
const { login, loading, error: authError } = useAuth()
const { success, error: showError } = useMessage()

// 表单数据
const form = ref({
  account: '',
  password: '',
  remember: false
})

// 错误信息
const error = ref('')

// 表单验证
const isFormValid = computed(() => {
  return form.value.account.trim() !== '' && form.value.password.trim() !== ''
})

// 处理登录提交
const handleSubmit = async () => {
  try {
    error.value = ''

    // 表单验证
    if (!form.value.account.trim()) {
      error.value = '请输入用户名或邮箱'
      return
    }

    if (!form.value.password.trim()) {
      error.value = '请输入密码'
      return
    }

    if (form.value.password.length < 6) {
      error.value = '密码长度不能少于6位'
      return
    }

    // 构造登录请求数据
    const loginData: LoginRequest = {
      account: form.value.account.trim(),
      password: form.value.password.trim()
    }

    // 调用登录API
    await login(loginData)

    // 登录成功提示
    success('登录成功！')

    // 获取重定向路径
    const redirectPath = (route.query.redirect as string) || '/'

    // 延迟跳转，让用户看到成功提示
    setTimeout(() => {
      router.push(redirectPath)
    }, 500)

  } catch (err: unknown) {
    // 处理登录错误
    const errorMessage = err instanceof Error ? err.message : '登录失败，请重试'
    error.value = errorMessage
    showError(errorMessage)
  }
}

// 监听认证错误变化
const unwatchAuthError = ref<(() => void) | null>(null)

onMounted(() => {
  // 清除之前的错误
  error.value = ''

  // 如果URL中有错误信息，显示出来
  if (route.query.error) {
    error.value = route.query.error as string
  }

  // 监听认证错误
  unwatchAuthError.value = () => {
    if (authError.value) {
      error.value = authError.value
    }
  }
})

// 清理函数
onBeforeUnmount(() => {
  if (unwatchAuthError.value) {
    unwatchAuthError.value()
  }
})
</script>
