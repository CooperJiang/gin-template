<template>
  <div
    class="flex items-center cursor-pointer hover:opacity-80 transition-opacity duration-150"
    @click="handleLogoClick"
  >
    <div class="flex-shrink-0">
      <div class="w-8 h-8 bg-blue-600 rounded-lg flex items-center justify-center">
        <span class="text-white font-bold text-sm">GM</span>
      </div>
    </div>
    <span v-if="showText" class="ml-3 text-lg font-semibold text-gray-900">{{ title }}</span>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { useRouter } from 'vue-router'

export default defineComponent({
  name: 'AppLogo',
  props: {
    title: {
      type: String,
      default: '管理后台',
    },
    showText: {
      type: Boolean,
      default: true,
    },
  },
  setup() {
    const router = useRouter()

    const handleLogoClick = () => {
      router.push('/dashboard')
    }

    return {
      handleLogoClick,
    }
  },
})
</script>
