<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {type DataTableColumns, type DataTableInst, NButton, NIcon, NTag, useMessage} from "naive-ui"
import {computed, h, onBeforeMount, onMounted, onUnmounted, ref} from 'vue'
import useThemeStore from "@/stores/useThemeStore";
import {handleDeletePreviousLog, handleFetchServerStatus, handleTestServerLatency} from "@/api/admin/server";

import {
  CheckmarkDoneOutline as status200,
  CloseOutline as status500,
  EnterOutline,
  RefreshOutline as status404,
  ReorderFourOutline as recordIcon,
  ServerOutline as sizeIcon,
} from "@vicons/ionicons5"
import {formatDate} from "@/utils/timeFormat";
import DataTableSuffix from "@/views/utils/DataTableSuffix.vue";

const {t} = useI18n();
const message = useMessage()
const themeStore = useThemeStore()
const tableRef = ref<DataTableInst>()
let animated = ref<boolean>(false)

let searchCode = ref<number>(0)
let pageCount = ref<number>(0)
let dataSize = ref<{ pageSize: number, page: number }>({
  pageSize: 20,
  page: 1,
})

let retryTestLatency = ref<boolean>(false)
let latencyTestIntervalId = ref<undefined | number>(undefined)
let latency = ref<number>(-1)
const callTestConnLatency = async () => {
  retryTestLatency.value = false;
  let testStartTime: Date | null | undefined = null;
  let testEndTime: Date | null | undefined = null;

  // 使用一个标志来避免请求堆积
  let isTesting = false;

  // 定时器 每隔 3 秒测试连接并更新延迟
  latencyTestIntervalId.value = setInterval(async () => {
    if (isTesting) return; // 如果上一个请求还没有完成，跳过当前请求
    isTesting = true;

    testStartTime = new Date(); // 每次请求前记录开始时间

    if (await handleTestServerLatency()) {
      testEndTime = new Date();
      latency.value = testEndTime.getTime() - testStartTime.getTime(); // 计算请求的延迟时间（毫秒）
    } else {
      latency.value = -1; // 如果请求失败，设置延迟为 -1
    }

    isTesting = false; // 请求完成，允许下一次请求
  }, 1500);

  // 30秒后停止测试
  setTimeout(() => {
    if (latencyTestIntervalId.value !== undefined) {
      clearInterval(latencyTestIntervalId.value);
      console.log('Stopped latency test');
      retryTestLatency.value = true;
    }
  }, 30000); // 30秒后停止
};

let getLatencyHint = computed<{
  title: string
  description: string
  classColor: string
}>(() => {
  if (latency.value > 0 && latency.value <= 20) {
    return {
      title: 'adminViews.queueMonit.latency.level.l1.title', // 对应 l1
      description: 'adminViews.queueMonit.latency.level.l1.description',
      classColor: 'green',
    };
  } else if (latency.value > 20 && latency.value <= 100) {
    return {
      title: 'adminViews.queueMonit.latency.level.l2.title', // 对应 l2
      description: 'adminViews.queueMonit.latency.level.l2.description',
      classColor: 'green',
    };
  } else if (latency.value > 100 && latency.value <= 250) {
    return {
      title: 'adminViews.queueMonit.latency.level.l3.title', // 对应 l3
      description: 'adminViews.queueMonit.latency.level.l3.description',
      classColor: 'yellow',
    };
  } else if (latency.value > 250 && latency.value <= 500) {
    return {
      title: 'adminViews.queueMonit.latency.level.l4.title', // 对应 l4
      description: 'adminViews.queueMonit.latency.level.l4.description',
      classColor: 'red',
    };
  } else if (latency.value == -1) {
    return {
      title: 'adminViews.queueMonit.latency.level.offline.title', // 对应 l4
      description: 'adminViews.queueMonit.latency.level.offline.description',
      classColor: 'offline',
    };
  } else {
    return {
      title: 'adminViews.queueMonit.latency.level.l5.title', // 对应 l5
      description: 'adminViews.queueMonit.latency.level.l5.description',
      classColor: 'red',
    };
  }
});

interface ApiLog {
  id: number,
  level: string,
  method: string,
  path: string,
  status_code: number,
  response_time: number,
  request_at: string,
  client_ip: string,
  "created_at": string,
  "deleted_at": string,
  "updated_at": string,
}

const columns = computed<DataTableColumns<ApiLog>>(() => [
  {
    title: t('adminViews.queueMonit.log.table.id'),
    key: 'id',
  },
  {
    title: t('adminViews.queueMonit.log.table.method'),
    key: 'method',
    render(row: ApiLog) {
      return h(NTag, {
        size: 'small',
        bordered: false,
        type: 'default',
      }, {
        default: () => row.method
      })
    }
  },
  {
    title: t('adminViews.queueMonit.log.table.path'),
    key: 'path',
    render(row: ApiLog) {
      return h('p', {style: {textDecoration: 'underline'}}, {
        default: () => row.path
      })
    }
  },
  {
    title: t('adminViews.queueMonit.log.table.code'),
    key: 'status_code',
    render(row: ApiLog) {
      return h(NTag, {
        size: 'small',
        bordered: false,
        type: computed(() => {
          switch (row.status_code) {
            case 200:
              return 'success';
            case 404:
              return 'warning';
            case 500:
              return 'error';
            default:
              return 'default';
          }
        }).value
      }, {
        default: () => row.status_code
      })
    }
  },
  {
    title: t('adminViews.queueMonit.log.table.clientIp'),
    key: 'client_ip',
  },
  {
    title: t('adminViews.queueMonit.log.table.responseTime'),
    key: 'response_time',
    render(row: ApiLog) {
      return h('p', {}, {
        default: () => row.response_time.toFixed(2) + 'ms'
      })
    }
  },
  {
    title: t('adminViews.queueMonit.log.table.requestAt'),
    key: 'request_at',
    render(row: ApiLog) {
      return h('p', {}, {
        default: () => formatDate(row.request_at)
      })
    }
  },
])

interface ApiReq {
  status200: number,
  status404: number,
  status500: number,
  login_req: number,
  reg_req: number,
  code: number,
  log_table_rows_count: number,
  log_table_size: number,
  // api_log_list: ApiLog[],
}

let apiLogList = ref<ApiLog[]>([])
let apiReq = ref<ApiReq>({
  status200: 0,
  status404: 0,
  status500: 0,
  login_req: 0,
  reg_req: 0,
  code: 200,
  log_table_rows_count: 0,
  log_table_size: 0,
  // api_log_list: []
})

const apiItems = computed<{
  title: string,
  content: any,
  unit?: string
}[]>(() => [
  {
    title: 'adminViews.queueMonit.api.items.ok.title',
    content: apiReq.value.status200,
    unit: 'adminViews.queueMonit.api.items.ok.unit',
  },
  {
    title: 'adminViews.queueMonit.api.items.notFound.title',
    content: apiReq.value.status404,
    unit: 'adminViews.queueMonit.api.items.notFound.unit',
  },
  {
    title: 'adminViews.queueMonit.api.items.internalErr.title',
    content: apiReq.value.status500,
    unit: 'adminViews.queueMonit.api.items.internalErr.unit',
  },
])

const downloadCsv = () => {
  tableRef.value?.downloadCsv({ fileName: 'data-table' })
}

const callFetchServerStatus = async () => {
  animated.value = false
  let data = await handleFetchServerStatus(0, dataSize.value.page, dataSize.value.pageSize)
  if (data.code === 200) {
    Object.assign(apiReq.value, data)
    apiLogList.value = []
    data.api_log_list.forEach((log: ApiLog) => apiLogList.value.push(log))
    pageCount.value = data.page_count
    animated.value = true
  } else {

  }
  console.log(data)
}

let showMention = ref<boolean>(true)

const callDeletePreviousLog = async () => {
  let data = await handleDeletePreviousLog()
  if (data.code === 200) {
    message.success(t('adminViews.queueMonit.log.deleteLogMsg', {nums: data.rows_deleted}))
    await callFetchServerStatus()
  } else {
    message.error(t('adminViews.queueMonit.log.deleteLogMsg') + data.msg || '')
  }
}

onMounted(() => {
  themeStore.contentPath = '/admin/dashboard/monitor'
  themeStore.menuSelected = 'queue-monitor'


  callTestConnLatency()

  animated.value = true
  // setTimeout(() => reloadCharts(), 1000)
  // reloadCharts()

  // setInterval(async () => handleTestServerLatency(), 10)
  setTimeout(() => showMention.value = false, 5000)

  callFetchServerStatus()
})

onBeforeMount(() => {
  // 调用接口获取数据
  // getSysInfo()
  // intervalId.value = setInterval(() => {
  //   getSysInfo()
  // }, 3000)
})

onUnmounted(() => {
  console.log('queue组件卸载')
  // clearInterval()
  clearInterval(latencyTestIntervalId.value)
})

</script>

<script lang="ts">
export default {
  name: 'QueueMonitor',
}
</script>

<template>
  <transition name="slide-fade">
    <div v-if="animated" class="root">

      <n-collapse-transition :show="showMention">
        <n-alert
            type="info"
            :bordered="!themeStore.enableDarkMode"
            style="margin-bottom: 14px"
        >
          {{ t('adminViews.queueMonit.headerHint') }}
        </n-alert>
      </n-collapse-transition>


      <n-card
          hoverable
          :embedded="true"
          :bordered="false"
          class="latency-card"
          :title="t('adminViews.queueMonit.latency.title')"
      >
        <div
            class="latency-inner"
        >
          <div class="latency-inner-mid">
            <div class="l-i-mid-left">
              <!--              <p class="latency-nums">{{ latency }}</p>-->
              <n-statistic tabular-nums>
                <n-number-animation
                    :to="latency"
                    :duration="0"
                />
                <template #suffix>
                  {{ t('adminViews.queueMonit.latency.unit') }}
                </template>
              </n-statistic>

            </div>
            <div class="l-i-mid-right">
              <div class="latency-title">
                <div :class="`mention-color-general ${getLatencyHint.classColor}`"></div>
                <div class="latency-mention-title">
                  {{ t(getLatencyHint?.title || '') }}
                </div>
                <n-button
                    v-if="retryTestLatency"
                    type="primary"
                    text
                    style="text-decoration: underline; margin-left: 10px; font-size: 0.8rem"
                    icon-placement="right"
                    @click="callTestConnLatency"
                >
                  {{ t('adminViews.queueMonit.latency.retry') }}
                </n-button>

              </div>
              <div class="latency-description">
                {{ t(getLatencyHint?.description || '') }}
              </div>
            </div>
          </div>
          <div class="latency-inner-btn">
            <p>{{ t('adminViews.queueMonit.latency.hint') }}</p>
          </div>
        </div>
      </n-card>

      <n-card
          hoverable
          :embedded="true"
          :bordered="false"
          class="api-card"
          :title="t('adminViews.queueMonit.api.title')"
      >
        <div>
          <n-grid
              cols="2 s:2 m:4"
              responsive="screen"
              :x-gap="15"
              :y-gap="15"
          >
            <n-grid-item v-for="(i, index) in apiItems" :key="index">
              <n-statistic :label="t(i.title)" :value="i.content">
                <template #prefix>
                  <n-icon v-if="index === 0">
                    <status200/>
                  </n-icon>
                  <n-icon v-if="index === 1">
                    <status404/>
                  </n-icon>
                  <n-icon v-if="index === 2">
                    <status500/>
                  </n-icon>
                </template>
                <template #suffix>
                  {{ t(i.unit || '') }}
                </template>
              </n-statistic>
            </n-grid-item>


            <n-grid-item>
              <div style="display: flex; flex-direction: row">
                <n-statistic
                    :label="t('adminViews.queueMonit.api.items.login2reg.title')"
                    :value="apiReq.login_req"
                >
                  <template #prefix>
                    <n-icon>
                      <EnterOutline/>
                    </n-icon>

                  </template>
                  <template #suffix>
                    /
                  </template>
                </n-statistic>
                <n-statistic
                    :label="'*'"
                    :value="apiReq.reg_req"
                >
                  <template #suffix>
                    {{ t('adminViews.queueMonit.api.items.internalErr.unit') }}
                  </template>
                </n-statistic>
              </div>
            </n-grid-item>
          </n-grid>
        </div>
      </n-card>

      <n-card
          hoverable
          :embedded="true"
          :bordered="false"
          class="log-card"
          :title="t('adminViews.queueMonit.log.title')"
      >
        <div class="log-card-inner">
          <div class="log-card-inner-l">
            <n-grid
                cols="1 s:1 m:2"
                responsive="screen"
                :x-gap="15"
                :y-gap="15"
            >
              <n-grid-item
                  style="display: flex; flex-direction: row;"
              >
                <n-statistic :label="t('adminViews.queueMonit.log.logTableRows')" tabular-nums>
                  <n-number-animation ref="numberAnimationInstRef" :from="0" :to="apiReq.log_table_rows_count"/>
                  <template #prefix>
                    <n-icon>
                      <recordIcon/>
                    </n-icon>
                  </template>
                  <template #suffix>
                    {{ t('adminViews.queueMonit.log.unit.lines') }}
                  </template>
                </n-statistic>
              </n-grid-item>
              <n-grid-item>
                <n-statistic :label="t('adminViews.queueMonit.log.logTableSize')" tabular-nums>
                  <n-number-animation
                      :from="0"
                      :to="apiReq.log_table_size"
                      :precision="2"
                  />
                  <template #prefix>
                    <n-icon>
                      <sizeIcon/>
                    </n-icon>
                  </template>
                  <template #suffix>
                    {{ t('adminViews.queueMonit.log.unit.mb') }}
                  </template>
                </n-statistic>
              </n-grid-item>
            </n-grid>
          </div>
          <div class="log-card-inner-r">
            <n-button
                text
                size="small"
                type="warning"
                @click="callDeletePreviousLog"
                class="clear-btn"
            >
              {{ t('adminViews.queueMonit.log.deleteLog') }}
            </n-button>
            <n-button
                text
                size="small"
                type="primary"
                @click="downloadCsv"
                class="export-btn"
            >
              {{ t('adminViews.queueMonit.log.exportCsv') }}
            </n-button>
            <p class="clear-hint">{{ t('adminViews.queueMonit.log.deleteLogHint') }}</p>
          </div>
        </div>
      </n-card>

      <n-card
          hoverable
          :embedded="true"
          :bordered="false"
          class="api-log-card"
          content-style="padding: 0;"
      >
        <n-data-table
            ref="tableRef"
            :bordered="false"
            :columns="columns"
            :data="apiLogList"
            striped
        />
      </n-card>

      <DataTableSuffix
          v-model:data-size="dataSize"
          v-model:page-count="pageCount"
          v-model:animated="animated"
          :update-data="callFetchServerStatus"
      />

    </div>

  </transition>


</template>

<style scoped>
.root {
  padding: 20px;

  .latency-card {
    .latency-inner {
      .latency-inner-up {

      }

      .latency-inner-mid {
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: flex-start;

        .l-i-mid-left {
          width: auto;

          .latency-nums {
            font-size: 1.6rem;
            font-weight: bold;
          }
        }

        .l-i-mid-right {
          display: flex;
          flex-direction: column;
          margin-left: 16px;

          .latency-title {
            display: flex;
            flex-direction: row;
            align-items: center;

            .mention-color-general {
              width: 8px;
              height: 8px;
              border-radius: 50%;
              margin-right: 8px;
            }

            .red {
              background-color: #ff6666;
            }

            .yellow {
              background-color: #ffcc66;
            }

            .green {
              background-color: #66cc66;
            }

            .offline {
              background-color: #979797;
            }

            .latency-mention-title {
              font-size: 0.9rem;
              font-weight: bold;
            }
          }

          .latency-description {
            font-size: 0.8rem;
            opacity: 0.7;
            /* text-decoration: underline; */
          }
        }
      }

      .latency-inner-btn {
        margin-top: 14px;
        font-size: 0.8rem;
        opacity: 0.8;

      }
    }
  }

  .api-card {
    margin-top: 14px;
  }

  .log-card {
    margin-top: 14px;

    .log-card-inner {
      display: flex;
      flex-direction: row;
      justify-content: space-between;

      .log-card-inner-l {
        flex: 1;
      }

      .log-card-inner-r {
        flex: 1;
        .clear-btn {
          font-size: 1rem;
          font-weight: bold !important;
          text-decoration: underline;
        }
        .export-btn {
          font-size: 1rem;
          margin-left: 10px;
          text-decoration: underline;
        }
        .clear-hint {
          margin-top: 4px;
          opacity: .8;
        }
      }
    }
  }

  .api-log-card {
    margin-top: 14px;
  }


}


</style>