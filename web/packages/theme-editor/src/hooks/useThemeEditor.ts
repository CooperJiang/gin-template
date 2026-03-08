import { useEffect } from 'react'
import { useEditorState, useEditorDispatch } from '../context/ThemeEditorContext'
import { applyTheme } from '@app/shared/theme'

const STORAGE_KEY = 'theme-editor-data'

interface StoredData {
  theme: import('@app/shared/theme').ThemeDefinition
  mode: 'light' | 'dark'
}

export function useThemeEditor() {
  const state = useEditorState()
  const dispatch = useEditorDispatch()

  // Load from localStorage on mount
  useEffect(() => {
    try {
      const raw = localStorage.getItem(STORAGE_KEY)
      if (raw) {
        const data: StoredData = JSON.parse(raw)
        if (data.theme) {
          dispatch({ type: 'SET_THEME', key: '' } as any)
          dispatch({ type: 'SET_THEME', theme: data.theme })
          dispatch({ type: 'SET_DIRTY', dirty: false })
        }
        if (data.mode) dispatch({ type: 'SET_MODE', mode: data.mode })
      }
    } catch { /* ignore */ }
  }, [])

  // Apply theme to DOM whenever theme or mode changes
  useEffect(() => {
    applyTheme(state.theme, state.mode)
  }, [state.theme, state.mode])

  const save = () => {
    const data: StoredData = { theme: state.theme, mode: state.mode }
    localStorage.setItem(STORAGE_KEY, JSON.stringify(data))
  }

  return { save }
}
