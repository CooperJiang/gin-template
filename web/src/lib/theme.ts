import SecureStorage, { STORAGE_KEYS } from './storage'

export type Theme = 'light' | 'dark' | 'system'

function getSystemTheme(): 'light' | 'dark' {
  return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light'
}

function applyTheme(theme: Theme) {
  const resolved = theme === 'system' ? getSystemTheme() : theme
  document.documentElement.classList.toggle('dark', resolved === 'dark')
}

export function getStoredTheme(): Theme {
  return SecureStorage.getItem<Theme>(STORAGE_KEYS.THEME, null, false) || 'light'
}

export function setTheme(theme: Theme) {
  SecureStorage.setItem(STORAGE_KEYS.THEME, theme, { encrypt: false })
  applyTheme(theme)
}

export function initTheme() {
  applyTheme(getStoredTheme())
  window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', () => {
    if (getStoredTheme() === 'system') {
      applyTheme('system')
    }
  })
}
