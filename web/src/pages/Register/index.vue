<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full space-y-8">
      <div>
        <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">
          创建新账户
        </h2>
        <p class="mt-2 text-center text-sm text-gray-600">
          或者
          <router-link to="/login" class="font-medium text-blue-600 hover:text-blue-500">
            登录已有账户
          </router-link>
        </p>
      </div>

      <form class="mt-8 space-y-6" @submit.prevent="handleRegister">
        <div class="space-y-4">
          <div>
            <label for="username" class="block text-sm font-medium text-gray-700">用户名</label>
            <input
              id="username"
              v-model="form.username"
              name="username"
              type="text"
              required
              class="mt-1 appearance-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
              placeholder="请输入用户名"
            />
          </div>

          <div>
            <label for="email" class="block text-sm font-medium text-gray-700">邮箱地址</label>
            <input
              id="email"
              v-model="form.email"
              name="email"
              type="email"
              required
              class="mt-1 appearance-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
              placeholder="请输入邮箱地址"
            />
          </div>

          <div>
            <label for="password" class="block text-sm font-medium text-gray-700">密码</label>
            <input
              id="password"
              v-model="form.password"
              name="password"
              type="password"
              required
              class="mt-1 appearance-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
              placeholder="请输入密码（6-20位）"
            />
          </div>

          <div>
            <label for="confirmPassword" class="block text-sm font-medium text-gray-700">确认密码</label>
            <input
              id="confirmPassword"
              v-model="form.confirmPassword"
              name="confirmPassword"
              type="password"
              required
              class="mt-1 appearance-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
              placeholder="请再次输入密码"
            />
          </div>

          <!-- 验证码输入 -->
          <div>
            <label for="verificationCode" class="block text-sm font-medium text-gray-700">验证码</label>
            <div class="flex space-x-2">
              <input
                id="verificationCode"
                v-model="form.verificationCode"
                name="verificationCode"
                type="text"
                required
                class="flex-1 mt-1 appearance-none relative block px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                placeholder="请输入验证码"
              />
              <button
                type="button"
                @click="handleSendCode"
                :disabled="!canSendCode || loading"
                class="px-4 py-2 bg-gray-600 text-white rounded-md hover:bg-gray-700 disabled:opacity-50 disabled:cursor-not-allowed"
              >
                {{ codeButtonText }}
              </button>
            </div>
          </div>
        </div>

        <div>
          <button
            type="submit"
            :disabled="loading || !isFormValid"
            class="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            <span v-if="loading" class="absolute left-0 inset-y-0 flex items-center pl-3">
              <div class="w-5 h-5 border-2 border-white border-t-transparent rounded-full animate-spin"></div>
            </span>
            {{ loading ? '注册中...' : '注册' }}
          </button>
        </div>
      </form>

      <!-- 错误提示 -->
      <div v-if="error" class="mt-4 p-3 bg-red-100 border border-red-400 text-red-700 rounded">
        {{ error }}
      </div>

      <!-- 成功提示 -->
      <div v-if="success" class="mt-4 p-3 bg-green-100 border border-green-400 text-green-700 rounded">
        {{ success }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuth } from '../../hooks/user/useAuth'
import type { RegisterRequest } from '../../types'

const router = useRouter()
const { register, sendRegistrationCode, loading } = useAuth()

const form = ref<{ username: string; email: string; password: string; confirmPassword: string; verificationCode: string }>({
  username: '',
  email: '',
  password: '',
  confirmPassword: '',
  verificationCode: ''
})

const error = ref('')
const success = ref('')
const countdown = ref(0)

// 计算属性
const isFormValid = computed(() => {
  return form.value.username.trim() !== '' &&
         form.value.email.trim() !== '' &&
         form.value.password.trim() !== '' &&
         form.value.confirmPassword.trim() !== '' &&
         form.value.verificationCode.trim() !== '' &&
         form.value.password === form.value.confirmPassword
})

const canSendCode = computed(() => {
  return form.value.email.trim() !== '' && countdown.value === 0
})

const codeButtonText = computed(() => {
  return countdown.value > 0 ? `${countdown.value}s后重发` : '发送验证码'
})

// 发送验证码
const handleSendCode = async () => {
  try {
    error.value = ''

    if (!form.value.email) {
      error.value = '请先输入邮箱地址'
      return
    }

    await sendRegistrationCode(form.value.email)
    success.value = '验证码已发送到您的邮箱'

    // 开始倒计时
    countdown.value = 60
    const timer = setInterval(() => {
      countdown.value--
      if (countdown.value <= 0) {
        clearInterval(timer)
      }
    }, 1000)

  } catch (err: any) {
    error.value = err.message || '发送验证码失败'
  }
}

const handleRegister = async () => {
  if (form.value.password !== form.value.confirmPassword) {
    error.value = '两次输入的密码不一致'
    return
  }

  try {
    error.value = ''
    success.value = ''

    await register({
      username: form.value.username,
      email: form.value.email,
      password: form.value.password,
      code: form.value.verificationCode
    })

    success.value = '注册成功！即将跳转到登录页面'

    // 延迟一下再跳转，让用户看到成功提示
    setTimeout(() => {
      router.push('/login')
    }, 2000)

  } catch (err: any) {
    error.value = err.message || '注册失败，请检查输入信息'
  }
}
</script>
