import type { PaginationParams, PaginationResponse } from '@/types/base'

export type MailAccountStatus = 'unused' | 'used'
export type MailboxStatus = 'unknown' | 'ready' | 'fetch_failed' | 'token_expired' | 'incomplete'
export type ImportSourceType = 'paste' | 'file'

export interface MailAccount {
  id: string
  email: string
  password: string
  client_id: string
  token_type: string
  folder: string
  source_format: string
  provider_ready: boolean
  status: MailAccountStatus
  mailbox_status: MailboxStatus
  last_permission_scope: string
  last_use_local_ip: boolean
  last_provider_check_at?: string | null
  last_mail_fetch_at?: string | null
  last_code_fetch_at?: string | null
  last_code: string
  last_code_at?: string | null
  last_mail_subject: string
  last_mail_from: string
  last_mail_received_at?: string | null
  allocated_at?: string | null
  remark: string
  tags: string
  import_batch_no: string
  created_at: string
  updated_at: string
  password_masked: string
  refresh_token_masked: string
  refresh_token_configured: boolean
}

export interface MailAccountListParams extends PaginationParams {
  keyword?: string
  status?: MailAccountStatus
  mailbox_status?: MailboxStatus
  provider_ready?: boolean
}

export type MailAccountListResponse = PaginationResponse<MailAccount>

export interface MailAccountImportRequest {
  content: string
  source_type: ImportSourceType
  filename?: string
  mode?: 'fill_missing'
}

export interface MailAccountImportError {
  line: number
  reason: string
  raw_line: string
}

export interface MailAccountImportResponse {
  total_lines: number
  created_count: number
  updated_count: number
  skipped_count: number
  failed_count: number
  batch_no: string
  errors: MailAccountImportError[]
}

export interface CreateMailAccountRequest {
  email: string
  password?: string
  client_id?: string
  refresh_token?: string
  status?: MailAccountStatus
  remark?: string
  tags?: string
  folder?: string
}

export interface UpdateMailAccountRequest {
  password?: string
  client_id?: string
  refresh_token?: string
  status?: MailAccountStatus
  remark?: string
  tags?: string
  folder?: string
}

export interface UpdateMailAccountStatusRequest {
  status: MailAccountStatus
}

export interface MailPreviewMessage {
  id: string
  subject: string
  from_address: string
  from_name: string
  received_time: string
  body_preview: string
  body?: string
  is_read: boolean
}

export interface MailPreviewResponse {
  messages: MailPreviewMessage[]
}

export interface FetchCodeRequest {
  keyword?: string
}

export interface BatchDeleteMailAccountRequest {
  ids: string[]
}

export interface FetchCodeResponse {
  code: string
  matched_subject: string
  matched_from: string
  received_at?: string | null
  body_preview: string
}

export interface MailAccountStats {
  total: number
  unused: number
  used: number
}

export interface DailyUsageItem {
  date: string
  count: number
}

export interface DailyUsageResponse {
  items: DailyUsageItem[]
}
