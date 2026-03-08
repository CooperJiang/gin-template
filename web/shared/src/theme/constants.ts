import type { VariableCategory } from './types'

/** Map from CSS variable suffix → ThemeModeColors dot path */
export const CSS_VAR_PREFIX = '--theme-'

const scaleSteps = [50, 100, 200, 300, 400, 500, 600, 700, 800, 900, 950] as const

function scaleItems(prefix: string, label: string): VariableCategory['items'] {
  return scaleSteps.map((s) => ({
    key: `${prefix}-${s}`,
    path: `${prefix}.${s}`,
    label: `${label} ${s}`,
  }))
}

function statusItems(prefix: string, label: string): VariableCategory['items'] {
  return [
    { key: `${prefix}-main`, path: `${prefix}.main`, label: `${label}` },
    { key: `${prefix}-light`, path: `${prefix}.light`, label: `${label} 浅色` },
  ]
}

export const VARIABLE_CATEGORIES: VariableCategory[] = [
  { id: 'primary', label: '主色调', items: scaleItems('primary', '主色') },
  { id: 'gray', label: '中性色', items: scaleItems('gray', '灰色') },
  { id: 'status', label: '状态色', items: [
    ...statusItems('success', '成功'),
    ...statusItems('warning', '警告'),
    ...statusItems('error', '错误'),
    ...statusItems('info', '信息'),
  ]},
  { id: 'background', label: '背景', items: [
    { key: 'bg', path: 'bg', label: '页面背景' },
    { key: 'bg-soft', path: 'bgSoft', label: '柔和背景' },
    { key: 'bg-mute', path: 'bgMute', label: '静默背景' },
  ]},
  { id: 'surface', label: '面板', items: [
    { key: 'surface', path: 'surface', label: '面板' },
    { key: 'surface-hover', path: 'surfaceHover', label: '面板悬停' },
    { key: 'surface-active', path: 'surfaceActive', label: '面板激活' },
  ]},
  { id: 'text', label: '文字', items: [
    { key: 'text', path: 'text', label: '主文字' },
    { key: 'text-secondary', path: 'textSecondary', label: '次要文字' },
    { key: 'text-muted', path: 'textMuted', label: '弱化文字' },
    { key: 'text-inverse', path: 'textInverse', label: '反色文字' },
    { key: 'heading', path: 'heading', label: '标题' },
  ]},
  { id: 'border', label: '边框', items: [
    { key: 'border', path: 'border', label: '边框' },
    { key: 'border-light', path: 'borderLight', label: '浅边框' },
    { key: 'border-hover', path: 'borderHover', label: '悬停边框' },
  ]},
  { id: 'ring', label: '焦点环', items: [
    { key: 'ring', path: 'ring', label: '焦点环' },
  ]},
  { id: 'input', label: '表单', items: [
    { key: 'input-bg', path: 'inputBg', label: '输入框背景' },
    { key: 'input-border', path: 'inputBorder', label: '输入框边框' },
    { key: 'input-text', path: 'inputText', label: '输入框文字' },
    { key: 'input-placeholder', path: 'inputPlaceholder', label: '占位符' },
  ]},
  { id: 'link', label: '链接', items: [
    { key: 'link', path: 'link', label: '链接' },
    { key: 'link-hover', path: 'linkHover', label: '链接悬停' },
  ]},
  { id: 'overlay', label: '遮罩', items: [
    { key: 'overlay', path: 'overlay', label: '遮罩' },
  ]},
  { id: 'scrollbar', label: '滚动条', items: [
    { key: 'scrollbar-track', path: 'scrollbarTrack', label: '滚动条轨道' },
    { key: 'scrollbar-thumb', path: 'scrollbarThumb', label: '滚动条滑块' },
    { key: 'scrollbar-thumb-hover', path: 'scrollbarThumbHover', label: '滑块悬停' },
  ]},
]

/** Flatten all variable items for iteration */
export const ALL_VARIABLES = VARIABLE_CATEGORIES.flatMap((c) => c.items)
