<script lang="ts" setup>
</script>

<script lang="ts">
export default {
  name: 'CommonContent'
}
</script>

<template>
  <div class="root">
    <n-card :embedded="true" :bordered="false" hoverable class="content-card" title="Redis配置"></n-card>
  </div>

</template>

<style lang="less" scoped>
.root {
  margin: 20px 20px 20px 0;
  .content-card {
    height: calc(100vh - 110px);
    background-color: #e3e5e7;
  }

}
</style>