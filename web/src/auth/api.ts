import { apiClient } from '@/api/client'
import type {
  LoginRequest,
  LoginResponse,
  RegisterRequest,
  SendCodeRequest,
  ResetPasswordRequest,
  ChangePasswordRequest,
  UpdateProfileRequest,
  User,
} from '@/types/auth'

export const authApi = {
  login: (data: LoginRequest) => apiClient.post<LoginResponse>('/user/login', data),
  register: (data: RegisterRequest) => apiClient.post('/user/register', data),
  sendRegistrationCode: (data: SendCodeRequest) =>
    apiClient.post('/user/send-registration-code', data),
  sendResetPasswordCode: (data: SendCodeRequest) =>
    apiClient.post('/user/send-reset-password-code', data),
  sendChangeEmailCode: (data: SendCodeRequest) =>
    apiClient.post('/user/send-change-email-code', data),
  resetPassword: (data: ResetPasswordRequest) => apiClient.post('/user/reset-password', data),
  getUserInfo: () => apiClient.get<User>('/user/info'),
  updateProfile: (data: UpdateProfileRequest) => apiClient.put<User>('/user/profile', data),
  changePassword: (data: ChangePasswordRequest) =>
    apiClient.post('/user/change-password', data),
}
