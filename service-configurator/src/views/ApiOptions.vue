<script setup lang="ts">
import useConfigStore from "@/store/useConfigStore";

const configStore = useConfigStore();
</script>

<script lang="ts">
export default {
  name: 'ApiOptions',
}
</script>

<template>
  <div class="root">
    <n-card class="card-bg" :embedded="true" :bordered="true" hoverable title="ApiServer"></n-card>

    <n-card class="card-bg" style="margin-top: 20px" :embedded="true" :bordered="true" hoverable>

      <div class="item">
        <div class="l-content">
          <span class="title">后端API服务器监听IP</span>
          <span class="describe">在这里设置后端API服务器的监听IP</span>
        </div>
        <div class="r-content">
          <n-input v-model:value="configStore.redisConfig.host" class="inp" size="large" placeholder="请输入监听IP"></n-input>
        </div>
      </div>

      <div class="item">
        <div class="l-content">
          <span class="title">后端API服务器监听端口</span>
          <span class="describe">在这里设置后端API服务器的监听端口</span>
        </div>
        <div class="r-content">
          <n-input-number v-model:value.number="configStore.redisConfig.port" class="inp" size="large" placeholder="请输入监听端口"></n-input-number>
        </div>
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