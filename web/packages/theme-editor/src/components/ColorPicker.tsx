import { useState, useEffect } from 'react'
import { useEditorState, useEditorDispatch } from '../context/ThemeEditorContext'
import { ALL_VARIABLES, getColorByPath, rgbTripletToHex, hexToRgbTriplet } from '@app/shared/theme'

export function ColorPicker() {
  const state = useEditorState()
  const dispatch = useEditorDispatch()
  const [hexInput, setHexInput] = useState('')

  const varItem = ALL_VARIABLES.find((v) => v.key === state.selectedVar)
  const colors = state.mode === 'light' ? state.theme.light : state.theme.dark
  const currentValue = varItem ? getColorByPath(colors, varItem.path) : ''
  const currentHex = currentValue ? rgbTripletToHex(currentValue) : '#000000'

  useEffect(() => {
    setHexInput(currentHex)
  }, [currentHex])

  if (!varItem) {
    return (
      <div className="p-4 text-center text-gray-400 text-[13px]">
        点击左侧变量或预览区域中的元素来选择颜色
      </div>
    )
  }

  const updateColor = (hex: string) => {
    const triplet = hexToRgbTriplet(hex)
    dispatch({ type: 'SET_COLOR', path: varItem.path, value: triplet })
  }

  const handleHexChange = (val: string) => {
    setHexInput(val)
    if (/^#[0-9a-fA-F]{6}$/.test(val)) {
      updateColor(val)
    }
  }

  const [r, g, b] = currentValue.split(' ').map(Number)

  const updateChannel = (channel: 'r' | 'g' | 'b', val: number) => {
    const rgb = { r, g, b }
    rgb[channel] = val
    const hex = rgbTripletToHex(`${rgb.r} ${rgb.g} ${rgb.b}`)
    setHexInput(hex)
    updateColor(hex)
  }

  return (
    <div className="p-4 space-y-4">
      <div>
        <div className="text-[12px] font-semibold text-gray-500 mb-2">{varItem.label}</div>
        <div className="text-[11px] text-gray-400 font-mono">--theme-{varItem.key}</div>
      </div>

      {/* Color preview + native picker */}
      <div className="flex items-center gap-3">
        <label className="relative w-12 h-12 rounded-lg border border-gray-200 overflow-hidden cursor-pointer shrink-0">
          <input
            type="color"
            value={currentHex}
            onChange={(e) => { setHexInput(e.target.value); updateColor(e.target.value) }}
            className="absolute inset-0 w-full h-full opacity-0 cursor-pointer"
          />
          <div className="w-full h-full" style={{ background: currentHex }} />
        </label>
        <input
          type="text"
          value={hexInput}
          onChange={(e) => handleHexChange(e.target.value)}
          className="flex-1 px-2.5 py-1.5 text-[13px] font-mono border border-gray-200 rounded-md bg-gray-50 focus:outline-none focus:border-primary-400 focus:ring-1 focus:ring-primary-400"
          placeholder="#000000"
        />
      </div>

      {/* RGB sliders */}
      <div className="space-y-2">
        {([['R', r, 'r'], ['G', g, 'g'], ['B', b, 'b']] as const).map(([label, val, ch]) => (
          <div key={label} className="flex items-center gap-2">
            <span className="text-[11px] font-mono text-gray-400 w-3">{label}</span>
            <input
              type="range"
              min={0}
              max={255}
              value={val}
              onChange={(e) => updateChannel(ch as 'r' | 'g' | 'b', Number(e.target.value))}
              className="flex-1 h-1.5 accent-primary-500"
            />
            <span className="text-[11px] font-mono text-gray-500 w-7 text-right">{val}</span>
          </div>
        ))}
      </div>

      {/* RGB triplet display */}
      <div className="text-[11px] font-mono text-gray-400 bg-gray-50 rounded px-2.5 py-1.5">
        rgb({r}, {g}, {b})
      </div>
    </div>
  )
}
