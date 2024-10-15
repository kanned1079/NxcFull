<script setup lang="ts">
import {onMounted} from "vue";
import useConfigStore from "@/store/useConfigStore";
import CommonHeader from "@/components/CommonHeader.vue"
import CommonAside from "@/components/CommonAside.vue";
import CommonContent from "@/components/CommonContent.vue";

const configStore = useConfigStore();

onMounted(() => {
  configStore.loadEtcdConfigFromLocalStorage()
})

</script>

<template>
  <n-layout>
    <n-layout-header style="position: fixed; top: 0; left: 0; z-index: 1000">
      <CommonHeader></CommonHeader>
    </n-layout-header>
    <n-layout has-sider style="margin-top: 50px">
      <n-layout-sider content-style="padding: 20px 20px 20px 20px;" style="position: fixed; z-index: 999;">
        <CommonAside></CommonAside>
      </n-layout-sider>
      <n-layout-content content-style="padding: 20px 20px 20px 20px;" style="margin-left: 260px">
        <CommonContent></CommonContent>
      </n-layout-content>
    </n-layout>
  </n-layout>
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