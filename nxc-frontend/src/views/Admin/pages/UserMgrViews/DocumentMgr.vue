<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, h, onMounted, ref} from "vue";
import {MdEditor} from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import useThemeStore from "@/stores/useThemeStore";
import useApiAddrStore from "@/stores/useApiAddrStore";
import {type DrawerPlacement, type FormInst, NButton, NSwitch, useMessage} from 'naive-ui'
import instance from "@/axios/index";
import {formatDate} from "@/utils/timeFormat";

const {t} = useI18n()
const apiAddrStore = useApiAddrStore();
const placement = ref<DrawerPlacement>('right')
let editType = ref<'add' | 'edit'>('add')
const active = ref(false)
const activate = (place: DrawerPlacement) => {
  active.value = true
  placement.value = place
}
const themeStore = useThemeStore();
const text = ref('# Hello Editor');
let animated = ref<boolean>(false)
// 表单
const formRef = ref<FormInst | null>(null)
const message = useMessage()

let pageCount = ref(10)

let dataSize = ref<{ pageSize: number, page: number }>({
  pageSize: 10,
  page: 1,
})

let dataCountOptions = [
  {
    label: '10条数据/页',
    value: 10,
  },
  {
    label: '20条数据/页',
    value: 20,
  },
  {
    label: '50条数据/页',
    value: 50,
  },
  {
    label: '100条数据/页',
    value: 100,
  },
]

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
    language: '',
    md_text: '',
    sort: 0
  }
})

let rules = {
  doc: {
    title: {
      required: true,
      message: '文档标题是必填的',
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
    sort: {
      required: true,
      trigger: 'blur'
    }
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
  Object.keys(formValue.value.doc).forEach(key => {
    if (typeof formValue.value.doc[key] === 'string')
      formValue.value.doc[key] = ''
    if (typeof formValue.value.doc[key] === 'number')
      formValue.value.doc[key] = 0
  })
}

let submitDoc = async () => {
  if (formValue.value.doc.title !== '') {
    animated.value = false
    active.value = false
    console.log(formValue.value)
    let {data} = await instance.post('/api/admin/v1/document', {
      ...formValue.value.doc
    })
    if (data.code === 200) {
      console.log(data)
      animated.value = true
      message.success('添加新的文档成功')
      // 清除表单值
      clearForm()
      await getAllDocumentList()
    } else {
      message.error('Add failure, please retry.')

    }
  } else {
    message.error('文档标题不能为空')
  }
}

// -----修改---------------------------------------------------------

const columns = [
  {
    title: computed(() => t('adminViews.docMgr.docId')).value,
    key: 'id',
    render(row: DocumentItem) {
      return h('p', {}, row.id);
    }
  },
  {
    title: computed(() => t('adminViews.docMgr.isShow')).value,
    key: 'show',
    render(row: DocumentItem) {
      return h(NSwitch, {
        value: row.show,
        onUpdateValue: () => updateDocumentShow(row),
      });
    }
  },
  {
    title: computed(() => t('adminViews.docMgr.sortAs')).value,
    key: 'title',
    render(row: DocumentItem) {
      return h('p', {}, row.sort);
    }
  },
  {
    title: computed(() => t('adminViews.docMgr.lang')).value,
    key: 'language',
    render(row: DocumentItem) {
      return h('p', {}, row.language);
    }
  },
  {
    title: computed(() => t('adminViews.docMgr.category')).value,
    key: ' category',
    render(row: DocumentItem) {
      return h('p', {}, row.category);
    }
  },
  {
    title: computed(() => t('adminViews.docMgr.title')).value,
    key: 'title',
    render(row: DocumentItem) {
      return h('p', {}, row.title);
    }
  },
  {
    title: computed(() => t('adminViews.docMgr.createdAt')).value,
    key: 'created_at',
    render(row: DocumentItem) {
      return h('p', {}, formatDate(row.created_at));
    }
  },
  {
    title: computed(() => t('adminViews.docMgr.updatedAt')).value,
    key: 'updated_at',
    render(row: DocumentItem) {
      return h('p', {}, formatDate(row.updated_at));
    }
  },
  {
    title: computed(() => t('adminViews.docMgr.operate')).value,
    key: 'actions',
    render(row: DocumentItem) {
      return h('div', {style: {display: 'flex', flexDirection: 'row'}}, [
        h(NButton, {
          size: 'small',
          type: 'primary',
          bordered: false,
          style: {marginLeft: '10px'},
          onClick: () => handleUpdateDocClick(row),
        }, {default: () => computed(() => t('adminViews.docMgr.edit')).value,}),
        h(NButton, {
          size: 'small',
          type: 'error',
          style: {marginLeft: '10px'},
          onClick: () => handleDeleteDoc(row)
        }, {default: () => computed(() => t('adminViews.docMgr.delete')).value,})
      ]);
    }
  }
];

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

let updateDocument = async () => {
  try {
    let {data} = await instance.put('/api/admin/v1/document', {
      id: updateDocId.value,  //  更新文档的id
      ...formValue.value.doc  // 更新的内容
    })
    if (data.code === 200) {
      message.success('Updated ok')
      await getAllDocumentList()
      active.value = false
    } else {
      message.error('Updated failure ' + data.msg || '')
    }
  } catch (error: any) {
    message.error(error.toString)
  }
}

let updateDocumentShow = async (row: DocumentItem) => {
  try {
    let {data} = await instance.patch('/api/admin/v1/document', {
      id: row.id,
    })
    if (data.code === 200) {
      message.success('Updated ok')
      await getAllDocumentList()
    } else {
      message.error('Updated failure ' + data.msg || '')
    }
  } catch (error: any) {
    message.error(error.toString)
  }
}

let getAllDocumentList = async () => {
  try {
    // animated.value = false
    let {data} = await instance.get('/api/admin/v1/document', {
      params: {
        page: dataSize.value.page,
        size: dataSize.value.pageSize,
      }
    })
    if (data.code === 200) {
      console.log(data.documents)
      documentItemList.value = []
      data.documents.forEach((doc: DocumentItem) => documentItemList.value.push(doc))
      pageCount.value = data.page_count as number
      message.success('Updated ok')
      animated.value = true
    } else {
      message.error('Updated failure ' + data.msg || '')
    }
  } catch (error: any) {
    message.error(error.toString)
  }
}

let handleDeleteDoc = async (row: DocumentItem) => {
  try {
    let {data} = await instance.delete('/api/admin/v1/document', {
      params: {
        id: row.id,
      }
    })
    if (data.code === 200 && data.deleted) {
      message.success('Delete doc ok, id: ' + row.id)
      await getAllDocumentList()
    } else {
      message.error('Delete doc failure' + data.msg || '')
    }
  } catch (error: any) {
    message.error('unknown err: '+ error)
  }
}

onMounted(async () => {
  themeStore.menuSelected = 'doc-manager'
  themeStore.contentPath = '/admin/dashboard/document';

  // request
  await getAllDocumentList()
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
  <div style="padding: 20px 20px 0 20px">
    <n-card class="body-card" hoverable :embedded="true" title="知识库管理" :bordered="false">
      <n-button type="primary" :bordered="false" @click="editType='add'; activate('right')">添加文档</n-button>
    </n-card>
  </div>

  <transition name="slide-fade">
    <div class="root" v-if="animated">
      <n-card hoverable :embedded="true" :bordered="false" content-style="padding: 0">
        <n-data-table
            class="table"
            :columns="columns"
            :data="documentItemList"
            :pagination="false"
            :bordered="true"
            style=""
            :scroll-x="800"
        />
      </n-card>

      <div style="margin-top: 20px; display: flex; flex-direction: row; justify-content: right;">
        <n-pagination
            size="medium"
            v-model:page.number="dataSize.page"
            :page-count="pageCount"
            @update:page="animated=false; getAllDocumentList()"
        />
        <n-select
            style="width: 160px; margin-left: 20px"
            v-model:value.number="dataSize.pageSize"
            size="small"
            :options="dataCountOptions"
            :remote="true"
            @update:value="animated=false; dataSize.page = 1; getAllDocumentList()"
        />
      </div>
    </div>
  </transition>


  <n-drawer v-model:show="active" width="80%" :placement="placement" @after-leave="clearForm()">
    <n-drawer-content :title="editType==='add'?'新增文档':'修改文档'">
      <n-form
          ref="formRef"
          :model="formValue"
          :rules="rules"
      >
        <n-form-item label="标题" path="doc.title">
          <n-input v-model:value="formValue.doc.title" placeholder="输入文档标题"/>
        </n-form-item>
        <n-form-item label="排序" path="doc.sort">
          <n-input v-model.number="formValue.doc.sort" placeholder="输入文档的排序级别"/>
        </n-form-item>
        <n-form-item label="分类" path="doc.category">
          <n-input v-model:value="formValue.doc.category" placeholder="输入文档分类"/>
        </n-form-item>
        <n-form-item label="语言" path="doc.language">
          <n-input v-model:value="formValue.doc.language" placeholder="选择文档语言"/>
        </n-form-item>
        <MdEditor style="background-color: rgba(0,0,0,0.0)" v-model="formValue.doc.md_text"
                  :theme="themeStore.enableDarkMode?'dark':'light'"/>
      </n-form>
      <div style="display: flex; flex-direction: row; justify-content: right; margin-top: 20px">
        <n-button style="margin-right: 15px; width: 80px" @click="cancelAdd" secondary type="primary">取消</n-button>
        <n-button style="width: 80px" type="primary" @click="editType==='add'?submitDoc():updateDocument()">
          {{ editType === 'add' ? '添加' : '修改' }}
        </n-button>
      </div>
    </n-drawer-content>
  </n-drawer>
</template>

<style scoped lang="less">
.root {
  padding: 20px;
}
</style>