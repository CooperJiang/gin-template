import type { BaseModel } from './base'

// 用户相关类型
export interface User {
  id?: string
  username: string
  email: string
  status: number
  role?: number
  avatar?: string
  bio?: string
  created_at?: string
  updated_at?: string
}

// 登录相关类型
export interface LoginRequest {
  account: string
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
  code?: string
}
