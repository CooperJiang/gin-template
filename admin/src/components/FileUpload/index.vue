<template>
  <div class="file-upload-container">
    <!-- 上传区域 -->
    <div
      class="upload-dragger"
      :class="{
        'upload-dragger--disabled': disabled,
        'upload-dragger--dragging': isDragging,
      }"
      @drop="handleDrop"
      @dragover="handleDragOver"
      @dragenter="handleDragEnter"
      @dragleave="handleDragLeave"
    >
      <input
        ref="fileInputRef"
        type="file"
        :name="`file-input-${componentId}`"
        :accept="accept"
        :multiple="multiple"
        :disabled="disabled"
        class="upload-input"
        @change="handleFileSelect"
        @click="handleInputClick"
      />

      <div class="upload-content">
        <div class="upload-icon-wrapper">
          <div class="upload-icon-container">
            <GlobalIcon name="cloud-arrow-up" size="3xl" class="upload-icon" />
            <div class="upload-icon-bg"></div>
            <div class="upload-icon-pulse"></div>
          </div>
        </div>
        <div class="upload-text">
          <h3 class="upload-title">
            <span class="upload-title-main">拖拽文件到这里</span>
            <span class="upload-title-sub">或点击选择文件</span>
          </h3>
          <p class="upload-subtitle">
            {{ uploadTips }}
          </p>
          <div class="upload-features">
            <div class="feature-item">
              <GlobalIcon name="lightning-bolt" size="sm" class="feature-icon" />
              <span>秒传检测</span>
            </div>
            <div class="feature-item">
              <GlobalIcon name="shield-check" size="sm" class="feature-icon" />
              <span>安全加密</span>
            </div>
            <div class="feature-item">
              <GlobalIcon name="clock" size="sm" class="feature-icon" />
              <span>断点续传</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 上传列表 -->
    <div v-if="showUploadList && fileList.length > 0" class="upload-list">
      <div
        v-for="file in fileList"
        :key="file.id"
        class="upload-item"
        :class="getUploadItemClass(file.status)"
      >
        <!-- 文件信息区域 -->
        <div class="upload-item-content">
          <div class="upload-item-info">
            <div class="file-icon-wrapper">
              <GlobalIcon :name="getFileIcon(file.type)" class="file-icon" />
            </div>
            <div class="file-details">
              <div class="file-name" :title="file.name">{{ file.name }}</div>
              <div class="file-size">{{ formatFileSize(file.size) }}</div>
            </div>
          </div>

          <!-- 状态区域 -->
          <div class="upload-status-area">
            <!-- 进度条 (仅上传中显示) -->
            <div v-if="file.status === UploadStatus.UPLOADING" class="upload-progress">
              <div class="progress-info">
                <span class="progress-text">上传中</span>
                <span class="progress-percentage">{{ file.progress.toFixed(1) }}%</span>
              </div>
              <div class="progress-bar">
                <div
                  class="progress-fill"
                  :style="{ width: `${file.progress}%` }"
                ></div>
              </div>
            </div>

            <!-- 状态图标和消息 -->
            <div class="upload-status" v-else>
              <!-- 成功状态 -->
              <div v-if="file.status === UploadStatus.SUCCESS" class="status-success">
                <GlobalIcon name="check-circle" variant="solid" class="status-icon" />
                <span class="status-text">上传成功</span>
              </div>

              <!-- 错误状态 -->
              <div v-else-if="file.status === UploadStatus.ERROR" class="status-error">
                <GlobalIcon name="exclamation-triangle" variant="solid" class="status-icon" />
                <span class="status-text">上传失败</span>
              </div>

              <!-- 暂停状态 -->
              <div v-else-if="file.status === UploadStatus.PAUSED" class="status-paused">
                <GlobalIcon name="pause" class="status-icon" />
                <span class="status-text">已暂停</span>
              </div>

              <!-- 等待状态 -->
              <div v-else-if="file.status === UploadStatus.PENDING" class="status-pending">
                <GlobalIcon name="clock" class="status-icon" />
                <span class="status-text">等待上传</span>
              </div>
            </div>
          </div>

          <!-- 操作按钮 -->
          <div class="upload-actions">
            <button
              v-if="file.status === UploadStatus.PENDING"
              @click="uploadFile(file)"
              class="action-btn action-btn--upload"
              title="开始上传"
            >
              <GlobalIcon name="arrow-up" size="sm" />
            </button>
            <button
              v-if="file.status === UploadStatus.UPLOADING"
              @click="pauseUpload(file)"
              class="action-btn action-btn--pause"
              title="暂停上传"
            >
              <GlobalIcon name="pause" size="sm" />
            </button>
            <button
              v-if="file.status === UploadStatus.PAUSED"
              @click="resumeUpload(file)"
              class="action-btn action-btn--resume"
              title="继续上传"
            >
              <GlobalIcon name="arrow-up" size="sm" />
            </button>
            <button
              @click="removeFile(file)"
              class="action-btn action-btn--remove"
              title="删除文件"
            >
              <GlobalIcon name="x-mark" size="sm" />
            </button>
          </div>
        </div>

        <!-- 错误信息 (完整宽度显示在底部) -->
        <div v-if="file.status === UploadStatus.ERROR && file.error" class="error-message">
          <GlobalIcon name="exclamation-triangle" size="sm" class="error-icon" />
          <span class="error-text">{{ file.error }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, defineOptions } from 'vue'
import type {
  FileUploadProps,
  FileUploadEmits,
  FileInfo,
  UploadConfig,
  UploadTask,
  ChunkInfo
} from './types'
import { UploadStatus } from './types'
import { useFileUpload } from '@/hooks/upload/useFileUpload'
import { formatFileSize } from '@/utils/format'
import { calculateMD5 } from '@/utils/crypto'

defineOptions({
  name: 'GlobalFileUpload',
})

// Props
const props = withDefaults(defineProps<FileUploadProps>(), {
  multiple: false,
  accept: '*/*',
  maxSize: 5 * 1024 * 1024 * 1024, // 5GB
  maxCount: 10,
  chunkSize: 2 * 1024 * 1024, // 2MB
  useChunkUpload: true,
  autoUpload: true,
  listType: 'text',
  showUploadList: true,
  disabled: false,
})

// Emits
const emit = defineEmits<FileUploadEmits>()

// Refs
const fileInputRef = ref<HTMLInputElement>()
const isDragging = ref(false)
const fileList = ref<FileInfo[]>([])
const uploadConfig = ref<UploadConfig | null>(null)
const uploadTasks = ref<Map<string, UploadTask>>(new Map())
const isProcessing = ref(false)
const componentId = ref(`upload-${Date.now()}-${Math.random().toString(36).substr(2, 9)}`)

// Hooks
const {
  getUploadConfig,
  simpleUpload,
  initChunkUpload,
  uploadChunk,
  mergeChunks,
  getUploadProgress
} = useFileUpload()

// 计算属性
const uploadTips = computed(() => {
  const tips = []
  if (props.maxSize) {
    tips.push(`文件大小不超过 ${formatFileSize(props.maxSize)}`)
  }
  if (props.maxCount > 1) {
    tips.push(`最多选择 ${props.maxCount} 个文件`)
  }
  return tips.join('，') || '支持任意格式文件'
})

// 生命周期
onMounted(async () => {
  try {
    uploadConfig.value = await getUploadConfig()
  } catch (error) {
    console.error('获取上传配置失败:', error)
    // 设置默认配置
    uploadConfig.value = {
      maxFileSize: 5 * 1024 * 1024 * 1024, // 5GB
      allowedMimeTypes: {},
      chunkSize: 2 * 1024 * 1024, // 2MB
    }
  }
})

// 方法
const triggerFileSelect = () => {
  console.log(`[${componentId.value}] triggerFileSelect 被调用, disabled: ${props.disabled}, isProcessing: ${isProcessing.value}`)
  if (!props.disabled && fileInputRef.value && !isProcessing.value) {
    console.log(`[${componentId.value}] 触发文件选择`)
    fileInputRef.value.click()
  }
}

const handleInputClick = (event: Event) => {
  console.log(`[${componentId.value}] input 点击事件触发`)
  // 阻止事件冒泡，避免重复触发
  event.stopPropagation()

  if (props.disabled || isProcessing.value) {
    event.preventDefault()
    return
  }
}

const handleFileSelect = async (event: Event) => {
  const target = event.target as HTMLInputElement
  const files = Array.from(target.files || [])

  console.log(`[${componentId.value}] handleFileSelect 触发，选择文件数量:`, files.length)

  // 立即清空input值，避免重复处理
  target.value = ''

  // 如果没有选择文件，直接返回
  if (files.length === 0) {
    console.log(`[${componentId.value}] 没有选择文件，返回`)
    return
  }

  // 简化防抖逻辑，只检查是否正在处理
  if (isProcessing.value) {
    console.log(`[${componentId.value}] 正在处理文件，跳过`)
    return
  }

  console.log(`[${componentId.value}] 开始处理文件`)
  isProcessing.value = true

  try {
    await processFiles(files)
  } catch (error) {
    console.error(`[${componentId.value}] 处理文件失败:`, error)
  } finally {
    // 立即重置状态，不延迟
    isProcessing.value = false
  }
}

const handleDrop = async (event: DragEvent) => {
  event.preventDefault()
  event.stopPropagation()
  isDragging.value = false

  if (props.disabled) return

  const files = Array.from(event.dataTransfer?.files || [])

  if (isProcessing.value) return

  isProcessing.value = true

  try {
    await processFiles(files)
  } catch (error) {
    console.error('处理文件失败:', error)
  } finally {
    isProcessing.value = false
  }
}

const handleDragOver = (event: DragEvent) => {
  event.preventDefault()
  event.stopPropagation()
}

const handleDragEnter = (event: DragEvent) => {
  event.preventDefault()
  event.stopPropagation()
  if (!props.disabled) {
    isDragging.value = true
  }
}

const handleDragLeave = (event: DragEvent) => {
  event.preventDefault()
  event.stopPropagation()
  // 只有当离开整个拖拽区域时才设置为false
  const currentTarget = event.currentTarget as Element
  const relatedTarget = event.relatedTarget as Element
  if (!currentTarget?.contains?.(relatedTarget)) {
    isDragging.value = false
  }
}

const processFiles = async (files: File[]) => {
  // 检查文件数量限制
  if (fileList.value.length + files.length > props.maxCount) {
    console.warn(`[${componentId.value}] 最多只能上传 ${props.maxCount} 个文件`)
    return
  }

  const newFiles: FileInfo[] = []

  for (const file of files) {
    try {
      // 验证文件
      const isValid = await validateFile(file)
      if (!isValid) {
        console.warn(`[${componentId.value}] 文件 "${file.name}" 验证失败`)
        continue
      }

      // 创建文件信息
      const fileInfo: FileInfo = {
        id: `${componentId.value}-${Date.now()}_${Math.random().toString(36).substr(2, 9)}`,
        name: file.name,
        size: file.size,
        type: file.type,
        status: UploadStatus.PENDING,
        progress: 0,
      }

      newFiles.push(fileInfo)

      // 创建上传任务
      const task: UploadTask = {
        file,
        fileInfo,
      }
      uploadTasks.value.set(fileInfo.id, task)

      // 如果启用自动上传，立即开始上传
      if (props.autoUpload) {
        setTimeout(() => uploadFile(fileInfo), 100)
      }
    } catch (error) {
      console.warn(`[${componentId.value}] 文件验证失败:`, error)
      // 创建失败的文件信息显示在列表中
      const fileInfo: FileInfo = {
        id: `${componentId.value}-${Date.now()}_${Math.random().toString(36).substr(2, 9)}`,
        name: file.name,
        size: file.size,
        type: file.type,
        status: UploadStatus.ERROR,
        progress: 0,
        error: error instanceof Error ? error.message : '文件验证失败'
      }
      newFiles.push(fileInfo)
    }
  }

  // 更新文件列表
  if (newFiles.length > 0) {
    fileList.value.push(...newFiles)
    emit('change', fileList.value)
  } else {
    console.log(`[${componentId.value}] 没有新文件需要添加到列表`)
  }
}

const validateFile = async (file: File): Promise<boolean> => {
  // 文件大小验证
  if (file.size > props.maxSize) {
    throw new Error(`文件 "${file.name}" 大小超过限制 ${formatFileSize(props.maxSize)}`)
  }

  // 文件类型验证
  if (props.accept !== '*/*') {
    const acceptTypes = props.accept.split(',').map(t => t.trim())
    const isAccepted = acceptTypes.some(accept => {
      if (accept.startsWith('.')) {
        return file.name.toLowerCase().endsWith(accept.toLowerCase())
      }
      return file.type.match(accept.replace('*', '.*'))
    })

    if (!isAccepted) {
      throw new Error(`文件 "${file.name}" 类型不支持`)
    }
  }

  // 服务端配置验证
  if (uploadConfig.value && uploadConfig.value.allowedMimeTypes) {
    if (file.size > uploadConfig.value.maxFileSize) {
      throw new Error(`文件 "${file.name}" 大小超过服务端限制`)
    }

    // 安全访问 allowedMimeTypes
    const allowedMimeTypes = uploadConfig.value.allowedMimeTypes || {}
    if (Object.keys(allowedMimeTypes).length > 0 && !allowedMimeTypes[file.type]) {
      throw new Error(`文件 "${file.name}" 类型不被服务端支持`)
    }
  }

  // 前置验证回调
  if (props.beforeUpload) {
    return await props.beforeUpload(file)
  }

  return true
}

const uploadFile = async (fileInfo: FileInfo) => {
  const task = uploadTasks.value.get(fileInfo.id)
  if (!task) {
    console.log('没有找到上传任务:', fileInfo.id)
    return
  }

  // 检查是否已经在上传或已完成
  if (fileInfo.status === UploadStatus.UPLOADING || fileInfo.status === UploadStatus.SUCCESS) {
    console.log('文件已在上传或已完成:', fileInfo.name, fileInfo.status)
    return
  }

  try {
    // 确保状态更新能被Vue追踪
    const fileIndex = fileList.value.findIndex(f => f.id === fileInfo.id)
    if (fileIndex !== -1) {
      fileList.value[fileIndex].status = UploadStatus.UPLOADING
      fileList.value[fileIndex].progress = 0
    }

    emit('upload:before', task.file)

    // 判断是否使用分片上传
    const shouldUseChunk = props.useChunkUpload &&
      task.file.size > (props.chunkSize || 2 * 1024 * 1024)

    if (shouldUseChunk) {
      console.log('使用分片上传')
      await chunkUpload(task)
    } else {
      console.log('使用简单上传')
      await simpleUploadFile(task)
    }

    // 确保状态更新能被Vue追踪
    if (fileIndex !== -1) {
      fileList.value[fileIndex].status = UploadStatus.SUCCESS
      fileList.value[fileIndex].progress = 100
    }

    emit('upload:success', fileInfo, task)
  } catch (error) {
    console.error('文件上传失败:', fileInfo.name, error)
    // 确保状态更新能被Vue追踪
    const fileIndex = fileList.value.findIndex(f => f.id === fileInfo.id)
    if (fileIndex !== -1) {
      fileList.value[fileIndex].status = UploadStatus.ERROR
      fileList.value[fileIndex].error = error instanceof Error ? error.message : '上传失败'
    }
    emit('upload:error', fileInfo, error as Error)
  }
}

const simpleUploadFile = async (task: UploadTask) => {
  const formData = new FormData()
  formData.append('file', task.file)

  const response = await simpleUpload(formData, (progress) => {
    task.fileInfo.progress = progress
    emit('upload:progress', task.fileInfo, progress)
  })

  task.fileInfo.url = response.filePath
}

const chunkUpload = async (task: UploadTask) => {
  const { file, fileInfo } = task
  const chunkSize = props.chunkSize || 2 * 1024 * 1024

  // 计算文件MD5
  const md5Hash = await calculateMD5(file)

  // 初始化分片上传
  const initResponse = await initChunkUpload({
    filename: file.name,
    fileSize: file.size,
    md5Hash,
    chunkSize,
  })

  task.fileID = initResponse.fileID

  // 如果是秒传，直接完成
  if (initResponse.chunkTotal === 1) {
    fileInfo.progress = 100
    return
  }

  // 创建分片
  const chunks: ChunkInfo[] = []
  const totalChunks = Math.ceil(file.size / chunkSize)

  for (let i = 0; i < totalChunks; i++) {
    const start = i * chunkSize
    const end = Math.min(start + chunkSize, file.size)
    const blob = file.slice(start, end)

    chunks.push({
      index: i,
      start,
      end,
      blob,
      uploaded: false,
    })
  }

  task.chunks = chunks
  task.uploadedChunks = 0

  // 上传分片
  for (const chunk of chunks) {
    if (fileInfo.status !== UploadStatus.UPLOADING) {
      break // 如果状态不是上传中，停止上传
    }

    const chunkMD5 = await calculateMD5(chunk.blob)
    const formData = new FormData()
    formData.append('fileID', task.fileID!)
    formData.append('chunkIndex', chunk.index.toString())
    formData.append('md5Hash', chunkMD5)
    formData.append('chunk', chunk.blob)

    await uploadChunk(formData)

    chunk.uploaded = true
    task.uploadedChunks!++

    // 更新进度
    const progress = (task.uploadedChunks! / totalChunks) * 100
    fileInfo.progress = progress
    emit('upload:progress', fileInfo, progress)
  }

  // 合并分片
  if (fileInfo.status === UploadStatus.UPLOADING) {
    const mergeResponse = await mergeChunks({ fileID: task.fileID! })
    fileInfo.url = mergeResponse.filePath
  }
}

const pauseUpload = (fileInfo: FileInfo) => {
  fileInfo.status = UploadStatus.PAUSED
}

const resumeUpload = (fileInfo: FileInfo) => {
  uploadFile(fileInfo)
}

const removeFile = (fileInfo: FileInfo) => {
  const index = fileList.value.findIndex(f => f.id === fileInfo.id)
  if (index > -1) {
    fileList.value.splice(index, 1)
    uploadTasks.value.delete(fileInfo.id)
    emit('remove', fileInfo)
    emit('change', fileList.value)
  }
}

const getFileIcon = (mimeType: string): string => {
  if (mimeType.startsWith('image/')) return 'photo'
  if (mimeType.startsWith('video/')) return 'video-camera'
  if (mimeType.startsWith('audio/')) return 'musical-note'
  if (mimeType.includes('pdf')) return 'document-text'
  if (mimeType.includes('zip') || mimeType.includes('rar') || mimeType.includes('7z')) return 'archive-box'
  if (mimeType.includes('word') || mimeType.includes('doc')) return 'document'
  if (mimeType.includes('excel') || mimeType.includes('sheet')) return 'table-cells'
  if (mimeType.includes('powerpoint') || mimeType.includes('presentation')) return 'presentation-chart-line'
  return 'document'
}

const getUploadItemClass = (status: UploadStatus): string => {
  const baseClass = 'upload-item--'
  switch (status) {
    case UploadStatus.SUCCESS:
      return baseClass + 'success'
    case UploadStatus.ERROR:
      return baseClass + 'error'
    case UploadStatus.UPLOADING:
      return baseClass + 'uploading'
    case UploadStatus.PAUSED:
      return baseClass + 'paused'
    case UploadStatus.PENDING:
      return baseClass + 'pending'
    default:
      return baseClass + 'default'
  }
}

// 暴露给父组件的方法
defineExpose({
  triggerFileSelect,
  uploadAll: () => {
    fileList.value
      .filter(f => f.status === UploadStatus.PENDING)
      .forEach(uploadFile)
  },
  clearAll: () => {
    fileList.value = []
    uploadTasks.value.clear()
    emit('change', fileList.value)
  },
})
</script>

<style scoped>
.file-upload-container {
  @apply w-full;
}

.upload-dragger {
  @apply relative border-2 border-dashed border-gray-300 rounded-xl p-12 text-center cursor-pointer transition-all duration-300 ease-in-out;
  @apply hover:border-primary-400 hover:bg-primary-50 focus-within:border-primary-500 focus-within:bg-primary-50;
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%);
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
  overflow: hidden;
}

.upload-dragger::before {
  content: '';
  @apply absolute inset-0 bg-gradient-to-br from-primary-50 to-transparent opacity-0 transition-opacity duration-300;
}

.upload-dragger:hover::before {
  @apply opacity-100;
}

.upload-dragger--disabled {
  @apply cursor-not-allowed opacity-50 grayscale;
  @apply hover:border-gray-300 hover:bg-transparent focus-within:border-gray-300 focus-within:bg-transparent;
}

.upload-dragger--dragging {
  @apply border-primary-500 bg-primary-100 scale-105;
  box-shadow: 0 20px 25px -5px rgba(59, 130, 246, 0.3), 0 10px 10px -5px rgba(59, 130, 246, 0.1);
  background: linear-gradient(135deg, #dbeafe 0%, #bfdbfe 100%);
}

.upload-dragger--dragging::before {
  @apply opacity-100;
}

.upload-input {
  @apply absolute inset-0 w-full h-full opacity-0 cursor-pointer;
}

.upload-content {
  @apply flex flex-col items-center justify-center;
}

.upload-icon-wrapper {
  @apply relative mb-6;
}

.upload-icon-container {
  @apply relative w-20 h-20 mx-auto mb-4;
}

.upload-icon {
  @apply absolute w-full h-full text-gray-400 transition-all duration-300 z-10;
}

.upload-icon-bg {
  @apply absolute inset-2 rounded-full bg-gradient-to-br from-gray-200 to-gray-300 transition-all duration-300;
}

.upload-icon-pulse {
  @apply absolute inset-1 rounded-full bg-gradient-to-br from-primary-200 to-primary-300 opacity-0 transition-all duration-300;
  animation: iconPulse 2s infinite ease-in-out;
}

.upload-dragger:hover .upload-icon {
  @apply text-primary-500 scale-110;
}

.upload-dragger:hover .upload-icon-bg {
  @apply from-primary-100 to-primary-200;
}

.upload-dragger:hover .upload-icon-pulse {
  @apply opacity-100;
}

.upload-dragger--dragging .upload-icon {
  @apply text-primary-600 scale-125;
}

.upload-dragger--dragging .upload-icon-pulse {
  @apply opacity-100;
  animation: iconPulse 1s infinite ease-in-out;
}

@keyframes iconPulse {
  0%, 100% {
    transform: scale(1);
    opacity: 0.3;
  }
  50% {
    transform: scale(1.1);
    opacity: 0.7;
  }
}

.upload-text {
  @apply text-center space-y-2;
}

.upload-title {
  @apply text-xl font-semibold text-gray-800 mb-2;
}

.upload-title-main {
  @apply block;
}

.upload-title-sub {
  @apply block text-sm text-gray-600 leading-relaxed;
}

.upload-subtitle {
  @apply text-sm text-gray-600 leading-relaxed;
}

.upload-features {
  @apply flex items-center justify-center space-x-6 mt-4 pt-4 border-t border-gray-200;
}

.feature-item {
  @apply flex items-center space-x-1.5 text-xs text-gray-600 font-medium;
  @apply transition-colors duration-200 hover:text-primary-600;
}

.feature-icon {
  @apply w-4 h-4;
}

.upload-list {
  @apply mt-6 space-y-3;
}

.upload-item {
  @apply bg-white border border-gray-200 rounded-lg shadow-sm transition-all duration-300;
  @apply hover:shadow-md;
}

.upload-item--success {
  @apply border-green-200 bg-gradient-to-r from-green-50 to-emerald-50;
}

.upload-item--error {
  @apply border-red-200 bg-gradient-to-r from-red-50 to-pink-50;
}

.upload-item--uploading {
  @apply border-blue-200 bg-gradient-to-r from-blue-50 to-indigo-50;
}

.upload-item--paused {
  @apply border-orange-200 bg-gradient-to-r from-orange-50 to-amber-50;
}

.upload-item--pending {
  @apply border-gray-200 bg-gradient-to-r from-gray-50 to-slate-50;
}

.upload-item-content {
  @apply flex items-center p-4;
}

.upload-item-info {
  @apply flex items-center flex-1 min-w-0;
}

.file-icon-wrapper {
  @apply mr-4 flex-shrink-0 p-2 bg-white rounded-lg shadow-sm;
}

.file-icon {
  @apply w-8 h-8 text-gray-600;
}

.file-details {
  @apply min-w-0 flex-1;
}

.file-name {
  @apply text-sm font-medium text-gray-900 truncate mb-1;
}

.file-size {
  @apply text-xs text-gray-500;
}

.upload-status-area {
  @apply flex items-center flex-1 max-w-xs mx-4;
}

.upload-progress {
  @apply w-full space-y-2;
}

.progress-info {
  @apply flex items-center justify-between text-xs;
}

.progress-text {
  @apply text-gray-600 font-medium;
}

.progress-percentage {
  @apply text-primary-600 font-semibold;
}

.progress-bar {
  @apply w-full h-2 bg-gray-200 rounded-full overflow-hidden;
}

.progress-fill {
  @apply h-full bg-gradient-to-r from-primary-500 to-primary-600 transition-all duration-300 ease-out;
  @apply relative;
}

.progress-fill::after {
  content: '';
  @apply absolute inset-0 bg-gradient-to-r from-transparent via-white to-transparent opacity-20;
  animation: shimmer 2s infinite;
}

@keyframes shimmer {
  0% { transform: translateX(-100%); }
  100% { transform: translateX(100%); }
}

.upload-status {
  @apply flex items-center space-x-2;
}

.status-success {
  @apply flex items-center space-x-2 text-green-600;
}

.status-error {
  @apply flex items-center space-x-2 text-red-600;
}

.status-paused {
  @apply flex items-center space-x-2 text-orange-600;
}

.status-pending {
  @apply flex items-center space-x-2 text-gray-600;
}

.status-icon {
  @apply w-5 h-5 flex-shrink-0;
}

.status-text {
  @apply text-xs font-medium;
}

.upload-actions {
  @apply flex items-center space-x-2 ml-auto;
}

.action-btn {
  @apply p-2 rounded-lg transition-all duration-200 border border-transparent;
  @apply hover:scale-105 active:scale-95 focus:outline-none focus:ring-2 focus:ring-offset-1;
}

.action-btn--upload {
  @apply text-primary-600 bg-primary-50 hover:bg-primary-100 hover:text-primary-700;
  @apply focus:ring-primary-500;
}

.action-btn--pause {
  @apply text-orange-600 bg-orange-50 hover:bg-orange-100 hover:text-orange-700;
  @apply focus:ring-orange-500;
}

.action-btn--resume {
  @apply text-green-600 bg-green-50 hover:bg-green-100 hover:text-green-700;
  @apply focus:ring-green-500;
}

.action-btn--remove {
  @apply text-red-600 bg-red-50 hover:bg-red-100 hover:text-red-700;
  @apply focus:ring-red-500;
}

.error-message {
  @apply flex items-start px-4 pb-4 pt-2 text-red-600 bg-red-50 border-t border-red-100;
}

.error-icon {
  @apply w-4 h-4 text-red-500 mr-2 mt-0.5 flex-shrink-0;
}

.error-text {
  @apply text-xs leading-relaxed;
}

/* 增强的响应式设计 */
@media (max-width: 640px) {
  .upload-dragger {
    @apply p-8;
  }

  .upload-title {
    @apply text-lg;
  }

  .upload-item-content {
    @apply flex-col items-start space-y-3 p-3;
  }

  .upload-status-area {
    @apply max-w-none mx-0 w-full;
  }

  .upload-actions {
    @apply ml-0 w-full justify-end;
  }
}

/* 动画效果 */
.upload-item {
  animation: slideInUp 0.3s ease-out;
}

@keyframes slideInUp {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* 拖拽时的特殊效果 */
.upload-dragger--dragging::before {
  content: '';
  @apply absolute inset-0 rounded-xl;
  background: linear-gradient(45deg, transparent 30%, rgba(59, 130, 246, 0.1) 50%, transparent 70%);
  animation: dragEffect 2s infinite;
}

@keyframes dragEffect {
  0%, 100% { opacity: 0; }
  50% { opacity: 1; }
}
</style>
