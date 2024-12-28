<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, h, onBeforeMount, onMounted, ref} from "vue";
import {NButton, NTag, NPopover, useMessage, NIcon} from "naive-ui"
import useThemeStore from "@/stores/useThemeStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import {formatDate} from "@/utils/timeFormat";
import instance from "@/axios";
import {useRouter} from "vue-router";
import {
  ChevronDownOutline as downIcon,
  PauseOutline as closeOrderIcon,
  CheckmarkDoneOutline as passOrderIcon,
} from "@vicons/ionicons5"
// import {
//   ChevronForwardOutline as toRight,
//   CheckmarkOutline as checkIcon,
//   SaveOutline as saveIcon,
// } from "@vicons/ionicons5"
import {RefreshOutline as refreshIcon, Search as searchIcon} from '@vicons/ionicons5'
import {handleGetAllActivationLog} from "@/api/admin/activation";

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
    title: computed(() => t('adminViews.activation.email')).value,
    key: 'email',
  },
  {
    title: computed(() => t('adminViews.activation.key')).value,
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
                  onMouseover(e: MouseEvent) {
                    const target = e.target as HTMLElement; // 使用类型断言，告诉 TypeScript target 是 HTMLElement
                    target.style.textDecoration = 'underline';
                    target.style.fontWeight = 'bold';
                  },
                  onMouseout(e: MouseEvent) {
                    const target = e.target as HTMLElement;
                    target.style.textDecoration = 'none';
                    target.style.fontWeight = 'normal';
                    Object.assign(keyDetails.value, {id: 0, key: '', created_at: ''});
                  },
                  onClick: (e: MouseEvent) => {
                    e.stopPropagation(); // 阻止事件冒泡
                    showDetails(row);  // 触发点击事件
                  },
                },
                { default: () => computed(() => t('adminViews.activation.showKey')).value },
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
              return h('div', {}, elements); // 返回包含条件渲染元素的 div
            },
          }
      );
    },
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
    render(row: ActivateRecord) {
      return h('div', {style: {display: 'flex', flexDirection: 'row'}}, [
        h(NButton, {
          size: 'small',
          secondary: true,
          type: 'primary',
          bordered: false,
          onClick: () => showDetails(row),
        }, {
          default: () => computed(() => t('adminViews.activation.turn2keyPage')).value,
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
  let data = await handleGetAllActivationLog(
      dataSize.value.page,
      dataSize.value.pageSize,
      searchForm.value.email.trim(),
      searchForm.value.valid,
      searchForm.value.sort
  );
  if (data.code === 200) {
    activateRecordList.value = []
    data.log.forEach((log: ActivateRecord) => activateRecordList.value.push(log))
    pageCount.value = data.page_count
    animated.value = true
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
      await callGetAllActivateLog()
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
  <div class="root-header">
<!--    <n-card-->
<!--        hoverable-->
<!--        :embedded="true"-->
<!--        :bordered="false"-->
<!--        :title="t('adminViews.activation.activateLog')"-->
<!--    ></n-card>-->
    <n-card
        hoverable
        :bordered="false"
        :embedded="true"
        class="card1"
        :title="t('adminViews.activation.activateLog')"
    >
      <n-button class="btn" secondary type="primary" size="medium" @click="showSearchModal=true">
        <n-icon size="14" style="padding-right: 8px">
          <searchIcon/>
        </n-icon>
        {{ t('adminViews.userMgr.query') }}
      </n-button>
      <n-button class="btn" tertiary type="primary" size="medium"
                @click="searchForm.email=''; animated=false; callGetAllActivateLog()">
        <n-icon size="14" style="padding-right: 8px">
          <refreshIcon/>
        </n-icon>
        {{ '重置搜索' }}
      </n-button>
    </n-card>
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

.show-key:hover {
  text-decoration: underline;
}

.btn {
  margin-right: 10px;
}
</style>