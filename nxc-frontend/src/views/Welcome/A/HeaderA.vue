<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {onBeforeUnmount, onMounted, ref} from "vue";
import useApiAddrStore from "@/stores/useApiAddrStore";
import useAppInfosStore from "@/stores/useAppInfosStore";
import {useRouter} from "vue-router";
import {
  Language as langIcon,
  Menu as menuIcon,
  InfiniteOutline as aboutIcon,
  CashOutline as pricingIcon,
  LogInOutline as loginIcon,
  PersonAddOutline as registerIcon,

} from '@vicons/ionicons5'
import renderIcon from "@/utils/iconFormator";


const {t, locale} = useI18n()

const router = useRouter();
const apiAddrStore = useApiAddrStore();
const appInfoStore = useAppInfosStore();

// let dropDownBtn = ref(null)
let windowWidth = ref<number>(0)

let langOptions = ref([
  {
    label: '简体中文 zh_CN',
    key: 'zh_CN',
    disabled: false
  },
  {
    label: '繁體中文 zh_HK',
    key: 'zh_HK',
    disabled: false
  },
  {
    label: 'English en_US',
    key: 'en_US',
    disabled: false
  },
  {
    label: '日本語 ja_JP',
    key: 'ja_JP',
    disabled: false
  },
])

let collapsedMenuOptions = ref([
  {
    label: '关于我们',
    key: 'about-us',
    disabled: false,
    icon: renderIcon(aboutIcon)

  },
  {
    label: '定价',
    key: 'pricing',
    disabled: false,
    icon: renderIcon(pricingIcon)
  },
  {
    label: '登录',
    key: 'login',
    disabled: false,
    icon: renderIcon(loginIcon)
  },
  {
    label: '注册',
    key: 'register',
    disabled: false,
    icon: renderIcon(registerIcon)
  },
])

// 设置切换语言
let handleSelectLang = (lang: string) => {
  console.log(lang)
  // lang: en_US, zh_CN, zh_HK, ja_JP
  // 动态切换语言
  locale.value = lang
  localStorage.setItem('locale', lang)  // 保存语言到 localStorage
}

let handleEnter = (link: string) => {
  console.log(link)
  switch (link) {
    case 'about-us': {
      break
    }
    case 'pricing': {

      break
    }
    case 'login': {
      router.push({path: '/login'})
      break
    }
    case 'register': {
      router.push({path: '/register'})
      break
    }
  }
}

let collapse = ref<boolean>(false)
let updateWidth = () => {
  windowWidth.value = window.innerWidth;
  windowWidth.value <= 1000 ? collapse.value = true : collapse.value = false
}

onMounted(() => {
  updateWidth()
  window.addEventListener('resize', updateWidth);
})

onBeforeUnmount(() => {
  window.removeEventListener('resize', updateWidth);
})

</script>

<script lang="ts">
export default {
  name: 'HeaderA',
}
</script>

<template>
  <div class="root">
    <div class="left-content">
      <span>
          <img :src="appInfoStore.appCommonConfig.logo_url" alt="icon">
      </span>
      <p class="name">{{ appInfoStore.appCommonConfig.app_name }}</p>
      <n-button v-if="!collapse" size="large" style="margin-right: 40px" color="#277aa8" text >
        {{ t('welcome.A.aboutUs') }}
      </n-button>
      <n-button v-if="!collapse" size="large" color="#277aa8" text>
        {{ t('welcome.A.pricing') }}
      </n-button>
    </div>
    <div class="right-content">
      <n-dropdown
          trigger="hover"
          :options="langOptions"
          @select="handleSelectLang"

      >
        <n-button
            style="margin-right: 20px; color: #277aa8"
            quaternary
            class="btn"
            size="medium"
        >
          <template #icon>
            <n-icon size="20">
              <langIcon/>
            </n-icon>
          </template>
        </n-button>
      </n-dropdown>

      <n-dropdown
          v-if="collapse"
          trigger="hover"
          :options="collapsedMenuOptions"
          @select="handleEnter"
          style="width: 200px;"
      >
        <n-button style="margin-right: 20px; color: #277aa8" quaternary class="btn" size="medium">
          <template #icon>
            <n-icon size="20">
              <menuIcon/>
            </n-icon>
          </template>
        </n-button>
      </n-dropdown>

      <n-button
          v-if="!collapse"
          text
          color="#277aa8"
          style="margin-right: 20px"
          size="large"
          type="primary"
          @click="router.push({path: '/login'})"
      >{{ t('welcome.A.login') }}
      </n-button>
      <n-button
          v-if="!collapse"
          text
          color="#277aa8"
          size="large"
          type="primary"
          @click="router.push({path: '/register'})"
      >{{ t('welcome.A.register') }}
      </n-button>
    </div>
  </div>
</template>

<style scoped lang="less">
.root {
  height: 70px;
  line-height: 70px;
  background-color: rgba(200, 217, 235, 0.8);
  backdrop-filter: blur(10px);
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
  padding: 0 20px;
  color: #277aa8;
  z-index: 10000;

  .left-content {
    display: flex;
    flex-direction: row;
    align-items: center;
    // 新增以下样式

    span {
      display: flex;

      flex-direction: row;
      align-items: center;
      margin-right: 10px;

      img {
        width: 30px;
      }
    }

    .name {
      font-size: 1.5rem;
      margin-right: 100px;
      color: #277aa8;
    }
  }

  .right-content {
    display: flex;
    flex-direction: row;
    align-items: center;
  }
}
</style>