<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {onBeforeUnmount, onMounted, ref} from "vue";

const {t} = useI18n();

let iconSize = ref<number>(40)

const rootRef = ref<HTMLElement | null>(null);

// let page = rootRef.value
// let scrollTop = page.getBoundingClientRect().top * -1
// console.log(scrollTop)

// 创建一个 ref 来获取 lay-2 元素
const lay2Ref = ref<HTMLElement | null>(null);

// 创建一个 ref 来获取 lay-1 元素
const lay1Ref = ref<HTMLElement | null>(null);

let advanceOptions = ref([
  {
    id: 0,
    icon_id: 0,
    title: "瀏覽安全",
    content: "優秀的防火牆過濾系統能有效防禦網路釣魚和惡意網站",
  },
  {
    id: 1,
    icon_id: 1,
    title: "端到端加密",
    content: "雙向 SSL 和端對端加密保護您的隱私安全，即使是服務器也無法讀取您的信息",
  },
  {
    id: 2,
    icon_id: 2,
    title: "高效管理",
    content: "一個用戶介面管理所有密鑰，管理功能完善豐富，無須擔心訂閱洩露問題",
  },
  {
    id: 3,
    icon_id: 3,
    title: "方便快捷",
    content: "提供完整的 API 文檔供 WebApp 或是嵌入到第三方軟件中",
  },
]);

let isScrolledToTarget = false;
let lastScrollTop = 0;

let onScrollFunction = () => {
  const lay1Element = lay1Ref.value;
  if (lay1Element) {
    const scrollTop = lay1Element.scrollTop;
    console.log(scrollTop);
    const isScrollingDown = scrollTop > lastScrollTop;
    lastScrollTop = scrollTop;
    if (!isScrolledToTarget && isScrollingDown && lay2Ref.value) {
      lay2Ref.value.scrollIntoView({behavior: "smooth"});
      isScrolledToTarget = true;
    } else if (isScrolledToTarget && !isScrollingDown) {
      // 添加反向滚动的处理逻辑，例如重置状态变量或执行其他操作
      isScrolledToTarget = false;
    }
  }
};


onMounted(() => {

});

onBeforeUnmount(() => {

});
</script>

<script lang="ts">
export default {
  name: "ContentB",
};
</script>

<template>
  <div class="root">

  </div>
</template>

<style scoped lang="less">
.root {

}
</style>
