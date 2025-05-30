<template>
  <div class="space-y-8">
    <!-- 基础对话框 -->
    <div class="demo-section">
      <h4 class="demo-title">基础对话框</h4>
      <p class="demo-description">最基本的对话框展示</p>

      <div class="demo-container">
        <div class="flex flex-wrap gap-4">
          <GlobalButton @click="basicDialog = true">
            基础对话框
          </GlobalButton>

          <GlobalButton @click="noHeaderDialog = true" variant="outline">
            无头部对话框
          </GlobalButton>

          <GlobalButton @click="customTitleDialog = true" variant="secondary">
            自定义标题
          </GlobalButton>
        </div>
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
          <div class="mx-auto flex items-center justify-center h-12 w-12 rounded-full bg-green-100 mb-4">
            <GlobalIcon name="check" size="lg" color="text-green-600" />
          </div>
          <h3 class="text-lg font-medium text-gray-900 mb-2">操作成功</h3>
          <p class="text-gray-500">您的操作已经成功完成！</p>
        </div>
      </GlobalDialog>

      <!-- 自定义标题对话框 -->
      <GlobalDialog
        v-model="customTitleDialog"
        :show-footer="true"
      >
        <template #title>
          <div class="flex items-center">
            <GlobalIcon name="exclamation-triangle" size="sm" color="text-orange-500" class="mr-2" />
            <span class="text-orange-700">重要提醒</span>
          </div>
        </template>

        <div class="text-gray-600">
          <p class="mb-3">这是一个带有自定义标题的对话框。</p>
          <div class="bg-orange-50 border border-orange-200 rounded p-3">
            <p class="text-orange-800 text-sm">
              ⚠️ 请注意：此操作不可逆转，请谨慎操作。
            </p>
          </div>
        </div>
      </GlobalDialog>
    </div>

    <!-- 不同尺寸 -->
    <div class="demo-section">
      <h4 class="demo-title">不同尺寸</h4>
      <p class="demo-description">提供多种预设尺寸</p>

      <div class="demo-container">
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

          <GlobalButton @click="customSizeDialog = true" variant="secondary">
            自定义尺寸
          </GlobalButton>
        </div>
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

      <!-- 自定义尺寸对话框 -->
      <GlobalDialog
        v-model="customSizeDialog"
        title="自定义尺寸对话框"
        width="800"
        height="500"
        :show-footer="true"
      >
        <p class="text-gray-600">这是一个自定义尺寸的对话框，宽度 800px，高度 500px。</p>
      </GlobalDialog>
    </div>

    <!-- 不同位置 -->
    <div class="demo-section">
      <h4 class="demo-title">不同位置</h4>
      <p class="demo-description">支持不同的显示位置</p>

      <div class="demo-container">
        <div class="flex flex-wrap gap-4">
          <GlobalButton @click="centerDialog = true">
            居中显示 (默认)
          </GlobalButton>

          <GlobalButton @click="topDialog = true" variant="outline">
            顶部显示
          </GlobalButton>

          <GlobalButton @click="bottomDialog = true" variant="secondary">
            底部显示
          </GlobalButton>
        </div>
      </div>

      <!-- 居中对话框 -->
      <GlobalDialog
        v-model="centerDialog"
        title="居中对话框"
        placement="center"
        :show-footer="true"
      >
        <p class="text-gray-600">这是一个居中显示的对话框。</p>
      </GlobalDialog>

      <!-- 顶部对话框 -->
      <GlobalDialog
        v-model="topDialog"
        title="顶部对话框"
        placement="top"
        :show-footer="true"
      >
        <p class="text-gray-600">这是一个在页面顶部显示的对话框。</p>
      </GlobalDialog>

      <!-- 底部对话框 -->
      <GlobalDialog
        v-model="bottomDialog"
        title="底部对话框"
        placement="bottom"
        :show-footer="true"
      >
        <p class="text-gray-600">这是一个在页面底部显示的对话框。</p>
      </GlobalDialog>
    </div>

    <!-- 行为控制 -->
    <div class="demo-section">
      <h4 class="demo-title">行为控制</h4>
      <p class="demo-description">控制对话框的交互行为</p>

      <div class="demo-container">
        <div class="flex flex-wrap gap-4">
          <GlobalButton @click="noMaskCloseDialog = true">
            禁止点击遮罩关闭
          </GlobalButton>

          <GlobalButton @click="noEscCloseDialog = true" variant="outline">
            禁止 ESC 关闭
          </GlobalButton>

          <GlobalButton @click="persistentDialog = true" variant="danger">
            持久对话框
          </GlobalButton>

          <GlobalButton @click="loadingDialog = true" variant="secondary">
            加载状态
          </GlobalButton>
        </div>
      </div>

      <!-- 禁止遮罩关闭 -->
      <GlobalDialog
        v-model="noMaskCloseDialog"
        title="禁止遮罩关闭"
        :mask-closable="false"
        :show-footer="true"
      >
        <p class="text-gray-600">点击遮罩层无法关闭此对话框，只能通过按钮关闭。</p>
      </GlobalDialog>

      <!-- 禁止 ESC 关闭 -->
      <GlobalDialog
        v-model="noEscCloseDialog"
        title="禁止 ESC 关闭"
        :esc-closable="false"
        :show-footer="true"
      >
        <p class="text-gray-600">按 ESC 键无法关闭此对话框。</p>
      </GlobalDialog>

      <!-- 持久对话框 -->
      <GlobalDialog
        v-model="persistentDialog"
        title="持久对话框"
        :persistent="true"
        :show-footer="true"
        confirm-text="强制关闭"
        :show-cancel="false"
        @confirm="persistentDialog = false"
      >
        <div class="text-center">
          <div class="mx-auto flex items-center justify-center h-12 w-12 rounded-full bg-red-100 mb-4">
            <GlobalIcon name="exclamation-triangle" size="lg" color="text-red-600" />
          </div>
          <p class="text-gray-600">这是一个持久对话框，无法通过常规方式关闭。</p>
          <p class="text-sm text-gray-500 mt-2">必须点击"强制关闭"按钮才能关闭。</p>
        </div>
      </GlobalDialog>

      <!-- 加载状态对话框 -->
      <GlobalDialog
        v-model="loadingDialog"
        title="处理中..."
        :loading="isProcessing"
        :show-footer="true"
        :show-close="!isProcessing"
        confirm-text="开始处理"
        @confirm="startProcessing"
      >
        <div class="text-center">
          <p class="text-gray-600 mb-4">点击按钮开始处理任务</p>
          <div v-if="isProcessing" class="text-blue-600">
            <p class="text-sm">正在处理中，请稍候...</p>
          </div>
        </div>
      </GlobalDialog>
    </div>

    <!-- 自定义内容 -->
    <div class="demo-section">
      <h4 class="demo-title">自定义内容</h4>
      <p class="demo-description">完全自定义对话框内容</p>

      <div class="demo-container">
        <div class="flex flex-wrap gap-4">
          <GlobalButton @click="formDialog = true">
            表单对话框
          </GlobalButton>

          <GlobalButton @click="listDialog = true" variant="outline">
            列表对话框
          </GlobalButton>

          <GlobalButton @click="imageDialog = true" variant="secondary">
            图片预览对话框
          </GlobalButton>
        </div>
      </div>

      <!-- 表单对话框 -->
      <GlobalDialog
        v-model="formDialog"
        title="编辑用户信息"
        :show-footer="true"
        size="large"
        confirm-text="保存"
        @confirm="saveForm"
      >
        <form class="space-y-4">
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">姓名</label>
              <input
                v-model="formData.name"
                type="text"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                placeholder="请输入姓名"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">邮箱</label>
              <input
                v-model="formData.email"
                type="email"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                placeholder="请输入邮箱"
              />
            </div>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">角色</label>
            <GlobalSelect
              v-model="formData.role"
              :options="roleOptions"
              placeholder="请选择角色"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">描述</label>
            <textarea
              v-model="formData.description"
              rows="3"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="请输入描述"
            ></textarea>
          </div>
        </form>
      </GlobalDialog>

      <!-- 列表对话框 -->
      <GlobalDialog
        v-model="listDialog"
        title="选择用户"
        size="large"
        :show-footer="true"
        confirm-text="确认选择"
      >
        <div class="space-y-3">
          <div class="text-sm text-gray-500 mb-3">请选择一个或多个用户：</div>
          <div
            v-for="user in users"
            :key="user.id"
            class="flex items-center p-3 border border-gray-200 rounded-lg hover:bg-gray-50 cursor-pointer"
            @click="toggleUser(user.id)"
          >
            <div class="flex-shrink-0 mr-3">
              <div class="w-10 h-10 bg-blue-500 rounded-full flex items-center justify-center text-white font-medium">
                {{ user.name.charAt(0) }}
              </div>
            </div>
            <div class="flex-1">
              <h4 class="text-sm font-medium text-gray-900">{{ user.name }}</h4>
              <p class="text-sm text-gray-500">{{ user.email }}</p>
            </div>
            <div class="flex-shrink-0">
              <div
                :class="[
                  'w-5 h-5 rounded border-2 transition-colors',
                  selectedUsers.includes(user.id)
                    ? 'bg-blue-500 border-blue-500'
                    : 'border-gray-300'
                ]"
              >
                <GlobalIcon
                  v-if="selectedUsers.includes(user.id)"
                  name="check"
                  size="xs"
                  color="text-white"
                />
              </div>
            </div>
          </div>
        </div>
      </GlobalDialog>

      <!-- 图片预览对话框 -->
      <GlobalDialog
        v-model="imageDialog"
        :show-header="false"
        :show-footer="false"
        size="large"
        @mask-click="imageDialog = false"
      >
        <div class="relative">
          <button
            @click="imageDialog = false"
            class="absolute top-4 right-4 z-10 p-2 bg-black bg-opacity-50 text-white rounded-full hover:bg-opacity-70 transition-all"
          >
            <GlobalIcon name="x-mark" size="sm" />
          </button>
          <img
            src="https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=800&h=600&fit=crop"
            alt="示例图片"
            class="w-full h-auto rounded-lg"
          />
          <div class="absolute bottom-4 left-4 bg-black bg-opacity-50 text-white px-3 py-1 rounded text-sm">
            美丽的风景图片
          </div>
        </div>
      </GlobalDialog>
    </div>

    <!-- 使用说明 -->
    <div class="bg-gradient-to-r from-purple-50 to-pink-50 p-6 rounded-lg border border-purple-200">
      <h4 class="text-lg font-semibold text-purple-900 mb-4 flex items-center gap-2">
        <GlobalIcon name="book-open" size="sm" color="text-purple-600" />
        使用说明
      </h4>

      <div class="space-y-4 text-sm text-purple-800">
        <div>
          <h5 class="font-semibold mb-2">基本用法</h5>
          <div class="bg-white p-3 rounded border border-purple-200 font-mono text-xs">
            <div>&lt;GlobalDialog</div>
            <div class="ml-2">v-model="visible"</div>
            <div class="ml-2">title="对话框标题"</div>
            <div class="ml-2">:show-footer="true"</div>
            <div class="ml-2">@confirm="handleConfirm"</div>
            <div>&gt;</div>
            <div class="ml-2">对话框内容</div>
            <div>&lt;/GlobalDialog&gt;</div>
          </div>
        </div>

        <div>
          <h5 class="font-semibold mb-2">插槽使用</h5>
          <div class="bg-white p-3 rounded border border-purple-200 font-mono text-xs">
            <div>&lt;template #title&gt;自定义标题&lt;/template&gt;</div>
            <div>&lt;template #header&gt;完全自定义头部&lt;/template&gt;</div>
            <div>&lt;template #footer&gt;自定义底部&lt;/template&gt;</div>
          </div>
        </div>
      </div>
    </div>

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

    <!-- Events API 说明 -->
    <div class="demo-section">
      <h4 class="demo-title">Events API 说明</h4>
      <p class="demo-description">对话框组件支持的事件回调</p>

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

defineOptions({
  name: 'DialogTest'
})

// 基础对话框状态
const basicDialog = ref(false)
const noHeaderDialog = ref(false)
const customTitleDialog = ref(false)

// 尺寸对话框状态
const smallDialog = ref(false)
const mediumDialog = ref(false)
const largeDialog = ref(false)
const customSizeDialog = ref(false)

// 位置对话框状态
const centerDialog = ref(false)
const topDialog = ref(false)
const bottomDialog = ref(false)

// 行为控制对话框状态
const noMaskCloseDialog = ref(false)
const noEscCloseDialog = ref(false)
const persistentDialog = ref(false)
const loadingDialog = ref(false)
const isProcessing = ref(false)

// 自定义内容对话框状态
const formDialog = ref(false)
const listDialog = ref(false)
const imageDialog = ref(false)

// 表单数据
const formData = ref({
  name: 'John Doe',
  email: 'john@example.com',
  role: 'user',
  description: '这是一个示例用户描述'
})

const roleOptions = [
  { label: '管理员', value: 'admin' },
  { label: '用户', value: 'user' },
  { label: '访客', value: 'guest' }
]

// 用户列表数据
const users = [
  { id: 1, name: 'Alice Johnson', email: 'alice@example.com' },
  { id: 2, name: 'Bob Smith', email: 'bob@example.com' },
  { id: 3, name: 'Charlie Brown', email: 'charlie@example.com' },
  { id: 4, name: 'Diana Davis', email: 'diana@example.com' }
]

const selectedUsers = ref<number[]>([1, 3])

// 事件处理
const handleConfirm = () => {
  console.log('确认操作')
}

const handleCancel = () => {
  console.log('取消操作')
}

const startProcessing = () => {
  isProcessing.value = true
  setTimeout(() => {
    isProcessing.value = false
    loadingDialog.value = false
  }, 3000)
}

const saveForm = () => {
  console.log('保存表单数据:', formData.value)
  formDialog.value = false
}

const toggleUser = (userId: number) => {
  const index = selectedUsers.value.indexOf(userId)
  if (index > -1) {
    selectedUsers.value.splice(index, 1)
  } else {
    selectedUsers.value.push(userId)
  }
}

// API 文档数据
const propsApi = [
  { name: 'modelValue', description: '是否显示对话框', type: 'boolean', default: 'false' },
  { name: 'title', description: '对话框标题', type: 'string', default: '""' },
  { name: 'width', description: '对话框宽度', type: 'string | number', default: '"auto"' },
  { name: 'height', description: '对话框高度', type: 'string | number', default: '"auto"' },
  { name: 'size', description: '预设尺寸', type: "'small' | 'medium' | 'large' | 'full'", default: "'medium'" },
  { name: 'placement', description: '显示位置', type: "'center' | 'top' | 'bottom'", default: "'center'" },
  { name: 'closable', description: '是否可关闭', type: 'boolean', default: 'true' },
  { name: 'maskClosable', description: '点击遮罩是否关闭', type: 'boolean', default: 'true' },
  { name: 'escClosable', description: 'ESC键是否关闭', type: 'boolean', default: 'true' },
  { name: 'showMask', description: '是否显示遮罩', type: 'boolean', default: 'true' },
  { name: 'showHeader', description: '是否显示头部', type: 'boolean', default: 'true' },
  { name: 'showFooter', description: '是否显示底部', type: 'boolean', default: 'false' },
  { name: 'showClose', description: '是否显示关闭按钮', type: 'boolean', default: 'true' },
  { name: 'loading', description: '是否加载中', type: 'boolean', default: 'false' },
  { name: 'persistent', description: '是否持久化', type: 'boolean', default: 'false' },
  { name: 'confirmText', description: '确认按钮文字', type: 'string', default: '"确认"' },
  { name: 'cancelText', description: '取消按钮文字', type: 'string', default: '"取消"' },
  { name: 'confirmType', description: '确认按钮类型', type: "'primary' | 'success' | 'warning' | 'danger'", default: "'primary'" },
  { name: 'zIndex', description: '层级', type: 'number', default: '1000' },
]

const eventsApi = [
  { name: 'update:modelValue', description: '显示状态变化', params: '(visible: boolean)' },
  { name: 'open', description: '开始打开时触发', params: '()' },
  { name: 'opened', description: '完全打开后触发', params: '()' },
  { name: 'close', description: '开始关闭时触发', params: '()' },
  { name: 'closed', description: '完全关闭后触发', params: '()' },
  { name: 'confirm', description: '点击确认按钮', params: '()' },
  { name: 'cancel', description: '点击取消按钮', params: '()' },
  { name: 'mask-click', description: '点击遮罩层', params: '()' },
  { name: 'esc-press', description: '按下ESC键', params: '()' },
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
    width: 200,
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
