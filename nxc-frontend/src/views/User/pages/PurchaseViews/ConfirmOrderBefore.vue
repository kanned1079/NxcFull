<script setup lang="ts">
import {onBeforeMount, onMounted, ref} from "vue"

let animated = ref<boolean>(false)
let showSpin = ref<boolean>(false)


onMounted(() => {
  animated.value = true
  showSpin.value = true
})

onBeforeMount(() => {
  animated.value = false
  showSpin.value = false
})


</script>

<script lang="ts">
export default {
  name: 'ConfirmOrderBefore'
}
</script>


<template>
  <transition name="slide-fade">
    <div v-if="animated" class="root">
      <n-spin :show="showSpin">
        <n-card class="load-body" hoverable :embedded="true" :bordered="false">
        </n-card>

        <template #description>
          我们正在处理您的订单，请稍后。
        </template>
      </n-spin>

    </div>
  </transition>
</template>

<style scoped lang="less">
.root {
  margin-top: 72px;
  padding: 0;

}

.load-body {
  height: 400px;
}
</style>