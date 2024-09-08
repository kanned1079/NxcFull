<script setup lang="ts" name="PrivilegeGroup">
import {defineComponent} from 'vue'
import useThemeStore from "@/stores/useThemeStore";
const themeStore = useThemeStore()

let handleAddGroup = () => {

}

</script>

<template>
  <div class="root">
    <n-card hoverable :embedded="true" title="权限组管理">
      <n-button class="add-btn" @click="handleAddGroup">添加权限组</n-button>
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