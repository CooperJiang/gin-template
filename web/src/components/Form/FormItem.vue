<template>
  <div
    :class="[
      'form-item mb-4',
      {
        'form-item--horizontal': layout === 'horizontal',
        'form-item--inline': layout === 'inline',
        'form-item--error': hasError,
        'form-item--required': item.required
      }
    ]"
  >
    <!-- 标签 -->
    <label
      v-if="item.label && layout !== 'inline'"
      :class="[
        'form-item__label block text-sm font-medium text-gray-700 mb-1',
        {
          'flex-shrink-0': layout === 'horizontal'
        }
      ]"
      :style="labelStyle"
    >
      {{ item.label }}
      <span v-if="item.required" class="text-red-500 ml-1">*</span>
    </label>

    <!-- 表单控件容器 -->
    <div
      :class="[
        'form-item__control',
        {
          'flex-1': layout === 'horizontal'
        }
      ]"
    >
      <!-- 输入框 -->
      <GlobalInput
        v-if="item.type === 'input'"
        :type="item.inputType || 'text'"
        :model-value="String(modelValue || '')"
        :placeholder="item.placeholder"
        :disabled="item.disabled || disabled"
        :maxlength="item.maxLength"
        :size="item.size || size"
        :error="hasError"
        @update:model-value="handleChange"
        @blur="handleBlur"
      />

      <!-- 密码框 -->
      <GlobalInput
        v-else-if="item.type === 'password'"
        type="password"
        :model-value="String(modelValue || '')"
        :placeholder="item.placeholder"
        :disabled="item.disabled || disabled"
        :maxlength="item.maxLength"
        :size="item.size || size"
        :error="hasError"
        :show-password="item.showPassword"
        @update:model-value="handleChange"
        @blur="handleBlur"
      />

      <!-- 文本域 -->
      <GlobalTextarea
        v-else-if="item.type === 'textarea'"
        :model-value="String(modelValue || '')"
        :placeholder="item.placeholder"
        :disabled="item.disabled || disabled"
        :maxlength="item.maxLength"
        :size="item.size || size"
        :error="hasError"
        rows="3"
        @update:model-value="handleChange"
        @blur="handleBlur"
      />

      <!-- 下拉选择 -->
      <GlobalSelect
        v-else-if="item.type === 'select'"
        :model-value="modelValue"
        :options="item.options || []"
        :placeholder="item.placeholder"
        :disabled="item.disabled || disabled"
        :multiple="item.multiple"
        :size="item.size || size"
        @update:model-value="handleChange"
        @blur="handleBlur"
      />

      <!-- 复选框 -->
      <div v-else-if="item.type === 'checkbox'" class="space-y-2">
        <label
          v-for="option in item.options"
          :key="option.value"
          class="flex items-center cursor-pointer hover:bg-gray-50 p-2 rounded"
        >
          <input
            type="checkbox"
            :value="option.value"
            :checked="Array.isArray(modelValue) ? modelValue.includes(option.value) : modelValue === option.value"
            :disabled="option.disabled || item.disabled || disabled"
            class="mr-2 h-4 w-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
            @change="handleCheckboxChange(option.value, $event)"
          />
          <GlobalIcon v-if="option.icon" :name="option.icon" size="sm" class="mr-2" />
          <span class="text-sm text-gray-700">{{ option.label }}</span>
        </label>
      </div>

      <!-- 单选框 -->
      <div v-else-if="item.type === 'radio'" class="space-y-2">
        <label
          v-for="option in item.options"
          :key="option.value"
          class="flex items-center cursor-pointer hover:bg-gray-50 p-2 rounded"
        >
          <input
            type="radio"
            :name="item.name"
            :value="option.value"
            :checked="modelValue === option.value"
            :disabled="option.disabled || item.disabled || disabled"
            class="mr-2 h-4 w-4 text-blue-600 border-gray-300 focus:ring-blue-500"
            @change="handleRadioChange(option.value)"
          />
          <GlobalIcon v-if="option.icon" :name="option.icon" size="sm" class="mr-2" />
          <span class="text-sm text-gray-700">{{ option.label }}</span>
        </label>
      </div>

      <!-- 错误信息 -->
      <div v-if="hasError && showErrorMessage" class="mt-1">
        <p
          v-for="error in errors"
          :key="error"
          class="text-sm text-red-600 flex items-center"
        >
          <GlobalIcon name="exclamation-circle" size="xs" class="mr-1" />
          {{ error }}
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, inject } from 'vue'
import type { FormItem, FormContext } from './types'

defineOptions({
  name: 'FormItem'
})

interface Props {
  item: FormItem
  modelValue?: any
  layout?: 'vertical' | 'horizontal' | 'inline'
  labelWidth?: string | number
  size?: 'small' | 'medium' | 'large'
  disabled?: boolean
  showErrorMessage?: boolean
  errors?: string[]
}

const props = withDefaults(defineProps<Props>(), {
  layout: 'vertical',
  size: 'medium',
  disabled: false,
  showErrorMessage: true,
  errors: () => []
})

const emit = defineEmits<{
  'update:model-value': [value: any]
  'field-change': [field: string, value: any]
  'field-blur': [field: string, value: any]
}>()

// 注入表单上下文
const formContext = inject<FormContext>('formContext')

// 密码显示状态
const showPasswordText = ref(false)

// 计算属性
const hasError = computed(() => props.errors.length > 0)

const labelStyle = computed(() => {
  if (props.layout === 'horizontal' && props.labelWidth) {
    const width = typeof props.labelWidth === 'number' ? `${props.labelWidth}px` : props.labelWidth
    return { width, minWidth: width }
  }
  return {}
})

// 事件处理
const handleInput = (event: Event) => {
  const target = event.target as HTMLInputElement | HTMLTextAreaElement
  handleChange(target.value)
}

const handleChange = (value: any) => {
  emit('update:model-value', value)
  emit('field-change', props.item.name, value)
  formContext?.updateField(props.item.name, value)
}

const handleBlur = () => {
  emit('field-blur', props.item.name, props.modelValue)
  formContext?.validateField(props.item.name)
}

const handleCheckboxChange = (value: any, event: Event) => {
  const target = event.target as HTMLInputElement
  let newValue: any

  if (props.item.multiple !== false) {
    // 多选模式
    const currentValue = Array.isArray(props.modelValue) ? [...props.modelValue] : []
    if (target.checked) {
      if (!currentValue.includes(value)) {
        currentValue.push(value)
      }
    } else {
      const index = currentValue.indexOf(value)
      if (index > -1) {
        currentValue.splice(index, 1)
      }
    }
    newValue = currentValue
  } else {
    // 单选模式
    newValue = target.checked ? value : null
  }

  handleChange(newValue)
}

const handleRadioChange = (value: any) => {
  handleChange(value)
}
</script>

<style scoped>
.form-item--horizontal {
  @apply flex items-start space-x-4;
}

.form-item--horizontal .form-item__label {
  @apply mb-0 pt-2;
}

.form-item--inline {
  @apply inline-block mr-4;
}

.form-item--error .form-item__label {
  @apply text-red-700;
}

.form-item--required .form-item__label {
  position: relative;
}
</style>
