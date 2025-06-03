import type { FormItemRule } from './types'

export interface ValidationResult {
  valid: boolean
  errors: string[]
}

// 内置验证规则
const builtInValidators = {
  required: (value: any): boolean => {
    if (Array.isArray(value)) {
      return value.length > 0
    }
    if (typeof value === 'string') {
      return value.trim().length > 0
    }
    return value !== null && value !== undefined && value !== ''
  },

  email: (value: string): boolean => {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
    return emailRegex.test(value)
  },

  url: (value: string): boolean => {
    try {
      new URL(value)
      return true
    } catch {
      return false
    }
  },

  number: (value: any): boolean => {
    if (value === '' || value === null || value === undefined) return false
    return !isNaN(Number(value)) && isFinite(Number(value))
  },

  integer: (value: any): boolean => {
    if (value === '' || value === null || value === undefined) return false
    return Number.isInteger(Number(value))
  },

  float: (value: any): boolean => {
    if (value === '' || value === null || value === undefined) return false
    const num = Number(value)
    return !isNaN(num) && isFinite(num) && !Number.isInteger(num)
  },

  array: (value: any): boolean => {
    return Array.isArray(value)
  },

  object: (value: any): boolean => {
    return typeof value === 'object' && value !== null && !Array.isArray(value)
  },

  boolean: (value: any): boolean => {
    return typeof value === 'boolean'
  },

  string: (value: any): boolean => {
    return typeof value === 'string'
  },

  date: (value: any): boolean => {
    return value instanceof Date || !isNaN(Date.parse(value))
  },

  hex: (value: string): boolean => {
    return /^#?([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$/.test(value)
  }
}

// 验证单个规则
export const validateRule = async (
  value: any,
  rule: FormItemRule,
  source: Record<string, any>
): Promise<ValidationResult> => {
  try {
    // required 验证
    if (rule.required && !builtInValidators.required(value)) {
      return {
        valid: false,
        errors: [rule.message || '此字段是必填项']
      }
    }

    // 如果值为空且不是必填项，则跳过其他验证
    if (!rule.required && (value === '' || value === null || value === undefined)) {
      return { valid: true, errors: [] }
    }

    // 类型验证
    if (rule.type && rule.type !== 'string') {
      const typeValidator = builtInValidators[rule.type as keyof typeof builtInValidators]
      if (typeValidator && !typeValidator(value)) {
        return {
          valid: false,
          errors: [rule.message || `请输入正确的${rule.type}格式`]
        }
      }
    }

    // 长度验证
    if (rule.len !== undefined) {
      const length = Array.isArray(value) ? value.length : String(value).length
      if (length !== rule.len) {
        return {
          valid: false,
          errors: [rule.message || `长度必须为 ${rule.len} 个字符`]
        }
      }
    }

    // 最小值/长度验证
    if (rule.min !== undefined) {
      if (rule.type === 'number' || rule.type === 'integer' || rule.type === 'float') {
        // 数字类型验证数值大小
        const numValue = Number(value)
        if (numValue < rule.min) {
          return {
            valid: false,
            errors: [rule.message || `数值不能小于 ${rule.min}`]
          }
        }
      } else {
        // 字符串和数组验证长度
        const length = Array.isArray(value) ? value.length : String(value).length
        if (length < rule.min) {
          return {
            valid: false,
            errors: [rule.message || `最少需要 ${rule.min} 个字符`]
          }
        }
      }
    }

    // 最大值/长度验证
    if (rule.max !== undefined) {
      if (rule.type === 'number' || rule.type === 'integer' || rule.type === 'float') {
        // 数字类型验证数值大小
        const numValue = Number(value)
        if (numValue > rule.max) {
          return {
            valid: false,
            errors: [rule.message || `数值不能大于 ${rule.max}`]
          }
        }
      } else {
        // 字符串和数组验证长度
        const length = Array.isArray(value) ? value.length : String(value).length
        if (length > rule.max) {
          return {
            valid: false,
            errors: [rule.message || `最多允许 ${rule.max} 个字符`]
          }
        }
      }
    }

    // 正则表达式验证
    if (rule.pattern && !rule.pattern.test(String(value))) {
      return {
        valid: false,
        errors: [rule.message || '格式不正确']
      }
    }

    // 枚举值验证
    if (rule.enum && !rule.enum.includes(value)) {
      return {
        valid: false,
        errors: [rule.message || '请选择有效的选项']
      }
    }

    // 自定义验证器
    if (rule.validator) {
      return new Promise((resolve) => {
        rule.validator!(
          rule,
          value,
          (error?: string | Error) => {
            if (error) {
              const message = error instanceof Error ? error.message : error
              resolve({
                valid: false,
                errors: [message]
              })
            } else {
              resolve({ valid: true, errors: [] })
            }
          },
          source
        )
      })
    }

    // 异步验证器
    if (rule.asyncValidator) {
      try {
        await rule.asyncValidator(rule, value, () => {}, source)
        return { valid: true, errors: [] }
      } catch (error) {
        const message = error instanceof Error ? error.message : 'Validation failed'
        return {
          valid: false,
          errors: [message]
        }
      }
    }

    return { valid: true, errors: [] }
  } catch (error) {
    console.error('Validation error:', error)
    return {
      valid: false,
      errors: [error instanceof Error ? error.message : 'Validation failed']
    }
  }
}

// 验证多个规则
export const validateRules = async (
  value: any,
  rules: FormItemRule[],
  source: Record<string, any>
): Promise<ValidationResult> => {
  const errors: string[] = []

  for (const rule of rules) {
    const result = await validateRule(value, rule, source)
    if (!result.valid) {
      errors.push(...result.errors)
      // 遇到第一个错误就停止验证（可配置）
      break
    }
  }

  return {
    valid: errors.length === 0,
    errors
  }
}

// 验证整个表单
export const validateForm = async (
  model: Record<string, any>,
  rules: Record<string, FormItemRule | FormItemRule[]>
): Promise<{ valid: boolean; errors: Record<string, string[]> }> => {
  const errors: Record<string, string[]> = {}
  let isValid = true

  const validationPromises = Object.keys(rules).map(async (field) => {
    const fieldRules = Array.isArray(rules[field]) ? rules[field] as FormItemRule[] : [rules[field] as FormItemRule]
    const result = await validateRules(model[field], fieldRules, model)

    if (!result.valid) {
      errors[field] = result.errors
      isValid = false
    }
  })

  await Promise.all(validationPromises)

  return { valid: isValid, errors }
}
