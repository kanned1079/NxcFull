<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, onBeforeMount, ref} from "vue"
import {useMessage} from "naive-ui";
import useThemeStore from "@/stores/useThemeStore";
import useUserInfoStore from "@/stores/useUserInfoStore";

const themeStore = useThemeStore()
const userInfoStore = useUserInfoStore()

let isAuthed = ref<boolean>(false)
isAuthed.value = userInfoStore.thisUser.isAdmin

interface ChatBobbleTheme {
  senderBgColor: string
  senderTextColor: string
  receiverBgColor: string
  receiverTextColor: string
  senderBgShallow: string
  receiverBgShallow: string
}

let chatBobbleColorTheme = computed(() => !themeStore.enableDarkMode ? {
  senderBgColor: '#5d8fc2',
  senderTextColor: '#fff',
  receiverBgColor: '#dadada',
  receiverTextColor: '#000',
  senderBgShallow: 'rgba(93,143,193,0.5)',
  receiverBgShallow: 'rgba(218,218,218,0.5)'
} as ChatBobbleTheme : {
  senderBgColor: '#486993',
  senderTextColor: '#fff',
  receiverBgColor: '#696969',
  receiverTextColor: '#fff',
  senderBgShallow: 'rgba(30,45,58,0.5)',
  receiverBgShallow: 'rgba(71,71,71,0.5)'
} as ChatBobbleTheme).value as ChatBobbleTheme

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
    body: '这是用户说的',
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
    body: '菠萝味百岁山文艺复兴菠萝味百岁山文艺复兴菠萝味百岁山文艺复兴 ',
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
    id: 1,
    sender: 0,
    body: '这不是很这不是很戏剧吗这不是很戏剧吗这不是很戏剧吗这不是很戏剧吗这不是很戏剧吗这不是很戏剧吗戏剧吗这不是很戏剧吗这不是很戏剧吗',
    date: '2024-10-28 14:13:23'
  },
  {
    id: 6,
    sender: 0,
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
    id: 5,
    sender: 0,
    body: '那我明天课也不去了吧',
    date: '2024-10-28 14:13:23'
  },
  {
    id: 4,
    sender: 1,
    body: '@菠萝味百岁山文艺复兴 ',
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
      <n-card :bordered="false" :embedded="true" hoverable content-style="padding: 0; border-radius: 0;" class="header">
        <p class="subject-name">充值未到帳問題</p>
      </n-card>
    </n-layout-header>

    <n-layout-content class="content" style="margin: 60px 0 70px 0" content-style="padding: 0;">
      <n-scrollbar class="sc">
        <div class="message-p" v-for="msg in messages" :key="msg.id">
          <div class="left" v-if="msg.sender === 0">
            <div class="left-message-body">
              <p class="left-message-body-tag">{{ msg.body }}</p>
            </div>
            <div class="left-date">
              <p class="left-date-tag">{{ msg.date }}</p>
            </div>
          </div>
          <div class="right" v-if="msg.sender === 1">
            <div class="right-message-body">
              <p class="right-message-body-tag">{{ msg.body }}</p>
            </div>
            <div class="right-date">
              <p class="right-date-tag">{{ msg.date }}</p>
            </div>
          </div>
        </div>
      </n-scrollbar>
    </n-layout-content>

    <n-layout-footer>

      <div class="send-box">
        <n-card class="send-box-btn" content-style="padding: 0">
          <div class="send-box-btn-body">
            <n-input size="large" class="text-input" placeholder="輸入要發送的消息"></n-input>
            <n-button size="large" secondary type="primary" class="send-btn">發送</n-button>
          </div>

        </n-card>
      </div>

    </n-layout-footer>
  </n-layout>

</template>

<style scoped lang="less">


.header {
  height: 60px;
  position: fixed;
  left: 0;
  right: 0;
  z-index: 1001;
  line-height: 60px;

  .subject-name {
    margin-left: 20px;
    font-size: 1.2rem;
    font-weight: 500;
  }
}

.content {
  height: 100vh;
}

.sc {
  max-height: 80vh;
  overflow-y: auto;

  /* 自定义滚动条样式 */
  scrollbar-width: thin;
  scrollbar-color: #cfcfcf #f5f5f5;

  .message-p {
    margin: 20px 0;
    display: flex;


    /* 控制左右消息的对齐 */

    .left,
    .right {
      max-width: 80%;
      display: flex;
      flex-direction: column;
    }

    .left {
      margin-left: 20px;
      align-items: flex-start;
      margin-right: auto; /* 保证左侧对齐 */
      //background-color: @primary-color;

      .left-message-body {
        .left-message-body-tag {
          max-width: 100%;
          background-color: v-bind('chatBobbleColorTheme.receiverBgColor');
          padding: 12px 16px;
          border-radius: 10px;
          font-size: 0.9rem;
          line-height: 1.5;
          color: v-bind('chatBobbleColorTheme.receiverTextColor');
          transition: ease 250ms;
        }

        .left-message-body-tag:hover {
          box-shadow: 5px 8px 16px v-bind('chatBobbleColorTheme.receiverBgShallow');
          transform: translateX(0px) translateY(-3px);
        }
      }

      .left-date {
        font-size: 0.7rem;
        color: #3a4754;
        margin-top: 5px;
        align-self: flex-start;
      }
    }

    .right {
      margin-right: 20px;
      align-items: flex-end;
      margin-left: auto; /* 保证右侧对齐 */
      //background-color: @secondary-color;

      .right-message-body {
        .right-message-body-tag {
          max-width: 100%;
          background-color: v-bind('chatBobbleColorTheme.senderBgColor');
          padding: 12px 16px;
          border-radius: 10px;
          font-size: 0.9rem;
          line-height: 1.5;
          color: v-bind('chatBobbleColorTheme.senderTextColor');
          transition: ease 250ms;
        }

        .right-message-body-tag:hover {
          box-shadow: 5px 8px 16px v-bind('chatBobbleColorTheme.senderBgShallow');
          transform: translateX(0px) translateY(-3px);
        }
      }

      .right-date {
        font-size: 0.7rem;
        color: #3a4754;
        margin-top: 5px;
        align-self: flex-end;
      }
    }
  }
}

.send-box {
  position: fixed;
  width: 100%;
  bottom: 0;

  .send-box-btn {
    padding: 15px;

    .send-box-btn-body {
      display: flex;
      flex-direction: row;

      .text-input {
        flex: 8;
        margin-right: 15px;
      }

      .send-btn {
        flex: 2;
      }
    }


  }


}

</style>