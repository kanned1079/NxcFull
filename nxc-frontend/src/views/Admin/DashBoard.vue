<script setup lang="ts">
import useThemeStore from "@/stores/useThemeStore";
import {RouterView} from 'vue-router';
import CommonHeader from "@/components/CommonHeader.vue";
import CommonAside from "@/components/CommonAside.vue";
const themeStore = useThemeStore();
</script>

<script lang="ts">
export default {
  name: 'DashBoard'
}
</script>

<template>

  <n-layout has-sider style="width: 100%; position: fixed; left: 0; top: 0; z-index: 1000">
    <n-layout-sider
        v-show="!themeStore.menuCollapsed"
        content-style="padding: 0px;"
        :collapsed-width="64"
        style="width: auto"
        :collapsed="themeStore.menuCollapsed"
    >
      <CommonAside></CommonAside>
    </n-layout-sider>
    <n-layout>
      <n-layout-header style="position: fixed; z-index: 1001;">
        <CommonHeader />
      </n-layout-header>
      <n-layout-content content-style="padding: 0;" class="content">
        <n-scrollbar style="max-height: 100vh;">
          <RouterView />
          <div style="height: 10rem"></div>
        </n-scrollbar>
      </n-layout-content>
      <!--      <n-layout-footer>成府路</n-layout-footer>-->
    </n-layout>
  </n-layout>


</template>

<style lang="less" scoped>
.n-layout-header,
.n-layout-footer {
  padding: 0;
  height: 3.25rem;
  width: 100%;
  overflow: hidden
}

@media screen and (min-width: 768px) {
  .n-layout-header {
    width: calc(100% - 272px);
  }
}

.n-layout-sider {
  height: 100vh;
  align-items: center;
  background-color: v-bind('themeStore.getTheme.globeTheme.asideBgColor');
}

.n-layout-content {
  height: 100vh;

}

.logo {
  padding: 0;
  overflow: hidden
}

.content {
  margin-top: 52px;
}

</style>