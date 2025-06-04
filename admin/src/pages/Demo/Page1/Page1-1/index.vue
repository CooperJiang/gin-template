<template>
  <div class="p-6">
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">页面1-1</h1>
      <p class="mt-2 text-gray-600">这是页面1的子页面1-1</p>
    </div>

    <div class="bg-white rounded-lg shadow p-6">
      <h2 class="text-lg font-semibold mb-4">页面内容</h2>
      <p class="text-gray-700">这是页面1-1的内容。您可以继续导航到更深层的子页面：</p>
      <ul class="mt-4 space-y-2">
        <li>
          <router-link to="/demo/page1/page1-1/page1-1-1" class="text-blue-600 hover:text-blue-800">
            → 页面1-1-1
          </router-link>
        </li>
        <li>
          <router-link to="/demo/page1/page1-1/page1-1-2" class="text-blue-600 hover:text-blue-800">
            → 页面1-1-2
          </router-link>
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup lang="ts">
// 页面1-1组件
</script>
