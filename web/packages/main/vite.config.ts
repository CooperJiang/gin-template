import { defineConfig } from 'vite'
import { createViteConfig } from '../../scripts/vite-config-base'

export default defineConfig(
  createViteConfig({
    port: 3000,
    isQiankunApp: false,
    projectRoot: import.meta.url,
  }),
)
