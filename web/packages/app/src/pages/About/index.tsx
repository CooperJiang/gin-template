import { Link } from 'react-router-dom'

const techStack = [
  {
    name: 'React 19',
    desc: '声明式 UI 框架，Concurrent 特性',
    icon: (
      <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round">
        <circle cx="12" cy="12" r="10" /><circle cx="12" cy="12" r="3" />
      </svg>
    ),
  },
  {
    name: 'TypeScript 5',
    desc: '完整的类型安全与智能提示',
    icon: (
      <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round">
        <polyline points="16 18 22 12 16 6" /><polyline points="8 6 2 12 8 18" />
      </svg>
    ),
  },
  {
    name: 'TailwindCSS',
    desc: '实用优先的原子化 CSS 框架',
    icon: (
      <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round">
        <path d="M12 2L2 7l10 5 10-5-10-5z" /><path d="M2 17l10 5 10-5" /><path d="M2 12l10 5 10-5" />
      </svg>
    ),
  },
  {
    name: 'Vite',
    desc: '极速的开发服务器与构建工具',
    icon: (
      <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round">
        <polygon points="13 2 3 14 12 14 11 22 21 10 12 10 13 2" />
      </svg>
    ),
  },
  {
    name: 'qiankun',
    desc: '成熟的微前端框架，技术栈无关',
    icon: (
      <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round">
        <rect x="2" y="3" width="20" height="14" rx="2" ry="2" /><line x1="8" y1="21" x2="16" y2="21" /><line x1="12" y1="17" x2="12" y2="21" />
      </svg>
    ),
  },
  {
    name: 'Go + Gin',
    desc: '高性能后端，embed 静态资源',
    icon: (
      <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round">
        <path d="M20 16V7a2 2 0 00-2-2H6a2 2 0 00-2 2v9m16 0H4m16 0l1.28 2.55A1 1 0 0120.39 21H3.61a1 1 0 01-.89-1.45L4 16" />
      </svg>
    ),
  },
]

const archLayers = [
  {
    label: '主基座 (Main)',
    desc: '路由分发、导航栏、认证页面、qiankun 注册与生命周期管理',
    color: '#3b82f6',
  },
  {
    label: '子应用 (Sub Apps)',
    desc: '独立开发部署的业务模块，通过 activeRule 匹配路由自动加载',
    color: '#6366f1',
  },
  {
    label: '共享层 (Shared / Core / Auth)',
    desc: '类型定义、HTTP 客户端、安全存储、认证状态、UI 组件等公共能力',
    color: '#8b5cf6',
  },
  {
    label: 'Go 后端',
    desc: 'Gin 框架，RESTful API，embed 静态资源，单二进制部署',
    color: '#a78bfa',
  },
]

export default function About() {
  return (
    <div className="bg-gray-50">
      {/* Header */}
      <section className="relative overflow-hidden">
        <div
          className="absolute inset-0"
          style={{
            background: 'linear-gradient(135deg, #eff6ff 0%, #eef2ff 50%, #f5f3ff 100%)',
          }}
        />
        <div
          className="absolute top-0 left-1/2 opacity-20"
          style={{
            width: 600,
            height: 600,
            borderRadius: '50%',
            background: 'radial-gradient(circle, #818cf8 0%, transparent 70%)',
            transform: 'translate(-50%, -60%)',
          }}
        />
        <div className="relative max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-16 sm:py-20 text-center">
          <h1
            className="text-3xl sm:text-4xl font-bold text-gray-900 mb-4 tracking-tight"
          >
            关于这个项目
          </h1>
          <p className="text-lg text-gray-500 max-w-2xl mx-auto leading-relaxed">
            一个基于现代技术栈构建的全栈微前端模板，为开发者提供从开发到部署的完整解决方案。
          </p>
        </div>
      </section>

      {/* Architecture */}
      <section className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-14">
        <h2 className="text-xl sm:text-2xl font-bold text-gray-900 mb-2">系统架构</h2>
        <p className="text-gray-500 mb-8 text-[15px]">自上而下的分层架构，各层职责清晰、松耦合</p>
        <div className="space-y-3">
          {archLayers.map((layer, i) => (
            <div
              key={layer.label}
              className="flex items-start gap-4 bg-white rounded-xl p-5"
              style={{
                border: '1px solid rgba(0,0,0,.06)',
                boxShadow: '0 1px 3px rgba(0,0,0,.04)',
              }}
            >
              <div
                className="shrink-0 w-9 h-9 rounded-lg flex items-center justify-center text-white text-sm font-bold"
                style={{ background: layer.color }}
              >
                {i + 1}
              </div>
              <div className="min-w-0">
                <div className="text-[15px] font-semibold text-gray-900">{layer.label}</div>
                <div className="text-[13px] text-gray-500 mt-0.5">{layer.desc}</div>
              </div>
            </div>
          ))}
        </div>
      </section>

      {/* Tech Stack */}
      <section
        className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 pb-14"
      >
        <h2 className="text-xl sm:text-2xl font-bold text-gray-900 mb-2">技术栈</h2>
        <p className="text-gray-500 mb-8 text-[15px]">精选成熟稳定的技术组合</p>
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
          {techStack.map((t) => (
            <div
              key={t.name}
              className="flex items-start gap-3.5 bg-white rounded-xl p-5"
              style={{
                border: '1px solid rgba(0,0,0,.06)',
                boxShadow: '0 1px 3px rgba(0,0,0,.04)',
              }}
            >
              <div
                className="shrink-0 w-10 h-10 rounded-lg flex items-center justify-center text-white"
                style={{ background: 'linear-gradient(135deg, #3b82f6, #6366f1)' }}
              >
                {t.icon}
              </div>
              <div className="min-w-0">
                <div className="text-[15px] font-semibold text-gray-900">{t.name}</div>
                <div className="text-[13px] text-gray-500 mt-0.5">{t.desc}</div>
              </div>
            </div>
          ))}
        </div>
      </section>

      {/* Features */}
      <section className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 pb-14">
        <h2 className="text-xl sm:text-2xl font-bold text-gray-900 mb-2">项目特性</h2>
        <p className="text-gray-500 mb-8 text-[15px]">开箱即用的能力，覆盖常见业务场景</p>
        <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
          {[
            { title: '完整认证流程', desc: '登录、注册、邮箱验证码、密码重置，安全存储 Token' },
            { title: '微前端子应用管理', desc: '子应用独立开发部署，主应用动态加载，CSS 隔离' },
            { title: '响应式设计', desc: '移动端优先，自适应桌面端，汉堡菜单与侧边栏' },
            { title: '子应用模板', desc: '提供标准化模板，一键创建新子应用，统一架构规范' },
            { title: '共享包体系', desc: 'shared 类型定义、core 工具库、auth 认证包，monorepo 共享' },
            { title: '灵活部署', desc: 'Go embed 单二进制 / Nginx 反向代理，两种方案自由选择' },
          ].map((f) => (
            <div
              key={f.title}
              className="flex items-start gap-3 bg-white rounded-xl p-5"
              style={{
                border: '1px solid rgba(0,0,0,.06)',
                boxShadow: '0 1px 3px rgba(0,0,0,.04)',
              }}
            >
              <div className="shrink-0 mt-0.5">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" strokeWidth="2.2" strokeLinecap="round" strokeLinejoin="round">
                  <path d="M20 6L9 17l-5-5" stroke="#3b82f6" />
                </svg>
              </div>
              <div>
                <div className="text-[15px] font-semibold text-gray-900">{f.title}</div>
                <div className="text-[13px] text-gray-500 mt-0.5">{f.desc}</div>
              </div>
            </div>
          ))}
        </div>
      </section>

      {/* CTA */}
      <section className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 pb-16">
        <div
          className="rounded-2xl p-8 sm:p-10 text-center text-white"
          style={{
            background: 'linear-gradient(135deg, #3b82f6, #6366f1)',
          }}
        >
          <h2 className="text-xl sm:text-2xl font-bold mb-3">开始使用</h2>
          <p className="text-blue-100 mb-6 text-[15px] max-w-lg mx-auto">
            克隆仓库，运行 pnpm install && pnpm dev，即可开始开发你的微前端应用。
          </p>
          <Link
            to="/"
            className="inline-flex items-center gap-2 px-5 py-2.5 text-[14px] font-semibold rounded-lg no-underline transition-colors"
            style={{
              background: 'rgba(255,255,255,.95)',
              color: '#4f46e5',
            }}
          >
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
              <line x1="19" y1="12" x2="5" y2="12" /><polyline points="12 19 5 12 12 5" />
            </svg>
            返回首页
          </Link>
        </div>
      </section>
    </div>
  )
}
