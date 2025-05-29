<template>
  <div class="p-6">
    <div class="mb-6">
      <h1 class="text-2xl font-bold text-gray-900">个人信息</h1>
      <p class="mt-2 text-gray-600">查看和管理您的个人资料</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- 用户信息卡片 -->
      <div class="bg-white rounded-lg shadow p-6">
        <h2 class="text-lg font-semibold mb-4">基本信息</h2>

        <div v-if="loading" class="flex justify-center py-8">
          <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
          <p class="ml-2">加载中...</p>
        </div>

        <div v-else-if="userInfo" class="space-y-4">
          <!-- 头像和基本信息 -->
          <div class="flex items-center space-x-4 mb-6">
            <div class="w-16 h-16 bg-gray-200 rounded-full flex items-center justify-center">
              <img
                v-if="userInfo.avatar"
                :src="userInfo.avatar"
                :alt="userInfo.username"
                class="w-16 h-16 rounded-full object-cover"
              />
              <span v-else class="text-xl font-semibold text-gray-600">
                {{ userInfo.username?.charAt(0).toUpperCase() }}
              </span>
            </div>
            <div>
              <h3 class="text-lg font-medium">{{ userInfo.username }}</h3>
              <p class="text-gray-600">{{ userInfo.email }}</p>
              <p class="text-sm text-gray-500">
                {{ getUserRoleText(userInfo.role) }} • {{ getUserStatusText(userInfo.status) }}
              </p>
            </div>
          </div>

          <!-- 可编辑字段 -->
          <div class="space-y-4">
            <!-- 用户名 -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">用户名</label>
              <div class="relative">
                <input
                  v-if="editingFields.username"
                  ref="usernameInput"
                  v-model="editData.username"
                  type="text"
                  class="w-full px-3 py-2 border border-blue-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  @blur="saveUsername"
                  @keyup.enter="saveUsername"
                  @keyup.escape="cancelEditUsername"
                />
                <div
                  v-else
                  class="flex items-center justify-between p-3 bg-gray-50 rounded-md cursor-pointer hover:bg-gray-100 transition-colors"
                  @click="startEditUsername"
                >
                  <span>{{ userInfo.username }}</span>
                  <GlobalIcon name="pencil" size="sm" class="text-gray-400" />
                </div>
              </div>
            </div>

            <!-- 邮箱 -->
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">邮箱</label>
              <div class="space-y-2">
                <div class="flex items-center justify-between p-3 bg-gray-50 rounded-md">
                  <span>{{ userInfo.email }}</span>
                  <button
                    @click="showEmailModal = true"
                    class="text-blue-600 hover:text-blue-700 text-sm"
                  >
                    修改邮箱
                  </button>
                </div>
              </div>
            </div>

            <!-- 其他信息 -->
            <div>
              <label class="block text-sm font-medium text-gray-700">注册时间</label>
              <div class="mt-1 p-3 bg-gray-50 rounded-md">
                {{ formatDate(userInfo.created_at) }}
              </div>
            </div>
          </div>
        </div>

        <div v-else class="text-center py-8 text-gray-500">无法加载用户信息</div>
      </div>

      <!-- 修改密码卡片 -->
      <div class="bg-white rounded-lg shadow p-6">
        <h2 class="text-lg font-semibold mb-4">修改密码</h2>

        <form @submit.prevent="handleChangePassword" class="space-y-4">
          <div>
            <label for="oldPassword" class="block text-sm font-medium text-gray-700">
              当前密码
            </label>
            <input
              id="oldPassword"
              v-model="passwordForm.oldPassword"
              type="password"
              required
              class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="请输入当前密码"
            />
          </div>

          <div>
            <label for="newPassword" class="block text-sm font-medium text-gray-700">
              新密码
            </label>
            <input
              id="newPassword"
              v-model="passwordForm.newPassword"
              type="password"
              required
              minlength="6"
              class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="请输入新密码（至少6位）"
            />
          </div>

          <div>
            <label for="confirmPassword" class="block text-sm font-medium text-gray-700">
              确认新密码
            </label>
            <input
              id="confirmPassword"
              v-model="passwordForm.confirmPassword"
              type="password"
              required
              class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="请再次输入新密码"
            />
          </div>

          <div v-if="validationError" class="text-red-600 text-sm">
            {{ validationError }}
          </div>

          <button
            type="submit"
            :disabled="passwordLoading"
            class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            <span v-if="passwordLoading" class="flex items-center">
              <GlobalIcon name="arrow-path" class="animate-spin -ml-1 mr-3 h-5 w-5" />
              修改中...
            </span>
            <span v-else>修改密码</span>
          </button>
        </form>
      </div>
    </div>

    <!-- 修改邮箱模态框 -->
    <div
      v-if="showEmailModal"
      class="fixed inset-0 z-50 overflow-y-auto"
      @click.self="closeEmailModal"
    >
      <div class="flex items-center justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity"></div>

        <div class="inline-block align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full">
          <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
            <div class="sm:flex sm:items-start">
              <div class="w-full">
                <h3 class="text-lg leading-6 font-medium text-gray-900 mb-4">
                  修改邮箱
                </h3>

                <form @submit.prevent="handleChangeEmail" class="space-y-4">
                  <div>
                    <label for="newEmail" class="block text-sm font-medium text-gray-700">
                      新邮箱
                    </label>
                    <input
                      id="newEmail"
                      v-model="emailForm.newEmail"
                      type="email"
                      required
                      class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                      placeholder="请输入新邮箱地址"
                    />
                  </div>

                  <div>
                    <label for="emailCode" class="block text-sm font-medium text-gray-700">
                      验证码
                    </label>
                    <div class="mt-1 flex rounded-md shadow-sm">
                      <input
                        id="emailCode"
                        v-model="emailForm.code"
                        type="text"
                        required
                        maxlength="6"
                        class="flex-1 block w-full px-3 py-2 border border-gray-300 rounded-l-md focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                        placeholder="请输入6位验证码"
                      />
                      <button
                        type="button"
                        :disabled="!canSendEmailCode || emailCodeLoading"
                        @click="sendEmailCode"
                        class="inline-flex items-center px-3 py-2 border border-l-0 border-gray-300 rounded-r-md bg-gray-50 text-gray-500 text-sm disabled:opacity-50 disabled:cursor-not-allowed hover:bg-gray-100"
                      >
                        <span v-if="emailCodeLoading">发送中...</span>
                        <span v-else-if="emailCodeCountdown > 0">{{ emailCodeCountdown }}s</span>
                        <span v-else>发送验证码</span>
                      </button>
                    </div>
                  </div>

                  <div v-if="emailError" class="text-red-600 text-sm">
                    {{ emailError }}
                  </div>
                </form>
              </div>
            </div>
          </div>

          <div class="bg-gray-50 px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse">
            <button
              @click="handleChangeEmail"
              :disabled="emailLoading"
              class="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-blue-600 text-base font-medium text-white hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 sm:ml-3 sm:w-auto sm:text-sm disabled:opacity-50 disabled:cursor-not-allowed"
            >
              <span v-if="emailLoading" class="flex items-center">
                <GlobalIcon name="arrow-path" class="animate-spin -ml-1 mr-2 h-4 w-4" />
                修改中...
              </span>
              <span v-else>确认修改</span>
            </button>
            <button
              @click="closeEmailModal"
              type="button"
              class="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm"
            >
              取消
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, nextTick, onMounted, defineOptions } from 'vue'
import { useAuth } from '../../hooks/user/useAuth'
import { authApi } from '../../api/auth'
import { useMessage } from '../../composables/useMessage'
import type { ChangePasswordRequest, User } from '../../types'

const { user, getUserInfo } = useAuth()
const { success, error } = useMessage()

// 用户信息
const userInfo = ref<User | null>(null)
const loading = ref(false)

// 密码表单
const passwordForm = ref<ChangePasswordRequest & { confirmPassword: string }>({
  oldPassword: '',
  newPassword: '',
  confirmPassword: '',
})

const passwordLoading = ref(false)
const validationError = ref('')

// 邮箱表单
const emailForm = ref({
  newEmail: '',
  code: '',
})

const emailLoading = ref(false)
const emailError = ref('')
const emailCodeLoading = ref(false)
const emailCodeCountdown = ref(0)

// 编辑相关
const editingFields = ref({
  username: false,
})

const editData = ref({
  username: '',
})

const usernameInput = ref<HTMLInputElement | null>(null)

// 计算属性
const canSendEmailCode = computed(() => {
  return emailForm.value.newEmail && emailForm.value.newEmail.includes('@') && emailCodeCountdown.value === 0
})

// 格式化日期
const formatDate = (dateString?: string) => {
  if (!dateString) return ''
  return new Date(dateString).toLocaleString('zh-CN')
}

// 获取用户角色文本
const getUserRoleText = (role?: number) => {
  switch (role) {
    case 1:
      return '超级管理员'
    case 2:
      return '管理员'
    case 3:
      return '普通用户'
    default:
      return '未知角色'
  }
}

// 获取用户状态文本
const getUserStatusText = (status?: number) => {
  switch (status) {
    case 1:
      return '正常'
    case 2:
      return '禁用'
    case 3:
      return '已删除'
    default:
      return '未知状态'
  }
}

// 获取用户信息
const fetchUserInfo = async () => {
  try {
    loading.value = true
    
    // 直接从API获取最新的用户信息
    const response = await authApi.getUserInfo()
    
    // 从响应中提取 data 字段
    userInfo.value = (response as any).data || response
  } catch (_error: unknown) {
    console.error('获取用户信息失败:', _error)
    error('获取用户信息失败')
  } finally {
    loading.value = false
  }
}

// 修改密码
const handleChangePassword = async () => {
  validationError.value = ''

  // 前端验证
  if (passwordForm.value.newPassword !== passwordForm.value.confirmPassword) {
    validationError.value = '两次输入的新密码不一致'
    return
  }

  if (passwordForm.value.newPassword.length < 6) {
    validationError.value = '新密码长度不能少于6位'
    return
  }

  try {
    passwordLoading.value = true

    await authApi.changePassword({
      oldPassword: passwordForm.value.oldPassword,
      newPassword: passwordForm.value.newPassword,
    })

    // 显示成功消息
    success('密码修改成功')

    // 清空表单
    passwordForm.value = {
      oldPassword: '',
      newPassword: '',
      confirmPassword: '',
    }
  } catch (_error: unknown) {
    // 错误已经由全局拦截器处理，这里不需要额外处理
    console.error('修改密码失败:', _error)
  } finally {
    passwordLoading.value = false
  }
}

// 修改邮箱相关
const showEmailModal = ref(false)

const closeEmailModal = () => {
  showEmailModal.value = false
  emailForm.value = { newEmail: '', code: '' }
  emailError.value = ''
}

// 发送邮箱验证码
const sendEmailCode = async () => {
  if (!canSendEmailCode.value) return

  try {
    emailCodeLoading.value = true
    emailError.value = ''

    // 这里需要调用发送修改邮箱验证码的 API
    await authApi.sendChangeEmailCode({ email: emailForm.value.newEmail })

    success('验证码已发送')

    // 开始倒计时
    startEmailCodeCountdown()
  } catch (err: unknown) {
    const errorMessage = err instanceof Error ? err.message : '发送验证码失败'
    emailError.value = errorMessage
    error(errorMessage)
  } finally {
    emailCodeLoading.value = false
  }
}

// 验证码倒计时
const startEmailCodeCountdown = () => {
  emailCodeCountdown.value = 60
  const timer = setInterval(() => {
    emailCodeCountdown.value--
    if (emailCodeCountdown.value <= 0) {
      clearInterval(timer)
    }
  }, 1000)
}

const handleChangeEmail = async () => {
  emailError.value = ''

  if (!emailForm.value.newEmail || !emailForm.value.code) {
    emailError.value = '请填写完整信息'
    return
  }

  try {
    emailLoading.value = true

    // 调用修改邮箱的接口（需要验证码）
    await authApi.updateProfile({
      email: emailForm.value.newEmail,
      code: emailForm.value.code
    })

    success('邮箱修改成功')

    // 刷新用户信息
    await fetchUserInfo()

    // 关闭模态框
    closeEmailModal()
  } catch (err: unknown) {
    const errorMessage = err instanceof Error ? err.message : '修改邮箱失败'
    emailError.value = errorMessage
  } finally {
    emailLoading.value = false
  }
}

// 编辑用户名相关
const startEditUsername = async () => {
  editingFields.value.username = true
  editData.value.username = userInfo.value?.username || ''

  // 等待DOM更新后聚焦输入框
  await nextTick()
  usernameInput.value?.focus()
}

const saveUsername = async () => {
  if (!editData.value.username.trim()) {
    error('用户名不能为空')
    cancelEditUsername()
    return
  }

  if (editData.value.username === userInfo.value?.username) {
    // 没有变化，直接取消编辑
    cancelEditUsername()
    return
  }

  try {
    // 调用更新用户资料接口
    await authApi.updateProfile({ username: editData.value.username })

    // 更新本地用户信息
    if (userInfo.value) {
      userInfo.value.username = editData.value.username
    }

    success('用户名修改成功')
  } catch (_error: unknown) {
    console.error('修改用户名失败:', _error)
    // 恢复原值
    editData.value.username = userInfo.value?.username || ''
  } finally {
    editingFields.value.username = false
  }
}

const cancelEditUsername = () => {
  editingFields.value.username = false
  editData.value.username = userInfo.value?.username || ''
}

onMounted(() => {
  fetchUserInfo()
})

defineOptions({
  name: 'ProfilePage',
})
</script>
