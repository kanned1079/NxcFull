<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {ref, reactive, onMounted, computed} from "vue";
import {useRouter} from "vue-router";
import useThemeStore from "@/stores/useThemeStore";
import useApiAddrStore from "@/stores/useApiAddrStore";
import type { DrawerPlacement, useMessage } from 'naive-ui'
const placement = ref<DrawerPlacement>('right')
// import MarkdownIt from 'markdown-it';
import {convertMd} from "@/utils/markdown";
import instance from "@/axios";
import {MdPreview} from "md-editor-v3";
import 'md-editor-v3/lib/style.css';

const {t} = useI18n()
const message = useMessage()
const apiAddrStore = useApiAddrStore();


let search = ref('')
let text = ref('')
let active = ref<boolean>(false)

const activate = (place: DrawerPlacement, title:string, body: string) => {
  drawTitle.value = title
  drawBody.value = body
  active.value = true
  placement.value = place
}

const themeStore = useThemeStore();
const router = useRouter();

let drawTitle = ref<string>('')
let drawBody = ref<string>('')

let doc_list = reactive([])

let getAllDocuments = async () => {
  console.log('拉取所有文档列表')
  try {
    let {data} = await instance.get(apiAddrStore.apiAddr.user.getAllDocumentList, {
      params: {
        lang: 'zh_CN',
        find: search.value
      }
    })
    if (data.code === 200) {
      console.log(data)
      Object.assign(doc_list, data.doc_list)
    } else {
      message.error(data.msg)
    }
  }catch (error) {
    message.error(error)
  }

}

let showDetail = () => {

}

onMounted(() => {
  console.log('document挂载');
  themeStore.menuSelected = 'user-doc'
  themeStore.userPath = '/dashboard/document'
  getAllDocuments()
  // getAllDocuments()

})


</script>

<script lang="ts">
export default {
  name: "UserDocument",
}
</script>

<template>
<div class="root">
  <n-card hoverable :embedded="true" content-style="padding: 0;" class="search-root">
    <n-input-group>
      <n-input v-model="search" :bordered="false" size="large" class="search-input" :placeholder="t('userDocument.searchPlaceholder')"></n-input>
      <n-button @click="getAllDocuments" strong :bordered="false" type="primary" size="large" class="search-btn">{{ t('userDocument.searchBtn') }}</n-button>
    </n-input-group>
  </n-card>

  <n-card v-for="item in doc_list" :key="item.category" class="doc-card" hoverable :embedded="true" :title="item.category" content-style="padding: 0;">
    <div
        class="doc-item"
        v-for="item in item.data"
        :key="item.category"
        @click="activate(themeStore.menuCollapsed?'bottom':'right', item.title, item.body)"
    >
      <p class="doc-title">{{ item.title }}</p>
      <p class="doc-release">{{ item.date }}</p>

    </div>
  </n-card>

  <n-drawer v-model:show="active" width="80%" height="80%" :placement="placement">
    <n-drawer-content :title="drawTitle">
<!--      <div v-html="convertMd(drawBody as string)"></div>-->
      <MdPreview :theme="themeStore.enableDarkMode?'dark':'light'" style="background-color: rgba(0,0,0,0.0)" v-model="drawBody"></MdPreview>
    </n-drawer-content>
  </n-drawer>

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

  .doc-card {
    margin-bottom: 20px;
    .doc-item {
      display: flex;
      flex-direction: column;
      padding: 10px 0 10px 25px;
      .doc-title {
        font-weight: bold;
        font-size: 1rem;
      }
      .doc-release {
        opacity: 0.8;
      }
    }
    .doc-item:hover {
      background-color: rgba(220, 220, 220, 0.2);
    }
  }
}

.n-card {
  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
  border: 0;
}

</style>