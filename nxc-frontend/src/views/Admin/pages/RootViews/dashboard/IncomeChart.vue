<script setup lang="ts">
import {useI18n} from "vue-i18n";
// TODO 待实现全球化
import Income from "@/views/Admin/pages/RootViews/dashboard/Income.vue"
import {ref, onMounted, onUnmounted, onBeforeMount, onBeforeUpdate, computed, watch, toRaw} from "vue"
import * as echarts from 'echarts';
import useAppInfosStore from "@/stores/useAppInfosStore";
import useThemeStore from "@/stores/useThemeStore";
import instance from "@/axios";
import {Podium as incomeIcon} from '@vicons/ionicons5'

type EChartsOption = echarts.EChartsOption;

const {t} = useI18n()
const themeStore = useThemeStore();
const appInfoStore = useAppInfosStore();

let apiAccessCount = ref<number[]>([])
let incomeCount = ref<number[]>([])
let yesterdayIncome = ref<number>(0.00)
let monthIncome = ref<number>(0.00)
let activeUserCount = ref<number>(0)
let inactiveUserCount = ref<number>(0)
let allUserCount = ref<number>(0)

let getAppOverview = async () => {
  try {
    let {data} = await instance.get('/api/admin/v1/app/overview')
    if (data.code === 200) {
      // console.log(data)
      data.api_access_count_history.forEach((item: number) => apiAccessCount.value.unshift(item))
      data.income_count_history.forEach((item: number) => incomeCount.value.unshift(item))
      yesterdayIncome.value = data.income_yesterday?data.income_yesterday:0.00
      monthIncome.value = data.income_this_month?data.income_this_month:0.00
      // console.log('a:', yesterdayIncome.value, monthIncome.value)
    }
  } catch (error) {
    console.log(error)
  }
}

let getLast7DaysWeekdays = (): string[] => {
  const weekdays = ["周日", "周一", "周二", "周三", "周四", "周五", "周六"];

  // 当前日期
  const now = new Date();
  const todayIndex = now.getDay() - 1; // 获取今天是星期几 (0=Sun, ..., 6=Sat)

  // 最近 7 天的星期数组
  const last7Days = [];
  for (let i = 0; i < 7; i++) {
    const index = (todayIndex - i + 7) % 7; // 循环计算星期
    last7Days.push(weekdays[index]);
  }

  return last7Days;
}


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
  color: {
    type: 'linear',
    x: 0,
    y: 0,
    x2: 0,
    y2: 1,
    colorStops: [
      {
        offset: 0,
        color: computed(() => themeStore.getTheme.selfOverride.common?.primaryColor).value as string,
      },
      {
        offset: 1,
        color: computed(() => themeStore.getTheme.selfOverride.common?.primaryColor).value as string,
      },
    ],
    global: false,
  },
  tooltip: {
    trigger: 'axis',
    formatter: '{b}: {c}',
    backgroundColor: 'rgba(50, 50, 50, 0.7)',
    borderColor: '#ccc',
    borderWidth: 1,
    textStyle: {
      color: '#fff',
    },
  },
  grid: {
    top: '10px', // 调整顶部间距
    left: '10px', // 调整左侧间距
    right: '10px', // 调整右侧间距
    bottom: '10px', // 调整底部间距
    containLabel: true, // 确保标签显示完整
  },
  xAxis: {
    type: 'category',
    data: getLast7DaysWeekdays(),
    axisLine: {
      lineStyle: {
        color: '#ccc',
      },
    },
    axisLabel: {
      fontSize: 12,
    },
  },
  yAxis: {
    type: 'value',
    axisLine: {
      lineStyle: {
        color: '#ccc',
      },
    },
    axisLabel: {
      fontSize: 12,
    },
    splitLine: {
      lineStyle: {
        type: 'dashed',
      },
    },
  },
  series: [
    {
      name: '访问次数',
      data: apiAccessCount.value,
      type: 'line',
      smooth: true,
      areaStyle: {
        color: {
          type: 'linear',
          x: 0,
          y: 0,
          x2: 0,
          y2: 1,
          colorStops: [
            {
              offset: 0,
              color: computed(() => themeStore.getTheme.selfOverride.common?.primaryColor).value as string,
            },
            {
              offset: 1,
              color: 'rgba(255, 255, 255, 0)',
            },
          ],
        },
      },
      lineStyle: {
        width: 2,
      },
      symbol: 'circle',
      symbolSize: 6,
    },
  ],
  backgroundColor: 'rgba(255, 255, 255, 0.0)',
};

let incomeChartOptions: EChartsOption = {
  color: {
    type: 'linear',
    x: 0,
    y: 0,
    x2: 0,
    y2: 1,
    colorStops: [
      {
        offset: 0,
        color: computed(() => themeStore.getTheme.selfOverride.common?.primaryColor).value as string,
      },
      {
        offset: 1,
        color: computed(() => themeStore.getTheme.selfOverride.common?.primaryColor).value as string,
      },
    ],
    global: false,
  },
  tooltip: {
    trigger: 'axis',
    axisPointer: {
      type: 'shadow',
    },
    formatter: '{b}: {c} ' + appInfoStore.appCommonConfig.currency,
    backgroundColor: 'rgba(50, 50, 50, 0.7)',
    borderColor: '#ccc',
    borderWidth: 1,
    textStyle: {
      color: '#fff',
    },
  },
  grid: {
    top: '10px', // 调整顶部间距
    left: '10px', // 调整左侧间距
    right: '10px', // 调整右侧间距
    bottom: '10px', // 调整底部间距
    containLabel: true, // 确保标签显示完整
  },
  xAxis: {
    type: 'category',
    data: getLast7DaysWeekdays(),
    axisLine: {
      lineStyle: {
        color: '#ccc',
      },
    },
    axisLabel: {
      fontSize: 12,
    },
    axisTick: {
      alignWithLabel: true,
    },
  },
  yAxis: {
    type: 'value',
    axisLine: {
      lineStyle: {
        color: '#ccc',
      },
    },
    axisLabel: {
      fontSize: 12,
    },
    splitLine: {
      lineStyle: {
        type: 'dashed',
      },
    },
  },
  series: [
    {
      name: '收入',
      type: 'bar',
      barWidth: '50%',
      data: incomeCount.value,
      itemStyle: {
        borderRadius: [4, 4, 0, 0], // 柱状图圆角
      },
    },
  ],
  backgroundColor: 'rgba(255, 255, 255, 0.0)',
};

let userActivityOption: EChartsOption = {
  tooltip: {
    trigger: 'item',
    formatter: '{b}: {c} ({d}%)',
    backgroundColor: 'rgba(50, 50, 50, 0.7)',
    borderColor: '#ccc',
    borderWidth: 1,
    textStyle: {
      color: '#fff',
    },
  },
  legend: {
    orient: 'vertical',
    left: 'left',
    itemGap: 15,  // 增加图例项间距
    textStyle: {
      fontSize: 12,
      color: '#666', // 设置图例字体颜色
    },
  },
  series: [
    {
      type: 'pie',
      radius: '50%',
      data: [
        { value: 31.3, name: '活跃' },
        { value: 61.1, name: '非活跃' },
        { value: 1.6, name: '其他' },
        { value: 6.0, name: '注销或封禁' },
      ],
      itemStyle: {
        borderRadius: 5, // 为每个扇区添加圆角
      },
      label: {
        show: true,
        position: 'outside', // 标签放在扇区外侧
        formatter: '{b}: {d}%', // 显示名称和百分比
        fontSize: 12,
        color: '#666',
      },
      emphasis: {
        itemStyle: {
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: 'rgba(0, 0, 0, 0.5)',
          borderWidth: 2,
          borderColor: '#fff', // 强调时的边框颜色
        },
      },
    },
  ],
  backgroundColor: 'rgba(255, 255, 255, 0.0)', // 保持透明背景
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
    visitedChart = echarts.init(visitedChartDOM.value, themeStore.enableDarkMode ? 'dark' : null);
    visitedChart.setOption(visitedChartOption);
    // 初始化第二个
    incomeChart = echarts.init(incomeChartDOM.value, themeStore.enableDarkMode ? 'dark' : null);
    incomeChart.setOption(incomeChartOptions)
    // 初始化第三个
    activityChart = echarts.init(userActivityDOM.value, themeStore.enableDarkMode ? 'dark' : null);
    activityChart.setOption(userActivityOption);
  }
}

let usersInfo = ref<{ name: string; data: number }[]>([
  {
    name: '总注册用户数',
    data: 2642,
  },
  {
    name: '活跃用户数',
    data: 82,
  },
  {
    name: '非活跃用户数',
    data: 24,
  },
  {
    name: '封禁或注销',
    data: 13,
  }
])

let generalInfo = ref<{ name: string; data: string }[]>([
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
    name: '货币单位',
    data: computed(() => appInfoStore.appCommonConfig.currency).value,
  },
  {
    name: '允许用户注册',
    data: computed(() => !appInfoStore.appCommonConfig.stop_register?'是':'否').value,
  },
])

let serverInfo = ref<{ name: string; data: string }[]>([
  {
    name: '服务器时间',
    data: '2024-10-18 15:23:03',
  },
  {
    name: 'API网关运行状态',
    data: '运行正常 (https)',
  },
  {
    name: '数据库运行状态',
    data: '运行正常',
  },
  {
    name: 'Redis运行状态',
    data: '运行正常',
  },
  {
    name: 'Etcd运行状态',
    data: '运行正常',
  },
  {
    name: '服务器操作系统类型',
    data: 'Linux 5.14.287',
  },
  {
    name: '后端程序操作系统架构',
    data: 'x86_64',
  },
  {
    name: '启用的支付方式',
    data: 'Wechat | Alipay | ApplePay',
  },

])

onBeforeMount(async () => {
  await getAppOverview()  // 先获取数据
  reloadCharts()  // 再渲染图表
})

onMounted(async () => {
  // await reloadCharts()
  // console.log(props.apiAccessCount)

  // for (let i = 0; i < props.apiAccessCount.length; i++) {
  //   console.log(props.apiAccessCount[i])
  // }

  // console.log(toRaw(props.apiAccessCount))

  console.log(yesterdayIncome.value, monthIncome.value)

});

onBeforeUpdate(() => {
  // reloadCharts()
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
      <n-card
          class="income-part-root"
          hoverable
          :embedded="true"
          :bordered="false"
          content-style="padding: 0;"
      >
        <n-flex
            style="justify-content: space-between;"
        >
          <div
              style="padding: 20px"
          >
            <n-statistic
                :label="t('adminViews.summary.incomeText')"
                tabular-nums
            >
              <n-number-animation
                  ref="numberAnimationInstRef"
                  :precision="2"
                  :from="0"
                  :to="yesterdayIncome"
              />
              /
              <n-number-animation
                  ref="numberAnimationInstRef"
                  :precision="2"
                  :from="0"
                  :to="monthIncome"
              />
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

<!--      <Income :yesterday-income="yesterdayIncome" :this-month-income="monthIncome"></Income>-->
      <div>
        <n-grid cols="1 s:2" responsive="screen" :x-gap="15" :y-gap="15">
          <n-grid-item>
            <n-card class="user-count-card" hoverable :embedded="true" :bordered="false" content-style="padding: 0" title="用户概览">
              <n-table
                  :bordered="false"
                  :bottom-bordered="false"
                  style="background-color: rgba(0,0,0,0.0); padding: 0 15px 0 15px"
              >
                <n-tr style=" margin-left: 200px" v-for="(item, index) in usersInfo" :key="index">
                  <n-td style="background-color: rgba(0,0,0,0.0);">{{ item.name }}</n-td>
                  <n-td style="background-color: rgba(0,0,0,0.0);">{{ item.data }}</n-td>
                </n-tr>
              </n-table>
            </n-card>
            <n-card
                title="最近一周API接口访问次数"
                hoverable class="card1" :embedded="true" :bordered="false">
              <div class="visited" style="height: 280px" ref="visitedChartDOM"></div>
            </n-card>
            <n-card
                title="最近一周收入金额"
                hoverable class="card2" :embedded="true" :bordered="false">
              <div class="internal" style="height: 280px" ref="incomeChartDOM"></div>
            </n-card>
<!--            <n-card-->
<!--                title="活跃用户占比"-->
<!--                hoverable class="card3" :embedded="true" :bordered="false">-->
<!--              <div class="internal" style="height: 360px" ref="userActivityDOM"></div>-->
<!--            </n-card>-->
          </n-grid-item>
          <n-grid-item>
            <n-card
                class="r-card"
                hoverable
                :embedded="true"
                :bordered="false"
                content-style="padding: 0"
                title="一般"
            >
              <n-table
                  :bordered="false"
                  :bottom-bordered="false"
                  style="background-color: rgba(0,0,0,0.0); padding: 0 15px 0 15px"
              >
                <n-tr
                    style=" margin-left: 200px"
                    v-for="(item, index) in generalInfo"
                    :key="index"
                >
                  <n-td
                      style="background-color: rgba(0,0,0,0.0);"
                  >
                    {{ item.name }}
                  </n-td>
                  <n-td
                      style="background-color: rgba(0,0,0,0.0);"
                  >
                    {{ item.data }}
                  </n-td>
                </n-tr>
              </n-table>
            </n-card>

            <n-card
                style="margin-top: 10px"
                class="r-card"
                hoverable
                :embedded="true"
                :bordered="false"
                content-style="padding: 0"
                title="系统配置"
            >
              <n-table
                  :bordered="false"
                  :bottom-bordered="false"
                  style="background-color: rgba(0,0,0,0.0); padding: 0 15px 10px 15px"
              >
                <n-tr
                    style=" margin-left: 200px"
                    v-for="(item, index) in serverInfo"
                    :key="index"
                >
                  <n-td
                      style="background-color: rgba(0,0,0,0.0);"
                  >
                    {{ item.name }}
                  </n-td>
                  <n-td
                      style="background-color: rgba(0,0,0,0.0);"
                  >
                    {{ item.data }}
                  </n-td>
                </n-tr>
              </n-table>

            </n-card>
          </n-grid-item>
        </n-grid>
      </div>
    </div>
</template>

<style scoped>
.root {
  padding: 15px 20px 20px 20px;
  display: flex;
  flex-direction: column;
  .income-part-root {
    display: flex;
    flex-direction: row;
    margin-bottom: 15px;
  }
}

.card1 {
  width: 100%;
  margin-bottom: 10px;

  .visited {
    background-color: rgba(0, 0, 0, 0.0);
  }
}

.card2 {
  width: 100%;
  margin-bottom: 10px;

  .internal {
    background-color: rgba(0, 0, 0, 0.0);
  }
}

.card3 {
  width: 100%;

  .internal {
    background-color: rgba(0, 0, 0, 0.0);
  }
}

.r-card {

}

.user-count-card {
  margin-bottom: 15px;
  padding-bottom: 10px;
}


</style>
