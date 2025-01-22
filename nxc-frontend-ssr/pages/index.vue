<script setup lang="ts">
 const btnClickFunc = () => {
   navigateTo('/admin/dash1')
 }
</script>

<template>
<div>
  index page
  <n-button
    type="success"
    tertiary
    @click="btnClickFunc"
  >To</n-button>
</div>
</template>

<style scoped lang="less">

</style>