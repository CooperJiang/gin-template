import type { ThemeModeColors, ThemeTokens, ThemeDefinition } from './types'

export const defaultLightColors: ThemeModeColors = {
  primary: {
    50: '239 246 255', 100: '219 234 254', 200: '191 219 254', 300: '147 197 253',
    400: '96 165 250', 500: '59 130 246', 600: '37 99 235', 700: '29 78 216',
    800: '30 64 175', 900: '30 58 138', 950: '23 37 84',
  },
  gray: {
    50: '249 250 251', 100: '243 244 246', 200: '229 231 235', 300: '209 213 219',
    400: '156 163 175', 500: '107 114 128', 600: '75 85 99', 700: '55 65 81',
    800: '31 41 55', 900: '17 24 39', 950: '3 7 18',
  },
  success: { main: '34 197 94', light: '220 252 231' },
  warning: { main: '234 179 8', light: '254 249 195' },
  error: { main: '239 68 68', light: '254 226 226' },
  info: { main: '59 130 246', light: '219 234 254' },

  bg: '249 250 251',           // gray-50
  bgSoft: '243 244 246',       // gray-100
  bgMute: '229 231 235',       // gray-200

  surface: '255 255 255',
  surfaceHover: '249 250 251',
  surfaceActive: '243 244 246',

  text: '17 24 39',            // gray-900
  textSecondary: '75 85 99',   // gray-600
  textMuted: '156 163 175',    // gray-400
  textInverse: '255 255 255',
  heading: '17 24 39',         // gray-900

  border: '229 231 235',       // gray-200
  borderLight: '243 244 246',  // gray-100
  borderHover: '209 213 219',  // gray-300

  ring: '59 130 246',          // primary-500

  inputBg: '255 255 255',
  inputBorder: '209 213 219',  // gray-300
  inputText: '17 24 39',       // gray-900
  inputPlaceholder: '156 163 175', // gray-400

  link: '37 99 235',           // primary-600
  linkHover: '29 78 216',      // primary-700

  overlay: '0 0 0',

  scrollbarTrack: '248 250 252',
  scrollbarThumb: '203 213 225',
  scrollbarThumbHover: '148 163 184',
}

export const defaultDarkColors: ThemeModeColors = {
  primary: {
    50: '23 37 84', 100: '30 58 138', 200: '30 64 175', 300: '29 78 216',
    400: '37 99 235', 500: '59 130 246', 600: '96 165 250', 700: '147 197 253',
    800: '191 219 254', 900: '219 234 254', 950: '239 246 255',
  },
  gray: {
    50: '3 7 18', 100: '17 24 39', 200: '31 41 55', 300: '55 65 81',
    400: '107 114 128', 500: '156 163 175', 600: '209 213 219', 700: '229 231 235',
    800: '243 244 246', 900: '249 250 251', 950: '255 255 255',
  },
  success: { main: '74 222 128', light: '20 83 45' },
  warning: { main: '250 204 21', light: '113 63 18' },
  error: { main: '248 113 113', light: '127 29 29' },
  info: { main: '96 165 250', light: '30 58 138' },

  bg: '3 7 18',               // gray-950
  bgSoft: '17 24 39',         // gray-900
  bgMute: '31 41 55',         // gray-800

  surface: '17 24 39',        // gray-900
  surfaceHover: '31 41 55',   // gray-800
  surfaceActive: '55 65 81',  // gray-700

  text: '243 244 246',        // gray-100
  textSecondary: '209 213 219', // gray-300
  textMuted: '107 114 128',   // gray-500
  textInverse: '17 24 39',    // gray-900
  heading: '249 250 251',     // gray-50

  border: '55 65 81',         // gray-700
  borderLight: '31 41 55',    // gray-800
  borderHover: '107 114 128', // gray-500

  ring: '96 165 250',         // primary-400

  inputBg: '31 41 55',        // gray-800
  inputBorder: '55 65 81',    // gray-700
  inputText: '243 244 246',   // gray-100
  inputPlaceholder: '107 114 128', // gray-500

  link: '96 165 250',         // primary-400
  linkHover: '147 197 253',   // primary-300

  overlay: '0 0 0',

  scrollbarTrack: '17 24 39',
  scrollbarThumb: '55 65 81',
  scrollbarThumbHover: '107 114 128',
}

export const defaultTokens: ThemeTokens = {
  radiusSm: '0.25rem',
  radiusBase: '0.5rem',
  radiusLg: '0.75rem',
  radiusXl: '1rem',
  radiusFull: '9999px',
  shadowSm: '0 1px 2px 0 rgb(0 0 0 / 0.05)',
  shadowBase: '0 1px 3px 0 rgb(0 0 0 / 0.1), 0 1px 2px -1px rgb(0 0 0 / 0.1)',
  shadowLg: '0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1)',
  fontSans: "'Inter', ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, sans-serif",
  fontMono: "ui-monospace, SFMono-Regular, 'SF Mono', Menlo, Consolas, monospace",
}

export const defaultTheme: ThemeDefinition = {
  name: 'Default Blue',
  light: defaultLightColors,
  dark: defaultDarkColors,
  tokens: defaultTokens,
}
