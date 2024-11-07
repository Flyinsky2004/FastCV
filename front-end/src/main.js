import '@/assets/css/main.css'
import 'animate.css';
import Antd from 'ant-design-vue';
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import 'ant-design-vue/dist/reset.css';

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router).use(ElementPlus).use(Antd)

app.mount('#app')
