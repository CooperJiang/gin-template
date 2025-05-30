export interface TabItem {
  /**
   * 标签页唯一标识
   */
  key: string

  /**
   * 标签页标题
   */
  title: string

  /**
   * 是否禁用
   */
  disabled?: boolean

  /**
   * 图标名称
   */
  icon?: string

  /**
   * 是否可关闭
   */
  closable?: boolean

  /**
   * 自定义内容
   */
  content?: string
}

export interface TabsProps {
  /**
   * 当前激活的标签页key
   */
  modelValue?: string

  /**
   * 标签页数据
   */
  items?: TabItem[]

  /**
   * 标签页类型
   */
  type?: 'line' | 'card' | 'button'

  /**
   * 标签页位置
   */
  position?: 'top' | 'bottom' | 'left' | 'right'

  /**
   * 是否显示添加按钮
   */
  addable?: boolean

  /**
   * 标签页大小
   */
  size?: 'small' | 'medium' | 'large'

  /**
   * 是否可以拖拽排序
   */
  sortable?: boolean

  /**
   * 是否懒加载标签页内容
   */
  lazy?: boolean

  /**
   * 是否动画
   */
  animated?: boolean
}

export interface TabsEmits {
  /**
   * 激活标签页改变事件
   */
  (e: 'update:modelValue', key: string): void

  /**
   * 标签页切换事件
   */
  (e: 'change', key: string, item: TabItem): void

  /**
   * 标签页关闭事件
   */
  (e: 'close', key: string, item: TabItem): void

  /**
   * 添加标签页事件
   */
  (e: 'add'): void

  /**
   * 标签页排序事件
   */
  (e: 'sort', items: TabItem[]): void
}

export interface TabsSlots {
  /**
   * 默认插槽 - 标签页内容
   */
  default?: () => any

  /**
   * 标签页标题插槽
   */
  title?: (props: { item: TabItem; active: boolean }) => any

  /**
   * 添加按钮插槽
   */
  addButton?: () => any

  /**
   * 额外操作插槽
   */
  extra?: () => any
}
