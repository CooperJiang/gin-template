import { registerMicroApps, start, addGlobalUncaughtErrorHandler } from 'qiankun'
import type { RegistrableApp, ObjectType } from 'qiankun'
import { syncFromStorage } from '@app/auth'
import { microAppRegistry } from '@/micro-app-registry'

// 同 tab：子应用通过 CustomEvent 通知主应用
window.addEventListener('auth-state-changed', () => syncFromStorage())
// 跨 tab：storage 事件同步
window.addEventListener('storage', () => syncFromStorage())

// --- 子应用加载/错误状态（供 React 消费） ---
let _loading = false
let _error: string | null = null
let _stateVersion = 0
const stateListeners = new Set<() => void>()

function notifyState() {
  _stateVersion++
  stateListeners.forEach((l) => l())
}

export function subscribeMicroState(cb: () => void) {
  stateListeners.add(cb)
  return () => stateListeners.delete(cb)
}

export function getMicroStateSnapshot() {
  return _stateVersion
}

export function getMicroLoading() { return _loading }
export function getMicroError() { return _error }

// --- 子应用注册 ---
const isDev = import.meta.env.DEV
const enablePrefetchAll = import.meta.env.PROD && import.meta.env.VITE_QIANKUN_PREFETCH === 'all'

const microApps: Array<RegistrableApp<ObjectType>> = microAppRegistry.map((app) => ({
  name: app.name,
  entry: isDev ? app.entryDev : app.entryProd,
  container: '#subapp-container',
  activeRule: app.activeRule,
}))

function getEntryHints(entry: string): string[] {
  const hints = new Set<string>([entry])

  if (entry.startsWith('//')) {
    hints.add(`http:${entry}`)
    hints.add(`https:${entry}`)
  }

  if (entry.startsWith('http://') || entry.startsWith('https://')) {
    hints.add(entry.replace(/^https?:/, ''))
    try {
      const url = new URL(entry)
      hints.add(url.origin)
      hints.add(url.pathname)
    } catch {
      // ignore invalid URL
    }
  }

  return Array.from(hints).filter(Boolean)
}

const subAppErrorHints = microApps.map((app) => ({
  name: app.name,
  entryHints: typeof app.entry === 'string' ? getEntryHints(app.entry) : [],
}))

function isSubAppError(msg: string, filename: string): boolean {
  if (msg.includes('LOADING_SOURCE_CODE')) return true

  return subAppErrorHints.some((app) => {
    if (msg.includes(`application '${app.name}'`) || msg.includes(`application ${app.name}`)) {
      return true
    }

    return app.entryHints.some((hint) => {
      if (!hint) return false
      return filename.includes(hint) || msg.includes(hint)
    })
  })
}

export function setupMicroApps() {
  registerMicroApps(microApps, {
    beforeLoad: [
      async (app) => {
        _loading = true
        _error = null
        notifyState()
        if (isDev) console.log('[主应用] 加载子应用:', app.name)
      },
    ],
    afterMount: [
      async (app) => {
        _loading = false
        notifyState()
        if (isDev) console.log('[主应用] 子应用已挂载:', app.name)
      },
    ],
  })

  addGlobalUncaughtErrorHandler((event: unknown) => {
    const msg = event instanceof Error
      ? event.message
      : (typeof (event as any)?.message === 'string' ? (event as any).message : '子应用加载异常')

    const filename = typeof (event as any)?.filename === 'string' ? (event as any).filename : ''
    const fromSubApp = isSubAppError(msg, filename)

    if (!fromSubApp) {
      if (isDev) {
        console.warn('[主应用] 忽略非子应用异常:', msg, event)
      }
      return
    }

    console.error('[主应用] 子应用异常:', msg)
    _loading = false
    _error = msg
    notifyState()
  })

  start({
    sandbox: { experimentalStyleIsolation: true },
    // 默认关闭预加载，减少首屏压力；如需恢复可设置 VITE_QIANKUN_PREFETCH=all
    prefetch: enablePrefetchAll ? 'all' : false,
  })
}

export function retryMicroApp() {
  _error = null
  _loading = false
  notifyState()
  window.location.reload()
}
