export interface TagProps {
  /**
   * 标签类型
   */
  type?: 'default' | 'primary' | 'success' | 'warning' | 'danger' | 'info'

  /**
   * 标签尺寸
   */
  size?: 'small' | 'medium' | 'large'

  /**
   * 标签变体
   */
  variant?: 'filled' | 'outlined' | 'ghost'

  /**
   * 是否可关闭
   */
  closable?: boolean

  /**
   * 是否禁用
   */
  disabled?: boolean

  /**
   * 标签文本
   */
  text?: string

  /**
   * 自定义颜色
   */
  color?: string

  /**
   * 前缀图标
   */
  icon?: string

  /**
   * 关闭图标
   */
  closeIcon?: string

  /**
   * 是否圆角
   */
  round?: boolean

  /**
   * 是否点击状态
   */
  checkable?: boolean

  /**
   * 是否选中（当checkable为true时有效）
   */
  checked?: boolean

  /**
   * 最大宽度
   */
  maxWidth?: string | number
}

export interface TagEmits {
  /**
   * 关闭事件
   */
  (e: 'close', event: MouseEvent): void

  /**
   * 点击事件
   */
  (e: 'click', event: MouseEvent): void

  /**
   * 选中状态改变事件（当checkable为true时）
   */
  (e: 'update:checked', checked: boolean): void

  /**
   * 选中状态改变事件（当checkable为true时）
   */
  (e: 'change', checked: boolean, event: MouseEvent): void
}

export interface TagInstance {
  /**
   * 获取焦点
   */
  focus: () => void

  /**
   * 失去焦点
   */
  blur: () => void
}
