<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, h, onBeforeMount, onMounted, ref} from "vue";
import {NButton, NIcon, NTag, useMessage, type DataTableColumns, useDialog} from "naive-ui"
import useThemeStore from "@/stores/useThemeStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
// import instance from "@/axios";
import {useRoute, useRouter} from "vue-router";
import {handleBlockKeyById, handleGetAllKeysByAdmin} from "@/api/admin/keys";
import PageHead from "@/views/utils/PageHead.vue";
import {RefreshOutline as refreshIcon, Search as searchIcon,} from "@vicons/ionicons5"
import DataTableSuffix from "@/views/utils/DataTableSuffix.vue";
import {formatDate} from "@/utils/timeFormat";
import useTablePagination from "@/hooks/useTablePagination";
// import {handleDeleteCouponById} from "@/api/admin/coupon";

interface KeyItem {
  key_id: number;
  user_id: number;
  client_id: string
  is_valid: boolean;
  is_used: boolean;
  key: string;
  order_id: string;
  plan_name: string;
  email: string;
  created_at: string;
  expiration_date: string;
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
const nowDate = new Date().getTime();
let showSearchModal = ref<boolean>(false)
let showDetailModal = ref<boolean>(false)
let detailKeyIndex = ref<number>(0)

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
  console.error('key_id is missing or invalid, skip');
  paramsKeyId = 0
}

const message = useMessage();
const dialog = useDialog();
const themeStore = useThemeStore()
const userInfoStore = useUserInfoStore()

let animated = ref<boolean>(false);

const [dataSize, pageCount] = useTablePagination()

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

const columns = computed<DataTableColumns<KeyItem>>(() => [
  {
    title: t('userActivation.id'),
    key: 'key_id',
  },
  {
    title: t('userActivation.orderId'),
    key: 'order_id',
  },
  {
    title: t('adminViews.keysMgr.table.email'),
    key: 'email',
  },
  {
    title: t('adminViews.keysMgr.table.key'),
    key: 'key',
    render(row: KeyItem) {
      return h(NTag, {size: 'small', bordered: false, type: row.is_valid && row.expiration_date_stamp >= nowDate?row.is_used?'success':'primary':'error'}, {default: () => row.key})
    },
  },
  {
    title: t('adminViews.keysMgr.table.isValid'),
    key: 'os_type',
    render(row: KeyItem) {
      return h(NTag, {
        size: 'small',
        bordered: false,
        type: row.is_valid ? 'success' : 'error'
      }, {default: () => row.is_valid ? t('adminViews.keysMgr.table.ok') : t('adminViews.keysMgr.table.blocked')});
    }
  },
  {
    title: t('adminViews.keysMgr.table.isExpired'),
    key: 'os_type',
    render(row: KeyItem) {
      return h(NTag, {
        size: 'small',
        bordered: false,
        type: row.expiration_date_stamp >= nowDate? 'success' : 'default',
        style: {textDecoration: row.expiration_date_stamp >= nowDate?'none':'line-through'}
      }, {default: () => {
          // row.is_used ? t('adminViews.keysMgr.table.used') : t('adminViews.keysMgr.table.unUsed')
          return row.expiration_date_stamp >= nowDate?t('adminViews.keysMgr.table.active'):t('adminViews.keysMgr.table.inactive')
        }});
    }
  },
  {
    title: t('adminViews.keysMgr.table.isUsed'),
    key: 'os_type',
    render(row: KeyItem) {
      return h(NTag, {
        size: 'small',
        bordered: false,
        type: row.is_used ? 'success' : 'default'
      }, {default: () => row.is_used ? t('adminViews.keysMgr.table.used') : t('adminViews.keysMgr.table.unUsed')});
    }
  },
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
          onClick: () => {
            // 查找 keyList 中与 row.key_id 相等的元素下标
            const index = keyList.value.findIndex((item: KeyItem) => item.key_id === row.key_id);
            // 如果找到了匹配的项，则设置 detailKeyIndex
            if (index !== -1) {
              detailKeyIndex.value = index;
              showDetailModal.value = true;
            } else {
              message.error('err')
            }
            // console.log(keyList.value[detailKeyIndex.value])
          },
        }, {
          default: () => t('adminViews.keysMgr.table.showDetail'),
        }),
        h(NButton, {
          size: 'small',
          secondary: true,
          type: row.is_valid?'warning' : 'success',
          style: {marginLeft: '10px'},
          bordered: false,
          onClick: () => callBlockKeyByKeyId(row),
        }, {
          default: () => row.is_valid?t('adminViews.keysMgr.table.blockKey'):t('adminViews.keysMgr.table.unblockKey')
        }),
      ]);
    },

  }
])

const callGetAllKeys = async () => {
  let data = await handleGetAllKeysByAdmin(
      dataSize.value.page,
      dataSize.value.pageSize,
      searchForm.value.keyId > 0 ? searchForm.value.keyId : paramsKeyId > 0 ? paramsKeyId : 0,
      searchForm.value.keyContent || '',
  )
  if (data && data.code === 200) {
    keyList.value = []
    console.log(data)
    if (data.keys) {
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

const callBlockKeyByKeyId = async (row: KeyItem) => {
  const runTask = async () => {
    let data = await handleBlockKeyById(row.user_id, row.key_id)
    if (data && data.code === 200) {
      message.success(t('adminViews.common.success'))
      await callGetAllKeys()
    } else {
      // TODO
      message.error(t('adminViews.common.failure'))
    }
  }
  if (!row.is_valid) return runTask()
  else dialog.warning({
    title: t('adminViews.keysMgr.mention.title'),
    content: () => {
      return h('div', {}, [
        h('p', {style: {fontWeight: 'bold', fontSize: '1rem', opacity: '0.8'}}, row.key),
        h('p', {style: {marginTop: '4px'}}, t('adminViews.keysMgr.mention.content'))
      ])
    },
    positiveText: t('adminViews.common.confirm'),
    negativeText: t('adminViews.common.cancel'),
    showIcon: false,
    actionStyle: {
      marginTop: '20px',
    },
    contentStyle: {
      marginTop: '20px',
    },
    onPositiveClick: async () => {
      await runTask()
    }
  })
}

type KeyDetailKey = 'user_id' | 'plan_name' | 'expiration_date' | 'created_at' | 'client_id'

interface KeyDetail {
  title: string;
  key: KeyDetailKey
}

const keyDetailItems: KeyDetail[] = [
  {
    title: 'adminViews.keysMgr.detailModal.userId',
    key: 'user_id',
  },
  {
    title: 'adminViews.keysMgr.detailModal.planName',
    key: 'plan_name',
  },
  {
    title: 'adminViews.keysMgr.detailModal.clientId',
    key: 'client_id',
  },
  {
    title: 'adminViews.keysMgr.detailModal.expiredAt',
    key: 'expiration_date',
  },
  {
    title: 'adminViews.keysMgr.detailModal.keyGeneratedAt',
    key: 'created_at',
  },
]

const formatDetailString = (key: string): string => {
  if (key === 'expiration_date' || key === 'created_at') {
    return formatDate(keyList.value[detailKeyIndex.value][key as KeyDetailKey] as string)
  } else if (key === 'user_id') {
    return `# ${keyList.value[detailKeyIndex.value][key as KeyDetailKey].toString()}`
  } else {
    return keyList.value[detailKeyIndex.value][key as KeyDetailKey].toString()
  }
}

onBeforeMount(() => {
  themeStore.menuSelected = 'key-manager'
  themeStore.breadcrumb = 'adminViews.keysMgr.keyMgr'
})

onMounted(async () => {
  // await callGetAllActivateLog()
  themeStore.userPath = '/dashboard/key'

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
          @click="Object.assign(searchForm, {keyId: 0, keyContent: ''}); animated=false ;callGetAllKeys()">
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
            :scroll-x="1200"
        />
      </n-card>
     <DataTableSuffix
         :data-size="dataSize"
         :page-count="pageCount"
         :animated="animated"
         :update-data="callGetAllKeys"
     />
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

  <n-modal
      v-model:show="showDetailModal"
      class="custom-card"
      preset="card"
      style="width: 600px"
      :style="{paddingBottom: '30px'}"
      :title="t('adminViews.keysMgr.detailModal.title')"
      size="medium"
      :bordered="false"
  >
    <div class="details-root" v-for="i in keyDetailItems" :key="i.title">
      <p class="detail-title">{{ t(i.title || '') }}</p>
      <p class="detail-content" :style="i.title==='adminViews.keysMgr.detailModal.clientId'?{fontSize: '0.9rem',}:null">{{ formatDetailString(i.key) }}</p>
    </div>

  </n-modal>


</template>

<style scoped lang="less">
.root-header {
  padding: 0;
}

.root {
  padding: 0 20px 20px 20px;
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
  margin-right: 10px;
}

.details-root {
  margin-top: 16px;
  .detail-title {
    font-size: 1rem;
    font-weight: bold;
    opacity: 0.7;
  }
  .detail-content {
    font-size: 1.1rem;
    margin-top: 4px;
  }
}
</style>