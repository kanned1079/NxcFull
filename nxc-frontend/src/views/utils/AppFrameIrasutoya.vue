<script setup lang="ts">
import {computed} from "vue";
import {useI18n} from "vue-i18n";
import {useMessage} from "naive-ui";
import useThemeStore from "@/stores/useThemeStore";
import useAppInfosStore from "@/stores/useAppInfosStore";
// import computer_tokui_boy from "@/assets/computer_tokui_boy.png"
// import smartphone_monku from "@/assets/smartphone_monku.png"
// import smartphone_couple_school from "@/assets/smartphone_couple_school.png"
// import smartphone_nidaimochi_man from "@/assets/smartphone_nidaimochi_man.png"
import smartphone_demo from "@/assets/smartphone_demo.png"
import desktop_demo from "@/assets/desktop_demo.png"
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
const {t} = useI18n();
const message = useMessage();
const themeStore = useThemeStore();
const appInfoStore = useAppInfosStore()

type AppItem = {
  id?: number
  platform: string
  link: string
}

const props = defineProps<{
  platform: 'desktop' | 'mobile',
  iosAppLink?: string,
  androidAppLink?: string,
  windowsAppLink?: string,
  osxAppLink?: string,
  LinuxAppLink?: string,
  appItemOptions: AppItem[]
}>()

let getDockBgColor = computed(() => themeStore.enableDarkMode ? 'rgba(37,37,37,0.3)' : 'rgba(255,255,255,0.3)')
let getIconBgColor = computed(() => themeStore.enableDarkMode ? '#252525' : '#fff')

let downloadClick = (platform: string) => {
  // 数据信息
  console.log(props.appItemOptions)

  // 查找对应平台的下载链接
  const downloadOption = props.appItemOptions.find((item: AppItem) => item.platform === platform);

  // 如果找到相应的下载链接，打开链接
  if (downloadOption && downloadOption.link) {
    window.open(downloadOption.link, '_blank');  // 打开链接到新标签页
  } else {
    // 如果没有找到对应的平台，显示提示
    console.error(`No download link found for platform: ${platform}`);
    message.error(t('userDocument.noContentTitle'));
  }
}

</script>

<script lang="ts">
export default {
  name: 'AppFrameIrasutoya',
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
      >
        <div class="img-root">
          <img :style="themeStore.enableDarkMode?{filter: 'brightness(0.5)'}:null" alt="mobile" class="img-body" :src="smartphone_demo" />
        </div>
      </n-card>

      <div class="opa-root">
        <p class="card-bottom-platform-title">{{ t('userAppDownload.card.mobile.designFor') }}</p>
        <p class="card-bottom-platform-shallow">{{ t('userAppDownload.card.mobile.designShallow') }}</p>
        <n-button
            type="primary"
            text
            class="app-link"
            size="medium"
            icon-placement="right"
            @click="downloadClick('Android')"
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
            @click="downloadClick('IOS')"
        >
          {{ t('userAppDownload.card.mobile.iosDownloadShallow') }}
          <template #icon>
            <n-icon size="20">
              <downloadIcon/>
            </n-icon>
          </template>
        </n-button>
      </div>

    </div>

    <div
        v-if="props.platform === 'desktop'"
        class="show-desktop-part"
    >
      <n-card
          hoverable
          :embedded="true"
          :bordered="false"
      >
        <div class="img-root">
          <img :style="themeStore.enableDarkMode?{filter: 'brightness(0.5)'}:null" alt="mobile" class="img-body" :src="desktop_demo" />
        </div>
      </n-card>

      <div class="opa-root">
        <p class="card-bottom-platform-title">{{ t('userAppDownload.card.desktop.designFor') }}</p>
        <p class="card-bottom-platform-shallow">{{ t('userAppDownload.card.desktop.designShallow') }}</p>
        <n-button
            type="primary"
            text
            class="app-link"
            size="medium"
            icon-placement="right"
            @click="downloadClick('Windows')"
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
            @click="downloadClick('OSX')"
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
            @click="downloadClick('Linux')"
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

  </div>

</template>

<style scoped lang="less">
.root {
  display: flex;
  flex-direction: column;
  align-items: flex-start;

  .show-mobile-part {
    width: 100%;

    .img-root {
      //background-color: #252525;
      //background-color: #e3e5e7;
      display: flex;
      flex-direction: row;
      justify-content: center;
      align-items: center;

      .img-body {
        width: 54%;
        height: auto;
        //filter: brightness(0.5);
      }
    }
    .opa-root {
      display: flex;
      flex-direction: column;
      align-items: flex-start;
    }
  }

  .show-desktop-part {
    width: 100%;

    .img-root {
      //background-color: #252525;
      //background-color: #e3e5e7;
      display: flex;
      flex-direction: row;
      justify-content: center;
      align-items: center;

      .img-body {
        width: 54%;
        height: auto;
      }
    }
    .opa-root {
      display: flex;
      flex-direction: column;
      align-items: flex-start;
    }
  }
}

//.root {
//  .show-mobile-part {
//    display: flex;
//    flex-direction: row;
//    align-items: flex-start;
//
//    .mobile-card-root {
//      background: linear-gradient(to bottom right, rgb(211, 149, 155), rgb(191, 230, 186));
//      height: 450px;
//
//      .mobile-card-inner {
//        height: 100%;
//        border-radius: 3px;
//        display: flex;
//        flex-direction: column;
//        justify-content: end;
//        align-items: center;
//
//        .mobile-inner-board {
//          width: 320px;
//          height: 360px; /* 4/5的高度 */
//          border-radius: 16px 16px 0 0;
//
//          .mobile-header-root {
//            display: flex;
//            flex-direction: row;
//            justify-content: space-between;
//            align-items: center;
//            opacity: 0.7;
//
//            .mobile-header-content-l {
//              .header-time {
//                font-size: 1rem;
//                font-weight: bold;
//              }
//            }
//
//            .mobile-header-content-r {
//              display: flex;
//              flex-direction: row;
//              align-items: center;
//            }
//          }
//
//          .mobile-content-root {
//            margin-top: 50px;
//
//            .mobile-content-title-shallow {
//              opacity: 0.6;
//            }
//
//            .mobile-content-title {
//              font-size: 1.4rem;
//              font-weight: bold;
//              margin-bottom: 8px;
//            }
//          }
//        }
//      }
//    }
//
//
//  }
//
//
//  .show-desktop-part {
//    display: flex;
//    flex-direction: column;
//    align-items: flex-start;
//
//    .desktop-card-root {
//      background: linear-gradient(to bottom right, rgb(185, 147, 214), rgb(140, 166, 219));
//      height: 450px;
//
//      .desktop-card-inner {
//        height: 100%;
//        border-radius: 3px;
//        display: flex;
//        flex-direction: column;
//        justify-content: space-between;
//        align-items: center;
//
//        .desktop-card-app-board {
//          flex: 9;
//          //background-color: #00a4ef;
//
//          display: flex; /* 设置父容器为flex布局 */
//          justify-content: center; /* 水平居中 */
//          align-items: center; /* 垂直居中 */
//
//          .desktop-card-app-content {
//            width: 400px;
//            border-radius: 12px;
//            /* 可选：如果需要内部元素也居中，可以加这行 */
//
//            .desktop-header {
//              display: flex;
//              padding: 10px;
//
//              .dot {
//                width: 10px;
//                height: 10px;
//                border-radius: 50%;
//              }
//
//              .red {
//                background-color: #ff6666;
//              }
//
//              .yellow {
//                background-color: #ffcc66;
//              }
//
//              .green {
//                background-color: #66cc66;
//              }
//
//              .right-gap {
//                margin-right: 5px;
//              }
//            }
//
//            .desktop-content-root {
//              display: flex;
//              flex-direction: row;
//              justify-content: space-between;
//              padding-bottom: 10px;
//
//              .desktop-content-l {
//                width: 100%;
//                display: flex;
//                flex: 1;
//                flex-direction: row;
//                .menu-part-l {
//                  flex: 9;
//                  padding: 0 8px 30px 8px;
//                }
//                .menu-part-r {
//                  flex: 1;
//                  .menu-divider {
//                    background-color: #999999;
//                    height: 100%;
//                    width: 1px;
//                    opacity: 0.2;
//                  }
//                }
//
//
//              }
//              .desktop-content-r {
//                flex: 4;
//                padding: 0 30px 10px 10px;
//                .desktop-content-root {
//                  display: flex;
//                  flex-direction: column;
//                  .desktop-content-title-shallow {
//                    opacity: 0.6;
//                  }
//
//                  .desktop-content-title {
//                    font-size: 1.4rem;
//                    font-weight: bold;
//                    margin-bottom: 8px;
//                  }
//                }
//              }
//            }
//          }
//
//
//        }
//        .desktop-card-dock-board {
//          flex: 1;
//          display: flex;
//          flex-direction: column;
//          justify-content: flex-end;
//
//          .dock-card {
//            margin: 16px;
//            border-radius: 8px;
//            height: 36px;
//            width: 173px;
//
//            .app-list-panel {
//              display: flex;
//              flex-direction: row;
//              align-items: center;
//              justify-content: flex-start;
//              width: 100%;
//              height: 100%;
//              padding-left: 4px;
//              .app-item {
//                width: 28px;
//                height: 28px;
//                border-radius: 6px;
//                background-color: #e3e5e7;
//                margin-right: 4px;
//                display: flex;
//                flex-direction: column;
//                justify-content: center;
//                align-items: center;
//              }
//            }
//          }
//        }
//      }
//    }
//  }
//}
//
.card-bottom-platform-title {
  font-size: 1.6rem;
  margin-top: 14px;

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