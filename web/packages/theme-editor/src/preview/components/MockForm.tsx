export function MockForm() {
  return (
    <div className="space-y-3">
      <div className="text-[12px] font-semibold text-gray-700">表单</div>
      <div data-theme-var="surface" className="bg-surface rounded-lg p-4 border border-gray-200 space-y-3">
        <div>
          <label data-theme-var="text" className="block text-[12px] font-medium text-gray-700 mb-1">用户名</label>
          <input
            data-theme-var="input-border"
            type="text"
            placeholder="请输入用户名"
            readOnly
            className="w-full px-3 py-1.5 text-[12px] border border-gray-300 rounded-md bg-white text-gray-900 placeholder-gray-400 focus:outline-none"
          />
        </div>
        <div>
          <label className="block text-[12px] font-medium text-gray-700 mb-1">邮箱</label>
          <input
            data-theme-var="input-bg"
            type="email"
            placeholder="请输入邮箱"
            readOnly
            className="w-full px-3 py-1.5 text-[12px] border border-gray-300 rounded-md bg-white text-gray-900 placeholder-gray-400 focus:outline-none"
          />
        </div>
        <div className="flex items-center gap-2">
          <div data-theme-var="primary-500" className="w-4 h-4 rounded bg-primary-500 flex items-center justify-center">
            <svg width="10" height="10" viewBox="0 0 24 24" fill="none" stroke="white" strokeWidth="3" strokeLinecap="round" strokeLinejoin="round">
              <polyline points="20 6 9 17 4 12" />
            </svg>
          </div>
          <span className="text-[12px] text-gray-600">记住我</span>
        </div>
      </div>
    </div>
  )
}
