<template>
  <div>
    <!-- 菜单项 -->
    <div v-if="!item.children">
      <router-link
        :to="item.href || '#'"
        class="group flex items-center rounded-md transition-colors relative"
        :class="[
          isCollapsed ? 'px-2 py-3 justify-center' : 'px-3 py-2',
          currentPath === item.href
            ? 'bg-blue-50 text-blue-700 border-r-2 border-blue-700'
            : 'text-gray-700 hover:text-gray-900 hover:bg-gray-50',
          'text-sm font-medium',
        ]"
        :title="isCollapsed ? item.name : ''"
      >
        <component
          :is="item.icon"
          class="flex-shrink-0"
          :class="[
            isCollapsed ? 'h-6 w-6' : 'mr-3 h-5 w-5',
            currentPath === item.href ? 'text-blue-500' : 'text-gray-400 group-hover:text-gray-500',
          ]"
        />
        <span v-if="!isCollapsed">{{ item.name }}</span>

        <!-- Tooltip for collapsed mode -->
        <div
          v-if="isCollapsed"
          class="absolute left-full ml-2 px-2 py-1 bg-gray-900 text-white text-xs rounded opacity-0 group-hover:opacity-100 transition-opacity pointer-events-none whitespace-nowrap z-50"
        >
          {{ item.name }}
        </div>
      </router-link>
    </div>

    <!-- 有子菜单的项目 -->
    <div v-else class="relative">
      <button
        ref="triggerRef"
        @click="handleMenuClick"
        class="group w-full flex items-center rounded-md transition-colors text-gray-700 hover:text-gray-900 hover:bg-gray-50 relative"
        :class="[
          isCollapsed ? 'px-2 py-3 justify-center' : 'px-3 py-2',
          { 'bg-gray-50': (isExpanded && !isCollapsed) || (showSubmenu && isCollapsed) },
          'text-sm font-medium',
        ]"
        :title="isCollapsed ? item.name : ''"
      >
        <component
          :is="item.icon"
          class="flex-shrink-0"
          :class="[
            isCollapsed ? 'h-6 w-6' : 'mr-3 h-5 w-5',
            'text-gray-400 group-hover:text-gray-500',
          ]"
        />
        <span v-if="!isCollapsed" class="flex-1 text-left">{{ item.name }}</span>
        <ChevronRightIcon
          v-if="!isCollapsed"
          class="ml-2 h-4 w-4 transition-transform"
          :class="{ 'rotate-90': isExpanded }"
        />
      </button>

      <!-- 使用PopoverMenu组件处理折叠状态的子菜单 -->
      <GlobalPopoverMenu
        v-if="isCollapsed && showSubmenu"
        :visible="true"
        :trigger-ref="triggerRef"
        placement="right"
        :offset="8"
      >
        <SidebarSubMenu :title="item.name" :items="item.children" @item-click="hideSubmenu" />
      </GlobalPopoverMenu>

      <!-- 子菜单 (only show when not collapsed) -->
      <div v-if="isExpanded && !isCollapsed" class="ml-6 mt-1 space-y-1">
        <SidebarMenuItem
          v-for="child in item.children"
          :key="child.name"
          :item="child"
          :current-path="currentPath"
          :level="level + 1"
          :is-collapsed="false"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { ref, computed, defineComponent } from 'vue'
import { ChevronRightIcon } from '@heroicons/vue/24/outline'
import SidebarSubMenu from './SidebarSubMenu.vue'

interface MenuItem {
  name: string
  href?: string
  icon: any
  children?: MenuItem[]
}

export default defineComponent({
  name: 'SidebarMenuItem',
  components: {
    ChevronRightIcon,
    SidebarSubMenu,
  },
  props: {
    item: {
      type: Object as () => MenuItem,
      required: true,
    },
    currentPath: {
      type: String,
      required: true,
    },
    level: {
      type: Number,
      default: 0,
    },
    isCollapsed: {
      type: Boolean,
      default: false,
    },
  },
  setup(props) {
    const isExpanded = ref(false)
    const showSubmenu = ref(false)
    const triggerRef = ref<HTMLElement>()

    // 检查是否有子项处于活跃状态
    const hasActiveChild = computed(() => {
      if (!props.item.children) return false

      const checkActive = (items: MenuItem[]): boolean => {
        return items.some((child) => {
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

    const handleMenuClick = () => {
      if (props.isCollapsed) {
        // 折叠状态下点击切换子菜单显示
        showSubmenu.value = !showSubmenu.value
      } else {
        // 展开状态下点击切换展开/收起
        toggleExpanded()
      }
    }

    const toggleExpanded = () => {
      isExpanded.value = !isExpanded.value
    }

    const hideSubmenu = () => {
      showSubmenu.value = false
    }

    return {
      isExpanded,
      showSubmenu,
      triggerRef,
      hasActiveChild,
      handleMenuClick,
      toggleExpanded,
      hideSubmenu,
    }
  },
})
</script>
