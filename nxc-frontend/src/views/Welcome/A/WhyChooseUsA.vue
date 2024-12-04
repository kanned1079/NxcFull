<script setup lang="ts">
import SnowFall from "@/views/utils/SnowFall.vue";
import {useI18n} from "vue-i18n";
import {ref, onMounted} from "vue"
import HeaderA from "@/views/Welcome/A/HeaderA.vue";
import ContentA from "@/views/Welcome/A/ContentA.vue";
import FooterA from "@/views/Welcome/A/FooterA.vue";

const {t} = useI18n()

let showSnowFall = ref<boolean>(false)

onMounted(() => {
  setTimeout(() => {
      showSnowFall.value = true
  } ,500)
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
        <ContentA/>
      </n-layout-content>
      <n-layout-footer>
        <FooterA/>
      </n-layout-footer>
    </n-layout>


  </n-space>
</template>

<style scoped lang="less">

</style>