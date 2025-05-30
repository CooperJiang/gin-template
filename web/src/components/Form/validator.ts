import type { FormRule, FormItem } from './types'

// 内置验证规则
export const builtInValidators = {
  required: (value: any): boolean => {
    if (Array.isArray(value)) {
      return value.length > 0
    }
    return value !== null && value !== undefined && String(value).trim() !== ''
  },

  email: (value: string): boolean => {
    if (!value) return true // 空值由 required 规则处理
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
    return emailRegex.test(value)
  },

  url: (value: string): boolean => {
    if (!value) return true
    try {
      new URL(value)
      return true
    } catch {
      return false
    }
  },

  number: (value: any): boolean => {
    if (!value) return true
    return !isNaN(Number(value))
  },

  pattern: (value: string, pattern: RegExp): boolean => {
    if (!value) return true
    return pattern.test(value)
  }
}

// 默认错误消息
export const defaultMessages = {
  required: '此字段为必填项',
  email: '请输入有效的邮箱地址',
  url: '请输入有效的URL地址',
  number: '请输入有效的数字',
  pattern: '输入格式不正确',
  min: '长度不能少于 {min} 个字符',
  max: '长度不能超过 {max} 个字符'
}

// 验证单个字段
export async function validateField(
  value: any,
  rules: FormRule[],
  formData: Record<string, any> = {}
): Promise<{ valid: boolean; errors: string[] }> {
  const errors: string[] = []

  for (const rule of rules) {
    let isValid = true
    let errorMessage = ''

    switch (rule.type) {
      case 'required':
        isValid = builtInValidators.required(value)
        errorMessage = rule.message || defaultMessages.required
        break

      case 'email':
        isValid = builtInValidators.email(value)
        errorMessage = rule.message || defaultMessages.email
        break

      case 'url':
        isValid = builtInValidators.url(value)
        errorMessage = rule.message || defaultMessages.url
        break

      case 'number':
        isValid = builtInValidators.number(value)
        errorMessage = rule.message || defaultMessages.number
        break

      case 'pattern':
        if (rule.pattern) {
          isValid = builtInValidators.pattern(value, rule.pattern)
          errorMessage = rule.message || defaultMessages.pattern
        }
        break

      case 'custom':
        if (rule.validator) {
          const result = rule.validator(value, formData)
          if (typeof result === 'boolean') {
            isValid = result
            errorMessage = rule.message || '验证失败'
          } else {
            isValid = false
            errorMessage = result
          }
        }
        break
    }

    // 长度验证
    if (isValid && (rule.min !== undefined || rule.max !== undefined)) {
      const length = Array.isArray(value) ? value.length : String(value || '').length

      if (rule.min !== undefined && length < rule.min) {
        isValid = false
        errorMessage = rule.message || defaultMessages.min.replace('{min}', String(rule.min))
      }

      if (rule.max !== undefined && length > rule.max) {
        isValid = false
        errorMessage = rule.message || defaultMessages.max.replace('{max}', String(rule.max))
      }
    }

    if (!isValid) {
      errors.push(errorMessage)
    }
  }

  return {
    valid: errors.length === 0,
    errors
  }
}

// 验证整个表单
export async function validateForm(
  formData: Record<string, any>,
  items: FormItem[]
): Promise<{ valid: boolean; errors: Record<string, string[]> }> {
  const errors: Record<string, string[]> = {}
  let isValid = true

  for (const item of items) {
    if (item.rules && item.rules.length > 0) {
      const fieldResult = await validateField(formData[item.name], item.rules, formData)
      if (!fieldResult.valid) {
        errors[item.name] = fieldResult.errors
        isValid = false
      }
    }
  }

  return {
    valid: isValid,
    errors
  }
}

// 根据触发时机过滤规则
export function getRulesByTrigger(rules: FormRule[], trigger: 'blur' | 'change' | 'submit'): FormRule[] {
  return rules.filter(rule => {
    if (!rule.trigger) return trigger === 'submit' // 默认在提交时验证
    return rule.trigger === trigger || rule.trigger === 'submit'
  })
}
