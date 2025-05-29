<template>
  <div class="w-52">
    <div class="px-3 py-2 text-xs font-semibold text-gray-900 border-b border-gray-100">
      {{ title }}
    </div>
    <div class="py-1">
      <div v-for="item in items" :key="item.name">
        <!-- 没有子菜单的项目 - 直接跳转 -->
        <router-link
          v-if="!item.children"
          :to="item.href || '#'"
          class="flex items-center px-3 py-2 text-sm text-gray-700 hover:bg-gray-50 hover:text-gray-900 transition-colors whitespace-nowrap"
          @click="$emit('itemClick')"
        >
          {{ item.name }}
        </router-link>

        <!-- 有子菜单的项目 - 点击展开下一级 -->
        <div v-else class="relative">
          <button
            :ref="(el) => setItemRef(item.name, el)"
            @click.stop="toggleSubmenu(item.name)"
            class="w-full flex items-center justify-between px-3 py-2 text-sm text-gray-700 hover:bg-gray-50 hover:text-gray-900 transition-colors whitespace-nowrap"
            :class="{ 'bg-gray-50': expandedItems[item.name] }"
          >
            <span class="flex-1 text-left">{{ item.name }}</span>
            <ChevronRightIcon
              class="h-4 w-4 transition-transform duration-200 flex-shrink-0 ml-2"
              :class="{ 'rotate-90': expandedItems[item.name] }"
            />
          </button>

          <!-- 子级弹出菜单 -->
          <GlobalPopoverMenu
            v-if="expandedItems[item.name]"
            :visible="true"
            :trigger-ref="itemRefs[item.name]"
            placement="right"
            :offset="4"
          >
            <SidebarSubMenu
              :title="item.name"
              :items="item.children"
              @item-click="handleSubItemClick"
            />
          </GlobalPopoverMenu>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive } from 'vue'
import { ChevronRightIcon } from '@heroicons/vue/24/outline'

interface MenuItem {
  name: string
  href?: string
  icon?: any
  children?: MenuItem[]
}

export default defineComponent({
  name: 'SidebarSubMenu',
  components: {
    ChevronRightIcon,
  },
  props: {
    title: {
      type: String,
      required: true,
    },
    items: {
      type: Array as () => MenuItem[],
      required: true,
    },
  },
  emits: ['itemClick'],
  setup(props, { emit }) {
    const expandedItems = reactive<Record<string, boolean>>({})
    const itemRefs = reactive<Record<string, HTMLElement>>({})

    const toggleSubmenu = (itemName: string) => {
      // 关闭其他同级菜单项
      Object.keys(expandedItems).forEach((key) => {
        if (key !== itemName) {
          expandedItems[key] = false
        }
      })

      // 切换当前菜单项
      expandedItems[itemName] = !expandedItems[itemName]
    }

    const handleSubItemClick = () => {
      // 子项被点击时，关闭所有子菜单并向上传播
      Object.keys(expandedItems).forEach((key) => {
        expandedItems[key] = false
      })
      emit('itemClick')
    }

    const setItemRef = (itemName: string, el: any) => {
      if (el && el.$el) {
        itemRefs[itemName] = el.$el as HTMLElement
      } else if (el) {
        itemRefs[itemName] = el as HTMLElement
      }
    }

    return {
      expandedItems,
      itemRefs,
      toggleSubmenu,
      handleSubItemClick,
      setItemRef,
    }
  },
})
</script>
