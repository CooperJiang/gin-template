/** RGB triplet string, e.g. "59 130 246" */
export type RGBTriplet = string

/** 50-950 color scale */
export interface ColorScale {
  50: RGBTriplet
  100: RGBTriplet
  200: RGBTriplet
  300: RGBTriplet
  400: RGBTriplet
  500: RGBTriplet
  600: RGBTriplet
  700: RGBTriplet
  800: RGBTriplet
  900: RGBTriplet
  950: RGBTriplet
}

export interface StatusColor {
  main: RGBTriplet
  light: RGBTriplet
}

/** All color variables for one mode (light or dark) */
export interface ThemeModeColors {
  primary: ColorScale
  gray: ColorScale

  success: StatusColor
  warning: StatusColor
  error: StatusColor
  info: StatusColor

  bg: RGBTriplet
  bgSoft: RGBTriplet
  bgMute: RGBTriplet

  surface: RGBTriplet
  surfaceHover: RGBTriplet
  surfaceActive: RGBTriplet

  text: RGBTriplet
  textSecondary: RGBTriplet
  textMuted: RGBTriplet
  textInverse: RGBTriplet
  heading: RGBTriplet

  border: RGBTriplet
  borderLight: RGBTriplet
  borderHover: RGBTriplet

  ring: RGBTriplet

  inputBg: RGBTriplet
  inputBorder: RGBTriplet
  inputText: RGBTriplet
  inputPlaceholder: RGBTriplet

  link: RGBTriplet
  linkHover: RGBTriplet

  overlay: RGBTriplet

  scrollbarTrack: RGBTriplet
  scrollbarThumb: RGBTriplet
  scrollbarThumbHover: RGBTriplet
}

export interface ThemeTokens {
  radiusSm: string
  radiusBase: string
  radiusLg: string
  radiusXl: string
  radiusFull: string
  shadowSm: string
  shadowBase: string
  shadowLg: string
  fontSans: string
  fontMono: string
}

export interface ThemeDefinition {
  name: string
  light: ThemeModeColors
  dark: ThemeModeColors
  tokens: ThemeTokens
}

export interface ThemeVersion {
  id: string
  name: string
  timestamp: number
  theme: ThemeDefinition
}

export interface ThemeExport {
  version: string
  exportedAt: string
  theme: ThemeDefinition
}

export interface VariableItem {
  key: string        // CSS variable suffix, e.g. "primary-500"
  path: string       // dot path in ThemeModeColors, e.g. "primary.500"
  label: string
}

export interface VariableCategory {
  id: string
  label: string
  items: VariableItem[]
}
