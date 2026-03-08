import { useRef } from 'react'
import { useEditorState, useEditorDispatch } from '../context/ThemeEditorContext'
import { useVersionManager } from '../hooks/useVersionManager'
import { useThemeEditor } from '../hooks/useThemeEditor'
import { exportTheme, importTheme } from '../utils/export'

export function TopBar() {
  const state = useEditorState()
  const dispatch = useEditorDispatch()
  const { saveVersion } = useVersionManager()
  const { save } = useThemeEditor()
  const fileRef = useRef<HTMLInputElement>(null)

  const handleSave = () => {
    save()
    const saved = saveVersion()
    if (!saved) {
      alert('主题未发生变化，无需保存')
    }
  }

  const handleImport = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0]
    if (!file) return
    try {
      const theme = await importTheme(file)
      dispatch({ type: 'SET_THEME', theme })
    } catch (err: any) {
      alert(err.message)
    }
    e.target.value = ''
  }

  return (
    <div className="h-12 border-b border-gray-200 bg-white flex items-center justify-between px-4 shrink-0">
      {/* Left: logo + theme name */}
      <div className="flex items-center gap-3">
        <div
          className="w-7 h-7 rounded-md flex items-center justify-center"
          style={{ background: 'linear-gradient(135deg, #3b82f6, #6366f1)' }}
        >
          <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="white" strokeWidth="2.2" strokeLinecap="round" strokeLinejoin="round">
            <circle cx="12" cy="12" r="10" /><path d="M12 2a15 15 0 010 20M2 12h20" />
          </svg>
        </div>
        <input
          type="text"
          value={state.theme.name}
          onChange={(e) => dispatch({ type: 'SET_THEME_NAME', name: e.target.value })}
          className="text-[14px] font-semibold text-gray-800 border-none bg-transparent focus:outline-none focus:bg-gray-50 rounded px-1.5 py-0.5 -ml-1.5"
          style={{ width: Math.max(100, state.theme.name.length * 10) }}
        />
        {state.dirty && (
          <span className="w-2 h-2 rounded-full bg-amber-400" title="未保存" />
        )}
      </div>

      {/* Center: mode toggle */}
      <div className="flex items-center bg-gray-100 rounded-lg p-0.5">
        <button
          onClick={() => dispatch({ type: 'SET_MODE', mode: 'light' })}
          className={`px-3 py-1 text-[12px] font-medium rounded-md border-none cursor-pointer transition-colors ${
            state.mode === 'light' ? 'bg-white text-gray-800 shadow-sm' : 'bg-transparent text-gray-500'
          }`}
        >
          <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="inline -mt-0.5 mr-1">
            <circle cx="12" cy="12" r="5" /><line x1="12" y1="1" x2="12" y2="3" /><line x1="12" y1="21" x2="12" y2="23" /><line x1="4.22" y1="4.22" x2="5.64" y2="5.64" /><line x1="18.36" y1="18.36" x2="19.78" y2="19.78" /><line x1="1" y1="12" x2="3" y2="12" /><line x1="21" y1="12" x2="23" y2="12" /><line x1="4.22" y1="19.78" x2="5.64" y2="18.36" /><line x1="18.36" y1="5.64" x2="19.78" y2="4.22" />
          </svg>
          Light
        </button>
        <button
          onClick={() => dispatch({ type: 'SET_MODE', mode: 'dark' })}
          className={`px-3 py-1 text-[12px] font-medium rounded-md border-none cursor-pointer transition-colors ${
            state.mode === 'dark' ? 'bg-white text-gray-800 shadow-sm' : 'bg-transparent text-gray-500'
          }`}
        >
          <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="inline -mt-0.5 mr-1">
            <path d="M21 12.79A9 9 0 1111.21 3 7 7 0 0021 12.79z" />
          </svg>
          Dark
        </button>
      </div>

      {/* Right: actions */}
      <div className="flex items-center gap-1.5">
        <button
          onClick={handleSave}
          className="px-3 py-1.5 text-[12px] font-medium text-white rounded-md border-none cursor-pointer"
          style={{ background: 'linear-gradient(135deg, #3b82f6, #6366f1)' }}
        >
          保存版本
        </button>
        <button
          onClick={() => exportTheme(state.theme)}
          className="px-3 py-1.5 text-[12px] font-medium text-gray-600 bg-gray-100 hover:bg-gray-200 rounded-md border-none cursor-pointer transition-colors"
        >
          导出
        </button>
        <button
          onClick={() => fileRef.current?.click()}
          className="px-3 py-1.5 text-[12px] font-medium text-gray-600 bg-gray-100 hover:bg-gray-200 rounded-md border-none cursor-pointer transition-colors"
        >
          导入
        </button>
        <input ref={fileRef} type="file" accept=".json" onChange={handleImport} className="hidden" />
      </div>
    </div>
  )
}
