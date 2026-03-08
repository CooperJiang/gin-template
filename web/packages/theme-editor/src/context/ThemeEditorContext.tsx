import { createContext, useContext, useReducer, type ReactNode, type Dispatch } from 'react'
import type { ThemeDefinition, ThemeVersion } from '@app/shared/theme'
import { defaultTheme } from '@app/shared/theme'

export interface EditorState {
  theme: ThemeDefinition
  mode: 'light' | 'dark'
  selectedVar: string | null   // CSS var key, e.g. "primary-500"
  versions: ThemeVersion[]
  dirty: boolean
}

export type EditorAction =
  | { type: 'SET_THEME'; theme: ThemeDefinition }
  | { type: 'SET_MODE'; mode: 'light' | 'dark' }
  | { type: 'SELECT_VAR'; key: string | null }
  | { type: 'SET_COLOR'; path: string; value: string }
  | { type: 'SET_VERSIONS'; versions: ThemeVersion[] }
  | { type: 'SET_DIRTY'; dirty: boolean }
  | { type: 'SET_THEME_NAME'; name: string }

function reducer(state: EditorState, action: EditorAction): EditorState {
  switch (action.type) {
    case 'SET_THEME':
      return { ...state, theme: action.theme, dirty: true }
    case 'SET_MODE':
      return { ...state, mode: action.mode }
    case 'SELECT_VAR':
      return { ...state, selectedVar: action.key }
    case 'SET_COLOR': {
      const { path, value } = action
      const modeKey = state.mode
      const clone: ThemeDefinition = JSON.parse(JSON.stringify(state.theme))
      const colors = clone[modeKey] as any
      const parts = path.split('.')
      let obj = colors
      for (let i = 0; i < parts.length - 1; i++) obj = obj[parts[i]]
      obj[parts[parts.length - 1]] = value
      return { ...state, theme: clone, dirty: true }
    }
    case 'SET_VERSIONS':
      return { ...state, versions: action.versions }
    case 'SET_DIRTY':
      return { ...state, dirty: action.dirty }
    case 'SET_THEME_NAME': {
      const clone = { ...state.theme, name: action.name }
      return { ...state, theme: clone, dirty: true }
    }
    default:
      return state
  }
}

const initialState: EditorState = {
  theme: JSON.parse(JSON.stringify(defaultTheme)),
  mode: 'light',
  selectedVar: null,
  versions: [],
  dirty: false,
}

const EditorContext = createContext<EditorState>(initialState)
const DispatchContext = createContext<Dispatch<EditorAction>>(() => {})

export function EditorProvider({ children }: { children: ReactNode }) {
  const [state, dispatch] = useReducer(reducer, initialState)
  return (
    <EditorContext.Provider value={state}>
      <DispatchContext.Provider value={dispatch}>
        {children}
      </DispatchContext.Provider>
    </EditorContext.Provider>
  )
}

export function useEditorState() {
  return useContext(EditorContext)
}

export function useEditorDispatch() {
  return useContext(DispatchContext)
}
