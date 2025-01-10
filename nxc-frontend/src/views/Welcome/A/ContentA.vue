<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, onBeforeUnmount, onMounted, ref} from "vue";
import useAppInfosStore from "@/stores/useAppInfosStore";
import TreeFront from "./assets/TreeFront.svg";
import TreeMedium from "./assets/TreeMedium.svg";
import Sun from "./assets/Sun.svg";
import {useRouter} from "vue-router";

import {
  ChevronForward as userIcon,
  FileTrayFullOutline as mgrIcon,
  GlobeOutline as safeIcon,
  LockClosedOutline as lockIcon,
  SparklesOutline as fastIcon,
} from "@vicons/ionicons5";

let animated = ref<{
  treeFrontAnimate: boolean
  treeMedAnimate: boolean
  sunAnimate: boolean
  appNameAnimate: boolean
  appDescription: boolean
}>({
  treeFrontAnimate: false,
  treeMedAnimate: false,
  sunAnimate: false,
  appNameAnimate: false,
  appDescription: false
})

const {t} = useI18n();
const router = useRouter();

const appInfosStore = useAppInfosStore();
let iconSize = ref<number>(40)

const rootRef = ref<HTMLElement | null>(null);

// let page = rootRef.value
// let scrollTop = page.getBoundingClientRect().top * -1
// console.log(scrollTop)

// 创建一个 ref 来获取 lay-2 元素
const lay2Ref = ref<HTMLElement | null>(null);

// 创建一个 ref 来获取 lay-1 元素
const lay1Ref = ref<HTMLElement | null>(null);

let advanceOptions = ref([
  {
    id: 0,
    icon_id: 0,
    title: computed(() => t('welcome.A.browseSafe')),
    content: computed(() => t('welcome.A.browseSafeSub')),
  },
  {
    id: 1,
    icon_id: 1,
    title: computed(() => t('welcome.A.encrypt')),
    content: computed(() => t('welcome.A.encryptSub')),
  },
  {
    id: 2,
    icon_id: 2,
    title: computed(() => t('welcome.A.mgr')),
    content: computed(() => t('welcome.A.mgrSub')),
  },
  {
    id: 3,
    icon_id: 3,
    title: computed(() => t('welcome.A.fast')),
    content: computed(() => t('welcome.A.fastSub')),
  },
]);

let isScrolledToTarget = false;
let lastScrollTop = 0;

let onScrollFunction = () => {
  const lay1Element = lay1Ref.value;
  if (lay1Element) {
    const scrollTop = lay1Element.scrollTop;
    console.log(scrollTop);
    const isScrollingDown = scrollTop > lastScrollTop;
    lastScrollTop = scrollTop;
    if (!isScrolledToTarget && isScrollingDown && lay2Ref.value) {
      lay2Ref.value.scrollIntoView({behavior: "smooth"});
      isScrolledToTarget = true;
    } else if (isScrolledToTarget && !isScrollingDown) {
      // 添加反向滚动的处理逻辑，例如重置状态变量或执行其他操作
      isScrolledToTarget = false;
    }
  }
};

const props = defineProps<{
  app_sub_description: string
  why_choose_us_hint: string
}>()

// interface WelcomeSettings {
//   app_sub_description: string
//   why_choose_us_hint: string
//   bilibili_official_link: string
//   youtube_official_link: string
//   instagram_link: string
//   wechat_official_link: string
//   filing_number: string
//   page_suffix: string
// }

// let settings = ref<WelcomeSettings>({
//   app_sub_description: '',
//   why_choose_us_hint: '',
//   bilibili_official_link: '',
//   youtube_official_link: '',
//   instagram_link: '',
//   wechat_official_link: '',
//   filing_number: '',
//   page_suffix: '',
// })

// let handleGetConfig = async () => {
//   try {
//     let {data} = await instance.get('/api/app/v1/welcome', {
//       params: {
//         lang: 'en',
//       }
//     })
//     if (data.code === 200) {
//       console.log("欢迎页面:", data)
//       Object.assign(settings.value, data.config)
//     }
//   } catch (err: any) {
//     console.log(err + '')
//   }
// }


onMounted(() => {
  // // 监听 lay-1 的滚动事件
  // if (lay1Ref.value) {
  //   lay1Ref.value.addEventListener('scroll', onScrollFunction, true);
  // }
  // // 监听窗口滚动事件，自动触发滚动
  // window.addEventListener('scroll', onScrollFunction, true);
  setTimeout(() => animated.value.treeFrontAnimate = true, 300)
  setTimeout(() => animated.value.treeMedAnimate = true, 400)
  setTimeout(() => animated.value.sunAnimate = true, 500)
  setTimeout(() => {
    let intervalId = ref<undefined | number>(undefined)
    intervalId.value = setInterval(() => {
      if (props.app_sub_description !== '') {
        animated.value.appNameAnimate = true;
        clearInterval(intervalId.value)
      }
    }, 200)
  }, 450)
  setTimeout(() => animated.value.appDescription = true, 700)
});

// onBeforeMount(async () => {
//   await handleGetConfig()
// })

onBeforeUnmount(() => {
//   // 移除 lay-1 的滚动事件监听器
//   if (lay1Ref.value) {
//     lay1Ref.value.removeEventListener('scroll', onScrollFunction, true);
//   }
//   window.removeEventListener('scroll', onScrollFunction, true);
});
</script>

<script lang="ts">
export default {
  name: "ContentA",
};
</script>

<template>
  <div class="root">
    <n-scrollbar style="scrollbar-width: none" content-style="width: 0">
      <div class="lay-1" ref="lay1Ref">
        <transition name="slide2-fade">
          <div v-if="animated.treeFrontAnimate" class="image-wrapper img1">
            <TreeFront/>
          </div>
        </transition>

        <transition name="slide2-fade">
          <div v-if="animated.treeMedAnimate" class="image-wrapper img2">
            <TreeMedium/>
          </div>
        </transition>

        <transition name="slide2-fade">
          <div v-if="animated.sunAnimate" class="image-wrapper img3">
            <Sun/>
          </div>
        </transition>

        <!-- 引入悬浮左侧的介绍内容 -->
        <transition name="slide2-fade">
          <div v-if="animated.appNameAnimate" class="intro" style="z-index: 1600">
            <div v-if="animated.appNameAnimate" class="web-name">
              <p class="title">{{ t('welcome.A.welcomeTo') }}</p>
              <p class="app-name">{{ appInfosStore.appCommonConfig.app_name }}</p>
            </div>
            <p class="desc">
              <!--              {{ t('welcome.A.welcomeToSub') }}-->
              {{ props.app_sub_description }}
            </p>
            <n-button
                @click="router.push({path: '/dashboard/summary'})"
                size="large"
                class="btn-start"
                text>
              {{ t('welcome.A.startingUse') }}
              <n-icon style="margin-left: 12px" size="18">
                <userIcon/>
              </n-icon>
            </n-button>
          </div>
        </transition>


      </div>

      <!--    lay-2每一行兩個卡片，如果是移動端則每行一個垂直排列-->
      <div class="lay-2" ref="lay2Ref">
        <div class="why-us-head">
          <p class="why-us-title">{{ t('welcome.A.whyUs') }}</p>
          <!--          <p class="why-us-title">{{ settings.why_choose_us_hint }}</p>-->
          <p class="why-us-sub">
            {{ props.why_choose_us_hint }}
          </p>
        </div>

        <div class="items">
          <div class="advance-item" v-for="option in advanceOptions" :key="option.id">
            <div class="card-content">
              <div class="icon-container">
                <n-icon color="#1d3144" v-if="option.icon_id === 0" :size="iconSize">
                  <safeIcon/>
                </n-icon>
                <n-icon color="#1d3144" v-if="option.icon_id === 1" :size="iconSize">
                  <lockIcon/>
                </n-icon>
                <n-icon color="#1d3144" v-if="option.icon_id === 2" :size="iconSize">
                  <mgrIcon/>
                </n-icon>
                <n-icon color="#1d3144" v-if="option.icon_id === 3" :size="iconSize">
                  <fastIcon/>
                </n-icon>
              </div>
              <div class="text-container">
                <p class="title">{{ option.title }}</p>
                <p class="desc2">{{ option.content }}</p>
              </div>
            </div>
          </div>
        </div>

      </div>

    </n-scrollbar>
  </div>
</template>

<style scoped lang="less">

.slide2-fade-enter-active {
  transition: all 500ms ease;
}

.slide2-fade-leave-active {
  transition: all 500ms ease;
}

.slide2-fade-enter-from,
.slide2-fade-leave-to {
  transform: translateY(150px);
  opacity: 0;
}

.root {
  height: 100vh;
}

.lay-1 {
  position: relative;
  display: flex;
  flex-direction: column;
  height: 100vh; // 修改为占满视窗高度
  background-color: #c8d9eb;
  overflow: hidden;

  .image-wrapper {
    width: 100%;
    position: absolute;
    background-repeat: no-repeat;
    background-size: cover;
    background-position-x: center;
    padding: 0;
    margin: 0;
  }

  .img1 {
    padding: 0;
    z-index: 1003;
    bottom: -1vh; // 固定到浏览器底部
    left: 0;
  }

  .img2 {
    z-index: 1002;
    bottom: 0; // 固定到浏览器底部
    left: 0;
  }

  .img3 {
    z-index: 1001;
    bottom: 0; // 固定到浏览器底部
    left: 200px;
  }

  .intro {
    position: absolute;
    top: 20%; // 上方留空部分
    left: 8%;
    z-index: 1004;
    padding: 20px 0 0 0;
    transition: all 0.3s ease-in-out;

    .web-name {
      display: flex;
      flex-direction: row;
      align-items: baseline;

      .title {
        color: #1f364e;
        font-size: 3rem;
        padding-right: 10px;
      }

      .app-name {
        color: #1d3144;
        font-size: 3rem;
      }
    }

    .desc {
      color: rgba(0, 23, 58, 0.9);
      margin-top: 10px;
      position: relative;
      max-width: 600px;
      font-size: 1.2rem;
      opacity: 0.7;
    }

    .btn-start {
      margin-top: 30px;
      width: 50%;
      height: 50px;
      border-radius: 10px;
      background-color: #1d324d;
      color: white;
      box-shadow: rgba(29, 50, 77, 0.2) 5px 5px 20px 0;
      transition: ease 300ms;
    }

    .btn-start:hover {
      box-shadow: rgba(29, 50, 77, 0.5) 5px 5px 30px 0;

    }
  }

  .intro:hover {
    transform: translateY(-10px) translateX(0);
  }

  @media (max-width: 768px) {
    .intro {
      left: 50%;
      top: 40%; // 调整移动端上方留空部分
      width: 70%;
      transform: translate(-50%, -50%);
      font-size: 1rem;

      .web-name {
        display: flex;
        flex-direction: column;

        .title {
          font-size: 2rem;
        }

        .app-name {
          font-size: 2rem;
        }
      }

    }

    @media (max-width: 768px) {
      .intro:hover {
        transform: translate(-50%, calc(-50% - 10px)); // 同时向上移动10px
      }
    }
  }
}

.lay-2 {
  background-color: #c8d9eb;
  height: 600px;


  .why-us-head {
    padding: 50px 30px 0 30px;
    display: flex;
    flex-direction: column;
    align-items: center;

    .why-us-title {
      margin: 20px 0 20px 0;
      font-size: 2rem;
      font-weight: 600;
      opacity: 0.9;
      color: #091b44;
    }

    .why-us-sub {
      color: rgba(9, 27, 68, 1);
      width: 80%;
      font-size: 1.2rem;
      font-weight: 400;
      opacity: 0.7;
    }

    @media (max-width: 768px) {
      .why-us-sub {
        width: 90%;
      }
    }
  }


  .items {
    height: 400px;
    background-color: #c8d9eb;
    display: flex;
    /* display: flex; */
    justify-items: center;
    flex-wrap: wrap;
    flex-direction: row;
    align-items: center;
    justify-content: space-evenly;
    align-content: center;

    .advance-item {
      width: 40%;
      height: 25%;
      border-radius: 10px;
      margin: 0 0 40px 0;
      display: flex;
      flex-direction: column;
      align-items: center;
      transition: ease 200ms;

    }

    .advance-item:hover {
      transform: translateX(0) translateY(-10px);
    }

    //@media (max-width: 768px) {
    //  .advance-item {
    //    width: 90%;
    //  }
    //}
  }

  @media (max-width: 768px) {
    .items {
      height: 800px;
      background-color: #c8d9eb;
      padding-top: 30px;
      display: flex;
      /* display: flex; */
      justify-items: center;
      flex-wrap: wrap;
      flex-direction: column;
      align-items: center;
      justify-content: space-evenly;
      align-content: center;

      .advance-item {
        margin: 0 30px;
        width: 90%;
        height: 20%;
      }
    }
  }
}

.items {
  display: flex;
  flex-direction: column;
  justify-content: center; /* 垂直居中 */
  align-items: center; /* 水平居中 */
  height: 100%; /* 占满父容器高度 */

  .advance-item {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 40%;
    height: 25%;
    margin-bottom: 20px;
    border-radius: 10px;

    .card-content {
      display: flex;
      flex-direction: row;
      align-items: center; /* 水平行内居中 */
      //padding: 10px;
      width: 100%;
      //background-color: #66afe9;

    }
  }
}

.text-container {
  flex: 4;
  padding-left: 30px;
  display: flex;
  flex-direction: column;


  .title {
    font-weight: bold;
    font-size: 1.3rem;
    color: #1d3144;
  }

  .desc2 {
    font-size: 1rem;
    opacity: 0.7;
    color: rgba(9, 27, 68, 0.9);
  }
}


</style>
