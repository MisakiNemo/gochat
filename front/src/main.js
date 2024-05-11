import './assets/main.css'
import pinia from '@/store'
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import {ElementPlus} from "@element-plus/icons-vue";



const app=createApp(App)
    app.use(ElementPlus)
    app.use(router)
    app.use(pinia)
    app.mount('#app')