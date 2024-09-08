<script setup lang="ts" name="CommonAside">
import {useRouter} from 'vue-router'
import useThemeStore from "@/stores/useThemeStore";
import renderIcon from "@/utils/iconFormator";
import {
  CardOutline as paymentIcon,
  ChatboxEllipsesOutline as noticeIcon,
  ClipboardOutline as subscriptionIcon,
  CogOutline as settingIcon,
  ColorPaletteOutline as themeIcon,
  ConstructOutline as privilegeIcon,
  GiftOutline as ticketIcon,
  HelpBuoyOutline as helpIcon,
  LayersOutline as nodeIcon,
  LibraryOutline as knowledgeIcon,
  ListOutline as orderIcon,
  PeopleOutline as usersIcon,
  PodiumOutline as queueIcon,
  ShuffleOutline as routerIcon,
  SpeedometerOutline as dashboardIcon,
} from '@vicons/ionicons5'

const themeStore = useThemeStore();

const router = useRouter();
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
        key: 'privilege-manager',
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
        key: 'ticket-manager',
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
        key: 'knowledge-manager',
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
  }
}
</script>

<template>
  <div class="root">
    <n-menu
        class="menu"
        :accordion="true"
        :options="MenuOption"
        @update:value="update"
        :value="themeStore.menuSelected"

    />
  </div>
</template>

<style lang="less" scoped>
.root {
  padding: 6px;

}
</style>