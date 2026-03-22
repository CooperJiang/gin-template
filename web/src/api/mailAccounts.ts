import { apiClient } from '@/api/client'
import type {
  BatchDeleteMailAccountRequest,
  CreateMailAccountRequest,
  DailyUsageResponse,
  FetchCodeRequest,
  FetchCodeResponse,
  MailAccount,
  MailAccountImportRequest,
  MailAccountImportResponse,
  MailAccountListParams,
  MailAccountListResponse,
  MailAccountStats,
  MailPreviewResponse,
  UpdateMailAccountRequest,
  UpdateMailAccountStatusRequest,
} from '@/types/mailAccount'

export const mailAccountsApi = {
  list: async (params: MailAccountListParams): Promise<MailAccountListResponse> => {
    const response = await apiClient.get<MailAccountListResponse | { items: MailAccountListResponse['items']; pagination: { page: number; page_size: number; total: number; total_pages: number } }>('/mail-accounts', { params })

    if ('pagination' in response) {
      return {
        items: response.items,
        total: response.pagination.total,
        page: response.pagination.page,
        page_size: response.pagination.page_size,
        total_pages: response.pagination.total_pages,
      }
    }

    return response
  },
  getById: (id: string) => apiClient.get<MailAccount>(`/mail-accounts/${id}`),
  importAccounts: (data: MailAccountImportRequest) =>
    apiClient.post<MailAccountImportResponse>('/mail-accounts/import', data),
  create: (data: CreateMailAccountRequest) => apiClient.post<MailAccount>('/mail-accounts', data),
  update: (id: string, data: UpdateMailAccountRequest) =>
    apiClient.put<MailAccount>(`/mail-accounts/${id}`, data),
  updateStatus: (id: string, data: UpdateMailAccountStatusRequest) =>
    apiClient.patch<MailAccount>(`/mail-accounts/${id}/status`, data),
  fetchMails: (id: string) =>
    apiClient.post<MailPreviewResponse>(`/mail-accounts/${id}/fetch-mails`, {}),
  fetchCode: (id: string, data?: FetchCodeRequest) =>
    apiClient.post<FetchCodeResponse>(`/mail-accounts/${id}/fetch-code`, data ?? {}),
  delete: (id: string) => apiClient.delete(`/mail-accounts/${id}`),
  batchDelete: (data: BatchDeleteMailAccountRequest) =>
    apiClient.post('/mail-accounts/batch-delete', data),
  stats: () => apiClient.get<MailAccountStats>('/mail-accounts/stats'),
  dailyUsage: () => apiClient.get<DailyUsageResponse>('/mail-accounts/daily-usage'),
}
