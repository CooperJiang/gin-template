#!/usr/bin/env node

/**
 * 子应用脚手架 CLI
 *
 * 用法：
 *   node scripts/create-app.mjs <name> [--port <port>] [--title <title>]
 *
 * 示例：
 *   node scripts/create-app.mjs admin --port 3002 --title "管理后台"
 */

import { existsSync, readFileSync, writeFileSync, readdirSync, statSync, cpSync } from 'node:fs'
import { resolve, join, dirname } from 'node:path'
import { fileURLToPath } from 'node:url'

const __dirname = dirname(fileURLToPath(import.meta.url))
const ROOT = resolve(__dirname, '..')
const TEMPLATE_DIR = join(ROOT, 'template')
const PACKAGES_DIR = join(ROOT, 'packages')
const MICRO_APP_REGISTRY_FILE = join(PACKAGES_DIR, 'main', 'src', 'micro-app-registry.ts')
const ROOT_PKG_FILE = join(ROOT, 'package.json')

// ─── 参数解析 ───

function parseArgs() {
  const rawArgs = process.argv.slice(2)
  const args = rawArgs[0] === '--' ? rawArgs.slice(1) : rawArgs
  if (args.length === 0 || args[0] === '--help' || args[0] === '-h') {
    printUsage()
    process.exit(0)
  }

  const name = args[0]
  let port = null
  let title = null

  for (let i = 1; i < args.length; i++) {
    if (args[i] === '--port' && args[i + 1]) {
      port = parseInt(args[++i], 10)
    } else if (args[i] === '--title' && args[i + 1]) {
      title = args[++i]
    }
  }

  return { name, port, title }
}

function printUsage() {
  console.log(`
  子应用脚手架 CLI

  用法：
    node scripts/create-app.mjs <name> [--port <port>] [--title <title>]

  参数：
    name          子应用名称（小写字母、数字、连字符）
    --port <port> 开发服务器端口（默认自动分配）
    --title <title> 应用标题（默认使用 name）

  示例：
    node scripts/create-app.mjs admin --port 3002 --title "管理后台"
    node scripts/create-app.mjs dashboard
  `)
}

// ─── 校验 ───

function validate(name, port) {
  if (!/^[a-z][a-z0-9-]*$/.test(name)) {
    console.error(`\x1b[31m错误：名称 "${name}" 不合法，只允许小写字母、数字和连字符，且以字母开头\x1b[0m`)
    process.exit(1)
  }

  if (name === 'main') {
    console.error(`\x1b[31m错误：不能使用 "main" 作为子应用名称（已被主基座占用）\x1b[0m`)
    process.exit(1)
  }

  const targetDir = join(PACKAGES_DIR, name)
  if (existsSync(targetDir)) {
    console.error(`\x1b[31m错误：packages/${name}/ 已存在\x1b[0m`)
    process.exit(1)
  }

  const usedPorts = getUsedPorts()
  if (usedPorts.includes(port)) {
    console.error(`\x1b[31m错误：端口 ${port} 已被其他子应用占用（已使用端口：${usedPorts.join(', ')}）\x1b[0m`)
    process.exit(1)
  }
}

// ─── 端口扫描 ───

function getUsedPorts() {
  const ports = []
  const dirs = readdirSync(PACKAGES_DIR)
  for (const dir of dirs) {
    const viteConfig = join(PACKAGES_DIR, dir, 'vite.config.ts')
    if (existsSync(viteConfig)) {
      const content = readFileSync(viteConfig, 'utf-8')
      const match = content.match(/port:\s*(\d+)/)
      if (match) {
        ports.push(parseInt(match[1], 10))
      }
    }
  }
  return ports
}

function autoAssignPort() {
  const ports = getUsedPorts()
  if (ports.length === 0) return 3001
  return Math.max(...ports) + 1
}

// ─── 模板复制 + 占位符替换 ───

function copyTemplate(name, port, title) {
  const targetDir = join(PACKAGES_DIR, name)

  // 递归复制
  cpSync(TEMPLATE_DIR, targetDir, { recursive: true })

  // 递归替换占位符
  replaceInDir(targetDir, {
    '__NAME__': name,
    '__PORT__': String(port),
    '__TITLE__': title,
  })

  return targetDir
}

function replaceInDir(dir, replacements) {
  const entries = readdirSync(dir)
  for (const entry of entries) {
    const fullPath = join(dir, entry)
    const stat = statSync(fullPath)
    if (stat.isDirectory()) {
      replaceInDir(fullPath, replacements)
    } else {
      let content = readFileSync(fullPath, 'utf-8')
      let changed = false
      for (const [placeholder, value] of Object.entries(replacements)) {
        if (content.includes(placeholder)) {
          content = content.replaceAll(placeholder, value)
          changed = true
        }
      }
      if (changed) {
        writeFileSync(fullPath, content, 'utf-8')
      }
    }
  }
}

// ─── 注册到主基座配置 ───

function getNextMenuOrder(content) {
  const matches = [...content.matchAll(/menuOrder:\s*(\d+)/g)]
  if (matches.length === 0) return 10
  const maxOrder = Math.max(...matches.map((m) => parseInt(m[1], 10)))
  return maxOrder + 10
}

function registerMicroAppConfig(name, port, title) {
  let content = readFileSync(MICRO_APP_REGISTRY_FILE, 'utf-8')
  const menuOrder = getNextMenuOrder(content)

  const newEntry = `  {
    name: '${name}',
    title: '${title}',
    activeRule: '/${name}',
    entryDev: 'http://localhost:${port}',
    entryProd: '/subapps/${name}/',
    showInMenu: true,
    menuOrder: ${menuOrder},
  },`

  // 在 microAppRegistry 数组的 ] 前插入
  const arrayEndRegex = /(\n)(]\s*\n\s*\nexport function getMicroAppMenuLinks)/
  const match = content.match(arrayEndRegex)

  if (match) {
    content = content.replace(arrayEndRegex, `\n${newEntry}\n$2`)
  } else {
    // fallback: 找最后一个 },\n] 模式
    const fallbackRegex = /(},?\s*\n)(]\s*)/
    const matches = [...content.matchAll(new RegExp(fallbackRegex, 'g'))]
    if (matches.length > 0) {
      const lastMatch = matches[matches.length - 1]
      const insertPos = lastMatch.index + lastMatch[1].length
      content = content.slice(0, insertPos) + newEntry + '\n' + content.slice(insertPos)
    } else {
      console.error('\x1b[31m错误：无法自动注册到 micro-app-registry.ts，请手动添加\x1b[0m')
      return false
    }
  }

  writeFileSync(MICRO_APP_REGISTRY_FILE, content, 'utf-8')
  return true
}

// ─── 更新 package.json ───

function updateRootPackageJson(name) {
  const pkg = JSON.parse(readFileSync(ROOT_PKG_FILE, 'utf-8'))
  const devFilter = `--filter @app/${name}`

  // dev 脚本追加
  if (pkg.scripts.dev && !pkg.scripts.dev.includes(`@app/${name}`)) {
    if (pkg.scripts.dev.includes('pnpm -r --parallel') && /\sdev\s*$/.test(pkg.scripts.dev)) {
      pkg.scripts.dev = pkg.scripts.dev.replace(/\sdev\s*$/, ` ${devFilter} dev`)
    } else {
      pkg.scripts.dev += ` & pnpm --filter @app/${name} dev`
    }
  }

  // build 脚本追加
  if (pkg.scripts.build && !pkg.scripts.build.includes(`@app/${name}`)) {
    pkg.scripts.build += ` && pnpm --filter @app/${name} build`
  }

  // 添加独立 dev/build 脚本
  pkg.scripts[`dev:${name}`] = `pnpm --filter @app/${name} dev`
  pkg.scripts[`build:${name}`] = `pnpm --filter @app/${name} build`

  // 添加 create 脚本（如果没有）
  if (!pkg.scripts.create) {
    pkg.scripts.create = 'node scripts/create-app.mjs'
  }
  if (!pkg.scripts['create:app']) {
    pkg.scripts['create:app'] = 'node scripts/create-app.mjs'
  }

  writeFileSync(ROOT_PKG_FILE, JSON.stringify(pkg, null, 2) + '\n', 'utf-8')
}

// ─── 主流程 ───

function main() {
  const { name, port: inputPort, title: inputTitle } = parseArgs()
  const port = inputPort || autoAssignPort()
  const title = inputTitle || name

  console.log(`\n\x1b[36m创建子应用：${name}\x1b[0m`)
  console.log(`  端口：${port}`)
  console.log(`  标题：${title}`)
  console.log()

  // 校验
  validate(name, port)

  // 1. 复制模板
  const targetDir = copyTemplate(name, port, title)
  console.log(`\x1b[32m✓\x1b[0m 创建 packages/${name}/`)

  // 2. 注册到主基座（微应用 + 菜单）
  const registered = registerMicroAppConfig(name, port, title)
  if (registered) {
    console.log(`\x1b[32m✓\x1b[0m 注册到 micro-app-registry.ts（包含菜单）`)
  }

  // 3. 更新 package.json
  updateRootPackageJson(name)
  console.log(`\x1b[32m✓\x1b[0m 更新 package.json`)

  // 完成
  console.log(`
\x1b[32m子应用 ${name} 创建成功！\x1b[0m

后续步骤：
  1. \x1b[33mpnpm install\x1b[0m              安装依赖
  2. \x1b[33mpnpm dev:${name}\x1b[0m           单独启动子应用
  3. \x1b[33mpnpm dev\x1b[0m                   启动所有应用

目录：packages/${name}/
端口：http://localhost:${port}
路由：/${name}
`)
}

main()
