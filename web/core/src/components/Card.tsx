import { HTMLAttributes, ReactNode } from 'react'

export interface CardProps extends HTMLAttributes<HTMLDivElement> {
  /**
   * 卡片标题
   */
  title?: string
  /**
   * 卡片额外操作
   */
  extra?: ReactNode
  /**
   * 是否有边框
   */
  bordered?: boolean
  /**
   * 是否可悬停
   */
  hoverable?: boolean
}

export function Card({
  title,
  extra,
  bordered = true,
  hoverable = false,
  className = '',
  children,
  ...props
}: CardProps) {
  const classes = [
    'rounded-lg',
    'bg-surface',
    bordered ? 'border border-app-border' : '',
    hoverable ? 'hover:shadow-md transition-shadow duration-200' : '',
    className,
  ]
    .filter(Boolean)
    .join(' ')

  return (
    <div className={classes} {...props}>
      {(title || extra) && (
        <div className="flex items-center justify-between px-6 py-4 border-b border-app-border">
          {title && <h3 className="text-lg font-semibold text-app-heading">{title}</h3>}
          {extra && <div>{extra}</div>}
        </div>
      )}
      <div className="p-6">{children}</div>
    </div>
  )
}

export interface CardSectionProps extends HTMLAttributes<HTMLDivElement> {
  title?: string
}

export function CardSection({ title, className = '', children, ...props }: CardSectionProps) {
  return (
    <div className={`py-4 ${className}`} {...props}>
      {title && <h4 className="text-sm font-medium text-app-text-secondary mb-2">{title}</h4>}
      {children}
    </div>
  )
}
