<template>
  <div class="table-loading">
    <div class="flex items-center justify-center py-8 text-gray-500">
      <GlobalIcon name="arrow-path" size="md" class="animate-spin mr-2" />
      正在加载...
    </div>
  </div>
</template>

<script setup lang="ts">
defineOptions({
  name: 'TableLoading'
})
</script>

<style scoped>
.table-loading {
  @apply border border-gray-200 rounded-lg bg-white;
}
</style>
