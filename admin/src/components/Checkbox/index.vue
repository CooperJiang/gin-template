<template>
  <label
    :class="wrapperClass"
    @click="handleLabelClick"
  >
    <input
      ref="checkboxRef"
      type="checkbox"
      :checked="isChecked"
      :disabled="disabled"
      :name="name"
      :id="id"
      :required="required"
      :class="inputClass"
      @change="handleChange"
      @click.stop="handleInputClick"
    />

    <!-- 自定义复选框外观 -->
    <span :class="checkboxClass">
      <GlobalIcon
        v-if="currentIcon"
        :name="currentIcon"
        :size="iconSize"
        :class="iconClass"
      />
    </span>

    <!-- 标签文本 -->
    <span v-if="label || $slots.default" :class="labelClass">
      <slot>{{ label }}</slot>
    </span>
  </label>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import type { CheckboxProps, CheckboxEmits, CheckboxInstance } from './types'

defineOptions({
  name: 'GlobalCheckbox'
})

const props = withDefaults(defineProps<CheckboxProps>(), {
  trueValue: true,
  falseValue: false,
  disabled: false,
  size: 'medium',
  color: 'primary',
  indeterminate: false,
  required: false,
  bordered: false,
  validateStatus: ''
})

const emit = defineEmits<CheckboxEmits>()

// 引用
const checkboxRef = ref<HTMLInputElement>()

// 计算属性
const isChecked = computed(() => {
  if (Array.isArray(props.modelValue)) {
    return props.modelValue.includes(props.trueValue)
  }
  return props.modelValue === props.trueValue
})

const wrapperClass = computed(() => {
  const classes = [
    'checkbox-wrapper',
    'inline-flex items-center gap-2 cursor-pointer transition-all duration-200'
  ]

  if (props.disabled) {
    classes.push('cursor-not-allowed opacity-60')
  }

  if (props.bordered) {
    classes.push('border border-gray-200 rounded-md p-2 hover:border-gray-300')
  }

  return classes.join(' ')
})

const inputClass = computed(() => {
  return 'sr-only' // 隐藏原生input，使用自定义外观
})

const checkboxClass = computed(() => {
  const classes = [
    'checkbox-custom',
    'inline-flex items-center justify-center rounded border-2 transition-all duration-200'
  ]

  // 尺寸样式
  const sizeClasses = {
    small: 'h-4 w-4',
    medium: 'h-5 w-5',
    large: 'h-6 w-6'
  }
  classes.push(sizeClasses[props.size])

  // 状态和颜色样式
  if (isChecked.value || props.indeterminate) {
    const colorClasses = {
      primary: 'bg-blue-600 border-blue-600',
      success: 'bg-green-600 border-green-600',
      warning: 'bg-yellow-600 border-yellow-600',
      danger: 'bg-red-600 border-red-600'
    }
    classes.push(colorClasses[props.color])
  } else {
    classes.push('bg-white border-gray-300')
  }

  // 验证状态
  if (props.validateStatus) {
    const validateClasses = {
      success: 'border-green-500',
      warning: 'border-yellow-500',
      error: 'border-red-500'
    }
    if (props.validateStatus in validateClasses) {
      classes.push(validateClasses[props.validateStatus as keyof typeof validateClasses])
    }
  }

  // 悬停效果
  if (!props.disabled) {
    classes.push('hover:border-gray-400')
  }

  return classes.join(' ')
})

const labelClass = computed(() => {
  const classes = ['checkbox-label', 'select-none']

  // 尺寸样式
  const sizeClasses = {
    small: 'text-sm',
    medium: 'text-sm',
    large: 'text-base'
  }
  classes.push(sizeClasses[props.size])

  if (props.disabled) {
    classes.push('text-gray-400')
  } else {
    classes.push('text-gray-700')
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

const iconClass = computed(() => {
  return 'text-white'
})

const currentIcon = computed(() => {
  if (props.indeterminate) {
    return props.indeterminateIcon || 'minus'
  } else if (isChecked.value) {
    return props.checkedIcon || 'check'
  } else if (props.uncheckedIcon) {
    return props.uncheckedIcon
  }
  return null
})

// 事件处理
const handleChange = (event: Event) => {
  if (props.disabled) return

  const target = event.target as HTMLInputElement
  let newValue: boolean | string | number | any[]

  if (Array.isArray(props.modelValue)) {
    const newArray = [...props.modelValue]
    if (target.checked) {
      if (!newArray.includes(props.trueValue)) {
        newArray.push(props.trueValue)
      }
    } else {
      const index = newArray.indexOf(props.trueValue)
      if (index > -1) {
        newArray.splice(index, 1)
      }
    }
    newValue = newArray
  } else {
    newValue = target.checked ? props.trueValue : props.falseValue
  }

  emit('update:modelValue', newValue)
  emit('change', newValue, event)
}

const handleLabelClick = (event: MouseEvent) => {
  if (props.disabled) {
    event.preventDefault()
    return
  }
  emit('click', event)
}

const handleInputClick = (event: MouseEvent) => {
  if (props.disabled) {
    event.preventDefault()
    return
  }
}

// 实例方法
const focus = () => {
  checkboxRef.value?.focus()
}

const blur = () => {
  checkboxRef.value?.blur()
}

const toggle = () => {
  if (props.disabled) return

  let newValue: boolean | string | number | any[]

  if (Array.isArray(props.modelValue)) {
    const newArray = [...props.modelValue]
    if (isChecked.value) {
      const index = newArray.indexOf(props.trueValue)
      if (index > -1) {
        newArray.splice(index, 1)
      }
    } else {
      if (!newArray.includes(props.trueValue)) {
        newArray.push(props.trueValue)
      }
    }
    newValue = newArray
  } else {
    newValue = isChecked.value ? props.falseValue : props.trueValue
  }

  emit('update:modelValue', newValue)
  emit('change', newValue, new Event('change'))
}

// 暴露实例方法
defineExpose<CheckboxInstance>({
  focus,
  blur,
  toggle
})
</script>

<style scoped>
.checkbox-wrapper:hover .checkbox-custom {
  @apply border-gray-400;
}

.checkbox-wrapper--disabled:hover .checkbox-custom {
  @apply border-gray-300;
}
</style>
