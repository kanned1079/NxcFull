<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, nextTick, onBeforeMount, onBeforeUnmount, onMounted, ref} from "vue";
import {NScrollbar, useMessage,} from "naive-ui";
import {formatDate} from "@/utils/timeFormat";
import useThemeStore from "@/stores/useThemeStore";
import useUserInfoStore from "@/stores/useUserInfoStore";

const themeStore = useThemeStore();
const userInfoStore = useUserInfoStore();
const {t} = useI18n();
const message = useMessage();

const scrollbar = ref<HTMLElement | null>(null);

let token = ref<string>(sessionStorage.getItem('token') || '');

interface ChatHistoryItem {
  id: number
  user_id: number
  sent_at: string
  sender_type: string
  content: string
}

let chatHistory = ref<ChatHistoryItem[]>([])

// WebSocket 连接实例
let socket: WebSocket | null = null;

// 创建和管理 WebSocket 连接
const createWebSocket = () => {
  if (socket) socket.close(); // 先关闭已有连接

  socket = new WebSocket(`ws://localhost:8081/ws/user/v1/chat?token=${token.value}&user_id=${paramsData.value.userId}&ticket_id=${paramsData.value.ticketId}`);

  socket.onopen = () => {
    console.log("WebSocket connection established");
  };

  socket.onmessage = (event: MessageEvent) => {
    try {
      const messageData = JSON.parse(event.data);
      // console.log(messageData.type)
      // 判断 Type 字段并根据类型执行不同操作
      if (messageData.type === "history") {
        // 清空现有聊天记录并添加新记录
        let previousLength = chatHistory.value.length
        // console.log('previousLength', previousLength)
        chatHistory.value = [];
        // console.log(messageData.history)
        // let chatMsgItem = messageData.history
        messageData.history.forEach((chatMsgItem: ChatHistoryItem) => {
          chatHistory.value.push(chatMsgItem);
        });
        let newLength = chatHistory.value.length
        // 使用 scrollTo 将滚动条滚动到底部
        if (previousLength != newLength) {
          setTimeout(() => {
            nextTick(() => {
              scrollbar.value?.scrollTo({top: 999999, behavior: 'smooth'});
            });
          }, 300)
        }
        // 如果工单已经结束 那么请求一次之后就停止ws连接
        if (paramsData.value.status === 204) {
          if (socket && socket.readyState === WebSocket.OPEN) {
            socket.close();
            console.log("WebSocket connection closed on unmount");
          }
        }

      } else if (messageData.type === "check") {
        // 处理单一消息（或成功响应）
        message.success("发送消息成功");
      } else {
        console.warn("Unknown message type received:", messageData.type);
      }

    } catch (error) {
      console.error("Failed to parse message:", error);
    }
  };

  socket.onclose = () => {
    if (paramsData.value.status !== 204) {
      console.log("WebSocket connection closed. Reconnecting...");
      setTimeout(createWebSocket, 5000); // 每5秒尝试重新连接
    }
  };

  socket.onerror = (error) => {
    console.error("WebSocket error:", error);
  };
};

// 发送消息函数
const sendWsMessage = (content: any) => {
  if (socket && socket.readyState === WebSocket.OPEN) {
    const msg = JSON.stringify(content);
    socket.send(msg);
    // console.log("发送的消息为:", msg);
    // message.success("发送消息成功");
  } else {
    message.error("WebSocket is not open. Unable to send message.");
  }
};

let isAuthed = ref<boolean>(userInfoStore.thisUser.isAdmin);

// 处理主题色的逻辑
interface ChatBobbleTheme {
  senderBgColor: string;
  senderTextColor: string;
  receiverBgColor: string;
  receiverTextColor: string;
  senderBgShallow: string;
  receiverBgShallow: string;
}

// let chatBobbleColorTheme = computed(
//     () =>
//         !themeStore.enableDarkMode
//             ? {
//               senderBgColor: "#5d8fc2",
//               senderTextColor: "#fff",
//               receiverBgColor: "#dadada",
//               receiverTextColor: "#000",
//               senderBgShallow: "rgba(93,143,193,0.5)",
//               receiverBgShallow: "rgba(218,218,218,0.5)",
//             }
//             : {
//               senderBgColor: "#486993",
//               senderTextColor: "#fff",
//               receiverBgColor: "#696969",
//               receiverTextColor: "#fff",
//               senderBgShallow: "rgba(30,45,58,0.5)",
//               receiverBgShallow: "rgba(71,71,71,0.5)",
//             }
// ).value as ChatBobbleTheme;

let chatBobbleColorTheme = computed(
    () =>
        !themeStore.enableDarkMode
            ? {
              senderBgColor: themeStore.getTheme.topHeaderBgColor,
              senderTextColor: "#fff",
              receiverBgColor: "#dadada",
              receiverTextColor: "#000",
              senderBgShallow: "rgba(25,25,25,0.2)",
              receiverBgShallow: "rgba(218,218,218,0.5)",
            }
            : {
              senderBgColor: themeStore.getTheme.topHeaderBgColor,
              senderTextColor: "#fff",
              receiverBgColor: "#696969",
              receiverTextColor: "#fff",
              senderBgShallow: "rgba(25,25,25,0.2)",
              receiverBgShallow: "rgba(71,71,71,0.5)",
            }
).value as ChatBobbleTheme;

let paramsData = ref<{
  userId: number;
  ticketId: number;
  subject: string;
  status: number;
  role: number;
}>({
  userId: 0,
  ticketId: 0,
  subject: '',
  status: 201,
  role: 1
});

// 消息输入和显示逻辑
let msgInput = ref<string>("");
let handleSendMsgBtnClicked = () => {
  if (msgInput.value.trim() !== "") {
    sendWsMessage({
      user_id: paramsData.value.userId,
      ticket_id: paramsData.value.ticketId,
      role: paramsData.value.role === 0 ? 'admin' : 'user',
      content: msgInput.value.trim(),
    });
    msgInput.value = ''; // 清空输入框
  } else {
    message.error("消息内容不能为空");
  }
};

// 生命周期函数
onBeforeMount(() => {
  const urlParams = new URLSearchParams(window.location.search);
  let user_id = Number(urlParams.get("user_id"));
  let ticket_id = Number(urlParams.get("ticket_id"));
  let subject = urlParams.get("subject")?.trim() as string;
  let status = Number(urlParams.get("status"));
  let role = Number(urlParams.get("role"));

  if (user_id <= 0 || ticket_id <= 0) {
    message.warning("非法訪問");
    setTimeout(() => window.close(), 2000)
  }
  paramsData.value.userId = user_id;
  paramsData.value.ticketId = ticket_id;
  paramsData.value.subject = subject;
  paramsData.value.status = status;
  paramsData.value.role = role

  console.log("params: ", paramsData)
});

onMounted(() => {
  if (token.value) {
    createWebSocket();
  } else {
    message.error("Token is missing. Cannot establish WebSocket connection.");
  }
});

// 挂载前关闭ws连接
onBeforeUnmount(() => {
  if (socket && socket.readyState === WebSocket.OPEN) {
    socket.close();
    console.log("WebSocket connection closed on unmount");
  }
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
        <p class="subject-name">{{ paramsData.subject }}</p>
      </n-card>
    </n-layout-header>

    <n-layout-content class="content" style="margin: 60px 0 70px 0" content-style="padding: 0;">
      <n-scrollbar class="sc" ref="scrollbar">
        <div class="message-p" v-for="msg in chatHistory" :key="msg.id">
          <div class="left" v-if=" paramsData.role===1?msg.sender_type === 'admin':msg.sender_type === 'user'">
            <div class="left-message-body">
              <p class="left-message-body-tag">{{ msg.content }}</p>
            </div>
            <div class="left-date">
              <p class="left-date-tag">{{ formatDate(msg.sent_at) }}</p>
            </div>
          </div>
          <div class="right" v-else>
            <div class="right-message-body">
              <p class="right-message-body-tag">{{ msg.content }}</p>
            </div>
            <div class="right-date">
              <p class="right-date-tag">{{ formatDate(msg.sent_at) }}</p>
            </div>
          </div>
        </div>
      </n-scrollbar>
    </n-layout-content>

    <n-layout-footer>
      <n-card content-style="padding: 0" class="send-box">
        <n-card class="send-box-btn" content-style="padding: 0">
          <div class="send-box-btn-body">
            <n-input
                @keyup.enter="handleSendMsgBtnClicked"
                :disabled="paramsData.status===204"
                size="medium" class="text-input"
                :placeholder="paramsData.status===204?'工单已经结束':'輸入要發送的消息'"
                v-model:value="msgInput"
            ></n-input>
            <n-button
                :disabled="paramsData.status===204"
                :bordered="false"
                size="medium"
                type="primary"
                class="send-btn"
                @click="handleSendMsgBtnClicked"
            >
              Send
            </n-button>
          </div>
        </n-card>
      </n-card>
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
        opacity: 0.5;
        font-size: 0.7rem;
        color: v-bind('chatBobbleColorTheme.receiverTextColor');
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
        opacity: 0.5;
        font-size: 0.7rem;
        color: v-bind('chatBobbleColorTheme.receiverTextColor');
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