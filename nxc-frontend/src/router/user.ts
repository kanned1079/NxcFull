import { type RouteRecordRaw } from 'vue-router';
import UserLogin from "@/views/User/Login/UserLogin.vue";
import UserRegister from "@/views/User/Login/UserRegister.vue"
import CommonAside from "@/components/CommonAside.vue";
import UserDashBoard from "@/views/User/UserDashBoard.vue";
import UserSummary from "@/views/User/pages/RootViews/UserSummary.vue";
import UserDocument from "@/views/User/pages/RootViews/UserDocument.vue";
import NewPurchase from "@/views/User/pages/PurchaseViews/NewPurchase.vue";
import NewSettlement from "@/views/User/pages/PurchaseViews/NewSettlement.vue";
import UserProfile from "@/views/User/pages/Profile/UserProfile.vue";
import Tickets from "@/views/User/pages/UserViews/Tickets.vue"
import UserKeys from "@/views/User/pages/PurchaseViews/UserKeys.vue";
import UserInvite from "@/views/User/pages/FinanceViews/UserInvite.vue";
import UserOrders from "@/views/User/pages/FinanceViews/UserOrders.vue";
import ConfirmOrder from "@/views/User/pages/PurchaseViews/ConfirmOrder.vue";

// WhyChooseUs页面
import WhyChooseUsA from "@/views/Welcome/A/WhyChooseUsA.vue";
import WhyChooseUsB from "@/views/Welcome/B/WhyChooseUsB.vue";

// test
// import ThemeEditor from "@/views/ThemeEditor.vue";

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
        redirect: '/welcome',
    },
    {
        path: '/welcome',
        component: WhyChooseUsA,
    },
    {
        path: '/welcome2',
        component: WhyChooseUsB,
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
            },
            {
                path: '/dashboard/purchase',
                component: NewPurchase,
            },
            {
                path: '/dashboard/purchase/settlement',
                component: NewSettlement,
            },
            {
                path: '/dashboard/purchase/confirm',
                component: ConfirmOrder,
            },
            {
                path: '/dashboard/profile',
                component: UserProfile,
            },
            {
                path: '/dashboard/tickets',
                component: Tickets,
            },
            // {
            //     path: '/dashboard/themeeditor',
            //     component: ThemeEditor,
            // },
            {
                path: '/dashboard/keys',
                component: UserKeys
            },
            {
                path: '/dashboard/orders',
                component: UserOrders
            },
            {
                path: '/dashboard/invite',
                component: UserInvite
            }

        ]
    }

];

export default userRoutes;