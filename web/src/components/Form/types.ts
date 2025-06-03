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
  required?: boolean
  message?: string
  trigger?: 'blur' | 'change' | 'submit'
  type?: 'string' | 'number' | 'boolean' | 'method' | 'regexp' | 'integer' | 'float' | 'array' | 'object' | 'enum' | 'date' | 'url' | 'hex' | 'email'
  pattern?: RegExp
  min?: number
  max?: number
  len?: number
  enum?: (string | number | boolean | null | undefined)[]
  validator?: (rule: FormRule, value: any, callback: (error?: string | Error) => void, source?: Record<string, any>, options?: any) => void
  asyncValidator?: (rule: FormRule, value: any, callback: (error?: string | Error) => void, source?: Record<string, any>, options?: any) => Promise<void>
}

export interface FormItemRule extends FormRule {
  trigger?: 'blur' | 'change' | 'submit' | ('blur' | 'change' | 'submit')[]
}

export type FormRules = Record<string, FormItemRule | FormItemRule[]>

export interface FormProps {
  model: Record<string, any>
  rules?: FormRules
  inline?: boolean
  labelPosition?: 'left' | 'right' | 'top'
  labelWidth?: string | number
  labelSuffix?: string
  hideRequiredAsterisk?: boolean
  showMessage?: boolean
  inlineMessage?: boolean
  statusIcon?: boolean
  validateOnRuleChange?: boolean
  size?: 'large' | 'default' | 'small'
  disabled?: boolean
}

export interface FormEmits {
  (e: 'validate', prop: string, isValid: boolean, message: string): void
}

export interface FormItemProps {
  label?: string
  prop?: string
  labelWidth?: string | number
  required?: boolean
  rules?: FormItemRule | FormItemRule[]
  error?: string
  showMessage?: boolean
  inlineMessage?: boolean
  size?: 'large' | 'default' | 'small'
  for?: string
}

export interface FormItemEmits {
  // FormItem 通常不需要额外的 emit
}

export interface FormInstance {
  validate: (callback?: (isValid: boolean, invalidFields?: Record<string, FormItemRule[]>) => void) => Promise<boolean>
  validateField: (props: string | string[], callback?: (errorMessage?: string) => void) => Promise<boolean>
  resetFields: () => void
  clearValidation: (props?: string | string[]) => void
  scrollToField: (prop: string) => void
}

export interface FormItemInstance {
  prop?: string
  errorMessage?: string
  $el?: HTMLElement
  validate: (trigger?: string, callback?: (errorMessage?: string) => void) => Promise<boolean>
  resetField: () => void
  clearValidation: () => void
}

export interface FormContext {
  model: Record<string, any>
  rules?: FormRules
  labelPosition?: 'left' | 'right' | 'top'
  labelWidth?: string | number
  labelSuffix?: string
  inline?: boolean
  size?: 'large' | 'default' | 'small'
  showMessage?: boolean
  inlineMessage?: boolean
  statusIcon?: boolean
  hideRequiredAsterisk?: boolean
  disabled?: boolean
  validateOnRuleChange?: boolean
  addField: (field: FormItemInstance) => void
  removeField: (field: FormItemInstance) => void
}
