import { Link } from 'react-router-dom'
import { useAuth } from '@app/auth'

const features = [
  {
    title: '现代技术栈',
    desc: '基于 React 19、TypeScript 5 和 TailwindCSS 构建，享受最新的开发体验与类型安全保障。',
    icon: (
      <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round">
        <polyline points="16 18 22 12 16 6" /><polyline points="8 6 2 12 8 18" />
      </svg>
    ),
  },
  {
    title: '安全认证体系',
    desc: '内置完整的用户认证系统，支持登录、注册、邮箱验证码、密码重置等功能。',
    icon: (
      <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round">
        <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z" />
      </svg>
    ),
  },
  {
    title: '微前端架构',
    desc: '基于 qiankun 的微前端方案，支持多应用独立开发、独立部署、运行时动态加载。',
    icon: (
      <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round">
        <rect x="2" y="3" width="20" height="14" rx="2" ry="2" /><line x1="8" y1="21" x2="16" y2="21" /><line x1="12" y1="17" x2="12" y2="21" />
      </svg>
    ),
  },
  {
    title: '工程化体系',
    desc: 'pnpm monorepo + Vite 构建，共享核心包与类型定义，统一代码规范与构建流程。',
    icon: (
      <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round">
        <path d="M22 19a2 2 0 01-2 2H4a2 2 0 01-2-2V5a2 2 0 012-2h5l2 3h9a2 2 0 012 2z" />
      </svg>
    ),
  },
  {
    title: '全栈部署',
    desc: 'Go 后端 embed 静态资源，单二进制部署；也支持 Nginx 反向代理的传统方案。',
    icon: (
      <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round">
        <path d="M20 16V7a2 2 0 00-2-2H6a2 2 0 00-2 2v9m16 0H4m16 0l1.28 2.55A1 1 0 0120.39 21H3.61a1 1 0 01-.89-1.45L4 16" />
      </svg>
    ),
  },
  {
    title: '开箱即用',
    desc: '提供子应用模板，一键创建新子应用，自带路由守卫、API 客户端、消息提示等基础设施。',
    icon: (
      <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round">
        <polyline points="20 12 20 22 4 22 4 12" /><rect x="2" y="7" width="20" height="5" /><line x1="12" y1="22" x2="12" y2="7" /><path d="M12 7H7.5a2.5 2.5 0 010-5C11 2 12 7 12 7z" /><path d="M12 7h4.5a2.5 2.5 0 000-5C13 2 12 7 12 7z" />
      </svg>
    ),
  },
]

const stats = [
  { value: 'React 19', label: '前端框架' },
  { value: 'Go 1.24', label: '后端语言' },
  { value: 'qiankun', label: '微前端方案' },
  { value: 'pnpm', label: '包管理器' },
]

export default function Home() {
  const { isAuthenticated } = useAuth()

  return (
    <div className="bg-gray-50">
      {/* Hero */}
      <section className="relative overflow-hidden">
        <div
          className="absolute inset-0"
          style={{
            background: 'linear-gradient(135deg, #eff6ff 0%, #eef2ff 50%, #f5f3ff 100%)',
          }}
        />
        {/* Decorative blobs */}
        <div
          className="absolute top-0 right-0 opacity-30"
          style={{
            width: 500,
            height: 500,
            borderRadius: '50%',
            background: 'radial-gradient(circle, #93c5fd 0%, transparent 70%)',
            transform: 'translate(30%, -40%)',
          }}
        />
        <div
          className="absolute bottom-0 left-0 opacity-20"
          style={{
            width: 400,
            height: 400,
            borderRadius: '50%',
            background: 'radial-gradient(circle, #a5b4fc 0%, transparent 70%)',
            transform: 'translate(-30%, 40%)',
          }}
        />

        <div className="relative max-w-5xl mx-auto px-4 sm:px-6 lg:px-8 py-20 sm:py-28">
          <div className="text-center">
            <div
              className="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full text-[13px] font-medium text-blue-700 mb-6"
              style={{ background: 'rgba(59,130,246,.1)', border: '1px solid rgba(59,130,246,.15)' }}
            >
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                <polygon points="13 2 3 14 12 14 11 22 21 10 12 10 13 2" />
              </svg>
              全栈微前端应用模板
            </div>
            <h1
              className="text-4xl sm:text-5xl font-bold text-gray-900 mb-5 tracking-tight"
              style={{ lineHeight: 1.15 }}
            >
              构建现代化的
              <span
                className="block mt-1"
                style={{
                  background: 'linear-gradient(135deg, #3b82f6, #6366f1)',
                  WebkitBackgroundClip: 'text',
                  WebkitTextFillColor: 'transparent',
                }}
              >
                企业级 Web 应用
              </span>
            </h1>
            <p className="text-lg text-gray-500 max-w-2xl mx-auto mb-10 leading-relaxed">
              基于 React + Go 的全栈微前端模板，提供完整的认证体系、子应用管理、
              工程化构建与一键部署能力，让你专注于业务开发。
            </p>
            <div className="flex items-center justify-center gap-3">
              <Link
                to="/about"
                className="inline-flex items-center gap-2 px-5 py-2.5 text-[14px] font-semibold text-white rounded-lg no-underline transition-all"
                style={{
                  background: 'linear-gradient(135deg, #3b82f6, #6366f1)',
                  boxShadow: '0 2px 8px rgba(99,102,241,.3)',
                }}
              >
                了解更多
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                  <line x1="5" y1="12" x2="19" y2="12" /><polyline points="12 5 19 12 12 19" />
                </svg>
              </Link>
              {isAuthenticated && (
                <Link
                  to="/profile"
                  className="inline-flex items-center gap-2 px-5 py-2.5 text-[14px] font-semibold text-gray-700 bg-white rounded-lg no-underline transition-colors hover:bg-gray-50"
                  style={{ boxShadow: '0 1px 3px rgba(0,0,0,.08)', border: '1px solid rgba(0,0,0,.08)' }}
                >
                  个人资料
                </Link>
              )}
            </div>
          </div>

          {/* Stats */}
          <div
            className="mt-16 grid grid-cols-2 sm:grid-cols-4 gap-4 max-w-3xl mx-auto"
          >
            {stats.map((s) => (
              <div
                key={s.label}
                className="text-center py-4 px-3 rounded-xl bg-white/70"
                style={{ border: '1px solid rgba(0,0,0,.05)', backdropFilter: 'blur(8px)' }}
              >
                <div className="text-lg font-bold text-gray-900">{s.value}</div>
                <div className="text-xs text-gray-500 mt-0.5">{s.label}</div>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Features */}
      <section className="max-w-5xl mx-auto px-4 sm:px-6 lg:px-8 py-16 sm:py-20">
        <div className="text-center mb-12">
          <h2 className="text-2xl sm:text-3xl font-bold text-gray-900 mb-3">核心能力</h2>
          <p className="text-gray-500 max-w-xl mx-auto">从开发到部署，提供完整的工程化解决方案</p>
        </div>
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-5">
          {features.map((f) => (
            <div
              key={f.title}
              className="group bg-white rounded-xl p-6 transition-all hover:-translate-y-0.5"
              style={{
                border: '1px solid rgba(0,0,0,.06)',
                boxShadow: '0 1px 3px rgba(0,0,0,.04)',
              }}
            >
              <div
                className="w-10 h-10 rounded-lg flex items-center justify-center text-white mb-4"
                style={{ background: 'linear-gradient(135deg, #3b82f6, #6366f1)' }}
              >
                {f.icon}
              </div>
              <h3 className="text-[15px] font-semibold text-gray-900 mb-2">{f.title}</h3>
              <p className="text-[13px] text-gray-500 leading-relaxed">{f.desc}</p>
            </div>
          ))}
        </div>
      </section>
    </div>
  )
}
