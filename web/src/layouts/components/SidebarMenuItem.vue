<template>
  <div>
    <!-- 菜单项 -->
    <div v-if="!item.children">
      <router-link
        :to="item.href || '#'"
        class="group flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors"
        :class="[
          currentPath === item.href
            ? 'bg-blue-50 text-blue-700 border-r-2 border-blue-700'
            : 'text-gray-700 hover:text-gray-900 hover:bg-gray-50'
        ]"
      >
        <component
          :is="item.icon"
          class="mr-3 h-5 w-5 flex-shrink-0"
          :class="[
            currentPath === item.href
              ? 'text-blue-500'
              : 'text-gray-400 group-hover:text-gray-500'
          ]"
        />
        {{ item.name }}
      </router-link>
    </div>

    <!-- 有子菜单的项目 -->
    <div v-else>
      <button
        @click="toggleExpanded"
        class="group w-full flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors text-gray-700 hover:text-gray-900 hover:bg-gray-50"
        :class="{ 'bg-gray-50': isExpanded }"
      >
        <component
          :is="item.icon"
          class="mr-3 h-5 w-5 flex-shrink-0 text-gray-400 group-hover:text-gray-500"
        />
        <span class="flex-1 text-left">{{ item.name }}</span>
        <ChevronRightIcon
          class="ml-2 h-4 w-4 transition-transform"
          :class="{ 'rotate-90': isExpanded }"
        />
      </button>

      <!-- 子菜单 -->
      <div v-if="isExpanded" class="ml-6 mt-1 space-y-1">
        <SidebarMenuItem
          v-for="child in item.children"
          :key="child.name"
          :item="child"
          :current-path="currentPath"
          :level="level + 1"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { ChevronRightIcon } from '@heroicons/vue/24/outline'

interface MenuItem {
  name: string
  href?: string
  icon: any
  children?: MenuItem[]
}

interface Props {
  item: MenuItem
  currentPath: string
  level?: number
}

const props = withDefaults(defineProps<Props>(), {
  level: 0
})

const isExpanded = ref(false)

// 检查是否有子项处于活跃状态
const hasActiveChild = computed(() => {
  if (!props.item.children) return false

  const checkActive = (items: MenuItem[]): boolean => {
    return items.some(child => {
      if (child.href === props.currentPath) return true
      if (child.children) return checkActive(child.children)
      return false
    })
  }

  return checkActive(props.item.children)
})

// 如果有活跃的子项，自动展开
if (hasActiveChild.value) {
  isExpanded.value = true
}

const toggleExpanded = () => {
  isExpanded.value = !isExpanded.value
}
</script>
