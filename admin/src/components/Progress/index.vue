<template>
  <div :class="progressClass">
    <!-- 线性进度条 -->
    <div v-if="type === 'line'" :class="lineProgressClass">
      <!-- 进度文字 (顶部) -->
      <div v-if="showText && textPosition === 'top'" :class="textClass">
        <slot :value="normalizedValue" :status="status">
          {{ progressText }}
        </slot>
      </div>

      <!-- 进度条容器 -->
      <div class="progress-line-container flex items-center">
        <!-- 进度条轨道 -->
        <div :class="lineTrackClass" :style="lineTrackStyle">
          <!-- 进度条填充 -->
          <div
            :class="lineFillClass"
            :style="lineFillStyle"
          >
            <!-- 内部文字 -->
            <div v-if="showText && textPosition === 'inside'" :class="insideTextClass">
              <slot name="inner" :value="normalizedValue" :status="status">
                {{ progressText }}
              </slot>
            </div>
          </div>

          <!-- 不确定进度动画 -->
          <div v-if="indeterminate" :class="indeterminateClass" />
        </div>

        <!-- 右侧文字 -->
        <div v-if="showText && textPosition === 'right'" :class="rightTextClass">
          <slot :value="normalizedValue" :status="status">
            {{ progressText }}
          </slot>
        </div>
      </div>

      <!-- 进度文字 (底部) -->
      <div v-if="showText && textPosition === 'bottom'" :class="textClass">
        <slot :value="normalizedValue" :status="status">
          {{ progressText }}
        </slot>
      </div>
    </div>

    <!-- 圆形进度条 -->
    <div v-else-if="type === 'circle' || type === 'dashboard'" :class="circleProgressClass">
      <svg :width="svgSize" :height="svgSize" viewBox="0 0 100 100">
        <!-- 背景圆环 -->
        <circle
          :cx="50"
          :cy="50"
          :r="radius"
          fill="none"
          :stroke="backgroundColor"
          :stroke-width="relativeStrokeWidth"
        />

        <!-- 进度圆环 -->
        <circle
          v-if="!indeterminate"
          :cx="50"
          :cy="50"
          :r="radius"
          fill="none"
          :stroke="currentColor"
          :stroke-width="relativeStrokeWidth"
          :stroke-dasharray="strokeDasharray"
          :stroke-dashoffset="strokeDashoffset"
          stroke-linecap="round"
          :style="circleStyle"
          :transform="circleTransform"
        />

        <!-- 不确定进度动画圆环 -->
        <circle
          v-if="indeterminate"
          :cx="50"
          :cy="50"
          :r="radius"
          fill="none"
          :stroke="currentColor"
          :stroke-width="relativeStrokeWidth"
          :stroke-dasharray="`${strokeDasharray * 0.25} ${strokeDasharray * 0.75}`"
          stroke-linecap="round"
          :style="indeterminateCircleStyle"
          :transform="circleTransform"
        />
      </svg>

      <!-- 圆形进度条中心文字 -->
      <div v-if="showText" :class="circleTextClass">
        <slot :value="normalizedValue" :status="status">
          {{ progressText }}
        </slot>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, watch } from 'vue'
import type { ProgressProps, ProgressEmits } from './types'

defineOptions({
  name: 'GlobalProgress'
})

const props = withDefaults(defineProps<ProgressProps>(), {
  value: 0,
  type: 'line',
  status: 'normal',
  size: 'medium',
  showText: true,
  textPosition: 'right',
  strokeWidth: undefined,
  width: 120,
  backgroundColor: '#f3f4f6',
  animated: true,
  striped: false,
  indeterminate: false,
  gapDegree: 75,
  gapPosition: 'bottom'
})

const emit = defineEmits<ProgressEmits>()

// 标准化进度值
const normalizedValue = computed(() => {
  return Math.max(0, Math.min(100, props.value))
})

// 进度文字
const progressText = computed(() => {
  if (props.format) {
    return props.format(normalizedValue.value)
  }
  return `${normalizedValue.value}%`
})

// 当前颜色
const currentColor = computed(() => {
  if (Array.isArray(props.color)) {
    const colorIndex = Math.floor((normalizedValue.value / 100) * (props.color.length - 1))
    return props.color[Math.min(colorIndex, props.color.length - 1)]
  }

  if (props.color) {
    return props.color
  }

  const statusColors = {
    normal: '#3b82f6',
    success: '#10b981',
    warning: '#f59e0b',
    error: '#ef4444',
    info: '#06b6d4'
  }
  return statusColors[props.status]
})

// 默认strokeWidth
const defaultStrokeWidth = computed(() => {
  const sizeMap = {
    small: props.type === 'line' ? 6 : 4,
    medium: props.type === 'line' ? 8 : 6,
    large: props.type === 'line' ? 10 : 8
  }
  return sizeMap[props.size]
})

// 实际strokeWidth
const actualStrokeWidth = computed(() => {
  return props.strokeWidth ?? defaultStrokeWidth.value
})

// SVG尺寸
const svgSize = computed(() => {
  return props.width
})

// 圆环半径
const radius = computed(() => {
  return 50 - actualStrokeWidth.value / 2
})

// 相对strokeWidth (相对于100x100的viewBox)
const relativeStrokeWidth = computed(() => {
  return (actualStrokeWidth.value / svgSize.value) * 100
})

// 圆环周长
const strokeDasharray = computed(() => {
  const circumference = 2 * Math.PI * radius.value
  if (props.type === 'dashboard') {
    const gapRadian = (props.gapDegree * Math.PI) / 180
    return circumference * (1 - gapRadian / (2 * Math.PI))
  }
  return circumference
})

// 圆环偏移
const strokeDashoffset = computed(() => {
  const offset = strokeDasharray.value * (1 - normalizedValue.value / 100)
  return offset
})

// 圆环变换
const circleTransform = computed(() => {
  if (props.type === 'dashboard') {
    const rotateMap = {
      top: 0,
      right: 90,
      bottom: 180,
      left: 270
    }
    const baseRotate = rotateMap[props.gapPosition]
    const gapRotate = props.gapDegree / 2
    return `rotate(${baseRotate + gapRotate} 50 50)`
  }
  return 'rotate(-90 50 50)'
})

// 样式计算
const progressClass = computed(() => {
  const classes = ['progress', `progress-${props.type}`, `progress-${props.size}`]

  if (props.status !== 'normal') {
    classes.push(`progress-${props.status}`)
  }

  return classes.join(' ')
})

const lineProgressClass = computed(() => {
  const classes = ['progress-line']

  if (props.textPosition === 'top' || props.textPosition === 'bottom') {
    classes.push('space-y-2')
  }

  return classes.join(' ')
})

const lineTrackClass = computed(() => {
  const classes = [
    'progress-line-track',
    'relative',
    'overflow-hidden',
    'rounded-full',
    'bg-gray-200'
  ]

  return classes.join(' ')
})

const lineTrackStyle = computed(() => {
  return {
    height: `${actualStrokeWidth.value}px`,
    backgroundColor: props.backgroundColor
  }
})

const lineFillClass = computed(() => {
  const classes = [
    'progress-line-fill',
    'h-full',
    'transition-all',
    'duration-300',
    'ease-out',
    'rounded-full',
    'relative',
    'overflow-hidden'
  ]

  if (props.striped) {
    classes.push('progress-striped')
  }

  if (props.animated && props.striped) {
    classes.push('progress-animated')
  }

  return classes.join(' ')
})

const lineFillStyle = computed(() => {
  return {
    width: props.indeterminate ? '100%' : `${normalizedValue.value}%`,
    backgroundColor: currentColor.value
  }
})

const indeterminateClass = computed(() => {
  return [
    'progress-indeterminate',
    'absolute',
    'inset-0',
    'rounded-full'
  ].join(' ')
})

const textClass = computed(() => {
  const sizeClasses = {
    small: 'text-xs',
    medium: 'text-sm',
    large: 'text-base'
  }
  return [
    'progress-text',
    'text-gray-600',
    'font-medium',
    sizeClasses[props.size]
  ].join(' ')
})

const rightTextClass = computed(() => {
  return [
    ...textClass.value.split(' '),
    'ml-3',
    'whitespace-nowrap'
  ].join(' ')
})

const insideTextClass = computed(() => {
  return [
    'progress-inside-text',
    'absolute',
    'inset-0',
    'flex',
    'items-center',
    'justify-center',
    'text-white',
    'text-xs',
    'font-medium'
  ].join(' ')
})

const circleProgressClass = computed(() => {
  return [
    'progress-circle',
    'relative',
    'inline-block'
  ].join(' ')
})

const circleTextClass = computed(() => {
  const sizeClasses = {
    small: 'text-xs',
    medium: 'text-sm',
    large: 'text-base'
  }
  return [
    'progress-circle-text',
    'absolute',
    'inset-0',
    'flex',
    'items-center',
    'justify-center',
    'text-gray-600',
    'font-medium',
    sizeClasses[props.size]
  ].join(' ')
})

const circleStyle = computed(() => {
  return {
    transition: props.animated ? 'stroke-dashoffset 0.3s ease-out' : 'none'
  }
})

const indeterminateCircleStyle = computed(() => {
  return {
    animation: 'progress-circle-spin 1.4s linear infinite'
  }
})

// 监听进度变化
watch(normalizedValue, (newValue, oldValue) => {
  emit('change', newValue)

  if (newValue === 100 && (oldValue ?? 0) < 100) {
    emit('complete')
  }
}, { immediate: true })
</script>

<style scoped>
/* 基础样式 */
.progress {
  @apply w-full;
}

.progress-line {
  @apply w-full;
}

.progress-line-track {
  @apply flex-1;
}

/* 条纹效果 */
.progress-striped {
  background-image: linear-gradient(
    45deg,
    rgba(255, 255, 255, 0.15) 25%,
    transparent 25%,
    transparent 50%,
    rgba(255, 255, 255, 0.15) 50%,
    rgba(255, 255, 255, 0.15) 75%,
    transparent 75%,
    transparent
  );
  background-size: 1rem 1rem;
}

/* 条纹动画 */
.progress-animated {
  animation: progress-stripes 1s linear infinite;
}

@keyframes progress-stripes {
  0% {
    background-position: 1rem 0;
  }
  100% {
    background-position: 0 0;
  }
}

/* 不确定进度动画 */
.progress-indeterminate {
  background: linear-gradient(
    90deg,
    transparent,
    currentColor,
    transparent
  );
  animation: progress-indeterminate 1.5s ease-in-out infinite;
}

@keyframes progress-indeterminate {
  0% {
    transform: translateX(-100%);
  }
  100% {
    transform: translateX(100%);
  }
}

/* 圆形进度条动画 */
@keyframes progress-circle-spin {
  0% {
    transform: rotate(-90deg);
  }
  100% {
    transform: rotate(270deg);
  }
}

/* 不同状态的颜色 */
.progress-success .progress-line-fill {
  @apply bg-green-500;
}

.progress-warning .progress-line-fill {
  @apply bg-yellow-500;
}

.progress-error .progress-line-fill {
  @apply bg-red-500;
}

.progress-info .progress-line-fill {
  @apply bg-blue-500;
}
</style>
