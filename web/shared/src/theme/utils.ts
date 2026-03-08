import type { ThemeModeColors, ThemeDefinition } from './types'
import { CSS_VAR_PREFIX, ALL_VARIABLES } from './constants'

/** "#3b82f6" → "59 130 246" */
export function hexToRgbTriplet(hex: string): string {
  const h = hex.replace('#', '')
  const n = parseInt(h, 16)
  return `${(n >> 16) & 255} ${(n >> 8) & 255} ${n & 255}`
}

/** "59 130 246" → "#3b82f6" */
export function rgbTripletToHex(triplet: string): string {
  const [r, g, b] = triplet.split(' ').map(Number)
  return '#' + [r, g, b].map((v) => v.toString(16).padStart(2, '0')).join('')
}

/** Get a color value from ThemeModeColors by dot path, e.g. "primary.500" */
export function getColorByPath(colors: ThemeModeColors, path: string): string {
  const parts = path.split('.')
  let obj: any = colors
  for (const p of parts) {
    obj = obj?.[p]
  }
  return typeof obj === 'string' ? obj : ''
}

/** Set a color value in ThemeModeColors by dot path */
export function setColorByPath(colors: ThemeModeColors, path: string, value: string): ThemeModeColors {
  const clone: any = JSON.parse(JSON.stringify(colors))
  const parts = path.split('.')
  let obj = clone
  for (let i = 0; i < parts.length - 1; i++) {
    obj = obj[parts[i]]
  }
  obj[parts[parts.length - 1]] = value
  return clone
}

/** Apply theme CSS variables to a target element */
export function applyTheme(
  theme: ThemeDefinition,
  mode: 'light' | 'dark',
  target: HTMLElement = document.documentElement,
) {
  const colors = mode === 'light' ? theme.light : theme.dark

  for (const v of ALL_VARIABLES) {
    const value = getColorByPath(colors, v.path)
    if (value) {
      target.style.setProperty(`${CSS_VAR_PREFIX}${v.key}`, value)
    }
  }

  // Apply tokens
  const t = theme.tokens
  target.style.setProperty('--theme-radius-sm', t.radiusSm)
  target.style.setProperty('--theme-radius-base', t.radiusBase)
  target.style.setProperty('--theme-radius-lg', t.radiusLg)
  target.style.setProperty('--theme-radius-xl', t.radiusXl)
  target.style.setProperty('--theme-radius-full', t.radiusFull)
  target.style.setProperty('--theme-shadow-sm', t.shadowSm)
  target.style.setProperty('--theme-shadow-base', t.shadowBase)
  target.style.setProperty('--theme-shadow-lg', t.shadowLg)
  target.style.setProperty('--theme-font-sans', t.fontSans)
  target.style.setProperty('--theme-font-mono', t.fontMono)

  // Toggle dark class
  if (mode === 'dark') {
    target.classList.add('dark')
  } else {
    target.classList.remove('dark')
  }
}

/** Generate full CSS string from a theme definition */
export function generateThemeCSS(theme: ThemeDefinition): string {
  const lines = (colors: ThemeModeColors, indent: string) => {
    return ALL_VARIABLES
      .map((v) => {
        const val = getColorByPath(colors, v.path)
        return val ? `${indent}${CSS_VAR_PREFIX}${v.key}: ${val};` : ''
      })
      .filter(Boolean)
  }

  const tokenLines = (indent: string) => {
    const t = theme.tokens
    return [
      `${indent}--theme-radius-sm: ${t.radiusSm};`,
      `${indent}--theme-radius-base: ${t.radiusBase};`,
      `${indent}--theme-radius-lg: ${t.radiusLg};`,
      `${indent}--theme-radius-xl: ${t.radiusXl};`,
      `${indent}--theme-radius-full: ${t.radiusFull};`,
      `${indent}--theme-shadow-sm: ${t.shadowSm};`,
      `${indent}--theme-shadow-base: ${t.shadowBase};`,
      `${indent}--theme-shadow-lg: ${t.shadowLg};`,
      `${indent}--theme-font-sans: ${t.fontSans};`,
      `${indent}--theme-font-mono: ${t.fontMono};`,
    ]
  }

  return [
    ':root {',
    ...lines(theme.light, '  '),
    ...tokenLines('  '),
    '}',
    '',
    '.dark {',
    ...lines(theme.dark, '  '),
    '}',
  ].join('\n')
}

/** Deep equality check for two theme definitions */
export function themesEqual(a: ThemeDefinition, b: ThemeDefinition): boolean {
  return JSON.stringify(a) === JSON.stringify(b)
}
