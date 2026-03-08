export function MockCards() {
  const cards = [
    { title: '数据概览', value: '12,345', change: '+12.5%', up: true },
    { title: '活跃用户', value: '8,901', change: '+5.2%', up: true },
    { title: '错误率', value: '0.3%', change: '-2.1%', up: false },
  ]

  return (
    <div className="space-y-3">
      <div className="text-[12px] font-semibold text-gray-700">卡片</div>
      <div className="grid grid-cols-3 gap-3">
        {cards.map((card) => (
          <div
            key={card.title}
            data-theme-var="surface"
            className="bg-surface rounded-lg p-4 border border-gray-200"
          >
            <div data-theme-var="text-muted" className="text-[11px] text-gray-400 mb-1">{card.title}</div>
            <div data-theme-var="heading" className="text-[18px] font-bold text-gray-900">{card.value}</div>
            <div className={`text-[11px] mt-1 ${card.up ? 'text-green-600' : 'text-red-500'}`}>
              {card.change}
            </div>
          </div>
        ))}
      </div>
    </div>
  )
}
