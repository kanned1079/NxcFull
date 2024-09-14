import { type RouteRecordRaw } from 'vue-router';
import UserLogin from "@/views/User/Login/UserLogin.vue";
import UserRegister from "@/views/User/Login/UserRegister.vue"
import CommonAside from "@/components/CommonAside.vue";
import UserDashBoard from "@/views/User/UserDashBoard.vue";
import UserSummary from "@/views/User/pages/RootViews/UserSummary.vue";
import UserDocument from "@/views/User/pages/RootViews/UserDocument.vue";

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
    {
        path: '/dashboard',
        component: UserDashBoard,
        children: [
            {
                path: '/dashboard/summary',
                component: UserSummary,
            },
            {
                path: '/dashboard/document',
                component: UserDocument,
            }
        ]
    }

];

export default userRoutes;