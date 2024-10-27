<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {onBeforeMount, ref} from "vue"
import useApiAddrStore from "@/stores/useApiAddrStore";
import useUserInfoStore from "@/stores/useUserInfoStore";
import useThemeStore from "@/stores/useThemeStore";
import {useMessage} from "naive-ui";
import instance from "@/axios";
import {CloseOutline as closeIcon, CopyOutline as copyIcon} from "@vicons/ionicons5"

const {t} = useI18n();
const message = useMessage()
const apiAddrStore = useApiAddrStore();
const userInfoStore = useUserInfoStore();
const themeStore = useThemeStore();

const showModal = ref<boolean>(false)
const keyIndex = ref<number>(0)


interface Key {
  order_id: number;
  key_id: number;
  key: string;        // show
  plan_id: number;
  plan_name: string;       // show
  is_active: boolean;
  is_valid: boolean;
  is_used: boolean;
  client_id: string;
  expiration_date: string;
  created_at: string;
}

let myKeys = ref<Key[]>([])


// getAllMyKeys 获取所有用户的密钥
let getAllMyKeys = async () => {
  myKeys.value = []
  try {
    let {data} = await instance.get('/api/user/v1/keys', {
      params: {
        user_id: userInfoStore.thisUser.id,
        page: 1,
        size: 100,
      }
    })
    if (data.code === 200) {
      console.log(data)
      data.my_keys.forEach((key: Key) => myKeys.value.push(key))
      //
    }
  } catch (error) {
    console.log(error)
  }
}

// 复制成功的反馈
// const copySuccess = ref(false);

// 复制文本的函数
const copyText = async (key: string, event: MouseEvent) => {
  event.stopPropagation()
  try {
    await navigator.clipboard.writeText(key);
    message.success('复制成功')

    // copySuccess.value = true;
    // setTimeout(() => {
    //   // copySuccess.value = false;
    //   message.success('复制成功')
    // }, 2000); // 显示2秒成功信息
  } catch (err) {
    console.error("复制失败:", err);
  }
};

onBeforeMount(() => {
  themeStore.userPath = '/dashboard/keys'
  themeStore.menuSelected = 'user-keys'
  getAllMyKeys()
})

</script>

<script lang="ts">
export default {
  name: 'UserKeys',
}
</script>

<template>
  <div class="root">
    <n-card class="title" :embedded="true" hoverable :bordered="false" :title="t('userKeys.myKeys')">
    </n-card>

    <n-card
        @click="showModal = true; keyIndex = index"
        class="key-card"
        :embedded="true"
        hoverable
        :bordered="false"
        content-style="padding: 0;"
        v-for="(item, index) in myKeys"
        :key="item.key_id"
    >
      <div class="key-item">
        <div class="l-part">
          <p class="plan-name">{{ item.plan_name }}</p>
          <!--          <p class="key-id">{{ item.key }}</p>-->
          <n-popover trigger="hover" placement="bottom" :show-arrow="false">
            <template #trigger>
              <!--              <n-button>悬浮</n-button>-->
              <!--                        <span class="key-id" style="width: auto">{{ item.key }}</span>-->
              <n-tag
                  class="key-tag"
                  style="margin: 5px 0 5px 0;"
                  size="large"
                  type="default"
                  :bordered="true"
                  checkable
                  @click="copyText(item.key, $event)"
              >

                <div class="key-inner" style="display: flex; flex-direction: row; align-items: center;">
                  {{ item.key }}
                  <n-icon class="copy-icon" size="16" style="margin: 0 10px">
                    <copyIcon/>
                  </n-icon>
                </div>

              </n-tag>
            </template>
            <span>点击可以复制到剪贴板</span>
          </n-popover>
          <div style="display: flex; flex-direction: row">
            <n-tag :bordered="false" size="small" style="margin-right: 10px" :type="item.is_active?'success':'error'">
              {{ item.is_active ? '订阅活躍' : '订阅非活跃' }}
            </n-tag>
            <n-tag :bordered="false" size="small" style="margin-right: 10px" :type="item.is_valid?'success':'warning'">
              {{ item.is_valid ? '密钥有效' : '密钥不可用' }}
            </n-tag>
            <n-tag :bordered="false" size="small" :type="item.is_used?'info':'success'">
              {{ item.is_used ? '已激活使用' : '待使用' }}
            </n-tag>
          </div>
        </div>
        <div class="r-part">
        </div>
      </div>
    </n-card>
  </div>

  <n-modal v-model:show="showModal">
    <n-card
        class="details-card"
        :title="t('userKeys.keyDetail')"
        :bordered="false"
        size="huge"
        role="dialog"
        aria-modal="true"
    >
      <template #header-extra>
        <n-button class="close-btn" @click="showModal = false" text :bordered="false" type="primary">
          <n-icon size="25">
            <closeIcon/>
          </n-icon>
        </n-button>
      </template>

      <div class="key-item-detail">
        <p class="key-item-title">{{ t('userKeys.keyId') }}</p>
        <p class="key-item-content">{{ myKeys[keyIndex].key_id }}</p>
      </div>

      <div class="key-item-detail">
        <p class="key-item-title">{{ t('userKeys.orderId') }}</p>
        <p class="key-item-content">{{ myKeys[keyIndex].order_id }}</p>
      </div>

      <div class="key-item-detail">
        <p class="key-item-title">{{ t('userKeys.releaseData') }}</p>
        <p class="key-item-content">{{ myKeys[keyIndex].created_at }}</p>
      </div>

      <div class="key-item-detail">
        <p class="key-item-title">{{ t('userKeys.expirationData') }}</p>
        <p class="key-item-content">{{ myKeys[keyIndex].expiration_date }}</p>
      </div>

      <div class="key-item-detail">
        <p class="key-item-title">{{ t('userKeys.clientId') }}</p>
        <p class="key-item-content">{{
            myKeys[keyIndex].client_id ? myKeys[keyIndex].client_id : t('userKeys.none')
          }}</p>
      </div>


      <template #footer>

      </template>
    </n-card>
  </n-modal>


</template>

<style scoped lang="less">
.root {
  padding: 20px;

  .title {

  }

  .key-card {
    margin-top: 20px;

    .key-item {
      padding: 20px;
      display: flex;
      flex-direction: row;
      flex-wrap: wrap;

      .l-part {
        flex: 3;

        .plan-name {
          font-size: 1rem;
          opacity: 0.8;
          font-weight: 400;
        }

        .key-id {
          font-size: 1.2rem;
          opacity: 0.8;
          padding-bottom: 5px;
        }

        .order-id {
          font-size: 1.1rem;
          opacity: 0.8;
        }
      }

      .r-part {
        flex: 0;
      }
    }

    transition: ease 300ms;
  }

  .key-card:hover {
    transform: translateX(0) translateY(-5px);
  }
}

.close-btn {
  transition: ease 300ms;
}

.close-btn:hover {
  transform: rotate(90deg);
}

.key-item-detail {
  //background-color: #66afe9;

  .key-item-title {
    font-size: 1rem;
    font-weight: 600;
    opacity: 0.7;
    margin-bottom: 5px;
  }

  .key-item-content {
    font-size: 1.2rem;
    font-weight: 600;
    opacity: 1;
    margin-bottom: 25px;
  }
}

.details-card {
  width: 50%;
}

@media (max-width: 768px) {
  .details-card {
    width: 80%;
  }
}

.key-tag {
  .key-inner {
    .copy-icon {
      transition: ease 400ms;
      opacity: 0;
    }
  }
}

.key-tag:hover {
  .key-inner {
    .copy-icon {
      opacity: 0.5;
    }
  }
}

</style>