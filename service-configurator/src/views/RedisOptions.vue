<script setup lang="ts">
import {onMounted} from "vue";
import useConfigStore from "@/store/useConfigStore";
import {instance} from "@/axios"
import {SaveOutline as saveIcon} from "@vicons/ionicons5"


const configStore = useConfigStore();

let handleSubmitRedisConfig = async () => {
  try {
    let {data} = await instance.post('/api/v1/config/redis', {
      ...configStore.redisConfig
    })
    if (data.code === 200) {
      console.log('ok')
    }
  } catch (error: any) {
    console.log(error)
  }
}

let handleGetRedisConfig = async () => {
  try {
    let {data} = await instance.get('/api/v1/config/redis')
    if (data.code === 200) {
      console.log('ok')
      Object.assign(configStore.redisConfig, data.redis_config)
    }
  } catch (error: any) {
    console.log(error)
  }
}

onMounted(() => {
  handleGetRedisConfig()
})

</script>

<script lang="ts">
export default {
  name: 'RedisOptions',
}
</script>

<template>
  <div class="root">
    <n-card class="card-bg" :embedded="true" :bordered="true" hoverable title="Redis配置"></n-card>

    <n-card class="card-bg" style="margin-top: 20px" :embedded="true" :bordered="true" hoverable>

      <div class="item">
        <div class="l-content">
          <span class="title">Redis连接地址</span>
          <span class="describe">这里是Redis服务器的域名或IP</span>
        </div>
        <div class="r-content">
          <n-input v-model:value="configStore.redisConfig.host" class="inp" size="large"
                   placeholder="请输入节点名，多个使用空格隔开"></n-input>
        </div>
      </div>

      <div class="item">
        <div class="l-content">
          <span class="title">Redis端口</span>
          <span class="describe">这里是Redis服务器的连接端口，它通常为6379</span>
        </div>
        <div class="r-content">
          <n-input-number v-model:value.number="configStore.redisConfig.port" class="inp" size="large"
                          placeholder="请输入数据库的连接端口，一般为6379"></n-input-number>
        </div>
      </div>

      <div class="item">
        <div class="l-content">
          <span class="title">用户名</span>
          <span class="describe">如果启用了认证，则需要提供用户名</span>
        </div>
        <div class="r-content">
          <n-input v-model:value="configStore.redisConfig.username" class="inp" size="large"
                   placeholder="输入用户名（非必填）"></n-input>
        </div>
      </div>

      <div class="item">
        <div class="l-content">
          <span class="title">密码</span>
          <span class="describe">如果启用了认证，则需要提供密码以通过认证</span>
        </div>
        <div class="r-content">
          <n-input v-model:value="configStore.redisConfig.password" class="inp" size="large"
                   placeholder="请输入用户的密码（非必填）"></n-input>
        </div>
      </div>

      <div class="item">
        <div class="l-content">
          <span class="title">数据库名</span>
          <span class="describe">这里是操作的Redis数据库名</span>
        </div>
        <div class="r-content">
          <n-input-number v-model:value.number="configStore.redisConfig.database" class="inp" size="large"
                          placeholder="请输入数据库名"></n-input-number>
        </div>
      </div>

      <!--      <div></div>-->

      <div class="save-part" style="display: flex; justify-content: right">
        <n-button style="width: 200px" size="large" type="primary" @click="handleSubmitRedisConfig">
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