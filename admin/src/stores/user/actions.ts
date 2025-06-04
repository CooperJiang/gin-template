// User Store Actions

import { userApi } from '@/api'
import { userStateUtils } from './state'
import type {
  UserState,
  FetchUsersParams,
  CreateUserRequest,
  UpdateUserRequest,
  User,
} from './types'

// 创建用户相关的actions
export const createUserActions = (state: UserState) => {
  // 获取用户列表
  const fetchUsers = async (params?: FetchUsersParams) => {
    try {
      userStateUtils.setLoading(state, true)
      const response = await userApi.getUsers(params)
      userStateUtils.setUsers(state, response.items, response.total)
      return response
    } catch (error) {
      console.error('获取用户列表失败:', error)
      throw error
    } finally {
      userStateUtils.setLoading(state, false)
    }
  }

  // 获取用户详情
  const fetchUser = async (id: string): Promise<User> => {
    try {
      userStateUtils.setLoading(state, true)
      const user = await userApi.getUser(id)
      userStateUtils.setCurrentUser(state, user)
      return user
    } catch (error) {
      console.error('获取用户详情失败:', error)
      throw error
    } finally {
      userStateUtils.setLoading(state, false)
    }
  }

  // 创建用户
  const createUser = async (userData: CreateUserRequest): Promise<User> => {
    try {
      userStateUtils.setLoading(state, true)
      const user = await userApi.createUser(userData)
      userStateUtils.addUser(state, user)
      console.log('创建用户成功:', user.username)
      return user
    } catch (error) {
      console.error('创建用户失败:', error)
      throw error
    } finally {
      userStateUtils.setLoading(state, false)
    }
  }

  // 更新用户
  const updateUser = async (id: string, userData: UpdateUserRequest): Promise<User> => {
    try {
      userStateUtils.setLoading(state, true)
      const user = await userApi.updateUser(id, userData)
      userStateUtils.updateUserInList(state, user)
      console.log('更新用户成功:', user.username)
      return user
    } catch (error) {
      console.error('更新用户失败:', error)
      throw error
    } finally {
      userStateUtils.setLoading(state, false)
    }
  }

  // 删除用户
  const deleteUser = async (id: string): Promise<User | undefined> => {
    try {
      userStateUtils.setLoading(state, true)
      const deletedUser = userStateUtils.removeUser(state, id)
      await userApi.deleteUser(id)
      console.log('删除用户成功:', deletedUser?.username)
      return deletedUser
    } catch (error) {
      console.error('删除用户失败:', error)
      // 如果删除失败，需要恢复状态
      // 这里可以添加更复杂的错误恢复逻辑
      throw error
    } finally {
      userStateUtils.setLoading(state, false)
    }
  }

  // 批量删除用户
  const batchDeleteUsers = async (ids: string[]): Promise<void> => {
    try {
      userStateUtils.setLoading(state, true)
      await userApi.batchDeleteUsers(ids)
      userStateUtils.removeUsers(state, ids)
      console.log('批量删除成功:', ids.length, '个用户')
    } catch (error) {
      console.error('批量删除失败:', error)
      throw error
    } finally {
      userStateUtils.setLoading(state, false)
    }
  }

  return {
    fetchUsers,
    fetchUser,
    createUser,
    updateUser,
    deleteUser,
    batchDeleteUsers,
  }
}
