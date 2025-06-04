<template>
  <div class="relative w-full">
    <!-- 文本域 -->
    <textarea
      ref="textareaRef"
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
      :rows="computedRows"
      :class="textareaClass"
      :style="textareaStyle"
      @input="handleInput"
      @blur="handleBlur"
      @focus="handleFocus"
      @keydown="handleKeydown"
    />

    <!-- 字符计数 -->
    <div
      v-if="showCount && maxlength"
      class="absolute bottom-2 right-2 text-xs text-gray-400 bg-white px-1"
    >
      {{ currentLength }}/{{ maxlength }}
    </div>

    <!-- 清空按钮 -->
    <button
      v-if="modelValue && !disabled && !readonly"
      type="button"
      class="absolute top-2 right-2 text-gray-400 hover:text-gray-600 transition-colors"
      @click="handleClear"
    >
      <GlobalIcon name="x-mark" size="sm" />
    </button>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, nextTick, watch } from 'vue'
import type { TextareaProps, TextareaEmits, TextareaInstance } from './types'

defineOptions({
  name: 'GlobalTextarea'
})

const props = withDefaults(defineProps<TextareaProps>(), {
  size: 'medium',
  rows: 3,
  disabled: false,
  readonly: false,
  error: false,
  required: false,
  autofocus: false,
  resize: 'vertical',
  showCount: false,
  autosize: false
})

const emit = defineEmits<TextareaEmits>()

// 引用
const textareaRef = ref<HTMLTextAreaElement>()

// 计算属性
const currentLength = computed(() => {
  return props.modelValue ? props.modelValue.length : 0
})

const computedRows = computed(() => {
  if (typeof props.autosize === 'object' && props.autosize.minRows) {
    return props.autosize.minRows
  }
  return props.rows
})

const textareaClass = computed(() => {
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

const textareaStyle = computed(() => {
  const style: Record<string, string> = {}

  // 调整大小
  style.resize = props.resize

  return style
})

// 自动调整高度
const adjustHeight = async () => {
  if (!props.autosize || !textareaRef.value) return

  await nextTick()

  const textarea = textareaRef.value
  const computed = window.getComputedStyle(textarea)
  const borderHeight = parseInt(computed.borderTopWidth) + parseInt(computed.borderBottomWidth)
  const paddingHeight = parseInt(computed.paddingTop) + parseInt(computed.paddingBottom)

  // 重置高度
  textarea.style.height = 'auto'

  // 计算内容高度
  const contentHeight = textarea.scrollHeight - borderHeight - paddingHeight
  const lineHeight = parseInt(computed.lineHeight)

  let rows = Math.ceil(contentHeight / lineHeight)

  // 限制行数
  if (typeof props.autosize === 'object') {
    if (props.autosize.minRows) {
      rows = Math.max(rows, props.autosize.minRows)
    }
    if (props.autosize.maxRows) {
      rows = Math.min(rows, props.autosize.maxRows)
    }
  }

  textarea.style.height = `${rows * lineHeight + paddingHeight + borderHeight}px`
}

// 事件处理
const handleInput = (event: Event) => {
  const target = event.target as HTMLTextAreaElement
  emit('update:modelValue', target.value)
  emit('input', target.value, event)

  if (props.autosize) {
    adjustHeight()
  }
}

const handleBlur = (event: FocusEvent) => {
  emit('blur', event)
}

const handleFocus = (event: FocusEvent) => {
  emit('focus', event)
}

const handleKeydown = (event: KeyboardEvent) => {
  emit('keydown', event)
}

const handleClear = () => {
  emit('update:modelValue', '')
  emit('clear')
  textareaRef.value?.focus()
}

// 实例方法
const focus = () => {
  textareaRef.value?.focus()
}

const blur = () => {
  textareaRef.value?.blur()
}

const select = () => {
  textareaRef.value?.select()
}

// 监听值变化，自动调整高度
watch(() => props.modelValue, () => {
  if (props.autosize) {
    adjustHeight()
  }
}, { immediate: true })

// 暴露实例方法
defineExpose<TextareaInstance>({
  focus,
  blur,
  select
})
</script>

<style scoped>
/* 防止在某些浏览器中出现多余的滚动条 */
textarea {
  overflow-y: hidden;
}

textarea:focus {
  overflow-y: auto;
}
</style>
