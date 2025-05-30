<template>
  <tbody class="table-body">
    <!-- 空数据状态 -->
    <tr v-if="!data.length">
      <td :colspan="totalColumns" class="text-center py-8 text-gray-500">
        <div class="flex flex-col items-center">
          <GlobalIcon name="inbox" size="lg" class="text-gray-300 mb-2" />
          {{ emptyText }}
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
      <td v-if="hasSelection" :class="selectionCellClass">
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
        :class="getCellClass(column)"
        :style="getColumnStyle(column)"
      >
        <div :class="getContentClass(column)">
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
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { TableColumn, TableData } from '../types'

defineOptions({
  name: 'TableBody'
})

interface Props {
  data: TableData[]
  columns: TableColumn[]
  size: 'small' | 'medium' | 'large'
  striped: boolean
  hasSelection: boolean
  selectedRowKeys: (string | number)[]
  rowSelection?: any
  rowKey: string | ((record: TableData) => string | number)
  emptyText: string
}

interface Emits {
  'row-click': [record: TableData, index: number, event: Event]
  'row-dblclick': [record: TableData, index: number, event: Event]
  'row-contextmenu': [record: TableData, index: number, event: Event]
  'row-select': [record: TableData, event: Event]
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const totalColumns = computed(() => {
  let count = props.columns.length
  if (props.hasSelection) count++
  return count
})

const selectionCellClass = computed(() => [
  'table-body-cell',
  `text-${props.size === 'small' ? 'xs' : props.size === 'large' ? 'base' : 'sm'}`,
])

const getRowKey = (record: TableData, index?: number): string | number => {
  if (typeof props.rowKey === 'function') {
    return props.rowKey(record)
  }
  return record[props.rowKey] ?? index ?? ''
}

const isRowSelected = (record: TableData): boolean => {
  return props.selectedRowKeys.includes(getRowKey(record))
}

const getCheckboxProps = (record: TableData) => {
  return props.rowSelection?.getCheckboxProps?.(record) || {}
}

// 行样式
const getRowClass = (record: TableData, index: number) => [
  'table-row',
  {
    'table-row-selected': isRowSelected(record),
    'table-row-disabled': record._disabled,
    'table-row-even': props.striped && index % 2 === 1,
  }
]

// 单元格相关函数
const getCellValue = (column: TableColumn, record: TableData) => {
  const dataIndex = column.dataIndex || column.key
  return record[dataIndex]
}

const getRenderContent = (column: TableColumn, record: TableData, index: number) => {
  if (column.render) {
    return column.render(getCellValue(column, record), record, index)
  }
  return null
}

const getCellClass = (column: TableColumn) => [
  'table-body-cell',
  `text-${props.size === 'small' ? 'xs' : props.size === 'large' ? 'base' : 'sm'}`,
  `text-${column.align || 'left'}`,
]

const getContentClass = (column: TableColumn) => [
  `text-${column.align || 'left'}`,
  {
    'truncate': column.ellipsis,
  }
]

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
  emit('row-select', record, event)
}
</script>

<style scoped>
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

.table-body-cell {
  @apply px-4 py-3 text-gray-700 border-b border-gray-200;
}

.table-checkbox {
  @apply w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 focus:ring-2;
}
</style>
