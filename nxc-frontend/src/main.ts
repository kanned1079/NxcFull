// 通用字体
import 'vfonts/Lato.css'
// 等宽字体
import 'vfonts/FiraCode.css'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'

// 语言国际化
import { createI18n } from 'vue-i18n';
import en_US from "@/language/en_US";
import zh_CN from "@/language/zh_CN";
import ja_JP from "@/language/ja_JP";
import zh_HK from "@/language/zh_HK";

// auto
import { createPinia } from 'pinia'
import { createApp } from 'vue'

import App from './App.vue'
import router from './router'
import naive, {zhCN} from "naive-ui";
// import bcrypt from 'bcryptjs'
// Vue.use(bcrypt)

const messages = {
    en_US,
    zh_CN,
    ja_JP,
    zh_HK,

}

const savedLocale = localStorage.getItem('locale') || 'zh_CN'

const i18n = createI18n({
    // locale: 'zh_CN',
    locale: savedLocale,
    // locale: 'en_US',
    // fallbackLocale: 'en',
    messages,
})

const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)

const app = createApp(App)

app.use(naive)

// app.use(bcrypt)

app.use(pinia)
app.use(i18n)
app.use(router)

app.mount('#app')
