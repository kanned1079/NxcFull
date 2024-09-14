<script setup lang="ts" name="UserDocument">
import {ref, reactive, onMounted, computed} from "vue";
import {useRouter} from "vue-router";
import useThemeStore from "@/stores/useThemeStore";
import MarkdownIt from 'markdown-it';

import instance from "@/axios";

const renderedMarkdown = (mdText: string) => {
  const md = new MarkdownIt();
  return md.render(mdText);
}

let text = ref('')

const themeStore = useThemeStore();
const router = useRouter();

let showData = [
  {
    category: '使用方式',
    data: [
      {
        title: '配置旁路由实现科学上网',
        body: `#### 什么是审计？

>审计，指封锁特定的关键词、域名的方式，让用户无法接触到互联网的某一部分。

#### NxcNetworks 是否会进行审计？

>是的，NxcNetworks 会进行审计。<br>我们尽力保护每一位用户的自由，但部分的审计是维持我们的服务稳定运作的必要条件。<br>NxcNetworks 审计（不允许访问）的内容仅包含：<br>  **1. 种子/磁力链/BT/PT/迅雷下载（正常的文件下载不在此类）<br> 2. 极端宗教组织，极端政治派别势力宣传网站（容易导致我们的服务被追查）**

#### 地区特例
> 为了配合台湾警方防止电信诈骗及洗钱活动，台湾地区的线路屏蔽了一些用于进行电信诈骗洗钱的网络游戏和用于行骗的社交媒体。<br>我们不保留这些网站的访问记录。`
      },
      {

      }
    ]
  }
]


let getAllDocuments = async () => {
  console.log('拉取所有文档列表')
  let {data} = await instance.get('')
  if (data.code === 200) {
    console.log(data)
  }
}

onMounted(() => {
  console.log('document挂载');
  themeStore.menuSelected = 'user-doc'
  themeStore.userPath = '/dashboard/document'
  // getAllDocuments()




})


</script>

<template>
<div class="root">
  <n-card hoverable :embedded="true" content-style="padding: 0;" class="search-root">
<!--    <n-input size="large" class="search-input" placeholder="请输入要搜索的内容（模糊查询）"></n-input>-->
<!--    <n-button :bordered="false" type="primary" size="large" class="search-btn">搜索</n-button>-->

    <n-input-group>
      <n-input :bordered="false" size="large" class="search-input" placeholder="请输入要搜索的内容（模糊查询）"></n-input>
      <n-button strong :bordered="false" type="primary" size="large" class="search-btn">搜索</n-button>
    </n-input-group>

  </n-card>

  <n-card class="card" hoverable :embedded="true" title="使用">
    <n-input type="textarea" :rows="10" v-model:value="text"></n-input>
    <div v-html="renderedMarkdown(text)"></div>
  </n-card>
</div>
</template>

<style scoped lang="less">
.root {
  padding: 20px;
  .search-root {
    display: flex;
    flex-direction: row;
    align-content: space-between;
    margin-bottom: 20px;
    .search-input {
      line-height: 50px;
      border: 10px;
    }
    .search-btn {
      height: 50px;
      width: 100px;
      margin-left: 20px;
    }
  }
}

.n-card {
  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
  border: 0;
}

</style>