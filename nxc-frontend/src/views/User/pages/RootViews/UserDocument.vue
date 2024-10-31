<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {ref, reactive, onMounted, onBeforeMount} from "vue";
// import {useRouter} from "vue-router";
import useThemeStore from "@/stores/useThemeStore";
// import useApiAddrStore from "@/stores/useApiAddrStore";
import  {type DrawerPlacement, useMessage } from 'naive-ui'
// import MarkdownIt from 'markdown-it';
// import {convertMd} from "@/utils/markdown";
import instance from "@/axios";
import {MdPreview} from "md-editor-v3";
import 'md-editor-v3/lib/style.css';

const placement = ref<DrawerPlacement>('right')

let animated = ref<boolean>(false)

const {t} = useI18n()
const message = useMessage()
// const apiAddrStore = useApiAddrStore();

let isEmpty = ref<boolean>(false)
let search = ref('')
// let text = ref('')
let active = ref<boolean>(false)

const activate = (place: DrawerPlacement, title:string, body: string) => {
  drawTitle.value = title
  drawBody.value = body
  active.value = true
  placement.value = place
}

const themeStore = useThemeStore();
// const router = useRouter();

let drawTitle = ref<string>('')
let drawBody = ref<string>('')

// 单个文档的接口定义
interface Document {
  id: number;
  language: string;
  category: string;
  title: string;
  body: string;
  sort: number;
  show: boolean;
  created_at: string; // 日期字符串
  updated_at: string; // 日期字符串
  deleted_at: string; // 日期字符串，可能为 null 或 undefined
}

// 类别及其文档列表的接口定义
interface CategoryDocuments {
  category: string;
  documents: Document[]; // 文档数组
}

let doc_list = ref<CategoryDocuments[]>([])

let getAllDocuments = async () => {
  console.log('拉取所有文档列表')
  doc_list.value = []
  try {
    let {data} = await instance.get('http://localhost:8081/api/user/v1/document', {
      params: {
        lang: 'zh_CN',
        find: search.value
      }
    })
    if (data.code === 200) {
      console.log(data)
      // Object.assign(doc_list, data.doc_list)
      if (!data.doc_list) {
        isEmpty.value = true
        return
      } else {
        isEmpty.value = false
      }
      data.doc_list.forEach((category: CategoryDocuments) => doc_list.value.push(category))
      console.log("doc_list", doc_list.value)

    } else {
      message.error(data.msg)
    }
  }catch (error: any) {
    message.error(error)
  }

}

// let showDetail = () => {
//
// }

onBeforeMount(() => {


})

onMounted(async () => {
  console.log('document挂载');
  themeStore.menuSelected = 'user-doc'
  themeStore.userPath = '/dashboard/document'
  // getAllDocuments()
  await getAllDocuments()
  animated.value = true

})


</script>

<script lang="ts">
export default {
  name: "UserDocument",
}
</script>

<template>
  <transition name="slide-fade">
    <div class="root" v-if="animated">
      <n-card hoverable :embedded="true" content-style="padding: 0;" class="search-root" :bordered="false">
        <n-input-group>
          <n-input v-model:value="search" :bordered="false" size="medium" class="search-input" :placeholder="t('userDocument.searchPlaceholder')"></n-input>
          <n-button @click="getAllDocuments" strong :bordered="false" type="primary" size="medium" class="search-btn">{{ t('userDocument.searchBtn') }}</n-button>
        </n-input-group>
      </n-card>

      <n-card v-for="item in doc_list" :key="item.category" class="doc-card" hoverable :embedded="true" :title="item.category" content-style="padding: 0;" :bordered="false">
        <div
            class="doc-item"
            v-for="doc in item.documents"
            :key="doc.category"
            @click="activate(themeStore.menuCollapsed?'bottom':'right', doc.title, doc.body)"
        >
          <p class="doc-title">{{ doc.title }}</p>
          <p class="doc-release">{{ doc.created_at }}</p>

        </div>
      </n-card>

      <n-card v-if="isEmpty" :embedded="true" hoverable :bordered="false" style="padding: 30px 0">
        <n-result status="404" title="查无结果" description="尝试换一个关键词吧">
        </n-result>
      </n-card>

      <n-drawer v-model:show="active" width="80%" height="80%" :placement="placement">
        <n-drawer-content :title="drawTitle">
          <!--      <div v-html="convertMd(drawBody as string)"></div>-->
          <MdPreview :theme="themeStore.enableDarkMode?'dark':'light'" style="background-color: rgba(0,0,0,0.0)" v-model="drawBody"></MdPreview>
        </n-drawer-content>
      </n-drawer>

    </div>
  </transition>

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
        font-weight: 400;
        font-size: 1rem;
      }
      .doc-release {
        opacity: 0.5;
      }
    }
    .doc-item:hover {
      background-color: rgba(220, 220, 220, 0.15);
    }
  }
}

</style>