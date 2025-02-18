// 创建一个路由并暴露
import {createRouter, createWebHistory} from "vue-router";
// 引入要呈现的组件
// 创建路由器
const router = createRouter({
    // 配置对象
    history: createWebHistory(),    // 指定路由时候一定要指定路由的工作模式
    routes: [   // 一个个的路由规则
        {
            path: '/',
            redirect: '/etcd',
        },
        {
            path: '/etcd',
            component: () => import("@/views/EtcdOptions.vue"),
        },
        {
            path: '/mysql',
            component: () => import("@/views/MysqlOptions.vue"),
        },
        {
            path: '/redis',
            component: () => import("@/views/RedisOptions.vue"),
        },
        {
            path: '/api',
            component: () => import("@/views/ApiOptions.vue"),
        },
        {
            path: '/mq',
            component: () => import("@/views/MqOptions.vue"),
        }
    ]
})
// 导出
export default router