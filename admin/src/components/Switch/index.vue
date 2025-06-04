<template>
  <div class="switch-wrapper" :class="{ 'switch-wrapper--disabled': disabled }">
    <!-- 开启时的文本 -->
    <span
      v-if="activeText && !isActive"
      class="switch-text switch-text--inactive"
      :class="textSizeClass"
    >
      {{ activeText }}
    </span>

    <!-- 开关本体 -->
    <button
      ref="switchRef"
      type="button"
      role="switch"
      :aria-checked="isActive"
      :disabled="disabled || loading"
      :name="name"
      :class="switchClass"
      @click="handleClick"
      @keydown.space.prevent="handleKeydown"
    >
      <!-- 开关滑块 -->
      <span :class="thumbClass">
        <!-- 加载图标 -->
        <GlobalIcon
          v-if="loading"
          name="loading"
          :size="iconSize"
          class="animate-spin text-white"
        />
        <!-- 自定义图标 -->
        <GlobalIcon
          v-else-if="showIcon && currentIcon"
          :name="currentIcon"
          :size="iconSize"
          class="text-white"
        />
      </span>

      <!-- 开启状态文本 -->
      <span
        v-if="activeText && isActive"
        class="absolute left-2 text-white font-medium"
        :class="textInnerClass"
      >
        {{ activeText }}
      </span>

      <!-- 关闭状态文本 -->
      <span
        v-if="inactiveText && !isActive"
        class="absolute right-2 text-gray-500 font-medium"
        :class="textInnerClass"
      >
        {{ inactiveText }}
      </span>
    </button>

    <!-- 关闭时的文本 -->
    <span
      v-if="inactiveText && isActive"
      class="switch-text switch-text--active"
      :class="textSizeClass"
    >
      {{ inactiveText }}
    </span>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import type { SwitchProps, SwitchEmits, SwitchInstance } from './types'

defineOptions({
  name: 'GlobalSwitch'
})

const props = withDefaults(defineProps<SwitchProps>(), {
  modelValue: false,
  disabled: false,
  size: 'medium',
  color: 'primary',
  activeValue: true,
  inactiveValue: false,
  showIcon: false,
  loading: false,
  required: false
})

const emit = defineEmits<SwitchEmits>()

// 引用
const switchRef = ref<HTMLButtonElement>()

// 计算属性
const isActive = computed(() => {
  return props.modelValue === props.activeValue
})

const switchClass = computed(() => {
  const classes = [
    'switch',
    'relative inline-flex items-center rounded-full transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2',
    'disabled:opacity-60 disabled:cursor-not-allowed'
  ]

  // 尺寸样式
  const sizeClasses = {
    small: 'h-5 w-9',
    medium: 'h-6 w-11',
    large: 'h-7 w-14'
  }
  classes.push(sizeClasses[props.size])

  // 颜色样式
  if (isActive.value) {
    const colorClasses = {
      primary: 'bg-blue-600 focus:ring-blue-500',
      success: 'bg-green-600 focus:ring-green-500',
      warning: 'bg-yellow-600 focus:ring-yellow-500',
      danger: 'bg-red-600 focus:ring-red-500'
    }
    classes.push(colorClasses[props.color])
  } else {
    classes.push('bg-gray-200 focus:ring-gray-500')
  }

  return classes.join(' ')
})

const thumbClass = computed(() => {
  const classes = [
    'switch-thumb',
    'inline-flex items-center justify-center rounded-full bg-white shadow transform transition-transform duration-200'
  ]

  // 尺寸样式
  const sizeClasses = {
    small: 'h-4 w-4',
    medium: 'h-5 w-5',
    large: 'h-6 w-6'
  }
  classes.push(sizeClasses[props.size])

  // 位置样式
  if (isActive.value) {
    const positionClasses = {
      small: 'translate-x-4',
      medium: 'translate-x-5',
      large: 'translate-x-7'
    }
    classes.push(positionClasses[props.size])
  } else {
    classes.push('translate-x-0')
  }

  return classes.join(' ')
})

const iconSize = computed(() => {
  const iconSizes = {
    small: 'xs',
    medium: 'sm',
    large: 'sm'
  }
  return iconSizes[props.size] as 'xs' | 'sm'
})

const textSizeClass = computed(() => {
  const textSizes = {
    small: 'text-xs',
    medium: 'text-sm',
    large: 'text-base'
  }
  return textSizes[props.size]
})

const textInnerClass = computed(() => {
  const textSizes = {
    small: 'text-xs',
    medium: 'text-xs',
    large: 'text-sm'
  }
  return textSizes[props.size]
})

const currentIcon = computed(() => {
  if (isActive.value && props.activeIcon) {
    return props.activeIcon
  } else if (!isActive.value && props.inactiveIcon) {
    return props.inactiveIcon
  } else if (props.showIcon) {
    return isActive.value ? 'check' : 'x-mark'
  }
  return null
})

// 事件处理
const handleClick = (event: MouseEvent) => {
  if (props.disabled || props.loading) return

  const newValue = isActive.value ? props.inactiveValue : props.activeValue
  emit('update:modelValue', newValue)
  emit('change', newValue)
  emit('click', event)
}

const handleKeydown = (event: KeyboardEvent) => {
  if (props.disabled || props.loading) return

  const newValue = isActive.value ? props.inactiveValue : props.activeValue
  emit('update:modelValue', newValue)
  emit('change', newValue)
}

// 实例方法
const focus = () => {
  switchRef.value?.focus()
}

const blur = () => {
  switchRef.value?.blur()
}

const toggle = () => {
  if (props.disabled || props.loading) return
  const newValue = isActive.value ? props.inactiveValue : props.activeValue
  emit('update:modelValue', newValue)
  emit('change', newValue)
}

// 暴露实例方法
defineExpose<SwitchInstance>({
  focus,
  blur,
  toggle
})
</script>

<style scoped>
.switch-wrapper {
  @apply inline-flex items-center gap-3;
}

.switch-wrapper--disabled {
  @apply opacity-60;
}

.switch-text {
  @apply font-medium text-gray-700;
}

.switch-text--inactive {
  @apply text-gray-400;
}

.switch-text--active {
  @apply text-gray-400;
}

.switch {
  @apply relative;
}

/* 确保开关内的文本不会被选中 */
.switch span {
  @apply select-none;
}
</style>
