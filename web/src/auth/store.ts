import SecureStorage, { STORAGE_KEYS } from '@/lib/storage'
import type { User } from '@/types/auth'

let _token: string | null = SecureStorage.getItem<string>(STORAGE_KEYS.AUTH_TOKEN) || null
let _user: User | null = SecureStorage.getItem<User>(STORAGE_KEYS.AUTH_USER) || null
let _version = 0
const listeners = new Set<() => void>()

function notify() {
  _version++
  listeners.forEach((l) => l())
}

export function setToken(t: string | null) {
  _token = t
  if (t) {
    SecureStorage.setItem(STORAGE_KEYS.AUTH_TOKEN, t, {
      encrypt: true,
      expiry: 7 * 24 * 60 * 60 * 1000,
    })
  } else {
    SecureStorage.removeItem(STORAGE_KEYS.AUTH_TOKEN)
  }
  notify()
}

export function setUser(u: User | null) {
  _user = u
  if (u) {
    SecureStorage.setItem(STORAGE_KEYS.AUTH_USER, u, {
      encrypt: true,
      expiry: 7 * 24 * 60 * 60 * 1000,
    })
  } else {
    SecureStorage.removeItem(STORAGE_KEYS.AUTH_USER)
  }
  notify()
}

export function clearAuth() {
  _token = null
  _user = null
  SecureStorage.removeItem(STORAGE_KEYS.AUTH_TOKEN)
  SecureStorage.removeItem(STORAGE_KEYS.AUTH_USER)
  notify()
}

export function subscribe(cb: () => void) {
  listeners.add(cb)
  return () => listeners.delete(cb)
}

export function getSnapshot() {
  return _version
}

export function getToken() {
  return _token
}
export function getUser() {
  return _user
}
export function getIsAuthenticated() {
  return !!_token && !!_user
}
