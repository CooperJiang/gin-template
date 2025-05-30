<template>
  <div class="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
    <h3 class="text-lg font-semibold mb-4">Button 按钮组件测试</h3>
    <p class="text-gray-600 text-sm mb-6">测试按钮组件的各种样式、状态和功能</p>

    <!-- 基础按钮类型 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">按钮类型</h4>
      <div class="flex flex-wrap gap-3">
        <GlobalButton type="primary" @click="handleClick('Primary')">
          Primary
        </GlobalButton>
        <GlobalButton type="secondary" @click="handleClick('Secondary')">
          Secondary
        </GlobalButton>
        <GlobalButton type="success" @click="handleClick('Success')">
          Success
        </GlobalButton>
        <GlobalButton type="warning" @click="handleClick('Warning')">
          Warning
        </GlobalButton>
        <GlobalButton type="danger" @click="handleClick('Danger')">
          Danger
        </GlobalButton>
        <GlobalButton type="info" @click="handleClick('Info')">
          Info
        </GlobalButton>
      </div>
    </div>

    <!-- 按钮变体 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">按钮变体</h4>
      <div class="space-y-4">
        <!-- Solid -->
        <div>
          <h5 class="text-sm font-medium text-gray-700 mb-2">Solid (默认)</h5>
          <div class="flex flex-wrap gap-2">
            <GlobalButton type="primary" variant="solid">Primary</GlobalButton>
            <GlobalButton type="success" variant="solid">Success</GlobalButton>
            <GlobalButton type="danger" variant="solid">Danger</GlobalButton>
          </div>
        </div>

        <!-- Outline -->
        <div>
          <h5 class="text-sm font-medium text-gray-700 mb-2">Outline</h5>
          <div class="flex flex-wrap gap-2">
            <GlobalButton type="primary" variant="outline">Primary</GlobalButton>
            <GlobalButton type="success" variant="outline">Success</GlobalButton>
            <GlobalButton type="danger" variant="outline">Danger</GlobalButton>
          </div>
        </div>

        <!-- Ghost -->
        <div>
          <h5 class="text-sm font-medium text-gray-700 mb-2">Ghost</h5>
          <div class="flex flex-wrap gap-2">
            <GlobalButton type="primary" variant="ghost">Primary</GlobalButton>
            <GlobalButton type="success" variant="ghost">Success</GlobalButton>
            <GlobalButton type="danger" variant="ghost">Danger</GlobalButton>
          </div>
        </div>

        <!-- Link -->
        <div>
          <h5 class="text-sm font-medium text-gray-700 mb-2">Link</h5>
          <div class="flex flex-wrap gap-2">
            <GlobalButton type="primary" variant="link">Primary</GlobalButton>
            <GlobalButton type="success" variant="link">Success</GlobalButton>
            <GlobalButton type="danger" variant="link">Danger</GlobalButton>
          </div>
        </div>
      </div>
    </div>

    <!-- 按钮尺寸 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">按钮尺寸</h4>
      <div class="flex flex-wrap items-center gap-3">
        <GlobalButton type="primary" size="small">Small</GlobalButton>
        <GlobalButton type="primary" size="medium">Medium</GlobalButton>
        <GlobalButton type="primary" size="large">Large</GlobalButton>
      </div>
    </div>

    <!-- 按钮形状 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">按钮形状</h4>
      <div class="flex flex-wrap items-center gap-3">
        <GlobalButton type="primary" shape="default">Default</GlobalButton>
        <GlobalButton type="primary" shape="round">Round</GlobalButton>
        <GlobalButton type="primary" shape="circle">+</GlobalButton>
      </div>
    </div>

    <!-- 按钮状态 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">按钮状态</h4>
      <div class="space-y-4">
        <!-- 正常状态 -->
        <div>
          <h5 class="text-sm font-medium text-gray-700 mb-2">正常状态</h5>
          <div class="flex flex-wrap gap-2">
            <GlobalButton type="primary">正常按钮</GlobalButton>
            <GlobalButton type="success">正常按钮</GlobalButton>
          </div>
        </div>

        <!-- 禁用状态 -->
        <div>
          <h5 class="text-sm font-medium text-gray-700 mb-2">禁用状态</h5>
          <div class="flex flex-wrap gap-2">
            <GlobalButton type="primary" :disabled="true">禁用按钮</GlobalButton>
            <GlobalButton type="success" variant="outline" :disabled="true">禁用按钮</GlobalButton>
          </div>
        </div>

        <!-- 加载状态 -->
        <div>
          <h5 class="text-sm font-medium text-gray-700 mb-2">加载状态</h5>
          <div class="flex flex-wrap gap-2">
            <GlobalButton type="primary" :loading="isLoading" @click="toggleLoading">
              {{ isLoading ? '加载中...' : '点击加载' }}
            </GlobalButton>
            <GlobalButton type="success" :loading="true">加载中...</GlobalButton>
          </div>
        </div>
      </div>
    </div>

    <!-- 块级按钮 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">块级按钮</h4>
      <div class="space-y-2">
        <GlobalButton type="primary" :block="true">块级主要按钮</GlobalButton>
        <GlobalButton type="secondary" variant="outline" :block="true">块级次要按钮</GlobalButton>
      </div>
    </div>

    <!-- 带图标的按钮 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">带图标的按钮</h4>
      <div class="flex flex-wrap gap-3">
        <GlobalButton
          type="primary"
          icon="eye"
          icon-position="left"
        >
          查看
        </GlobalButton>
        <GlobalButton
          type="success"
          icon="check"
          icon-position="left"
        >
          确认
        </GlobalButton>
        <GlobalButton
          type="danger"
          icon="trash"
          icon-position="right"
        >
          删除
        </GlobalButton>
        <GlobalButton
          type="secondary"
          icon="download"
          icon-position="left"
          variant="outline"
        >
          下载
        </GlobalButton>
        <GlobalButton
          type="info"
          icon="settings"
          icon-position="left"
          variant="ghost"
        >
          设置
        </GlobalButton>
        <GlobalButton
          type="warning"
          icon="heart"
          icon-variant="solid"
          icon-position="left"
        >
          收藏
        </GlobalButton>
      </div>
    </div>

    <!-- 互动测试 -->
    <div class="bg-gray-50 p-4 rounded-md">
      <h4 class="font-medium text-gray-900 mb-3">互动测试区域</h4>
      <p class="text-sm text-gray-600 mb-3">点击计数: {{ clickCount }}</p>
      <div class="flex flex-wrap gap-2">
        <GlobalButton type="primary" @click="incrementCount">点击 +1</GlobalButton>
        <GlobalButton type="secondary" variant="outline" @click="resetCount">重置</GlobalButton>
        <GlobalButton type="warning" @click="simulateAsyncAction" :loading="isAsyncLoading">
          异步操作
        </GlobalButton>
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

    <!-- Slots API 说明 -->
    <div class="demo-section">
      <h4 class="demo-title">Slots API 说明</h4>
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
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useMessage } from '../../../composables/useMessage'

defineOptions({
  name: 'ButtonTest'
})

const { success, info } = useMessage()

// 状态管理
const isLoading = ref(false)
const isAsyncLoading = ref(false)
const clickCount = ref(0)

// 处理点击事件
const handleClick = (buttonType: string) => {
  info(`你点击了 ${buttonType} 按钮`)
}

// 切换加载状态
const toggleLoading = () => {
  isLoading.value = !isLoading.value
  if (isLoading.value) {
    setTimeout(() => {
      isLoading.value = false
      success('加载完成！')
    }, 2000)
  }
}

// 增加计数
const incrementCount = () => {
  clickCount.value++
}

// 重置计数
const resetCount = () => {
  clickCount.value = 0
  success('计数已重置')
}

// 模拟异步操作
const simulateAsyncAction = async () => {
  isAsyncLoading.value = true
  try {
    // 模拟延迟
    await new Promise(resolve => setTimeout(resolve, 2000))
    success('异步操作完成！')
  } finally {
    isAsyncLoading.value = false
  }
}

// API 文档数据
const propsApi = [
  { name: 'type', description: '按钮类型', type: "'primary' | 'secondary' | 'success' | 'warning' | 'danger' | 'info'", default: "'primary'" },
  { name: 'size', description: '按钮尺寸', type: "'small' | 'medium' | 'large'", default: "'medium'" },
  { name: 'loading', description: '是否为加载状态', type: 'boolean', default: 'false' },
  { name: 'disabled', description: '是否禁用', type: 'boolean', default: 'false' },
  { name: 'shape', description: '按钮形状', type: "'default' | 'round' | 'circle'", default: "'default'" },
  { name: 'variant', description: '按钮样式变体', type: "'solid' | 'outline' | 'ghost' | 'link'", default: "'solid'" },
  { name: 'block', description: '是否为块级元素', type: 'boolean', default: 'false' },
  { name: 'icon', description: '图标名称', type: 'IconName', default: 'undefined' },
  { name: 'iconPosition', description: '图标位置', type: "'left' | 'right'", default: "'left'" },
  { name: 'iconVariant', description: '图标变体', type: "'outline' | 'solid'", default: "'outline'" },
]

const eventsApi = [
  { name: 'click', description: '点击按钮时触发', params: '(event: MouseEvent)' },
]

const slotsApi = [
  { name: 'default', description: '按钮内容', params: '-' },
]
</script>
