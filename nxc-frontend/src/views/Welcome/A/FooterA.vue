<script setup lang="ts">

</script>

<script lang="ts">
export default {
  name: 'FooterA',
}
</script>

<template>
<div class="root" style="color: white">
  <p class="title">快速連結</p>
  <n-hr style="margin-bottom: 30px"></n-hr>

  <div class="subscribe-us">
    <p class="subscribe-us-title">關注我們</p>
    <div class="subscribe-us-media">
      <div class="sub-item">BiliBili1</div>
      <div class="sub-item">BiliBili2</div>
      <div class="sub-item">BiliBili3</div>
      <div class="sub-item">BiliBili4</div>
    </div>
  </div>
  <n-hr style="margin-bottom: 30px"></n-hr>


  <p>©2024 HP Development Company, L.P. The information contained herein is subject to change without notice.</p>
</div>
</template>

<style scoped lang="less">
.root {
  height: 300px;
  background-color: #0e1b42;
  padding: 30px;
  .title {
    font-size: 1.2rem;
    font-weight: 700;
    opacity: 0.8;
    margin-top: 30px;
  }
  .subscribe-us {
    .subscribe-us-title {
      font-size: 1.2rem;
      font-weight: 300;

    }
    .subscribe-us-media {
      display: flex;
      flex-direction: row;
      .sub-item {
        margin-right: 20px;
      }
    }
  }
}
</style>