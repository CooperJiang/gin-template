<template>
  <div class="space-y-8">
    <!-- 页面标题 -->
    <div class="border-b border-gray-200 pb-4">
      <h2 class="text-2xl font-bold text-gray-900">Card 卡片组件测试</h2>
      <p class="mt-2 text-gray-600">测试卡片组件的各种功能和样式</p>
    </div>

    <!-- 基础用法 -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <GlobalCard title="基础卡片" subtitle="简单的卡片示例">
        <p class="text-gray-600">这是卡片的基础内容，支持任意的HTML内容。</p>
      </GlobalCard>

      <GlobalCard title="带操作的卡片">
        <template #extra>
          <GlobalButton size="small" variant="ghost" icon="ellipsis-horizontal" />
        </template>
        <p class="text-gray-600">卡片头部可以包含额外的操作按钮。</p>
      </GlobalCard>

      <GlobalCard>
        <template #cover>
          <div class="h-32 bg-gradient-to-r from-blue-400 to-purple-500"></div>
        </template>
        <template #title>自定义标题</template>
        <p class="text-gray-600">带有封面图片的卡片示例。</p>
      </GlobalCard>
    </div>

    <!-- 不同阴影 -->
    <GlobalCard title="不同阴影效果" class="w-full">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
        <GlobalCard shadow="none" title="无阴影">
          <p class="text-sm text-gray-600">shadow="none"</p>
        </GlobalCard>

        <GlobalCard shadow="small" title="小阴影">
          <p class="text-sm text-gray-600">shadow="small"</p>
        </GlobalCard>

        <GlobalCard shadow="medium" title="中等阴影">
          <p class="text-sm text-gray-600">shadow="medium"</p>
        </GlobalCard>

        <GlobalCard shadow="large" title="大阴影">
          <p class="text-sm text-gray-600">shadow="large"</p>
        </GlobalCard>
      </div>
    </GlobalCard>

    <!-- 不同背景色 -->
    <GlobalCard title="不同背景色" class="w-full">
      <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
        <GlobalCard background="white" title="白色背景" size="small">
          <p class="text-xs">默认白色</p>
        </GlobalCard>

        <GlobalCard background="gray" title="灰色背景" size="small">
          <p class="text-xs">淡灰色背景</p>
        </GlobalCard>

        <GlobalCard background="blue" title="蓝色背景" size="small">
          <p class="text-xs">淡蓝色背景</p>
        </GlobalCard>

        <GlobalCard background="green" title="绿色背景" size="small">
          <p class="text-xs">淡绿色背景</p>
        </GlobalCard>
      </div>
    </GlobalCard>

    <!-- 可悬停和加载状态 -->
    <GlobalCard title="交互状态" class="w-full">
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <GlobalCard title="可悬停卡片" hoverable @click="handleCardClick">
          <p class="text-gray-600">鼠标悬停时会有动画效果，点击查看事件。</p>
        </GlobalCard>

        <GlobalCard title="加载状态" :loading="isLoading">
          <p class="text-gray-600">卡片可以显示加载状态。</p>
          <template #footer>
            <GlobalButton @click="toggleLoading" size="small">
              {{ isLoading ? '停止加载' : '开始加载' }}
            </GlobalButton>
          </template>
        </GlobalCard>

        <GlobalCard title="自定义加载" :loading="isCustomLoading">
          <template #loading>
            <div class="flex flex-col items-center justify-center space-y-2">
              <div class="animate-pulse text-blue-500">
                <GlobalIcon name="heart" size="lg" />
              </div>
              <span class="text-sm text-gray-500">自定义加载中...</span>
            </div>
          </template>
          <p class="text-gray-600">可以自定义加载状态的显示内容。</p>
          <template #footer>
            <GlobalButton @click="toggleCustomLoading" size="small">
              {{ isCustomLoading ? '停止' : '开始' }}
            </GlobalButton>
          </template>
        </GlobalCard>
      </div>
    </GlobalCard>

    <!-- 复杂布局示例 -->
    <GlobalCard title="复杂布局示例" class="w-full">
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- 用户资料卡片 -->
        <GlobalCard shadow="medium" hoverable>
          <template #cover>
            <div class="h-24 bg-gradient-to-r from-purple-400 via-pink-500 to-red-500"></div>
          </template>

          <div class="relative -mt-8 px-4">
            <div class="flex items-center space-x-4">
              <div class="w-16 h-16 bg-white rounded-full border-4 border-white shadow-lg flex items-center justify-center">
                <GlobalIcon name="user" size="lg" class="text-gray-600" />
              </div>
              <div class="flex-1">
                <h3 class="text-lg font-semibold text-gray-900">张三</h3>
                <p class="text-sm text-gray-500">高级开发工程师</p>
              </div>
            </div>

            <div class="mt-4 space-y-2">
              <div class="flex items-center text-sm text-gray-600">
                <GlobalIcon name="envelope" size="sm" class="mr-2" />
                zhangsan@example.com
              </div>
              <div class="flex items-center text-sm text-gray-600">
                <GlobalIcon name="phone" size="sm" class="mr-2" />
                +86 138-0000-0000
              </div>
            </div>
          </div>

          <template #footer>
            <div class="flex space-x-2">
              <GlobalButton size="small" type="primary" class="flex-1">发消息</GlobalButton>
              <GlobalButton size="small" variant="outline" class="flex-1">查看资料</GlobalButton>
            </div>
          </template>
        </GlobalCard>

        <!-- 统计卡片 -->
        <GlobalCard title="数据统计" subtitle="实时数据展示">
          <template #extra>
            <span class="text-xs text-green-600 bg-green-100 px-2 py-1 rounded">实时</span>
          </template>

          <div class="grid grid-cols-2 gap-4">
            <div class="text-center p-3 bg-blue-50 rounded-lg">
              <div class="text-2xl font-bold text-blue-600">1,234</div>
              <div class="text-sm text-gray-500">总用户</div>
            </div>
            <div class="text-center p-3 bg-green-50 rounded-lg">
              <div class="text-2xl font-bold text-green-600">89</div>
              <div class="text-sm text-gray-500">在线用户</div>
            </div>
            <div class="text-center p-3 bg-yellow-50 rounded-lg">
              <div class="text-2xl font-bold text-yellow-600">456</div>
              <div class="text-sm text-gray-500">今日访问</div>
            </div>
            <div class="text-center p-3 bg-purple-50 rounded-lg">
              <div class="text-2xl font-bold text-purple-600">78%</div>
              <div class="text-sm text-gray-500">转化率</div>
            </div>
          </div>

          <template #footer>
            <div class="text-xs text-gray-500 text-center">
              最后更新: {{ new Date().toLocaleTimeString() }}
            </div>
          </template>
        </GlobalCard>
      </div>
    </GlobalCard>

    <!-- 事件测试 -->
    <GlobalCard title="事件测试" class="w-full">
      <div class="space-y-4">
        <div class="bg-gray-50 p-4 rounded-lg">
          <h4 class="text-sm font-medium text-gray-700 mb-2">事件日志:</h4>
          <div class="space-y-1 max-h-32 overflow-y-auto">
            <div
              v-for="(log, index) in eventLogs"
              :key="index"
              class="text-xs text-gray-600 font-mono"
            >
              [{{ log.time }}] {{ log.message }}
            </div>
          </div>
          <GlobalButton size="small" variant="outline" class="mt-2" @click="clearLogs">
            清空日志
          </GlobalButton>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
          <GlobalCard
            title="点击事件"
            hoverable
            @click="addLog('卡片被点击')"
            @mouseenter="addLog('鼠标进入卡片')"
            @mouseleave="addLog('鼠标离开卡片')"
          >
            <p class="text-sm text-gray-600">点击或悬停查看事件</p>
          </GlobalCard>

          <GlobalCard title="鼠标事件" @mouseenter="addLog('鼠标进入')">
            <p class="text-sm text-gray-600">鼠标进入时触发事件</p>
          </GlobalCard>

          <GlobalCard title="组合事件" hoverable @click="addLog('组合事件触发')">
            <p class="text-sm text-gray-600">组合多种交互效果</p>
          </GlobalCard>
        </div>
      </div>
    </GlobalCard>

    <!-- Props API 说明 -->
    <GlobalCard title="Props API 说明" class="w-full">
      <div class="overflow-x-auto">
        <table class="w-full border-collapse border border-gray-300 text-sm">
          <thead>
            <tr class="bg-gray-50">
              <th class="border border-gray-300 px-4 py-2 text-left">参数</th>
              <th class="border border-gray-300 px-4 py-2 text-left">说明</th>
              <th class="border border-gray-300 px-4 py-2 text-left">类型</th>
              <th class="border border-gray-300 px-4 py-2 text-left">默认值</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="prop in propsApi" :key="prop.name" class="hover:bg-gray-50">
              <td class="border border-gray-300 px-4 py-2 font-mono text-blue-600">{{ prop.name }}</td>
              <td class="border border-gray-300 px-4 py-2">{{ prop.description }}</td>
              <td class="border border-gray-300 px-4 py-2 font-mono text-green-600">{{ prop.type }}</td>
              <td class="border border-gray-300 px-4 py-2 font-mono text-orange-600">{{ prop.default }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </GlobalCard>

    <!-- Events API 说明 -->
    <GlobalCard title="Events API 说明" class="w-full">
      <div class="overflow-x-auto">
        <table class="w-full border-collapse border border-gray-300 text-sm">
          <thead>
            <tr class="bg-gray-50">
              <th class="border border-gray-300 px-4 py-2 text-left">事件名</th>
              <th class="border border-gray-300 px-4 py-2 text-left">说明</th>
              <th class="border border-gray-300 px-4 py-2 text-left">参数</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="event in eventsApi" :key="event.name" class="hover:bg-gray-50">
              <td class="border border-gray-300 px-4 py-2 font-mono text-blue-600">{{ event.name }}</td>
              <td class="border border-gray-300 px-4 py-2">{{ event.description }}</td>
              <td class="border border-gray-300 px-4 py-2 font-mono text-green-600">{{ event.params }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </GlobalCard>

    <!-- Slots API 说明 -->
    <GlobalCard title="Slots API 说明" class="w-full">
      <div class="overflow-x-auto">
        <table class="w-full border-collapse border border-gray-300 text-sm">
          <thead>
            <tr class="bg-gray-50">
              <th class="border border-gray-300 px-4 py-2 text-left">插槽名</th>
              <th class="border border-gray-300 px-4 py-2 text-left">说明</th>
              <th class="border border-gray-300 px-4 py-2 text-left">参数</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="slot in slotsApi" :key="slot.name" class="hover:bg-gray-50">
              <td class="border border-gray-300 px-4 py-2 font-mono text-blue-600">{{ slot.name }}</td>
              <td class="border border-gray-300 px-4 py-2">{{ slot.description }}</td>
              <td class="border border-gray-300 px-4 py-2 font-mono text-green-600">{{ slot.params }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </GlobalCard>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

defineOptions({
  name: 'CardTest'
})

// 状态管理
const isLoading = ref(false)
const isCustomLoading = ref(false)
const eventLogs = ref<Array<{ time: string; message: string }>>([])

// 方法
const toggleLoading = () => {
  isLoading.value = !isLoading.value
}

const toggleCustomLoading = () => {
  isCustomLoading.value = !isCustomLoading.value
}

const handleCardClick = () => {
  addLog('可悬停卡片被点击')
}

const addLog = (message: string) => {
  const time = new Date().toLocaleTimeString()
  eventLogs.value.unshift({ time, message })
  if (eventLogs.value.length > 10) {
    eventLogs.value.pop()
  }
}

const clearLogs = () => {
  eventLogs.value = []
}

// API 文档数据
const propsApi = [
  { name: 'title', description: '卡片标题', type: 'string', default: 'undefined' },
  { name: 'subtitle', description: '卡片副标题', type: 'string', default: 'undefined' },
  { name: 'size', description: '卡片大小', type: "'small' | 'medium' | 'large'", default: "'medium'" },
  { name: 'shadow', description: '卡片阴影', type: "'none' | 'small' | 'medium' | 'large' | 'hover'", default: "'small'" },
  { name: 'bordered', description: '是否显示边框', type: 'boolean', default: 'true' },
  { name: 'hoverable', description: '是否可悬停', type: 'boolean', default: 'false' },
  { name: 'background', description: '卡片背景色', type: "'white' | 'gray' | 'blue' | 'green' | 'yellow' | 'red' | 'purple' | 'indigo'", default: "'white'" },
  { name: 'headerDivider', description: '是否显示头部分割线', type: 'boolean', default: 'true' },
  { name: 'footerDivider', description: '是否显示底部分割线', type: 'boolean', default: 'true' },
  { name: 'rounded', description: '卡片圆角', type: "'none' | 'small' | 'medium' | 'large' | 'full'", default: "'medium'" },
  { name: 'headerPadding', description: '头部内边距', type: "'none' | 'small' | 'medium' | 'large'", default: "'medium'" },
  { name: 'bodyPadding', description: '内容内边距', type: "'none' | 'small' | 'medium' | 'large'", default: "'medium'" },
  { name: 'footerPadding', description: '底部内边距', type: "'none' | 'small' | 'medium' | 'large'", default: "'medium'" },
  { name: 'loading', description: '是否加载中', type: 'boolean', default: 'false' }
]

const eventsApi = [
  { name: 'click', description: '点击卡片时触发', params: '(event: MouseEvent)' },
  { name: 'mouseenter', description: '鼠标进入卡片时触发', params: '(event: MouseEvent)' },
  { name: 'mouseleave', description: '鼠标离开卡片时触发', params: '(event: MouseEvent)' }
]

const slotsApi = [
  { name: 'default', description: '卡片内容', params: '-' },
  { name: 'header', description: '头部插槽', params: '-' },
  { name: 'title', description: '标题插槽', params: '-' },
  { name: 'subtitle', description: '副标题插槽', params: '-' },
  { name: 'extra', description: '额外操作插槽', params: '-' },
  { name: 'footer', description: '底部插槽', params: '-' },
  { name: 'cover', description: '封面插槽', params: '-' },
  { name: 'loading', description: '加载状态插槽', params: '-' }
]
</script>
