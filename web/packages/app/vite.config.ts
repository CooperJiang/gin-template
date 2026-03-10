import { defineConfig } from 'vite'
import { createViteConfig } from '../../scripts/vite-config-base'

export default defineConfig(
  createViteConfig({
    appName: 'app',
    port: 3001,
    isQiankunApp: true,
    productionBase: '/subapps/app/',
    projectRoot: import.meta.url,
  }),
)
