<script setup lang="ts">
import {type Component, h} from "vue"
import {useRouter} from "vue-router";
import type {MenuOption} from 'naive-ui'
import {NIcon} from 'naive-ui'
import mysqlIcon from "@/assets/mysql-line.svg"
import redisIcon from "@/assets/redis-line.svg"
import etcdIcon from "@/assets/icon_etcd.svg"
import apiIcon from "@/assets/api-line.svg"
import mqIcon from "@/assets/rabbitmq.svg"
import useThemeStore from "@/store/useThemeStore";

const themeStore = useThemeStore()
const router = useRouter()

let renderIcon = (icon: Component) => {
  return () =>
      h('div', {style: {marginRight: '20px', display: 'flex', alignItems: 'center'}}, [
        h(NIcon, null, {default: () => h(icon)})
      ])
}

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
    label: 'RabbitMQ',
    key: 'mq',
    icon: renderIcon(mqIcon)
  },
  {
    label: 'API接口',
    key: 'api',
    icon: renderIcon(apiIcon)
  },

]

let handleSelected = () => {
  console.log('to: ', themeStore.path)
  themeStore.path = `/${themeStore.menuSelected}`
  router.push({
    path: themeStore.path
  })
}


</script>

<script lang="ts">
export default {
  name: 'CommonAside',
}
</script>

<template>
  <div class="root">
    <n-card :embedded="true" :bordered="true" title="配置项" class="aside-card" content-style="padding: 5px">
      <n-menu
          v-model:value="themeStore.menuSelected"
          :root-indent="36"
          :indent="12"
          :options="menuOptions"
          @update:value="handleSelected"
      />
    </n-card>
    <!--    <n-icon><mysqlIcon/></n-icon>-->
  </div>

</template>

<style scoped lang="less">
.root {
  width: 240px;

  .aside-card {
    height: calc(100vh - 90px);
    background-color: #f4f4f4;
  }
}
</style>