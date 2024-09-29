<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, ref} from "vue";
import {useRouter} from 'vue-router'
import useThemeStore from "@/stores/useThemeStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import renderIcon from "@/utils/iconFormator";
import CommonLogo from "@/components/CommonLogo.vue";
import {
  BookOutline as manualIcon,
  CardOutline as paymentIcon,
  CartOutline as buyIcon,
  ChatboxEllipsesOutline as noticeIcon,
  ClipboardOutline as subscriptionIcon,
  CogOutline as settingIcon,
  ColorPaletteOutline as themeIcon,
  ConstructOutline as privilegeIcon,
  GiftOutline as ticketIcon,
  HelpBuoyOutline as helpIcon,
  HelpBuoyOutline as supportIcon,
  KeyOutline as activeIcon,
  LayersOutline as nodeIcon,
  LibraryOutline as knowledgeIcon,
  ListOutline as orderIcon,
  ListOutline as myOrderIcon,
  PeopleOutline as usersIcon,
  PeopleOutline as inviteIcon,
  PersonOutline as profileIcon,
  PodiumOutline as queueIcon,
  ShuffleOutline as routerIcon,
  SpeedometerOutline as dashboardIcon,
} from '@vicons/ionicons5'

const {t} = useI18n()
const themeStore = useThemeStore();
const userInfoStore = useUserInfoStore();

const router = useRouter();

// 管理员 ----------------------------------
let MenuOption = [
  {
    label: '仪表盘',
    key: 'dashboard',
    icon: renderIcon(dashboardIcon)
  },
  {
    label: '服务端监控',
    key: 'queue-monitor',
    icon: renderIcon(queueIcon)
  },
  {
    label: '设置',
    key: 'pinball-1973',
    // icon: renderIcon(settingIcon),
    disabled: false,
    children: [
      {
        label: '系统配置',
        key: 'system-config',
        icon: renderIcon(settingIcon),
      },
      {
        label: '支付设置',
        key: 'payment-config',
        icon: renderIcon(paymentIcon),
      },
      {
        label: '主题配置',
        key: 'theme-config',
        icon: renderIcon(themeIcon),
      },
    ]
  },
  {
    label: '服务器',
    key: 'server',
    // icon: renderIcon(BookIcon),
    disabled: false,
    children: [
      {
        label: '节点管理',
        key: 'node-manager',
        icon: renderIcon(nodeIcon),
      },
      {
        label: '权限组管理',
        key: 'privilege-group-mgr',
        icon: renderIcon(privilegeIcon),
      },
      {
        label: '路由管理',
        key: 'router-config',
        icon: renderIcon(routerIcon),
      },
    ]
  },
  {
    label: '财务',
    key: 'finance',
    // icon: renderIcon(BookIcon),
    disabled: false,
    children: [
      {
        label: '订阅管理',
        key: 'subscribe-manager',
        icon: renderIcon(subscriptionIcon),
      },
      {
        label: '订单管理',
        key: 'order-manager',
        icon: renderIcon(orderIcon),
      },
      {
        label: '优惠券管理',
        key: 'coupon-mgr',
        icon: renderIcon(ticketIcon),
      },
    ]
  },
  {
    label: '用户',
    key: 'users',
    // icon: renderIcon(BookIcon),
    disabled: false,
    children: [
      {
        label: '用户管理',
        key: 'user-manager',
        icon: renderIcon(usersIcon),
      },
      {
        label: '公告管理',
        key: 'notice-manager',
        icon: renderIcon(noticeIcon),
      },
      {
        label: '工单管理',
        key: 'prob-manager',
        icon: renderIcon(helpIcon),
      },
      {
        label: '知识库管理',
        key: 'doc-manager',
        icon: renderIcon(knowledgeIcon),
      },
    ]
  },

]

// let handleUpdateExpandedKeys = (keys: string[]) => {
//   console.info(`[onUpdate:expandedKeys]: ${JSON.stringify(keys)}`)
// }
let update = (key: string) => {
  console.log(key)
  switch (key) {
    case 'dashboard': {
      router.push({path: '/admin/dashboard/summary'})
      themeStore.menuSelected = 'dashboard'
      // themeStore.contentPath = '/admin/dashboard/summary'
      break
    }
    case 'queue-monitor': {
      themeStore.menuSelected = 'queue-monitor'
      router.push({path: '/admin/dashboard/monitor'})
      // themeStore.contentPath = '/admin/dashboard/monitor'
      break
    }
    case 'system-config': {
      themeStore.menuSelected = 'system-config'
      router.push({path: '/admin/dashboard/systemconfig'})
      break
    }
    case 'payment-config': {
      themeStore.menuSelected = 'payment-config'
      router.push({path: '/admin/dashboard/payment'})
      break
    }
    case 'theme-config': {
      themeStore.menuSelected = 'theme-config'
      router.push({path: '/admin/dashboard/theme'})
      break
    }

      // part2
    case 'router-config' : {
      themeStore.menuSelected = 'router-mgr'
      router.push({path: '/admin/dashboard/routermgr'})
      break
    }

      // part4
    case 'user-manager': {
      themeStore.menuSelected = 'user-manager'
      router.push({path: '/admin/dashboard/usermanager'})
      break
    }
    case 'notice-manager': {
      themeStore.menuSelected = 'notice-manager'
      router.push({path: '/admin/dashboard/noticemanager'})
      break
    }
    case 'subscribe-manager': {
      themeStore.menuSelected = 'subscribe-manager'
      router.push({path: '/admin/dashboard/subscribemanager'})
      break
    }
    case 'doc-manager': {
      themeStore.menuSelected = 'doc-manager'
      router.push({path: '/admin/dashboard/document'})
      break
    }
    case 'coupon-mgr': {
      // themeStore.menuSelected = 'publicNotice-manager'
      router.push({path: '/admin/dashboard/coupon'})
      break
    }
    case 'privilege-group-mgr': {
      themeStore.menuSelected = 'privilege-group-mgr'
      router.push({path: '/admin/dashboard/group'})
      break
    }
  }
}


// 普通用户 ----------------------------------
let UserMenuOption = ref([
  {
    label: computed(() => t('commonAside.user.dashboard')),
    key: 'user-dashboard',
    icon: renderIcon(dashboardIcon)
  },
  {
    label: computed(() => t('commonAside.user.document')),
    key: 'user-doc',
    icon: renderIcon(manualIcon)
  },
  {
    label: computed(() => t('commonAside.user.subscribe')),
    key: 'user-sub',
    // icon: renderIcon(settingIcon),
    disabled: false,
    children: [
      {
        label: computed(() => t('commonAside.user.purchase')),
        key: 'user-buy-plan',
        icon: renderIcon(buyIcon),
      },
      {
        label: computed(() => t('commonAside.user.surplus')),
        key: 'user-keys',
        icon: renderIcon(paymentIcon),
      },
    ]
  },
  {
    label: computed(() => t('commonAside.user.fiance')),
    key: 'user-fiance',
    // icon: renderIcon(settingIcon),
    disabled: false,
    children: [
      {
        label: computed(() => t('commonAside.user.myOrder')),
        key: 'user-orders',
        icon: renderIcon(myOrderIcon),
      },
      {
        label: computed(() => t('commonAside.user.myInvite')),
        key: 'user-invite',
        icon: renderIcon(inviteIcon),
      },
    ]
  },
  {
    label: computed(() => t('commonAside.user.user')),
    key: 'my-profile',
    // icon: renderIcon(settingIcon),
    disabled: false,
    children: [
      {
        label: computed(() => t('commonAside.user.profile')),
        key: 'user-profile',
        icon: renderIcon(profileIcon),
      },
      {
        label: computed(() => t('commonAside.user.support')),
        key: 'user-tickets',
        icon: renderIcon(supportIcon),
      },
      {
        label: computed(() => t('commonAside.user.activateLog')),
        key: 'user-active-record',
        icon: renderIcon(activeIcon),
      },
    ]
  },
])


let userUpdate = (key: string) => {
  switch (key) {
    case 'user-dashboard': {
      themeStore.menuSelected = 'user-dashboard'
      router.push({path: '/dashboard/summary'})
      break
    }
    case 'user-doc': {
      themeStore.menuSelected = 'user-doc'
      router.push({path: '/dashboard/document'})
      break
    }
    case 'user-buy-plan': {
      themeStore.menuSelected = 'user-buy-plan'
      router.push({path: '/dashboard/purchase'})
      break
    }
    case 'user-orders': {
      router.push({path: '/dashboard/orders'})
      break
    }


    case 'user-profile': {
      // themeStore.menuSelected = 'user-profile'
      router.push({path: '/dashboard/profile'})
      break
    }


    case 'user-tickets': {
      router.push({path: '/dashboard/tickets'})
      break
    }
    case 'user-keys': {
      router.push({path: '/dashboard/keys'})
      break
    }
    case 'user-invite': {
      router.push({path: '/dashboard/invite'})
      break
    }
  }
}

</script>

<script lang="ts">
export default {
  name: 'CommonAside',
}
</script>

<template>
  <div class="root">
    <CommonLogo class="logo"></CommonLogo>

<!--    <n-layout-sider-->
<!--        bordered-->
<!--        collapse-mode="width"-->
<!--        :collapsed-width="64"-->
<!--        :width="240"-->
<!--        :collapsed="collapsed"-->
<!--        show-trigger-->
<!--        @collapse="collapsed = true"-->
<!--        @expand="collapsed = false"-->
<!--    >-->

      <n-menu
          v-if="userInfoStore.thisUser.isAdmin"
          class="menu"
          :accordion="true"
          :options="MenuOption"
          @update:value="update"
          :value="themeStore.menuSelected"

      />
      <n-menu
          v-else
          class="menu"
          :default-expand-all="true"
          :options="UserMenuOption"
          @update:value="userUpdate"
          :value="themeStore.menuSelected"
          :root-indent="36"
          :indent="0"

      />

<!--    </n-layout-sider>-->


  </div>
</template>

<style lang="less" scoped>
.root {
  display: flex;
  flex-direction: column;
  .logo {
  }
  .menu {
    padding: 15px;
  }
}
</style>