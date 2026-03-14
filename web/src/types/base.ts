export interface BaseResponse<T = unknown> {
  code: number
  message: string
  data: T
  request_id?: string
  timestamp?: string
}

export interface PaginationParams {
  page?: number
  size?: number
}

export interface PaginationResponse<T> {
  items: T[]
  total: number
  page: number
  page_size: number
  total_pages: number
}
