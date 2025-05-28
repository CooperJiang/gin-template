<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50">
    <div class="max-w-md w-full text-center">
      <div class="mb-8">
        <h1 class="text-9xl font-bold text-gray-300">404</h1>
        <h2 class="text-2xl font-semibold text-gray-900 mb-4">页面未找到</h2>
        <p class="text-gray-600 mb-8">
          抱歉，您访问的页面不存在或已被移除。
        </p>
      </div>

      <div class="space-y-4">
        <router-link
          to="/dashboard"
          class="btn btn-primary w-full"
        >
          返回首页
        </router-link>

        <button
          @click="$router.go(-1)"
          class="btn btn-secondary w-full"
        >
          返回上一页
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// 404页面不需要额外逻辑
</script>
