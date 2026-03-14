import axios, { type AxiosInstance, type AxiosRequestConfig, type AxiosResponse } from 'axios'
import type { BaseResponse } from '@/types/base'
import SecureStorage, { STORAGE_KEYS } from './storage'
import { message } from './message'

export interface HttpClientOptions {
  baseURL?: string
  timeout?: number
  onUnauthorized?: () => void
}

export function createHttpClient(options: HttpClientOptions = {}) {
  const instance: AxiosInstance = axios.create({
    baseURL: options.baseURL || '/api',
    timeout: options.timeout || 10000,
    headers: { 'Content-Type': 'application/json' },
  })

  instance.interceptors.request.use(
    (config) => {
      const token = SecureStorage.getItem<string>(STORAGE_KEYS.AUTH_TOKEN)
      if (token) {
        config.headers.Authorization = `Bearer ${token}`
      }
      if (config.method === 'get') {
        config.params = { ...config.params, _t: Date.now() }
      }
      return config
    },
    (error) => Promise.reject(error),
  )

  instance.interceptors.response.use(
    (response: AxiosResponse<BaseResponse>) => {
      const { data } = response
      if (data.code !== 200) {
        message.error(data.message || '请求失败')
        return Promise.reject(new Error(data.message || '请求失败'))
      }
      return { ...response, data: data.data } as AxiosResponse
    },
    (error) => {
      let msg = '网络错误'
      if (error.response) {
        const { status, data } = error.response
        switch (status) {
          case 401:
            msg = '未授权，请重新登录'
            SecureStorage.removeItem(STORAGE_KEYS.AUTH_TOKEN)
            SecureStorage.removeItem(STORAGE_KEYS.AUTH_USER)
            options.onUnauthorized?.()
            break
          case 403:
            msg = '拒绝访问'
            break
          case 404:
            msg = '请求的资源不存在'
            break
          case 500:
            msg = '服务器内部错误'
            break
          default:
            msg = data?.message || `请求失败 (${status})`
        }
      } else if (error.request) {
        msg = '网络连接失败，请检查网络连接'
      }
      message.error(msg)
      return Promise.reject(new Error(msg))
    },
  )

  return {
    instance,
    async get<T = unknown>(url: string, config?: AxiosRequestConfig): Promise<T> {
      const response = await instance.get(url, config)
      return response.data
    },
    async post<T = unknown>(url: string, data?: unknown, config?: AxiosRequestConfig): Promise<T> {
      const response = await instance.post(url, data, config)
      return response.data
    },
    async put<T = unknown>(url: string, data?: unknown, config?: AxiosRequestConfig): Promise<T> {
      const response = await instance.put(url, data, config)
      return response.data
    },
    async delete<T = unknown>(url: string, config?: AxiosRequestConfig): Promise<T> {
      const response = await instance.delete(url, config)
      return response.data
    },
    async patch<T = unknown>(
      url: string,
      data?: unknown,
      config?: AxiosRequestConfig,
    ): Promise<T> {
      const response = await instance.patch(url, data, config)
      return response.data
    },
  }
}

export type HttpClient = ReturnType<typeof createHttpClient>
