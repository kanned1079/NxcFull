import {type RouteRecordRaw} from 'vue-router';
import DashBoard from "@/views/Admin/DashBoard.vue";
import Summary from "@/views/Admin/pages/RootViews/Summary.vue";
import QueueMonitor from "@/views/Admin/pages/RootViews/QueueMonitor.vue";
import SystemConfig from "@/views/Admin/pages/SettingViews/SystemConfig.vue";
import PaymentConfig from "@/views/Admin/pages/SettingViews/PaymentConfig.vue";
import ThemeConfig from "@/views/Admin/pages/SettingViews/ThemeConfig.vue";
import UserManager from "@/views/Admin/pages/UserMgrViews/UserManager.vue";
import AdminLogin from "@/views/Admin/Login/AdminLogin.vue";
import RouterMgr from "@/views/Admin/pages/ServerViews/RouterMgr.vue"
import PrivilegeGroup from "@/views/Admin/pages/ServerViews/PrivilegeGroupMgr.vue"
import CouponMgr from "@/views/Admin/pages/FinanceViews/CouponMgr.vue";
import NoticeManager from "@/views/Admin/pages/UserMgrViews/NoticeManager.vue";
import SubscribeMgr from "@/views/Admin/pages/FinanceViews/SubscribeMgr.vue";
import DocumentMgr from "@/views/Admin/pages/UserMgrViews/DocumentMgr.vue";
import PrivilegeGroupMgr from "@/views/Admin/pages/ServerViews/PrivilegeGroupMgr.vue";
import TicketMgr from "@/views/Admin/pages/UserMgrViews/TicketMgr.vue";
import OrderMgr from "@/views/Admin/pages/FinanceViews/OrderMgr.vue"
import ActivationMgr from "@/views/Admin/pages/FinanceViews/ActivationMgr.vue";
import keyMgr from "@/views/Admin/pages/FinanceViews/KeyMgr.vue";

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
                path: 'summary',
                name: 'summary',
                component: Summary,
            },
            {
                path: 'monitor',
                name: 'monitor',
                component: QueueMonitor
            },
            {
                path: 'systemconfig',
                name: 'system-config',
                component: SystemConfig,
            },
            {
                path: 'payment',
                name: 'payment',
                component: PaymentConfig,
            },
            {
                path: 'theme',
                name: 'theme',
                component: ThemeConfig,
            },
            {
                path: 'node',
                name: 'node',
                component: ThemeConfig,
            },

            // part4
            {
                path: 'usermanager',
                name: 'user-manager',
                component: UserManager,
            },

            {
                path: 'routermgr',
                name: 'router-mgr',
                component: RouterMgr
            },
            {
                path: 'privilegegroup',
                name: 'privilege-mgr',
                component: PrivilegeGroup,
            },
            {
                path: 'noticemanager',
                name: 'notice-manager',
                component: NoticeManager,
            },
            {
                path: 'subscribemanager',
                name: 'subscription-manager',
                component: SubscribeMgr,
            },
            {
                path: 'document',
                name: 'document-mgr',
                component: DocumentMgr,
            },
            {
                path: 'coupon',
                name: 'coupon',
                component: CouponMgr,
            },
            {
                path: 'group',
                name: 'privilege-group-mgr',
                component: PrivilegeGroupMgr,
            },
            {
                path: 'ticket',
                component: TicketMgr,
            },
            {
                path: 'order',
                component: OrderMgr,
            },
            {
                path: 'activation',
                component: ActivationMgr,
            },
            {
                path: 'key',
                component: keyMgr,
            }

        ]
    },
    // {
    //     path: '/dashboard',
    //     component: DashBoard,
    // },
    {
        path: '/admin/login',
        name: 'admin-login',
        component: AdminLogin
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
    },

];

export default adminRoutes;