// 通用字体
import 'vfonts/Lato.css'
// 等宽字体
import 'vfonts/FiraCode.css'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'

// import useAppInfosStore from "@/stores/useAppInfosStore";
import {handleFetchRootRuntimeEnvConfig} from "@/api/common/env"

// 语言国际化
import {createI18n} from 'vue-i18n';

// auto
import {createPinia} from 'pinia'
import {createApp} from 'vue'

import App from './App.vue'
import router from './router'
import naive from "naive-ui";

import {langs} from "@/language"

const savedLocale = localStorage.getItem('locale') || 'zh_CN'

const i18n = createI18n({
    locale: savedLocale,
    fallbackLocale: 'en_US',
    messages: langs
})

const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)

const app = createApp(App)

app.use(naive)

// app.use(bcrypt)

app.use(pinia)
app.use(i18n)
app.use(router)

const initApp = async () => {
    if (await handleFetchRootRuntimeEnvConfig()) console.log('pull root config success.')
    else console.log('pull root config success.')
    app.mount('#app')
}

// app.mount('#app')
initApp()