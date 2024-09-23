<script setup lang="ts" name="ThemeConfig">
import useThemeStore from "@/stores/useThemeStore";
import {computed, onMounted} from "vue";

const themeStore = useThemeStore();

interface ItemConfig {
  name: string,
  describe: string,
  backgroundImageUrl: string
}

let themeBg = computed(() => ({
  backgroundSize: 'cover',
  backgroundImage: `url(https://ikanned.com:24444/d/Upload/NXC/IMG_1152.png)`,
  backgroundPosition: 'center'
}))

// coverColor 设置不同颜色模式下的遮罩颜色
let coverColor = computed(() => {
  if (!themeStore.enableDarkMode) {
    return `linear-gradient(to right, rgba(255, 255, 255, 1) 10%, rgba(255, 255, 255, 0) 100%)`
  } else {
    return `linear-gradient(to right, ${themeStore.getTheme.globeTheme.cardBgColor} 10%, rgba(255, 255, 255, 0) 100%)`
  }
})

let itemBgImage = computed(() => {})

let allThemes = {
  theme1: {
    name: 'Journey\'s terminus.',
    describe: '这是使用的默认主题',
    backgroundImageUrl: 'https://ikanned.com:24444/d/Upload/NXC/ani24420.jpg',
    enabled: true,
  },
  theme2: {
    name: 'Naniwa',
    describe: '这是使用的默认主题',
    backgroundImageUrl: 'https://ikanned.com:24444/d/Upload/NXC/IMG_1152.png',
    enabled: false,
  },
  theme3: {
    name: 'Sky',
    describe: '这是使用的默认主题',
    backgroundImageUrl: 'https://ikanned.com:24444/d/Upload/NXC/101209570_p0.jpg',
    enabled: false,
  }
}

let handleSetTheme = (name: string) => {
  console.log(name)
}

let darkerImage = computed(() => {
  if (themeStore.enableDarkMode)
    return { opacity: 0.5 }
})

onMounted(() => {
  themeStore.menuSelected = 'theme-config'
  themeStore.contentPath = '/admin/dashboard/theme'
})

</script>

<template>
  <div class="root">
    <n-card hoverable :embedded="true" class="root-card" title="主题设置" :bordered="false">
    </n-card>

    <n-alert type="warning" :bordered="false" style="margin-bottom: 20px">
      如果你不采用前后分离的方式部署，那么主题配置将不会生效。
      <link rel="了解前后分离">
    </n-alert>

    <n-card v-for="item in allThemes" :key="item.name" class="theme-card" :embedded="true" hoverable :bordered="false" content-style="padding: 0;">
      <div class="content">
        <div class="l-content">
          <p class="theme-name">{{ item.name }}<n-tag v-if="item.enabled" style="margin-left: 10px" :bordered="false" type="success">使用中</n-tag></p>
          <p class="theme-desc">{{ item.describe }}</p>
        </div>
        <div class="r-content" :style="{backgroundImage: `url(${item.backgroundImageUrl})`, opacity: themeStore.enableDarkMode?0.5:1,}">
          <n-button tertiary type="primary" class="btn" @click="handleSetTheme(item.name)">设置</n-button>
        </div>
      </div>
    </n-card>



  </div>

</template>

<style lang="less" scoped>
.root {
  padding: 20px;

  .root-card {
    margin-bottom: 10px;
  }

  .r-content::before {
    content: "";
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    //background: linear-gradient(to right, rgba(255, 255, 255, 1) 0%, rgba(255, 255, 255, 0) 100%);
    background: v-bind('coverColor');
    //background-color: #66afe9;
    border-radius: 0 3px 3px 0;
  }

  .theme-card {
    display: flex;
    height: 120px;
    margin-bottom: 20px;
    .content {
      display: flex;
      height: 120px;

      .l-content {
        flex: 2;
        height: 120px;
        display: flex;
        flex-direction: column;
        justify-content: center;
        padding-left: 30px;
        //background-color: #66afe9;

        .theme-name {
          font-size: 16px;
          font-weight: bold;
          margin-bottom: 10px;
        }

        .theme-desc {
          opacity: 0.7;
        }
      }

      .r-content {
        flex: 3;
        text-align: right;
        display: flex;
        border-radius: 0 3px 3px 0;
        line-height: 120px;
        justify-content: right;
        position: relative;
        background-size: cover;
        background-repeat: no-repeat;
        background-position: center;

        .btn {
          width: 100px;
          margin-right: 50px;
          position: absolute;
          top: 50%;
          transform: translateY(-50%);
          opacity: 1;
          backdrop-filter: blur(50px);
          -webkit-backdrop-filter: blur(5px);
        }
      }

    }


  }
  .theme-card:hover {
    transition: transform 200ms ease;
    transform: translateY(-5px);
  }

  //.n-card:hover {
  //  transform: translateY(-2px);
  //  transition: ease 200ms;
  //}

  .test {
    width: 100%;
    height: 100px;
  }

}

</style>