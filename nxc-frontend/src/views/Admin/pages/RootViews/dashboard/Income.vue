<script setup lang="ts" name="Income">
import {
  Pricetags as incomeIcon,
} from '@vicons/ionicons5'
</script>

<template>
  <n-card hoverable class="root" content-style="padding: 0" :embedded="true" :bordered="false">
    <div class="layer-up">
      <p class="txt">今日收入</p>
      <n-icon size="40" style="opacity: 10%; margin: 10px 10px 0 0 ">
        <incomeIcon/>
      </n-icon>
    </div>
    <div class="layer-bottom">

      <p class="fund">634563.23</p>
      <p class="sign">¥</p>
    </div>
  </n-card>
</template>

<style scoped>
.root {
  padding: 0;
  margin: 0 20px 20px 20px;
  display: flex;
  flex-direction: column;
  .layer-up {
    display: flex;
    justify-content: space-between;

    .txt {
      font-size: 16px;
      margin-left: 30px;
      margin-top: 10px;
    }
  }

  .layer-bottom {
    display: flex;
    margin: 10px 10px 10px 10px;
    position: relative;
    .fund {
      margin-left: 20px;
      font-size: 45px;
    }
    .sign {
      font-size: 24px;
      text-align: start;
    }
  }
}
</style>