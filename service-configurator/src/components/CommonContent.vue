<script lang="ts" setup>
import {RouterView} from "vue-router";
</script>

<script lang="ts">
export default {
  name: 'CommonContent'
}
</script>

<template>
<!--    <n-card :embedded="true" :bordered="false" hoverable class="content-card" title="Redis配置"></n-card>-->
  <RouterView></RouterView>


</template>

<style lang="less" scoped>

</style>