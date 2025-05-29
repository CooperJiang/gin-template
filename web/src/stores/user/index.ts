// User Store 主入口文件

import { defineStore } from 'pinia'
import { createUserState } from './state'
import { createUserActions } from './actions'
import type { UserStore } from './types'

// 定义User Store
export const useUserStore = defineStore('user', (): UserStore => {
  // 创建状态
  const state = createUserState()

  // 创建actions
  const actions = createUserActions(state)

  // 返回store
  return {
    // 状态
    ...state,
    // 方法
    ...actions,
  }
})

// 导出类型和工具
export * from './types'
export * from './state'
export * from './actions'
