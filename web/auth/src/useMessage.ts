import { message } from '@app/core'

export function useMessage() {
  return {
    success: message.success,
    error: message.error,
    warning: message.warning,
    info: message.info,
    remove: message.remove,
    clear: message.clear,
  }
}
