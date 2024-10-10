<script setup lang="ts">
import {ref, h} from "vue"
</script>

<script lang="ts">
export default {
  name: "OrderMgr",
}
</script>

<template>
<div class="root">
  <n-card :embedded="trur" hoverable :bordered="false" :title="'订单管理'">

  </n-card>
</div>
</template>

<style scoped lang="less">
.root {
   padding: 20px;

}
</style>