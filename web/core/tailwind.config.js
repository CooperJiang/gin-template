import forms from '@tailwindcss/forms'
import typography from '@tailwindcss/typography'

/** Helper: reference a --theme-* RGB triplet CSS variable with alpha support */
const themeColor = (name) => `rgb(var(--theme-${name}) / <alpha-value>)`

/** Generate a full color scale object from a prefix */
const colorScale = (prefix) => ({
  50: themeColor(`${prefix}-50`),
  100: themeColor(`${prefix}-100`),
  200: themeColor(`${prefix}-200`),
  300: themeColor(`${prefix}-300`),
  400: themeColor(`${prefix}-400`),
  500: themeColor(`${prefix}-500`),
  600: themeColor(`${prefix}-600`),
  700: themeColor(`${prefix}-700`),
  800: themeColor(`${prefix}-800`),
  900: themeColor(`${prefix}-900`),
  950: themeColor(`${prefix}-950`),
})

/** @type {import('tailwindcss').Config} */
export default {
  darkMode: 'class',
  content: [
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: colorScale('primary'),
        gray: colorScale('gray'),
        success: { DEFAULT: themeColor('success-main'), light: themeColor('success-light') },
        warning: { DEFAULT: themeColor('warning-main'), light: themeColor('warning-light') },
        error: { DEFAULT: themeColor('error-main'), light: themeColor('error-light') },
        info: { DEFAULT: themeColor('info-main'), light: themeColor('info-light') },
        surface: {
          DEFAULT: themeColor('surface'),
          hover: themeColor('surface-hover'),
          active: themeColor('surface-active'),
        },
        'app-bg': {
          DEFAULT: themeColor('bg'),
          soft: themeColor('bg-soft'),
          mute: themeColor('bg-mute'),
        },
        'app-text': {
          DEFAULT: themeColor('text'),
          secondary: themeColor('text-secondary'),
          muted: themeColor('text-muted'),
          inverse: themeColor('text-inverse'),
        },
        'app-heading': themeColor('heading'),
        'app-border': {
          DEFAULT: themeColor('border'),
          light: themeColor('border-light'),
          hover: themeColor('border-hover'),
        },
        'app-ring': themeColor('ring'),
        'app-input': {
          bg: themeColor('input-bg'),
          border: themeColor('input-border'),
          text: themeColor('input-text'),
          placeholder: themeColor('input-placeholder'),
        },
        'app-link': { DEFAULT: themeColor('link'), hover: themeColor('link-hover') },
        'app-overlay': themeColor('overlay'),
      },
      fontFamily: {
        sans: ['var(--theme-font-sans)'],
        mono: ['var(--theme-font-mono)'],
      },
      borderRadius: {
        sm: 'var(--theme-radius-sm)',
        DEFAULT: 'var(--theme-radius-base)',
        md: 'var(--theme-radius-base)',
        lg: 'var(--theme-radius-lg)',
        xl: 'var(--theme-radius-xl)',
        full: 'var(--theme-radius-full)',
      },
      boxShadow: {
        sm: 'var(--theme-shadow-sm)',
        DEFAULT: 'var(--theme-shadow-base)',
        md: 'var(--theme-shadow-base)',
        lg: 'var(--theme-shadow-lg)',
      },
    },
  },
  plugins: [
    forms,
    typography,
  ],
}
