<template>
  <div :class="tabsClass">
    <!-- 标签页头部 -->
    <div :class="tabsHeaderClass">
      <!-- 标签页列表 -->
      <div :class="tabsNavClass">
        <div
          v-for="item in items"
          :key="item.key"
          :class="getTabClass(item)"
          @click="handleTabClick(item)"
        >
          <!-- 图标 -->
          <GlobalIcon
            v-if="item.icon"
            :name="item.icon"
            size="sm"
            :class="[
              'transition-colors duration-200',
              { 'mr-2': item.title }
            ]"
          />

          <!-- 标题 -->
          <slot name="title" :item="item" :active="isActive(item)">
            <span class="truncate">{{ item.title }}</span>
          </slot>

          <!-- 关闭按钮 -->
          <GlobalButton
            v-if="item.closable && !item.disabled"
            type="secondary"
            variant="ghost"
            size="small"
            shape="circle"
            icon="x-mark"
            class="ml-2 opacity-0 group-hover:opacity-100 transition-opacity duration-200"
            @click.stop="handleTabClose(item)"
          />
        </div>

        <!-- 添加按钮 -->
        <div v-if="addable" class="flex-shrink-0 ml-2">
          <slot name="addButton">
            <GlobalButton
              type="secondary"
              variant="ghost"
              size="small"
              shape="circle"
              icon="plus"
              class="border-dashed"
              @click="handleAdd"
            />
          </slot>
        </div>

        <!-- 线条指示器 (仅line类型) -->
        <div
          v-if="type === 'line'"
          :class="indicatorClass"
          :style="indicatorStyle"
        />
      </div>

      <!-- 额外操作区域 -->
      <div v-if="$slots.extra" class="flex-shrink-0 ml-4">
        <slot name="extra" />
      </div>
    </div>

    <!-- 标签页内容 -->
    <div :class="tabsContentClass">
      <div
        v-for="item in items"
        v-show="isActive(item)"
        :key="item.key"
        :class="tabPaneClass"
      >
        <slot :name="item.key" :item="item">
          <div v-if="item.content" v-html="item.content" />
        </slot>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, nextTick } from 'vue'
import type { TabsProps, TabsEmits, TabItem } from './types'

defineOptions({
  name: 'GlobalTabs'
})

const props = withDefaults(defineProps<TabsProps>(), {
  type: 'line',
  position: 'top',
  size: 'medium',
  addable: false,
  sortable: false,
  lazy: false,
  animated: true,
  items: () => []
})

const emit = defineEmits<TabsEmits>()

// 内部状态
const tabsNavRef = ref<HTMLElement>()
const indicatorStyle = ref({})

// 计算当前激活的标签页
const activeKey = computed({
  get: () => props.modelValue || (props.items[0]?.key ?? ''),
  set: (value) => emit('update:modelValue', value)
})

// 标签页容器样式
const tabsClass = computed(() => {
  const classes = ['tabs-container']

  if (['left', 'right'].includes(props.position)) {
    classes.push('flex')
    if (props.position === 'left') {
      classes.push('flex-row')
    } else {
      classes.push('flex-row-reverse')
    }
  }

  return classes.join(' ')
})

// 标签页头部样式
const tabsHeaderClass = computed(() => {
  const classes = ['tabs-header', 'flex', 'items-center']

  if (['left', 'right'].includes(props.position)) {
    classes.push('flex-col', 'border-b-0')
    if (props.position === 'left') {
      classes.push('border-r', 'border-gray-200')
    } else {
      classes.push('border-l', 'border-gray-200')
    }
  } else {
    classes.push('flex-row')
    if (props.position === 'top') {
      classes.push('border-b', 'border-gray-200')
    } else {
      classes.push('border-t', 'border-gray-200')
    }
  }

  return classes.join(' ')
})

// 标签页导航样式
const tabsNavClass = computed(() => {
  const classes = ['tabs-nav', 'relative', 'flex']

  if (['left', 'right'].includes(props.position)) {
    classes.push('flex-col', 'space-y-1')
  } else {
    classes.push('flex-row', 'space-x-1')
  }

  return classes.join(' ')
})

// 标签页内容样式
const tabsContentClass = computed(() => {
  const classes = ['tabs-content']

  if (['left', 'right'].includes(props.position)) {
    classes.push('flex-1', 'min-w-0')
  }

  return classes.join(' ')
})

// 标签页面板样式
const tabPaneClass = computed(() => {
  const classes = ['tab-pane']

  if (props.animated) {
    classes.push('transition-all', 'duration-200')
  }

  const sizeClasses = {
    small: 'p-3',
    medium: 'p-4',
    large: 'p-6'
  }
  classes.push(sizeClasses[props.size])

  return classes.join(' ')
})

// 指示器样式
const indicatorClass = computed(() => {
  const classes = [
    'tabs-indicator',
    'absolute',
    'bg-blue-600',
    'transition-all',
    'duration-200',
    'ease-in-out'
  ]

  if (['top', 'bottom'].includes(props.position)) {
    classes.push('h-0.5')
    if (props.position === 'top') {
      classes.push('bottom-0')
    } else {
      classes.push('top-0')
    }
  } else {
    classes.push('w-0.5')
    if (props.position === 'left') {
      classes.push('right-0')
    } else {
      classes.push('left-0')
    }
  }

  return classes.join(' ')
})

// 获取单个标签页样式
const getTabClass = (item: TabItem) => {
  const classes = [
    'tab-item',
    'group',
    'relative',
    'flex',
    'items-center',
    'cursor-pointer',
    'transition-all',
    'duration-200',
    'select-none'
  ]

  // 尺寸样式
  const sizeClasses = {
    small: 'px-3 py-2 text-sm',
    medium: 'px-4 py-2.5 text-sm',
    large: 'px-6 py-3 text-base'
  }
  classes.push(sizeClasses[props.size])

  // 类型样式
  if (props.type === 'card') {
    classes.push('border', 'rounded-t-md', 'bg-white')
    if (isActive(item)) {
      classes.push('border-gray-300', 'border-b-white', '-mb-px')
    } else {
      classes.push('border-transparent', 'hover:border-gray-300')
    }
  } else if (props.type === 'button') {
    classes.push('rounded-md', 'border')
    if (isActive(item)) {
      classes.push('bg-blue-100', 'border-blue-300', 'text-blue-700')
    } else {
      classes.push('border-transparent', 'hover:bg-gray-100')
    }
  } else {
    // line 类型
    if (isActive(item)) {
      classes.push('text-blue-600')
    } else {
      classes.push('text-gray-600', 'hover:text-gray-900')
    }
  }

  // 禁用状态
  if (item.disabled) {
    classes.push('opacity-50', 'cursor-not-allowed')
  }

  return classes.join(' ')
}

// 判断是否为激活状态
const isActive = (item: TabItem) => {
  return activeKey.value === item.key
}

// 处理标签页点击
const handleTabClick = (item: TabItem) => {
  if (item.disabled) return

  activeKey.value = item.key
  emit('change', item.key, item)
}

// 处理标签页关闭
const handleTabClose = (item: TabItem) => {
  emit('close', item.key, item)
}

// 处理添加标签页
const handleAdd = () => {
  emit('add')
}

// 更新指示器位置
const updateIndicator = async () => {
  if (props.type !== 'line' || !tabsNavRef.value) return

  await nextTick()

  const activeTab = tabsNavRef.value.querySelector('.tab-item.text-blue-600') as HTMLElement
  if (!activeTab) return

  const navRect = tabsNavRef.value.getBoundingClientRect()
  const tabRect = activeTab.getBoundingClientRect()

  if (['top', 'bottom'].includes(props.position)) {
    indicatorStyle.value = {
      left: `${tabRect.left - navRect.left}px`,
      width: `${tabRect.width}px`
    }
  } else {
    indicatorStyle.value = {
      top: `${tabRect.top - navRect.top}px`,
      height: `${tabRect.height}px`
    }
  }
}

// 监听激活标签页变化
watch(activeKey, updateIndicator, { immediate: true })
watch(() => props.items, updateIndicator, { deep: true })
</script>

<style scoped>
/* 自定义样式 */
.tabs-container {
  @apply w-full;
}

.tabs-header {
  @apply bg-white;
}

.tabs-nav {
  @apply flex-1;
}

.tab-item {
  @apply whitespace-nowrap;
}

.tabs-content {
  @apply bg-white;
}

.tab-pane {
  @apply w-full;
}

/* 指示器动画 */
.tabs-indicator {
  @apply rounded-full;
}

/* 拖拽排序样式 (预留) */
.sortable-ghost {
  @apply opacity-50;
}

.sortable-chosen {
  @apply transform scale-105;
}
</style>
