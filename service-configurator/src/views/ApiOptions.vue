<script setup lang="ts">
import {onMounted} from "vue";
import useConfigStore from "@/store/useConfigStore";
import {instance} from "@/axios";
import {useMessage} from "naive-ui";
import {SaveOutline as saveIcon} from "@vicons/ionicons5"

const message = useMessage()
const configStore = useConfigStore();


let handleSubmitApiOptions = async () => {
  try {
    let {data} = await instance.post('/api/v1/config/server', {
      ...configStore.apiServerConfig
    })
    if (data.code === 200) {
      console.log('ok')
      message.success('保存ApiServer设置成功')
      // await handleGetRedisConfig()
    } else {
      message.error("保存设置出错：", data.msg)
    }
  } catch (error: any) {
    console.log(error)
    message.error(error)
  }
}

let handleGetApiOptions = async () => {
  try {
    let {data} = await instance.get('/api/v1/config/server')
    if (data.code === 200) {
      Object.assign(configStore.apiServerConfig, data.api_server_config)
      console.log('ok')
    }
  } catch (error: any) {
    console.log(error)
  }
}

onMounted(() => {
  handleGetApiOptions()
})

</script>

<script lang="ts">
export default {
  name: 'ApiOptions',
}
</script>

<template>
  <div class="root">
    <n-card class="card-bg" :embedded="true" :bordered="false" hoverable title="ApiServer"></n-card>

    <n-card class="card-bg" style="margin-top: 20px" :embedded="true" :bordered="false" hoverable>

      <div class="item">
        <div class="l-content">
          <span class="title">后端API服务器监听IP</span>
          <span class="describe">在这里设置后端API服务器的监听IP</span>
        </div>
        <div class="r-content">
          <n-input v-model:value="configStore.apiServerConfig.listen_addr" class="inp" size="large"
                   placeholder="请输入监听IP"></n-input>
        </div>
      </div>

      <div class="item">
        <div class="l-content">
          <span class="title">后端API服务器监听端口</span>
          <span class="describe">在这里设置后端API服务器的监听端口</span>
        </div>
        <div class="r-content">
          <n-input-number v-model:value.number="configStore.apiServerConfig.listen_port" class="inp" size="large"
                          placeholder="请输入监听端口"></n-input-number>
        </div>
      </div>

      <div class="save-part" style="display: flex; justify-content: right">
        <n-button style="width: 200px" size="large" type="primary" @click="handleSubmitApiOptions">
          <n-icon style="margin-right: 8px" size="18">
            <saveIcon/>
          </n-icon>
          保存设置
        </n-button>
      </div>


    </n-card>


  </div>
</template>

<style scoped lang="less">
.root {
  //margin: 20px;
}


.card-bg {
  //background-color: #f4f4f4;
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

  @media (max-width: 1000px) {
    .l-content {
      flex-direction: row;
      align-items: center;

      .describe {
        margin-left: 10px;
      }
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