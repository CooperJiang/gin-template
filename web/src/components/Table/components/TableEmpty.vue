<template>
  <tr>
    <td :colspan="colspan" class="text-center py-8 text-gray-500">
      <div class="flex flex-col items-center">
        <GlobalIcon name="inbox" size="lg" class="text-gray-300 mb-2" />
        {{ emptyText }}
      </div>
    </td>
  </tr>
</template>

<script setup lang="ts">
defineOptions({
  name: 'TableEmpty'
})

interface Props {
  colspan: number
  emptyText: string
}

withDefaults(defineProps<Props>(), {
  emptyText: '暂无数据'
})
</script>
