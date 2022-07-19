import './plugins/axios'
import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import App from './App.vue'
import router from './router'
import store from './store'
import 'element-plus/theme-chalk/dark/css-vars.css'
import 'element-plus/dist/index.css'
createApp(App).use(store).use(store).use(router).use(ElementPlus).mount('#app')
