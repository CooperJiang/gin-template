import { InputHTMLAttributes, forwardRef } from 'react'

export interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  /**
   * 标签文本
   */
  label?: string
  /**
   * 错误信息
   */
  error?: string
  /**
   * 帮助文本
   */
  helperText?: string
  /**
   * 是否全宽
   */
  fullWidth?: boolean
}

export const Input = forwardRef<HTMLInputElement, InputProps>(
  (
    { label, error, helperText, fullWidth = false, className = '', ...props },
    ref
  ) => {
    const inputClasses = [
      'block px-3 py-2 rounded-md',
      'border border-app-input-border',
      'bg-app-input-bg text-app-input-text',
      'placeholder:text-app-input-placeholder',
      'focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent',
      'transition-colors duration-200',
      'disabled:opacity-50 disabled:cursor-not-allowed',
      error ? 'border-error focus:ring-error' : '',
      fullWidth ? 'w-full' : '',
      className,
    ]
      .filter(Boolean)
      .join(' ')

    return (
      <div className={fullWidth ? 'w-full' : ''}>
        {label && (
          <label className="block text-sm font-medium text-app-text mb-1">
            {label}
            {props.required && <span className="text-error ml-1">*</span>}
          </label>
        )}
        <input ref={ref} className={inputClasses} {...props} />
        {error && <p className="mt-1 text-sm text-error">{error}</p>}
        {helperText && !error && (
          <p className="mt-1 text-sm text-app-text-muted">{helperText}</p>
        )}
      </div>
    )
  }
)

Input.displayName = 'Input'
