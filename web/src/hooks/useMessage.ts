import { useCallback } from 'react'
import { message } from '@/lib/message'

export function useMessage() {
  const success = useCallback((content: string, duration?: number) => {
    message.success(content, duration)
  }, [])

  const error = useCallback((content: string, duration?: number) => {
    message.error(content, duration)
  }, [])

  const warning = useCallback((content: string, duration?: number) => {
    message.warning(content, duration)
  }, [])

  const info = useCallback((content: string, duration?: number) => {
    message.info(content, duration)
  }, [])

  return { success, error, warning, info }
}
