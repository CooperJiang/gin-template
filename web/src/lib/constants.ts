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
  'block w-full border-nb border-nb-border bg-nb-surface px-3.5 py-2.5 text-sm text-nb-text placeholder:text-nb-text-muted focus:ring-2 focus:ring-nb-primary focus:border-nb-border outline-none transition-shadow disabled:opacity-60 rounded-[var(--nb-radius)] shadow-nb-sm focus:shadow-nb'
