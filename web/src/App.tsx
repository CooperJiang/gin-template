import { useEffect } from 'react'
import ErrorBoundary from '@/components/ErrorBoundary'
import AppRouter from '@/router'
import { initTheme } from '@/lib/theme'

export default function App() {
  useEffect(() => {
    initTheme()
  }, [])

  return (
    <ErrorBoundary>
      <AppRouter />
    </ErrorBoundary>
  )
}
