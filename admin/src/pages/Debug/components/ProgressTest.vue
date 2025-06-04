<template>
  <div class="space-y-8">
    <!-- 页面标题 -->
    <div class="border-b border-gray-200 pb-4">
      <h2 class="text-2xl font-bold text-gray-900">Progress 进度条组件测试</h2>
      <p class="mt-2 text-gray-600">测试进度条组件的各种功能和样式</p>
    </div>

    <!-- 基础线性进度条 -->
    <GlobalCard title="基础线性进度条" class="w-full">
      <div class="space-y-6">
        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-3">默认样式</h4>
          <GlobalProgress :value="basicProgress" />
        </div>

        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-3">隐藏文字</h4>
          <GlobalProgress :value="basicProgress" :show-text="false" />
        </div>

        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-3">文字在顶部</h4>
          <GlobalProgress :value="basicProgress" text-position="top" />
        </div>

        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-3">文字在内部</h4>
          <GlobalProgress :value="basicProgress" text-position="inside" stroke-width="20" />
        </div>

        <div class="flex items-center space-x-4">
          <GlobalButton @click="decreaseProgress" size="small" :disabled="basicProgress <= 0">
            <GlobalIcon name="minus" size="sm" />
          </GlobalButton>
          <span class="text-sm text-gray-600 min-w-0">{{ basicProgress }}%</span>
          <GlobalButton @click="increaseProgress" size="small" :disabled="basicProgress >= 100">
            <GlobalIcon name="plus" size="sm" />
          </GlobalButton>
          <GlobalButton @click="resetProgress" size="small" variant="outline">重置</GlobalButton>
        </div>
      </div>
    </GlobalCard>

    <!-- 不同状态 -->
    <GlobalCard title="不同状态" class="w-full">
      <div class="space-y-4">
        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-2">Normal (默认)</h4>
          <GlobalProgress :value="75" status="normal" />
        </div>

        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-2">Success</h4>
          <GlobalProgress :value="100" status="success" />
        </div>

        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-2">Warning</h4>
          <GlobalProgress :value="60" status="warning" />
        </div>

        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-2">Error</h4>
          <GlobalProgress :value="30" status="error" />
        </div>

        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-2">Info</h4>
          <GlobalProgress :value="85" status="info" />
        </div>
      </div>
    </GlobalCard>

    <!-- 不同尺寸 -->
    <GlobalCard title="不同尺寸" class="w-full">
      <div class="space-y-4">
        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-2">Small</h4>
          <GlobalProgress :value="65" size="small" />
        </div>

        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-2">Medium (默认)</h4>
          <GlobalProgress :value="65" size="medium" />
        </div>

        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-2">Large</h4>
          <GlobalProgress :value="65" size="large" />
        </div>

        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-2">自定义粗细</h4>
          <GlobalProgress :value="65" :stroke-width="15" />
        </div>
      </div>
    </GlobalCard>

    <!-- 特殊效果 -->
    <GlobalCard title="特殊效果" class="w-full">
      <div class="space-y-6">
        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-3">条纹效果</h4>
          <GlobalProgress :value="70" striped />
        </div>

        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-3">条纹 + 动画</h4>
          <GlobalProgress :value="70" striped animated />
        </div>

        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-3">不确定进度</h4>
          <GlobalProgress indeterminate />
        </div>

        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-3">自定义颜色</h4>
          <GlobalProgress :value="80" color="#ff6b6b" />
        </div>

        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-3">渐变颜色</h4>
          <GlobalProgress :value="90" :color="['#ff6b6b', '#ffa500', '#32cd32']" />
        </div>
      </div>
    </GlobalCard>

    <!-- 圆形进度条 -->
    <GlobalCard title="圆形进度条" class="w-full">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div class="text-center">
          <h4 class="text-sm font-medium text-gray-700 mb-3">默认圆形</h4>
          <GlobalProgress type="circle" :value="circleProgress" />
        </div>

        <div class="text-center">
          <h4 class="text-sm font-medium text-gray-700 mb-3">不同尺寸</h4>
          <div class="flex items-center justify-center space-x-4">
            <GlobalProgress type="circle" :value="circleProgress" size="small" :width="80" />
            <GlobalProgress type="circle" :value="circleProgress" size="large" :width="150" />
          </div>
        </div>

        <div class="text-center">
          <h4 class="text-sm font-medium text-gray-700 mb-3">不同状态</h4>
          <div class="grid grid-cols-2 gap-4">
            <GlobalProgress type="circle" :value="100" status="success" :width="80" />
            <GlobalProgress type="circle" :value="30" status="error" :width="80" />
          </div>
        </div>

        <div class="text-center">
          <h4 class="text-sm font-medium text-gray-700 mb-3">仪表盘</h4>
          <GlobalProgress type="dashboard" :value="circleProgress" />
        </div>

        <div class="text-center">
          <h4 class="text-sm font-medium text-gray-700 mb-3">不确定进度</h4>
          <GlobalProgress type="circle" indeterminate />
        </div>

        <div class="text-center">
          <h4 class="text-sm font-medium text-gray-700 mb-3">自定义格式</h4>
          <GlobalProgress
            type="circle"
            :value="circleProgress"
            :format="circleProgressFormat"
          />
        </div>
      </div>

      <div class="mt-6 flex items-center justify-center space-x-4">
        <GlobalButton @click="decreaseCircleProgress" size="small" :disabled="circleProgress <= 0">
          <GlobalIcon name="minus" size="sm" />
        </GlobalButton>
        <span class="text-sm text-gray-600 min-w-0">圆形进度: {{ circleProgress }}%</span>
        <GlobalButton @click="increaseCircleProgress" size="small" :disabled="circleProgress >= 100">
          <GlobalIcon name="plus" size="sm" />
        </GlobalButton>
      </div>
    </GlobalCard>

    <!-- 自定义插槽 -->
    <GlobalCard title="自定义插槽" class="w-full">
      <div class="space-y-6">
        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-3">自定义文字</h4>
          <GlobalProgress :value="slotProgress" text-position="right">
            <template #default="{ value }">
              <div class="flex items-center space-x-2">
                <GlobalIcon name="star" size="sm" class="text-yellow-500" />
                <span class="font-medium">{{ value }}% 完成</span>
                <GlobalIcon name="star" size="sm" class="text-yellow-500" />
              </div>
            </template>
          </GlobalProgress>
        </div>

        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-3">圆形自定义文字</h4>
          <div class="flex justify-center">
            <GlobalProgress type="circle" :value="slotProgress">
              <template #default="{ value }">
                <div class="text-center">
                  <div class="text-lg font-bold text-blue-600">{{ value }}%</div>
                  <div class="text-xs text-gray-500">已完成</div>
                </div>
              </template>
            </GlobalProgress>
          </div>
        </div>

        <div class="flex items-center justify-center space-x-4">
          <GlobalButton @click="decreaseSlotProgress" size="small" :disabled="slotProgress <= 0">
            <GlobalIcon name="minus" size="sm" />
          </GlobalButton>
          <span class="text-sm text-gray-600 min-w-0">插槽进度: {{ slotProgress }}%</span>
          <GlobalButton @click="increaseSlotProgress" size="small" :disabled="slotProgress >= 100">
            <GlobalIcon name="plus" size="sm" />
          </GlobalButton>
        </div>
      </div>
    </GlobalCard>

    <!-- 动态演示 -->
    <GlobalCard title="动态演示" class="w-full">
      <div class="space-y-6">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <h4 class="text-sm font-medium text-gray-700 mb-3">模拟下载</h4>
            <div class="space-y-3">
              <GlobalProgress :value="downloadProgress" status="info" striped animated>
                <template #default="{ value }">
                  {{ Math.round(value * 2.56) }}MB / 256MB
                </template>
              </GlobalProgress>
              <div class="flex space-x-2">
                <GlobalButton
                  @click="startDownload"
                  size="small"
                  :disabled="isDownloading"
                  type="primary"
                >
                  {{ isDownloading ? '下载中...' : '开始下载' }}
                </GlobalButton>
                <GlobalButton @click="stopDownload" size="small" variant="outline">
                  停止
                </GlobalButton>
              </div>
            </div>
          </div>

          <div>
            <h4 class="text-sm font-medium text-gray-700 mb-3">模拟上传</h4>
            <div class="space-y-3">
              <GlobalProgress
                type="circle"
                :value="uploadProgress"
                status="success"
                :width="100"
              >
                <template #default="{ value }">
                  <div class="text-center">
                    <GlobalIcon name="cloud-arrow-up" size="md" class="text-green-500 mb-1" />
                    <div class="text-xs">{{ value }}%</div>
                  </div>
                </template>
              </GlobalProgress>
              <div class="text-center">
                <GlobalButton
                  @click="startUpload"
                  size="small"
                  :disabled="isUploading"
                  type="success"
                >
                  {{ isUploading ? '上传中...' : '开始上传' }}
                </GlobalButton>
              </div>
            </div>
          </div>
        </div>

        <div>
          <h4 class="text-sm font-medium text-gray-700 mb-3">多步骤进度</h4>
          <div class="space-y-2">
            <div class="flex justify-between text-sm text-gray-600">
              <span>步骤 {{ currentStep }}/5</span>
              <span>{{ stepNames[currentStep - 1] }}</span>
            </div>
            <GlobalProgress :value="(currentStep / 5) * 100" />
            <div class="flex space-x-2">
              <GlobalButton
                @click="previousStep"
                size="small"
                :disabled="currentStep <= 1"
                variant="outline"
              >
                上一步
              </GlobalButton>
              <GlobalButton
                @click="nextStep"
                size="small"
                :disabled="currentStep >= 5"
                type="primary"
              >
                下一步
              </GlobalButton>
              <GlobalButton @click="resetSteps" size="small" variant="outline">
                重置
              </GlobalButton>
            </div>
          </div>
        </div>
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
import { ref, onUnmounted } from 'vue'

defineOptions({
  name: 'ProgressTest'
})

// 基础进度
const basicProgress = ref(65)
const circleProgress = ref(75)
const slotProgress = ref(50)

// 动态演示
const downloadProgress = ref(0)
const uploadProgress = ref(0)
const isDownloading = ref(false)
const isUploading = ref(false)
const currentStep = ref(1)

const stepNames = ['初始化', '数据准备', '处理中', '验证', '完成']

let downloadTimer: number | null = null
let uploadTimer: number | null = null

// 基础进度控制
const increaseProgress = () => {
  if (basicProgress.value < 100) {
    basicProgress.value += 10
  }
}

const decreaseProgress = () => {
  if (basicProgress.value > 0) {
    basicProgress.value -= 10
  }
}

const resetProgress = () => {
  basicProgress.value = 0
}

// 圆形进度控制
const increaseCircleProgress = () => {
  if (circleProgress.value < 100) {
    circleProgress.value += 10
  }
}

const decreaseCircleProgress = () => {
  if (circleProgress.value > 0) {
    circleProgress.value -= 10
  }
}

// 格式化函数
const circleProgressFormat = (val: number) => `${val}/100`

// 插槽进度控制
const increaseSlotProgress = () => {
  if (slotProgress.value < 100) {
    slotProgress.value += 10
  }
}

const decreaseSlotProgress = () => {
  if (slotProgress.value > 0) {
    slotProgress.value -= 10
  }
}

// 模拟下载
const startDownload = () => {
  isDownloading.value = true
  downloadProgress.value = 0

  downloadTimer = window.setInterval(() => {
    downloadProgress.value += Math.random() * 5
    if (downloadProgress.value >= 100) {
      downloadProgress.value = 100
      stopDownload()
    }
  }, 200)
}

const stopDownload = () => {
  isDownloading.value = false
  if (downloadTimer) {
    clearInterval(downloadTimer)
    downloadTimer = null
  }
}

// 模拟上传
const startUpload = () => {
  isUploading.value = true
  uploadProgress.value = 0

  uploadTimer = window.setInterval(() => {
    uploadProgress.value += Math.random() * 3
    if (uploadProgress.value >= 100) {
      uploadProgress.value = 100
      isUploading.value = false
      if (uploadTimer) {
        clearInterval(uploadTimer)
        uploadTimer = null
      }
    }
  }, 300)
}

// 步骤控制
const nextStep = () => {
  if (currentStep.value < 5) {
    currentStep.value++
  }
}

const previousStep = () => {
  if (currentStep.value > 1) {
    currentStep.value--
  }
}

const resetSteps = () => {
  currentStep.value = 1
}

// 清理定时器
onUnmounted(() => {
  if (downloadTimer) {
    clearInterval(downloadTimer)
  }
  if (uploadTimer) {
    clearInterval(uploadTimer)
  }
})

// API 文档数据
const propsApi = [
  { name: 'value', description: '当前进度值 (0-100)', type: 'number', default: '0' },
  { name: 'type', description: '进度条类型', type: "'line' | 'circle' | 'dashboard'", default: "'line'" },
  { name: 'status', description: '进度条状态', type: "'normal' | 'success' | 'warning' | 'error' | 'info'", default: "'normal'" },
  { name: 'size', description: '进度条大小', type: "'small' | 'medium' | 'large'", default: "'medium'" },
  { name: 'showText', description: '是否显示进度文字', type: 'boolean', default: 'true' },
  { name: 'textPosition', description: '进度文字位置 (仅line类型有效)', type: "'right' | 'inside' | 'top' | 'bottom'", default: "'right'" },
  { name: 'strokeWidth', description: '进度条粗细 (线条宽度)', type: 'number', default: 'undefined' },
  { name: 'width', description: '圆形进度条的直径', type: 'number', default: '120' },
  { name: 'color', description: '进度条颜色', type: 'string | string[]', default: 'undefined' },
  { name: 'backgroundColor', description: '背景颜色', type: 'string', default: "'#f3f4f6'" },
  { name: 'animated', description: '是否显示动画', type: 'boolean', default: 'true' },
  { name: 'striped', description: '是否条纹效果', type: 'boolean', default: 'false' },
  { name: 'format', description: '自定义格式化函数', type: '(value: number) => string', default: 'undefined' },
  { name: 'indeterminate', description: '是否不确定进度 (显示滚动动画)', type: 'boolean', default: 'false' },
  { name: 'steps', description: '步骤数 (分段进度条)', type: 'number', default: 'undefined' },
  { name: 'gapDegree', description: '圆形进度条的缺口角度 (仅dashboard类型)', type: 'number', default: '75' },
  { name: 'gapPosition', description: '圆形进度条的缺口位置', type: "'top' | 'bottom' | 'left' | 'right'", default: "'bottom'" }
]

const eventsApi = [
  { name: 'change', description: '进度值改变事件', params: '(value: number)' },
  { name: 'complete', description: '进度完成事件', params: '()' }
]

const slotsApi = [
  { name: 'default', description: '自定义进度文字', params: '{ value: number, status: string }' },
  { name: 'inner', description: '进度条内部内容 (仅line类型内部文字位置有效)', params: '{ value: number, status: string }' }
]
</script>
