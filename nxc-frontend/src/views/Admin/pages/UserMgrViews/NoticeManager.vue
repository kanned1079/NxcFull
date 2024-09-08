<script setup lang="ts" name="NoticeManager">
import {onMounted, ref, reactive} from "vue";
import useThemeStore from "@/stores/useThemeStore.ts"
import {NButton} from "naive-ui";

const themeStore = useThemeStore()

let showModal = ref(false)

let fromData = reactive({
  title: '',
  content: '',
  tags: '',
  img_url: '',
})

let handleAddNotice = () => {
  console.log('处理添加一条通知')
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