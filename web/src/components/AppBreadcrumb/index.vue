<template>
  <nav class="flex" aria-label="Breadcrumb">
    <ol class="flex items-center space-x-1">
      <li
        v-for="(item, index) in breadcrumbs"
        :key="item.path || item.name"
        class="flex items-center"
      >
        <!-- 首页图标 -->
        <HomeIcon v-if="index === 0" class="h-4 w-4 text-gray-400 mr-1" />

        <!-- 分隔符 -->
        <ChevronRightIcon v-if="index > 0" class="h-3 w-3 text-gray-300 mx-1" />

        <!-- 面包屑项 -->
        <div class="flex items-center">
          <!-- 可点击的链接 -->
          <router-link
            v-if="item.path && index < breadcrumbs.length - 1"
            :to="item.path"
            class="text-sm font-medium text-blue-600 hover:text-blue-800 transition-colors duration-150 px-1 py-0.5 rounded hover:bg-blue-50 cursor-pointer"
          >
            {{ item.name }}
          </router-link>

          <!-- 不可点击的文件夹类型 -->
          <span
            v-else-if="index < breadcrumbs.length - 1"
            class="text-sm font-medium text-gray-400 px-1 py-0.5 cursor-default select-none"
            :title="`${item.name} 是一个目录`"
          >
            {{ item.name }}
          </span>

          <!-- 当前页面（不可点击） -->
          <span v-else class="text-sm font-semibold text-gray-800 px-1 py-0.5 cursor-default">
            {{ item.name }}
          </span>
        </div>
      </li>
    </ol>
  </nav>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { HomeIcon, ChevronRightIcon } from '@heroicons/vue/24/outline'
import { navigationConfig, type MenuItem } from '../../layouts/config/navigation'

const route = useRoute()
const router = useRouter()

interface BreadcrumbItem {
  name: string
  path?: string
}

// 检查路由是否有重定向（说明是文件夹类型）
const hasRedirect = (path: string): boolean => {
  const routeRecord = router.getRoutes().find((r) => r.path === path)
  return !!routeRecord?.redirect
}

// 检查导航配置中是否有子菜单但没有href（文件夹类型）
const isNavigationFolder = (items: MenuItem[], targetPath: string): boolean => {
  for (const item of items) {
    // 如果当前项匹配目标路径
    if (item.href === targetPath) {
      return false // 有href说明是页面，不是文件夹
    }

    // 递归检查子菜单
    if (item.children) {
      // 检查当前项是否是目标路径的父级，且没有href
      const isParentFolder =
        !item.href &&
        item.children.some((child) => {
          return (
            targetPath.startsWith(child.href || '') ||
            (child.children &&
              child.children.some((grandchild) => targetPath.startsWith(grandchild.href || '')))
          )
        })

      if (isParentFolder) {
        return true
      }

      // 递归检查子项
      const result = isNavigationFolder(item.children, targetPath)
      if (result) return result
    }
  }
  return false
}

// 检查路径是否为文件夹类型（综合判断）
const isFolderType = (path: string): boolean => {
  // 1. 首先检查路由是否有重定向
  if (hasRedirect(path)) {
    return true
  }

  // 2. 检查导航配置
  return isNavigationFolder(navigationConfig, path)
}

// 递归查找路径对应的导航项，返回完整路径
const findNavigationPath = (
  items: MenuItem[],
  targetPath: string,
  currentPath: BreadcrumbItem[] = [],
): BreadcrumbItem[] | null => {
  for (const item of items) {
    const newPath = [...currentPath, { name: item.name, path: item.href }]

    if (item.href === targetPath) {
      return newPath
    }

    if (item.children) {
      const result = findNavigationPath(item.children, targetPath, newPath)
      if (result) {
        return result
      }
    }
  }
  return null
}

// 根据路径生成面包屑
const generateBreadcrumbsFromPath = (path: string): BreadcrumbItem[] => {
  const segments = path.split('/').filter(Boolean)
  const breadcrumbs: BreadcrumbItem[] = []

  // 特殊路由处理映射
  const routeNameMap: Record<string, string> = {
    dashboard: '概览',
    profile: '个人资料',
    demo: '页面管理',
    page1: '页面1',
    'page1-1': '页面1-1',
    'page1-2': '页面1-2',
    'page1-3': '页面1-3',
    'page1-1-1': '页面1-1-1',
    'page1-1-2': '页面1-1-2',
    'page1-2-1': '页面1-2-1',
    'page1-2-2': '页面1-2-2',
  }

  // 添加首页
  if (path !== '/dashboard') {
    breadcrumbs.push({ name: '概览', path: '/dashboard' })
  }

  let currentPath = ''
  for (let i = 0; i < segments.length; i++) {
    currentPath += '/' + segments[i]
    const segment = segments[i]

    // 如果是首页，直接添加
    if (currentPath === '/dashboard') {
      if (path === '/dashboard') {
        breadcrumbs.push({ name: '概览', path: currentPath })
      }
      continue
    }

    const name = routeNameMap[segment] || segment.charAt(0).toUpperCase() + segment.slice(1)

    // 检查当前路径是否有重定向或者是导航文件夹
    const isFolder = isFolderType(currentPath)

    breadcrumbs.push({
      name,
      path: isFolder ? undefined : currentPath, // 如果是文件夹类型，不提供路径（不可点击）
    })
  }

  return breadcrumbs
}

// 计算面包屑
const breadcrumbs = computed(() => {
  const currentPath = route.path

  // 首先尝试从导航配置中查找
  const navPath = findNavigationPath(navigationConfig, currentPath)
  if (navPath && navPath.length > 0) {
    // 检查每个路径项是否为文件夹类型
    return navPath.map((item, index) => {
      if (!item.path || index === navPath.length - 1) {
        return item
      }

      // 检查是否有重定向或者是导航文件夹
      const isFolder = isFolderType(item.path)

      return {
        ...item,
        path: isFolder ? undefined : item.path,
      }
    })
  }

  // 如果导航配置中没有，则根据路径生成
  return generateBreadcrumbsFromPath(currentPath)
})

defineOptions({
  name: 'AppBreadcrumb',
})
</script>

<style scoped>
/* 面包屑悬停效果 */
.router-link-active {
  @apply text-blue-700 bg-blue-100;
}

/* 响应式设计 */
@media (max-width: 640px) {
  nav {
    @apply text-xs;
  }

  .space-x-1 > :not([hidden]) ~ :not([hidden]) {
    margin-left: 0.125rem;
  }

  .mx-1 {
    margin-left: 0.125rem;
    margin-right: 0.125rem;
  }

  .mr-1 {
    margin-right: 0.125rem;
  }
}

/* 面包屑项的悬停效果 */
nav a:hover {
  @apply transform scale-105;
}

/* 确保不可点击项目没有指针样式 */
.cursor-default {
  cursor: default !important;
}

/* 可点击项目的指针样式 */
.cursor-pointer {
  cursor: pointer !important;
}
</style>
