// PopoverMenu 组件的类型定义

export interface PopoverMenuProps {
  placement?: 'bottom-start' | 'bottom-end' | 'top-start' | 'top-end'
  trigger?: 'click' | 'hover'
  disabled?: boolean
  offset?: [number, number]
}
