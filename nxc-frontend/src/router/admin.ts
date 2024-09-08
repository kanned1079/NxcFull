import { type RouteRecordRaw } from 'vue-router';
import DashBoard from "@/views/Admin/DashBoard.vue";
import Summary from "@/views/Admin/pages/RootViews/Summary.vue";
import QueueMonitor from "@/views/Admin/pages/RootViews/QueueMonitor.vue";
import SystemConfig from "@/views/Admin/pages/SettingViews/SystemConfig.vue";
import PaymentConfig from "@/views/Admin/pages/SettingViews/SystemConfig/PaymentConfig.vue";
import ThemeConfig from "@/views/Admin/pages/SettingViews/ThemeConfig.vue";
import UserManager from "@/views/Admin/pages/UserMgrViews/UserManager.vue";
import AdminLogin from "@/views/Admin/Login/AdminLogin.vue";
import RouterMgr from "@/views/Admin/pages/ServerViews/RouterMgr.vue"
import PrivilegeGroup from "@/views/Admin/pages/ServerViews/PrivilegeGroup.vue"
import CouponMgr from "@/views/Admin/pages/FinanceViews/CouponMgr.vue";
import UserLogin from "@/views/User/Login/UserLogin.vue";

const adminRoutes: RouteRecordRaw[] = [
    {
        path: '/admin/dashboard',
        name: 'admin-dashboard',
        component: DashBoard,
        meta: {
            requireAuth: true,
        },
        children: [
            {
                path: '/admin/dashboard/summary',
                name: 'summary',
                component: Summary,
            },
            {
                path: '/admin/dashboard/monitor',
                name: 'monitor',
                component: QueueMonitor
            },
            {
                path: '/admin/dashboard/systemconfig',
                name: 'system-config',
                component: SystemConfig,
            },
            {
                path: '/admin/dashboard/payment',
                name: 'payment',
                component: PaymentConfig,
            },
            {
                path: '/admin/dashboard/theme',
                name: 'theme',
                component: ThemeConfig,
            },
            {
                path: '/admin/dashboard/node',
                name: 'node',
                component: ThemeConfig,
            },

            // part4
            {
                path: '/admin/dashboard/usermanager',
                name: 'user-manager',
                component: UserManager,
            },

            {
                path: '/admin/dashboard/routermgr',
                name: 'router-mgr',
                component: RouterMgr
            },
            {
                path: '/admin/dashboard/privilegegroup',
                name: 'privilege-mgr',
                component: PrivilegeGroup,
            }

        ]
    },
    {
        path: '/dashboard',
        component: DashBoard,
    },
    {
        path: '/admin/login',
        name: 'admin-login',
        component: AdminLogin
    },
    {
        path: '/admin/dashboard/couponmgr',
        name: 'coupon-mgr',
        component: CouponMgr,
    },
    // {
    //     path: '/login',
    //     name: 'login',
    //     component: UserLogin
    // },
    // {
    //     path: '/home',
    //     redirect: '/',
    // },
    // {
    //     path: '/',
    //     redirect: '/login',
    // },
    {
        path: '/admin',
        redirect: '/admin/dashboard',
    }
];

export default adminRoutes;