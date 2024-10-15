import router from '@/router'
import { createApp } from 'vue'
import App from './App.vue'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'

const app = createApp(App)
const pinia = createPinia()

pinia.use(piniaPluginPersistedstate)

// createApp(App).mount('#app')
app.use(pinia)
app.use(router)


app.mount('#app')
