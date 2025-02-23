<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, ref} from "vue";
import {useRouter} from 'vue-router'
import useAppInfosStore from "@/stores/useAppInfosStore";
import useThemeStore from "@/stores/useThemeStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import renderIcon from "@/utils/iconFormator";
import CommonLogo from "@/components/CommonLogo.vue";
import {type MenuOption} from "naive-ui"
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
  KeyOutline as keyIcon,
  LayersOutline as nodeIcon,
  LibraryOutline as knowledgeIcon,
  ListOutline as orderIcon,
  ListOutline as myOrderIcon,
  PeopleOutline as usersIcon,
  PeopleOutline as inviteIcon,
  PersonOutline as profileIcon,
  PodiumOutline as queueIcon,
  ShuffleOutline as routerIcon,
  TimeOutline as historyIcon,
  SpeedometerOutline as dashboardIcon,
  CloudDownloadOutline as downloadIcon,
} from '@vicons/ionicons5'

const {t} = useI18n()
const appInfoStore = useAppInfosStore()
const themeStore = useThemeStore();
const userInfoStore = useUserInfoStore();

const router = useRouter();

// 管理员 ----------------------------------
let MenuOption = ref<MenuOption[]>([
  {
    label: t('commonAside.admin.dashboard'),
    key: 'dashboard',
    icon: renderIcon(dashboardIcon),
  },
  {
    label: t('commonAside.admin.queueMonit'),
    key: 'queue-monitor',
    icon: renderIcon(queueIcon)
  },
  {
    label: t('commonAside.admin.settings'),
    key: 'pinball-1973',
    // icon: renderIcon(settingIcon),
    disabled: false,
    children: [
      {
        label: t('commonAside.admin.systemConfig'),
        key: 'system-config',
        icon: renderIcon(settingIcon),
      },
      {
        label: t('commonAside.admin.paymentConfig'),
        key: 'payment-config',
        icon: renderIcon(paymentIcon),
      },
      {
        label: t('commonAside.admin.themeConfig'),
        key: 'theme-config',
        icon: renderIcon(themeIcon),
      },
    ]
  },
  {
    label: t('commonAside.admin.server'),
    key: 'server',
    // icon: renderIcon(BookIcon),
    disabled: false,
    children: [
      // {
      //   label: '节点管理',
      //   key: 'node-manager',
      //   icon: renderIcon(nodeIcon),
      // },
      {
        label: t('commonAside.admin.privilege'),
        key: 'privilege-group-mgr',
        icon: renderIcon(privilegeIcon),
      },
      // {
      //   label: '路由管理',
      //   key: 'router-config',
      //   icon: renderIcon(routerIcon),
      // },
    ]
  },
  {
    label: t('commonAside.admin.finance'),
    key: 'finance',
    // icon: renderIcon(BookIcon),
    disabled: false,
    children: [
      {
        label: t('commonAside.admin.subscription'),
        key: 'subscription-manager',
        icon: renderIcon(subscriptionIcon),
      },
      {
        label: t('commonAside.admin.coupon'),
        key: 'coupon-mgr',
        icon: renderIcon(ticketIcon),
      },
      {
        label: t('commonAside.admin.order'),
        key: 'order-manager',
        icon: renderIcon(orderIcon),
      },
      {
        label: t('commonAside.admin.activate'),
        key: 'activate-manager',
        icon: renderIcon(historyIcon),
      },
      {
        label: t('commonAside.admin.key'),
        key: 'key-manager',
        icon: renderIcon(keyIcon),
      },
    ]
  },
  {
    label: t('commonAside.admin.user'),
    key: 'users',
    // icon: renderIcon(BookIcon),
    disabled: false,
    children: [
      {
        label: t('commonAside.admin.userMgr'),
        key: 'user-manager',
        icon: renderIcon(usersIcon),
      },
      {
        label: t('commonAside.admin.notice'),
        key: 'notice-manager',
        icon: renderIcon(noticeIcon),
      },
      {
        label: t('commonAside.admin.ticket'),
        key: 'ticket-mgr',
        icon: renderIcon(helpIcon),
      },
      {
        label: t('commonAside.admin.doc'),
        key: 'doc-manager',
        icon: renderIcon(knowledgeIcon),
      },
    ]
  },
])


let adminUpdate = (key: string) => {
  themeStore.menuIsFlippActive = false
  switch (key) {
    case 'dashboard': return router.push({path: '/admin/dashboard/summary'})
    case 'queue-monitor': return router.push({path: '/admin/dashboard/monitor'})
    case 'system-config': return router.push({path: '/admin/dashboard/systemconfig'})
    case 'payment-config': return router.push({path: '/admin/dashboard/payment'})
    case 'theme-config': return router.push({path: '/admin/dashboard/theme'})
    case 'router-config' : return router.push({path: '/admin/dashboard/routermgr'})
    case 'user-manager': return router.push({path: '/admin/dashboard/usermanager'})
    case 'notice-manager': return router.push({path: '/admin/dashboard/noticemanager'})
    case 'subscription-manager': return router.push({path: '/admin/dashboard/subscribemanager'})
    case 'doc-manager':return router.push({path: '/admin/dashboard/document'})
    case 'coupon-mgr':return router.push({path: '/admin/dashboard/coupon'})
    case 'privilege-group-mgr':return router.push({path: '/admin/dashboard/group'})
    case 'ticket-mgr':return router.push({path: '/admin/dashboard/ticket'})
    case 'order-manager':return router.push({path: '/admin/dashboard/order'})
    case 'activate-manager': return router.push({path: '/admin/dashboard/activation'})
    case 'key-manager': return router.push({path: '/admin/dashboard/key'})
  }
}


// 普通用户 ----------------------------------
let UserMenuOption = ref<MenuOption[]>([
  {
    label: t('commonAside.user.dashboard'),
    key: 'user-dashboard',
    icon: renderIcon(dashboardIcon)
  },
  {
    label: t('commonAside.user.document'),
    key: 'user-doc',
    icon: renderIcon(manualIcon)
  },
  {
    label: t('commonAside.user.app'),
    key: 'user-app-download',
    icon: renderIcon(downloadIcon)
  },
  {
    label: t('commonAside.user.subscription'),
    key: 'user-sub',
    disabled: false,
    children: [
      {
        label: t('commonAside.user.purchase'),
        key: 'user-buy-plan',
        icon: renderIcon(buyIcon),
      },
      {
        label: t('commonAside.user.surplus'),
        key: 'user-keys',
        icon: renderIcon(keyIcon),
      },
      {
        label: t('commonAside.user.activateLog'),
        key: 'user-activate-log',
        icon: renderIcon(historyIcon),
      },
    ]
  },
  {
    label: t('commonAside.user.fiance'),
    key: 'user-fiance',
    disabled: false,
    children: [
      {
        label: t('commonAside.user.topUp'),
        key: 'user-top-up',
        icon: renderIcon(paymentIcon),
      },
      {
        label: t('commonAside.user.myOrder'),
        key: 'user-orders',
        icon: renderIcon(myOrderIcon),
      },
      {
        label: t('commonAside.user.myInvite'),
        key: 'user-invite',
        icon: renderIcon(inviteIcon),
      },
    ]
  },
  {
    label: t('commonAside.user.user'),
    key: 'my-profile',
    disabled: false,
    children: [
      {
        label: t('commonAside.user.profile'),
        key: 'user-profile',
        icon: renderIcon(profileIcon),
      },
      {
        label: t('commonAside.user.support'),
        key: 'user-tickets',
        icon: renderIcon(supportIcon),
      },

    ]
  },
])

let userUpdate = (key: string) => {
  themeStore.menuIsFlippActive = false
  switch (key) {
    case 'user-dashboard': return router.push({path: '/dashboard/summary'})
    case 'user-doc': return router.push({path: '/dashboard/document'})
    case 'user-app-download': return router.push({path: '/dashboard/app'})
    case 'user-buy-plan': return router.push({path: '/dashboard/purchase'})
    case 'user-orders': return router.push({path: '/dashboard/orders'})
    case 'user-profile': return router.push({path: '/dashboard/profile'})
    case 'user-top-up': return router.push({path: '/dashboard/topup'})
    case 'user-tickets': return router.push({path: '/dashboard/tickets'})
    case 'user-keys': return router.push({path: '/dashboard/keys'})
    case 'user-invite': return router.push({path: '/dashboard/invite'})
    case 'user-activate-log': return router.push({path: '/dashboard/log'})
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
    <n-scrollbar style="max-height: calc(100vh - 52px)" :size="0">
      <n-menu
          v-if="userInfoStore.thisUser.isAdmin"
          class="menu"
          :accordion="false"
          :default-expand-all="true"
          :options="MenuOption"
          @update:value="adminUpdate"
          :value="themeStore.menuSelected"
          v-model="themeStore.menuSelected"
      />
      <n-menu
          v-else
          class="menu"
          :default-expand-all="true"
          :options="UserMenuOption"
          @update:value="userUpdate"
          :value="themeStore.menuSelected"
          v-model="themeStore.menuSelected"
          :root-indent="36"
          :indent="0"
      />
    </n-scrollbar>
    <p class="left-suffix">{{ appInfoStore.appCommonConfig.app_name }} {{ appInfoStore.appVersion }}</p>

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
  .left-suffix {
    font-size: 0.7rem;
    font-weight: bold;
    opacity: 0.2;
    position: absolute;
    left: 20px;
    bottom: 8px;
  }
}
</style>