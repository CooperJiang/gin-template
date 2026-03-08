export function MockAlerts() {
  const alerts = [
    { type: 'success', label: '成功', msg: '操作已成功完成', varKey: 'success-main', bg: 'bg-green-50', text: 'text-green-800', border: 'border-green-200' },
    { type: 'warning', label: '警告', msg: '请注意检查配置', varKey: 'warning-main', bg: 'bg-yellow-50', text: 'text-yellow-800', border: 'border-yellow-200' },
    { type: 'error', label: '错误', msg: '请求处理失败', varKey: 'error-main', bg: 'bg-red-50', text: 'text-red-800', border: 'border-red-200' },
    { type: 'info', label: '信息', msg: '系统将于今晚维护', varKey: 'info-main', bg: 'bg-blue-50', text: 'text-blue-800', border: 'border-blue-200' },
  ]

  return (
    <div className="space-y-3">
      <div className="text-[12px] font-semibold text-gray-700">提示</div>
      <div className="space-y-2">
        {alerts.map((a) => (
          <div
            key={a.type}
            data-theme-var={a.varKey}
            className={`px-3 py-2 rounded-md text-[12px] border ${a.bg} ${a.text} ${a.border}`}
          >
            <span className="font-semibold">{a.label}：</span>{a.msg}
          </div>
        ))}
      </div>
    </div>
  )
}
