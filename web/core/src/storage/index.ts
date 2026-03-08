import { SecureStorage } from '@app/shared'

interface StorageOptions {
  encrypt?: boolean
  expiry?: number
}

/** 纯 JS 存储便捷函数（无框架依赖） */
export function createStorage<T>(key: string, defaultValue: T, options: StorageOptions = {}) {
  return {
    get(): T {
      return SecureStorage.getItem<T>(key, defaultValue) ?? defaultValue
    },
    set(value: T) {
      SecureStorage.setItem(key, value, options)
    },
    remove() {
      SecureStorage.removeItem(key)
    },
  }
}

/** 预配置：认证存储（加密 + 7天过期） */
export function createAuthStorage<T>(key: string, defaultValue: T) {
  return createStorage(key, defaultValue, {
    encrypt: true,
    expiry: 7 * 24 * 60 * 60 * 1000,
  })
}

/** 预配置：会话存储（加密 + 1天过期） */
export function createSessionStorage<T>(key: string, defaultValue: T) {
  return createStorage(key, defaultValue, {
    encrypt: true,
    expiry: 24 * 60 * 60 * 1000,
  })
}

/** 预配置：持久存储（加密，无过期） */
export function createPersistentStorage<T>(key: string, defaultValue: T) {
  return createStorage(key, defaultValue, { encrypt: true })
}
