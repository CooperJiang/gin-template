import { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import { Card, Button, Input, Loading } from '@app/core'
import { createHttpClient } from '@app/core'
import type { PaginationResponse } from '@app/shared/types'

interface User {
  id: number
  name: string
  email: string
  role: string
  createdAt: string
}

/**
 * 用户列表页面示例
 * 展示：表格、分页、搜索、加载状态
 */
export function UserList() {
  const navigate = useNavigate()
  const [users, setUsers] = useState<User[]>([])
  const [loading, setLoading] = useState(false)
  const [page, setPage] = useState(1)
  const [total, setTotal] = useState(0)
  const [search, setSearch] = useState('')
  const pageSize = 10

  const http = createHttpClient({ baseURL: '/api/v1' })

  useEffect(() => {
    fetchUsers()
  }, [page, search])

  const fetchUsers = async () => {
    setLoading(true)
    try {
      const response = await http.get<PaginationResponse<User>>('/users', {
        params: {
          page,
          pageSize,
          search,
        },
      })
      setUsers(response.items)
      setTotal(response.total)
    } catch (error) {
      // 错误已由全局处理器处理
    } finally {
      setLoading(false)
    }
  }

  const totalPages = Math.ceil(total / pageSize)

  return (
    <div className="max-w-6xl mx-auto p-6">
      <Card
        title="用户管理"
        extra={
          <Button onClick={() => navigate('/examples/user/new')}>
            + 添加用户
          </Button>
        }
      >
        {/* 搜索栏 */}
        <div className="mb-4">
          <Input
            placeholder="搜索用户名或邮箱..."
            value={search}
            onChange={(e) => {
              setSearch(e.target.value)
              setPage(1) // 重置到第一页
            }}
            fullWidth
          />
        </div>

        {/* 表格 */}
        {loading ? (
          <div className="py-12">
            <Loading />
          </div>
        ) : users.length === 0 ? (
          <div className="text-center py-12 text-app-text-muted">
            暂无数据
          </div>
        ) : (
          <>
            <div className="overflow-x-auto">
              <table className="w-full">
                <thead className="bg-app-bg-soft border-b border-app-border">
                  <tr>
                    <th className="px-4 py-3 text-left text-sm font-medium text-app-text">
                      用户名
                    </th>
                    <th className="px-4 py-3 text-left text-sm font-medium text-app-text">
                      邮箱
                    </th>
                    <th className="px-4 py-3 text-left text-sm font-medium text-app-text">
                      角色
                    </th>
                    <th className="px-4 py-3 text-left text-sm font-medium text-app-text">
                      创建时间
                    </th>
                    <th className="px-4 py-3 text-right text-sm font-medium text-app-text">
                      操作
                    </th>
                  </tr>
                </thead>
                <tbody className="divide-y divide-app-border">
                  {users.map((user) => (
                    <tr
                      key={user.id}
                      className="hover:bg-app-bg-soft transition-colors"
                    >
                      <td className="px-4 py-3 text-sm text-app-text">
                        {user.name}
                      </td>
                      <td className="px-4 py-3 text-sm text-app-text">
                        {user.email}
                      </td>
                      <td className="px-4 py-3 text-sm">
                        <span className="inline-flex px-2 py-1 text-xs font-medium rounded-full bg-primary-100 text-primary-700">
                          {user.role}
                        </span>
                      </td>
                      <td className="px-4 py-3 text-sm text-app-text-muted">
                        {new Date(user.createdAt).toLocaleDateString()}
                      </td>
                      <td className="px-4 py-3 text-sm text-right">
                        <Button
                          variant="ghost"
                          size="sm"
                          onClick={() => navigate(`/examples/user/${user.id}`)}
                        >
                          查看
                        </Button>
                        <Button
                          variant="ghost"
                          size="sm"
                          onClick={() =>
                            navigate(`/examples/user/${user.id}/edit`)
                          }
                        >
                          编辑
                        </Button>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>

            {/* 分页 */}
            <div className="flex items-center justify-between mt-4 pt-4 border-t border-app-border">
              <div className="text-sm text-app-text-muted">
                共 {total} 条记录，第 {page} / {totalPages} 页
              </div>
              <div className="flex gap-2">
                <Button
                  variant="outline"
                  size="sm"
                  disabled={page === 1}
                  onClick={() => setPage(page - 1)}
                >
                  上一页
                </Button>
                <Button
                  variant="outline"
                  size="sm"
                  disabled={page === totalPages}
                  onClick={() => setPage(page + 1)}
                >
                  下一页
                </Button>
              </div>
            </div>
          </>
        )}
      </Card>
    </div>
  )
}
