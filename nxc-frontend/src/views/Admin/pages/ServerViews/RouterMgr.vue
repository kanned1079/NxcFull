<script setup lang="ts">
import {defineComponent, onMounted} from 'vue'
import useThemeStore from "@/stores/useThemeStore";
const themeStore = useThemeStore()

let handleAddRouter = () => {

}

onMounted(() => {
  themeStore.menuSelected = 'router-config'
  themeStore.contentPath = '/admin/dashboard/routermgr'
})

</script>

<script lang="ts">
export default {
  name: 'RouterMgr'
}
</script>

<template>
  <div class="root">
    <n-card hoverable :embedded="true" title="路由管理" :bordered="false">
      <n-button type="primary" :bordered="false" class="add-btn" @click="handleAddRouter">添加路由</n-button>
    </n-card>
  </div>
</template>

<style lang="less" scoped>
.root {
  padding: 20px;
  min-width: 900px;
  .add-btn {
    margin-top: 10px;
  }
}

</style>