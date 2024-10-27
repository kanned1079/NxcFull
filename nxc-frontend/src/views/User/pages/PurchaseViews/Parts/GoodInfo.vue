<script setup lang="ts">
import {defineProps} from "vue"
const props = defineProps(['goodInfoData'])
</script>

<template>
  <n-card
      class="good-info"
      :embedded="true"
      hoverable
      content-style="padding: 0"
      :bordered="false"
  >
    <n-p class="title">商品信息</n-p>
    <div class="good-info-item" v-for="item in props.goodInfoData" :key="item.title">
      <p class="item-head">{{ item.title }}</p>
      <p class="item-content">{{ item.content }}</p>
    </div>
  </n-card>
</template>

<style scoped lang="less">
.good-info {
  padding-bottom: 15px;
  margin-bottom: 20px;

  .title {
    font-size: 1.1rem;
    font-weight: 400;
    background-color: rgba(216, 216, 216, 0.1);
    padding: 10px 20px 15px 20px;
    border-radius: 3px 3px 0 0;
  }

  .good-info-item {
    display: flex;
    flex-direction: row;
    margin: 0 0 10px 20px;
    //background-color: #66afe9;

    .item-head {
      font-size: 0.9rem;
      opacity: 0.8;
      font-weight: 400;
      min-width: 100px;
      //background-color: #66afe9;
    }

    .item-content {
      font-size: 0.9rem;
      margin-left: 120px;
      opacity: 1;
      font-weight: 400;
    }

  }
}
</style>