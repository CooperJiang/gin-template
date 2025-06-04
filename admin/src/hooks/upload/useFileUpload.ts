import { ApiClient } from '@/utils/request'
import SecureStorage, { STORAGE_KEYS } from '@/utils/storage'
import type {
  UploadConfig,
  SimpleUploadResponse,
  ChunkUploadInitResponse,
  ChunkUploadResponse,
  ChunkMergeResponse,
  UploadProgressResponse,
} from '@/components/FileUpload/types'

/**
 * 文件上传Hook
 */
export function useFileUpload() {
  /**
   * 获取上传配置
   */
  const getUploadConfig = async (): Promise<UploadConfig> => {
    return ApiClient.get<UploadConfig>('/upload/config')
  }

  /**
   * 简单文件上传
   * @param formData 表单数据
   * @param onProgress 进度回调
   */
  const simpleUpload = async (
    formData: FormData,
    onProgress?: (progress: number) => void
  ): Promise<SimpleUploadResponse> => {
    return new Promise((resolve, reject) => {
      const xhr = new XMLHttpRequest()

      // 监听上传进度
      if (onProgress) {
        xhr.upload.addEventListener('progress', (event) => {
          if (event.lengthComputable) {
            const progress = (event.loaded / event.total) * 100
            onProgress(progress)
          }
        })
      }

      // 监听响应
      xhr.addEventListener('load', () => {
        if (xhr.status >= 200 && xhr.status < 300) {
          try {
            const response = JSON.parse(xhr.responseText)
            if (response.code === 200) {
              resolve(response.data)
            } else {
              reject(new Error(response.message || '上传失败'))
            }
          } catch (error) {
            reject(new Error('解析响应失败'))
          }
        } else {
          reject(new Error(`HTTP ${xhr.status}: ${xhr.statusText}`))
        }
      })

      // 监听错误
      xhr.addEventListener('error', () => {
        reject(new Error('网络错误'))
      })

      // 先打开连接
      xhr.open('POST', `${import.meta.env.VITE_API_BASE_URL || 'http://localhost:9000/api/v1'}/upload/simple`)

      // 然后设置请求头
      const token = SecureStorage.getItem<string>(STORAGE_KEYS.AUTH_TOKEN)
      if (token) {
        xhr.setRequestHeader('Authorization', `Bearer ${token}`)
      }

      // 发送请求
      xhr.send(formData)
    })
  }

  /**
   * 初始化分片上传
   * @param params 初始化参数
   */
  const initChunkUpload = async (params: {
    filename: string
    fileSize: number
    md5Hash: string
    chunkSize: number
  }): Promise<ChunkUploadInitResponse> => {
    return ApiClient.post<ChunkUploadInitResponse>('/upload/chunk/init', params)
  }

  /**
   * 上传分片
   * @param formData 分片数据
   */
  const uploadChunk = async (formData: FormData): Promise<ChunkUploadResponse> => {
    return new Promise((resolve, reject) => {
      const xhr = new XMLHttpRequest()

      // 监听响应
      xhr.addEventListener('load', () => {
        if (xhr.status >= 200 && xhr.status < 300) {
          try {
            const response = JSON.parse(xhr.responseText)
            if (response.code === 200) {
              resolve(response.data)
            } else {
              reject(new Error(response.message || '分片上传失败'))
            }
          } catch (error) {
            reject(new Error('解析响应失败'))
          }
        } else {
          reject(new Error(`HTTP ${xhr.status}: ${xhr.statusText}`))
        }
      })

      // 监听错误
      xhr.addEventListener('error', () => {
        reject(new Error('网络错误'))
      })

      // 先打开连接
      xhr.open('POST', `${import.meta.env.VITE_API_BASE_URL || 'http://localhost:9000/api/v1'}/upload/chunk`)

      // 然后设置请求头
      const token = SecureStorage.getItem<string>(STORAGE_KEYS.AUTH_TOKEN)
      if (token) {
        xhr.setRequestHeader('Authorization', `Bearer ${token}`)
      }

      // 发送请求
      xhr.send(formData)
    })
  }

  /**
   * 合并分片
   * @param params 合并参数
   */
  const mergeChunks = async (params: {
    fileID: string
  }): Promise<ChunkMergeResponse> => {
    return ApiClient.post<ChunkMergeResponse>('/upload/chunk/merge', params)
  }

  /**
   * 获取上传进度
   * @param fileID 文件ID
   */
  const getUploadProgress = async (fileID: string): Promise<UploadProgressResponse> => {
    return ApiClient.get<UploadProgressResponse>(`/upload/progress/${fileID}`)
  }

  return {
    getUploadConfig,
    simpleUpload,
    initChunkUpload,
    uploadChunk,
    mergeChunks,
    getUploadProgress,
  }
}
