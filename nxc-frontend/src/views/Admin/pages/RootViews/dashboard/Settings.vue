<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {useRouter} from 'vue-router';
import {
  BagHandleSharp as SubIcon,
  PeopleSharp as UserIcon,
  ReorderFourSharp as OrderIcon,
  SettingsSharp as settingIcon,
  Card as payIcon,
  Key as keyIcon,
} from '@vicons/ionicons5'

const {t} = useI18n()
const router = useRouter();

</script>

<script lang="ts">
export default {
  name: 'Settings',
}
</script>

<template>
    <div class="root">
      <n-grid
          cols="2 s:2 m:4"
          responsive="screen"
          :x-gap="15"
          :y-gap="15"
      >
        <n-grid-item>
          <n-card
              hoverable
              @click="router.push({ path: '/admin/dashboard/systemconfig'})"
              :embedded="true"
              :bordered="false"
              class="shortcut-item"
          >
            <div>
              <n-icon size="25">
                <settingIcon/>
              </n-icon>
              <p style="margin-top: 8px; font-size: 15px">{{ t('adminViews.summary.systemConfig') }}</p>
            </div>
          </n-card>
        </n-grid-item>
<!--        <n-grid-item>-->
<!--          <n-card-->
<!--              hoverable style="margin-right: 20px"-->
<!--              :embedded="true"-->
<!--              :bordered="false"-->
<!--              class="shortcut-item"-->
<!--          >-->
<!--            <div>-->
<!--              <n-icon size="25">-->
<!--                <payIcon/>-->
<!--              </n-icon>-->
<!--              <p style="margin-top: 8px; font-size: 15px">{{ t('adminViews.summary.paymentConfig') }}</p>-->
<!--            </div>-->
<!--          </n-card>-->
<!--        </n-grid-item>-->
        <n-grid-item>
          <n-card
              hoverable
              :embedded="true"
              :bordered="false"
              class="shortcut-item"
              @click="router.push({path: '/admin/dashboard/subscribemanager'})"
          >
            <div>
              <n-icon size="25">
                <OrderIcon/>
              </n-icon>
              <p style="margin-top: 8px; font-size: 15px">{{ t('adminViews.summary.planMgr') }}</p>
            </div>
          </n-card>
        </n-grid-item>
        <n-grid-item>
          <n-card
              hoverable
              style="margin-right: 20px"
              :embedded="true"
              :bordered="false"
              @click="router.push({ path: '/admin/dashboard/usermanager'})"
              class="shortcut-item"
          >
            <div>
              <n-icon size="25">
                <UserIcon/>
              </n-icon>
              <p style="margin-top: 8px; font-size: 15px">{{ t('adminViews.summary.userMgr') }}</p>
            </div>
          </n-card>
        </n-grid-item>
        <n-grid-item>
          <n-card
              hoverable
              :embedded="true"
              :bordered="false"
              @click="router.push({ path: '/admin/dashboard/order'})"
              class="shortcut-item"
          >
            <div>
              <n-icon size="25">
                <SubIcon/>
              </n-icon>
              <p style="margin-top: 8px; font-size: 15px">{{ t('adminViews.summary.orderMgr') }}</p>
            </div>
          </n-card>
        </n-grid-item>
<!--        <n-grid-item>-->
<!--          <n-card-->
<!--              hoverable style="margin-right: 20px"-->
<!--              :embedded="true"-->
<!--              :bordered="false"-->
<!--              class="shortcut-item"-->
<!--          >-->
<!--            <div>-->
<!--              <n-icon size="25">-->
<!--                <keyIcon/>-->
<!--              </n-icon>-->
<!--              <p style="margin-top: 8px; font-size: 15px">{{ t('adminViews.summary.keyMgr') }}</p>-->
<!--            </div>-->
<!--          </n-card>-->
<!--        </n-grid-item>-->
      </n-grid>
    </div>

</template>

<style lang="less" scoped>
.root {
  padding: 20px 20px 0 20px;
  display: flex;
  flex-direction: row;
  text-align: center;
}
.shortcut-item {
  display: flex;
  transition: ease 200ms;
}

.shortcut-item:hover {
  transform: translateX(0px) translateY(-3px);
  filter: brightness(0.98);
}

</style>