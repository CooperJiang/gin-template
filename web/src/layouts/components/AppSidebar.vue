<template>
  <aside
    class="fixed inset-y-0 left-0 z-50 w-64 bg-white shadow-lg transform transition-transform duration-300 ease-in-out lg:translate-x-0 lg:static lg:inset-0"
    :class="{ '-translate-x-full': !isOpen, 'translate-x-0': isOpen }"
  >
    <!-- 侧边栏头部 -->
    <div class="flex items-center justify-between h-16 px-6 border-b border-gray-200">
      <AppLogo />
      <button
        @click="$emit('toggle')"
        class="lg:hidden p-1 rounded-md text-gray-400 hover:text-gray-500 hover:bg-gray-100"
      >
        <XMarkIcon class="h-6 w-6" />
      </button>
    </div>

    <!-- 导航菜单 -->
    <nav class="mt-6 px-3 flex-1 overflow-y-auto">
      <div class="space-y-1">
        <SidebarMenuItem
          v-for="item in navigation"
          :key="item.name"
          :item="item"
          :current-path="currentPath"
        />
      </div>
    </nav>

    <!-- 侧边栏底部提示 -->
    <div class="absolute bottom-0 left-0 right-0 p-4 border-t border-gray-200">
      <div class="bg-blue-50 rounded-lg p-3">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center">
              <span class="text-blue-600 text-xs">💡</span>
            </div>
          </div>
          <div class="ml-3">
            <p class="text-xs font-medium text-blue-900">小提示</p>
            <p class="text-xs text-blue-700">点击菜单可展开子项</p>
          </div>
        </div>
      </div>
    </div>
  </aside>
</template>

<script setup lang="ts">
import { computed } from 'vue'
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

// 导航配置
const navigation = navigationConfig
</script>
