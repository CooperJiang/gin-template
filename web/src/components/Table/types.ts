export interface TableColumn {
  key: string
  title: string
  dataIndex?: string
  width?: number | string
  minWidth?: number
  maxWidth?: number
  align?: 'left' | 'center' | 'right'
  sortable?: boolean
  filterable?: boolean
  fixed?: 'left' | 'right'
  render?: (value: any, record: any, index: number) => any
  ellipsis?: boolean
  resizable?: boolean
}

export interface TableData {
  [key: string]: any
  _id?: string | number
  _disabled?: boolean
  _selected?: boolean
  _expanded?: boolean
}

export interface SortInfo {
  column: string
  order: 'asc' | 'desc' | null
}

export interface FilterInfo {
  [key: string]: any
}

export interface TableProps {
  columns: TableColumn[]
  data: TableData[]
  loading?: boolean
  size?: 'small' | 'medium' | 'large'
  bordered?: boolean
  striped?: boolean
  hoverable?: boolean
  showHeader?: boolean
  height?: number | string
  maxHeight?: number | string
  rowKey?: string | ((record: TableData) => string | number)
  pagination?: boolean | object
  // 选择相关
  rowSelection?: {
    type?: 'checkbox' | 'radio'
    selectedRowKeys?: (string | number)[]
    onChange?: (selectedRowKeys: (string | number)[], selectedRows: TableData[]) => void
    onSelect?: (record: TableData, selected: boolean, selectedRows: TableData[]) => void
    onSelectAll?: (selected: boolean, selectedRows: TableData[], changeRows: TableData[]) => void
    getCheckboxProps?: (record: TableData) => { disabled?: boolean }
  }
  // 展开相关
  expandable?: {
    expandedRowKeys?: (string | number)[]
    expandRowByClick?: boolean
    onExpand?: (expanded: boolean, record: TableData) => void
    onExpandedRowsChange?: (expandedRows: (string | number)[]) => void
    expandedRowRender?: (record: TableData, index: number) => any
  }
  // 排序筛选
  sortInfo?: SortInfo
  filterInfo?: FilterInfo
  // 样式
  tableLayout?: 'auto' | 'fixed'
  scroll?: { x?: number | string; y?: number | string }
  sticky?: boolean
  emptyText?: string
}

export interface TableEmits {
  'update:sortInfo': [sortInfo: SortInfo]
  'update:filterInfo': [filterInfo: FilterInfo]
  'change': [pagination: any, filters: FilterInfo, sorter: SortInfo]
  'row-click': [record: TableData, index: number, event: Event]
  'row-dblclick': [record: TableData, index: number, event: Event]
  'row-contextmenu': [record: TableData, index: number, event: Event]
  'header-click': [column: TableColumn, event: Event]
  'selection-change': [selectedRowKeys: (string | number)[], selectedRows: TableData[]]
  'expand': [expanded: boolean, record: TableData]
}
