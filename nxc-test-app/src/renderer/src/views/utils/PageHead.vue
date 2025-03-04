<script setup lang="ts">
import {type Component} from "vue"
const props = defineProps<{
  title: string;
  description: string;
  icon?: Component;
  secondary?: boolean;
}>()
</script>

<template>
  <div class="root">
    <div class="t-part">
      <n-icon class="page-icon" size="20" v-if="props.icon" :component="props.icon?props.icon:null"/>
      <p class="title" :style="secondary?{fontSize: '1.1rem !important'}:null">{{ props.title }}</p>
    </div>
    <div class="b-part">
      <p class="description">{{ props.description }}</p>
    </div>
    <div class="slot-elements">

      <slot></slot>
    </div>
  </div>
</template>

<style scoped lang="less">
.root {
  margin-top: 20px;
  .t-part {
    display: flex;
    flex-direction: row;
    align-items: center;
    .page-icon {
      margin-right: 4px;
    }
    /*
    .back-btn {
      margin-right: 6px;
    }
     */
    .title {
      font-size: 1.4rem;
      font-weight: bold;
      opacity: 0.9;
    }
  }

  .b-part {
    .description {
      font-size: 0.9rem;
      opacity: 0.8;
    }
  }
  .slot-elements {
    display: flex;
    flex-direction: row;
    align-items: center;
    margin-top: 10px;
    /*
    .operate-prefix {
      margin-right: 10px;
    }

     */
  }
}
</style>
