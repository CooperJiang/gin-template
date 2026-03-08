import { ApiClient } from './request'
import type {
  LoginRequest,
  LoginResponse,
  RegisterRequest,
  SendCodeRequest,
  ResetPasswordRequest,
  ChangePasswordRequest,
  UpdateProfileRequest,
  User,
} from '@app/shared'

export const authApi = {
  login: (data: LoginRequest) => ApiClient.post<LoginResponse>('/user/login', data),
  register: (data: RegisterRequest) => ApiClient.post('/user/register', data),
  sendRegistrationCode: (data: SendCodeRequest) => ApiClient.post('/user/send-registration-code', data),
  sendResetPasswordCode: (data: SendCodeRequest) => ApiClient.post('/user/send-reset-password-code', data),
  sendChangeEmailCode: (data: SendCodeRequest) => ApiClient.post('/user/send-change-email-code', data),
  resetPassword: (data: ResetPasswordRequest) => ApiClient.post('/user/reset-password', data),
  getUserInfo: () => ApiClient.get<User>('/user/info'),
  updateProfile: (data: UpdateProfileRequest) => ApiClient.put<User>('/user/profile', data),
  changePassword: (data: ChangePasswordRequest) => ApiClient.post('/user/change-password', data),
}
