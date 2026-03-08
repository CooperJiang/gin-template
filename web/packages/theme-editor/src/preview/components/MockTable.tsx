export function MockTable() {
  const rows = [
    { name: '张三', role: '管理员', status: '活跃' },
    { name: '李四', role: '编辑', status: '活跃' },
    { name: '王五', role: '访客', status: '禁用' },
  ]

  return (
    <div className="space-y-3">
      <div className="text-[12px] font-semibold text-gray-700">表格</div>
      <div data-theme-var="surface" className="bg-surface rounded-lg border border-gray-200 overflow-hidden">
        <table className="w-full text-[12px]" style={{ borderCollapse: 'collapse' }}>
          <thead>
            <tr data-theme-var="bg-soft" className="bg-gray-50">
              <th className="text-left px-3 py-2 font-semibold text-gray-600 border-b border-gray-200">姓名</th>
              <th className="text-left px-3 py-2 font-semibold text-gray-600 border-b border-gray-200">角色</th>
              <th className="text-left px-3 py-2 font-semibold text-gray-600 border-b border-gray-200">状态</th>
            </tr>
          </thead>
          <tbody>
            {rows.map((row, i) => (
              <tr key={row.name} className={i < rows.length - 1 ? 'border-b border-gray-100' : ''}>
                <td data-theme-var="text" className="px-3 py-2 text-gray-800">{row.name}</td>
                <td data-theme-var="text-secondary" className="px-3 py-2 text-gray-500">{row.role}</td>
                <td className="px-3 py-2">
                  <span
                    data-theme-var={row.status === '活跃' ? 'success-main' : 'error-main'}
                    className={`inline-block px-2 py-0.5 rounded-full text-[11px] font-medium ${
                      row.status === '活跃' ? 'bg-green-50 text-green-700' : 'bg-red-50 text-red-600'
                    }`}
                  >
                    {row.status}
                  </span>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  )
}
