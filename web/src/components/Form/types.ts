export interface FormItem {
  // 基础属性
  name: string
  label: string
  type: 'input' | 'password' | 'select' | 'checkbox' | 'radio' | 'textarea'

  // 表单项配置
  placeholder?: string
  required?: boolean
  disabled?: boolean

  // 校验规则
  rules?: FormRule[]

  // 输入框特定属性
  inputType?: 'text' | 'email' | 'number' | 'tel' | 'url'
  maxLength?: number
  minLength?: number

  // 选择器特定属性
  options?: FormOption[]
  multiple?: boolean

  // 尺寸和样式
  size?: 'small' | 'medium' | 'large'

  // 其他属性
  defaultValue?: any
  showPassword?: boolean // 密码框是否显示切换按钮
}

export interface FormOption {
  label: string
  value: string | number | boolean
  disabled?: boolean
  icon?: string
}

export interface FormRule {
  // 校验类型
  type?: 'required' | 'email' | 'url' | 'number' | 'pattern' | 'custom'

  // 校验消息
  message?: string

  // 长度校验
  min?: number
  max?: number

  // 正则表达式
  pattern?: RegExp

  // 自定义校验函数
  validator?: (value: any, formData: Record<string, any>) => boolean | string

  // 触发时机
  trigger?: 'blur' | 'change' | 'submit'
}

export interface FormProps {
  // 表单数据
  modelValue?: Record<string, any>

  // 表单项配置
  items: FormItem[]

  // 布局配置
  layout?: 'vertical' | 'horizontal' | 'inline'
  labelWidth?: string | number

  // 尺寸
  size?: 'small' | 'medium' | 'large'

  // 行为控制
  showResetButton?: boolean
  showSubmitButton?: boolean
  submitText?: string
  resetText?: string
  submitLoading?: boolean

  // 校验配置
  validateOnChange?: boolean
  validateOnBlur?: boolean
  showErrorMessage?: boolean

  // 样式配置
  bordered?: boolean
  disabled?: boolean
}

export interface FormEmits {
  'update:modelValue': [value: Record<string, any>]
  'submit': [value: Record<string, any>]
  'reset': []
  'validate': [valid: boolean, errors: Record<string, string[]>]
  'field-change': [field: string, value: any]
  'field-blur': [field: string, value: any]
}

export interface FormInstance {
  validate: () => Promise<{ valid: boolean; errors: Record<string, string[]> }>
  validateField: (field: string) => Promise<{ valid: boolean; errors: string[] }>
  resetValidation: () => void
  reset: () => void
  getFieldValue: (field: string) => any
  setFieldValue: (field: string, value: any) => void
}

export interface FormContext {
  formData: Record<string, any>
  errors: Record<string, string[]>
  layout: string
  size: string
  disabled: boolean
  validateField: (field: string) => void
  updateField: (field: string, value: any) => void
}
