<script setup lang="ts">
import {useI18n} from "vue-i18n";
// TODO 待实现全球化
import {computed, onBeforeMount, onBeforeUpdate, onMounted, onUnmounted, ref, watch} from "vue"
import {useRouter} from "vue-router"
import * as echarts from 'echarts';
import useAppInfosStore from "@/stores/useAppInfosStore";
import useThemeStore from "@/stores/useThemeStore";
// import instance from "@/axios";
import {useMessage} from "naive-ui"
import {
  Podium as incomeIcon,
  CheckmarkCircle as successIcon,
  CloseCircle as errIcon,
  FunnelOutline as userPrefixIcon,
} from '@vicons/ionicons5'
import {handleGetUserLayout, handleFetchAppCommonConfig, handleGetAppOverview} from "@/api/admin/server";
import {config as backendConfig} from "@/config"
// import {type IsNetRequestedSuccess} from "@/hooks/useNetworkResultHook"
import PageHead from "@/views/utils/PageHead.vue";
// import ErrorPage from "@/views/ErrorPage.vue";

type EChartsOption = echarts.EChartsOption;

// const props = defineProps<{
//   setNetStatus: (status: IsNetRequestedSuccess) => void
// }>()


const {t, locale} = useI18n()
const themeStore = useThemeStore();
const appInfoStore = useAppInfosStore();
const message = useMessage()
const router = useRouter()

let animated = ref<boolean>(false)
let apiAccessCount = ref<number[]>([])
let incomeCount = ref<number[]>([])
let yesterdayIncome = ref<number>(0.00)
let monthIncome = ref<number>(0.00)
let netFetchResult = ref<number[]>([])

let getLast7DaysWeekdays = (): string[] => {
  const weekdays = [
    t('week.sunday'),
    t('week.monday'),
    t('week.tuesday'),
    t('week.wednesday'),
    t('week.thursday'),
    t('week.friday'),
    t('week.saturday'),
  ];

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

watch(() => locale.value, () => {
  clearCharts()
  reloadCharts()  // 再渲染图表
})

// 历史访问人数
let visitedChartDOM = ref(null)
// 历史收入
let incomeChartDOM = ref(null)

let visitedChart: echarts.ECharts | null = null;
let incomeChart: echarts.ECharts | null = null;


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

let clearCharts = () => {
  if (visitedChart) {
    echarts.dispose(visitedChart);  // 销毁第一个图表实例
    visitedChart = null;  // 清除引用
  }

  if (incomeChart) {
    echarts.dispose(incomeChart);  // 销毁第二个图表实例
    incomeChart = null;  // 清除引用
  }
}

let reloadCharts = () => {
  if (visitedChartDOM.value && incomeChartDOM.value) {
    // 初始化第一个图表
    visitedChart = echarts.init(visitedChartDOM.value, themeStore.enableDarkMode ? 'dark' : null);
    visitedChart.setOption({
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
    } as EChartsOption);
    // 初始化第二个
    incomeChart = echarts.init(incomeChartDOM.value, themeStore.enableDarkMode ? 'dark' : null);
    incomeChart.setOption({
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
          name: 'income',
          type: 'bar',
          barWidth: '50%',
          data: incomeCount.value,
          itemStyle: {
            borderRadius: [4, 4, 0, 0], // 柱状图圆角
          },
        },
      ],
      backgroundColor: 'rgba(255, 255, 255, 0.0)',
    } as EChartsOption)
  }
}

type MiscellaneousConfig = {
  userAll: number;
  userActive: number;
  userInactive: number;
  userBlocked: number;
  config: {
    serverTime: string;
    apiServerStatus: boolean;
    mysqlServerStatus: boolean;
    redisServerStatus: boolean;
    osType: string;
    osArch: string;
    uptime: string;
    ginMode: string;
    numCpu: number;
    numCgoCall: number;
    enabledPayMethods?: string[]
  }
}

let miscellaneousConfig = ref<MiscellaneousConfig>({
  userAll: 100,
  userActive: 20,
  userInactive: 10,
  userBlocked: 30,
  config: {
    serverTime: '',
    apiServerStatus: false,
    mysqlServerStatus: false,
    redisServerStatus: false,
    osType: '',
    osArch: '',
    uptime: '',
    ginMode: '',
    numCpu: 0,
    numCgoCall: 0,
    enabledPayMethods: []
  }
})

// TODO
let usersInfo = computed<{ name: string; data: number; suffix?: string; color: 'success'|'default'|'error' }[]>(() => [
  {
    name: 'adminViews.summary.userCard.allRegisteredUsers',
    data: miscellaneousConfig.value.userAll,
    color: 'success',
  },
  {
    name: 'adminViews.summary.userCard.activeUsers',
    data: miscellaneousConfig.value.userActive,
    // suffix: `${((miscellaneousConfig.value.userActive / miscellaneousConfig.value.userAll) * 100).toFixed(1)}%}`,
    color: 'success',
  },
  {
    name: 'adminViews.summary.userCard.inactiveUsers',
    data: miscellaneousConfig.value.userInactive,
    color: 'default',
  },
  {
    name: 'adminViews.summary.userCard.blockedUsers',
    data: miscellaneousConfig.value.userBlocked,
    color: 'error',
  }
])

let generalInfo = computed<{ name: string; data: string }[]>(() => [
  {
    name: 'adminViews.summary.general.localTime',
    data: computed(() => {
      return (new Date()).toString()
    }).value,
  },
  {
    name: 'adminViews.summary.general.osType',
    data: getOperatingSystem(),
  },
  {
    name: 'adminViews.summary.general.appName',
    data: appInfoStore.appCommonConfig.app_name,
  },
  {
    name: 'adminViews.summary.general.appUrl',
    data: appInfoStore.appCommonConfig.app_url,
  },
  {
    name: 'adminViews.summary.general.currency',
    data: appInfoStore.appCommonConfig.currency,
  },
  {
    name: 'adminViews.summary.general.allowRegister',
    data: !appInfoStore.appCommonConfig.stop_register ? t('adminViews.common.yes') : t('adminViews.common.no'),
  },
])

type ServerInfo = {
  type: 'text' | 'list' | 'tag';
  name: string;
  data: string | string[] | undefined | null;
}

let serverInfo = computed<ServerInfo[]>(() => [
  {
    type: 'text',
    name: 'adminViews.summary.system.axiosAddr',
    data: backendConfig.apiAddr.axiosAddr,
  },
  {
    type: 'text',
    name: 'adminViews.summary.system.wsAddr',
    data: backendConfig.apiAddr.wsAddr,
  },
  {
    type: 'text',
    name: 'adminViews.summary.system.serverTime',
    data: miscellaneousConfig.value.config.serverTime,
  },
  {
    type: 'text',
    name: 'adminViews.summary.system.uptime',
    data: miscellaneousConfig.value.config.uptime,
  },
  {
    type: 'tag',
    name: 'adminViews.summary.system.gatewayStatus',
    data: `${miscellaneousConfig.value.config.apiServerStatus}`,
  },
  {
    type: 'tag',
    name: 'adminViews.summary.system.dbStatus',
    data: `${miscellaneousConfig.value.config.mysqlServerStatus}`,
  },
  {
    type: 'tag',
    name: 'adminViews.summary.system.redisStatus',
    data: `${miscellaneousConfig.value.config.redisServerStatus}`,
  },
  {
    type: 'text',
    name: 'adminViews.summary.system.serverOsType',
    data: miscellaneousConfig.value.config.osType,
  },
  {
    type: 'text',
    name: 'adminViews.summary.system.serverOsArch',
    data: miscellaneousConfig.value.config.osArch,
  },
  {
    type: 'text',
    name: 'adminViews.summary.system.runMode',
    data: miscellaneousConfig.value.config.ginMode,
  },
  {
    type: 'text',
    name: 'adminViews.summary.system.cpuNums',
    data: `${miscellaneousConfig.value.config.numCpu.toString()} ${t('adminViews.summary.core')}`,
  },
  {
    type: 'text',
    name: 'adminViews.summary.system.numCgoCall',
    data: `${miscellaneousConfig.value.config.numCgoCall.toString()} ${t('adminViews.summary.system.time')}`,
  },
  {
    type: 'list',
    name: 'adminViews.summary.system.paymentMethods',
    data: miscellaneousConfig.value.config.enabledPayMethods,
  },

])

let getAppOverview = async () => {
  let data = await handleGetAppOverview()
  if (data && data.code === 200) {
    // console.log(data)
    if (data.api_access_count_history && data.income_count_history) {
      data.api_access_count_history.forEach((item: number) => apiAccessCount.value.unshift(item))
      data.income_count_history.forEach((item: number) => incomeCount.value.unshift(item))
    }
    yesterdayIncome.value = data.income_yesterday ? data.income_yesterday : 0.00
    monthIncome.value = data.income_this_month ? data.income_this_month : 0.00
    netFetchResult.value.push(1)
  } else {
    message.error(t('adminViews.common.fetchDataFailure'))
    netFetchResult.value.push(-1)
  }
}

let callGetUsersLayout = async () => {
  let data = await handleGetUserLayout()
  if (data && data.code === 200) {
    miscellaneousConfig.value.userAll = data.user_all
    miscellaneousConfig.value.userActive = data.user_active
    miscellaneousConfig.value.userInactive = data.user_inactive
    miscellaneousConfig.value.userBlocked = data.user_blocked
    netFetchResult.value.push(1)
  } else {
    message.error(t('adminViews.common.fetchDataFailure'))
    netFetchResult.value.push(-1)
  }
}

let callGetAppCommonConfig = async () => {
  let data = await handleFetchAppCommonConfig()
  console.log(data)
  if (data && data.code === 200) {
    miscellaneousConfig.value.config.serverTime = data.config.server_time
    miscellaneousConfig.value.config.enabledPayMethods = data.config.enabled_pay_methods
    miscellaneousConfig.value.config.apiServerStatus = data.config.api_server_status
    miscellaneousConfig.value.config.mysqlServerStatus = data.config.mysql_server_status
    miscellaneousConfig.value.config.redisServerStatus = data.config.redis_server_status
    miscellaneousConfig.value.config.osType = data.config.os_type
    miscellaneousConfig.value.config.osArch = data.config.os_arch
    miscellaneousConfig.value.config.uptime = data.config.uptime
    miscellaneousConfig.value.config.ginMode = data.config.gin_mode
    miscellaneousConfig.value.config.numCpu = data.config.num_cpu
    miscellaneousConfig.value.config.numCgoCall = data.config.num_cgo_call
    // console.log(data)
    netFetchResult.value.push(1)
  } else {
    message.error(t('adminViews.common.fetchDataFailure'))
    netFetchResult.value.push(-1)
  }
}

// let isNetDataFetchedSuccess = computed(() => netFetchResult.value.length === 3)

let showErrMessage = ref<boolean>(false)
onBeforeMount(async () => {
  await getAppOverview()  // 先获取数据
  await callGetUsersLayout()
  await callGetAppCommonConfig()
  let sum = 0;
  netFetchResult.value.forEach((i: number) => {
    if (i === -1) {
      return showErrMessage.value = true
    }
    sum += i
  })
  if (sum === 3) {
    animated.value = true
  }

})

onMounted(() => {

  console.log(netFetchResult)
  setTimeout(() => {
    reloadCharts()
  }, 500)

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

  <transition name="slide-fade">
    <PageHead
        v-if="showErrMessage"
        :title="t('adminViews.summary.reqErr')"
        :description="t('adminViews.summary.reqErrHint')"/>
  </transition>

  <transition name="slide-fade">
    <div class="root" v-if="animated">
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
            <n-icon size="50">
              <incomeIcon/>
            </n-icon>
          </div>
        </n-flex>
      </n-card>

      <!--      <Income :yesterday-income="yesterdayIncome" :this-month-income="monthIncome"></Income>-->
      <div>
        <n-grid cols="1 s:2" responsive="screen" :x-gap="14" :y-gap="14">
          <n-grid-item>
            <n-card class="user-count-card" hoverable :embedded="true" :bordered="false" content-style="padding: 0"
                    :title="t('adminViews.summary.userCard.title')">
              <n-table
                  :bordered="false"
                  :bottom-bordered="false"
                  style="background-color: rgba(0,0,0,0.0); padding: 0 14px 0 14px"
              >
                <n-tr style=" margin-left: 200px" v-for="(item, index) in usersInfo" :key="index">
                  <n-td style="background-color: rgba(0,0,0,0.0);">{{ t(item.name) }}</n-td>
                  <n-td style="background-color: rgba(0,0,0,0.0);">
                    <div style="display: flex; flex-direction: row">
                      <p style="opacity: 0.9">{{ item.data }}</p>
                      <NTag
                          v-if="index >= 1"
                          style="margin-left: 10px"
                          :type="item.color==='success'?'success':item.color==='error'?'error':'default'"
                          size="small"
                          :bordered="!themeStore.enableDarkMode">
                        {{ `${((item.data / miscellaneousConfig.userAll)*100).toFixed(2)} %` }}
                      </NTag>
                    </div>
                  </n-td>
                </n-tr>
              </n-table>
            </n-card>
            <n-card
                hoverable class="card1" :embedded="true" :bordered="false">
              <PageHead
                  style="margin: 0 !important; padding: 0 !important;"
                  :title="t('adminViews.summary.apiAccessCard')"
                  :description="t('adminViews.summary.apiAccessCardHint')"
                  secondary
              />
              <div class="visited" style="height: 280px" ref="visitedChartDOM"></div>
            </n-card>
            <n-card
                hoverable class="card2" :embedded="true" :bordered="false">
              <PageHead
                  style="margin: 0 !important; padding: 0 !important;"
                  :title="t('adminViews.summary.incomeWeek')"
                  :description="t('adminViews.summary.incomeCardHint')"
                  secondary
              />
              <div class="internal" style="height: 280px" ref="incomeChartDOM"></div>
            </n-card>
          </n-grid-item>
          <n-grid-item>

            <n-alert
                type="info"
                style="margin-bottom: 10px"
                :closable="true"
                :bordered="!themeStore.enableDarkMode"
                v-if="appInfoStore.appCommonConfig.stop_register"
            >
              {{ t('adminViews.summary.system.stopRegisterHint') }}
              <n-button
                  text
                  type="primary"
                  icon-placement="right"
                  @click="router.push({path: '/admin/dashboard/systemconfig'})"
                  style="text-decoration: underline"
              >
                {{ t('adminViews.summary.system.toSetting') }}
              </n-button>
            </n-alert>

            <n-card
                class="r-card"
                hoverable
                :embedded="true"
                :bordered="false"
                style="padding-bottom: 10px;"
                content-style="padding: 0"
                :title="t('adminViews.summary.general.title')"
            >
              <n-table
                  :bordered="false"
                  :bottom-bordered="false"
                  style="background-color: rgba(0,0,0,0.0); padding: 0 14px 0 14px"
              >
                <n-tr
                    style=" margin-left: 200px"
                    v-for="(item, index) in generalInfo"
                    :key="index"
                >
                  <n-td
                      style="background-color: rgba(0,0,0,0.0);"
                  >
                    {{ t(item.name) }}
                  </n-td>
                  <n-td
                      style="background-color: rgba(0,0,0,0.0);"
                  >
                    <p style="opacity: 0.9">{{ item.data }}</p>
                  </n-td>
                </n-tr>
              </n-table>
            </n-card>

            <n-alert
                type="error"
                style="margin-top: 10px"
                :closable="true"
                :bordered="!themeStore.enableDarkMode"
                v-if="!miscellaneousConfig.config.apiServerStatus ||
              !miscellaneousConfig.config.redisServerStatus ||
              !miscellaneousConfig.config.mysqlServerStatus"
            >
              {{ t('adminViews.summary.system.checkServer') }}
            </n-alert>

            <n-card
                style="margin-top: 10px"
                class="r-card"
                hoverable
                :embedded="true"
                :bordered="false"
                content-style="padding: 0"
                :title="t('adminViews.summary.system.title')"
            >
              <n-table
                  :bordered="false"
                  :bottom-bordered="false"
                  style="background-color: rgba(0,0,0,0.0); padding: 0 14px 10px 14px"
              >
                <n-tr
                    style=" margin-left: 200px"
                    v-for="(item, index) in serverInfo"
                    :key="index"
                >
                  <n-td
                      style="background-color: rgba(0,0,0,0.0);"
                  >
                    {{ t(item.name) }}
                  </n-td>
                  <n-td style="background-color: rgba(0,0,0,0.0);">
                    <p style="opacity: 0.9" v-if="item.type === 'text'">{{ item.data }}</p>
                    <NTag
                        v-if="item.type === 'tag' && item.data === 'true' || item.data === 'false'"
                        :bordered="!themeStore.enableDarkMode"
                        :type="item.data === 'true' ? 'success' : 'error'"
                        size="small"
                    >
                      {{ t(`adminViews.summary.system.${item.data === 'true' ? 'runOK' : 'runErr'}`) }}
                      <template #icon>
                        <n-icon v-if="item.data === 'true'"><successIcon/></n-icon>
                        <n-icon v-else><errIcon/></n-icon>
                      </template>
                    </NTag>
                    <div
                        style="display: flex; flex-direction: row;"
                        v-if="item.type === 'list'"
                    >
                      <NTag
                          :bordered="!themeStore.enableDarkMode"
                          size="small"
                          type="primary"
                          style="margin-right: 14px"
                          v-for="(method) in item.data"
                      >
                        {{ method }}
                      <template #icon><n-icon><successIcon/></n-icon></template>
                      </NTag>
                    </div>
                  </n-td>
                </n-tr>
              </n-table>

            </n-card>
          </n-grid-item>
        </n-grid>
      </div>
    </div>
  </transition>

</template>

<style scoped>
.root {
  padding: 14px 20px 20px 20px;
  display: flex;
  flex-direction: column;

  .income-part-root {
    display: flex;
    flex-direction: row;
    margin-bottom: 14px;
  }
}

.card1 {
  width: 100%;
  margin-bottom: 10px;

  .visited {
    background-color: rgba(0, 0, 0, 0.0);
    margin-top: 10px;
  }
}

.card2 {
  width: 100%;
  margin-bottom: 10px;

  .internal {
    background-color: rgba(0, 0, 0, 0.0);
    margin-top: 10px;

  }
}

.card3 {
  width: 100%;

  .internal {
    background-color: rgba(0, 0, 0, 0.0);
  }
}



.user-count-card {
  margin-bottom: 10px;
  padding-bottom: 10px;
}

.to-management-root {
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  font-weight: bold;
}


</style>
