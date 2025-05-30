<template>
  <div class="space-y-8">
    <!-- 基础表格 -->
    <div class="demo-section">
      <h4 class="demo-title">基础表格</h4>
      <p class="demo-description">最基本的表格展示，支持排序</p>

      <div class="demo-container">
        <GlobalTable
          :columns="basicColumns"
          :data="basicData"
          :loading="basicLoading"
        />
      </div>

      <div class="mt-4 flex gap-2">
        <GlobalButton
          size="sm"
          variant="primary"
          @click="refreshBasicData"
          :loading="basicLoading"
        >
          刷新数据
        </GlobalButton>
        <GlobalButton
          size="sm"
          variant="secondary"
          @click="toggleBasicLoading"
        >
          切换加载状态
        </GlobalButton>
      </div>
    </div>

    <!-- 选择功能 -->
    <div class="demo-section">
      <h4 class="demo-title">行选择功能</h4>
      <p class="demo-description">支持多选和单选，可以禁用特定行</p>

      <div class="mb-4 space-y-2">
        <div class="text-sm text-gray-600">
          已选择: {{ selectedRowKeys.length }} 项
          <span v-if="selectedRowKeys.length" class="ml-2">
            ({{ selectedRowKeys.join(', ') }})
          </span>
        </div>
        <div class="flex gap-2">
          <GlobalButton
            size="sm"
            variant="outline"
            @click="selectionType = selectionType === 'checkbox' ? 'radio' : 'checkbox'"
          >
            切换为{{ selectionType === 'checkbox' ? '单选' : '多选' }}
          </GlobalButton>
          <GlobalButton
            size="sm"
            variant="outline"
            @click="clearSelection"
          >
            清空选择
          </GlobalButton>
        </div>
      </div>

      <div class="demo-container">
        <GlobalTable
          :columns="selectionColumns"
          :data="selectionData"
          :row-selection="{
            type: selectionType,
            selectedRowKeys: selectedRowKeys,
            onChange: handleSelectionChange,
            getCheckboxProps: getCheckboxProps
          }"
        />
      </div>
    </div>

    <!-- 表格配置 -->
    <div class="demo-section">
      <h4 class="demo-title">表格配置</h4>
      <p class="demo-description">不同尺寸、边框、斑马纹等样式配置</p>

      <!-- 配置选项 -->
      <div class="mb-4 grid grid-cols-2 md:grid-cols-4 gap-4 p-4 bg-gray-50 rounded-lg">
        <div>
          <label class="block text-xs font-medium text-gray-700 mb-1">尺寸</label>
          <select
            v-model="tableConfig.size"
            class="w-full px-2 py-1 text-xs border border-gray-300 rounded"
          >
            <option value="small">Small</option>
            <option value="medium">Medium</option>
            <option value="large">Large</option>
          </select>
        </div>
        <div>
          <label class="block text-xs font-medium text-gray-700 mb-1">边框</label>
          <GlobalButton
            size="sm"
            :variant="tableConfig.bordered ? 'primary' : 'outline'"
            @click="tableConfig.bordered = !tableConfig.bordered"
            class="w-full"
          >
            {{ tableConfig.bordered ? '开启' : '关闭' }}
          </GlobalButton>
        </div>
        <div>
          <label class="block text-xs font-medium text-gray-700 mb-1">斑马纹</label>
          <GlobalButton
            size="sm"
            :variant="tableConfig.striped ? 'primary' : 'outline'"
            @click="tableConfig.striped = !tableConfig.striped"
            class="w-full"
          >
            {{ tableConfig.striped ? '开启' : '关闭' }}
          </GlobalButton>
        </div>
        <div>
          <label class="block text-xs font-medium text-gray-700 mb-1">悬停效果</label>
          <GlobalButton
            size="sm"
            :variant="tableConfig.hoverable ? 'primary' : 'outline'"
            @click="tableConfig.hoverable = !tableConfig.hoverable"
            class="w-full"
          >
            {{ tableConfig.hoverable ? '开启' : '关闭' }}
          </GlobalButton>
        </div>
      </div>

      <div class="demo-container">
        <GlobalTable
          :columns="configColumns"
          :data="configData"
          v-bind="tableConfig"
        />
      </div>
    </div>

    <!-- 固定高度和滚动 -->
    <div class="demo-section">
      <h4 class="demo-title">固定高度和滚动</h4>
      <p class="demo-description">表格固定高度，内容区域可滚动</p>

      <div class="demo-container">
        <GlobalTable
          :columns="scrollColumns"
          :data="scrollData"
          height="300"
          :scroll="{ x: 1200 }"
        />
      </div>
    </div>

    <!-- 分页表格 -->
    <div class="demo-section">
      <h4 class="demo-title">分页表格</h4>
      <p class="demo-description">集成分页组件，支持数据分页展示</p>

      <div class="demo-container">
        <GlobalTable
          :columns="paginationColumns"
          :data="paginationData"
          :pagination="{
            current: currentPage,
            pageSize: pageSize,
            total: totalItems,
            showSizeChanger: true,
            showQuickJumper: true,
            showTotal: true
          }"
          @change="handleTableChange"
        />
      </div>
    </div>

    <!-- 自定义渲染 -->
    <div class="demo-section">
      <h4 class="demo-title">自定义渲染</h4>
      <p class="demo-description">自定义单元格内容渲染，支持组件和HTML</p>

      <div class="demo-container">
        <GlobalTable
          :columns="customColumns"
          :data="customData"
          @row-click="handleRowClick"
        />
      </div>
    </div>

    <!-- 空数据状态 -->
    <div class="demo-section">
      <h4 class="demo-title">空数据状态</h4>
      <p class="demo-description">没有数据时的展示效果</p>

      <div class="demo-container">
        <GlobalTable
          :columns="basicColumns"
          :data="[]"
          empty-text="没有找到任何数据"
        />
      </div>
    </div>

    <!-- 使用说明 -->
    <div class="bg-gradient-to-r from-indigo-50 to-purple-50 p-6 rounded-lg border border-indigo-200">
      <h4 class="text-lg font-semibold text-indigo-900 mb-4 flex items-center gap-2">
        <GlobalIcon name="book-open" size="sm" color="text-indigo-600" />
        使用说明
      </h4>

      <div class="space-y-4 text-sm text-indigo-800">
        <div>
          <h5 class="font-semibold mb-2">基本用法</h5>
          <div class="bg-white p-3 rounded border border-indigo-200 font-mono text-xs">
            <div>&lt;GlobalTable</div>
            <div class="ml-2">:columns="columns"</div>
            <div class="ml-2">:data="tableData"</div>
            <div class="ml-2">:loading="loading"</div>
            <div class="ml-2">@row-click="handleRowClick"</div>
            <div>/&gt;</div>
          </div>
        </div>

        <div>
          <h5 class="font-semibold mb-2">列定义示例</h5>
          <div class="bg-white p-3 rounded border border-indigo-200 font-mono text-xs">
            <div>const columns = [</div>
            <div class="ml-2">{'{'} key: 'name', title: '姓名', sortable: true {'}'},</div>
            <div class="ml-2">{'{'} key: 'age', title: '年龄', align: 'center' {'}'},</div>
            <div class="ml-2">{'{'} key: 'action', title: '操作', render: (_, record) => h(Button, {'{'} onClick: () => edit(record) {'}'}, '编辑') {'}'}</div>
            <div>]</div>
          </div>
        </div>
      </div>
    </div>

    <!-- API 说明 -->
    <div class="demo-section">
      <h4 class="demo-title">Props API 说明</h4>
      <p class="demo-description">表格组件支持的属性配置</p>

      <div class="demo-container">
        <GlobalTable
          :columns="propsApiColumns"
          :data="propsApi"
          :bordered="true"
          size="small"
        />
      </div>
    </div>

    <!-- Events API 说明 -->
    <div class="demo-section">
      <h4 class="demo-title">Events API 说明</h4>
      <p class="demo-description">表格组件支持的事件回调</p>

      <div class="demo-container">
        <GlobalTable
          :columns="eventsApiColumns"
          :data="eventsApi"
          :bordered="true"
          size="small"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, h } from 'vue'

defineOptions({
  name: 'TableTest'
})

// 基础表格数据
const basicLoading = ref(false)
const basicColumns = [
  { key: 'id', title: 'ID', width: 80, sortable: true },
  { key: 'name', title: '姓名', sortable: true },
  { key: 'age', title: '年龄', width: 100, align: 'center' as const, sortable: true },
  { key: 'email', title: '邮箱', ellipsis: true },
  { key: 'department', title: '部门' },
]

const basicData = ref([
  { id: 1, name: '张三', age: 28, email: 'zhangsan@example.com', department: '技术部' },
  { id: 2, name: '李四', age: 32, email: 'lisi@example.com', department: '产品部' },
  { id: 3, name: '王五', age: 25, email: 'wangwu@example.com', department: '设计部' },
  { id: 4, name: '赵六', age: 29, email: 'zhaoliu@example.com', department: '运营部' },
  { id: 5, name: '孙七', age: 31, email: 'sunqi@example.com', department: '市场部' },
])

// 选择功能
const selectedRowKeys = ref<(string | number)[]>([])
const selectionType = ref<'checkbox' | 'radio'>('checkbox')

const selectionColumns = [
  { key: 'id', title: 'ID', width: 80 },
  { key: 'name', title: '姓名' },
  { key: 'status', title: '状态', width: 100 },
  { key: 'role', title: '角色' },
]

const selectionData = ref([
  { id: 1, name: '张三', status: '正常', role: '管理员' },
  { id: 2, name: '李四', status: '正常', role: '用户' },
  { id: 3, name: '王五', status: '禁用', role: '用户' },
  { id: 4, name: '赵六', status: '正常', role: '编辑' },
  { id: 5, name: '孙七', status: '正常', role: '用户' },
])

// 表格配置
const tableConfig = reactive({
  size: 'medium' as 'small' | 'medium' | 'large',
  bordered: true,
  striped: false,
  hoverable: true,
})

const configColumns = [
  { key: 'name', title: '名称' },
  { key: 'type', title: '类型' },
  { key: 'size', title: '大小', align: 'right' as const },
  { key: 'date', title: '日期' },
]

const configData = ref([
  { name: 'document.pdf', type: 'PDF文档', size: '2.3MB', date: '2024-01-15' },
  { name: 'image.jpg', type: '图片', size: '1.8MB', date: '2024-01-14' },
  { name: 'video.mp4', type: '视频', size: '15.6MB', date: '2024-01-13' },
  { name: 'archive.zip', type: '压缩包', size: '8.9MB', date: '2024-01-12' },
])

// 滚动表格
const scrollColumns = [
  { key: 'id', title: 'ID', width: 80, fixed: 'left' as const },
  { key: 'name', title: '姓名', width: 120, fixed: 'left' as const },
  { key: 'field1', title: '字段1', width: 150 },
  { key: 'field2', title: '字段2', width: 150 },
  { key: 'field3', title: '字段3', width: 150 },
  { key: 'field4', title: '字段4', width: 150 },
  { key: 'field5', title: '字段5', width: 150 },
  { key: 'field6', title: '字段6', width: 150 },
  { key: 'action', title: '操作', width: 120, fixed: 'right' as const },
]

const scrollData = ref(
  Array.from({ length: 20 }, (_, i) => ({
    id: i + 1,
    name: `用户${i + 1}`,
    field1: `数据${i + 1}-1`,
    field2: `数据${i + 1}-2`,
    field3: `数据${i + 1}-3`,
    field4: `数据${i + 1}-4`,
    field5: `数据${i + 1}-5`,
    field6: `数据${i + 1}-6`,
    action: '操作',
  }))
)

// 分页数据
const currentPage = ref(1)
const pageSize = ref(10)
const totalItems = ref(50)

const paginationColumns = [
  { key: 'id', title: 'ID', width: 80 },
  { key: 'title', title: '标题' },
  { key: 'category', title: '分类' },
  { key: 'views', title: '浏览量', align: 'right' as const },
  { key: 'date', title: '创建时间' },
]

const paginationData = ref([
  { id: 1, title: '文章标题1', category: '技术', views: 1234, date: '2024-01-15' },
  { id: 2, title: '文章标题2', category: '产品', views: 856, date: '2024-01-14' },
  { id: 3, title: '文章标题3', category: '设计', views: 743, date: '2024-01-13' },
  { id: 4, title: '文章标题4', category: '运营', views: 692, date: '2024-01-12' },
  { id: 5, title: '文章标题5', category: '市场', views: 581, date: '2024-01-11' },
])

// 自定义渲染
const customColumns = [
  { key: 'name', title: '姓名' },
  {
    key: 'avatar',
    title: '头像',
    width: 80,
    render: () => h('div', { class: 'w-8 h-8 bg-blue-500 rounded-full flex items-center justify-center text-white text-xs' }, 'A')
  },
  {
    key: 'status',
    title: '状态',
    render: (value: string) => {
      const colors = {
        '正常': 'bg-green-100 text-green-800',
        '禁用': 'bg-red-100 text-red-800',
        '待审': 'bg-yellow-100 text-yellow-800',
      }
      return h('span', {
        class: `px-2 py-1 rounded-full text-xs ${colors[value as keyof typeof colors] || 'bg-gray-100 text-gray-800'}`
      }, value)
    }
  },
  {
    key: 'action',
    title: '操作',
    render: (_, record: any) => h('div', { class: 'space-x-2' }, [
      h('button', {
        class: 'text-blue-600 hover:text-blue-800 text-sm',
        onClick: () => handleEdit(record)
      }, '编辑'),
      h('button', {
        class: 'text-red-600 hover:text-red-800 text-sm',
        onClick: () => handleDelete(record)
      }, '删除'),
    ])
  },
]

const customData = ref([
  { id: 1, name: '张三', status: '正常' },
  { id: 2, name: '李四', status: '禁用' },
  { id: 3, name: '王五', status: '待审' },
  { id: 4, name: '赵六', status: '正常' },
])

// 事件处理
const refreshBasicData = () => {
  basicLoading.value = true
  setTimeout(() => {
    // 模拟数据刷新
    basicData.value = basicData.value.map(item => ({
      ...item,
      age: Math.floor(Math.random() * 40) + 20
    }))
    basicLoading.value = false
  }, 1000)
}

const toggleBasicLoading = () => {
  basicLoading.value = !basicLoading.value
}

const handleSelectionChange = (keys: (string | number)[], rows: any[]) => {
  selectedRowKeys.value = keys
  console.log('选择变化:', keys, rows)
}

const clearSelection = () => {
  selectedRowKeys.value = []
}

const getCheckboxProps = (record: any) => {
  return {
    disabled: record.status === '禁用'
  }
}

const handleTableChange = (pagination: any, filters: any, sorter: any) => {
  currentPage.value = pagination.page
  pageSize.value = pagination.pageSize
  console.log('表格变化:', pagination, filters, sorter)
}

const handleRowClick = (record: any, index: number, event: Event) => {
  console.log('行点击:', record, index, event)
}

const handleEdit = (record: any) => {
  console.log('编辑:', record)
}

const handleDelete = (record: any) => {
  console.log('删除:', record)
}

// API 文档数据
const propsApi = [
  { name: 'columns', description: '表格列的配置', type: 'TableColumn[]', default: '[]' },
  { name: 'data', description: '表格数据', type: 'TableData[]', default: '[]' },
  { name: 'loading', description: '是否加载中', type: 'boolean', default: 'false' },
  { name: 'size', description: '表格大小', type: "'small' | 'medium' | 'large'", default: "'medium'" },
  { name: 'bordered', description: '是否有边框', type: 'boolean', default: 'true' },
  { name: 'striped', description: '是否显示斑马纹', type: 'boolean', default: 'false' },
  { name: 'hoverable', description: '是否可悬停', type: 'boolean', default: 'true' },
  { name: 'height', description: '表格高度', type: 'number | string', default: 'undefined' },
  { name: 'rowSelection', description: '行选择配置', type: 'object', default: 'undefined' },
  { name: 'pagination', description: '分页配置', type: 'boolean | object', default: 'false' },
  { name: 'emptyText', description: '空数据文案', type: 'string', default: "'暂无数据'" },
  { name: 'rowKey', description: '行数据的Key', type: 'string | function', default: "'id'" },
]

const eventsApi = [
  { name: 'row-click', description: '点击行时触发', params: '(record, index, event)' },
  { name: 'row-dblclick', description: '双击行时触发', params: '(record, index, event)' },
  { name: 'selection-change', description: '选择变化时触发', params: '(selectedRowKeys, selectedRows)' },
  { name: 'change', description: '分页、排序、筛选变化时触发', params: '(pagination, filters, sorter)' },
  { name: 'header-click', description: '点击表头时触发', params: '(column, event)' },
]

// API 表格列定义
const propsApiColumns = [
  {
    key: 'name',
    title: '参数',
    width: 150,
    render: (value: string) => h('code', { class: 'text-blue-600 bg-blue-50 px-2 py-1 rounded text-sm' }, value)
  },
  { key: 'description', title: '说明' },
  {
    key: 'type',
    title: '类型',
    width: 200,
    render: (value: string) => h('code', { class: 'text-green-600 bg-green-50 px-2 py-1 rounded text-sm' }, value)
  },
  {
    key: 'default',
    title: '默认值',
    width: 120,
    render: (value: string) => h('code', { class: 'text-orange-600 bg-orange-50 px-2 py-1 rounded text-sm' }, value)
  },
]

const eventsApiColumns = [
  {
    key: 'name',
    title: '事件名',
    width: 150,
    render: (value: string) => h('code', { class: 'text-blue-600 bg-blue-50 px-2 py-1 rounded text-sm' }, value)
  },
  { key: 'description', title: '说明' },
  {
    key: 'params',
    title: '参数',
    width: 250,
    render: (value: string) => h('code', { class: 'text-green-600 bg-green-50 px-2 py-1 rounded text-sm' }, value)
  },
]
</script>

<style scoped>
.demo-section {
  @apply bg-white rounded-lg border border-gray-200 p-6 mb-6;
}

.demo-title {
  @apply text-lg font-semibold text-gray-900 mb-4;
}

.demo-description {
  @apply text-sm text-gray-600 mb-4;
}

.demo-container {
  @apply bg-gray-50 p-4 rounded border border-gray-200;
}
</style>
