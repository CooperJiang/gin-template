export type MessageType = 'success' | 'error' | 'warning' | 'info'

interface MessageOptions {
  type: MessageType
  content: string
  duration?: number
}

interface MessageItem extends MessageOptions {
  id: string
  timer?: ReturnType<typeof setTimeout>
  el: HTMLElement
}

const icons: Record<MessageType, string> = {
  success: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" width="24" height="24"><path stroke-linecap="round" stroke-linejoin="round" d="M9 12.75 11.25 15 15 9.75M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z" /></svg>`,
  error: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" width="24" height="24"><path stroke-linecap="round" stroke-linejoin="round" d="M12 9v3.75m9-.75a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 3.75h.008v.008H12v-.008Z" /></svg>`,
  warning: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" width="24" height="24"><path stroke-linecap="round" stroke-linejoin="round" d="M12 9v3.75m-9.303 3.376c-.866 1.5.217 3.374 1.948 3.374h14.71c1.73 0 2.813-1.874 1.948-3.374L13.949 3.378c-.866-1.5-3.032-1.5-3.898 0L2.697 16.126ZM12 15.75h.007v.008H12v-.008Z" /></svg>`,
  info: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" width="24" height="24"><path stroke-linecap="round" stroke-linejoin="round" d="m11.25 11.25.041-.02a.75.75 0 0 1 1.063.852l-.708 2.836a.75.75 0 0 0 1.063.853l.041-.021M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9-3.75h.008v.008H12V8.25Z" /></svg>`,
}

const typeStyles: Record<MessageType, string> = {
  success: 'background:#ecfdf5;color:#065f46;border-color:#a7f3d0;',
  error: 'background:#fef2f2;color:#991b1b;border-color:#fecaca;',
  warning: 'background:#fffbeb;color:#92400e;border-color:#fde68a;',
  info: 'background:#eff6ff;color:#1e40af;border-color:#bfdbfe;',
}

let messageId = 0
const messages: MessageItem[] = []
let container: HTMLElement | null = null
let styleInjected = false

function injectStyles() {
  if (styleInjected) return
  styleInjected = true
  const style = document.createElement('style')
  style.textContent = `
    .core-msg-container{position:fixed;top:24px;left:50%;transform:translateX(-50%);z-index:9999;pointer-events:none;display:flex;flex-direction:column;align-items:center;gap:12px}
    .core-msg-item{pointer-events:auto;display:flex;align-items:center;gap:12px;padding:12px 20px;border-radius:12px;border:1px solid;min-width:320px;max-width:480px;font-size:14px;font-weight:500;line-height:1.5;box-shadow:0 10px 25px -3px rgba(0,0,0,.1),0 4px 6px -2px rgba(0,0,0,.05);backdrop-filter:blur(8px);animation:core-msg-in .4s cubic-bezier(.175,.885,.32,1.275)}
    .core-msg-item.removing{animation:core-msg-out .3s ease-in forwards}
    .core-msg-icon{flex-shrink:0;width:24px;height:24px}
    .core-msg-content{flex:1}
    @keyframes core-msg-in{from{opacity:0;transform:translateY(-50px) scale(.9)}to{opacity:1;transform:translateY(0) scale(1)}}
    @keyframes core-msg-out{from{opacity:1;transform:translateY(0) scale(1)}to{opacity:0;transform:translateY(-30px) scale(.95)}}
  `
  document.head.appendChild(style)
}

function getContainer(): HTMLElement {
  if (container && document.body.contains(container)) return container
  injectStyles()
  container = document.createElement('div')
  container.className = 'core-msg-container'
  document.body.appendChild(container)
  return container
}

function removeMessage(id: string) {
  const index = messages.findIndex((m) => m.id === id)
  if (index === -1) return
  const msg = messages[index]
  if (msg.timer) clearTimeout(msg.timer)
  msg.el.classList.add('removing')
  msg.el.addEventListener('animationend', () => {
    msg.el.remove()
    if (container && container.children.length === 0) {
      container.remove()
      container = null
    }
  })
  messages.splice(index, 1)
}

function addMessage(options: MessageOptions): string {
  const id = `core_msg_${++messageId}_${Date.now()}`
  const duration = options.duration ?? 3000
  const c = getContainer()

  const el = document.createElement('div')
  el.className = 'core-msg-item'
  el.setAttribute('style', typeStyles[options.type])

  const iconEl = document.createElement('span')
  iconEl.className = 'core-msg-icon'
  iconEl.innerHTML = icons[options.type]
  el.appendChild(iconEl)

  const contentEl = document.createElement('span')
  contentEl.className = 'core-msg-content'
  contentEl.textContent = options.content
  el.appendChild(contentEl)

  c.appendChild(el)

  const item: MessageItem = { ...options, id, el }
  if (duration > 0) {
    item.timer = setTimeout(() => removeMessage(id), duration)
  }
  messages.push(item)

  return id
}

export const message = {
  success: (content: string, duration = 3000) => addMessage({ type: 'success', content, duration }),
  error: (content: string, duration = 4000) => addMessage({ type: 'error', content, duration }),
  warning: (content: string, duration = 3500) =>
    addMessage({ type: 'warning', content, duration }),
  info: (content: string, duration = 3000) => addMessage({ type: 'info', content, duration }),
  remove: removeMessage,
}
