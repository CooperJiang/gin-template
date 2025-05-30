export interface CheckboxProps {
  /**
   * 绑定值
   */
  modelValue?: boolean | string | number | any[]

  /**
   * 选中时的值
   */
  trueValue?: boolean | string | number

  /**
   * 未选中时的值
   */
  falseValue?: boolean | string | number

  /**
   * 复选框标签
   */
  label?: string

  /**
   * 是否禁用
   */
  disabled?: boolean

  /**
   * 复选框尺寸
   */
  size?: 'small' | 'medium' | 'large'

  /**
   * 复选框颜色
   */
  color?: 'primary' | 'success' | 'warning' | 'danger'

  /**
   * 是否为中间状态
   */
  indeterminate?: boolean

  /**
   * 复选框名称
   */
  name?: string

  /**
   * 复选框ID
   */
  id?: string

  /**
   * 是否必填
   */
  required?: boolean

  /**
   * 自定义选中图标
   */
  checkedIcon?: string

  /**
   * 自定义未选中图标
   */
  uncheckedIcon?: string

  /**
   * 自定义中间状态图标
   */
  indeterminateIcon?: string

  /**
   * 是否显示边框
   */
  bordered?: boolean

  /**
   * 校验状态
   */
  validateStatus?: 'success' | 'warning' | 'error' | ''
}

export interface CheckboxEmits {
  /**
   * 更新值事件
   */
  (e: 'update:modelValue', value: boolean | string | number | any[]): void

  /**
   * 状态改变事件
   */
  (e: 'change', value: boolean | string | number | any[], event: Event): void

  /**
   * 点击事件
   */
  (e: 'click', event: MouseEvent): void
}

export interface CheckboxInstance {
  /**
   * 获取焦点
   */
  focus: () => void

  /**
   * 失去焦点
   */
  blur: () => void

  /**
   * 切换选中状态
   */
  toggle: () => void
}
