import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import App from './App.vue'
import axios from 'axios'
import VueClipboard from 'vue-clipboard2'
import {router} from './router'

axios.defaults.baseURL = 'http://127.0.0.1:10370/updown'

const app = createApp(App)
app.use(router)
app.use(VueClipboard)
app.use(ElementPlus)
app.mount('#app')