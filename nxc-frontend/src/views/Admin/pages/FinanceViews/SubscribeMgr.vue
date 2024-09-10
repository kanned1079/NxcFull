<script setup lang="ts" name="SubscribeMgr">
import {ref, reactive, onMounted} from "vue";
import useThemeStore from "@/stores/useThemeStore";
import instance from "@/axios";
const themeStore = useThemeStore();

let handleAddSubscribe = () => {
  console.log('处理添加一个订阅')
}

let getAllSubscribeItems = async () => {
  let {data} = await instance.get('/api/admin/get-all-subscribes')
  console.log('获取所有的订阅', data)


}

let aaa = () => {

}

onMounted(() => {
  console.log('SubscribeMgr挂载')
  getAllSubscribeItems()

  themeStore.menuSelected = 'subscribe-manager'
  themeStore.contentPath = '/admin/dashboard/subscribemanager'

  // themeStore.
})

</script>

<template>
<div class="root">
  <n-card class="card" hoverable :embedded="true" title="订阅管理">
    <n-button class="add-btn" @click="handleAddSubscribe">添加订阅</n-button>
  </n-card>
</div>
</template>

<style scoped lang="less">
.root {
  padding: 20px;
  .card {
    .add-btn {
      margin-top: 10px;
    }
  }
}
</style>