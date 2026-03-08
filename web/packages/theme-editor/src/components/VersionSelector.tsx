import { useEditorState } from '../context/ThemeEditorContext'
import { useVersionManager } from '../hooks/useVersionManager'

export function VersionSelector() {
  const state = useEditorState()
  const { loadVersion, deleteVersion } = useVersionManager()

  if (state.versions.length === 0) {
    return (
      <div className="p-4 text-center text-gray-400 text-[13px]">
        暂无保存的版本
      </div>
    )
  }

  return (
    <div className="p-4">
      <div className="text-[12px] font-semibold text-gray-500 mb-3">历史版本</div>
      <div className="space-y-1.5">
        {state.versions.map((v) => (
          <div
            key={v.id}
            className="flex items-center justify-between px-3 py-2 rounded-lg border border-gray-200 bg-white"
          >
            <button
              onClick={() => loadVersion(v.id)}
              className="flex-1 text-left border-none bg-transparent cursor-pointer p-0"
            >
              <div className="text-[13px] text-gray-700">{v.name}</div>
              <div className="text-[11px] text-gray-400">
                {new Date(v.timestamp).toLocaleString('zh-CN')}
              </div>
            </button>
            <button
              onClick={() => deleteVersion(v.id)}
              className="text-gray-300 hover:text-red-500 border-none bg-transparent cursor-pointer p-1 transition-colors"
            >
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                <line x1="18" y1="6" x2="6" y2="18" /><line x1="6" y1="6" x2="18" y2="18" />
              </svg>
            </button>
          </div>
        ))}
      </div>
    </div>
  )
}
