import { createRouter, createWebHistory, createWebHashHistory, type RouteRecordRaw } from 'vue-router'
// import DashBoard from "@/views/Admin/DashBoard.vue";
// import Summary from "@/views/Admin/UserSummary.vue";
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


// ok1
// router.beforeEach((to, from, next) => {
//   console.log(from.path, to.path)
//   // const themeStore = useThemeStore()
//   const userInfoStore = useUserInfoStore()
//   // themeStore.contentPath = to.path
//   console.log('跳转到登录页?', to.meta.requireAuth && !userInfoStore.isAuthed)
//   if (to.meta.requireAuth && !userInfoStore.isAuthed) {// 判断用户是否已登录
//     // userInfoStore.isAdmin 此字段默认是false 如果使用此字段判断则进入不了管理员登录界面 访问/admin时会被定向到/login
//     next('/admin/login'); // 跳转到登录页
//   } else {
//     next(); // 放行
//   }
// });

// v2
// router.beforeEach((to, from, next) => {
//   console.log(from.path, to.path);
//
//   const userInfoStore = useUserInfoStore();
//
//   // 检查是否已登录
//   if (!userInfoStore.isAuthed) {
//     // 用户未登录
//     if (to.path.startsWith('/admin')) {
//       // 如果当前路径已经是 /admin/login，避免无限循环重定向
//       if (to.path !== '/admin/login') {
//         next('/admin/login');
//       } else {
//         next();
//       }
//     } else {
//       // 非管理员路径，重定向到普通用户登录页面，但避免重定向到当前页面
//       if (to.path !== '/login') {
//         next('/login');
//       } else {
//         next();
//       }
//     }
//   } else {
//     // 用户已登录
//     if (to.path.startsWith('/admin')) {
//       // 如果是访问/admin相关的路径
//       if (!userInfoStore.thisUser.isAdmin) {
//         // 非管理员用户禁止访问/admin路径，重定向到普通用户的 dashboard，但避免重定向到当前页面
//         if (to.path !== '/dashboard/summary') {
//           next('/dashboard/summary');
//         } else {
//           next();
//         }
//       } else {
//         // 管理员用户访问/admin相关的路径，避免无限循环
//         if (to.path === '/admin/login') {
//           // 管理员已经登录，避免再次访问登录页面，重定向到/admin/dashboard/summary
//           if (to.path !== '/admin/dashboard/summary') {
//             next('/admin/dashboard/summary');
//           } else {
//             next();
//           }
//         } else {
//           next(); // 正常放行
//         }
//       }
//     } else {
//       // 非/admin路径，且已登录为普通用户
//       if (to.path === '/login') {
//         // 避免已登录用户访问登录页面，重定向到普通用户的 dashboard
//         if (to.path !== '/dashboard/summary') {
//           next('/dashboard/summary');
//         } else {
//           next();
//         }
//       } else {
//         next(); // 正常放行
//       }
//     }
//   }
// });

router.beforeEach((to, from, next) => {
  console.log(from.path, to.path);

  const userInfoStore = useUserInfoStore();

  // 检查用户是否已登录
  if (!userInfoStore.isAuthed) {
    // 未登录用户
    if (to.path.startsWith('/admin')) {
      // 访问 /admin 路径时，未登录用户重定向到管理员登录页，但如果是访问 /admin/login 则放行
      if (to.path !== '/admin/login') {
        next('/admin/login');
      } else {
        next();
      }
    } else if (to.path === '/register') {
      // 允许未登录用户访问注册页面
      next();
    } else {
      // 非管理员路径，且不是注册页面，重定向到普通用户登录页面
      if (to.path !== '/login') {
        next('/login');
      } else {
        next();
      }
    }
  } else {
    // 用户已登录
    if (to.path.startsWith('/admin')) {
      // 如果是访问 /admin 路径
      if (!userInfoStore.thisUser.isAdmin) {
        // 非管理员用户禁止访问 /admin 路径，重定向到普通用户的 dashboard
        if (to.path !== '/dashboard/summary') {
          next('/dashboard/summary');
        } else {
          next();
        }
      } else {
        // 管理员用户访问 /admin 相关的路径
        if (to.path === '/admin/login') {
          // 管理员已经登录，避免再次访问登录页面，重定向到 /admin/dashboard/summary
          if (to.path !== '/admin/dashboard/summary') {
            next('/admin/dashboard/summary');
          } else {
            next();
          }
        } else {
          next(); // 正常放行
        }
      }
    } else {
      // 已登录用户访问非管理员路径
      if (to.path === '/login') {
        // 已登录用户访问普通用户登录页面，重定向到 dashboard
        if (to.path !== '/dashboard/summary') {
          next('/dashboard/summary');
        } else {
          next();
        }
      } else {
        next(); // 正常放行
      }
    }
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
