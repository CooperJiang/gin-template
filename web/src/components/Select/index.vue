<template>
  <div
    ref="selectRef"
    :class="[
      'relative w-full',
      {
        'cursor-not-allowed': disabled,
        'cursor-pointer': !disabled
      }
    ]"
  >
    <!-- 选择器输入框 -->
    <div
      ref="triggerRef"
      :class="[
        'flex items-center min-h-[40px] px-3 border rounded-lg transition-all duration-200 focus-within:ring-2 focus-within:ring-blue-500 focus-within:ring-opacity-50',
        sizeClasses,
        variantClasses,
        statusClasses,
        {
          'cursor-not-allowed bg-gray-50': disabled,
          'cursor-pointer': !disabled,
          'border-blue-500': visible && !error && !success,
        }
      ]"
      @click="handleTriggerClick"
      @mousedown.prevent
    >
      <!-- 多选标签 -->
      <div v-if="multiple && selectedOptions.length" class="flex flex-wrap gap-1 flex-1 min-w-0">
        <span
          v-for="option in displayTags"
          :key="option.value"
          class="inline-flex items-center px-2 py-1 bg-blue-100 text-blue-800 text-xs rounded-md"
        >
          <GlobalIcon v-if="option.icon" :name="option.icon" size="xs" class="mr-1" />
          {{ option.label }}
          <button
            v-if="!disabled"
            @click.stop="removeTag(option)"
            @mousedown.stop
            class="ml-1 hover:text-blue-600 focus:outline-none"
          >
            <GlobalIcon name="x-mark" size="xs" />
          </button>
        </span>
        <span
          v-if="hasMoreTags"
          class="inline-flex items-center px-2 py-1 bg-gray-100 text-gray-600 text-xs rounded-md"
        >
          +{{ selectedOptions.length - maxTagCount }}
        </span>
      </div>

      <!-- 单选显示 -->
      <div v-else-if="!multiple && selectedOption" class="flex items-center flex-1 min-w-0">
        <GlobalIcon v-if="selectedOption.icon" :name="selectedOption.icon" size="sm" class="mr-2 flex-shrink-0" />
        <span class="truncate">{{ selectedOption.label }}</span>
      </div>

      <!-- 搜索输入框 -->
      <input
        v-if="filterable && visible"
        ref="inputRef"
        v-model="searchQuery"
        :placeholder="searchPlaceholder"
        :disabled="disabled"
        class="flex-1 min-w-0 outline-none bg-transparent text-sm"
        @input="handleSearch"
        @keydown="handleKeydown"
        @focus="handleFocus"
        @blur="handleBlur"
        @mousedown.stop
      />

      <!-- 占位符 -->
      <span
        v-else-if="!hasSelected"
        class="flex-1 text-gray-400 text-sm truncate select-none"
        @click="handleTriggerClick"
      >
        {{ placeholder }}
      </span>

      <!-- 右侧图标区域 -->
      <div class="flex items-center ml-2 flex-shrink-0 space-x-1">
        <!-- 清除按钮 -->
        <button
          v-if="clearable && hasSelected && !disabled"
          @click.stop="handleClear"
          @mousedown.stop
          class="p-1 hover:text-gray-600 focus:outline-none transition-colors"
        >
          <GlobalIcon name="x-mark" size="xs" />
        </button>

        <!-- 加载图标 -->
        <div v-if="loading" class="animate-spin">
          <GlobalIcon name="arrow-path" size="sm" />
        </div>

        <!-- 箭头图标 -->
        <div
          v-else-if="showArrow"
          :class="[
            'transition-transform duration-200',
            { 'rotate-180': visible }
          ]"
        >
          <GlobalIcon name="chevron-down" size="sm" />
        </div>
      </div>
    </div>

    <!-- 下拉选项面板 -->
    <Teleport to="body">
      <div
        v-if="visible"
        ref="dropdownRef"
        :class="[
          'fixed bg-white border border-gray-200 rounded-lg shadow-lg py-1',
          'transform transition-all duration-200 ease-out'
        ]"
        :style="{
          ...dropdownStyle,
          width: dropdownStyle.width || initialDropdownWidth
        }"
      >
        <!-- 选项列表 -->
        <div
          :class="[
            'overflow-y-auto',
            { 'max-h-60': !maxHeight }
          ]"
          :style="{ maxHeight: maxHeight ? `${maxHeight}px` : '240px' }"
        >
          <div
            v-for="(option, index) in filteredOptions"
            :key="option.value"
            :class="[
              'flex items-center px-3 py-2 cursor-pointer text-sm transition-colors',
              {
                'bg-blue-50 text-blue-600': isSelected(option),
                'hover:bg-gray-50': !option.disabled && !isSelected(option),
                'text-gray-400 cursor-not-allowed': option.disabled,
                'bg-blue-100': index === highlightIndex
              }
            ]"
            @click="handleOptionClick(option)"
          >
            <!-- 多选复选框 -->
            <div v-if="multiple" class="mr-2">
              <div
                :class="[
                  'w-4 h-4 border-2 rounded transition-colors',
                  {
                    'bg-blue-500 border-blue-500': isSelected(option),
                    'border-gray-300': !isSelected(option)
                  }
                ]"
              >
                <GlobalIcon
                  v-if="isSelected(option)"
                  name="check"
                  size="xs"
                  color="text-white"
                />
              </div>
            </div>

            <!-- 选项图标 -->
            <GlobalIcon
              v-if="option.icon"
              :name="option.icon"
              size="sm"
              class="mr-2 flex-shrink-0"
            />

            <!-- 选项文本 -->
            <span class="flex-1 truncate">{{ option.label }}</span>

            <!-- 单选选中图标 -->
            <GlobalIcon
              v-if="!multiple && isSelected(option)"
              name="check"
              size="sm"
              color="text-blue-500"
              class="ml-2 flex-shrink-0"
            />
          </div>

          <!-- 无选项提示 -->
          <div
            v-if="filteredOptions.length === 0"
            class="px-3 py-8 text-center text-gray-400 text-sm"
          >
            <GlobalIcon name="magnifying-glass" size="lg" class="mx-auto mb-2 opacity-50" />
            <div>{{ searchQuery ? '无匹配选项' : '暂无选项' }}</div>
          </div>
        </div>
      </div>
    </Teleport>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, nextTick, onMounted, onUnmounted } from 'vue'
import type { SelectProps, SelectEmits, SelectOption } from './types'

defineOptions({
  name: 'GlobalSelect'
})

const props = withDefaults(defineProps<SelectProps>(), {
  options: () => [],
  placeholder: '请选择',
  size: 'medium',
  variant: 'default',
  clearable: false,
  multiple: false,
  filterable: false,
  maxTagCount: 3,
  showArrow: true,
  loading: false,
  maxHeight: 240,
  placement: 'auto',
  error: false,
  success: false
})

const emit = defineEmits<SelectEmits>()

// 模板引用
const selectRef = ref<HTMLDivElement>()
const triggerRef = ref<HTMLDivElement>()
const dropdownRef = ref<HTMLDivElement>()
const inputRef = ref<HTMLInputElement>()

// 状态管理
const visible = ref(false)
const searchQuery = ref('')
const highlightIndex = ref(-1)

// 初始化下拉样式
const dropdownStyle = ref({})

// 计算属性
const selectedOptions = computed(() => {
  if (!props.multiple) return []
  const values = Array.isArray(props.modelValue) ? props.modelValue as (string | number)[] : []
  return props.options.filter(option => values.includes(option.value))
})

const selectedOption = computed(() => {
  if (props.multiple) return null
  return props.options.find(option => option.value === props.modelValue) || null
})

const hasSelected = computed(() => {
  return props.multiple ? selectedOptions.value.length > 0 : !!selectedOption.value
})

const displayTags = computed(() => {
  if (!props.maxTagCount) return selectedOptions.value
  return selectedOptions.value.slice(0, props.maxTagCount)
})

const hasMoreTags = computed(() => {
  return props.maxTagCount && selectedOptions.value.length > props.maxTagCount
})

const searchPlaceholder = computed(() => {
  if (hasSelected.value) return ''
  return props.placeholder
})

const filteredOptions = computed(() => {
  if (!props.filterable || !searchQuery.value) {
    return props.options
  }

  const query = searchQuery.value.toLowerCase()
  return props.options.filter(option => {
    if (props.filterMethod) {
      return props.filterMethod(searchQuery.value, option)
    }
    return option.label.toLowerCase().includes(query)
  })
})

// 计算初始宽度，确保下拉面板从一开始就有正确的宽度
const initialDropdownWidth = computed(() => {
  if (!triggerRef.value) return '200px'
  const rect = triggerRef.value.getBoundingClientRect()
  return rect.width > 0 ? `${rect.width}px` : '200px'
})

// 样式计算
const sizeClasses = computed(() => {
  const sizes = {
    small: 'min-h-[32px] text-sm',
    medium: 'min-h-[40px] text-sm',
    large: 'min-h-[48px] text-base'
  }
  return sizes[props.size]
})

const variantClasses = computed(() => {
  const variants = {
    default: 'bg-white border-gray-300',
    bordered: 'bg-white border-2 border-gray-300',
    filled: 'bg-gray-50 border-gray-300'
  }
  return variants[props.variant]
})

const statusClasses = computed(() => {
  if (props.error) return 'border-red-500'
  if (props.success) return 'border-green-500'
  return ''
})

// 方法
const updateDropdownPosition = () => {
  if (!triggerRef.value) return

  const triggerRect = triggerRef.value.getBoundingClientRect()
  const viewportHeight = window.innerHeight
  const viewportWidth = window.innerWidth

  // 基础位置：触发器下方
  let top = triggerRect.bottom + window.scrollY + 4
  let left = triggerRect.left + window.scrollX

  // 检查垂直空间
  const spaceBelow = viewportHeight - triggerRect.bottom
  const spaceAbove = triggerRect.top
  const estimatedDropdownHeight = props.maxHeight || 240

  // 如果下方空间不够且上方空间更多，则显示在上方
  if (spaceBelow < estimatedDropdownHeight && spaceAbove > spaceBelow) {
    top = triggerRect.top + window.scrollY - estimatedDropdownHeight - 4
  }

  // 检查水平空间，防止超出右边界
  const dropdownWidth = triggerRect.width
  if (left + dropdownWidth > viewportWidth) {
    left = viewportWidth - dropdownWidth - 8 + window.scrollX
  }

  // 防止超出左边界
  if (left < 8) {
    left = 8 + window.scrollX
  }

  dropdownStyle.value = {
    position: 'absolute',
    top: `${top}px`,
    left: `${left}px`,
    width: `${dropdownWidth}px`,
    zIndex: '9999'
  }
}

// 监听滚动事件更新位置
const handleScroll = () => {
  if (visible.value) {
    updateDropdownPosition()
  }
}

// 监听窗口大小变化
const handleResize = () => {
  if (visible.value) {
    updateDropdownPosition()
  }
}

const isSelected = (option: SelectOption) => {
  if (props.multiple) {
    const values = Array.isArray(props.modelValue) ? props.modelValue as (string | number)[] : []
    return values.includes(option.value)
  }
  return props.modelValue === option.value
}

const handleTriggerClick = (event: MouseEvent) => {
  if (props.disabled) return

  // 检查是否点击的是清除按钮或其他控制按钮
  const target = event.target as HTMLElement
  if (target.closest('button')) {
    return
  }

  // 如果是可搜索模式且面板未显示，或者不是可搜索模式，则切换显示状态
  if (!props.filterable || !visible.value) {
    if (!visible.value) {
      // 在显示前先计算位置
      updateDropdownPosition()
    }

    visible.value = !visible.value
    emit('visible-change', visible.value)
  }

  if (visible.value && props.filterable) {
    nextTick(() => {
      inputRef.value?.focus()
    })
  }
}

const handleOptionClick = (option: SelectOption) => {
  if (option.disabled) return

  if (props.multiple) {
    const values = Array.isArray(props.modelValue) ? [...(props.modelValue as (string | number)[])] : []
    const index = values.findIndex(v => v === option.value)

    if (index > -1) {
      values.splice(index, 1)
    } else {
      values.push(option.value)
    }

    emit('update:modelValue', values)
    emit('change', values, selectedOptions.value)
  } else {
    emit('update:modelValue', option.value)
    emit('change', option.value, option)
    visible.value = false
    emit('visible-change', false)
  }
}

const removeTag = (option: SelectOption) => {
  if (props.disabled) return

  const values = Array.isArray(props.modelValue) ? [...(props.modelValue as (string | number)[])] : []
  const index = values.findIndex(v => v === option.value)

  if (index > -1) {
    values.splice(index, 1)
    emit('update:modelValue', values)
    emit('change', values, selectedOptions.value)
  }
}

const handleClear = () => {
  const newValue = props.multiple ? [] : ''
  emit('update:modelValue', newValue)
  emit('change', newValue, props.multiple ? [] : null)
  emit('clear')
  searchQuery.value = ''
}

const handleSearch = () => {
  highlightIndex.value = -1
  emit('search', searchQuery.value)
}

const handleKeydown = (event: KeyboardEvent) => {
  if (!visible.value) return

  switch (event.key) {
    case 'ArrowDown':
      event.preventDefault()
      highlightIndex.value = Math.min(highlightIndex.value + 1, filteredOptions.value.length - 1)
      break
    case 'ArrowUp':
      event.preventDefault()
      highlightIndex.value = Math.max(highlightIndex.value - 1, 0)
      break
    case 'Enter':
      event.preventDefault()
      if (highlightIndex.value >= 0 && filteredOptions.value[highlightIndex.value]) {
        handleOptionClick(filteredOptions.value[highlightIndex.value])
      }
      break
    case 'Escape':
      visible.value = false
      emit('visible-change', false)
      break
  }
}

const handleFocus = (event: FocusEvent) => {
  emit('focus', event)
}

const handleBlur = (event: FocusEvent) => {
  emit('blur', event)
}

const handleClickOutside = (event: Event) => {
  if (!selectRef.value || !dropdownRef.value) return

  const target = event.target as Node
  if (!selectRef.value.contains(target) && !dropdownRef.value.contains(target)) {
    visible.value = false
    emit('visible-change', false)
    searchQuery.value = ''
  }
}

// 暴露方法
const focus = () => {
  if (props.filterable && inputRef.value) {
    inputRef.value.focus()
  } else {
    triggerRef.value?.focus()
  }
}

const blur = () => {
  if (props.filterable && inputRef.value) {
    inputRef.value.blur()
  } else {
    triggerRef.value?.blur()
  }
}

const clear = () => {
  handleClear()
}

defineExpose({
  focus,
  blur,
  clear
})

// 生命周期
onMounted(() => {
  document.addEventListener('click', handleClickOutside)
  window.addEventListener('scroll', handleScroll, true)
  window.addEventListener('resize', handleResize)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
  window.removeEventListener('scroll', handleScroll, true)
  window.removeEventListener('resize', handleResize)
})

// 监听器
watch(visible, (newVisible) => {
  if (!newVisible) {
    searchQuery.value = ''
    highlightIndex.value = -1
  }
})
</script>

<style scoped>
/* 自定义滚动条 */
.overflow-y-auto::-webkit-scrollbar {
  width: 6px;
}

.overflow-y-auto::-webkit-scrollbar-track {
  background: #f1f5f9;
}

.overflow-y-auto::-webkit-scrollbar-thumb {
  background: #cbd5e1;
  border-radius: 3px;
}

.overflow-y-auto::-webkit-scrollbar-thumb:hover {
  background: #94a3b8;
}
</style>
