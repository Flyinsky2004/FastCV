import '@/assets/css/main.css'
import 'animate.css';
import Antd from 'ant-design-vue';
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import axios from "axios";
import 'ant-design-vue/dist/reset.css';
import { MotionPlugin } from '@vueuse/motion'
import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router).use(ElementPlus).use(Antd).use(MotionPlugin)
//开发环境
//axios.defaults.baseURL="http://localhost:8080"
//生产环境3
axios.defaults.baseURL= window.location.protocol + '//' + window.location.host
app.mount('#app')
