<template>
  <div>
    <!-- 欢迎信息 -->
    <div class="mb-8">
      <h2 class="text-2xl font-bold text-gray-900 mb-2">欢迎回来，{{ user?.username }}！</h2>
      <p class="text-gray-600">
        今天是
        {{
          new Date().toLocaleDateString('zh-CN', {
            year: 'numeric',
            month: 'long',
            day: 'numeric',
            weekday: 'long',
          })
        }}
      </p>
    </div>

    <!-- 统计卡片 -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
      <div class="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
        <div class="flex items-center">
          <div class="p-2 bg-blue-100 rounded-lg">
            <svg
              class="w-6 h-6 text-blue-600"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"
              ></path>
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-600">总用户数</p>
            <p class="text-2xl font-bold text-gray-900">1,234</p>
          </div>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
        <div class="flex items-center">
          <div class="p-2 bg-green-100 rounded-lg">
            <svg
              class="w-6 h-6 text-green-600"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"
              ></path>
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-600">月收入</p>
            <p class="text-2xl font-bold text-gray-900">¥56,789</p>
          </div>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
        <div class="flex items-center">
          <div class="p-2 bg-purple-100 rounded-lg">
            <svg
              class="w-6 h-6 text-purple-600"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"
              ></path>
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-600">活跃度</p>
            <p class="text-2xl font-bold text-gray-900">89.5%</p>
          </div>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
        <div class="flex items-center">
          <div class="p-2 bg-red-100 rounded-lg">
            <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L4.082 15.5c-.77.833.192 2.5 1.732 2.5z"
              ></path>
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-600">错误数</p>
            <p class="text-2xl font-bold text-gray-900">23</p>
          </div>
        </div>
      </div>
    </div>

    <!-- 快捷操作和系统信息 -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- 快捷操作 -->
      <div class="bg-white rounded-lg shadow p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">快捷操作</h3>
        <div class="space-y-3">
          <router-link
            to="/users"
            class="flex items-center p-3 rounded-lg border border-gray-200 hover:bg-gray-50 transition-colors"
          >
            <UsersIcon class="h-5 w-5 text-gray-400 mr-3" />
            <span class="text-sm font-medium text-gray-900">用户管理</span>
          </router-link>

          <router-link
            to="/users/create"
            class="flex items-center p-3 rounded-lg border border-gray-200 hover:bg-gray-50 transition-colors"
          >
            <PlusIcon class="h-5 w-5 text-gray-400 mr-3" />
            <span class="text-sm font-medium text-gray-900">创建用户</span>
          </router-link>

          <router-link
            to="/profile"
            class="flex items-center p-3 rounded-lg border border-gray-200 hover:bg-gray-50 transition-colors"
          >
            <UserIcon class="h-5 w-5 text-gray-400 mr-3" />
            <span class="text-sm font-medium text-gray-900">个人资料</span>
          </router-link>
        </div>
      </div>

      <!-- 系统信息 -->
      <div class="bg-white rounded-lg shadow p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">系统信息</h3>
        <div class="space-y-3">
          <div class="flex justify-between">
            <span class="text-sm text-gray-500">系统版本</span>
            <span class="text-sm font-medium text-gray-900">v1.0.0</span>
          </div>
          <div class="flex justify-between">
            <span class="text-sm text-gray-500">运行时间</span>
            <span class="text-sm font-medium text-gray-900">{{ uptime }}</span>
          </div>
          <div class="flex justify-between">
            <span class="text-sm text-gray-500">最后登录</span>
            <span class="text-sm font-medium text-gray-900">{{ lastLogin }}</span>
          </div>
          <div class="flex justify-between">
            <span class="text-sm text-gray-500">当前用户</span>
            <span class="text-sm font-medium text-gray-900">{{ user?.username }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- 最近活动 -->
    <div class="mt-8">
      <div class="bg-white rounded-lg shadow p-6">
        <h3 class="text-lg font-medium text-gray-900 mb-4">最近活动</h3>
        <div class="space-y-4">
          <div class="flex items-center p-3 bg-gray-50 rounded-lg">
            <div class="flex-shrink-0">
              <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
                <UserIcon class="h-4 w-4 text-blue-600" />
              </div>
            </div>
            <div class="ml-3">
              <p class="text-sm font-medium text-gray-900">用户登录</p>
              <p class="text-xs text-gray-500">{{ user?.username }} 于 {{ lastLogin }} 登录系统</p>
            </div>
          </div>

          <div class="flex items-center p-3 bg-gray-50 rounded-lg">
            <div class="flex-shrink-0">
              <div class="w-8 h-8 bg-green-100 rounded-full flex items-center justify-center">
                <CheckCircleIcon class="h-4 w-4 text-green-600" />
              </div>
            </div>
            <div class="ml-3">
              <p class="text-sm font-medium text-gray-900">系统状态</p>
              <p class="text-xs text-gray-500">系统运行正常，所有服务可用</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useAuth } from '../../hooks/user/useAuth'
import { UsersIcon, CheckCircleIcon, PlusIcon, UserIcon } from '@heroicons/vue/24/outline'

const { user } = useAuth()

const uptime = computed(() => {
  // 简单的运行时间计算
  const now = new Date()
  const start = new Date(now.getFullYear(), now.getMonth(), now.getDate())
  const diff = now.getTime() - start.getTime()
  const hours = Math.floor(diff / (1000 * 60 * 60))
  const minutes = Math.floor((diff % (1000 * 60 * 60)) / (1000 * 60))
  return `${hours}小时${minutes}分钟`
})

const lastLogin = computed(() => {
  return new Date().toLocaleString('zh-CN')
})

defineOptions({
  name: 'DashboardPage',
})
</script>
