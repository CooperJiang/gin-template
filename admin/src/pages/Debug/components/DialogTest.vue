<template>
  <div class="space-y-8">
    <!-- 基础对话框 -->
    <GlobalCodeDemo
      title="基础对话框"
      description="最基本的对话框展示，支持自定义标题和内容"
      :code="basicDialogCode"
    >
      <div class="flex flex-wrap gap-4">
        <GlobalButton @click="basicDialog = true">
          基础对话框
        </GlobalButton>

        <GlobalButton @click="noHeaderDialog = true" variant="outline">
          无头部对话框
        </GlobalButton>
      </div>

      <!-- 基础对话框 -->
      <GlobalDialog
        v-model="basicDialog"
        title="基础对话框"
        :show-footer="true"
        @confirm="handleConfirm"
        @cancel="handleCancel"
      >
        <p class="text-gray-600">这是一个基础的对话框内容，您可以在这里放置任何内容。</p>
        <p class="text-gray-600 mt-2">支持多行文本和各种元素。</p>
      </GlobalDialog>

      <!-- 无头部对话框 -->
      <GlobalDialog
        v-model="noHeaderDialog"
        :show-header="false"
        :show-footer="true"
      >
        <div class="text-center">
          <h3 class="text-lg font-medium text-gray-900 mb-2">操作成功</h3>
          <p class="text-gray-500">您的操作已经成功完成！</p>
        </div>
      </GlobalDialog>
    </GlobalCodeDemo>

    <!-- 不同尺寸 -->
    <GlobalCodeDemo
      title="不同尺寸"
      description="提供多种预设尺寸：small、medium、large"
      :code="sizeDialogCode"
    >
      <div class="flex flex-wrap gap-4">
        <GlobalButton @click="smallDialog = true" size="sm">
          小型对话框
        </GlobalButton>

        <GlobalButton @click="mediumDialog = true">
          中型对话框 (默认)
        </GlobalButton>

        <GlobalButton @click="largeDialog = true" variant="outline">
          大型对话框
        </GlobalButton>
      </div>

      <!-- 小型对话框 -->
      <GlobalDialog
        v-model="smallDialog"
        title="小型对话框"
        size="small"
        :show-footer="true"
      >
        <p class="text-gray-600">这是一个小尺寸的对话框，适合简单的确认操作。</p>
      </GlobalDialog>

      <!-- 中型对话框 -->
      <GlobalDialog
        v-model="mediumDialog"
        title="中型对话框"
        size="medium"
        :show-footer="true"
      >
        <p class="text-gray-600">这是一个中等尺寸的对话框，是默认的尺寸。</p>
        <p class="text-gray-600 mt-2">适合大多数使用场景。</p>
      </GlobalDialog>

      <!-- 大型对话框 -->
      <GlobalDialog
        v-model="largeDialog"
        title="大型对话框"
        size="large"
        :show-footer="true"
      >
        <div class="space-y-4">
          <p class="text-gray-600">这是一个大尺寸的对话框，适合展示更多内容。</p>
          <div class="grid grid-cols-2 gap-4">
            <div class="bg-gray-50 p-4 rounded">
              <h4 class="font-medium mb-2">左侧内容</h4>
              <p class="text-sm text-gray-600">这里可以放置更多的内容和信息。</p>
            </div>
            <div class="bg-gray-50 p-4 rounded">
              <h4 class="font-medium mb-2">右侧内容</h4>
              <p class="text-sm text-gray-600">大型对话框为复杂布局提供更多空间。</p>
            </div>
          </div>
        </div>
      </GlobalDialog>
    </GlobalCodeDemo>

    <!-- API 说明 -->
    <div class="demo-section">
      <h4 class="demo-title">Props API 说明</h4>
      <p class="demo-description">对话框组件支持的属性配置</p>

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
import { ref } from 'vue'

defineOptions({
  name: 'DialogTest'
})

// 基础对话框状态
const basicDialog = ref(false)
const noHeaderDialog = ref(false)

// 尺寸对话框状态
const smallDialog = ref(false)
const mediumDialog = ref(false)
const largeDialog = ref(false)

// 事件处理
const handleConfirm = () => {
  console.log('确认操作')
}

const handleCancel = () => {
  console.log('取消操作')
}

// API 文档数据
const propsApi = [
  { name: 'modelValue', description: '是否显示对话框', type: 'boolean', default: 'false' },
  { name: 'title', description: '对话框标题', type: 'string', default: '""' },
  { name: 'width', description: '对话框宽度', type: 'string | number', default: '"auto"' },
  { name: 'height', description: '对话框高度', type: 'string | number', default: '"auto"' },
  { name: 'size', description: '预设尺寸', type: "'small' | 'medium' | 'large'", default: "'medium'" },
  { name: 'showHeader', description: '是否显示头部', type: 'boolean', default: 'true' },
  { name: 'showFooter', description: '是否显示底部', type: 'boolean', default: 'false' },
  { name: 'confirmText', description: '确认按钮文字', type: 'string', default: '"确认"' },
  { name: 'cancelText', description: '取消按钮文字', type: 'string', default: '"取消"' }
]

// API 表格列定义
const propsApiColumns = [
  { key: 'name', title: '参数', width: 150 },
  { key: 'description', title: '说明' },
  { key: 'type', title: '类型', width: 200 },
  { key: 'default', title: '默认值', width: 120 },
]

// 代码示例字符串 - 简化版本，避免Vue模板语法冲突
const basicDialogCode = `// 基础对话框示例
<GlobalButton @click="basicDialog = true">基础对话框</GlobalButton>

<GlobalDialog
  v-model="basicDialog"
  title="基础对话框"
  :show-footer="true"
  @confirm="handleConfirm"
>
  <p>这是一个基础的对话框内容</p>
</GlobalDialog>`

const sizeDialogCode = `// 不同尺寸对话框示例
<GlobalButton @click="smallDialog = true">小型对话框</GlobalButton>
<GlobalButton @click="largeDialog = true">大型对话框</GlobalButton>

<GlobalDialog v-model="smallDialog" title="小型对话框" size="small">
  <p>小尺寸对话框内容</p>
</GlobalDialog>

<GlobalDialog v-model="largeDialog" title="大型对话框" size="large">
  <p>大尺寸对话框内容</p>
</GlobalDialog>`
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
