import { defineConfig } from 'vite'
import { createViteConfig } from '../../scripts/vite-config-base'

export default defineConfig(
  createViteConfig({
    appName: '__NAME__',
    port: __PORT__,
    isQiankunApp: true,
    productionBase: '/subapps/__NAME__/',
    projectRoot: import.meta.url,
  }),
)
