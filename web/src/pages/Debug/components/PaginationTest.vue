<template>
  <div class="space-y-8">
    <!-- 基础分页 -->
    <GlobalCodeDemo
      title="基础分页"
      description="最简单的分页组件使用方式，只需要提供总数据量"
      :code="basicPaginationCode"
    >
      <GlobalPagination
        :current="basicPagination.current"
        :total="basicPagination.total"
        :showSizeChanger="false"
        :showQuickJumper="false"
        @change="handleBasicChange"
      />
    </GlobalCodeDemo>

    <!-- 带页数选择器 -->
    <GlobalCodeDemo
      title="带页数选择器"
      description="可以改变每页显示的条数，适用于用户希望自定义每页显示数量的场景"
      :code="sizeChangerPaginationCode"
    >
      <GlobalPagination
        :current="sizeChangerPagination.current"
        :pageSize="sizeChangerPagination.pageSize"
        :total="sizeChangerPagination.total"
        :showSizeChanger="true"
        :pageSizeOptions="[10, 20, 50, 100]"
        @change="handleSizeChangerChange"
        @showSizeChange="handleSizeChange"
      />

      <div class="mt-2 text-sm text-gray-600">
        当前页: {{ sizeChangerPagination.current }}, 每页: {{ sizeChangerPagination.pageSize }}
      </div>
    </GlobalCodeDemo>

    <!-- 带快速跳转 -->
    <GlobalCodeDemo
      title="带快速跳转"
      description="支持直接输入页码快速跳转，提高大数据量场景下的操作效率"
      :code="quickJumperPaginationCode"
    >
      <GlobalPagination
        :current="quickJumperPagination.current"
        :total="quickJumperPagination.total"
        :showQuickJumper="true"
        @change="handleQuickJumperChange"
      />
    </GlobalCodeDemo>

    <!-- 简洁模式 -->
    <GlobalCodeDemo
      title="简洁模式"
      description="简化的分页显示，适用于空间有限的场景"
      :code="simplePaginationCode"
    >
      <GlobalPagination
        :current="simplePagination.current"
        :total="simplePagination.total"
        :simple="true"
        @change="handleSimpleChange"
      />
    </GlobalCodeDemo>

    <!-- API 说明 -->
    <div class="demo-section">
      <h4 class="demo-title">Props API 说明</h4>
      <div class="overflow-x-auto">
        <table class="w-full border-collapse border border-gray-300 text-sm">
          <thead>
            <tr class="bg-gray-50">
              <th class="border border-gray-300 px-4 py-2 text-left">参数</th>
              <th class="border border-gray-300 px-4 py-2 text-left">说明</th>
              <th class="border border-gray-300 px-4 py-2 text-left">类型</th>
              <th class="border border-gray-300 px-4 py-2 text-left">默认值</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="prop in propsApi" :key="prop.name" class="hover:bg-gray-50">
              <td class="border border-gray-300 px-4 py-2 font-mono text-blue-600">{{ prop.name }}</td>
              <td class="border border-gray-300 px-4 py-2">{{ prop.description }}</td>
              <td class="border border-gray-300 px-4 py-2 font-mono text-green-600">{{ prop.type }}</td>
              <td class="border border-gray-300 px-4 py-2 font-mono text-orange-600">{{ prop.default }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue'

defineOptions({
  name: 'PaginationTest'
})

// 各种分页状态
const basicPagination = reactive({
  current: 1,
  total: 100
})

const sizeChangerPagination = reactive({
  current: 1,
  pageSize: 10,
  total: 200
})

const quickJumperPagination = reactive({
  current: 1,
  total: 500
})

const simplePagination = reactive({
  current: 1,
  total: 80
})

// 事件处理
const handleBasicChange = (page: number, pageSize: number) => {
  basicPagination.current = page
  console.log('基础分页变化:', { page, pageSize })
}

const handleSizeChangerChange = (page: number, pageSize: number) => {
  sizeChangerPagination.current = page
  sizeChangerPagination.pageSize = pageSize
  console.log('带页数选择器分页变化:', { page, pageSize })
}

const handleSizeChange = (current: number, size: number) => {
  console.log('页数大小变化:', { current, size })
}

const handleQuickJumperChange = (page: number, pageSize: number) => {
  quickJumperPagination.current = page
  console.log('快速跳转分页变化:', { page, pageSize })
}

const handleSimpleChange = (page: number, pageSize: number) => {
  simplePagination.current = page
  console.log('简洁模式分页变化:', { page, pageSize })
}

// API 文档数据
const propsApi = [
  { name: 'current', description: '当前页码', type: 'number', default: '1' },
  { name: 'pageSize', description: '每页条数', type: 'number', default: '10' },
  { name: 'total', description: '数据总数', type: 'number', default: '必填' },
  { name: 'pageSizeOptions', description: '指定每页可以显示多少条', type: 'number[]', default: '[10, 20, 50, 100]' },
  { name: 'showSizeChanger', description: '是否可以改变 pageSize', type: 'boolean', default: 'true' },
  { name: 'showQuickJumper', description: '是否显示快速跳转', type: 'boolean', default: 'false' },
  { name: 'showTotal', description: '是否显示总数', type: 'boolean', default: 'true' },
  { name: 'simple', description: '是否简洁模式', type: 'boolean', default: 'false' },
  { name: 'size', description: '分页器大小', type: "'small' | 'medium' | 'large'", default: "'medium'" },
  { name: 'disabled', description: '是否禁用', type: 'boolean', default: 'false' },
]

// 代码示例字符串 - 简化版本
const basicPaginationCode = `// 基础分页示例
<GlobalPagination
  :current="basicPagination.current"
  :total="basicPagination.total"
  :showSizeChanger="false"
  :showQuickJumper="false"
  @change="handleBasicChange"
/>`

const sizeChangerPaginationCode = `// 带页数选择器示例
<GlobalPagination
  :current="sizeChangerPagination.current"
  :pageSize="sizeChangerPagination.pageSize"
  :total="sizeChangerPagination.total"
  :showSizeChanger="true"
  :pageSizeOptions="[10, 20, 50, 100]"
  @change="handleSizeChangerChange"
  @showSizeChange="handleSizeChange"
/>`

const quickJumperPaginationCode = `// 快速跳转示例
<GlobalPagination
  :current="quickJumperPagination.current"
  :total="quickJumperPagination.total"
  :showQuickJumper="true"
  @change="handleQuickJumperChange"
/>`

const simplePaginationCode = `// 简洁模式示例
<GlobalPagination
  :current="simplePagination.current"
  :total="simplePagination.total"
  :simple="true"
  @change="handleSimpleChange"
/>`
</script>

<style scoped>
.demo-section {
  @apply bg-white rounded-lg border border-gray-200 p-6;
}

.demo-title {
  @apply text-lg font-semibold text-gray-900 mb-2;
}

.demo-description {
  @apply text-sm text-gray-600 mb-4;
}

.demo-container {
  @apply bg-gray-50 p-4 rounded border border-gray-200 mb-4;
}
</style>
