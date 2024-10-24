<script setup lang="ts">
import {ref, h} from "vue"

</script>

<script lang="ts">
export default {
  name: 'PaymentResult'
}
</script>

<template>
<div class="root">
  <n-card :embedded="true" :bordered="false" hoverable title="支付结果">

  </n-card>

</div>
</template>

<style scoped lang="less">
.root {
  margin: 20px;
}
</style>