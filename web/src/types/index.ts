// 基础响应类型
export interface BaseResponse<T = any> {
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

// 用户相关类型
export interface User {
  id: string
  username: string
  email: string
  status: 'active' | 'inactive' | 'banned'
  role: 'admin' | 'user'
  avatar?: string
  created_at: string
  updated_at: string
}

export interface CreateUserRequest {
  username: string
  email: string
  password: string
  role?: 'admin' | 'user'
  status?: 'active' | 'inactive'
}

export interface UpdateUserRequest {
  username?: string
  email?: string
  password?: string
  role?: 'admin' | 'user'
  status?: 'active' | 'inactive'
  avatar?: string
}

export interface QueryUserRequest extends PaginationParams {
  username?: string
  email?: string
  status?: number
  role?: number
}

// 登录相关类型
export interface LoginRequest {
  account: string // 支持用户名或邮箱
  password: string
}

export interface LoginResponse {
  token: string
  user: User
}

// 注册相关类型
export interface RegisterRequest {
  username: string
  email: string
  password: string
  code: string
}

export interface RegisterResponse {
  message: string
  user: User
}

// 发送验证码请求
export interface SendCodeRequest {
  email: string
}

export interface SendCodeResponse {
  message: string
}

// 重置密码请求
export interface ResetPasswordRequest {
  email: string
  code: string
  newPassword: string
}

// 修改密码请求
export interface ChangePasswordRequest {
  oldPassword: string
  newPassword: string
}

// 更新用户资料请求
export interface UpdateProfileRequest {
  username?: string
  email?: string
  avatar?: string
}

// 通用状态枚举
export enum Status {
  INACTIVE = 0,
  ACTIVE = 1
}

export enum UserRole {
  SUPER_ADMIN = 1,
  ADMIN = 2,
  USER = 3
}

// 表格列定义
export interface TableColumn {
  key: string
  title: string
  width?: string
  sortable?: boolean
  render?: (value: any, record: any) => string
}

// 表单字段定义
export interface FormField {
  name: string
  label: string
  type: 'text' | 'email' | 'password' | 'number' | 'select' | 'textarea'
  required?: boolean
  placeholder?: string
  options?: { label: string; value: any }[]
  rules?: any[]
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
export interface ApiResponse<T = any> {
  code: number
  message: string
  data: T
}

// 错误类型
export interface ApiError {
  code: number
  message: string
  details?: any
}

export interface PaginatedResponse<T> {
  data: T[]
  total: number
  page: number
  limit: number
}
