// 导入所有组件
import AppBreadcrumb from './AppBreadcrumb'
import PopoverMenu from './PopoverMenu'
import Button from './Button'
import Icon from './Icon'
import FileUpload from './FileUpload'

// 组件库插件安装函数
import type { App, Plugin } from 'vue'

const GinComponentsPlugin: Plugin = {
  install(app: App) {
    app.component('GlobalAppBreadcrumb', AppBreadcrumb)
    app.component('GlobalPopoverMenu', PopoverMenu)
    app.component('GlobalButton', Button)
    app.component('GlobalIcon', Icon)
    app.component('GlobalFileUpload', FileUpload)
  },
}

export default GinComponentsPlugin

// 单独导出组件
export { AppBreadcrumb, PopoverMenu, Button, Icon, FileUpload }

export { default as MessageContainer } from './MessageContainer'
