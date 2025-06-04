export interface SwitchProps {
  /**
   * 开关状态
   */
  modelValue?: boolean

  /**
   * 是否禁用
   */
  disabled?: boolean

  /**
   * 开关尺寸
   */
  size?: 'small' | 'medium' | 'large'

  /**
   * 开启时的颜色
   */
  color?: 'primary' | 'success' | 'warning' | 'danger'

  /**
   * 开启时的文本
   */
  activeText?: string

  /**
   * 关闭时的文本
   */
  inactiveText?: string

  /**
   * 开启时的值
   */
  activeValue?: boolean | string | number

  /**
   * 关闭时的值
   */
  inactiveValue?: boolean | string | number

  /**
   * 是否显示图标
   */
  showIcon?: boolean

  /**
   * 自定义开启图标
   */
  activeIcon?: string

  /**
   * 自定义关闭图标
   */
  inactiveIcon?: string

  /**
   * 是否加载中
   */
  loading?: boolean

  /**
   * 开关名称
   */
  name?: string

  /**
   * 是否必填
   */
  required?: boolean
}

export interface SwitchEmits {
  /**
   * 更新值事件
   */
  (e: 'update:modelValue', value: boolean | string | number): void

  /**
   * 状态改变事件
   */
  (e: 'change', value: boolean | string | number): void

  /**
   * 点击事件
   */
  (e: 'click', event: MouseEvent): void
}

export interface SwitchInstance {
  /**
   * 获取焦点
   */
  focus: () => void

  /**
   * 失去焦点
   */
  blur: () => void

  /**
   * 切换状态
   */
  toggle: () => void
}
