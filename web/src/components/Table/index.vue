<template>
  <div class="table-container" :class="containerClass">
    <!-- 加载状态 -->
    <div v-if="loading" class="table-loading">
      <div class="flex items-center justify-center py-8 text-gray-500">
        <GlobalIcon name="arrow-path" size="md" class="animate-spin mr-2" />
        正在加载...
      </div>
    </div>

    <!-- 表格内容 -->
    <div v-else class="table-wrapper" :style="wrapperStyle">
      <table :class="tableClass" :style="tableStyle">
        <!-- 表头 -->
        <thead v-if="showHeader" class="table-header">
          <tr>
            <!-- 选择列 -->
            <th v-if="hasSelection" :class="headerCellClass" class="text-center" style="width: 50px;">
              <input
                v-if="rowSelection?.type === 'checkbox'"
                type="checkbox"
                :checked="isAllSelected"
                :indeterminate="isIndeterminate"
                @change="handleSelectAll"
                :disabled="!data.length"
                class="table-checkbox"
              />
            </th>

            <!-- 数据列 -->
            <th
              v-for="column in columns"
              :key="column.key"
              :class="getHeaderCellClass(column)"
              :style="getColumnStyle(column)"
              @click="handleHeaderClick(column, $event)"
            >
              <div class="flex items-center" :class="getHeaderAlignClass(column)">
                <span class="truncate">{{ column.title }}</span>

                <!-- 排序图标 -->
                <div v-if="column.sortable" class="ml-1 flex flex-col">
                  <GlobalIcon
                    name="chevron-up"
                    size="xs"
                    :class="getSortIconClass(column, 'asc')"
                    class="sort-icon sort-icon-up"
                  />
                  <GlobalIcon
                    name="chevron-down"
                    size="xs"
                    :class="getSortIconClass(column, 'desc')"
                    class="sort-icon sort-icon-down"
                  />
                </div>

                <!-- 筛选图标 -->
                <GlobalIcon
                  v-if="column.filterable"
                  name="funnel"
                  size="xs"
                  class="ml-1 text-gray-400 hover:text-gray-600 cursor-pointer"
                />
              </div>
            </th>
          </tr>
        </thead>

        <!-- 表体 -->
        <tbody class="table-body">
          <!-- 空数据状态 -->
          <tr v-if="!data.length">
            <td :colspan="totalColumns" class="p-0">
              <div class="table-empty-state">
                <div class="table-empty-icon">
                  <GlobalIcon name="inbox" size="xl" class="w-full h-full" />
                </div>
                <div class="table-empty-title">暂无数据</div>
                <div class="table-empty-description">
                  {{ emptyText || '当前没有可显示的数据，请尝试添加一些内容或调整筛选条件' }}
                </div>
              </div>
            </td>
          </tr>

          <!-- 数据行 -->
          <tr
            v-else
            v-for="(record, index) in data"
            :key="getRowKey(record, index)"
            :class="getRowClass(record, index)"
            @click="handleRowClick(record, index, $event)"
            @dblclick="handleRowDblClick(record, index, $event)"
            @contextmenu="handleRowContextMenu(record, index, $event)"
          >
            <!-- 选择列 -->
            <td v-if="hasSelection" :class="bodyCellClass" class="text-center">
              <input
                :type="rowSelection?.type || 'checkbox'"
                :checked="isRowSelected(record)"
                @change="handleRowSelect(record, $event)"
                :disabled="getCheckboxProps(record)?.disabled"
                class="table-checkbox"
              />
            </td>

            <!-- 数据列 -->
            <td
              v-for="column in columns"
              :key="column.key"
              :class="getBodyCellClass(column)"
              :style="getColumnStyle(column)"
            >
              <div :class="getCellContentClass(column)">
                <!-- 自定义渲染 -->
                <component
                  v-if="column.render"
                  :is="getRenderContent(column, record, index)"
                />
                <!-- 默认渲染 -->
                <span v-else :class="{ 'truncate': column.ellipsis }">
                  {{ getCellValue(column, record) }}
                </span>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- 分页器 -->
    <div v-if="showPagination" class="table-pagination">
      <GlobalPagination
        v-bind="paginationProps"
        @change="handlePaginationChange"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import type { TableProps, TableEmits, SortInfo, TableData, TableColumn } from './types'

defineOptions({
  name: 'GlobalTable',
})

const props = withDefaults(defineProps<TableProps>(), {
  loading: false,
  size: 'medium',
  bordered: true,
  striped: false,
  hoverable: true,
  showHeader: true,
  rowKey: 'id',
  pagination: false,
  tableLayout: 'auto',
  sticky: false,
  emptyText: '暂无数据',
})

const emit = defineEmits<TableEmits>()

// 内部状态
const selectedRowKeys = ref<(string | number)[]>(props.rowSelection?.selectedRowKeys || [])
const currentSortInfo = ref<SortInfo>({ column: '', order: null })

// 计算属性
const hasSelection = computed(() => !!props.rowSelection)
const totalColumns = computed(() => {
  let count = props.columns.length
  if (hasSelection.value) count++
  return count
})

const showPagination = computed(() => {
  return typeof props.pagination === 'object' || props.pagination === true
})

const paginationProps = computed(() => {
  if (typeof props.pagination === 'object') {
    return props.pagination
  }
  return {}
})

// 样式计算
const containerClass = computed(() => [
  'table-container',
  `table-size-${props.size}`,
  {
    'table-bordered': props.bordered,
    'table-striped': props.striped,
    'table-hoverable': props.hoverable,
  }
])

const tableClass = computed(() => [
  'table',
  {
    'table-fixed': props.tableLayout === 'fixed',
  }
])

const wrapperStyle = computed(() => {
  const style: any = {}
  if (props.height) {
    style.height = typeof props.height === 'number' ? `${props.height}px` : props.height
    style.overflowY = 'auto'
  }
  if (props.maxHeight) {
    style.maxHeight = typeof props.maxHeight === 'number' ? `${props.maxHeight}px` : props.maxHeight
    style.overflowY = 'auto'
  }
  if (props.scroll?.x) {
    style.overflowX = 'auto'
  }
  return style
})

const tableStyle = computed(() => {
  const style: any = {}
  if (props.scroll?.x) {
    style.minWidth = typeof props.scroll.x === 'number' ? `${props.scroll.x}px` : props.scroll.x
  }
  return style
})

const headerCellClass = computed(() => [
  'table-header-cell',
  `text-${props.size === 'small' ? 'xs' : props.size === 'large' ? 'sm' : 'xs'}`,
])

const bodyCellClass = computed(() => [
  'table-body-cell',
  `text-${props.size === 'small' ? 'xs' : props.size === 'large' ? 'base' : 'sm'}`,
])

// 选择相关
const isAllSelected = computed(() => {
  if (!props.data.length) return false
  const selectableRows = props.data.filter((record: TableData) => !getCheckboxProps(record)?.disabled)
  return selectableRows.length > 0 && selectableRows.every((record: TableData) =>
    selectedRowKeys.value.includes(getRowKey(record))
  )
})

const isIndeterminate = computed(() => {
  const selectableRows = props.data.filter((record: TableData) => !getCheckboxProps(record)?.disabled)
  const selectedCount = selectableRows.filter((record: TableData) =>
    selectedRowKeys.value.includes(getRowKey(record))
  ).length
  return selectedCount > 0 && selectedCount < selectableRows.length
})

// 工具函数
const getRowKey = (record: TableData, index?: number): string | number => {
  if (typeof props.rowKey === 'function') {
    return props.rowKey(record)
  }
  return record[props.rowKey] ?? index ?? ''
}

const getCheckboxProps = (record: TableData) => {
  return props.rowSelection?.getCheckboxProps?.(record) || {}
}

const isRowSelected = (record: TableData): boolean => {
  return selectedRowKeys.value.includes(getRowKey(record))
}

// 样式函数
const getColumnStyle = (column: TableColumn) => {
  const style: any = {}
  if (column.width) {
    style.width = typeof column.width === 'number' ? `${column.width}px` : column.width
  }
  if (column.minWidth) {
    style.minWidth = `${column.minWidth}px`
  }
  if (column.maxWidth) {
    style.maxWidth = `${column.maxWidth}px`
  }
  return style
}

const getHeaderCellClass = (column: TableColumn) => [
  ...headerCellClass.value,
  {
    'cursor-pointer': column.sortable || column.filterable,
    'select-none': column.sortable || column.filterable,
  }
]

const getBodyCellClass = (column: TableColumn) => [
  ...bodyCellClass.value,
  `text-${column.align || 'left'}`,
]

const getHeaderAlignClass = (column: TableColumn) => {
  const align = column.align || 'left'
  return {
    'justify-start': align === 'left',
    'justify-center': align === 'center',
    'justify-end': align === 'right',
  }
}

const getCellContentClass = (column: TableColumn) => [
  `text-${column.align || 'left'}`,
  {
    'truncate': column.ellipsis,
  }
]

const getRowClass = (record: TableData, index: number) => [
  'table-row',
  {
    'table-row-selected': isRowSelected(record),
    'table-row-disabled': record._disabled,
    'table-row-even': props.striped && index % 2 === 1,
  }
]

const getSortIconClass = (column: TableColumn, order: 'asc' | 'desc') => {
  const isActive = currentSortInfo.value.column === column.key && currentSortInfo.value.order === order
  return {
    'text-blue-500': isActive,
    'text-gray-300': !isActive,
  }
}

const getCellValue = (column: TableColumn, record: TableData) => {
  const dataIndex = column.dataIndex || column.key
  return record[dataIndex]
}

const getRenderContent = (column: TableColumn, record: TableData, index: number) => {
  if (column.render) {
    const value = getCellValue(column, record)
    return column.render(value, record, index)
  }
  return null
}

// 事件处理
const handleHeaderClick = (column: TableColumn, event: Event) => {
  emit('header-click', column, event)

  if (column.sortable) {
    let newOrder: 'asc' | 'desc' | null = 'asc'

    if (currentSortInfo.value.column === column.key) {
      if (currentSortInfo.value.order === 'asc') {
        newOrder = 'desc'
      } else if (currentSortInfo.value.order === 'desc') {
        newOrder = null
      }
    }

    const newSortInfo: SortInfo = {
      column: newOrder ? column.key : '',
      order: newOrder
    }

    currentSortInfo.value = newSortInfo
    emit('update:sortInfo', newSortInfo)
  }
}

const handleSelectAll = (event: Event) => {
  const target = event.target as HTMLInputElement
  const checked = target.checked

  const selectableRows = props.data.filter((record: TableData) => !getCheckboxProps(record)?.disabled)

  if (checked) {
    const newSelectedKeys = [...selectedRowKeys.value]
    selectableRows.forEach(record => {
      const key = getRowKey(record)
      if (!newSelectedKeys.includes(key)) {
        newSelectedKeys.push(key)
      }
    })
    selectedRowKeys.value = newSelectedKeys
  } else {
    const selectableKeys = selectableRows.map(record => getRowKey(record))
    selectedRowKeys.value = selectedRowKeys.value.filter(key => !selectableKeys.includes(key))
  }

  const selectedRows = props.data.filter((record: TableData) => selectedRowKeys.value.includes(getRowKey(record)))
  props.rowSelection?.onChange?.(selectedRowKeys.value, selectedRows)
  props.rowSelection?.onSelectAll?.(checked, selectedRows, selectableRows)
  emit('selection-change', selectedRowKeys.value, selectedRows)
}

const handleRowClick = (record: TableData, index: number, event: Event) => {
  emit('row-click', record, index, event)
}

const handleRowDblClick = (record: TableData, index: number, event: Event) => {
  emit('row-dblclick', record, index, event)
}

const handleRowContextMenu = (record: TableData, index: number, event: Event) => {
  emit('row-contextmenu', record, index, event)
}

const handleRowSelect = (record: TableData, event: Event) => {
  const target = event.target as HTMLInputElement
  const checked = target.checked
  const rowKey = getRowKey(record)

  if (props.rowSelection?.type === 'radio') {
    selectedRowKeys.value = checked ? [rowKey] : []
  } else {
    if (checked) {
      if (!selectedRowKeys.value.includes(rowKey)) {
        selectedRowKeys.value.push(rowKey)
      }
    } else {
      const index = selectedRowKeys.value.indexOf(rowKey)
      if (index > -1) {
        selectedRowKeys.value.splice(index, 1)
      }
    }
  }

  const selectedRows = props.data.filter((record: TableData) => selectedRowKeys.value.includes(getRowKey(record)))
  props.rowSelection?.onChange?.(selectedRowKeys.value, selectedRows)
  props.rowSelection?.onSelect?.(record, checked, selectedRows)
  emit('selection-change', selectedRowKeys.value, selectedRows)
}

const handlePaginationChange = (page: number, pageSize: number) => {
  emit('change', { page, pageSize }, props.filterInfo || {}, currentSortInfo.value)
}

// 监听外部变化
watch(() => props.rowSelection?.selectedRowKeys, (newKeys) => {
  if (newKeys) {
    selectedRowKeys.value = [...newKeys]
  }
}, { immediate: true })

watch(() => props.sortInfo, (newSortInfo) => {
  if (newSortInfo) {
    currentSortInfo.value = { ...newSortInfo }
  }
}, { immediate: true })
</script>

<style scoped>
/* 表格容器 */
.table-container {
  @apply bg-white rounded-lg border border-gray-200 overflow-hidden;
}

/* 表格样式 */
.table {
  @apply w-full border-collapse;
}

.table-fixed {
  table-layout: fixed;
}

/* 表头样式 */
.table-header {
  @apply bg-gray-50;
}

.table-header-cell {
  @apply px-4 py-3 text-left font-medium text-gray-900 border-b border-gray-200;
}

/* 表体样式 */
.table-body-cell {
  @apply px-4 py-3 text-gray-700 border-b border-gray-200;
}

/* 行样式 */
.table-row {
  @apply transition-colors duration-150;
}

.table-row-selected {
  @apply bg-blue-50;
}

.table-row-disabled {
  @apply opacity-50 cursor-not-allowed;
}

.table-row-even {
  @apply bg-gray-50;
}

/* 尺寸变体 */
.table-size-small .table-header-cell,
.table-size-small .table-body-cell {
  @apply px-2 py-2;
}

.table-size-large .table-header-cell,
.table-size-large .table-body-cell {
  @apply px-6 py-4;
}

/* 边框样式增强 */
.table-bordered {
  @apply border border-gray-300 rounded-lg overflow-hidden;
}

.table-bordered .table-header-cell,
.table-bordered .table-body-cell {
  @apply border-r border-gray-300 last:border-r-0;
}

.table-bordered .table-row {
  @apply border-b border-gray-200;
}

.table-bordered .table-row:last-child {
  @apply border-b-0;
}

/* 无边框样式 */
.table-container:not(.table-bordered) {
  @apply border-0;
}

.table-container:not(.table-bordered) .table-header-cell,
.table-container:not(.table-bordered) .table-body-cell {
  @apply border-r-0;
}

.table-container:not(.table-bordered) .table-row {
  @apply border-b border-gray-100;
}

/* 悬停效果 */
.table-hoverable .table-row:hover {
  @apply bg-gray-50;
}

/* 复选框样式 */
.table-checkbox {
  @apply w-4 h-4 text-blue-600 bg-white border-2 border-gray-300 rounded focus:ring-blue-500 focus:ring-2 cursor-pointer;
}

.table-checkbox:checked {
  @apply bg-blue-600 border-blue-600;
}

.table-checkbox:checked::before {
  content: '';
  @apply block w-2 h-2 bg-white;
  mask: url("data:image/svg+xml,%3csvg viewBox='0 0 16 16' fill='white' xmlns='http://www.w3.org/2000/svg'%3e%3cpath d='m13.854 3.646-7.5 7.5a.5.5 0 0 1-.708 0l-3.5-3.5a.5.5 0 1 1 .708-.708L6 10.293l7.146-7.147a.5.5 0 0 1 .708.708z'/%3e%3c/svg%3e") no-repeat center;
  mask-size: contain;
}

.table-checkbox:indeterminate {
  @apply bg-blue-600 border-blue-600;
}

.table-checkbox:indeterminate::before {
  content: '';
  @apply block w-2 h-1 bg-white;
  mask: url("data:image/svg+xml,%3csvg viewBox='0 0 16 16' fill='white' xmlns='http://www.w3.org/2000/svg'%3e%3cpath d='M4 8a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7A.5.5 0 0 1 4 8z'/%3e%3c/svg%3e") no-repeat center;
  mask-size: contain;
}

/* 排序图标 */
.sort-icon {
  @apply transition-colors duration-150;
}

.sort-icon-up {
  @apply -mb-1;
}

.sort-icon-down {
  @apply -mt-1;
}

/* 加载状态 */
.table-loading {
  @apply border border-gray-200 rounded-lg bg-white;
}

/* 分页器 */
.table-pagination {
  @apply px-4 py-3 border-t border-gray-200 bg-gray-50;
}

/* 空数据状态美化 */
.table-empty-state {
  @apply flex flex-col items-center justify-center py-16 px-6 text-gray-500;
}

.table-empty-icon {
  @apply w-20 h-20 text-gray-300 mb-4;
}

.table-empty-title {
  @apply text-lg font-medium text-gray-900 mb-2;
}

.table-empty-description {
  @apply text-sm text-gray-500 text-center max-w-sm;
}

/* 滚动条样式 */
.table-wrapper::-webkit-scrollbar {
  width: 6px;
  height: 6px;
}

.table-wrapper::-webkit-scrollbar-track {
  background: #f1f5f9;
}

.table-wrapper::-webkit-scrollbar-thumb {
  background: #cbd5e1;
  border-radius: 3px;
}

.table-wrapper::-webkit-scrollbar-thumb:hover {
  background: #94a3b8;
}
</style>
