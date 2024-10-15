<script setup lang="ts">
import {ref} from "vue"
import useConfigStore from "@/store/useConfigStore";
import {SaveOutline as saveIcon} from "@vicons/ionicons5"

const configStore = useConfigStore();

</script>

<script lang="ts">
export default {
  name: 'EtcdOptions',
}
</script>

<template>
  <div class="root" style="">
    <n-card class="card-bg" :embedded="true" :bordered="true" hoverable title="Etcd配置"></n-card>

    <n-alert type="info" style="margin: 20px 0">
      Etcd的配置仅保存在本地localStorage
    </n-alert>

    <n-card class="card-bg" style="margin-top: 20px" :embedded="true" :bordered="true" hoverable>


      <div class="item">
        <div class="l-content">
          <span class="title">终端节点地址</span>
          <span class="describe">这里是Etcd的客户端域名或IP</span>
        </div>
        <div class="r-content">
          <n-input v-model:value="configStore.etcdConfig.endpoints" class="inp" size="large" placeholder="请输入节点名，多个使用空格隔开"></n-input>
        </div>
      </div>

      <div class="item">
        <div class="l-content">
          <span class="title">通信端口</span>
          <span class="describe">这里是Etcd服务器的连接端口，它通常为2379或2380</span>
        </div>
        <div class="r-content">
          <n-input-number v-model:value.number="configStore.etcdConfig.port" class="inp" size="large" placeholder="请输入节点的端口，一般为2379"></n-input-number>
        </div>
      </div>

      <div class="item">
        <div class="l-content">
          <span class="title">用户名</span>
          <span class="describe">如果启用了认证，则需要提供用户名</span>
        </div>
        <div class="r-content">
          <n-input v-model:value="configStore.etcdConfig.username" class="inp" size="large" placeholder="输入用户名（非必填）"></n-input>
        </div>
      </div>

      <div class="item">
        <div class="l-content">
          <span class="title">密码</span>
          <span class="describe">如果启用了认证，则需要提供密码以通过认证</span>
        </div>
        <div class="r-content">
          <n-input v-model:value="configStore.etcdConfig.password" class="inp" size="large" placeholder="输入密码（非必填）"></n-input>
        </div>
      </div>

      <div class="save-part" style="display: flex; justify-content: right">
        <n-button style="width: 200px" size="large" type="primary" @click="configStore.saveEtcdConfig2LocalStorage">
          <n-icon style="margin-right: 8px" size="18"><saveIcon/></n-icon>
          保存设置
        </n-button>
      </div>

    </n-card>




  </div>
</template>

<style scoped lang="less">
.root {

}

.card-bg {
  background-color: #f4f4f4;
}

.item {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  margin: 10px 0 40px 0;

  .l-content {
    flex: 3;
    //background-color: #66afe9;
    display: flex;
    flex-direction: column;

    .title {
      font-size: 1rem;
      font-weight: 700;
    }

    .describe {
      opacity: 0.6;
    }
  }

  .r-content {
    flex: 3;
    //background-color: #4cae4c;
    display: flex;
    justify-content: center;
    align-items: center;

    .inp {
      width: 100%;
    }
  }
}

@media (max-width: 1000px) {
  .item {
    flex-direction: column;
    .r-content {
      margin-top: 10px;
    }
  }
}

</style>