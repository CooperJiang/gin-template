import './styles/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

// 导入全局组件插件
import GinComponentsPlugin from './components'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(GinComponentsPlugin)

app.mount('#app')
