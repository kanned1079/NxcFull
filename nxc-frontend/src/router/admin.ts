import {type RouteRecordRaw} from 'vue-router';
import DashBoard from "@/views/Admin/DashBoard.vue";
import Summary from "@/views/Admin/pages/RootViews/Summary.vue";
import QueueMonitor from "@/views/Admin/pages/RootViews/QueueMonitor.vue";
import SystemConfig from "@/views/Admin/pages/SettingViews/SystemConfig.vue";
import PaymentConfig from "@/views/Admin/pages/SettingViews/SystemConfig/PaymentConfig.vue";
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
            },
            {
                path: '/admin/dashboard/noticemanager',
                name: 'notice-manager',
                component: NoticeManager,
            },
            {
                path: '/admin/dashboard/subscribemanager',
                name: 'subscribe-manager',
                component: SubscribeMgr,
            },
            {
                path: '/admin/dashboard/document',
                name: 'document-mgr',
                component: DocumentMgr,
            },
            {
                path: '/admin/dashboard/publicNotice',
                name: 'publicNotice-mgr',
                component: CouponMgr,
            },
            {
                path: '/admin/dashboard/group',
                name: 'privilege-group-mgr',
                component: PrivilegeGroupMgr,
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