export interface SelectOption {
  label: string
  value: string | number
  disabled?: boolean
  icon?: string
}

export interface SelectProps {
  // 基础属性
  modelValue?: string | number | (string | number)[]
  options?: SelectOption[]
  placeholder?: string
  disabled?: boolean
  clearable?: boolean
  multiple?: boolean

  // 尺寸和样式
  size?: 'small' | 'medium' | 'large'
  variant?: 'default' | 'bordered' | 'filled'

  // 搜索和过滤
  filterable?: boolean
  filterMethod?: (query: string, option: SelectOption) => boolean

  // 显示控制
  maxTagCount?: number
  showArrow?: boolean
  loading?: boolean

  // 下拉面板
  maxHeight?: number
  placement?: 'bottom' | 'top' | 'auto'

  // 验证状态
  error?: boolean
  success?: boolean
}

export interface SelectEmits {
  'update:modelValue': [value: string | number | (string | number)[]]
  'change': [value: string | number | (string | number)[], option: SelectOption | SelectOption[] | null]
  'clear': []
  'focus': [event: FocusEvent]
  'blur': [event: FocusEvent]
  'search': [query: string]
  'visible-change': [visible: boolean]
}

export interface SelectInstance {
  focus: () => void
  blur: () => void
  clear: () => void
}
