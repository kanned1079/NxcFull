<script setup lang="ts">
import CommonHeader from "@/components/CommonHeader.vue"
import CommonAside from "@/components/CommonAside.vue";
import CommonContent from "@/components/CommonContent.vue";
</script>

<template>
  <n-layout>
    <n-layout-header>
      <CommonHeader></CommonHeader>
    </n-layout-header>
    <n-layout has-sider>
      <n-layout-sider content-style="padding: 0px;">
        <CommonAside></CommonAside>
      </n-layout-sider>
      <n-layout-content content-style="padding: 0px;">
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