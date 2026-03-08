import './styles.css'
import { createRoot } from 'react-dom/client'
import {
  renderWithQiankun,
  qiankunWindow,
  type QiankunProps,
} from 'vite-plugin-qiankun/dist/helper'
import App from './App'
import { SecureStorage } from '@app/shared'

let root: ReturnType<typeof createRoot> | null = null

function render(props: QiankunProps = {}) {
  const { container } = props
  const mountEl = container
    ? (container.querySelector('#app') as HTMLElement)
    : document.getElementById('app')!

  SecureStorage.migrateFromOldStorage()
  SecureStorage.cleanExpiredItems()

  root = createRoot(mountEl)
  root.render(<App />)
}

renderWithQiankun({
  bootstrap() {
    if (import.meta.env.DEV) {
      console.log('[子应用 app] bootstrap')
    }
  },
  mount(props) {
    if (import.meta.env.DEV) {
      console.log('[子应用 app] mount', props)
    }
    render(props)
  },
  unmount() {
    if (import.meta.env.DEV) {
      console.log('[子应用 app] unmount')
    }
    if (root) {
      root.unmount()
      root = null
    }
  },
  update(props) {
    if (import.meta.env.DEV) {
      console.log('[子应用 app] update', props)
    }
  },
})

if (!qiankunWindow.__POWERED_BY_QIANKUN__) {
  // @ts-ignore CSS import
  import('@app/core/styles').then(async () => {
    // Load saved theme
    try {
      const raw = localStorage.getItem('theme-editor-data')
      if (raw) {
        const { theme, mode } = JSON.parse(raw)
        if (theme) {
          const { applyTheme } = await import('@app/shared/theme')
          applyTheme(theme, mode || 'light')
        }
      }
    } catch { /* use CSS defaults */ }
    render()
  })
}
