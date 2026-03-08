// 基础响应类型
export interface BaseResponse<T = unknown> {
  code: number
  message: string
  data: T
  request_id?: string
  timestamp?: string
}

// 分页请求参数
export interface PaginationParams {
  page?: number
  size?: number
}

// 分页响应数据
export interface PaginationResponse<T> {
  items: T[]
  total: number
  page: number
  page_size: number
  total_pages: number
}

// 基础模型
export interface BaseModel {
  id: string
  created_at: string
  updated_at: string
}

// 通用状态枚举
export enum Status {
  INACTIVE = 0,
  ACTIVE = 1,
}

export enum UserRole {
  SUPER_ADMIN = 1,
  ADMIN = 2,
  USER = 3,
}

// 表单字段定义
export interface FormField {
  name: string
  label: string
  type: 'text' | 'email' | 'password' | 'number' | 'select' | 'textarea'
  required?: boolean
  placeholder?: string
  options?: { label: string; value: string | number }[]
  rules?: Array<(value: unknown) => boolean | string>
}

// 菜单项类型
export interface MenuItem {
  id: string
  title: string
  icon?: string
  path?: string
  children?: MenuItem[]
}

// API响应类型
export interface ApiResponse<T = unknown> {
  code: number
  message: string
  data: T
  success: boolean
}

// 错误类型
export interface ApiError {
  code: number
  message: string
  details?: Record<string, unknown>
}

export interface PaginatedResponse<T> {
  data: T[]
  total: number
  page: number
  limit: number
}

// API 分页查询参数
export interface PageQuery {
  page?: number
  limit?: number
  search?: string
}

// API 分页响应
export interface PageResponse<T> {
  list: T[]
  total: number
  page: number
  limit: number
}

// HTTP 请求配置
export interface RequestConfig {
  timeout?: number
  headers?: Record<string, string>
  params?: Record<string, unknown>
  data?: unknown
}

// HTTP 响应类型
export interface HttpResponse<T = unknown> {
  data: T
  status: number
  statusText: string
  headers: Record<string, string>
}
