<template>
  <div class="space-y-8">
    <!-- 基础上传 -->
    <div class="bg-white p-6 rounded-lg border border-gray-200">
      <h4 class="text-lg font-semibold text-gray-900 mb-4 flex items-center gap-2">
        <GlobalIcon name="cloud-arrow-up" size="sm" color="text-indigo-500" />
        基础文件上传
      </h4>
      <div class="space-y-4">
        <p class="text-sm text-gray-600">支持拖拽上传，自动上传，支持多文件选择</p>

        <GlobalFileUpload
          :multiple="true"
          :maxCount="5"
          accept="image/*,video/*,.pdf,.txt"
          @upload:success="handleUploadSuccess"
          @upload:error="handleUploadError"
          @change="handleFileListChange"
        />
      </div>
    </div>

    <!-- 高级配置上传 -->
    <div class="bg-white p-6 rounded-lg border border-gray-200">
      <h4 class="text-lg font-semibold text-gray-900 mb-4 flex items-center gap-2">
        <GlobalIcon name="cog-6-tooth" size="sm" color="text-purple-500" />
        高级配置上传
      </h4>
      <div class="space-y-4">
        <p class="text-sm text-gray-600">自定义分片大小、禁用自动上传、限制文件类型和大小</p>

        <!-- 配置选项 -->
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4 p-4 bg-gray-50 rounded-lg">
          <div>
            <label class="block text-xs font-medium text-gray-700 mb-1">自动上传</label>
            <button
              @click="advancedConfig.autoUpload = !advancedConfig.autoUpload"
              :class="[
                'w-full px-3 py-1 text-xs rounded transition-colors',
                advancedConfig.autoUpload
                  ? 'bg-green-100 text-green-700 border border-green-300'
                  : 'bg-gray-100 text-gray-700 border border-gray-300'
              ]"
            >
              {{ advancedConfig.autoUpload ? '开启' : '关闭' }}
            </button>
          </div>
          <div>
            <label class="block text-xs font-medium text-gray-700 mb-1">分片上传</label>
            <button
              @click="advancedConfig.useChunkUpload = !advancedConfig.useChunkUpload"
              :class="[
                'w-full px-3 py-1 text-xs rounded transition-colors',
                advancedConfig.useChunkUpload
                  ? 'bg-blue-100 text-blue-700 border border-blue-300'
                  : 'bg-gray-100 text-gray-700 border border-gray-300'
              ]"
            >
              {{ advancedConfig.useChunkUpload ? '开启' : '关闭' }}
            </button>
          </div>
          <div>
            <label class="block text-xs font-medium text-gray-700 mb-1">分片大小</label>
            <select
              v-model="advancedConfig.chunkSize"
              class="w-full px-2 py-1 text-xs border border-gray-300 rounded"
            >
              <option :value="1024 * 1024">1MB</option>
              <option :value="2 * 1024 * 1024">2MB</option>
              <option :value="5 * 1024 * 1024">5MB</option>
              <option :value="10 * 1024 * 1024">10MB</option>
            </select>
          </div>
          <div>
            <label class="block text-xs font-medium text-gray-700 mb-1">最大文件数</label>
            <select
              v-model="advancedConfig.maxCount"
              class="w-full px-2 py-1 text-xs border border-gray-300 rounded"
            >
              <option :value="1">1个</option>
              <option :value="3">3个</option>
              <option :value="5">5个</option>
              <option :value="10">10个</option>
            </select>
          </div>
        </div>

        <GlobalFileUpload
          ref="advancedUploadRef"
          :multiple="advancedConfig.maxCount > 1"
          :maxCount="advancedConfig.maxCount"
          :autoUpload="advancedConfig.autoUpload"
          :useChunkUpload="advancedConfig.useChunkUpload"
          :chunkSize="advancedConfig.chunkSize"
          accept="*/*"
          @upload:progress="handleUploadProgress"
          @upload:success="handleUploadSuccess"
          @upload:error="handleUploadError"
        />

        <!-- 手动操作按钮 -->
        <div v-if="!advancedConfig.autoUpload" class="flex gap-2">
          <GlobalButton
            size="sm"
            variant="primary"
            @click="uploadAll"
          >
            <GlobalIcon name="cloud-arrow-up" size="xs" class="mr-1" />
            开始上传
          </GlobalButton>
          <GlobalButton
            size="sm"
            variant="secondary"
            @click="clearAll"
          >
            <GlobalIcon name="trash" size="xs" class="mr-1" />
            清空列表
          </GlobalButton>
        </div>
      </div>
    </div>

    <!-- 上传状态监控 -->
    <div class="bg-white p-6 rounded-lg border border-gray-200">
      <h4 class="text-lg font-semibold text-gray-900 mb-4 flex items-center gap-2">
        <GlobalIcon name="chart-bar" size="sm" color="text-green-500" />
        上传状态监控
      </h4>

      <!-- 统计信息 -->
      <div class="grid grid-cols-2 md:grid-cols-4 gap-4 mb-6">
        <div class="bg-blue-50 p-4 rounded-lg text-center">
          <div class="text-2xl font-bold text-blue-600">{{ uploadStats.total }}</div>
          <div class="text-sm text-blue-600">总文件数</div>
        </div>
        <div class="bg-green-50 p-4 rounded-lg text-center">
          <div class="text-2xl font-bold text-green-600">{{ uploadStats.success }}</div>
          <div class="text-sm text-green-600">成功上传</div>
        </div>
        <div class="bg-yellow-50 p-4 rounded-lg text-center">
          <div class="text-2xl font-bold text-yellow-600">{{ uploadStats.uploading }}</div>
          <div class="text-sm text-yellow-600">上传中</div>
        </div>
        <div class="bg-red-50 p-4 rounded-lg text-center">
          <div class="text-2xl font-bold text-red-600">{{ uploadStats.failed }}</div>
          <div class="text-sm text-red-600">上传失败</div>
        </div>
      </div>

      <!-- 最近上传记录 -->
      <div class="space-y-2">
        <h5 class="font-medium text-gray-900">最近上传记录</h5>
        <div class="max-h-40 overflow-y-auto border border-gray-200 rounded">
          <div v-if="uploadLogs.length === 0" class="p-4 text-center text-gray-500 text-sm">
            暂无上传记录
          </div>
          <div
            v-for="log in uploadLogs"
            :key="log.id"
            class="flex items-center justify-between p-3 border-b border-gray-100 last:border-b-0"
          >
            <div class="flex items-center gap-3 min-w-0 flex-1">
              <div
                :class="[
                  'w-2 h-2 rounded-full flex-shrink-0',
                  log.status === 'success' ? 'bg-green-500' :
                  log.status === 'error' ? 'bg-red-500' :
                  log.status === 'uploading' ? 'bg-blue-500' :
                  'bg-gray-300'
                ]"
              ></div>
              <div class="min-w-0 flex-1">
                <div class="text-sm font-medium text-gray-900 truncate">{{ log.fileName }}</div>
                <div class="text-xs text-gray-500">{{ formatFileSize(log.fileSize) }} • {{ formatTime(log.timestamp) }}</div>
              </div>
            </div>
            <div class="flex items-center gap-2 flex-shrink-0">
              <div v-if="log.status === 'uploading'" class="text-xs text-blue-600">
                {{ log.progress?.toFixed(1) }}%
              </div>
              <GlobalIcon
                :name="
                  log.status === 'success' ? 'check-circle' :
                  log.status === 'error' ? 'x-circle' :
                  log.status === 'uploading' ? 'arrow-path' :
                  'clock'
                "
                size="xs"
                :class="[
                  log.status === 'success' ? 'text-green-500' :
                  log.status === 'error' ? 'text-red-500' :
                  log.status === 'uploading' ? 'text-blue-500 animate-spin' :
                  'text-gray-400'
                ]"
              />
            </div>
          </div>
        </div>
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
          <h5 class="font-semibold mb-2">基本功能</h5>
          <ul class="list-disc list-inside space-y-1 ml-4">
            <li>支持拖拽上传和点击选择文件</li>
            <li>自动检测文件类型和大小限制</li>
            <li>实时显示上传进度</li>
            <li>支持暂停和恢复上传</li>
          </ul>
        </div>

        <div>
          <h5 class="font-semibold mb-2">高级特性</h5>
          <ul class="list-disc list-inside space-y-1 ml-4">
            <li>大文件自动分片上传，支持断点续传</li>
            <li>MD5哈希去重，相同文件实现秒传</li>
            <li>可配置分片大小和上传并发数</li>
            <li>完善的错误处理和重试机制</li>
          </ul>
        </div>

        <div>
          <h5 class="font-semibold mb-2">组件属性</h5>
          <div class="bg-white p-3 rounded border border-indigo-200 font-mono text-xs">
            <div>&lt;GlobalFileUpload</div>
            <div class="ml-2">:multiple="true"</div>
            <div class="ml-2">:maxCount="5"</div>
            <div class="ml-2">:maxSize="50 * 1024 * 1024"</div>
            <div class="ml-2">:chunkSize="2 * 1024 * 1024"</div>
            <div class="ml-2">:useChunkUpload="true"</div>
            <div class="ml-2">:autoUpload="true"</div>
            <div class="ml-2">accept="image/*,video/*"</div>
            <div class="ml-2">@upload:success="handleSuccess"</div>
            <div class="ml-2">@upload:error="handleError"</div>
            <div>/&gt;</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue'
import type { FileInfo } from '@/components/FileUpload/types'
import { formatFileSize } from '@/utils/format'
import { useMessage } from '@/composables/useMessage'

defineOptions({
  name: 'UploadTest'
})

const { success, error } = useMessage()

// 高级上传配置
const advancedConfig = reactive({
  autoUpload: true,
  useChunkUpload: true,
  chunkSize: 2 * 1024 * 1024, // 2MB
  maxCount: 3,
})

// 上传组件引用
const advancedUploadRef = ref()

// 上传记录
interface UploadLog {
  id: string
  fileName: string
  fileSize: number
  status: 'pending' | 'uploading' | 'success' | 'error'
  progress?: number
  timestamp: number
  error?: string
}

const uploadLogs = ref<UploadLog[]>([])

// 上传统计
const uploadStats = computed(() => {
  const total = uploadLogs.value.length
  const success = uploadLogs.value.filter(log => log.status === 'success').length
  const uploading = uploadLogs.value.filter(log => log.status === 'uploading').length
  const failed = uploadLogs.value.filter(log => log.status === 'error').length

  return { total, success, uploading, failed }
})

// 处理文件列表变化
const handleFileListChange = (fileList: FileInfo[]) => {
  console.log('文件列表变化:', fileList)

  // 添加新文件到日志
  fileList.forEach(file => {
    const existingLog = uploadLogs.value.find(log => log.id === file.id)
    if (!existingLog) {
      uploadLogs.value.unshift({
        id: file.id,
        fileName: file.name,
        fileSize: file.size,
        status: 'pending',
        timestamp: Date.now(),
      })
    }
  })
}

// 处理上传进度
const handleUploadProgress = (file: FileInfo, progress: number) => {
  console.log(`文件 ${file.name} 上传进度: ${progress}%`)

  const log = uploadLogs.value.find(log => log.id === file.id)
  if (log) {
    log.status = 'uploading'
    log.progress = progress
  }
}

// 处理上传成功
const handleUploadSuccess = (file: FileInfo, response: any) => {
  console.log('上传成功:', file, response)
  success(`文件 "${file.name}" 上传成功`)

  const log = uploadLogs.value.find(log => log.id === file.id)
  if (log) {
    log.status = 'success'
    log.progress = 100
  }
}

// 处理上传错误
const handleUploadError = (file: FileInfo, err: Error) => {
  console.error('上传失败:', file, err)
  error(`文件 "${file.name}" 上传失败: ${err.message}`)

  const log = uploadLogs.value.find(log => log.id === file.id)
  if (log) {
    log.status = 'error'
    log.error = err.message
  }
}

// 手动上传所有文件
const uploadAll = () => {
  if (advancedUploadRef.value) {
    advancedUploadRef.value.uploadAll()
  }
}

// 清空所有文件
const clearAll = () => {
  if (advancedUploadRef.value) {
    advancedUploadRef.value.clearAll()
  }
}

// 格式化时间
const formatTime = (timestamp: number) => {
  const date = new Date(timestamp)
  return date.toLocaleTimeString('zh-CN', {
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit'
  })
}
</script>
