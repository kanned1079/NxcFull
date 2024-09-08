import { createRouter, createWebHistory, createWebHashHistory, type RouteRecordRaw } from 'vue-router'
// import DashBoard from "@/views/Admin/DashBoard.vue";
// import Summary from "@/views/Admin/Summary.vue";
// import UserLogin from '@/views/User/Login/UserLogin.vue'  // 普通用户的登录窗口
// import AdminLogin from "@/views/Admin/Login/AdminLogin.vue";  // 管理员的登录窗口
// import QueueMonitor from "@/views/Admin/QueueMonitor.vue";
// import SystemConfig from "@/views/Admin/SystemConfig.vue";
// import PaymentConfig from "@/views/Admin/PaymentConfig.vue";
// import ThemeConfig from "@/views/Admin/ThemeConfig.vue";
// import UserManager from "@/views/Admin/UserManager.vue";

import useUserInfoStore from '@/stores/useUserInfoStore'
import useThemeStore from '@/stores/useThemeStore'

import adminRoutes from "@/router/admin";
import userRoutes from "@/router/user"

const routes: RouteRecordRaw[] = [
    ...adminRoutes,
    ...userRoutes
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

router.beforeEach((to, from, next) => {
  console.log(from.path, to.path)
  // const themeStore = useThemeStore()
  const userInfoStore = useUserInfoStore()
  // themeStore.contentPath = to.path
  console.log('跳转到登录页?', to.meta.requireAuth && !userInfoStore.isAuthed)
  if (to.meta.requireAuth && !userInfoStore.isAuthed) {// 判断用户是否已登录
    next('/admin/login'); // 跳转到登录页
  } else {
    next(); // 放行
    // console.log('存储当前页面路径', to.path)
    // themeStore.contentPath = to.path
  }
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
