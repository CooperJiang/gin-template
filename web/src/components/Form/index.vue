<template>
  <form
    :class="[
      'global-form',
      {
        'form--horizontal': layout === 'horizontal',
        'form--inline': layout === 'inline',
        'form--bordered': bordered,
        'form--disabled': disabled
      }
    ]"
    @submit.prevent="handleSubmit"
  >
    <!-- 表单项 -->
    <FormItem
      v-for="item in items"
      :key="item.name"
      :item="item"
      :model-value="formData[item.name]"
      :layout="layout"
      :label-width="labelWidth"
      :size="item.size || size"
      :disabled="item.disabled || disabled"
      :show-error-message="showErrorMessage"
      :errors="errors[item.name] || []"
      @update:model-value="updateField(item.name, $event)"
      @field-change="handleFieldChange"
      @field-blur="handleFieldBlur"
    />

    <!-- 表单操作按钮 -->
    <div
      v-if="showSubmitButton || showResetButton"
      :class="[
        'form-actions mt-6',
        {
          'text-center': layout === 'vertical',
          'ml-auto': layout === 'horizontal' && labelWidth,
          'inline-block': layout === 'inline'
        }
      ]"
      :style="actionStyle"
    >
      <div class="space-x-3">
        <GlobalButton
          v-if="showResetButton"
          type="secondary"
          variant="outline"
          :size="size"
          :disabled="disabled"
          @click="handleReset"
        >
          {{ resetText }}
        </GlobalButton>

        <GlobalButton
          v-if="showSubmitButton"
          type="primary"
          htmlType="submit"
          :size="size"
          :loading="submitLoading"
          :disabled="disabled"
        >
          {{ submitText }}
        </GlobalButton>
      </div>
    </div>
  </form>
</template>

<script setup lang="ts">
import { ref, computed, provide, watch, nextTick } from 'vue'
import FormItem from './FormItem.vue'
import { validateForm, validateField, getRulesByTrigger } from './validator'
import type { FormProps, FormEmits, FormInstance, FormContext } from './types'

defineOptions({
  name: 'GlobalForm'
})

const props = withDefaults(defineProps<FormProps>(), {
  modelValue: () => ({}),
  layout: 'vertical',
  size: 'medium',
  showResetButton: true,
  showSubmitButton: true,
  submitText: '提交',
  resetText: '重置',
  submitLoading: false,
  validateOnChange: false,
  validateOnBlur: true,
  showErrorMessage: true,
  bordered: false,
  disabled: false
})

const emit = defineEmits<FormEmits>()

// 表单数据
const formData = ref<Record<string, any>>({ ...props.modelValue })
const errors = ref<Record<string, string[]>>({})

// 计算属性
const actionStyle = computed(() => {
  if (props.layout === 'horizontal' && props.labelWidth) {
    const width = typeof props.labelWidth === 'number' ? `${props.labelWidth}px` : props.labelWidth
    return { paddingLeft: width }
  }
  return {}
})

// 初始化表单数据
const initializeFormData = () => {
  const data: Record<string, any> = { ...props.modelValue }

  // 设置默认值
  props.items.forEach(item => {
    if (!(item.name in data) && item.defaultValue !== undefined) {
      data[item.name] = item.defaultValue
    }
  })

  formData.value = data
  emit('update:modelValue', data)
}

// 更新字段值
const updateField = (field: string, value: any) => {
  formData.value[field] = value
  emit('update:modelValue', { ...formData.value })

  if (props.validateOnChange) {
    nextTick(() => {
      validateFieldWithTrigger(field, 'change')
    })
  }
}

// 验证字段（根据触发时机）
const validateFieldWithTrigger = async (field: string, trigger: 'blur' | 'change' | 'submit') => {
  const item = props.items.find(item => item.name === field)
  if (!item?.rules) return

  const rules = getRulesByTrigger(item.rules, trigger)
  if (rules.length === 0) return

  const result = await validateField(formData.value[field], rules, formData.value)

  if (!result.valid) {
    errors.value[field] = result.errors
  } else {
    delete errors.value[field]
  }

  errors.value = { ...errors.value }
}

// 验证单个字段（公开方法）
const validateSingleField = async (field: string) => {
  const item = props.items.find(item => item.name === field)
  if (!item?.rules) {
    return { valid: true, errors: [] }
  }

  const result = await validateField(formData.value[field], item.rules, formData.value)

  if (!result.valid) {
    errors.value[field] = result.errors
  } else {
    delete errors.value[field]
  }

  errors.value = { ...errors.value }
  return result
}

// 验证整个表单
const validate = async () => {
  const result = await validateForm(formData.value, props.items)
  errors.value = result.errors
  emit('validate', result.valid, result.errors)
  return result
}

// 重置验证状态
const resetValidation = () => {
  errors.value = {}
}

// 重置表单
const reset = () => {
  // 重置为默认值
  const data: Record<string, any> = {}
  props.items.forEach(item => {
    if (item.defaultValue !== undefined) {
      data[item.name] = item.defaultValue
    } else {
      data[item.name] = item.type === 'checkbox' && item.multiple !== false ? [] : ''
    }
  })

  formData.value = data
  emit('update:modelValue', data)
  resetValidation()
  emit('reset')
}

// 获取字段值
const getFieldValue = (field: string) => {
  return formData.value[field]
}

// 设置字段值
const setFieldValue = (field: string, value: any) => {
  updateField(field, value)
}

// 事件处理
const handleSubmit = async () => {
  const result = await validate()
  if (result.valid) {
    emit('submit', { ...formData.value })
  }
}

const handleReset = () => {
  reset()
}

const handleFieldChange = (field: string, value: any) => {
  emit('field-change', field, value)
}

const handleFieldBlur = (field: string, value: any) => {
  if (props.validateOnBlur) {
    validateFieldWithTrigger(field, 'blur')
  }
  emit('field-blur', field, value)
}

// 提供表单上下文
const formContext: FormContext = {
  formData: formData.value,
  errors: errors.value,
  layout: props.layout,
  size: props.size,
  disabled: props.disabled,
  validateField: (field: string) => validateFieldWithTrigger(field, 'blur'),
  updateField
}

provide('formContext', formContext)

// 暴露方法
defineExpose<FormInstance>({
  validate,
  validateField: validateSingleField,
  resetValidation,
  reset,
  getFieldValue,
  setFieldValue
})

// 初始化
initializeFormData()

// 监听外部数据变化
watch(() => props.modelValue, (newValue) => {
  if (JSON.stringify(newValue) !== JSON.stringify(formData.value)) {
    formData.value = { ...newValue }
  }
}, { deep: true })
</script>

<style scoped>
.form--horizontal .form-actions {
  @apply flex justify-end;
}

.form--inline {
  @apply flex flex-wrap items-end;
}

.form--inline .form-actions {
  @apply ml-4;
}

.form--bordered {
  @apply border border-gray-200 rounded-lg p-6;
}

.form--disabled {
  @apply opacity-60 pointer-events-none;
}
</style>
