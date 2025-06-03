<template>
  <div
    ref="formItemRef"
    :class="[
      'global-form-item',
      {
        'global-form-item--feedback': validateState,
        'global-form-item--error': validateState === 'error',
        'global-form-item--validating': validateState === 'validating',
        'global-form-item--success': validateState === 'success',
        'global-form-item--required': isRequired,
        'global-form-item--no-asterisk': formContext?.hideRequiredAsterisk,
        [`global-form-item--${sizeClass}`]: sizeClass,
        'global-form-item--inline': formContext?.inline,
        'global-form-item--label-top': formContext?.labelPosition === 'top'
      }
    ]"
  >
    <!-- 水平布局（默认） -->
    <div v-if="formContext?.labelPosition !== 'top'" class="flex items-start">
      <!-- 标签区域 -->
      <div
        v-if="label || $slots.label"
        class="flex-shrink-0 flex items-center"
        :style="labelStyle"
      >
        <label
          :for="labelFor"
          :class="[
            'global-form-item__label',
            {
              'global-form-item__label--required': isRequired && !formContext?.hideRequiredAsterisk
            }
          ]"
        >
          <slot name="label">{{ label }}{{ formContext?.labelSuffix || '' }}</slot>
        </label>
      </div>

      <!-- 内容区域 -->
      <div class="flex-1 min-w-0">
        <slot />

        <!-- 错误消息 -->
        <div
          v-if="shouldShowError"
          :class="[
            'global-form-item__error',
            {
              'global-form-item__error--inline': formContext?.inlineMessage
            }
          ]"
        >
          {{ validateMessage }}
        </div>
      </div>
    </div>

    <!-- 顶部布局 -->
    <div v-else>
      <!-- 标签 -->
      <label
        v-if="label || $slots.label"
        :for="labelFor"
        :class="[
          'global-form-item__label global-form-item__label--top',
          {
            'global-form-item__label--required': isRequired && !formContext?.hideRequiredAsterisk
          }
        ]"
      >
        <slot name="label">{{ label }}{{ formContext?.labelSuffix || '' }}</slot>
      </label>

      <!-- 内容区域 -->
      <div class="global-form-item__content">
        <slot />

        <!-- 错误消息 -->
        <div
          v-if="shouldShowError"
          :class="[
            'global-form-item__error',
            {
              'global-form-item__error--inline': formContext?.inlineMessage
            }
          ]"
        >
          {{ validateMessage }}
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, inject, onMounted, onUnmounted, provide, nextTick, reactive } from 'vue'
import type { FormItemProps, FormItemEmits, FormItemInstance, FormContext, FormItemRule } from './types'
import { validateRules } from './validator'

defineOptions({
  name: 'GlobalFormItem'
})

const props = withDefaults(defineProps<FormItemProps>(), {
  showMessage: true,
  size: 'default'
})

const emit = defineEmits<FormItemEmits>()

// 获取表单上下文
const formContext = inject<FormContext>('formContext')

// 模板引用
const formItemRef = ref<HTMLDivElement>()

// 验证状态
const validateState = ref<'' | 'validating' | 'success' | 'error'>('')
const validateMessage = ref('')
const validateDisabled = ref(false)

// 计算属性
const labelFor = computed(() => props.for || props.prop)

const isRequired = computed(() => {
  let requiredByRule = false

  if (formContext?.rules && props.prop) {
    const rules = formContext.rules[props.prop]
    if (rules) {
      const rulesArray = Array.isArray(rules) ? rules : [rules]
      requiredByRule = rulesArray.some(rule => rule.required)
    }
  }

  if (props.rules) {
    const rulesArray = Array.isArray(props.rules) ? props.rules : [props.rules]
    requiredByRule = rulesArray.some(rule => rule.required)
  }

  return props.required || requiredByRule
})

const sizeClass = computed(() => {
  return props.size || formContext?.size || 'default'
})

const fieldValue = computed(() => {
  if (formContext?.model && props.prop) {
    return formContext.model[props.prop]
  }
  return undefined
})

const labelStyle = computed(() => {
  if (formContext?.labelPosition === 'top' || formContext?.inline) {
    return {}
  }

  const labelWidth = props.labelWidth || formContext?.labelWidth
  if (!labelWidth) return {}

  return {
    width: typeof labelWidth === 'number' ? `${labelWidth}px` : labelWidth,
    flexShrink: '0'
  }
})

const contentStyle = computed(() => {
  // flex布局下不需要margin-left，由flex-1自动填充
  return {}
})

const shouldShowError = computed(() => {
  return (
    validateState.value === 'error' &&
    (props.showMessage ?? formContext?.showMessage ?? true) &&
    validateMessage.value
  )
})

// 获取验证规则
const getRules = (): FormItemRule[] => {
  let rules: FormItemRule[] = []

  // 从 props.rules 获取
  if (props.rules) {
    rules = Array.isArray(props.rules) ? props.rules : [props.rules]
  }

  // 从表单的 rules 获取
  if (formContext?.rules && props.prop) {
    const formRules = formContext.rules[props.prop]
    if (formRules) {
      const formRulesArray = Array.isArray(formRules) ? formRules : [formRules]
      rules = rules.concat(formRulesArray)
    }
  }

  return rules
}

// 过滤特定触发时机的规则
const getFilteredRules = (trigger?: string): FormItemRule[] => {
  const rules = getRules()

  if (!trigger) return rules

  // 当trigger为submit时，应该验证所有规则
  if (trigger === 'submit') {
    return rules
  }

  return rules.filter(rule => {
    if (!rule.trigger) return true

    const triggers = Array.isArray(rule.trigger) ? rule.trigger : [rule.trigger]
    return triggers.includes(trigger as 'blur' | 'change' | 'submit')
  })
}

// 验证字段
const validate = async (trigger?: string, callback?: (errorMessage?: string) => void): Promise<boolean> => {
  if (validateDisabled.value) {
    callback?.()
    return true
  }

  const rules = getFilteredRules(trigger)

  if (rules.length === 0) {
    callback?.()
    return true
  }

  validateState.value = 'validating'

  try {
    const result = await validateRules(fieldValue.value, rules, formContext?.model || {})

    if (result.valid) {
      validateState.value = 'success'
      validateMessage.value = ''
      callback?.()
      return true
    } else {
      validateState.value = 'error'
      validateMessage.value = result.errors[0] || 'Validation failed'
      callback?.(validateMessage.value)
      return false
    }
  } catch (error) {
    validateState.value = 'error'
    const errorMsg = error instanceof Error ? error.message : 'Validation failed'
    validateMessage.value = errorMsg
    callback?.(errorMsg)
    return false
  }
}

// 重置字段
const resetField = () => {
  validateState.value = ''
  validateMessage.value = ''
  validateDisabled.value = false

  if (formContext?.model && props.prop) {
    // 重置为合适的默认值
    formContext.model[props.prop] = ''
  }
}

// 获取字段的默认值
const getDefaultValue = () => {
  // 根据字段类型返回合适的默认值
  return ''
}

// 清除验证
const clearValidation = () => {
  validateState.value = ''
  validateMessage.value = ''
}

// 创建实例对象
const formItemInstance: FormItemInstance = reactive({
  prop: props.prop,
  get errorMessage() { return validateMessage.value },
  get $el() { return formItemRef.value },
  validate,
  resetField,
  clearValidation
})

// 注册到表单
onMounted(() => {
  if (props.prop) {
    formContext?.addField(formItemInstance)
  }
})

onUnmounted(() => {
  formContext?.removeField(formItemInstance)
})

// 暴露实例方法
defineExpose(formItemInstance)
</script>

<style scoped>
.global-form-item {
  @apply mb-4;
}

.global-form-item__label {
  @apply text-sm font-medium text-gray-700;
  line-height: 1.5;
  padding: 0.5rem 0.75rem 0.5rem 0;
  word-break: break-all;
  position: relative;
}

/* 所有标签前都有*号占位 */
.global-form-item__label::before {
  content: '*';
  @apply mr-1;
  color: transparent; /* 默认透明 */
  width: 8px; /* 固定宽度确保占位 */
  display: inline-block;
}

/* 必填时显示红色*号 */
.global-form-item__label--required::before {
  @apply text-red-500;
}

.global-form-item__label--top {
  @apply block mb-2;
  padding: 0 0 0.5rem 0;
}

.global-form-item__content {
  position: relative;
  font-size: 14px;
}

.global-form-item__error {
  @apply text-red-500 text-xs mt-1 transition-all duration-300;
  line-height: 1;
}

.global-form-item__error--inline {
  @apply inline-block ml-2;
}

/* 内联布局 */
.global-form-item--inline {
  @apply inline-block mr-4 mb-2;
  vertical-align: top;
}

.global-form-item--inline .global-form-item__label {
  padding: 0.5rem 0.5rem 0.5rem 0;
}

/* 不同尺寸 */
.global-form-item--small .global-form-item__label {
  @apply text-xs;
  padding: 0.375rem 0.75rem 0.375rem 0;
}

.global-form-item--large .global-form-item__label {
  @apply text-base;
  padding: 0.75rem 0.75rem 0.75rem 0;
}

/* 确保输入框和标签高度匹配 */
.global-form-item :deep(.global-input),
.global-form-item :deep(.global-textarea),
.global-form-item :deep(.global-select) {
  min-height: 40px;
}

.global-form-item--small :deep(.global-input),
.global-form-item--small :deep(.global-textarea),
.global-form-item--small :deep(.global-select) {
  min-height: 32px;
}

.global-form-item--large :deep(.global-input),
.global-form-item--large :deep(.global-textarea),
.global-form-item--large :deep(.global-select) {
  min-height: 48px;
}

/* 验证状态样式 */
.global-form-item--error :deep(.global-input),
.global-form-item--error :deep(.global-textarea),
.global-form-item--error :deep(.global-select) {
  border-color: #f56565;
}

.global-form-item--success :deep(.global-input),
.global-form-item--success :deep(.global-textarea),
.global-form-item--success :deep(.global-select) {
  border-color: #48bb78;
}

/* 顶部标签布局 */
.global-form-item--label-top .global-form-item__content {
  margin-left: 0;
}
</style>
