<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50">
    <div class="text-center">
      <h1 class="text-9xl font-bold text-gray-300">404</h1>
      <h2 class="text-2xl font-bold text-gray-900 mt-4 mb-2">页面未找到</h2>
      <p class="text-gray-600 mb-8">抱歉，您访问的页面不存在。</p>
      <router-link
        to="/"
        class="inline-flex items-center px-6 py-3 border border-transparent text-base font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700"
      >
        返回首页
      </router-link>
    </div>
  </div>
</template>

<script setup lang="ts">
// 静态页面，无需额外逻辑
</script>
