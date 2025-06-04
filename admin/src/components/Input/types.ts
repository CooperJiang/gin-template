export interface InputProps {
  /**
   * 输入框类型
   */
  type?: 'text' | 'email' | 'password' | 'number' | 'tel' | 'url' | 'search'

  /**
   * 输入值
   */
  modelValue?: string | number

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
   * 最大长度
   */
  maxlength?: number

  /**
   * 最小长度
   */
  minlength?: number

  /**
   * 是否显示密码切换按钮（仅password类型有效）
   */
  showPassword?: boolean

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
   * 前缀图标
   */
  prefixIcon?: string

  /**
   * 后缀图标
   */
  suffixIcon?: string

  /**
   * 是否自动获取焦点
   */
  autofocus?: boolean

  /**
   * 输入框名称
   */
  name?: string
}

export interface InputEmits {
  /**
   * 更新值事件
   */
  (e: 'update:modelValue', value: string | number): void

  /**
   * 输入事件
   */
  (e: 'input', value: string | number, event: Event): void

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
   * 回车事件
   */
  (e: 'enter', event: KeyboardEvent): void

  /**
   * 清空事件
   */
  (e: 'clear'): void
}

export interface InputInstance {
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
