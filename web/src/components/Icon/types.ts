export interface IconProps {
  /**
   * 图标名称
   */
  name: string

  /**
   * 图标大小
   */
  size?: 'xs' | 'sm' | 'md' | 'lg' | 'xl' | '2xl'

  /**
   * 图标风格
   */
  variant?: 'outline' | 'solid'

  /**
   * 自定义类名
   */
  class?: string

  /**
   * 图标颜色 (Tailwind CSS 颜色类)
   */
  color?: string
}

// 可用图标映射
export type IconName =
  // 基础操作
  | 'plus' | 'minus' | 'x-mark' | 'check'
  | 'pencil' | 'trash' | 'eye' | 'eye-slash'
  | 'heart' | 'star' | 'share' | 'download'
  | 'upload' | 'refresh' | 'settings' | 'search'

  // 箭头和导航
  | 'arrow-up' | 'arrow-down' | 'arrow-left' | 'arrow-right'
  | 'chevron-up' | 'chevron-down' | 'chevron-left' | 'chevron-right'
  | 'arrow-up-right' | 'arrow-down-right' | 'arrow-path'

  // 用户和社交
  | 'user' | 'user-group' | 'user-plus' | 'user-minus'
  | 'chat-bubble-left' | 'chat-bubble-left-right' | 'phone'
  | 'envelope' | 'bell' | 'bell-slash'

  // 文件和文档
  | 'document' | 'document-text' | 'folder' | 'folder-open'
  | 'archive-box' | 'clipboard' | 'clipboard-document'
  | 'photo' | 'video-camera' | 'musical-note'

  // 界面和布局
  | 'home' | 'building-office' | 'map-pin' | 'globe-alt'
  | 'calendar' | 'clock' | 'sun' | 'moon'
  | 'list-bullet' | 'squares-2x2' | 'table-cells'
  | 'bars-3' | 'bars-4' | 'ellipsis-horizontal'

  // 状态和反馈
  | 'check-circle' | 'x-circle' | 'exclamation-triangle'
  | 'information-circle' | 'question-mark-circle'
  | 'shield-check' | 'shield-exclamation' | 'lock-closed'
  | 'lock-open' | 'key' | 'finger-print'

  // 电商和商业
  | 'shopping-cart' | 'shopping-bag' | 'credit-card'
  | 'banknotes' | 'gift' | 'tag' | 'ticket'
  | 'chart-bar' | 'chart-pie' | 'presentation-chart-line'

  // 技术和工具
  | 'cog-6-tooth' | 'wrench-screwdriver' | 'cpu-chip'
  | 'server' | 'cloud' | 'wifi' | 'signal'
  | 'device-phone-mobile' | 'computer-desktop' | 'tv'
