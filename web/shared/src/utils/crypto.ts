import CryptoJS from 'crypto-js'

/**
 * 计算文件或Blob的MD5哈希值
 * @param file 文件或Blob对象
 */
export function calculateMD5(file: File | Blob): Promise<string> {
  return new Promise((resolve, reject) => {
    const reader = new FileReader()

    reader.onload = (event) => {
      try {
        const arrayBuffer = event.target?.result as ArrayBuffer
        const wordArray = CryptoJS.lib.WordArray.create(arrayBuffer)
        const hash = CryptoJS.MD5(wordArray)
        resolve(hash.toString(CryptoJS.enc.Hex))
      } catch (error) {
        reject(new Error('计算MD5失败'))
      }
    }

    reader.onerror = () => {
      reject(new Error('读取文件失败'))
    }

    reader.readAsArrayBuffer(file)
  })
}

/**
 * 简单的字符串哈希函数（作为降级方案）
 * @param buffer ArrayBuffer
 */
function simpleStringHash(buffer: ArrayBuffer): string {
  const uint8Array = new Uint8Array(buffer)
  let hash = 0

  for (let i = 0; i < uint8Array.length; i++) {
    const char = uint8Array[i]
    hash = ((hash << 5) - hash) + char
    hash = hash & hash // 转换为32位整数
  }

  // 转换为16进制字符串并确保长度为32位
  return Math.abs(hash).toString(16).padStart(8, '0').repeat(4)
}

/**
 * 计算字符串的SHA-256哈希值
 * @param text 要哈希的文本
 */
export async function calculateSHA256(text: string): Promise<string> {
  const encoder = new TextEncoder()
  const data = encoder.encode(text)
  const hashBuffer = await crypto.subtle.digest('SHA-256', data)
  const hashArray = Array.from(new Uint8Array(hashBuffer))
  return hashArray.map(b => b.toString(16).padStart(2, '0')).join('')
}

/**
 * 生成随机字符串
 * @param length 长度
 * @param charset 字符集
 */
export function generateRandomString(
  length = 16,
  charset = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
): string {
  let result = ''
  for (let i = 0; i < length; i++) {
    result += charset.charAt(Math.floor(Math.random() * charset.length))
  }
  return result
}

/**
 * Base64编码
 * @param str 要编码的字符串
 */
export function base64Encode(str: string): string {
  return btoa(unescape(encodeURIComponent(str)))
}

/**
 * Base64解码
 * @param str 要解码的字符串
 */
export function base64Decode(str: string): string {
  return decodeURIComponent(escape(atob(str)))
}
