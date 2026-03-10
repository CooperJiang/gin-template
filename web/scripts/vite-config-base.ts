import { fileURLToPath, URL } from 'node:url'
import { loadEnv, type UserConfig } from 'vite'
import react from '@vitejs/plugin-react'

export interface ViteConfigOptions {
  /**
   * 应用名称（用于qiankun子应用）
   */
  appName?: string
  /**
   * 开发服务器端口
   */
  port: number
  /**
   * 是否为qiankun子应用
   */
  isQiankunApp?: boolean
  /**
   * 生产环境的base路径
   */
  productionBase?: string
  /**
   * 项目根目录（用于解析路径别名）
   */
  projectRoot: string
}

/**
 * 创建统一的Vite配置
 */
export function createViteConfig(options: ViteConfigOptions) {
  const { appName, port, isQiankunApp = false, productionBase, projectRoot } = options

  return ({ mode }: { mode: string }): UserConfig => {
    const env = loadEnv(mode, process.cwd(), '')
    const apiProxyTarget = env.VITE_API_PROXY_TARGET || 'http://localhost:9000'

    const config: UserConfig = {
      plugins: [react()],
      resolve: {
        alias: {
          '@': fileURLToPath(new URL('./src', projectRoot)),
        },
      },
      server: {
        port,
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

    // qiankun子应用特定配置
    if (isQiankunApp && appName) {
      const qiankun = require('vite-plugin-qiankun').default

      config.plugins!.push(qiankun(appName, { useDevMode: true }))

      // 生产环境子应用路径
      config.base = mode === 'production' && productionBase ? productionBase : '/'

      // qiankun子应用需要关闭HMR，避免模块导入冲突
      if (config.server) {
        config.server.hmr = false
        config.server.cors = true
        config.server.origin = `http://localhost:${port}`
      }
    }

    return config
  }
}
