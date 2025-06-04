/**
 * 安全的localStorage封装工具
 * 支持加密、有效期管理、统一接口
 */

// 简单的Base64编码/解码（可以替换为更强的加密算法）
class SimpleEncryption {
  private static key = 'gin-template-secret-key-2024'

  // 简单的异或加密
  static encrypt(text: string): string {
    const keyBytes = new TextEncoder().encode(this.key)
    const textBytes = new TextEncoder().encode(text)
    const encrypted = new Uint8Array(textBytes.length)

    for (let i = 0; i < textBytes.length; i++) {
      encrypted[i] = textBytes[i] ^ keyBytes[i % keyBytes.length]
    }

    return btoa(String.fromCharCode(...encrypted))
  }

  static decrypt(encryptedText: string): string {
    try {
      const keyBytes = new TextEncoder().encode(this.key)
      const encryptedBytes = new Uint8Array(
        atob(encryptedText)
          .split('')
          .map((char) => char.charCodeAt(0)),
      )
      const decrypted = new Uint8Array(encryptedBytes.length)

      for (let i = 0; i < encryptedBytes.length; i++) {
        decrypted[i] = encryptedBytes[i] ^ keyBytes[i % keyBytes.length]
      }

      return new TextDecoder().decode(decrypted)
    } catch (error) {
      console.warn('解密失败:', error)
      return ''
    }
  }
}

// 存储项接口
interface StorageItem<T = unknown> {
  value: T
  timestamp: number
  expiry?: number // 过期时间戳
}

// 存储配置选项
interface StorageOptions {
  encrypt?: boolean // 是否加密
  expiry?: number // 有效期（毫秒）
}

class SecureStorage {
  private static PREFIX = 'gin_template_'

  /**
   * 设置存储项
   * @param key 键名
   * @param value 值
   * @param options 配置选项
   */
  static setItem<T>(key: string, value: T, options: StorageOptions = {}): void {
    try {
      const { encrypt = true, expiry } = options
      const fullKey = this.PREFIX + key

      const item: StorageItem<T> = {
        value,
        timestamp: Date.now(),
        expiry: expiry ? Date.now() + expiry : undefined,
      }

      let jsonString = JSON.stringify(item)

      if (encrypt) {
        jsonString = SimpleEncryption.encrypt(jsonString)
      }

      localStorage.setItem(fullKey, jsonString)
    } catch (error) {
      console.error('存储失败:', error)
    }
  }

  /**
   * 获取存储项
   * @param key 键名
   * @param defaultValue 默认值
   * @param decrypt 是否解密
   */
  static getItem<T>(key: string, defaultValue: T | null = null, decrypt = true): T | null {
    try {
      const fullKey = this.PREFIX + key
      const stored = localStorage.getItem(fullKey)

      if (!stored) {
        return defaultValue
      }

      let jsonString = stored

      if (decrypt) {
        jsonString = SimpleEncryption.decrypt(stored)
        if (!jsonString) {
          // 解密失败，可能是旧数据或损坏的数据
          this.removeItem(key)
          return defaultValue
        }
      }

      const item: StorageItem<T> = JSON.parse(jsonString)

      // 检查是否过期
      if (item.expiry && Date.now() > item.expiry) {
        this.removeItem(key)
        return defaultValue
      }

      return item.value
    } catch (error) {
      console.warn('获取存储项失败:', error)
      // 清除可能损坏的数据
      this.removeItem(key)
      return defaultValue
    }
  }

  /**
   * 移除存储项
   * @param key 键名
   */
  static removeItem(key: string): void {
    try {
      const fullKey = this.PREFIX + key
      localStorage.removeItem(fullKey)
    } catch (error) {
      console.error('移除存储项失败:', error)
    }
  }

  /**
   * 清除所有存储项
   */
  static clear(): void {
    try {
      const keys = Object.keys(localStorage)
      keys.forEach((key) => {
        if (key.startsWith(this.PREFIX)) {
          localStorage.removeItem(key)
        }
      })
    } catch (error) {
      console.error('清除存储失败:', error)
    }
  }

  /**
   * 检查存储项是否存在且未过期
   * @param key 键名
   */
  static hasItem(key: string): boolean {
    return this.getItem(key) !== null
  }

  /**
   * 获取所有存储的键名
   */
  static getAllKeys(): string[] {
    try {
      const keys = Object.keys(localStorage)
      return keys
        .filter((key) => key.startsWith(this.PREFIX))
        .map((key) => key.replace(this.PREFIX, ''))
    } catch (error) {
      console.error('获取键名失败:', error)
      return []
    }
  }

  /**
   * 获取存储项的元信息
   * @param key 键名
   */
  static getItemInfo(key: string): {
    exists: boolean
    expired: boolean
    timestamp?: number
    expiry?: number
  } {
    try {
      const fullKey = this.PREFIX + key
      const stored = localStorage.getItem(fullKey)

      if (!stored) {
        return { exists: false, expired: false }
      }

      const jsonString = SimpleEncryption.decrypt(stored)
      if (!jsonString) {
        return { exists: false, expired: false }
      }

      const item: StorageItem = JSON.parse(jsonString)
      const expired = item.expiry ? Date.now() > item.expiry : false

      return {
        exists: true,
        expired,
        timestamp: item.timestamp,
        expiry: item.expiry,
      }
    } catch (error) {
      console.warn('获取存储项信息失败:', error)
      return { exists: false, expired: false }
    }
  }

  /**
   * 清理旧的localStorage数据
   * 用于从旧版本迁移到新的安全存储
   */
  static migrateFromOldStorage(): void {
    try {
      const oldKeys = [
        'auth_token',
        'auth_user',
        'login_remember',
        'user_preferences',
        'theme',
        'language'
      ]

      let cleanedCount = 0
      oldKeys.forEach(key => {
        if (localStorage.getItem(key)) {
          localStorage.removeItem(key)
          cleanedCount++
        }
      })

      if (cleanedCount > 0) {
        console.log(`已清理 ${cleanedCount} 个旧的localStorage项`)
      }
    } catch (error) {
      console.warn('清理旧数据失败:', error)
    }
  }

  /**
   * 获取存储使用统计
   */
  static getStorageStats(): {
    totalItems: number
    encryptedItems: number
    expiredItems: number
    totalSize: string
  } {
    try {
      const keys = this.getAllKeys()
      let encryptedItems = 0
      let expiredItems = 0
      let totalSize = 0

      keys.forEach(key => {
        const info = this.getItemInfo(key)
        if (info.exists) {
          if (info.expired) {
            expiredItems++
          }

          const fullKey = this.PREFIX + key
          const stored = localStorage.getItem(fullKey)
          if (stored) {
            totalSize += stored.length
            // 尝试解密，如果成功说明是加密的
            try {
              SimpleEncryption.decrypt(stored)
              encryptedItems++
            } catch {
              // 非加密数据
            }
          }
        }
      })

      return {
        totalItems: keys.length,
        encryptedItems,
        expiredItems,
        totalSize: `${(totalSize / 1024).toFixed(2)} KB`
      }
    } catch (error) {
      console.warn('获取存储统计失败:', error)
      return {
        totalItems: 0,
        encryptedItems: 0,
        expiredItems: 0,
        totalSize: '0 KB'
      }
    }
  }

  /**
   * 清理过期数据
   */
  static cleanExpiredItems(): number {
    try {
      const keys = this.getAllKeys()
      let cleanedCount = 0

      keys.forEach(key => {
        const info = this.getItemInfo(key)
        if (info.expired) {
          this.removeItem(key)
          cleanedCount++
        }
      })

      if (cleanedCount > 0) {
        console.log(`已清理 ${cleanedCount} 个过期的存储项`)
      }

      return cleanedCount
    } catch (error) {
      console.warn('清理过期数据失败:', error)
      return 0
    }
  }
}

// 预定义的存储键名常量
export const STORAGE_KEYS = {
  AUTH_TOKEN: 'auth_token',
  AUTH_USER: 'auth_user',
  USER_PREFERENCES: 'user_preferences',
  THEME: 'theme',
  LANGUAGE: 'language',
  REMEMBER_LOGIN: 'remember_login',
} as const

// 常用的有效期常量（毫秒）
export const EXPIRY_TIME = {
  MINUTE: 60 * 1000,
  HOUR: 60 * 60 * 1000,
  DAY: 24 * 60 * 60 * 1000,
  WEEK: 7 * 24 * 60 * 60 * 1000,
  MONTH: 30 * 24 * 60 * 60 * 1000,
} as const

export default SecureStorage
