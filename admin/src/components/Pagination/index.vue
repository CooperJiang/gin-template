<template>
  <div
    v-if="!hideOnSinglePage || totalPages > 1"
    :class="[
      'flex items-center gap-2',
      positionClass,
      sizeClass,
      { 'opacity-50 pointer-events-none': disabled }
    ]"
  >
    <!-- 总数显示 -->
    <div v-if="showTotal && !simple" class="text-sm text-gray-600 mr-4">
      <slot name="total" :total="total" :range="currentRange">
        {{ totalText }}
      </slot>
    </div>

    <!-- 简洁模式 -->
    <template v-if="simple">
      <button
        :disabled="disabled || current <= 1"
        @click="handlePageChange(current - 1)"
        :class="[
          'pagination-btn',
          buttonSizeClass,
          { 'pagination-btn--disabled': disabled || current <= 1 }
        ]"
      >
        <GlobalIcon name="chevron-left" :size="iconSize" />
      </button>

      <span :class="['text-gray-600 px-3', textSizeClass]">
        {{ current }} / {{ totalPages }}
      </span>

      <button
        :disabled="disabled || current >= totalPages"
        @click="handlePageChange(current + 1)"
        :class="[
          'pagination-btn',
          buttonSizeClass,
          { 'pagination-btn--disabled': disabled || current >= totalPages }
        ]"
      >
        <GlobalIcon name="chevron-right" :size="iconSize" />
      </button>
    </template>

    <!-- 完整模式 -->
    <template v-else>
      <!-- 上一页 -->
      <button
        :disabled="disabled || current <= 1"
        @click="handlePageChange(current - 1)"
        :class="[
          'pagination-btn',
          buttonSizeClass,
          { 'pagination-btn--disabled': disabled || current <= 1 }
        ]"
      >
        <GlobalIcon name="chevron-left" :size="iconSize" />
      </button>

      <!-- 页码列表 -->
      <div class="flex items-center gap-1">
        <!-- 第一页 -->
        <button
          v-if="showFirstLast"
          @click="handlePageChange(1)"
          :class="[
            'pagination-number',
            numberSizeClass,
            { 'pagination-number--active': current === 1 }
          ]"
        >
          1
        </button>

        <!-- 省略号（前） -->
        <span v-if="showPrevEllipsis" :class="['pagination-ellipsis', buttonSizeClass]">
          <GlobalIcon name="ellipsis-horizontal" :size="iconSize" />
        </span>

        <!-- 可见页码 -->
        <button
          v-for="page in visiblePages"
          :key="page"
          @click="handlePageChange(page)"
          :class="[
            'pagination-number',
            numberSizeClass,
            { 'pagination-number--active': current === page }
          ]"
        >
          {{ page }}
        </button>

        <!-- 省略号（后） -->
        <span v-if="showNextEllipsis" :class="['pagination-ellipsis', buttonSizeClass]">
          <GlobalIcon name="ellipsis-horizontal" :size="iconSize" />
        </span>

        <!-- 最后一页 -->
        <button
          v-if="showLastPage"
          @click="handlePageChange(totalPages)"
          :class="[
            'pagination-number',
            numberSizeClass,
            { 'pagination-number--active': current === totalPages }
          ]"
        >
          {{ totalPages }}
        </button>
      </div>

      <!-- 下一页 -->
      <button
        :disabled="disabled || current >= totalPages"
        @click="handlePageChange(current + 1)"
        :class="[
          'pagination-btn',
          buttonSizeClass,
          { 'pagination-btn--disabled': disabled || current >= totalPages }
        ]"
      >
        <GlobalIcon name="chevron-right" :size="iconSize" />
      </button>

      <!-- 每页条数选择器 -->
      <div v-if="showSizeChanger" :class="['flex items-center gap-2 ml-4', textSizeClass]">
        <span class="text-gray-600">每页</span>
        <Select
          :modelValue="pageSize"
          :options="pageSizeOptions"
          :disabled="disabled"
          :size="size"
          @update:modelValue="handleSizeChange"
        />
        <span class="text-gray-600">条</span>
      </div>

      <!-- 快速跳转 -->
      <div v-if="showQuickJumper" :class="['flex items-center gap-2 ml-4', textSizeClass]">
        <span class="text-gray-600">跳至</span>
        <input
          v-model="jumpPageInput"
          @keyup.enter="handleJumpPage"
          @blur="handleJumpPage"
          :disabled="disabled"
          type="number"
          min="1"
          :max="totalPages"
          :class="[
            'pagination-input',
            inputSizeClass,
            { 'pagination-input--disabled': disabled }
          ]"
        />
        <span class="text-gray-600">页</span>
      </div>
    </template>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import type { PaginationProps, PaginationEmits } from './types'
import Select from './components/Select.vue'

defineOptions({
  name: 'GlobalPagination',
})

const props = withDefaults(defineProps<PaginationProps>(), {
  current: 1,
  pageSize: 10,
  pageSizeOptions: () => [10, 20, 50, 100],
  showSizeChanger: true,
  showQuickJumper: false,
  showTotal: true,
  showLessItems: false,
  disabled: false,
  size: 'medium',
  simple: false,
  position: 'right',
  hideOnSinglePage: false,
})

const emit = defineEmits<PaginationEmits>()

// 跳转页面输入
const jumpPageInput = ref<string>('')

// 计算属性
const totalPages = computed(() => Math.ceil(props.total / props.pageSize))

const currentRange = computed((): [number, number] => {
  const start = (props.current - 1) * props.pageSize + 1
  const end = Math.min(props.current * props.pageSize, props.total)
  return [start, end]
})

const totalText = computed(() => {
  if (props.totalTemplate) {
    return props.totalTemplate(props.total, currentRange.value)
  }
  const [start, end] = currentRange.value
  return `显示 ${start}-${end} 条，共 ${props.total} 条`
})

// 样式计算 - 更明显的尺寸区别
const positionClass = computed(() => {
  switch (props.position) {
    case 'left':
      return 'justify-start'
    case 'center':
      return 'justify-center'
    default:
      return 'justify-end'
  }
})

const sizeClass = computed(() => {
  switch (props.size) {
    case 'small':
      return 'gap-1'
    case 'large':
      return 'gap-3'
    default:
      return 'gap-2'
  }
})

const textSizeClass = computed(() => {
  switch (props.size) {
    case 'small':
      return 'text-xs'
    case 'large':
      return 'text-base'
    default:
      return 'text-sm'
  }
})

const buttonSizeClass = computed(() => {
  switch (props.size) {
    case 'small':
      return 'w-6 h-6'
    case 'large':
      return 'w-10 h-10'
    default:
      return 'w-8 h-8'
  }
})

const numberSizeClass = computed(() => {
  switch (props.size) {
    case 'small':
      return 'min-w-6 h-6 px-1 text-xs'
    case 'large':
      return 'min-w-10 h-10 px-3 text-base'
    default:
      return 'min-w-8 h-8 px-2 text-sm'
  }
})

const inputSizeClass = computed(() => {
  switch (props.size) {
    case 'small':
      return 'w-10 px-1 py-0.5 text-xs'
    case 'large':
      return 'w-16 px-3 py-2 text-base'
    default:
      return 'w-12 px-2 py-1 text-sm'
  }
})

const iconSize = computed(() => {
  switch (props.size) {
    case 'small':
      return 'xs' as const
    case 'large':
      return 'md' as const
    default:
      return 'sm' as const
  }
})

// 优化可见页码计算 - 减少中间页码数量
const visiblePages = computed(() => {
  const delta = props.showLessItems ? 1 : 2
  const range = []

  for (
    let i = Math.max(2, props.current - delta);
    i <= Math.min(totalPages.value - 1, props.current + delta);
    i++
  ) {
    range.push(i)
  }

  return range
})

const showFirstLast = computed(() => {
  return totalPages.value > 1 && (props.current > 3 || totalPages.value > 5)
})

const showPrevEllipsis = computed(() => {
  return props.current > (props.showLessItems ? 3 : 4)
})

const showNextEllipsis = computed(() => {
  return props.current < totalPages.value - (props.showLessItems ? 2 : 3)
})

const showLastPage = computed(() => {
  return totalPages.value > 1 &&
         props.current < totalPages.value - (props.showLessItems ? 1 : 2) &&
         totalPages.value > (props.showLessItems ? 3 : 5)
})

// 事件处理
const handlePageChange = (page: number) => {
  if (page === props.current || page < 1 || page > totalPages.value) {
    return
  }
  emit('change', page, props.pageSize)
}

const handleSizeChange = (size: number) => {
  const newCurrent = Math.min(props.current, Math.ceil(props.total / size))
  emit('showSizeChange', newCurrent, size)
  emit('change', newCurrent, size)
}

const handleJumpPage = () => {
  const page = parseInt(jumpPageInput.value)
  if (!isNaN(page) && page >= 1 && page <= totalPages.value) {
    handlePageChange(page)
  }
  jumpPageInput.value = ''
}

// 监听当前页变化，清空跳转输入
watch(() => props.current, () => {
  jumpPageInput.value = ''
})
</script>

<style scoped>
.pagination-btn {
  @apply flex items-center justify-center rounded border border-gray-300 bg-white text-gray-600 hover:bg-gray-50 hover:border-gray-400 transition-colors duration-200;
}

.pagination-btn--disabled {
  @apply opacity-50 cursor-not-allowed hover:bg-white hover:border-gray-300;
}

.pagination-number {
  @apply flex items-center justify-center rounded border border-gray-300 bg-white text-gray-700 hover:bg-gray-50 hover:border-gray-400 transition-colors duration-200;
}

.pagination-number--active {
  @apply border-blue-500 bg-blue-500 text-white hover:bg-blue-600 hover:border-blue-600;
}

.pagination-ellipsis {
  @apply flex items-center justify-center text-gray-400;
}

.pagination-input {
  @apply border border-gray-300 rounded text-center bg-white text-gray-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent;
}

.pagination-input--disabled {
  @apply opacity-50 cursor-not-allowed;
}
</style>
