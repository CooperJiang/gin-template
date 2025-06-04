<template>
  <div class="min-h-screen bg-gray-50 py-12">
    <div class="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="bg-white shadow rounded-lg">
        <div class="px-6 py-4 border-b border-gray-200">
          <h1 class="text-2xl font-bold text-gray-900">个人资料</h1>
        </div>
        <div class="p-6">
          <!-- 加载状态 -->
          <div v-if="loading" class="flex justify-center items-center py-8">
            <svg class="animate-spin h-8 w-8 text-blue-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            <span class="ml-2 text-gray-600">加载中...</span>
          </div>

          <!-- 用户信息 -->
          <div v-else-if="user" class="space-y-6">
            <div>
              <label class="block text-sm font-medium text-gray-700">用户名</label>
              <div class="mt-1">
                <input
                  type="text"
                  :value="user.username"
                  readonly
                  class="block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-gray-50 text-gray-500 sm:text-sm"
                />
              </div>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">邮箱</label>
              <div class="mt-1">
                <input
                  type="email"
                  :value="user.email"
                  readonly
                  class="block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-gray-50 text-gray-500 sm:text-sm"
                />
              </div>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">用户状态</label>
              <div class="mt-1">
                <span :class="[
                  'inline-flex items-center px-3 py-1 rounded-full text-xs font-medium',
                  user.status === 1 ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
                ]">
                  {{ user.status === 1 ? '正常' : '禁用' }}
                </span>
              </div>
            </div>
            <div v-if="user.bio">
              <label class="block text-sm font-medium text-gray-700">个人简介</label>
              <div class="mt-1">
                <textarea
                  :value="user.bio"
                  readonly
                  rows="3"
                  class="block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-gray-50 text-gray-500 sm:text-sm"
                />
              </div>
            </div>
            <div v-if="user.created_at">
              <label class="block text-sm font-medium text-gray-700">注册时间</label>
              <div class="mt-1">
                <input
                  type="text"
                  :value="formatDate(user.created_at)"
                  readonly
                  class="block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-gray-50 text-gray-500 sm:text-sm"
                />
              </div>
            </div>
          </div>

          <!-- 错误状态 -->
          <div v-else class="text-center py-8">
            <div class="text-red-600 mb-4">
              <ExclamationCircleIcon class="h-12 w-12 mx-auto" />
            </div>
            <p class="text-gray-600">无法加载用户信息</p>
            <button
              @click="refreshUserInfo"
              class="mt-4 px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700"
            >
              重新加载
            </button>
          </div>

          <!-- 操作按钮 -->
          <div v-if="user" class="mt-8 flex space-x-4">
            <button
              type="button"
              class="px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700"
            >
              编辑资料
            </button>
            <router-link
              to="/settings"
              class="px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50"
            >
              账户设置
            </router-link>
            <router-link
              to="/"
              class="px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50"
            >
              返回首页
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { ExclamationCircleIcon } from '@heroicons/vue/24/outline'
import { useAuth } from '@/hooks/user/useAuth'
import { formatDate } from '@/utils/format'

const { user, loading, getUserInfo } = useAuth()

// 刷新用户信息
const refreshUserInfo = async () => {
  try {
    await getUserInfo()
  } catch (error) {
    console.error('Failed to refresh user info:', error)
  }
}

// 组件挂载时获取用户信息
onMounted(async () => {
  if (!user.value) {
    await refreshUserInfo()
  }
})
</script>
