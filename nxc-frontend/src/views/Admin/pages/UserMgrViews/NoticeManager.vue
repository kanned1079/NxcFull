<script setup lang="ts">
import  {useI18n} from "vue-i18n";
import {computed, h, onMounted, onBeforeMount, ref} from "vue";
import useThemeStore from "@/stores/useThemeStore"
import instance from "@/axios";
import {AddOutline as AddIcon} from "@vicons/ionicons5"
import {NButton, NIcon, NSwitch, useMessage, type DataTableColumns, type FormInst, type FormRules} from 'naive-ui'
// import useNoticesStore from "@/stores/useNoticesStore";
import {formatDate} from "@/utils/timeFormat"
import DataTableSuffix from "@/views/utils/DataTableSuffix.vue";
import PageHead from "@/views/utils/PageHead.vue";
import {handleFetchAllNotices} from "@/api/admin/notice";

// const apiAddrStore = useApiAddrStore();
const {t} = useI18n()
const i18nPrefix = 'adminViews.noticeMgr'
const themeStore = useThemeStore()
// const noticesStore = useNoticesStore()
const message = useMessage()

let animated = ref<boolean>(false)


let editType = ref<'add' | 'edit'>('add')

let pageCount = ref(10)

let dataSize = ref<{ pageSize: number, page: number }>({
  pageSize: 10,
  page: 1,
})

interface formData {
  id?: number
  title: string
  content: string
  tags: string
  img_url?: string
}

let formRef = ref<FormInst | null>(null)
let formData = ref<formData>({
  title: '',
  content: '',
  tags: '',
  img_url: '',
})

const rules: FormRules = {
  title: {
    required: true,
    trigger: 'blur',
    type: 'string',
    validator() {
      return formData.value.title.trim()?true:new Error('empty not allowed')
    }
  },
  content: {
    required: true,
    trigger: 'blur',
    type: 'string',
    validator() {
      return formData.value.content.trim()?true:new Error('empty not allowed')
    }
  },
}

// noticesStore.getAllNotices()

// 显示表单
let showModal = ref(false)

// 接口
interface Notice {
  id: number,
  title: string,
  content: string,
  show: boolean,
  img_url?: string,
  tags?: string,
  created_at: string,
  updated_at?: string,
  deleted_at?: string,
}

let noticesArr = ref<Notice[]>([])

// 添加新的
let handleAddNotice = () => {
  // formData.value = {}
  Object.assign(formData.value, {
    title: '',
    content: '',
    tags: '',
    img_url: '',
  })
  console.log('处理添加一条通知')
  editType.value = 'add'
  showModal.value = true
}

let handleEditNotice = (row: Notice) => {
  // formData.value = {}
  Object.assign(formData.value, {
    title: '',
    content: '',
    tags: '',
    img_url: '',
  })
  console.log('处理修改一条通知')
  editType.value = 'edit'
  Object.assign(formData.value, row)
  showModal.value = true
}

let closeModal = () => {
  console.log('close')
}

// 这里是按钮按下
let submitModal = async () => {
  if (formRef.value) {
   await formRef.value.validate(async (err: any) => {
      if (!err) {
        switch (editType.value) {
          case 'add': {
            await addNewNotice()
            break
          }
          case 'edit': {
            await editNotice()
            break
          }
        }
      } else {
        message.warning(t('userLogin.checkForm'))
      }
    })
  }
}

let addNewNotice = async () => {
  try {
    animated.value = false
    let {data} = await instance.post('/api/admin/v1/notice', {
      ...formData.value
    })
    if (data.code === 200) {
      message.success("添加成功")
    }
    await getAllNotices()
  } catch (error: any) {
    console.log(error)
  }
}

const columns = computed<DataTableColumns<Notice>>(() => [
  {
    title: t(`${i18nPrefix}.table.id`),
    key: 'id',
    render(row: Notice) {
      return h('p', {}, row.id);
    }
  },
  {
    title: t(`${i18nPrefix}.table.show`),
    key: 'show',
    render(row: Notice) {
      return h(NSwitch, {
        value: row.show,
        onUpdateValue: (value) => updateNoticeEnabled(row.id, value),
      });
    }
  },
  {
    title: t(`${i18nPrefix}.table.title`),
    key: 'title',
    render(row: Notice) {
      return h('p', {}, row.title);
    }
  },
  {
    title: t(`${i18nPrefix}.table.createdAt`),
    key: 'created_at',
    render(row: Notice) {
      return h('p', {}, formatDate(row.created_at));
    }
  },
  {
    title: t(`${i18nPrefix}.table.action.title`),
    key: 'actions',
    fixed: 'right',
    render(row: Notice) {
      return h('div', {style: {display: 'flex', flexDirection: 'row'}}, [
        h(NButton, {
          size: 'small',
          type: 'primary',
          secondary: true,
          bordered: false,
          style: {marginLeft: '10px'},
          onClick: () => handleEditNotice(row)
        }, {default: () => t(`${i18nPrefix}.table.action.edit`)}),
        h(NButton, {
          size: 'small',
          type: 'error',
          secondary: true,
          style: {marginLeft: '10px'},
          onClick: () => deleteNotice(row.id as number)
        }, {default: () => t(`${i18nPrefix}.table.action.delete`)})
      ]);
    },
  }
])

// let pagination = {
//   pageSize: 10
// }

// 获取所有的通知
let getAllNotices = async () => {
  // try {
  //   let {data} = await instance.get('/api/admin/v1/notice', {
  //     params: {
  //       page: dataSize.value.page,
  //       size: dataSize.value.pageSize,
  //       is_user: false,
  //     }
  //   })
  let data = await handleFetchAllNotices(dataSize.value.page, dataSize.value.pageSize, false)
    if (data && data.code === 200) {
      noticesArr.value = []
      // console.log(1111)
      data.notices.forEach((notice: Notice) => noticesArr.value.push(notice))
      pageCount.value = data.page_count
      animated.value = true
    } else {
      message.error('获取失败')
    }
  // } catch (err) {
  //   message.error('未知错误 ' + err)
  // }
}

let updateNoticeEnabled = async (id: number, enabled: boolean) => {
  console.log("new: ", id, enabled)
  try {
    // animated.value = false
    let {data} = await instance.put('/api/admin/v1/notice/status', {
      id,
      is_show: enabled,
    })
    if (data.code === 200) {
      console.log('更新成功')
    }
    await getAllNotices()
  } catch (error: any) {
    console.log(error)
  }
}

// 删除一条通知
let deleteNotice = async (id: number) => {
  const message = useMessage()
  try {
    animated.value = false
    let {data} = await instance.delete('/api/admin/v1/notice',
        {
          params: {notice_id: id,},
        },
    )
    if (data.code === 200) {
      message.success("删除成功")
      await getAllNotices()
    } else {
      message.error("出现错误")
    }
    // data.code === 200 ? message.success("删除成功") : message.error("出现错误")
  } catch (err) {
    message.error("未知错误" + err)
  }
  // console.log(data)
}

// let handleEditNotice = () => {
//
// }

let editNotice = async () => {
  try {
    animated.value = false
    let {data} = await instance.put('/api/admin/v1/notice', {
      ...formData.value
    })
    if (data.code === 200) {
      message.success("修改成功")
    }
    await getAllNotices()
  } catch (error: any) {
    console.log(error)
  }
}

onBeforeMount(() => {
  themeStore.breadcrumb = `${i18nPrefix}.title`
  themeStore.menuSelected = 'notice-manager'

})

onMounted(async () => {
  // 保存历史记录
  themeStore.contentPath = '/admin/dashboard/noticemanager'
  // console.log('NoticeManager挂载')
  await getAllNotices()
  // console.log(noticesStore.noticesArr)
  // noticesStore.aa()
  animated.value = true
})


let handleDeleteNotice = async () => {
  await deleteNotice(7)
}

</script>

<script lang="ts">
export default {
  name: 'NoticeManager',
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
        @click="handleAddNotice">
      <template #icon>
        <n-icon>
          <AddIcon/>
        </n-icon>
      </template>
      {{ t(`${i18nPrefix}.addNotice`) }}
    </n-button>

  </PageHead>
<!--  <div style="padding: 20px 20px 0 20px">-->
<!--    <n-card title="公告管理" hoverable :embedded="true" class="card" :bordered="false">-->
<!--      <n-button secondary type="primary" :bordered="false" class="add-btn" @click="handleAddNotice">添加公告</n-button>-->
<!--      <n-button secondary type="primary" :bordered="false" style="margin-left: 20px" @click="handleDeleteNotice">测试-->
<!--      </n-button>-->
<!--    </n-card>-->
<!--  </div>-->

  <transition name="slide-fade">
    <div class="root" v-if="animated">
      <n-card :embedded="true" hoverable :bordered="false" content-style="padding: 0px" style="margin-top: 20px">
        <n-data-table
            striped
            class="table"
            :columns="columns"
            :data="noticesArr"
            :pagination="false"
            :bordered="true"
            style=""
            :scroll-x="1000"
        />
      </n-card>

      <DataTableSuffix
          v-model:data-size="dataSize"
          v-model:page-count="pageCount"
          v-model:animated="animated"
          :update-data="getAllNotices"
      />

<!--      <div style="margin-top: 20px; display: flex; flex-direction: row; justify-content: right;">-->
<!--        <n-pagination-->
<!--            size="medium"-->
<!--            v-model:page.number="dataSize.page"-->
<!--            :page-count="pageCount"-->
<!--            @update:page="getAllNotices() "-->
<!--        />-->
<!--        <n-select-->
<!--            style="width: 160px; margin-left: 20px"-->
<!--            v-model:value.number="dataSize.pageSize"-->
<!--            size="small"-->
<!--            :options="dataCountOptions"-->
<!--            :remote="true"-->
<!--            @update:value="dataSize.page = 1; getAllNotices()"-->
<!--        />-->
<!--      </div>-->
    </div>
  </transition>


  <n-modal
      :title="t(`${i18nPrefix}.modal.addNew`)"
      v-model:show="showModal"
      preset="dialog"
      :positive-text="t('adminViews.common.confirm')"
      :negative-text="t('adminViews.common.cancel')"
      style="width: 480px"
      @positive-click="submitModal"
      @negative-click="closeModal"
      :show-icon="false"
      :block-scroll="1600"
  >
    <!--          pass-->
    <div style="margin-top: 30px"></div>

    <n-form
      :show-feedback="false"
      :show-label="true"
      :show-require-mark="true"
      :model="formData"
      :rules="rules"
      size="medium"
      ref="formRef"
    >

      <n-form-item class="form-gap-btn" path="title" :label="t(`${i18nPrefix}.modal.title.title`)">
        <n-input @keydown.enter.prevent :placeholder="t(`${i18nPrefix}.modal.title.placeholder`)" v-model:value="formData.title"/>
      </n-form-item>

      <n-form-item class="form-gap-btn" path="content" :label="t(`${i18nPrefix}.modal.content.title`)">
        <n-input type="textarea" :placeholder="t(`${i18nPrefix}.modal.content.placeholder`)" :rows="8" show-count v-model:value="formData.content"/>
      </n-form-item>

      <n-form-item class="form-gap-btn" path="tags" :label="t(`${i18nPrefix}.modal.tag.title`)">
        <n-input @keydown.enter.prevent :placeholder="t(`${i18nPrefix}.modal.tag.placeholder`)" v-model:value="formData.tags"/>
      </n-form-item>

      <n-form-item class="form-gap-btn" path="img_url" :label="t(`${i18nPrefix}.modal.img.title`)">
        <n-input @keydown.enter.prevent :placeholder="t(`${i18nPrefix}.modal.img.placeholder`)" v-model:value="formData.img_url"/>
      </n-form-item>
    </n-form>


  </n-modal>
</template>

<style lang="less" scoped>

.card {
  border: 0;

  .add-btn {
    margin-top: 10px;
  }
}

.root {
  padding: 0 20px 0 20px;


}

.form-root {
  width: auto;

  .line {
    width: 100%;
    display: flex;
    align-content: center;

    .title {
      width: 30px;
    }
  }
}

.form-gap-btn {
  margin-bottom: 20px;
}

</style>