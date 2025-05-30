<template>
  <span
    ref="tagRef"
    :class="tagClass"
    :style="tagStyle"
    @click="handleClick"
    @keydown.enter="handleClick"
    @keydown.space.prevent="handleClick"
    :tabindex="checkable ? 0 : -1"
  >
    <!-- 前缀图标 -->
    <GlobalIcon
      v-if="icon"
      :name="icon"
      :size="iconSize"
      class="tag-icon"
    />

    <!-- 标签内容 -->
    <span class="tag-content">
      <slot>{{ text }}</slot>
    </span>

    <!-- 关闭按钮 -->
    <button
      v-if="closable && !disabled"
      type="button"
      class="tag-close"
      :class="closeButtonClass"
      @click.stop="handleClose"
      @keydown.enter.stop="handleClose"
      @keydown.space.stop.prevent="handleClose"
    >
      <GlobalIcon
        :name="closeIcon || 'x-mark'"
        :size="closeIconSize"
      />
    </button>
  </span>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import type { TagProps, TagEmits, TagInstance } from './types'

defineOptions({
  name: 'GlobalTag'
})

const props = withDefaults(defineProps<TagProps>(), {
  type: 'default',
  size: 'medium',
  variant: 'filled',
  closable: false,
  disabled: false,
  round: false,
  checkable: false,
  checked: false
})

const emit = defineEmits<TagEmits>()

// 引用
const tagRef = ref<HTMLSpanElement>()

// 计算属性
const tagClass = computed(() => {
  const classes = [
    'tag',
    'inline-flex items-center gap-1 font-medium transition-all duration-200 select-none'
  ]

  // 尺寸样式
  const sizeClasses = {
    small: 'px-2 py-0.5 text-xs',
    medium: 'px-2.5 py-1 text-sm',
    large: 'px-3 py-1.5 text-base'
  }
  classes.push(sizeClasses[props.size])

  // 圆角样式
  if (props.round) {
    classes.push('rounded-full')
  } else {
    classes.push('rounded')
  }

  // 状态样式
  if (props.disabled) {
    classes.push('opacity-60 cursor-not-allowed')
  } else if (props.checkable) {
    classes.push('cursor-pointer hover:shadow-sm')
  }

  // 选中状态（当可选择时）
  if (props.checkable && props.checked) {
    classes.push('ring-2 ring-offset-1')
  }

  // 类型和变体样式
  const typeVariantClass = getTypeVariantClass()
  classes.push(typeVariantClass)

  return classes.join(' ')
})

const tagStyle = computed(() => {
  const style: Record<string, string> = {}

  if (props.maxWidth) {
    style.maxWidth = typeof props.maxWidth === 'number' ? `${props.maxWidth}px` : props.maxWidth
  }

  if (props.color) {
    if (props.variant === 'filled') {
      style.backgroundColor = props.color
      style.color = 'white'
      style.borderColor = props.color
    } else if (props.variant === 'outlined') {
      style.borderColor = props.color
      style.color = props.color
    } else if (props.variant === 'ghost') {
      style.color = props.color
      style.backgroundColor = `${props.color}20` // 20% opacity
    }
  }

  return style
})

const iconSize = computed(() => {
  const iconSizes = {
    small: 'xs',
    medium: 'sm',
    large: 'sm'
  }
  return iconSizes[props.size] as 'xs' | 'sm'
})

const closeIconSize = computed(() => {
  const iconSizes = {
    small: 'xs',
    medium: 'xs',
    large: 'sm'
  }
  return iconSizes[props.size] as 'xs' | 'sm'
})

const closeButtonClass = computed(() => {
  const classes = [
    'inline-flex items-center justify-center rounded-full transition-colors duration-200 hover:bg-black hover:bg-opacity-10'
  ]

  const sizeClasses = {
    small: 'w-3 h-3',
    medium: 'w-4 h-4',
    large: 'w-5 h-5'
  }
  classes.push(sizeClasses[props.size])

  return classes.join(' ')
})

// 获取类型和变体组合的样式类
const getTypeVariantClass = () => {
  if (props.color) {
    // 如果有自定义颜色，则使用基础样式
    if (props.variant === 'filled') {
      return 'text-white border border-transparent'
    } else if (props.variant === 'outlined') {
      return 'bg-white border'
    } else if (props.variant === 'ghost') {
      return 'border border-transparent'
    }
  }

  const typeStyles = {
    default: {
      filled: 'bg-gray-500 text-white border border-gray-500',
      outlined: 'bg-white text-gray-700 border border-gray-300',
      ghost: 'bg-gray-100 text-gray-700 border border-transparent'
    },
    primary: {
      filled: 'bg-blue-600 text-white border border-blue-600',
      outlined: 'bg-white text-blue-600 border border-blue-600',
      ghost: 'bg-blue-50 text-blue-600 border border-transparent'
    },
    success: {
      filled: 'bg-green-600 text-white border border-green-600',
      outlined: 'bg-white text-green-600 border border-green-600',
      ghost: 'bg-green-50 text-green-600 border border-transparent'
    },
    warning: {
      filled: 'bg-yellow-600 text-white border border-yellow-600',
      outlined: 'bg-white text-yellow-600 border border-yellow-600',
      ghost: 'bg-yellow-50 text-yellow-600 border border-transparent'
    },
    danger: {
      filled: 'bg-red-600 text-white border border-red-600',
      outlined: 'bg-white text-red-600 border border-red-600',
      ghost: 'bg-red-50 text-red-600 border border-transparent'
    },
    info: {
      filled: 'bg-cyan-600 text-white border border-cyan-600',
      outlined: 'bg-white text-cyan-600 border border-cyan-600',
      ghost: 'bg-cyan-50 text-cyan-600 border border-transparent'
    }
  }

  return typeStyles[props.type][props.variant]
}

// 事件处理
const handleClick = (event: MouseEvent | KeyboardEvent) => {
  if (props.disabled) return

  if (props.checkable) {
    const newChecked = !props.checked
    emit('update:checked', newChecked)
    emit('change', newChecked, event as MouseEvent)
  }

  emit('click', event as MouseEvent)
}

const handleClose = (event: MouseEvent | KeyboardEvent) => {
  if (props.disabled) return
  emit('close', event as MouseEvent)
}

// 实例方法
const focus = () => {
  tagRef.value?.focus()
}

const blur = () => {
  tagRef.value?.blur()
}

// 暴露实例方法
defineExpose<TagInstance>({
  focus,
  blur
})
</script>

<style scoped>
.tag-content {
  @apply truncate;
}

.tag:focus {
  @apply outline-none ring-2 ring-offset-1 ring-blue-500;
}
</style>
