import {type RouteRecordRaw} from 'vue-router';
// // import DashBoard from "@/views/Admin/DashBoard.vue";
// import Summary from "@/views/Admin/pages/RootViews/Summary.vue";
// import QueueMonitor from "@/views/Admin/pages/RootViews/QueueMonitor.vue";
// import SystemConfig from "@/views/Admin/pages/SettingViews/SystemConfig.vue";
// import PaymentConfig from "@/views/Admin/pages/SettingViews/PaymentConfig.vue";
// import ThemeConfig from "@/views/Admin/pages/SettingViews/ThemeConfig.vue";
// import UserManager from "@/views/Admin/pages/UserMgrViews/UserManager.vue";
// // import AdminLogin from "@/views/Admin/Login/AdminLogin.vue";
// import RouterMgr from "@/views/Admin/pages/ServerViews/RouterMgr.vue"
// import PrivilegeGroup from "@/views/Admin/pages/ServerViews/PrivilegeGroupMgr.vue"
// import CouponMgr from "@/views/Admin/pages/FinanceViews/CouponMgr.vue";
// import NoticeManager from "@/views/Admin/pages/UserMgrViews/NoticeManager.vue";
// import SubscribeMgr from "@/views/Admin/pages/FinanceViews/SubscribeMgr.vue";
// import DocumentMgr from "@/views/Admin/pages/UserMgrViews/DocumentMgr.vue";
// // import PrivilegeGroupMgr from "@/views/Admin/pages/ServerViews/PrivilegeGroupMgr.vue";
// import TicketMgr from "@/views/Admin/pages/UserMgrViews/TicketMgr.vue";
// import OrderMgr from "@/views/Admin/pages/FinanceViews/OrderMgr.vue"
// import ActivationMgr from "@/views/Admin/pages/FinanceViews/ActivationMgr.vue";
// import keyMgr from "@/views/Admin/pages/FinanceViews/KeyMgr.vue";
// import useAppInfosStore from "@/stores/useAppInfosStore";
// import {useRouter} from "vue-router";

// const appInfo = useAppInfosStore();

const adminRoutes: RouteRecordRaw[] = [
    {
        path: '/admin/login/:path?',
        name: 'admin-login',
        // component: AdminLogin
        component: () => import("@/views/Admin/Login/AdminLogin.vue")
    },
    {
        path: '/admin',
        redirect: '/admin/dashboard',
    },
    {
        path: '/admin/dashboard',
        name: 'admin-dashboard',
        component: () => import("@/views/Admin/DashBoard.vue"),
        meta: {
            requireAuth: true,
        },
        children: [
            {
                path: 'summary',
                name: 'summary',
                component: () => import("@/views/Admin/pages/RootViews/Summary.vue"),
            },
            {
                path: 'monitor',
                name: 'monitor',
                component: () => import("@/views/Admin/pages/RootViews/QueueMonitor.vue")
            },
            {
                path: 'systemconfig',
                name: 'system-config',
                component: () => import("@/views/Admin/pages/SettingViews/SystemConfig.vue"),
            },
            {
                path: 'payment',
                name: 'payment',
                component: () => import("@/views/Admin/pages/SettingViews/PaymentConfig.vue"),
            },
            {
                path: 'theme',
                name: 'theme',
                component: () => import("@/views/Admin/pages/SettingViews/ThemeConfig.vue"),
            },
            // {
            //     path: 'node',
            //     name: 'node',
            //     component: ThemeConfig,
            // },

            // part4
            {
                path: 'usermanager',
                name: 'user-manager',
                component: () => import("@/views/Admin/pages/UserMgrViews/UserManager.vue"),
            },

            // {
            //     path: 'routermgr',
            //     name: 'router-mgr',
            //     component: RouterMgr
            // },
            // {
            //     path: 'privilegegroup',
            //     name: 'privilege-mgr',
            //     component: () => import("@/views/Admin/pages/ServerViews/PrivilegeGroupMgr.vue"),
            // },
            {
                path: 'noticemanager',
                name: 'notice-manager',
                component: () => import("@/views/Admin/pages/UserMgrViews/NoticeManager.vue"),
            },
            {
                path: 'subscribemanager',
                name: 'subscription-manager',
                component: () => import("@/views/Admin/pages/FinanceViews/SubscribeMgr.vue"),
            },
            {
                path: 'document',
                name: 'document-mgr',
                component: () => import("@/views/Admin/pages/UserMgrViews/DocumentMgr.vue"),
            },
            {
                path: 'coupon',
                name: 'coupon',
                component: () => import("@/views/Admin/pages/FinanceViews/CouponMgr.vue"),
            },
            {
                path: 'group',
                name: 'privilege-group-mgr',
                component: () => import("@/views/Admin/pages/ServerViews/PrivilegeGroupMgr.vue"),
            },
            {
                path: 'ticket',
                component: () => import("@/views/Admin/pages/UserMgrViews/TicketMgr.vue"),
            },
            {
                path: 'order',
                component: () => import("@/views/Admin/pages/FinanceViews/OrderMgr.vue"),
            },
            {
                path: 'activation',
                component: () => import("@/views/Admin/pages/FinanceViews/ActivationMgr.vue"),
            },
            {
                path: 'key',
                component: () => import("@/views/Admin/pages/FinanceViews/KeyMgr.vue"),
            }

        ]
    },

];


export default adminRoutes;