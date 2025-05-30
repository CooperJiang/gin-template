<template>
  <div class="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
    <h3 class="text-lg font-semibold mb-4">Textarea 文本域组件测试</h3>
    <p class="text-gray-600 text-sm mb-6">测试文本域组件的各种功能、状态和特性</p>

    <!-- 基础文本域 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">基础文本域</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">普通文本域</label>
            <GlobalTextarea
              v-model="basicTextarea"
              placeholder="请输入多行文本内容..."
              :rows="3"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">带清空按钮</label>
            <GlobalTextarea
              v-model="clearableTextarea"
              placeholder="输入内容后显示清空按钮"
              :rows="3"
            />
          </div>
        </div>
        <div class="mt-4 p-3 bg-white rounded border">
          <h5 class="text-sm font-medium mb-2">输入值：</h5>
          <pre class="text-xs text-gray-600">基础文本域: {{ basicTextarea }}
清空文本域: {{ clearableTextarea }}</pre>
        </div>
      </div>
    </div>

    <!-- 文本域尺寸 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">文本域尺寸</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="space-y-3">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">小尺寸 (small)</label>
            <GlobalTextarea
              v-model="smallSizeTextarea"
              size="small"
              placeholder="小尺寸文本域"
              :rows="3"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">中等尺寸 (medium) - 默认</label>
            <GlobalTextarea
              v-model="mediumSizeTextarea"
              size="medium"
              placeholder="中等尺寸文本域"
              :rows="3"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">大尺寸 (large)</label>
            <GlobalTextarea
              v-model="largeSizeTextarea"
              size="large"
              placeholder="大尺寸文本域"
              :rows="3"
            />
          </div>
        </div>
      </div>
    </div>

    <!-- 文本域状态 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">文本域状态</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">正常状态</label>
            <GlobalTextarea
              v-model="normalTextarea"
              placeholder="正常状态文本域"
              :rows="3"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">禁用状态</label>
            <GlobalTextarea
              v-model="disabledTextarea"
              placeholder="禁用状态文本域"
              :disabled="true"
              :rows="3"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">只读状态</label>
            <GlobalTextarea
              v-model="readonlyTextarea"
              placeholder="只读状态文本域"
              :readonly="true"
              :rows="3"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">错误状态</label>
            <GlobalTextarea
              v-model="errorTextarea"
              placeholder="错误状态文本域"
              :error="true"
              :rows="3"
            />
          </div>
        </div>
      </div>
    </div>

    <!-- 字符计数功能 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">字符计数功能</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">限制100字符</label>
            <GlobalTextarea
              v-model="countTextarea100"
              placeholder="最多输入100个字符"
              :maxlength="100"
              :show-count="true"
              :rows="4"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">限制200字符</label>
            <GlobalTextarea
              v-model="countTextarea200"
              placeholder="最多输入200个字符"
              :maxlength="200"
              :show-count="true"
              :rows="4"
            />
          </div>
        </div>
        <div class="mt-4 p-3 bg-white rounded border">
          <h5 class="text-sm font-medium mb-2">当前字符数：</h5>
          <pre class="text-xs text-gray-600">100字符限制: {{ countTextarea100.length }}/100
200字符限制: {{ countTextarea200.length }}/200</pre>
        </div>
      </div>
    </div>

    <!-- 自动调整高度 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">自动调整高度</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">自动高度（无限制）</label>
            <GlobalTextarea
              v-model="autosizeTextarea"
              placeholder="输入更多内容，高度会自动调整..."
              :autosize="true"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">自动高度（3-6行）</label>
            <GlobalTextarea
              v-model="autosizeLimitedTextarea"
              placeholder="高度限制在3-6行之间..."
              :autosize="{ minRows: 3, maxRows: 6 }"
            />
          </div>
        </div>
      </div>
    </div>

    <!-- 调整大小模式 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">调整大小模式</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">垂直调整 (默认)</label>
            <GlobalTextarea
              v-model="resizeVerticalTextarea"
              placeholder="可以垂直调整大小"
              resize="vertical"
              :rows="3"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">水平调整</label>
            <GlobalTextarea
              v-model="resizeHorizontalTextarea"
              placeholder="可以水平调整大小"
              resize="horizontal"
              :rows="3"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">双向调整</label>
            <GlobalTextarea
              v-model="resizeBothTextarea"
              placeholder="可以双向调整大小"
              resize="both"
              :rows="3"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">禁止调整</label>
            <GlobalTextarea
              v-model="resizeNoneTextarea"
              placeholder="不可调整大小"
              resize="none"
              :rows="3"
            />
          </div>
        </div>
      </div>
    </div>

    <!-- 功能测试 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">功能测试</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">自动聚焦</label>
            <GlobalTextarea
              v-model="autofocusTextarea"
              placeholder="页面加载时自动聚焦"
              :autofocus="true"
              :rows="3"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">多行内容</label>
            <GlobalTextarea
              v-model="multilineTextarea"
              placeholder="可以输入多行内容..."
              :rows="5"
            />
          </div>
        </div>

        <div class="space-y-4">
          <div>
            <h5 class="text-sm font-medium text-gray-700 mb-2">事件测试</h5>
            <GlobalTextarea
              v-model="eventTestTextarea"
              placeholder="输入内容测试事件"
              :rows="3"
              @input="handleInput"
              @focus="handleFocus"
              @blur="handleBlur"
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
          <label class="block text-sm font-medium text-gray-700 mb-2">方法测试文本域</label>
          <GlobalTextarea
            ref="apiTestTextareaRef"
            v-model="apiTestTextarea"
            placeholder="用于测试API方法"
            :rows="3"
          />
        </div>
        <div class="flex flex-wrap gap-2">
          <GlobalButton
            type="primary"
            size="small"
            @click="focusTextarea"
          >
            获取焦点
          </GlobalButton>
          <GlobalButton
            type="secondary"
            size="small"
            @click="blurTextarea"
          >
            失去焦点
          </GlobalButton>
          <GlobalButton
            type="warning"
            size="small"
            @click="selectTextarea"
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
                <td class="px-4 py-2 text-sm text-gray-500">string</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">文本域的值</td>
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
                <td class="px-4 py-2 text-sm text-gray-500">文本域尺寸</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">rows</td>
                <td class="px-4 py-2 text-sm text-gray-500">number</td>
                <td class="px-4 py-2 text-sm text-gray-500">3</td>
                <td class="px-4 py-2 text-sm text-gray-500">文本域行数</td>
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
                <td class="px-4 py-2 text-sm font-medium text-gray-900">maxlength</td>
                <td class="px-4 py-2 text-sm text-gray-500">number</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">最大输入长度</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">showCount</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean</td>
                <td class="px-4 py-2 text-sm text-gray-500">false</td>
                <td class="px-4 py-2 text-sm text-gray-500">是否显示字符计数</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">autosize</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean | {minRows?: number, maxRows?: number}</td>
                <td class="px-4 py-2 text-sm text-gray-500">false</td>
                <td class="px-4 py-2 text-sm text-gray-500">是否自动调整高度</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">resize</td>
                <td class="px-4 py-2 text-sm text-gray-500">'none' | 'both' | 'horizontal' | 'vertical'</td>
                <td class="px-4 py-2 text-sm text-gray-500">'vertical'</td>
                <td class="px-4 py-2 text-sm text-gray-500">调整大小方式</td>
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
                <td class="px-4 py-2 text-sm text-gray-500">(value: string)</td>
                <td class="px-4 py-2 text-sm text-gray-500">输入值变化时触发</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">input</td>
                <td class="px-4 py-2 text-sm text-gray-500">(value: string, event: Event)</td>
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
                <td class="px-4 py-2 text-sm text-gray-500">让文本域获得焦点</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">blur</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">让文本域失去焦点</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">select</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">选中文本域的所有内容</td>
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
  &lt;GlobalTextarea
    v-model="textareaValue"
    placeholder="请输入内容..."
    :rows="4"
    :maxlength="200"
    :show-count="true"
    :autosize="{ minRows: 3, maxRows: 8 }"
    size="medium"
    :error="hasError"
    @input="handleInput"
    @focus="handleFocus"
  /&gt;
&lt;/template&gt;

&lt;script setup&gt;
import { ref } from 'vue'

const textareaValue = ref('')
const hasError = ref(false)

const handleInput = (value, event) => {
  console.log('输入值:', value)
  // 内容验证
  hasError.value = value.length > 180
}

const handleFocus = () => {
  console.log('文本域获得焦点')
}
&lt;/script&gt;</code></pre>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import type { TextareaInstance } from '../../../components/Textarea/types'

defineOptions({
  name: 'TextareaTest'
})

// 基础文本域
const basicTextarea = ref('')
const clearableTextarea = ref('这里有一些初始内容\n可以进行多行编辑')

// 文本域尺寸
const smallSizeTextarea = ref('')
const mediumSizeTextarea = ref('')
const largeSizeTextarea = ref('')

// 文本域状态
const normalTextarea = ref('')
const disabledTextarea = ref('这是一个禁用状态的文本域\n内容不可编辑')
const readonlyTextarea = ref('这是一个只读状态的文本域\n内容不可修改')
const errorTextarea = ref('这是一个错误状态的文本域\n显示红色边框')

// 字符计数功能
const countTextarea100 = ref('')
const countTextarea200 = ref('')

// 自动调整高度
const autosizeTextarea = ref('')
const autosizeLimitedTextarea = ref('')

// 调整大小模式
const resizeVerticalTextarea = ref('')
const resizeHorizontalTextarea = ref('')
const resizeBothTextarea = ref('')
const resizeNoneTextarea = ref('')

// 功能测试
const autofocusTextarea = ref('')
const multilineTextarea = ref('这是第一行内容\n这是第二行内容\n这是第三行内容')
const eventTestTextarea = ref('')
const eventLogs = ref<string[]>([])

// API测试
const apiTestTextarea = ref('这是一些测试内容\n用于演示API方法')
const apiTestTextareaRef = ref<TextareaInstance>()

// 事件处理
const handleInput = (value: string, event: Event) => {
  const timestamp = new Date().toLocaleTimeString()
  eventLogs.value.unshift(`${timestamp} - Input: ${value.length} 字符`)
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

const handleClear = () => {
  const timestamp = new Date().toLocaleTimeString()
  eventLogs.value.unshift(`${timestamp} - Clear 清空内容`)
  if (eventLogs.value.length > 5) {
    eventLogs.value = eventLogs.value.slice(0, 5)
  }
}

// API 方法测试
const focusTextarea = () => {
  apiTestTextareaRef.value?.focus()
}

const blurTextarea = () => {
  apiTestTextareaRef.value?.blur()
}

const selectTextarea = () => {
  apiTestTextareaRef.value?.select()
}
</script>

<style scoped>
/* 样式定制 */
</style>
