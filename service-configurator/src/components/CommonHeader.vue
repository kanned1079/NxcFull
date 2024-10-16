<script lang="ts" setup>
import yamlIcon from "@/assets/yaml.svg"
import useThemeStore from "@/store/useThemeStore";

const themeStore = useThemeStore();
</script>

<script lang="ts">
export default {
  name: 'CommonHeader',
}
</script>

<template>
  <div class="root">
    <div class="header-content" :style="!themeStore.darkThemeEnabled?{backgroundColor: '#69b0ac'}:{backgroundColor: '#577c8a'}">

      <span class="name">
        <n-icon size="20" style="margin-right: 10px"><yamlIcon/></n-icon>
        微服务配置
      </span>
    </div>
  </div>

</template>

<style lang="less" scoped>
.root {
  .header-content {
    width: 100%;
    height: 50px;
    line-height: 50px;
    //background-color: #66afe9;
    display: flex;
    align-items: center;
    justify-content: left;
    .name {
      font-size: 1rem;
      color: white;
      margin-left: 30px;
      display: flex;
      flex-direction: row;
      align-items: center;
      justify-content: center;
    }
  }
}
</style>