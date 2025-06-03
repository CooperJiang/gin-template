<template>
  <div class="p-6 space-y-6">
    <h2 class="text-2xl font-bold">Element UI 风格表格测试</h2>

    <!-- 简单测试 -->
    <div class="bg-white p-4 rounded-lg border">
      <h3 class="text-lg font-semibold mb-4">基础功能测试</h3>
      <GlobalTable :data="testData" :bordered="true">
        <GlobalTableColumn prop="name" label="姓名" />
        <GlobalTableColumn prop="age" label="年龄" :width="100" />
        <GlobalTableColumn label="操作" :width="120">
          <template #default="{ row }">
            <button
              @click="handleClick(row)"
              class="px-2 py-1 text-sm bg-blue-500 text-white rounded hover:bg-blue-600"
            >
              点击
            </button>
          </template>
        </GlobalTableColumn>
      </GlobalTable>
    </div>

    <!-- 完整示例 -->
    <div class="bg-white p-4 rounded-lg border">
      <h3 class="text-lg font-semibold mb-4">完整功能示例</h3>
      <GlobalTable :data="fullData" :bordered="true" size="medium">
        <GlobalTableColumn prop="id" label="ID" :width="80" />
        <GlobalTableColumn prop="name" label="姓名" :sortable="true" />
        <GlobalTableColumn prop="email" label="邮箱" :ellipsis="true" />
        <GlobalTableColumn prop="department" label="部门" />
        <GlobalTableColumn label="状态" :width="120" align="center">
          <template #default="{ row }">
            <span
              :class="[
                'px-2 py-1 rounded-full text-xs',
                row.status === '正常' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800'
              ]"
            >
              {{ row.status }}
            </span>
          </template>
        </GlobalTableColumn>
        <GlobalTableColumn label="操作" :width="180" align="center">
          <template #default="{ row }">
            <div class="space-x-2">
              <button
                @click="editUser(row)"
                class="px-3 py-1 text-sm bg-blue-500 text-white rounded hover:bg-blue-600"
              >
                编辑
              </button>
              <button
                @click="deleteUser(row)"
                class="px-3 py-1 text-sm bg-red-500 text-white rounded hover:bg-red-600"
              >
                删除
              </button>
            </div>
          </template>
        </GlobalTableColumn>
      </GlobalTable>
    </div>

    <!-- 日志显示 -->
    <div class="bg-gray-50 p-4 rounded-lg border">
      <h3 class="text-lg font-semibold mb-2">操作日志</h3>
      <div v-if="logs.length === 0" class="text-gray-500">暂无操作</div>
      <div v-else class="space-y-1">
        <div
          v-for="(log, index) in logs"
          :key="index"
          class="text-sm bg-white p-2 rounded border"
        >
          {{ log }}
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

defineOptions({
  name: 'ElementTableTest'
})

// 测试数据
const testData = ref([
  { name: '张三', age: 28 },
  { name: '李四', age: 32 },
  { name: '王五', age: 25 }
])

const fullData = ref([
  { id: 1, name: '张三', email: 'zhangsan@example.com', department: '技术部', status: '正常' },
  { id: 2, name: '李四', email: 'lisi@example.com', department: '产品部', status: '正常' },
  { id: 3, name: '王五', email: 'wangwu@example.com', department: '设计部', status: '正常' }
])

// 操作日志
const logs = ref([])

// 事件处理
const handleClick = (row) => {
  logs.value.unshift(`点击了用户: ${row.name}`)
}

const editUser = (row) => {
  logs.value.unshift(`编辑用户: ${row.name} (ID: ${row.id})`)
}

const deleteUser = (row) => {
  logs.value.unshift(`删除用户: ${row.name} (ID: ${row.id})`)
}
</script>

<style scoped>
/* 组件样式 */
</style>
