import { useEditorState, useEditorDispatch } from '../context/ThemeEditorContext'
import { rgbTripletToHex, getColorByPath, VARIABLE_CATEGORIES, ALL_VARIABLES } from '@app/shared/theme'

export function LeftPanel() {
  const state = useEditorState()
  const dispatch = useEditorDispatch()
  const colors = state.mode === 'light' ? state.theme.light : state.theme.dark

  return (
    <div className="w-64 shrink-0 border-r border-gray-200 bg-white overflow-y-auto h-full">
      <div className="p-3 border-b border-gray-100">
        <div className="text-xs font-semibold text-gray-400 uppercase tracking-wider">变量列表</div>
      </div>
      <div className="p-2 space-y-1">
        {VARIABLE_CATEGORIES.map((cat) => (
          <CategoryGroup key={cat.id} category={cat} colors={colors} selectedVar={state.selectedVar} dispatch={dispatch} />
        ))}
      </div>
    </div>
  )
}

function CategoryGroup({
  category,
  colors,
  selectedVar,
  dispatch,
}: {
  category: typeof VARIABLE_CATEGORIES[number]
  colors: any
  selectedVar: string | null
  dispatch: ReturnType<typeof useEditorDispatch>
}) {
  return (
    <div className="mb-2">
      <div className="text-[11px] font-semibold text-gray-400 uppercase tracking-wider px-2 py-1.5">
        {category.label}
      </div>
      {category.items.map((item) => {
        const value = getColorByPath(colors, item.path)
        const hex = value ? rgbTripletToHex(value) : '#000000'
        const isSelected = selectedVar === item.key
        return (
          <button
            key={item.key}
            onClick={() => dispatch({ type: 'SELECT_VAR', key: item.key })}
            className={`w-full flex items-center gap-2 px-2 py-1.5 rounded text-left text-[12px] border-none cursor-pointer transition-colors ${
              isSelected ? 'bg-primary-50 text-primary-700' : 'bg-transparent text-gray-600 hover:bg-gray-50'
            }`}
          >
            <span
              className="w-4 h-4 rounded border border-gray-200 shrink-0"
              style={{ background: hex }}
            />
            <span className="truncate">{item.label}</span>
          </button>
        )
      })}
    </div>
  )
}
