import { createRouter, createWebHistory } from 'vue-router'
import { useAuth } from '../hooks/user/useAuth'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: '/dashboard',
    },
    // 游客页面 - 未登录用户可访问，已登录用户会被重定向
    {
      path: '/login',
      name: 'login',
      component: () => import('@/pages/Login/index.vue'),
      meta: {
        requiresGuest: true,
        title: '登录',
      },
    },
    {
      path: '/register',
      name: 'register',
      component: () => import('@/pages/Register/index.vue'),
      meta: {
        requiresGuest: true,
        title: '注册',
      },
    },
    {
      path: '/forgot-password',
      name: 'forgot-password',
      component: () => import('@/pages/ForgotPassword/index.vue'),
      meta: {
        requiresGuest: true,
        title: '忘记密码',
      },
    },
    // 需要认证的页面
    {
      path: '/dashboard',
      name: 'dashboard',
      component: () => import('@/pages/Dashboard/index.vue'),
      meta: {
        requiresAuth: true,
        layout: 'admin',
        breadcrumb: '概览',
      },
    },
    {
      path: '/debug',
      name: 'debug',
      component: () => import('@/pages/Debug/index.vue'),
      meta: {
        requiresAuth: true,
        layout: 'admin',
        breadcrumb: 'Debug调试',
      },
    },
    {
      path: '/element-table-test',
      name: 'element-table-test',
      component: () => import('@/pages/Debug/components/ElementTableTest.vue'),
      meta: {
        requiresAuth: true,
        layout: 'admin',
        breadcrumb: 'Element表格测试',
      },
    },
    {
      path: '/simple-table-test',
      name: 'simple-table-test',
      component: () => import('@/pages/Debug/components/SimpleTableTest.vue'),
      meta: {
        requiresAuth: true,
        layout: 'admin',
        breadcrumb: '简单表格测试',
      },
    },
    {
      path: '/demo/page1',
      name: 'demo-page1',
      component: () => import('@/pages/Demo/Page1/index.vue'),
      meta: {
        requiresAuth: true,
        layout: 'admin',
        breadcrumb: '页面1',
      },
    },
    {
      path: '/demo/page1/page1-1',
      redirect: '/demo/page1/page1-1/page1-1-1',
    },
    {
      path: '/demo/page1/page1-2',
      redirect: '/demo/page1/page1-2/page1-2-1',
    },
    {
      path: '/demo/page1/page1-3',
      name: 'demo-page1-3',
      component: () => import('@/pages/Demo/Page1/Page1-3/index.vue'),
      meta: {
        requiresAuth: true,
        layout: 'admin',
        breadcrumb: '页面1-3',
      },
    },
    {
      path: '/demo/page1/page1-2/page1-2-1',
      name: 'demo-page1-2-1',
      component: () => import('@/pages/Demo/Page1/Page1-2/Page1-2-1/index.vue'),
      meta: {
        requiresAuth: true,
        layout: 'admin',
        breadcrumb: '页面1-2-1',
      },
    },
    {
      path: '/demo/page1/page1-2/page1-2-2',
      name: 'demo-page1-2-2',
      component: () => import('@/pages/Demo/Page1/Page1-2/Page1-2-2/index.vue'),
      meta: {
        requiresAuth: true,
        layout: 'admin',
        breadcrumb: '页面1-2-2',
      },
    },
    {
      path: '/demo/page1/page1-1/page1-1-1',
      name: 'demo-page1-1-1',
      component: () => import('@/pages/Demo/Page1/Page1-1/Page1-1-1/index.vue'),
      meta: {
        requiresAuth: true,
        layout: 'admin',
        breadcrumb: '页面1-1-1',
      },
    },
    {
      path: '/demo/page1/page1-1/page1-1-2',
      name: 'demo-page1-1-2',
      component: () => import('@/pages/Demo/Page1/Page1-1/Page1-1-2/index.vue'),
      meta: {
        requiresAuth: true,
        layout: 'admin',
        breadcrumb: '页面1-1-2',
      },
    },
    {
      path: '/profile',
      name: 'profile',
      component: () => import('@/pages/Profile/index.vue'),
      meta: {
        requiresAuth: true,
        layout: 'admin',
        breadcrumb: '个人资料',
      },
    },
    // 公共页面 - 任何人都可以访问
    {
      path: '/:pathMatch(.*)*',
      name: 'not-found',
      component: () => import('@/pages/NotFound/index.vue'),
      meta: {
        title: '页面未找到',
      },
    },
  ],
})

// 路由守卫
router.beforeEach(async (to, from, next) => {
  const { isAuthenticated, getUserInfo, token, logout } = useAuth()

  // 如果有token但没有用户信息，先获取用户信息
  // 但如果是从登录页跳转过来的，跳过这个检查（因为登录时已经设置了用户信息）
  if (token.value && !isAuthenticated.value && from.name !== 'login') {
    try {
      await getUserInfo()
    } catch (error) {
      console.error('Auth initialization failed:', error)
      // 获取用户信息失败，清除可能已过期的token和用户信息
      logout(false)
      // 如果当前要访问的是需要认证的页面，则跳转到登录页
      if (to.meta.requiresAuth) {
        next({
          name: 'login',
          query: { redirect: to.fullPath },
          replace: true,
        })
        return
      }
    }
  }

  // 检查是否需要认证的页面
  if (to.meta.requiresAuth && !isAuthenticated.value) {
    // 记录用户想要访问的页面，登录后跳转回来
    next({
      name: 'login',
      query: { redirect: to.fullPath },
      replace: true,
    })
    return
  }

  // 检查是否是游客页面（登录、注册等）
  if (to.meta.requiresGuest && isAuthenticated.value) {
    // 如果已登录用户访问登录页，重定向到来源页面或首页
    const redirectPath = (from.query?.redirect as string) || '/dashboard'
    next({
      path: redirectPath,
      replace: true,
    })
    return
  }

  // 404页面不需要认证，任何人都可以访问
  if (to.name === 'not-found') {
    next()
    return
  }

  next()
})

// 路由后置守卫 - 用于设置页面标题等
router.afterEach((to) => {
  // 设置页面标题
  const defaultTitle = '管理后台'
  const pageTitle = to.meta.breadcrumb as string

  if (pageTitle) {
    document.title = `${pageTitle} - ${defaultTitle}`
  } else {
    document.title = defaultTitle
  }

  // 可以在这里添加其他全局逻辑，如页面访问统计等
})

export default router
