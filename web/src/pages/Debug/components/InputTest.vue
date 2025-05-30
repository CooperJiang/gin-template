<template>
  <div class="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
    <h3 class="text-lg font-semibold mb-4">Input 输入框组件测试</h3>
    <p class="text-gray-600 text-sm mb-6">测试输入框组件的各种类型、状态和功能</p>

    <!-- 基础输入框 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">基础输入框</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">普通输入框</label>
            <GlobalInput
              v-model="basicInput"
              placeholder="请输入内容"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">带清空按钮</label>
            <GlobalInput
              v-model="clearableInput"
              placeholder="输入内容后显示清空按钮"
            />
          </div>
        </div>
        <div class="mt-4 p-3 bg-white rounded border">
          <h5 class="text-sm font-medium mb-2">输入值：</h5>
          <pre class="text-xs text-gray-600">基础输入: {{ basicInput }}
清空输入: {{ clearableInput }}</pre>
        </div>
      </div>
    </div>

    <!-- 输入框类型 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">输入框类型</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">文本输入</label>
            <GlobalInput
              v-model="textInput"
              type="text"
              placeholder="文本输入"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">邮箱输入</label>
            <GlobalInput
              v-model="emailInput"
              type="email"
              placeholder="请输入邮箱"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">密码输入</label>
            <GlobalInput
              v-model="passwordInput"
              type="password"
              placeholder="请输入密码"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">密码（可显示）</label>
            <GlobalInput
              v-model="showPasswordInput"
              type="password"
              placeholder="密码可切换显示"
              :show-password="true"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">数字输入</label>
            <GlobalInput
              v-model="numberInput"
              type="number"
              placeholder="请输入数字"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">电话输入</label>
            <GlobalInput
              v-model="telInput"
              type="tel"
              placeholder="请输入电话号码"
            />
          </div>
        </div>
        <div class="mt-4 p-3 bg-white rounded border">
          <h5 class="text-sm font-medium mb-2">输入值：</h5>
          <pre class="text-xs text-gray-600">文本: {{ textInput }}
邮箱: {{ emailInput }}
密码: {{ passwordInput }}
显示密码: {{ showPasswordInput }}
数字: {{ numberInput }}
电话: {{ telInput }}</pre>
        </div>
      </div>
    </div>

    <!-- 图标和前后缀 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">图标和前后缀</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">前缀图标</label>
            <GlobalInput
              v-model="prefixIconInput"
              placeholder="用户名"
              prefix-icon="user"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">后缀图标</label>
            <GlobalInput
              v-model="suffixIconInput"
              placeholder="搜索内容"
              suffix-icon="magnifying-glass"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">邮箱（前缀图标）</label>
            <GlobalInput
              v-model="emailPrefixInput"
              type="email"
              placeholder="请输入邮箱"
              prefix-icon="envelope"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">链接（后缀图标）</label>
            <GlobalInput
              v-model="urlInput"
              type="url"
              placeholder="请输入网址"
              suffix-icon="link"
            />
          </div>
        </div>
        <div class="mt-4 p-3 bg-white rounded border">
          <h5 class="text-sm font-medium mb-2">输入值：</h5>
          <pre class="text-xs text-gray-600">前缀图标: {{ prefixIconInput }}
后缀图标: {{ suffixIconInput }}
邮箱: {{ emailPrefixInput }}
链接: {{ urlInput }}</pre>
        </div>
      </div>
    </div>

    <!-- 输入框尺寸 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">输入框尺寸</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="space-y-3">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">小尺寸 (small)</label>
            <GlobalInput
              v-model="smallSizeInput"
              size="small"
              placeholder="小尺寸输入框"
              prefix-icon="user"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">中等尺寸 (medium) - 默认</label>
            <GlobalInput
              v-model="mediumSizeInput"
              size="medium"
              placeholder="中等尺寸输入框"
              prefix-icon="user"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">大尺寸 (large)</label>
            <GlobalInput
              v-model="largeSizeInput"
              size="large"
              placeholder="大尺寸输入框"
              prefix-icon="user"
            />
          </div>
        </div>
      </div>
    </div>

    <!-- 状态演示 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">输入框状态</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">正常状态</label>
            <GlobalInput
              v-model="normalInput"
              placeholder="正常状态输入框"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">禁用状态</label>
            <GlobalInput
              v-model="disabledInput"
              placeholder="禁用状态输入框"
              :disabled="true"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">只读状态</label>
            <GlobalInput
              v-model="readonlyInput"
              placeholder="只读状态输入框"
              :readonly="true"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">错误状态</label>
            <GlobalInput
              v-model="errorInput"
              placeholder="错误状态输入框"
              :error="true"
            />
          </div>
        </div>
        <div class="flex gap-2 mt-4">
          <GlobalButton
            type="secondary"
            size="small"
            @click="toggleStates"
          >
            切换状态
          </GlobalButton>
        </div>
      </div>
    </div>

    <!-- 功能测试 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">功能测试</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">限制长度（最大10字符）</label>
            <GlobalInput
              v-model="maxLengthInput"
              placeholder="最多输入10个字符"
              :maxlength="10"
            />
            <p class="text-xs text-gray-500 mt-1">当前: {{ maxLengthInput.length }}/10</p>
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">自动聚焦</label>
            <GlobalInput
              v-model="autofocusInput"
              placeholder="页面加载时自动聚焦"
              :autofocus="true"
            />
          </div>
        </div>

        <div class="space-y-4">
          <div>
            <h5 class="text-sm font-medium text-gray-700 mb-2">事件测试</h5>
            <GlobalInput
              v-model="eventTestInput"
              placeholder="输入内容测试事件"
              @input="handleInput"
              @focus="handleFocus"
              @blur="handleBlur"
              @enter="handleEnter"
              @clear="handleClear"
            />
            <div class="mt-2 p-2 bg-white rounded border text-xs">
              <p><strong>事件日志:</strong></p>
              <div class="max-h-20 overflow-y-auto">
                <div v-for="(log, index) in eventLogs" :key="index" class="text-gray-600">
                  {{ log }}
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- API测试 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">API方法测试</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">方法测试输入框</label>
          <GlobalInput
            ref="apiTestInputRef"
            v-model="apiTestInput"
            placeholder="用于测试API方法"
            prefix-icon="wrench-screwdriver"
          />
        </div>
        <div class="flex flex-wrap gap-2">
          <GlobalButton
            type="primary"
            size="small"
            @click="focusInput"
          >
            获取焦点
          </GlobalButton>
          <GlobalButton
            type="secondary"
            size="small"
            @click="blurInput"
          >
            失去焦点
          </GlobalButton>
          <GlobalButton
            type="warning"
            size="small"
            @click="selectInput"
          >
            选中内容
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
                <td class="px-4 py-2 text-sm text-gray-500">string | number</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">输入框的值</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">type</td>
                <td class="px-4 py-2 text-sm text-gray-500">'text' | 'email' | 'password' | 'number' | 'tel' | 'url' | 'search'</td>
                <td class="px-4 py-2 text-sm text-gray-500">'text'</td>
                <td class="px-4 py-2 text-sm text-gray-500">输入框类型</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">placeholder</td>
                <td class="px-4 py-2 text-sm text-gray-500">string</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">占位符文本</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">size</td>
                <td class="px-4 py-2 text-sm text-gray-500">'small' | 'medium' | 'large'</td>
                <td class="px-4 py-2 text-sm text-gray-500">'medium'</td>
                <td class="px-4 py-2 text-sm text-gray-500">输入框尺寸</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">disabled</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean</td>
                <td class="px-4 py-2 text-sm text-gray-500">false</td>
                <td class="px-4 py-2 text-sm text-gray-500">是否禁用</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">readonly</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean</td>
                <td class="px-4 py-2 text-sm text-gray-500">false</td>
                <td class="px-4 py-2 text-sm text-gray-500">是否只读</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">error</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean</td>
                <td class="px-4 py-2 text-sm text-gray-500">false</td>
                <td class="px-4 py-2 text-sm text-gray-500">是否显示错误状态</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">showPassword</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean</td>
                <td class="px-4 py-2 text-sm text-gray-500">false</td>
                <td class="px-4 py-2 text-sm text-gray-500">是否显示密码切换按钮</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">prefixIcon</td>
                <td class="px-4 py-2 text-sm text-gray-500">string</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">前缀图标名称</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">suffixIcon</td>
                <td class="px-4 py-2 text-sm text-gray-500">string</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">后缀图标名称</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">maxlength</td>
                <td class="px-4 py-2 text-sm text-gray-500">number</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">最大输入长度</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">autofocus</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean</td>
                <td class="px-4 py-2 text-sm text-gray-500">false</td>
                <td class="px-4 py-2 text-sm text-gray-500">是否自动获取焦点</td>
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
                <td class="px-4 py-2 text-sm text-gray-500">(value: string | number)</td>
                <td class="px-4 py-2 text-sm text-gray-500">输入值变化时触发</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">input</td>
                <td class="px-4 py-2 text-sm text-gray-500">(value: string | number, event: Event)</td>
                <td class="px-4 py-2 text-sm text-gray-500">输入时触发</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">focus</td>
                <td class="px-4 py-2 text-sm text-gray-500">(event: FocusEvent)</td>
                <td class="px-4 py-2 text-sm text-gray-500">获得焦点时触发</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">blur</td>
                <td class="px-4 py-2 text-sm text-gray-500">(event: FocusEvent)</td>
                <td class="px-4 py-2 text-sm text-gray-500">失去焦点时触发</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">enter</td>
                <td class="px-4 py-2 text-sm text-gray-500">(event: KeyboardEvent)</td>
                <td class="px-4 py-2 text-sm text-gray-500">按下回车键时触发</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">clear</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">清空内容时触发</td>
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
                <td class="px-4 py-2 text-sm text-gray-500">让输入框获得焦点</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">blur</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">让输入框失去焦点</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">select</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">选中输入框的所有内容</td>
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
  &lt;GlobalInput
    v-model="inputValue"
    type="email"
    placeholder="请输入邮箱"
    prefix-icon="envelope"
    size="medium"
    :error="hasError"
    @input="handleInput"
    @enter="handleSubmit"
  /&gt;
&lt;/template&gt;

&lt;script setup&gt;
import { ref } from 'vue'

const inputValue = ref('')
const hasError = ref(false)

const handleInput = (value, event) => {
  console.log('输入值:', value)
  // 邮箱验证
  hasError.value = !value.includes('@')
}

const handleSubmit = () => {
  console.log('提交:', inputValue.value)
}
&lt;/script&gt;</code></pre>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import type { InputInstance } from '../../../components/Input/types'

defineOptions({
  name: 'InputTest'
})

// 基础输入框
const basicInput = ref('')
const clearableInput = ref('这里有初始内容')

// 输入框类型
const textInput = ref('')
const emailInput = ref('')
const passwordInput = ref('')
const showPasswordInput = ref('')
const numberInput = ref('')
const telInput = ref('')

// 图标和前后缀
const prefixIconInput = ref('')
const suffixIconInput = ref('')
const emailPrefixInput = ref('')
const urlInput = ref('')

// 输入框尺寸
const smallSizeInput = ref('')
const mediumSizeInput = ref('')
const largeSizeInput = ref('')

// 状态演示
const normalInput = ref('')
const disabledInput = ref('禁用状态的值')
const readonlyInput = ref('只读状态的值')
const errorInput = ref('错误状态的值')

// 功能测试
const maxLengthInput = ref('')
const autofocusInput = ref('')
const eventTestInput = ref('')
const eventLogs = ref<string[]>([])

// API测试
const apiTestInput = ref('测试内容')
const apiTestInputRef = ref<InputInstance>()

// 事件处理
const handleInput = (value: string | number, event: Event) => {
  const timestamp = new Date().toLocaleTimeString()
  eventLogs.value.unshift(`${timestamp} - Input: ${value}`)
  if (eventLogs.value.length > 5) {
    eventLogs.value = eventLogs.value.slice(0, 5)
  }
}

const handleFocus = (event: FocusEvent) => {
  const timestamp = new Date().toLocaleTimeString()
  eventLogs.value.unshift(`${timestamp} - Focus 获得焦点`)
  if (eventLogs.value.length > 5) {
    eventLogs.value = eventLogs.value.slice(0, 5)
  }
}

const handleBlur = (event: FocusEvent) => {
  const timestamp = new Date().toLocaleTimeString()
  eventLogs.value.unshift(`${timestamp} - Blur 失去焦点`)
  if (eventLogs.value.length > 5) {
    eventLogs.value = eventLogs.value.slice(0, 5)
  }
}

const handleEnter = (event: KeyboardEvent) => {
  const timestamp = new Date().toLocaleTimeString()
  eventLogs.value.unshift(`${timestamp} - Enter 回车键`)
  if (eventLogs.value.length > 5) {
    eventLogs.value = eventLogs.value.slice(0, 5)
  }
}

const handleClear = () => {
  const timestamp = new Date().toLocaleTimeString()
  eventLogs.value.unshift(`${timestamp} - Clear 清空内容`)
  if (eventLogs.value.length > 5) {
    eventLogs.value = eventLogs.value.slice(0, 5)
  }
}

// API 方法测试
const focusInput = () => {
  apiTestInputRef.value?.focus()
}

const blurInput = () => {
  apiTestInputRef.value?.blur()
}

const selectInput = () => {
  apiTestInputRef.value?.select()
}

// 状态切换
const toggleStates = () => {
  // 这里可以添加状态切换逻辑
  console.log('切换状态')
}
</script>

<style scoped>
/* 样式定制 */
</style>
