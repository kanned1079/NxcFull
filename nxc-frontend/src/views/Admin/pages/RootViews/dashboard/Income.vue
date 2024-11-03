<script setup lang="ts">
import {ref} from "vue"
import {useI18n} from "vue-i18n";
import useAppInfosStore from "@/stores/useAppInfosStore";
import {Podium as incomeIcon} from '@vicons/ionicons5'

let income = ref<{
  yesterday: number
  month: number
}>({
  yesterday: 138.09,
  month: 2494.17,
})

const {t} = useI18n()
const appInfoStore = useAppInfosStore();
</script>

<script lang="ts">
export default {
  name: 'Income'
}
</script>

<template>
  <n-card class="root" hoverable :embedded="true" :bordered="false" content-style="padding: 0;">
    <n-flex style="justify-content: space-between;">
      <div style="padding: 20px">
              <n-statistic :label="t('adminViews.summary.incomeText')" tabular-nums>
                <n-number-animation ref="numberAnimationInstRef" :precision="2" :from="0" :to="income.yesterday"/>
                /
                <n-number-animation ref="numberAnimationInstRef" :precision="2" :from="0" :to="income.month"/>
                <template #suffix>
                  {{ appInfoStore.appCommonConfig.currency }}
                </template>
              </n-statistic>
      </div>
      <div style="opacity: 0.05; padding: 5px 20px 0 0">
       <n-icon size="50"><incomeIcon/></n-icon>
      </div>
    </n-flex>
  </n-card>
</template>

<style scoped lang="less">
.root {
  width: 100%;
  display: flex;
  flex-direction: row;
  margin: 0 20px;
  .l-content {
    flex: 3;
    background-color: #3a4754;
  }
  .r-content {
    flex: 1;
    background-color: #0e1b42;
  }
}
</style>