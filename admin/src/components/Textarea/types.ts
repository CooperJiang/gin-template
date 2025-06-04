export interface TextareaProps {
  /**
   * 输入值
   */
  modelValue?: string

  /**
   * 占位符
   */
  placeholder?: string

  /**
   * 是否禁用
   */
  disabled?: boolean

  /**
   * 是否只读
   */
  readonly?: boolean

  /**
   * 输入框尺寸
   */
  size?: 'small' | 'medium' | 'large'

  /**
   * 行数
   */
  rows?: number

  /**
   * 最大长度
   */
  maxlength?: number

  /**
   * 最小长度
   */
  minlength?: number

  /**
   * 是否有错误状态
   */
  error?: boolean

  /**
   * 是否必填
   */
  required?: boolean

  /**
   * 自动完成
   */
  autocomplete?: string

  /**
   * 是否自动获取焦点
   */
  autofocus?: boolean

  /**
   * 输入框名称
   */
  name?: string

  /**
   * 是否可调整大小
   */
  resize?: 'none' | 'both' | 'horizontal' | 'vertical'

  /**
   * 是否显示字符计数
   */
  showCount?: boolean

  /**
   * 是否自动调整高度
   */
  autosize?: boolean | { minRows?: number; maxRows?: number }
}

export interface TextareaEmits {
  /**
   * 更新值事件
   */
  (e: 'update:modelValue', value: string): void

  /**
   * 输入事件
   */
  (e: 'input', value: string, event: Event): void

  /**
   * 失去焦点事件
   */
  (e: 'blur', event: FocusEvent): void

  /**
   * 获得焦点事件
   */
  (e: 'focus', event: FocusEvent): void

  /**
   * 按键事件
   */
  (e: 'keydown', event: KeyboardEvent): void

  /**
   * 清空事件
   */
  (e: 'clear'): void
}

export interface TextareaInstance {
  /**
   * 获取焦点
   */
  focus: () => void

  /**
   * 失去焦点
   */
  blur: () => void

  /**
   * 选中输入框内容
   */
  select: () => void
}
