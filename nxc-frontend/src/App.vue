<script setup lang="ts">
import {onMounted, toRaw} from 'vue';
import {RouterView} from 'vue-router';
import useThemeStore from "@/stores/useThemeStore";
import useAppInfosStore from "@/stores/useAppInfosStore";
import type {NConfigProvider} from 'naive-ui'

const themeStore = useThemeStore();
const appInfosStore = useAppInfosStore()

onMounted(() => {
  console.log('app挂载');
  // 处理深色主题设置
  themeStore.readEnableDarkMode()
  // 从总设置中获取主题并应用
  themeStore.setThemeFromSetting()

  // 获取总体

});
</script>

<template>
  <n-config-provider :theme="themeStore.getMainTheme" :theme-overrides="toRaw(themeStore.getTheme.selfOverride)">
    <n-message-provider>
      <n-notification-provider>
        <n-dialog-provider>
          <RouterView></RouterView>

        </n-dialog-provider>
      </n-notification-provider>
    </n-message-provider>
  </n-config-provider>

</template>


<style lang="less">
* {
  padding: 0;
  margin: 0;
}

.n-card {
  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
  border: 0;
}
</style>
