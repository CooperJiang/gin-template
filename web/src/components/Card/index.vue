<template>
  <div
    :class="cardClass"
    @click="handleClick"
    @mouseenter="handleMouseEnter"
    @mouseleave="handleMouseLeave"
  >
    <!-- 加载状态遮罩 -->
    <div v-if="loading" :class="loadingOverlayClass">
      <slot name="loading">
        <div class="flex items-center justify-center">
          <GlobalIcon name="arrow-path" size="lg" class="animate-spin text-blue-500" />
        </div>
      </slot>
    </div>

    <!-- 封面图片 -->
    <div v-if="$slots.cover" class="card-cover overflow-hidden">
      <slot name="cover" />
    </div>

    <!-- 卡片头部 -->
    <div v-if="hasHeader" :class="headerClass">
      <slot name="header">
        <div class="flex items-center justify-between">
          <div class="flex-1 min-w-0">
            <!-- 标题 -->
            <div v-if="title || $slots.title" class="card-title">
              <slot name="title">
                <h3 class="text-lg font-semibold text-gray-900 truncate">{{ title }}</h3>
              </slot>
            </div>

            <!-- 副标题 -->
            <div v-if="subtitle || $slots.subtitle" class="card-subtitle mt-1">
              <slot name="subtitle">
                <p class="text-sm text-gray-500 truncate">{{ subtitle }}</p>
              </slot>
            </div>
          </div>

          <!-- 额外操作 -->
          <div v-if="$slots.extra" class="card-extra flex-shrink-0 ml-4">
            <slot name="extra" />
          </div>
        </div>
      </slot>
    </div>

    <!-- 卡片内容 -->
    <div :class="bodyClass">
      <slot />
    </div>

    <!-- 卡片底部 -->
    <div v-if="$slots.footer" :class="footerClass">
      <slot name="footer" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { CardProps, CardEmits } from './types'

defineOptions({
  name: 'GlobalCard'
})

const props = withDefaults(defineProps<CardProps>(), {
  size: 'medium',
  shadow: 'small',
  bordered: true,
  hoverable: false,
  background: 'white',
  headerDivider: true,
  footerDivider: true,
  rounded: 'medium',
  headerPadding: 'medium',
  bodyPadding: 'medium',
  footerPadding: 'medium',
  loading: false
})

const emit = defineEmits<CardEmits>()

// 判断是否有头部内容
const hasHeader = computed(() => {
  return props.title || props.subtitle || slots.header || slots.title || slots.subtitle || slots.extra
})

// 获取插槽
const slots = defineSlots()

// 卡片主容器样式
const cardClass = computed(() => {
  const classes = ['card', 'relative', 'transition-all', 'duration-200']

  // 背景色
  const backgroundClasses = {
    white: 'bg-white',
    gray: 'bg-gray-50',
    blue: 'bg-blue-50',
    green: 'bg-green-50',
    yellow: 'bg-yellow-50',
    red: 'bg-red-50',
    purple: 'bg-purple-50',
    indigo: 'bg-indigo-50'
  }
  classes.push(backgroundClasses[props.background])

  // 圆角
  const roundedClasses = {
    none: 'rounded-none',
    small: 'rounded',
    medium: 'rounded-lg',
    large: 'rounded-xl',
    full: 'rounded-2xl'
  }
  classes.push(roundedClasses[props.rounded])

  // 边框
  if (props.bordered) {
    classes.push('border', 'border-gray-200')
  }

  // 阴影
  if (props.shadow !== 'none') {
    const shadowClasses = {
      small: 'shadow-sm',
      medium: 'shadow-md',
      large: 'shadow-lg',
      hover: 'shadow-sm hover:shadow-md'
    }
    classes.push(shadowClasses[props.shadow])
  }

  // 可悬停效果
  if (props.hoverable) {
    classes.push('cursor-pointer', 'hover:shadow-lg', 'hover:-translate-y-1')
  }

  // 加载状态
  if (props.loading) {
    classes.push('overflow-hidden')
  }

  return classes.join(' ')
})

// 加载遮罩样式
const loadingOverlayClass = computed(() => {
  return [
    'absolute inset-0 bg-white bg-opacity-75 flex items-center justify-center z-10',
    props.rounded !== 'none' ? roundedClasses[props.rounded] : ''
  ].join(' ')
})

// 头部样式
const headerClass = computed(() => {
  const classes = ['card-header']

  // 内边距
  const paddingClasses = {
    none: '',
    small: 'p-3',
    medium: 'p-4',
    large: 'p-6'
  }
  classes.push(paddingClasses[props.headerPadding])

  // 分割线
  if (props.headerDivider) {
    classes.push('border-b', 'border-gray-200')
  }

  return classes.join(' ')
})

// 内容区域样式
const bodyClass = computed(() => {
  const classes = ['card-body']

  // 内边距
  const paddingClasses = {
    none: '',
    small: 'p-3',
    medium: 'p-4',
    large: 'p-6'
  }
  classes.push(paddingClasses[props.bodyPadding])

  return classes.join(' ')
})

// 底部样式
const footerClass = computed(() => {
  const classes = ['card-footer']

  // 内边距
  const paddingClasses = {
    none: '',
    small: 'p-3',
    medium: 'p-4',
    large: 'p-6'
  }
  classes.push(paddingClasses[props.footerPadding])

  // 分割线
  if (props.footerDivider) {
    classes.push('border-t', 'border-gray-200')
  }

  return classes.join(' ')
})

// 圆角类名映射
const roundedClasses = {
  none: 'rounded-none',
  small: 'rounded',
  medium: 'rounded-lg',
  large: 'rounded-xl',
  full: 'rounded-2xl'
}

// 事件处理
const handleClick = (event: MouseEvent) => {
  if (!props.loading) {
    emit('click', event)
  }
}

const handleMouseEnter = (event: MouseEvent) => {
  emit('mouseenter', event)
}

const handleMouseLeave = (event: MouseEvent) => {
  emit('mouseleave', event)
}
</script>

<style scoped>
/* 卡片样式 */
.card {
  @apply overflow-hidden;
}

.card-cover {
  @apply first:rounded-t-lg last:rounded-b-lg;
}

.card-header {
  @apply bg-inherit;
}

.card-body {
  @apply flex-1 bg-inherit;
}

.card-footer {
  @apply bg-inherit;
}

.card-title {
  @apply flex-1 min-w-0;
}

.card-subtitle {
  @apply flex-1 min-w-0;
}

.card-extra {
  @apply flex items-center;
}

/* 悬停效果 */
.card:hover .card-cover img {
  @apply scale-105;
  transition: transform 0.3s ease;
}

/* 加载状态下禁用所有交互 */
.card:has(.absolute.inset-0) {
  @apply pointer-events-none;
}
</style>
