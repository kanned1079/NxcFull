<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, h, onBeforeMount, onMounted, ref} from "vue";
// import useApiAddrStore from "@/stores/useApiAddrStore";
import {type DataTableColumns, type DrawerPlacement, type FormInst, NButton, NIcon, NSwitch, useMessage} from "naive-ui"
import {MdEditor} from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import useThemeStore from "@/stores/useThemeStore";
import {AddOutline as AddIcon} from "@vicons/ionicons5"
import {formatDate} from "@/utils/timeFormat";
import DataTableSuffix from "@/views/utils/DataTableSuffix.vue";
import PageHead from "@/views/utils/PageHead.vue";
import {
  handleAddNewDoc,
  handleDeleteDocById,
  handleEditDocById,
  handleGetAllDoc,
  handleUpdateDocShow
} from "@/api/admin/doc";

const i18nPrefix = 'adminViews.docMgr'
const {t, locale} = useI18n()
// const apiAddrStore = useApiAddrStore();
const placement = ref<DrawerPlacement>('right')
let editType = ref<'add' | 'edit'>('add')
const active = ref(false)
const activate = (place: DrawerPlacement) => {
  active.value = true
  placement.value = place
}
const themeStore = useThemeStore();
// const text = ref('# Hello Editor');
let animated = ref<boolean>(false)
// 表单
const formRef = ref<FormInst | null>(null)
const message = useMessage()

let pageCount = ref(10)

let dataSize = ref<{ pageSize: number, page: number }>({
  pageSize: 10,
  page: 1,
})

const mdTheme = computed(() => themeStore.enableDarkMode?'dark':'light')
const codeTheme: string = 'github'

let previewTheme = ref<string>('github')

interface DocumentItem {
  id: number
  show: boolean
  sort: number
  language: string
  title: string
  category: string
  body: string
  created_at: string
  updated_at: string
}

let documentItemList = ref<DocumentItem[]>([])

let formValue = ref<{
  doc: {
    title: string,
    category: string,
    language: string,
    md_text: string,
    sort: number,
  }
}>({
  doc: {
    title: '',
    category: '',
    language: 'en_US',
    md_text: '',
    sort: 1000
  }
})

let rules = {
  doc: {
    title: {
      required: true,
      trigger: 'blur'
    },
    category: {
      required: false,
      trigger: 'blur',
    },
    language: {
      required: false,
      trigger: 'blur'
    },

  }
}

let cancelAdd = () => {
  active.value = false
  formValue.value.doc.title = ''
  formValue.value.doc.category = ''
  formValue.value.doc.language = ''
  formValue.value.doc.md_text = ''
  formValue.value.doc.sort = 0
}

let clearForm = () => {
  Object.assign(formValue.value.doc, {
    title: '',
    category: '',
    language: '',
    md_text: '',
    sort: 0
  })
}

// -----修改---------------------------------------------------------

const columns = computed<DataTableColumns<DocumentItem>>(() => [
  {
    title: t(`${i18nPrefix}.table.docId`),
    key: 'id',
    render(row: DocumentItem) {
      return h('p', {}, row.id);
    }
  },
  {
    title: t(`${i18nPrefix}.table.isShow`),
    key: 'show',
    render(row: DocumentItem) {
      return h(NSwitch, {
        value: row.show,
        onUpdateValue: () => callUpdateDocumentShow(row),
      });
    }
  },
  {
    title: t('adminViews.docMgr.table.sortAs'),
    key: 'title',
    render(row: DocumentItem) {
      return h('p', {}, row.sort);
    }
  },
  {
    title: t('adminViews.docMgr.table.lang'),
    key: 'language',
    render(row: DocumentItem) {
      return h('p', {}, row.language);
    }
  },
  {
    title: t('adminViews.docMgr.table.category'),
    key: ' category',
    render(row: DocumentItem) {
      return h('p', {}, row.category);
    }
  },
  {
    title: t('adminViews.docMgr.table.title'),
    key: 'title',
    render(row: DocumentItem) {
      return h('p', {}, row.title);
    }
  },
  {
    title: t('adminViews.docMgr.table.createdAt'),
    key: 'created_at',
    render(row: DocumentItem) {
      return h('p', {}, formatDate(row.created_at));
    }
  },
  {
    title: t('adminViews.docMgr.table.updatedAt'),
    key: 'updated_at',
    render(row: DocumentItem) {
      return h('p', {}, formatDate(row.updated_at));
    }
  },
  {
    title: t('adminViews.docMgr.table.operate'),
    key: 'actions',
    render(row: DocumentItem) {
      return h('div', {style: {display: 'flex', flexDirection: 'row'}}, [
        h(NButton, {
          size: 'small',
          type: 'primary',
          secondary: true,
          bordered: false,
          style: {marginLeft: '10px'},
          onClick: () => handleUpdateDocClick(row),
        }, {default: () => computed(() => t('adminViews.docMgr.table.edit')).value,}),
        h(NButton, {
          size: 'small',
          type: 'error',
          secondary: true,
          style: {marginLeft: '10px'},
          onClick: () => handleDeleteDoc(row)
        }, {default: () => computed(() => t('adminViews.docMgr.table.delete')).value,})
      ]);
    },
    fixed: 'right',
  }
])

let updateDocId = ref<number>(0)
let handleUpdateDocClick = (row: DocumentItem) => {
  formValue.value.doc.title = row.title// pass not fixed
  formValue.value.doc.sort = row.sort
  formValue.value.doc.category = row.category
  formValue.value.doc.language = row.language
  formValue.value.doc.md_text = row.body
  updateDocId.value = row.id
  editType.value = 'edit'
  active.value = true
}

// callHandleAddNewDoc 提交新的文档
const callHandleAddNewDoc = async () => {
  if (formValue.value.doc.title !== '') {
    animated.value = false
    active.value = false
    let data = await handleAddNewDoc(formValue.value.doc)
    if (data && data.code === 200) {
      animated.value = true
      message.success(t('adminViews.common.addSuccess'))
      // 清除表单值
      clearForm()
      await callGetAllDocumentList()
    } else {
      message.error(t('adminViews.common.addFailure'))
    }
  } else {
    message.error(t(`${i18nPrefix}.titleNotEmpty`))
  }
}

let callGetAllDocumentList = async () => {
  let data = await handleGetAllDoc(dataSize.value.page, dataSize.value.pageSize)
  if (data && data.code === 200) {
    documentItemList.value = []
    data.documents?.forEach((doc: DocumentItem) => documentItemList.value.push(doc))
    pageCount.value = data.page_count as number
    animated.value = true
  } else {
    message.error(t('adminViews.common.fetchDataFailure') + data.msg || '')
  }
}

let updateDocument = async () => {
  let data = await handleEditDocById(updateDocId.value, formValue.value.doc)
    if (data && data.code === 200) {
      await callGetAllDocumentList()
      active.value = false
    } else {
      message.error(t('adminViews.common.updateFailure') + data.msg || '')
    }
}

let callUpdateDocumentShow = async (row: DocumentItem) => {
  let data = await handleUpdateDocShow(row.id)
  if (data && data.code === 200) {
    await callGetAllDocumentList()
  } else {
    message.error(t('adminViews.common.updateFailure') + data.msg || '')
  }
}

let handleDeleteDoc = async (row: DocumentItem) => {
  let data = await handleDeleteDocById(row.id)
    if (data && data.code === 200 && data.deleted) {
      message.success(t('adminViews.common.deleteSuccess'))
      await callGetAllDocumentList()
    } else {
      message.error(t('adminViews.common.deleteFailure') + data.msg || '')
    }
}

onBeforeMount(() => {
  themeStore.menuSelected = 'doc-manager'
  themeStore.breadcrumb = 'adminViews.docMgr.title'
  formValue.value.doc.language = locale.value || 'en_US'
})

onMounted(async () => {
  themeStore.contentPath = '/admin/dashboard/document';


  // request
  await callGetAllDocumentList()
  setTimeout(() => animated.value = true, 50)
  // animated.value = true
})

</script>

<script lang="ts">
export default {
  name: 'DocumentMgr',
}
</script>

<template>
  <PageHead
      :title="t(`${i18nPrefix}.title`)"
      :description="t(`${i18nPrefix}.description`)"
  >
    <n-button
        tertiary
        type="primary"
        size="medium"
        class="btn-right"
        @click="editType='add'; activate('right')">
      <template #icon>
        <n-icon>
          <AddIcon/>
        </n-icon>
      </template>
      {{ t(`${i18nPrefix}.addDoc`) }}
    </n-button>

  </PageHead>


  <!--  <div style="padding: 20px 20px 0 20px">-->
  <!--    <n-card class="body-card" hoverable :embedded="true" title="知识库管理" :bordered="false">-->
  <!--      <n-button type="primary" :bordered="false" @click="editType='add'; activate('right')">添加文档</n-button>-->
  <!--    </n-card>-->
  <!--  </div>-->

  <transition name="slide-fade">
    <div class="root" v-if="animated">
      <n-card hoverable :embedded="true" :bordered="false" content-style="padding: 0">
        <n-data-table
            striped
            class="table"
            :columns="columns"
            :data="documentItemList"
            :pagination="false"
            :bordered="true"
            style=""
            :scroll-x="1600"
        />
      </n-card>

      <DataTableSuffix
          v-model:data-size="dataSize"
          v-model:page-count="pageCount"
          v-model:animated="animated"
          :update-data="callGetAllDocumentList"
      />


    </div>
  </transition>


  <n-drawer
      v-model:show="active"
      width="80%"
      :placement="placement"
      @after-leave="clearForm();"
  >
    <n-drawer-content :title="editType==='add'?t(`${i18nPrefix}.form.add`):t(`${i18nPrefix}.form.edit`)">
      <n-form
          ref="formRef"
          :model="formValue"
          :rules="rules"
          :show-feedback="false"
      >
        <n-form-item class="form-bt-gap" :label="t(`${i18nPrefix}.form.title.title`)" path="doc.title">
          <n-input v-model:value="formValue.doc.title" :placeholder="t(`${i18nPrefix}.form.title.placeholder`)"/>
        </n-form-item>
        <n-form-item class="form-bt-gap" :label="t(`${i18nPrefix}.form.sort.title`)" path="doc.sort">
          <n-input-number style="width: 100%" v-model:value.number="formValue.doc.sort"
                          :placeholder="t(`${i18nPrefix}.form.sort.placeholder`)"/>
        </n-form-item>
        <n-form-item class="form-bt-gap" :label="t(`${i18nPrefix}.form.category.title`)" path="doc.category">
          <n-input v-model:value="formValue.doc.category" :placeholder="t(`${i18nPrefix}.form.category.placeholder`)"/>
        </n-form-item>
        <n-form-item class="form-bt-gap" :label="t(`${i18nPrefix}.form.lang.title`)" path="doc.language">
          <!--          <n-input v-model:value="formValue.doc.language" placeholder="选择文档语言"/>-->
          <n-select
              :options="[{label: '简体中文', value: 'zh_CN'},{label: '繁體中文', value: 'zh_HK'},{label: 'English', value: 'en_US'},{label: '日本語', value: 'ja_JP'},]"
               v-model:value="formValue.doc.language" :placeholder="t(`${i18nPrefix}.form.lang.placeholder`)"/>
        </n-form-item>
        <MdEditor
            v-model="formValue.doc.md_text"
            :theme="mdTheme"
            :code-theme="codeTheme"
            :preview-theme="previewTheme"

        />
      </n-form>
      <div style="display: flex; flex-direction: row; justify-content: right; margin-top: 20px">
        <n-button style="margin-right: 15px;" @click="cancelAdd" secondary type="primary">
          {{ t(`${i18nPrefix}.form.cancel`) }}
        </n-button>
        <n-button style="" type="primary" @click="editType==='add'?callHandleAddNewDoc():updateDocument()">
          {{ editType === 'add' ? t(`${i18nPrefix}.form.addBtn`) : t(`${i18nPrefix}.form.editBtn`) }}
        </n-button>
      </div>
    </n-drawer-content>
  </n-drawer>
</template>

<style scoped lang="less">
.root {
  margin: 20px;
}

.form-bt-gap {
  margin-bottom: 14px;
}

.md-editor-dark, .md-editor-modal-container[data-theme='light'] {
  background-color: #f1f1f1;
}


.md-editor-dark, .md-editor-modal-container[data-theme='dark'] {
  background-color: #252525;
}



</style>