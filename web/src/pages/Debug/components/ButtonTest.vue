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
</script>
