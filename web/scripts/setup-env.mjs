#!/usr/bin/env node

/**
 * 环境变量设置脚本
 * 自动检查并创建 .env.local 文件
 */

import { existsSync, readFileSync, writeFileSync } from 'node:fs'
import { resolve, dirname } from 'node:path'
import { fileURLToPath } from 'node:url'

const __dirname = dirname(fileURLToPath(import.meta.url))
const ROOT = resolve(__dirname, '..')
const ENV_LOCAL = resolve(ROOT, '.env.local')
const ENV_EXAMPLE = resolve(ROOT, '.env.example')

function setupEnv() {
  // 如果 .env.local 已存在，跳过
  if (existsSync(ENV_LOCAL)) {
    console.log('\x1b[32m✓\x1b[0m .env.local 已存在')
    return
  }

  console.log('\x1b[33m!\x1b[0m .env.local 不存在，正在创建...')

  // 从 .env.example 复制
  if (existsSync(ENV_EXAMPLE)) {
    const content = readFileSync(ENV_EXAMPLE, 'utf-8')
    writeFileSync(ENV_LOCAL, content, 'utf-8')
    console.log('\x1b[32m✓\x1b[0m 已从 .env.example 创建 .env.local')
    return
  }

  // 创建默认配置
  const defaultContent = `# 后端 API 代理目标
VITE_API_PROXY_TARGET=http://localhost:9000
`
  writeFileSync(ENV_LOCAL, defaultContent, 'utf-8')
  console.log('\x1b[32m✓\x1b[0m 已创建默认 .env.local')
}

setupEnv()
