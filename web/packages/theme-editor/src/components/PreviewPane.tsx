import { useCallback, useRef, useEffect } from 'react'
import { useEditorState, useEditorDispatch } from '../context/ThemeEditorContext'
import { ALL_VARIABLES } from '@app/shared/theme'
import { MockPage } from '../preview/MockPage'

export function PreviewPane() {
  const state = useEditorState()
  const dispatch = useEditorDispatch()
  const containerRef = useRef<HTMLDivElement>(null)

  const handleClick = useCallback((e: React.MouseEvent) => {
    const target = e.target as HTMLElement
    const el = target.closest('[data-theme-var]') as HTMLElement | null
    if (el) {
      e.stopPropagation()
      const varKey = el.getAttribute('data-theme-var')
      if (varKey && ALL_VARIABLES.some((v) => v.key === varKey)) {
        dispatch({ type: 'SELECT_VAR', key: varKey })
      }
    }
  }, [dispatch])

  // Highlight selected element
  useEffect(() => {
    const container = containerRef.current
    if (!container) return

    // Remove previous highlights
    container.querySelectorAll('.theme-editor-highlight').forEach((el) => {
      el.classList.remove('theme-editor-highlight')
      el.removeAttribute('data-theme-label')
    })

    if (state.selectedVar) {
      const els = container.querySelectorAll(`[data-theme-var="${state.selectedVar}"]`)
      const item = ALL_VARIABLES.find((v) => v.key === state.selectedVar)
      els.forEach((el) => {
        el.classList.add('theme-editor-highlight')
        if (item) el.setAttribute('data-theme-label', item.label)
      })
    }
  }, [state.selectedVar])

  return (
    <div
      ref={containerRef}
      onClick={handleClick}
      className="flex-1 overflow-auto bg-gray-100"
    >
      <MockPage />
    </div>
  )
}
