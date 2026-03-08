import baseConfig from '@app/core/tailwind-config'

/** @type {import('tailwindcss').Config} */
export default {
  ...baseConfig,
  // 主应用所有工具类加 !important，防止被子应用 CSS 覆盖
  important: true,
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
    "../../auth/src/**/*.{js,ts,jsx,tsx}",
  ],
}
