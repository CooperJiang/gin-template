import { fileURLToPath, URL } from 'node:url'
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import qiankun from 'vite-plugin-qiankun'

export default defineConfig(({ mode }) => ({
  plugins: [
    react(),
    qiankun('app', { useDevMode: true }),
  ],
  // 生产环境子应用从 /subapps/app/ 加载
  base: mode === 'production' ? '/subapps/app/' : '/',
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
  server: {
    port: 3001,
    // 关闭 HMR，避免 @vitejs/plugin-react 注入 <script type="module"> preamble，
    // 该脚本会被 qiankun 以普通脚本执行并触发 "Cannot use import statement outside a module"
    hmr: false,
    cors: true,
    origin: 'http://localhost:3001',
    proxy: {
      '/api': {
        target: 'http://localhost:9000',
        changeOrigin: true,
      },
    },
  },
  build: {
    outDir: 'dist',
  },
}))
