<template>
  <button
    :type="htmlType"
    :class="buttonClass"
    :disabled="disabled || loading"
    @click="handleClick"
  >
    <!-- 加载状态图标 -->
    <GlobalIcon
      v-if="loading"
      name="arrow-path"
      size="sm"
      class="animate-spin"
      :class="{ 'mr-2': $slots.default && iconPosition === 'left' }"
    />

    <!-- 左侧图标 -->
    <GlobalIcon
      v-if="!loading && icon && iconPosition === 'left'"
      :name="icon"
      size="sm"
      :variant="iconVariant"
      class="mr-2"
    />

    <!-- 按钮内容 -->
    <span v-if="$slots.default">
      <slot />
    </span>

    <!-- 右侧图标 -->
    <GlobalIcon
      v-if="!loading && icon && iconPosition === 'right'"
      :name="icon"
      size="sm"
      :variant="iconVariant"
      class="ml-2"
    />
  </button>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { ButtonProps, ButtonEmits } from './types'

defineOptions({
  name: 'GlobalButton'
})

const props = withDefaults(defineProps<ButtonProps>(), {
  type: 'primary',
  size: 'medium',
  shape: 'default',
  variant: 'solid',
  loading: false,
  disabled: false,
  block: false,
  iconPosition: 'left',
  iconVariant: 'outline',
  htmlType: 'button'
})

const emit = defineEmits<ButtonEmits>()

// 计算按钮样式类
const buttonClass = computed(() => {
  const classes = [
    'inline-flex items-center justify-center font-medium transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2'
  ]

  // 尺寸样式
  const sizeClasses = {
    small: 'px-3 py-1.5 text-sm leading-4',
    medium: 'px-4 py-2 text-sm',
    large: 'px-6 py-3 text-base'
  }
  classes.push(sizeClasses[props.size])

  // 形状样式
  const shapeClasses = {
    default: 'rounded-md',
    round: 'rounded-full',
    circle: 'rounded-full w-10 h-10 p-0'
  }
  classes.push(shapeClasses[props.shape])

  // 类型和变体样式
  const typeVariantClasses = getTypeVariantClasses()
  classes.push(typeVariantClasses)

  // 块级元素
  if (props.block) {
    classes.push('w-full')
  }

  // 禁用状态
  if (props.disabled || props.loading) {
    classes.push('cursor-not-allowed opacity-50')
  } else {
    classes.push('cursor-pointer')
  }

  return classes.join(' ')
})

// 获取类型和变体组合的样式
const getTypeVariantClasses = (): string => {
  const { type, variant } = props

  // 定义颜色映射
  const colorMap = {
    primary: {
      solid: 'bg-blue-600 text-white hover:bg-blue-700 focus:ring-blue-500 border-transparent',
      outline: 'border-blue-600 text-blue-600 bg-transparent hover:bg-blue-50 focus:ring-blue-500',
      ghost: 'text-blue-600 bg-transparent hover:bg-blue-50 focus:ring-blue-500 border-transparent',
      link: 'text-blue-600 bg-transparent hover:text-blue-700 focus:ring-blue-500 border-transparent underline-offset-4 hover:underline'
    },
    secondary: {
      solid: 'bg-gray-600 text-white hover:bg-gray-700 focus:ring-gray-500 border-transparent',
      outline: 'border-gray-600 text-gray-600 bg-transparent hover:bg-gray-50 focus:ring-gray-500',
      ghost: 'text-gray-600 bg-transparent hover:bg-gray-50 focus:ring-gray-500 border-transparent',
      link: 'text-gray-600 bg-transparent hover:text-gray-700 focus:ring-gray-500 border-transparent underline-offset-4 hover:underline'
    },
    success: {
      solid: 'bg-green-600 text-white hover:bg-green-700 focus:ring-green-500 border-transparent',
      outline: 'border-green-600 text-green-600 bg-transparent hover:bg-green-50 focus:ring-green-500',
      ghost: 'text-green-600 bg-transparent hover:bg-green-50 focus:ring-green-500 border-transparent',
      link: 'text-green-600 bg-transparent hover:text-green-700 focus:ring-green-500 border-transparent underline-offset-4 hover:underline'
    },
    warning: {
      solid: 'bg-yellow-600 text-white hover:bg-yellow-700 focus:ring-yellow-500 border-transparent',
      outline: 'border-yellow-600 text-yellow-600 bg-transparent hover:bg-yellow-50 focus:ring-yellow-500',
      ghost: 'text-yellow-600 bg-transparent hover:bg-yellow-50 focus:ring-yellow-500 border-transparent',
      link: 'text-yellow-600 bg-transparent hover:text-yellow-700 focus:ring-yellow-500 border-transparent underline-offset-4 hover:underline'
    },
    danger: {
      solid: 'bg-red-600 text-white hover:bg-red-700 focus:ring-red-500 border-transparent',
      outline: 'border-red-600 text-red-600 bg-transparent hover:bg-red-50 focus:ring-red-500',
      ghost: 'text-red-600 bg-transparent hover:bg-red-50 focus:ring-red-500 border-transparent',
      link: 'text-red-600 bg-transparent hover:text-red-700 focus:ring-red-500 border-transparent underline-offset-4 hover:underline'
    },
    info: {
      solid: 'bg-cyan-600 text-white hover:bg-cyan-700 focus:ring-cyan-500 border-transparent',
      outline: 'border-cyan-600 text-cyan-600 bg-transparent hover:bg-cyan-50 focus:ring-cyan-500',
      ghost: 'text-cyan-600 bg-transparent hover:bg-cyan-50 focus:ring-cyan-500 border-transparent',
      link: 'text-cyan-600 bg-transparent hover:text-cyan-700 focus:ring-cyan-500 border-transparent underline-offset-4 hover:underline'
    }
  }

  const baseClasses = variant === 'link' ? '' : 'border'
  const variantClasses = colorMap[type][variant]

  return `${baseClasses} ${variantClasses}`
}

// 处理点击事件
const handleClick = (event: MouseEvent) => {
  if (!props.disabled && !props.loading) {
    emit('click', event)
  }
}
</script>
