<script lang="ts" setup>
import {computed, h, onBeforeMount, onMounted, ref} from "vue";
import useUserInfoStore from "@/stores/useUserInfoStore";
import {
  CheckmarkOutline as copiedIcon,
  CopyOutline as copyIcon,
  GiftOutline as giftIcon,
  MedalOutline as rewardIcon,
} from "@vicons/ionicons5"
import useAppInfosStore from "@/stores/useAppInfosStore";
import {useI18n} from "vue-i18n";
import useThemeStore from "@/stores/useThemeStore";
// import instance from "@/axios";
import {
  handleCreateMyInviteCode,
  handleGetInviteMsg,
  handleGetMyInviteCode,
  handleGetMyInvitedUserList
} from "@/api/user/invite";
import {NButton, NIcon, NTag, useMessage, type DataTableColumns, useDialog} from "naive-ui";
import {formatDate} from "@/utils/timeFormat";
import DataTableSuffix from "@/views/utils/DataTableSuffix.vue";
import PageHead from "@/views/utils/PageHead.vue";
import useTablePagination from "@/hooks/useTablePagination";

const {t} = useI18n();
const userInfoStore = useUserInfoStore();
const appInfoStore = useAppInfosStore();
const themeStore = useThemeStore();
const message = useMessage();
const dialog = useDialog();

let animated = ref<boolean>(false);
let listAnimated = ref<boolean>(false)
let showCreateInviteCodeMention = ref<boolean>(false);
let showTableLoading = ref<boolean>(true)

let myFaCode = ref<string>('')
let inviteMsg = ref<string>('')

interface MyInvitedUser {
  id: number
  email: string
  created_at: string
}

let myInvitedUserList = ref<MyInvitedUser[]>([]);

// let pageCount = ref(10)
//
// let dataSize = ref<{ pageSize: number, page: number }>({
//   pageSize: 10,
//   page: 1,
// })

let [dataSize, pageCount] = useTablePagination()


const columns = computed<DataTableColumns<MyInvitedUser>>(() => [
  {
    title: '#',
    key: 'id'
  },
  {
    title: t('userInvite.email'),
    key: 'email'
  },
  {
    title: t('userInvite.createdAt'),
    key: 'created_at',
    render(row: MyInvitedUser) {
      return h('span', {}, {default: () => formatDate(row.created_at)});
    }
  },
]);

let createMyCodeClick = () => {
  dialog.info({
    title: t('userInvite.generateCodeConfirm'),
    content: t('userInvite.generateCodeHint'),
    positiveText: t('adminViews.common.confirm'),
    negativeText: t('adminViews.common.cancel'),
    showIcon: false,
    actionStyle: {
      marginTop: '20px',
    },
    contentStyle: {
      marginTop: '20px',
    },
    onPositiveClick: () => callHandleCreateMyInviteCode(),
  })
}

let callHandleCreateMyInviteCode = async () => {
  let data = await handleCreateMyInviteCode(userInfoStore.thisUser.id)
  if (data.code === 200) {
    console.log('created ok.');
    await callHandleGetMyInviteCode()
    await callHandleGetMyInvitedUserList()
    // showCreateInviteCodeMention.value = false

  } else {
    message.error('create failure: ' + data.msg || '');
  }
}

let callHandleGetInviteMsg = async () => {
  let data = await handleGetInviteMsg()
  if (data.code === 200) {
    console.log(data)
    inviteMsg.value = data.invite_msg || ''
    // animated.value = true
  } else if (data.code === 404) {
  } else {
    console.log(data.msg)
  }
}

let callHandleGetMyInviteCode = async () => {
  let data = await handleGetMyInviteCode(userInfoStore.thisUser.id)
  if (data.code === 200) {
    console.log(data)
    myFaCode.value = data.my_invite_code
    animated.value = true
    showTableLoading.value = false
  } else if (data.code === 404) {
  } else {
    console.log(data.msg)
  }
}

let callHandleGetMyInvitedUserList = async () => {
  let data = await handleGetMyInvitedUserList(userInfoStore.thisUser.id, dataSize.value.page, dataSize.value.pageSize);
  if (data && data.code === 200 && data.user_list) {
    listAnimated.value = true
    myInvitedUserList.value = []
    // console.log(data)
    pageCount.value = data.page_count
    data.user_list.forEach((item: MyInvitedUser) => myInvitedUserList.value.push(item))
    animated.value = true
  } else if (data.code === 404) {
  } else {
    console.log(data.msg)
  }
  // listAnimated.value = true

}

let faLink = computed(() => `${appInfoStore.appCommonConfig.app_url}/register?code=${myFaCode.value}`)

// 复制成功的反馈
const copySuccess = ref(false);

// 复制文本的函数
const copyText = async (key: string, event: MouseEvent) => {
  event.stopPropagation()
  try {
    await navigator.clipboard.writeText(key);
    // message.success(t('userKeys.copiedSuccessMessage'))
    message.success(t('userKeys.hoverCopiedSuccessMention'))
    copySuccess.value = true
    // copySuccess.value = true;
    setTimeout(() => {
      copySuccess.value = false
      // copySuccess.value = false;
      // message.success('复制成功')
    }, 1200); // 显示2秒成功信息
  } catch (err) {
    console.error(t('userKeys.copyFailure'), err);
  }
};

onBeforeMount(async () => {
  themeStore.breadcrumb = 'userInvite.myInvite'
  themeStore.menuSelected = 'user-invite'
  await callHandleGetInviteMsg()
})


onMounted(async () => {
  themeStore.userPath = '/dashboard/invite'
  // await handleGetInviteMsg()
  await callHandleGetMyInviteCode()

  if (myFaCode.value !== '') {
    await callHandleGetMyInvitedUserList()
  } else {
    showTableLoading.value = false
  }

  animated.value = true;
})

</script>

<script lang="ts">
export default {
  name: "UserInvite",
}
</script>

<template>

  <PageHead
      :title="t('userInvite.myInvite')"
      :description="t('userActivation.description')"
      style="margin: 20px 20px 0 20px"
  />

  <div style="padding: 0px 20px 14px 20px">
<!--    <n-card-->
<!--        hoverable-->
<!--        :embedded="true"-->
<!--        :bordered="false"-->
<!--        :title="t('userInvite.myInvite')"-->
<!--    ></n-card>-->

    <n-collapse-transition :show="inviteMsg.trim() !== ''">
      <n-card
          hoverable
          :bordered="false"
          :embedded="true"
          content-style="padding: 0"
      >
        <n-alert type="warning" :bordered="false">
          {{ inviteMsg }}
          <template #icon>
            <n-icon><rewardIcon /></n-icon>
          </template>
        </n-alert>
      </n-card>
    </n-collapse-transition>


  </div>


  <transition name="slide-fade">
    <div style="padding: 0 20px 15px 20px" v-if="animated">

      <n-card
          hoverable
          :embedded="true"
          :bordered="false"
      >

        <div class="code-mgr-main">
          <p class="code-mgr-main-title">{{ t('userInvite.faCodeManage') }}</p>
          <n-button
              type="primary"
              secondary
              size="small"
              @click="createMyCodeClick"
          >
            {{ myFaCode === '' ? t('userInvite.generateFaCode') : t('userInvite.flushFaCode') }}
          </n-button>
        </div>

        <n-table
            :bordered="false"
            :bottom-bordered="false"
            class="code-info"
        >
          <n-tr style=" margin-left: 200px">
            <n-td style="background-color: rgba(0,0,0,0.0);">{{ t('userInvite.faCode') }}</n-td>
            <n-td style="background-color: rgba(0,0,0,0.0);">
              <n-tag
                  type="default"
                  checkable
              >
                {{ myFaCode !== '' ? myFaCode : t('userInvite.noFaCode') }}
              </n-tag>
            </n-td>
          </n-tr>
          <n-tr style=" margin-left: 200px">
            <n-td style="background-color: rgba(0,0,0,0.0);">{{ t('userInvite.faLink') }}</n-td>
            <n-td style="background-color: rgba(0,0,0,0.0);">
              <n-tag
                  type="default"
                  checkable
                  @click="myFaCode !== ''?copyText(faLink, $event):null"
              >

                <div class="fa-link-main">
                  <p class="fa-link">
                    {{ myFaCode !== '' ? faLink : t('userInvite.generateFaCodePlease') }}
                  </p>
                  <n-icon v-if="myFaCode !== '' && !copySuccess" class="fa-link-copy-icon">
                    <copyIcon/>
                  </n-icon>
                  <n-icon v-if="myFaCode !== '' && copySuccess" class="fa-link-copy-icon">
                    <copiedIcon/>
                  </n-icon>
                </div>


              </n-tag>
            </n-td>
          </n-tr>
        </n-table>


      </n-card>


      <transition name="slide-fade">
        <div v-if="listAnimated">
          <n-card
              hoverable
              :embedded="true"
              :bordered="false"
              style="margin-top: 15px"
              content-style="padding: 0"
          >
            <div style="padding: 20px">
              <p style="font-weight: bold; font-size: 1.1rem">{{ t('userInvite.usersMyInvited') }}</p>

            </div>

            <n-data-table
                striped
                class="table"
                :columns="columns"
                :data="myInvitedUserList"
                :pagination="false"
                :bordered="true"
                :loading="showTableLoading"
            />
          </n-card>

          <DataTableSuffix
              v-model:data-size="dataSize"
              v-model:page-count="pageCount"
              v-model:animated="animated"
              :update-data="callHandleGetMyInvitedUserList"
          />
<!--          <div style="margin-top: 20px; display: flex; flex-direction: row; justify-content: right;">-->
<!--            <n-pagination-->
<!--                size="medium"-->
<!--                v-model:page.number="dataSize.page"-->
<!--                :page-count="pageCount"-->
<!--                @update:page="animated=false; callHandleGetMyInvitedUserList()"-->
<!--            />-->
<!--            <n-select-->
<!--                style="width: 160px; margin-left: 20px"-->
<!--                v-model:value.number="dataSize.pageSize"-->
<!--                size="small"-->
<!--                :options="dataCountOptions"-->
<!--                :remote="true"-->
<!--                @update:value="animated=false; dataSize.page = 1; callHandleGetMyInvitedUserList()"-->
<!--            />-->
<!--          </div>-->
        </div>
      </transition>


    </div>
  </transition>

<!--  <n-modal-->
<!--      v-model:show="showCreateInviteCodeMention"-->
<!--      preset="dialog"-->
<!--      :title="t('userInvite.generateCodeConfirm')"-->
<!--      :content="t('userInvite.generateCodeHint')"-->
<!--      :positive-text="t('userInvite.positiveClick')"-->
<!--      :negative-text="t('userInvite.negativeClick')"-->
<!--      @positive-click="callHandleCreateMyInviteCode"-->
<!--      @negative-click="showCreateInviteCodeMention=false"-->
<!--  />-->

</template>

<style lang="less" scoped>

.code-mgr-main {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;

  .code-mgr-main-title {
    font-size: 1.1rem;
    font-weight: bold;
  }
}

.code-info {
  margin-top: 15px;
  background-color: rgba(0, 0, 0, 0.0);
}

.code-div {
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  align-items: baseline;
  font-size: 1rem;

  .code-title {
    margin-right: 6px;
  }
}

.fa-link-main {
  display: flex;
  flex-direction: row;
  align-items: center;

  .fa-link {
    margin-right: 6px;
    transition: ease 200ms;
  }

  .fa-link-copy-icon {
    transition: ease 200ms;
    opacity: 0;
  }
}

.fa-link-main:hover {
  .fa-link {
    text-decoration: underline;
  }

  .fa-link-copy-icon {
    opacity: 0.7;
  }
}
</style>