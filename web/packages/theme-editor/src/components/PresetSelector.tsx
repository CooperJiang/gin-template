import { useEditorState, useEditorDispatch } from '../context/ThemeEditorContext'
import { presetThemes } from '@app/shared/theme'

export function PresetSelector() {
  const dispatch = useEditorDispatch()

  return (
    <div className="p-4">
      <div className="text-[12px] font-semibold text-gray-500 mb-3">预设主题</div>
      <div className="grid grid-cols-1 gap-2">
        {presetThemes.map((preset) => {
          const p500 = preset.light.primary[500].split(' ')
          const bg = `rgb(${p500.join(',')})`
          return (
            <button
              key={preset.name}
              onClick={() => dispatch({ type: 'SET_THEME', theme: JSON.parse(JSON.stringify(preset)) })}
              className="flex items-center gap-2.5 px-3 py-2 rounded-lg border border-gray-200 bg-white hover:bg-gray-50 cursor-pointer text-left transition-colors"
            >
              <span className="w-5 h-5 rounded-full shrink-0" style={{ background: bg }} />
              <span className="text-[13px] text-gray-700">{preset.name}</span>
            </button>
          )
        })}
      </div>
    </div>
  )
}
