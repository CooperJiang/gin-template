<template>
  <div class="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
    <h3 class="text-lg font-semibold mb-4">Form 表单组件测试</h3>
    <p class="text-gray-600 text-sm mb-6">测试表单组件的各种配置、验证规则和功能</p>

    <!-- 基础表单 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">基础表单</h4>
      <div class="bg-gray-50 p-4 rounded-md">
        <GlobalForm
          v-model="basicFormData"
          :items="basicFormItems"
          @submit="handleBasicSubmit"
          @reset="handleBasicReset"
          @validate="handleValidate"
        />
        <div class="mt-4 p-3 bg-white rounded border">
          <h5 class="text-sm font-medium mb-2">表单数据：</h5>
          <pre class="text-xs text-gray-600">{{ JSON.stringify(basicFormData, null, 2) }}</pre>
        </div>
      </div>
    </div>

    <!-- 水平布局表单 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">水平布局表单</h4>
      <div class="bg-gray-50 p-4 rounded-md">
        <GlobalForm
          v-model="horizontalFormData"
          :items="horizontalFormItems"
          layout="horizontal"
          label-width="120px"
          @submit="handleHorizontalSubmit"
        />
      </div>
    </div>

    <!-- 内联表单 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">内联表单</h4>
      <div class="bg-gray-50 p-4 rounded-md">
        <GlobalForm
          v-model="inlineFormData"
          :items="inlineFormItems"
          layout="inline"
          submit-text="搜索"
          :show-reset-button="false"
          @submit="handleInlineSubmit"
        />
      </div>
    </div>

    <!-- 验证表单 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">表单验证</h4>
      <div class="bg-gray-50 p-4 rounded-md">
        <GlobalForm
          ref="validationFormRef"
          v-model="validationFormData"
          :items="validationFormItems"
          :validate-on-change="true"
          :validate-on-blur="true"
          @submit="handleValidationSubmit"
        />
        <div class="mt-4 flex gap-2">
          <GlobalButton
            type="secondary"
            variant="outline"
            @click="validateForm"
          >
            手动验证
          </GlobalButton>
          <GlobalButton
            type="warning"
            variant="outline"
            @click="resetValidation"
          >
            重置验证
          </GlobalButton>
        </div>
      </div>
    </div>

    <!-- 复杂表单 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">复杂表单示例</h4>
      <div class="bg-gray-50 p-4 rounded-md">
        <GlobalForm
          v-model="complexFormData"
          :items="complexFormItems"
          :submit-loading="isSubmitting"
          bordered
          @submit="handleComplexSubmit"
        />
      </div>
    </div>

    <!-- 禁用表单 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">禁用状态</h4>
      <div class="bg-gray-50 p-4 rounded-md">
        <div class="mb-4">
          <GlobalButton
            type="secondary"
            variant="outline"
            @click="toggleDisabled"
          >
            {{ formDisabled ? '启用表单' : '禁用表单' }}
          </GlobalButton>
        </div>
        <GlobalForm
          v-model="disabledFormData"
          :items="disabledFormItems"
          :disabled="formDisabled"
        />
      </div>
    </div>

    <!-- 表单尺寸 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">表单尺寸</h4>
      <div class="space-y-4">
        <div>
          <h5 class="text-sm font-medium text-gray-700 mb-2">小尺寸</h5>
          <div class="bg-gray-50 p-4 rounded-md">
            <GlobalForm
              v-model="sizeFormData"
              :items="sizeFormItems"
              size="small"
              layout="inline"
              :show-reset-button="false"
            />
          </div>
        </div>
        <div>
          <h5 class="text-sm font-medium text-gray-700 mb-2">中等尺寸（默认）</h5>
          <div class="bg-gray-50 p-4 rounded-md">
            <GlobalForm
              v-model="sizeFormData"
              :items="sizeFormItems"
              size="medium"
              layout="inline"
              :show-reset-button="false"
            />
          </div>
        </div>
        <div>
          <h5 class="text-sm font-medium text-gray-700 mb-2">大尺寸</h5>
          <div class="bg-gray-50 p-4 rounded-md">
            <GlobalForm
              v-model="sizeFormData"
              :items="sizeFormItems"
              size="large"
              layout="inline"
              :show-reset-button="false"
            />
          </div>
        </div>
      </div>
    </div>

    <!-- 独立组件测试 -->
    <div class="mb-8">
      <h4 class="font-medium text-gray-900 mb-4">独立组件测试</h4>
      <div class="bg-gray-50 p-4 rounded-md space-y-4">
        <!-- Input 组件测试 -->
        <div>
          <h5 class="text-sm font-medium text-gray-700 mb-2">GlobalInput 组件</h5>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <GlobalInput
              v-model="inputValue"
              placeholder="普通输入框"
              prefix-icon="user"
            />
            <GlobalInput
              v-model="passwordValue"
              type="password"
              placeholder="密码输入框"
              :show-password="true"
            />
            <GlobalInput
              v-model="emailValue"
              type="email"
              placeholder="邮箱输入框"
              suffix-icon="envelope"
            />
            <GlobalInput
              v-model="disabledValue"
              placeholder="禁用状态"
              :disabled="true"
            />
          </div>
        </div>

        <!-- Textarea 组件测试 -->
        <div>
          <h5 class="text-sm font-medium text-gray-700 mb-2">GlobalTextarea 组件</h5>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <GlobalTextarea
              v-model="textareaValue"
              placeholder="普通文本域"
              :rows="3"
            />
            <GlobalTextarea
              v-model="textareaWithCountValue"
              placeholder="带字符计数的文本域"
              :maxlength="200"
              :show-count="true"
              :rows="3"
            />
          </div>
        </div>

        <div class="mt-4 p-3 bg-white rounded border">
          <h5 class="text-sm font-medium mb-2">独立组件数据：</h5>
          <pre class="text-xs text-gray-600">输入框: {{ inputValue }}
密码: {{ passwordValue }}
邮箱: {{ emailValue }}
文本域: {{ textareaValue }}
带计数文本域: {{ textareaWithCountValue }}</pre>
        </div>
      </div>
    </div>

    <!-- API 测试区域 -->
    <div class="bg-gray-50 p-4 rounded-md">
      <h4 class="font-medium text-gray-900 mb-3">API 测试区域</h4>
      <div class="space-y-4">
        <div>
          <h5 class="text-sm font-medium text-gray-700 mb-2">表单方法测试</h5>
          <div class="flex flex-wrap gap-2">
            <GlobalButton
              type="primary"
              size="small"
              @click="setFieldValue"
            >
              设置用户名为 "admin"
            </GlobalButton>
            <GlobalButton
              type="secondary"
              size="small"
              @click="getFieldValue"
            >
              获取邮箱值
            </GlobalButton>
            <GlobalButton
              type="warning"
              size="small"
              @click="validateField"
            >
              验证邮箱字段
            </GlobalButton>
          </div>
        </div>
        <div>
          <h5 class="text-sm font-medium text-gray-700 mb-2">操作结果：</h5>
          <div class="p-3 bg-white rounded border text-xs">
            <pre>{{ apiTestResult || '暂无操作结果' }}</pre>
          </div>
        </div>
      </div>
    </div>

    <!-- API 文档 -->
    <div class="mt-8 border-t pt-8">
      <h3 class="text-lg font-semibold mb-4">API 文档</h3>

      <!-- Props -->
      <div class="mb-6">
        <h4 class="font-medium text-gray-900 mb-3">Props</h4>
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200 border border-gray-200 rounded">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">属性</th>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">类型</th>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">默认值</th>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">说明</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">modelValue</td>
                <td class="px-4 py-2 text-sm text-gray-500">Record&lt;string, any&gt;</td>
                <td class="px-4 py-2 text-sm text-gray-500">{}</td>
                <td class="px-4 py-2 text-sm text-gray-500">表单数据对象</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">items</td>
                <td class="px-4 py-2 text-sm text-gray-500">FormItem[]</td>
                <td class="px-4 py-2 text-sm text-gray-500">[]</td>
                <td class="px-4 py-2 text-sm text-gray-500">表单项配置数组</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">layout</td>
                <td class="px-4 py-2 text-sm text-gray-500">'vertical' | 'horizontal' | 'inline'</td>
                <td class="px-4 py-2 text-sm text-gray-500">'vertical'</td>
                <td class="px-4 py-2 text-sm text-gray-500">表单布局方式</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">size</td>
                <td class="px-4 py-2 text-sm text-gray-500">'small' | 'medium' | 'large'</td>
                <td class="px-4 py-2 text-sm text-gray-500">'medium'</td>
                <td class="px-4 py-2 text-sm text-gray-500">表单组件尺寸</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">labelWidth</td>
                <td class="px-4 py-2 text-sm text-gray-500">string | number</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">标签宽度（水平布局时有效）</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">validateOnChange</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean</td>
                <td class="px-4 py-2 text-sm text-gray-500">false</td>
                <td class="px-4 py-2 text-sm text-gray-500">值变化时是否验证</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">validateOnBlur</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean</td>
                <td class="px-4 py-2 text-sm text-gray-500">true</td>
                <td class="px-4 py-2 text-sm text-gray-500">失去焦点时是否验证</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">showSubmitButton</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean</td>
                <td class="px-4 py-2 text-sm text-gray-500">true</td>
                <td class="px-4 py-2 text-sm text-gray-500">是否显示提交按钮</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">showResetButton</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean</td>
                <td class="px-4 py-2 text-sm text-gray-500">true</td>
                <td class="px-4 py-2 text-sm text-gray-500">是否显示重置按钮</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">submitLoading</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean</td>
                <td class="px-4 py-2 text-sm text-gray-500">false</td>
                <td class="px-4 py-2 text-sm text-gray-500">提交按钮加载状态</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">disabled</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean</td>
                <td class="px-4 py-2 text-sm text-gray-500">false</td>
                <td class="px-4 py-2 text-sm text-gray-500">是否禁用整个表单</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">bordered</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean</td>
                <td class="px-4 py-2 text-sm text-gray-500">false</td>
                <td class="px-4 py-2 text-sm text-gray-500">是否显示表单边框</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Events -->
      <div class="mb-6">
        <h4 class="font-medium text-gray-900 mb-3">Events</h4>
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200 border border-gray-200 rounded">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">事件名</th>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">参数</th>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">说明</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">update:modelValue</td>
                <td class="px-4 py-2 text-sm text-gray-500">(value: Record&lt;string, any&gt;)</td>
                <td class="px-4 py-2 text-sm text-gray-500">表单数据变化时触发</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">submit</td>
                <td class="px-4 py-2 text-sm text-gray-500">(value: Record&lt;string, any&gt;)</td>
                <td class="px-4 py-2 text-sm text-gray-500">表单提交时触发</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">reset</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">表单重置时触发</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">validate</td>
                <td class="px-4 py-2 text-sm text-gray-500">(valid: boolean, errors: Record&lt;string, string[]&gt;)</td>
                <td class="px-4 py-2 text-sm text-gray-500">表单验证时触发</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">field-change</td>
                <td class="px-4 py-2 text-sm text-gray-500">(field: string, value: any)</td>
                <td class="px-4 py-2 text-sm text-gray-500">字段值变化时触发</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">field-blur</td>
                <td class="px-4 py-2 text-sm text-gray-500">(field: string, value: any)</td>
                <td class="px-4 py-2 text-sm text-gray-500">字段失去焦点时触发</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Methods -->
      <div class="mb-6">
        <h4 class="font-medium text-gray-900 mb-3">Methods</h4>
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200 border border-gray-200 rounded">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">方法名</th>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">参数</th>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">返回值</th>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">说明</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">validate</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">Promise&lt;{valid: boolean, errors: Record&lt;string, string[]&gt;}&gt;</td>
                <td class="px-4 py-2 text-sm text-gray-500">验证整个表单</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">validateField</td>
                <td class="px-4 py-2 text-sm text-gray-500">(field: string)</td>
                <td class="px-4 py-2 text-sm text-gray-500">Promise&lt;{valid: boolean, errors: string[]}&gt;</td>
                <td class="px-4 py-2 text-sm text-gray-500">验证单个字段</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">resetValidation</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">void</td>
                <td class="px-4 py-2 text-sm text-gray-500">重置验证状态</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">reset</td>
                <td class="px-4 py-2 text-sm text-gray-500">-</td>
                <td class="px-4 py-2 text-sm text-gray-500">void</td>
                <td class="px-4 py-2 text-sm text-gray-500">重置表单数据</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">getFieldValue</td>
                <td class="px-4 py-2 text-sm text-gray-500">(field: string)</td>
                <td class="px-4 py-2 text-sm text-gray-500">any</td>
                <td class="px-4 py-2 text-sm text-gray-500">获取字段值</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">setFieldValue</td>
                <td class="px-4 py-2 text-sm text-gray-500">(field: string, value: any)</td>
                <td class="px-4 py-2 text-sm text-gray-500">void</td>
                <td class="px-4 py-2 text-sm text-gray-500">设置字段值</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- FormItem -->
      <div class="mb-6">
        <h4 class="font-medium text-gray-900 mb-3">FormItem 配置</h4>
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200 border border-gray-200 rounded">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">属性</th>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">类型</th>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">说明</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">name</td>
                <td class="px-4 py-2 text-sm text-gray-500">string</td>
                <td class="px-4 py-2 text-sm text-gray-500">字段名称，必填</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">label</td>
                <td class="px-4 py-2 text-sm text-gray-500">string</td>
                <td class="px-4 py-2 text-sm text-gray-500">字段标签</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">type</td>
                <td class="px-4 py-2 text-sm text-gray-500">'input' | 'password' | 'select' | 'checkbox' | 'radio' | 'textarea'</td>
                <td class="px-4 py-2 text-sm text-gray-500">组件类型</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">placeholder</td>
                <td class="px-4 py-2 text-sm text-gray-500">string</td>
                <td class="px-4 py-2 text-sm text-gray-500">占位符文本</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">required</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean</td>
                <td class="px-4 py-2 text-sm text-gray-500">是否必填</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">rules</td>
                <td class="px-4 py-2 text-sm text-gray-500">FormRule[]</td>
                <td class="px-4 py-2 text-sm text-gray-500">验证规则数组</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">options</td>
                <td class="px-4 py-2 text-sm text-gray-500">FormOption[]</td>
                <td class="px-4 py-2 text-sm text-gray-500">选项数据（select/checkbox/radio）</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">defaultValue</td>
                <td class="px-4 py-2 text-sm text-gray-500">any</td>
                <td class="px-4 py-2 text-sm text-gray-500">默认值</td>
              </tr>
              <tr>
                <td class="px-4 py-2 text-sm font-medium text-gray-900">disabled</td>
                <td class="px-4 py-2 text-sm text-gray-500">boolean</td>
                <td class="px-4 py-2 text-sm text-gray-500">是否禁用</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- 使用示例 -->
      <div>
        <h4 class="font-medium text-gray-900 mb-3">使用示例</h4>
        <div class="bg-gray-50 p-4 rounded-md">
          <pre class="text-sm text-gray-800 overflow-x-auto"><code>&lt;template&gt;
  &lt;GlobalForm
    v-model="formData"
    :items="formItems"
    layout="vertical"
    :validate-on-blur="true"
    @submit="handleSubmit"
    @validate="handleValidate"
  /&gt;
&lt;/template&gt;

&lt;script setup&gt;
import { ref } from 'vue'

const formData = ref({})

const formItems = [
  {
    name: 'username',
    label: '用户名',
    type: 'input',
    placeholder: '请输入用户名',
    rules: [
      { type: 'required', message: '用户名不能为空' },
      { min: 3, max: 20, message: '用户名长度为3-20个字符' }
    ]
  },
  {
    name: 'email',
    label: '邮箱',
    type: 'input',
    inputType: 'email',
    placeholder: '请输入邮箱',
    rules: [
      { type: 'required', message: '邮箱不能为空' },
      { type: 'email', message: '请输入正确的邮箱格式' }
    ]
  }
]

const handleSubmit = (data) => {
  console.log('表单提交:', data)
}

const handleValidate = (valid, errors) => {
  console.log('验证结果:', valid, errors)
}
&lt;/script&gt;</code></pre>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import type { FormInstance, FormItem } from '../../../components/Form/types'

defineOptions({
  name: 'FormTest'
})

// 引用
const validationFormRef = ref<FormInstance>()

// 表单数据
const basicFormData = ref({})
const horizontalFormData = ref({})
const inlineFormData = ref({})
const validationFormData = ref({})
const complexFormData = ref({})
const disabledFormData = ref({})
const sizeFormData = ref({})

// 独立组件测试数据
const inputValue = ref('')
const passwordValue = ref('')
const emailValue = ref('')
const disabledValue = ref('禁用状态示例')
const textareaValue = ref('')
const textareaWithCountValue = ref('')

// 状态
const isSubmitting = ref(false)
const formDisabled = ref(false)
const apiTestResult = ref('')

// 基础表单配置
const basicFormItems: FormItem[] = [
  {
    name: 'username',
    label: '用户名',
    type: 'input',
    placeholder: '请输入用户名',
    defaultValue: ''
  },
  {
    name: 'email',
    label: '邮箱',
    type: 'input',
    inputType: 'email',
    placeholder: '请输入邮箱地址'
  },
  {
    name: 'description',
    label: '描述',
    type: 'textarea',
    placeholder: '请输入描述信息'
  }
]

// 水平布局表单配置
const horizontalFormItems: FormItem[] = [
  {
    name: 'title',
    label: '标题',
    type: 'input',
    placeholder: '请输入标题',
    required: true
  },
  {
    name: 'category',
    label: '分类',
    type: 'select',
    placeholder: '请选择分类',
    options: [
      { label: '技术文章', value: 'tech' },
      { label: '生活随笔', value: 'life' },
      { label: '工作总结', value: 'work' }
    ]
  },
  {
    name: 'tags',
    label: '标签',
    type: 'select',
    placeholder: '请选择标签',
    multiple: true,
    options: [
      { label: 'Vue', value: 'vue' },
      { label: 'React', value: 'react' },
      { label: 'TypeScript', value: 'typescript' },
      { label: 'JavaScript', value: 'javascript' }
    ]
  }
]

// 内联表单配置
const inlineFormItems: FormItem[] = [
  {
    name: 'keyword',
    label: '关键词',
    type: 'input',
    placeholder: '搜索关键词'
  },
  {
    name: 'type',
    label: '类型',
    type: 'select',
    placeholder: '全部',
    options: [
      { label: '全部', value: '' },
      { label: '文章', value: 'article' },
      { label: '视频', value: 'video' },
      { label: '图片', value: 'image' }
    ]
  }
]

// 验证表单配置
const validationFormItems: FormItem[] = [
  {
    name: 'username',
    label: '用户名',
    type: 'input',
    placeholder: '请输入用户名',
    rules: [
      { type: 'required', message: '用户名不能为空' },
      { min: 3, max: 20, message: '用户名长度为3-20个字符' }
    ]
  },
  {
    name: 'email',
    label: '邮箱',
    type: 'input',
    inputType: 'email',
    placeholder: '请输入邮箱地址',
    rules: [
      { type: 'required', message: '邮箱不能为空' },
      { type: 'email', message: '请输入正确的邮箱格式' }
    ]
  },
  {
    name: 'password',
    label: '密码',
    type: 'password',
    placeholder: '请输入密码',
    showPassword: true,
    rules: [
      { type: 'required', message: '密码不能为空' },
      { min: 6, message: '密码长度不能少于6位' },
      {
        type: 'pattern',
        pattern: /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d@$!%*?&]{6,}$/,
        message: '密码必须包含大小写字母和数字'
      }
    ]
  },
  {
    name: 'confirmPassword',
    label: '确认密码',
    type: 'password',
    placeholder: '请再次输入密码',
    rules: [
      { type: 'required', message: '确认密码不能为空' },
      {
        type: 'custom',
        validator: (value: any, formData: Record<string, any>) => {
          return value === formData.password || '两次输入的密码不一致'
        }
      }
    ]
  },
  {
    name: 'agree',
    label: '同意条款',
    type: 'checkbox',
    rules: [
      {
        type: 'custom',
        validator: (value: any) => {
          return value === true || '请同意用户协议'
        }
      }
    ],
    options: [
      { label: '我已阅读并同意用户协议', value: true }
    ]
  }
]

// 复杂表单配置
const complexFormItems: FormItem[] = [
  {
    name: 'fullName',
    label: '姓名',
    type: 'input',
    placeholder: '请输入真实姓名',
    required: true,
    rules: [
      { type: 'required', message: '姓名不能为空' }
    ]
  },
  {
    name: 'gender',
    label: '性别',
    type: 'radio',
    options: [
      { label: '男', value: 'male' },
      { label: '女', value: 'female' }
    ],
    defaultValue: 'male'
  },
  {
    name: 'age',
    label: '年龄',
    type: 'input',
    inputType: 'number',
    placeholder: '请输入年龄',
    rules: [
      { type: 'required', message: '年龄不能为空' },
      { type: 'number', message: '请输入有效数字' },
      { min: 18, max: 120, message: '年龄必须在18-120之间' }
    ]
  },
  {
    name: 'hobbies',
    label: '兴趣爱好',
    type: 'checkbox',
    multiple: true,
    options: [
      { label: '阅读', value: 'reading' },
      { label: '运动', value: 'sports' },
      { label: '音乐', value: 'music' },
      { label: '旅行', value: 'travel' },
      { label: '编程', value: 'coding' }
    ]
  },
  {
    name: 'bio',
    label: '个人简介',
    type: 'textarea',
    placeholder: '请简单介绍一下自己...',
    maxLength: 500
  }
]

// 禁用表单配置
const disabledFormItems: FormItem[] = [
  {
    name: 'readonly',
    label: '只读字段',
    type: 'input',
    defaultValue: '这是一个只读字段',
    disabled: true
  },
  {
    name: 'normal',
    label: '正常字段',
    type: 'input',
    placeholder: '这个字段受全局禁用状态控制'
  }
]

// 尺寸表单配置
const sizeFormItems: FormItem[] = [
  {
    name: 'name',
    label: '姓名',
    type: 'input',
    placeholder: '请输入姓名'
  },
  {
    name: 'status',
    label: '状态',
    type: 'select',
    options: [
      { label: '启用', value: 'active' },
      { label: '禁用', value: 'inactive' }
    ]
  }
]

// 事件处理
const handleBasicSubmit = (data: Record<string, any>) => {
  console.log('基础表单提交:', data)
  apiTestResult.value = `基础表单提交: ${JSON.stringify(data, null, 2)}`
}

const handleBasicReset = () => {
  console.log('基础表单重置')
  apiTestResult.value = '基础表单已重置'
}

const handleValidate = (valid: boolean, errors: Record<string, string[]>) => {
  console.log('表单验证结果:', valid, errors)
  apiTestResult.value = `验证结果: ${valid ? '通过' : '失败'}\n错误信息: ${JSON.stringify(errors, null, 2)}`
}

const handleHorizontalSubmit = (data: Record<string, any>) => {
  console.log('水平布局表单提交:', data)
  apiTestResult.value = `水平布局表单提交: ${JSON.stringify(data, null, 2)}`
}

const handleInlineSubmit = (data: Record<string, any>) => {
  console.log('内联表单提交:', data)
  apiTestResult.value = `搜索参数: ${JSON.stringify(data, null, 2)}`
}

const handleValidationSubmit = (data: Record<string, any>) => {
  console.log('验证表单提交:', data)
  apiTestResult.value = `验证表单提交成功: ${JSON.stringify(data, null, 2)}`
}

const handleComplexSubmit = async (data: Record<string, any>) => {
  console.log('复杂表单提交:', data)
  isSubmitting.value = true

  // 模拟提交延迟
  await new Promise(resolve => setTimeout(resolve, 2000))

  isSubmitting.value = false
  apiTestResult.value = `复杂表单提交成功: ${JSON.stringify(data, null, 2)}`
}

// API 测试方法
const validateForm = async () => {
  if (!validationFormRef.value) return

  const result = await validationFormRef.value.validate()
  apiTestResult.value = `完整表单验证结果: ${JSON.stringify(result, null, 2)}`
}

const resetValidation = () => {
  if (!validationFormRef.value) return

  validationFormRef.value.resetValidation()
  apiTestResult.value = '验证状态已重置'
}

const setFieldValue = () => {
  if (!validationFormRef.value) return

  validationFormRef.value.setFieldValue('username', 'admin')
  apiTestResult.value = '已将用户名设置为 "admin"'
}

const getFieldValue = () => {
  if (!validationFormRef.value) return

  const email = validationFormRef.value.getFieldValue('email')
  apiTestResult.value = `邮箱字段当前值: ${email || '(空)'}`
}

const validateField = async () => {
  if (!validationFormRef.value) return

  const result = await validationFormRef.value.validateField('email')
  apiTestResult.value = `邮箱字段验证结果: ${JSON.stringify(result, null, 2)}`
}

const toggleDisabled = () => {
  formDisabled.value = !formDisabled.value
  apiTestResult.value = `表单${formDisabled.value ? '已禁用' : '已启用'}`
}
</script>

<style scoped>
/* 表单项间距 */
:deep(.global-form) {
  @apply space-y-4;
}

/* 内联表单样式调整 */
:deep(.form--inline .global-form) {
  @apply space-y-0 space-x-4;
}
</style>
