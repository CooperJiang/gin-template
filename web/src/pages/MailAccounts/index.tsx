import { useEffect, useMemo, useRef, useState, type ChangeEvent } from 'react'
import { Copy, Check, Loader2, Mail, Search, Upload, Plus, Pencil, KeyRound, Trash2, ChevronDown, AlertTriangle } from 'lucide-react'
import { mailAccountsApi } from '@/api/mailAccounts'
import { Button } from '@/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog'
import { INPUT_CLASS } from '@/lib/constants'
import { message } from '@/lib/message'
import { cn } from '@/lib/utils'
import type {
  CreateMailAccountRequest,
  FetchCodeResponse,
  MailAccount,
  MailAccountImportResponse,
  MailAccountListParams,
  MailAccountStatus,
  MailPreviewMessage,
  MailboxStatus,
  UpdateMailAccountRequest,
} from '@/types/mailAccount'

const PAGE_SIZE = 20
const STATUS_OPTIONS: { label: string; value: '' | MailAccountStatus }[] = [
  { label: '全部状态', value: '' },
  { label: '未使用', value: 'unused' },
  { label: '已使用', value: 'used' },
]
const MAILBOX_OPTIONS = [
  { label: '全部拉取状态', value: '' },
  { label: '待检测', value: 'unknown' },
  { label: '可拉取', value: 'ready' },
  { label: '拉取失败', value: 'fetch_failed' },
  { label: '令牌过期', value: 'token_expired' },
  { label: '待补全', value: 'incomplete' },
] as const
const PROVIDER_OPTIONS = [
  { label: '全部账号', value: '' },
  { label: 'Provider 完整', value: 'true' },
  { label: '待补全', value: 'false' },
] as const

type ImportDialogState = {
  open: boolean
  sourceType: 'paste' | 'file'
  content: string
  filename: string
  summary: MailAccountImportResponse | null
  submitting: boolean
}

type EditorState = {
  open: boolean
  mode: 'create' | 'edit'
  account: MailAccount | null
  submitting: boolean
  form: {
    email: string
    password: string
    client_id: string
    refresh_token: string
    status: MailAccountStatus
    remark: string
    tags: string
    folder: string
  }
}

type MailPreviewState = {
  open: boolean
  loading: boolean
  account: MailAccount | null
  messages: MailPreviewMessage[]
}

type FetchCodeState = {
  open: boolean
  loading: boolean
  account: MailAccount | null
  result: FetchCodeResponse | null
}

type HtmlPreviewState = {
  open: boolean
  subject: string
  html: string
}

function formatDateTime(value?: string | null) {
  if (!value) return '—'
  const date = new Date(value)
  if (Number.isNaN(date.getTime())) return value
  return date.toLocaleString('zh-CN', { hour12: false })
}

function statusLabel(status: MailAccountStatus) {
  switch (status) {
    case 'unused':
      return '未使用'
    case 'used':
      return '已使用'
  }
}

function mailboxLabel(status: MailAccount['mailbox_status']) {
  switch (status) {
    case 'ready':
      return '可拉取'
    case 'fetch_failed':
      return '拉取失败'
    case 'token_expired':
      return '令牌过期'
    case 'incomplete':
      return '待补全'
    default:
      return '待检测'
  }
}

const emptyEditorForm: EditorState['form'] = {
  email: '',
  password: '',
  client_id: '',
  refresh_token: '',
  status: 'unused',
  remark: '',
  tags: '',
  folder: 'inbox',
}

const SELECT_CLASS =
  'appearance-none h-8 rounded-md border border-slate-200 dark:border-slate-700 bg-white dark:bg-slate-800 pl-2.5 pr-7 text-sm text-slate-700 dark:text-slate-300 focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition cursor-pointer'

function CopyAccountButton({ email, password }: { email: string; password: string }) {
  const [copied, setCopied] = useState(false)

  const handleCopy = async () => {
    const text = `邮箱:${email} 密码:${password || '无'}`
    await navigator.clipboard.writeText(text)
    setCopied(true)
    setTimeout(() => setCopied(false), 1500)
  }

  return (
    <button
      type="button"
      onClick={() => void handleCopy()}
      className="inline-flex items-center justify-center rounded p-0.5 text-slate-400 transition hover:text-slate-600 dark:hover:text-slate-200"
      title="复制邮箱和密码"
    >
      {copied ? <Check className="h-3.5 w-3.5 text-emerald-500" /> : <Copy className="h-3.5 w-3.5" />}
    </button>
  )
}

function CopyButton({ text }: { text: string }) {
  const [copied, setCopied] = useState(false)

  const handleCopy = async () => {
    if (!text) return
    await navigator.clipboard.writeText(text)
    setCopied(true)
    setTimeout(() => setCopied(false), 1500)
  }

  return (
    <button
      type="button"
      onClick={() => void handleCopy()}
      className="inline-flex items-center justify-center rounded p-0.5 text-slate-400 transition hover:text-slate-600 dark:hover:text-slate-200"
      title="复制"
    >
      {copied ? <Check className="h-3.5 w-3.5 text-emerald-500" /> : <Copy className="h-3.5 w-3.5" />}
    </button>
  )
}

export default function MailAccounts() {
  const [filters, setFilters] = useState({
    keyword: '',
    status: '',
    mailbox_status: '',
    provider_ready: '',
    page: 1,
  })
  const [loading, setLoading] = useState(false)
  const [accounts, setAccounts] = useState<MailAccount[]>([])
  const [total, setTotal] = useState(0)
  const [totalPages, setTotalPages] = useState(1)
  const [importDialog, setImportDialog] = useState<ImportDialogState>({
    open: false,
    sourceType: 'paste',
    content: '',
    filename: '',
    summary: null,
    submitting: false,
  })
  const [editor, setEditor] = useState<EditorState>({
    open: false,
    mode: 'create',
    account: null,
    submitting: false,
    form: emptyEditorForm,
  })
  const [mailPreview, setMailPreview] = useState<MailPreviewState>({
    open: false,
    loading: false,
    account: null,
    messages: [],
  })
  const [fetchCode, setFetchCode] = useState<FetchCodeState>({
    open: false,
    loading: false,
    account: null,
    result: null,
  })
  const [htmlPreview, setHtmlPreview] = useState<HtmlPreviewState>({
    open: false,
    subject: '',
    html: '',
  })
  const fileInputRef = useRef<HTMLInputElement | null>(null)
  const [confirmStatusId, setConfirmStatusId] = useState<string | null>(null)
  const [confirmDeleteId, setConfirmDeleteId] = useState<string | null>(null)

  const listParams = useMemo<MailAccountListParams>(() => {
    const params: MailAccountListParams = {
      page: filters.page,
      size: PAGE_SIZE,
    }
    if (filters.keyword.trim()) params.keyword = filters.keyword.trim()
    if (filters.status) params.status = filters.status as MailAccountStatus
    if (filters.mailbox_status) params.mailbox_status = filters.mailbox_status as MailboxStatus
    if (filters.provider_ready === 'true') params.provider_ready = true
    if (filters.provider_ready === 'false') params.provider_ready = false
    return params
  }, [filters])

  const loadAccounts = async () => {
    setLoading(true)
    try {
      const response = await mailAccountsApi.list(listParams)
      setAccounts(response.items)
      setTotal(response.total)
      setTotalPages(response.total_pages || 1)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    void loadAccounts()
  }, [listParams])

  const openCreateDialog = () => {
    setEditor({
      open: true,
      mode: 'create',
      account: null,
      submitting: false,
      form: emptyEditorForm,
    })
  }

  const openEditDialog = (account: MailAccount) => {
    setEditor({
      open: true,
      mode: 'edit',
      account,
      submitting: false,
      form: {
        email: account.email,
        password: '',
        client_id: account.client_id,
        refresh_token: '',
        status: account.status,
        remark: account.remark,
        tags: account.tags,
        folder: account.folder || 'inbox',
      },
    })
  }

  const handleImportSubmit = async () => {
    if (!importDialog.content.trim()) {
      message.warning('请先粘贴内容或选择文件')
      return
    }

    setImportDialog((current) => ({ ...current, submitting: true }))
    try {
      const summary = await mailAccountsApi.importAccounts({
        content: importDialog.content,
        source_type: importDialog.sourceType,
        filename: importDialog.filename || undefined,
        mode: 'fill_missing',
      })
      setImportDialog((current) => ({ ...current, summary, submitting: false }))
      message.success('导入任务已完成')
      void loadAccounts()
    } catch {
      setImportDialog((current) => ({ ...current, submitting: false }))
    }
  }

  const handleEditorSubmit = async () => {
    if (!editor.form.email.trim() && editor.mode === 'create') {
      message.warning('邮箱不能为空')
      return
    }

    const payload: CreateMailAccountRequest | UpdateMailAccountRequest = {
      password: editor.form.password || undefined,
      client_id: editor.form.client_id || undefined,
      refresh_token: editor.form.refresh_token || undefined,
      status: editor.form.status,
      remark: editor.form.remark || undefined,
      tags: editor.form.tags || undefined,
      folder: editor.form.folder || undefined,
      ...(editor.mode === 'create' ? { email: editor.form.email.trim() } : {}),
    }

    setEditor((current) => ({ ...current, submitting: true }))
    try {
      if (editor.mode === 'create') {
        await mailAccountsApi.create(payload as CreateMailAccountRequest)
        message.success('邮箱账号已新增')
      } else if (editor.account) {
        await mailAccountsApi.update(editor.account.id, payload as UpdateMailAccountRequest)
        message.success('邮箱账号已更新')
      }

      setEditor((current) => ({ ...current, open: false, submitting: false }))
      void loadAccounts()
    } catch {
      setEditor((current) => ({ ...current, submitting: false }))
    }
  }

  const handleFetchMails = async (account: MailAccount) => {
    setMailPreview({ open: true, loading: true, account, messages: [] })
    try {
      const response = await mailAccountsApi.fetchMails(account.id)
      setMailPreview({ open: true, loading: false, account, messages: response.messages })
      message.success('邮件预览已更新')
      void loadAccounts()
    } catch {
      setMailPreview((current) => ({ ...current, loading: false }))
    }
  }

  const handleFetchCode = async (account: MailAccount) => {
    setFetchCode({ open: true, loading: true, account, result: null })
    try {
      const result = await mailAccountsApi.fetchCode(account.id)
      setFetchCode({ open: true, loading: false, account, result })
      message.success('最新验证码已获取')
      void loadAccounts()
    } catch {
      setFetchCode((current) => ({ ...current, loading: false }))
    }
  }

  const handleFileSelected = async (event: ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0]
    if (!file) return

    const content = await file.text()
    setImportDialog((current) => ({
      ...current,
      sourceType: 'file',
      content,
      filename: file.name,
      summary: null,
    }))
    event.target.value = ''
  }

  const handleDelete = async (account: MailAccount) => {
    try {
      await mailAccountsApi.delete(account.id)
      message.success('删除成功')
      setConfirmDeleteId(null)
      void loadAccounts()
    } catch {
      setConfirmDeleteId(null)
    }
  }

  const handleToggleStatus = async (account: MailAccount) => {
    const newStatus: MailAccountStatus = account.status === 'unused' ? 'used' : 'unused'
    try {
      await mailAccountsApi.updateStatus(account.id, { status: newStatus })
      message.success(`已切换为${newStatus === 'unused' ? '未使用' : '已使用'}`)
      setConfirmStatusId(null)
      void loadAccounts()
    } catch {
      setConfirmStatusId(null)
    }
  }

  return (
    <div className="space-y-4">
      {/* 页面标题 */}
      <div>
        <h1 className="text-xl font-semibold text-slate-900 dark:text-slate-100">邮箱管理</h1>
        <p className="mt-1 text-sm text-slate-500 dark:text-slate-400">管理您的所有邮箱账号</p>
      </div>

      {/* 筛选 + 操作：一行 */}
      <div className="flex flex-wrap items-center gap-2">
        <label className="relative">
          <Search className="pointer-events-none absolute left-2.5 top-1/2 h-4 w-4 -translate-y-1/2 text-gray-400" />
          <input
            value={filters.keyword}
            onChange={(event) =>
              setFilters((current) => ({ ...current, keyword: event.target.value, page: 1 }))
            }
            className={cn(INPUT_CLASS, 'h-8 w-52 pl-8 text-sm')}
            placeholder="搜索邮箱地址..."
          />
        </label>

        {/* 状态单选按钮组 */}
        <div className="flex items-center rounded-md border border-slate-200 dark:border-slate-700 overflow-hidden h-8">
          {STATUS_OPTIONS.map((option) => (
            <button
              key={option.label}
              type="button"
              onClick={() => setFilters((current) => ({ ...current, status: option.value, page: 1 }))}
              className={cn(
                'px-3 h-full text-xs font-medium transition-colors border-r last:border-r-0 border-slate-200 dark:border-slate-700',
                filters.status === option.value
                  ? 'bg-blue-500 text-white'
                  : 'bg-white dark:bg-slate-800 text-slate-600 dark:text-slate-400 hover:bg-slate-50 dark:hover:bg-slate-750',
              )}
            >
              {option.label}
            </button>
          ))}
        </div>

        <div className="relative">
          <select
            value={filters.mailbox_status}
            onChange={(event) =>
              setFilters((current) => ({ ...current, mailbox_status: event.target.value, page: 1 }))
            }
            className={SELECT_CLASS}
          >
            {MAILBOX_OPTIONS.map((option) => (
              <option key={option.label} value={option.value}>
                {option.label}
              </option>
            ))}
          </select>
          <ChevronDown className="pointer-events-none absolute right-2 top-1/2 h-3.5 w-3.5 -translate-y-1/2 text-slate-400" />
        </div>
        <div className="relative">
          <select
            value={filters.provider_ready}
            onChange={(event) =>
              setFilters((current) => ({ ...current, provider_ready: event.target.value, page: 1 }))
            }
            className={SELECT_CLASS}
          >
            {PROVIDER_OPTIONS.map((option) => (
              <option key={option.label} value={option.value}>
                {option.label}
              </option>
            ))}
          </select>
          <ChevronDown className="pointer-events-none absolute right-2 top-1/2 h-3.5 w-3.5 -translate-y-1/2 text-slate-400" />
        </div>

        <div className="ml-auto flex items-center gap-2">
          <Button
            size="sm"
            variant="outline"
            onClick={() =>
              setImportDialog({
                open: true,
                sourceType: 'paste',
                content: '',
                filename: '',
                summary: null,
                submitting: false,
              })
            }
          >
            <Upload className="h-3.5 w-3.5" />
            导入
          </Button>
          <Button size="sm" onClick={openCreateDialog}>
            <Plus className="h-3.5 w-3.5" />
            新增
          </Button>
        </div>
      </div>

      {/* 卡片网格列表 */}
      {loading ? (
        <div className="flex justify-center py-16">
          <Loader2 className="h-6 w-6 animate-spin text-slate-400" />
        </div>
      ) : accounts.length === 0 ? (
        <div className="rounded-lg border border-dashed border-slate-300 dark:border-slate-700 py-16 text-center text-slate-400 text-sm">
          暂无数据
        </div>
      ) : (
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
          {accounts.map((account) => (
            <div
              key={account.id}
              className={cn(
                'rounded-lg border shadow-sm flex flex-col',
                account.status === 'used'
                  ? 'border-red-200 bg-red-50/40 dark:border-red-900/50 dark:bg-red-950/20'
                  : 'border-slate-200 bg-white dark:border-slate-700 dark:bg-slate-900',
              )}
            >
              {/* Card top: email + clickable status badge + copy */}
              <div className={cn(
                'px-4 pt-4 pb-3 border-b',
                account.status === 'used'
                  ? 'border-red-100 dark:border-red-900/30'
                  : 'border-slate-100 dark:border-slate-800',
              )}>
                <div className="flex items-start justify-between gap-2">
                  <div className="min-w-0 flex-1">
                    <div className="flex items-center gap-1.5">
                      <Mail className="h-3.5 w-3.5 shrink-0 text-slate-400" />
                      <span className="truncate text-sm font-medium text-slate-900 dark:text-slate-100">
                        {account.email}
                      </span>
                      <CopyAccountButton email={account.email} password={account.password} />
                    </div>
                  </div>
                  {/* Status badge with popover confirm */}
                  <div className="relative shrink-0">
                    <button
                      type="button"
                      onClick={() => setConfirmStatusId(confirmStatusId === account.id ? null : account.id)}
                      title="点击切换状态"
                      className={cn(
                        'inline-block rounded px-2 py-0.5 text-xs font-medium cursor-pointer transition-colors',
                        account.status === 'unused'
                          ? 'bg-emerald-50 text-emerald-600 hover:bg-emerald-100 dark:bg-emerald-500/15 dark:text-emerald-400 dark:hover:bg-emerald-500/25'
                          : 'bg-red-100 text-red-600 hover:bg-red-200 dark:bg-red-500/20 dark:text-red-400 dark:hover:bg-red-500/30',
                      )}
                    >
                      {statusLabel(account.status)}
                    </button>
                    {confirmStatusId === account.id && (
                      <>
                        <div className="fixed inset-0 z-40" onClick={() => setConfirmStatusId(null)} />
                        <div className="absolute right-0 top-full mt-1.5 z-50 w-52 rounded-lg border border-slate-200 dark:border-slate-700 bg-white dark:bg-slate-800 shadow-lg p-3">
                          <p className="text-xs text-slate-600 dark:text-slate-300 mb-2.5">
                            确定将状态切换为「{account.status === 'unused' ? '已使用' : '未使用'}」吗？
                          </p>
                          <div className="flex items-center justify-end gap-1.5">
                            <Button
                              size="sm"
                              variant="outline"
                              className="h-6 px-2 text-xs"
                              onClick={() => setConfirmStatusId(null)}
                            >
                              取消
                            </Button>
                            <Button
                              size="sm"
                              className="h-6 px-2 text-xs"
                              onClick={() => void handleToggleStatus(account)}
                            >
                              确认
                            </Button>
                          </div>
                        </div>
                      </>
                    )}
                  </div>
                </div>
              </div>

              {/* Card middle: details */}
              <div className="px-4 py-3 space-y-2 text-sm flex-1">
                {/* Password */}
                <div className="flex items-center gap-2">
                  <span className="text-slate-400 text-xs w-10 shrink-0">密码</span>
                  <span className="font-mono text-slate-700 dark:text-slate-300 truncate">
                    {account.password || '—'}
                  </span>
                </div>
                {/* Token status */}
                <div className="flex items-center gap-2">
                  <span className="text-slate-400 text-xs w-10 shrink-0">令牌</span>
                  <span
                    className={cn(
                      'inline-block rounded px-1.5 py-0.5 text-xs font-medium',
                      account.refresh_token_configured
                        ? 'bg-emerald-50 text-emerald-600 dark:bg-emerald-500/15 dark:text-emerald-400'
                        : 'bg-orange-50 text-orange-500 dark:bg-orange-500/15 dark:text-orange-400',
                    )}
                  >
                    {account.refresh_token_configured ? '已配置' : '未配置'}
                  </span>
                </div>
                {/* Mailbox status */}
                <div className="flex items-center gap-2">
                  <span className="text-slate-400 text-xs w-10 shrink-0">拉取</span>
                  <span
                    className={cn(
                      'inline-block rounded px-1.5 py-0.5 text-xs font-medium',
                      account.mailbox_status === 'ready'
                        ? 'bg-cyan-50 text-cyan-600 dark:bg-cyan-500/15 dark:text-cyan-400'
                        : account.mailbox_status === 'fetch_failed' || account.mailbox_status === 'token_expired'
                          ? 'bg-rose-50 text-rose-500 dark:bg-rose-500/15 dark:text-rose-400'
                          : 'bg-slate-100 text-slate-500 dark:bg-slate-700 dark:text-slate-400',
                    )}
                  >
                    {mailboxLabel(account.mailbox_status)}
                  </span>
                </div>
                <div className="flex items-center gap-2">
                  <span className="text-slate-400 text-xs w-10 shrink-0">使用</span>
                  <span className="text-slate-700 dark:text-slate-300 truncate">
                    {formatDateTime(account.allocated_at)}
                  </span>
                </div>
              </div>

              {/* Card bottom: actions */}
              <div className={cn(
                'px-3 py-2.5 border-t flex items-center gap-1',
                account.status === 'used'
                  ? 'border-red-100 dark:border-red-900/30'
                  : 'border-slate-100 dark:border-slate-800',
              )}>
                <Button
                  size="sm"
                  className="h-7 px-2 text-xs flex-1"
                  onClick={() => void handleFetchCode(account)}
                >
                  <KeyRound className="h-3 w-3" />
                  取码
                </Button>
                <Button
                  size="sm"
                  variant="outline"
                  className="h-7 px-2 text-xs flex-1"
                  onClick={() => void handleFetchMails(account)}
                >
                  <Mail className="h-3 w-3" />
                  邮件
                </Button>
                <Button
                  size="sm"
                  variant="outline"
                  className="h-7 px-2 text-xs"
                  onClick={() => openEditDialog(account)}
                >
                  <Pencil className="h-3 w-3" />
                </Button>
                {/* Delete button with popover confirm */}
                <div className="relative">
                  <Button
                    size="sm"
                    variant="outline"
                    className="h-7 px-2 text-xs text-red-500 hover:bg-red-50 hover:text-red-600 dark:text-red-400 dark:hover:bg-red-950/30 dark:hover:text-red-300"
                    onClick={() => setConfirmDeleteId(confirmDeleteId === account.id ? null : account.id)}
                  >
                    <Trash2 className="h-3 w-3" />
                  </Button>
                  {confirmDeleteId === account.id && (
                    <>
                      <div className="fixed inset-0 z-40" onClick={() => setConfirmDeleteId(null)} />
                      <div className="absolute right-0 bottom-full mb-1.5 z-50 w-56 rounded-lg border border-red-200 dark:border-red-900/50 bg-white dark:bg-slate-800 shadow-lg p-3">
                        <div className="flex items-start gap-2 mb-2.5">
                          <AlertTriangle className="h-4 w-4 text-red-500 shrink-0 mt-0.5" />
                          <p className="text-xs text-slate-600 dark:text-slate-300">
                            确定删除 <span className="font-medium text-slate-900 dark:text-slate-100">{account.email}</span> 吗？此操作不可撤销。
                          </p>
                        </div>
                        <div className="flex items-center justify-end gap-1.5">
                          <Button
                            size="sm"
                            variant="outline"
                            className="h-6 px-2 text-xs"
                            onClick={() => setConfirmDeleteId(null)}
                          >
                            取消
                          </Button>
                          <Button
                            size="sm"
                            className="h-6 px-2 text-xs bg-red-500 hover:bg-red-600 text-white"
                            onClick={() => void handleDelete(account)}
                          >
                            删除
                          </Button>
                        </div>
                      </div>
                    </>
                  )}
                </div>
              </div>
            </div>
          ))}
        </div>
      )}

      {/* 分页 */}
      <div className="flex items-center justify-between text-sm">
        <span className="text-slate-400">
          共 {total} 条 · 第 {filters.page}/{Math.max(totalPages, 1)} 页
        </span>
        <div className="flex items-center gap-1.5">
          <Button
            size="sm"
            variant="outline"
            className="h-7 px-3 text-xs"
            disabled={filters.page <= 1}
            onClick={() => setFilters((current) => ({ ...current, page: current.page - 1 }))}
          >
            上一页
          </Button>
          <Button
            size="sm"
            variant="outline"
            className="h-7 px-3 text-xs"
            disabled={filters.page >= Math.max(totalPages, 1)}
            onClick={() => setFilters((current) => ({ ...current, page: current.page + 1 }))}
          >
            下一页
          </Button>
        </div>
      </div>

      {/* 导入弹窗 */}
      <Dialog
        open={importDialog.open}
        onOpenChange={(open) => setImportDialog((current) => ({ ...current, open }))}
      >
        <DialogContent className="max-w-3xl">
          <DialogHeader>
            <DialogTitle>导入邮箱账号</DialogTitle>
            <DialogDescription>
              支持粘贴导入和本地文件读取导入。文件内容不会走单独上传存储，只会在浏览器读取后以文本形式提交。
            </DialogDescription>
          </DialogHeader>
          <div className="space-y-4">
            <div className="flex flex-wrap gap-2">
              <Button
                variant={importDialog.sourceType === 'paste' ? 'default' : 'outline'}
                onClick={() =>
                  setImportDialog((current) => ({
                    ...current,
                    sourceType: 'paste',
                    filename: '',
                    summary: null,
                  }))
                }
              >
                粘贴导入
              </Button>
              <Button
                variant={importDialog.sourceType === 'file' ? 'default' : 'outline'}
                onClick={() => {
                  setImportDialog((current) => ({ ...current, sourceType: 'file', summary: null }))
                  fileInputRef.current?.click()
                }}
              >
                选择文件
              </Button>
              <input
                ref={fileInputRef}
                type="file"
                accept=".txt,.csv"
                className="hidden"
                onChange={(event) => void handleFileSelected(event)}
              />
            </div>

            <div className="rounded-lg border border-dashed border-slate-300 bg-slate-50/70 p-4 dark:border-slate-700 dark:bg-slate-900/40">
              <div className="mb-2 flex items-center justify-between text-sm text-slate-500 dark:text-slate-400">
                <span>导入内容</span>
                <span>{importDialog.filename ? `文件：${importDialog.filename}` : '支持 3 种账号格式'}</span>
              </div>
              <textarea
                rows={10}
                value={importDialog.content}
                onChange={(event) =>
                  setImportDialog((current) => ({
                    ...current,
                    content: event.target.value,
                    summary: null,
                  }))
                }
                placeholder={
                  '邮箱----密码----ID----token\n邮箱<TAB>密码<TAB>id<TAB>token\n邮箱----密码'
                }
                className={cn(INPUT_CLASS, 'min-h-[200px] resize-y font-mono text-xs leading-6')}
              />
            </div>

            {importDialog.summary && (
              <div className="grid gap-4 lg:grid-cols-[1fr_1.2fr]">
                <div className="rounded-lg border border-emerald-200 bg-emerald-50/70 p-4 dark:border-emerald-900/60 dark:bg-emerald-950/20">
                  <div className="text-sm font-medium text-emerald-700 dark:text-emerald-300">
                    导入批次 {importDialog.summary.batch_no}
                  </div>
                  <div className="mt-3 grid grid-cols-2 gap-3 text-sm text-emerald-900 dark:text-emerald-100">
                    <div>
                      <div className="text-xs text-emerald-600/80 dark:text-emerald-300/70">总行数</div>
                      <div className="mt-1 text-xl font-semibold">{importDialog.summary.total_lines}</div>
                    </div>
                    <div>
                      <div className="text-xs text-emerald-600/80 dark:text-emerald-300/70">新增</div>
                      <div className="mt-1 text-xl font-semibold">{importDialog.summary.created_count}</div>
                    </div>
                    <div>
                      <div className="text-xs text-emerald-600/80 dark:text-emerald-300/70">补空更新</div>
                      <div className="mt-1 text-xl font-semibold">{importDialog.summary.updated_count}</div>
                    </div>
                    <div>
                      <div className="text-xs text-emerald-600/80 dark:text-emerald-300/70">失败</div>
                      <div className="mt-1 text-xl font-semibold">{importDialog.summary.failed_count}</div>
                    </div>
                  </div>
                </div>
                <div className="rounded-lg border border-slate-200 bg-white p-4 dark:border-slate-700 dark:bg-slate-900/50">
                  <div className="mb-2 text-sm font-medium text-slate-900 dark:text-slate-100">失败明细</div>
                  <div className="max-h-44 space-y-2 overflow-auto pr-1 text-xs text-slate-600 dark:text-slate-300">
                    {importDialog.summary.errors.length === 0 ? (
                      <div className="rounded bg-slate-50 px-3 py-2 dark:bg-slate-800/70">本次没有失败行。</div>
                    ) : (
                      importDialog.summary.errors.map((item) => (
                        <div key={`${item.line}-${item.reason}`} className="rounded bg-slate-50 px-3 py-2 dark:bg-slate-800/70">
                          <div className="font-medium text-slate-900 dark:text-slate-100">
                            第 {item.line} 行 · {item.reason}
                          </div>
                          <div className="mt-1 break-all font-mono text-[11px] text-slate-500 dark:text-slate-400">
                            {item.raw_line}
                          </div>
                        </div>
                      ))
                    )}
                  </div>
                </div>
              </div>
            )}
          </div>
          <DialogFooter>
            <Button variant="outline" onClick={() => setImportDialog((current) => ({ ...current, open: false }))}>
              关闭
            </Button>
            <Button disabled={importDialog.submitting} onClick={() => void handleImportSubmit()}>
              {importDialog.submitting && <Loader2 className="h-4 w-4 animate-spin" />}
              执行导入
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      {/* 编辑弹窗 */}
      <Dialog open={editor.open} onOpenChange={(open) => setEditor((current) => ({ ...current, open }))}>
        <DialogContent className="max-w-2xl">
          <DialogHeader>
            <DialogTitle>{editor.mode === 'create' ? '新增邮箱账号' : '编辑邮箱账号'}</DialogTitle>
            <DialogDescription>
              {editor.mode === 'create'
                ? '手工录入单个邮箱账号，支持直接补齐 provider 凭据。'
                : '编辑时不会回显明文密码和 refresh token，如需更新请重新输入。'}
            </DialogDescription>
          </DialogHeader>
          <div className="grid gap-4 md:grid-cols-2">
            <label className="space-y-2 md:col-span-2">
              <span className="text-sm font-medium text-slate-700 dark:text-slate-300">邮箱</span>
              <input
                value={editor.form.email}
                disabled={editor.mode === 'edit'}
                onChange={(event) =>
                  setEditor((current) => ({
                    ...current,
                    form: { ...current.form, email: event.target.value },
                  }))
                }
                className={INPUT_CLASS}
                placeholder="name@outlook.com"
              />
            </label>
            <label className="space-y-2">
              <span className="text-sm font-medium text-slate-700 dark:text-slate-300">邮箱密码</span>
              <input
                type="password"
                value={editor.form.password}
                onChange={(event) =>
                  setEditor((current) => ({
                    ...current,
                    form: { ...current.form, password: event.target.value },
                  }))
                }
                className={INPUT_CLASS}
                placeholder={editor.mode === 'edit' ? '留空则不更新' : '请输入邮箱密码'}
              />
            </label>
            <label className="space-y-2">
              <span className="text-sm font-medium text-slate-700 dark:text-slate-300">状态</span>
              <select
                value={editor.form.status}
                onChange={(event) =>
                  setEditor((current) => ({
                    ...current,
                    form: { ...current.form, status: event.target.value as MailAccountStatus },
                  }))
                }
                className={INPUT_CLASS}
              >
                {STATUS_OPTIONS.filter((option) => option.value).map((option) => (
                  <option key={option.value} value={option.value}>
                    {option.label}
                  </option>
                ))}
              </select>
            </label>
            <label className="space-y-2">
              <span className="text-sm font-medium text-slate-700 dark:text-slate-300">Client ID</span>
              <input
                value={editor.form.client_id}
                onChange={(event) =>
                  setEditor((current) => ({
                    ...current,
                    form: { ...current.form, client_id: event.target.value },
                  }))
                }
                className={INPUT_CLASS}
                placeholder="provider client id"
              />
            </label>
            <label className="space-y-2">
              <span className="text-sm font-medium text-slate-700 dark:text-slate-300">Folder</span>
              <input
                value={editor.form.folder}
                onChange={(event) =>
                  setEditor((current) => ({
                    ...current,
                    form: { ...current.form, folder: event.target.value },
                  }))
                }
                className={INPUT_CLASS}
                placeholder="inbox"
              />
            </label>
            <label className="space-y-2 md:col-span-2">
              <span className="text-sm font-medium text-slate-700 dark:text-slate-300">Refresh Token</span>
              <textarea
                rows={4}
                value={editor.form.refresh_token}
                onChange={(event) =>
                  setEditor((current) => ({
                    ...current,
                    form: { ...current.form, refresh_token: event.target.value },
                  }))
                }
                className={cn(INPUT_CLASS, 'resize-y font-mono text-xs')}
                placeholder={editor.mode === 'edit' ? '留空则不更新' : '请输入 refresh token'}
              />
            </label>
            <label className="space-y-2">
              <span className="text-sm font-medium text-slate-700 dark:text-slate-300">标签</span>
              <input
                value={editor.form.tags}
                onChange={(event) =>
                  setEditor((current) => ({
                    ...current,
                    form: { ...current.form, tags: event.target.value },
                  }))
                }
                className={INPUT_CLASS}
                placeholder="例如：chatgpt, internal"
              />
            </label>
            <label className="space-y-2">
              <span className="text-sm font-medium text-slate-700 dark:text-slate-300">备注</span>
              <input
                value={editor.form.remark}
                onChange={(event) =>
                  setEditor((current) => ({
                    ...current,
                    form: { ...current.form, remark: event.target.value },
                  }))
                }
                className={INPUT_CLASS}
                placeholder="记录用途或来源"
              />
            </label>
          </div>
          <DialogFooter>
            <Button variant="outline" onClick={() => setEditor((current) => ({ ...current, open: false }))}>
              取消
            </Button>
            <Button disabled={editor.submitting} onClick={() => void handleEditorSubmit()}>
              {editor.submitting && <Loader2 className="h-4 w-4 animate-spin" />}
              {editor.mode === 'create' ? '创建账号' : '保存修改'}
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      {/* 邮件预览弹窗 */}
      <Dialog
        open={mailPreview.open}
        onOpenChange={(open) => setMailPreview((current) => ({ ...current, open }))}
      >
        <DialogContent className="max-w-4xl">
          <DialogHeader>
            <DialogTitle>邮件预览</DialogTitle>
            <DialogDescription>
              {mailPreview.account ? `当前账号：${mailPreview.account.email}` : '查看标准化后的最近邮件列表'}
            </DialogDescription>
          </DialogHeader>
          <div className="max-h-[70vh] space-y-3 overflow-auto pr-1">
            {mailPreview.loading ? (
              <div className="flex min-h-[220px] items-center justify-center text-slate-500 dark:text-slate-400">
                <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                正在获取邮件预览...
              </div>
            ) : mailPreview.messages.length === 0 ? (
              <div className="rounded-lg border border-dashed border-slate-300 p-10 text-center text-slate-500 dark:border-slate-700 dark:text-slate-400">
                暂无邮件预览结果
              </div>
            ) : (
              mailPreview.messages.map((item) => (
                <div
                  key={item.id}
                  className="rounded-lg border border-slate-200 bg-white p-4 shadow-sm dark:border-slate-700 dark:bg-slate-900/60"
                >
                  <div className="flex flex-col gap-3 sm:flex-row sm:items-start sm:justify-between">
                    <div>
                      <div className="flex items-center gap-2 text-sm text-slate-500 dark:text-slate-400">
                        <Mail className="h-4 w-4" />
                        {item.from_name || item.from_address}
                        <span>·</span>
                        <span>{item.from_address}</span>
                      </div>
                      <div className="mt-2 text-base font-semibold text-slate-900 dark:text-slate-100">
                        {item.subject || '无主题'}
                      </div>
                    </div>
                    <div className="text-xs text-slate-500 dark:text-slate-400">
                      {formatDateTime(item.received_time)}
                    </div>
                  </div>
                  <div className="mt-3 rounded bg-slate-50 px-3 py-3 text-sm leading-6 text-slate-700 dark:bg-slate-800/80 dark:text-slate-300">
                    {item.body_preview || '无预览内容'}
                  </div>
                  {item.body && (
                    <div className="mt-2 flex justify-end">
                      <Button
                        size="sm"
                        variant="outline"
                        className="h-7 text-xs"
                        onClick={() =>
                          setHtmlPreview({
                            open: true,
                            subject: item.subject || '无主题',
                            html: item.body || '',
                          })
                        }
                      >
                        查看完整邮件
                      </Button>
                    </div>
                  )}
                </div>
              ))
            )}
          </div>
          <DialogFooter>
            <Button variant="outline" onClick={() => setMailPreview((current) => ({ ...current, open: false }))}>
              关闭
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      {/* HTML 邮件预览弹窗 */}
      <Dialog
        open={htmlPreview.open}
        onOpenChange={(open) => setHtmlPreview((current) => ({ ...current, open }))}
      >
        <DialogContent className="max-w-4xl h-[85vh] flex flex-col">
          <DialogHeader className="shrink-0">
            <DialogTitle className="flex items-center justify-between pr-8">
              <span className="truncate">{htmlPreview.subject}</span>
            </DialogTitle>
            <DialogDescription>HTML 格式邮件内容</DialogDescription>
          </DialogHeader>
          <div className="flex-1 min-h-0 rounded-lg border border-slate-200 dark:border-slate-700 overflow-hidden bg-white">
            <iframe
              srcDoc={htmlPreview.html}
              sandbox="allow-same-origin"
              title="邮件内容"
              className="w-full h-full border-0"
            />
          </div>
          <DialogFooter className="shrink-0">
            <Button
              variant="outline"
              onClick={() => setHtmlPreview((current) => ({ ...current, open: false }))}
            >
              关闭
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      {/* 验证码弹窗 */}
      <Dialog open={fetchCode.open} onOpenChange={(open) => setFetchCode((current) => ({ ...current, open }))}>
        <DialogContent className="max-w-2xl">
          <DialogHeader>
            <DialogTitle>最新验证码</DialogTitle>
            <DialogDescription>
              {fetchCode.account ? `当前账号：${fetchCode.account.email}` : '展示最近一次验证码提取结果'}
            </DialogDescription>
          </DialogHeader>
          {fetchCode.loading ? (
            <div className="flex min-h-[160px] items-center justify-center text-slate-500 dark:text-slate-400">
              <Loader2 className="mr-2 h-4 w-4 animate-spin" />
              正在提取验证码...
            </div>
          ) : fetchCode.result ? (
            <div className="space-y-4">
              <div className="rounded-lg border border-cyan-200 bg-cyan-50/60 p-5 dark:border-cyan-900/60 dark:bg-cyan-950/20">
                <div className="text-xs text-cyan-600 dark:text-cyan-400">验证码</div>
                <div className="mt-2 flex items-center gap-3">
                  <span className="font-mono text-3xl font-semibold tracking-[0.3em] text-slate-900 dark:text-white">
                    {fetchCode.result.code || '—'}
                  </span>
                  {fetchCode.result.code && <CopyButton text={fetchCode.result.code} />}
                </div>
              </div>
              <div className="grid gap-3 text-sm">
                <div className="flex gap-2">
                  <span className="w-16 shrink-0 text-slate-400">主题</span>
                  <span className="text-slate-800 dark:text-slate-200">{fetchCode.result.matched_subject || '—'}</span>
                </div>
                <div className="flex gap-2">
                  <span className="w-16 shrink-0 text-slate-400">发件人</span>
                  <span className="text-slate-800 dark:text-slate-200">{fetchCode.result.matched_from || '—'}</span>
                </div>
                <div className="flex gap-2">
                  <span className="w-16 shrink-0 text-slate-400">时间</span>
                  <span className="text-slate-800 dark:text-slate-200">{formatDateTime(fetchCode.result.received_at)}</span>
                </div>
                <div className="flex gap-2">
                  <span className="w-16 shrink-0 text-slate-400">正文</span>
                  <span className="text-slate-600 dark:text-slate-300">{fetchCode.result.body_preview || '—'}</span>
                </div>
              </div>
            </div>
          ) : (
            <div className="rounded-lg border border-dashed border-slate-300 p-10 text-center text-slate-400 dark:border-slate-700">
              当前没有可展示的验证码结果
            </div>
          )}
          <DialogFooter>
            <Button variant="outline" onClick={() => setFetchCode((current) => ({ ...current, open: false }))}>
              关闭
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  )
}
