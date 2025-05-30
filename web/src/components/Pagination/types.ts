export interface PaginationProps {
  /**
   * 当前页码
   */
  current?: number

  /**
   * 每页条数
   */
  pageSize?: number

  /**
   * 数据总数
   */
  total: number

  /**
   * 指定每页可以显示多少条
   */
  pageSizeOptions?: number[]

  /**
   * 是否可以改变 pageSize
   */
  showSizeChanger?: boolean

  /**
   * 是否显示快速跳转
   */
  showQuickJumper?: boolean

  /**
   * 是否显示总数
   */
  showTotal?: boolean

  /**
   * 总数显示模板函数
   */
  totalTemplate?: (total: number, range: [number, number]) => string

  /**
   * 当两侧页码较多时，只显示部分页码，设置显示的页码按钮数量
   */
  showLessItems?: boolean

  /**
   * 是否禁用
   */
  disabled?: boolean

  /**
   * 分页器大小
   */
  size?: 'small' | 'medium' | 'large'

  /**
   * 是否简洁模式
   */
  simple?: boolean

  /**
   * 分页器位置
   */
  position?: 'left' | 'center' | 'right'

  /**
   * 当页数少于等于此值时不显示分页器
   */
  hideOnSinglePage?: boolean
}

export interface PaginationEmits {
  /**
   * 页码改变的回调，参数是改变后的页码及每页条数
   */
  change: [page: number, pageSize: number]

  /**
   * pageSize 变化的回调
   */
  showSizeChange: [current: number, size: number]
}

export interface PaginationSlots {
  /**
   * 自定义总数显示
   */
  total?: {
    total: number
    range: [number, number]
  }
}
