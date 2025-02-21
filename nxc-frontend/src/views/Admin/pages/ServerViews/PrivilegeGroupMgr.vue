<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, h, onBeforeMount, onMounted, ref} from 'vue'
import {
  type DataTableColumns,
  type FormInst,
  NButton,
  NIcon,
  type NotificationType,
  useMessage,
  useNotification
} from "naive-ui";
import {AddOutline as AddIcon} from "@vicons/ionicons5"
import useThemeStore from "@/stores/useThemeStore";
import instance from "@/axios";
import {BarChartOutlined, UserOutlined} from '@vicons/antd'
import DataTableSuffix from "@/views/utils/DataTableSuffix.vue";
import PageHead from "@/views/utils/PageHead.vue"; // 引入所需圖標

const {t} = useI18n()
const message = useMessage()

let modifyFunc = ref<string>('add')
let modifyId = ref<number | null>(null)
let pageCount = ref(10)
let showLoading = ref<boolean>(false)
let animated = ref<boolean>(false)
let dataSize = ref<{ pageSize: number, page: number }>({
  pageSize: 10,
  page: 1,
})

const formRef = ref<FormInst | null>(null)
const notification = useNotification();
const themeStore = useThemeStore()

interface PrivilegeGroup {
  id: number
  name: string
  plan_count: number
  user_count: number
}

interface GroupItem {
  id: number
  name: string
  user_count?: number
  plan_count?: number
  updated_at?: string
  created_at?: string
  deleted_at?: string
}

const groupList = ref<GroupItem[]>([])

// 定義表格列
const columns = computed<DataTableColumns<GroupItem>>(() =>
    [
      {title: t('adminViews.groupMgr.groupId'), key: 'id'},
      {title: t('adminViews.groupMgr.groupName'), key: 'name'},
      {
        title: t('adminViews.groupMgr.userCount'),
        key: 'user_count',
        render(row: GroupItem) {
          return h('div', {style: {display: 'flex', flexDirection: 'row',}}, [
            h(NIcon, {size: 18, style: {marginRight: 5}}, {default: () => h(UserOutlined)}), // 用戶圖標
            `${row.hasOwnProperty('user_count') ? row.user_count : 0}`,
          ])
        }
      },
      {
        title: t('adminViews.groupMgr.planCount'),
        key: 'plan_count',
        render(row: GroupItem) {
          return h('div', {style: {display: 'flex', flexDirection: 'row'}}, [
            h(NIcon, {size: 18, style: {marginRight: 5}}, {default: () => h(BarChartOutlined)}), // 計畫圖標
            `${row.hasOwnProperty('plan_count') ? row.plan_count : 0}`,
          ])
        }
      },
      {
        title: t('adminViews.groupMgr.operate'),
        key: 'actions',
        fixed: 'right',
        render(row: GroupItem) {
          return h('div', {style: 'display: flex; gap: 8px'}, [
            h(NButton, {
              size: 'small',
              type: 'primary',
              secondary: true,
              onClick: () => editGroup(row)
            }, {default: () => t('adminViews.groupMgr.editBtnHint')}),
            h(NButton, {
              size: 'small',
              type: 'error',
              secondary: true,
              onClick: () => deleteGroup(row.id)
            }, {default: () => t('adminViews.groupMgr.deleteBtnHint')})
          ])
        }
      }
    ]
)

// 編輯組操作
const editGroup = (row: GroupItem) => {
  // message.info(`編輯權限組：${id}`)
  modifyFunc.value = 'edit'
  newGroupName.value = row.name
  showModal.value = true
  modifyId.value = row.id
}

// 刪除組操作
const deleteGroup = async (id: number) => {
  animated.value = false
  let {data} = await instance.delete('http://localhost:8081/api/admin/v1/groups', {
    data: {
      id,
    }
  })
  if (data.code === 200) {
    message.success(t('adminViews.groupMgr.common.delSuccess'))
    await getAllGroups()
  }
}

let showModal = ref(false)

let formValue = ref({
  privilege_group: {
    name: '',
  }
})

let rules = {
  privilege_group: {
    required: true,
    message: '权限组名称是必填的',
    trigger: 'blur'
  }
}


let newGroupName = ref<string>('')

let notify = (type: NotificationType, title: string, meta?: string) => {
  notification[type]({
    content: title,
    meta: meta,
    duration: 2500,
    keepAliveOnHover: true
  })
}

let handleAddGroup = async () => {
  modifyFunc.value = 'add'
  newGroupName.value = ''
  showModal.value = true
}

let submitNewGroup = async () => {
  try {
    let {data} = await instance.post('http://localhost:8081/api/admin/v1/groups', {
      group_name: newGroupName.value
    })
    if (data.code === 200) {
      message.success(t('adminViews.groupMgr.common.addSuccess'))
    } else {
      message.error(t('adminViews.groupMgr.common.addFailure') + data.msg || '')
    }
  } catch (error) {
    console.log(error)
  }
}

let closeModal = () => {
  showModal.value = false
}

// let groupList = ref<PrivilegeGroup>([])

let getAllGroups = async () => {
  showLoading.value = true
  try {
    let {data} = await instance.get('http://localhost:8081/api/admin/v1/groups', {
      params: {
        page: dataSize.value.page,
        size: dataSize.value.pageSize,
      }
    })
    if (data.code === 200) {
      groupList.value = []
      showLoading.value = false
      data.group_list.forEach((group: PrivilegeGroup) => groupList.value.push(group))
      pageCount.value = data.page_count
      // groupList.value = data.group_list
      animated.value = true
    }
  } catch (error) {
    console.log(error)
  }
}

let submit = async () => {
  if (modifyFunc.value === 'add') {
    await submitNewGroup()
  } else {
    await submitUpdate()
  }
  await getAllGroups()
}

let submitUpdate = async () => {
  let {data} = await instance.put('/api/admin/v1/groups', {
    id: modifyId.value,
    name: newGroupName.value,
  })
  if (data.code === 200) {
    message.success(t('adminViews.groupMgr.common.delSuccess'))
  } else {
    message.error(t('adminViews.groupMgr.common.delFailure'))
  }
}

onBeforeMount(() => {
  themeStore.breadcrumb = t('adminViews.groupMgr.title')
  themeStore.menuSelected = 'privilege-group-mgr'
})

onMounted(async () => {
  themeStore.contentPath = '/admin/dashboard/group'
  await getAllGroups()
  animated.value = true
})

</script>

<script lang="ts">
export default {
  name: 'PrivilegeGroupMgr',
}
</script>

<template>
  <!--  <div style="padding: 20px 20px 0 20px">-->
  <!--    <n-card hoverable :embedded="true" :title="t('adminViews.groupMgr.title')" :bordered="false">-->
  <!--      <n-button type="primary" :bordered="false" class="add-btn" @click="handleAddGroup">{{ t('adminViews.groupMgr.common.addNewGroup') }}</n-button>-->
  <!--    </n-card>-->
  <!--  </div>-->

  <PageHead
      :title="t('adminViews.groupMgr.title')"
      :description="t('adminViews.groupMgr.description')"
  >
    <template #default>
      <n-button
          tertiary
          type="primary"
          size="medium"
          class="btn-right"
          @click="handleAddGroup"
      >
        {{ t('adminViews.groupMgr.common.addNewGroup') }}
        <template #icon>
          <n-icon>
            <AddIcon/>
          </n-icon>
        </template>
      </n-button>
    </template>
  </PageHead>

  <!--  <div style="padding: 20px 20px 0 20px">-->

  <!--    <n-page-header :subtitle="t('adminViews.groupMgr.description')" @back="useRouter().back()">-->
  <!--      <template #title>-->
  <!--        <p style="font-size: 1.4rem; font-weight: bold; margin-bottom: 10px;">{{ t('adminViews.groupMgr.title') }}</p>-->
  <!--      </template>-->
  <!--&lt;!&ndash;      <template #avatar>&ndash;&gt;-->
  <!--&lt;!&ndash;        <n-icon><privilegeIcon /></n-icon>&ndash;&gt;-->
  <!--&lt;!&ndash;      </template>&ndash;&gt;-->

  <!--    </n-page-header>-->

  <!--    <n-flex vertical>-->
  <!--      <div>-->
  <!--              <n-tag-->
  <!--                  type="default"-->
  <!--                  :bordered="false"-->
  <!--                  checkable-->
  <!--                  style="font-size: 1.4rem; font-weight: bold; margin-bottom: 10px; width: auto"-->
  <!--              >{{ t('adminViews.groupMgr.title') }}-->
  <!--                <template #icon>-->
  <!--                  <n-icon><privilegeIcon /></n-icon>-->
  <!--                </template>-->
  <!--              </n-tag>-->
  <!--      </div>-->
  <!--    </n-flex>-->

  <!--    <n-flex vertical>-->
  <!--      <n-tag-->
  <!--          type="default"-->
  <!--          :bordered="false"-->
  <!--          checkable-->
  <!--          style="font-size: 1.4rem; font-weight: bold; margin-bottom: 10px; width: auto"-->
  <!--      >{{ t('adminViews.groupMgr.title') }}-->
  <!--        <template #icon>-->
  <!--          <n-icon><privilegeIcon /></n-icon>-->
  <!--        </template>-->
  <!--      </n-tag>-->
  <!--      <n-tag-->
  <!--          type="primary"-->
  <!--          checkable-->
  <!--          disabled-->
  <!--          style="font-size: 0.8rem; opacity: 0.8">-->
  <!--        > {{ t('adminViews.groupMgr.description') }}-->
  <!--      </n-tag>-->
  <!--    </n-flex>-->

  <!--  </div>-->

  <transition name="slide-fade">
    <div class="root" v-if="animated">
      <n-card style="margin-top: 20px" :embedded="true" hoverable :bordered="false" content-style="padding: 0;">
        <!--      在這裏放置表格-->
        <!--      表格有五列 列名分別是 權限組ID， 組名稱， 用戶數量（前面使用n-icon放置一個用戶圖標）， 計畫數量（前面使用n-icon放置一個圖標， 操作（內部為兩個按鈕 水平排列分別為編輯和刪除）-->
        <n-spin :show="showLoading">
          <n-data-table
              striped
              :columns="columns"
              :data="groupList"
              :scroll-x="900"
              :remote="true"
          />
        </n-spin>
      </n-card>

      <DataTableSuffix
          v-model:data-size="dataSize"
          v-model:page-count="pageCount"
          v-model:animated="animated"
          :update-data="getAllGroups"
      />

    </div>
  </transition>

  <n-modal
      :title="modifyFunc==='add'?t('adminViews.groupMgr.common.addNewGroup'):t('adminViews.groupMgr.common.alterGroupName')"
      v-model:show="showModal"
      preset="dialog"
      :positive-text="modifyFunc==='add'?t('adminViews.groupMgr.common.addConfirmed'):t('adminViews.groupMgr.common.alterConfirmed')"
      :negative-text="t('adminViews.groupMgr.common.cancel')"
      style="width: 480px; border: 0"
      @positive-click="submit"
      @negative-click="closeModal"
      :show-icon="false"
  >
    <!--          pass-->
    <div style="margin-top: 30px"></div>

    <n-form
        ref="formRef"
        :model="formValue"
        :rules="rules"
    >
      <n-form-item :label="t('adminViews.groupMgr.groupName')" path="privilege_group.name">
        <n-input v-model:value="newGroupName" :placeholder="t('adminViews.groupMgr.groupPlaceholder')"/>
      </n-form-item>
    </n-form>

  </n-modal>
</template>

<style lang="less" scoped>
.root {
  //padding: 20px;
  padding: 0 20px;

  //.add-btn {
  //  margin-top: 10px;
  //}
  //
  //.btn-pagination {
  //  margin-top: 20px;
  //  display: flex;
  //  flex-direction: row;
  //  justify-content: right;
  //}
}

</style>