# Troubleshooting (web)

## 1. `ENOTFOUND registry.npmjs.org` when running pnpm install

Symptoms:
- `pnpm install` fails with DNS resolve errors
- `curl https://registry.npmjs.org/react` cannot resolve host

Checks:
```bash
pnpm -C web doctor
```

Common causes:
- DNS server cannot resolve npm registry
- corporate proxy/firewall restrictions
- local network profile mismatch (VPN/proxy DNS split)

What to verify:
1. Registry config points to npmjs:
```bash
pnpm config get registry
```
Expected: `https://registry.npmjs.org`
2. Local `.npmrc` exists (`web/.npmrc`) and has correct registry.
3. DNS can resolve registry host.
4. Try different network or DNS profile if needed.

## 2. Login API returns 404 (`/api/user/login`)

Expected local flow:
- Frontend host: `http://localhost:3000`
- Frontend child: `http://localhost:3001`
- Backend API: `http://localhost:9000`

`/api/*` is proxied by Vite to `VITE_API_PROXY_TARGET` (default `http://localhost:9000`).

Verify:
1. Backend process is running on `9000`.
2. In `web/.env.local`, if backend uses another port, set:
```bash
VITE_API_PROXY_TARGET=http://localhost:<your-backend-port>
```
3. Restart frontend dev servers after env changes.

## 3. Child app load error with `Cannot use import statement outside a module`

This is usually caused by Vite HMR preamble scripts being executed by qiankun as normal scripts.

Fix in child app Vite config:
- `server.hmr = false`
- `server.cors = true`
- `server.origin = 'http://localhost:<child-port>'`

The default app and template already include these settings.

## 4. Child app not mounting in host

Checklist:
1. `entryDev` in `web/packages/main/src/micro-app-registry.ts` matches the child app dev port.
2. `activeRule` is unique and starts with `/`.
3. Child app runs and is reachable directly in browser.
4. Host app is running and route matches child `activeRule`.

## 5. Runtime errors from browser extensions

If errors reference unknown globals/functions from minified bundle but code looks correct, test in incognito mode or disable extensions first.

A previous `useLogo is not defined` issue was caused by browser extension injection, not workspace code.
