export function MockNav() {
  return (
    <nav
      data-theme-var="surface"
      className="bg-surface border-b border-gray-200 px-5 py-3 flex items-center justify-between"
    >
      <div className="flex items-center gap-6">
        <div className="flex items-center gap-2">
          <div
            data-theme-var="primary-500"
            className="w-7 h-7 rounded-md flex items-center justify-center bg-primary-500"
          >
            <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="white" strokeWidth="2.2" strokeLinecap="round" strokeLinejoin="round">
              <path d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
            </svg>
          </div>
          <span data-theme-var="heading" className="text-[14px] font-bold text-gray-900">应用平台</span>
        </div>
        <div className="flex items-center gap-1">
          {['首页', '文档', '关于'].map((label, i) => (
            <span
              key={label}
              data-theme-var={i === 0 ? 'primary-600' : 'text-secondary'}
              className={`px-2.5 py-1 text-[12px] font-medium rounded-md cursor-default ${
                i === 0 ? 'text-primary-600 bg-primary-50' : 'text-gray-500'
              }`}
            >
              {label}
            </span>
          ))}
        </div>
      </div>
      <div className="flex items-center gap-2">
        <span data-theme-var="text-muted" className="text-[12px] text-gray-400">用户名</span>
        <div data-theme-var="gray-300" className="w-7 h-7 rounded-full bg-gray-300" />
      </div>
    </nav>
  )
}
