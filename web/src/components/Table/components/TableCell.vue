<template>
  <td :class="cellClass" :style="columnStyle">
    <div :class="contentClass">
      <!-- 自定义渲染 -->
      <component
        v-if="column.render"
        :is="renderContent"
      />
      <!-- 默认渲染 -->
      <span v-else :class="{ 'truncate': column.ellipsis }">
        {{ cellValue }}
      </span>
    </div>
  </td>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { TableColumn, TableData } from '../types'

defineOptions({
  name: 'TableCell'
})

interface Props {
  column: TableColumn
  record: TableData
  index: number
  size: 'small' | 'medium' | 'large'
}

const props = defineProps<Props>()

const cellValue = computed(() => {
  const dataIndex = props.column.dataIndex || props.column.key
  return props.record[dataIndex]
})

const renderContent = computed(() => {
  if (props.column.render) {
    return props.column.render(cellValue.value, props.record, props.index)
  }
  return null
})

const cellClass = computed(() => [
  'table-body-cell',
  `text-${props.size === 'small' ? 'xs' : props.size === 'large' ? 'base' : 'sm'}`,
  `text-${props.column.align || 'left'}`,
])

const contentClass = computed(() => [
  `text-${props.column.align || 'left'}`,
  {
    'truncate': props.column.ellipsis,
  }
])

const columnStyle = computed(() => {
  const style: any = {}
  if (props.column.width) {
    style.width = typeof props.column.width === 'number' ? `${props.column.width}px` : props.column.width
  }
  if (props.column.minWidth) {
    style.minWidth = `${props.column.minWidth}px`
  }
  if (props.column.maxWidth) {
    style.maxWidth = `${props.column.maxWidth}px`
  }
  return style
})
</script>

<style scoped>
.table-body-cell {
  @apply px-4 py-3 text-gray-700 border-b border-gray-200;
}
</style>
