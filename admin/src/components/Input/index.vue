<template>
  <div
    :class="[
      'relative w-full',
      {
        'input-wrapper--disabled': disabled,
        'input-wrapper--error': error,
        'input-wrapper--small': size === 'small',
        'input-wrapper--medium': size === 'medium',
        'input-wrapper--large': size === 'large'
      }
    ]"
  >
    <!-- 前缀图标 -->
    <div
      v-if="prefixIcon"
      class="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 pointer-events-none"
    >
      <GlobalIcon :name="prefixIcon" size="sm" />
    </div>

    <!-- 输入框 -->
    <input
      ref="inputRef"
      :type="actualType"
      :value="modelValue"
      :placeholder="placeholder"
      :disabled="disabled"
      :readonly="readonly"
      :maxlength="maxlength"
      :minlength="minlength"
      :required="required"
      :autocomplete="autocomplete"
      :autofocus="autofocus"
      :name="name"
      :class="inputClass"
      @input="handleInput"
      @blur="handleBlur"
      @focus="handleFocus"
      @keydown="handleKeydown"
    />

    <!-- 密码显示切换按钮 -->
    <button
      v-if="type === 'password' && showPassword"
      type="button"
      class="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-600 transition-colors"
      @click="togglePasswordVisibility"
    >
      <GlobalIcon :name="showPasswordText ? 'eye-slash' : 'eye'" size="sm" />
    </button>

    <!-- 后缀图标 -->
    <div
      v-else-if="suffixIcon"
      class="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 pointer-events-none"
    >
      <GlobalIcon :name="suffixIcon" size="sm" />
    </div>

    <!-- 清空按钮 -->
    <button
      v-else-if="modelValue && !disabled && !readonly && type !== 'password'"
      type="button"
      class="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-600 transition-colors"
      @click="handleClear"
    >
      <GlobalIcon name="x-mark" size="sm" />
    </button>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import type { InputProps, InputEmits, InputInstance } from './types'

defineOptions({
  name: 'GlobalInput'
})

const props = withDefaults(defineProps<InputProps>(), {
  type: 'text',
  size: 'medium',
  disabled: false,
  readonly: false,
  error: false,
  required: false,
  showPassword: false,
  autofocus: false
})

const emit = defineEmits<InputEmits>()

// 引用
const inputRef = ref<HTMLInputElement>()

// 内部状态
const showPasswordText = ref(false)

// 计算属性
const actualType = computed(() => {
  if (props.type === 'password' && props.showPassword) {
    return showPasswordText.value ? 'text' : 'password'
  }
  return props.type
})

const inputClass = computed(() => {
  const classes = [
    'w-full border rounded-md transition-all duration-200 focus:outline-none focus:ring-2',
    'disabled:bg-gray-50 disabled:text-gray-500 disabled:cursor-not-allowed',
    'readonly:bg-gray-50 readonly:cursor-default'
  ]

  // 尺寸样式
  const sizeClasses = {
    small: 'px-3 py-1.5 text-sm',
    medium: 'px-3 py-2 text-sm',
    large: 'px-4 py-3 text-base'
  }
  classes.push(sizeClasses[props.size])

  // 前缀图标时的左边距
  if (props.prefixIcon) {
    classes.push('pl-10')
  }

  // 后缀元素时的右边距
  if (props.suffixIcon || (props.type === 'password' && props.showPassword) || (!props.disabled && !props.readonly)) {
    classes.push('pr-10')
  }

  // 状态样式
  if (props.error) {
    classes.push('border-red-300 focus:border-red-500 focus:ring-red-500')
  } else if (props.disabled) {
    classes.push('border-gray-300')
  } else {
    classes.push('border-gray-300 focus:border-blue-500 focus:ring-blue-500')
  }

  return classes.join(' ')
})

// 事件处理
const handleInput = (event: Event) => {
  const target = event.target as HTMLInputElement
  const value = props.type === 'number' ? Number(target.value) : target.value
  emit('update:modelValue', value)
  emit('input', value, event)
}

const handleBlur = (event: FocusEvent) => {
  emit('blur', event)
}

const handleFocus = (event: FocusEvent) => {
  emit('focus', event)
}

const handleKeydown = (event: KeyboardEvent) => {
  emit('keydown', event)
  if (event.key === 'Enter') {
    emit('enter', event)
  }
}

const handleClear = () => {
  emit('update:modelValue', '')
  emit('clear')
  inputRef.value?.focus()
}

const togglePasswordVisibility = () => {
  showPasswordText.value = !showPasswordText.value
}

// 实例方法
const focus = () => {
  inputRef.value?.focus()
}

const blur = () => {
  inputRef.value?.blur()
}

const select = () => {
  inputRef.value?.select()
}

// 暴露实例方法
defineExpose<InputInstance>({
  focus,
  blur,
  select
})
</script>

<style scoped>
.input-wrapper--disabled {
  @apply opacity-60;
}
</style>
