<template>
  <form
    :class="[
      'global-form',
      {
        'global-form--inline': inline,
        'global-form--label-left': labelPosition === 'left',
        'global-form--label-right': labelPosition === 'right',
        'global-form--label-top': labelPosition === 'top'
      }
    ]"
    @submit.prevent
  >
    <slot />
  </form>
</template>

<script setup lang="ts">
import { provide, reactive, ref, computed, nextTick } from 'vue'
import type { FormProps, FormEmits, FormInstance, FormItemInstance, FormContext, FormRules, FormItemRule } from './types'
import { validateRules } from './validator'

defineOptions({
  name: 'GlobalForm'
})

const props = withDefaults(defineProps<FormProps>(), {
  labelPosition: 'right',
  showMessage: true,
  size: 'default',
  validateOnRuleChange: true,
  disabled: false,
  hideRequiredAsterisk: false,
  inlineMessage: false,
  statusIcon: false
})

const emit = defineEmits<FormEmits>()

// 存储所有表单项实例
const fields = ref<FormItemInstance[]>([])

// 表单验证状态
const isValidating = ref(false)

// 计算属性
const formSize = computed(() => props.size)

// 添加表单项
const addField = (field: FormItemInstance) => {
  if (field) {
    fields.value.push(field)
  }
}

// 移除表单项
const removeField = (field: FormItemInstance) => {
  if (field) {
    const index = fields.value.indexOf(field)
    if (index > -1) {
      fields.value.splice(index, 1)
    }
  }
}

// 获取指定字段的表单项
const getField = (prop: string) => {
  return fields.value.find(field => field.prop === prop)
}

// 过滤需要验证的字段
const filterFields = (props?: string | string[]) => {
  if (!props) return fields.value
  const propsArray = Array.isArray(props) ? props : [props]
  return fields.value.filter(field => field.prop && propsArray.includes(field.prop))
}

// 验证整个表单
const validate = async (callback?: (isValid: boolean, invalidFields?: Record<string, FormItemRule[]>) => void): Promise<boolean> => {
  if (!props.model) {
    console.warn('[GlobalForm] model is required for validate.')
    return false
  }

  isValidating.value = true

  const invalidFields: Record<string, FormItemRule[]> = {}
  let isValid = true

  // 并发验证所有字段
  const validationPromises = fields.value.map(async (field) => {
    try {
      const valid = await field.validate('submit')
      if (!valid && field.prop) {
        isValid = false
        // 这里可以收集具体的错误信息
      }
      return valid
    } catch (error) {
      if (field.prop) {
        isValid = false
      }
      return false
    }
  })

  try {
    await Promise.all(validationPromises)
  } catch (error) {
    console.error('[GlobalForm] Validation error:', error)
    isValid = false
  }

  isValidating.value = false

  // 执行回调
  if (callback) {
    callback(isValid, invalidFields)
  }

  return isValid
}

// 验证指定字段
const validateField = async (props: string | string[], callback?: (errorMessage?: string) => void): Promise<boolean> => {
  const fieldsToValidate = filterFields(props)

  if (fieldsToValidate.length === 0) {
    console.warn('[GlobalForm] Please pass correct props!')
    return false
  }

  let isValid = true
  let firstError = ''

  for (const field of fieldsToValidate) {
    try {
      const fieldValid = await field.validate('submit')
      if (!fieldValid) {
        isValid = false
        if (!firstError && field.errorMessage) {
          firstError = field.errorMessage
        }
      }
    } catch (error) {
      isValid = false
      if (!firstError) {
        firstError = error instanceof Error ? error.message : 'Validation failed'
      }
    }
  }

  if (callback) {
    callback(isValid ? undefined : firstError)
  }

  return isValid
}

// 重置表单字段
const resetFields = () => {
  fields.value.forEach(field => {
    field.resetField()
  })
}

// 清除验证信息
const clearValidation = (props?: string | string[]) => {
  const fieldsToReset = filterFields(props)
  fieldsToReset.forEach(field => {
    field.clearValidation()
  })
}

// 滚动到指定字段
const scrollToField = (prop: string) => {
  const field = getField(prop)
  if (field && field.$el) {
    field.$el.scrollIntoView({
      behavior: 'smooth',
      block: 'center'
    })
  }
}

// 提供表单上下文
const formContext: FormContext = reactive({
  model: computed(() => props.model),
  rules: computed(() => props.rules),
  labelPosition: computed(() => props.labelPosition),
  labelWidth: computed(() => props.labelWidth),
  labelSuffix: computed(() => props.labelSuffix),
  inline: computed(() => props.inline),
  size: formSize,
  showMessage: computed(() => props.showMessage),
  inlineMessage: computed(() => props.inlineMessage),
  statusIcon: computed(() => props.statusIcon),
  hideRequiredAsterisk: computed(() => props.hideRequiredAsterisk),
  disabled: computed(() => props.disabled),
  validateOnRuleChange: computed(() => props.validateOnRuleChange),
  addField,
  removeField
})

provide('formContext', formContext)

// 暴露方法
defineExpose<FormInstance>({
  validate,
  validateField,
  resetFields,
  clearValidation,
  scrollToField
})
</script>

<style scoped>
.global-form {
  @apply w-full;
}

.global-form--inline {
  @apply flex flex-wrap items-start;
}

.global-form--inline :deep(.global-form-item) {
  @apply inline-block mr-4 mb-4;
}

.global-form--label-left :deep(.global-form-item__label) {
  @apply text-left;
}

.global-form--label-right :deep(.global-form-item__label) {
  @apply text-right;
}

.global-form--label-top :deep(.global-form-item__label) {
  @apply text-left mb-1;
}

.global-form--label-top :deep(.global-form-item__content) {
  @apply ml-0;
}
</style>
