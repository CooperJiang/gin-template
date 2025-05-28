<template>
  <div class="flex items-center">
    <div class="flex-shrink-0">
      <div class="w-8 h-8 bg-blue-600 rounded-lg flex items-center justify-center">
        <span class="text-white font-bold text-sm">GM</span>
      </div>
    </div>
    <span v-if="showText" class="ml-3 text-lg font-semibold text-gray-900">{{ title }}</span>
  </div>
</template>

<script setup lang="ts">
interface Props {
  title?: string
  showText?: boolean
}

withDefaults(defineProps<Props>(), {
  title: '管理后台',
  showText: true
})
</script>
