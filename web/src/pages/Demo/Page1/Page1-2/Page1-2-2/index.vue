<template>
  <div class="p-6">
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h1 class="text-2xl font-bold text-gray-900 mb-4">页面1-2-2</h1>
      <p class="text-gray-600 mb-4">这是页面1-2-2的内容，属于页面1-2的子页面。</p>

      <div class="bg-purple-50 border border-purple-200 rounded-lg p-4">
        <h2 class="text-lg font-semibold text-purple-900 mb-2">页面信息</h2>
        <ul class="text-purple-800 space-y-1">
          <li>• 当前路径: /demo/page1/page1-2/page1-2-2</li>
          <li>• 页面类型: 子页面</li>
          <li>• 父页面: 页面1-2</li>
          <li>• 面包屑: 概览 > 页面管理 > 页面1 > 页面1-2 > 页面1-2-2</li>
        </ul>
      </div>

      <div class="mt-6">
        <h3 class="text-lg font-medium text-gray-900 mb-3">导航测试</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <router-link
            to="/demo/page1/page1-2/page1-2-1"
            class="block p-4 bg-gray-50 hover:bg-gray-100 rounded-lg border border-gray-200 transition-colors"
          >
            <div class="font-medium text-gray-900">页面1-2-1</div>
            <div class="text-sm text-gray-600">跳转到兄弟页面</div>
          </router-link>

          <router-link
            to="/demo/page1/page1-2"
            class="block p-4 bg-gray-50 hover:bg-gray-100 rounded-lg border border-gray-200 transition-colors"
          >
            <div class="font-medium text-gray-900">页面1-2</div>
            <div class="text-sm text-gray-600">返回父页面</div>
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// 页面1-2-2组件
</script>
