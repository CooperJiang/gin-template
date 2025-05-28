<template>
  <div class="p-6">
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">个人信息</h1>
      <p class="mt-2 text-gray-600">查看和管理您的个人资料</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- 用户信息卡片 -->
      <div class="bg-white rounded-lg shadow p-6">
        <h2 class="text-lg font-semibold mb-4">基本信息</h2>

        <div v-if="loading" class="flex justify-center py-8">
          <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
        </div>

        <div v-else-if="userInfo" class="space-y-4">
          <div class="flex items-center space-x-4">
            <div class="w-16 h-16 bg-gray-200 rounded-full flex items-center justify-center">
              <img
                v-if="userInfo.avatar"
                :src="userInfo.avatar"
                :alt="userInfo.username"
                class="w-16 h-16 rounded-full object-cover"
              >
              <span v-else class="text-xl font-semibold text-gray-600">
                {{ userInfo.username?.charAt(0).toUpperCase() }}
              </span>
            </div>
            <div>
              <h3 class="text-lg font-medium">{{ userInfo.username }}</h3>
              <p class="text-gray-600">{{ userInfo.email }}</p>
            </div>
          </div>

          <div class="grid grid-cols-1 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700">用户名</label>
              <div class="mt-1 p-3 bg-gray-50 rounded-md">{{ userInfo.username }}</div>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700">邮箱</label>
              <div class="mt-1 p-3 bg-gray-50 rounded-md">{{ userInfo.email }}</div>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700">状态</label>
              <div class="mt-1 p-3 bg-gray-50 rounded-md">
                <span
                  :class="{
                    'text-green-600': userInfo.status === 1,
                    'text-red-600': userInfo.status !== 1
                  }"
                >
                  {{ userInfo.status === 1 ? '正常' : '禁用' }}
                </span>
              </div>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700">角色</label>
              <div class="mt-1 p-3 bg-gray-50 rounded-md">
                {{ userInfo.role === 1 ? '管理员' : '普通用户' }}
              </div>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700">注册时间</label>
              <div class="mt-1 p-3 bg-gray-50 rounded-md">
                {{ formatDate(userInfo.created_at) }}
              </div>
            </div>
          </div>
        </div>

        <div v-else class="text-center py-8 text-gray-500">
          无法加载用户信息
        </div>
      </div>

      <!-- 修改密码卡片 -->
      <div class="bg-white rounded-lg shadow p-6">
        <h2 class="text-lg font-semibold mb-4">修改密码</h2>

        <form @submit.prevent="handleChangePassword" class="space-y-4">
          <div>
            <label for="oldPassword" class="block text-sm font-medium text-gray-700">
              当前密码
            </label>
            <input
              id="oldPassword"
              v-model="passwordForm.oldPassword"
              type="password"
              required
              class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="请输入当前密码"
            />
          </div>

          <div>
            <label for="newPassword" class="block text-sm font-medium text-gray-700">
              新密码
            </label>
            <input
              id="newPassword"
              v-model="passwordForm.newPassword"
              type="password"
              required
              minlength="6"
              class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="请输入新密码（至少6位）"
            />
          </div>

          <div>
            <label for="confirmPassword" class="block text-sm font-medium text-gray-700">
              确认新密码
            </label>
            <input
              id="confirmPassword"
              v-model="passwordForm.confirmPassword"
              type="password"
              required
              class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="请再次输入新密码"
            />
          </div>

          <div v-if="passwordError" class="text-red-600 text-sm">
            {{ passwordError }}
          </div>

          <div v-if="passwordSuccess" class="text-green-600 text-sm">
            {{ passwordSuccess }}
          </div>

          <button
            type="submit"
            :disabled="passwordLoading"
            class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            <span v-if="passwordLoading" class="flex items-center">
              <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              修改中...
            </span>
            <span v-else>修改密码</span>
          </button>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useAuth } from '../../hooks/user/useAuth'
import { authApi } from '../../api/auth'
import type { ChangePasswordRequest } from '../../types'

const { user, getUserInfo } = useAuth()

// 用户信息
const userInfo = ref<any>(null)
const loading = ref(false)

// 密码表单
const passwordForm = ref<ChangePasswordRequest & { confirmPassword: string }>({
  oldPassword: '',
  newPassword: '',
  confirmPassword: ''
})

const passwordLoading = ref(false)
const passwordError = ref('')
const passwordSuccess = ref('')

// 格式化日期
const formatDate = (dateString: string) => {
  if (!dateString) return ''
  return new Date(dateString).toLocaleString('zh-CN')
}

// 获取用户信息
const fetchUserInfo = async () => {
  try {
    loading.value = true
    // 先尝试从useAuth获取用户信息
    if (user.value) {
      userInfo.value = user.value
    } else {
      // 如果没有，则从API获取
      const response = await getUserInfo()
      userInfo.value = response
    }
  } catch (error: any) {
    console.error('获取用户信息失败:', error)
  } finally {
    loading.value = false
  }
}

// 修改密码
const handleChangePassword = async () => {
  passwordError.value = ''
  passwordSuccess.value = ''

  // 验证密码
  if (passwordForm.value.newPassword !== passwordForm.value.confirmPassword) {
    passwordError.value = '两次输入的新密码不一致'
    return
  }

  if (passwordForm.value.newPassword.length < 6) {
    passwordError.value = '新密码长度不能少于6位'
    return
  }

  try {
    passwordLoading.value = true

    await authApi.changePassword({
      oldPassword: passwordForm.value.oldPassword,
      newPassword: passwordForm.value.newPassword
    })

    passwordSuccess.value = '密码修改成功'

    // 清空表单
    passwordForm.value = {
      oldPassword: '',
      newPassword: '',
      confirmPassword: ''
    }
  } catch (error: any) {
    passwordError.value = error.response?.data?.message || '密码修改失败'
  } finally {
    passwordLoading.value = false
  }
}

onMounted(() => {
  fetchUserInfo()
})
</script>
