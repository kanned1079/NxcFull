<script setup lang="ts" name="IncomeChart">
import {ref, onMounted, onUnmounted, onBeforeUpdate} from "vue"
import * as echarts from 'echarts';
import useThemeStore from "@/stores/useThemeStore";
type EChartsOption = echarts.EChartsOption;

const themeStore = useThemeStore();


// 历史访问人数
let visitedChartDOM = ref(null)
// 历史收入
let incomeChartDOM = ref(null)

let visitedChart: echarts.ECharts | null = null;
let incomeChart: echarts.ECharts | null = null;

let visitedChartOption: EChartsOption = {
  title: {
    text: '访问人数'
  },
  xAxis: {
    type: 'category',
    data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
  },
  yAxis: {
    type: 'value'
  },
  series: [
    {
      data: [120, 932, 345, 934, 1290, 140, 543],
      type: 'line',
      smooth: true
    }
  ],
  backgroundColor: 'rgba(255, 255, 255, 0.0)'
};

let incomeChartOptions: EChartsOption = {
  title: {
    text: '历史收入'
  },
  tooltip: {
    trigger: 'axis',
    axisPointer: {
      type: 'shadow'
    }
  },
  grid: {
    left: '3%',
    right: '4%',
    // bottom: '3%',
    containLabel: true
  },
  xAxis: [
    {
      type: 'category',
      data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
      axisTick: {
        alignWithLabel: true
      }
    }
  ],
  yAxis: [
    {
      type: 'value'
    }
  ],
  series: [
    {
      name: 'Direct',
      type: 'bar',
      barWidth: '50%',
      data: [10, 52, 200, 334, 390, 330, 220]
    }
  ],
  backgroundColor: 'rgba(255, 255, 255, 0.0)'
}

let reloadCharts = () => {
  if (visitedChartDOM.value && incomeChartDOM.value) {
    // 初始化第一个图表
    visitedChart = echarts.init(visitedChartDOM.value, themeStore.enableDarkMode?'dark':null);
    visitedChart.setOption(visitedChartOption);
    // 初始化第二个
    incomeChart = echarts.init(incomeChartDOM.value, themeStore.enableDarkMode?'dark':null);
    incomeChart.setOption(incomeChartOptions)
  }
}

onMounted(async () => {
  reloadCharts()

});

onBeforeUpdate(() => {
  reloadCharts()
})

onUnmounted(() => {
  if (visitedChart && incomeChart) {
    visitedChart.dispose();
    incomeChart.dispose();
  }
});
</script>

<template>
  <div class="root">
    <n-card hoverable class="card1" :embedded="true">
      <div class="visited" style="height: 280px" ref="visitedChartDOM"></div>
    </n-card>
    <n-card hoverable class="card2" :embedded="true">
      <div class="income" style="height: 280px" ref="incomeChartDOM"></div>
    </n-card>
  </div>

</template>

<style scoped>
.root {
  height: 300px;
  margin: 0 20px 20px 20px;
  display: flex;
  .card1 {
    height: 300px;
    margin-right: 10px;
    .visited {
      background-color: rgba(0,0,0,0.0);
    }
  }
  .card2 {
    height: 300px;
    margin-left: 10px;
    .income {
      background-color: rgba(0,0,0,0.0);
    }
  }


}

.n-card {
  //background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
  border: 0;
}
</style>
