// User Store 相关类型定义

import type { Ref } from 'vue'
import type { User, CreateUserRequest, UpdateUserRequest, PaginationResponse } from '@/types'

// Store State 类型
export interface UserState {
  users: Ref<User[]>
  currentUser: Ref<User | null>
  loading: Ref<boolean>
  total: Ref<number>
}

// 获取用户列表的参数类型
export interface FetchUsersParams {
  page?: number
  limit?: number
  search?: string
  status?: number
}

// 批量操作的参数类型
export interface BatchDeleteParams {
  ids: string[]
}

// Store Actions 类型
export interface UserActions {
  fetchUsers: (params?: FetchUsersParams) => Promise<PaginationResponse<User>>
  fetchUser: (id: string) => Promise<User>
  createUser: (userData: CreateUserRequest) => Promise<User>
  updateUser: (id: string, userData: UpdateUserRequest) => Promise<User>
  deleteUser: (id: string) => Promise<User | undefined>
  batchDeleteUsers: (ids: string[]) => Promise<void>
}

// 完整的User Store类型
export interface UserStore extends UserState, UserActions {}

// 导出常用类型
export type { User, CreateUserRequest, UpdateUserRequest }
