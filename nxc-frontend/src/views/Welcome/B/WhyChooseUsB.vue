<script setup lang="ts">
import {useI18n} from "vue-i18n";

const {t} = useI18n()
</script>

<script lang="ts">
export default {
  name: 'WhyChooseUsA',
}
</script>


<template>
  <div class="root">
    <div class="l-content">
      <div class="l-content-top"></div>
      <div class="l-content-color1"></div>
      <div class="l-content-color2"></div>
      <div class="l-content-color3"></div>
      <div class="l-container">
        <p class="title">加入我們 Join us.</p>
        <p class="sub">If you have no account.</p>
      </div>
    </div>
    <div class="r-content"></div>
  </div>
</template>

<style scoped lang="less">
.root {
  height: 100vh;
  display: flex;
  position: relative;
}

.l-content {
  flex: 1;
  position: relative;
  overflow: hidden;
  //background-color: #292929;

  .l-container {
    position: absolute;
    z-index: 1001;
    top: 40%;
    left: 15%;
    //background-color: #4cae4c;
    width: 70%;

    .title {
      font-size: 2.5rem;
      font-weight: 700;
      opacity: 0.8;
      color: #1d3144;
    }

    .sub {
      margin-top: 20px;
      font-size: 1.2rem;
      opacity: 0.5;
    }
  }

  .l-content-top {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    backdrop-filter: blur(100px);
    z-index: 1000;
  }

  .l-content-color1 {
    width: 500px;
    height: 500px;
    border-radius: 50%;
    background-color: rgba(225,107,140, 0.5);
    position: absolute;
    //top: calc(70% - 400px / 2);
    //left: calc(30% - 400px / 2);
    //bottom: calc(100% - 100px);
    left: -250px;
    bottom: -250px;
    z-index: 1;
  }

  .l-content-color2 {
    width: 300px;
    height: 300px;
    border-radius: 50%;
    //background-color: rgba(84, 234, 255, 0);
    background: linear-gradient(to right, rgb(127, 127, 213), rgb(134, 168, 231), rgb(145, 234, 228));
    position: absolute;
    top: calc(40% - 400px / 2);
    left: calc(45% - 400px / 2);
    z-index: 2;
  }

  .l-content-color3 {
    width: 300px;
    height: 300px;
    border-radius: 50%;
    //background-color: rgba(220, 135, 227, 0);
    background: linear-gradient(to right, rgb(194, 229, 156), rgb(100, 179, 244));
    position: absolute;
    top: calc(50% - 400px / 2);
    left: calc(60% - 400px / 2);
    z-index: 3;
  }
}

.r-content {
  flex: 1;
  //background-color: #c8d9eb;

  //background-color: #e3e5e7;

}
</style>