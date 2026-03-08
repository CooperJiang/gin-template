#!/usr/bin/env node

import { existsSync, readFileSync } from 'node:fs'
import { spawnSync } from 'node:child_process'
import { dirname, resolve } from 'node:path'
import { fileURLToPath } from 'node:url'
import dns from 'node:dns/promises'
import https from 'node:https'
import net from 'node:net'

const __dirname = dirname(fileURLToPath(import.meta.url))
const ROOT = resolve(__dirname, '..')

/** @type {Array<{name: string, status: 'PASS' | 'WARN' | 'FAIL', detail: string}>} */
const checks = []

function addCheck(name, status, detail) {
  checks.push({ name, status, detail })
}

function run(bin, args) {
  const result = spawnSync(bin, args, {
    cwd: ROOT,
    encoding: 'utf8',
    stdio: ['ignore', 'pipe', 'pipe'],
  })

  if (result.error) {
    return {
      ok: false,
      stdout: '',
      stderr: result.error.message,
    }
  }

  return {
    ok: result.status === 0,
    stdout: (result.stdout || '').trim(),
    stderr: (result.stderr || '').trim(),
  }
}

function parseRegistryFromNpmrc(content) {
  const line = content
    .split('\n')
    .map((item) => item.trim())
    .find((item) => item.startsWith('registry='))
  return line ? line.replace('registry=', '').trim() : ''
}

function checkHttpsRegistry() {
  return new Promise((resolvePromise) => {
    const req = https.request(
      'https://registry.npmjs.org/react',
      {
        method: 'HEAD',
        timeout: 5000,
      },
      (res) => {
        const code = res.statusCode || 0
        if (code >= 200 && code < 500) {
          resolvePromise({ ok: true, detail: `HTTP ${code}` })
        } else {
          resolvePromise({ ok: false, detail: `HTTP ${code}` })
        }
      },
    )

    req.on('error', (err) => {
      resolvePromise({ ok: false, detail: err.message })
    })

    req.on('timeout', () => {
      req.destroy()
      resolvePromise({ ok: false, detail: 'Request timeout (5s)' })
    })

    req.end()
  })
}

function checkPort(port) {
  return new Promise((resolvePromise) => {
    const socket = new net.Socket()
    let done = false

    const finish = (isOpen) => {
      if (done) return
      done = true
      socket.destroy()
      resolvePromise(isOpen)
    }

    socket.setTimeout(700)
    socket.once('connect', () => finish(true))
    socket.once('timeout', () => finish(false))
    socket.once('error', () => finish(false))
    socket.connect(port, '127.0.0.1')
  })
}

async function main() {
  addCheck('workspace root', 'PASS', ROOT)

  const nodeMajor = Number(process.versions.node.split('.')[0] || 0)
  if (nodeMajor >= 18) {
    addCheck('node version', 'PASS', process.version)
  } else {
    addCheck('node version', 'FAIL', `${process.version} (Node >= 18 required)`)
  }

  const pnpmVersion = run('pnpm', ['--version'])
  if (pnpmVersion.ok) {
    addCheck('pnpm', 'PASS', pnpmVersion.stdout)
  } else {
    addCheck('pnpm', 'FAIL', pnpmVersion.stderr || 'pnpm not available in PATH')
  }

  if (pnpmVersion.ok) {
    const pnpmRegistry = run('pnpm', ['config', 'get', 'registry'])
    const registryValue = pnpmRegistry.stdout || '(empty)'
    if (registryValue.includes('registry.npmjs.org')) {
      addCheck('pnpm registry', 'PASS', registryValue)
    } else {
      addCheck('pnpm registry', 'WARN', registryValue)
    }
  }

  const npmrcPath = resolve(ROOT, '.npmrc')
  if (existsSync(npmrcPath)) {
    const npmrc = readFileSync(npmrcPath, 'utf8')
    const registry = parseRegistryFromNpmrc(npmrc)
    addCheck('.npmrc', 'PASS', registry || 'found')
  } else {
    addCheck('.npmrc', 'WARN', 'missing (optional)')
  }

  try {
    const dnsLookup = await dns.lookup('registry.npmjs.org')
    addCheck('dns registry.npmjs.org', 'PASS', `${dnsLookup.address}`)
  } catch (err) {
    addCheck('dns registry.npmjs.org', 'FAIL', err instanceof Error ? err.message : String(err))
  }

  const httpsRegistry = await checkHttpsRegistry()
  if (httpsRegistry.ok) {
    addCheck('https registry access', 'PASS', httpsRegistry.detail)
  } else {
    addCheck('https registry access', 'FAIL', httpsRegistry.detail)
  }

  const requiredPaths = [
    'pnpm-workspace.yaml',
    'packages/main/src/micro-app-registry.ts',
    'packages/main/vite.config.ts',
    'packages/app/vite.config.ts',
    'template/vite.config.ts',
  ]
  const missing = requiredPaths.filter((item) => !existsSync(resolve(ROOT, item)))
  if (missing.length === 0) {
    addCheck('workspace files', 'PASS', `${requiredPaths.length} required files found`)
  } else {
    addCheck('workspace files', 'FAIL', `missing: ${missing.join(', ')}`)
  }

  const port3000Open = await checkPort(3000)
  const port3001Open = await checkPort(3001)
  const port9000Open = await checkPort(9000)

  addCheck('port 3000 (main)', port3000Open ? 'PASS' : 'WARN', port3000Open ? 'running' : 'not running')
  addCheck('port 3001 (app)', port3001Open ? 'PASS' : 'WARN', port3001Open ? 'running' : 'not running')
  addCheck('port 9000 (backend)', port9000Open ? 'PASS' : 'WARN', port9000Open ? 'running' : 'not running')

  const statusOrder = { PASS: 0, WARN: 1, FAIL: 2 }
  checks.sort((a, b) => statusOrder[a.status] - statusOrder[b.status])

  console.log('\nWeb workspace doctor\n')
  for (const check of checks) {
    console.log(`[${check.status}] ${check.name}: ${check.detail}`)
  }

  const failCount = checks.filter((item) => item.status === 'FAIL').length
  const warnCount = checks.filter((item) => item.status === 'WARN').length

  console.log(`\nSummary: ${checks.length} checks, ${failCount} fail, ${warnCount} warn.`)
  if (failCount > 0) {
    console.log('Action: fix FAIL items first. For registry DNS problems, check local DNS/proxy/firewall settings.')
    process.exit(1)
  }

  console.log('Action: ready to continue development.')
}

main().catch((err) => {
  console.error('[FAIL] doctor crashed:', err)
  process.exit(1)
})
