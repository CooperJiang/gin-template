<template>
  <div class="space-y-8">
    <!-- 简单功能测试 -->
    <GlobalCodeDemo
      title="功能测试 - Element UI 风格"
      description="测试 GlobalTableColumn 子组件是否正常工作"
      :code="simpleTestCode"
    >
      <GlobalTable :data="simpleTestData" :bordered="true">
        <GlobalTableColumn prop="name" label="姓名" />
        <GlobalTableColumn prop="age" label="年龄" :width="100" />
        <GlobalTableColumn label="操作" :width="100">
          <template #default="{ row }">
            <button @click="alert(`点击了${row.name}`)" class="text-blue-600">点击</button>
          </template>
        </GlobalTableColumn>
      </GlobalTable>
    </GlobalCodeDemo>

    <!-- Slot 方式表格 -->
    <GlobalCodeDemo
      title="自定义渲染表格 (类似 Element UI 风格)"
      description="使用 render 函数自定义列渲染，支持灵活的内容展示和操作按钮"
      :code="slotTableCode"
    >
      <GlobalTable
        :data="slotTableData"
        :columns="slotTableColumns"
        :loading="false"
        :bordered="true"
        size="medium"
      />
    </GlobalCodeDemo>

    <!-- 真正的 Element UI 风格表格 -->
    <GlobalCodeDemo
      title="Element UI 风格表格 (子组件方式)"
      description="使用 GlobalTableColumn 子组件定义表格列，完全仿照 Element UI 的使用方式"
      :code="elementStyleTableCode"
    >
      <div class="p-4 bg-green-50 border border-green-200 rounded-lg mb-4">
        <h4 class="text-green-800 font-semibold mb-2">✅ Element UI 风格已支持</h4>
        <div class="text-sm text-green-700 space-y-1">
          <div>✅ 子组件方式定义列 (GlobalTableColumn)</div>
          <div>✅ slot 自定义渲染</div>
          <div>✅ 完整的列属性支持</div>
          <div>✅ 与 Element UI 使用方式一致</div>
        </div>
      </div>

      <!-- 真正的 Element UI 风格实现 -->
      <GlobalTable
        :data="elementTableData"
        :loading="false"
        :bordered="true"
        size="medium"
        style="width: 100%"
      >
        <GlobalTableColumn prop="id" label="ID" :width="80" />
        <GlobalTableColumn prop="name" label="姓名" :sortable="true" />
        <GlobalTableColumn prop="age" label="年龄" :width="100" align="center" />
        <GlobalTableColumn prop="email" label="邮箱" :ellipsis="true" />
        <GlobalTableColumn prop="department" label="部门" />
        <GlobalTableColumn label="状态" :width="120" align="center">
          <template #default="{ row }">
            <span
              :class="[
                'px-2 py-1 rounded-full text-xs',
                row.status === '正常' ? 'bg-green-100 text-green-800' :
                row.status === '禁用' ? 'bg-red-100 text-red-800' :
                'bg-yellow-100 text-yellow-800'
              ]"
            >
              {{ row.status || '正常' }}
            </span>
          </template>
        </GlobalTableColumn>
        <GlobalTableColumn label="操作" :width="180" align="center">
          <template #default="{ row }">
            <GlobalButton size="sm" @click="handleEdit(row)" class="mr-2">
              编辑
            </GlobalButton>
            <GlobalButton size="sm" variant="outline" @click="handleDelete(row)">
              删除
            </GlobalButton>
          </template>
        </GlobalTableColumn>
      </GlobalTable>
    </GlobalCodeDemo>

    <!-- 基础表格 -->
    <GlobalCodeDemo
      title="基础表格"
      description="最基本的表格展示，支持排序和加载状态"
      :code="basicTableCode"
    >
      <GlobalTable
        :columns="basicColumns"
        :data="basicData"
        :loading="basicLoading"
      />

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
    </GlobalCodeDemo>

    <!-- 选择功能 -->
    <GlobalCodeDemo
      title="行选择功能"
      description="支持多选和单选，可以禁用特定行"
      :code="selectionTableCode"
    >
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
    </GlobalCodeDemo>

    <!-- 表格配置 -->
    <GlobalCodeDemo
      title="表格配置"
      description="不同尺寸、边框、斑马纹等样式配置"
      :code="configTableCode"
    >
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

      <GlobalTable
        :columns="configColumns"
        :data="configData"
        v-bind="tableConfig"
      />
    </GlobalCodeDemo>

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
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, h } from 'vue'

defineOptions({
  name: 'TableTest'
})

// Slot 方式表格数据
const slotTableData = ref([
  { id: 1, name: '张三', age: 28, email: 'zhangsan@example.com', status: '正常' },
  { id: 2, name: '李四', age: 32, email: 'lisi@example.com', status: '禁用' },
  { id: 3, name: '王五', age: 25, email: 'wangwu@example.com', status: '待审' },
  { id: 4, name: '赵六', age: 29, email: 'zhaoliu@example.com', status: '正常' },
])

// Element UI 风格表格数据
const elementTableData = ref([
  { id: 1, name: '张三', age: 28, email: 'zhangsan@example.com', department: '技术部' },
  { id: 2, name: '李四', age: 32, email: 'lisi@example.com', department: '产品部' },
  { id: 3, name: '王五', age: 25, email: 'wangwu@example.com', department: '设计部' },
  { id: 4, name: '赵六', age: 29, email: 'zhaoliu@example.com', department: '运营部' },
])

// Element UI 风格表格列定义
const elementStyleColumns = [
  { key: 'id', title: 'ID', width: 80 },
  { key: 'name', title: '姓名' },
  { key: 'age', title: '年龄', width: 100, align: 'center' },
  { key: 'email', title: '邮箱' },
  { key: 'department', title: '部门' },
  {
    key: 'action',
    title: '操作',
    width: 180,
    align: 'center',
    render: (_: any, record: any) => {
      return h('div', {
        class: 'space-x-2'
      }, [
        h('button', {
          class: 'px-3 py-1 text-sm bg-blue-500 text-white rounded hover:bg-blue-600',
          onClick: () => handleEdit(record)
        }, '编辑'),
        h('button', {
          class: 'px-3 py-1 text-sm bg-gray-500 text-white rounded hover:bg-gray-600',
          onClick: () => handleDelete(record)
        }, '删除')
      ])
    }
  }
]

// Slot 方式表格列定义 (使用 render 函数)
const slotTableColumns = [
  { key: 'id', title: 'ID', width: 80, sortable: true },
  { key: 'name', title: '姓名', sortable: true },
  { key: 'age', title: '年龄', width: 100, align: 'center' },
  { key: 'email', title: '邮箱', ellipsis: true },
  {
    key: 'status',
    title: '状态',
    width: 120,
    align: 'center',
    render: (value: string) => {
      const colors = {
        '正常': 'bg-green-100 text-green-800',
        '禁用': 'bg-red-100 text-red-800',
        '待审': 'bg-yellow-100 text-yellow-800',
      }
      // 使用Vue的h函数创建VNode
      return h('span', {
        class: `px-2 py-1 rounded-full text-xs ${colors[value as keyof typeof colors] || 'bg-gray-100 text-gray-800'}`
      }, value)
    }
  },
  {
    key: 'action',
    title: '操作',
    width: 150,
    align: 'center',
    render: (_: any, record: any) => {
      // 使用Vue的h函数创建VNode
      return h('div', {
        class: 'space-x-2'
      }, [
        h('button', {
          class: 'px-2 py-1 text-xs text-blue-600 border border-blue-300 rounded hover:bg-blue-50',
          onClick: () => handleEdit(record)
        }, '编辑'),
        h('button', {
          class: 'px-2 py-1 text-xs text-red-600 border border-red-300 rounded hover:bg-red-50',
          onClick: () => handleDelete(record)
        }, '删除')
      ])
    }
  }
]

// 基础表格数据
const basicLoading = ref(false)
const basicColumns = [
  { key: 'id', title: 'ID', width: 80, sortable: true },
  { key: 'name', title: '姓名', sortable: true },
  { key: 'age', title: '年龄', width: 100, align: 'center', sortable: true },
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
const selectionType = ref('checkbox')

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
  size: 'medium',
  bordered: true,
  striped: false,
  hoverable: true,
})

const configColumns = [
  { key: 'name', title: '名称' },
  { key: 'type', title: '类型' },
  { key: 'size', title: '大小', align: 'right' },
  { key: 'date', title: '日期' },
]

const configData = ref([
  { name: 'document.pdf', type: 'PDF文档', size: '2.3MB', date: '2024-01-15' },
  { name: 'image.jpg', type: '图片', size: '1.8MB', date: '2024-01-14' },
  { name: 'video.mp4', type: '视频', size: '15.6MB', date: '2024-01-13' },
  { name: 'archive.zip', type: '压缩包', size: '8.9MB', date: '2024-01-12' },
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

const handleSelectionChange = (keys: any[], rows: any[]) => {
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
  { name: 'rowSelection', description: '行选择配置', type: 'object', default: 'undefined' },
  { name: 'emptyText', description: '空数据文案', type: 'string', default: "'暂无数据'" },
]

// API 表格列定义
const propsApiColumns = [
  { key: 'name', title: '参数', width: 150 },
  { key: 'description', title: '说明' },
  { key: 'type', title: '类型', width: 200 },
  { key: 'default', title: '默认值', width: 120 },
]

// 代码示例字符串 - 简化版本
const slotTableCode = `// 自定义渲染表格示例 (类似 Element UI 风格)
const columns = [
  { key: 'id', title: 'ID', width: 80, sortable: true },
  { key: 'name', title: '姓名', sortable: true },
  { key: 'age', title: '年龄', width: 100, align: 'center' },
  {
    key: 'status',
    title: '状态',
    width: 120,
    align: 'center',
    render: (value) => {
      const colors = {
        '正常': 'bg-green-100 text-green-800',
        '禁用': 'bg-red-100 text-red-800'
      }
      // 返回 VNode 对象而不是 HTML 字符串
      return {
        type: 'span',
        props: {
          class: \`px-2 py-1 rounded-full text-xs \${colors[value]}\`
        },
        children: value
      }
    }
  },
  {
    key: 'action',
    title: '操作',
    width: 150,
    align: 'center',
    render: (_, record) => {
      // 返回 VNode 对象，支持事件处理
      return {
        type: 'div',
        props: { class: 'space-x-2' },
        children: [
          {
            type: 'button',
            props: {
              class: 'px-2 py-1 text-xs text-blue-600 border border-blue-300 rounded',
              onClick: () => handleEdit(record)
            },
            children: '编辑'
          }
        ]
      }
    }
  }
]

<GlobalTable :data="tableData" :columns="columns" />`

const basicTableCode = `// 基础表格示例
<GlobalTable
  :columns="basicColumns"
  :data="basicData"
  :loading="basicLoading"
/>

const basicColumns = [
  { key: 'id', title: 'ID', width: 80, sortable: true },
  { key: 'name', title: '姓名', sortable: true },
  { key: 'age', title: '年龄', width: 100, align: 'center' },
  { key: 'email', title: '邮箱', ellipsis: true }
]`

const selectionTableCode = `// 行选择表格示例
<GlobalTable
  :columns="selectionColumns"
  :data="selectionData"
  :row-selection="{
    type: 'checkbox',
    selectedRowKeys: selectedRowKeys,
    onChange: handleSelectionChange
  }"
/>`

const configTableCode = `// 表格配置示例
<GlobalTable
  :columns="configColumns"
  :data="configData"
  :size="tableConfig.size"
  :bordered="tableConfig.bordered"
  :striped="tableConfig.striped"
  :hoverable="tableConfig.hoverable"
/>`

const elementStyleTableCode = `// Element UI 风格表格示例 (子组件方式)
<GlobalTable
  :data="elementTableData"
  :loading="false"
  :bordered="true"
  size="medium"
  style="width: 100%"
>
  <GlobalTableColumn prop="id" label="ID" :width="80" />
  <GlobalTableColumn prop="name" label="姓名" :sortable="true" />
  <GlobalTableColumn prop="age" label="年龄" :width="100" align="center" />
  <GlobalTableColumn prop="email" label="邮箱" :ellipsis="true" />
  <GlobalTableColumn prop="department" label="部门" />

  <!-- 自定义状态列 -->
  <GlobalTableColumn label="状态" :width="120" align="center">
    <template #default="{ row }">
      <span
        :class="[
          'px-2 py-1 rounded-full text-xs',
          row.status === '正常' ? 'bg-green-100 text-green-800' :
          row.status === '禁用' ? 'bg-red-100 text-red-800' :
          'bg-yellow-100 text-yellow-800'
        ]"
      >
        \{{ row.status || '正常' }}
      </span>
    </template>
  </GlobalTableColumn>

  <!-- 操作列 -->
  <GlobalTableColumn label="操作" :width="180" align="center">
    <template #default="{ row }">
      <GlobalButton size="sm" @click="handleEdit(row)" class="mr-2">
        编辑
      </GlobalButton>
      <GlobalButton size="sm" variant="outline" @click="handleDelete(row)">
        删除
      </GlobalButton>
    </template>
  </GlobalTableColumn>
</GlobalTable>

// 数据定义
const elementTableData = ref([
  { id: 1, name: '张三', age: 28, email: 'zhangsan@example.com', department: '技术部' },
  { id: 2, name: '李四', age: 32, email: 'lisi@example.com', department: '产品部' },
  { id: 3, name: '王五', age: 25, email: 'wangwu@example.com', department: '设计部' },
  { id: 4, name: '赵六', age: 29, email: 'zhaoliu@example.com', department: '运营部' },
])`

const simpleTestCode = `// 简单功能测试示例
<GlobalTable :data="simpleTestData" :bordered="true">
  <GlobalTableColumn prop="name" label="姓名" />
  <GlobalTableColumn prop="age" label="年龄" :width="100" />
  <GlobalTableColumn label="操作" :width="100">
    <template #default="{ row }">
      <button @click="alert(\`点击了\${row.name}\`)" class="text-blue-600">点击</button>
    </template>
  </GlobalTableColumn>
</GlobalTable>`

const simpleTestData = ref([
  { name: '张三', age: 28 },
  { name: '李四', age: 32 },
  { name: '王五', age: 25 },
  { name: '赵六', age: 29 },
])
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