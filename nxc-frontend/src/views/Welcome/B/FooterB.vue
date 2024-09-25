<script setup lang="ts">
import {useI18n} from "vue-i18n";
import youtube from "../icons/youtube.svg"
import bilibili from "../icons/bilibili.svg"
import instagram from "../icons/instagram.svg"
import wechat from "../icons/wechat.svg"
import {ref, reactive} from "vue";

const {t} = useI18n();

let contactMethods = reactive([
  {
    id: 0,
    name: 'Bilibili',
  },
  {
    id: 1,
    name: 'YouTube',
  },
  {
    id: 2,
    name: 'Instagram',
  },
  {
    id: 3,
    name: 'WechatOfficial',
  },
])

</script>

<script lang="ts">
export default {
  name: 'FooterB',
}
</script>

<template>
<div class="root" style="color: white">
  <p class="title">快速連結</p>
  <n-hr style="margin-bottom: 30px"></n-hr>

  <div class="subscribe-us">
    <p class="subscribe-us-title">關注我們</p>
    <div class="subscribe-us-media">
      <n-button
          text
          color="white"
          class="sub-item"
          v-for="item in contactMethods"
          :key="item.id"
      >
        <n-icon v-if="item.name === 'Bilibili'" style="margin-right: 5px" size="25"><bilibili /></n-icon>
        <n-icon v-if="item.name === 'YouTube'" style="margin-right: 5px" size="25"><youtube /></n-icon>
        <n-icon v-if="item.name === 'Instagram'" style="margin-right: 5px" size="25"><instagram /></n-icon>
        <n-icon v-if="item.name === 'WechatOfficial'" style="margin-right: 5px" size="25"><wechat /></n-icon>
        {{ item.name }}
      </n-button>

    </div>
  </div>
  <n-hr style="margin-bottom: 30px"></n-hr>
  <p class="code">備案號  12345677</p>
  <p class="right-policy">Copyright © 2013 - 2024 Tencent Cloud. All Rights Reserved. NxcNetworks 版權所有</p>
</div>
</template>

<style scoped lang="less">
.root {
  height: 240px;
  background-color: #0e1b42;
  padding: 20px 20px 70px 20px;
  .title {
    font-size: 1.2rem;
    font-weight: 700;
    opacity: 0.8;
    margin-top: 30px;
  }
  .subscribe-us {
    .subscribe-us-title {
      margin-bottom: 10px;
      font-size: 1.25rem;
      font-weight: 400;

    }
    .subscribe-us-media {
      display: flex;
      flex-direction: row;
      .sub-item {
        margin-right: 20px;
        font-size: 1.1rem;
      }
    }
  }
  .code {
      font-size: 1rem;
      opacity: 0.8;
    margin-bottom: 5px;
  }
  .right-policy {
    font-size: 1rem;
    opacity: 0.5;
  }
}
</style>