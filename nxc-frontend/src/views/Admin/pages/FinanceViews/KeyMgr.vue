<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, h, onBeforeMount, onMounted, ref} from "vue";
import {NButton, NIcon, NTag, useMessage} from "naive-ui"
import useThemeStore from "@/stores/useThemeStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
// import instance from "@/axios";
import {useRoute, useRouter} from "vue-router";
import {handleGetAllKeysByAdmin} from "@/api/admin/keys";
import PageHead from "@/views/utils/PageHead.vue";
import {RefreshOutline as refreshIcon, Search as searchIcon,} from "@vicons/ionicons5"


interface KeyItem {
  key_id: number;
  client_id: string
  is_valid: boolean;
  is_used: boolean;
  key: string;
  order_id: string;
  plan_name: string;
  email: string;
  expiration_date_stamp: number
}

let keyList = ref<KeyItem[]>([])
let searchMethod = ref<'id' | 'content'>('id')
let searchForm = ref<{
  keyId: number;
  keyContent: string;
}>({
  keyId: 0,
  keyContent: "",
})

const {t} = useI18n();
const router = useRouter();
const route = useRoute()

let showSearchModal = ref<boolean>(false)

let paramsKeyId: number = 0
const keyIdParam = route.params.key_id as string;
if (keyIdParam) {
  // 解析成功后检查是否为有效的数字
  const parsedId = parseInt(keyIdParam, 10);
  if (!isNaN(parsedId)) {
    paramsKeyId = parsedId;
  } else {
    console.error('Invalid key_id:', keyIdParam);
    paramsKeyId = 0; // 确保无效时默认为 0
  }
} else {
  console.error('key_id is missing or invalid');
  paramsKeyId = 0
}


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

let keyDetails = ref<{
  id: number;
  key: string;
  created_at: string;
}>({
  id: 0,
  key: '',
  created_at: '',
})

// let showDetails = async (row: ActivateRecord) => {
//   try {
//     let {data} = await instance.get('/api/user/v1/key/details', {
//       params: {
//         key_id: row.key_id,
//       }
//     })
//     if (data.code === 200) {
//       Object.assign(currentRecord.value, row)
//       Object.assign(keyDetails.value, data.details)
//     } else {
//       message.error('err:' + data.msg)
//     }
//   } catch (err: any) {
//     console.log(err)
//   }
// }

const columns = [
  {
    title: computed(() => t('userActivation.id')).value,
    key: 'key_id',
  },
  // {
  //   title: computed(() => t('userActivation.keyId')).value,
  //   key: 'key_id',
  //   render(row: KeyItem) {
  //     return h('p', {}, {default: () => '# ' + row.key_id});
  //   }
  // },
  // {
  //   title: computed(() => t('userActivation.isBind')).value,
  //   key: 'is_bind',
  //   render(row: KeyItem) {
  //     return h(
  //         NTag,
  //         {
  //           size: 'small',
  //           bordered: false,
  //           type: row.is_bind ? 'success' : 'error',
  //         },
  //         {
  //           default: () => row.is_bind ? t('userActivation.active') : t('userActivation.inactive')
  //         });
  //   }
  // },
  {
    title: computed(() => t('userActivation.orderId')).value,
    key: 'order_id',
  },
  {
    title: computed(() => '邮箱地址').value,
    key: 'email',
  },
  {
    title: '密钥',
    key: 'key',
    render(row: KeyItem) {
      return h(NTag, {size: 'small', bordered: false, type: 'primary'}, {default: () => row.key})
    },
  },
  {
    title: computed(() => t('adminViews.keysMgr.table.isValid')).value,
    key: 'os_type',
    render(row: KeyItem) {
      return h(NTag, {
        size: 'small',
        bordered: false,
        type: row.is_valid ? 'success' : 'default'
      }, {default: () => row.is_valid ? t('adminViews.keysMgr.table.ok') : t('adminViews.keysMgr.table.blocked')});
    }
  },
  {
    title: computed(() => t('adminViews.keysMgr.table.isUsed')).value,
    key: 'os_type',
    render(row: KeyItem) {
      return h(NTag, {
        size: 'small',
        bordered: false,
        type: row.is_used ? 'success' : 'default'
      }, {default: () => row.is_used ? t('adminViews.keysMgr.table.used') : t('adminViews.keysMgr.table.unUsed')});
    }
  },
  // {
  //   title: computed(() => t('userActivation.osType')).value,
  //   key: 'os_type',
  //   render(row: KeyItem) {
  //     return h(NTag, {size: 'small', bordered: false, type: 'default'}, {default: () => row.os_type});
  //   }
  // },
  {
    title: computed(() => t('userActivation.requestAt')).value,
    key: 'request_at',
    render(row: KeyItem) {
      return h('span', {}, {default: () => row.expiration_date_stamp});
    },
  },
  // {
  //   title: computed(() => t('userActivation.createdAt')).value,
  //   key: 'created_at',
  //   render(row: ActivateRecord) {
  //     return h('span', {}, {default: () => formatDate(row.created_at)});
  //   },
  // },
  {
    title: computed(() => t('userActivation.actions')).value,
    key: 'actions',
    fixed: 'right',
    render(row: KeyItem) {
      return h('div', {style: {display: 'flex', flexDirection: 'row'}}, [
        h(NButton, {
          size: 'small',
          secondary: true,
          type: 'primary',
          bordered: false,
          onClick: () => console.log('ssss'),
        }, {
          default: () => computed(() => '转到密钥').value,
        }),
        // h(NButton, {
        //   size: 'small',
        //   secondary: true,
        //   type: 'warning',
        //   style: {marginLeft: '10px'},
        //   bordered: false,
        //   disabled: !row.is_bind,
        //   onClick: () => handleUnbindById(row),
        // }, {
        //   default: () => computed(() => t('userActivation.cancelBind')).value,
        // }),
      ]);
    },

  }
];

let callGetAllActivateLog = async () => {
  // TODO: fix fetch keys
  // let data = await handleGetAllActivationLog(dataSize.value.page, dataSize.value.pageSize)
  // if (data.code === 200) {
  //   activateRecordList.value = []
  //   data.log.forEach((log: ActivateRecord) => activateRecordList.value.push(log))
  //   pageCount.value = data.page_count
  //   animated.value = true
  // }
}

const callGetAllKeys = async () => {
  let data = await handleGetAllKeysByAdmin(dataSize.value.page, dataSize.value.pageSize, paramsKeyId)
  if (data && data.code === 200) {
    console.log(data)
    if (data.keys) {
      keyList.value = []
      data.keys.forEach((i: KeyItem) => keyList.value.push(i))
      pageCount.value = data.page_count
    } else {
      pageCount.value = 0
    }
  } else {
    message.error('err' + data.msg || '')
    pageCount.value = 0
  }
  animated.value = true

}


// let handleUnbindById = async (row: ActivateRecord) => {
//   animated.value = false
//   try {
//     let {data} = await instance.delete('/api/user/v1/activation', {
//       params: {
//         user_id: userInfoStore.thisUser.id,
//         activation_id: row.id,
//       }
//     })
//     if (data.code === 200) {
//       await callGetAllActivateLog()
//     }
//
//   } catch (error: any) {
//     console.log(error)
//     message.error("unknown err: " + error)
//   }
// }

let enableAlterRemark = ref<boolean>(false)

// let handleCommitNewRemark = async () => {
//   enableAlterRemark.value = false
//   try {
//     let {data} = await instance.patch('/api/user/v1/activation/remark', {
//       user_id: userInfoStore.thisUser.id,
//       record_id: currentRecord.value.id,
//       remark: currentRecord.value.remark,
//     })
//     if (data.code === 200) {
//       message.success(t('userActivation.updateSuccess'))
//     } else {
//       message.error('err: ' + data.msg || '')
//     }
//   } catch (err: any) {
//     console.log(err)
//   }
// }


onBeforeMount(() => {
  themeStore.userPath = '/dashboard/key'
  themeStore.menuSelected = 'key-manager'
})

onMounted(async () => {
  // await callGetAllActivateLog()

  await callGetAllKeys()
  animated.value = true
})

</script>

<script lang="ts">
export default {
  name: "KeyMgr",
}
</script>

<template>
  <div class="root-header">
    <PageHead
        :title="t('adminViews.keysMgr.keyMgr')"
        :description="t('adminViews.activation.description')"
    >
      <n-button
          tertiary
          type="primary"
          size="medium"
          class="btn-right"
          @click="showSearchModal=true">
        <template #icon>
          <n-icon>
            <searchIcon/>
          </n-icon>
        </template>
        {{ t('adminViews.userMgr.query') }}
      </n-button>
      <n-button
          tertiary
          type="primary"
          size="medium"
          class="btn-right"
          @click="Object.assign(searchForm, {KeyId: 0, keyContent: ''}); callGetAllKeys()">
        <template #icon>
          <n-icon>
            <refreshIcon/>
          </n-icon>
        </template>
        {{ t('adminViews.orderMgr.resetSearch') }}
      </n-button>
    </PageHead>
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
            :data="keyList"
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
            @update:page="animated=false; callGetAllActivateLog()"
        />
        <n-select
            style="width: 200px; margin-left: 20px"
            v-model:value.number="dataSize.pageSize"
            size="small"
            :options="dataCountOptions"
            :remote="true"
            @update:value="animated=false; dataSize.page = 1; callGetAllActivateLog()"
        />
      </div>
    </div>
  </transition>

  <n-modal
      :title="t('adminViews.activation.search')"
      v-model:show="showSearchModal"
      preset="dialog"
      :positive-text="t('adminViews.userMgr.query')"
      :negative-text="t('adminViews.userMgr.cancel')"
      style="width: 480px"
      @positive-click="animated=false; callGetAllKeys()"
      @negative-click="showSearchModal=false;"
      @before-leave="Object.assign(searchForm, {KeyId: 0, keyContent: ''})"
      :show-icon="false"
  >
    <div style="margin-top: 30px"></div>
    <!--    <n-form-item path="email" :label="t('adminViews.userMgr.userEmail')">-->
    <!--      <n-input-->
    <!--          autofocus-->
    <!--          @keyup.enter="showSearchModal=false;animated=false; callGetAllActivateLog()"-->
    <!--          :placeholder="t('adminViews.userMgr.inputUserEmail')"-->
    <!--          v-model:value="searchForm.email"-->
    <!--      ></n-input>-->
    <!--    </n-form-item>-->

    <n-form-item path="valid" :label="t('adminViews.keysMgr.searchModal.searchMethod')">
      <n-select
          :options="[{label: t('adminViews.keysMgr.searchModal.byId'), value: 'id'},{label: t('adminViews.keysMgr.searchModal.byContent'), value: 'content'}]"
          :default-value="false"
          v-model:value="searchMethod"
      >
      </n-select>
    </n-form-item>

    <n-form-item v-if="searchMethod==='id'" path="id" :label="t('adminViews.keysMgr.searchModal.keyId')">
      <n-input
          autofocus
          @keyup.enter="showSearchModal=false;animated=false; callGetAllKeys()"
          :placeholder="t('adminViews.keysMgr.searchModal.idPlaceholder')"
          v-model:value="searchForm.keyId"
      ></n-input>
    </n-form-item>
    <n-form-item v-if="searchMethod==='content'" path="content" :label="t('adminViews.keysMgr.searchModal.keyContent')">
      <n-input
          autofocus
          @keyup.enter="showSearchModal=false;animated=false; callGetAllKeys()"
          :placeholder="t('adminViews.keysMgr.searchModal.contentPlaceholder')"
          v-model:value="searchForm.keyContent"
      ></n-input>
    </n-form-item>

  </n-modal>

</template>

<style scoped lang="less">
.root-header {
  padding: 0;
}

.root {
  padding: 0 0 20px 20px;
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

.show-key:hover {
  text-decoration: underline;
}

.btn-right {
  margin-right: 20px;
}
</style>