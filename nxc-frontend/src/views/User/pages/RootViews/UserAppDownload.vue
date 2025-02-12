<script setup lang="ts">
import AppFrame from "@/views/utils/AppFrame.vue";
import {useI18n} from "vue-i18n";
import {useRoute} from "vue-router"
import {onBeforeMount, onMounted, ref} from "vue";
import useThemeStore from "@/stores/useThemeStore";
import {useMessage} from "naive-ui"
import {handleFetchAllAppDownloadList} from "@/api/user/app";
import AppFrameIrasutoya from "@/views/utils/AppFrameIrasutoya.vue";

const {t} = useI18n()
const route = useRoute()
const message = useMessage()
const themeStore = useThemeStore()
let animated = ref<boolean>(false)
let downloadEnabled = ref<boolean>(false)

interface AppItem {
  id?: number
  platform: string
  link: string
}

let appDownloadList = ref<AppItem[]>([])

let callFetchAllDownloadLink = async () => {
  let data = await handleFetchAllAppDownloadList('en')
  if (data.code === 200) {
    console.log(data)
    downloadEnabled.value = data.download_enabled
    data.download_links.forEach((item: AppItem) => appDownloadList.value.push(item))
    animated.value = true
  } else {
    message.error(data.message || '')
  }
}

onBeforeMount(() => {
  themeStore.breadcrumb = t('userAppDownload.title')
  themeStore.menuSelected = 'user-app-download'
  themeStore.userPath = route.fullPath
})

onMounted(async () => {
  await callFetchAllDownloadLink()
})

</script>

<script lang="ts">
export default {
  name: 'UserAppDownload',
}
</script>

<template>
  <div class="root">
    <transition name="slide-fade">
      <div class="inner-root" v-if="animated">

        <n-alert
            type="warning"
            class="mention"
            v-if="!downloadEnabled"
        >
          {{ t('userAppDownload.common.noDownload') }}
        </n-alert>

        <div class="up-title-root">
          <p class="download-title">
            {{ t('userAppDownload.common.title') }}
          </p>
          <p class="download-sub-title">
            {{ t('userAppDownload.common.shallow') }}
          </p>
        </div>

        <div
            v-if="downloadEnabled"
            style="width: 100%; display: flex; justify-content: center;"
        >
<!--          <n-grid-->
<!--              cols="1 s:1 m:2"-->
<!--              responsive="screen"-->
<!--              x-gap="20px"-->
<!--              y-gap="20px"-->
<!--              style="max-width: 1368px"-->
<!--              v-if="false"-->
<!--          >-->
<!--            <n-grid-item>-->

<!--              <AppFrame-->
<!--                  platform="mobile"-->
<!--              ></AppFrame>-->
<!--            </n-grid-item>-->
<!--            <n-grid-item>-->

<!--              <AppFrame-->
<!--                  platform="desktop"-->
<!--              ></AppFrame>-->
<!--            </n-grid-item>-->
<!--          </n-grid>-->

          <n-grid
              cols="1 s:1 m:2"
              responsive="screen"
              x-gap="20px"
              y-gap="20px"
              style="max-width: 1368px"
              v-if="true"
          >
            <n-grid-item>

              <AppFrameIrasutoya
                  platform="mobile"
              ></AppFrameIrasutoya>
            </n-grid-item>
            <n-grid-item>

              <AppFrameIrasutoya
                  platform="desktop"
              ></AppFrameIrasutoya>
            </n-grid-item>
          </n-grid>

        </div>

<!--        <AppFrameIrasutoya platform="desktop"></AppFrameIrasutoya>-->
<!--        <AppFrameIrasutoya platform="mobile"></AppFrameIrasutoya>-->




        <div class="app-page-all-suffix">
          <div class="suffix-hint-inner">
            <p
              v-for="item in 2"
              :key="item"
              class="suffix-font"
            >
              {{ t(`userAppDownload.suffix.p${item}`) }}
            </p>
          </div>
        </div>
      </div>


    </transition>
  </div>
</template>

<style scoped lang="less">
.root {
  padding: 20px;

  .inner-root {
    .mention {
      margin-bottom: 20px;
    }

    .up-title-root {
      text-align: center;

      .download-title {
        font-size: 1.5rem;
        margin: 20px;
      }

      .download-sub-title {
        font-size: 0.9rem;
        opacity: 0.7;
        margin: 0 30px 20px 30px;
        text-align: center !important;
      }
    }

    .platform-card {
      margin-bottom: 20px;
    }

    .app-page-all-suffix {
      padding: 0 100px;
    }
  }
}

.app-page-all-suffix{
  display: flex;
  flex-direction: row;
  justify-content: center;
  margin-top: 40px;
  .suffix-hint-inner {
    .suffix-font {
      font-size: 0.7rem;
      opacity: 0.8;
      margin-top: 4px;
    }
  }
}


</style>