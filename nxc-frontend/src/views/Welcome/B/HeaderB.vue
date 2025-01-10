<script setup lang="ts">
import {useI18n} from "vue-i18n";
import useApiAddrStore from "@/stores/useApiAddrStore";
import {useRouter} from "vue-router";

const {t} = useI18n();
const router = useRouter();
const apiAddrStore = useApiAddrStore();

</script>

<script lang="ts">
export default {
  name: 'HeaderB',
}
</script>

<template>
  <div class="root">
    <div class="left-content">
      <span>
          <img src="https://ikanned.com:24444/d/Upload/NXC/links3.svg" alt="SVG Image">
      </span>
      <p class="name">Nxc Networks</p>
      <n-button size="large" style="margin-right: 40px" color="#277aa8" text type="primary">关于我们</n-button>
      <n-button size="large" color="#277aa8" text type="primary">定价</n-button>
    </div>
    <div class="right-content">
      <n-button
          secondary
          color="#277aa8"
          style="margin-right: 20px"
          size="large"
          type="primary"
          @click="router.push({path: '/login'})"
      >登录
      </n-button>
      <n-button
          secondary
          color="#277aa8"
          size="large"
          type="primary"
          @click="router.push({path: '/register'})"
      >注册新账号
      </n-button>
    </div>
  </div>
</template>

<style scoped lang="less">
.root {
  height: 70px;
  line-height: 70px;
  background-color: rgba(200, 217, 235, 0.8);
  backdrop-filter: blur(10px);
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
  padding: 0 20px;
  color: #277aa8;

  .left-content {
    display: flex;
    flex-direction: row;
    align-items: center;
    // 新增以下样式

    span {
      display: flex;
      align-items: center;
      margin-right: 10px;

      img {
        width: 30px;

      }
    }

    .name {
      font-size: 1.5rem;
      margin-right: 100px;
      color: #277aa8;
    }
  }

  .right-content {
    display: flex;
    flex-direction: row;
    align-items: center;
  }
}
</style>