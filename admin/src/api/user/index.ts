import { ApiClient } from '@/utils/request'
import type {
  User,
  CreateUserRequest,
  UpdateUserRequest,
  QueryUserRequest,
  PaginationResponse,
} from '@/types'

export const userApi = {
  // 获取用户列表
  getUsers(params?: QueryUserRequest): Promise<PaginationResponse<User>> {
    return ApiClient.get('/users', { params })
  },

  // 获取用户详情
  getUser(id: string): Promise<User> {
    return ApiClient.get(`/users/${id}`)
  },

  // 创建用户
  createUser(data: CreateUserRequest): Promise<User> {
    return ApiClient.post('/users', data)
  },

  // 更新用户
  updateUser(id: string, data: UpdateUserRequest): Promise<User> {
    return ApiClient.put(`/users/${id}`, data)
  },

  // 删除用户
  deleteUser(id: string): Promise<void> {
    return ApiClient.delete(`/users/${id}`)
  },

  // 批量删除用户
  batchDeleteUsers(ids: string[]): Promise<void> {
    return ApiClient.delete('/users/batch', { data: { ids } })
  },

  // 启用/禁用用户
  toggleUserStatus(id: string, status: number): Promise<User> {
    return ApiClient.patch(`/users/${id}/status`, { status })
  },

  // 重置用户密码
  resetUserPassword(id: string, password: string): Promise<void> {
    return ApiClient.put(`/users/${id}/password`, { password })
  },

  // 获取用户统计信息
  getUserStats(): Promise<{
    total: number
    active: number
    inactive: number
    new_today: number
  }> {
    return ApiClient.get('/users/stats')
  },
}
