<template>
  <div v-if="loading" class="loading-overlay">
    <div class="loading-container">
      <div class="loading-spinner"></div>
      <p v-if="text" class="loading-text">{{ text }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
interface Props {
  loading?: boolean
  text?: string
}

withDefaults(defineProps<Props>(), {
  loading: true,
  text: '加载中...'
})
</script>

<style scoped>
.loading-overlay {
  @apply fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50;
}

.loading-container {
  @apply bg-white rounded-lg p-6 flex flex-col items-center space-y-4;
}

.loading-spinner {
  @apply w-8 h-8 border-4 border-primary-200 border-t-primary-600 rounded-full animate-spin;
}

.loading-text {
  @apply text-gray-600 text-sm;
}
</style>