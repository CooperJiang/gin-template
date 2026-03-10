export interface LoadingProps {
  /**
   * 加载文本
   */
  text?: string
  /**
   * 大小
   */
  size?: 'sm' | 'md' | 'lg'
  /**
   * 是否全屏
   */
  fullscreen?: boolean
}

const sizeClasses = {
  sm: 'h-4 w-4',
  md: 'h-8 w-8',
  lg: 'h-12 w-12',
}

export function Loading({ text = '加载中...', size = 'md', fullscreen = false }: LoadingProps) {
  const spinner = (
    <svg
      className={`animate-spin text-primary-500 ${sizeClasses[size]}`}
      xmlns="http://www.w3.org/2000/svg"
      fill="none"
      viewBox="0 0 24 24"
    >
      <circle
        className="opacity-25"
        cx="12"
        cy="12"
        r="10"
        stroke="currentColor"
        strokeWidth="4"
      />
      <path
        className="opacity-75"
        fill="currentColor"
        d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
      />
    </svg>
  )

  if (fullscreen) {
    return (
      <div className="fixed inset-0 flex items-center justify-center bg-app-overlay z-50">
        <div className="flex flex-col items-center gap-3 bg-surface px-8 py-6 rounded-lg shadow-lg">
          {spinner}
          {text && <p className="text-app-text">{text}</p>}
        </div>
      </div>
    )
  }

  return (
    <div className="flex flex-col items-center gap-2">
      {spinner}
      {text && <p className="text-sm text-app-text-muted">{text}</p>}
    </div>
  )
}
