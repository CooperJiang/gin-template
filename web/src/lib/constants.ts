export const VALIDATION = {
  PASSWORD_MIN: 6,
  PASSWORD_MAX: 20,
  USERNAME_MIN: 2,
  USERNAME_MAX: 20,
  CODE_LENGTH: 6,
  CODE_COUNTDOWN: 60,
} as const

export const EMAIL_REGEX = /^[^\s@]+@[^\s@]+\.[^\s@]+$/

export const INPUT_CLASS =
  'block w-full rounded-lg border border-gray-200 dark:border-gray-700 px-3.5 py-2.5 text-sm text-gray-900 dark:text-gray-100 bg-white dark:bg-gray-800 placeholder:text-gray-400 dark:placeholder:text-gray-500 focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 outline-none transition disabled:bg-gray-50 dark:disabled:bg-gray-700 disabled:opacity-60'
