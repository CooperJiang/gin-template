<template>
  <div class="min-h-screen bg-gray-50 py-12">
    <div class="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="bg-white shadow rounded-lg">
        <div class="px-6 py-4 border-b border-gray-200">
          <h1 class="text-2xl font-bold text-gray-900">设置</h1>
        </div>
        <div class="p-6">
          <div class="space-y-8">
            <!-- 账户设置 -->
            <div>
              <h3 class="text-lg font-medium text-gray-900 mb-4">账户设置</h3>
              <div class="space-y-4">
                <div class="flex items-center justify-between">
                  <div>
                    <h4 class="text-sm font-medium text-gray-900">修改密码</h4>
                    <p class="text-sm text-gray-500">更新您的登录密码</p>
                  </div>
                  <button class="px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50">
                    修改
                  </button>
                </div>
                <div class="flex items-center justify-between">
                  <div>
                    <h4 class="text-sm font-medium text-gray-900">修改邮箱</h4>
                    <p class="text-sm text-gray-500">更新您的邮箱地址</p>
                  </div>
                  <button class="px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50">
                    修改
                  </button>
                </div>
              </div>
            </div>

            <!-- 通知设置 -->
            <div>
              <h3 class="text-lg font-medium text-gray-900 mb-4">通知设置</h3>
              <div class="space-y-4">
                <div class="flex items-center justify-between">
                  <div>
                    <h4 class="text-sm font-medium text-gray-900">邮件通知</h4>
                    <p class="text-sm text-gray-500">接收重要更新的邮件通知</p>
                  </div>
                  <button class="relative inline-flex h-6 w-11 flex-shrink-0 cursor-pointer rounded-full border-2 border-transparent bg-blue-600 transition-colors duration-200 ease-in-out focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2">
                    <span class="translate-x-5 inline-block h-5 w-5 transform rounded-full bg-white shadow ring-0 transition duration-200 ease-in-out"></span>
                  </button>
                </div>
              </div>
            </div>

            <!-- 隐私设置 -->
            <div>
              <h3 class="text-lg font-medium text-gray-900 mb-4">隐私设置</h3>
              <div class="space-y-4">
                <div class="flex items-center justify-between">
                  <div>
                    <h4 class="text-sm font-medium text-gray-900">个人资料可见性</h4>
                    <p class="text-sm text-gray-500">控制其他用户是否可以查看您的个人资料</p>
                  </div>
                  <select class="block w-32 px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm">
                    <option>公开</option>
                    <option>私密</option>
                  </select>
                </div>
              </div>
            </div>
          </div>

          <div class="mt-8 flex space-x-4">
            <button
              type="button"
              class="px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700"
            >
              保存设置
            </button>
            <router-link
              to="/"
              class="px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50"
            >
              返回首页
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// TODO: 实现设置逻辑
</script>
