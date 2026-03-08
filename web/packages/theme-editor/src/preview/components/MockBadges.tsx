export function MockBadges() {
  return (
    <div className="space-y-3">
      <div className="text-[12px] font-semibold text-gray-700">标签</div>
      <div className="flex flex-wrap gap-2">
        <span data-theme-var="primary-100" className="px-2.5 py-0.5 text-[11px] font-medium rounded-full bg-primary-100 text-primary-700">
          React
        </span>
        <span data-theme-var="success-light" className="px-2.5 py-0.5 text-[11px] font-medium rounded-full bg-green-100 text-green-700">
          已发布
        </span>
        <span data-theme-var="warning-light" className="px-2.5 py-0.5 text-[11px] font-medium rounded-full bg-yellow-100 text-yellow-700">
          审核中
        </span>
        <span data-theme-var="error-light" className="px-2.5 py-0.5 text-[11px] font-medium rounded-full bg-red-100 text-red-700">
          已拒绝
        </span>
        <span data-theme-var="gray-200" className="px-2.5 py-0.5 text-[11px] font-medium rounded-full bg-gray-200 text-gray-600">
          草稿
        </span>
      </div>
    </div>
  )
}
