<script setup lang="ts">
definePageMeta({
  path: '/admin/dash1'
})

if (import.meta.server) {
  console.log('server')
} else {
  console.log('client')
}
</script>

<script lang="ts">
export default {
  name: "Dashboard",
}
</script>

<template>
dash
</template>

<style scoped lang="less">

</style>