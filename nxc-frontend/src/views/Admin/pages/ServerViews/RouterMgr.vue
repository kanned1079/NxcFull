<script setup lang="ts" name="RouterMgr">
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

<template>
  <div class="root">
    <n-card hoverable :embedded="true" title="路由管理">
      <n-button class="add-btn" @click="handleAddRouter">添加路由</n-button>
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
.n-card {
  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
  padding: 0;
  border: 0;
}
</style>