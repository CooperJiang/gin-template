import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { Card, Button, Input } from '@app/core'
import { createHttpClient, message } from '@app/core'

interface UserFormData {
  name: string
  email: string
  role: string
  password: string
}

interface FormErrors {
  name?: string
  email?: string
  role?: string
  password?: string
}

/**
 * 用户表单页面示例
 * 展示：表单验证、错误处理、提交状态
 */
export function UserForm() {
  const navigate = useNavigate()
  const [formData, setFormData] = useState<UserFormData>({
    name: '',
    email: '',
    role: 'user',
    password: '',
  })
  const [errors, setErrors] = useState<FormErrors>({})
  const [submitting, setSubmitting] = useState(false)

  const http = createHttpClient({ baseURL: '/api/v1' })

  const validate = (): boolean => {
    const newErrors: FormErrors = {}

    if (!formData.name.trim()) {
      newErrors.name = '请输入用户名'
    } else if (formData.name.length < 2) {
      newErrors.name = '用户名至少2个字符'
    }

    if (!formData.email.trim()) {
      newErrors.email = '请输入邮箱'
    } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(formData.email)) {
      newErrors.email = '邮箱格式不正确'
    }

    if (!formData.password) {
      newErrors.password = '请输入密码'
    } else if (formData.password.length < 6) {
      newErrors.password = '密码至少6个字符'
    }

    setErrors(newErrors)
    return Object.keys(newErrors).length === 0
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()

    if (!validate()) {
      return
    }

    setSubmitting(true)
    try {
      await http.post('/users', formData)
      message.success('用户创建成功')
      navigate('/examples/user/list')
    } catch (error) {
      // 错误已由全局处理器处理
    } finally {
      setSubmitting(false)
    }
  }

  const handleChange = (field: keyof UserFormData, value: string) => {
    setFormData({ ...formData, [field]: value })
    // 清除该字段的错误
    if (errors[field]) {
      setErrors({ ...errors, [field]: undefined })
    }
  }

  return (
    <div className="max-w-2xl mx-auto p-6">
      <Card title="添加用户">
        <form onSubmit={handleSubmit} className="space-y-4">
          <Input
            label="用户名"
            value={formData.name}
            onChange={(e) => handleChange('name', e.target.value)}
            error={errors.name}
            placeholder="请输入用户名"
            required
            fullWidth
          />

          <Input
            label="邮箱"
            type="email"
            value={formData.email}
            onChange={(e) => handleChange('email', e.target.value)}
            error={errors.email}
            placeholder="user@example.com"
            required
            fullWidth
          />

          <Input
            label="密码"
            type="password"
            value={formData.password}
            onChange={(e) => handleChange('password', e.target.value)}
            error={errors.password}
            helperText="至少6个字符"
            placeholder="请输入密码"
            required
            fullWidth
          />

          <div className="fullWidth">
            <label className="block text-sm font-medium text-app-text mb-1">
              角色 <span className="text-error ml-1">*</span>
            </label>
            <select
              value={formData.role}
              onChange={(e) => handleChange('role', e.target.value)}
              className="block w-full px-3 py-2 rounded-md border border-app-input-border bg-app-input-bg text-app-input-text focus:outline-none focus:ring-2 focus:ring-primary-500"
              required
            >
              <option value="user">普通用户</option>
              <option value="admin">管理员</option>
              <option value="moderator">版主</option>
            </select>
          </div>

          <div className="flex gap-3 pt-4">
            <Button
              type="submit"
              loading={submitting}
              disabled={submitting}
              fullWidth
            >
              创建用户
            </Button>
            <Button
              type="button"
              variant="outline"
              onClick={() => navigate('/examples/user/list')}
              disabled={submitting}
              fullWidth
            >
              取消
            </Button>
          </div>
        </form>
      </Card>

      {/* 表单数据预览（仅用于演示） */}
      <Card title="表单数据预览" className="mt-6">
        <pre className="text-sm bg-app-bg-soft p-4 rounded overflow-auto">
          {JSON.stringify(formData, null, 2)}
        </pre>
      </Card>
    </div>
  )
}
