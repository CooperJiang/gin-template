export function MockButtons() {
  return (
    <div className="space-y-3">
      <div className="text-[12px] font-semibold text-gray-700">按钮</div>
      <div className="flex flex-wrap gap-2">
        <button data-theme-var="primary-600" className="px-3.5 py-1.5 text-[12px] font-medium text-white bg-primary-600 rounded-md border-none cursor-default">
          主要按钮
        </button>
        <button data-theme-var="primary-100" className="px-3.5 py-1.5 text-[12px] font-medium text-primary-700 bg-primary-100 rounded-md border-none cursor-default">
          次要按钮
        </button>
        <button data-theme-var="gray-200" className="px-3.5 py-1.5 text-[12px] font-medium text-gray-700 bg-gray-200 rounded-md border-none cursor-default">
          默认按钮
        </button>
        <button data-theme-var="error-main" className="px-3.5 py-1.5 text-[12px] font-medium text-white bg-error rounded-md border-none cursor-default">
          危险按钮
        </button>
        <button data-theme-var="border" className="px-3.5 py-1.5 text-[12px] font-medium text-gray-700 bg-white rounded-md border border-gray-200 cursor-default">
          边框按钮
        </button>
      </div>
    </div>
  )
}
