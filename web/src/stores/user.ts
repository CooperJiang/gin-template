import { defineStore } from 'pinia'
import { ref } from 'vue'
import { userApi } from '@/api'
import type { User, CreateUserRequest, UpdateUserRequest } from '@/types'

export const useUserStore = defineStore('user', () => {
  const users = ref<User[]>([])
  const currentUser = ref<User | null>(null)
  const loading = ref(false)
  const total = ref(0)

  // 获取用户列表
  const fetchUsers = async (params?: {
    page?: number
    limit?: number
    search?: string
    status?: number
  }) => {
    try {
      loading.value = true
      const response = await userApi.getUsers(params)
      users.value = response.items
      total.value = response.total
      return response
    } catch (error) {
      console.error('获取用户列表失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  // 获取用户详情
  const fetchUser = async (id: string) => {
    try {
      loading.value = true
      const user = await userApi.getUser(id)
      currentUser.value = user
      return user
    } catch (error) {
      console.error('获取用户详情失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  // 创建用户
  const createUser = async (userData: CreateUserRequest) => {
    try {
      loading.value = true
      const user = await userApi.createUser(userData)
      users.value.unshift(user)
      total.value += 1
      console.log('创建用户成功:', user.username)
      return user
    } catch (error) {
      console.error('创建用户失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  // 更新用户
  const updateUser = async (id: string, userData: UpdateUserRequest) => {
    try {
      loading.value = true
      const user = await userApi.updateUser(id, userData)

      // 更新列表中的用户
      const index = users.value.findIndex(u => u.id === id)
      if (index !== -1) {
        users.value[index] = user
      }

      // 更新当前用户
      if (currentUser.value?.id === id) {
        currentUser.value = user
      }

      console.log('更新用户成功:', user.username)
      return user
    } catch (error) {
      console.error('更新用户失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  // 删除用户
  const deleteUser = async (id: string) => {
    try {
      loading.value = true
      const deletedUser = users.value.find(u => u.id === id)
      await userApi.deleteUser(id)

      // 从列表中移除
      const index = users.value.findIndex(u => u.id === id)
      if (index !== -1) {
        users.value.splice(index, 1)
        total.value -= 1
      }

      console.log('删除用户成功:', deletedUser?.username)
      return deletedUser
    } catch (error) {
      console.error('删除用户失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  // 批量删除用户
  const batchDeleteUsers = async (ids: string[]) => {
    try {
      loading.value = true
      await userApi.batchDeleteUsers(ids)

      // 从列表中移除
      users.value = users.value.filter(u => !ids.includes(u.id))
      total.value -= ids.length

      console.log('批量删除成功:', ids.length, '个用户')
    } catch (error) {
      console.error('批量删除失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  return {
    // 状态
    users,
    currentUser,
    loading,
    total,

    // 方法
    fetchUsers,
    fetchUser,
    createUser,
    updateUser,
    deleteUser,
    batchDeleteUsers
  }
})
