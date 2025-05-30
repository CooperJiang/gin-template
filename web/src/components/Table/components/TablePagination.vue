<template>
  <div v-if="showPagination" class="table-pagination">
    <GlobalPagination
      v-bind="paginationProps"
      @change="handlePaginationChange"
    />
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

defineOptions({
  name: 'TablePagination'
})

interface Props {
  pagination: boolean | object
}

interface Emits {
  'pagination-change': [page: number, pageSize: number]
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const showPagination = computed(() => {
  return typeof props.pagination === 'object' || props.pagination === true
})

const paginationProps = computed(() => {
  if (typeof props.pagination === 'object') {
    return props.pagination
  }
  return {}
})

const handlePaginationChange = (page: number, pageSize: number) => {
  emit('pagination-change', page, pageSize)
}
</script>

<style scoped>
.table-pagination {
  @apply px-4 py-3 border-t border-gray-200 bg-gray-50;
}
</style>
