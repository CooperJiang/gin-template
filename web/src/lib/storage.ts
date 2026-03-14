class SimpleEncryption {
  private static key = 'gin-template-secret-key-2024'

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
    } catch {
      console.warn('Decryption failed')
      return ''
    }
  }
}

interface StorageItem<T = unknown> {
  value: T
  timestamp: number
  expiry?: number
}

interface StorageOptions {
  encrypt?: boolean
  expiry?: number
}

class SecureStorage {
  private static PREFIX = 'gin_template_frontend_'

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
      console.error('Storage setItem failed:', error)
    }
  }

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
          this.removeItem(key)
          return defaultValue
        }
      }

      const item: StorageItem<T> = JSON.parse(jsonString)

      if (item.expiry && Date.now() > item.expiry) {
        this.removeItem(key)
        return defaultValue
      }

      return item.value
    } catch {
      this.removeItem(key)
      return defaultValue
    }
  }

  static removeItem(key: string): void {
    try {
      const fullKey = this.PREFIX + key
      localStorage.removeItem(fullKey)
    } catch (error) {
      console.error('Storage removeItem failed:', error)
    }
  }

  static clear(): void {
    try {
      const keys = Object.keys(localStorage)
      keys.forEach((key) => {
        if (key.startsWith(this.PREFIX)) {
          localStorage.removeItem(key)
        }
      })
    } catch (error) {
      console.error('Storage clear failed:', error)
    }
  }
}

export const STORAGE_KEYS = {
  AUTH_TOKEN: 'auth_token',
  AUTH_USER: 'auth_user',
  THEME: 'theme',
  LANGUAGE: 'language',
} as const

export default SecureStorage
