<template>
  <div class="min-h-screen bg-gray-50">
    <div class="max-w-4xl mx-auto py-12 px-4 sm:px-6 lg:px-8">
      <div class="text-center mb-12">
        <h1 class="text-4xl font-bold text-gray-900 mb-4">关于我们</h1>
        <p class="text-xl text-gray-600">
          了解更多关于这个用户端模板的信息
        </p>
      </div>

      <div class="bg-white rounded-lg shadow-md p-8 mb-8">
        <h2 class="text-2xl font-bold text-gray-900 mb-6">项目介绍</h2>
        <div class="prose max-w-none">
          <p class="text-gray-600 mb-4">
            这是一个基于现代前端技术栈构建的用户端模板，旨在为开发者提供一个快速启动项目的基础框架。
          </p>
          <p class="text-gray-600 mb-4">
            该模板采用了最新的技术和最佳实践，确保代码的可维护性、可扩展性和性能。
          </p>
        </div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-8 mb-8">
        <div class="bg-white rounded-lg shadow-md p-6">
          <h3 class="text-xl font-bold text-gray-900 mb-4">技术栈</h3>
          <ul class="space-y-2 text-gray-600">
            <li>• Vue 3 - 渐进式 JavaScript 框架</li>
            <li>• TypeScript - 类型安全的 JavaScript</li>
            <li>• TailwindCSS - 实用优先的 CSS 框架</li>
            <li>• Vue Router - 官方路由管理器</li>
            <li>• Pinia - 状态管理库</li>
            <li>• Vite - 快速的构建工具</li>
          </ul>
        </div>

        <div class="bg-white rounded-lg shadow-md p-6">
          <h3 class="text-xl font-bold text-gray-900 mb-4">特性</h3>
          <ul class="space-y-2 text-gray-600">
            <li>• 响应式设计</li>
            <li>• 用户认证系统</li>
            <li>• 安全的数据存储</li>
            <li>• 现代化的 UI 组件</li>
            <li>• 代码规范和格式化</li>
            <li>• 开发工具集成</li>
          </ul>
        </div>
      </div>

      <div class="text-center">
        <router-link
          to="/"
          class="inline-flex items-center px-6 py-3 border border-transparent text-base font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700"
        >
          返回首页
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// 静态页面，无需额外逻辑
</script>
