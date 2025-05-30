<template>
  <div class="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
    <h3 class="text-lg font-semibold mb-4">InputTag 输入标签组件测试</h3>
    <p class="text-gray-600 text-sm mb-6">测试输入标签组件的各种功能和特性</p>

    <!-- 基础标签输入 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">基础标签输入</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">基础使用</label>
          <GlobalInputTag
            v-model="basicTags"
            placeholder="输入标签后按回车添加"
          />
        </div>
        <div class="mt-4 p-3 bg-white rounded border">
          <h5 class="text-sm font-medium mb-2">当前标签：</h5>
          <pre class="text-xs text-gray-600">{{ basicTags }}</pre>
        </div>
      </div>
    </div>

    <!-- 不同尺寸 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">不同尺寸</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">小尺寸</label>
          <GlobalInputTag
            v-model="smallSizeTags"
            size="small"
            tag-size="small"
            placeholder="小尺寸标签输入"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">中等尺寸</label>
          <GlobalInputTag
            v-model="mediumSizeTags"
            size="medium"
            tag-size="medium"
            placeholder="中等尺寸标签输入"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">大尺寸</label>
          <GlobalInputTag
            v-model="largeSizeTags"
            size="large"
            tag-size="large"
            placeholder="大尺寸标签输入"
          />
        </div>
      </div>
    </div>

    <!-- 标签样式 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">标签样式</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">主要色标签</label>
          <GlobalInputTag
            v-model="primaryTags"
            tag-type="primary"
            placeholder="主要色标签"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">成功色标签</label>
          <GlobalInputTag
            v-model="successTags"
            tag-type="success"
            placeholder="成功色标签"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">轮廓样式</label>
          <GlobalInputTag
            v-model="outlinedTags"
            tag-variant="outlined"
            tag-type="primary"
            placeholder="轮廓样式标签"
          />
        </div>
      </div>
    </div>

    <!-- 功能配置 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">功能配置</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">限制最大数量（最多3个）</label>
          <GlobalInputTag
            v-model="limitedTags"
            :max="3"
            placeholder="最多添加3个标签"
            :show-count="true"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">允许重复标签</label>
          <GlobalInputTag
            v-model="duplicateTags"
            :allow-duplicates="true"
            placeholder="允许重复的标签"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">自定义分隔符（空格和逗号）</label>
          <GlobalInputTag
            v-model="separatorTags"
            separator=" ,"
            placeholder="输入后按空格或逗号分隔"
          />
        </div>
      </div>
    </div>

    <!-- 验证功能 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">验证功能</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">长度限制（3-10字符）</label>
          <GlobalInputTag
            v-model="lengthValidatedTags"
            :min-tag-length="3"
            :max-tag-length="10"
            placeholder="标签长度必须3-10字符"
            @invalid="handleInvalid"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">邮箱验证</label>
          <GlobalInputTag
            v-model="emailTags"
            input-type="email"
            :validate="validateEmail"
            placeholder="输入邮箱地址"
            @invalid="handleInvalid"
          />
        </div>
        <div v-if="validationError" class="text-red-500 text-xs">
          {{ validationError }}
        </div>
      </div>
    </div>

    <!-- 建议功能 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">建议功能</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">技能标签（带建议）</label>
          <GlobalInputTag
            v-model="skillTags"
            :suggestions="skillSuggestions"
            placeholder="输入技能名称，会显示建议"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">编程语言（带图标）</label>
          <GlobalInputTag
            v-model="languageTags"
            :suggestions="languageSuggestions"
            icon="code-bracket"
            tag-type="info"
            placeholder="输入编程语言"
          />
        </div>
      </div>
    </div>

    <!-- 状态演示 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">状态演示</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">禁用状态</label>
          <GlobalInputTag
            v-model="disabledTags"
            :disabled="true"
            placeholder="禁用状态"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">只读状态</label>
          <GlobalInputTag
            v-model="readonlyTags"
            :readonly="true"
            placeholder="只读状态"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">错误状态</label>
          <GlobalInputTag
            v-model="errorTags"
            :error="true"
            placeholder="错误状态"
          />
        </div>
      </div>
    </div>

    <!-- 事件测试 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">事件测试</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">事件测试</label>
          <GlobalInputTag
            v-model="eventTags"
            placeholder="测试各种事件"
            @add="handleAdd"
            @remove="handleRemove"
            @max-reached="handleMaxReached"
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
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">API测试</label>
          <GlobalInputTag
            ref="apiInputTagRef"
            v-model="apiTags"
            placeholder="用于测试API方法"
          />
        </div>
        <div class="flex flex-wrap gap-2">
          <GlobalButton
            type="primary"
            size="small"
            @click="focusInputTag"
          >
            获取焦点
          </GlobalButton>
          <GlobalButton
            type="secondary"
            size="small"
            @click="addTestTag"
          >
            添加测试标签
          </GlobalButton>
          <GlobalButton
            type="warning"
            size="small"
            @click="clearTags"
          >
            清空标签
          </GlobalButton>
        </div>
      </div>
    </div>

    <!-- API 文档 -->
    <div class="border-t pt-8">
      <h3 class="text-lg font-semibold mb-4">API 文档</h3>

      <!-- Props - 简化版 -->
      <div class="mb-6">
        <h4 class="font-medium text-gray-900 mb-3">主要属性</h4>
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
                <td class="px-4 py-2 text-sm text-gray-500">string[]</td>
                <td class="px-4 py-2 text-sm text-gray-500">[]</td>
                <td class="px-4 py-2 text-sm text-gray-500">绑定值</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">placeholder</td>
                <td class="px-4 py-2 text-sm text-gray-500">string</td>
                <td class="px-4 py-2 text-sm text-gray-500">'请输入标签'</td>
                <td class="px-4 py-2 text-sm text-gray-500">占位符</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">max</td>
                <td class="px-4 py-2 text-sm text-gray-500">number</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">最大标签数量</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">allowDuplicates</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean</td>
                <td class="px-4 py-2 text-sm text-gray-500">false</td>
                <td class="px-4 py-2 text-sm text-gray-500">是否允许重复标签</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">separator</td>
                <td class="px-4 py-2 text-sm text-gray-500">string | RegExp</td>
                <td class="px-4 py-2 text-sm text-gray-500">','</td>
                <td class="px-4 py-2 text-sm text-gray-500">分隔符</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">validate</td>
                <td class="px-4 py-2 text-sm text-gray-500">(tag: string) => boolean | string</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">标签验证函数</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">suggestions</td>
                <td class="px-4 py-2 text-sm text-gray-500">string[]</td>
                <td class="px-4 py-2 text-sm text-gray-500">[]</td>
                <td class="px-4 py-2 text-sm text-gray-500">预设标签建议</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">showCount</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean</td>
                <td class="px-4 py-2 text-sm text-gray-500">false</td>
                <td class="px-4 py-2 text-sm text-gray-500">是否显示计数</td>
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
  &lt;GlobalInputTag
    v-model="tags"
    placeholder="输入标签"
    :max="5"
    :suggestions="['Vue', 'React', 'Angular']"
    :show-count="true"
    @add="handleAdd"
    @remove="handleRemove"
  /&gt;
&lt;/template&gt;

&lt;script setup&gt;
import { ref } from 'vue'

const tags = ref(['JavaScript', 'TypeScript'])

const handleAdd = (tag) => {
  console.log('添加标签:', tag)
}

const handleRemove = (tag, index) => {
  console.log('移除标签:', tag, index)
}
&lt;/script&gt;</code></pre>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import type { InputTagInstance } from '../../../components/InputTag/types'

defineOptions({
  name: 'InputTagTest'
})

// 基础标签输入
const basicTags = ref(['Vue.js', 'TypeScript'])

// 不同尺寸
const smallSizeTags = ref(['小', '标签'])
const mediumSizeTags = ref(['中等', '标签'])
const largeSizeTags = ref(['大', '标签'])

// 标签样式
const primaryTags = ref(['主要', '标签'])
const successTags = ref(['成功', '标签'])
const outlinedTags = ref(['轮廓', '标签'])

// 功能配置
const limitedTags = ref(['标签1', '标签2'])
const duplicateTags = ref(['重复', '重复', '标签'])
const separatorTags = ref(['空格', '逗号', '分隔'])

// 验证功能
const lengthValidatedTags = ref(['JavaScript'])
const emailTags = ref(['user@example.com'])
const validationError = ref('')

// 建议功能
const skillTags = ref(['Vue.js', 'React'])
const languageTags = ref(['JavaScript', 'Python'])

const skillSuggestions = [
  'Vue.js', 'React', 'Angular', 'Svelte', 'Node.js', 'Express', 'Koa',
  'MongoDB', 'MySQL', 'PostgreSQL', 'Redis', 'Docker', 'Kubernetes'
]

const languageSuggestions = [
  'JavaScript', 'TypeScript', 'Python', 'Java', 'C++', 'Go', 'Rust',
  'PHP', 'Ruby', 'Swift', 'Kotlin', 'C#', 'Dart'
]

// 状态演示
const disabledTags = ref(['禁用', '标签'])
const readonlyTags = ref(['只读', '标签'])
const errorTags = ref(['错误', '标签'])

// 事件测试
const eventTags = ref(['事件', '测试'])
const eventLogs = ref<string[]>([])

// API测试
const apiTags = ref(['API', '测试'])
const apiInputTagRef = ref<InputTagInstance>()

// 验证函数
const validateEmail = (tag: string): boolean | string => {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  return emailRegex.test(tag) || '请输入有效的邮箱地址'
}

// 事件处理
const handleInvalid = (tag: string, error: string) => {
  validationError.value = `标签 "${tag}" 验证失败: ${error}`
  setTimeout(() => {
    validationError.value = ''
  }, 3000)
}

const handleAdd = (tag: string) => {
  const timestamp = new Date().toLocaleTimeString()
  eventLogs.value.unshift(`${timestamp} - 添加标签: ${tag}`)
  if (eventLogs.value.length > 5) {
    eventLogs.value = eventLogs.value.slice(0, 5)
  }
}

const handleRemove = (tag: string, index: number) => {
  const timestamp = new Date().toLocaleTimeString()
  eventLogs.value.unshift(`${timestamp} - 移除标签: ${tag} (索引: ${index})`)
  if (eventLogs.value.length > 5) {
    eventLogs.value = eventLogs.value.slice(0, 5)
  }
}

const handleMaxReached = (tags: string[]) => {
  const timestamp = new Date().toLocaleTimeString()
  eventLogs.value.unshift(`${timestamp} - 达到最大数量限制`)
  if (eventLogs.value.length > 5) {
    eventLogs.value = eventLogs.value.slice(0, 5)
  }
}

// API 方法测试
const focusInputTag = () => {
  apiInputTagRef.value?.focus()
}

const addTestTag = () => {
  const result = apiInputTagRef.value?.addTag('新标签')
  console.log('添加结果:', result)
}

const clearTags = () => {
  apiInputTagRef.value?.clear()
}
</script>

<style scoped>
/* 样式定制 */
</style>
