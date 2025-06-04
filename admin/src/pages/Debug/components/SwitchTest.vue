<template>
  <div class="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
    <h3 class="text-lg font-semibold mb-4">Switch 开关组件测试</h3>
    <p class="text-gray-600 text-sm mb-6">测试开关组件的各种状态、尺寸和功能</p>

    <!-- 基础开关 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">基础开关</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="flex flex-wrap items-center gap-6">
          <div class="flex items-center gap-3">
            <span class="text-sm font-medium text-gray-700">基础开关:</span>
            <GlobalSwitch v-model="basicSwitch" />
          </div>
          <div class="flex items-center gap-3">
            <span class="text-sm font-medium text-gray-700">默认开启:</span>
            <GlobalSwitch v-model="defaultOnSwitch" />
          </div>
          <div class="flex items-center gap-3">
            <span class="text-sm font-medium text-gray-700">禁用状态:</span>
            <GlobalSwitch v-model="disabledSwitch" :disabled="true" />
          </div>
          <div class="flex items-center gap-3">
            <span class="text-sm font-medium text-gray-700">加载状态:</span>
            <GlobalSwitch v-model="loadingSwitch" :loading="true" />
          </div>
        </div>
        <div class="mt-4 p-3 bg-white rounded border">
          <h5 class="text-sm font-medium mb-2">开关状态：</h5>
          <pre class="text-xs text-gray-600">基础开关: {{ basicSwitch }}
默认开启: {{ defaultOnSwitch }}
禁用状态: {{ disabledSwitch }}
加载状态: {{ loadingSwitch }}</pre>
        </div>
      </div>
    </div>

    <!-- 开关尺寸 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">开关尺寸</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="flex flex-wrap items-center gap-6">
          <div class="flex items-center gap-3">
            <span class="text-sm font-medium text-gray-700">小尺寸:</span>
            <GlobalSwitch v-model="smallSwitch" size="small" />
          </div>
          <div class="flex items-center gap-3">
            <span class="text-sm font-medium text-gray-700">中等尺寸:</span>
            <GlobalSwitch v-model="mediumSwitch" size="medium" />
          </div>
          <div class="flex items-center gap-3">
            <span class="text-sm font-medium text-gray-700">大尺寸:</span>
            <GlobalSwitch v-model="largeSwitch" size="large" />
          </div>
        </div>
      </div>
    </div>

    <!-- 开关颜色 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">开关颜色</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
          <div class="flex items-center gap-3">
            <span class="text-sm font-medium text-gray-700">主要色:</span>
            <GlobalSwitch v-model="primarySwitch" color="primary" />
          </div>
          <div class="flex items-center gap-3">
            <span class="text-sm font-medium text-gray-700">成功色:</span>
            <GlobalSwitch v-model="successSwitch" color="success" />
          </div>
          <div class="flex items-center gap-3">
            <span class="text-sm font-medium text-gray-700">警告色:</span>
            <GlobalSwitch v-model="warningSwitch" color="warning" />
          </div>
          <div class="flex items-center gap-3">
            <span class="text-sm font-medium text-gray-700">危险色:</span>
            <GlobalSwitch v-model="dangerSwitch" color="danger" />
          </div>
        </div>
      </div>
    </div>

    <!-- 带文本的开关 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">带文本的开关</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="space-y-3">
          <div class="flex items-center gap-3">
            <span class="text-sm font-medium text-gray-700">开启/关闭文本:</span>
            <GlobalSwitch
              v-model="textSwitch1"
              active-text="开启"
              inactive-text="关闭"
            />
          </div>
          <div class="flex items-center gap-3">
            <span class="text-sm font-medium text-gray-700">是/否文本:</span>
            <GlobalSwitch
              v-model="textSwitch2"
              active-text="是"
              inactive-text="否"
              color="success"
            />
          </div>
          <div class="flex items-center gap-3">
            <span class="text-sm font-medium text-gray-700">启用/禁用:</span>
            <GlobalSwitch
              v-model="textSwitch3"
              active-text="启用"
              inactive-text="禁用"
              color="warning"
            />
          </div>
        </div>
      </div>
    </div>

    <!-- 带图标的开关 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">带图标的开关</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="flex flex-wrap items-center gap-6">
          <div class="flex items-center gap-3">
            <span class="text-sm font-medium text-gray-700">默认图标:</span>
            <GlobalSwitch v-model="iconSwitch1" :show-icon="true" />
          </div>
          <div class="flex items-center gap-3">
            <span class="text-sm font-medium text-gray-700">自定义图标:</span>
            <GlobalSwitch
              v-model="iconSwitch2"
              :show-icon="true"
              active-icon="sun"
              inactive-icon="moon"
              color="warning"
            />
          </div>
          <div class="flex items-center gap-3">
            <span class="text-sm font-medium text-gray-700">音量图标:</span>
            <GlobalSwitch
              v-model="iconSwitch3"
              :show-icon="true"
              active-icon="speaker-wave"
              inactive-icon="speaker-x-mark"
              color="success"
            />
          </div>
        </div>
      </div>
    </div>

    <!-- 自定义值 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">自定义值</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="flex flex-wrap items-center gap-6">
          <div class="flex items-center gap-3">
            <span class="text-sm font-medium text-gray-700">数字值:</span>
            <GlobalSwitch
              v-model="numberSwitch"
              :active-value="1"
              :inactive-value="0"
              active-text="1"
              inactive-text="0"
            />
          </div>
          <div class="flex items-center gap-3">
            <span class="text-sm font-medium text-gray-700">字符串值:</span>
            <GlobalSwitch
              v-model="stringSwitch"
              active-value="yes"
              inactive-value="no"
              active-text="Yes"
              inactive-text="No"
            />
          </div>
        </div>
        <div class="mt-4 p-3 bg-white rounded border">
          <h5 class="text-sm font-medium mb-2">自定义值：</h5>
          <pre class="text-xs text-gray-600">数字值: {{ numberSwitch }} (类型: {{ typeof numberSwitch }})
字符串值: {{ stringSwitch }} (类型: {{ typeof stringSwitch }})</pre>
        </div>
      </div>
    </div>

    <!-- 事件测试 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">事件测试</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="flex items-center gap-3">
          <span class="text-sm font-medium text-gray-700">事件开关:</span>
          <GlobalSwitch
            v-model="eventSwitch"
            @change="handleChange"
            @click="handleClick"
          />
        </div>
        <div class="mt-4 p-2 bg-white rounded border text-xs">
          <p><strong>事件日志:</strong></p>
          <div class="max-h-20 overflow-y-auto">
            <div v-for="(log, index) in eventLogs" :key="index" class="text-gray-600">
              {{ log }}
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- API方法测试 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">API方法测试</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="flex items-center gap-3">
          <span class="text-sm font-medium text-gray-700">API测试开关:</span>
          <GlobalSwitch
            ref="apiSwitchRef"
            v-model="apiSwitch"
          />
        </div>
        <div class="flex flex-wrap gap-2">
          <GlobalButton
            type="primary"
            size="small"
            @click="focusSwitch"
          >
            获取焦点
          </GlobalButton>
          <GlobalButton
            type="secondary"
            size="small"
            @click="blurSwitch"
          >
            失去焦点
          </GlobalButton>
          <GlobalButton
            type="warning"
            size="small"
            @click="toggleSwitch"
          >
            切换状态
          </GlobalButton>
        </div>
      </div>
    </div>

    <!-- API 文档 -->
    <div class="border-t pt-8">
      <h3 class="text-lg font-semibold mb-4">API 文档</h3>

      <!-- Props -->
      <div class="mb-6">
        <h4 class="font-medium text-gray-900 mb-3">Props</h4>
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200 border border-gray-200 rounded">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">属性</th>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">类型</th>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">默认值</th>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">说明</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">modelValue</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean</td>
                <td class="px-4 py-2 text-sm text-gray-500">false</td>
                <td class="px-4 py-2 text-sm text-gray-500">开关状态</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">disabled</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean</td>
                <td class="px-4 py-2 text-sm text-gray-500">false</td>
                <td class="px-4 py-2 text-sm text-gray-500">是否禁用</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">size</td>
                <td class="px-4 py-2 text-sm text-gray-500">'small' | 'medium' | 'large'</td>
                <td class="px-4 py-2 text-sm text-gray-500">'medium'</td>
                <td class="px-4 py-2 text-sm text-gray-500">开关尺寸</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">color</td>
                <td class="px-4 py-2 text-sm text-gray-500">'primary' | 'success' | 'warning' | 'danger'</td>
                <td class="px-4 py-2 text-sm text-gray-500">'primary'</td>
                <td class="px-4 py-2 text-sm text-gray-500">开启时的颜色</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">activeText</td>
                <td class="px-4 py-2 text-sm text-gray-500">string</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">开启时的文本</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">inactiveText</td>
                <td class="px-4 py-2 text-sm text-gray-500">string</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">关闭时的文本</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">activeValue</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean | string | number</td>
                <td class="px-4 py-2 text-sm text-gray-500">true</td>
                <td class="px-4 py-2 text-sm text-gray-500">开启时的值</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">inactiveValue</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean | string | number</td>
                <td class="px-4 py-2 text-sm text-gray-500">false</td>
                <td class="px-4 py-2 text-sm text-gray-500">关闭时的值</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">showIcon</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean</td>
                <td class="px-4 py-2 text-sm text-gray-500">false</td>
                <td class="px-4 py-2 text-sm text-gray-500">是否显示图标</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">activeIcon</td>
                <td class="px-4 py-2 text-sm text-gray-500">string</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">自定义开启图标</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">inactiveIcon</td>
                <td class="px-4 py-2 text-sm text-gray-500">string</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">自定义关闭图标</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">loading</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean</td>
                <td class="px-4 py-2 text-sm text-gray-500">false</td>
                <td class="px-4 py-2 text-sm text-gray-500">是否加载中</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Events -->
      <div class="mb-6">
        <h4 class="font-medium text-gray-900 mb-3">Events</h4>
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200 border border-gray-200 rounded">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">事件名</th>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">参数</th>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">说明</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">update:modelValue</td>
                <td class="px-4 py-2 text-sm text-gray-500">(value: boolean | string | number)</td>
                <td class="px-4 py-2 text-sm text-gray-500">状态改变时触发</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">change</td>
                <td class="px-4 py-2 text-sm text-gray-500">(value: boolean | string | number)</td>
                <td class="px-4 py-2 text-sm text-gray-500">状态改变时触发</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">click</td>
                <td class="px-4 py-2 text-sm text-gray-500">(event: MouseEvent)</td>
                <td class="px-4 py-2 text-sm text-gray-500">点击时触发</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Methods -->
      <div class="mb-6">
        <h4 class="font-medium text-gray-900 mb-3">Methods</h4>
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200 border border-gray-200 rounded">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">方法名</th>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">参数</th>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">说明</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">focus</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">让开关获得焦点</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">blur</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">让开关失去焦点</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">toggle</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">切换开关状态</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- 使用示例 -->
      <div>
        <h4 class="font-medium text-gray-900 mb-3">使用示例</h4>
        <div class="bg-gray-50 p-4 rounded-md">
          <pre class="text-sm text-gray-800 overflow-x-auto"><code>&lt;template&gt;
  &lt;GlobalSwitch
    v-model="switchValue"
    size="medium"
    color="primary"
    active-text="开启"
    inactive-text="关闭"
    :show-icon="true"
    @change="handleChange"
  /&gt;
&lt;/template&gt;

&lt;script setup&gt;
import { ref } from 'vue'

const switchValue = ref(false)

const handleChange = (value) => {
  console.log('开关状态变化:', value)
}
&lt;/script&gt;</code></pre>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import type { SwitchInstance } from '../../../components/Switch/types'

defineOptions({
  name: 'SwitchTest'
})

// 基础开关
const basicSwitch = ref(false)
const defaultOnSwitch = ref(true)
const disabledSwitch = ref(true)
const loadingSwitch = ref(false)

// 开关尺寸
const smallSwitch = ref(true)
const mediumSwitch = ref(true)
const largeSwitch = ref(true)

// 开关颜色
const primarySwitch = ref(true)
const successSwitch = ref(true)
const warningSwitch = ref(true)
const dangerSwitch = ref(true)

// 带文本的开关
const textSwitch1 = ref(false)
const textSwitch2 = ref(true)
const textSwitch3 = ref(false)

// 带图标的开关
const iconSwitch1 = ref(true)
const iconSwitch2 = ref(false)
const iconSwitch3 = ref(true)

// 自定义值
const numberSwitch = ref(0)
const stringSwitch = ref('no')

// 事件测试
const eventSwitch = ref(false)
const eventLogs = ref<string[]>([])

// API测试
const apiSwitch = ref(false)
const apiSwitchRef = ref<SwitchInstance>()

// 事件处理
const handleChange = (value: boolean | string | number) => {
  const timestamp = new Date().toLocaleTimeString()
  eventLogs.value.unshift(`${timestamp} - Change: ${value}`)
  if (eventLogs.value.length > 5) {
    eventLogs.value = eventLogs.value.slice(0, 5)
  }
}

const handleClick = (event: MouseEvent) => {
  const timestamp = new Date().toLocaleTimeString()
  eventLogs.value.unshift(`${timestamp} - Click`)
  if (eventLogs.value.length > 5) {
    eventLogs.value = eventLogs.value.slice(0, 5)
  }
}

// API 方法测试
const focusSwitch = () => {
  apiSwitchRef.value?.focus()
}

const blurSwitch = () => {
  apiSwitchRef.value?.blur()
}

const toggleSwitch = () => {
  apiSwitchRef.value?.toggle()
}
</script>

<style scoped>
/* 样式定制 */
</style>
