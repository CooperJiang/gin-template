<template>
  <div class="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
    <h3 class="text-lg font-semibold mb-4">Checkbox 复选框组件测试</h3>
    <p class="text-gray-600 text-sm mb-6">测试复选框组件的各种状态、尺寸和功能</p>

    <!-- 基础复选框 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">基础复选框</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="space-y-3">
          <GlobalCheckbox v-model="basicCheckbox" label="基础复选框" />
          <GlobalCheckbox v-model="defaultChecked" label="默认选中" />
          <GlobalCheckbox v-model="disabledUnchecked" label="禁用未选中" :disabled="true" />
          <GlobalCheckbox v-model="disabledChecked" label="禁用已选中" :disabled="true" />
          <GlobalCheckbox v-model="indeterminateCheckbox" label="中间状态" :indeterminate="true" />
        </div>
        <div class="mt-4 p-3 bg-white rounded border">
          <h5 class="text-sm font-medium mb-2">复选框状态：</h5>
          <pre class="text-xs text-gray-600">基础复选框: {{ basicCheckbox }}
默认选中: {{ defaultChecked }}
禁用未选中: {{ disabledUnchecked }}
禁用已选中: {{ disabledChecked }}
中间状态: {{ indeterminateCheckbox }}</pre>
        </div>
      </div>
    </div>

    <!-- 复选框尺寸 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">复选框尺寸</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="space-y-3">
          <GlobalCheckbox v-model="smallCheckbox" label="小尺寸复选框" size="small" />
          <GlobalCheckbox v-model="mediumCheckbox" label="中等尺寸复选框" size="medium" />
          <GlobalCheckbox v-model="largeCheckbox" label="大尺寸复选框" size="large" />
        </div>
      </div>
    </div>

    <!-- 复选框颜色 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">复选框颜色</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="space-y-3">
          <GlobalCheckbox v-model="primaryCheckbox" label="主要色复选框" color="primary" />
          <GlobalCheckbox v-model="successCheckbox" label="成功色复选框" color="success" />
          <GlobalCheckbox v-model="warningCheckbox" label="警告色复选框" color="warning" />
          <GlobalCheckbox v-model="dangerCheckbox" label="危险色复选框" color="danger" />
        </div>
      </div>
    </div>

    <!-- 自定义图标 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">自定义图标</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="space-y-3">
          <GlobalCheckbox
            v-model="heartCheckbox"
            label="爱心图标"
            checked-icon="heart"
            color="danger"
          />
          <GlobalCheckbox
            v-model="starCheckbox"
            label="星星图标"
            checked-icon="star"
            color="warning"
          />
          <GlobalCheckbox
            v-model="thumbsUpCheckbox"
            label="点赞图标"
            checked-icon="hand-thumb-up"
            color="success"
          />
        </div>
      </div>
    </div>

    <!-- 边框样式 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">边框样式</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="space-y-3">
          <GlobalCheckbox v-model="borderedCheckbox1" label="带边框复选框" :bordered="true" />
          <GlobalCheckbox v-model="borderedCheckbox2" label="带边框大尺寸" :bordered="true" size="large" />
        </div>
      </div>
    </div>

    <!-- 验证状态 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">验证状态</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="space-y-3">
          <GlobalCheckbox v-model="successValidation" label="验证成功" validate-status="success" />
          <GlobalCheckbox v-model="warningValidation" label="验证警告" validate-status="warning" />
          <GlobalCheckbox v-model="errorValidation" label="验证错误" validate-status="error" />
        </div>
      </div>
    </div>

    <!-- 自定义值 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">自定义值</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="space-y-3">
          <GlobalCheckbox
            v-model="numberValue"
            label="数字值 (1/0)"
            :true-value="1"
            :false-value="0"
          />
          <GlobalCheckbox
            v-model="stringValue"
            label="字符串值 (yes/no)"
            true-value="yes"
            false-value="no"
          />
        </div>
        <div class="mt-4 p-3 bg-white rounded border">
          <h5 class="text-sm font-medium mb-2">自定义值：</h5>
          <pre class="text-xs text-gray-600">数字值: {{ numberValue }} (类型: {{ typeof numberValue }})
字符串值: {{ stringValue }} (类型: {{ typeof stringValue }})</pre>
        </div>
      </div>
    </div>

    <!-- 复选框组 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">复选框组</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="space-y-3">
          <div class="font-medium text-gray-700 mb-2">选择你喜欢的水果：</div>
          <GlobalCheckbox
            v-model="fruits"
            true-value="apple"
            label="苹果"
          />
          <GlobalCheckbox
            v-model="fruits"
            true-value="banana"
            label="香蕉"
          />
          <GlobalCheckbox
            v-model="fruits"
            true-value="orange"
            label="橙子"
          />
          <GlobalCheckbox
            v-model="fruits"
            true-value="grape"
            label="葡萄"
          />
        </div>
        <div class="mt-4 p-3 bg-white rounded border">
          <h5 class="text-sm font-medium mb-2">选中的水果：</h5>
          <pre class="text-xs text-gray-600">{{ fruits }}</pre>
        </div>
      </div>
    </div>

    <!-- 事件测试 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">事件测试</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <GlobalCheckbox
          v-model="eventCheckbox"
          label="事件测试复选框"
          @change="handleChange"
          @click="handleClick"
        />
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
        <GlobalCheckbox
          ref="apiCheckboxRef"
          v-model="apiCheckbox"
          label="API测试复选框"
        />
        <div class="flex flex-wrap gap-2">
          <GlobalButton
            type="primary"
            size="small"
            @click="focusCheckbox"
          >
            获取焦点
          </GlobalButton>
          <GlobalButton
            type="secondary"
            size="small"
            @click="blurCheckbox"
          >
            失去焦点
          </GlobalButton>
          <GlobalButton
            type="warning"
            size="small"
            @click="toggleCheckbox"
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
                <td class="px-4 py-2 text-sm text-gray-500">boolean | string | number | any[]</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">绑定值</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">trueValue</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean | string | number</td>
                <td class="px-4 py-2 text-sm text-gray-500">true</td>
                <td class="px-4 py-2 text-sm text-gray-500">选中时的值</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">falseValue</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean | string | number</td>
                <td class="px-4 py-2 text-sm text-gray-500">false</td>
                <td class="px-4 py-2 text-sm text-gray-500">未选中时的值</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">label</td>
                <td class="px-4 py-2 text-sm text-gray-500">string</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">复选框标签</td>
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
                <td class="px-4 py-2 text-sm text-gray-500">复选框尺寸</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">color</td>
                <td class="px-4 py-2 text-sm text-gray-500">'primary' | 'success' | 'warning' | 'danger'</td>
                <td class="px-4 py-2 text-sm text-gray-500">'primary'</td>
                <td class="px-4 py-2 text-sm text-gray-500">复选框颜色</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">indeterminate</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean</td>
                <td class="px-4 py-2 text-sm text-gray-500">false</td>
                <td class="px-4 py-2 text-sm text-gray-500">是否为中间状态</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">checkedIcon</td>
                <td class="px-4 py-2 text-sm text-gray-500">string</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">自定义选中图标</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">bordered</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean</td>
                <td class="px-4 py-2 text-sm text-gray-500">false</td>
                <td class="px-4 py-2 text-sm text-gray-500">是否显示边框</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">validateStatus</td>
                <td class="px-4 py-2 text-sm text-gray-500">'success' | 'warning' | 'error' | ''</td>
                <td class="px-4 py-2 text-sm text-gray-500">''</td>
                <td class="px-4 py-2 text-sm text-gray-500">校验状态</td>
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
                <td class="px-4 py-2 text-sm text-gray-500">(value: boolean | string | number | any[])</td>
                <td class="px-4 py-2 text-sm text-gray-500">值变化时触发</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">change</td>
                <td class="px-4 py-2 text-sm text-gray-500">(value: boolean | string | number | any[], event: Event)</td>
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
                <td class="px-4 py-2 text-sm text-gray-500">让复选框获得焦点</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">blur</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">让复选框失去焦点</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">toggle</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">切换选中状态</td>
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
  &lt;!-- 基础使用 --&gt;
  &lt;GlobalCheckbox v-model="checked" label="同意协议" /&gt;

  &lt;!-- 复选框组 --&gt;
  &lt;GlobalCheckbox v-model="hobbies" true-value="reading" label="阅读" /&gt;
  &lt;GlobalCheckbox v-model="hobbies" true-value="music" label="音乐" /&gt;
  &lt;GlobalCheckbox v-model="hobbies" true-value="sports" label="运动" /&gt;
&lt;/template&gt;

&lt;script setup&gt;
import { ref } from 'vue'

const checked = ref(false)
const hobbies = ref([])
&lt;/script&gt;</code></pre>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import type { CheckboxInstance } from '../../../components/Checkbox/types'

defineOptions({
  name: 'CheckboxTest'
})

// 基础复选框
const basicCheckbox = ref(false)
const defaultChecked = ref(true)
const disabledUnchecked = ref(false)
const disabledChecked = ref(true)
const indeterminateCheckbox = ref(false)

// 复选框尺寸
const smallCheckbox = ref(true)
const mediumCheckbox = ref(true)
const largeCheckbox = ref(true)

// 复选框颜色
const primaryCheckbox = ref(true)
const successCheckbox = ref(true)
const warningCheckbox = ref(true)
const dangerCheckbox = ref(true)

// 自定义图标
const heartCheckbox = ref(false)
const starCheckbox = ref(true)
const thumbsUpCheckbox = ref(false)

// 边框样式
const borderedCheckbox1 = ref(true)
const borderedCheckbox2 = ref(false)

// 验证状态
const successValidation = ref(true)
const warningValidation = ref(false)
const errorValidation = ref(true)

// 自定义值
const numberValue = ref(1)
const stringValue = ref('no')

// 复选框组
const fruits = ref<string[]>(['apple', 'orange'])

// 事件测试
const eventCheckbox = ref(false)
const eventLogs = ref<string[]>([])

// API测试
const apiCheckbox = ref(false)
const apiCheckboxRef = ref<CheckboxInstance>()

// 事件处理
const handleChange = (value: boolean | string | number | any[], event: Event) => {
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
const focusCheckbox = () => {
  apiCheckboxRef.value?.focus()
}

const blurCheckbox = () => {
  apiCheckboxRef.value?.blur()
}

const toggleCheckbox = () => {
  apiCheckboxRef.value?.toggle()
}
</script>

<style scoped>
/* 样式定制 */
</style>
