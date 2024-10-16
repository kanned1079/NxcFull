<script setup lang="ts">
import {onMounted} from "vue";
import useConfigStore from "@/store/useConfigStore";
import useThemeStore from "@/store/useThemeStore";
import CommonHeader from "@/components/CommonHeader.vue"
import CommonAside from "@/components/CommonAside.vue";
import CommonContent from "@/components/CommonContent.vue";
import {darkTheme, GlobalThemeOverrides, NConfigProvider} from 'naive-ui'
// import {instance} from "@/axios";

const themeOverrides: GlobalThemeOverrides = {
  common: {
    primaryColor: '#69b0ac',
    primaryColorHover: '#78c2c4',
    primaryColorPressed: '#305a56',
  },
  Button: {},
  Switch: {
    railColorActive: '#69b0ac'
  },
  Input: {
    caretColor: '#69b0ac'
  }
}

const configStore = useConfigStore();
const themeStore = useThemeStore();


onMounted(async () => {
  configStore.loadEtcdConfigFromLocalStorage()
  // checkEtcd()
  await configStore.checkEtcd()
})

</script>

<template>
  <n-message-provider>

    <n-loading-bar-provider>
      <n-config-provider :theme-overrides="themeOverrides" :theme="themeStore.darkThemeEnabled?darkTheme:null">
        <n-layout>
          <n-layout-header style="position: fixed; top: 0; left: 0; z-index: 1000">
            <CommonHeader></CommonHeader>
          </n-layout-header>
          <n-layout-content>
            <div style="display: flex; flex-direction: row; margin-top: 50px">
              <div style="flex: 1;">
                <CommonAside></CommonAside>
              </div>
              <div style="flex: 5; margin: 20px 20px 0 0">
                <CommonContent></CommonContent>
              </div>
            </div>
          </n-layout-content>
        </n-layout>
      </n-config-provider>
    </n-loading-bar-provider>
  </n-message-provider>

</template>

<style lang="less" scoped>
* {
  padding: 0;
  margin: 0;
  scrollbar-width: none;
  -ms-overflow-style: none;
}

*::-webkit-scrollbar {
  display: none;
}

//.n-layout-header,
//.n-layout-footer {
//  background: rgba(128, 128, 128, 0.2);
//  padding: 24px;
//}

//.n-layout-sider {
//  background: rgba(128, 128, 128, 0.3);
//}

//.n-layout-content {
//  background: rgba(128, 128, 128, 0.4);
//}
</style>