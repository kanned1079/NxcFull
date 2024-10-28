<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {onBeforeMount, ref} from "vue"
import {useMessage} from "naive-ui";

let paramsData = ref<{
  userId: number
  ticketId: number
}>({
  userId: 0,
  ticketId: 0,
})

let messages = ref<{
  id: number
  sender: number
  body: string
  date: string
}[]>([
  {
    id: 0,
    sender: 1,
    body: '不是你们真没赶上车啊',
    date: '2024-10-28 14:13:23'
  },
  {
    id: 1,
    sender: 0,
    body: '这不是很戏剧吗这不是很戏剧吗这不是很戏剧吗',
    date: '2024-10-28 14:13:23'
  },
  {
    id: 3,
    sender: 0,
    body: '别走啊',
    date: '2024-10-28 14:13:23'
  },
  {
    id: 4,
    sender: 1,
    body: '@菠萝味百岁山文艺复兴 ',
    date: '2024-10-28 14:13:23'
  },
  {
    id: 5,
    sender: 0,
    body: '那我明天课也不去了吧',
    date: '2024-10-28 14:13:23'
  },
  {
    id: 6,
    sender: 1,
    body: '在你身后',
    date: '2024-10-28 14:13:23'
  },
  {
    id: 6,
    sender: 1,
    body: '在你身后',
    date: '2024-10-28 14:13:23'
  },
  {
    id: 6,
    sender: 1,
    body: '在你身后',
    date: '2024-10-28 14:13:23'
  },
  {
    id: 6,
    sender: 1,
    body: '在你身后',
    date: '2024-10-28 14:13:23'
  },
  {
    id: 6,
    sender: 1,
    body: '在你身后',
    date: '2024-10-28 14:13:23'
  },
])

const {t} = useI18n()
const message = useMessage()

onBeforeMount(() => {
  message.info('beforeMount')
  const urlParams = new URLSearchParams(window.location.search);
  let user_id = Number(urlParams.get('user_id'));
  let ticket_id = Number(urlParams.get('ticket_id'));
  if (user_id <= 0 || ticket_id <= 0) {
    message.warning('非法訪問')
  }
  paramsData.value.userId = user_id
  paramsData.value.ticketId = ticket_id
})
</script>

<script lang="ts">
export default {
  name: 'ChatDialog',
}
</script>

<template>


  <n-layout>
    <n-layout-header>
      <div class="header">
        <p class="subject-name">充值未到帳問題</p>
      </div>
    </n-layout-header>
    <n-layout-content style="margin-top: 50px" content-style="padding: 24px;">
      <n-scrollbar class="sc">
        <div class="message-p" v-for="msg in messages" :key="msg.id">
          <div class="left" v-if="msg.sender === 0">
            <div class="l-text">
              <div class="l-text-body">{{ msg.body }}</div>
              <p style="margin-left: 5px; opacity: 0.3">{{ msg.date }}</p>
            </div>

          </div>
          <div class="right" v-if="msg.sender === 1">
            <div style="display: flex; flex-direction: column; justify-content: end">
              <p style="margin-right: 5px; opacity: 0.3">{{ msg.date }}</p>
            </div>
            <div class="r-text">{{ msg.body }}</div>
          </div>

        </div>
      </n-scrollbar>
    </n-layout-content>
    <n-layout-footer>
      <div class="send-box">
        <div>
          <n-input style="width: 240px"></n-input>
          <n-button>發送</n-button>
        </div>
      </div>
    </n-layout-footer>
  </n-layout>


  <!--  <n-card :bordered="false" class="root" content-style="padding: 0">-->

  <!--    <div class="message-box">-->
  <!--      -->

  <!--    </div>-->
  <!--    -->
  <!--  </n-card>-->
</template>

<style scoped lang="less">


.header {
  height: 60px;
  background-color: #454545;
  position: fixed;
  left: 0;
  right: 0;
  z-index: 1001;
  line-height: 60px;

  .subject-name {
    margin-left: 10px;
    font-size: 1.25rem;
    font-weight: 400;
  }
}

.sc {
  margin-top: 60px;
}

.message-box {
  padding-top: 60px;

  .message-p {
    height: 30px;
    margin: 40px 0;


    .left {
      display: flex;
      flex-direction: row;
      justify-content: left;

      .l-text {
        display: flex;
        flex-direction: column;

        .l-text-body {
          font-size: 1rem;
          background-color: #656565;
          padding: 5px;
          border-radius: 3px;
        }

      }
    }

    .right {
      display: flex;
      flex-direction: row;
      justify-content: right;

      .r-text {
        font-size: 1rem;
        background-color: #858585;
        color: white;
        padding: 5px;
        border-radius: 3px;
      }
    }
  }
}

.send-box {
  height: 50px;
  background-color: #0e1b42;
}

</style>