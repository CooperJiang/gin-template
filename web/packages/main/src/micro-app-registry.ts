export interface MicroAppRegistryItem {
  name: string
  title: string
  activeRule: string
  entryDev: string
  entryProd: string
  showInMenu?: boolean
  menuOrder?: number
}

export const microAppRegistry: MicroAppRegistryItem[] = [
  {
    name: 'app',
    title: '首页',
    activeRule: '/app',
    entryDev: 'http://localhost:3001',
    entryProd: '/subapps/app/',
    showInMenu: true,
    menuOrder: 10,
  },
]

export function getMicroAppMenuLinks() {
  return microAppRegistry
    .filter((app) => app.showInMenu !== false)
    .sort((a, b) => (a.menuOrder ?? 999) - (b.menuOrder ?? 999))
    .map((app) => ({ href: app.activeRule, label: app.title }))
}
