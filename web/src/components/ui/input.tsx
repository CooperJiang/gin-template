import * as React from 'react'
import { cn } from '@/lib/utils'

const Input = React.forwardRef<HTMLInputElement, React.InputHTMLAttributes<HTMLInputElement>>(
  ({ className, type, ...props }, ref) => {
    return (
      <input
        type={type}
        className={cn(
          'flex h-10 w-full border-nb border-nb-border bg-nb-surface px-3 py-2 text-sm text-nb-text shadow-nb-sm rounded-[var(--nb-radius)] placeholder:text-nb-text-muted focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-nb-primary focus-visible:shadow-nb disabled:cursor-not-allowed disabled:opacity-50 transition-shadow file:border-0 file:bg-transparent file:text-sm file:font-bold',
          className,
        )}
        ref={ref}
        {...props}
      />
    )
  },
)
Input.displayName = 'Input'

export { Input }
