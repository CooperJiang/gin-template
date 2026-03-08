import baseConfig from '@app/core/tailwind-config'

/** @type {import('tailwindcss').Config} */
export default {
  ...baseConfig,
  corePlugins: { preflight: false },
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
    "../../auth/src/**/*.{js,ts,jsx,tsx}",
  ],
}
