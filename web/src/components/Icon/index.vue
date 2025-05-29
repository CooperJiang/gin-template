<template>
  <component
    :is="iconComponent"
    :class="iconClass"
  />
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { IconProps } from './types'

// Heroicons Outline 图标
import {
  PlusIcon,
  MinusIcon,
  XMarkIcon,
  CheckIcon,
  PencilIcon,
  TrashIcon,
  EyeIcon,
  EyeSlashIcon,
  HeartIcon,
  StarIcon,
  ShareIcon,
  ArrowDownTrayIcon as DownloadIcon,
  ArrowUpTrayIcon as UploadIcon,
  ArrowPathIcon as RefreshIcon,
  Cog6ToothIcon as SettingsIcon,
  MagnifyingGlassIcon as SearchIcon,
  ArrowUpIcon,
  ArrowDownIcon,
  ArrowLeftIcon,
  ArrowRightIcon,
  ChevronUpIcon,
  ChevronDownIcon,
  ChevronLeftIcon,
  ChevronRightIcon,
  ArrowTopRightOnSquareIcon as ArrowUpRightIcon,
  ArrowUturnDownIcon as ArrowDownRightIcon,
  UserIcon,
  UserGroupIcon,
  UserPlusIcon,
  UserMinusIcon,
  ChatBubbleLeftIcon,
  ChatBubbleLeftRightIcon,
  PhoneIcon,
  EnvelopeIcon,
  BellIcon,
  BellSlashIcon,
  DocumentIcon,
  DocumentTextIcon,
  FolderIcon,
  FolderOpenIcon,
  ArchiveBoxIcon,
  ClipboardIcon,
  ClipboardDocumentIcon,
  PhotoIcon,
  VideoCameraIcon,
  MusicalNoteIcon,
  HomeIcon,
  BuildingOfficeIcon,
  MapPinIcon,
  GlobeAltIcon,
  CalendarIcon,
  ClockIcon,
  SunIcon,
  MoonIcon,
  ListBulletIcon,
  Squares2X2Icon,
  TableCellsIcon,
  Bars3Icon,
  Bars4Icon,
  EllipsisHorizontalIcon,
  CheckCircleIcon,
  XCircleIcon,
  ExclamationTriangleIcon,
  InformationCircleIcon,
  QuestionMarkCircleIcon,
  ShieldCheckIcon,
  ShieldExclamationIcon,
  LockClosedIcon,
  LockOpenIcon,
  KeyIcon,
  FingerPrintIcon,
  ShoppingCartIcon,
  ShoppingBagIcon,
  CreditCardIcon,
  BanknotesIcon,
  GiftIcon,
  TagIcon,
  TicketIcon,
  ChartBarIcon,
  ChartPieIcon,
  PresentationChartLineIcon,
  WrenchScrewdriverIcon,
  CpuChipIcon,
  ServerIcon,
  CloudIcon,
  WifiIcon,
  SignalIcon,
  DevicePhoneMobileIcon,
  ComputerDesktopIcon,
  TvIcon
} from '@heroicons/vue/24/outline'

// Heroicons Solid 图标
import {
  PlusIcon as PlusIconSolid,
  MinusIcon as MinusIconSolid,
  XMarkIcon as XMarkIconSolid,
  CheckIcon as CheckIconSolid,
  PencilIcon as PencilIconSolid,
  TrashIcon as TrashIconSolid,
  EyeIcon as EyeIconSolid,
  EyeSlashIcon as EyeSlashIconSolid,
  HeartIcon as HeartIconSolid,
  StarIcon as StarIconSolid,
  ShareIcon as ShareIconSolid,
  ArrowDownTrayIcon as DownloadIconSolid,
  ArrowUpTrayIcon as UploadIconSolid,
  ArrowPathIcon as RefreshIconSolid,
  Cog6ToothIcon as SettingsIconSolid,
  MagnifyingGlassIcon as SearchIconSolid,
  ArrowUpIcon as ArrowUpIconSolid,
  ArrowDownIcon as ArrowDownIconSolid,
  ArrowLeftIcon as ArrowLeftIconSolid,
  ArrowRightIcon as ArrowRightIconSolid,
  ChevronUpIcon as ChevronUpIconSolid,
  ChevronDownIcon as ChevronDownIconSolid,
  ChevronLeftIcon as ChevronLeftIconSolid,
  ChevronRightIcon as ChevronRightIconSolid,
  UserIcon as UserIconSolid,
  UserGroupIcon as UserGroupIconSolid,
  ChatBubbleLeftIcon as ChatBubbleLeftIconSolid,
  PhoneIcon as PhoneIconSolid,
  EnvelopeIcon as EnvelopeIconSolid,
  BellIcon as BellIconSolid,
  HomeIcon as HomeIconSolid,
  CheckCircleIcon as CheckCircleIconSolid,
  XCircleIcon as XCircleIconSolid,
  ExclamationTriangleIcon as ExclamationTriangleIconSolid,
  InformationCircleIcon as InformationCircleIconSolid
} from '@heroicons/vue/24/solid'

defineOptions({
  name: 'GlobalIcon'
})

const props = withDefaults(defineProps<IconProps>(), {
  size: 'md',
  variant: 'outline',
  class: '',
  color: ''
})

// 图标映射
const iconMap = {
  outline: {
    'plus': PlusIcon,
    'minus': MinusIcon,
    'x-mark': XMarkIcon,
    'check': CheckIcon,
    'pencil': PencilIcon,
    'trash': TrashIcon,
    'eye': EyeIcon,
    'eye-slash': EyeSlashIcon,
    'heart': HeartIcon,
    'star': StarIcon,
    'share': ShareIcon,
    'download': DownloadIcon,
    'upload': UploadIcon,
    'refresh': RefreshIcon,
    'settings': SettingsIcon,
    'search': SearchIcon,
    'arrow-up': ArrowUpIcon,
    'arrow-down': ArrowDownIcon,
    'arrow-left': ArrowLeftIcon,
    'arrow-right': ArrowRightIcon,
    'chevron-up': ChevronUpIcon,
    'chevron-down': ChevronDownIcon,
    'chevron-left': ChevronLeftIcon,
    'chevron-right': ChevronRightIcon,
    'arrow-up-right': ArrowUpRightIcon,
    'arrow-down-right': ArrowDownRightIcon,
    'arrow-path': RefreshIcon,
    'user': UserIcon,
    'user-group': UserGroupIcon,
    'user-plus': UserPlusIcon,
    'user-minus': UserMinusIcon,
    'chat-bubble-left': ChatBubbleLeftIcon,
    'chat-bubble-left-right': ChatBubbleLeftRightIcon,
    'phone': PhoneIcon,
    'envelope': EnvelopeIcon,
    'bell': BellIcon,
    'bell-slash': BellSlashIcon,
    'document': DocumentIcon,
    'document-text': DocumentTextIcon,
    'folder': FolderIcon,
    'folder-open': FolderOpenIcon,
    'archive-box': ArchiveBoxIcon,
    'clipboard': ClipboardIcon,
    'clipboard-document': ClipboardDocumentIcon,
    'photo': PhotoIcon,
    'video-camera': VideoCameraIcon,
    'musical-note': MusicalNoteIcon,
    'home': HomeIcon,
    'building-office': BuildingOfficeIcon,
    'map-pin': MapPinIcon,
    'globe-alt': GlobeAltIcon,
    'calendar': CalendarIcon,
    'clock': ClockIcon,
    'sun': SunIcon,
    'moon': MoonIcon,
    'list-bullet': ListBulletIcon,
    'squares-2x2': Squares2X2Icon,
    'table-cells': TableCellsIcon,
    'bars-3': Bars3Icon,
    'bars-4': Bars4Icon,
    'ellipsis-horizontal': EllipsisHorizontalIcon,
    'check-circle': CheckCircleIcon,
    'x-circle': XCircleIcon,
    'exclamation-triangle': ExclamationTriangleIcon,
    'information-circle': InformationCircleIcon,
    'question-mark-circle': QuestionMarkCircleIcon,
    'shield-check': ShieldCheckIcon,
    'shield-exclamation': ShieldExclamationIcon,
    'lock-closed': LockClosedIcon,
    'lock-open': LockOpenIcon,
    'key': KeyIcon,
    'finger-print': FingerPrintIcon,
    'shopping-cart': ShoppingCartIcon,
    'shopping-bag': ShoppingBagIcon,
    'credit-card': CreditCardIcon,
    'banknotes': BanknotesIcon,
    'gift': GiftIcon,
    'tag': TagIcon,
    'ticket': TicketIcon,
    'chart-bar': ChartBarIcon,
    'chart-pie': ChartPieIcon,
    'presentation-chart-line': PresentationChartLineIcon,
    'cog-6-tooth': SettingsIcon,
    'wrench-screwdriver': WrenchScrewdriverIcon,
    'cpu-chip': CpuChipIcon,
    'server': ServerIcon,
    'cloud': CloudIcon,
    'wifi': WifiIcon,
    'signal': SignalIcon,
    'device-phone-mobile': DevicePhoneMobileIcon,
    'computer-desktop': ComputerDesktopIcon,
    'tv': TvIcon
  },
  solid: {
    'plus': PlusIconSolid,
    'minus': MinusIconSolid,
    'x-mark': XMarkIconSolid,
    'check': CheckIconSolid,
    'pencil': PencilIconSolid,
    'trash': TrashIconSolid,
    'eye': EyeIconSolid,
    'eye-slash': EyeSlashIconSolid,
    'heart': HeartIconSolid,
    'star': StarIconSolid,
    'share': ShareIconSolid,
    'download': DownloadIconSolid,
    'upload': UploadIconSolid,
    'refresh': RefreshIconSolid,
    'settings': SettingsIconSolid,
    'search': SearchIconSolid,
    'arrow-up': ArrowUpIconSolid,
    'arrow-down': ArrowDownIconSolid,
    'arrow-left': ArrowLeftIconSolid,
    'arrow-right': ArrowRightIconSolid,
    'chevron-up': ChevronUpIconSolid,
    'chevron-down': ChevronDownIconSolid,
    'chevron-left': ChevronLeftIconSolid,
    'chevron-right': ChevronRightIconSolid,
    'user': UserIconSolid,
    'user-group': UserGroupIconSolid,
    'chat-bubble-left': ChatBubbleLeftIconSolid,
    'phone': PhoneIconSolid,
    'envelope': EnvelopeIconSolid,
    'bell': BellIconSolid,
    'home': HomeIconSolid,
    'check-circle': CheckCircleIconSolid,
    'x-circle': XCircleIconSolid,
    'exclamation-triangle': ExclamationTriangleIconSolid,
    'information-circle': InformationCircleIconSolid
  }
}

// 计算图标组件
const iconComponent = computed(() => {
  const icons = iconMap[props.variant]
  const iconName = props.name as keyof typeof icons

  // 如果当前变体中有该图标，直接返回
  if (icons[iconName]) {
    return icons[iconName]
  }

  // 如果当前变体中没有，使用plus图标作为fallback
  return icons['plus'] || iconMap.outline['plus']
})

// 计算图标样式类
const iconClass = computed(() => {
  const classes = []

  // 尺寸样式
  const sizeClasses = {
    'xs': 'h-3 w-3',
    'sm': 'h-4 w-4',
    'md': 'h-5 w-5',
    'lg': 'h-6 w-6',
    'xl': 'h-8 w-8',
    '2xl': 'h-10 w-10'
  }
  classes.push(sizeClasses[props.size])

  // 颜色样式
  if (props.color) {
    classes.push(props.color)
  }

  // 自定义类名
  if (props.class) {
    classes.push(props.class)
  }

  return classes.join(' ')
})
</script>
