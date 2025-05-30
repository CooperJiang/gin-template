<template>
  <div class="space-y-8">
    <!-- 基础分页 -->
    <div class="demo-section">
      <h4 class="demo-title">基础分页</h4>
      <p class="demo-description">
        最简单的分页组件使用方式，只需要提供总数据量。默认不开启快速跳转和页数选择器。
      </p>
      <div class="demo-container">
        <GlobalPagination
          :current="basicPagination.current"
          :total="basicPagination.total"
          :showSizeChanger="false"
          :showQuickJumper="false"
          @change="handleBasicChange"
        />
      </div>
      <div class="demo-code">
        <code>
&lt;GlobalPagination
  :current="{{ basicPagination.current }}"
  :total="{{ basicPagination.total }}"
  @change="handleBasicChange"
/&gt;
        </code>
      </div>
    </div>

    <!-- 带页数选择器 -->
    <div class="demo-section">
      <h4 class="demo-title">带页数选择器</h4>
      <p class="demo-description">
        可以改变每页显示的条数，适用于用户希望自定义每页显示数量的场景。
      </p>
      <div class="demo-container">
        <GlobalPagination
          :current="sizeChangerPagination.current"
          :pageSize="sizeChangerPagination.pageSize"
          :total="sizeChangerPagination.total"
          :showSizeChanger="true"
          :pageSizeOptions="[10, 20, 50, 100]"
          @change="handleSizeChangerChange"
          @showSizeChange="handleSizeChange"
        />
      </div>
      <div class="demo-info">
        当前页: {{ sizeChangerPagination.current }}, 每页: {{ sizeChangerPagination.pageSize }}
      </div>
    </div>

    <!-- 带快速跳转 -->
    <div class="demo-section">
      <h4 class="demo-title">带快速跳转</h4>
      <p class="demo-description">
        支持直接输入页码快速跳转，提高大数据量场景下的操作效率。
      </p>
      <div class="demo-container">
        <GlobalPagination
          :current="quickJumperPagination.current"
          :total="quickJumperPagination.total"
          :showQuickJumper="true"
          @change="handleQuickJumperChange"
        />
      </div>
    </div>

    <!-- 简洁模式 -->
    <div class="demo-section">
      <h4 class="demo-title">简洁模式</h4>
      <p class="demo-description">
        简化的分页显示，适用于空间有限的场景。
      </p>
      <div class="demo-container">
        <GlobalPagination
          :current="simplePagination.current"
          :total="simplePagination.total"
          :simple="true"
          @change="handleSimpleChange"
        />
      </div>
    </div>

    <!-- 不同尺寸 -->
    <div class="demo-section">
      <h4 class="demo-title">不同尺寸</h4>
      <p class="demo-description">
        提供小、中、大三种尺寸，适应不同的设计需求。
      </p>
      <div class="demo-container space-y-4">
        <div>
          <div class="text-sm text-gray-600 mb-2">小尺寸 (Small)</div>
          <GlobalPagination
            :current="sizePagination.current"
            :total="200"
            size="small"
            @change="handleSizeChange"
          />
        </div>
        <div>
          <div class="text-sm text-gray-600 mb-2">中等尺寸 (Medium) - 默认</div>
          <GlobalPagination
            :current="sizePagination.current"
            :total="200"
            size="medium"
            @change="handleSizeChange"
          />
        </div>
        <div>
          <div class="text-sm text-gray-600 mb-2">大尺寸 (Large)</div>
          <GlobalPagination
            :current="sizePagination.current"
            :total="200"
            size="large"
            @change="handleSizeChange"
          />
        </div>
      </div>
    </div>

    <!-- 不同位置 -->
    <div class="demo-section">
      <h4 class="demo-title">不同位置</h4>
      <p class="demo-description">
        支持左对齐、居中、右对齐三种位置设置。
      </p>
      <div class="demo-container space-y-4">
        <div>
          <div class="text-sm text-gray-600 mb-2">左对齐</div>
          <GlobalPagination
            :current="positionPagination.current"
            :total="150"
            position="left"
            @change="handlePositionChange"
          />
        </div>
        <div>
          <div class="text-sm text-gray-600 mb-2">居中对齐</div>
          <GlobalPagination
            :current="positionPagination.current"
            :total="150"
            position="center"
            @change="handlePositionChange"
          />
        </div>
        <div>
          <div class="text-sm text-gray-600 mb-2">右对齐 (默认)</div>
          <GlobalPagination
            :current="positionPagination.current"
            :total="150"
            position="right"
            @change="handlePositionChange"
          />
        </div>
      </div>
    </div>

    <!-- 自定义总数显示 -->
    <div class="demo-section">
      <h4 class="demo-title">自定义总数显示</h4>
      <p class="demo-description">
        可以自定义总数显示格式，或通过插槽完全自定义显示内容。
      </p>
      <div class="demo-container space-y-4">
        <div>
          <div class="text-sm text-gray-600 mb-2">自定义模板函数</div>
          <GlobalPagination
            :current="customTotalPagination.current"
            :total="customTotalPagination.total"
            :totalTemplate="customTotalTemplate"
            @change="handleCustomTotalChange"
          />
        </div>
        <div>
          <div class="text-sm text-gray-600 mb-2">自定义插槽</div>
          <GlobalPagination
            :current="customTotalPagination.current"
            :total="customTotalPagination.total"
            @change="handleCustomTotalChange"
          >
            <template #total="{ total, range }">
              <div class="flex items-center gap-2 text-sm">
                <GlobalIcon name="document-text" size="sm" color="text-blue-500" />
                <span class="text-blue-600 font-medium">
                  第 {{ range[0] }}-{{ range[1] }} 项，总共 {{ total }} 项数据
                </span>
              </div>
            </template>
          </GlobalPagination>
        </div>
      </div>
    </div>

    <!-- 禁用状态 -->
    <div class="demo-section">
      <h4 class="demo-title">禁用状态</h4>
      <p class="demo-description">
        分页器的禁用状态，适用于数据加载中或其他不可操作的场景。
      </p>
      <div class="demo-container">
        <GlobalPagination
          :current="3"
          :total="100"
          :disabled="true"
        />
      </div>
    </div>

    <!-- 隐藏单页 -->
    <div class="demo-section">
      <h4 class="demo-title">隐藏单页分页器</h4>
      <p class="demo-description">
        当总页数为1时自动隐藏分页器，避免不必要的UI元素。
      </p>
      <div class="demo-container space-y-4">
        <div>
          <div class="text-sm text-gray-600 mb-2">单页数据 (隐藏分页器)</div>
          <GlobalPagination
            :current="1"
            :total="5"
            :hideOnSinglePage="true"
          />
          <div class="text-xs text-gray-500 mt-2">↑ 由于只有1页，分页器被隐藏</div>
        </div>
        <div>
          <div class="text-sm text-gray-600 mb-2">多页数据 (显示分页器)</div>
          <GlobalPagination
            :current="1"
            :total="25"
            :hideOnSinglePage="true"
          />
        </div>
      </div>
    </div>

    <!-- 完整功能示例 -->
    <div class="demo-section">
      <h4 class="demo-title">完整功能示例</h4>
      <p class="demo-description">
        包含所有功能的分页器示例：页数选择、快速跳转、自定义总数显示等。
      </p>
      <div class="demo-container">
        <GlobalPagination
          :current="fullFeaturePagination.current"
          :pageSize="fullFeaturePagination.pageSize"
          :total="fullFeaturePagination.total"
          :showSizeChanger="true"
          :showQuickJumper="true"
          :showTotal="true"
          :pageSizeOptions="[10, 20, 50, 100]"
          @change="handleFullFeatureChange"
          @showSizeChange="handleFullFeatureSizeChange"
        />
      </div>
      <div class="demo-info">
        <div class="text-sm text-gray-600">
          当前状态: 第 {{ fullFeaturePagination.current }} 页，每页 {{ fullFeaturePagination.pageSize }} 条，
          总共 {{ fullFeaturePagination.total }} 条数据，共 {{ Math.ceil(fullFeaturePagination.total / fullFeaturePagination.pageSize) }} 页
        </div>
      </div>
    </div>

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

    <!-- Events API 说明 -->
    <div class="demo-section">
      <h4 class="demo-title">Events API 说明</h4>
      <div class="overflow-x-auto">
        <table class="w-full border-collapse border border-gray-300 text-sm">
          <thead>
            <tr class="bg-gray-50">
              <th class="border border-gray-300 px-4 py-2 text-left">事件名</th>
              <th class="border border-gray-300 px-4 py-2 text-left">说明</th>
              <th class="border border-gray-300 px-4 py-2 text-left">参数</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="event in eventsApi" :key="event.name" class="hover:bg-gray-50">
              <td class="border border-gray-300 px-4 py-2 font-mono text-blue-600">{{ event.name }}</td>
              <td class="border border-gray-300 px-4 py-2">{{ event.description }}</td>
              <td class="border border-gray-300 px-4 py-2 font-mono text-green-600">{{ event.params }}</td>
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

const sizePagination = reactive({
  current: 3
})

const positionPagination = reactive({
  current: 2
})

const customTotalPagination = reactive({
  current: 1,
  total: 95
})

const fullFeaturePagination = reactive({
  current: 1,
  pageSize: 20,
  total: 1000
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

const handlePositionChange = (page: number, pageSize: number) => {
  positionPagination.current = page
  console.log('位置分页变化:', { page, pageSize })
}

const handleCustomTotalChange = (page: number, pageSize: number) => {
  customTotalPagination.current = page
  console.log('自定义总数分页变化:', { page, pageSize })
}

const handleFullFeatureChange = (page: number, pageSize: number) => {
  fullFeaturePagination.current = page
  fullFeaturePagination.pageSize = pageSize
  console.log('完整功能分页变化:', { page, pageSize })
}

const handleFullFeatureSizeChange = (current: number, size: number) => {
  fullFeaturePagination.current = current
  fullFeaturePagination.pageSize = size
  console.log('完整功能页数变化:', { current, size })
}

// 自定义总数模板
const customTotalTemplate = (total: number, range: [number, number]) => {
  return `共 ${total} 条记录，当前显示第 ${range[0]} - ${range[1]} 条`
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
  { name: 'position', description: '分页器位置', type: "'left' | 'center' | 'right'", default: "'right'" },
  { name: 'disabled', description: '是否禁用', type: 'boolean', default: 'false' },
  { name: 'hideOnSinglePage', description: '只有一页时是否隐藏分页器', type: 'boolean', default: 'false' },
]

const eventsApi = [
  { name: 'change', description: '页码或每页条数改变的回调', params: '(page: number, pageSize: number)' },
  { name: 'showSizeChange', description: 'pageSize 变化的回调', params: '(current: number, size: number)' },
]
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

.demo-code {
  @apply bg-gray-900 text-gray-100 p-3 rounded text-xs font-mono overflow-x-auto;
}

.demo-info {
  @apply text-xs text-gray-500 bg-blue-50 p-2 rounded border border-blue-200;
}
</style>
