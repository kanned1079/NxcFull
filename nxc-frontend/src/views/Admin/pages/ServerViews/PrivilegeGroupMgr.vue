<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {computed, h, onBeforeMount, onMounted, ref} from 'vue'
import {
  type DataTableColumns,
  type FormInst,
  NButton,
  NIcon,
  useMessage,
  useNotification,
  useDialog
} from "naive-ui";
import {AddOutline as AddIcon} from "@vicons/ionicons5"
import useThemeStore from "@/stores/useThemeStore";
import {BarChartOutlined, UserOutlined} from '@vicons/antd'
import DataTableSuffix from "@/views/utils/DataTableSuffix.vue";
import PageHead from "@/views/utils/PageHead.vue";
import {handleDeleteGroup, handleGetAllGroups, handleSubmitNewGroup, handleUpdateGroup} from "@/api/admin/groups";
import useTablePagination from "@/hooks/useTablePagination";

const {t} = useI18n()
const message = useMessage()
const dialog = useDialog()

let modifyFunc = ref<string>('add')
let modifyId = ref<number | null>(null)
let showLoading = ref<boolean>(false)
let animated = ref<boolean>(false)

const [dataSize, pageCount] = useTablePagination()

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
              onClick: () => deleteGroup(row)
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
const deleteGroup = async (row: GroupItem) => {
  if (row.plan_count && row.user_count)
    if (row.plan_count > 0 || row.user_count > 0)
      return message.error(t('adminViews.groupMgr.common.delNotAllowed'))
  dialog.error({
    title: t('adminViews.common.dialog.delete'),
    content: () => {
      return h('div', {}, [
        h('p', {style: {fontWeight: 'bold', fontSize: '1rem', opacity: '0.8'}}, row.name),
        h('p', {style: {marginTop: '4px'}}, t('adminViews.groupMgr.common.delMention'))
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
      animated.value = false
      let data = await handleDeleteGroup(row.id)
      if (data && data.code === 200) {
        message.success(t('adminViews.common.deleteSuccess'))
        await getAllGroups()
      } else {
        message.error(t('adminViews.common.deleteFailure'))
      }
    }
  })
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
    message: '',
    trigger: 'blur'
  }
}


let newGroupName = ref<string>('')

let handleAddGroup = async () => {
  modifyFunc.value = 'add'
  newGroupName.value = ''
  showModal.value = true
}

let submitNewGroup = async () => {
  animated.value = false
  if (newGroupName.value.trim().length <= 0)
    return message.error(t('adminViews.common.formatNotAllowed'))
  let data = await handleSubmitNewGroup(newGroupName.value.trim())
  if (data && data.code === 200) {
    message.success(t('adminViews.common.addSuccess'))
    await getAllGroups()
  } else {
    message.error(t('adminViews.common.addFailure') + data.msg || '')
  }
}

let closeModal = () => {
  showModal.value = false
}

let getAllGroups = async () => {
  showLoading.value = true
  let data = await handleGetAllGroups(dataSize.value.page, dataSize.value.pageSize)
  if (data && data.code === 200) {
    groupList.value = []
    showLoading.value = false
    if (data.group_list)
      data.group_list.forEach((group: PrivilegeGroup) => groupList.value.push(group))
    pageCount.value = data.page_count || 0
    animated.value = true
  } else {
    message.error(t('adminViews.common.fetchDataFailure'))
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
  if (modifyId.value === null) return message.error(t('adminViews.common.updateFailure'))
  let data = await handleUpdateGroup(modifyId.value, newGroupName.value.trim())
  if (data && data.code === 200) {
    message.success(t('adminViews.common.updateSuccess'))
  } else {
    message.error(t('adminViews.common.updateFailure'))
  }
}

onBeforeMount(() => {
  themeStore.breadcrumb = 'adminViews.groupMgr.title'
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
              :scroll-x="1200"
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
        <n-input clearable v-model:value="newGroupName" :placeholder="t('adminViews.groupMgr.groupPlaceholder')"/>
      </n-form-item>
    </n-form>

  </n-modal>
</template>

<style lang="less" scoped>
.root {
  //padding: 20px;
  padding: 0 20px;

}

</style>