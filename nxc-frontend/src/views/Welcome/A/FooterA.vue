<script setup lang="ts">
import {useI18n} from "vue-i18n";
import youtube from "../icons/youtube.svg"
import bilibili from "../icons/bilibili.svg"
import instagram from "../icons/instagram.svg"
import wechat from "../icons/wechat.svg"
import {reactive} from "vue";

const {t} = useI18n();
const props = defineProps<{
  bilibili_official_link: string,
  youtube_official_link: string,
  instagram_link: string,
  filing_number: string,
  page_suffix: string,
}>()

let contactMethods = reactive([
  {
    id: 0,
    name: 'Bilibili',
    link: props.bilibili_official_link || '',
  },
  {
    id: 1,
    name: 'YouTube',
    link: props.youtube_official_link || '',
  },
  {
    id: 2,
    name: 'Instagram',
    link: props.instagram_link || '',
  },
  // {
  //   id: 3,
  //   name: 'WechatOfficial',
  // },
])

let clickToLink = (link: string) => {
  // 新开启一个窗口并加载指定的链接
  console.log(link)
  window.open(link, '_blank'); // '_blank' 表示在新标签页或窗口中打开链接
}

</script>

<script lang="ts">
export default {
  name: 'FooterA',
}
</script>

<template>
  <div class="root" style="color: white">
    <p class="title">{{ t('welcome.A.fastLink') }}</p>
    <n-hr style="margin-bottom: 30px"></n-hr>

    <div class="subscribe-us">
      <p class="subscribe-us-title">{{ t('welcome.A.subscribeUs') }}</p>
      <div class="subscribe-us-media">
        <n-button
            text
            color="white"
            class="sub-item"
            v-for="(item, index) in contactMethods"
            :key="item.id"
            @click="clickToLink(item.link)"
        >

          <n-icon v-if="item.name === 'Bilibili'" style="margin-right: 5px" size="20">
            <bilibili/>
          </n-icon>
          <n-icon v-if="item.name === 'YouTube'" style="margin-right: 5px" size="20">
            <youtube/>
          </n-icon>
          <n-icon v-if="item.name === 'Instagram'" style="margin-right: 5px" size="20">
            <instagram/>
          </n-icon>
          <n-icon v-if="item.name === 'WechatOfficial'" style="margin-right: 5px" size="25">
            <wechat/>
          </n-icon>
          {{ item.name }}
        </n-button>
      </div>
    </div>
    <n-hr style="margin-bottom: 30px"></n-hr>
    <!--    <p class="code">{{ t('welcome.A.filingsCode', {code: '11223'}) }}</p>-->
    <p class="code">{{ props.filing_number }}</p>
    <p class="right-policy">{{ props.page_suffix }}</p>
  </div>
</template>

<style scoped lang="less">
.root {
  height: 240px;
  background-color: #0e1b42;
  padding: 20px 20px 70px 20px;

  .title {
    font-size: 1.5rem;
    font-weight: 300;
    opacity: 18;
    margin-top: 30px;
  }

  .subscribe-us {
    .subscribe-us-title {
      margin-bottom: 10px;
      font-size: 1rem;
      opacity: 0.8;
      font-weight: 400;

    }

    .subscribe-us-media {
      display: flex;
      flex-direction: row;
      margin-top: 10px;

      .sub-item {
        margin-right: 20px;
        font-size: 1rem;
        transition: ease 300ms;
      }

      .sub-item:hover {
        opacity: 0.5;
      }
    }
  }

  .code {
    font-size: 1rem;
    opacity: 0.8;
  }

  .right-policy {
    margin-top: 10px;
    font-size: 1rem;
    opacity: 0.5;
  }
}
</style>