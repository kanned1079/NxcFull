<script setup lang="ts">
import {computed} from "vue";
import {useI18n} from "vue-i18n";
import {
  BatteryHalfOutline as batteryIcon,
  Cellular as dataIcon,
  ReturnDownForwardOutline as downloadIcon,
  Home as homeDirIcon,
  Apps as appsIcon,
  Globe as netIcon,
  Terminal as terminalIcon,
  Trash as trashIcon,
} from "@vicons/ionicons5"
import {useMessage} from "naive-ui";
import useThemeStore from "@/stores/useThemeStore";
import useAppInfosStore from "@/stores/useAppInfosStore";

const {t} = useI18n();
const message = useMessage();
const themeStore = useThemeStore();
const appInfoStore = useAppInfosStore()
const props = defineProps<{
  platform: 'desktop' | 'mobile',
  iosAppLink?: string,
  androidAppLink?: string,
  windowsAppLink?: string,
  osxAppLink?: string,
  LinuxAppLink?: string,
}>()

let getDockBgColor = computed(() => themeStore.enableDarkMode?'rgba(37,37,37,0.3)':'rgba(255,255,255,0.3)')
let getIconBgColor = computed(() => themeStore.enableDarkMode?'#252525':'#fff')

let downloadClick = (platform: string) => {
  message.info(platform);
  window.open('https://ikanned.com:2444/d/R730xd_SSD/61256422_p0.jpg')
}

</script>

<script lang="ts">
export default {
  name: 'AppFrame',
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
                <n-icon size="24">
                  <batteryIcon/>
                </n-icon>
              </div>
            </div>
            <div class="mobile-content-root">
              <p class="mobile-content-title-shallow">{{ t('userAppDownload.card.common.welcome') }}</p>
              <p class="mobile-content-title">{{ appInfoStore.appCommonConfig.app_name || 'App' }}</p>
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
      <p class="card-bottom-platform-title">{{ t('userAppDownload.card.mobile.designFor') }}</p>
      <p class="card-bottom-platform-shallow">{{ t('userAppDownload.card.mobile.designShallow') }}</p>
      <n-button
          type="primary"
          text
          class="app-link"
          size="medium"
          icon-placement="right"
          @click="downloadClick('android')"
      >
        {{ t('userAppDownload.card.mobile.androidDownloadShallow') }}
        <template #icon>
          <n-icon size="20">
            <downloadIcon/>
          </n-icon>
        </template>
      </n-button>
      <n-button
          type="primary"
          text
          class="app-link"
          size="medium"
          icon-placement="right"
          @click="downloadClick('ios')"
      >
        {{ t('userAppDownload.card.mobile.iosDownloadShallow') }}
        <template #icon>
          <n-icon size="20">
            <downloadIcon/>
          </n-icon>
        </template>
      </n-button>
    </div>

    <div
        v-if="props.platform === 'desktop'"
        class="show-desktop-part"
    >
      <n-card
          hoverable
          :embedded="true"
          :bordered="false"
          content-style="padding: 0;"
          class="desktop-card-root"
      >
        <div class="desktop-card-inner">
          <div class="desktop-card-app-board">
            <n-card
                hoverable
                :embedded="true"
                :bordered="false"
                content-style="padding: 0"
                class="desktop-card-app-content"
            >
              <div class="desktop-header">
                <div class="dot red right-gap"></div>
                <div class="dot yellow right-gap"></div>
                <div class="dot green right-gap"></div>
              </div>
              <div class="desktop-content-root">
                <div class="desktop-content-l">
                  <div class="menu-part-l">
                    <n-skeleton :height="6" style="margin: 5px 0 20px 0"></n-skeleton>
                    <div v-for="i in 3" :key="i">
                      <n-skeleton :height="3" width="40%" style="margin-top: 14px"></n-skeleton>
                      <n-skeleton :height="2" style="margin-top: 6px"></n-skeleton>
                      <n-skeleton :height="2" style="margin-top: 6px"></n-skeleton>
                      <n-skeleton :height="2" style="margin-top: 6px"></n-skeleton>
                    </div>
                  </div>
                  <div class="menu-part-r">
                    <div class="menu-divider"></div>
                  </div>
                </div>
                <div class="desktop-content-r">
                  <div class="desktop-content-root">
                    <p class="desktop-content-title-shallow">{{ t('userAppDownload.card.common.welcome') }}</p>
                    <p class="desktop-content-title">{{ appInfoStore.appCommonConfig.app_name || 'App' }}</p>
                    <n-skeleton width="30%" :height="8" style="margin-top: 20px"></n-skeleton>
                    <n-skeleton :height="4" style="margin-top: 10px"></n-skeleton>
                    <n-skeleton :height="4" width="78%" style="margin-top: 8px"></n-skeleton>
                    <n-skeleton :height="4" style="margin-top: 8px"></n-skeleton>
                    <n-skeleton :height="4" width="47%" style="margin-top: 8px"></n-skeleton>
                  </div>
                </div>
              </div>
            </n-card>
          </div>
          <div class="desktop-card-dock-board">
            <n-card
              hoverable
              :embedded="true"
              :bordered="false"
              content-style="padding: 0"
              class="dock-card"
              :style="{backgroundColor: getDockBgColor,}"
            >
              <div class="app-list-panel">
                <div class="app-item" :style="{backgroundColor: getIconBgColor}">
                  <n-icon size="16" style="opacity: 0.5"><homeDirIcon /></n-icon>
                </div>
                <div class="app-item" :style="{backgroundColor: getIconBgColor}">
                  <n-icon size="16" style="opacity: 0.5"><netIcon /></n-icon>
                </div>
                <div class="app-item" :style="{backgroundColor: getIconBgColor}">
                  <n-icon size="16" style="opacity: 0.5"><appsIcon /></n-icon>
                </div>
                <div style="width: 1px; height: calc(100% - 6px); background-color: rgba(51, 51, 51, 0.1); margin: 0 4px;" />
                <div class="app-item" :style="{backgroundColor: getIconBgColor}">
                  <n-icon size="16" style="opacity: 0.5"><terminalIcon /></n-icon>
                </div>
                <div class="app-item" :style="{backgroundColor: getIconBgColor}">
                  <n-icon size="16" style="opacity: 0.5"><trashIcon /></n-icon>
                </div>
              </div>

            </n-card>
          </div>
        </div>
        <!--          <p style="font-size: 1.5rem">桌面端</p>-->

      </n-card>
      <p class="card-bottom-platform-title">{{ t('userAppDownload.card.desktop.designFor') }}</p>
      <p class="card-bottom-platform-shallow">{{ t('userAppDownload.card.desktop.designShallow') }}</p>
      <n-button
          type="primary"
          text
          class="app-link"
          size="medium"
          icon-placement="right"
          @click="downloadClick('windows')"
      >
        {{ t('userAppDownload.card.desktop.windowsDownloadShallow') }}
        <template #icon>
          <n-icon size="20">
            <downloadIcon/>
          </n-icon>
        </template>
      </n-button>
      <n-button
          type="primary"
          text
          class="app-link"
          size="medium"
          icon-placement="right"
          @click="downloadClick('osx')"
      >
        {{ t('userAppDownload.card.desktop.osxDownloadShallow') }}
        <template #icon>
          <n-icon size="20">
            <downloadIcon/>
          </n-icon>
        </template>
      </n-button>
      <n-button
          type="primary"
          text
          class="app-link"
          size="medium"
          icon-placement="right"
          @click="downloadClick('linux')"
      >
        {{ t('userAppDownload.card.desktop.linuxDownloadShallow') }}
        <template #icon>
          <n-icon size="20">
            <downloadIcon/>
          </n-icon>
        </template>
      </n-button>
    </div>

  </div>

</template>

<style scoped lang="less">
.root {
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


  }


  .show-desktop-part {
    display: flex;
    flex-direction: column;
    align-items: flex-start;

    .desktop-card-root {
      background: linear-gradient(to bottom right, rgb(185, 147, 214), rgb(140, 166, 219));
      height: 450px;

      .desktop-card-inner {
        height: 100%;
        border-radius: 3px;
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        align-items: center;

        .desktop-card-app-board {
          flex: 9;
          //background-color: #00a4ef;

          display: flex; /* 设置父容器为flex布局 */
          justify-content: center; /* 水平居中 */
          align-items: center; /* 垂直居中 */

          .desktop-card-app-content {
            width: 400px;
            border-radius: 12px;
            /* 可选：如果需要内部元素也居中，可以加这行 */

            .desktop-header {
              display: flex;
              padding: 10px;

              .dot {
                width: 10px;
                height: 10px;
                border-radius: 50%;
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

              .right-gap {
                margin-right: 5px;
              }
            }

            .desktop-content-root {
              display: flex;
              flex-direction: row;
              justify-content: space-between;
              padding-bottom: 10px;

              .desktop-content-l {
                width: 100%;
                display: flex;
                flex: 1;
                flex-direction: row;
                .menu-part-l {
                  flex: 9;
                  padding: 0 8px 30px 8px;
                }
                .menu-part-r {
                  flex: 1;
                  .menu-divider {
                    background-color: #999999;
                    height: 100%;
                    width: 1px;
                    opacity: 0.2;
                  }
                }


              }
              .desktop-content-r {
                flex: 4;
                padding: 0 30px 10px 10px;
                .desktop-content-root {
                  display: flex;
                  flex-direction: column;
                  .desktop-content-title-shallow {
                    opacity: 0.6;
                  }

                  .desktop-content-title {
                    font-size: 1.4rem;
                    font-weight: bold;
                    margin-bottom: 8px;
                  }
                }
              }
            }
          }


        }
        .desktop-card-dock-board {
          flex: 1;
          display: flex;
          flex-direction: column;
          justify-content: flex-end;

          .dock-card {
            margin: 16px;
            border-radius: 8px;
            height: 36px;
            width: 173px;

            .app-list-panel {
              display: flex;
              flex-direction: row;
              align-items: center;
              justify-content: flex-start;
              width: 100%;
              height: 100%;
              padding-left: 4px;
              .app-item {
                width: 28px;
                height: 28px;
                border-radius: 6px;
                background-color: #e3e5e7;
                margin-right: 4px;
                display: flex;
                flex-direction: column;
                justify-content: center;
                align-items: center;
              }
            }
          }
        }
      }
    }
  }
}

.card-bottom-platform-title {
  font-size: 1.6rem;
  margin-top: 24px;

}

.card-bottom-platform-shallow {
  font-size: 0.9rem;
  opacity: 0.7;
  margin-top: 10px;
  margin-bottom: 30px;
}

.app-link {
  margin-bottom: 10px;
}

</style>