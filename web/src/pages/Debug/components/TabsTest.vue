<template>
  <div class="space-y-8">
    <!-- 页面标题 -->
    <div class="border-b border-gray-200 pb-4">
      <h2 class="text-2xl font-bold text-gray-900">Tabs 标签页组件测试</h2>
      <p class="mt-2 text-gray-600">测试标签页组件的各种功能和样式</p>
    </div>

    <!-- 基础用法 -->
    <GlobalCard title="基础用法" class="w-full">
      <div class="space-y-4">
        <GlobalTabs
          v-model="basicActiveKey"
          :items="basicTabs"
          @change="handleTabChange"
        >
          <template #tab1>
            <div class="p-4">
              <h3 class="text-lg font-medium text-gray-900 mb-2">首页内容</h3>
              <p class="text-gray-600">这是首页的内容区域，展示网站的主要信息和功能入口。</p>
            </div>
          </template>

          <template #tab2>
            <div class="p-4">
              <h3 class="text-lg font-medium text-gray-900 mb-2">产品列表</h3>
              <div class="grid grid-cols-2 gap-4">
                <div class="border rounded-lg p-3">
                  <h4 class="font-medium">产品 A</h4>
                  <p class="text-sm text-gray-500">产品描述信息</p>
                </div>
                <div class="border rounded-lg p-3">
                  <h4 class="font-medium">产品 B</h4>
                  <p class="text-sm text-gray-500">产品描述信息</p>
                </div>
              </div>
            </div>
          </template>

          <template #tab3>
            <div class="p-4">
              <h3 class="text-lg font-medium text-gray-900 mb-2">关于我们</h3>
              <p class="text-gray-600">公司简介和联系方式等信息。</p>
            </div>
          </template>
        </GlobalTabs>
      </div>
    </GlobalCard>

    <!-- 不同类型 -->
    <GlobalCard title="不同类型" class="w-full">
      <div class="space-y-6">
        <!-- Line 类型 -->
        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-3">Line 类型 (默认)</h4>
          <GlobalTabs v-model="typeActiveKey" type="line" :items="typeTabs">
            <template #type1><div class="p-4">Line 类型的内容</div></template>
            <template #type2><div class="p-4">带有下划线指示器</div></template>
            <template #type3><div class="p-4">简洁的设计风格</div></template>
          </GlobalTabs>
        </div>

        <!-- Card 类型 -->
        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-3">Card 类型</h4>
          <GlobalTabs v-model="typeActiveKey" type="card" :items="typeTabs">
            <template #type1><div class="p-4">Card 类型的内容</div></template>
            <template #type2><div class="p-4">类似卡片的设计</div></template>
            <template #type3><div class="p-4">有明显的边框</div></template>
          </GlobalTabs>
        </div>

        <!-- Button 类型 -->
        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-3">Button 类型</h4>
          <GlobalTabs v-model="typeActiveKey" type="button" :items="typeTabs">
            <template #type1><div class="p-4">Button 类型的内容</div></template>
            <template #type2><div class="p-4">按钮样式的标签页</div></template>
            <template #type3><div class="p-4">现代化的外观</div></template>
          </GlobalTabs>
        </div>
      </div>
    </GlobalCard>

    <!-- 不同尺寸 -->
    <GlobalCard title="不同尺寸" class="w-full">
      <div class="space-y-6">
        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-3">Small 尺寸</h4>
          <GlobalTabs v-model="sizeActiveKey" size="small" :items="sizeTabs">
            <template #size1><div class="p-3">小尺寸的内容区域</div></template>
            <template #size2><div class="p-3">更紧凑的设计</div></template>
          </GlobalTabs>
        </div>

        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-3">Medium 尺寸 (默认)</h4>
          <GlobalTabs v-model="sizeActiveKey" size="medium" :items="sizeTabs">
            <template #size1><div class="p-4">中等尺寸的内容区域</div></template>
            <template #size2><div class="p-4">标准的内边距</div></template>
          </GlobalTabs>
        </div>

        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-3">Large 尺寸</h4>
          <GlobalTabs v-model="sizeActiveKey" size="large" :items="sizeTabs">
            <template #size1><div class="p-6">大尺寸的内容区域</div></template>
            <template #size2><div class="p-6">更宽松的内边距</div></template>
          </GlobalTabs>
        </div>
      </div>
    </GlobalCard>

    <!-- 带图标和可关闭 -->
    <GlobalCard title="带图标和可关闭" class="w-full">
      <div class="space-y-4">
        <GlobalTabs
          v-model="iconActiveKey"
          :items="iconTabs"
          addable
          @close="handleTabClose"
          @add="handleTabAdd"
        >
          <template #icon1>
            <div class="p-4">
              <div class="flex items-center space-x-3">
                <GlobalIcon name="home" size="lg" class="text-blue-500" />
                <div>
                  <h3 class="text-lg font-medium">主页</h3>
                  <p class="text-gray-600">欢迎来到主页面</p>
                </div>
              </div>
            </div>
          </template>

          <template #icon2>
            <div class="p-4">
              <div class="flex items-center space-x-3">
                <GlobalIcon name="user-group" size="lg" class="text-green-500" />
                <div>
                  <h3 class="text-lg font-medium">用户管理</h3>
                  <p class="text-gray-600">管理系统用户</p>
                </div>
              </div>
            </div>
          </template>

          <template #icon3>
            <div class="p-4">
              <div class="flex items-center space-x-3">
                <GlobalIcon name="cog-6-tooth" size="lg" class="text-orange-500" />
                <div>
                  <h3 class="text-lg font-medium">系统设置</h3>
                  <p class="text-gray-600">配置系统参数</p>
                </div>
              </div>
            </div>
          </template>
        </GlobalTabs>
      </div>
    </GlobalCard>

    <!-- 禁用状态 -->
    <GlobalCard title="禁用状态" class="w-full">
      <div class="space-y-4">
        <GlobalTabs v-model="disabledActiveKey" :items="disabledTabs">
          <template #disabled1>
            <div class="p-4">正常标签页内容</div>
          </template>
          <template #disabled2>
            <div class="p-4">这个标签页被禁用了</div>
          </template>
          <template #disabled3>
            <div class="p-4">另一个正常的标签页</div>
          </template>
        </GlobalTabs>
      </div>
    </GlobalCard>

    <!-- 自定义标题插槽 -->
    <GlobalCard title="自定义标题插槽" class="w-full">
      <div class="space-y-4">
        <GlobalTabs v-model="customActiveKey" :items="customTabs">
          <template #title="{ item, active }">
            <div class="flex items-center space-x-2">
              <GlobalIcon :name="item.icon" size="sm" />
              <span :class="{ 'font-bold': active }">{{ item.title }}</span>
              <span v-if="item.key === 'custom2'" class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-red-100 text-red-800">
                New
              </span>
            </div>
          </template>

          <template #custom1>
            <div class="p-4">
              <h3 class="text-lg font-medium mb-2">自定义标题样式</h3>
              <p class="text-gray-600">可以完全自定义标签页标题的显示方式</p>
            </div>
          </template>

          <template #custom2>
            <div class="p-4">
              <h3 class="text-lg font-medium mb-2">带徽章的标题</h3>
              <p class="text-gray-600">标题可以包含徽章、图标等元素</p>
            </div>
          </template>

          <template #custom3>
            <div class="p-4">
              <h3 class="text-lg font-medium mb-2">动态样式</h3>
              <p class="text-gray-600">根据激活状态应用不同的样式</p>
            </div>
          </template>
        </GlobalTabs>
      </div>
    </GlobalCard>

    <!-- 额外操作区域 -->
    <GlobalCard title="额外操作区域" class="w-full">
      <div class="space-y-4">
        <GlobalTabs v-model="extraActiveKey" :items="extraTabs">
          <template #extra>
            <div class="flex items-center space-x-2">
              <GlobalButton size="small" variant="ghost" icon="cog-6-tooth">设置</GlobalButton>
              <GlobalButton size="small" variant="ghost" icon="question-mark-circle">帮助</GlobalButton>
            </div>
          </template>

          <template #extra1>
            <div class="p-4">
              <h3 class="text-lg font-medium mb-2">内容区域 1</h3>
              <p class="text-gray-600">右侧有额外的操作按钮</p>
            </div>
          </template>

          <template #extra2>
            <div class="p-4">
              <h3 class="text-lg font-medium mb-2">内容区域 2</h3>
              <p class="text-gray-600">可以放置设置、帮助等操作</p>
            </div>
          </template>
        </GlobalTabs>
      </div>
    </GlobalCard>

    <!-- 事件测试 -->
    <GlobalCard title="事件测试" class="w-full">
      <div class="space-y-4">
        <div class="bg-gray-50 p-4 rounded-lg">
          <h4 class="text-sm font-medium text-gray-700 mb-2">事件日志:</h4>
          <div class="space-y-1 max-h-32 overflow-y-auto">
            <div
              v-for="(log, index) in eventLogs"
              :key="index"
              class="text-xs text-gray-600 font-mono"
            >
              [{{ log.time }}] {{ log.message }}
            </div>
          </div>
          <GlobalButton
            size="small"
            variant="outline"
            class="mt-2"
            @click="clearLogs"
          >
            清空日志
          </GlobalButton>
        </div>

        <GlobalTabs
          v-model="eventActiveKey"
          :items="eventTabs"
          @change="handleEventTabChange"
          @close="handleEventTabClose"
          @add="handleEventTabAdd"
        >
          <template #event1>
            <div class="p-4">监听标签页切换事件</div>
          </template>
          <template #event2>
            <div class="p-4">监听标签页关闭事件</div>
          </template>
          <template #event3>
            <div class="p-4">监听添加事件</div>
          </template>
        </GlobalTabs>
      </div>
    </GlobalCard>

    <!-- Props API 说明 -->
    <GlobalCard title="Props API 说明" class="w-full">
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
    </GlobalCard>

    <!-- Events API 说明 -->
    <GlobalCard title="Events API 说明" class="w-full">
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
    </GlobalCard>

    <!-- Slots API 说明 -->
    <GlobalCard title="Slots API 说明" class="w-full">
      <div class="overflow-x-auto">
        <table class="w-full border-collapse border border-gray-300 text-sm">
          <thead>
            <tr class="bg-gray-50">
              <th class="border border-gray-300 px-4 py-2 text-left">插槽名</th>
              <th class="border border-gray-300 px-4 py-2 text-left">说明</th>
              <th class="border border-gray-300 px-4 py-2 text-left">参数</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="slot in slotsApi" :key="slot.name" class="hover:bg-gray-50">
              <td class="border border-gray-300 px-4 py-2 font-mono text-blue-600">{{ slot.name }}</td>
              <td class="border border-gray-300 px-4 py-2">{{ slot.description }}</td>
              <td class="border border-gray-300 px-4 py-2 font-mono text-green-600">{{ slot.params }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </GlobalCard>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import type { TabItem } from '../../../components/Tabs/types'

defineOptions({
  name: 'TabsTest'
})

// 基础用法
const basicActiveKey = ref('tab1')
const basicTabs: TabItem[] = [
  { key: 'tab1', title: '首页', icon: 'home' },
  { key: 'tab2', title: '产品', icon: 'cube' },
  { key: 'tab3', title: '关于', icon: 'information-circle' }
]

// 不同类型
const typeActiveKey = ref('type1')
const typeTabs: TabItem[] = [
  { key: 'type1', title: '标签一' },
  { key: 'type2', title: '标签二' },
  { key: 'type3', title: '标签三' }
]

// 不同尺寸
const sizeActiveKey = ref('size1')
const sizeTabs: TabItem[] = [
  { key: 'size1', title: '选项 A' },
  { key: 'size2', title: '选项 B' }
]

// 带图标和可关闭
const iconActiveKey = ref('icon1')
const iconTabs = ref<TabItem[]>([
  { key: 'icon1', title: '主页', icon: 'home', closable: false },
  { key: 'icon2', title: '用户', icon: 'user-group', closable: true },
  { key: 'icon3', title: '设置', icon: 'cog-6-tooth', closable: true }
])

// 禁用状态
const disabledActiveKey = ref('disabled1')
const disabledTabs: TabItem[] = [
  { key: 'disabled1', title: '正常标签' },
  { key: 'disabled2', title: '禁用标签', disabled: true },
  { key: 'disabled3', title: '另一个标签' }
]

// 自定义标题
const customActiveKey = ref('custom1')
const customTabs: TabItem[] = [
  { key: 'custom1', title: '基础', icon: 'document-text' },
  { key: 'custom2', title: '高级', icon: 'star' },
  { key: 'custom3', title: '专业', icon: 'fire' }
]

// 额外操作
const extraActiveKey = ref('extra1')
const extraTabs: TabItem[] = [
  { key: 'extra1', title: '概览' },
  { key: 'extra2', title: '详情' }
]

// 事件测试
const eventActiveKey = ref('event1')
const eventTabs = ref<TabItem[]>([
  { key: 'event1', title: '事件 1', closable: true },
  { key: 'event2', title: '事件 2', closable: true },
  { key: 'event3', title: '事件 3', closable: true }
])

// 事件日志
const eventLogs = ref<Array<{ time: string; message: string }>>([])

// 添加日志
const addLog = (message: string) => {
  const time = new Date().toLocaleTimeString()
  eventLogs.value.unshift({ time, message })
  if (eventLogs.value.length > 10) {
    eventLogs.value.pop()
  }
}

// 清空日志
const clearLogs = () => {
  eventLogs.value = []
}

// 事件处理函数
const handleTabChange = (key: string, item: TabItem) => {
  console.log('Tab changed:', key, item)
}

const handleTabClose = (key: string, item: TabItem) => {
  console.log('Tab closed:', key, item)
  iconTabs.value = iconTabs.value.filter(tab => tab.key !== key)

  // 如果关闭的是当前激活的标签页，切换到第一个
  if (iconActiveKey.value === key && iconTabs.value.length > 0) {
    iconActiveKey.value = iconTabs.value[0].key
  }
}

const handleTabAdd = () => {
  console.log('Add tab clicked')
  const newKey = `icon${iconTabs.value.length + 1}`
  iconTabs.value.push({
    key: newKey,
    title: `新标签 ${iconTabs.value.length + 1}`,
    icon: 'plus',
    closable: true
  })
  iconActiveKey.value = newKey
}

// 事件测试处理函数
const handleEventTabChange = (key: string, item: TabItem) => {
  addLog(`标签页切换: ${item.title} (${key})`)
}

const handleEventTabClose = (key: string, item: TabItem) => {
  addLog(`标签页关闭: ${item.title} (${key})`)
  eventTabs.value = eventTabs.value.filter(tab => tab.key !== key)

  if (eventActiveKey.value === key && eventTabs.value.length > 0) {
    eventActiveKey.value = eventTabs.value[0].key
  }
}

const handleEventTabAdd = () => {
  addLog('点击添加标签页按钮')
  const newKey = `event${eventTabs.value.length + 1}`
  eventTabs.value.push({
    key: newKey,
    title: `新事件 ${eventTabs.value.length + 1}`,
    closable: true
  })
  eventActiveKey.value = newKey
}

// API 文档数据
const propsApi = [
  { name: 'modelValue', description: '当前激活的标签页key', type: 'string', default: 'undefined' },
  { name: 'items', description: '标签页数据数组', type: 'TabItem[]', default: '[]' },
  { name: 'type', description: '标签页类型', type: "'line' | 'card' | 'button'", default: "'line'" },
  { name: 'position', description: '标签页位置', type: "'top' | 'bottom' | 'left' | 'right'", default: "'top'" },
  { name: 'addable', description: '是否显示添加按钮', type: 'boolean', default: 'false' },
  { name: 'size', description: '标签页大小', type: "'small' | 'medium' | 'large'", default: "'medium'" },
  { name: 'sortable', description: '是否可以拖拽排序', type: 'boolean', default: 'false' },
  { name: 'lazy', description: '是否懒加载标签页内容', type: 'boolean', default: 'false' },
  { name: 'animated', description: '是否显示动画', type: 'boolean', default: 'true' }
]

const eventsApi = [
  { name: 'update:modelValue', description: '激活标签页改变事件', params: '(key: string)' },
  { name: 'change', description: '标签页切换事件', params: '(key: string, item: TabItem)' },
  { name: 'close', description: '标签页关闭事件', params: '(key: string, item: TabItem)' },
  { name: 'add', description: '添加标签页事件', params: '()' },
  { name: 'sort', description: '标签页排序事件', params: '(items: TabItem[])' }
]

const slotsApi = [
  { name: 'default', description: '标签页内容', params: '-' },
  { name: 'title', description: '标签页标题插槽', params: '{ item: TabItem, active: boolean }' },
  { name: 'addButton', description: '添加按钮插槽', params: '-' },
  { name: 'extra', description: '额外操作插槽', params: '-' },
  { name: '[key]', description: '动态插槽，key为标签页的key', params: '{ item: TabItem }' }
]
</script>
