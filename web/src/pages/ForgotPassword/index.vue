<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full space-y-8">
      <div>
        <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">
          重置密码
        </h2>
        <p class="mt-2 text-center text-sm text-gray-600">
          输入您的邮箱地址和验证码，设置新密码
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

      <!-- 成功提示 -->
      <div v-if="successMessage" class="bg-green-50 border border-green-200 rounded-md p-4">
        <div class="flex">
          <div class="flex-shrink-0">
            <CheckCircleIcon class="h-5 w-5 text-green-400" />
          </div>
          <div class="ml-3">
            <p class="text-green-800">{{ successMessage }}</p>
          </div>
        </div>
      </div>

      <form class="mt-8 space-y-6" @submit.prevent="handleSubmit">
        <div class="space-y-4">
          <div>
            <label for="email" class="block text-sm font-medium text-gray-700">邮箱地址</label>
            <input
              id="email"
              v-model="form.email"
              type="email"
              required
              :disabled="loading"
              class="mt-1 appearance-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm disabled:opacity-50 disabled:cursor-not-allowed"
              placeholder="请输入您的邮箱"
            />
          </div>

          <div>
            <label for="code" class="block text-sm font-medium text-gray-700">验证码</label>
            <div class="mt-1 flex space-x-2">
              <input
                id="code"
                v-model="form.code"
                type="text"
                required
                :disabled="loading"
                class="flex-1 appearance-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm disabled:opacity-50 disabled:cursor-not-allowed"
                placeholder="请输入验证码"
              />
              <button
                type="button"
                :disabled="sendCodeLoading || countdown > 0 || !isEmailValid"
                @click="handleSendCode"
                class="px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50 disabled:cursor-not-allowed"
              >
                <span v-if="sendCodeLoading">发送中...</span>
                <span v-else-if="countdown > 0">{{ countdown }}s</span>
                <span v-else>发送验证码</span>
              </button>
            </div>
          </div>

          <div>
            <label for="newPassword" class="block text-sm font-medium text-gray-700">新密码</label>
            <input
              id="newPassword"
              v-model="form.newPassword"
              type="password"
              required
              :disabled="loading"
              class="mt-1 appearance-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm disabled:opacity-50 disabled:cursor-not-allowed"
              placeholder="请输入新密码"
            />
          </div>

          <div>
            <label for="confirmPassword" class="block text-sm font-medium text-gray-700">确认新密码</label>
            <input
              id="confirmPassword"
              v-model="form.confirmPassword"
              type="password"
              required
              :disabled="loading"
              class="mt-1 appearance-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm disabled:opacity-50 disabled:cursor-not-allowed"
              placeholder="请再次输入新密码"
            />
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
            {{ loading ? '重置中...' : '重置密码' }}
          </button>
        </div>

        <div class="text-center">
          <router-link to="/login" class="font-medium text-blue-600 hover:text-blue-500">
            返回登录
          </router-link>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onBeforeUnmount } from 'vue'
import { useRouter } from 'vue-router'
import { ExclamationCircleIcon, CheckCircleIcon } from '@heroicons/vue/24/outline'
import { useAuth } from '@/hooks/user/useAuth'
import { useMessage } from '@/composables/useMessage'
import type { ResetPasswordRequest } from '@/types'

const router = useRouter()
const { resetPassword, sendResetPasswordCode, loading } = useAuth()
const { success, error: showError } = useMessage()

// 表单数据
const form = ref({
  email: '',
  code: '',
  newPassword: '',
  confirmPassword: ''
})

// 状态
const error = ref('')
const successMessage = ref('')
const sendCodeLoading = ref(false)
const countdown = ref(0)
let countdownTimer: NodeJS.Timeout | null = null

// 表单验证
const isEmailValid = computed(() => {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  return emailRegex.test(form.value.email)
})

const isFormValid = computed(() => {
  return (
    form.value.email.trim() !== '' &&
    form.value.code.trim() !== '' &&
    form.value.newPassword.trim() !== '' &&
    form.value.confirmPassword.trim() !== '' &&
    form.value.newPassword === form.value.confirmPassword &&
    isEmailValid.value
  )
})

// 发送验证码
const handleSendCode = async () => {
  try {
    error.value = ''

    if (!form.value.email.trim()) {
      error.value = '请输入邮箱地址'
      return
    }

    if (!isEmailValid.value) {
      error.value = '请输入有效的邮箱地址'
      return
    }

    sendCodeLoading.value = true
    await sendResetPasswordCode(form.value.email.trim())

    success('验证码已发送到您的邮箱')

    // 开始倒计时
    startCountdown()

  } catch (err: unknown) {
    const errorMessage = err instanceof Error ? err.message : '发送验证码失败'
    error.value = errorMessage
    showError(errorMessage)
  } finally {
    sendCodeLoading.value = false
  }
}

// 开始倒计时
const startCountdown = () => {
  countdown.value = 60
  countdownTimer = setInterval(() => {
    countdown.value--
    if (countdown.value <= 0) {
      clearInterval(countdownTimer!)
      countdownTimer = null
    }
  }, 1000)
}

// 处理重置密码提交
const handleSubmit = async () => {
  try {
    error.value = ''
    successMessage.value = ''

    // 表单验证
    if (!form.value.email.trim()) {
      error.value = '请输入邮箱地址'
      return
    }

    if (!isEmailValid.value) {
      error.value = '请输入有效的邮箱地址'
      return
    }

    if (!form.value.code.trim()) {
      error.value = '请输入验证码'
      return
    }

    if (!form.value.newPassword.trim()) {
      error.value = '请输入新密码'
      return
    }

    if (form.value.newPassword.length < 6) {
      error.value = '密码长度不能少于6位'
      return
    }

    if (form.value.newPassword !== form.value.confirmPassword) {
      error.value = '两次输入的密码不一致'
      return
    }

    // 构造重置密码请求数据
    const resetData: ResetPasswordRequest = {
      email: form.value.email.trim(),
      code: form.value.code.trim(),
      newPassword: form.value.newPassword.trim()
    }

    // 调用重置密码API
    await resetPassword(resetData)

    // 重置成功
    successMessage.value = '密码重置成功！请使用新密码登录'
    success('密码重置成功！')

    // 延迟跳转到登录页
    setTimeout(() => {
      router.push('/login')
    }, 2000)

  } catch (err: unknown) {
    // 处理重置密码错误
    const errorMessage = err instanceof Error ? err.message : '重置密码失败，请重试'
    error.value = errorMessage
    showError(errorMessage)
  }
}

// 清理函数
onBeforeUnmount(() => {
  if (countdownTimer) {
    clearInterval(countdownTimer)
  }
})
</script>
