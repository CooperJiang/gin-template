<template>
  <div>
    <!-- 欢迎信息 -->
    <div class="mb-8">
      <h2 class="text-2xl font-bold text-gray-900 mb-2">
        欢迎回来，{{ user?.username }}！
      </h2>
      <p class="text-gray-600">
        今天是 {{ new Date().toLocaleDateString('zh-CN', {
          year: 'numeric',
          month: 'long',
          day: 'numeric',
          weekday: 'long'
        }) }}
      </p>
    </div>

    <!-- 统计卡片 -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
      <div class="bg-white rounded-lg shadow p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <UsersIcon class="h-8 w-8 text-blue-600" />
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-500">总用户数</p>
            <p class="text-2xl font-semibold text-gray-900">--</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <CheckCircleIcon class="h-8 w-8 text-green-600" />
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-500">活跃用户</p>
            <p class="text-2xl font-semibold text-gray-900">--</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <XCircleIcon class="h-8 w-8 text-red-600" />
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-500">禁用用户</p>
            <p class="text-2xl font-semibold text-gray-900">--</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <PlusCircleIcon class="h-8 w-8 text-purple-600" />
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-500">今日新增</p>
            <p class="text-2xl font-semibold text-gray-900">--</p>
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
import {
  UsersIcon,
  CheckCircleIcon,
  XCircleIcon,
  PlusCircleIcon,
  PlusIcon,
  UserIcon
} from '@heroicons/vue/24/outline'

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
</script>
