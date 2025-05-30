<template>
  <div class="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
    <h3 class="text-lg font-semibold mb-4">Tag 标签组件测试</h3>
    <p class="text-gray-600 text-sm mb-6">测试标签组件的各种类型、尺寸和功能</p>

    <!-- 基础标签 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">基础标签</h4>
      <div class="bg-gray-50 p-4 rounded-md">
        <div class="flex flex-wrap gap-2">
          <GlobalTag text="默认标签" />
          <GlobalTag text="主要标签" type="primary" />
          <GlobalTag text="成功标签" type="success" />
          <GlobalTag text="警告标签" type="warning" />
          <GlobalTag text="危险标签" type="danger" />
          <GlobalTag text="信息标签" type="info" />
        </div>
      </div>
    </div>

    <!-- 标签变体 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">标签变体</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-3">
        <div>
          <h5 class="text-sm font-medium text-gray-700 mb-2">填充样式 (filled)</h5>
          <div class="flex flex-wrap gap-2">
            <GlobalTag text="默认" variant="filled" />
            <GlobalTag text="主要" type="primary" variant="filled" />
            <GlobalTag text="成功" type="success" variant="filled" />
            <GlobalTag text="警告" type="warning" variant="filled" />
          </div>
        </div>
        <div>
          <h5 class="text-sm font-medium text-gray-700 mb-2">轮廓样式 (outlined)</h5>
          <div class="flex flex-wrap gap-2">
            <GlobalTag text="默认" variant="outlined" />
            <GlobalTag text="主要" type="primary" variant="outlined" />
            <GlobalTag text="成功" type="success" variant="outlined" />
            <GlobalTag text="警告" type="warning" variant="outlined" />
          </div>
        </div>
        <div>
          <h5 class="text-sm font-medium text-gray-700 mb-2">幽灵样式 (ghost)</h5>
          <div class="flex flex-wrap gap-2">
            <GlobalTag text="默认" variant="ghost" />
            <GlobalTag text="主要" type="primary" variant="ghost" />
            <GlobalTag text="成功" type="success" variant="ghost" />
            <GlobalTag text="警告" type="warning" variant="ghost" />
          </div>
        </div>
      </div>
    </div>

    <!-- 标签尺寸 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">标签尺寸</h4>
      <div class="bg-gray-50 p-4 rounded-md">
        <div class="flex flex-wrap items-center gap-2">
          <GlobalTag text="小标签" size="small" type="primary" />
          <GlobalTag text="中等标签" size="medium" type="primary" />
          <GlobalTag text="大标签" size="large" type="primary" />
        </div>
      </div>
    </div>

    <!-- 可关闭标签 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">可关闭标签</h4>
      <div class="bg-gray-50 p-4 rounded-md">
        <div class="flex flex-wrap gap-2">
          <GlobalTag
            v-for="(tag, index) in closableTags"
            :key="tag"
            :text="tag"
            :closable="true"
            type="primary"
            @close="removeTag(index)"
          />
        </div>
        <GlobalButton
          v-if="closableTags.length === 0"
          type="secondary"
          size="small"
          @click="resetClosableTags"
          class="mt-2"
        >
          重置标签
        </GlobalButton>
      </div>
    </div>

    <!-- 图标标签 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">图标标签</h4>
      <div class="bg-gray-50 p-4 rounded-md">
        <div class="flex flex-wrap gap-2">
          <GlobalTag text="用户" icon="user" type="primary" />
          <GlobalTag text="星标" icon="star" type="warning" />
          <GlobalTag text="邮件" icon="envelope" type="info" />
          <GlobalTag text="设置" icon="cog-6-tooth" type="success" />
        </div>
      </div>
    </div>

    <!-- 圆角标签 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">圆角标签</h4>
      <div class="bg-gray-50 p-4 rounded-md">
        <div class="flex flex-wrap gap-2">
          <GlobalTag text="圆角标签" :round="true" type="primary" />
          <GlobalTag text="圆角可关闭" :round="true" :closable="true" type="success" />
          <GlobalTag text="圆角图标" :round="true" icon="heart" type="danger" />
        </div>
      </div>
    </div>

    <!-- 可选择标签 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">可选择标签</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <div class="flex flex-wrap gap-2">
          <GlobalTag
            v-for="tag in selectableTags"
            :key="tag.name"
            :text="tag.name"
            :checkable="true"
            :checked="tag.checked"
            type="primary"
            @update:checked="tag.checked = $event"
          />
        </div>
        <div class="mt-4 p-3 bg-white rounded border">
          <h5 class="text-sm font-medium mb-2">选中的标签：</h5>
          <pre class="text-xs text-gray-600">{{ selectableTags.filter(t => t.checked).map(t => t.name) }}</pre>
        </div>
      </div>
    </div>

    <!-- 自定义颜色 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">自定义颜色</h4>
      <div class="bg-gray-50 p-4 rounded-md">
        <div class="flex flex-wrap gap-2">
          <GlobalTag text="紫色" color="#9333ea" variant="filled" />
          <GlobalTag text="粉色" color="#ec4899" variant="outlined" />
          <GlobalTag text="青色" color="#06b6d4" variant="ghost" />
          <GlobalTag text="橙色" color="#f97316" variant="filled" :closable="true" />
        </div>
      </div>
    </div>

    <!-- 禁用状态 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">禁用状态</h4>
      <div class="bg-gray-50 p-4 rounded-md">
        <div class="flex flex-wrap gap-2">
          <GlobalTag text="禁用标签" :disabled="true" />
          <GlobalTag text="禁用主要" type="primary" :disabled="true" />
          <GlobalTag text="禁用可关闭" :closable="true" :disabled="true" />
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
                <td class="px-4 py-2 text-sm font-medium text-gray-900">text</td>
                <td class="px-4 py-2 text-sm text-gray-500">string</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">标签文本</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">type</td>
                <td class="px-4 py-2 text-sm text-gray-500">'default' | 'primary' | 'success' | 'warning' | 'danger' | 'info'</td>
                <td class="px-4 py-2 text-sm text-gray-500">'default'</td>
                <td class="px-4 py-2 text-sm text-gray-500">标签类型</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">size</td>
                <td class="px-4 py-2 text-sm text-gray-500">'small' | 'medium' | 'large'</td>
                <td class="px-4 py-2 text-sm text-gray-500">'medium'</td>
                <td class="px-4 py-2 text-sm text-gray-500">标签尺寸</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">variant</td>
                <td class="px-4 py-2 text-sm text-gray-500">'filled' | 'outlined' | 'ghost'</td>
                <td class="px-4 py-2 text-sm text-gray-500">'filled'</td>
                <td class="px-4 py-2 text-sm text-gray-500">标签变体</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">closable</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean</td>
                <td class="px-4 py-2 text-sm text-gray-500">false</td>
                <td class="px-4 py-2 text-sm text-gray-500">是否可关闭</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">checkable</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean</td>
                <td class="px-4 py-2 text-sm text-gray-500">false</td>
                <td class="px-4 py-2 text-sm text-gray-500">是否可选择</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">checked</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean</td>
                <td class="px-4 py-2 text-sm text-gray-500">false</td>
                <td class="px-4 py-2 text-sm text-gray-500">是否选中</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">icon</td>
                <td class="px-4 py-2 text-sm text-gray-500">string</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">前缀图标</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">round</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean</td>
                <td class="px-4 py-2 text-sm text-gray-500">false</td>
                <td class="px-4 py-2 text-sm text-gray-500">是否圆角</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">color</td>
                <td class="px-4 py-2 text-sm text-gray-500">string</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">自定义颜色</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">disabled</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean</td>
                <td class="px-4 py-2 text-sm text-gray-500">false</td>
                <td class="px-4 py-2 text-sm text-gray-500">是否禁用</td>
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
                <td class="px-4 py-2 text-sm font-medium text-gray-900">close</td>
                <td class="px-4 py-2 text-sm text-gray-500">(event: MouseEvent)</td>
                <td class="px-4 py-2 text-sm text-gray-500">关闭时触发</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">click</td>
                <td class="px-4 py-2 text-sm text-gray-500">(event: MouseEvent)</td>
                <td class="px-4 py-2 text-sm text-gray-500">点击时触发</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">update:checked</td>
                <td class="px-4 py-2 text-sm text-gray-500">(checked: boolean)</td>
                <td class="px-4 py-2 text-sm text-gray-500">选中状态改变时触发</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">change</td>
                <td class="px-4 py-2 text-sm text-gray-500">(checked: boolean, event: MouseEvent)</td>
                <td class="px-4 py-2 text-sm text-gray-500">选中状态改变时触发</td>
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
  &lt;!-- 基础标签 --&gt;
  &lt;GlobalTag text="标签文本" type="primary" /&gt;

  &lt;!-- 可关闭标签 --&gt;
  &lt;GlobalTag
    text="可关闭标签"
    :closable="true"
    @close="handleClose"
  /&gt;

  &lt;!-- 可选择标签 --&gt;
  &lt;GlobalTag
    text="可选择标签"
    :checkable="true"
    :checked="isChecked"
    @update:checked="isChecked = $event"
  /&gt;
&lt;/template&gt;

&lt;script setup&gt;
import { ref } from 'vue'

const isChecked = ref(false)

const handleClose = () => {
  console.log('标签被关闭')
}
&lt;/script&gt;</code></pre>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

defineOptions({
  name: 'TagTest'
})

// 可关闭标签
const closableTags = ref(['标签1', '标签2', '标签3', '标签4'])

// 可选择标签
const selectableTags = ref([
  { name: 'Vue.js', checked: true },
  { name: 'React', checked: false },
  { name: 'Angular', checked: true },
  { name: 'Svelte', checked: false },
])

// 方法
const removeTag = (index: number) => {
  closableTags.value.splice(index, 1)
}

const resetClosableTags = () => {
  closableTags.value = ['标签1', '标签2', '标签3', '标签4']
}
</script>

<style scoped>
/* 样式定制 */
</style>
