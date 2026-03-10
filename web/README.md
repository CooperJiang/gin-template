# Web Micro-Frontend Workspace (React + qiankun)

This directory is a React-first micro-frontend workspace. It includes a qiankun host app, one child app scaffold, shared packages, and a CLI to generate new child apps.

## 1. Quick Start

1. Install dependencies:
```bash
pnpm -C web setup
```
2. Optional environment file:
```bash
cp web/.env.example web/.env.local
```
3. Start host + default child app:
```bash
pnpm -C web dev
```
4. Open host app:
- `http://localhost:3000`

Backend API default target is `http://localhost:9000` via `/api` proxy.

## 2. Useful Commands

- `pnpm -C web start`: alias of `dev`
- `pnpm -C web dev`: run host (`main`) + default child (`app`)
- `pnpm -C web dev:full`: run host + all locally registered child apps + theme-editor
- `pnpm -C web dev:main`: run host only
- `pnpm -C web dev:app`: run default child only
- `pnpm -C web dev:theme`: run theme editor only
- `pnpm -C web build`: build host + locally registered child apps
- `pnpm -C web type-check`: type-check all workspace packages
- `pnpm -C web doctor`: run environment diagnostics
- `pnpm -C web create:app -- <name> --port <port> --title <title>`: scaffold a new child app

## 3. Workspace Structure

- `web/packages/main`: qiankun host app (React)
- `web/packages/app`: default child app template instance (React)
- `web/packages/theme-editor`: optional tool app (React)
- `web/shared`: shared types/utils/theme package
- `web/core`: core services/styles package
- `web/auth`: auth state and auth pages package
- `web/template`: source template used by the app generator
- `web/scripts/create-app.mjs`: child-app generator CLI
- `web/scripts/doctor.mjs`: local environment checks

## 4. How qiankun Is Wired

1. Child app registry lives in `web/packages/main/src/micro-app-registry.ts`.
2. Host app converts registry to qiankun registrations in `web/packages/main/src/micro-apps.ts`.
3. Host layout menu links are generated from registry via `getMicroAppMenuLinks()`.
4. New child apps are auto-registered by `create:app` (registry + root scripts).

## 5. Create a New Child App

Example:
```bash
pnpm -C web create:app -- admin --port 3002 --title "Admin"
```

This command will:
1. Copy `web/template` to `web/packages/admin`
2. Replace placeholders (name/port/title)
3. Append app config into `web/packages/main/src/micro-app-registry.ts`
4. Update root scripts (`dev`, `dev:full`, `build`, `dev:admin`, `build:admin`)
5. Include the app in production packaging flow (`scripts/build_web.sh`) via registry-based discovery

Then run:
```bash
pnpm -C web dev:admin
pnpm -C web dev
```

No manual update is required in Go client routes for new sub-apps. `internal/routes/client_routes.go` now serves `/subapps/:app/*` dynamically.

## 6. Environment Variables

Use `web/.env.local` (or shell env):

- `VITE_API_PROXY_TARGET`: backend proxy target, default `http://localhost:9000`
- `VITE_QIANKUN_PREFETCH`: host prefetch mode, use `all` to enable full prefetch in production

## 7. Troubleshooting

See `web/docs/TROUBLESHOOTING.md`.

If dependency install fails with DNS/registry errors, run:
```bash
pnpm -C web doctor
```
and follow the FAIL/WARN hints.
