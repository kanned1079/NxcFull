<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, h, onBeforeMount, onMounted, ref} from "vue"
import {type FormInst, NButton, NTag, useMessage} from "naive-ui"
import useThemeStore from "@/stores/useThemeStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import useAppInfosStore from "@/stores/useAppInfosStore";
import {formatDate} from "@/utils/timeFormat";
import {closeTicket, commitNewTicket, getAllMyTickets} from "@/api/user/tickets";

interface Ticket {
  subject: string
  urgency: number | null
  body: string
}

// 表格的字段名
interface TicketItem {
  id: number  // 工单Id
  subject: string // 工单主题
  urgency: number // 紧急程度 1低 2中 3高
  status: number // true工单未关闭 false工单关闭
  created_at: string  // 工单创建时间
  last_reply: string // 最后一次回复
}

const {t} = useI18n();
const appInfoStore = useAppInfosStore();
const userInfoStore = useUserInfoStore();
const themeStore = useThemeStore();
const message = useMessage();
const formRef = ref<FormInst | null>(null)

let animated = ref<boolean>(false)

let selectValidationStatus = computed(() => {
  return createStatus(newTicket.value.urgency)
})

let createStatus = (value: any) => {
  switch (value) {
    case 1:
      return undefined
    case 2:
      return undefined
    case 3:
      return undefined
    default:
      return 'error'
  }
}

const rules = {
  subject: {
    required: true,
    trigger: 'blur',
    message: ''
  },
  body: {
    required: true,
    trigger: 'blur',
    message: '',
  }
}

let ticketList = ref<TicketItem[]>([])

let formIsValid = ref<boolean>(false)


let validateClick = () => {
  formRef.value?.validate((errors) => {
    if (!errors) {
      // message.success('Valid')
      formIsValid.value = true
      // return true
    } else {
      console.log(errors)
      // message.error('Invalid')
      formIsValid.value = false
      // return false
    }
  })
}


let urgencyOptions = [
  {
    label: computed(() => t('userTickets.ticketUrgencyLevel.low')).value,
    value: 1,
  },
  {
    label: computed(() => t('userTickets.ticketUrgencyLevel.med')).value,
    value: 2,
  },
  {
    label: computed(() => t('userTickets.ticketUrgencyLevel.high')).value,
    value: 3,
  }
]


let newTicket = ref<Ticket>({
  subject: '',
  urgency: null,
  body: ''
})

let showNewTicketModal = ref<boolean>(false)

let cancelCreateNewTicket = () => {
  showNewTicketModal.value = false
  newTicket.value.subject = ''
  newTicket.value.urgency = null
  newTicket.value.body = ''
}

let callCloseTicket = async (ticket_id: number) => {
  let data = await closeTicket(userInfoStore.thisUser.id, ticket_id)
  if (data.code === 200 && data.closed) {
    message.success(computed(() => t('userTickets.ticketCloseSuccess')).value)
    await callGetAllMyTickets()
  } else {
    message.error(computed(() => t('userTickets.ticketCloseFailure')).value + data.msg as string || '')
  }
}

// let closeTicket = async (ticket_id: number) => {
//   try {
//     let {data} = await instance.delete('/api/user/v1/ticket', {
//       params: {
//         user_id: userInfoStore.thisUser.id,
//         ticket_id,
//       }
//     })
//     if (data.code === 200 && data.closed) {
//       message.success(computed(() => t('userTickets.ticketCloseSuccess')).value)
//       await callGetAllMyTickets()
//     } else {
//       message.error(computed(() => t('userTickets.ticketCloseFailure')).value + data.msg as string || '')
//     }
//   } catch (error: any) {
//     console.log(error)
//   }
// }

// -------------------------------------------------------------------------------------------------------------------------------------------------x


let callGetAllMyTickets = async () => {
  animated.value = false
  let data = await getAllMyTickets(userInfoStore.thisUser.id)
  if (data.code === 200) {
    ticketList.value = []
    data.tickets.forEach((ticket: TicketItem) => ticketList.value.push(ticket))
    animated.value = true
  } else if (data.code === 404) {
    message.info(computed(() => t('userTickets.noTickets')).value)
  } else {
    message.error(data.msg || '' as string)
  }
}

let callCommitNewTicket = async () => {
  validateClick()
  if (formIsValid.value) {
    console.log('表單驗證通過')
    let data = await commitNewTicket(userInfoStore.thisUser.id, newTicket.value)
    if (data.code === 200 && data.created) {
      message.success(computed(() => t('userTickets.commitNewTicketSuccess')).value)
      cancelCreateNewTicket()
      // await callGetAllMyTickets()
      await callGetAllMyTickets()
      console.log(data)
    } else {
      message.error(computed(() => t('userTickets.commitNewTicketFailure')).value + data.msg as string || '')
    }
  }
}

// let getAllMyTickets = async () => {
//   try {
//     animated.value = false
//     let {data} = await instance.get('/api/user/v1/ticket', {
//       params: {
//         user_id: userInfoStore.thisUser.id
//       }
//     })
//     if (data.code === 200) {
//       ticketList.value = []
//       data.tickets.forEach((ticket: TicketItem) => ticketList.value.push(ticket))
//       animated.value = true
//     } else if (data.code === 404) {
//       message.info(computed(() => t('userTickets.noTickets')).value)
//     } else {
//       message.error(data.msg || '' as string)
//     }
//   } catch (error: any) {
//     message.error(error.toString)
//   }
//
// }


// 点击打开工单后 单独打开一个聊天窗口
let openTicket = (ticket: TicketItem) => {
  message.info(`查看工单：${ticket.subject}`);
  const chatDialogWindow = window.open(
      `${appInfoStore.appCommonConfig.app_url}/user/tickets/chat?user_id=${userInfoStore.thisUser.id}&ticket_id=${ticket.id}&subject=${ticket.subject}&status=${ticket.status}&role=1`,
      'Chat',
      'width=480,height=640,resizable=yes,scrollbars=yes',
  );
  chatDialogWindow ? chatDialogWindow.addEventListener('onload', () => {
    chatDialogWindow.postMessage({greeting: 'Hello from the parent window!'}, '*'); // * 表示接受来自任何来源的消息
  }) : message.error("Cannot open Dialog.")

}


// 定义表格的列
const columns = ref([
  {title: computed(() => t('userTickets.ticketId')).value, key: 'id'},
  {title: computed(() => t('userTickets.ticketSubject')).value, key: 'subject'},
  {
    title: computed(() => t('userTickets.ticketUrgency')).value, key: 'urgency', render(row: TicketItem) {
      let level = '';
      switch (row.urgency) {
        case 1:
          level = computed(() => t('userTickets.ticketUrgencyLevel.low')).value
          break;
        case 2:
          level = computed(() => t('userTickets.ticketUrgencyLevel.med')).value
          break;
        case 3:
          level = computed(() => t('userTickets.ticketUrgencyLevel.high')).value
          break;
      }
      return h(NTag, {
        type: row.urgency === 3 ? 'error' : row.urgency === 2 ? 'warning' : 'success',
        bordered: false,
      }, {default: () => level});
    }
  },
  {
    title: computed(() => t('userTickets.ticketStatus')).value, key: 'status', render(row: TicketItem) {
      return h(NTag, {
        type: row.status !== 204 ? 'info' : 'default',
        bordered: false,
      }, {default: () => row.status !== 204 ? computed(() => t('userTickets.ticketActive')).value : computed(() => t('userTickets.ticketInActive')).value});
    }
  },
  {
    title: computed(() => t('userTickets.ticketCreatedAt')), key: 'created_at', render(row: TicketItem) {
      return formatDate(row.created_at)
    }
  },
  {
    title: computed(() => t('userTickets.lastResponse')), key: 'last_response', render(row: TicketItem) {
      return row.last_reply ? formatDate(row.last_reply) : computed(() => t('userTickets.noReply')).value
    }
  },
  {
    title: computed(() => t('userTickets.operate')).value, key: 'actions', fixed: 'right', render(row: TicketItem) {
      return h('div', {style: {display: 'flex', flexDirection: 'row'}}, [
        h(NButton, {
          size: 'small',
          type: 'primary',
          secondary: true,
          bordered: false,
          disabled: false,
          style: {marginLeft: '10px'},
          onClick: () => openTicket(row)
        }, {
          default: () => computed(() => t('userTickets.checkTicket')).value
        }),
        h(NButton, {
          size: 'small',
          type: 'error',
          secondary: true,
          disabled: row.status === 204,
          style: {marginLeft: '10px'},
          onClick: () => callCloseTicket(row.id)
        }, {
          default: () => computed(() => t('userTickets.closeTicket')).value
        })
      ]);
    }
  }
])

onBeforeMount(() => {
  themeStore.menuSelected = 'user-tickets'

})

onMounted(async () => {
  themeStore.userPath = '/dashboard/tickets'

  await callGetAllMyTickets()
  animated.value = true
})

</script>

<script lang="ts">
export default {
  name: 'Tickets',
}

</script>

<template>
  <div class="root">
    <n-card class="tickets-history" :embedded="true" hoverable content-style="padding: 0;" :bordered="false">
      <div class="header">
        <p class="title">{{ t('userTickets.userTickets') }}</p>
        <n-button :bordered="false" class="add-btn" type="primary" @click="showNewTicketModal = true">{{
            t('userTickets.addTicket')
          }}
        </n-button>
      </div>
    </n-card>

    <transition name="slide-fade">
      <n-card v-if="animated" :embedded="true" hoverable content-style="padding: 0;" :bordered="false">
        <n-data-table
            striped
            class="table"
            :columns="columns"
            :data="ticketList"
            :pagination="false"
            :bordered="true"
            style=""
            :scroll-x="800"
        />
      </n-card>
    </transition>


    <n-modal v-model:show="showNewTicketModal">
      <n-card
          style="width: 400px"
          :title="t('userTickets.addTicket')"
          :bordered="false"
      >
        <n-form
            ref="formRef"
            :rules="rules"
            :model="newTicket"
        >
          <n-form-item
              :label="t('userTickets.ticketSubject')"
              path="subject"
              required
              :show-require-mark="true"
          >
            <n-input v-model:value="newTicket.subject" clearable
                     :placeholder="t('userTickets.form.ticketSubjectDescription')"/>
          </n-form-item>
          <n-form-item
              :label="t('userTickets.ticketUrgency')"
              path="urgency"
              required
              :show-require-mark="true"
              :validation-status="selectValidationStatus"
          >
            <n-select
                :options="urgencyOptions" v-model:value.number="newTicket.urgency as number"
                :placeholder="t('userTickets.form.ticketUrgencyDescription')"

            ></n-select>
          </n-form-item>
          <n-form-item
              :label="t('userTickets.ticketContent')"
              path="body"
              required
              :show-require-mark="true"
          >
            <n-input v-model:value="newTicket.body" show-count type="textarea" :rows="6"
                     :placeholder="t('userTickets.form.ticketBody')"></n-input>
          </n-form-item>
        </n-form>
        <div style="display: flex; flex-direction: row; justify-content: right">
          <n-button style="width: 80px; margin-right: 10px;" :bordered="false" secondary type="primary"
                    @click="cancelCreateNewTicket">{{ t('userTickets.cancel') }}
          </n-button>
          <n-button style="width: 80px" :bordered="false" type="primary" @click="callCommitNewTicket">
            {{ t('userTickets.submit') }}
          </n-button>

        </div>
      </n-card>
    </n-modal>


  </div>
</template>

<style scoped lang="less">
.root {
  padding: 20px;

  .tickets-history {
    .header {
      display: flex;
      flex-direction: row;
      justify-content: space-between;
      padding: 15px 25px 15px 25px;
      align-items: center;

      .title {
        font-size: 1.1rem;
        font-weight: bold;
        opacity: 0.8;
      }
    }

    margin-bottom: 15px;
  }
}


</style>
