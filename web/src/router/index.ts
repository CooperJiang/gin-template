import { createRouter, createWebHistory } from 'vue-router'
import { useAuth } from '../hooks/user/useAuth'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: '/dashboard'
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/pages/Login/index.vue'),
      meta: { requiresGuest: true }
    },
    {
      path: '/register',
      name: 'register',
      component: () => import('@/pages/Register/index.vue'),
      meta: { requiresGuest: true }
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: () => import('@/pages/Dashboard/index.vue'),
      meta: { requiresAuth: true, layout: 'admin' }
    },
    {
      path: '/demo/page1',
      name: 'demo-page1',
      component: () => import('@/pages/Demo/Page1/index.vue'),
      meta: { requiresAuth: true, layout: 'admin' }
    },
    {
      path: '/demo/page1/page1-1',
      name: 'demo-page1-1',
      component: () => import('@/pages/Demo/Page1/Page1-1/index.vue'),
      meta: { requiresAuth: true, layout: 'admin' }
    },
    {
      path: '/demo/page1/page1-2',
      name: 'demo-page1-2',
      component: () => import('@/pages/Demo/Page1/Page1-2/index.vue'),
      meta: { requiresAuth: true, layout: 'admin' }
    },
    {
      path: '/demo/page1/page1-1/page1-1-1',
      name: 'demo-page1-1-1',
      component: () => import('@/pages/Demo/Page1/Page1-1/Page1-1-1/index.vue'),
      meta: { requiresAuth: true, layout: 'admin' }
    },
    {
      path: '/demo/page1/page1-1/page1-1-2',
      name: 'demo-page1-1-2',
      component: () => import('@/pages/Demo/Page1/Page1-1/Page1-1-2/index.vue'),
      meta: { requiresAuth: true, layout: 'admin' }
    },
    {
      path: '/profile',
      name: 'profile',
      component: () => import('@/pages/Profile/index.vue'),
      meta: { requiresAuth: true, layout: 'admin' }
    },
    {
      path: '/forgot-password',
      name: 'forgot-password',
      component: () => import('@/pages/ForgotPassword/index.vue'),
      meta: { requiresGuest: true }
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'not-found',
      component: () => import('@/pages/NotFound/index.vue')
    }
  ]
})

// 路由守卫
router.beforeEach(async (to, from, next) => {
  const { isAuthenticated, getUserInfo, token } = useAuth()

  // 如果有token但没有用户信息，先获取用户信息
  if (token.value && !isAuthenticated.value) {
    try {
      await getUserInfo()
    } catch (error) {
      console.error('Auth initialization failed:', error)
      // 获取用户信息失败，可能token已过期
    }
  }

  // 检查是否需要认证
  if (to.meta.requiresAuth && !isAuthenticated.value) {
    next({ name: 'login', query: { redirect: to.fullPath } })
    return
  }

  // 检查是否需要游客状态（已登录用户不能访问登录页等）
  if (to.meta.requiresGuest && isAuthenticated.value) {
    next({ name: 'dashboard' })
    return
  }

  next()
})

export default router
