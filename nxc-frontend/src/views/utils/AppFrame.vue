<script setup lang="ts">
import {
  BatteryHalfOutline as batteryIcon,
  Cellular as dataIcon,
  Wifi as wifiIcon,
  ReturnDownForwardOutline as downloadIcon,
} from "@vicons/ionicons5"
import useAppInfosStore from "@/stores/useAppInfosStore";

const appInfoStore = useAppInfosStore()
const props = defineProps<{
  platform: 'desktop' | 'mobile'

}>()


</script>

<script lang="ts">
export default {
  name: 'DesktopAppFrame',
}
</script>

<template>
  <div class="root">
    <div
        v-if="props.platform === 'mobile'"
        class="show-mobile-part"
    >
      <n-card
          hoverable
          :embedded="true"
          :bordered="false"
          class="mobile-card-root"
          content-style="padding: 0;"
      >
        <div class="mobile-card-inner">
          <n-card
              :embedded="true"
              :bordered="false"
              class="mobile-inner-board"
              content-style="padding: 10px 20px"
          >
            <div class="mobile-header-root">
              <div class="mobile-header-content-l">
                <p class="header-time">23:59</p>
              </div>
              <div class="mobile-header-content-r">
                <n-icon size="20" style="margin-right: 10px">
                  <dataIcon/>
                </n-icon>
                <n-icon size="20" style="margin-right: 10px">
                  <wifiIcon/>
                </n-icon>
                <n-icon size="24">
                  <batteryIcon/>
                </n-icon>
              </div>
            </div>
            <div class="mobile-content-root">
              <p class="mobile-content-title-shallow">欢迎来到</p>
              <p class="mobile-content-title">{{ appInfoStore.appCommonConfig.app_name || 'Nxc Networks' }}</p>
              <n-skeleton :height="10" width="60%" style="margin-top: 10px"></n-skeleton>
              <n-skeleton width="30%" :height="14" style="margin-top: 50px"></n-skeleton>
              <n-skeleton :height="10" style="margin-top: 10px"></n-skeleton>
              <n-skeleton :height="10" width="78%" style="margin-top: 8px"></n-skeleton>
              <n-skeleton :height="10" style="margin-top: 8px"></n-skeleton>
              <n-skeleton :height="10" width="47%" style="margin-top: 8px"></n-skeleton>
            </div>
          </n-card>
        </div>
      </n-card>
      <p class="mobile-card-bottom-platform-title">For Mobile</p>
      <p class="mobile-card-bottom-platform-shallow">
        习近平总书记十分重视历史文化资源的挖掘、利用，多次发表重要讲话，满载对传承和发展中华优秀传统文化的期许。</p>

      <n-button
          type="primary"
          text
          class="mobile-app-link"
          size="large"
          icon-placement="right"
      >
        下载 Android 客户端
        <template #icon>
          <n-icon size="20">
            <downloadIcon />
          </n-icon>
        </template>
      </n-button>
      <n-button
          type="primary"
          text
          class="mobile-app-link"
          size="large"
          icon-placement="right"
      >
        下载 IOS 客户端
        <template #icon>
          <n-icon size="20">
            <downloadIcon />
          </n-icon>
        </template>
      </n-button>
    </div>
    <div v-if="props.platform === 'desktop'">
      <n-card
          hoverable
          :embedded="true"
          :bordered="false"
          content-style="padding: 0;"
          style="background: linear-gradient(to bottom right, rgb(185, 147, 214), rgb(140, 166, 219)); height: 500px;"
      >
          <p style="font-size: 1.5rem">桌面端</p>
<!--        <div style="width: 100%; margin-top: 100px; display: flex; flex-direction: row; justify-content: center">-->
<!--          <n-card style="width: 400px; border-radius: 10px" content-style="padding: 0;">-->
<!--            <div-->
<!--                style="display: flex; flex-direction: row; justify-content: space-between; align-items: center; margin-top: 10px">-->
<!--              <div style="opacity: 0.9; display: flex; margin-left: 15px">-->
<!--                &lt;!&ndash;                        <p style="font-size: 0.8rem; margin: 0 0 0 15px;">10:13</p>&ndash;&gt;-->
<!--                <div-->
<!--                    style="width: 10px; height: 10px; border-radius: 50%; background-color: #ff6666; margin-right: 5px"></div>-->
<!--                <div-->
<!--                    style="width: 10px; height: 10px; border-radius: 50%; background-color: #ffcc66; margin-right: 5px"></div>-->
<!--                <div-->
<!--                    style="width: 10px; height: 10px; border-radius: 50%; background-color: #66cc66; margin-right: 5px"></div>-->
<!--              </div>-->

<!--            </div>-->
<!--            <div style="margin-top: 30px; padding: 20px">-->
<!--              <n-skeleton :width="100" :height="20"></n-skeleton>-->
<!--              <n-skeleton :width="150" :height="12" style="margin-top: 10px"></n-skeleton>-->
<!--              <n-skeleton :height="10" style="margin-top: 50px"></n-skeleton>-->
<!--              <n-skeleton :height="10" style="margin-top: 10px"></n-skeleton>-->
<!--              <n-skeleton :height="10" style="margin-top: 10px"></n-skeleton>-->
<!--              <n-skeleton :height="10" style="margin-top: 10px"></n-skeleton>-->
<!--            </div>-->
<!--          </n-card>-->
<!--        </div>-->
      </n-card>
    </div>
  </div>

</template>

<style scoped lang="less">
.show-mobile-part {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  .mobile-card-root {
    background: linear-gradient(to bottom right, rgb(211, 149, 155), rgb(191, 230, 186));
    height: 450px;

    .mobile-card-inner {
      height: 100%;
      border-radius: 3px;
      display: flex;
      flex-direction: column;
      justify-content: end;
      align-items: center;

      .mobile-inner-board {
        width: 320px;
        height: 360px; /* 4/5的高度 */
        border-radius: 16px 16px 0 0;

        .mobile-header-root {
          display: flex;
          flex-direction: row;
          justify-content: space-between;
          align-items: center;
          opacity: 0.7;

          .mobile-header-content-l {
            .header-time {
              font-size: 1rem;
              font-weight: bold;
            }
          }

          .mobile-header-content-r {
            display: flex;
            flex-direction: row;
            align-items: center;
          }
        }

        .mobile-content-root {
          margin-top: 50px;

          .mobile-content-title-shallow {
            opacity: 0.6;
          }

          .mobile-content-title {
            font-size: 1.4rem;
            font-weight: bold;
            margin-bottom: 8px;
          }
        }
      }
    }
  }
  .mobile-card-bottom-platform-title {
    font-size: 1.5rem;
    margin-top: 24px;

  }
  .mobile-card-bottom-platform-shallow {
    font-size: 1rem;
    opacity: 0.7;
    margin-top: 10px;
    margin-bottom: 30px;
  }
  .mobile-app-link {
    margin-bottom: 20px;
  }
}


.desktop-card-root {

}


</style>