<script setup lang="ts">
import {onMounted, toRaw, onBeforeUnmount, onBeforeMount, watch, computed} from 'vue';
import {RouterView} from 'vue-router';
import useThemeStore from "@/stores/useThemeStore";
import useAppInfosStore from "@/stores/useAppInfosStore";
import {type NConfigProvider} from 'naive-ui'

const themeStore = useThemeStore();
const appInfosStore = useAppInfosStore()

const handleResize = () => {
  // console.log('宽度小于768', window.innerWidth < 768)
  themeStore.menuCollapsed = window.innerWidth < 768;

};

let isDark = computed(() => window.matchMedia('(prefers-color-scheme: dark)').matches)

onMounted(() => {
  handleResize(); // 初始时根据当前宽度设置状
  window.addEventListener('resize', handleResize); // 动态监听窗口大小
  console.log('app挂载');
  // 处理深色主题设置
  themeStore.readEnableDarkMode()
  // 从总设置中获取主题并应用
  themeStore.setThemeFromSetting()
  // 监听深色模式的变化
  console.log(isDark.value)
  themeStore.enableDarkMode = isDark.value

});

watch(isDark, (val: boolean) => {
  themeStore.enableDarkMode = val
})



onBeforeMount(() => {
  // appInfosStore.getCommonConfig()
  themeStore.setThemeFromSetting()
})

// onMounted(() => {
//   themeStore.setThemeFromSetting()
// })

onBeforeUnmount(() => {
  window.removeEventListener('resize', handleResize); // 组件卸载时移除监听器

});


</script>

<template>

<!--<div style="scrollbar-width: none">-->
  <n-config-provider
      :theme="themeStore.getMainTheme"
      :theme-overrides="toRaw(themeStore.getTheme.selfOverride)"
  >

    <!--    <n-theme-editor>-->
    <n-message-provider>
      <n-notification-provider>
        <n-dialog-provider>
          <RouterView></RouterView>
        </n-dialog-provider>
      </n-notification-provider>
    </n-message-provider>
    <!--    </n-theme-editor>-->


  </n-config-provider>
<!--</div>-->


</template>


<style lang="less">

.slide-fade-enter-active {
  transition: all 250ms ease;
}

.slide-fade-leave-active {
  transition: all 250ms ease;
}

.slide-fade-enter-from {
  transform: translateY(10px);
  opacity: 0;
}
.slide-fade-leave-to {
  //transform: translateY(10px);
  opacity: 0;
}

* {
  padding: 0;
  margin: 0;
  scrollbar-width: none;
  -ms-overflow-style: none;
}

*::-webkit-scrollbar {
  display: none;
}
</style>
