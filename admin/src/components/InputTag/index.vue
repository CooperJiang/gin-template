<template>
  <div :class="wrapperClass">
    <!-- 标签列表 -->
    <div class="input-tag-tags" :class="tagsClass">
      <GlobalTag
        v-for="(tag, index) in modelValue"
        :key="`${tag}-${index}`"
        :text="tag"
        :type="tagType"
        :size="tagSize"
        :variant="tagVariant"
        :closable="closable && !disabled && !readonly"
        @close="removeTag(index)"
      />

      <!-- 输入框 -->
      <div class="input-tag-input-wrapper" v-if="!readonly">
        <input
          ref="inputRef"
          v-model="inputValue"
          :type="inputType"
          :placeholder="currentPlaceholder"
          :disabled="disabled"
          :autofocus="autofocus"
          :class="inputClass"
          @input="handleInput"
          @keydown="handleKeydown"
          @focus="handleFocus"
          @blur="handleBlur"
          @paste="handlePaste"
        />
      </div>
    </div>

    <!-- 右侧操作区域 -->
    <div class="input-tag-actions" v-if="!readonly">
      <!-- 自定义图标 -->
      <GlobalIcon
        v-if="icon"
        :name="icon"
        :size="iconSize"
        class="text-gray-400"
      />

      <!-- 计数显示 -->
      <span v-if="showCount" :class="countClass">
        {{ modelValue?.length || 0 }}{{ max ? `/${max}` : '' }}
      </span>
    </div>

    <!-- 建议列表 -->
    <div
      v-if="showSuggestionList"
      class="input-tag-suggestions"
      :class="suggestionsClass"
    >
      <div
        v-for="(suggestion, index) in filteredSuggestions"
        :key="suggestion"
        :class="suggestionItemClass(index)"
        @click="addSuggestionTag(suggestion)"
        @mouseenter="activeSuggestionIndex = index"
      >
        {{ suggestion }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, nextTick, watch } from 'vue'
import type { InputTagProps, InputTagEmits, InputTagInstance } from './types'

defineOptions({
  name: 'GlobalInputTag'
})

const props = withDefaults(defineProps<InputTagProps>(), {
  modelValue: () => [],
  placeholder: '请输入标签',
  disabled: false,
  readonly: false,
  size: 'medium',
  allowDuplicates: false,
  separator: ',',
  tagType: 'primary',
  tagSize: 'small',
  tagVariant: 'filled',
  closable: true,
  inputType: 'text',
  autofocus: false,
  minTagLength: 1,
  sortable: false,
  error: false,
  showCount: false,
  suggestions: () => [],
  showSuggestions: true
})

const emit = defineEmits<InputTagEmits>()

// 引用
const inputRef = ref<HTMLInputElement>()

// 响应式数据
const inputValue = ref('')
const activeSuggestionIndex = ref(-1)
const showSuggestionList = ref(false)

// 计算属性
const wrapperClass = computed(() => {
  const classes = [
    'input-tag-wrapper',
    'relative w-full border rounded-md transition-all duration-200 focus-within:ring-2 focus-within:ring-blue-500 focus-within:ring-offset-1'
  ]

  // 尺寸样式
  const sizeClasses = {
    small: 'min-h-[32px]',
    medium: 'min-h-[40px]',
    large: 'min-h-[48px]'
  }
  classes.push(sizeClasses[props.size])

  // 状态样式
  if (props.disabled) {
    classes.push('bg-gray-100 border-gray-300 cursor-not-allowed')
  } else if (props.readonly) {
    classes.push('bg-gray-50 border-gray-300')
  } else if (props.error) {
    classes.push('border-red-500 focus-within:ring-red-500')
  } else {
    classes.push('bg-white border-gray-300 hover:border-gray-400')
  }

  return classes.join(' ')
})

const tagsClass = computed(() => {
  const classes = ['flex flex-wrap items-center gap-1 flex-1']

  // 内边距
  const paddingClasses = {
    small: 'p-1',
    medium: 'p-2',
    large: 'p-3'
  }
  classes.push(paddingClasses[props.size])

  return classes.join(' ')
})

const inputClass = computed(() => {
  const classes = [
    'outline-none bg-transparent border-0 flex-1 min-w-[80px]'
  ]

  // 文字尺寸
  const textSizeClasses = {
    small: 'text-sm',
    medium: 'text-sm',
    large: 'text-base'
  }
  classes.push(textSizeClasses[props.size])

  if (props.disabled) {
    classes.push('cursor-not-allowed')
  }

  return classes.join(' ')
})

const countClass = computed(() => {
  const classes = ['text-xs select-none']

  if (props.max && (modelValue.value?.length || 0) >= props.max) {
    classes.push('text-red-500')
  } else {
    classes.push('text-gray-500')
  }

  return classes.join(' ')
})

const suggestionsClass = computed(() => {
  return [
    'absolute top-full left-0 right-0 z-10 mt-1 bg-white border border-gray-300 rounded-md shadow-lg max-h-48 overflow-y-auto'
  ].join(' ')
})

const iconSize = computed(() => {
  const iconSizes = {
    small: 'xs',
    medium: 'sm',
    large: 'sm'
  }
  return iconSizes[props.size] as 'xs' | 'sm'
})

const currentPlaceholder = computed(() => {
  if (modelValue.value?.length) {
    return ''
  }
  return props.placeholder
})

const filteredSuggestions = computed(() => {
  if (!props.showSuggestions || !inputValue.value.trim()) {
    return []
  }

  const input = inputValue.value.toLowerCase()
  return props.suggestions
    .filter(suggestion =>
      suggestion.toLowerCase().includes(input) &&
      (props.allowDuplicates || !modelValue.value?.includes(suggestion))
    )
    .slice(0, 10) // 限制显示数量
})

const modelValue = computed(() => props.modelValue || [])

// 方法
const suggestionItemClass = (index: number) => {
  const classes = [
    'px-3 py-2 cursor-pointer transition-colors duration-200'
  ]

  if (index === activeSuggestionIndex.value) {
    classes.push('bg-blue-100 text-blue-700')
  } else {
    classes.push('hover:bg-gray-100')
  }

  return classes.join(' ')
}

const validateTag = (tag: string): string | true => {
  if (!tag.trim()) {
    return '标签不能为空'
  }

  if (props.minTagLength && tag.length < props.minTagLength) {
    return `标签长度不能少于${props.minTagLength}个字符`
  }

  if (props.maxTagLength && tag.length > props.maxTagLength) {
    return `标签长度不能超过${props.maxTagLength}个字符`
  }

  if (!props.allowDuplicates && modelValue.value.includes(tag)) {
    return '标签已存在'
  }

  if (props.max && modelValue.value.length >= props.max) {
    return `最多只能添加${props.max}个标签`
  }

  if (props.validate) {
    const result = props.validate(tag)
    if (result !== true) {
      return typeof result === 'string' ? result : '标签格式不正确'
    }
  }

  return true
}

const addTag = (tag: string): boolean => {
  const trimmedTag = tag.trim()
  if (!trimmedTag) return false

  const validation = validateTag(trimmedTag)
  if (validation !== true) {
    emit('invalid', trimmedTag, validation)
    return false
  }

  if (props.max && modelValue.value.length >= props.max) {
    emit('max-reached', modelValue.value)
    return false
  }

  const newTags = [...modelValue.value, trimmedTag]
  emit('update:modelValue', newTags)
  emit('add', trimmedTag)

  inputValue.value = ''
  hideSuggestions()
  return true
}

const removeTag = (index: number) => {
  const tag = modelValue.value[index]
  const newTags = modelValue.value.filter((_, i) => i !== index)
  emit('update:modelValue', newTags)
  emit('remove', tag, index)
}

const addSuggestionTag = (tag: string) => {
  if (addTag(tag)) {
    inputRef.value?.focus()
  }
}

const showSuggestions = () => {
  if (filteredSuggestions.value.length > 0) {
    showSuggestionList.value = true
    activeSuggestionIndex.value = -1
  }
}

const hideSuggestions = () => {
  showSuggestionList.value = false
  activeSuggestionIndex.value = -1
}

// 事件处理
const handleInput = (event: Event) => {
  const target = event.target as HTMLInputElement
  const value = target.value

  // 检查分隔符
  if (typeof props.separator === 'string' && value.includes(props.separator)) {
    const tags = value.split(props.separator)
    const lastTag = tags.pop() || ''

    tags.forEach(tag => {
      if (tag.trim()) {
        addTag(tag.trim())
      }
    })

    inputValue.value = lastTag
  } else if (props.separator instanceof RegExp && props.separator.test(value)) {
    const tags = value.split(props.separator)
    const lastTag = tags.pop() || ''

    tags.forEach(tag => {
      if (tag.trim()) {
        addTag(tag.trim())
      }
    })

    inputValue.value = lastTag
  }

  emit('input', inputValue.value)

  if (props.showSuggestions) {
    nextTick(() => {
      if (inputValue.value.trim()) {
        showSuggestions()
      } else {
        hideSuggestions()
      }
    })
  }
}

const handleKeydown = (event: KeyboardEvent) => {
  if (props.disabled || props.readonly) return

  switch (event.key) {
    case 'Enter':
      event.preventDefault()
      if (showSuggestionList.value && activeSuggestionIndex.value >= 0) {
        addSuggestionTag(filteredSuggestions.value[activeSuggestionIndex.value])
      } else if (inputValue.value.trim()) {
        addTag(inputValue.value.trim())
      }
      break

    case 'Backspace':
      if (!inputValue.value && modelValue.value.length > 0) {
        removeTag(modelValue.value.length - 1)
      }
      break

    case 'ArrowDown':
      if (showSuggestionList.value) {
        event.preventDefault()
        activeSuggestionIndex.value = Math.min(
          activeSuggestionIndex.value + 1,
          filteredSuggestions.value.length - 1
        )
      }
      break

    case 'ArrowUp':
      if (showSuggestionList.value) {
        event.preventDefault()
        activeSuggestionIndex.value = Math.max(activeSuggestionIndex.value - 1, -1)
      }
      break

    case 'Escape':
      hideSuggestions()
      break
  }
}

const handleFocus = (event: FocusEvent) => {
  emit('focus', event)
  if (props.showSuggestions && inputValue.value.trim()) {
    showSuggestions()
  }
}

const handleBlur = (event: FocusEvent) => {
  // 延迟隐藏建议，允许点击建议项
  setTimeout(() => {
    hideSuggestions()
  }, 200)

  // 如果输入框有内容，尝试添加标签
  if (inputValue.value.trim()) {
    addTag(inputValue.value.trim())
  }

  emit('blur', event)
}

const handlePaste = (event: ClipboardEvent) => {
  event.preventDefault()
  const pasteData = event.clipboardData?.getData('text') || ''

  if (typeof props.separator === 'string') {
    const tags = pasteData.split(props.separator)
    tags.forEach(tag => {
      const trimmedTag = tag.trim()
      if (trimmedTag) {
        addTag(trimmedTag)
      }
    })
  } else {
    addTag(pasteData.trim())
  }
}

// 实例方法
const focus = () => {
  inputRef.value?.focus()
}

const blur = () => {
  inputRef.value?.blur()
}

const clear = () => {
  emit('update:modelValue', [])
  inputValue.value = ''
  hideSuggestions()
}

const getInputValue = () => {
  return inputValue.value
}

const setInputValue = (value: string) => {
  inputValue.value = value
}

// 监听建议变化
watch(filteredSuggestions, (newSuggestions) => {
  if (newSuggestions.length === 0) {
    hideSuggestions()
  }
})

// 暴露实例方法
defineExpose<InputTagInstance>({
  focus,
  blur,
  addTag,
  removeTag,
  clear,
  getInputValue,
  setInputValue
})
</script>

<style scoped>
.input-tag-wrapper {
  display: flex;
  align-items: stretch;
}

.input-tag-actions {
  @apply flex items-center gap-2 px-2;
}

.input-tag-input-wrapper {
  @apply flex-1 min-w-0;
}

/* 建议列表滚动条样式 */
.input-tag-suggestions::-webkit-scrollbar {
  width: 6px;
}

.input-tag-suggestions::-webkit-scrollbar-track {
  background: #f1f5f9;
}

.input-tag-suggestions::-webkit-scrollbar-thumb {
  background: #cbd5e1;
  border-radius: 3px;
}

.input-tag-suggestions::-webkit-scrollbar-thumb:hover {
  background: #94a3b8;
}
</style>
