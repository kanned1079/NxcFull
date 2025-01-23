import {fileURLToPath, URL} from 'node:url'

import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'

// 安装NaiveUI
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { NaiveUiResolver } from 'unplugin-vue-components/resolvers'
// 第一步导入插件
import VueSetupExtend from 'vite-plugin-vue-setup-extend'

import svgLoader from 'vite-svg-loader';

import vueDevTools from 'vite-plugin-vue-devtools'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        vue(),
        svgLoader(),
        VueSetupExtend(), // 第二步使用插件
        vueDevTools(),  // 开发工具
        AutoImport({
            imports: [
                'vue',
                {
                    'naive-ui': [
                        'useDialog',
                        'useMessage',
                        'useNotification',
                        'useLoadingBar'
                    ]
                }
            ]
        }),
        Components({
            resolvers: [NaiveUiResolver()]
        })
    ],
    // define: {
    //     __APP_VERSION__: 'v1.0.0_patch4',
    // },
    resolve: {
        alias: {
            '@': fileURLToPath(new URL('./src', import.meta.url))
        }
    },
})
