import { createHttpClient } from '@app/core'

export const ApiClient = createHttpClient({
  baseURL: import.meta.env.VITE_API_BASE_URL || '/api/v1',
  onUnauthorized: () => {
    if (!window.location.pathname.endsWith('/login')) {
      window.location.replace('/login')
    }
  },
})
