<script setup lang="ts">
import {useI18n} from "vue-i18n"
import Site from "@/views/Admin/pages/SettingViews/SystemConfig/Site.vue";
import Security from "@/views/Admin/pages/SettingViews/SystemConfig/Security.vue";
import Personalization from "@/views/Admin/pages/SettingViews/SystemConfig/Personalization.vue";
import BackendConfig from "@/views/Admin/pages/SettingViews/SystemConfig/Welcome.vue";
import SendMail from "@/views/Admin/pages/SettingViews/SystemConfig/SendMail.vue";
import WeApp from "@/views/Admin/pages/SettingViews/SystemConfig/WeApp.vue"
import Notice from "@/views/Admin/pages/SettingViews/SystemConfig/Notice.vue"
import {onBeforeMount, onMounted, ref} from "vue"
import useThemeStore from "@/stores/useThemeStore";
import Subscribe from "@/views/Admin/pages/SettingViews/SystemConfig/Invite.vue";
import useSettingStore from "@/stores/useSettingStore";

const {t} = useI18n()
let animated = ref<boolean>(false)
const settingStore = useSettingStore()
const themeStore = useThemeStore();

onBeforeMount(async () => {
  themeStore.menuSelected = 'system-config'
  await settingStore.loadSetting()
})

onMounted(() => {
  console.log('menuSelected:', themeStore.menuSelected)
  themeStore.contentPath = '/admin/dashboard/systemconfig'
  animated.value = true
})

</script>

<script lang="ts">
export default {
  name: 'SystemConfig'
}
</script>

<template>
  <transition name="slide-fade">
    <div class="root" v-if="animated" style="margin: 20px">
      <n-card hoverable style="padding: 0" :bordered="false" :embedded="true">
        <n-tabs type="segment" animated>
          <n-tab-pane name="chap1" :tab="t('adminViews.systemConfig.site.common.title')">
            <Site></Site>
          </n-tab-pane>
          <n-tab-pane name="chap2" :tab="t('adminViews.systemConfig.security.common.title')">
            <Security></Security>
          </n-tab-pane>
          <n-tab-pane name="chap3" :tab="t('adminViews.systemConfig.frontend.common.title')">
            <Personalization></Personalization>
          </n-tab-pane>
          <n-tab-pane name="chap4" :tab="t('adminViews.systemConfig.inviteAndRebate.common.title')">
            <Subscribe></Subscribe>
          </n-tab-pane>
          <n-tab-pane name="chap5" :tab="t('adminViews.systemConfig.welcome.common.title')">
            <BackendConfig></BackendConfig>
          </n-tab-pane>
          <n-tab-pane name="chap6" :tab="t('adminViews.systemConfig.sendMail.common.title')">
              <SendMail></SendMail>
          </n-tab-pane>
<!--          <n-tab-pane name="chap7" :tab="t('adminViews.systemConfig.notice.common.title')">-->
<!--            <Notice></Notice>-->
<!--          </n-tab-pane>-->
          <n-tab-pane name="chap7" :tab="t('adminViews.systemConfig.appDownload.common.title')">
            <WeApp></WeApp>
          </n-tab-pane>
        </n-tabs>
      </n-card>
    </div>
  </transition>

</template>

<style lang="less" scoped>
</style>