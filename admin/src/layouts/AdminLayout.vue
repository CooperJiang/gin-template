<template>
  <div class="h-screen bg-gray-50 flex overflow-hidden">
    <!-- 侧边栏 -->
    <AppSidebar :is-open="sidebarOpen" @toggle="toggleSidebar" />

    <!-- 主内容区域 -->
    <div class="flex-1 flex flex-col lg:ml-0 h-full">
      <!-- 顶部导航栏 -->
      <AppNavbar @toggle-sidebar="toggleSidebar" />

      <!-- 主内容 -->
      <main class="flex-1 overflow-y-auto overflow-x-hidden custom-scrollbar">
        <div class="p-4 sm:p-6 lg:p-8">
          <router-view />
        </div>
      </main>

      <!-- Footer -->
      <footer class="bg-white border-t border-gray-200 px-4 py-3 flex-shrink-0">
        <div class="flex items-center justify-between text-sm text-gray-500">
          <div>© 2025 管理系统. All rights reserved.</div>
          <div class="flex flex-col sm:flex-row sm:space-x-4 space-y-1 sm:space-y-0 text-right sm:text-left">
            <div class="flex items-center space-x-2">
              <span>版本 v{{ versionInfo?.version || '1.0.0' }}</span>
              <span v-if="!loading" class="w-2 h-2 bg-green-500 rounded-full" title="版本信息已加载"></span>
              <span v-else class="w-2 h-2 bg-yellow-500 rounded-full animate-pulse" title="加载中..."></span>
            </div>
            <span class="hidden sm:inline">•</span>
            <div class="text-xs">
              <div>构建时间: {{ versionInfo?.buildDate || new Date().toLocaleDateString('zh-CN') }}</div>
            </div>
          </div>
        </div>
      </footer>
    </div>

    <!-- 移动端遮罩 -->
    <div
      v-if="sidebarOpen"
      class="fixed inset-0 z-40 bg-gray-600 bg-opacity-75 lg:hidden"
      @click="toggleSidebar"
    ></div>
  </div>
</template>

<script setup lang="ts">
import { ref, defineOptions } from 'vue'
import AppNavbar from './components/AppNavbar.vue'
import AppSidebar from './components/AppSidebar.vue'
import { useVersion } from '@/composables/useVersion'

const sidebarOpen = ref(false)

const toggleSidebar = () => {
  sidebarOpen.value = !sidebarOpen.value
}

// 使用版本信息
const { versionInfo, loading } = useVersion()

defineOptions({
  name: 'AdminLayout',
})
</script>

<style scoped>
/* 确保侧边栏在移动端正确显示 */
@media (max-width: 1024px) {
  aside {
    position: fixed;
  }
}
</style>
