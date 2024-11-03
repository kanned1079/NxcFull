<script setup lang="ts">
import {ref, onMounted, onUnmounted, onBeforeUpdate, computed, watch} from "vue"
import * as echarts from 'echarts';
import useAppInfosStore from "@/stores/useAppInfosStore";
import useThemeStore from "@/stores/useThemeStore";
type EChartsOption = echarts.EChartsOption;

const themeStore = useThemeStore();
const appInfoStore = useAppInfosStore();


// 历史访问人数
let visitedChartDOM = ref(null)
// 历史收入
let incomeChartDOM = ref(null)
// 用户活跃情况
let userActivityDOM = ref(null)

let visitedChart: echarts.ECharts | null = null;
let incomeChart: echarts.ECharts | null = null;
let activityChart: echarts.ECharts | null = null;

let visitedChartOption: EChartsOption = {
  title: {
    text: '访问人数趋势图'
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

let userActivityOption: EChartsOption = {
  title: {
    text: '用户占比',
  },
  tooltip: {
    trigger: 'item'
  },
  legend: {
    top: '5%',
    left: 'center'
  },
  series: [
    {
      name: 'Access From',
      type: 'pie',
      radius: ['40%', '70%'],
      avoidLabelOverlap: false,
      padAngle: 5,
      itemStyle: {
        borderRadius: 10
      },
      label: {
        show: false,
        position: 'center',
        fontSize: '1.5rem'
      },
      emphasis: {
        label: {
          show: true,
          fontSize: 40,
          fontWeight: 'bold'
        }
      },
      labelLine: {
        show: false
      },
      data: [
        { value: 1048, name: 'Search Engine' },
        { value: 735, name: 'Direct' },
        { value: 580, name: 'Email' },
        { value: 484, name: 'Union Ads' },
        { value: 300, name: 'Video Ads' }
      ]
    }
  ],
  backgroundColor: 'rgba(255, 255, 255, 0.0)'
};

let getOperatingSystem = (): string => {
  const userAgent = navigator.userAgent;

  if (userAgent.includes("Windows")) {
    return "Windows";
  } else if (userAgent.includes("Macintosh")) {
    return "Macintosh";
  } else if (userAgent.includes("Linux")) {
    return "Linux";
  } else if (userAgent.includes("Android")) {
    return "Android";
  } else if (userAgent.includes("iOS") || userAgent.includes("iPad") || userAgent.includes("iPhone")) {
    return "iOS";
  } else {
    return "Unknown OS";
  }
}


let reloadCharts = () => {
  if (visitedChartDOM.value && incomeChartDOM.value) {
    // 初始化第一个图表
    visitedChart = echarts.init(visitedChartDOM.value, themeStore.enableDarkMode?'dark':null);
    visitedChart.setOption(visitedChartOption);
    // 初始化第二个
    incomeChart = echarts.init(incomeChartDOM.value, themeStore.enableDarkMode?'dark':null);
    incomeChart.setOption(incomeChartOptions)
    // 初始化第三个
    activityChart = echarts.init(userActivityDOM.value, themeStore.enableDarkMode?'dark':null);
    activityChart.setOption(userActivityOption);
  }
}

let generalInfo = ref<{name: string; data: string}[]>([
  {
    name: '浏览器时间',
    data: computed(() => {
      // let timeNow = new Date()
      // return `${timeNow.getUTCFullYear()}-${timeNow.getUTCMonth()}-${timeNow.getUTCDay()} ${timeNow.getUTCHours()}:${timeNow.getMinutes()}:${timeNow.getSeconds()} UTC`
      return (new Date()).toString()
    }).value,
  },
  {
    name: '当前操作系统类型',
    data: getOperatingSystem(),
  },
  {
    name: '系统名',
    data: computed(() => appInfoStore.appCommonConfig.app_name).value,
  },
  {
    name: '最新网址',
    data: computed(() => appInfoStore.appCommonConfig.app_url).value,
  },
  {
    name: '当前系统时间',
    data: '2024-10-18 15:23:03',
  },
  {
    name: '当前系统时间',
    data: '2024-10-18 15:23:03',
  },
])

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

<script lang="ts">
export default {
  name: 'IncomeChart'
}
</script>

<template>
  <div class="root">
    <n-grid cols="1 s:2" responsive="screen" :x-gap="15" :y-gap="15">
      <n-grid-item>
        <n-card hoverable class="card1" :embedded="true" :bordered="false">
          <div class="visited" style="height: 280px" ref="visitedChartDOM"></div>
        </n-card>
        <n-card hoverable class="card2" :embedded="true" :bordered="false">
          <div class="internal" style="height: 280px" ref="incomeChartDOM"></div>
        </n-card>
        <n-card hoverable class="card3" :embedded="true" :bordered="false">
          <div class="internal" style="height: 360px" ref="userActivityDOM"></div>
        </n-card>
      </n-grid-item>
      <n-grid-item>
        <n-card class="r-card" hoverable :embedded="true" :bordered="false" content-style="padding: 0" title="一般">
          <n-table
            :bordered="false"
            :bottom-bordered="false"
            style="background-color: rgba(0,0,0,0.0); padding: 0 15px 10px 15px"
          >
            <n-tr style=" margin-left: 200px" v-for="(item, index) in generalInfo" :key="index">
              <n-td style="background-color: rgba(0,0,0,0.0);">{{ item.name }}</n-td>
              <n-td style="background-color: rgba(0,0,0,0.0);">{{ item.data }}</n-td>
            </n-tr>
          </n-table>
        </n-card>
      </n-grid-item>
    </n-grid>
  </div>

</template>

<style scoped>
.root {
  padding: 20px;
  display: flex;
}
.card1 {
  height: 300px;
  width: 100%;
  margin-bottom: 10px;
  .visited {
    background-color: rgba(0,0,0,0.0);
  }
}
.card2 {
  height: 300px;
  width: 100%;
  margin-bottom: 10px;
  .internal {
    background-color: rgba(0,0,0,0.0);
  }
}

.card3 {
  height: 360px;
  width: 100%;
  .internal {
    background-color: rgba(0,0,0,0.0);
  }
}

.r-card {

}
</style>
