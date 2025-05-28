import { ApiClient } from '@/utils/request'
import type {
  LoginRequest,
  LoginResponse,
  RegisterRequest,
  SendCodeRequest,
  ResetPasswordRequest,
  ChangePasswordRequest,
  UpdateProfileRequest,
  User
} from '@/types'

export const authApi = {
  // 登录
  login: (data: LoginRequest) =>
    ApiClient.post<LoginResponse>('/user/login', data),

  // 注册
  register: (data: RegisterRequest) =>
    ApiClient.post('/user/register', data),

  // 发送注册验证码
  sendRegistrationCode: (data: SendCodeRequest) =>
    ApiClient.post('/user/send-registration-code', data),

  // 发送重置密码验证码
  sendResetPasswordCode: (data: SendCodeRequest) =>
    ApiClient.post('/user/send-reset-password-code', data),

  // 重置密码
  resetPassword: (data: ResetPasswordRequest) =>
    ApiClient.post('/user/reset-password', data),

  // 获取用户信息
  getUserInfo: () =>
    ApiClient.get<User>('/user/info'),

  // 更新用户资料
  updateProfile: (data: UpdateProfileRequest) =>
    ApiClient.put<User>('/user/profile', data),

  // 修改密码
  changePassword: (data: ChangePasswordRequest) =>
    ApiClient.post('/user/change-password', data)
}
