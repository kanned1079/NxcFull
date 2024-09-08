<script setup lang="ts" name="CouponMgr">
import {onMounted} from "vue";

onMounted(() => {
  console.log('couponMgr挂载')

})

</script>

<template>
<div class="root">
  <n-card title="优惠券管理" hoverable :embedded="true" class="card">

  </n-card>
</div>
</template>

<style lang="less" scoped>
.root {
  .card {
    border: 0;
  }
}
</style>