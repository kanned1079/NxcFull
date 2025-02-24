<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {onBeforeMount, onMounted, ref, watch} from "vue";
import useThemeStore from "@/stores/useThemeStore";
import {formatDate} from "@/utils/timeFormat";
import {type DrawerPlacement, useMessage} from 'naive-ui'
import {getAllDocuments} from "@/api/user/doc";
import {MdPreview} from "md-editor-v3";
import 'md-editor-v3/lib/style.css';
import PageHead from "@/views/utils/PageHead.vue";

const placement = ref<DrawerPlacement>('right')

let animated = ref<boolean>(false)

const {t, locale} = useI18n()
const message = useMessage()

let isEmpty = ref<boolean>(false)
let search = ref('')
let active = ref<boolean>(false)

const activate = (place: DrawerPlacement, title: string, body: string) => {
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

let callGetAllDocuments = async () => {
  animated.value = false
  console.log(locale.value)
  let data = await getAllDocuments(locale.value || 'en_US', search.value)
  if (data.code === 200) {
    doc_list.value = []

    // 处理查询为空的情况
    if (!data.doc_list || data.doc_list.length === 0) {
      isEmpty.value = true
      animated.value = true
      return
    } else {
      isEmpty.value = false
      // 如果有数据，填充文档列表
      data.doc_list.forEach((category: CategoryDocuments) => {
        doc_list.value.push(category)
      })
      animated.value = true
    }
  } else {
    message.error(data.msg)
  }
}

watch(() => locale.value, async () => {
  await callGetAllDocuments() // 切换语言的时候重新获取文档
})


onBeforeMount(() => {
  themeStore.breadcrumb = 'userDocument.title'
  themeStore.menuSelected = 'user-doc'
})

themeStore.userPath = '/dashboard/document'

onMounted(async () => {
  // getAllDocuments()
  await callGetAllDocuments()
  animated.value = true

})


</script>

<script lang="ts">
export default {
  name: "UserDocument",
}
</script>

<template>
  <div>
    <PageHead
        :title="t('userDocument.title')"
        :description="t('userDocument.description')"
    >

      <n-card
          hoverable
          :embedded="true"
          content-style="padding: 0;"
          class="search-root"
          :bordered="false"
      >
        <n-input-group>
          <n-input
              v-model:value="search"
              :bordered="false"
              size="medium"
              class="search-input"
              :placeholder="t('userDocument.searchPlaceholder')"
          ></n-input>
          <n-button
              @click="callGetAllDocuments"
              strong
              :bordered="false"
              type="primary"
              size="medium"
              class="search-btn"
          >
            {{ t('userDocument.searchBtn') }}
          </n-button>
        </n-input-group>
      </n-card>

    </PageHead>
  </div>


  <transition name="slide-fade">
    <div class="root" v-if="animated">
      <n-card
          v-for="item in doc_list"
          :key="item.category"
          class="doc-card"
          hoverable
          :embedded="true"
          :title="item.category"
          content-style="padding: 0;"
          :bordered="false"
      >
        <div
            class="doc-item"
            v-for="doc in item.documents"
            :key="doc.category"
            @click="activate(themeStore.menuCollapsed?'bottom':'right', doc.title, doc.body)"
        >
          <p class="doc-title">{{ doc.title }}</p>
          <p class="doc-release">{{ doc.created_at ? formatDate(doc.created_at) : '' }}</p>

        </div>
      </n-card>

      <n-card
          v-if="isEmpty"
          :embedded="true"
          hoverable
          :bordered="false"
          style="padding: 30px 0"
      >
        <n-result
            status="404"
            :title="t('userDocument.noContentTitle')"
            :description="t('userDocument.noContentTitleHint')"
        >
        </n-result>
      </n-card>

      <n-drawer
          v-model:show="active"
          width="80%"
          height="80%"
          :placement="placement"
      >
        <n-drawer-content
            :title="drawTitle"
        >
          <!--      <div v-html="convertMd(drawBody as string)"></div>-->
          <MdPreview
              :theme="themeStore.enableDarkMode?'dark':'light'"
              style="background-color: rgba(0,0,0,0.0)"
              v-model="drawBody"
              :preview-theme="'github'"
          ></MdPreview>
        </n-drawer-content>
      </n-drawer>

    </div>
  </transition>

</template>

<style scoped lang="less">
.root {
  padding: 0 20px;

  .search-root {
    display: flex;
    flex-direction: row;
    align-content: space-between;
    margin-bottom: 15px;

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
    margin-bottom: 10px;

    .doc-item {
      display: flex;
      flex-direction: column;
      padding: 10px 0 10px 25px;

      .doc-title {
        font-weight: 400;
        font-size: 0.9rem;
      }

      .doc-release {
        opacity: 0.5;
        font-size: 0.8rem;
      }
    }

    .doc-item:hover {
      background-color: rgba(220, 220, 220, 0.15);
    }
  }
}

</style>