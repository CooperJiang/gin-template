<template>
  <Teleport to="body">
    <Transition
      :name="transition"
      @enter="handleEnter"
      @after-enter="handleAfterEnter"
      @leave="handleLeave"
      @after-leave="handleAfterLeave"
    >
      <div
        v-if="modelValue"
        :class="[
          'fixed inset-0 z-50 flex items-center justify-center',
          wrapClass,
          placementClasses
        ]"
        :style="{ zIndex }"
      >
        <!-- 遮罩层 -->
        <div
          v-if="showMask"
          :class="[
            'absolute inset-0 bg-black bg-opacity-50 transition-opacity duration-300',
            maskClass
          ]"
          @click="handleMaskClick"
        />

        <!-- 对话框容器 -->
        <div
          :class="[
            'relative bg-white rounded-lg shadow-xl transform transition-all duration-300',
            'max-h-screen overflow-hidden flex flex-col',
            sizeClasses,
            customClass
          ]"
          :style="dialogStyle"
          @click.stop
        >
          <!-- 加载层 -->
          <div
            v-if="loading"
            class="absolute inset-0 bg-white bg-opacity-75 flex items-center justify-center z-10"
          >
            <div class="flex flex-col items-center space-y-2">
              <div class="animate-spin w-6 h-6">
                <GlobalIcon name="arrow-path" size="lg" color="text-blue-500" />
              </div>
              <span class="text-sm text-gray-500">加载中...</span>
            </div>
          </div>

          <!-- 头部 -->
          <div
            v-if="showHeader"
            class="flex items-center justify-between px-6 py-4 border-b border-gray-200 flex-shrink-0"
          >
            <div class="flex items-center">
              <!-- 自定义头部插槽 -->
              <slot name="header">
                <!-- 标题插槽 -->
                <slot name="title">
                  <h3 class="text-lg font-semibold text-gray-900">
                    {{ title }}
                  </h3>
                </slot>
              </slot>
            </div>

            <!-- 关闭按钮 -->
            <button
              v-if="showClose && closable && !persistent"
              @click="handleClose"
              class="p-1 hover:bg-gray-100 rounded-md transition-colors focus:outline-none focus:ring-2 focus:ring-blue-500"
            >
              <slot name="close">
                <GlobalIcon name="x-mark" size="lg" color="text-gray-400" />
              </slot>
            </button>
          </div>

          <!-- 内容区域 -->
          <div
            :class="[
              'flex-1 overflow-y-auto',
              bodyClass,
              { 'p-6': !$slots.default || !hasCustomPadding }
            ]"
          >
            <slot />
          </div>

          <!-- 底部 -->
          <div
            v-if="showFooter"
            class="flex items-center justify-end space-x-3 px-6 py-4 border-t border-gray-200 bg-gray-50 flex-shrink-0"
          >
            <slot name="footer">
              <!-- 取消按钮 -->
              <GlobalButton
                v-if="showCancel"
                variant="outline"
                @click="handleCancel"
              >
                {{ cancelText }}
              </GlobalButton>

              <!-- 确认按钮 -->
              <GlobalButton
                v-if="showConfirm"
                :variant="confirmType"
                :loading="confirmLoading"
                @click="handleConfirm"
              >
                {{ confirmText }}
              </GlobalButton>
            </slot>
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup lang="ts">
import { ref, computed, watch, onMounted, onUnmounted, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import type { DialogProps, DialogEmits } from './types'

defineOptions({
  name: 'GlobalDialog'
})

const props = withDefaults(defineProps<DialogProps>(), {
  modelValue: false,
  title: '',
  width: 'auto',
  height: 'auto',
  placement: 'center',
  size: 'medium',
  closable: true,
  maskClosable: true,
  escClosable: true,
  closeOnRouteChange: true,
  showMask: true,
  showHeader: true,
  showFooter: false,
  showClose: true,
  transition: 'dialog-fade',
  zIndex: 1000,
  confirmText: '确认',
  cancelText: '取消',
  confirmType: 'primary',
  showConfirm: true,
  showCancel: true,
  confirmLoading: false,
  persistent: false,
  loading: false
})

const emit = defineEmits<DialogEmits>()

// 路由监听
const router = useRouter()

// 计算属性
const placementClasses = computed(() => {
  const placements = {
    center: 'items-center justify-center',
    top: 'items-start justify-center pt-20',
    bottom: 'items-end justify-center pb-20'
  }
  return placements[props.placement]
})

const sizeClasses = computed(() => {
  const sizes = {
    small: 'w-full max-w-md',
    medium: 'w-full max-w-lg',
    large: 'w-full max-w-2xl',
    full: 'w-full h-full max-w-none m-0 rounded-none'
  }
  return sizes[props.size]
})

const dialogStyle = computed(() => {
  const style: Record<string, string> = {}

  // 自定义宽度
  if (props.width && props.width !== 'auto') {
    if (typeof props.width === 'number') {
      style.width = `${props.width}px`
    } else {
      style.width = props.width
    }
  }

  // 自定义高度
  if (props.height && props.height !== 'auto') {
    if (typeof props.height === 'number') {
      style.height = `${props.height}px`
    } else {
      style.height = props.height
    }
  }

  return style
})

const hasCustomPadding = computed(() => {
  // 检查是否有自定义内容需要取消默认内边距
  return false // 可以根据实际需求实现
})

// 方法
const handleMaskClick = () => {
  if (props.maskClosable && !props.persistent) {
    emit('mask-click')
    handleClose()
  }
}

const handleClose = () => {
  if (!props.persistent) {
    emit('update:modelValue', false)
    emit('close')
  }
}

const handleConfirm = () => {
  emit('confirm')
}

const handleCancel = () => {
  emit('cancel')
  if (!props.persistent) {
    handleClose()
  }
}

const handleEscape = (event: KeyboardEvent) => {
  if (event.key === 'Escape' && props.escClosable && props.modelValue && !props.persistent) {
    emit('esc-press')
    handleClose()
  }
}

// 动画事件
const handleEnter = () => {
  emit('open')
}

const handleAfterEnter = () => {
  emit('opened')
}

const handleLeave = () => {
  // 开始关闭动画
}

const handleAfterLeave = () => {
  emit('closed')
}

// 暴露方法
const open = () => {
  emit('update:modelValue', true)
}

const close = () => {
  handleClose()
}

const toggle = () => {
  emit('update:modelValue', !props.modelValue)
}

defineExpose({
  open,
  close,
  toggle
})

// 生命周期
onMounted(() => {
  document.addEventListener('keydown', handleEscape)

  // 路由变化时关闭对话框
  if (props.closeOnRouteChange && router) {
    router.afterEach(() => {
      if (props.modelValue) {
        handleClose()
      }
    })
  }
})

onUnmounted(() => {
  document.removeEventListener('keydown', handleEscape)
})

// 监听器
watch(() => props.modelValue, (newValue) => {
  if (newValue) {
    // 阻止页面滚动
    document.body.style.overflow = 'hidden'
  } else {
    // 恢复页面滚动
    document.body.style.overflow = ''
  }
}, { immediate: true })

// 组件卸载时恢复滚动
onUnmounted(() => {
  document.body.style.overflow = ''
})
</script>

<style scoped>
/* 对话框动画 */
.dialog-fade-enter-active,
.dialog-fade-leave-active {
  transition: opacity 0.3s ease;
}

.dialog-fade-enter-active .relative,
.dialog-fade-leave-active .relative {
  transition: transform 0.3s ease;
}

.dialog-fade-enter-from,
.dialog-fade-leave-to {
  opacity: 0;
}

.dialog-fade-enter-from .relative {
  transform: scale(0.9) translateY(-20px);
}

.dialog-fade-leave-to .relative {
  transform: scale(0.9) translateY(-20px);
}

/* 自定义滚动条 */
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
</style>
