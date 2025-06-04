<template>
  <div class="h-[calc(100vh-175px)] bg-gray-50 overflow-hidden">
      <!-- 左右布局容器 -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
        <div class="flex h-[calc(100vh-175px)]">
          <!-- 左侧组件导航 - 固定不滚动 -->
          <div class="w-64 bg-gray-50 border-r border-gray-200 flex flex-col">
            <!-- 导航头部 -->
            <div class="px-4 py-4 border-b border-gray-200 bg-white">
              <h2 class="text-lg font-semibold text-gray-900">组件列表</h2>
              <p class="text-xs text-gray-500 mt-1">选择要测试的组件</p>
            </div>

            <!-- 导航菜单 - 固定不滚动 -->
            <nav class="flex-1 p-3 overflow-y-scroll">
              <div class="space-y-1">
                <button
                  v-for="tab in tabs"
                  :key="tab.key"
                  @click="activeTab = tab.key"
                  :class="[
                    'w-full flex items-center px-3 py-2.5 text-sm font-medium rounded-lg transition-all duration-200 text-left group',
                    activeTab === tab.key
                      ? 'bg-blue-100 text-blue-700 shadow-sm border border-blue-200'
                      : 'text-gray-600 hover:text-gray-900 hover:bg-white hover:shadow-sm border border-transparent'
                  ]"
                >
                  <div
                    :class="[
                      'flex items-center justify-center w-8 h-8 rounded-md mr-3 transition-colors',
                      activeTab === tab.key
                        ? `bg-${getColorClass(tab.color)}-500`
                        : 'bg-gray-300 group-hover:bg-gray-400'
                    ]"
                  >
                    <GlobalIcon
                      :name="tab.icon"
                      size="sm"
                      :color="activeTab === tab.key ? 'text-white' : 'text-white'"
                    />
                  </div>
                  <div class="flex-1">
                    <div class="font-medium">{{ tab.name }}</div>
                    <div class="text-xs text-gray-500 mt-0.5 truncate">
                      {{ tab.shortDesc }}
                    </div>
                  </div>
                </button>
              </div>
            </nav>

            <!-- 导航底部信息 -->
            <div class="px-4 py-3 border-t border-gray-200 bg-white">
              <div class="text-xs text-gray-500">
                <div class="flex justify-between items-center">
                  <span>总计组件</span>
                  <span class="font-medium">{{ tabs.length }}</span>
                </div>
                <div class="flex justify-between items-center mt-1">
                  <span>当前组件</span>
                  <span class="font-medium text-blue-600">{{ currentTab.name }}</span>
                </div>
              </div>
            </div>
          </div>

          <!-- 右侧内容区域 - 可滚动 -->
          <div class="flex-1 flex flex-col">
            <!-- 内容头部 - 固定 -->
            <div class="px-6 py-4 border-b border-gray-200 bg-white">
              <div class="flex items-center">
                <div
                  :class="[
                    'flex items-center justify-center w-12 h-12 rounded-xl mr-4 shadow-sm',
                    `bg-${getColorClass(currentTab.color)}-100`
                  ]"
                >
                  <GlobalIcon
                    :name="currentTab.icon"
                    size="lg"
                    :color="currentTab.color"
                  />
                </div>
                <div>
                  <h3 class="text-xl font-bold text-gray-900">{{ currentTab.title }}</h3>
                  <p class="text-gray-600 mt-1">{{ currentTab.description }}</p>
                </div>
              </div>
            </div>

            <!-- 组件内容区域 - 可滚动 -->
            <div class="flex-1 overflow-y-auto">
              <div class="p-6">
                <KeepAlive>
                  <component :is="currentComponent" />
                </KeepAlive>
              </div>
            </div>
          </div>
        </div>
      </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import * as IconTestModule from './components/IconTest.vue'
import * as ButtonTestModule from './components/ButtonTest.vue'
import * as SelectTestModule from './components/SelectTest.vue'
import * as DialogTestModule from './components/DialogTest.vue'
import * as UploadTestModule from './components/UploadTest.vue'
import * as MessageTestModule from './components/MessageTest.vue'
import * as StorageTestModule from './components/StorageTest.vue'
import * as PaginationTestModule from './components/PaginationTest.vue'
import * as TableTestModule from './components/TableTest.vue'
import * as FormTestModule from './components/FormTest.vue'
import * as InputTestModule from './components/InputTest.vue'
import * as TextareaTestModule from './components/TextareaTest.vue'
import * as SwitchTestModule from './components/SwitchTest.vue'
import * as CheckboxTestModule from './components/CheckboxTest.vue'
import * as TagTestModule from './components/TagTest.vue'
import * as InputTagTestModule from './components/InputTagTest.vue'
import * as TabsTestModule from './components/TabsTest.vue'
import * as CardTestModule from './components/CardTest.vue'
import * as ProgressTestModule from './components/ProgressTest.vue'

// 获取默认组件
const IconTest = (IconTestModule as any).default || IconTestModule
const ButtonTest = (ButtonTestModule as any).default || ButtonTestModule
const SelectTest = (SelectTestModule as any).default || SelectTestModule
const DialogTest = (DialogTestModule as any).default || DialogTestModule
const TableTest = (TableTestModule as any).default || TableTestModule
const PaginationTest = (PaginationTestModule as any).default || PaginationTestModule
const UploadTest = (UploadTestModule as any).default || UploadTestModule
const MessageTest = (MessageTestModule as any).default || MessageTestModule
const StorageTest = (StorageTestModule as any).default || StorageTestModule
const FormTest = (FormTestModule as any).default || FormTestModule
const InputTest = (InputTestModule as any).default || InputTestModule
const TextareaTest = (TextareaTestModule as any).default || TextareaTestModule
const SwitchTest = (SwitchTestModule as any).default || SwitchTestModule
const CheckboxTest = (CheckboxTestModule as any).default || CheckboxTestModule
const TagTest = (TagTestModule as any).default || TagTestModule
const InputTagTest = (InputTagTestModule as any).default || InputTagTestModule
const TabsTest = (TabsTestModule as any).default || TabsTestModule
const CardTest = (CardTestModule as any).default || CardTestModule
const ProgressTest = (ProgressTestModule as any).default || ProgressTestModule

defineOptions({
  name: 'DebugPage'
})

// Tab配置
const tabs = [
  {
    key: 'icon',
    name: 'Icon 图标',
    shortDesc: '基础图标组件',
    icon: 'star',
    title: 'Icon 图标组件',
    description: '基于 Heroicons 的图标组件，支持 outline 和 solid 两种风格',
    color: 'text-blue-500',
    component: IconTest
  },
  {
    key: 'button',
    name: 'Button 按钮',
    shortDesc: '多功能按钮',
    icon: 'check',
    title: 'Button 按钮组件',
    description: '支持多种类型、尺寸和状态的按钮组件',
    color: 'text-green-500',
    component: ButtonTest
  },
  {
    key: 'select',
    name: 'Select 选择器',
    shortDesc: '下拉选择组件',
    icon: 'chevron-down',
    title: 'Select 选择器组件',
    description: '功能丰富的下拉选择组件，支持单选、多选、搜索等功能',
    color: 'text-violet-500',
    component: SelectTest
  },
  {
    key: 'dialog',
    name: 'Dialog 对话框',
    shortDesc: '模态对话框',
    icon: 'rectangle-stack',
    title: 'Dialog 对话框组件',
    description: '功能强大的模态对话框组件，支持多种尺寸和交互方式',
    color: 'text-emerald-500',
    component: DialogTest
  },
  {
    key: 'table',
    name: 'Table 表格',
    shortDesc: '数据展示表格',
    icon: 'table-cells',
    title: 'Table 表格组件',
    description: '功能强大的数据表格组件，支持排序、筛选、分页等功能',
    color: 'text-indigo-500',
    component: TableTest
  },
  {
    key: 'pagination',
    name: 'Pagination 分页',
    shortDesc: '分页导航组件',
    icon: 'squares-2x2',
    title: 'Pagination 分页组件',
    description: '功能完整的分页组件，支持多种配置和自定义样式',
    color: 'text-pink-500',
    component: PaginationTest
  },
  {
    key: 'upload',
    name: 'Upload 上传',
    shortDesc: '文件上传组件',
    icon: 'cloud-arrow-up',
    title: 'Upload 文件上传组件',
    description: '支持简单上传和大文件分片上传，带有进度显示和MD5去重',
    color: 'text-purple-500',
    component: UploadTest
  },
  {
    key: 'message',
    name: 'Message 消息',
    shortDesc: '消息提示组件',
    icon: 'chat-bubble-left',
    title: 'Message 消息组件',
    description: '全局消息提示组件，支持多种类型和自动关闭',
    color: 'text-orange-500',
    component: MessageTest
  },
  {
    key: 'storage',
    name: 'Storage 存储',
    shortDesc: '安全存储方案',
    icon: 'archive-box',
    title: 'Storage 安全存储组件',
    description: '支持加密和有效期管理的安全存储解决方案',
    color: 'text-cyan-500',
    component: StorageTest
  },
  {
    key: 'form',
    name: 'Form 表单',
    shortDesc: '表单组件',
    icon: 'document-text',
    title: 'Form 表单组件',
    description: '功能强大的表单组件，支持多种类型和验证',
    color: 'text-teal-500',
    component: FormTest
  },
  {
    key: 'input',
    name: 'Input 输入框',
    shortDesc: '输入框组件',
    icon: 'pencil',
    title: 'Input 输入框组件',
    description: '功能强大的输入框组件，支持多种类型、状态和交互',
    color: 'text-blue-600',
    component: InputTest
  },
  {
    key: 'textarea',
    name: 'Textarea 文本域',
    shortDesc: '文本域组件',
    icon: 'document-text',
    title: 'Textarea 文本域组件',
    description: '功能丰富的多行文本输入组件，支持自动调整和字符计数',
    color: 'text-emerald-600',
    component: TextareaTest
  },
  {
    key: 'switch',
    name: 'Switch 开关',
    shortDesc: '开关组件',
    icon: 'toggle-right',
    title: 'Switch 开关组件',
    description: '功能丰富的开关组件，支持多种状态和交互',
    color: 'text-indigo-600',
    component: SwitchTest
  },
  {
    key: 'checkbox',
    name: 'Checkbox 复选框',
    shortDesc: '复选框组件',
    icon: 'check-square',
    title: 'Checkbox 复选框组件',
    description: '功能丰富的复选框组件，支持多种状态和交互',
    color: 'text-pink-600',
    component: CheckboxTest
  },
  {
    key: 'tag',
    name: 'Tag 标签',
    shortDesc: '标签组件',
    icon: 'tag',
    title: 'Tag 标签组件',
    description: '功能丰富的标签组件，支持多种状态和交互',
    color: 'text-purple-600',
    component: TagTest
  },
  {
    key: 'input-tag',
    name: 'InputTag 输入标签',
    shortDesc: '输入标签组件',
    icon: 'hashtag',
    title: 'InputTag 输入标签组件',
    description: '功能丰富的输入标签组件，支持多种状态和交互',
    color: 'text-orange-600',
    component: InputTagTest
  },
  {
    key: 'tabs',
    name: 'Tabs 标签页',
    shortDesc: '标签页导航组件',
    icon: 'folder',
    title: 'Tabs 标签页组件',
    description: '功能强大的标签页组件，支持多种类型和交互方式',
    color: 'text-slate-600',
    component: TabsTest
  },
  {
    key: 'card',
    name: 'Card 卡片',
    shortDesc: '内容容器组件',
    icon: 'rectangle-group',
    title: 'Card 卡片组件',
    description: '灵活的内容容器组件，支持多种布局和交互效果',
    color: 'text-zinc-600',
    component: CardTest
  },
  {
    key: 'progress',
    name: 'Progress 进度条',
    shortDesc: '进度显示组件',
    icon: 'chart-bar',
    title: 'Progress 进度条组件',
    description: '多样式进度条组件，支持线性、圆形、仪表盘等形式',
    color: 'text-sky-600',
    component: ProgressTest
  }
]

// 当前激活的Tab
const activeTab = ref('icon')

// 当前Tab的配置
const currentTab = computed(() => {
  return tabs.find(tab => tab.key === activeTab.value) || tabs[0]
})

// 当前显示的组件
const currentComponent = computed(() => {
  return currentTab.value.component
})

// 获取颜色类名
const getColorClass = (colorClass: string) => {
  const colorMap: Record<string, string> = {
    'text-blue-500': 'blue',
    'text-green-500': 'green',
    'text-violet-500': 'violet',
    'text-indigo-500': 'indigo',
    'text-pink-500': 'pink',
    'text-purple-500': 'purple',
    'text-orange-500': 'orange',
    'text-cyan-500': 'cyan',
    'text-teal-500': 'teal',
    'text-emerald-500': 'emerald',
    'text-blue-600': 'blue',
    'text-emerald-600': 'emerald',
    'text-indigo-600': 'indigo',
    'text-pink-600': 'pink',
    'text-purple-600': 'purple',
    'text-orange-600': 'orange',
    'text-slate-600': 'slate',
    'text-zinc-600': 'zinc',
    'text-sky-600': 'sky'
  }
  return colorMap[colorClass] || 'gray'
}

// 根据颜色获取对应的样式类
const getColorClasses = (color: string) => {
  const colorMap: Record<string, string> = {
    'text-blue-500': 'bg-blue-100 text-blue-800 border-blue-200',
    'text-green-500': 'bg-green-100 text-green-800 border-green-200',
    'text-violet-500': 'bg-violet-100 text-violet-800 border-violet-200',
    'text-emerald-500': 'bg-emerald-100 text-emerald-800 border-emerald-200',
    'text-indigo-500': 'bg-indigo-100 text-indigo-800 border-indigo-200',
    'text-pink-500': 'bg-pink-100 text-pink-800 border-pink-200',
    'text-purple-500': 'bg-purple-100 text-purple-800 border-purple-200',
    'text-orange-500': 'bg-orange-100 text-orange-800 border-orange-200',
    'text-cyan-500': 'bg-cyan-100 text-cyan-800 border-cyan-200',
    'text-teal-500': 'bg-teal-100 text-teal-800 border-teal-200',
    'text-blue-600': 'bg-blue-100 text-blue-800 border-blue-200',
    'text-emerald-600': 'bg-emerald-100 text-emerald-800 border-emerald-200',
    'text-indigo-600': 'bg-indigo-100 text-indigo-800 border-indigo-200',
    'text-pink-600': 'bg-pink-100 text-pink-800 border-pink-200',
    'text-purple-600': 'bg-purple-100 text-purple-800 border-purple-200',
    'text-orange-600': 'bg-orange-100 text-orange-800 border-orange-200'
  }
  return colorMap[color] || 'bg-gray-100 text-gray-800 border-gray-200'
}
</script>

<style scoped>
/* 自定义滚动条样式 - 只应用于右侧内容区域 */
.overflow-y-auto::-webkit-scrollbar {
  width: 6px;
}

.overflow-y-auto::-webkit-scrollbar-track {
  background: #f1f5f9;
}

.overflow-y-auto::-webkit-scrollbar-thumb {
  background: #cbd5e1;
  border-radius: 3px;
}

.overflow-y-auto::-webkit-scrollbar-thumb:hover {
  background: #94a3b8;
}

/* 导航按钮动画 */
nav button {
  transition: all 0.2s ease;
}

nav button:hover {
  transform: translateX(2px);
}
</style>
