import {createRouter, createWebHistory, type RouteRecordRaw} from 'vue-router'
import useUserInfoStore from '@/stores/useUserInfoStore'

import ErrorPage from "@/views/ErrorPage.vue";

import adminRoutes from "@/router/admin";
import userRoutes from "@/router/user"

const routes: RouteRecordRaw[] = [
    ...adminRoutes,
    ...userRoutes,
    {
        path: '/error',
        component: ErrorPage,
    }
]

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes,
})


// router.beforeEach((to, from, next) => {
//     console.log(from.path, to.path);
//
//     const userInfoStore = useUserInfoStore();
//
//     // 检查用户是否已登录
//     if (!userInfoStore.isAuthed) {
//         // 未登录用户
//         if (to.path.startsWith('/admin')) {
//             if (to.path !== '/admin/login') {
//                 return next('/admin/login');
//             } else {
//                 return next();
//             }
//         } else if (['/register', '/welcome', '/welcome2', '/dashboard/ticket/chat'].includes(to.path)) {
//             // 允许未登录用户访问注册页面、欢迎页面、和聊天页面
//             return next();
//         } else {
//             // 未登录用户重定向到普通用户登录页面
//             if (to.path !== '/login') {
//                 return next('/login');
//             } else {
//                 return next();
//             }
//         }
//     } else {
//         // 用户已登录
//         if (to.path.startsWith('/admin')) {
//             if (!userInfoStore.thisUser.isAdmin) {
//                 // 非管理员用户禁止访问 /admin 路径，重定向到普通用户的 dashboard
//                 if (to.path !== '/dashboard/summary') {
//                     return next('/dashboard/summary');
//                 } else {
//                     return next();
//                 }
//             } else {
//                 if (to.path === '/admin/login') {
//                     // 管理员已经登录，避免再次访问登录页面，重定向到 /admin/dashboard/summary
//                     if (to.path !== '/admin/dashboard/summary') {
//                         return next('/admin/dashboard/summary');
//                     } else {
//                         return next();
//                     }
//                 } else {
//                     return next();
//                 }
//             }
//         } else if (to.path === '/login') {
//             // 已登录用户访问普通用户登录页面，重定向到 dashboard
//             if (to.path !== '/dashboard/summary') {
//                 return next('/dashboard/summary');
//             } else {
//                 return next();
//             }
//         } else if (to.path === '/user/tickets/chat') {
//             // 允许已登录用户访问聊天页面
//             return next();
//         } else {
//             return next(); // 正常放行
//         }
//     }
// });

router.beforeEach((to, from, next) => {
    console.log(from.path, to.path);

    const userInfoStore = useUserInfoStore();

    // 检查用户是否已登录
    if (!userInfoStore.isAuthed) {
        // 未登录用户处理逻辑
        if (to.path.startsWith('/admin')) {
            // 未登录用户访问 /admin，重定向到管理员登录页面
            return to.path === '/admin/login' ? next() : next('/admin/login');
        }

        // 允许未登录用户访问的路径
        const publicPaths = ['/register', '/welcome', '/welcome2', '/dashboard/ticket/chat'];
        if (publicPaths.includes(to.path)) {
            return next();
        }

        // 未登录用户重定向到普通登录页面
        return to.path === '/login' ? next() : next('/login');
    }

    // 已登录用户处理逻辑
    if (to.path.startsWith('/admin')) {
        // 非管理员禁止访问 /admin
        if (!userInfoStore.thisUser.isAdmin) {
            return to.path === '/dashboard/summary' ? next() : next('/dashboard/summary');
        }

        // 管理员登录后禁止访问登录页面，重定向到管理员 dashboard
        return to.path === '/admin/login' ? next('/admin/dashboard/summary') : next();
    }

    // 已登录用户禁止访问普通用户登录页面，重定向到用户 dashboard
    if (to.path === '/login') {
        return next('/dashboard/summary');
    }

    // 允许已登录用户访问的特殊页面
    if (to.path === '/user/tickets/chat') {
        return next();
    }

    // 默认放行
    return next();
});


// router.beforeEach(async (to, from) => {
//   console.log(from.path, to.path)
//
//   const userInfoStore = useUserInfoStore()
//   console.log('鉴权状态', userInfoStore.isAuthed)
//   if (userInfoStore.isAuthed) {
//     if (to.path !== '/admin/dashboard') {
//       return {path: '/admin/dashboard'}
//     }
//
//   } else {
//
//   }
//   //
//   // if (userInfoStore.isAuthed) {
//   //   console.log('允许访问')
//   // if (userInfoStore.isAuthed) {
//   //   console.log('允许访问')
//   //   console.log('用户角色是否是管理员', userInfoStore.thisUser.isAdmin)
//   //   if (userInfoStore.thisUser.isAdmin && to.path !== '/admin/dashboard') {
//   //     console.log('前往管理员面板')
//   //     return {
//   //       path: '/admin/dashboard'
//   //     }
//   //   } else {
//   //     console.log('非管理员')
//   //     return {
//   //       path: '/admin/login'
//   //     }
//   //   }
//   // } else {
//   //   console.log('禁止访问')
//   //   // return {
//   //   //   path: '/login'
//   //   // }
//   // }
//   //   console.log('用户角色是否是管理员', userInfoStore.thisUser.isAdmin)
//   //   if (userInfoStore.thisUser.isAdmin && to.path !== '/admin/dashboard') {
//   //     console.log('前往管理员面板')
//   //     return {
//   //       path: '/admin/dashboard'
//   //     }
//   //   } else {
//   //     console.log('非管理员')
//   //     return {
//   //       path: '/admin/login'
//   //     }
//   //   }
//   // } else {
//   //   console.log('禁止访问')
//   //   // return {
//   //   //   path: '/login'
//   //   // }
//   // }
// });

export default router
