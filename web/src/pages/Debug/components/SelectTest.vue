<template>
  <div class="space-y-8">
    <!-- 基础选择器 -->
    <div class="demo-section">
      <h4 class="demo-title">基础选择器</h4>
      <p class="demo-description">最基本的选择器使用</p>

      <div class="demo-container">
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
      </div>
    </div>

    <!-- 多选选择器 -->
    <div class="demo-section">
      <h4 class="demo-title">多选选择器</h4>
      <p class="demo-description">支持选择多个选项</p>

      <div class="demo-container">
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
      </div>
    </div>

    <!-- 可搜索选择器 -->
    <div class="demo-section">
      <h4 class="demo-title">可搜索选择器</h4>
      <p class="demo-description">支持输入关键词过滤选项</p>

      <div class="demo-container">
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
      </div>
    </div>

    <!-- 不同尺寸 -->
    <div class="demo-section">
      <h4 class="demo-title">不同尺寸</h4>
      <p class="demo-description">提供 small、medium、large 三种尺寸</p>

      <div class="demo-container">
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
      </div>
    </div>

    <!-- 不同样式 -->
    <div class="demo-section">
      <h4 class="demo-title">不同样式</h4>
      <p class="demo-description">提供不同的视觉样式</p>

      <div class="demo-container">
        <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Default</label>
            <GlobalSelect
              v-model="variantValue.default"
              :options="basicOptions"
              variant="default"
              placeholder="默认样式"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Bordered</label>
            <GlobalSelect
              v-model="variantValue.bordered"
              :options="basicOptions"
              variant="bordered"
              placeholder="边框加粗样式"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">Filled</label>
            <GlobalSelect
              v-model="variantValue.filled"
              :options="basicOptions"
              variant="filled"
              placeholder="填充背景样式"
            />
          </div>
        </div>
      </div>
    </div>

    <!-- 带图标选择器 -->
    <div class="demo-section">
      <h4 class="demo-title">带图标选择器</h4>
      <p class="demo-description">选项可以带有图标</p>

      <div class="demo-container">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">图标选项</label>
            <GlobalSelect
              v-model="iconValue"
              :options="iconOptions"
              placeholder="选择带图标的选项"
              clearable
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">多选图标</label>
            <GlobalSelect
              v-model="multipleIconValue"
              :options="iconOptions"
              placeholder="多选图标选项"
              multiple
              clearable
            />
          </div>
        </div>
      </div>
    </div>

    <!-- 状态演示 -->
    <div class="demo-section">
      <h4 class="demo-title">状态演示</h4>
      <p class="demo-description">不同状态下的选择器</p>

      <div class="demo-container">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">禁用状态</label>
            <GlobalSelect
              v-model="disabledValue"
              :options="basicOptions"
              placeholder="禁用的选择器"
              disabled
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">加载状态</label>
            <GlobalSelect
              v-model="loadingValue"
              :options="basicOptions"
              placeholder="加载中..."
              loading
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">错误状态</label>
            <GlobalSelect
              v-model="errorValue"
              :options="basicOptions"
              placeholder="错误状态"
              error
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">成功状态</label>
            <GlobalSelect
              v-model="successValue"
              :options="basicOptions"
              placeholder="成功状态"
              success
            />
          </div>
        </div>
      </div>
    </div>

    <!-- 使用说明 -->
    <div class="bg-gradient-to-r from-blue-50 to-indigo-50 p-6 rounded-lg border border-blue-200">
      <h4 class="text-lg font-semibold text-blue-900 mb-4 flex items-center gap-2">
        <GlobalIcon name="book-open" size="sm" color="text-blue-600" />
        使用说明
      </h4>

      <div class="space-y-4 text-sm text-blue-800">
        <div>
          <h5 class="font-semibold mb-2">基本用法</h5>
          <div class="bg-white p-3 rounded border border-blue-200 font-mono text-xs">
            <div>&lt;GlobalSelect</div>
            <div class="ml-2">v-model="value"</div>
            <div class="ml-2">:options="options"</div>
            <div class="ml-2">placeholder="请选择"</div>
            <div class="ml-2">@change="handleChange"</div>
            <div>/&gt;</div>
          </div>
        </div>

        <div>
          <h5 class="font-semibold mb-2">选项数据格式</h5>
          <div class="bg-white p-3 rounded border border-blue-200 font-mono text-xs">
            <div>const options = [</div>
            <div class="ml-2">{'{'} label: '选项1', value: 'option1' {'}'},</div>
            <div class="ml-2">{'{'} label: '选项2', value: 'option2', disabled: true {'}'},</div>
            <div class="ml-2">{'{'} label: '图标选项', value: 'option3', icon: 'star' {'}'}</div>
            <div>]</div>
          </div>
        </div>
      </div>
    </div>

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

    <!-- Events API 说明 -->
    <div class="demo-section">
      <h4 class="demo-title">Events API 说明</h4>
      <p class="demo-description">选择器组件支持的事件回调</p>

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
import { ref, h } from 'vue'
import type { SelectOption } from '../../../components/Select/types'

defineOptions({
  name: 'SelectTest'
})

// 基础数据
const basicValue = ref('')
const clearableValue = ref('option2')

const basicOptions: SelectOption[] = [
  { label: '选项一', value: 'option1' },
  { label: '选项二', value: 'option2' },
  { label: '选项三', value: 'option3' },
  { label: '禁用选项', value: 'disabled', disabled: true },
  { label: '选项五', value: 'option5' },
]

// 多选数据
const multipleValue = ref<(string | number)[]>(['option1', 'option3'])
const limitedTagsValue = ref<(string | number)[]>(['option1', 'option2', 'option3', 'option5'])

// 搜索数据
const searchableValue = ref('')
const customFilterValue = ref('')

const searchableOptions: SelectOption[] = [
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
const sizeValue = ref({
  small: '',
  medium: 'option2',
  large: ''
})

// 样式数据
const variantValue = ref({
  default: '',
  bordered: 'option1',
  filled: ''
})

// 图标数据
const iconValue = ref('')
const multipleIconValue = ref<(string | number)[]>([])

const iconOptions: SelectOption[] = [
  { label: '首页', value: 'home', icon: 'home' },
  { label: '用户', value: 'user', icon: 'user' },
  { label: '设置', value: 'settings', icon: 'cog-6-tooth' },
  { label: '收藏', value: 'favorite', icon: 'heart' },
  { label: '消息', value: 'message', icon: 'chat-bubble-left' },
  { label: '搜索', value: 'search', icon: 'magnifying-glass' },
]

// 状态数据
const disabledValue = ref('option1')
const loadingValue = ref('')
const errorValue = ref('')
const successValue = ref('option2')

// 事件处理
const handleBasicChange = (value: string | number | (string | number)[], option: SelectOption | SelectOption[] | null) => {
  console.log('选择变化:', value, option)
}

// 自定义过滤方法
const customFilter = (query: string, option: SelectOption) => {
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

  const pinyin = pinyinMap[option.value as string] || ''

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
  { name: 'variant', description: '样式变体', type: "'default' | 'bordered' | 'filled'", default: "'default'" },
  { name: 'filterable', description: '是否可搜索', type: 'boolean', default: 'false' },
  { name: 'filterMethod', description: '自定义过滤方法', type: 'function', default: 'undefined' },
  { name: 'maxTagCount', description: '最大标签显示数量', type: 'number', default: '3' },
  { name: 'showArrow', description: '是否显示箭头', type: 'boolean', default: 'true' },
  { name: 'loading', description: '是否加载中', type: 'boolean', default: 'false' },
  { name: 'maxHeight', description: '下拉面板最大高度', type: 'number', default: '240' },
  { name: 'placement', description: '下拉面板位置', type: "'bottom' | 'top' | 'auto'", default: "'auto'" },
  { name: 'error', description: '错误状态', type: 'boolean', default: 'false' },
  { name: 'success', description: '成功状态', type: 'boolean', default: 'false' },
]

const eventsApi = [
  { name: 'update:modelValue', description: '值变化时触发', params: '(value)' },
  { name: 'change', description: '选择变化时触发', params: '(value, option)' },
  { name: 'clear', description: '清除时触发', params: '()' },
  { name: 'focus', description: '获得焦点时触发', params: '(event)' },
  { name: 'blur', description: '失去焦点时触发', params: '(event)' },
  { name: 'search', description: '搜索时触发', params: '(query)' },
  { name: 'visible-change', description: '下拉面板显示状态变化', params: '(visible)' },
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
    width: 150,
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
