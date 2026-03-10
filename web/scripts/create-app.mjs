#!/usr/bin/env node

/**
 * 子应用脚手架 CLI
 *
 * 用法：
 *   node scripts/create-app.mjs <name> [--port <port>] [--title <title>] [--install]
 *
 * 示例：
 *   node scripts/create-app.mjs admin --port 3002 --title "管理后台" --install
 */

import { existsSync, readFileSync, writeFileSync, readdirSync, statSync, cpSync } from 'node:fs'
import { resolve, join, dirname } from 'node:path'
import { fileURLToPath } from 'node:url'
import { execSync } from 'node:child_process'

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
  let install = false

  for (let i = 1; i < args.length; i++) {
    if (args[i] === '--port' && args[i + 1]) {
      port = parseInt(args[++i], 10)
    } else if (args[i] === '--title' && args[i + 1]) {
      title = args[++i]
    } else if (args[i] === '--install') {
      install = true
    }
  }

  return { name, port, title, install }
}

function printUsage() {
  console.log(`
  子应用脚手架 CLI

  用法：
    node scripts/create-app.mjs <name> [--port <port>] [--title <title>] [--install]

  参数：
    name            子应用名称（小写字母、数字、连字符）
    --port <port>   开发服务器端口（默认自动分配）
    --title <title> 应用标题（默认使用 name）
    --install       创建后自动运行 pnpm install

  示例：
    node scripts/create-app.mjs admin --port 3002 --title "管理后台" --install
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

  // 检查注册表中是否已存在同名应用
  if (existsSync(MICRO_APP_REGISTRY_FILE)) {
    const registryContent = readFileSync(MICRO_APP_REGISTRY_FILE, 'utf-8')
    const namePattern = new RegExp(`name:\\s*['"]${name}['"]`)
    if (namePattern.test(registryContent)) {
      console.error(`\x1b[31m错误：应用名称 "${name}" 已在 micro-app-registry.ts 中注册\x1b[0m`)
      process.exit(1)
    }
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
      // 支持两种格式：
      // 1. createViteConfig({ port: 3001 })
      // 2. server: { port: 3001 }
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

  function appendFilterToParallelDevScript(script) {
    if (!script || script.includes(`@app/${name}`)) return script
    if (script.includes('pnpm -r --parallel') && /\sdev\s*$/.test(script)) {
      return script.replace(/\sdev\s*$/, ` ${devFilter} dev`)
    }
    return `${script} & pnpm ${devFilter} dev`
  }

  // dev 脚本追加
  pkg.scripts.dev = appendFilterToParallelDevScript(pkg.scripts.dev)
  pkg.scripts['dev:full'] = appendFilterToParallelDevScript(pkg.scripts['dev:full'])

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

// ─── 创建 .env.local ───

function createEnvLocal() {
  const envLocalPath = join(ROOT, '.env.local')

  // 如果已存在，不覆盖
  if (existsSync(envLocalPath)) {
    return false
  }

  // 从 .env.example 复制，如果存在的话
  const envExamplePath = join(ROOT, '.env.example')
  if (existsSync(envExamplePath)) {
    const envExample = readFileSync(envExamplePath, 'utf-8')
    writeFileSync(envLocalPath, envExample, 'utf-8')
    return true
  }

  // 否则创建默认配置
  const defaultEnv = `# 后端 API 代理目标
VITE_API_PROXY_TARGET=http://localhost:9000
`
  writeFileSync(envLocalPath, defaultEnv, 'utf-8')
  return true
}

// ─── 安装依赖 ───

function installDependencies() {
  console.log(`\n\x1b[36m正在安装依赖...\x1b[0m`)
  try {
    execSync('pnpm install', { cwd: ROOT, stdio: 'inherit' })
    console.log(`\x1b[32m✓\x1b[0m 依赖安装完成`)
    return true
  } catch (error) {
    console.error(`\x1b[31m✗\x1b[0m 依赖安装失败：${error.message}`)
    return false
  }
}

// ─── 验证模板完整性 ───

function validateTemplate() {
  const requiredFiles = [
    'package.json',
    'vite.config.ts',
    'tsconfig.json',
    'tailwind.config.js',
    'src/main.tsx',
    'src/App.tsx',
  ]

  for (const file of requiredFiles) {
    const filePath = join(TEMPLATE_DIR, file)
    if (!existsSync(filePath)) {
      console.error(`\x1b[31m错误：模板文件缺失 - ${file}\x1b[0m`)
      return false
    }
  }
  return true
}

// ─── Rollback 机制 ───

function rollback(name, registryBackup, pkgBackup) {
  console.log(`\n\x1b[33m正在回滚更改...\x1b[0m`)

  // 删除创建的目录
  const targetDir = join(PACKAGES_DIR, name)
  if (existsSync(targetDir)) {
    try {
      execSync(`rm -rf "${targetDir}"`, { stdio: 'ignore' })
      console.log(`\x1b[32m✓\x1b[0m 已删除 packages/${name}/`)
    } catch (error) {
      console.error(`\x1b[31m✗\x1b[0m 删除目录失败：${error.message}`)
    }
  }

  // 恢复注册表
  if (registryBackup && existsSync(MICRO_APP_REGISTRY_FILE)) {
    try {
      writeFileSync(MICRO_APP_REGISTRY_FILE, registryBackup, 'utf-8')
      console.log(`\x1b[32m✓\x1b[0m 已恢复 micro-app-registry.ts`)
    } catch (error) {
      console.error(`\x1b[31m✗\x1b[0m 恢复注册表失败：${error.message}`)
    }
  }

  // 恢复 package.json
  if (pkgBackup && existsSync(ROOT_PKG_FILE)) {
    try {
      writeFileSync(ROOT_PKG_FILE, pkgBackup, 'utf-8')
      console.log(`\x1b[32m✓\x1b[0m 已恢复 package.json`)
    } catch (error) {
      console.error(`\x1b[31m✗\x1b[0m 恢复 package.json 失败：${error.message}`)
    }
  }

  console.log(`\x1b[33m回滚完成\x1b[0m\n`)
}

// ─── 主流程 ───

function main() {
  const { name, port: inputPort, title: inputTitle, install } = parseArgs()
  const port = inputPort || autoAssignPort()
  const title = inputTitle || name

  console.log(`\n\x1b[36m创建子应用：${name}\x1b[0m`)
  console.log(`  端口：${port}`)
  console.log(`  标题：${title}`)
  console.log()

  // 校验模板
  if (!validateTemplate()) {
    console.error(`\x1b[31m模板验证失败，请检查 template/ 目录\x1b[0m`)
    process.exit(1)
  }

  // 校验参数
  validate(name, port)

  // 备份文件（用于回滚）
  let registryBackup = null
  let pkgBackup = null

  try {
    // 读取备份
    if (existsSync(MICRO_APP_REGISTRY_FILE)) {
      registryBackup = readFileSync(MICRO_APP_REGISTRY_FILE, 'utf-8')
    }
    if (existsSync(ROOT_PKG_FILE)) {
      pkgBackup = readFileSync(ROOT_PKG_FILE, 'utf-8')
    }

    // 1. 复制模板
    const targetDir = copyTemplate(name, port, title)
    console.log(`\x1b[32m✓\x1b[0m 创建 packages/${name}/`)

    // 2. 注册到主基座（微应用 + 菜单）
    const registered = registerMicroAppConfig(name, port, title)
    if (!registered) {
      throw new Error('注册到 micro-app-registry.ts 失败')
    }
    console.log(`\x1b[32m✓\x1b[0m 注册到 micro-app-registry.ts（包含菜单）`)

    // 3. 更新 package.json
    updateRootPackageJson(name)
    console.log(`\x1b[32m✓\x1b[0m 更新 package.json`)

    // 4. 创建 .env.local（如果不存在）
    const envCreated = createEnvLocal()
    if (envCreated) {
      console.log(`\x1b[32m✓\x1b[0m 创建 .env.local`)
    }

    // 5. 验证创建结果
    const requiredFiles = ['package.json', 'vite.config.ts', 'src/main.tsx']
    for (const file of requiredFiles) {
      const filePath = join(PACKAGES_DIR, name, file)
      if (!existsSync(filePath)) {
        throw new Error(`关键文件缺失：${file}`)
      }
    }

    // 6. 安装依赖（如果指定了 --install）
    if (install) {
      const installed = installDependencies()
      if (!installed) {
        console.log(`\x1b[33m⚠\x1b[0m 依赖安装失败，但子应用已创建成功`)
        console.log(`请手动运行：pnpm install`)
      }
    }

    // 完成
    const nextSteps = install
      ? `后续步骤：
  1. \x1b[33mpnpm dev:${name}\x1b[0m           单独启动子应用
  2. \x1b[33mpnpm dev\x1b[0m                   启动所有应用`
      : `后续步骤：
  1. \x1b[33mpnpm install\x1b[0m              安装依赖
  2. \x1b[33mpnpm dev:${name}\x1b[0m           单独启动子应用
  3. \x1b[33mpnpm dev\x1b[0m                   启动所有应用`

    console.log(`
\x1b[32m子应用 ${name} 创建成功！\x1b[0m

${nextSteps}

目录：packages/${name}/
端口：http://localhost:${port}
路由：/${name}
`)
  } catch (error) {
    console.error(`\n\x1b[31m错误：${error.message}\x1b[0m`)
    rollback(name, registryBackup, pkgBackup)
    process.exit(1)
  }
}

main()
