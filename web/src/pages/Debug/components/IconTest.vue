<template>
  <div class="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
    <h3 class="text-lg font-semibold mb-4">Icon 图标组件测试</h3>
    <p class="text-gray-600 text-sm mb-6">基于 Heroicons 的图标组件，支持 outline 和 solid 两种风格</p>

    <!-- 图标尺寸 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">图标尺寸</h4>
      <div class="flex flex-wrap items-center gap-4">
        <div class="flex items-center gap-2">
          <GlobalIcon name="star" size="xs" />
          <span class="text-sm text-gray-600">xs (12px)</span>
        </div>
        <div class="flex items-center gap-2">
          <GlobalIcon name="star" size="sm" />
          <span class="text-sm text-gray-600">sm (16px)</span>
        </div>
        <div class="flex items-center gap-2">
          <GlobalIcon name="star" size="md" />
          <span class="text-sm text-gray-600">md (20px)</span>
        </div>
        <div class="flex items-center gap-2">
          <GlobalIcon name="star" size="lg" />
          <span class="text-sm text-gray-600">lg (24px)</span>
        </div>
        <div class="flex items-center gap-2">
          <GlobalIcon name="star" size="xl" />
          <span class="text-sm text-gray-600">xl (32px)</span>
        </div>
        <div class="flex items-center gap-2">
          <GlobalIcon name="star" size="2xl" />
          <span class="text-sm text-gray-600">2xl (40px)</span>
        </div>
      </div>
    </div>

    <!-- 图标风格 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">图标风格</h4>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <h5 class="text-sm font-medium text-gray-700 mb-3">Outline (默认)</h5>
          <div class="flex flex-wrap gap-3">
            <GlobalIcon name="heart" variant="outline" size="lg" />
            <GlobalIcon name="star" variant="outline" size="lg" />
            <GlobalIcon name="user" variant="outline" size="lg" />
            <GlobalIcon name="bell" variant="outline" size="lg" />
            <GlobalIcon name="home" variant="outline" size="lg" />
          </div>
        </div>
        <div>
          <h5 class="text-sm font-medium text-gray-700 mb-3">Solid</h5>
          <div class="flex flex-wrap gap-3">
            <GlobalIcon name="heart" variant="solid" size="lg" color="text-red-500" />
            <GlobalIcon name="star" variant="solid" size="lg" color="text-yellow-500" />
            <GlobalIcon name="user" variant="solid" size="lg" color="text-blue-500" />
            <GlobalIcon name="bell" variant="solid" size="lg" color="text-green-500" />
            <GlobalIcon name="home" variant="solid" size="lg" color="text-purple-500" />
          </div>
        </div>
      </div>
    </div>

    <!-- 彩色图标 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">彩色图标</h4>
      <div class="flex flex-wrap gap-3">
        <GlobalIcon name="check-circle" size="lg" color="text-green-500" />
        <GlobalIcon name="x-circle" size="lg" color="text-red-500" />
        <GlobalIcon name="exclamation-triangle" size="lg" color="text-yellow-500" />
        <GlobalIcon name="information-circle" size="lg" color="text-blue-500" />
        <GlobalIcon name="heart" variant="solid" size="lg" color="text-pink-500" />
        <GlobalIcon name="star" variant="solid" size="lg" color="text-yellow-400" />
      </div>
    </div>

    <!-- 常用图标分类展示 -->
    <div class="space-y-8">
      <!-- 基础操作图标 -->
      <div>
        <h4 class="font-medium text-gray-900 mb-4">基础操作</h4>
        <div class="grid grid-cols-8 md:grid-cols-12 lg:grid-cols-16 gap-4">
          <div v-for="iconName in basicIcons" :key="iconName" class="text-center">
            <div class="p-2 hover:bg-gray-50 rounded-md cursor-pointer" @click="copyIconName(iconName)">
              <GlobalIcon :name="iconName" size="lg" class="mx-auto mb-1" />
              <span class="text-xs text-gray-600 block truncate">{{ iconName }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- 箭头和导航 -->
      <div>
        <h4 class="font-medium text-gray-900 mb-4">箭头和导航</h4>
        <div class="grid grid-cols-8 md:grid-cols-12 lg:grid-cols-16 gap-4">
          <div v-for="iconName in navigationIcons" :key="iconName" class="text-center">
            <div class="p-2 hover:bg-gray-50 rounded-md cursor-pointer" @click="copyIconName(iconName)">
              <GlobalIcon :name="iconName" size="lg" class="mx-auto mb-1" />
              <span class="text-xs text-gray-600 block truncate">{{ iconName }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- 用户和社交 -->
      <div>
        <h4 class="font-medium text-gray-900 mb-4">用户和社交</h4>
        <div class="grid grid-cols-8 md:grid-cols-12 lg:grid-cols-16 gap-4">
          <div v-for="iconName in socialIcons" :key="iconName" class="text-center">
            <div class="p-2 hover:bg-gray-50 rounded-md cursor-pointer" @click="copyIconName(iconName)">
              <GlobalIcon :name="iconName" size="lg" class="mx-auto mb-1" />
              <span class="text-xs text-gray-600 block truncate">{{ iconName }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- 状态和反馈 -->
      <div>
        <h4 class="font-medium text-gray-900 mb-4">状态和反馈</h4>
        <div class="grid grid-cols-8 md:grid-cols-12 lg:grid-cols-16 gap-4">
          <div v-for="iconName in statusIcons" :key="iconName" class="text-center">
            <div class="p-2 hover:bg-gray-50 rounded-md cursor-pointer" @click="copyIconName(iconName)">
              <GlobalIcon :name="iconName" size="lg" class="mx-auto mb-1" />
              <span class="text-xs text-gray-600 block truncate">{{ iconName }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- 更多常用图标 -->
      <div>
        <h4 class="font-medium text-gray-900 mb-4">更多常用图标</h4>
        <div class="grid grid-cols-8 md:grid-cols-12 lg:grid-cols-16 gap-4">
          <div v-for="iconName in moreIcons" :key="iconName" class="text-center">
            <div class="p-2 hover:bg-gray-50 rounded-md cursor-pointer" @click="copyIconName(iconName)">
              <GlobalIcon :name="iconName" size="lg" class="mx-auto mb-1" />
              <span class="text-xs text-gray-600 block truncate">{{ iconName }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 使用说明 -->
    <div class="mt-8 bg-blue-50 border border-blue-200 rounded-lg p-4">
      <h4 class="font-medium text-blue-900 mb-2">使用方式</h4>
      <div class="text-sm text-blue-800 space-y-1">
        <p>• 基础用法: <code class="bg-blue-100 px-1 rounded">&lt;GlobalIcon name="star" /&gt;</code></p>
        <p>• 设置尺寸: <code class="bg-blue-100 px-1 rounded">&lt;GlobalIcon name="star" size="lg" /&gt;</code></p>
        <p>• 设置风格: <code class="bg-blue-100 px-1 rounded">&lt;GlobalIcon name="star" variant="solid" /&gt;</code></p>
        <p>• 设置颜色: <code class="bg-blue-100 px-1 rounded">&lt;GlobalIcon name="star" color="text-red-500" /&gt;</code></p>
        <p>• 点击图标名称可复制到剪贴板</p>
      </div>
    </div>

    <!-- 扩展更多图标说明 -->
    <div class="mt-6 bg-amber-50 border border-amber-200 rounded-lg p-6">
      <div class="flex items-center gap-2 mb-4">
        <GlobalIcon name="plus" size="lg" color="text-amber-600" />
        <h4 class="font-semibold text-amber-900">如何添加更多图标</h4>
      </div>

      <div class="space-y-4 text-sm text-amber-800">
        <div>
          <h5 class="font-medium text-amber-900 mb-2">📋 方法一：添加 Heroicons 现有图标</h5>
          <ol class="list-decimal list-inside space-y-2 ml-4">
            <li>
              <strong>浏览图标库：</strong>访问
              <a href="https://heroicons.com/" target="_blank" class="text-amber-700 underline hover:text-amber-800">
                heroicons.com
              </a>
              查看所有可用图标
            </li>
            <li>
              <strong>更新类型定义：</strong>在
              <code class="bg-amber-100 px-1 rounded">src/components/Icon/types.ts</code>
              中的 <code class="bg-amber-100 px-1 rounded">IconName</code> 类型添加新图标名称
            </li>
            <li>
              <strong>导入图标组件：</strong>在
              <code class="bg-amber-100 px-1 rounded">src/components/Icon/index.vue</code>
              中导入对应的图标组件
            </li>
            <li>
              <strong>添加到映射：</strong>在 <code class="bg-amber-100 px-1 rounded">iconMap</code>
              中添加 outline 和 solid 版本（如果有）
            </li>
          </ol>
        </div>

        <div>
          <h5 class="font-medium text-amber-900 mb-2">🔧 方法二：添加自定义 SVG 图标</h5>
          <ol class="list-decimal list-inside space-y-2 ml-4">
            <li>
              <strong>准备 SVG：</strong>确保 SVG 图标符合 24x24 尺寸规范，使用
              <code class="bg-amber-100 px-1 rounded">currentColor</code> 作为颜色
            </li>
            <li>
              <strong>创建 Vue 组件：</strong>将 SVG 转换为 Vue 函数式组件
            </li>
            <li>
              <strong>按照上述步骤：</strong>添加类型定义、导入组件、更新映射
            </li>
          </ol>
        </div>

        <div>
          <h5 class="font-medium text-amber-900 mb-2">💡 示例：添加新图标</h5>
          <div class="bg-amber-100 p-3 rounded border">
            <p class="mb-2 font-medium">1. 更新 types.ts:</p>
            <pre class="text-xs bg-white p-2 rounded overflow-x-auto"><code>export type IconName =
  // ... 现有图标
  | 'new-icon-name'  // 添加新图标名称</code></pre>

            <p class="mb-2 font-medium mt-3">2. 在 index.vue 中导入:</p>
            <pre class="text-xs bg-white p-2 rounded overflow-x-auto"><code>import { NewIconName } from '@heroicons/vue/24/outline'</code></pre>

            <p class="mb-2 font-medium mt-3">3. 添加到映射:</p>
            <pre class="text-xs bg-white p-2 rounded overflow-x-auto"><code>const iconMap = {
  outline: {
    // ... 现有图标
    'new-icon-name': NewIconName
  }
}</code></pre>
          </div>
        </div>

        <div class="bg-amber-100 border border-amber-300 rounded p-3">
          <div class="flex items-start gap-2">
            <GlobalIcon name="exclamation-triangle" size="sm" color="text-amber-700" class="mt-0.5 flex-shrink-0" />
            <div>
              <p class="font-medium text-amber-900 mb-1">⚠️ 注意事项</p>
              <ul class="text-xs space-y-1">
                <li>• 确保图标名称使用 kebab-case 格式（如 arrow-up）</li>
                <li>• Solid 版本的图标可能不是所有图标都有，可以只添加 outline 版本</li>
                <li>• 添加后记得在这个测试页面中展示新图标</li>
                <li>• 重启开发服务器以确保 TypeScript 类型更新生效</li>
              </ul>
            </div>
          </div>
        </div>

        <div>
          <h5 class="font-medium text-amber-900 mb-2">🚀 推荐资源</h5>
          <ul class="space-y-1 ml-4">
            <li>
              • <strong>Heroicons:</strong>
              <a href="https://heroicons.com/" target="_blank" class="text-amber-700 underline hover:text-amber-800">
                官方图标库
              </a>
            </li>
            <li>
              • <strong>Tabler Icons:</strong>
              <a href="https://tabler-icons.io/" target="_blank" class="text-amber-700 underline hover:text-amber-800">
                免费 SVG 图标库
              </a>
            </li>
            <li>
              • <strong>Lucide:</strong>
              <a href="https://lucide.dev/" target="_blank" class="text-amber-700 underline hover:text-amber-800">
                美观的开源图标库
              </a>
            </li>
            <li>
              • <strong>Feather Icons:</strong>
              <a href="https://feathericons.com/" target="_blank" class="text-amber-700 underline hover:text-amber-800">
                简洁的开源图标
              </a>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useMessage } from '../../../composables/useMessage'

defineOptions({
  name: 'IconTest'
})

const { success } = useMessage()

// 图标分类
const basicIcons = [
  'plus', 'minus', 'x-mark', 'check', 'pencil', 'trash',
  'eye', 'eye-slash', 'heart', 'star', 'share', 'download',
  'upload', 'refresh', 'settings', 'search'
]

const navigationIcons = [
  'arrow-up', 'arrow-down', 'arrow-left', 'arrow-right',
  'chevron-up', 'chevron-down', 'chevron-left', 'chevron-right',
  'arrow-up-right', 'arrow-down-right', 'arrow-path'
]

const socialIcons = [
  'user', 'user-group', 'user-plus', 'user-minus',
  'chat-bubble-left', 'chat-bubble-left-right', 'phone',
  'envelope', 'bell', 'bell-slash'
]

const statusIcons = [
  'check-circle', 'x-circle', 'exclamation-triangle',
  'information-circle', 'question-mark-circle', 'shield-check',
  'shield-exclamation', 'lock-closed', 'lock-open'
]

// 更多分类图标示例
const moreIcons = [
  'document', 'folder', 'photo', 'home', 'calendar', 'clock',
  'shopping-cart', 'gift', 'chart-bar', 'server', 'cloud', 'wifi'
]

// 复制图标名称
const copyIconName = async (iconName: string) => {
  try {
    await navigator.clipboard.writeText(iconName)
    success(`已复制图标名称: ${iconName}`)
  } catch (err) {
    console.error('复制失败:', err)
  }
}
</script>
