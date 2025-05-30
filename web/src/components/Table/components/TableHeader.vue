<template>
  <thead v-if="showHeader" class="table-header">
    <tr>
      <!-- 选择列 -->
      <th v-if="hasSelection" :class="headerCellClass" style="width: 50px;">
        <input
          v-if="rowSelection?.type === 'checkbox'"
          type="checkbox"
          :checked="isAllSelected"
          :indeterminate="isIndeterminate"
          @change="handleSelectAll"
          :disabled="!hasData"
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
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { TableColumn, SortInfo } from '../types'

defineOptions({
  name: 'TableHeader'
})

interface Props {
  columns: TableColumn[]
  showHeader: boolean
  hasSelection: boolean
  hasData: boolean
  size: 'small' | 'medium' | 'large'
  isAllSelected: boolean
  isIndeterminate: boolean
  rowSelection?: any
  currentSortInfo: SortInfo
}

interface Emits {
  'header-click': [column: TableColumn, event: Event]
  'select-all': [event: Event]
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const headerCellClass = computed(() => [
  'table-header-cell',
  `text-${props.size === 'small' ? 'xs' : props.size === 'large' ? 'sm' : 'xs'}`,
])

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

const getHeaderAlignClass = (column: TableColumn) => {
  const align = column.align || 'left'
  return {
    'justify-start': align === 'left',
    'justify-center': align === 'center',
    'justify-end': align === 'right',
  }
}

const getSortIconClass = (column: TableColumn, order: 'asc' | 'desc') => {
  const isActive = props.currentSortInfo.column === column.key && props.currentSortInfo.order === order
  return {
    'text-blue-500': isActive,
    'text-gray-300': !isActive,
  }
}

const handleHeaderClick = (column: TableColumn, event: Event) => {
  emit('header-click', column, event)
}

const handleSelectAll = (event: Event) => {
  emit('select-all', event)
}
</script>

<style scoped>
.table-header {
  @apply bg-gray-50;
}

.table-header-cell {
  @apply px-4 py-3 text-left font-medium text-gray-900 border-b border-gray-200;
}

.table-checkbox {
  @apply w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 focus:ring-2;
}

.sort-icon {
  @apply transition-colors duration-150;
}

.sort-icon-up {
  @apply -mb-1;
}

.sort-icon-down {
  @apply -mt-1;
}
</style>
