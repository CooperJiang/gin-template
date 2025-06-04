<template>
  <div class="p-6">
    <h2 class="text-xl font-bold mb-4">Simple Table Test</h2>

    <div class="mb-4">
      <p>测试数据: {{ testData }}</p>
      <p>Merged columns: {{ mergedColumns }}</p>
    </div>

    <!-- 基础测试 -->
    <div class="mb-6">
      <h3 class="text-lg font-semibold mb-2">基础 GlobalTable 测试</h3>
      <GlobalTable :data="testData" :bordered="true">
        <GlobalTableColumn prop="name" label="姓名" />
        <GlobalTableColumn prop="age" label="年龄" :width="100" />
      </GlobalTable>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

const testData = ref([
  { name: '张三', age: 28 },
  { name: '李四', age: 32 }
])

// 暴露一些数据用于调试
const mergedColumns = computed(() => {
  return 'Debug info here'
})
</script>