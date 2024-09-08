<script setup lang="ts" name="CommonLogo">
import {defineComponent} from 'vue'
import {useRouter} from 'vue-router'
import useThemeStore from '@/stores/useThemeStore'
import theme from "@/stores/useThemeStore";

const router = useRouter()
const themeStore = useThemeStore()

let backDashBoard = () => {
  console.log('回到主页')
  router.push({
    path: '/dashboard/summary'
  })

}

</script>

<template>
<div class="root" @click="backDashBoard">
  <div class="txt-root">
    <button class="btn">
      <p class="txt">Nxc Networks</p>
    </button>
  </div>
</div>
</template>

<style scoped>
.root {
  text-align: center;
  justify-content: center;
  line-height: 52px;
  background-color: v-bind('themeStore.getTheme.topLogoBgColor');
  .txt-root {
    .btn {
      background-color: rgba(255, 255, 255, 0.0);
      border: none;
      line-height: 30px;
      .txt {
        position: relative;
        font-size: 1.25rem;
        font-weight: 100;
        color: v-bind('themeStore.getTheme.topLogoTextColor')
      }
      .txt::after {
        content: '';
        position: absolute;
        width: 100%;
        transform: scaleX(0);
        height: 2px;
        bottom: 0;
        left: 0;
        background-color: #7995c6;
        transform-origin: bottom right;
        transition: transform 0.25s ease-out;
      }
      .txt:hover::after {
        transform: scaleX(1);
        transform-origin: bottom left;
      }
    }
  }
}
</style>