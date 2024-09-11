import { type RouteRecordRaw } from 'vue-router';
import UserLogin from "@/views/User/Login/UserLogin.vue";
import UserRegister from "@/views/User/Login/UserRegister.vue"

const userRoutes: RouteRecordRaw[] = [
    {
        path: '/login',
        name: 'login',
        component: UserLogin
    },
    {
        path: '/register',
        name: 'register',
        component: UserRegister
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