<script setup lang="ts">
import useThemeStore from "@/stores/useThemeStore";
import {computed, onMounted, ref} from "vue";
import {type FormInst, useMessage} from "naive-ui";

let animated = ref<boolean>(false)

const message = useMessage()
let showModal = ref<boolean>(false)
const themeStore = useThemeStore();
const formRef = ref<FormInst | null>(null)

interface ThemeConfig {
  id: number
  general_color_name: string
  login_page_bg_url: string
}

let formValue = ref<ThemeConfig>({
  id: 0,
  general_color_name: '',
  login_page_bg_url: '',

})

let rules = [
  {}
]

interface ItemConfig {
  name: string,
  describe: string,
  backgroundImageUrl: string
}

let themeBg = computed(() => ({
  backgroundSize: 'cover',
  backgroundImage: `url(https://ikanned.com:24444/d/Upload/NXC/IMG_1152.png)`,
  backgroundPosition: 'center'
}))

// coverColor 设置不同颜色模式下的遮罩颜色
let coverColor = computed(() => {
  if (!themeStore.enableDarkMode) {
    return `linear-gradient(to right, rgba(255, 255, 255, 1) 10%, rgba(255, 255, 255, 0) 100%)`
  } else {
    return `linear-gradient(to right, ${themeStore.getTheme.globeTheme.cardBgColor} 10%, rgba(255, 255, 255, 0) 100%)`
  }
})

let itemBgImage = computed(() => {
})

let allThemes = {
  theme1: {
    name: 'Journey\'s terminus.',
    describe: '这是使用的默认主题',
    backgroundImageUrl: 'https://ikanned.com:2444/d/Upload/NXC/ani24420.jpg',
    enabled: true,
  },
  theme2: {
    name: 'Naniwa',
    describe: '这是使用的默认主题',
    backgroundImageUrl: 'https://ikanned.com:2444/d/Upload/NXC/IMG_1152.png',
    enabled: false,
  },
  theme3: {
    name: 'Sky',
    describe: '这是使用的默认主题',
    backgroundImageUrl: 'https://ikanned.com:2444/d/Upload/NXC/101209570_p0.jpg',
    enabled: false,
  }
}

let handleSetTheme = (name: string) => {
  console.log(name)
}

let darkerImage = computed(() => {
  if (themeStore.enableDarkMode)
    return {opacity: 0.5}
})

onMounted(() => {
  themeStore.menuSelected = 'theme-config'
  themeStore.contentPath = '/admin/dashboard/theme'

  animated.value = true
})

interface SaleItem {
  id: number,
  name: string,
  price: number
}

let preSaleItems = ref<SaleItem[]>([
  {
    id: 0,
    name: '东西1',
    price: 10.21,
  },
  {
    id: 1,
    name: '东西2',
    price: 6.40,
  },
  {
    id: 2,
    name: '东西3',
    price: 16.02,
  }
])

let addedItems = ref<SaleItem[]>([])

let addItem2List = (item: SaleItem) => {
  addedItems.value.push(item)
}

let addItem2ListByIndex = (index: number) => {
  addedItems.value.push(preSaleItems.value[index])
}

let getTotalPrice = computed(() => {
  let amount = 0;
  addedItems.value.forEach((item: SaleItem) => {
    amount += item.price;
  })
  return amount;
})

</script>

<script lang="ts">
export default {
  name: 'ThemeConfig'
}
</script>

<template>
  <div class="root">
    <n-card hoverable :embedded="true" class="root-card" title="主题设置" :bordered="false">
    </n-card>

    <n-alert type="warning" :bordered="false" style="margin-bottom: 20px">
      如果你不采用前后分离的方式部署，那么主题配置将不会生效。
      <link rel="了解前后分离">
    </n-alert>


  </div>

  <transition name="slide-fade">
    <div style="padding: 0 20px 20px 20px" v-if="animated">
      <n-card v-for="item in allThemes" :key="item.name" class="theme-card" :embedded="true" hoverable :bordered="false"
              content-style="padding: 0;">
        <div class="content">
          <div class="l-content">
            <p class="theme-name">{{ item.name }}
              <n-tag v-if="item.enabled" style="margin-left: 10px" :bordered="false" type="success">使用中</n-tag>
            </p>
            <p class="theme-desc">{{ item.describe }}</p>
          </div>
          <div class="r-content"
               :style="{backgroundImage: `url(${item.backgroundImageUrl})`, opacity: themeStore.enableDarkMode?0.5:1,}">
            <div class="r-content-inner"></div>
            <n-button tertiary type="primary" class="btn" @click="showModal=true">设为当前</n-button>
            <!--          <n-button tertiary type="primary" class="btn" @click="handleSetTheme(item.name)">配置</n-button>-->
          </div>
        </div>
      </n-card>
    </div>
  </transition>

  <n-modal
      title="编辑主题"
      v-model:show="showModal"
      preset="dialog"
      positive-text="确认"
      negative-text="取消"
      style="width: 480px;"
      @positive-click="message.info('positive')"
      @negative-click="message.warning('negative')"
      :show-icon="false"
  >

    <div style="margin-top: 30px"></div>
    <n-form
        ref="formRef"
        :model="formValue"
        :rules="rules"
    >

      <n-form-item
          :label="'给主题起一个名字'"
      >
        <n-input>

        </n-input>
      </n-form-item>

    </n-form>

  </n-modal>

  <div
      style="padding: 30px"
  >
    <n-table>
      <n-tr>
        <n-td
            style="column-span: 3; font-weight: bold; font-size: 1rem"
        >
          商品
        </n-td>
      </n-tr>
      <n-tr
          v-for="(item, index) in preSaleItems"
          :key="item.id"
      >
        <n-td>{{ item.name }}</n-td>
        <n-td>{{ item.price.toFixed(2) }}</n-td>
        <n-td>
          <n-button
              type="primary"
              secondary
              style="width: 100%"
              @click="addItem2ListByIndex(index)"
          >
            添加
          </n-button>
        </n-td>
      </n-tr>
    </n-table>

    <div style="height: 30px"></div>

    <n-table>
      <n-tr>
        <n-td style="column-span: 3; font-weight: bold; font-size: 1rem">
          商品列表
        </n-td>
      </n-tr>
      <n-tr
          v-for="item in addedItems"
          :key="item.id"
      >
        <n-td>{{ item.name }}</n-td>
        <n-td>{{ item.price.toFixed(2) }}</n-td>
      </n-tr>
    </n-table>
    <n-h3>总价 {{ getTotalPrice.toFixed(2) }}</n-h3>
  </div>


</template>

<style lang="less" scoped>
.root {
  padding: 20px;

  .root-card {
    margin-bottom: 10px;
  }


  //.n-card:hover {
  //  transform: translateY(-2px);
  //  transition: ease 200ms;
  //}

  .test {
    width: 100%;
    height: 100px;
  }

}

.theme-card {
  display: flex;
  height: 120px;
  margin-bottom: 20px;
  transition: ease 200ms;

  .content {
    display: flex;
    height: 120px;

    .l-content {
      flex: 2;
      height: 120px;
      display: flex;
      flex-direction: column;
      justify-content: center;
      padding-left: 30px;
      //background-color: #66afe9;

      .theme-name {
        font-size: 16px;
        font-weight: bold;
        margin-bottom: 10px;
      }

      .theme-desc {
        opacity: 0.7;
      }
    }

    .r-content::before {
      content: "";
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
      //background: linear-gradient(to right, rgba(255, 255, 255, 1) 0%, rgba(255, 255, 255, 0) 100%);
      background: v-bind('coverColor');
      //background-color: #66afe9;
      border-radius: 0 3px 3px 0;
    }

    .r-content {
      flex: 3;
      text-align: right;
      display: flex;
      border-radius: 0 3px 3px 0;
      line-height: 120px;
      justify-content: right;
      position: relative;
      background-size: cover;
      background-repeat: no-repeat;
      background-position: center;

      .r-content-inner {
        display: flex;
        flex-direction: row;
      }

      .btn {
        width: 100px;
        margin-right: 50px;
        position: absolute;
        top: 50%;
        transform: translateY(-50%);
        opacity: 1;
        backdrop-filter: blur(50px);
        -webkit-backdrop-filter: blur(5px);
      }
    }

  }


}

.theme-card:hover {
  transition: transform 200ms ease;
  transform: translateY(-5px);
}

</style>