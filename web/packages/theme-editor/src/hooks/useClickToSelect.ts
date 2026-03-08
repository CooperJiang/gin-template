import { useCallback } from 'react'
import { useEditorDispatch } from '../context/ThemeEditorContext'

export function useClickToSelect() {
  const dispatch = useEditorDispatch()

  const handlePreviewClick = useCallback((e: React.MouseEvent) => {
    const target = e.target as HTMLElement
    const el = target.closest('[data-theme-var]') as HTMLElement | null
    if (el) {
      e.stopPropagation()
      const varKey = el.getAttribute('data-theme-var')
      dispatch({ type: 'SELECT_VAR', key: varKey })
    }
  }, [dispatch])

  return { handlePreviewClick }
}
