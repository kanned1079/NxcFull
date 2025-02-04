<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, onBeforeMount, onMounted, onUnmounted, reactive, ref} from 'vue'
import useThemeStore from "@/stores/useThemeStore";
import useApiAddrStore from "@/stores/useApiAddrStore";
import {handleTestServerLatency} from "@/api/admin/test";
import {RefreshOutline as retryIcon} from "@vicons/ionicons5"

// let animated = ref<boolean>(false)
const {t} = useI18n();
const apiAddrStore = useApiAddrStore();

let animated = ref<boolean>(false)

// const titlePart1 = ['当前作业量', '近一小时处理量', '7日内报错数量']
const themeStore = useThemeStore()
let status = ref(true)
const statusColor = {
  error: '#f95555',
  running: '#3bc65c'
}


let hardwareInfo = reactive({
  cpu: {
    title: 'CPU型号',
    content: '',
    unit: '',
  },
  cores: {
    title: '总核心数',
    content: 0,
    unit: 'core(s)',
  },
  memory: {
    title: '内存大小',
    content: '',
    unit: 'GiB',
  },
  disk: {
    title: '根目录大小',
    content: '',
    unit: 'GiB',
  },
})

let osInfo = reactive({
  release: {
    title: '操作系统版本',
    content: ''
  },
  kernel: {
    title: '内核版本',
    content: ''
  },
  architecture: {
    title: '系统架构',
    content: ''
  },
  processNum: {
    title: '进程号',
    content: '',
  },
  goroutine: {
    title: '系统Go协程数',
    content: '',
  },
  gc: {
    title: '上一次GC回收时间',
    content: ''
  }
})

let intervalId = ref()

// let refresh = setInterval(() => {
//   getSysInfo()
// }, 3000)

let retryTestLatency = ref<boolean>(false)
let latencyTestIntervalId = ref<undefined | number>(undefined)
let latency = ref<number>(-1)
const callTestConnLatency = async () => {
  // const startTime = new Date();
  retryTestLatency.value = false
  let testStartTime: Date | null | undefined = null
  let testEndTime: Date | null | undefined = null

  // 定时器，每隔 1 秒测试连接并更新延迟
  latencyTestIntervalId.value = setInterval(async () => {
    testStartTime = new Date(); // 每次请求前记录开始时间
    if (await handleTestServerLatency()) {
      testEndTime = new Date(); // 每次请求完成后记录结束时间
      latency.value = testEndTime.getTime() - testStartTime.getTime(); // 计算请求的延迟时间（毫秒）
    } else {
      return latency.value = -1
    }
  }, 3000); // 每 1 秒执行一次请求

  setTimeout(() => {
    if (latencyTestIntervalId.value !== undefined) {
      clearInterval(latencyTestIntervalId.value);
      console.log('Stopped latency test');
      retryTestLatency.value = true

    }
  }, 30000); // 10 秒后自动停止测试
};

// const callTestConnLatency = () => {
//   // 初始化重试测试延迟标志为 false
//   retryTestLatency.value = false;
//
//   // 初始化测试开始时间和结束时间为 null
//   let testStartTime: Date | null | undefined = null;
//   let testEndTime = null;
//
//   // 定时器 ID
//   let latencyTestIntervalId;
//
//   // 定义一个内部函数来执行延迟测试
//   const performLatencyTest = () => {
//     // 记录测试开始时间
//     testStartTime = new Date();
//
//     // 调用处理服务器延迟测试的函数
//     handleTestServerLatency()
//         .then((success) => {
//           if (success) {
//             // 记录测试结束时间
//             testEndTime = new Date();
//             // 计算延迟时间（毫秒）
//             latency.value = testEndTime.getTime() - testStartTime.getTime();
//           } else {
//             // 测试失败，将延迟时间设为 -1
//             latency.value = -1;
//           }
//         })
//         .catch(() => {
//           // 处理错误，将延迟时间设为 -1
//           latency.value = -1;
//         });
//   };
//
//   // 每隔 5 秒执行一次延迟测试
//   latencyTestIntervalId = setInterval(performLatencyTest, 5000);
//
//   // 30 秒后自动停止测试
//   setTimeout(() => {
//     if (latencyTestIntervalId) {
//       // 清除定时器
//       clearInterval(latencyTestIntervalId);
//       console.log('Stopped latency test');
//       // 设置重试测试延迟标志为 true
//       retryTestLatency.value = true;
//     }
//   }, 30000);
// };

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

onMounted(() => {
  themeStore.contentPath = '/admin/dashboard/monitor'
  themeStore.menuSelected = 'queue-monitor'
  // themeStore.menuSelected = 'queue-monitor'
  // // 调用接口获取数据
  // getSysInfo()
  // intervalId.value = setInterval(() => {
  //   getSysInfo()
  // }, 3000)

  callTestConnLatency()

  animated.value = true
  // setTimeout(() => reloadCharts(), 1000)
  // reloadCharts()

  // setInterval(async () => handleTestServerLatency(), 10)

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
  clearInterval(intervalId.value)
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
          <!--          <div class="latency-inner-up">-->
          <!--            <p>连接状态： OK</p>-->
          <!--          </div>-->
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

  /*
  .card1 {
    width: 100%;

    .inner-card {
      display: flex;

      .part1 {
        flex: 1;
        height: 100px;
        border-radius: 4px;
        position: relative;

        .title {
          position: absolute;
          top: 3px;
          left: 3px;
          font-size: 1rem;
          opacity: 0.8;
        }

        .num {
          position: absolute;
          bottom: 10px;
          left: 3px;
          font-size: 30px;
        }

        .status {
          position: absolute;
          bottom: 10px;
          left: 3px;
          font-size: 30px;
          color: v-bind(nowStatusColor);
        }
      }
    }
  }

  .card2 {
    width: 100%;
    margin-top: 20px;

    .card1-inner {
      width: 100%;
      height: 150px;
      display: flex;

      .cpu-panel {
        flex: 1;
        text-align: center;
        width: 200px;
      }

      .mem-panel {
        flex: 1;
        text-align: center;
        width: 200px;

      }

      .disk-panel {
        flex: 1;
        text-align: center;
        width: 200px;
      }
    }

    .card3 {
      display: flex;

      .card2-inner {
        flex: 1;
        height: 320px;
        border-radius: 5px;
        margin-top: 30px;

        .title-card2 {
          font-size: 1.15rem;
          opacity: 0.8;
          margin-bottom: 20px;
        }

        .item-box {
          display: flex;
          margin-left: 5px;
          justify-content: left;
          line-height: 30px;

          .title {
            font-size: 1rem;
            margin-right: 5px;
            margin-bottom: 10px;
          }

        }
      }
    }


  }

   */
}


</style>