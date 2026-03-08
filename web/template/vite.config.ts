import { fileURLToPath, URL } from 'node:url'
import { defineConfig, loadEnv } from 'vite'
import react from '@vitejs/plugin-react'
import qiankun from 'vite-plugin-qiankun'

export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), '')
  const apiProxyTarget = env.VITE_API_PROXY_TARGET || 'http://localhost:9000'

  return {
    plugins: [
      react(),
      qiankun('__NAME__', { useDevMode: true }),
    ],
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url)),
      },
    },
    base: mode === 'production' ? '/subapps/__NAME__/' : '/',
    server: {
      port: __PORT__,
      // 关闭 HMR，避免 React 插件注入 module preamble 与 qiankun 执行方式冲突
      hmr: false,
      cors: true,
      origin: 'http://localhost:__PORT__',
      proxy: {
        '/api': {
          target: apiProxyTarget,
          changeOrigin: true,
        },
      },
    },
    build: {
      outDir: 'dist',
    },
  }
})
