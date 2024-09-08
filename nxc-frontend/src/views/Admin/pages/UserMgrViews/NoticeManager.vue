<script setup lang="ts" name="NoticeManager">
import {onMounted, ref, reactive} from "vue";
import useThemeStore from "@/stores/useThemeStore.ts"
import {NButton} from "naive-ui";
import instance from "@/axios";

const themeStore = useThemeStore()

let showModal = ref(false)

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

let noticesData = reactive([]<Notices>)

// let fromData = reactive<Notices>({
//   id: 2,
//   title: '2024夏季优惠券发放 20%OFF',
//   content: '优惠券码：HappySummer2024Pre\n\t\t\t此优惠码使用截至日期2024-08-31，主站点与常州站点同步皆可使用，请阁下在确认订单时在页面上方填入此优惠码。祝大家2024夏天快乐！',
//   show: true,
//   img_url: '',
//   tags: "coupons",
//   createdAt: "2024-09-08T21:23:43+08:00",
//   updatedAt: "2024-09-08T21:23:43+08:00",
//   deletedAt: null
// })

let getAllNotices = async () => {
  let {data} = await instance.get('http://localhost:8080/api/admin/get-all-notices')
  // console.log(data)
  Object.assign(data, noticesData)
  console.log(noticesData)
}

let handleAddNotice = () => {
  console.log('处理添加一条通知')
  getAllNotices()

  showModal.value = true

}

let closeModal = () => {
  console.log('close')
}

let submitModal = () => {
  console.log('submit')
}

let model = reactive({

})


onMounted(() => {
  console.log('NoticeManager mounted.')

  themeStore.menuSelected = 'notice-manager'
  themeStore.contentPath = '/admin/dashboard/noticemanager'
})

</script>

<template>
  <div class="root">
    <n-card title="公告管理" hoverable :embedded="true" class="card">
      <n-button class="add-btn" @click="handleAddNotice">添加公告</n-button>

      <n-modal
          title="新建通知"
          v-model:show="showModal"
          preset="dialog"
          positive-text="确认"
          negative-text="算了"
          style="width: 480px"
          @positive-click="closeModal"
          @negative-click="submitModal"
          :show-icon="false"
      >

          <!--          pass-->
        <div style="margin-top: 30px"></div>

          <n-form-item path="age" label="标题">
            <n-input @keydown.enter.prevent placeholder="输入通知标题" v-model:value="fromData.title"/>
          </n-form-item>

          <n-form-item path="age" label="公告内容">
            <n-input type="textarea" placeholder="输入公告内容" :rows="8" show-count  />
          </n-form-item>

          <n-form-item path="age" label="公告标签">
            <n-input @keydown.enter.prevent placeholder="输入公告的标签"/>
          </n-form-item>

          <n-form-item path="age" label="图片URL">
            <n-input @keydown.enter.prevent placeholder="输入背景图片的URL"/>
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

.n-card {
  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
  border: 0;
}
</style>