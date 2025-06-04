export interface CardProps {
  /**
   * 卡片标题
   */
  title?: string

  /**
   * 卡片副标题
   */
  subtitle?: string

  /**
   * 卡片大小
   */
  size?: 'small' | 'medium' | 'large'

  /**
   * 卡片阴影
   */
  shadow?: 'none' | 'small' | 'medium' | 'large' | 'hover'

  /**
   * 卡片边框
   */
  bordered?: boolean

  /**
   * 是否可悬停
   */
  hoverable?: boolean

  /**
   * 卡片背景色
   */
  background?: 'white' | 'gray' | 'blue' | 'green' | 'yellow' | 'red' | 'purple' | 'indigo'

  /**
   * 是否显示头部分割线
   */
  headerDivider?: boolean

  /**
   * 是否显示底部分割线
   */
  footerDivider?: boolean

  /**
   * 卡片圆角
   */
  rounded?: 'none' | 'small' | 'medium' | 'large' | 'full'

  /**
   * 头部内边距
   */
  headerPadding?: 'none' | 'small' | 'medium' | 'large'

  /**
   * 内容内边距
   */
  bodyPadding?: 'none' | 'small' | 'medium' | 'large'

  /**
   * 底部内边距
   */
  footerPadding?: 'none' | 'small' | 'medium' | 'large'

  /**
   * 是否加载中
   */
  loading?: boolean
}

export interface CardEmits {
  /**
   * 点击事件
   */
  (e: 'click', event: MouseEvent): void

  /**
   * 鼠标进入事件
   */
  (e: 'mouseenter', event: MouseEvent): void

  /**
   * 鼠标离开事件
   */
  (e: 'mouseleave', event: MouseEvent): void
}

export interface CardSlots {
  /**
   * 默认插槽 - 卡片内容
   */
  default?: () => any

  /**
   * 头部插槽
   */
  header?: () => any

  /**
   * 标题插槽
   */
  title?: () => any

  /**
   * 副标题插槽
   */
  subtitle?: () => any

  /**
   * 额外操作插槽
   */
  extra?: () => any

  /**
   * 底部插槽
   */
  footer?: () => any

  /**
   * 封面插槽
   */
  cover?: () => any

  /**
   * 加载状态插槽
   */
  loading?: () => any
}
