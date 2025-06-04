export interface ProgressProps {
  /**
   * 当前进度值 (0-100)
   */
  value?: number

  /**
   * 进度条类型
   */
  type?: 'line' | 'circle' | 'dashboard'

  /**
   * 进度条状态
   */
  status?: 'normal' | 'success' | 'warning' | 'error' | 'info'

  /**
   * 进度条大小
   */
  size?: 'small' | 'medium' | 'large'

  /**
   * 是否显示进度文字
   */
  showText?: boolean

  /**
   * 进度文字位置 (仅line类型有效)
   */
  textPosition?: 'right' | 'inside' | 'top' | 'bottom'

  /**
   * 进度条粗细 (线性进度条的高度或圆形进度条的线条宽度)
   */
  strokeWidth?: number

  /**
   * 圆形进度条的直径
   */
  width?: number

  /**
   * 进度条颜色
   */
  color?: string | string[]

  /**
   * 背景颜色
   */
  backgroundColor?: string

  /**
   * 是否显示动画
   */
  animated?: boolean

  /**
   * 是否条纹效果
   */
  striped?: boolean

  /**
   * 自定义格式化函数
   */
  format?: (value: number) => string

  /**
   * 是否不确定进度 (显示滚动动画)
   */
  indeterminate?: boolean

  /**
   * 步骤数 (分段进度条)
   */
  steps?: number

  /**
   * 圆形进度条的缺口角度 (仅dashboard类型)
   */
  gapDegree?: number

  /**
   * 圆形进度条的缺口位置
   */
  gapPosition?: 'top' | 'bottom' | 'left' | 'right'
}

export interface ProgressEmits {
  /**
   * 进度值改变事件
   */
  (e: 'change', value: number): void

  /**
   * 进度完成事件
   */
  (e: 'complete'): void
}

export interface ProgressSlots {
  /**
   * 自定义进度文字
   */
  default?: (props: { value: number; status: string }) => any

  /**
   * 进度条内部内容 (仅line类型内部文字位置有效)
   */
  inner?: (props: { value: number; status: string }) => any
}
