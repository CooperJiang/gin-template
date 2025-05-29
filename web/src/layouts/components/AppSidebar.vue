<template>
  <div class="relative">
    <aside
      class="fixed inset-y-0 left-0 z-50 bg-white shadow-lg transform transition-all duration-300 ease-in-out lg:translate-x-0 lg:static lg:inset-0 lg:h-screen flex flex-col overflow-hidden"
      :class="[
        { '-translate-x-full': !isOpen && !isCollapsed, 'translate-x-0': isOpen || isCollapsed },
        isCollapsed ? 'w-16' : 'w-64',
      ]"
    >
      <!-- 侧边栏头部 -->
      <div class="flex items-center h-16 px-3 border-b border-gray-200 flex-shrink-0">
        <div v-if="!isCollapsed" class="flex-1">
          <AppLogo :class="{ 'animate-fade-in': !isCollapsed }" />
        </div>
        <div v-else class="flex justify-center w-full">
          <AppLogo :show-text="false" />
        </div>

        <!-- 移动端关闭按钮 -->
        <button
          v-if="!isCollapsed"
          @click="$emit('toggle')"
          class="lg:hidden p-1 rounded-md text-gray-400 hover:text-gray-500 hover:bg-gray-100"
        >
          <XMarkIcon class="h-6 w-6" />
        </button>
      </div>

      <!-- 导航菜单 -->
      <nav
        class="flex-1 overflow-y-auto overflow-x-hidden"
        :class="isCollapsed ? 'px-2 pt-6' : 'px-3 pt-6'"
      >
        <div class="space-y-1">
          <SidebarMenuItem
            v-for="item in navigation"
            :key="item.name"
            :item="item"
            :current-path="currentPath"
            :is-collapsed="isCollapsed"
            :class="{ 'animate-fade-in': !isCollapsed }"
          />
        </div>
      </nav>

      <!-- 底部区域 -->
      <div class="flex-shrink-0 border-t border-gray-200 p-3">
        <!-- 展开状态 -->
        <div
          v-if="!isCollapsed"
          class="bg-gray-50 rounded-lg p-3 flex items-center justify-between animate-fade-in"
        >
          <!-- 小提示信息 -->
          <div class="flex items-center flex-1">
            <div
              class="w-6 h-6 bg-blue-100 rounded-full flex items-center justify-center flex-shrink-0"
            >
              <svg class="w-3 h-3 text-blue-600" fill="currentColor" viewBox="0 0 20 20">
                <path
                  fill-rule="evenodd"
                  d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"
                  clip-rule="evenodd"
                />
              </svg>
            </div>
            <span class="ml-2 text-xs text-gray-600 font-medium">点击菜单展开子项</span>
          </div>

          <!-- 折叠按钮 -->
          <button
            @click="toggleCollapse"
            class="group relative p-1.5 rounded-md border border-gray-300 bg-white hover:bg-gray-50 hover:border-gray-400 transition-all duration-200 shadow-sm hover:shadow"
            title="折叠侧边栏"
          >
            <div class="w-4 h-4 text-gray-500 group-hover:text-gray-700">
              <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M4 6h16M4 12h16M4 18h16"
                />
              </svg>
            </div>

            <!-- 悬停提示 -->
            <div
              class="absolute -top-8 left-1/2 transform -translate-x-1/2 px-2 py-1 bg-gray-900 text-white text-xs rounded opacity-0 group-hover:opacity-100 transition-opacity pointer-events-none whitespace-nowrap z-50"
            >
              折叠侧边栏
            </div>
          </button>
        </div>

        <!-- 折叠状态 -->
        <div v-else class="flex justify-center">
          <button
            @click="toggleCollapse"
            class="group relative p-2 rounded-md border border-gray-300 bg-white hover:bg-gray-50 hover:border-gray-400 transition-all duration-200 shadow-sm hover:shadow"
            title="展开侧边栏"
          >
            <div class="w-4 h-4 text-gray-500 group-hover:text-gray-700">
              <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M4 8V4a1 1 0 011-1h4m5 0h4a1 1 0 011 1v4m0 5v4a1 1 0 01-1 1h-4m-5 0H5a1 1 0 01-1-1v-4m2-5a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H8a2 2 0 01-2-2v-2z"
                />
              </svg>
            </div>

            <!-- 悬停提示 -->
            <div
              class="absolute -top-8 left-1/2 transform -translate-x-1/2 px-2 py-1 bg-gray-900 text-white text-xs rounded opacity-0 group-hover:opacity-100 transition-opacity pointer-events-none whitespace-nowrap z-50"
            >
              展开侧边栏
            </div>
          </button>
        </div>
      </div>
    </aside>

    <!-- 遮罩层 (移动端) -->
    <div
      v-if="isOpen && !isCollapsed"
      class="fixed inset-0 bg-gray-600 bg-opacity-75 lg:hidden z-40"
      @click="$emit('toggle')"
    ></div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { useRoute } from 'vue-router'
import { XMarkIcon } from '@heroicons/vue/24/outline'
import AppLogo from './AppLogo.vue'
import SidebarMenuItem from './SidebarMenuItem.vue'
import { navigationConfig } from '../config/navigation'

interface Props {
  isOpen: boolean
}

defineProps<Props>()

defineEmits<{
  toggle: []
}>()

const route = useRoute()
const currentPath = computed(() => route.path)

// 折叠状态
const isCollapsed = ref(false)

const toggleCollapse = () => {
  isCollapsed.value = !isCollapsed.value
}

// 导航配置
const navigation = navigationConfig
</script>

<style scoped>
.animate-fade-in {
  animation: fadeIn 0.4s ease-in-out 0.2s both;
}

@keyframes fadeIn {
  0% {
    opacity: 0;
  }
  100% {
    opacity: 1;
  }
}
</style>
