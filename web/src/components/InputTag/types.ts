export interface InputTagProps {
  /**
   * 绑定值
   */
  modelValue?: string[]

  /**
   * 占位符
   */
  placeholder?: string

  /**
   * 是否禁用
   */
  disabled?: boolean

  /**
   * 是否只读
   */
  readonly?: boolean

  /**
   * 尺寸
   */
  size?: 'small' | 'medium' | 'large'

  /**
   * 最大标签数量
   */
  max?: number

  /**
   * 是否允许重复标签
   */
  allowDuplicates?: boolean

  /**
   * 分隔符
   */
  separator?: string | RegExp

  /**
   * 标签验证函数
   */
  validate?: (tag: string) => boolean | string

  /**
   * 标签类型
   */
  tagType?: 'default' | 'primary' | 'success' | 'warning' | 'danger' | 'info'

  /**
   * 标签尺寸
   */
  tagSize?: 'small' | 'medium' | 'large'

  /**
   * 标签变体
   */
  tagVariant?: 'filled' | 'outlined' | 'ghost'

  /**
   * 是否可关闭标签
   */
  closable?: boolean

  /**
   * 输入框类型
   */
  inputType?: 'text' | 'email' | 'url'

  /**
   * 自动聚焦
   */
  autofocus?: boolean

  /**
   * 最大标签长度
   */
  maxTagLength?: number

  /**
   * 最小标签长度
   */
  minTagLength?: number

  /**
   * 是否可排序
   */
  sortable?: boolean

  /**
   * 错误状态
   */
  error?: boolean

  /**
   * 是否显示计数
   */
  showCount?: boolean

  /**
   * 自定义图标
   */
  icon?: string

  /**
   * 预设标签
   */
  suggestions?: string[]

  /**
   * 是否显示建议
   */
  showSuggestions?: boolean
}

export interface InputTagEmits {
  /**
   * 更新值事件
   */
  (e: 'update:modelValue', value: string[]): void

  /**
   * 标签添加事件
   */
  (e: 'add', tag: string): void

  /**
   * 标签移除事件
   */
  (e: 'remove', tag: string, index: number): void

  /**
   * 输入事件
   */
  (e: 'input', value: string): void

  /**
   * 聚焦事件
   */
  (e: 'focus', event: FocusEvent): void

  /**
   * 失焦事件
   */
  (e: 'blur', event: FocusEvent): void

  /**
   * 验证失败事件
   */
  (e: 'invalid', tag: string, error: string): void

  /**
   * 达到最大数量事件
   */
  (e: 'max-reached', value: string[]): void

  /**
   * 标签排序事件
   */
  (e: 'sort', tags: string[]): void
}

export interface InputTagInstance {
  /**
   * 获取焦点
   */
  focus: () => void

  /**
   * 失去焦点
   */
  blur: () => void

  /**
   * 添加标签
   */
  addTag: (tag: string) => boolean

  /**
   * 移除标签
   */
  removeTag: (index: number) => void

  /**
   * 清空所有标签
   */
  clear: () => void

  /**
   * 获取当前输入值
   */
  getInputValue: () => string

  /**
   * 设置输入值
   */
  setInputValue: (value: string) => void
}
