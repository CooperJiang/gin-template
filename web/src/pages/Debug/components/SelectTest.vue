<template>
  <div class="space-y-8">
    <!-- 基础选择器 -->
    <GlobalCodeDemo
      title="基础选择器"
      description="最基本的选择器使用，支持单选和可清除功能"
      :code="basicSelectCode"
    >
      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">单选</label>
          <GlobalSelect
            v-model="basicValue"
            :options="basicOptions"
            placeholder="请选择一个选项"
            @change="handleBasicChange"
          />
          <div class="mt-2 text-xs text-gray-500">
            当前值: {{ basicValue || '未选择' }}
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">可清除</label>
          <GlobalSelect
            v-model="clearableValue"
            :options="basicOptions"
            placeholder="可清除的选择器"
            clearable
          />
          <div class="mt-2 text-xs text-gray-500">
            当前值: {{ clearableValue || '未选择' }}
          </div>
        </div>
      </div>
    </GlobalCodeDemo>

    <!-- 多选选择器 -->
    <GlobalCodeDemo
      title="多选选择器"
      description="支持选择多个选项，可以限制标签显示数量"
      :code="multipleSelectCode"
    >
      <div class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">多选模式</label>
          <GlobalSelect
            v-model="multipleValue"
            :options="basicOptions"
            placeholder="请选择多个选项"
            multiple
            clearable
          />
          <div class="mt-2 text-xs text-gray-500">
            当前值: {{ multipleValue.length ? multipleValue.join(', ') : '未选择' }}
          </div>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">限制标签显示数量</label>
          <GlobalSelect
            v-model="limitedTagsValue"
            :options="basicOptions"
            placeholder="最多显示2个标签"
            multiple
            :max-tag-count="2"
            clearable
          />
          <div class="mt-2 text-xs text-gray-500">
            已选择: {{ limitedTagsValue.length }} 项
          </div>
        </div>
      </div>
    </GlobalCodeDemo>

    <!-- 可搜索选择器 -->
    <GlobalCodeDemo
      title="可搜索选择器"
      description="支持输入关键词过滤选项，可自定义过滤方法"
      :code="searchableSelectCode"
    >
      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">可搜索</label>
          <GlobalSelect
            v-model="searchableValue"
            :options="searchableOptions"
            placeholder="输入关键词搜索"
            filterable
            clearable
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">自定义过滤方法</label>
          <GlobalSelect
            v-model="customFilterValue"
            :options="searchableOptions"
            placeholder="按拼音首字母搜索"
            filterable
            :filter-method="customFilter"
            clearable
          />
        </div>
      </div>
    </GlobalCodeDemo>

    <!-- 不同尺寸 -->
    <GlobalCodeDemo
      title="不同尺寸"
      description="提供 small、medium、large 三种尺寸"
      :code="sizeSelectCode"
    >
      <div class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">Small</label>
          <GlobalSelect
            v-model="sizeValue.small"
            :options="basicOptions"
            size="small"
            placeholder="小尺寸选择器"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">Medium (默认)</label>
          <GlobalSelect
            v-model="sizeValue.medium"
            :options="basicOptions"
            size="medium"
            placeholder="中等尺寸选择器"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">Large</label>
          <GlobalSelect
            v-model="sizeValue.large"
            :options="basicOptions"
            size="large"
            placeholder="大尺寸选择器"
          />
        </div>
      </div>
    </GlobalCodeDemo>

    <!-- API 说明 -->
    <div class="demo-section">
      <h4 class="demo-title">Props API 说明</h4>
      <p class="demo-description">选择器组件支持的属性配置</p>

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
import { ref, reactive } from 'vue'

defineOptions({
  name: 'SelectTest'
})

// 基础数据
const basicValue = ref('')
const clearableValue = ref('option2')

const basicOptions = [
  { label: '选项一', value: 'option1' },
  { label: '选项二', value: 'option2' },
  { label: '选项三', value: 'option3' },
  { label: '禁用选项', value: 'disabled', disabled: true },
  { label: '选项五', value: 'option5' },
]

// 多选数据
const multipleValue = ref(['option1', 'option3'])
const limitedTagsValue = ref(['option1', 'option2', 'option3', 'option5'])

// 搜索数据
const searchableValue = ref('')
const customFilterValue = ref('')

const searchableOptions = [
  { label: '北京 Beijing', value: 'beijing' },
  { label: '上海 Shanghai', value: 'shanghai' },
  { label: '广州 Guangzhou', value: 'guangzhou' },
  { label: '深圳 Shenzhen', value: 'shenzhen' },
  { label: '杭州 Hangzhou', value: 'hangzhou' },
  { label: '南京 Nanjing', value: 'nanjing' },
  { label: '武汉 Wuhan', value: 'wuhan' },
  { label: '成都 Chengdu', value: 'chengdu' },
]

// 尺寸数据
const sizeValue = reactive({
  small: '',
  medium: 'option2',
  large: ''
})

// 事件处理
const handleBasicChange = (value: any, option: any) => {
  console.log('选择变化:', value, option)
}

// 自定义过滤方法
const customFilter = (query: string, option: any) => {
  const searchText = option.label.toLowerCase()
  const queryLower = query.toLowerCase()

  // 按拼音首字母搜索的简单实现
  const pinyinMap: Record<string, string> = {
    'beijing': 'bj',
    'shanghai': 'sh',
    'guangzhou': 'gz',
    'shenzhen': 'sz',
    'hangzhou': 'hz',
    'nanjing': 'nj',
    'wuhan': 'wh',
    'chengdu': 'cd'
  }

  const pinyin = pinyinMap[option.value] || ''

  return searchText.includes(queryLower) || pinyin.includes(queryLower)
}

// API 文档数据
const propsApi = [
  { name: 'modelValue', description: '绑定值', type: 'string | number | (string | number)[]', default: 'undefined' },
  { name: 'options', description: '选项数据', type: 'SelectOption[]', default: '[]' },
  { name: 'placeholder', description: '占位符文本', type: 'string', default: "'请选择'" },
  { name: 'disabled', description: '是否禁用', type: 'boolean', default: 'false' },
  { name: 'clearable', description: '是否可清除', type: 'boolean', default: 'false' },
  { name: 'multiple', description: '是否多选', type: 'boolean', default: 'false' },
  { name: 'size', description: '尺寸大小', type: "'small' | 'medium' | 'large'", default: "'medium'" },
  { name: 'filterable', description: '是否可搜索', type: 'boolean', default: 'false' },
  { name: 'filterMethod', description: '自定义过滤方法', type: 'function', default: 'undefined' },
  { name: 'maxTagCount', description: '最大标签显示数量', type: 'number', default: '3' },
]

// API 表格列定义
const propsApiColumns = [
  { key: 'name', title: '参数', width: 150 },
  { key: 'description', title: '说明' },
  { key: 'type', title: '类型', width: 200 },
  { key: 'default', title: '默认值', width: 120 },
]

// 代码示例字符串 - 简化版本，避免Vue模板语法冲突
const basicSelectCode = `// 基础选择器示例
<GlobalSelect
  v-model="basicValue"
  :options="basicOptions"
  placeholder="请选择一个选项"
  clearable
/>

const basicOptions = [
  { label: '选项一', value: 'option1' },
  { label: '选项二', value: 'option2' },
  { label: '选项三', value: 'option3' }
]`

const multipleSelectCode = `// 多选选择器示例
<GlobalSelect
  v-model="multipleValue"
  :options="basicOptions"
  placeholder="请选择多个选项"
  multiple
  clearable
  :max-tag-count="2"
/>`

const searchableSelectCode = `// 可搜索选择器示例
<GlobalSelect
  v-model="searchableValue"
  :options="searchableOptions"
  placeholder="输入关键词搜索"
  filterable
  clearable
  :filter-method="customFilter"
/>`

const sizeSelectCode = `// 不同尺寸示例
<GlobalSelect v-model="value" :options="options" size="small" />
<GlobalSelect v-model="value" :options="options" size="medium" />
<GlobalSelect v-model="value" :options="options" size="large" />`
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
