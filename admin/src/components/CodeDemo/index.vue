<template>
  <div class="code-demo">
    <!-- 演示区域 -->
    <div class="code-demo__showcase">
      <div class="code-demo__header">
        <h4 v-if="title" class="code-demo__title">{{ title }}</h4>
        <p v-if="description" class="code-demo__description">{{ description }}</p>
      </div>

      <div class="code-demo__content">
        <slot />
      </div>
    </div>

    <!-- 操作按钮 -->
    <div class="code-demo__actions">
      <button
        @click="toggleCode"
        class="code-demo__toggle-btn"
        :class="{ 'code-demo__toggle-btn--active': showCode }"
      >
        <GlobalIcon :name="showCode ? 'chevron-up' : 'code-bracket'" size="sm" />
        <span>{{ showCode ? '隐藏代码' : '查看代码' }}</span>
      </button>

      <button
        v-if="showCode"
        @click="copyCode"
        class="code-demo__copy-btn"
        :class="{ 'code-demo__copy-btn--copied': copied }"
      >
        <GlobalIcon :name="copied ? 'check' : 'document-duplicate'" size="sm" />
        <span>{{ copied ? '已复制' : '复制代码' }}</span>
      </button>
    </div>

    <!-- 代码区域 -->
    <Transition name="code-expand">
      <div v-if="showCode" class="code-demo__code-area">
        <div class="code-demo__code-header">
          <span class="code-demo__code-title">{{ codeTitle || 'Vue Template' }}</span>
          <div class="code-demo__code-lang">{{ language || 'vue' }}</div>
        </div>

        <div class="code-demo__code-content">
          <pre><code ref="codeRef" :class="`language-${language || 'vue'}`" v-html="highlightedCode"></code></pre>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, nextTick, onMounted, watch } from 'vue'
import { useClipboard } from '@vueuse/core'
import Prism from 'prismjs'
import 'prismjs/themes/prism-dark.css'
import 'prismjs/components/prism-javascript'
import 'prismjs/components/prism-typescript'
import 'prismjs/components/prism-css'
import 'prismjs/components/prism-markup'
import 'prismjs/components/prism-markup-templating'

interface CodeDemoProps {
  title?: string
  description?: string
  code: string
  language?: 'vue' | 'javascript' | 'typescript' | 'html' | 'css'
  codeTitle?: string
}

const props = withDefaults(defineProps<CodeDemoProps>(), {
  language: 'vue'
})

const showCode = ref(false)
const codeRef = ref<HTMLElement>()

// 使用 VueUse 的 clipboard 功能
const { copy, copied } = useClipboard({ source: computed(() => props.code) })

// 格式化代码（去除多余的缩进）
const formattedCode = computed(() => {
  const lines = props.code.split('\n')

  // 移除开头和结尾的空行
  while (lines.length > 0 && lines[0].trim() === '') {
    lines.shift()
  }
  while (lines.length > 0 && lines[lines.length - 1].trim() === '') {
    lines.pop()
  }

  if (lines.length === 0) return ''

  // 找到最小的缩进量
  const minIndent = lines
    .filter(line => line.trim() !== '')
    .reduce((min, line) => {
      const indent = line.match(/^(\s*)/)?.[1].length || 0
      return Math.min(min, indent)
    }, Infinity)

  // 移除多余的缩进
  return lines
    .map(line => line.slice(minIndent))
    .join('\n')
})

// 语法高亮后的代码
const highlightedCode = computed(() => {
  if (!formattedCode.value) return ''

  try {
    // 对于Vue模板，使用HTML语法高亮
    const language = props.language === 'vue' ? 'html' : props.language
    return Prism.highlight(formattedCode.value, Prism.languages[language] || Prism.languages.html, language)
  } catch (error) {
    console.warn('Code highlighting failed:', error)
    return formattedCode.value
  }
})

const toggleCode = () => {
  showCode.value = !showCode.value
}

const copyCode = async () => {
  try {
    await copy()
  } catch (error) {
    console.error('复制失败:', error)
  }
}

// 监听代码变化，重新高亮
watch([showCode, formattedCode], () => {
  if (showCode.value) {
    nextTick(() => {
      if (codeRef.value) {
        // Prism 已经在计算属性中处理了高亮
      }
    })
  }
})
</script>

<style scoped>
.code-demo {
  @apply border border-gray-200 rounded-lg overflow-hidden bg-white shadow-sm mb-6;
}

.code-demo__showcase {
  @apply p-6;
}

.code-demo__header {
  @apply mb-4;
}

.code-demo__title {
  @apply text-lg font-semibold text-gray-900 mb-2;
}

.code-demo__description {
  @apply text-gray-600 text-sm;
}

.code-demo__content {
  @apply min-h-[60px];
}

.code-demo__actions {
  @apply flex items-center justify-center gap-2 px-6 py-3 bg-gray-50 border-t border-gray-200;
}

.code-demo__toggle-btn,
.code-demo__copy-btn {
  @apply inline-flex items-center gap-2 px-3 py-1.5 text-sm font-medium rounded-md transition-all duration-200;
}

.code-demo__toggle-btn {
  @apply text-gray-600 hover:text-blue-600 hover:bg-blue-50;
}

.code-demo__toggle-btn--active {
  @apply text-blue-600 bg-blue-50;
}

.code-demo__copy-btn {
  @apply text-gray-600 hover:text-green-600 hover:bg-green-50;
}

.code-demo__copy-btn--copied {
  @apply text-green-600 bg-green-50;
}

.code-demo__code-area {
  @apply border-t border-gray-200 bg-gray-900;
}

.code-demo__code-header {
  @apply flex items-center justify-between px-4 py-2 bg-gray-800 border-b border-gray-700;
}

.code-demo__code-title {
  @apply text-sm text-gray-300 font-medium;
}

.code-demo__code-lang {
  @apply text-xs text-gray-400 bg-gray-700 px-2 py-1 rounded;
}

.code-demo__code-content {
  @apply overflow-x-auto;
}

.code-demo__code-content pre {
  @apply m-0 p-4 bg-transparent text-sm leading-relaxed;
}

.code-demo__code-content code {
  @apply font-mono;
}

/* 代码展开动画 */
.code-expand-enter-active,
.code-expand-leave-active {
  @apply transition-all duration-300 ease-in-out;
  overflow: hidden;
}

.code-expand-enter-from,
.code-expand-leave-to {
  max-height: 0;
  opacity: 0;
}

.code-expand-enter-to,
.code-expand-leave-from {
  max-height: 1000px;
  opacity: 1;
}

/* Prism.js 会自动处理高亮样式 */
.code-demo__code-content :deep(.token.comment),
.code-demo__code-content :deep(.token.prolog),
.code-demo__code-content :deep(.token.doctype),
.code-demo__code-content :deep(.token.cdata) {
  color: #6a737d;
}

.code-demo__code-content :deep(.token.punctuation) {
  color: #d1d5da;
}

.code-demo__code-content :deep(.token.property),
.code-demo__code-content :deep(.token.tag),
.code-demo__code-content :deep(.token.boolean),
.code-demo__code-content :deep(.token.number),
.code-demo__code-content :deep(.token.constant),
.code-demo__code-content :deep(.token.symbol),
.code-demo__code-content :deep(.token.deleted) {
  color: #f97583;
}

.code-demo__code-content :deep(.token.selector),
.code-demo__code-content :deep(.token.attr-name),
.code-demo__code-content :deep(.token.string),
.code-demo__code-content :deep(.token.char),
.code-demo__code-content :deep(.token.builtin),
.code-demo__code-content :deep(.token.inserted) {
  color: #9ecbff;
}

.code-demo__code-content :deep(.token.operator),
.code-demo__code-content :deep(.token.entity),
.code-demo__code-content :deep(.token.url),
.code-demo__code-content :deep(.language-css .token.string),
.code-demo__code-content :deep(.style .token.string) {
  color: #79b8ff;
}

.code-demo__code-content :deep(.token.atrule),
.code-demo__code-content :deep(.token.attr-value),
.code-demo__code-content :deep(.token.keyword) {
  color: #b392f0;
}

.code-demo__code-content :deep(.token.function),
.code-demo__code-content :deep(.token.class-name) {
  color: #ffab70;
}

.code-demo__code-content :deep(.token.regex),
.code-demo__code-content :deep(.token.important),
.code-demo__code-content :deep(.token.variable) {
  color: #ffab70;
}
</style>
