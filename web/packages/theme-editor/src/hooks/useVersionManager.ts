import { useEffect } from 'react'
import { useEditorState, useEditorDispatch } from '../context/ThemeEditorContext'
import type { ThemeVersion } from '@app/shared/theme'
import { themesEqual } from '@app/shared/theme'

const VERSIONS_KEY = 'theme-editor-versions'
const MAX_VERSIONS = 10

function genId() {
  return Date.now().toString(36) + Math.random().toString(36).slice(2, 8)
}

export function useVersionManager() {
  const state = useEditorState()
  const dispatch = useEditorDispatch()

  // Load versions on mount
  useEffect(() => {
    try {
      const raw = localStorage.getItem(VERSIONS_KEY)
      if (raw) {
        dispatch({ type: 'SET_VERSIONS', versions: JSON.parse(raw) })
      }
    } catch { /* ignore */ }
  }, [])

  const saveVersion = (name?: string) => {
    const versions = [...state.versions]
    // Check if theme changed from latest version
    if (versions.length > 0 && themesEqual(versions[0].theme, state.theme)) {
      return false // no changes
    }

    const version: ThemeVersion = {
      id: genId(),
      name: name || `版本 ${versions.length + 1}`,
      timestamp: Date.now(),
      theme: JSON.parse(JSON.stringify(state.theme)),
    }

    versions.unshift(version)
    if (versions.length > MAX_VERSIONS) versions.pop()

    localStorage.setItem(VERSIONS_KEY, JSON.stringify(versions))
    dispatch({ type: 'SET_VERSIONS', versions })
    dispatch({ type: 'SET_DIRTY', dirty: false })
    return true
  }

  const loadVersion = (id: string) => {
    const version = state.versions.find((v) => v.id === id)
    if (version) {
      dispatch({ type: 'SET_THEME', theme: JSON.parse(JSON.stringify(version.theme)) })
      dispatch({ type: 'SET_DIRTY', dirty: false })
    }
  }

  const deleteVersion = (id: string) => {
    const versions = state.versions.filter((v) => v.id !== id)
    localStorage.setItem(VERSIONS_KEY, JSON.stringify(versions))
    dispatch({ type: 'SET_VERSIONS', versions })
  }

  return { saveVersion, loadVersion, deleteVersion }
}
