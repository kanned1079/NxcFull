<script lang="ts" setup>
</script>

<script lang="ts">
export default {
  name: 'CommonHeader',
}
</script>

<template>
  <div class="root">
    <n-card :embedded="true" :bordered="false" hoverable class="header-content" content-style="padding: 0px" title="微服务配置">
      <p class="name"></p>
    </n-card>
  </div>

</template>

<style lang="less" scoped>
.root {
  margin: 20px 20px 0 20px;
  .header-content {
    width: 100%;
    background-color: #e3e5e7;
    .name {
      font-size: 1.2rem;
      padding-left: 30px;

    }
  }
}
</style>