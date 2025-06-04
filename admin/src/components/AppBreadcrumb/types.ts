// AppBreadcrumb 组件的类型定义

export interface BreadcrumbItem {
  name: string
  href?: string
  current?: boolean
}

export interface AppBreadcrumbProps {
  items?: BreadcrumbItem[]
  separator?: string
  showHome?: boolean
}
