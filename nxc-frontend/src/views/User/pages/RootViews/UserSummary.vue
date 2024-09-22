<script setup lang="ts" name="UserSummary">
import {useI18n} from 'vue-i18n'
import {onMounted, ref} from "vue";
import {useRouter} from "vue-router";
import useUserInfoStore from "@/stores/useUserInfoStore";
import useAppInfosStore from "@/stores/useAppInfosStore";
import useThemeStore from "@/stores/useThemeStore";
import { useDialog, useMessage } from 'naive-ui'
const message = useMessage()
const dialog = useDialog()
import {
  ArrowBack,
  ArrowForward,
  KeyOutline as keyIcon,
  CartOutline as cartIcon,
  BookOutline as bookIcon,
  ArrowForwardCircleOutline as subIcon,
  BagHandleOutline as buyIcon,
  HelpBuoyOutline as supportIcon,
}from '@vicons/ionicons5'
import instance from "@/axios";
import useApiAddrStore from "@/stores/useApiAddrStore";

const {t, locale} = useI18n()
const apiAddrStore = useApiAddrStore();
const appInfoStore = useAppInfosStore();
const themeStore = useThemeStore();
const userInfoStore = useUserInfoStore();
const router = useRouter();

let thisNotices = ref([])

// let noticeBg = computed(() => ({
//   backgroundSize: 'cover',
//   backgroundImage: `url(https://ikanned.com:24444/d/Upload/NXC/IMG_1152.png)`,
//   backgroundPosition: 'center'
// }))
//
// let setNoticeBg = (imgUrl: string) => imgUrl === '' ? ({
//   backgroundImage: `url('https://ikanned.com:24444/d/Upload/NXC/noticeUniversalBgDay.svg')`,
//   // backgroundImage: `url('./noticeUniversalBgDay.svg')`,
// }) : ({
//   backgroundImage: `url(${imgUrl})`,
// })

// if (imgUrl === '') {
//   console.log('无默认背景')
//   return ({
//   })
// } else {
//   return ({
//   })
// }


let helpData = [
  {
    id: 0,
    title: computed(() => t('userSummary.tutorial.title')),
    content: computed(() => t('userSummary.tutorial.content', {name: appInfoStore.appCommonConfig.app_name})),
    icon_id: 1,
  },
  {
    id: 1,
    title: computed(() => t('userSummary.checkKey.title')),
    content: computed(() => t('userSummary.checkKey.content')),
    icon_id: 2,
  },
  {
    id: 2,
    title: computed(() => t('userSummary.renewPlan.title')),
    content: computed(() => t('userSummary.renewPlan.content')),
    icon_id: 3,
  },
  {
    id: 3,
    title: computed(() => t('userSummary.support.title')),
    content: computed(() => t('userSummary.support.content')),
    icon_id: 4,
  }
]

let pathById = [
  {
    id: 0,
    path: '/dashboard/document',
  },
  {
    id: 1,
    path: '/dashboard/purchase',
  },
  {
    id: 2,
    path: '/dashboard/purchase',
  },
  {
    id: 3,
    path: '/dashboard/support',
  },
]

let showDetail = () => {
  console.log('显示通知详情')
}

let getAllNotices = async () => {
  console.log('获取所有有效通知')
  try {
    let {data} = await instance.get(apiAddrStore.apiAddr.user.getAllNoticesList)
    if (data.code === 200) {
      thisNotices.value = data.notices
    }
  } catch (error) {
    console.log(error)
  }
}

let handleConfirm = (title:string, content: string) => {
  dialog.success({
    showIcon: false,
    title: title,
    content: content,
  })
}

interface MyActivePlans {
  expiration_date: string
  plan_name: string
}

let myActivePlans = ref<MyActivePlans[]>([]);
let haveActive = ref<boolean>(false)

let getActivePlanList = async () => {
  try {
    let {data} = await instance.get(apiAddrStore.apiAddr.user.getMyPlanList, {
      params: {
        user_id: userInfoStore.thisUser.id,
        // user_id: 3,
      }
    })
    if (data.code === 200) {
      haveActive.value = true
      myActivePlans.value = data.my_plans
      console.log(myActivePlans.value)
    }
  }catch (error) {
    haveActive.value = false
    console.log(error)
  }
}

let goCartPage = () => {
  console.log('to cart page')
}

onMounted(() => {
  console.log('用户summary挂载')
  themeStore.userPath = '/dashboard/summary'
  themeStore.menuSelected = 'user-dashboard'

  getAllNotices()
  getActivePlanList()


})
</script>

<template>
  <div class="root">
    <n-card content-style="padding: 0;" :embedded="true" hoverable>
      <n-carousel show-arrow autoplay style="border-radius: 3px">
        <n-card
            v-for="item in thisNotices"
            :key="item.id"
            class="item"
            :style="item.img_url===''?({backgroundImage: `url('https://ikanned.com:24444/d/Upload/NXC/noticeUniversalBgDay.svg')`,}):({backgroundImage: `url(${item.img_url})`,backgroundSize: 'cover'})"
            @click="handleConfirm(item.title, item.content)"
        >
          <div class="content">
            <n-tag class="tag" type="warning">{{ item.tags }}</n-tag>
            <div class="inner">
              <p class="title">{{ item.title }}</p>
              <p class="date">{{ item.created_at }}</p>
            </div>
          </div>

        </n-card>

        <template #arrow="{ prev, next }">
          <div class="custom-arrow">
            <button type="button" class="custom-arrow--left" @click="prev">
              <n-icon>
                <ArrowBack/>
              </n-icon>
            </button>
            <button type="button" class="custom-arrow--right" @click="next">
              <n-icon>
                <ArrowForward/>
              </n-icon>
            </button>
          </div>
        </template>
        <template #dots="{ total, currentIndex, to }">
          <ul class="custom-dots">
            <li
                v-for="index of total"
                :key="index"
                :class="{ ['is-active']: currentIndex === index - 1 }"
                @click="to(index - 1)"
            />
          </ul>
        </template>
      </n-carousel>
    </n-card>


    <n-card
        class="my-subscribe"
        :embedded="true"
        hoverable
        :title="t('userSummary.myPlan')"
        content-style="padding: 0"
    >

      <!--      <h1 v-for="i in 3" :key="i">{{i}}</h1>-->
      <!--      {{userInfoStore.thisUser}}-->

      <n-card
          v-if="!haveActive"
          class="no-license"
          content-style="padding: 0;"
          @click="goCartPage"
      >
        <div
            style="display: flex;text-align: center;justify-content: flex-end;line-height: 20px;flex-direction: column;align-items: center;padding-bottom: 20px;">
          <n-icon style="text-align: center;font-size: 30px;">
            <cartIcon/>
          </n-icon>
          <p class="describe" style="margin-top: 5px;opacity: 0.8;font-size: 1rem;font-weight: bold;">{{ t('userSummary.toPurchase') }}</p>
        </div>
      </n-card>

      <n-card
          v-else-if="haveActive"
          v-for="(plan, index) in myActivePlans"
          :key="index"
          class="license-active"
          content-style="padding: 0;"
          style="padding: 0 25px 0 25px"
      >
        <div class="plan-item">
          <p style="font-size: 1.1rem; font-weight: bold;opacity: 0.9;">{{ plan.plan_name }}</p>
          <p style="font-size: 12px; opacity: 0.6;margin-top: 3px">{{ t('userSummary.timeLeft', {msg: plan.expiration_date, })}}</p>
        </div>

        <n-hr v-if="!(index === (myActivePlans.length - 1))"></n-hr>
      </n-card>


    </n-card>

    <n-card
        style="margin: 20px 0; padding-right: 20px"
        :embedded="true"
        hoverable
        :title="t('userSummary.shortcut')"
        content-style="padding: 0;">
      <div
          :class="themeStore.enableDarkMode?'help-day':'help-night'"
          v-for="item in helpData"
          :key="item.id"
          style="display: flex; justify-content: space-between; width: 100%; height: auto; padding: 10px"
        @click="router.push({ path: pathById[item.id].path})"
      >
        <div style="height: 100%; text-align: left; padding-left: 20px">
          <p style="font-size: 16px;">{{ item.title }}</p>
          <p style="font-size: 12px; margin-top: 5px; opacity: 0.5">{{ item.content }}</p>
        </div>
        <div style="width: 40px; font-size: 30px; margin-right: 20px; display: flex; flex-direction: column; justify-content: center">
          <n-icon v-if="item.icon_id===1" style="opacity: 0.5; width: 40px; "><bookIcon/></n-icon>
          <n-icon v-if="item.icon_id===2" style="opacity: 0.5; width: 40px; "><keyIcon/></n-icon>
          <n-icon v-if="item.icon_id===3" style="opacity: 0.5; width: 40px; "><buyIcon/></n-icon>
          <n-icon v-if="item.icon_id===4" style="opacity: 0.5; width: 40px; "><supportIcon/></n-icon>
        </div>
      </div>
    </n-card>


  </div>
</template>

<style scoped lang="less">

.root {
  padding: 20px;

  .my-subscribe {
    margin-top: 20px;
    padding-bottom: 20px;
  }
}

.item {
  background-repeat: repeat;
  background-position: center;

  .content {
    .tag {
      width: auto;
      padding: 0 5px 0 5px;
      text-align: center;
    }

    .inner {
      margin: 50px 0 20px 0;

      .title {
        font-size: 1.5rem;
        font-weight: bold;
        opacity: 0.8;
      }

      .date {
        font-size: 16px;
        opacity: 0.8;
      }
    }
  }
}

.item::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-repeat: repeat;
  background-position: center;
  opacity: 0.3;
  background-color: rgba(0, 0, 0, 0.5);
}


//.item::before {
//  content: "";
//  height: 440px;
//
//  background-color: rgba(130,130,130,0.5);
//
//}


.carousel-img {
  width: 100%;
  height: 240px;
  object-fit: cover;
}

.custom-arrow {
  display: flex;
  position: absolute;
  bottom: 25px;
  right: 10px;
}

.custom-arrow button {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 28px;
  height: 28px;
  margin-right: 12px;
  color: #fff;
  background-color: rgba(255, 255, 255, 0.1);
  border-width: 0;
  border-radius: 8px;
  transition: background-color 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  cursor: pointer;
}

.custom-arrow button:hover {
  background-color: rgba(255, 255, 255, 0.2);
}

.custom-arrow button:active {
  transform: scale(0.95);
  transform-origin: center;
}

.custom-dots {
  display: flex;
  margin: 0;
  padding: 0;
  position: absolute;
  bottom: 20px;
  left: 20px;
}

.custom-dots li {
  display: inline-block;
  width: 12px;
  height: 4px;
  margin: 0 3px;
  border-radius: 4px;
  background-color: rgba(255, 255, 255, 0.4);
  transition: width 0.3s,
  background-color 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  cursor: pointer;
}

.custom-dots li.is-active {
  width: 40px;
  background: #fff;
}

.n-card {
  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
  border: 0;
}

.help-day:hover {
  background-color: rgba(220, 220, 220, 0.1);
}

.help-night:hover {
  background-color: rgba(69, 69, 69, 0.1);
}


.n-card {
  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
  border: 0;
}

.plan-item {
  transition: transform 200ms ease;
}

.plan-item:hover {
  transform: translateY(-3px);
}

</style>