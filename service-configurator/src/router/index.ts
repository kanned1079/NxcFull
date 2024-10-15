// 创建一个路由并暴露
import {createRouter, createWebHistory} from "vue-router";
// 引入要呈现的组件
import EtcdOptions from "@/views/EtcdOptions.vue";
import MysqlOptions from "@/views/MysqlOptions.vue";
import RedisOptions from "@/views/RedisOptions.vue";
import MqOptions from "@/views/MqOptions.vue";
import ApiOptions from "@/views/ApiOptions.vue";
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
            component: EtcdOptions,
        },
        {
            path: '/mysql',
            component: MysqlOptions,
        },
        {
            path: '/redis',
            component: RedisOptions,
        },
        {
            path: '/api',
            component: ApiOptions,
        },
        {
            path: '/mq',
            component: MqOptions,
        }
    ]
})
// 导出
export default router