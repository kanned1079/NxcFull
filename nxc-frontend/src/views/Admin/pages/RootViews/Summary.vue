<script setup lang="ts">
import {onMounted, ref} from "vue";
import Settings from '@/views/Admin/pages/RootViews/dashboard/Settings.vue'
import useThemeStore from "@/stores/useThemeStore";
import Income from '@/views/Admin/pages/RootViews/dashboard/Income.vue'
import IncomeChart from "@/views/Admin/pages/RootViews/dashboard/IncomeChart.vue";

const themeStore = useThemeStore();

let animated = ref<boolean>(false)

onMounted(() => {
  themeStore.contentPath = '/admin/dashboard/summary'
  themeStore.menuSelected = 'dashboard'

  animated.value = true
})

</script>

<script lang="ts">
export default {
  name: 'Summary'
}
</script>

<template>
  <transition name="slide-fade">
    <div class="root" v-if="animated">
      <div class="setting-panel">
        <Settings></Settings>
      </div>
      <div class="income-panel">
        <Income></Income>
      </div>
      <div class="income-chart">
        <IncomeChart></IncomeChart>
      </div>
    </div>
  </transition>

</template>

<style lang="less" scoped>
.root {
  width: 100%;
  border-radius: 3px;
  min-width: 800px;
  height: 140px;
  text-align: center;

  .setting-panel {
    display: flex;
    align-items: center;
    line-height: 160px;
    justify-content: center;
  }

  .income-panel {
    display: flex;
    align-items: center;
    padding: 0;
    line-height: 160px;
    justify-content: center;
  }

  .income-chart {
    display: flex;
    align-items: center;
    padding: 0;
    justify-content: center;
  }
}
//.n-card {
//  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
//  border: 0;
//}
</style>