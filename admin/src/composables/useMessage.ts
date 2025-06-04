import { ref } from 'vue'

export type MessageType = 'success' | 'error' | 'warning' | 'info'

export interface Message {
  id: string
  type: MessageType
  content: string
  duration?: number
  showClose?: boolean
  timer?: ReturnType<typeof setTimeout>
}

// 全局消息队列
const messages = ref<Message[]>([])

let messageId = 0

const generateId = () => `message_${++messageId}_${Date.now()}`

// 添加消息
const addMessage = (options: Omit<Message, 'id' | 'timer'>) => {
  const message: Message = {
    id: generateId(),
    duration: 3000, // 默认3秒
    showClose: false,
    ...options,
  }

  messages.value.push(message)

  // 自动移除消息
  if (message.duration && message.duration > 0) {
    message.timer = setTimeout(() => {
      removeMessage(message.id)
    }, message.duration)
  }

  return message.id
}

// 移除消息
const removeMessage = (id: string) => {
  const index = messages.value.findIndex((msg) => msg.id === id)
  if (index > -1) {
    const message = messages.value[index]
    // 清除定时器
    if (message.timer) {
      clearTimeout(message.timer)
    }
    messages.value.splice(index, 1)
  }
}

// 清空所有消息
const clearMessages = () => {
  // 清除所有定时器
  messages.value.forEach((message) => {
    if (message.timer) {
      clearTimeout(message.timer)
    }
  })
  messages.value = []
}

export const useMessage = () => {
  const success = (content: string, duration: number = 3000) => {
    return addMessage({ type: 'success', content, duration })
  }

  const error = (content: string, duration: number = 4000) => {
    return addMessage({ type: 'error', content, duration })
  }

  const warning = (content: string, duration: number = 3500) => {
    return addMessage({ type: 'warning', content, duration })
  }

  const info = (content: string, duration: number = 3000) => {
    return addMessage({ type: 'info', content, duration })
  }

  return {
    messages: messages.value,
    success,
    error,
    warning,
    info,
    removeMessage,
    clearMessages,
  }
}

// 导出全局状态供组件使用
export { messages, removeMessage }
