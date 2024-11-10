<script setup lang="ts">
</script>

<script lang="ts">
export default {
  name: 'LoadingBefore',
}
</script>

<template>
  <div class="root">
    <n-card class="loading-notice" hoverable :embedded="true" :bordered="false">
      <div class="loading-notice-head">
        <n-skeleton style="margin-bottom: 20px" height="30px" width="30%" :sharp="false"/>
        <n-skeleton style="margin-bottom: 0" width="40%" :sharp="false" text/>
      </div>
    </n-card>
    <n-card class="loading-plans" hoverable :embedded="true" :bordered="false">
      <div class="loading-plans-head">
        <n-skeleton style="margin-bottom: 20px" height="30px" width="20%" :sharp="false"/>
      </div>
      <div v-for="(i, k) in 3" :key="k" class="loading-plans-body">
        <n-skeleton style="margin-bottom: 10px" height="20px" width="100%" :sharp="false"/>
        <n-skeleton style="margin-bottom: 20px" height="12px" width="33%" :sharp="false"/>
        <n-hr v-if="i !== 3"></n-hr>
      </div>
    </n-card>
  </div>
</template>

<style scoped lang="less">
.root {
  padding: 20px;
  opacity: 0.5;

  .loading-notice {
    height: 200px;
    position: relative;
    margin-bottom: 20px;

    .loading-notice-head {
      height: 100%;
      display: flex;
      flex-direction: column;
      justify-content: end;

    }
  }

  .loading-plans {
    height: 100%;
    display: flex;

    .loading-plans-head {
      margin-bottom: 40px;
    }

  }
}
</style>