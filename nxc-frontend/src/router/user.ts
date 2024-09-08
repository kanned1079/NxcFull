import { type RouteRecordRaw } from 'vue-router';
import UserLogin from "@/views/User/Login/UserLogin.vue";

const userRoutes: RouteRecordRaw[] = [
    {
        path: '/login',
        name: 'login',
        component: UserLogin
    },
    {
        path: '/home',
        redirect: '/',
    },
    {
        path: '/',
        redirect: '/login',
    },

];

export default userRoutes;