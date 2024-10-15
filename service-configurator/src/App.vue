<script setup lang="ts">
import {onMounted} from "vue";
import useConfigStore from "@/store/useConfigStore";
import useThemeStore from "@/store/useThemeStore";
import CommonHeader from "@/components/CommonHeader.vue"
import CommonAside from "@/components/CommonAside.vue";
import CommonContent from "@/components/CommonContent.vue";
import { darkTheme } from 'naive-ui'

const configStore = useConfigStore();
const themeStore = useThemeStore();

onMounted(() => {
  configStore.loadEtcdConfigFromLocalStorage()
})

</script>

<template>

<!--  <n-config-provider :theme="darkTheme">-->
<!--    <n-layout>-->
<!--      <n-layout-header style="position: fixed; top: 0; left: 0; z-index: 1000">-->
<!--        <CommonHeader></CommonHeader>-->
<!--      </n-layout-header>-->
<!--      <n-layout has-sider style="margin-top: 50px">-->
<!--        <n-layout-sider style="position: fixed; z-index: 999;" content-style="padding: 0; ">-->
<!--          <CommonAside></CommonAside>-->
<!--        </n-layout-sider>-->
<!--        <n-layout-content content-style="padding: 20px 20px 20px 20px;" style="margin-left: 260px">-->
<!--          <CommonContent></CommonContent>-->
<!--        </n-layout-content>-->
<!--      </n-layout>-->
<!--    </n-layout>-->
<!--  </n-config-provider>-->

  <n-config-provider :theme="themeStore.darkThemeEnabled?darkTheme:null">
    <n-layout>
      <n-layout-header style="position: fixed; top: 0; left: 0; z-index: 1000">
        <CommonHeader></CommonHeader>
      </n-layout-header>
      <n-layout-content >
        <div style="display: flex; flex-direction: row; margin-top: 50px">
          <div style="flex: 1;">
            <CommonAside></CommonAside>
          </div>
          <div style="flex: 5; margin: 20px 20px 0 0">
            <CommonContent></CommonContent>

          </div>

        </div>

<!--        <n-layout-sider style="position: fixed; z-index: 999;" content-style="padding: 0; ">-->
<!--          <CommonAside></CommonAside>-->
<!--        </n-layout-sider>-->
<!--        <n-layout-content content-style="padding: 20px 20px 20px 20px;" style="margin-left: 260px">-->
<!--          <CommonContent></CommonContent>-->
<!--        </n-layout-content>-->
      </n-layout-content>
    </n-layout>
  </n-config-provider>


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