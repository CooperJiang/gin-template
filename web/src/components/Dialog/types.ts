export interface DialogProps {
  // 基础属性
  modelValue?: boolean
  title?: string
  width?: string | number
  height?: string | number

  // 位置和尺寸
  placement?: 'center' | 'top' | 'bottom'
  size?: 'small' | 'medium' | 'large' | 'full'

  // 行为控制
  closable?: boolean
  maskClosable?: boolean
  escClosable?: boolean
  closeOnRouteChange?: boolean

  // 样式控制
  showMask?: boolean
  maskClass?: string
  wrapClass?: string
  bodyClass?: string

  // 内容控制
  showHeader?: boolean
  showFooter?: boolean
  showClose?: boolean

  // 动画
  transition?: string

  // 层级
  zIndex?: number

  // 自定义
  customClass?: string

  // 加载状态
  loading?: boolean

  // 确认取消按钮
  confirmText?: string
  cancelText?: string
  confirmType?: 'primary' | 'success' | 'warning' | 'danger'
  showConfirm?: boolean
  showCancel?: boolean
  confirmLoading?: boolean

  // 状态
  persistent?: boolean // 是否持久化，不可关闭
}

export interface DialogEmits {
  'update:modelValue': [value: boolean]
  'open': []
  'opened': []
  'close': []
  'closed': []
  'confirm': []
  'cancel': []
  'mask-click': []
  'esc-press': []
}

export interface DialogInstance {
  open: () => void
  close: () => void
  toggle: () => void
}

export interface DialogSlots {
  default?: any
  header?: any
  title?: any
  footer?: any
  close?: any
}
