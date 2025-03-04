import {createRouter, createWebHistory, type RouteRecordRaw} from 'vue-router'


// import adminRoutes from "@/router/admin";
// import userRoutes from "@/router/user"

const routes: RouteRecordRaw[] = [
  // ...adminRoutes,
  // ...userRoutes,
  // {
  //     path: '/error',
  //     component: ErrorPage,
  // },
  // {
  //   path: '/:pathMatch(.*)*', // 使用 pathMatch(.*)* 来匹配任何未定义的路径
  //   name: 'NotFound',
  //   component: () => import("@/views/utils/NotFound.vue"),
  // },
  {
    path: '/',
    component: () => import("./../views/Login.vue"),
  },
  {
    path: '/home',
    component: () => import("./../views/Home.vue")

  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

export default router
