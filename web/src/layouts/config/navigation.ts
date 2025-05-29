import {
  HomeIcon,
  FolderIcon,
  DocumentIcon,
  UserIcon,
  WrenchScrewdriverIcon,
} from '@heroicons/vue/24/outline'

export interface MenuItem {
  name: string
  href?: string
  icon: any
  children?: MenuItem[]
}

export const navigationConfig: MenuItem[] = [
  {
    name: '概览',
    href: '/dashboard',
    icon: HomeIcon,
  },
  {
    name: 'Debug调试',
    href: '/debug',
    icon: WrenchScrewdriverIcon,
  },
  {
    name: '页面管理',
    icon: FolderIcon,
    children: [
      {
        name: '页面1',
        href: '/demo/page1',
        icon: DocumentIcon,
        children: [
          {
            name: '页面1-1',
            icon: FolderIcon,
            children: [
              {
                name: '页面1-1-1',
                href: '/demo/page1/page1-1/page1-1-1',
                icon: DocumentIcon,
              },
              {
                name: '页面1-1-2',
                href: '/demo/page1/page1-1/page1-1-2',
                icon: DocumentIcon,
              },
            ],
          },
          {
            name: '页面1-2',
            icon: FolderIcon,
            children: [
              {
                name: '页面1-2-1',
                href: '/demo/page1/page1-2/page1-2-1',
                icon: DocumentIcon,
              },
              {
                name: '页面1-2-2',
                href: '/demo/page1/page1-2/page1-2-2',
                icon: DocumentIcon,
              },
            ],
          },
          {
            name: '页面1-3',
            href: '/demo/page1/page1-3',
            icon: DocumentIcon,
          },
        ],
      },
    ],
  },
  {
    name: '个人资料',
    href: '/profile',
    icon: UserIcon,
  },
]
