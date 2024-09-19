// 通用字体
import 'vfonts/Lato.css'
// 等宽字体
import 'vfonts/FiraCode.css'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
// auto
import { createPinia } from 'pinia'
import { createApp } from 'vue'
//
// // md编辑和展示
// import VueMarkdownEditor from '@kangc/v-md-editor';
// import '@kangc/v-md-editor/lib/style/base-editor.css';
// import vuepressTheme from '@kangc/v-md-editor/lib/theme/vuepress.js';
// import '@kangc/v-md-editor/lib/theme/style/vuepress.css';
//
// // md预览
// import VMdPreview from '@kangc/v-md-editor/lib/preview';
// import '@kangc/v-md-editor/lib/style/preview.css';
// import githubTheme from '@kangc/v-md-editor/lib/theme/github.js';
// import '@kangc/v-md-editor/lib/theme/style/github.css';

// import Prism from 'prismjs';
//
// // highlightjs
// import hljs from 'highlight.js';

import App from './App.vue'
import router from './router'

// VueMarkdownEditor.use(vuepressTheme, {
//     Prism,
// });
//
// VMdPreview.use(githubTheme, {
//     Hljs: hljs,
// });


const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)

const app = createApp(App)

// app.use(VueMarkdownEditor)
// app.use(VMdPreview)
app.use(pinia)
app.use(router)

app.mount('#app')
