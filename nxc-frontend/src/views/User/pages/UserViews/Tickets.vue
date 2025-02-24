<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, h, onBeforeMount, onMounted, ref} from "vue"
import {type FormInst, NButton, NTag, useMessage, type DataTableColumns, NIcon} from "naive-ui"
import useThemeStore from "@/stores/useThemeStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import useAppInfosStore from "@/stores/useAppInfosStore";
import {formatDate} from "@/utils/timeFormat";
import {closeTicket, commitNewTicket, getAllMyTickets} from "@/api/user/tickets";
import {AddOutline as AddIcon} from "@vicons/ionicons5"

import DataTableSuffix from "@/views/utils/DataTableSuffix.vue";
import PageHead from "@/views/utils/PageHead.vue";

let pageCount = ref<number>(0)
let dataSize = ref<{ pageSize: number, page: number }>({
  pageSize: 10,
  page: 1,
})

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

// let formIsValid = ref<boolean>(false)

let validateClick = async (): Promise<boolean> => {
  if (!formRef.value) return false;
  try {
    await formRef.value.validate(); // validate() 本身会抛出错误
    // formIsValid.value = true;
    return true;
  } catch (errors) {
    message.error(t('userTickets.checkForm'))
    console.log(errors);
    // formIsValid.value = false;
    return false;
  }
};


let urgencyOptions = computed<{
  label: string;
  value: number;
}[]>(() => [
  {
    label: t('userTickets.ticketUrgencyLevel.low'),
    value: 1,
  },
  {
    label: t('userTickets.ticketUrgencyLevel.med'),
    value: 2,
  },
  {
    label: t('userTickets.ticketUrgencyLevel.high'),
    value: 3,
  }
])


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

// -------------------------------------------------------------------------------------------------------------------------------------------------x

let callGetAllMyTickets = async () => {
  animated.value = false
  let data = await getAllMyTickets(userInfoStore.thisUser.id, dataSize.value.page, dataSize.value.pageSize)
  if (data.code === 200) {
    ticketList.value = []
    data.tickets.forEach((ticket: TicketItem) => ticketList.value.push(ticket))
    pageCount.value = data.page_count
    animated.value = true
  } else if (data.code === 404) {
    message.info(computed(() => t('userTickets.noTickets')).value)
  } else {
    message.error(data.msg || '' as string)
  }
}

let callCommitNewTicket = async () => {
  if (await validateClick()) {
    // if (formIsValid.value) {
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

const columns = computed<DataTableColumns<TicketItem>>(() => [
  {
    title: t('userTickets.ticketId'),
    key: 'id',
  },
  {
    title: t('userTickets.ticketSubject'),
    key: 'subject',
  },
  {
    title: t('userTickets.ticketUrgency'),
    key: 'urgency',
    render(row: TicketItem, rowIndex: number) {
      let level = '';
      switch (row.urgency) {
        case 1:
          level = t('userTickets.ticketUrgencyLevel.low');
          break;
        case 2:
          level = t('userTickets.ticketUrgencyLevel.med');
          break;
        case 3:
          level = t('userTickets.ticketUrgencyLevel.high');
          break;
      }
      return h(
          NTag,
          {
            type: row.urgency === 3 ? 'error' : row.urgency === 2 ? 'warning' : 'success',
            bordered: false,
          },
          { default: () => level }
      );
    },
  },
  {
    title: t('userTickets.ticketStatus'),
    key: 'status',
    render(row: TicketItem, rowIndex: number) {
      return h(
          NTag,
          {
            type: row.status !== 204 ? 'info' : 'default',
            bordered: false,
          },
          {
            default: () =>
                row.status !== 204
                    ? t('userTickets.ticketActive')
                    : t('userTickets.ticketInActive'),
          }
      );
    },
  },
  {
    title: t('userTickets.ticketCreatedAt'),
    key: 'created_at',
    render(row: TicketItem, rowIndex: number) {
      return formatDate(row.created_at);
    },
  },
  {
    title: t('userTickets.lastResponse'),
    key: 'last_response',
    render(row: TicketItem, rowIndex: number) {
      return row.last_reply
          ? formatDate(row.last_reply)
          : t('userTickets.noReply');
    },
  },
  {
    title: t('userTickets.operate'),
    key: 'actions',
    fixed: 'right',
    render(row: TicketItem, rowIndex: number) {
      return h(
          'div',
          { style: { display: 'flex', flexDirection: 'row' } },
          [
            h(
                NButton,
                {
                  size: 'small',
                  type: 'primary',
                  secondary: true,
                  bordered: false,
                  style: { marginLeft: '10px' },
                  onClick: () => openTicket(row),
                },
                { default: () => t('userTickets.checkTicket') }
            ),
            h(
                NButton,
                {
                  size: 'small',
                  type: 'error',
                  secondary: true,
                  disabled: row.status === 204,
                  style: { marginLeft: '10px' },
                  onClick: () => callCloseTicket(row.id),
                },
                { default: () => t('userTickets.closeTicket') }
            ),
          ]
      );
    },
  },
]);

onBeforeMount(() => {
  themeStore.breadcrumb = 'userTickets.userTickets'
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

  <PageHead
      :title="t('userTickets.userTickets')"
      :description="t('userTickets.description')"
  >
    <n-button
        tertiary
        type="primary"
        size="medium"
        class="btn-right"
        @click="showNewTicketModal = true">
      <template #icon>
        <n-icon>
          <AddIcon/>
        </n-icon>
      </template>

      {{ t('userTickets.addTicket') }}
    </n-button>
  </PageHead>

  <div class="root">
<!--    <n-card class="tickets-history" :embedded="true" hoverable content-style="padding: 0;" :bordered="false">-->
<!--      <div class="header">-->
<!--        <p class="title">{{ t('userTickets.userTickets') }}</p>-->
<!--        <n-button :bordered="false" class="add-btn" type="primary" @click="showNewTicketModal = true">{{-->
<!--            t('userTickets.addTicket')-->
<!--          }}-->
<!--        </n-button>-->
<!--      </div>-->
<!--    </n-card>-->

    <transition name="slide-fade">
      <div v-if="animated">
        <n-card :embedded="true" hoverable content-style="padding: 0;" :bordered="false">
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

        <DataTableSuffix
            v-model:data-size="dataSize"
            v-model:page-count="pageCount"
            v-model:animated="animated"
            :update-data="callGetAllMyTickets"
        />
      </div>

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
  padding: 0 20px 20px 20px;

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
