export type MessageType = 'success' | 'error' | 'warning' | 'info'

export interface Message {
  id: string
  type: MessageType
  content: string
  duration?: number
  showClose?: boolean
}

export interface MessageContainerProps {
  maxCount?: number
}
