<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, h, onBeforeMount, onMounted, ref} from "vue";
import {NButton, NIcon, NPopover, NTag, useMessage, type DataTableColumns} from "naive-ui"
import useThemeStore from "@/stores/useThemeStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import {formatDate} from "@/utils/timeFormat";
import {useRouter} from "vue-router";
import {RefreshOutline as refreshIcon, Search as searchIcon,} from "@vicons/ionicons5"

import {handleGetAllActivationLog, handleShowKeyDetails} from "@/api/admin/activation";
import DataTableSuffix from "@/views/utils/DataTableSuffix.vue";
import PageHead from "@/views/utils/PageHead.vue";
import useTablePagination from "@/hooks/useTablePagination";

interface ActivateRecord {
  id: number;
  user_id: number;
  email: string;
  order_id: string;
  key?: string;
  key_id: number;
  request_at: string;
  client_version: string;
  os_type: string;
  remark: string;
  is_bind: boolean;
  created_at: string;
  updated_at: string;
}

let showSearchModal = ref<boolean>(false)

let searchForm = ref<{
  email: string
  valid: boolean
  sort: 'ASC' | 'DESC'
}>({
  email: '',
  valid: false,
  sort: 'DESC',   // 默认按照created_at降序排序
})


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

let showDetails = async (row: ActivateRecord) => {
  let data = await handleShowKeyDetails(row.key_id)
  if (data.code === 200) {
    Object.assign(currentRecord.value, row)
    Object.assign(keyDetails.value, data.details)
  } else {
    message.error(t('adminViews.common.fetchDataFailure') + data.msg || '')
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
            type: row.is_bind ? 'success' : 'error',
          },
          {
            default: () => row.is_bind ? t('userActivation.active') : t('userActivation.inactive')
          });
    }
  },
  {
    title: t('userActivation.orderId'),
    key: 'order_id',
  },
  {
    title: t('adminViews.activation.email'),
    key: 'email',
  },
  {
    title: t('adminViews.activation.key'),
    key: 'key',
    render(row: ActivateRecord) {
      return h(
          NPopover,
          {
            trigger: 'hover',
            showArrow: false,
          },
          {
            // 这里定义了 'trigger' 插槽内容
            trigger: () => h(
                NTag,
                {
                  size: 'small',
                  bordered: false,
                  type: 'primary',
                  style: {
                    textDecoration: 'none', // 默认不显示下划线
                    cursor: 'pointer',      // 鼠标悬浮时指针变为手型
                  },
                  onMouseover: (e: MouseEvent) => {
                    const target = e.target as HTMLElement;
                    target.style.textDecoration = 'underline';
                    target.style.fontWeight = 'bold';
                  },
                  onMouseout: (e: MouseEvent) => {
                    const target = e.target as HTMLElement;
                    target.style.textDecoration = 'none';
                    target.style.fontWeight = 'normal';
                    try {
                      Object.assign(keyDetails.value, {id: 0, key: '', created_at: ''});
                    } catch (error) {
                      console.error("Error assigning to keyDetails.value:", error);
                    }
                  },
                  onClick: (e: MouseEvent) => {
                    e.stopPropagation();
                    try {
                      showDetails(row);
                    } catch (error) {
                      console.error("Error calling showDetails:", error);
                    }
                  },
                },
                {default: () => t('adminViews.activation.showKey')},
            ),
            default: () => {
              const elements = [];
              if (keyDetails.value.key !== "") {
                elements.push(h('p', {
                  style: {
                    textDecoration: 'underline',
                  }
                }, `${keyDetails.value.key} (#${keyDetails.value.id})`));
                elements.push(h('p', {
                  style: {opacity: 0.6,}
                }, `${formatDate(keyDetails.value.created_at)} (${t('adminViews.activation.createdAt')})`));
              } else {
                elements.push(h('p', null, t('adminViews.activation.click2getKey')));
              }
              return h('div', {}, elements);
            },
          }
      );
    },
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
          onClick: () => router.push({name: 'adminKeysMgr', params: {key_id: row.key_id}}),
        }, {
          default: () => t('adminViews.activation.turn2keyPage'),
        }),
        // h(NButton, {
        //   size: 'small',
        //   secondary: true,
        //   type: 'warning',
        //   style: {marginLeft: '10px'},
        //   bordered: false,
        //   disabled:!row.is_bind,
        //   onClick: () => handleUnbindById(row),
        // }, {
        //   default: () => t('userActivation.cancelBind'),
        // }),
      ]);
    },

  }
]);

const callGetAllActivateLog = async () => {
  let data = await handleGetAllActivationLog(
      dataSize.value.page,
      dataSize.value.pageSize,
      searchForm.value.email.trim(),
      searchForm.value.valid,
      searchForm.value.sort
  );
  if (data && data.code === 200) {
    activateRecordList.value = []
    if (data.log)
      data.log.forEach((log: ActivateRecord) => activateRecordList.value.push(log))
    pageCount.value = data.page_count
  } else {
    message.error(t('adminViews.common.fetchDataFailure') + data.msg || '')
  }
  animated.value = true
  
}

// let enableAlterRemark = ref<boolean>(false)

onBeforeMount(() => {
  themeStore.breadcrumb = t('adminViews.activation.activateLog')
  themeStore.userPath = '/dashboard/activation'
  themeStore.menuSelected = 'activate-manager'
})

onMounted(async () => {
  await callGetAllActivateLog()
  animated.value = true
})

</script>

<script lang="ts">
export default {
  name: "ActivationMgr",
}
</script>

<template>
  <PageHead
      :title="t('adminViews.activation.activateLog')"
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
        @click="searchForm.email=''; animated=false; callGetAllActivateLog()">
      <template #icon>
        <n-icon>
          <refreshIcon/>
        </n-icon>
      </template>
      {{ t('adminViews.orderMgr.resetSearch') }}
    </n-button>
  </PageHead>

  <!--  <div class="root-header">-->
  <!--    &lt;!&ndash;    <n-card&ndash;&gt;-->
  <!--    &lt;!&ndash;        hoverable&ndash;&gt;-->
  <!--    &lt;!&ndash;        :embedded="true"&ndash;&gt;-->
  <!--    &lt;!&ndash;        :bordered="false"&ndash;&gt;-->
  <!--    &lt;!&ndash;        :title="t('adminViews.activation.activateLog')"&ndash;&gt;-->
  <!--    &lt;!&ndash;    ></n-card>&ndash;&gt;-->
  <!--    <n-card-->
  <!--        hoverable-->
  <!--        :bordered="false"-->
  <!--        :embedded="true"-->
  <!--        class="card1"-->
  <!--        :title="t('adminViews.activation.activateLog')"-->
  <!--    >-->
  <!--      <n-button class="btn" secondary type="primary" size="medium" @click="showSearchModal=true">-->
  <!--        <n-icon size="14" style="padding-right: 8px">-->
  <!--          <searchIcon/>-->
  <!--        </n-icon>-->
  <!--        {{ t('adminViews.userMgr.query') }}-->
  <!--      </n-button>-->
  <!--      <n-button class="btn" tertiary type="primary" size="medium"-->
  <!--                @click="searchForm.email=''; animated=false; callGetAllActivateLog()">-->
  <!--        <n-icon size="14" style="padding-right: 8px">-->
  <!--          <refreshIcon/>-->
  <!--        </n-icon>-->
  <!--        {{ '重置搜索' }}-->
  <!--      </n-button>-->
  <!--    </n-card>-->
  <!--  </div>-->
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
            :scroll-x="1200"
        />
      </n-card>
      <DataTableSuffix
          v-model:data-size="dataSize"
          v-model:page-count="pageCount"
          v-model:animated="animated"
          :update-data="callGetAllActivateLog"
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
      @positive-click="animated=false; callGetAllActivateLog()"
      @negative-click="showSearchModal=false;"
      @before-leave="searchForm.email=''"
      :show-icon="false"
  >
    <div style="margin-top: 30px"></div>
    <n-form-item path="email" :label="t('adminViews.userMgr.userEmail')">
      <n-input
          autofocus
          @keyup.enter="showSearchModal=false;animated=false; callGetAllActivateLog()"
          :placeholder="t('adminViews.userMgr.inputUserEmail')"
          v-model:value="searchForm.email"
      ></n-input>
    </n-form-item>

    <n-form-item path="valid" :label="t('adminViews.activation.filter')">
      <n-select
          :options="[{label: t('adminViews.activation.filterAll'), value: false},{label: t('adminViews.activation.filterActive'), value: true}]"
          :default-value="false"
          v-model:value="searchForm.valid"
      >
      </n-select>
    </n-form-item>

    <n-form-item path="sort" :label="t('adminViews.activation.sortAlgorithm')">
      <n-select
          :options="[{label: t('adminViews.activation.sortASC'), value: 'ASC'},{label: t('adminViews.activation.sortDESC'), value: 'DESC'}]"
          :default-value="'DESC'"
          v-model:value="searchForm.sort"
      >
      </n-select>
    </n-form-item>

  </n-modal>

</template>

<style scoped lang="less">
.root-header {
  padding: 0 20px 0 20px;
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

.btn {
  margin-right: 10px;
}

.btn-right {
  margin-right: 10px;
}

</style>