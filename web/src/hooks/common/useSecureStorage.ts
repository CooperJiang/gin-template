import { ref, watch, type Ref } from 'vue'
import SecureStorage from '../../utils/storage'

interface SecureStorageOptions {
  encrypt?: boolean
  expiry?: number
}

export function useSecureStorage<T>(
  key: string,
  defaultValue: T,
  options: SecureStorageOptions = {},
): [Ref<T>, (newValue: T) => void, () => void] {
  // 尝试从安全存储获取初始值
  const initialValue = SecureStorage.getItem<T>(key, defaultValue) ?? defaultValue
  const value = ref<T>(initialValue) as Ref<T>

  // 监听变化并更新存储
  watch(
    value,
    (newValue) => {
      SecureStorage.setItem(key, newValue, options)
    },
    { deep: true },
  )

  const setValue = (newValue: T) => {
    value.value = newValue
    // 立即更新存储，不依赖 watch 的异步执行
    SecureStorage.setItem(key, newValue, options)
  }

  const removeValue = () => {
    SecureStorage.removeItem(key)
    value.value = defaultValue
  }

  return [value, setValue, removeValue]
}

// 便捷的预配置hook
export function useAuthStorage<T>(
  key: string,
  defaultValue: T,
): [Ref<T>, (newValue: T) => void, () => void] {
  return useSecureStorage(key, defaultValue, {
    encrypt: true,
    expiry: 7 * 24 * 60 * 60 * 1000, // 7天过期
  })
}

export function useSessionStorage<T>(
  key: string,
  defaultValue: T,
): [Ref<T>, (newValue: T) => void, () => void] {
  return useSecureStorage(key, defaultValue, {
    encrypt: true,
    expiry: 24 * 60 * 60 * 1000, // 1天过期
  })
}

export function usePersistentStorage<T>(
  key: string,
  defaultValue: T,
): [Ref<T>, (newValue: T) => void, () => void] {
  return useSecureStorage(key, defaultValue, {
    encrypt: true,
    // 无过期时间
  })
}
