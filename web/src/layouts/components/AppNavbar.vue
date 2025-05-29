<template>
  <header class="bg-white shadow-sm border-b border-gray-200">
    <div class="flex items-center justify-between h-16 px-4 sm:px-6 lg:px-8">
      <!-- 移动端菜单按钮 -->
      <button
        @click="$emit('toggleSidebar')"
        class="lg:hidden p-2 rounded-md text-gray-400 hover:text-gray-500 hover:bg-gray-100"
      >
        <Bars3Icon class="h-6 w-6" />
      </button>

      <!-- 面包屑导航 -->
      <div class="flex-1 lg:flex-none">
        <GlobalAppBreadcrumb />
      </div>

      <!-- 右侧工具栏 -->
      <div class="flex items-center space-x-4">
        <!-- 通知 -->
        <button class="p-2 text-gray-400 hover:text-gray-500 hover:bg-gray-100 rounded-md">
          <BellIcon class="h-5 w-5" />
        </button>

        <!-- 用户菜单 -->
        <div class="relative">
          <button
            @click="toggleUserMenu"
            class="flex items-center text-sm rounded-full focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            <div class="w-8 h-8 bg-gray-300 rounded-full flex items-center justify-center">
              <UserIcon class="h-5 w-5 text-gray-600" />
            </div>
          </button>

          <!-- 用户下拉菜单 -->
          <div
            v-if="userMenuOpen"
            class="absolute right-0 mt-2 w-48 bg-white rounded-md shadow-lg py-1 z-50 border border-gray-200"
          >
            <router-link
              to="/profile"
              class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
              @click="handleProfileClick"
            >
              个人资料
            </router-link>
            <button
              @click="handleLogout"
              class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
            >
              退出登录
            </button>
          </div>
        </div>
      </div>
    </div>
  </header>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuth } from '../../hooks/user/useAuth'
import { Bars3Icon, BellIcon, UserIcon } from '@heroicons/vue/24/outline'

defineEmits<{
  toggleSidebar: []
}>()

const router = useRouter()
const { logout } = useAuth()

const userMenuOpen = ref(false)

const toggleUserMenu = () => {
  userMenuOpen.value = !userMenuOpen.value
}

const handleLogout = async () => {
  userMenuOpen.value = false
  logout()
}

const handleProfileClick = () => {
  userMenuOpen.value = false
  router.push('/profile')
}

// 点击外部关闭用户菜单
const handleClickOutside = (event: Event) => {
  const target = event.target as Element
  if (!target.closest('.relative')) {
    userMenuOpen.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>
