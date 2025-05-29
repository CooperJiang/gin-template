<template>
  <teleport to="body">
    <div
      v-if="visible"
      ref="popoverRef"
      class="fixed z-[9999] bg-white shadow-lg border border-gray-200 rounded-md py-1"
      :style="popoverStyle"
      style="min-width: 200px; opacity: 0; transform: scale(0.95)"
      :class="{ 'animate-popover-in': visible }"
    >
      <slot />
    </div>
  </teleport>
</template>

<script lang="ts">
import { defineComponent, ref, computed, nextTick, watch } from 'vue'

export default defineComponent({
  name: 'PopoverMenu',
  props: {
    visible: {
      type: Boolean,
      required: true,
    },
    triggerRef: {
      type: Object as () => HTMLElement | null,
      default: null,
    },
    placement: {
      type: String as () => 'right' | 'bottom' | 'left' | 'top',
      default: 'right',
    },
    offset: {
      type: Number,
      default: 8,
    },
  },
  setup(props) {
    const popoverRef = ref<HTMLElement>()

    // 计算popover位置
    const popoverStyle = computed(() => {
      if (!props.triggerRef) return {}

      const triggerRect = props.triggerRef.getBoundingClientRect()
      const { placement, offset } = props

      let left = 0
      let top = 0

      switch (placement) {
        case 'right':
          left = triggerRect.right + offset
          top = triggerRect.top
          break
        case 'bottom':
          left = triggerRect.left
          top = triggerRect.bottom + offset
          break
        case 'left':
          left = triggerRect.left - offset
          top = triggerRect.top
          break
        case 'top':
          left = triggerRect.left
          top = triggerRect.top - offset
          break
      }

      // 边界检测，确保popover不会超出视窗
      const popoverEl = popoverRef.value
      if (popoverEl) {
        const popoverRect = popoverEl.getBoundingClientRect()
        const viewportWidth = window.innerWidth
        const viewportHeight = window.innerHeight

        // 水平边界检测
        if (left + popoverRect.width > viewportWidth) {
          left = viewportWidth - popoverRect.width - 10
        }
        if (left < 10) {
          left = 10
        }

        // 垂直边界检测
        if (top + popoverRect.height > viewportHeight) {
          top = viewportHeight - popoverRect.height - 10
        }
        if (top < 10) {
          top = 10
        }
      }

      return {
        left: `${left}px`,
        top: `${top}px`,
        transform: placement === 'left' ? 'translateX(-100%)' : undefined,
      }
    })

    // 监听visible变化，重新计算位置
    watch(
      () => props.visible,
      (newVisible) => {
        if (newVisible) {
          nextTick(() => {
            // 触发重新计算位置
            if (popoverRef.value) {
              popoverRef.value.style.visibility = 'hidden'
              popoverRef.value.style.visibility = 'visible'
            }
          })
        }
      },
    )

    return {
      popoverRef,
      popoverStyle,
    }
  },
})
</script>

<style scoped>
.animate-popover-in {
  animation: popoverIn 150ms ease-out forwards;
}

@keyframes popoverIn {
  0% {
    opacity: 0;
    transform: scale(0.95);
  }
  100% {
    opacity: 1;
    transform: scale(1);
  }
}
</style>
