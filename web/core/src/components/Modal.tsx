import { ReactNode, useEffect } from 'react'
import { Button } from './Button'

export interface ModalProps {
  /**
   * 是否打开
   */
  open: boolean
  /**
   * 关闭回调
   */
  onClose: () => void
  /**
   * 标题
   */
  title?: string
  /**
   * 内容
   */
  children: ReactNode
  /**
   * 底部操作
   */
  footer?: ReactNode
  /**
   * 宽度
   */
  width?: 'sm' | 'md' | 'lg' | 'xl'
  /**
   * 点击遮罩是否关闭
   */
  closeOnOverlayClick?: boolean
}

const widthClasses = {
  sm: 'max-w-sm',
  md: 'max-w-md',
  lg: 'max-w-lg',
  xl: 'max-w-xl',
}

export function Modal({
  open,
  onClose,
  title,
  children,
  footer,
  width = 'md',
  closeOnOverlayClick = true,
}: ModalProps) {
  useEffect(() => {
    if (open) {
      document.body.style.overflow = 'hidden'
    } else {
      document.body.style.overflow = 'unset'
    }
    return () => {
      document.body.style.overflow = 'unset'
    }
  }, [open])

  if (!open) return null

  return (
    <div
      className="fixed inset-0 z-50 flex items-center justify-center bg-app-overlay"
      onClick={(e) => {
        if (closeOnOverlayClick && e.target === e.currentTarget) {
          onClose()
        }
      }}
    >
      <div className={`bg-surface rounded-lg shadow-xl w-full mx-4 ${widthClasses[width]}`}>
        {/* Header */}
        {title && (
          <div className="flex items-center justify-between px-6 py-4 border-b border-app-border">
            <h3 className="text-lg font-semibold text-app-heading">{title}</h3>
            <button
              onClick={onClose}
              className="text-app-text-muted hover:text-app-text transition-colors"
            >
              <svg
                className="h-5 w-5"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M6 18L18 6M6 6l12 12"
                />
              </svg>
            </button>
          </div>
        )}

        {/* Body */}
        <div className="px-6 py-4">{children}</div>

        {/* Footer */}
        {footer && (
          <div className="flex items-center justify-end gap-3 px-6 py-4 border-t border-app-border">
            {footer}
          </div>
        )}
      </div>
    </div>
  )
}

export interface ConfirmModalProps {
  open: boolean
  onClose: () => void
  onConfirm: () => void | Promise<void>
  title: string
  description: string
  confirmText?: string
  cancelText?: string
  confirmVariant?: 'primary' | 'danger'
  loading?: boolean
}

export function ConfirmModal({
  open,
  onClose,
  onConfirm,
  title,
  description,
  confirmText = '确认',
  cancelText = '取消',
  confirmVariant = 'primary',
  loading = false,
}: ConfirmModalProps) {
  return (
    <Modal
      open={open}
      onClose={onClose}
      title={title}
      width="sm"
      footer={
        <>
          <Button variant="ghost" onClick={onClose} disabled={loading}>
            {cancelText}
          </Button>
          <Button variant={confirmVariant} onClick={onConfirm} loading={loading}>
            {confirmText}
          </Button>
        </>
      }
    >
      <p className="text-app-text">{description}</p>
    </Modal>
  )
}
