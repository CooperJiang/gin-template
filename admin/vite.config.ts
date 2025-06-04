import { fileURLToPath, URL } from 'node:url'
import fs from 'fs'
import path from 'path'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'

// 读取版本信息
const getVersionInfo = () => {
  try {
    const packageJsonPath = path.join(process.cwd(), 'package.json')
    const versionConfigPath = path.join(process.cwd(), 'src/config/version.json')

    // 优先从version.json读取，如果不存在则从package.json读取
    if (fs.existsSync(versionConfigPath)) {
      const versionConfig = JSON.parse(fs.readFileSync(versionConfigPath, 'utf8'))
      return {
        version: versionConfig.version,
        buildTime: versionConfig.buildTime
      }
    } else {
      const packageJson = JSON.parse(fs.readFileSync(packageJsonPath, 'utf8'))
      return {
        version: packageJson.version,
        buildTime: new Date().toISOString()
      }
    }
  } catch (error) {
    console.warn('获取版本信息失败:', error)
    return {
      version: '1.0.0',
      buildTime: new Date().toISOString()
    }
  }
}

// https://vite.dev/config/
export default defineConfig(() => {
  const versionInfo = getVersionInfo()

  return {
    plugins: [
      vue(),
      vueDevTools(),
    ],
    base: '/admin/',
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url))
      },
    },
    server: {
      port: 5000,
    },
    define: {
      __APP_VERSION__: JSON.stringify(versionInfo.version),
      __BUILD_TIME__: JSON.stringify(versionInfo.buildTime),
    },
    envPrefix: ['VITE_', 'VUE_'],
    // 定义环境变量
    customEnv: {
      VITE_APP_VERSION: versionInfo.version,
      VITE_BUILD_TIME: versionInfo.buildTime,
    },
    assetsInclude: ['**/*.json'],
    json: {
      stringify: false
    }
  }
})
