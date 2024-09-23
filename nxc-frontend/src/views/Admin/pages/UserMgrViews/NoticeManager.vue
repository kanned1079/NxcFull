<script setup lang="ts">
import {onMounted, reactive, ref} from "vue";
import useThemeStore from "@/stores/useThemeStore.ts"
import instance from "@/axios";
import type {DataTableColumns} from 'naive-ui'
import {NButton, NSwitch, useMessage} from 'naive-ui'
import useNoticesStore from "@/stores/useNoticesStore";
import useApiAddrStore from "@/stores/useApiAddrStore";
const apiAddrStore = useApiAddrStore();

const themeStore = useThemeStore()
const noticesStore = useNoticesStore()
const message = useMessage()

interface formData {
  title: string
  content: string
  tags: string
  img_url?: string
}

let formData = reactive<formData>({
  title: '',
  content: '',
  tags: '',
  img_url: '',
})

// noticesStore.getAllNotices()

// 显示表单
let showModal = ref(false)

// 接口
interface Notices {
  id: int,
  title: string,
  content: string,
  show: boolean,
  img_url?: string,
  tags?: string,
  created_at: string,
  updated_at?: string,
  deleted_at?: string,
}

// 添加新的
let handleAddNotice = async () => {
  await console.log('处理添加一条通知')


  showModal.value = true

}

let closeModal = () => {
  console.log('close')
}

let submitModal = async () => {
  console.log('submit')
  console.log(formData)
  let {data} = await instance.post(apiAddrStore.apiAddr.admin.addANotice, formData)
  if (data.code === 200) {
    message.success("添加成功")
  }

}


let createColumns = ({sendMail}: { sendMail: (rowData: Notices) => void }): DataTableColumns<Notices> => {
  return [
    {
      title: 'ID',
      key: 'id',
    },
    {
      title: '是否显示',
      key: 'show',
      render(row) {
        return h(NSwitch, {
          size: 'small',
          value: row.show,
        })
      }
    },
    {
      title: '标题',
      key: 'title'
    },
    {
      title: '创建时间',
      key: 'created_at',

    },
    {
      title: '操作',
      key: 'operate',
      render(row) {
        return h(
            NButton,
            {
              size: 'small',
              onClick: () => sendMail(row)
            },
            {default: () => 'Send Email'}
        )
      }
    },

  ]
}

let columns = createColumns({
  sendMail(rowData) {
    message.info(`send mail to ${rowData.name}`)
  }
})

let pagination = {
  pageSize: 10
}

// 获取所有的通知
let getAllNotices = async () => {
  try {
    let {data} = await instance.get(apiAddrStore.apiAddr.admin.getAllNoticesList)
    if (data.code === 200) {
      // console.log(data.Document)
      noticesStore.noticesArr = data.notices
      console.log('axios获取的数据', noticesStore.noticesArr)
      // createData()
    } else {
      message.error('获取失败')
    }
  } catch (err) {
    message.error('未知错误 ' + err)
  }
}

// 删除一条通知
let deleteNotice = async (id: number) => {
  const message = useMessage()
  try {
    let {data} = await instance.delete(apiAddrStore.apiAddr.admin.deleteANotice,
        {
          params: {notice_id: notice_id,},
        },
    )
    data.code === 200 ? message.success("删除成功") : message.error("出现错误")
  } catch (err) {
    message.error("奇怪的错误" + err)
  }
  // console.log(data)
}

onMounted(() => {
  console.log('NoticeManager挂载')
  getAllNotices()
  console.log(noticesStore.noticesArr)
  // noticesStore.aa()

  // 保存历史记录
  themeStore.menuSelected = 'notice-manager'
  themeStore.contentPath = '/admin/dashboard/noticemanager'
})


let handleDeleteNotice = async () => {
  await deleteNotice(2)
}

</script>

<script lang="ts">
export default {
  name: 'NoticeManager',
}
</script>

<template>
  <div class="root">
    <n-card title="公告管理" hoverable :embedded="true" class="card" :bordered="false">
      <n-button type="primary" :bordered="false" class="add-btn" @click="handleAddNotice">添加公告</n-button>
      <n-button type="primary" :bordered="false" style="margin-left: 20px" @click="handleDeleteNotice">测试</n-button>
      <n-data-table
          style="margin-top: 20px"
          :bordered="false"
          :columns="columns"
          :data="noticesStore.noticesArr"
          :pagination="pagination"
      />

      <n-modal
          title="新建通知"
          v-model:show="showModal"
          preset="dialog"
          positive-text="确认添加"
          negative-text="算了"
          style="width: 480px"
          @positive-click="submitModal"
          @negative-click="closeModal"
          :show-icon="false"
      >
        <!--          pass-->
        <div style="margin-top: 30px"></div>

        <n-form-item path="age" label="标题">
          <n-input @keydown.enter.prevent placeholder="输入通知标题" v-model:value="formData.title"/>
        </n-form-item>

        <n-form-item path="age" label="公告内容">
          <n-input type="textarea" placeholder="输入公告内容" :rows="8" show-count v-model:value="formData.content"/>
        </n-form-item>

        <n-form-item path="age" label="公告标签">
          <n-input @keydown.enter.prevent placeholder="输入公告的标签" v-model:value="formData.tags"/>
        </n-form-item>

        <n-form-item path="age" label="图片URL">
          <n-input @keydown.enter.prevent placeholder="输入背景图片的URL" v-model:value="formData.img_url"/>
        </n-form-item>


        <!--          pass-->
        <template #footer>
          尾部
        </template>


      </n-modal>

    </n-card>
  </div>
</template>

<style lang="less" scoped>
.root {
  padding: 20px;

  .card {
    border: 0;

    .add-btn {
      margin-top: 10px;
    }
  }
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

</style>