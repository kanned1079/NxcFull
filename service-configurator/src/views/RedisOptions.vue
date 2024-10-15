<script setup lang="ts">

</script>

<script lang="ts">
export default {
  name: 'RedisOptions',
}
</script>

<template>
  <div class="root">
    <n-card :embedded="true" :bordered="false" hoverable title="Redis配置"></n-card>
  </div>
</template>

<style scoped lang="less">
.root {
  margin: 20px;
}
</style>