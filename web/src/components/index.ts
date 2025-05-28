// 导入所有组件
import { AppLoading } from './AppLoading'

// 组件库插件安装函数
import type { App, Plugin } from 'vue'

const GinComponentsPlugin: Plugin = {
  install(app: App) {
    app.component('GinLoading', AppLoading)
  }
}

export default GinComponentsPlugin
