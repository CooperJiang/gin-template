// 上传配置类型
export interface UploadConfig {
  maxFileSize: number
  allowedMimeTypes: Record<string, boolean>
  chunkSize: number
}

// 文件上传状态
export enum UploadStatus {
  PENDING = 'pending',    // 等待上传
  UPLOADING = 'uploading', // 上传中
  SUCCESS = 'success',     // 上传成功
  ERROR = 'error',         // 上传失败
  PAUSED = 'paused',       // 暂停
}

// 文件信息
export interface FileInfo {
  id: string
  name: string
  size: number
  type: string
  status: UploadStatus
  progress: number
  url?: string
  error?: string
}

// 简单上传响应
export interface SimpleUploadResponse {
  fileID: string
  filename: string
  storedName: string
  fileSize: number
  mimeType: string
  extension: string
  md5Hash: string
  filePath: string
  uploadedAt: string
}

// 分片上传初始化响应
export interface ChunkUploadInitResponse {
  fileID: string
  chunkSize: number
  chunkTotal: number
  uploadToken?: string
}

// 分片上传响应
export interface ChunkUploadResponse {
  fileID: string
  chunkIndex: number
  chunkUploaded: number
  chunkTotal: number
  isCompleted: boolean
}

// 分片合并响应
export interface ChunkMergeResponse {
  fileID: string
  filename: string
  storedName: string
  fileSize: number
  mimeType: string
  extension: string
  md5Hash: string
  filePath: string
  uploadedAt: string
}

// 上传进度响应
export interface UploadProgressResponse {
  fileID: string
  filename: string
  fileSize: number
  chunkTotal: number
  chunkUploaded: number
  progress: number
  status: number
}

// 组件Props
export interface FileUploadProps {
  // 基础配置
  multiple?: boolean              // 是否支持多文件
  accept?: string                 // 接受的文件类型
  maxSize?: number               // 最大文件大小（字节）
  maxCount?: number              // 最大文件数量

  // 上传模式
  chunkSize?: number             // 分片大小（字节）
  useChunkUpload?: boolean       // 是否使用分片上传
  autoUpload?: boolean           // 是否自动上传

  // 样式配置
  listType?: 'text' | 'picture' | 'picture-card'  // 列表类型
  showUploadList?: boolean       // 是否显示上传列表
  disabled?: boolean             // 是否禁用

  // 回调事件
  beforeUpload?: (file: File) => boolean | Promise<boolean>
  onProgress?: (file: FileInfo, progress: number) => void
  onSuccess?: (file: FileInfo, response: any) => void
  onError?: (file: FileInfo, error: Error) => void
  onChange?: (fileList: FileInfo[]) => void
}

// 分片信息
export interface ChunkInfo {
  index: number
  start: number
  end: number
  blob: Blob
  md5?: string
  uploaded: boolean
}

// 上传任务
export interface UploadTask {
  file: File
  fileInfo: FileInfo
  chunks?: ChunkInfo[]
  uploadedChunks?: number
  fileID?: string
}

export interface FileUploadEmits {
  'update:fileList': [fileList: FileInfo[]]
  'change': [fileList: FileInfo[]]
  'upload:before': [file: File]
  'upload:progress': [file: FileInfo, progress: number]
  'upload:success': [file: FileInfo, response: any]
  'upload:error': [file: FileInfo, error: Error]
  'remove': [file: FileInfo]
}
