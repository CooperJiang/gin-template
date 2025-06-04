import axios, { type AxiosInstance, type AxiosRequestConfig, type AxiosResponse } from 'axios'
import type { BaseResponse } from '@/types'
import { useMessage } from '@/composables/useMessage'
import SecureStorage, { STORAGE_KEYS } from './storage'

// 创建axios实例
const request: AxiosInstance = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:9000/api/v1',
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
})

// 获取消息实例
const { error: showError } = useMessage()

// 请求拦截器
request.interceptors.request.use(
  (config) => {
    // 添加认证token
    const token = SecureStorage.getItem<string>(STORAGE_KEYS.AUTH_TOKEN)
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }

    // 添加请求时间戳（防止缓存）
    if (config.method === 'get') {
      config.params = {
        ...config.params,
        _t: Date.now(),
      }
    }

    return config
  },
  (error) => {
    return Promise.reject(error)
  },
)

// 响应拦截器
request.interceptors.response.use(
  (response: AxiosResponse<BaseResponse>) => {
    const { data } = response

    // 检查业务状态码
    if (data.code !== 200) {
      const error = new Error(data.message || '请求失败')
      // 显示错误消息
      showError(data.message || '请求失败')
      return Promise.reject(error)
    }

    // 返回修改后的响应，将实际数据放在 data 字段中
    return {
      ...response,
      data: data.data,
    } as AxiosResponse
  },
  (error) => {
    // 处理HTTP错误
    let message = '网络错误'

    if (error.response) {
      const { status, data } = error.response

      switch (status) {
        case 401:
          message = '未授权，请重新登录'
          // 清除token并跳转到登录页
          SecureStorage.removeItem(STORAGE_KEYS.AUTH_TOKEN)
          SecureStorage.removeItem(STORAGE_KEYS.AUTH_USER)
          // 使用更优雅的方式跳转，避免整个页面重新加载
          if (window.location.pathname !== '/login') {
            window.location.replace('/login')
          }
          break
        case 403:
          message = '拒绝访问'
          break
        case 404:
          message = '请求的资源不存在'
          break
        case 500:
          message = '服务器内部错误'
          break
        default:
          message = data?.message || `请求失败 (${status})`
      }
    } else if (error.request) {
      message = '网络连接失败'
    }

    // 显示错误消息
    showError(message)

    const customError = new Error(message)
    return Promise.reject(customError)
  },
)

// 封装常用请求方法
export const ApiClient = {
  async get<T = unknown>(url: string, config?: AxiosRequestConfig): Promise<T> {
    const response = await request.get(url, config)
    return response.data
  },

  async post<T = unknown>(url: string, data?: unknown, config?: AxiosRequestConfig): Promise<T> {
    const response = await request.post(url, data, config)
    return response.data
  },

  async put<T = unknown>(url: string, data?: unknown, config?: AxiosRequestConfig): Promise<T> {
    const response = await request.put(url, data, config)
    return response.data
  },

  async delete<T = unknown>(url: string, config?: AxiosRequestConfig): Promise<T> {
    const response = await request.delete(url, config)
    return response.data
  },

  async patch<T = unknown>(url: string, data?: unknown, config?: AxiosRequestConfig): Promise<T> {
    const response = await request.patch(url, data, config)
    return response.data
  },
}

export default request
