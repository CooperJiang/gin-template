<template>
  <div class="min-h-screen bg-gray-50">
    <!-- 导航栏 -->
    <nav class="bg-white shadow-sm">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <div class="flex items-center">
            <h1 class="text-xl font-bold text-gray-900">用户端</h1>
          </div>
          <div class="flex items-center space-x-4">
            <template v-if="isAuthenticated">
              <router-link
                to="/profile"
                class="text-gray-700 hover:text-blue-600 px-3 py-2 rounded-md text-sm font-medium"
              >
                个人资料
              </router-link>
              <router-link
                to="/settings"
                class="text-gray-700 hover:text-blue-600 px-3 py-2 rounded-md text-sm font-medium"
              >
                设置
              </router-link>
              <button
                @click="handleLogout"
                class="bg-red-600 text-white px-4 py-2 rounded-md text-sm font-medium hover:bg-red-700"
              >
                退出登录
              </button>
            </template>
            <template v-else>
              <router-link
                to="/login"
                class="text-gray-700 hover:text-blue-600 px-3 py-2 rounded-md text-sm font-medium"
              >
                登录
              </router-link>
              <router-link
                to="/register"
                class="bg-blue-600 text-white px-4 py-2 rounded-md text-sm font-medium hover:bg-blue-700"
              >
                注册
              </router-link>
            </template>
          </div>
        </div>
      </div>
    </nav>

    <!-- 主要内容 -->
    <main class="max-w-7xl mx-auto py-12 px-4 sm:px-6 lg:px-8">
      <div class="text-center">
        <h1 class="text-4xl font-bold text-gray-900 mb-8">
          欢迎来到用户端
        </h1>
        <p class="text-xl text-gray-600 mb-12">
          这是一个基于 Vue 3 + TypeScript + TailwindCSS 的用户端模板
        </p>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-8 mt-12">
          <div class="bg-white p-6 rounded-lg shadow-md">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">现代技术栈</h3>
            <p class="text-gray-600">
              使用 Vue 3 Composition API、TypeScript 和 TailwindCSS 构建
            </p>
          </div>
          <div class="bg-white p-6 rounded-lg shadow-md">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">安全认证</h3>
            <p class="text-gray-600">
              内置用户认证系统，支持登录、注册、密码重置等功能
            </p>
          </div>
          <div class="bg-white p-6 rounded-lg shadow-md">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">响应式设计</h3>
            <p class="text-gray-600">
              完全响应式设计，在各种设备上都有良好的用户体验
            </p>
          </div>
        </div>

        <div class="mt-12">
          <router-link
            to="/about"
            class="inline-flex items-center px-6 py-3 border border-transparent text-base font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700"
          >
            了解更多
          </router-link>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { useAuth } from '@/hooks/user/useAuth'
import { useMessage } from '@/composables/useMessage'

const { isAuthenticated, logout, user } = useAuth()
const { success } = useMessage()

const handleLogout = () => {
  logout()
  success('已成功退出登录')
}
</script>
