<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, onBeforeMount, onMounted, onUnmounted, reactive, ref} from 'vue'
import useThemeStore from "@/stores/useThemeStore";
import useApiAddrStore from "@/stores/useApiAddrStore";
import {handleTestServerLatency} from "@/api/admin/test";
import instance from "@/axios/index"
const {t} = useI18n();
const apiAddrStore = useApiAddrStore();
const themeStore = useThemeStore()
let animated = ref<boolean>(false)

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
      testEndTime = new Date(); // 每次请求完成后记录结束时间
      latency.value = testEndTime.getTime() - testStartTime.getTime(); // 计算请求的延迟时间（毫秒）
    } else {
      latency.value = -1; // 如果请求失败，设置延迟为 -1
    }

    isTesting = false; // 请求完成，允许下一次请求
  }, 3000);

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

const callFetchServerStatus = async () => {
  let {data} = await instance.get('/api/admin/v1/server/status/fetch')
  console.log(data)
}

let showMention = ref<boolean>(true)

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
      <!--      <n-card-->
      <!--          title="最近一周API接口访问次数"-->
      <!--          hoverable class="card1" :embedded="true" :bordered="false">-->
      <!--        <div class="visited" style="height: 280px" ref="visitedChartDOM">-->

      <!--        </div>-->
      <!--      </n-card>-->


      <n-collapse-transition :show="showMention">
        <n-alert
            type="info"
            :bordered="!themeStore.enableDarkMode"
            style="margin-bottom: 14px"
        >
          请勿长时间停留此页面，在网络高峰时段频繁查询将些许影响网关性能和数据库吞吐量。
        </n-alert>
      </n-collapse-transition>


      <n-card
          hoverable
          :embedded="true"
          :bordered="false"
          class="latency-card"
          title="服务器延迟"
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
                  ms
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
                  <!--                  <template #icon>-->
                  <!--                    <n-icon size="14">-->
                  <!--                      <retryIcon/>-->
                  <!--                    </n-icon>-->
                  <!--                  </template>-->
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


}


</style>