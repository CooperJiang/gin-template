<template>
  <Teleport to="body">
    <div class="fixed top-6 left-1/2 transform -translate-x-1/2 z-[9999] pointer-events-none">
      <TransitionGroup name="message" tag="div" class="flex flex-col items-center space-y-3">
        <div v-for="message in messages" :key="message.id" class="message-item pointer-events-auto">
          <div
            :class="[
              'flex items-center gap-3 px-5 py-4 rounded-xl shadow-xl backdrop-blur-sm',
              'transition-all duration-300 ease-out min-w-80 max-w-md',
              'border border-opacity-30',
              messageClasses[message.type],
            ]"
          >
            <!-- 图标 -->
            <div class="flex-shrink-0">
              <component :is="messageIcons[message.type]" class="w-6 h-6" />
            </div>

            <!-- 消息内容 -->
            <div class="flex-1 text-sm font-medium leading-relaxed">
              {{ message.content }}
            </div>

            <!-- 关闭按钮 -->
            <button
              v-if="message.showClose"
              @click="removeMessage(message.id)"
              class="flex-shrink-0 p-1 opacity-60 hover:opacity-100 transition-opacity rounded-full hover:bg-black hover:bg-opacity-10"
            >
              <XMarkIcon class="w-4 h-4" />
            </button>
          </div>
        </div>
      </TransitionGroup>
    </div>
  </Teleport>
</template>

<script setup lang="ts">
import { messages, removeMessage } from '../../composables/useMessage'
import {
  CheckCircleIcon,
  ExclamationCircleIcon,
  ExclamationTriangleIcon,
  InformationCircleIcon,
  XMarkIcon,
} from '@heroicons/vue/24/outline'
// import type { MessageContainerProps } from './types'

// Props (currently not used but structure ready for future use)
// const props = withDefaults(defineProps<MessageContainerProps>(), {
//   maxCount: 5
// })

// 消息类型对应的样式 - 优化颜色和对比度
const messageClasses = {
  success: 'bg-emerald-50 text-emerald-800 border-emerald-200 shadow-emerald-100',
  error: 'bg-red-50 text-red-800 border-red-200 shadow-red-100',
  warning: 'bg-amber-50 text-amber-800 border-amber-200 shadow-amber-100',
  info: 'bg-sky-50 text-sky-800 border-sky-200 shadow-sky-100',
}

// 消息类型对应的图标
const messageIcons = {
  success: CheckCircleIcon,
  error: ExclamationCircleIcon,
  warning: ExclamationTriangleIcon,
  info: InformationCircleIcon,
}

defineOptions({
  name: 'MessageContainer',
})
</script>

<style scoped>
/* 消息进入动画 - 从上方滑入 */
.message-enter-active {
  transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275);
}

.message-leave-active {
  transition: all 0.3s ease-in;
}

.message-enter-from {
  opacity: 0;
  transform: translateY(-50px) scale(0.9);
}

.message-leave-to {
  opacity: 0;
  transform: translateY(-30px) scale(0.95);
}

.message-move {
  transition: transform 0.4s ease;
}

/* 消息项悬停效果 */
.message-item:hover {
  transform: scale(1.02);
}

/* 自定义阴影效果 */
.message-item > div {
  box-shadow:
    0 10px 25px -3px rgba(0, 0, 0, 0.1),
    0 4px 6px -2px rgba(0, 0, 0, 0.05),
    0 0 0 1px rgba(255, 255, 255, 0.1);
}
</style>
