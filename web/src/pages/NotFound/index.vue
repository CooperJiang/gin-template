<template>
  <div
    class="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-50 to-indigo-100"
  >
    <div class="max-w-lg w-full text-center px-6">
      <!-- 404图标动画 -->
      <div class="mb-8 animate-bounce">
        <div
          class="mx-auto w-32 h-32 bg-gradient-to-r from-blue-500 to-purple-600 rounded-full flex items-center justify-center mb-6 shadow-lg"
        >
          <svg class="w-16 h-16 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M9.172 16.172a4 4 0 015.656 0M9 12h6m-6-4h6m2 5.291A7.962 7.962 0 0112 15c-2.87 0-5.407 1.516-6.834 3.791M6 15.813A7.963 7.963 0 014 10c0-4.411 3.589-8 8-8s8 3.589 8 8c0 2.192-.893 4.165-2.332 5.581"
            />
          </svg>
        </div>
        <h1
          class="text-8xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-blue-600 to-purple-600 animate-pulse"
        >
          404
        </h1>
      </div>

      <!-- 错误信息 -->
      <div class="mb-8 animate-fade-in-up">
        <h2 class="text-3xl font-bold text-gray-800 mb-4">页面走丢了...</h2>
        <p class="text-gray-600 text-lg mb-6">抱歉，您访问的页面不存在或已被移除。</p>
        <p class="text-sm text-gray-500 mb-4">
          将在 <span class="font-bold text-blue-600">{{ countdown }}</span> 秒后自动返回首页
        </p>

        <!-- 倒计时进度条 -->
        <div class="w-full bg-gray-200 rounded-full h-2 mb-6">
          <div
            class="bg-gradient-to-r from-blue-500 to-purple-500 h-2 rounded-full transition-all duration-1000 ease-linear"
            :style="{ width: `${(5 - countdown) * 20}%` }"
          ></div>
        </div>
      </div>

      <!-- 操作按钮 -->
      <div class="space-y-4 animate-fade-in-up-delay">
        <router-link
          to="/dashboard"
          class="inline-flex items-center justify-center w-full px-6 py-3 bg-gradient-to-r from-blue-500 to-purple-600 text-white font-semibold rounded-lg shadow-lg hover:shadow-xl transform hover:scale-105 transition-all duration-200"
        >
          <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"
            />
          </svg>
          立即返回首页
        </router-link>

        <button
          @click="goBack"
          class="inline-flex items-center justify-center w-full px-6 py-3 bg-white text-gray-700 font-semibold rounded-lg border-2 border-gray-300 hover:border-gray-400 hover:shadow-md transform hover:scale-105 transition-all duration-200"
        >
          <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M10 19l-7-7m0 0l7-7m-7 7h18"
            />
          </svg>
          返回上一页
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const countdown = ref(5)
let timer: number | null = null

onMounted(() => {
  // 启动倒计时
  timer = window.setInterval(() => {
    countdown.value--

    if (countdown.value <= 0) {
      // 倒计时结束，跳转到首页
      router.push('/dashboard')
    }
  }, 1000)
})

onUnmounted(() => {
  // 清理定时器
  if (timer) {
    clearInterval(timer)
    timer = null
  }
})

const goBack = () => {
  router.go(-1)
}

defineOptions({
  name: 'NotFoundPage',
})
</script>

<style scoped>
@keyframes fadeInUp {
  0% {
    opacity: 0;
    transform: translateY(30px);
  }
  100% {
    opacity: 1;
    transform: translateY(0);
  }
}

.animate-fade-in-up {
  animation: fadeInUp 0.6s ease-out;
}

.animate-fade-in-up-delay {
  animation: fadeInUp 0.6s ease-out 0.3s both;
}
</style>
