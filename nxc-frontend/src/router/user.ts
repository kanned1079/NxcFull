import { type RouteRecordRaw } from 'vue-router';


const userRoutes: RouteRecordRaw[] = [
    {
        path: '/login',
        name: 'login',
        component: () => import("@/views/User/Login/UserLogin.vue")
    },
    {
        path: '/register',
        name: 'register',
        component: () => import("@/views/User/Login/UserRegister.vue")
    },
    {
        path: '/home',
        redirect: '/',
    },
    {
        path: '/',
        redirect: '/welcome',
    },
    {
        path: '/welcome',
        component: () => import("@/views/Welcome/A/WhyChooseUsA.vue"),
    },
    {
        path: '/welcome2',
        component: () => import("@/views/Welcome/B/WhyChooseUsB.vue"),
    },
    {
        path: '/user/tickets/chat',
        component: () => import("@/views/utils/ChatDialog.vue"),
    },
    {
        path: '/dashboard',
        component: () => import("@/views/User/UserDashBoard.vue"),
        children: [
            {
                path: 'summary', // 去掉前导斜杠
                component: () => import("@/views/User/pages/RootViews/UserSummary.vue"),
            },
            {
                path: 'document',
                component: () => import("@/views/User/pages/RootViews/UserDocument.vue"),
            },
            {
                path: 'app',
                component: () => import("@/views/User/pages/RootViews/UserAppDownload.vue"),
            },
            {
                path: 'purchase',
                component: () => import("@/views/User/pages/PurchaseViews/NewPurchase.vue"),
            },
            {
                path: 'purchase/settlement',
                component: () => import("@/views/User/pages/PurchaseViews/NewSettlement.vue"),
            },
            {
                path: 'purchase/confirm',
                component: () => import("@/views/User/pages/PurchaseViews/ConfirmOrder.vue"),
            },
            {
                path: 'orders/details',
                component: () => import("@/views/User/pages/PurchaseViews/PaymentResult/PaymentResult.vue"),
            },
            {
                path: 'profile',
                component: () => import("@/views/User/pages/UserViews/UserProfile.vue"),
            },
            {
                path: 'topup',
                component: () => import("@/views/User/pages/FinanceViews/UserTopUp.vue"),
            },
            {
                path: 'tickets',
                component: () => import("@/views/User/pages/UserViews/Tickets.vue"),
            },

            {
                path: 'keys',
                component: () => import("@/views/User/pages/PurchaseViews/UserKeys.vue")
            },
            {
                path: 'orders',
                component: () => import("@/views/User/pages/FinanceViews/UserOrders.vue")
            },
            {
                path: 'invite',
                component: () => import("@/views/User/pages/FinanceViews/UserInvite.vue")
            },
            {
                path: 'log',
                component: () => import("@/views/User/pages/UserViews/ActivateLog.vue")
            },
        ]
    }
];

export default userRoutes;