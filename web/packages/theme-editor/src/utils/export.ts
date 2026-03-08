import type { ThemeDefinition, ThemeExport } from '@app/shared/theme'

export function exportTheme(theme: ThemeDefinition) {
  const data: ThemeExport = {
    version: '1.0',
    exportedAt: new Date().toISOString(),
    theme,
  }
  const blob = new Blob([JSON.stringify(data, null, 2)], { type: 'application/json' })
  const url = URL.createObjectURL(blob)
  const a = document.createElement('a')
  a.href = url
  a.download = `theme-${theme.name.toLowerCase().replace(/\s+/g, '-')}.json`
  a.click()
  URL.revokeObjectURL(url)
}

export function importTheme(file: File): Promise<ThemeDefinition> {
  return new Promise((resolve, reject) => {
    const reader = new FileReader()
    reader.onload = () => {
      try {
        const data = JSON.parse(reader.result as string)
        const theme = data.theme || data
        if (!theme.name || !theme.light || !theme.dark || !theme.tokens) {
          reject(new Error('无效的主题配置文件'))
          return
        }
        resolve(theme)
      } catch {
        reject(new Error('JSON 解析失败'))
      }
    }
    reader.onerror = () => reject(new Error('文件读取失败'))
    reader.readAsText(file)
  })
}
