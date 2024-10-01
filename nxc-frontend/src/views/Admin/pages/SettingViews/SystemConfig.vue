<script setup lang="ts">
import Site from "@/views/Admin/pages/SettingViews/SystemConfig/Site.vue";
import Security from "@/views/Admin/pages/SettingViews/SystemConfig/Security.vue";
import Personalization from "@/views/Admin/pages/SettingViews/SystemConfig/Personalization.vue";
import NodesConfig from "@/views/Admin/pages/SettingViews/SystemConfig/BackendConfig.vue";
import SendMail from "@/views/Admin/pages/SettingViews/SystemConfig/SendMail.vue";
import WeApp from "@/views/Admin/pages/SettingViews/SystemConfig/WeApp.vue"
import Notice from "@/views/Admin/pages/SettingViews/SystemConfig/Notice.vue"
import BackendConfig from "@/views/Admin/pages/SettingViews/SystemConfig/BackendConfig.vue"

import useThemeStore from "@/stores/useThemeStore";
import Subscribe from "@/views/Admin/pages/SettingViews/SystemConfig/Subscribe.vue";
import useSettingStore from "@/stores/useSettingStore";
import {onBeforeMount, onMounted} from "vue";
import instance from "@/axios";

const settingStore = useSettingStore()

const themeStore = useThemeStore();
//
// let initSystemConfigData = async () => {
//   let { data } = await instance.get('http://localhost:8080/admin/get-setting')
//   console.log(data)
// }

onBeforeMount(() => {
  settingStore.loadSetting()

})

onMounted(() => {
  console.log('menuSelected:', themeStore.menuSelected)
  themeStore.contentPath = '/admin/dashboard/systemconfig'
  themeStore.menuSelected = 'system-config'
})

</script>

<script lang="ts">
export default {
  name: 'SystemConfig'
}
</script>

<template>

  <div class="root" style="margin: 20px">
    <n-card hoverable style="padding: 0" :bordered="false" :embedded="true">
      <n-tabs type="segment" animated>
        <n-tab-pane name="chap1" tab="站点">
          <Site></Site>
        </n-tab-pane>
        <n-tab-pane name="chap2" tab="安全">
          <Security></Security>
        </n-tab-pane>
        <n-tab-pane name="chap3" tab="个性化">
          <Personalization></Personalization>
        </n-tab-pane>
        <n-tab-pane name="chap4" tab="订阅">
          <Subscribe></Subscribe>
        </n-tab-pane>
        <n-tab-pane name="chap5" tab="后端设置">
          <BackendConfig></BackendConfig>
        </n-tab-pane>
        <n-tab-pane name="chap6" tab="邮件">
          <n-message-provider>
            <SendMail></SendMail>
          </n-message-provider>
        </n-tab-pane>
        <n-tab-pane name="chap7" tab="通知">
          <Notice></Notice>
        </n-tab-pane>
        <n-tab-pane name="chap8" tab="APP">
          <WeApp></WeApp>
        </n-tab-pane>
      </n-tabs>
    </n-card>
  </div>


  <!--  <div style="display: flex; flex-direction: column; max-height: 100vh; overflow-y: auto;">-->
  <!--    <Personalization></Personalization>-->
  <!--    <Site></Site>-->
  <!--    <h1>1</h1>-->
  <!--  </div>-->


</template>

<style lang="less" scoped>
//.n-card {
//  //background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
//  border: 0;
//}
</style>