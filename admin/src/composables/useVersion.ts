import { ref, onMounted } from 'vue'
import type { VersionInfo } from '@/types/version'

// 全局变量声明
declare const __APP_VERSION__: string
declare const __BUILD_TIME__: string

export function useVersion() {
  const versionInfo = ref<VersionInfo | null>(null)
  const loading = ref(true)
  const error = ref<string | null>(null)

  const loadVersionInfo = async () => {
    try {
      loading.value = true
      error.value = null

      // 使用全局定义的版本变量
      const version = typeof __APP_VERSION__ !== 'undefined' ? __APP_VERSION__ : '1.0.0'
      const buildTime = typeof __BUILD_TIME__ !== 'undefined' ? __BUILD_TIME__ : new Date().toISOString()

      // 构建版本信息
      versionInfo.value = {
        version: version,
        buildTime: buildTime,
        buildTimestamp: new Date(buildTime).getTime(),
        buildDate: new Date(buildTime).toLocaleDateString('zh-CN', {
          year: 'numeric',
          month: '2-digit',
          day: '2-digit',
          hour: '2-digit',
          minute: '2-digit',
          second: '2-digit'
        })
      }
    } catch (err) {
      console.warn('无法加载版本信息:', err)
      error.value = '版本信息加载失败'
      // 提供默认值
      versionInfo.value = {
        version: '1.0.0',
        buildTime: new Date().toISOString(),
        buildTimestamp: Date.now(),
        buildDate: new Date().toLocaleDateString('zh-CN')
      }
    } finally {
      loading.value = false
    }
  }

  onMounted(() => {
    loadVersionInfo()
  })

  return {
    versionInfo,
    loading,
    error,
    loadVersionInfo
  }
}
