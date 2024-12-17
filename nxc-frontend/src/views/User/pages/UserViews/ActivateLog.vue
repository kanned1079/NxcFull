<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {onBeforeMount, onMounted, ref, computed, h} from "vue";
import {NButton, useMessage} from "naive-ui"
import useThemeStore from "@/stores/useThemeStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import {formatDate} from "@/utils/timeFormat";
import instance from "@/axios";

const {t} = useI18n();
const message = useMessage();
const themeStore = useThemeStore()
const userInfoStore = useUserInfoStore()

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

let showDetails = (index: number) => {
  console.log(index)
  message.info("show " + index)
}

const columns = [
  {
    title: computed(() => t('userActivation.id')).value,
    key: 'id',
  },
  {
    title: computed(() => t('userActivation.orderId')).value,
    key: 'order_id',
  },
  {
    title: computed(() => t('userActivation.keyId')).value,
    key: 'key_id',
  },
  {
    title: computed(() => t('userActivation.requestAt')).value,
    key: 'request_at',
    render(row: ActivateRecord) {
      return h('span', {}, { default: () => formatDate(row.request_at) });
    },
  },
  {
    title: computed(() => t('userActivation.clientVersion')).value,
    key: 'client_version',
  },
  {
    title: computed(() => t('userActivation.osType')).value,
    key: 'os_type',
  },
  {
    title: computed(() => t('userActivation.createdAt')).value,
    key: 'created_at',
    render(row: ActivateRecord) {
      return h('span', {}, { default: () => formatDate(row.created_at) });
    },
  },
  {
    title: computed(() => t('userActivation.actions')).value,
    key: 'actions',
    render(row: ActivateRecord) {
      return h('div', { style: { display: 'flex', flexDirection: 'row' } }, [
        h(NButton, {
          size: 'small',
          type: 'primary',
          onClick: () => showDetails(row.id),
        }, {
          default: () => computed(() => t('userActivation.showDetail')).value,
        }),
      ]);
    },
    width: 200,
    fixed: 'right',
  }
];

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
  created_at: string;
  updated_at: string;
}

let activateRecordList = ref<ActivateRecord[]>([])

let handleGetAllMyActivateLog = async () => {
  try {
    let {data} = await instance.get('/api/user/v1/activation', {
      params: {
        user_id: userInfoStore.thisUser.id,
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
  name: "ActivateLog",
}
</script>

<template>
  <div class="root-header">
    <n-card
        hoverable
        :embedded="true"
        :bordered="false"
        :title="t('userActivation.activateLog')"
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
        />
      </n-card>
      <div style="margin-top: 20px; display: flex; flex-direction: row; justify-content: right;">
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


</template>

<style scoped lang="less">
.root-header {
  padding: 20px 20px 0 20px;
}

.root {
  padding: 20px;
}
</style>