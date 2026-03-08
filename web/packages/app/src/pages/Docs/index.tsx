import { useState } from 'react'

const sections = [
  { id: 'quickstart', label: '快速开始' },
  { id: 'structure', label: '项目结构' },
  { id: 'dev', label: '开发指南' },
  { id: 'add-subapp', label: '新增子应用' },
  { id: 'auth', label: '认证系统' },
  { id: 'build', label: '构建部署' },
  { id: 'env', label: '环境变量' },
]

export default function Docs() {
  const [active, setActive] = useState('quickstart')

  const scrollTo = (id: string) => {
    setActive(id)
    document.getElementById(id)?.scrollIntoView({ behavior: 'smooth', block: 'start' })
  }

  return (
    <div className="bg-gray-50 min-h-screen">
      <div className="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
        <div className="mb-10">
          <h1 className="text-3xl font-bold text-gray-900 tracking-tight">使用文档</h1>
          <p className="text-gray-500 mt-2">从零开始搭建、开发和部署你的微前端应用</p>
        </div>

        <div className="flex gap-8">
          {/* Sidebar */}
          <nav className="hidden lg:block shrink-0" style={{ width: 180 }}>
            <div className="sticky top-20 space-y-0.5">
              {sections.map((s) => (
                <button
                  key={s.id}
                  onClick={() => scrollTo(s.id)}
                  className={`block w-full text-left px-3 py-2 rounded-lg text-[13px] font-medium border-none cursor-pointer transition-colors ${
                    active === s.id
                      ? 'text-blue-600 bg-blue-50'
                      : 'text-gray-500 bg-transparent hover:text-gray-700 hover:bg-gray-100'
                  }`}
                >
                  {s.label}
                </button>
              ))}
            </div>
          </nav>

          {/* Content */}
          <div className="min-w-0 flex-1 space-y-6">

            <Section id="quickstart" title="快速开始">
              <P>确保本地已安装 Node.js 18+、pnpm 和 Go 1.22+。</P>
              <Code>{`# 克隆项目
git clone <repo-url>
cd gin-template

# 安装前端依赖
cd web && pnpm install

# 启动前端开发服务器（主应用 + 所有子应用）
pnpm dev

# 另一个终端，启动 Go 后端
cd gin-template
go run cmd/server/main.go`}</Code>
              <P>启动后访问 <Mono>http://localhost:3000</Mono>，主应用运行在 3000 端口，子应用 app 运行在 3001 端口，Go 后端运行在 9000 端口。</P>
            </Section>

            <Section id="structure" title="项目结构">
              <Code>{`gin-template/
├── cmd/server/          # Go 入口
├── internal/            # Go 业务代码
│   ├── routes/          # 路由定义
│   ├── static/          # embed 静态资源（构建后生成）
│   └── ...
├── scripts/             # 构建脚本
│   └── build_web.sh     # 前端打包脚本
├── web/                 # 前端 monorepo
│   ├── shared/          # @app/shared - 类型定义、工具函数
│   ├── core/            # @app/core - HTTP 客户端、消息提示、样式
│   ├── auth/            # @app/auth - 认证状态管理、登录注册页面
│   ├── packages/
│   │   ├── main/        # 主基座（端口 3000）
│   │   └── app/         # 子应用（端口 3001）
│   └── template/        # 子应用模板
└── Makefile`}</Code>
              <P>前端采用 pnpm workspace monorepo，<Mono>shared</Mono>、<Mono>core</Mono>、<Mono>auth</Mono> 是共享包，所有子应用通过 workspace 协议引用。</P>
            </Section>

            <Section id="dev" title="开发指南">
              <H3>常用命令</H3>
              <Code>{`# 启动所有前端应用
pnpm dev

# 只启动主应用
pnpm --filter @app/main dev

# 只启动子应用 app
pnpm --filter @app/app dev

# 类型检查
pnpm type-check

# 构建所有前端应用
pnpm build`}</Code>

              <H3>开发模式说明</H3>
              <P>开发时主应用（3000）通过 qiankun 加载子应用（3001）。子应用也可以独立运行，此时会自动提供登录/注册页面。</P>
              <P>子应用的 Vite 配置关闭了 HMR（<Mono>hmr: false</Mono>），因为 qiankun 的沙箱与 React HMR 的 module preamble 注入存在冲突。修改代码后需手动刷新。</P>

              <H3>API 代理</H3>
              <P>开发模式下，所有 <Mono>/api</Mono> 请求会被 Vite 代理到 <Mono>http://localhost:9000</Mono>（Go 后端）。</P>
            </Section>

            <Section id="add-subapp" title="新增子应用">
              <P>项目提供了标准化的子应用模板，按以下步骤创建新子应用：</P>

              <H3>1. 复制模板</H3>
              <Code>{`cp -r web/template web/packages/<name>
# 例如：cp -r web/template web/packages/admin`}</Code>

              <H3>2. 替换占位符</H3>
              <P>在新子应用目录中，全局替换以下占位符：</P>
              <Table
                headers={['占位符', '说明', '示例']}
                rows={[
                  ['__NAME__', '子应用名称（英文）', 'admin'],
                  ['__PORT__', '开发服务器端口', '3002'],
                  ['__TITLE__', '页面标题（中文）', '管理后台'],
                ]}
              />
              <Code>{`# macOS / Linux 批量替换
cd web/packages/admin
find . -type f \\( -name "*.ts" -o -name "*.tsx" -o -name "*.json" -o -name "*.html" \\) \\
  -exec sed -i '' 's/__NAME__/admin/g; s/__PORT__/3002/g; s/__TITLE__/管理后台/g' {} +`}</Code>

              <H3>3. 注册到 workspace</H3>
              <P>编辑 <Mono>web/pnpm-workspace.yaml</Mono>，确认 <Mono>packages/*</Mono> 已包含新目录（通常通配符已覆盖）。然后在 web 根目录执行：</P>
              <Code>{`pnpm install`}</Code>

              <H3>4. 注册到主基座</H3>
              <P>编辑 <Mono>web/packages/main/src/micro-apps.ts</Mono>，在 <Mono>microApps</Mono> 数组中添加：</P>
              <Code>{`{
  name: 'admin',
  entry: isDev ? '//localhost:3002' : '/subapps/admin/',
  container: '#subapp-container',
  activeRule: '/admin',
  props: { globalStateActions },
}`}</Code>

              <H3>5. 添加导航链接</H3>
              <P>编辑 <Mono>web/packages/main/src/components/Layout.tsx</Mono>，在 <Mono>navLinks</Mono> 中添加：</P>
              <Code>{`{ href: '/admin', label: '管理后台' }`}</Code>

              <H3>6. 更新构建脚本</H3>
              <P>编辑 <Mono>scripts/build_web.sh</Mono>，添加新子应用的 dist 复制逻辑：</P>
              <Code>{`ADMIN_DIST_DIR="\${WEB_DIR}/packages/admin/dist"

# 在构建检查后添加
mkdir -p "$TARGET_WEB_DIR/subapps/admin"
cp -r "$ADMIN_DIST_DIR"/* "$TARGET_WEB_DIR/subapps/admin/"`}</Code>

              <H3>7. 添加 Go 路由（生产部署）</H3>
              <P>编辑 <Mono>internal/routes/client_routes.go</Mono>，参照 <Mono>/subapps/app/*filepath</Mono> 的模式，添加 <Mono>/subapps/admin/*filepath</Mono> 路由。</P>
            </Section>

            <Section id="auth" title="认证系统">
              <P>认证由 <Mono>@app/auth</Mono> 包统一管理，所有子应用通过 <Mono>useAuth()</Mono> hook 获取认证状态。</P>

              <H3>核心 API</H3>
              <Code>{`import { useAuth } from '@app/auth'

const {
  user,              // 当前用户信息
  isAuthenticated,   // 是否已登录
  loading,           // 请求中
  login,             // 登录
  register,          // 注册
  logout,            // 登出
  getUserInfo,       // 刷新用户信息
} = useAuth()`}</Code>

              <H3>路由守卫</H3>
              <P>在路由中使用 <Mono>RequireAuth</Mono> 组件包裹需要登录的页面：</P>
              <Code>{`<Route path="/profile" element={
  <RequireAuth><Profile /></RequireAuth>
} />`}</Code>
              <P>未登录用户会被重定向到 <Mono>/login?redirect=原路径</Mono>，登录后自动跳回。</P>

              <H3>跨应用状态同步</H3>
              <P>认证状态通过三种机制同步：</P>
              <ul className="list-disc list-inside text-[14px] text-gray-600 space-y-1 ml-1">
                <li>qiankun globalState — 主应用与子应用间</li>
                <li>CustomEvent — 同 tab 内子应用通知主应用</li>
                <li>storage 事件 — 跨 tab 同步</li>
              </ul>
            </Section>

            <Section id="build" title="构建部署">
              <H3>方式一：Go embed 单二进制</H3>
              <Code>{`# 一键构建（前端 + 后端）
make build

# 或分步执行
bash scripts/build_web.sh   # 打包前端到 internal/static/web/
go build -o server cmd/server/main.go

# 运行
./server`}</Code>
              <P>构建后前端静态文件被 embed 到 Go 二进制中，部署时只需一个可执行文件。</P>

              <H3>方式二：Nginx 反向代理</H3>
              <Code>{`# 只构建前端
cd web && pnpm build`}</Code>
              <P>构建产物在 <Mono>web/packages/main/dist</Mono>（主应用）和 <Mono>web/packages/app/dist</Mono>（子应用）。Nginx 配置示例：</P>
              <Code>{`server {
    listen 80;

    # 主应用
    location / {
        root /path/to/main/dist;
        try_files $uri $uri/ /index.html;
    }

    # 子应用
    location /subapps/app/ {
        alias /path/to/app/dist/;
        try_files $uri $uri/ /subapps/app/index.html;
    }

    # API 代理
    location /api/ {
        proxy_pass http://127.0.0.1:9000;
    }
}`}</Code>
            </Section>

            <Section id="env" title="环境变量">
              <P>每个子应用根目录有 <Mono>.env</Mono> 文件：</P>
              <Table
                headers={['变量', '说明', '默认值']}
                rows={[
                  ['VITE_API_BASE_URL', 'API 基础路径', '/api'],
                  ['VITE_APP_TITLE', '应用标题', '应用平台'],
                ]}
              />
              <P>可创建 <Mono>.env.local</Mono> 覆盖本地配置（已被 .gitignore 忽略）。Vite 只会暴露 <Mono>VITE_</Mono> 前缀的变量到客户端代码。</P>
            </Section>

          </div>
        </div>
      </div>
    </div>
  )
}

/* ---- Sub-components ---- */

function Section({ id, title, children }: { id: string; title: string; children: React.ReactNode }) {
  return (
    <section
      id={id}
      className="bg-white rounded-xl p-6 sm:p-8"
      style={{ border: '1px solid rgba(0,0,0,.06)', boxShadow: '0 1px 3px rgba(0,0,0,.04)' }}
    >
      <h2 className="text-xl font-bold text-gray-900 mb-4">{title}</h2>
      <div className="space-y-4">{children}</div>
    </section>
  )
}

function H3({ children }: { children: React.ReactNode }) {
  return <h3 className="text-[15px] font-semibold text-gray-900 mt-2">{children}</h3>
}

function P({ children }: { children: React.ReactNode }) {
  return <p className="text-[14px] text-gray-600 leading-relaxed">{children}</p>
}

function Mono({ children }: { children: React.ReactNode }) {
  return (
    <code
      className="text-[13px] font-mono px-1.5 py-0.5 rounded"
      style={{ background: '#f1f5f9', color: '#334155' }}
    >
      {children}
    </code>
  )
}

function Code({ children }: { children: React.ReactNode }) {
  return (
    <pre
      className="text-[13px] font-mono leading-relaxed overflow-x-auto rounded-lg p-4"
      style={{ background: '#0f172a', color: '#e2e8f0' }}
    >
      <code>{children}</code>
    </pre>
  )
}

function Table({ headers, rows }: { headers: string[]; rows: string[][] }) {
  return (
    <div className="overflow-x-auto rounded-lg" style={{ border: '1px solid rgba(0,0,0,.08)' }}>
      <table className="w-full text-[13px]" style={{ borderCollapse: 'collapse' }}>
        <thead>
          <tr style={{ background: '#f8fafc' }}>
            {headers.map((h) => (
              <th key={h} className="text-left px-4 py-2.5 font-semibold text-gray-700" style={{ borderBottom: '1px solid rgba(0,0,0,.08)' }}>
                {h}
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {rows.map((row, i) => (
            <tr key={i}>
              {row.map((cell, j) => (
                <td key={j} className="px-4 py-2.5 text-gray-600" style={{ borderBottom: i < rows.length - 1 ? '1px solid rgba(0,0,0,.05)' : undefined }}>
                  {j === 0 ? <Mono>{cell}</Mono> : cell}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}
