<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, h, onBeforeMount, onMounted, ref} from "vue";
import {NButton, NTag, useMessage, type DataTableColumns} from "naive-ui"
import useThemeStore from "@/stores/useThemeStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import {formatDate} from "@/utils/timeFormat";
import instance from "@/axios";
import {useRouter} from "vue-router";
import {ChevronForwardOutline as toRight,} from "@vicons/ionicons5"
import {handleGetAllMyActivateLog, handleUnbindById} from "@/api/user/record"
import DataTableSuffix from "@/views/utils/DataTableSuffix.vue";
import PageHead from "@/views/utils/PageHead.vue";
import useTablePagination from "@/hooks/useTablePagination"

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
// let pageCount = ref(10)
//
// let dataSize = ref<{ pageSize: number, page: number }>({
//   pageSize: 10,
//   page: 1,
// })


let [dataSize, pageCount] = useTablePagination()

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

const columns = computed<DataTableColumns<ActivateRecord>>(() => [
  {
    title: t('userActivation.id'),
    key: 'id',
  },
  {
    title: t('userActivation.keyId'),
    key: 'key_id',
    render(row: ActivateRecord) {
      return h('p', {}, {default: () => '# ' + row.key_id});
    }
  },
  {
    title: t('userActivation.isBind'),
    key: 'is_bind',
    render(row: ActivateRecord) {
      return h(
          NTag,
          {
            size: 'small',
            bordered: false,
            type: row.is_bind? 'success' : 'error',
          },
          {
            default: () => row.is_bind? t('userActivation.active') : t('userActivation.inactive')
          });
    }
  },
  {
    title: t('userActivation.orderId'),
    key: 'order_id',
  },
  {
    title: t('userActivation.clientVersion'),
    key: 'client_version',
    render(row: ActivateRecord) {
      return h(NTag, {size: 'small', bordered: false, type: 'default'}, {default: () => row.client_version});
    }
  },
  {
    title: t('userActivation.osType'),
    key: 'os_type',
    render(row: ActivateRecord) {
      return h(NTag, {size: 'small', bordered: false, type: 'default'}, {default: () => row.os_type});
    }
  },
  {
    title: t('userActivation.requestAt'),
    key: 'request_at',
    render(row: ActivateRecord) {
      return h('span', {}, {default: () => formatDate(row.request_at)});
    },
  },
  {
    title: t('userActivation.actions'),
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
          default: () => t('userActivation.showDetail'),
        }),
        h(NButton, {
          size: 'small',
          secondary: true,
          type: 'warning',
          style: {marginLeft: '10px'},
          bordered: false,
          disabled:!row.is_bind,
          onClick: () => callHandleUnbindById(row),
        }, {
          default: () => t('userActivation.cancelBind'),
        }),
      ]);
    },
  }
]);

let callHandleGetAllMyActivateLog = async () => {
  let data = await handleGetAllMyActivateLog(
      userInfoStore.thisUser.id,
      dataSize.value.page,
      dataSize.value.pageSize
  )
  if (data.code === 200) {
    activateRecordList.value = []
    data.log.forEach((log: ActivateRecord) => activateRecordList.value.push(log))
    pageCount.value = data.page_count
    animated.value = true
  }
}

let callHandleUnbindById = async (row: ActivateRecord) => {
  let data = await handleUnbindById(
      userInfoStore.thisUser.id,
      row.id,
  )
  if (data.code === 200) await callHandleGetAllMyActivateLog()
  else message.error('err:' + data.message || '')
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
  themeStore.breadcrumb = 'userActivation.activateLog'
  themeStore.userPath = '/dashboard/log'
  themeStore.menuSelected = 'user-activate-log'
})

onMounted(async () => {
  await callHandleGetAllMyActivateLog()
  animated.value = true
})

</script>

<script lang="ts">
export default {
  name: "ActivateLog",
}
</script>

<template>
  <PageHead
      :title="t('userActivation.activateLog')"
      :description="t('userActivation.description')"
  />

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
      <DataTableSuffix
          v-model:data-size="dataSize"
          v-model:page-count="pageCount"
          v-model:animated="animated"
          :update-data="callHandleGetAllMyActivateLog"
      />
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
          <div
              style="display: flex; flex-direction: row; align-items: center; justify-content: center; margin-left: 10px;">
            <p>{{ !enableAlterRemark ? t('userActivation.alterRemark') : t('userActivation.commitRemark') }}</p>

          </div>
        </n-button>
      </div>
      <n-input
          :rows="2"
          type="textarea"
          size="large"
          :disabled="!enableAlterRemark"
          :placeholder="t('userActivation.setRemark')"
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
.root {
  padding: 0 20px;
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