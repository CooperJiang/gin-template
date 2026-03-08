import '@app/core/styles'
import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { SecureStorage } from '@app/shared'
import { applyTheme } from '@app/shared/theme'
import App from './App'
import { setupMicroApps } from './micro-apps'

// Load saved theme from localStorage
try {
  const raw = localStorage.getItem('theme-editor-data')
  if (raw) {
    const { theme, mode } = JSON.parse(raw)
    if (theme) applyTheme(theme, mode || 'light')
  }
} catch { /* use CSS defaults */ }

function suppressKnownExternalScriptErrors() {
  if (!import.meta.env.DEV) return

  const shouldIgnore = (msg: string) => msg.includes('useLogo is not defined')

  window.addEventListener('error', (event) => {
    if (!shouldIgnore(event.message || '')) return
    event.preventDefault()
    console.warn('[主应用] 已忽略外部脚本错误:', event.message)
  }, true)

  window.addEventListener('unhandledrejection', (event) => {
    const reason = event.reason
    const message = typeof reason === 'string'
      ? reason
      : (reason && typeof reason.message === 'string' ? reason.message : '')
    if (!shouldIgnore(message)) return
    event.preventDefault()
    console.warn('[主应用] 已忽略外部脚本 Promise 异常:', message)
  })
}

suppressKnownExternalScriptErrors()

SecureStorage.migrateFromOldStorage()
SecureStorage.cleanExpiredItems()

createRoot(document.getElementById('app')!).render(
  <StrictMode>
    <App />
  </StrictMode>,
)

setupMicroApps()
