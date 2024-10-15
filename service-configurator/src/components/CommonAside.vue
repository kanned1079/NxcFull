<script setup lang="ts">
import {ref, h, type Component} from "vue"
import { NIcon } from 'naive-ui'
import type { MenuOption } from 'naive-ui'
import mysqlIcon from "@/assets/mysql-line.svg"
import redisIcon from "@/assets/redis-line.svg"
import etcdIcon from "@/assets/icon_etcd.svg"
import apiIcon from "@/assets/api-line.svg"
import {
  BookOutline as BookIcon,
  PersonOutline as PersonIcon,
  WineOutline as WineIcon
} from '@vicons/ionicons5'


const menuOptions: MenuOption[] = [
  {
    label: 'Etcd',
    key: 'etcd',
    icon: renderIcon(etcdIcon)
  },
  {
    label: 'MySQL',
    key: 'mysql',
    icon: renderIcon(mysqlIcon)
  },
  {
    label: 'Redis',
    key: 'redis',
    icon: renderIcon(redisIcon)
  },
  {
    label: 'API接口',
    key: 'api',
    icon: renderIcon(apiIcon)
  },

]



function renderIcon(icon: Component) {
  return () => h(NIcon, null, { default: () => h(icon) })
}

let activeKey = ref<string>('')



</script>

<script lang="ts">
export default {
  name: 'CommonAside',
}
</script>

<template>
  <div class="root">
    <n-card :embedded="true" :bordered="false" hoverable title="配置项" class="aside-card" content-style="padding: 5px">
      <n-menu
          v-model:value="activeKey"
          :root-indent="36"
          :indent="12"
          :options="menuOptions"
      />
    </n-card>
<!--    <n-icon><mysqlIcon/></n-icon>-->
  </div>
  
</template>

<style scoped lang="less">
.root {
  margin: 20px;
  .aside-card {
    height: calc(100vh - 110px);
    background-color: #e3e5e7;
  }
}
</style>