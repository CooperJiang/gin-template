<script setup lang="ts">
import { onMounted } from 'vue'
import { useAuth } from './hooks/user/useAuth'
import MessageContainer from './components/MessageContainer.vue'

const { getUserInfo, isAuthenticated } = useAuth()

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
    <!-- 主要内容 -->
    <router-view />

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
