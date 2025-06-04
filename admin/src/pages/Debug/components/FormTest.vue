<template>
  <div class="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
    <h3 class="text-lg font-semibold mb-4">Form 表单组件测试</h3>
    <p class="text-gray-600 text-sm mb-6">基于 Element UI 设计的表单组件，支持完整的验证功能</p>

    <!-- 基础表单 -->
    <GlobalCodeDemo
      title="基础表单"
      description="最简单的表单使用方式，包含基本的输入框和文本域"
      :code="basicFormCode"
    >
      <div class="bg-gray-50 p-4 rounded-md">
        <GlobalForm
          ref="basicFormRef"
          :model="basicForm"
          :rules="basicRules"
          label-width="80px"
        >
          <GlobalFormItem label="用户名" prop="username">
            <GlobalInput v-model="basicForm.username" placeholder="请输入用户名" />
          </GlobalFormItem>

          <GlobalFormItem label="邮箱" prop="email">
            <GlobalInput v-model="basicForm.email" type="email" placeholder="请输入邮箱" />
          </GlobalFormItem>

          <GlobalFormItem label="描述" prop="description">
            <GlobalTextarea v-model="basicForm.description" placeholder="请输入描述" />
          </GlobalFormItem>

          <GlobalFormItem>
            <GlobalButton type="primary" @click="submitBasicForm">提交</GlobalButton>
            <GlobalButton @click="resetBasicForm" style="margin-left: 10px;">重置</GlobalButton>
          </GlobalFormItem>
        </GlobalForm>

        <div class="mt-4 p-3 bg-white rounded border">
          <h5 class="text-sm font-medium mb-2">表单数据：</h5>
          <pre class="text-xs text-gray-600">{{ JSON.stringify(basicForm, null, 2) }}</pre>
        </div>
      </div>
    </GlobalCodeDemo>

    <!-- 验证表单 -->
    <GlobalCodeDemo
      title="表单验证"
      description="包含各种验证规则的表单示例，支持必填、格式验证、自定义验证等"
      :code="validationFormCode"
    >
      <div class="bg-gray-50 p-4 rounded-md">
        <GlobalForm
          ref="validationFormRef"
          :model="validationForm"
          :rules="validationRules"
          label-width="120px"
          label-position="right"
        >
          <GlobalFormItem label="用户名" prop="username">
            <GlobalInput v-model="validationForm.username" placeholder="请输入用户名" />
          </GlobalFormItem>

          <GlobalFormItem label="密码" prop="password">
            <GlobalInput
              v-model="validationForm.password"
              type="password"
              placeholder="请输入密码"
              show-password
            />
          </GlobalFormItem>

          <GlobalFormItem label="确认密码" prop="confirmPassword">
            <GlobalInput
              v-model="validationForm.confirmPassword"
              type="password"
              placeholder="请再次输入密码"
              show-password
            />
          </GlobalFormItem>

          <GlobalFormItem label="年龄" prop="age">
            <GlobalInput v-model="validationForm.age" type="number" placeholder="请输入年龄" />
          </GlobalFormItem>

          <GlobalFormItem label="网站" prop="website">
            <GlobalInput v-model="validationForm.website" placeholder="请输入网站地址" />
          </GlobalFormItem>

          <GlobalFormItem label="性别" prop="gender">
            <GlobalSelect
              v-model="validationForm.gender"
              placeholder="请选择性别"
              :options="genderOptions"
            />
          </GlobalFormItem>

          <GlobalFormItem>
            <GlobalButton type="primary" @click="submitValidationForm">提交</GlobalButton>
            <GlobalButton @click="resetValidationForm" style="margin-left: 10px;">重置</GlobalButton>
            <GlobalButton type="warning" @click="validateSpecificField" style="margin-left: 10px;">验证用户名</GlobalButton>
          </GlobalFormItem>
        </GlobalForm>
      </div>
    </GlobalCodeDemo>

    <!-- 内联表单 -->
    <GlobalCodeDemo
      title="内联表单"
      description="表单项在一行内排列，适合简单的搜索表单"
      :code="inlineFormCode"
    >
      <div class="bg-gray-50 p-4 rounded-md">
        <GlobalForm
          :model="inlineForm"
          :inline="true"
        >
          <GlobalFormItem label="关键词">
            <GlobalInput v-model="inlineForm.keyword" placeholder="搜索关键词" />
          </GlobalFormItem>

          <GlobalFormItem label="类型">
            <GlobalSelect
              v-model="inlineForm.type"
              placeholder="请选择类型"
              :options="typeOptions"
            />
          </GlobalFormItem>

          <GlobalFormItem>
            <GlobalButton type="primary" @click="submitInlineForm">搜索</GlobalButton>
          </GlobalFormItem>
        </GlobalForm>
      </div>
    </GlobalCodeDemo>

    <!-- 顶部标签表单 -->
    <GlobalCodeDemo
      title="顶部标签表单"
      description="标签位置在输入框上方，适合较宽的表单布局"
      :code="topLabelFormCode"
    >
      <div class="bg-gray-50 p-4 rounded-md">
        <GlobalForm
          ref="topLabelFormRef"
          :model="topLabelForm"
          :rules="topLabelRules"
          label-position="top"
        >
          <div class="grid grid-cols-2 gap-4">
            <GlobalFormItem label="姓名" prop="name">
              <GlobalInput v-model="topLabelForm.name" placeholder="请输入姓名" />
            </GlobalFormItem>

            <GlobalFormItem label="电话" prop="phone">
              <GlobalInput v-model="topLabelForm.phone" placeholder="请输入电话号码" />
            </GlobalFormItem>
          </div>

          <GlobalFormItem label="地址" prop="address">
            <GlobalInput v-model="topLabelForm.address" placeholder="请输入详细地址" />
          </GlobalFormItem>

          <GlobalFormItem label="备注" prop="remark">
            <GlobalTextarea v-model="topLabelForm.remark" placeholder="请输入备注" :rows="3" />
          </GlobalFormItem>

          <GlobalFormItem>
            <GlobalButton type="primary" @click="submitTopLabelForm">提交</GlobalButton>
            <GlobalButton @click="resetTopLabelForm" style="margin-left: 10px;">重置</GlobalButton>
          </GlobalFormItem>
        </GlobalForm>
      </div>
    </GlobalCodeDemo>

    <!-- 禁用状态 -->
    <GlobalCodeDemo
      title="禁用状态"
      description="通过disabled属性禁用整个表单"
      :code="disabledFormCode"
    >
      <div class="bg-gray-50 p-4 rounded-md">
        <div class="mb-4">
          <GlobalButton @click="toggleDisabled">
            {{ formDisabled ? '启用表单' : '禁用表单' }}
          </GlobalButton>
        </div>

        <GlobalForm
          :model="disabledForm"
          :disabled="formDisabled"
          label-width="80px"
        >
          <GlobalFormItem label="用户名">
            <GlobalInput v-model="disabledForm.username" placeholder="用户名" />
          </GlobalFormItem>

          <GlobalFormItem label="状态">
            <GlobalSelect
              v-model="disabledForm.status"
              placeholder="请选择状态"
              :options="statusOptions"
            />
          </GlobalFormItem>

          <GlobalFormItem>
            <GlobalButton type="primary">提交</GlobalButton>
          </GlobalFormItem>
        </GlobalForm>
      </div>
    </GlobalCodeDemo>

    <!-- 尺寸示例 -->
    <GlobalCodeDemo
      title="不同尺寸"
      description="表单支持small、default、large三种尺寸"
      :code="sizeFormCode"
    >
      <div class="space-y-4">
        <div>
          <h5 class="text-sm font-medium mb-2">小尺寸 (small)</h5>
          <div class="bg-gray-50 p-4 rounded-md">
            <GlobalForm :model="sizeForm" size="small" label-width="60px">
              <GlobalFormItem label="姓名">
                <GlobalInput v-model="sizeForm.name" placeholder="请输入姓名" />
              </GlobalFormItem>
              <GlobalFormItem label="年龄">
                <GlobalInput v-model="sizeForm.age" placeholder="请输入年龄" />
              </GlobalFormItem>
            </GlobalForm>
          </div>
        </div>

        <div>
          <h5 class="text-sm font-medium mb-2">默认尺寸 (default)</h5>
          <div class="bg-gray-50 p-4 rounded-md">
            <GlobalForm :model="sizeForm" size="default" label-width="60px">
              <GlobalFormItem label="姓名">
                <GlobalInput v-model="sizeForm.name" placeholder="请输入姓名" />
              </GlobalFormItem>
              <GlobalFormItem label="年龄">
                <GlobalInput v-model="sizeForm.age" placeholder="请输入年龄" />
              </GlobalFormItem>
            </GlobalForm>
          </div>
        </div>

        <div>
          <h5 class="text-sm font-medium mb-2">大尺寸 (large)</h5>
          <div class="bg-gray-50 p-4 rounded-md">
            <GlobalForm :model="sizeForm" size="large" label-width="60px">
              <GlobalFormItem label="姓名">
                <GlobalInput v-model="sizeForm.name" placeholder="请输入姓名" />
              </GlobalFormItem>
              <GlobalFormItem label="年龄">
                <GlobalInput v-model="sizeForm.age" placeholder="请输入年龄" />
              </GlobalFormItem>
            </GlobalForm>
          </div>
        </div>
      </div>
    </GlobalCodeDemo>

    <!-- API 测试 -->
    <div class="mt-8 p-4 bg-blue-50 rounded-md">
      <h4 class="font-medium text-blue-900 mb-4">API 测试</h4>
      <div class="flex gap-2 flex-wrap mb-4">
        <GlobalButton size="small" @click="validateAllFields">验证全部字段</GlobalButton>
        <GlobalButton size="small" @click="clearValidation">清除验证</GlobalButton>
        <GlobalButton size="small" @click="resetFields">重置字段</GlobalButton>
      </div>
      <div v-if="apiTestResult" class="mt-4 p-3 bg-white rounded border text-sm">
        <strong>测试结果：</strong>
        <pre>{{ apiTestResult }}</pre>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import type { FormInstance, FormRules } from '../../../components/Form/types'

defineOptions({
  name: 'FormTest'
})

// 表单引用
const basicFormRef = ref<FormInstance>()
const validationFormRef = ref<FormInstance>()
const topLabelFormRef = ref<FormInstance>()

// 状态
const formDisabled = ref(false)
const apiTestResult = ref('')

// 基础表单
const basicForm = reactive({
  username: '',
  email: '',
  description: ''
})

const basicRules: FormRules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' }
  ],
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
  ],
  description: [
    { required: true, message: '请输入描述', trigger: 'blur' }
  ]
}

// 验证表单
const validationForm = reactive({
  username: '',
  password: '',
  confirmPassword: '',
  age: '',
  website: '',
  gender: ''
})

const validationRules: FormRules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 3, max: 20, message: '长度在 3 到 20 个字符', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '密码长度不能少于6位', trigger: 'blur' },
    {
      pattern: /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d@$!%*?&]{6,}$/,
      message: '密码必须包含大小写字母和数字',
      trigger: 'blur'
    }
  ],
  confirmPassword: [
    { required: true, message: '请再次输入密码', trigger: 'blur' },
    {
      validator: (rule, value, callback, source) => {
        if (value !== source?.password) {
          callback('两次输入的密码不一致')
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
  ],
  age: [
    { required: true, message: '请输入年龄', trigger: 'blur' },
    { type: 'number', message: '请输入有效数字', trigger: 'blur' },
    { min: 18, max: 120, message: '年龄必须在18-120之间', trigger: 'blur' }
  ],
  website: [
    { type: 'url', message: '请输入正确的网站地址', trigger: 'blur' }
  ],
  gender: [
    { required: true, message: '请选择性别', trigger: 'change' }
  ]
}

// 内联表单
const inlineForm = reactive({
  keyword: '',
  type: ''
})

// 顶部标签表单
const topLabelForm = reactive({
  name: '',
  phone: '',
  address: '',
  remark: ''
})

const topLabelRules: FormRules = {
  name: [
    { required: true, message: '请输入姓名', trigger: 'blur' }
  ],
  phone: [
    { required: true, message: '请输入电话号码', trigger: 'blur' },
    { pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号码', trigger: 'blur' }
  ]
}

// 禁用表单
const disabledForm = reactive({
  username: 'admin',
  status: 'active'
})

// 尺寸表单
const sizeForm = reactive({
  name: '',
  age: ''
})

// 选项数据
const genderOptions = [
  { label: '男', value: 'male' },
  { label: '女', value: 'female' }
]

const typeOptions = [
  { label: '全部', value: '' },
  { label: '文章', value: 'article' },
  { label: '视频', value: 'video' },
  { label: '图片', value: 'image' }
]

const statusOptions = [
  { label: '启用', value: 'active' },
  { label: '禁用', value: 'inactive' }
]

// 代码示例字符串
const basicFormCode = `<template>
  <GlobalForm
    ref="basicFormRef"
    :model="basicForm"
    :rules="basicRules"
    label-width="80px"
  >
    <GlobalFormItem label="用户名" prop="username">
      <GlobalInput v-model="basicForm.username" placeholder="请输入用户名" />
    </GlobalFormItem>

    <GlobalFormItem label="邮箱" prop="email">
      <GlobalInput v-model="basicForm.email" type="email" placeholder="请输入邮箱" />
    </GlobalFormItem>

    <GlobalFormItem label="描述" prop="description">
      <GlobalTextarea v-model="basicForm.description" placeholder="请输入描述" />
    </GlobalFormItem>

    <GlobalFormItem>
      <GlobalButton type="primary" @click="submitBasicForm">提交</GlobalButton>
      <GlobalButton @click="resetBasicForm">重置</GlobalButton>
    </GlobalFormItem>
  </GlobalForm>
</template>`

const validationFormCode = `<template>
  <GlobalForm
    ref="validationFormRef"
    :model="validationForm"
    :rules="validationRules"
    label-width="120px"
    label-position="right"
  >
    <GlobalFormItem label="用户名" prop="username">
      <GlobalInput v-model="validationForm.username" placeholder="请输入用户名" />
    </GlobalFormItem>

    <GlobalFormItem label="密码" prop="password">
      <GlobalInput
        v-model="validationForm.password"
        type="password"
        placeholder="请输入密码"
        show-password
      />
    </GlobalFormItem>

    <GlobalFormItem label="确认密码" prop="confirmPassword">
      <GlobalInput
        v-model="validationForm.confirmPassword"
        type="password"
        placeholder="请再次输入密码"
        show-password
      />
    </GlobalFormItem>

    <GlobalFormItem label="年龄" prop="age">
      <GlobalInput v-model="validationForm.age" type="number" placeholder="请输入年龄" />
    </GlobalFormItem>

    <GlobalFormItem label="性别" prop="gender">
      <GlobalSelect
        v-model="validationForm.gender"
        placeholder="请选择性别"
        :options="genderOptions"
      />
    </GlobalFormItem>

    <GlobalFormItem>
      <GlobalButton type="primary" @click="submitValidationForm">提交</GlobalButton>
      <GlobalButton @click="resetValidationForm">重置</GlobalButton>
    </GlobalFormItem>
  </GlobalForm>
</template>`

const inlineFormCode = `<template>
  <GlobalForm :model="inlineForm" :inline="true">
    <GlobalFormItem label="关键词">
      <GlobalInput v-model="inlineForm.keyword" placeholder="搜索关键词" />
    </GlobalFormItem>

    <GlobalFormItem label="类型">
      <GlobalSelect
        v-model="inlineForm.type"
        placeholder="请选择类型"
        :options="typeOptions"
      />
    </GlobalFormItem>

    <GlobalFormItem>
      <GlobalButton type="primary" @click="submitInlineForm">搜索</GlobalButton>
    </GlobalFormItem>
  </GlobalForm>
</template>`

const topLabelFormCode = `<template>
  <GlobalForm
    ref="topLabelFormRef"
    :model="topLabelForm"
    :rules="topLabelRules"
    label-position="top"
  >
    <div class="grid grid-cols-2 gap-4">
      <GlobalFormItem label="姓名" prop="name">
        <GlobalInput v-model="topLabelForm.name" placeholder="请输入姓名" />
      </GlobalFormItem>

      <GlobalFormItem label="电话" prop="phone">
        <GlobalInput v-model="topLabelForm.phone" placeholder="请输入电话号码" />
      </GlobalFormItem>
    </div>

    <GlobalFormItem label="地址" prop="address">
      <GlobalInput v-model="topLabelForm.address" placeholder="请输入详细地址" />
    </GlobalFormItem>

    <GlobalFormItem label="备注" prop="remark">
      <GlobalTextarea v-model="topLabelForm.remark" placeholder="请输入备注" :rows="3" />
    </GlobalFormItem>

    <GlobalFormItem>
      <GlobalButton type="primary" @click="submitTopLabelForm">提交</GlobalButton>
      <GlobalButton @click="resetTopLabelForm">重置</GlobalButton>
    </GlobalFormItem>
  </GlobalForm>
</template>`

const disabledFormCode = `<template>
  <div>
    <GlobalButton @click="toggleDisabled">
      {{ formDisabled ? '启用表单' : '禁用表单' }}
    </GlobalButton>

    <GlobalForm
      :model="disabledForm"
      :disabled="formDisabled"
      label-width="80px"
    >
      <GlobalFormItem label="用户名">
        <GlobalInput v-model="disabledForm.username" placeholder="用户名" />
      </GlobalFormItem>

      <GlobalFormItem label="状态">
        <GlobalSelect
          v-model="disabledForm.status"
          placeholder="请选择状态"
          :options="statusOptions"
        />
      </GlobalFormItem>

      <GlobalFormItem>
        <GlobalButton type="primary">提交</GlobalButton>
      </GlobalFormItem>
    </GlobalForm>
  </div>
</template>`

const sizeFormCode = `<template>
  <div>
    <!-- 小尺寸 -->
    <GlobalForm :model="sizeForm" size="small" label-width="60px">
      <GlobalFormItem label="姓名">
        <GlobalInput v-model="sizeForm.name" placeholder="请输入姓名" />
      </GlobalFormItem>
      <GlobalFormItem label="年龄">
        <GlobalInput v-model="sizeForm.age" placeholder="请输入年龄" />
      </GlobalFormItem>
    </GlobalForm>

    <!-- 默认尺寸 -->
    <GlobalForm :model="sizeForm" size="default" label-width="60px">
      <GlobalFormItem label="姓名">
        <GlobalInput v-model="sizeForm.name" placeholder="请输入姓名" />
      </GlobalFormItem>
      <GlobalFormItem label="年龄">
        <GlobalInput v-model="sizeForm.age" placeholder="请输入年龄" />
      </GlobalFormItem>
    </GlobalForm>

    <!-- 大尺寸 -->
    <GlobalForm :model="sizeForm" size="large" label-width="60px">
      <GlobalFormItem label="姓名">
        <GlobalInput v-model="sizeForm.name" placeholder="请输入姓名" />
      </GlobalFormItem>
      <GlobalFormItem label="年龄">
        <GlobalInput v-model="sizeForm.age" placeholder="请输入年龄" />
      </GlobalFormItem>
    </GlobalForm>
  </div>
</template>`

// 方法
const submitBasicForm = async () => {
  try {
    const valid = await basicFormRef.value?.validate()
    if (valid) {
      apiTestResult.value = `基础表单提交成功: ${JSON.stringify(basicForm, null, 2)}`
    } else {
      apiTestResult.value = '基础表单验证失败'
    }
  } catch (error) {
    apiTestResult.value = `基础表单验证失败: ${error}`
  }
}

const resetBasicForm = () => {
  basicFormRef.value?.resetFields()
  apiTestResult.value = '基础表单已重置'
}

const submitValidationForm = async () => {
  try {
    const valid = await validationFormRef.value?.validate()
    if (valid) {
      apiTestResult.value = `验证表单提交成功: ${JSON.stringify(validationForm, null, 2)}`
    } else {
      apiTestResult.value = '验证表单验证失败'
    }
  } catch (error) {
    apiTestResult.value = `验证表单验证失败: ${error}`
  }
}

const resetValidationForm = () => {
  validationFormRef.value?.resetFields()
  apiTestResult.value = '验证表单已重置'
}

const validateSpecificField = async () => {
  try {
    const valid = await validationFormRef.value?.validateField('username')
    apiTestResult.value = `用户名字段验证结果: ${valid ? '通过' : '失败'}`
  } catch (error) {
    apiTestResult.value = `用户名字段验证失败: ${error}`
  }
}

const submitInlineForm = () => {
  apiTestResult.value = `内联表单搜索: ${JSON.stringify(inlineForm, null, 2)}`
}

const submitTopLabelForm = async () => {
  try {
    const valid = await topLabelFormRef.value?.validate()
    if (valid) {
      apiTestResult.value = `顶部标签表单提交成功: ${JSON.stringify(topLabelForm, null, 2)}`
    }
  } catch (error) {
    apiTestResult.value = `顶部标签表单验证失败: ${error}`
  }
}

const resetTopLabelForm = () => {
  topLabelFormRef.value?.resetFields()
  apiTestResult.value = '顶部标签表单已重置'
}

const toggleDisabled = () => {
  formDisabled.value = !formDisabled.value
  apiTestResult.value = `表单${formDisabled.value ? '已禁用' : '已启用'}`
}

// API 测试方法
const validateAllFields = async () => {
  try {
    const valid = await validationFormRef.value?.validate()
    apiTestResult.value = `整个表单验证结果: ${valid ? '通过' : '失败'}`
  } catch (error) {
    apiTestResult.value = `表单验证出错: ${error}`
  }
}

const clearValidation = () => {
  validationFormRef.value?.clearValidation()
  apiTestResult.value = '已清除验证状态'
}

const resetFields = () => {
  validationFormRef.value?.resetFields()
  apiTestResult.value = '已重置所有字段'
}
</script>

<style scoped>
/* 表单样式调整 */
:deep(.global-form-item) {
  margin-bottom: 1rem;
}

:deep(.global-form--inline .global-form-item) {
  margin-bottom: 0;
  margin-right: 1rem;
}

/* 网格布局样式 */
.grid {
  display: grid;
}

.grid-cols-2 {
  grid-template-columns: repeat(2, minmax(0, 1fr));
}

.gap-4 {
  gap: 1rem;
}
</style>
