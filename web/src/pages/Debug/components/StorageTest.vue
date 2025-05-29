<template>
  <div class="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
    <h3 class="text-lg font-semibold mb-4">安全存储测试</h3>

    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <!-- 存储统计 -->
      <div>
        <h4 class="font-medium text-gray-900 mb-3">存储统计</h4>
        <div class="bg-gray-50 p-4 rounded-md space-y-2">
          <div class="flex justify-between">
            <span class="text-sm text-gray-600">总项目数:</span>
            <span class="text-sm font-medium">{{ storageStats.totalItems }}</span>
          </div>
          <div class="flex justify-between">
            <span class="text-sm text-gray-600">加密项目:</span>
            <span class="text-sm font-medium">{{ storageStats.encryptedItems }}</span>
          </div>
          <div class="flex justify-between">
            <span class="text-sm text-gray-600">过期项目:</span>
            <span class="text-sm font-medium text-red-600">{{ storageStats.expiredItems }}</span>
          </div>
          <div class="flex justify-between">
            <span class="text-sm text-gray-600">总大小:</span>
            <span class="text-sm font-medium">{{ storageStats.totalSize }}</span>
          </div>
        </div>
      </div>

      <!-- 存储操作 -->
      <div>
        <h4 class="font-medium text-gray-900 mb-3">存储操作</h4>
        <div class="space-y-2">
          <button
            @click="refreshStats"
            class="w-full px-3 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors text-sm"
          >
            刷新统计
          </button>
          <button
            @click="cleanExpired"
            class="w-full px-3 py-2 bg-yellow-600 text-white rounded-md hover:bg-yellow-700 transition-colors text-sm"
          >
            清理过期数据
          </button>
          <button
            @click="testStorage"
            class="w-full px-3 py-2 bg-green-600 text-white rounded-md hover:bg-green-700 transition-colors text-sm"
          >
            测试存储功能
          </button>
          <button
            @click="clearAllStorage"
            class="w-full px-3 py-2 bg-red-600 text-white rounded-md hover:bg-red-700 transition-colors text-sm"
          >
            清空所有存储
          </button>
        </div>
      </div>
    </div>

    <!-- 存储项列表 -->
    <div class="mt-6">
      <h4 class="font-medium text-gray-900 mb-3">存储项列表</h4>
      <div class="bg-gray-50 p-4 rounded-md max-h-60 overflow-y-auto">
        <div v-if="storageKeys.length === 0" class="text-gray-500 text-sm text-center py-4">
          暂无存储项
        </div>
        <div v-else class="space-y-2">
          <div
            v-for="key in storageKeys"
            :key="key"
            class="flex items-center justify-between p-2 bg-white rounded border text-sm"
          >
            <span class="font-medium">{{ key }}</span>
            <div class="flex items-center space-x-2">
              <span class="text-xs text-gray-500" :title="getKeyInfo(key)">
                {{ getKeyStatus(key) }}
              </span>
              <button
                @click="removeStorageItem(key)"
                class="text-red-600 hover:text-red-800 text-xs"
              >
                删除
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useMessage } from '../../../composables/useMessage'
import SecureStorage from '../../../utils/storage'

defineOptions({
  name: 'StorageTest'
})

const { success, error, info } = useMessage()

// 存储相关状态
const storageStats = ref({
  totalItems: 0,
  encryptedItems: 0,
  expiredItems: 0,
  totalSize: '0 KB'
})

const storageKeys = ref<string[]>([])

// 刷新存储统计
const refreshStats = () => {
  storageStats.value = SecureStorage.getStorageStats()
  storageKeys.value = SecureStorage.getAllKeys()
  success('存储统计已刷新')
}

// 清理过期数据
const cleanExpired = () => {
  const count = SecureStorage.cleanExpiredItems()
  if (count > 0) {
    success(`已清理 ${count} 个过期项目`)
    refreshStats()
  } else {
    info('没有过期的项目需要清理')
  }
}

// 测试存储功能
const testStorage = () => {
  try {
    // 测试基本存储
    SecureStorage.setItem('test_basic', 'Hello World', { encrypt: true })

    // 测试有效期存储
    SecureStorage.setItem('test_expiry', 'Will expire soon', {
      encrypt: true,
      expiry: 5000 // 5秒后过期
    })

    // 测试复杂对象
    SecureStorage.setItem('test_object', {
      user: 'test',
      timestamp: Date.now(),
      config: { theme: 'dark', lang: 'zh' }
    }, { encrypt: true })

    success('存储功能测试完成，已创建3个测试项')

    // 延迟刷新统计
    setTimeout(() => {
      refreshStats()
    }, 100)
  } catch (err) {
    error('存储测试失败: ' + (err instanceof Error ? err.message : '未知错误'))
  }
}

// 清空所有存储
const clearAllStorage = () => {
  if (confirm('确定要清空所有存储数据吗？这将删除所有保存的数据（包括登录信息）！')) {
    SecureStorage.clear()
    success('所有存储数据已清空')
    refreshStats()
  }
}

// 获取键的状态
const getKeyStatus = (key: string): string => {
  const info = SecureStorage.getItemInfo(key)
  if (info.expired) return '已过期'
  if (info.expiry) {
    const remaining = info.expiry - Date.now()
    if (remaining > 0) {
      const hours = Math.floor(remaining / (1000 * 60 * 60))
      if (hours > 24) {
        const days = Math.floor(hours / 24)
        return `${days}天后过期`
      }
      return `${hours}小时后过期`
    }
  }
  return '永久'
}

// 获取键的详细信息
const getKeyInfo = (key: string): string => {
  const info = SecureStorage.getItemInfo(key)
  const created = info.timestamp ? new Date(info.timestamp).toLocaleString() : '未知'
  const expiry = info.expiry ? new Date(info.expiry).toLocaleString() : '永不过期'
  return `创建时间: ${created}\n过期时间: ${expiry}`
}

// 删除存储项
const removeStorageItem = (key: string) => {
  if (confirm(`确定要删除存储项 "${key}" 吗？`)) {
    SecureStorage.removeItem(key)
    success(`已删除存储项: ${key}`)
    refreshStats()
  }
}

// 组件挂载时刷新统计
onMounted(() => {
  refreshStats()
})
</script>
