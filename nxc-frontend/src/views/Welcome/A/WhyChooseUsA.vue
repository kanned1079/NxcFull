<script setup lang="ts">
import SnowFall from "@/views/utils/SnowFall.vue";
import {useI18n} from "vue-i18n";
import {ref, onMounted, onBeforeMount} from "vue"
import HeaderA from "@/views/Welcome/A/HeaderA.vue";
import ContentA from "@/views/Welcome/A/ContentA.vue";
import FooterA from "@/views/Welcome/A/FooterA.vue";
import instance from "@/axios/authInstance"

const {t} = useI18n()

let getConfigSuccess = ref<boolean>(false)

interface WelcomeSettings {
  app_sub_description: string
  why_choose_us_hint: string
  bilibili_official_link: string
  youtube_official_link: string
  instagram_link: string
  wechat_official_link: string
  filing_number: string
  page_suffix: string
}

let settings = ref<WelcomeSettings>({
  app_sub_description: '',
  why_choose_us_hint: '',
  bilibili_official_link: '',
  youtube_official_link: '',
  instagram_link: '',
  wechat_official_link: '',
  filing_number: '',
  page_suffix: '',
})

let handleGetConfig = async () => {
  try {
    let {data} = await instance.get('/api/app/v1/welcome', {
      params: {
        lang: 'en',
      }
    })
    if (data.code === 200) {
      getConfigSuccess.value = true
      Object.assign(settings.value, data.config)
    }
  } catch (err: any) {
    console.log(err + '')
  }
}

let showSnowFall = ref<boolean>(false)

onBeforeMount(() => {
  handleGetConfig()
})

onMounted(() => {
  setTimeout(() => {
      showSnowFall.value = true
  } ,1000)
})
</script>

<script lang="ts">
export default {
  name: 'WhyChooseUsA',
}
</script>

<template>
  <SnowFall :show="showSnowFall"></SnowFall>
  <n-space vertical size="large" style="scrollbar-width: none">

    <n-layout style="scrollbar-width: none">
      <n-layout-header style="position: fixed; top: 0; left: 0; right: 0; z-index: 2000; background-color: rgba(0,0,0,0)">
<!--        <HeaderA/>-->
        <HeaderA style="z-index: 16000"></HeaderA>
      </n-layout-header>
      <n-layout-content content-style="padding: 0px;" >
<!--        <ContentA/>-->
        <ContentA
            :app_sub_description="settings.app_sub_description"
            :why_choose_us_hint="settings.why_choose_us_hint"
        ></ContentA>
      </n-layout-content>
      <n-layout-footer>
<!--        <FooterA/>-->
        <FooterA
            v-if="getConfigSuccess"
            :bilibili_official_link="settings.bilibili_official_link"
            :youtube_official_link="settings.youtube_official_link"
            :instagram_link="settings.instagram_link"
            :filing_number="settings.filing_number"
            :page_suffix="settings.page_suffix"
        ></FooterA>
      </n-layout-footer>
    </n-layout>


  </n-space>
</template>

<style scoped lang="less">

</style>