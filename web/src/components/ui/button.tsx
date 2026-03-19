import * as React from 'react'
import { Slot } from '@radix-ui/react-slot'
import { cva, type VariantProps } from 'class-variance-authority'
import { cn } from '@/lib/utils'

const buttonVariants = cva(
  'inline-flex items-center justify-center gap-2 whitespace-nowrap text-sm font-bold transition-all border-nb border-nb-border rounded-[var(--nb-radius)] nb-interactive disabled:pointer-events-none disabled:opacity-50 [&_svg]:pointer-events-none [&_svg]:size-4 [&_svg]:shrink-0 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-nb-primary',
  {
    variants: {
      variant: {
        default: 'bg-nb-primary text-[var(--nb-primary-text)] shadow-nb hover:bg-[var(--nb-primary-hover)]',
        destructive: 'bg-nb-red text-white shadow-nb',
        outline: 'bg-nb-surface text-nb-text shadow-nb-sm hover:bg-nb-bg-soft',
        secondary: 'bg-nb-surface-alt text-nb-text shadow-nb-sm',
        ghost: 'border-transparent shadow-none hover:bg-nb-bg-soft hover:border-nb-border hover:shadow-nb-sm',
        link: 'text-nb-text underline-offset-4 hover:underline border-transparent shadow-none',
      },
      size: {
        default: 'h-10 px-5 py-2',
        sm: 'h-8 px-3 text-xs',
        lg: 'h-12 px-8 text-base',
        icon: 'h-10 w-10',
      },
    },
    defaultVariants: {
      variant: 'default',
      size: 'default',
    },
  },
)

export interface ButtonProps
  extends React.ButtonHTMLAttributes<HTMLButtonElement>,
    VariantProps<typeof buttonVariants> {
  asChild?: boolean
}

const Button = React.forwardRef<HTMLButtonElement, ButtonProps>(
  ({ className, variant, size, asChild = false, ...props }, ref) => {
    const Comp = asChild ? Slot : 'button'
    return (
      <Comp className={cn(buttonVariants({ variant, size, className }))} ref={ref} {...props} />
    )
  },
)
Button.displayName = 'Button'

export { Button, buttonVariants }
