<template>
  <div class="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
    <h3 class="text-lg font-semibold mb-4">消息提示组件测试</h3>
    <p class="text-gray-600 text-sm mb-4">测试不同类型的消息提示效果和自定义持续时间</p>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      <!-- 基础消息测试 -->
      <div class="space-y-3">
        <h4 class="font-medium text-gray-900">基础消息</h4>
        <div class="space-y-2">
          <button
            @click="testSuccess"
            class="w-full px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700 transition-colors text-sm"
          >
            成功消息 (2秒)
          </button>
          <button
            @click="testError"
            class="w-full px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700 transition-colors text-sm"
          >
            错误消息 (5秒)
          </button>
          <button
            @click="testWarning"
            class="w-full px-4 py-2 bg-yellow-600 text-white rounded-md hover:bg-yellow-700 transition-colors text-sm"
          >
            警告消息 (默认)
          </button>
          <button
            @click="testInfo"
            class="w-full px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors text-sm"
          >
            信息消息 (1秒)
          </button>
        </div>
      </div>

      <!-- 队列测试 -->
      <div class="space-y-3">
        <h4 class="font-medium text-gray-900">队列测试</h4>
        <div class="space-y-2">
          <button
            @click="testMultiple"
            class="w-full px-4 py-2 bg-purple-600 text-white rounded-md hover:bg-purple-700 transition-colors text-sm"
          >
            多条消息队列
          </button>
          <button
            @click="testLongMessage"
            class="w-full px-4 py-2 bg-indigo-600 text-white rounded-md hover:bg-indigo-700 transition-colors text-sm"
          >
            长文本消息
          </button>
          <button
            @click="testRapidMessages"
            class="w-full px-4 py-2 bg-pink-600 text-white rounded-md hover:bg-pink-700 transition-colors text-sm"
          >
            快速连续消息
          </button>
        </div>
      </div>

      <!-- 特殊测试 -->
      <div class="space-y-3">
        <h4 class="font-medium text-gray-900">特殊场景</h4>
        <div class="space-y-2">
          <button
            @click="testPersistent"
            class="w-full px-4 py-2 bg-gray-600 text-white rounded-md hover:bg-gray-700 transition-colors text-sm"
          >
            持久消息 (10秒)
          </button>
          <button
            @click="testQuick"
            class="w-full px-4 py-2 bg-teal-600 text-white rounded-md hover:bg-teal-700 transition-colors text-sm"
          >
            快速消息 (0.5秒)
          </button>
          <button
            @click="clearAllMessages"
            class="w-full px-4 py-2 bg-orange-600 text-white rounded-md hover:bg-orange-700 transition-colors text-sm"
          >
            清空所有消息
          </button>
        </div>
      </div>
    </div>

    <!-- API 说明 -->
    <div class="demo-section">
      <h4 class="demo-title">Composable API 说明</h4>
      <p class="demo-description">
        Message组件通过 <code class="bg-gray-100 px-1 rounded">useMessage()</code> composable使用，提供以下方法：
      </p>
      <div class="overflow-x-auto">
        <table class="w-full border-collapse border border-gray-300 text-sm">
          <thead>
            <tr class="bg-gray-50">
              <th class="border border-gray-300 px-4 py-2 text-left">方法名</th>
              <th class="border border-gray-300 px-4 py-2 text-left">说明</th>
              <th class="border border-gray-300 px-4 py-2 text-left">参数</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="method in methodsApi" :key="method.name" class="hover:bg-gray-50">
              <td class="border border-gray-300 px-4 py-2 font-mono text-blue-600">{{ method.name }}</td>
              <td class="border border-gray-300 px-4 py-2">{{ method.description }}</td>
              <td class="border border-gray-300 px-4 py-2 font-mono text-green-600">{{ method.params }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- 使用示例 -->
    <div class="demo-section">
      <h4 class="demo-title">使用示例</h4>
      <div class="bg-gray-900 text-gray-100 p-4 rounded text-sm font-mono overflow-x-auto">
        <div class="text-green-400">// 导入composable</div>
        <div class="text-blue-400">import</div> <div class="text-white">{ useMessage }</div> <div class="text-blue-400">from</div> <div class="text-yellow-400">'@/composables/useMessage'</div>
        <br /><br />
        <div class="text-green-400">// 在setup中使用</div>
        <div class="text-blue-400">const</div> <div class="text-white">{ success, error, warning, info, clearMessages } = useMessage()</div>
        <br /><br />
        <div class="text-green-400">// 显示消息</div>
        <div class="text-white">success('操作成功！')</div>
        <div class="text-white">error('操作失败！', 5000) // 5秒后自动关闭</div>
        <div class="text-white">warning('注意事项')</div>
        <div class="text-white">info('提示信息')</div>
        <br />
        <div class="text-green-400">// 清空所有消息</div>
        <div class="text-white">clearMessages()</div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useMessage } from '../../../composables/useMessage'

defineOptions({
  name: 'MessageTest'
})

const { success, error, warning, info, clearMessages } = useMessage()

// 基础消息测试
const testSuccess = () => {
  success('这是一条成功消息！', 2000)
}

const testError = () => {
  error('这是一条错误消息！', 5000)
}

const testWarning = () => {
  warning('这是一条警告消息！')
}

const testInfo = () => {
  info('这是一条信息消息！', 1000)
}

// 队列测试
const testMultiple = () => {
  success('第一条消息（2秒后消失）', 2000)
  setTimeout(() => error('第二条消息（4秒后消失）', 4000), 300)
  setTimeout(() => warning('第三条消息（默认时间）'), 600)
  setTimeout(() => info('第四条消息（1秒后消失）', 1000), 900)
}

const testLongMessage = () => {
  warning('这是一条很长很长的消息，用来测试消息框在显示长文本时的表现效果，看看是否能够正确地自适应宽度和换行。', 6000)
}

const testRapidMessages = () => {
  for (let i = 1; i <= 5; i++) {
    setTimeout(() => {
      info(`快速消息 ${i}/5`, 2000)
    }, i * 100)
  }
}

// 特殊场景测试
const testPersistent = () => {
  error('这是一条持久消息，将显示10秒钟', 10000)
}

const testQuick = () => {
  success('快速消息！', 500)
}

const clearAllMessages = () => {
  clearMessages()
  info('所有消息已清空', 1500)
}

// API 文档数据
const propsApi = [
  { name: '-', description: 'Message组件通过composable使用，无直接Props', type: '-', default: '-' },
]

const eventsApi = [
  { name: '-', description: 'Message组件通过composable使用，无直接Events', type: '-', default: '-' },
]

const methodsApi = [
  { name: 'success', description: '显示成功消息', params: '(message: string, duration?: number)' },
  { name: 'error', description: '显示错误消息', params: '(message: string, duration?: number)' },
  { name: 'warning', description: '显示警告消息', params: '(message: string, duration?: number)' },
  { name: 'info', description: '显示信息消息', params: '(message: string, duration?: number)' },
  { name: 'clearMessages', description: '清空所有消息', params: '()' },
]
</script>
