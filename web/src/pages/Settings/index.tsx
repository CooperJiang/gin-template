import { Card, CardContent, CardHeader, CardTitle, CardDescription } from '@/components/ui/card'

export default function Settings() {
  return (
    <div>
      <div className="mb-8">
        <h1 className="text-2xl font-bold text-gray-900">系统设置</h1>
        <p className="text-gray-500 mt-1">管理应用程序设置和偏好</p>
      </div>

      <div className="max-w-2xl space-y-6">
        <Card>
          <CardHeader>
            <CardTitle>通用设置</CardTitle>
            <CardDescription>管理应用程序的基本配置</CardDescription>
          </CardHeader>
          <CardContent>
            <p className="text-sm text-gray-500">设置功能正在开发中，敬请期待。</p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>安全设置</CardTitle>
            <CardDescription>管理密码和安全选项</CardDescription>
          </CardHeader>
          <CardContent>
            <p className="text-sm text-gray-500">安全设置功能正在开发中，敬请期待。</p>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
