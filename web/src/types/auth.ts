export interface User {
  id?: string
  username: string
  email: string
  status: number
  role?: number
  avatar?: string
  bio?: string
  created_at?: string
  updated_at?: string
}

export interface LoginRequest {
  account: string
  password: string
}

export interface LoginResponse {
  token: string
  user: User
}

export interface RegisterRequest {
  username: string
  email: string
  password: string
  code: string
}

export interface SendCodeRequest {
  email: string
}

export interface ResetPasswordRequest {
  email: string
  code: string
  newPassword: string
}

export interface ChangePasswordRequest {
  oldPassword: string
  newPassword: string
}

export interface UpdateProfileRequest {
  username?: string
  email?: string
  avatar?: string
  code?: string
}
