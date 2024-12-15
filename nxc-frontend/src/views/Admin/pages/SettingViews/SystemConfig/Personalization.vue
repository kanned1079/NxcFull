<script setup lang="ts" name="">
import {onMounted, reactive} from 'vue'
import useThemeStore from '@/stores/useThemeStore'
import useSettingStore from "@/stores/useSettingStore";

const themeStore = useThemeStore()
const settingStore = useSettingStore()

const options = [
  {
    label: '默认',
    value: 'defaultDay',
    disabled: false,
  },
  {
    label: '深蓝色',
    value: 'darkBlueDay',
    disabled: false,
  },
  {
    label: '奶绿色',
    value: 'milkGreenDay',
    disabled: false,
  },
  {
    label: '若竹',
    value: 'bambooGreen',
    disabled: false,
  },

  {
    label: '雾松',
    value: 'mistyPine',
    disabled: false,
  },
  {
    label: '冰川蓝',
    value: 'glacierBlue',
    disabled: false,
  },


    /*
    *  case 'darkBlueDay': {
                return darkBlueDay.value;
            }
            case 'milkGreenDay': {
                return milkGreenDay.value;
            }
            case 'biliPink': {
                return biliPink.value;
            }
            case 'bambooGreen': {
                return bambooGreen.value;
            }
            case 'mistyPine': {
                return mistyPine.value;
            }
            case 'glacierBlue': {
    * */
]

const themeSettings = reactive({
  aside_theme: true,
  header_theme: false,
  background_image: ''
})

// let handok = () => {
//   console.log('ok')
// }

// onMounted(() => {
//   console.log(themeStore.allTheme)
// })

</script>

<template>
  <div class="root">
    <n-card :embedded="true" class="security-panel" title="个性化" :bordered="false">
      <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">边栏风格</p>
            <p class="shallow">设置侧边栏的颜色。</p>
          </div>
        </span>
        <span class="r-content" style="text-align: right">
          <n-switch
              size="medium"
              v-model:value="settingStore.settings.frontend.frontend_theme_sidebar"
              @update:value="settingStore.saveOption('frontend','frontend_theme_sidebar', settingStore.settings.frontend.frontend_theme_sidebar)"
          ></n-switch>
        </span>
      </div>
      <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">头部风格</p>
            <p class="shallow">设置顶部的颜色。</p>
          </div>
        </span>
        <span class="r-content" style="text-align: right">
          <n-switch
              size="medium"
              v-model:value="settingStore.settings.frontend.frontend_theme_header"
              @update:value="settingStore.saveOption('frontend','frontend_theme_header', settingStore.settings.frontend.frontend_theme_header)"
          ></n-switch>
        </span>
      </div>
      <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">主题色</p>
            <p class="shallow">设置整个网页的主题色。</p>
          </div>
        </span>
        <span class="r-content">
<!--          <n-select size="large" :options="options" v-model:value="themeStore.selectedTheme"/>-->
          <n-select
              size="large"
              :options="options"
              v-model:value="settingStore.settings.frontend.frontend_theme"
              @update:value="themeStore.setThemeFromSetting();settingStore.saveOption('frontend','frontend_theme', settingStore.settings.frontend.frontend_theme)"/>
        </span>
      </div>
      <div class="item">
        <span class="l-content">
          <div class="describe">
            <p class="title">背景</p>
            <p class="shallow">将会在后台登录页面进行展示。</p>
          </div>
        </span>
        <span class="r-content">
          <n-input
              size="large"
              v-model:value="settingStore.settings.frontend.frontend_background_url"
              placeholder="https://x.com/logo.jpeg"
              @blur="settingStore.saveOption('frontend','frontend_background_url', settingStore.settings.frontend.frontend_background_url)"
          ></n-input>
        </span>
      </div>
    </n-card>
  </div>
</template>

<style lang="less" scoped>
.root {
  min-width: 900px;

  .security-panel {
    .item {
      height: 50px;
      display: flex;
      margin-bottom: 30px;

      .l-content {
        flex: 1;

        .describe {
          .title {
            font-weight: bold;
          }

          .shallow {
            margin-top: 5px;
            opacity: 0.5;
          }
        }
      }

      .r-content {
        margin-left: 30px;
        flex: 0.8;
        justify-content: center;
        line-height: 50px;
      }
    }
  }
}


</style>