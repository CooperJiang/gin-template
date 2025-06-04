<template>
  <tr
    :class="rowClass"
    @click="handleRowClick"
    @dblclick="handleRowDblClick"
    @contextmenu="handleRowContextMenu"
  >
    <!-- 选择列 -->
    <td v-if="hasSelection" :class="selectionCellClass">
      <input
        :type="rowSelection?.type || 'checkbox'"
        :checked="isSelected"
        @change="handleRowSelect"
        :disabled="checkboxProps?.disabled"
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
          :is="getRenderContent(column)"
        />
        <!-- 默认渲染 -->
        <span v-else :class="{ 'truncate': column.ellipsis }">
          {{ getCellValue(column) }}
        </span>
      </div>
    </td>
  </tr>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { TableColumn, TableData } from '../types'

defineOptions({
  name: 'TableRow'
})

interface Props {
  record: TableData
  index: number
  columns: TableColumn[]
  size: 'small' | 'medium' | 'large'
  striped: boolean
  hasSelection: boolean
  isSelected: boolean
  rowSelection?: any
  checkboxProps?: { disabled?: boolean }
  rowKey: string | ((record: TableData) => string | number)
}

interface Emits {
  'row-click': [record: TableData, index: number, event: Event]
  'row-dblclick': [record: TableData, index: number, event: Event]
  'row-contextmenu': [record: TableData, index: number, event: Event]
  'row-select': [record: TableData, event: Event]
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const rowClass = computed(() => [
  'table-row',
  {
    'table-row-selected': props.isSelected,
    'table-row-disabled': props.record._disabled,
    'table-row-even': props.striped && props.index % 2 === 1,
  }
])

const selectionCellClass = computed(() => [
  'table-body-cell',
  `text-${props.size === 'small' ? 'xs' : props.size === 'large' ? 'base' : 'sm'}`,
])

// 单元格相关函数
const getCellValue = (column: TableColumn) => {
  const dataIndex = column.dataIndex || column.key
  return props.record[dataIndex]
}

const getRenderContent = (column: TableColumn) => {
  if (column.render) {
    return column.render(getCellValue(column), props.record, props.index)
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

const handleRowClick = (event: Event) => {
  emit('row-click', props.record, props.index, event)
}

const handleRowDblClick = (event: Event) => {
  emit('row-dblclick', props.record, props.index, event)
}

const handleRowContextMenu = (event: Event) => {
  emit('row-contextmenu', props.record, props.index, event)
}

const handleRowSelect = (event: Event) => {
  emit('row-select', props.record, event)
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
