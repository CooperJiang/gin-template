// 导入所有组件
import AppBreadcrumb from './AppBreadcrumb'
import PopoverMenu from './PopoverMenu'
import Button from './Button'
import Icon from './Icon'
import Input from './Input'
import Textarea from './Textarea'
import Switch from './Switch'
import Checkbox from './Checkbox'
import Tag from './Tag'
import InputTag from './InputTag'
import FileUpload from './FileUpload'
import Pagination from './Pagination'
import Table from './Table'
import { TableColumn } from './Table/components'
import Select from './Select'
import Dialog from './Dialog'
import Form from './Form'
import FormItem from './Form/FormItem.vue'
import Tabs from './Tabs'
import Card from './Card'
import Progress from './Progress'
import CodeDemo from './CodeDemo'

// 组件库插件安装函数
import type { App, Plugin } from 'vue'

const GinComponentsPlugin: Plugin = {
  install(app: App) {
    app.component('GlobalAppBreadcrumb', AppBreadcrumb)
    app.component('GlobalPopoverMenu', PopoverMenu)
    app.component('GlobalButton', Button)
    app.component('GlobalIcon', Icon)
    app.component('GlobalInput', Input)
    app.component('GlobalTextarea', Textarea)
    app.component('GlobalSwitch', Switch)
    app.component('GlobalCheckbox', Checkbox)
    app.component('GlobalTag', Tag)
    app.component('GlobalInputTag', InputTag)
    app.component('GlobalFileUpload', FileUpload)
    app.component('GlobalPagination', Pagination)
    app.component('GlobalTable', Table)
    app.component('GlobalTableColumn', TableColumn)
    app.component('GlobalSelect', Select)
    app.component('GlobalDialog', Dialog)
    app.component('GlobalForm', Form)
    app.component('GlobalFormItem', FormItem)
    app.component('GlobalTabs', Tabs)
    app.component('GlobalCard', Card)
    app.component('GlobalProgress', Progress)
    app.component('GlobalCodeDemo', CodeDemo)
  },
}

export default GinComponentsPlugin

// 单独导出组件
export {
  AppBreadcrumb,
  PopoverMenu,
  Button,
  Icon,
  Input,
  Textarea,
  Switch,
  Checkbox,
  Tag,
  InputTag,
  FileUpload,
  Pagination,
  Table,
  TableColumn,
  Select,
  Dialog,
  Form,
  FormItem,
  Tabs,
  Card,
  Progress,
  CodeDemo
}

export { default as MessageContainer } from './MessageContainer'
