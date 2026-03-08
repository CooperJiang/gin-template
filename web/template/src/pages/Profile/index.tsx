import { useEffect } from 'react'
import { Link } from 'react-router-dom'
import { useAuth } from '@app/auth'
import { formatDate } from '@app/shared'

export default function Profile() {
  const { user, loading, getUserInfo } = useAuth()

  useEffect(() => {
    if (!user) {
      getUserInfo().catch(() => {})
    }
  }, [])

  return (
    <div className="min-h-screen bg-gray-50 py-12">
      <div className="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="bg-white shadow rounded-lg">
          <div className="px-6 py-4 border-b border-gray-200">
            <h1 className="text-2xl font-bold text-gray-900">个人资料</h1>
          </div>
          <div className="p-6">
            {loading ? (
              <div className="flex justify-center items-center py-8">
                <div className="animate-spin h-8 w-8 border-4 border-blue-600 border-t-transparent rounded-full" />
                <span className="ml-2 text-gray-600">加载中...</span>
              </div>
            ) : user ? (
              <>
                <div className="space-y-6">
                  <div>
                    <label className="block text-sm font-medium text-gray-700">用户名</label>
                    <input type="text" value={user.username} readOnly className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-gray-50 text-gray-500 sm:text-sm" />
                  </div>
                  <div>
                    <label className="block text-sm font-medium text-gray-700">邮箱</label>
                    <input type="email" value={user.email} readOnly className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-gray-50 text-gray-500 sm:text-sm" />
                  </div>
                  <div>
                    <label className="block text-sm font-medium text-gray-700">用户状态</label>
                    <div className="mt-1">
                      <span className={`inline-flex items-center px-3 py-1 rounded-full text-xs font-medium ${user.status === 1 ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'}`}>
                        {user.status === 1 ? '正常' : '禁用'}
                      </span>
                    </div>
                  </div>
                  {user.bio && (
                    <div>
                      <label className="block text-sm font-medium text-gray-700">个人简介</label>
                      <textarea value={user.bio} readOnly rows={3} className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-gray-50 text-gray-500 sm:text-sm" />
                    </div>
                  )}
                  {user.created_at && (
                    <div>
                      <label className="block text-sm font-medium text-gray-700">注册时间</label>
                      <input type="text" value={formatDate(user.created_at)} readOnly className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm bg-gray-50 text-gray-500 sm:text-sm" />
                    </div>
                  )}
                </div>
                <div className="mt-8 flex space-x-4">
                  <button type="button" className="px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700">编辑资料</button>
                  <Link to="/settings" className="px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50">账户设置</Link>
                  <Link to="/" className="px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50">返回首页</Link>
                </div>
              </>
            ) : (
              <div className="text-center py-8">
                <p className="text-gray-600">无法加载用户信息</p>
                <button onClick={() => getUserInfo().catch(() => {})} className="mt-4 px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700">重新加载</button>
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  )
}
