<script setup lang="ts">
import { onMounted, computed } from 'vue'
import { useRoute } from 'vue-router'
import { useAuth } from './hooks/user/useAuth'
import AdminLayout from './layouts/AdminLayout.vue'
import MessageContainer from './components/MessageContainer'

const route = useRoute()
const { getUserInfo, isAuthenticated } = useAuth()

// 根据路由meta决定使用哪个布局
const layout = computed(() => {
  return route.meta.layout || 'default'
})

onMounted(async () => {
  // 应用启动时初始化认证状态
  if (isAuthenticated.value) {
    try {
      await getUserInfo()
    } catch (error) {
      console.error('Failed to get user info:', error)
    }
  }
})
</script>

<template>
  <div id="app">
    <!-- 管理员布局 -->
    <AdminLayout v-if="layout === 'admin'">
      <router-view />
    </AdminLayout>

    <!-- 默认布局（登录、注册等页面） -->
    <router-view v-else />

    <!-- 全局消息提示容器 -->
    <MessageContainer />
  </div>
</template>

<style>
/* 全局样式已在 main.css 中定义 */
#app {
  width: 100%;
  min-height: 100vh;
}
</style>
