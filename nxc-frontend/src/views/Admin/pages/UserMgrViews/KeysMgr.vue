<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, h, onBeforeMount, onMounted, ref} from "vue";
import {NButton, NTag, useMessage} from "naive-ui"
import useThemeStore from "@/stores/useThemeStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import {formatDate} from "@/utils/timeFormat";
import instance from "@/axios";
import {useRouter} from "vue-router";
import {
  ChevronForwardOutline as toRight,
  CheckmarkOutline as checkIcon,
  SaveOutline as saveIcon,
} from "@vicons/ionicons5"

interface ActivateRecord {
  id: number;
  user_id: number;
  email: string;
  order_id: string;
  key_id: number;
  request_at: string;
  client_version: string;
  os_type: string;
  remark: string;
  is_bind: boolean;
  created_at: string;
  updated_at: string;
}

let activateRecordList = ref<ActivateRecord[]>([])
let currentRecord = ref<ActivateRecord>({
  id: 0,
  user_id: 0,
  email: '',
  order_id: '',
  key_id: 0,
  request_at: '',
  client_version: '',
  os_type: '',
  remark: '',
  is_bind: false,
  created_at: '',
  updated_at: '',
})

const {t} = useI18n();
const router = useRouter();
const message = useMessage();
const themeStore = useThemeStore()
const userInfoStore = useUserInfoStore()

let showModal = ref<boolean>(false)
let animated = ref<boolean>(false);
let pageCount = ref(10)

let dataSize = ref<{ pageSize: number, page: number }>({
  pageSize: 10,
  page: 1,
})

let dataCountOptions = [
  {
    label: computed(() => t('pagination.perPage10')).value,
    value: 10,
  },
  {
    label: computed(() => t('pagination.perPage20')).value,
    value: 20,
  },
  {
    label: computed(() => t('pagination.perPage50')).value,
    value: 50,
  },
  {
    label: computed(() => t('pagination.perPage100')).value,
    value: 100,
  },
]

let keyDetails = ref<{
  id: number;
  key: string;
  created_at: string;
}>({
  id: 0,
  key: '',
  created_at: '',
})
let showDetails = async (row: ActivateRecord) => {
  try {
    let {data} = await instance.get('/api/user/v1/key/details', {
      params: {
        key_id: row.key_id,
      }
    })
    if (data.code === 200) {
      Object.assign(currentRecord.value, row)
      Object.assign(keyDetails.value, data.details)
      showModal.value = true
    } else {
      message.error('err:' + data.msg)
    }
  } catch (err: any) {
    console.log(err)
  }
}

const columns = [
  {
    title: computed(() => t('userActivation.id')).value,
    key: 'id',
  },
  {
    title: computed(() => t('userActivation.keyId')).value,
    key: 'key_id',
    render(row: ActivateRecord) {
      return h('p', {}, {default: () => '# ' + row.key_id});
    }
  },
  {
    title: computed(() => t('userActivation.isBind')).value,
    key: 'is_bind',
    render(row: ActivateRecord) {
      return h(
          NTag,
          {
            size: 'small',
            bordered: false,
            type: row.is_bind ? 'success' : 'error',
          },
          {
            default: () => row.is_bind ? t('userActivation.active') : t('userActivation.inactive')
          });
    }
  },
  {
    title: computed(() => t('userActivation.orderId')).value,
    key: 'order_id',
  },
  {
    title: computed(() => t('userActivation.clientVersion')).value,
    key: 'client_version',
    render(row: ActivateRecord) {
      return h(NTag, {size: 'small', bordered: false, type: 'default'}, {default: () => row.client_version});
    }
  },
  {
    title: computed(() => t('userActivation.osType')).value,
    key: 'os_type',
    render(row: ActivateRecord) {
      return h(NTag, {size: 'small', bordered: false, type: 'default'}, {default: () => row.os_type});
    }
  },
  {
    title: computed(() => t('userActivation.requestAt')).value,
    key: 'request_at',
    render(row: ActivateRecord) {
      return h('span', {}, {default: () => formatDate(row.request_at)});
    },
  },
  {
    title: computed(() => t('userActivation.createdAt')).value,
    key: 'created_at',
    render(row: ActivateRecord) {
      return h('span', {}, {default: () => formatDate(row.created_at)});
    },
  },
  {
    title: computed(() => t('userActivation.actions')).value,
    key: 'actions',
    fixed: 'right',
    render(row: ActivateRecord) {
      return h('div', {style: {display: 'flex', flexDirection: 'row'}}, [
        h(NButton, {
          size: 'small',
          secondary: true,
          type: 'primary',
          bordered: false,
          onClick: () => showDetails(row),
        }, {
          default: () => computed(() => t('userActivation.showDetail')).value,
        }),
        h(NButton, {
          size: 'small',
          secondary: true,
          type: 'warning',
          style: {marginLeft: '10px'},
          bordered: false,
          disabled: !row.is_bind,
          onClick: () => handleUnbindById(row),
        }, {
          default: () => computed(() => t('userActivation.cancelBind')).value,
        }),
      ]);
    },

  }
];

let handleGetAllMyActivateLog = async () => {
  try {
    let {data} = await instance.get('/api/user/v1/activation', {
      params: {
        user_id: userInfoStore.thisUser.id,
        page: dataSize.value.page,
        size: dataSize.value.pageSize,
      }
    })
    if (data.code === 200) {
      activateRecordList.value = []
      data.log.forEach((log: ActivateRecord) => activateRecordList.value.push(log))
      pageCount.value = data.page_count
      animated.value = true
    }

  } catch (error: any) {
    console.log(error)
    message.error("unknown err: " + error)
  }
}

let handleUnbindById = async (row: ActivateRecord) => {
  animated.value = false
  try {
    let {data} = await instance.delete('/api/user/v1/activation', {
      params: {
        user_id: userInfoStore.thisUser.id,
        activation_id: row.id,
      }
    })
    if (data.code === 200) {
      await handleGetAllMyActivateLog()
    }

  } catch (error: any) {
    console.log(error)
    message.error("unknown err: " + error)
  }
}

let enableAlterRemark = ref<boolean>(false)

let handleCommitNewRemark = async () => {
  enableAlterRemark.value = false
  try {
    let {data} = await instance.patch('/api/user/v1/activation/remark', {
      user_id: userInfoStore.thisUser.id,
      record_id: currentRecord.value.id,
      remark: currentRecord.value.remark,
    })
    if (data.code === 200) {
      message.success(t('userActivation.updateSuccess'))
    } else {
      message.error('err: ' + data.msg || '')
    }
  } catch (err: any) {
    console.log(err)
  }
}


onBeforeMount(() => {
  themeStore.userPath = '/dashboard/log'
  themeStore.menuSelected = 'user-activate-log'
})

onMounted(async () => {
  await handleGetAllMyActivateLog()
  animated.value = true
})

</script>

<script lang="ts">
export default {
  name: "KeysMgr",
}
</script>

<template>
  <div class="root-header">
    <n-card
        hoverable
        :embedded="true"
        :bordered="false"
        :title="t('adminViews.keysMgr.keyMgr')"
    ></n-card>
  </div>
  <transition name="slide-fade">
    <div class="root" v-if="animated">
      <n-card
          hoverable
          :embedded="true"
          :bordered="false"
          content-style="padding: 0"
      >
        <n-data-table
            striped
            class="table"
            :columns="columns"
            :data="activateRecordList"
            :pagination="false"
            :bordered="true"
            :scroll-x="1000"
        />
      </n-card>
      <div style="margin-top: 15px; display: flex; flex-direction: row; justify-content: right;">
        <n-pagination
            size="medium"
            v-model:page.number="dataSize.page"
            :page-count="pageCount"
            @update:page="animated=false; handleGetAllMyActivateLog()"
        />
        <n-select
            style="width: 200px; margin-left: 20px"
            v-model:value.number="dataSize.pageSize"
            size="small"
            :options="dataCountOptions"
            :remote="true"
            @update:value="animated=false; dataSize.page = 1; handleGetAllMyActivateLog()"
        />
      </div>
    </div>
  </transition>

  <n-modal
      v-model:show="showModal"
      class="custom-card"
      preset="card"
      style="width: 640px"
      :title="t('userActivation.details')"
      size="huge"
      :bordered="false"
  >

    <div class="details-item-detail">
      <p class="details-item-title">{{ t('userActivation.keyId') }}</p>
      <p class="details-item-content">{{ '# ' + keyDetails.id }}</p>
    </div>

    <div class="details-item-detail">
      <p class="details-item-title">{{ t('userActivation.keyContent') }}</p>
      <p class="details-item-content">{{ keyDetails.key }}</p>
    </div>

    <div class="details-item-detail">
      <p class="details-item-title">{{ t('userActivation.keyGeneratedAt') }}</p>
      <p class="details-item-content">{{ keyDetails.created_at ? formatDate(keyDetails.created_at) : 'NULL' }}</p>
    </div>

    <div class="details-item-detail">
      <p class="details-item-title">{{ t('userActivation.activateRequestAt') }}</p>
      <p class="details-item-content">{{ currentRecord.request_at ? formatDate(currentRecord.request_at) : 'NULL' }}</p>
    </div>

    <div class="details-item-detail">
      <div class="details-item-detail-remark-title">
        <p class="details-item-title">{{ t('userActivation.remark') }}</p>
        <n-button
            type="primary"
            text
            size="small"
            style="text-decoration: underline"
            @click="!enableAlterRemark?enableAlterRemark=true:handleCommitNewRemark()"
        >
          <div style="display: flex; flex-direction: row; align-items: center; justify-content: center; margin-left: 10px;">
            <p>{{ !enableAlterRemark?t('userActivation.alterRemark'):t('userActivation.commitRemark') }}</p>

          </div>
        </n-button>
      </div>
      <n-input
          :rows="2"
          type="textarea"
          size="large"
          :disabled="!enableAlterRemark"
          :placeholder="'在這裡設置備注信息'"
          v-model:value.trim="currentRecord.remark"
          :bordered="true"
          style="margin-top: 5px"
      >
      </n-input>
    </div>

    <template #footer>
      <div style="width: 100%; display: flex; flex-direction: row; justify-content: flex-end; margin-top: 10px">
        <div>
          {{ t('userActivation.useIssueOccur') }}
          <n-button
              type="primary"
              text
              @click="router.push('/dashboard/tickets')"
          >
            {{ t('userActivation.chatWithUs') }}
            <n-icon style="margin-left: 4px">
              <toRight/>
            </n-icon>
          </n-button>
        </div>
      </div>
    </template>
  </n-modal>


</template>

<style scoped lang="less">
.root-header {
  padding: 20px 20px 0 20px;
}

.root {
  padding: 15px 20px;
}

.details-item-detail {
  //background-color: #66afe9;

  .details-item-title {
    font-size: 1rem;
    font-weight: 600;
    opacity: 0.7;
    margin-bottom: 5px;
  }

  .details-item-content {
    font-size: 1.2rem;
    font-weight: 600;
    opacity: 1;
    margin-bottom: 25px;
  }

  .details-item-detail-remark-title {
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    align-items: baseline;
  }

  .details-item-content-remark {
    font-size: 1rem;
    font-weight: 400;
    opacity: 1;
    margin-bottom: 25px;
  }

}
</style>