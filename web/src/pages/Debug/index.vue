<template>
  <div>
    <!-- 固定导航栏 -->
    <div class="sticky top-0 z-10 bg-white border-b border-gray-200 mb-8">
      <div class="px-6 py-4">
        <!-- 导航菜单 -->
        <nav class="flex flex-wrap gap-2">
          <a
            href="#icon-test"
            class="inline-flex items-center px-3 py-2 text-sm font-medium text-blue-600 bg-blue-50 rounded-md hover:bg-blue-100 transition-colors"
            @click="scrollToSection('icon-test')"
          >
            <GlobalIcon name="star" size="sm" class="mr-2" />
            Icon 图标
          </a>
          <a
            href="#button-test"
            class="inline-flex items-center px-3 py-2 text-sm font-medium text-green-600 bg-green-50 rounded-md hover:bg-green-100 transition-colors"
            @click="scrollToSection('button-test')"
          >
            <GlobalIcon name="check" size="sm" class="mr-2" />
            Button 按钮
          </a>
          <a
            href="#upload-test"
            class="inline-flex items-center px-3 py-2 text-sm font-medium text-indigo-600 bg-indigo-50 rounded-md hover:bg-indigo-100 transition-colors"
            @click="scrollToSection('upload-test')"
          >
            <GlobalIcon name="upload" size="sm" class="mr-2" />
            Upload 上传
          </a>
          <a
            href="#message-test"
            class="inline-flex items-center px-3 py-2 text-sm font-medium text-purple-600 bg-purple-50 rounded-md hover:bg-purple-100 transition-colors"
            @click="scrollToSection('message-test')"
          >
            <GlobalIcon name="chat-bubble-left" size="sm" class="mr-2" />
            Message 消息
          </a>
          <a
            href="#storage-test"
            class="inline-flex items-center px-3 py-2 text-sm font-medium text-orange-600 bg-orange-50 rounded-md hover:bg-orange-100 transition-colors"
            @click="scrollToSection('storage-test')"
          >
            <GlobalIcon name="archive-box" size="sm" class="mr-2" />
            Storage 存储
          </a>
          <a
            href="#developer-info"
            class="inline-flex items-center px-3 py-2 text-sm font-medium text-gray-600 bg-gray-50 rounded-md hover:bg-gray-100 transition-colors"
            @click="scrollToSection('developer-info')"
          >
            <GlobalIcon name="information-circle" size="sm" class="mr-2" />
            开发说明
          </a>
        </nav>
      </div>
    </div>

    <!-- Icon 组件测试 -->
    <section id="icon-test" class="mb-12 scroll-mt-32">
      <div class="flex items-center gap-3 mb-6">
        <GlobalIcon name="star" size="lg" color="text-blue-500" />
        <h3 class="text-xl font-semibold text-gray-900">Icon 图标组件</h3>
      </div>
      <IconTest />
    </section>

    <!-- Button 组件测试 -->
    <section id="button-test" class="mb-12 scroll-mt-32">
      <div class="flex items-center gap-3 mb-6">
        <GlobalIcon name="check" size="lg" color="text-green-500" />
        <h3 class="text-xl font-semibold text-gray-900">Button 按钮组件</h3>
      </div>
      <ButtonTest />
    </section>

    <!-- 文件上传组件测试 -->
    <section id="upload-test" class="mb-12 scroll-mt-32">
      <div class="flex items-center gap-3 mb-6">
        <GlobalIcon name="upload" size="lg" color="text-indigo-500" />
        <h3 class="text-xl font-semibold text-gray-900">Upload 文件上传组件</h3>
      </div>
      <UploadTest />
    </section>

    <!-- 消息组件测试 -->
    <section id="message-test" class="mb-12 scroll-mt-32">
      <div class="flex items-center gap-3 mb-6">
        <GlobalIcon name="chat-bubble-left" size="lg" color="text-purple-500" />
        <h3 class="text-xl font-semibold text-gray-900">Message 消息组件</h3>
      </div>
      <MessageTest />
    </section>

    <!-- 存储组件测试 -->
    <section id="storage-test" class="mb-12 scroll-mt-32">
      <div class="flex items-center gap-3 mb-6">
        <GlobalIcon name="archive-box" size="lg" color="text-orange-500" />
        <h3 class="text-xl font-semibold text-gray-900">Storage 存储组件</h3>
      </div>
      <StorageTest />
    </section>

    <!-- 开发者信息 -->
    <section id="developer-info" class="scroll-mt-32">
      <div class="flex items-center gap-3 mb-6">
        <GlobalIcon name="information-circle" size="lg" color="text-gray-500" />
        <h3 class="text-xl font-semibold text-gray-900">开发者说明</h3>
      </div>
      <div class="bg-gradient-to-r from-blue-50 to-indigo-50 p-6 rounded-lg border border-blue-200">
        <h4 class="text-lg font-semibold text-blue-900 mb-3">项目说明</h4>
        <ul class="text-blue-800 text-sm space-y-2">
          <li class="flex items-start gap-2">
            <GlobalIcon name="check-circle" size="sm" color="text-blue-600" class="mt-0.5 flex-shrink-0" />
            此页面用于测试和预览组件效果
          </li>
          <li class="flex items-start gap-2">
            <GlobalIcon name="check-circle" size="sm" color="text-blue-600" class="mt-0.5 flex-shrink-0" />
            可以在这里验证组件的各种状态和行为
          </li>
          <li class="flex items-start gap-2">
            <GlobalIcon name="check-circle" size="sm" color="text-blue-600" class="mt-0.5 flex-shrink-0" />
            全局组件使用 Global 前缀，无需导入直接使用
          </li>
          <li class="flex items-start gap-2">
            <GlobalIcon name="check-circle" size="sm" color="text-blue-600" class="mt-0.5 flex-shrink-0" />
            Icon组件基于 Heroicons，支持 outline 和 solid 两种风格
          </li>
          <li class="flex items-start gap-2">
            <GlobalIcon name="check-circle" size="sm" color="text-blue-600" class="mt-0.5 flex-shrink-0" />
            安全存储支持加密和有效期管理
          </li>
          <li class="flex items-start gap-2">
            <GlobalIcon name="check-circle" size="sm" color="text-blue-600" class="mt-0.5 flex-shrink-0" />
            所有存储数据都经过加密处理
          </li>
          <li class="flex items-start gap-2">
            <GlobalIcon name="check-circle" size="sm" color="text-blue-600" class="mt-0.5 flex-shrink-0" />
            文件上传支持简单上传和大文件分片上传
          </li>
          <li class="flex items-start gap-2">
            <GlobalIcon name="check-circle" size="sm" color="text-blue-600" class="mt-0.5 flex-shrink-0" />
            支持断点续传、进度显示和MD5去重
          </li>
          <li class="flex items-start gap-2">
            <GlobalIcon name="check-circle" size="sm" color="text-blue-600" class="mt-0.5 flex-shrink-0" />
            方便开发和调试组件功能
          </li>
        </ul>

        <div class="mt-4 pt-4 border-t border-blue-200">
          <h5 class="font-medium text-blue-900 mb-2">快速导航</h5>
          <div class="flex flex-wrap gap-2 text-sm">
            <button
              @click="scrollToTop"
              class="inline-flex items-center px-2 py-1 bg-blue-100 text-blue-700 rounded hover:bg-blue-200 transition-colors"
            >
              <GlobalIcon name="arrow-up" size="xs" class="mr-1" />
              返回顶部
            </button>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
import IconTest from './components/IconTest.vue'
import ButtonTest from './components/ButtonTest.vue'
import UploadTest from './components/UploadTest.vue'
import MessageTest from './components/MessageTest.vue'
import StorageTest from './components/StorageTest.vue'

defineOptions({
  name: 'DebugPage'
})

// 平滑滚动到指定区域
const scrollToSection = (sectionId: string) => {
  const element = document.getElementById(sectionId)
  if (element) {
    element.scrollIntoView({
      behavior: 'smooth',
      block: 'start'
    })
  }
}

// 返回顶部
const scrollToTop = () => {
  window.scrollTo({
    top: 0,
    behavior: 'smooth'
  })
}
</script>
