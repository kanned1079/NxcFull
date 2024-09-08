<script setup lang="ts" name="DashBoard">
import {onMounted} from 'vue'
import useThemeStore from "@/stores/useThemeStore";
import {RouterView, useRouter} from 'vue-router';
import CommonLogo from "@/components/CommonLogo.vue";
import CommonHeader from "@/components/CommonHeader.vue";
import CommonAside from "@/components/CommonAside.vue";

const themeStore = useThemeStore();

const router = useRouter();

onMounted(() => {
  console.log('原来的dashboard')
  router.push({
    path: '/admin/dashboard/summary',
  })
})
</script>

<template>


  <n-layout has-sider style="width: 100%; position: fixed; left: 0; top: 0; z-index: 1000">
    <div>
      <n-layout-header class="logo" >
        <CommonLogo></CommonLogo>
      </n-layout-header>

    </div>

    <n-layout>
      <n-layout-header>
        <!--        颐和园路-->
        <CommonHeader></CommonHeader>
      </n-layout-header>

      <!--      <n-layout-footer>成府路</n-layout-footer>-->
    </n-layout>
  </n-layout>


  <n-layout has-sider style="width: 100%; position: fixed; left: 0; top: 52px; z-index: 1001">
    <div>

      <n-layout-sider content-style="" width="240px">
        <!--        海淀桥-->
        <CommonAside></CommonAside>
      </n-layout-sider>
    </div>

    <n-layout>

      <n-layout-content content-style="padding: 0" class="content">
        <!--        平山道-->
        <n-scrollbar style="max-height: 100vh">
          <RouterView></RouterView>
          <div style="height: 3rem"></div>
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

.n-layout-sider {
  //background-color: v-bind('nowTheme.commonAsideBgColor');
  height: 100vh;
  align-items: center;
  background-color: v-bind('themeStore.getTheme.globeTheme.asideBgColor');
}

.n-layout-content {
  //background-color: v-bind('nowTheme.commonContentBgColor');
  height: 100%;
}

.logo {
  width: 240px;
  padding: 0;
  overflow: hidden
}

.content {
  background-color: v-bind('themeStore.getTheme.globeTheme.contentBgColor');
}

</style>