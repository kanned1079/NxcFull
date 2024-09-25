<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {onBeforeUnmount, onMounted, ref} from "vue";
import TreeFront from "./assets/TreeFront.svg";
import TreeMedium from "./assets/TreeMedium.svg";
import Sun from "./assets/Sun.svg";
import {
  FileTrayFullOutline as mgrIcon,
  GlobeOutline as safeIcon,
  LockClosedOutline as lockIcon,
  SparklesOutline as fastIcon,
} from "@vicons/ionicons5";

const {t} = useI18n();

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
    title: "瀏覽安全",
    content: "優秀的防火牆過濾系統能有效防禦網路釣魚和惡意網站",
  },
  {
    id: 1,
    icon_id: 1,
    title: "端到端加密",
    content: "雙向 SSL 和端對端加密保護您的隱私安全，即使是服務器也無法讀取您的信息",
  },
  {
    id: 2,
    icon_id: 2,
    title: "高效管理",
    content: "一個用戶介面管理所有密鑰，管理功能完善豐富，無須擔心訂閱洩露問題",
  },
  {
    id: 3,
    icon_id: 3,
    title: "方便快捷",
    content: "提供完整的 API 文檔供 WebApp 或是嵌入到第三方軟件中",
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
      lay2Ref.value.scrollIntoView({ behavior: "smooth" });
      isScrolledToTarget = true;
    } else if (isScrolledToTarget &&!isScrollingDown) {
      // 添加反向滚动的处理逻辑，例如重置状态变量或执行其他操作
      isScrolledToTarget = false;
    }
  }
};


onMounted(() => {
  // 监听 lay-1 的滚动事件
  if (lay1Ref.value) {
    lay1Ref.value.addEventListener('scroll', onScrollFunction, true);
  }
  // 监听窗口滚动事件，自动触发滚动
  window.addEventListener('scroll', onScrollFunction, true);
});

onBeforeUnmount(() => {
  // 移除 lay-1 的滚动事件监听器
  if (lay1Ref.value) {
    lay1Ref.value.removeEventListener('scroll', onScrollFunction, true);
  }
  window.removeEventListener('scroll', onScrollFunction, true);
});
</script>

<script lang="ts">
export default {
  name: "ContentA",
};
</script>

<template>
  <div class="root">
    <!--    <n-scrollbar style="height: 100vh"  >-->
    <div class="lay-1" ref="lay1Ref">
      <div class="image-wrapper img1"><img :src="TreeFront" alt="Image 1"/></div>
      <div class="image-wrapper img2"><img :src="TreeMedium" alt="Image 2"/></div>
      <div class="image-wrapper img3"><img :src="Sun" alt="Image 3"/></div>
      <!-- 引入悬浮左侧的介绍内容 -->
      <div class="intro">
        <div class="web-name">
          <p class="title">{{ t('welcome.A.welcomeTo') }}</p>
          <p class="app-name">{{ 'app_name' }}</p>
        </div>
        <p class="desc">
          “穿过县境上长长的隧道，便是雪国。夜空下，大地一片莹白，火车在信号所前停下来。”在这里川端康成用几近吝啬的简洁文字，拉开了《雪国》的序幕。
        </p>
        <n-button size="large" class="btn-start" text>开始使用</n-button>
      </div>

    </div>

    <!--    lay-2每一行兩個卡片，如果是移動端則每行一個垂直排列-->
    <div class="lay-2" ref="lay2Ref">
      <div class="why-us-head">
        <p class="why-us-title">為什麼選擇我們</p>
        <p class="why-us-sub">
          為什麼選擇我們為什麼選擇我們為什麼選擇我們為什麼選擇我們為什麼選擇我們為什麼選擇我們為什麼選擇我們為什麼選擇我們為什麼選擇我們為什麼選擇我們為什麼選擇我們
        </p>
      </div>

      <div class="items">
        <div class="advance-item" v-for="option in advanceOptions" :key="option.id">
          <div class="card-content">
            <div class="icon-container">
              <n-icon v-if="option.icon_id === 0" size="50">
                <safeIcon/>
              </n-icon>
              <n-icon v-if="option.icon_id === 1" size="50">
                <lockIcon/>
              </n-icon>
              <n-icon v-if="option.icon_id === 2" size="50">
                <mgrIcon/>
              </n-icon>
              <n-icon v-if="option.icon_id === 3" size="50">
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

    <!--        </n-scrollbar>-->
  </div>
</template>

<style scoped lang="less">
.root {
  height: 100vh;
}

.lay-1 {
  position: relative;
  display: flex;
  flex-direction: column;
  height: 100vh; // 修改为占满视窗高度
  background-color: #c8d9eb;
  //overflow: hidden;

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
      margin-top: 10px;
      position: relative;
      max-width: 600px;
      font-size: 1.2rem;
      opacity: 0.8;
    }

    .btn-start {
      margin-top: 30px;
      width: 50%;
      height: 50px;
      border-radius: 10px;
      background-color: #1d324d;
      color: white;
      box-shadow: rgba(29, 50, 77, 0.2) 10px 10px 50px 0;
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
      margin-top: 20px;
      font-size: 2rem;
      font-weight: 600;
      opacity: 0.9;
    }

    .why-us-sub {
      width: auto;
      font-size: 1rem;
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
    font-size: 1.5rem
  }

  .desc2 {
    font-size: 1.2rem;
    opacity: 0.7;
  }
}


</style>
