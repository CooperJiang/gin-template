import { useState } from 'react'
import { EditorProvider } from './context/ThemeEditorContext'
import { useThemeEditor } from './hooks/useThemeEditor'
import { TopBar } from './components/TopBar'
import { LeftPanel } from './components/LeftPanel'
import { ColorPicker } from './components/ColorPicker'
import { PreviewPane } from './components/PreviewPane'
import { PresetSelector } from './components/PresetSelector'
import { VersionSelector } from './components/VersionSelector'

function EditorLayout() {
  useThemeEditor()
  const [rightTab, setRightTab] = useState<'color' | 'presets' | 'versions'>('color')

  return (
    <div className="h-screen flex flex-col bg-gray-50">
      <TopBar />
      <div className="flex flex-1 overflow-hidden">
        {/* Left: variable tree */}
        <LeftPanel />

        {/* Center: preview */}
        <PreviewPane />

        {/* Right: color picker / presets / versions */}
        <div className="w-72 shrink-0 border-l border-gray-200 bg-white flex flex-col overflow-hidden">
          <div className="flex border-b border-gray-200 shrink-0">
            {([
              ['color', '颜色'],
              ['presets', '预设'],
              ['versions', '版本'],
            ] as const).map(([key, label]) => (
              <button
                key={key}
                onClick={() => setRightTab(key)}
                className={`flex-1 py-2 text-[12px] font-medium border-none cursor-pointer transition-colors ${
                  rightTab === key
                    ? 'text-primary-600 bg-primary-50 border-b-2 border-primary-500'
                    : 'text-gray-500 bg-transparent hover:text-gray-700'
                }`}
              >
                {label}
              </button>
            ))}
          </div>
          <div className="flex-1 overflow-y-auto">
            {rightTab === 'color' && <ColorPicker />}
            {rightTab === 'presets' && <PresetSelector />}
            {rightTab === 'versions' && <VersionSelector />}
          </div>
        </div>
      </div>
    </div>
  )
}

export default function App() {
  return (
    <EditorProvider>
      <EditorLayout />
    </EditorProvider>
  )
}
