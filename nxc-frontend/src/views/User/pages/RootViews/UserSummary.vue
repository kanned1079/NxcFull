<script setup lang="ts">
import {useI18n} from 'vue-i18n'
import {computed, onMounted, ref} from "vue";
import {useRouter} from "vue-router";
import useUserInfoStore from "@/stores/useUserInfoStore";
import useAppInfosStore from "@/stores/useAppInfosStore";
import useThemeStore from "@/stores/useThemeStore";
import {useDialog, useMessage} from 'naive-ui'
import {formatDate} from "@/utils/timeFormat"

import {
  ArrowBack,
  ArrowForward,
  BagHandleOutline as buyIcon,
  BookOutline as bookIcon,
  CartOutline as cartIcon,
  ChevronForwardOutline as toRightIcon,
  HelpBuoyOutline as supportIcon,
  KeyOutline as keyIcon,
} from '@vicons/ionicons5'
import instance from "@/axios";
// import useApiAddrStore from "@/stores/useApiAddrStore";

let animated = ref<boolean>(false)

// const message = useMessage()
const dialog = useDialog()
const {t, locale} = useI18n()
// const apiAddrStore = useApiAddrStore();
const appInfoStore = useAppInfosStore();
const themeStore = useThemeStore();
const userInfoStore = useUserInfoStore();
const router = useRouter();

interface Notice {
  id: number
  tags: string
  title: string
  content: string
  img_url?: string
  created_at: string
}

let thisNotices = ref<Notice[]>([])

let helpData = [
  {
    id: 0,
    title: computed(() => t('userSummary.tutorial.title')).value,
    content: computed(() => t('userSummary.tutorial.content', {name: appInfoStore.appCommonConfig.app_name})).value,
    icon_id: 1,
  },
  {
    id: 1,
    title: computed(() => t('userSummary.checkKey.title')).value,
    content: computed(() => t('userSummary.checkKey.content')).value,
    icon_id: 2,
  },
  {
    id: 2,
    title: computed(() => t('userSummary.renewPlan.title')).value,
    content: computed(() => t('userSummary.renewPlan.content')).value,
    icon_id: 3,
  },
  {
    id: 3,
    title: computed(() => t('userSummary.support.title')).value,
    content: computed(() => t('userSummary.support.content')).value,
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
    path: '/dashboard/keys',
  },
  {
    id: 2,
    path: '/dashboard/purchase',
  },
  {
    id: 3,
    path: '/dashboard/tickets',
  },
]

// let showDetail = () => {
//   console.log('显示通知详情')
// }

let getAllNotices = async () => {
  console.log('获取所有有效通知')
  try {
    let {data} = await instance.get('/api/user/v1/notice', {
      params: {
        is_user: true,
      }
    })
    if (data.code === 200) {
      thisNotices.value = data.notices
      data.notices.forEach((notice: Notice) => thisNotices.value.push(notice))
    }
  } catch (error) {
    console.log(error)
  }
}

let handleConfirm = (title: string, content: string) => {
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
    let {data} = await instance.get('/api/user/v1/plan/summary/fetch', {
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
  } catch (error) {
    haveActive.value = false
    console.log(error)
  }
}

// let goCartPage = () => {
//   console.log('to cart page')
//
// }

onMounted(async () => {
  console.log('用户summary挂载')
  themeStore.userPath = '/dashboard/summary'
  themeStore.menuSelected = 'user-dashboard'

  await getAllNotices()
  await getActivePlanList()

  animated.value = true

})
</script>

<script lang="ts">
export default {
  name: 'UserSummary',
}
</script>

<template>
  <transition name="slide-fade">
    <div class="root" v-if="animated">
      <n-alert :bordered="false" style="margin-bottom: 20px" title="" type="warning">
        <div style="display: flex; flex-direction: row; align-items: center">
          您有待处理的工单
          <n-button style="margin-left: 5px" text type="primary" @click="router.push({path: '/dashboard/tickets'})">
            去查看
            <n-icon>
              <toRightIcon/>
            </n-icon>
          </n-button>
        </div>
      </n-alert>
      <n-card content-style="padding: 0;" :embedded="true" hoverable :bordered="false">
        <n-carousel show-arrow autoplay style="border-radius: 3px">
          <n-card
              v-for="item in thisNotices"
              :key="item.id"
              class="item"
              @click="handleConfirm(item.title, item.content)"
              :bordered="false"
          >
            <div class="content">
              <n-tag
                  class="tag"
                  type="warning"
              >
                {{ item.tags }}
              </n-tag>
              <div class="inner">
                <p :style="item.img_url !== ''? { color: '#fff' }: null" class="title">
                  {{ item.title }}
                </p>
                <p :style="item.img_url !== ''? { color: '#fff' }: null" class="date">
                  {{ formatDate(item.created_at) }}
                </p>
              </div>
            </div>
            <div
                :class="item.img_url===''?'content-bg':'content-bg-img'"
                :style="item.img_url === '' ? { filter: 'brightness(0.9) contrast(0.5)' } : {
                  backgroundImage: `url(${item.img_url})`,
                  ...(themeStore.enableDarkMode ? { filter: 'brightness(0.6) contrast(0.8)' } : {filter: 'brightness(1) contrast(0.6) grayscale(0.2)'})}"
            >
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
          :bordered="false"
      >
        <n-card
            v-if="!haveActive"
            class="no-license"
            style="background-color: rgba(0,0,0,0.0)"
            content-style="padding: 0;"
            @click="router.push({path: '/dashboard/purchase'})"
            :bordered="false"
        >
          <div
              style="display: flex;text-align: center;justify-content: flex-end;line-height: 20px;flex-direction: column;align-items: center;padding: 20px; background-color: rgba(0,0,0,0.0)">
            <n-icon style="text-align: center" size="40">
              <cartIcon/>
            </n-icon>
            <p class="describe" style="margin-top: 5px;opacity: 0.8;font-size: 1rem;font-weight: bold;">
              {{ t('userSummary.toPurchase') }}</p>
          </div>
        </n-card>

<!--        <n-card-->
<!--            v-else-if="haveActive"-->
<!--            v-for="(plan, index) in myActivePlans"-->
<!--            :key="index"-->
<!--            class="license-active"-->
<!--            content-style="padding: 0;"-->
<!--            style="padding: 0 25px 0 25px; background-color: rgba(0,0,0,0.0)"-->
<!--            :bordered="false"-->
<!--        >-->
<!--          <div class="plan-item">-->
<!--            <p style="font-size: 1.1rem; font-weight: bold;opacity: 0.9;">{{ plan.plan_name }}</p>-->
<!--            <p style="font-size: 12px; opacity: 0.6;margin-top: 3px">-->
<!--              {{ t('userSummary.timeLeft', {msg: plan.expiration_date,}) }}</p>-->
<!--          </div>-->

<!--          <n-hr v-if="!(index === (myActivePlans.length - 1))"></n-hr>-->
<!--        </n-card>-->

        <n-card
            v-else-if="haveActive"
            v-for="(plan, index) in myActivePlans.slice(0, 5)"
            :key="index"
            class="license-active"
            content-style="padding: 0;"
            style="padding: 0 25px 0 25px; background-color: rgba(0,0,0,0.0)"
            :bordered="false"
        >
          <div class="plan-item">
            <p style="font-size: 1.1rem; font-weight: bold; opacity: 0.9;">{{ plan.plan_name }}</p>
            <p style="font-size: 12px; opacity: 0.6; margin-top: 3px">
              {{ t('userSummary.timeLeft', { msg: plan.expiration_date }) }}
            </p>
          </div>

          <n-hr v-if="index < myActivePlans.slice(0, 5).length - 1"></n-hr>
        </n-card>

        <div
            v-if="myActivePlans.length > 5"
            style="margin: 10px 20px 0 20px;"
        >
          <div style="display: flex; flex-direction: row; align-items: center; justify-content: end;">
            <n-button type="primary" text @click="router.push({path: '/dashboard/keys'})">
              查看所有密鑰
            <n-icon style="margin-left: 5px"><toRightIcon/></n-icon>
            </n-button>
          </div>
        </div>


      </n-card>

      <n-card
          style="margin: 20px 0; padding-right: 20px"
          :embedded="true"
          hoverable
          :title="t('userSummary.shortcut')"
          content-style="padding: 0;"
          :bordered="false"
      >
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
          <div
              style="width: 40px; font-size: 30px; margin-right: 20px; display: flex; flex-direction: column; justify-content: center">
            <n-icon v-if="item.icon_id===1" style="opacity: 0.5; width: 40px; ">
              <bookIcon/>
            </n-icon>
            <n-icon v-if="item.icon_id===2" style="opacity: 0.5; width: 40px; ">
              <keyIcon/>
            </n-icon>
            <n-icon v-if="item.icon_id===3" style="opacity: 0.5; width: 40px; ">
              <buyIcon/>
            </n-icon>
            <n-icon v-if="item.icon_id===4" style="opacity: 0.5; width: 40px; ">
              <supportIcon/>
            </n-icon>
          </div>
        </div>
      </n-card>


    </div>
  </transition>


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
  background-repeat: no-repeat;
  background-position: center;
  position: relative;
  width: 100%;

  background-size: cover;
  background-attachment: fixed;
  background-position-x: center;

  .content {
    position: relative;
    z-index: 1; /* 确保内容在背景图之上 */
    width: 100%;

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
        opacity: 0.9;
      }

      .date {
        font-size: 16px;
        opacity: 0.8;
      }
    }
  }

  .content-bg {
    //grid-auto-rows: 300px;
    //display: grid;
    //grid-template-columns: repeat(auto-fill, minmax(100px, 1fr)); /* 根据每个组件的大小进行调整 */
    background-image: url("@/assets/noticeUniversalBgDay.svg");
    //display: flex;
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: 0; /* 设置在底层 */
    //background: url('your-background-image-url') center center / cover no-repeat; /* 替换为实际背景图 */
    opacity: 1; /* 背景透明度 */
  }

  .content-bg-img {
    //grid-auto-rows: 300px;
    //display: grid;
    //grid-template-columns: repeat(auto-fill, minmax(100px, 1fr)); /* 根据每个组件的大小进行调整 */
    //background-image: url("@/assets/noticeUniversalBgDay.svg");
    //display: flex;
    //filter: brightness(0.7) contrast(0.8);
    background-repeat: no-repeat;
    background-position: center;
    background-size: cover;
    background-attachment: fixed;
    background-position-x: center;
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: 0; /* 设置在底层 */
    //background: url('your-background-image-url') center center / cover no-repeat; /* 替换为实际背景图 */
    opacity: 1; /* 背景透明度 */
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
  background-color: rgba(129, 129, 129, 0.2);
}

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
  background-color: rgba(255, 255, 255, 0.8);
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

//.n-card {
//  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
//  border: 0;
//}

.help-day:hover {
  background-color: rgba(220, 220, 220, 0.1);
}

.help-night:hover {
  background-color: rgba(69, 69, 69, 0.1);
}


//.n-card {
//  background-color: v-bind('themeStore.getTheme.globeTheme.cardBgColor');
//  border: 0;
//}

.plan-item {
  transition: transform 200ms ease;
}

.plan-item:hover {
  transform: translateY(-3px);
}

</style>