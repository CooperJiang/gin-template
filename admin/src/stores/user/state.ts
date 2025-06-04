// User Store 状态管理

import { ref } from 'vue'
import type { User, UserState } from './types'

// 创建状态
export const createUserState = (): UserState => {
  return {
    users: ref<User[]>([]),
    currentUser: ref<User | null>(null),
    loading: ref(false),
    total: ref(0),
  }
}

// 状态重置功能
export const resetUserState = (state: UserState) => {
  state.users.value = []
  state.currentUser.value = null
  state.loading.value = false
  state.total.value = 0
}

// 状态工具函数
export const userStateUtils = {
  // 设置加载状态
  setLoading: (state: UserState, loading: boolean) => {
    state.loading.value = loading
  },

  // 设置用户列表
  setUsers: (state: UserState, users: User[], total?: number) => {
    state.users.value = users
    if (total !== undefined) {
      state.total.value = total
    }
  },

  // 设置当前用户
  setCurrentUser: (state: UserState, user: User | null) => {
    state.currentUser.value = user
  },

  // 添加用户到列表
  addUser: (state: UserState, user: User) => {
    state.users.value.unshift(user)
    state.total.value += 1
  },

  // 更新列表中的用户
  updateUserInList: (state: UserState, updatedUser: User) => {
    const index = state.users.value.findIndex((u) => u.id === updatedUser.id)
    if (index !== -1) {
      state.users.value[index] = updatedUser
    }

    // 如果是当前用户，也更新
    if (state.currentUser.value?.id === updatedUser.id) {
      state.currentUser.value = updatedUser
    }
  },

  // 从列表中删除用户
  removeUser: (state: UserState, userId: string) => {
    const index = state.users.value.findIndex((u) => u.id === userId)
    if (index !== -1) {
      const removedUser = state.users.value[index]
      state.users.value.splice(index, 1)
      state.total.value -= 1
      return removedUser
    }
    return undefined
  },

  // 批量删除用户
  removeUsers: (state: UserState, userIds: string[]) => {
    const removedUsers = state.users.value.filter((u) => u.id && userIds.includes(u.id))
    state.users.value = state.users.value.filter((u) => !u.id || !userIds.includes(u.id))
    state.total.value -= userIds.length
    return removedUsers
  },
}
