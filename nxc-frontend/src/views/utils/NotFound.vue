<script setup lang="ts">
defineOptions({
  name: 'NotFound',
})

import {onBeforeMount, onMounted, ref} from "vue";
import {useRouter} from "vue-router";
import useUserInfoStore from "@/stores/useUserInfoStore";
import {useI18n} from "vue-i18n";
import PageHead from "@/views/utils/PageHead.vue";

const {t} = useI18n();
const i18nPrefix: string = 'notFound';
const userInfoStore = useUserInfoStore();
const router = useRouter();

let secLeft = ref<number>(1000);
let timer: number | undefined = undefined

const callBack = () => {
  if (userInfoStore.isAuthed) {
    if (userInfoStore.thisUser.isAdmin) return router.replace({path: '/admin/dashboard/summary'});
    else router.replace({path: '/dashboard/summary'});
  } else {
    return router.replace({path: '/'});
  }
};

// 启动定时器
onMounted(() => {
  timer = setInterval(() => {
    if (secLeft.value > 1) {
      secLeft.value--;
    } else {
      clearInterval(timer); // 停止定时器
      callBack(); // 执行回调
    }
  }, 1000);
});

onBeforeMount(() => {
  // 你可以在这里进行一些初始化操作，当前没有需要的操作。
  timer?clearInterval(timer):null
});

</script>

<template>
  <div class="root">
    <n-layout-content style="height: 100vh">
      <div class="inner">
        <PageHead
            :title="t(`${i18nPrefix}.title`)"
            :description="t(`${i18nPrefix}.description`)"
        >
         <div>
           <p style="opacity: 0.8">{{ t(`${i18nPrefix}.p1`, { sec: secLeft }) }}</p>
           <n-button
               type="primary"
               text
               style="text-decoration: underline;"
               @click="callBack"
           >
             {{ t(`${i18nPrefix}.manualBack`) }}
           </n-button>
         </div>
        </PageHead>
      </div>
    </n-layout-content>
  </div>
</template>

<style scoped lang="less">
.root {
  height: 100vh;
  width: 100%;
}

.inner {
  padding-left: 20px;
  height: calc(100vh - 30%);
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;

  .title {
    font-size: 4rem;
  }
}
</style>