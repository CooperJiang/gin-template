import type { IconName } from '../Icon/types'

export interface ButtonProps {
  /**
   * 按钮类型
   */
  type?: 'primary' | 'secondary' | 'success' | 'warning' | 'danger' | 'info'

  /**
   * 按钮尺寸
   */
  size?: 'small' | 'medium' | 'large'

  /**
   * 是否为加载状态
   */
  loading?: boolean

  /**
   * 是否禁用
   */
  disabled?: boolean

  /**
   * 按钮形状
   */
  shape?: 'default' | 'round' | 'circle'

  /**
   * 按钮样式变体
   */
  variant?: 'solid' | 'outline' | 'ghost' | 'link'

  /**
   * 是否为块级元素
   */
  block?: boolean

  /**
   * 图标名称
   */
  icon?: IconName

  /**
   * 图标位置
   */
  iconPosition?: 'left' | 'right'

  /**
   * 图标变体
   */
  iconVariant?: 'outline' | 'solid'

  /**
   * HTML button type
   */
  htmlType?: 'button' | 'submit' | 'reset'
}

export interface ButtonEmits {
  /**
   * 点击事件
   */
  (e: 'click', event: MouseEvent): void
}
