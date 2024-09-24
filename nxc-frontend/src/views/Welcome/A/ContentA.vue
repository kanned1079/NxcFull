<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {ref, onMounted, onBeforeMount} from "vue";
import TreeFront from "./assets/TreeFront.svg"
import TreeMedium from "./assets/TreeMedium.svg"
import Sun from "./assets/Sun.svg"


// 创建一个 ref 来获取 lay-2 元素
const lay1Ref = ref<HTMLElement | null>(null);

// 监听滚动事件，自动触发滚动

window.addEventListener('scroll', function (event){
  console.log("scroll", event);
  console.log(document.documentElement.scrollTop);
});
</script>

<script lang="ts">
export default {
  name: 'ContentA',
}
</script>

<template>
  <div class="root" >
<!--    <n-scrollbar style="height: 100vh"  >-->
      <div class="lay-1" ref="lay1Ref">
        <div class="image-wrapper img1"><img :src="TreeFront" alt="Image 1"/></div>
        <div class="image-wrapper img2"><img :src="TreeMedium" alt="Image 2"/></div>
        <div class="image-wrapper img3"><img :src="Sun" alt="Image 3"/></div>
        <!-- 引入悬浮左侧的介绍内容 -->
        <div class="intro">
          <div class="web-name">
            <p class="title" id="appdemo">欢迎来到</p>
            <p class="app-name">{{ 'app_name' }}</p>
          </div>
          <p class="desc">
            “穿过县境上长长的隧道，便是雪国。夜空下，大地一片莹白，火车在信号所前停下来。”在这里川端康成用几近吝啬的简洁文字，拉开了《雪国》的序幕。</p>
          <n-button size="large" class="btn-start" text>开始使用</n-button>
        </div>

      </div>

    <!--       此处有空隙-->

      <div class="lay-2" ></div>

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
  height: 100vh;
  background-color: #c8d9eb;
}


</style>
