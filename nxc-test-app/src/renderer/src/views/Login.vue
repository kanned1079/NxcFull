<script setup lang="ts">
defineOptions({
  name: 'Login'
})
import {ref, onMounted} from "vue"
import {type FormInst, type FormRules, useMessage} from "naive-ui"
// import PageHead from "./utils/PageHead.vue"
// import {useRouter} from "vue-router";

// import {ipcRenderer} from "electron"


// import instance from "./../axios"
//
// const router = useRouter()
type LoginValues = {
  email: string
  password: string
  key_content: string

}
//

const message = useMessage()
let loginValue  = ref<LoginValues>({
  email: '',
  password: '',
  key_content: '',
})



// const submitLogin = () => {
//   console.log(loginValue.value)
//   try {
//     router.push({path: '/home'})
//     let {data} = instance.post('/api/user/v1/activation/bind', {
//       ...loginValue.value
//     })
//     console.log(data)
//     if (data && data.code === 200) {
//
//     }
//   }
//
//   catch (err: any) {
//   }
// }
let formRef = ref<FormInst | null>(null)

let formRules: FormRules = {
  email: {
    type: 'email',
    required: true,
    trigger: ['blur', 'change']
  },
  password: {
    type: 'string',
    required: true,
    trigger: ['blur', 'change']
  },
  key_content: {
    type: 'string',
    required: true,
    trigger: ['blur', 'change']
  }
}

const submitClick = async () => {
  if (formRef.value) {
    await formRef.value.validate(async (err: any) => {
      if (!err) {
        console.log('ok')
        try {
          // let {data} = await instance.post('/api/user/v1/activation/bind', {...loginValue.value})
          // const resp = await ipcRenderer.invoke('login', {...loginValue.value})

          // const resp = await window.electron.login({...loginValue.value})
           window.electron.ipcRenderer.send('login', {...loginValue.value})

          // console.log(resp)

        } catch (err:any) {
          console.error(err)
          message.error('请求发生错误请重新尝试')
        }
      } else {
        message.error('表单验证失败')
        // console.log('err')
      }
    })
  }
}


onMounted(() => {
  // submitLogin()
})

</script>

<template>
<div class="root">

  <div
    class="root-inner"
  >
    <n-card
      hoverable
      :embedded="true"
      :bordered="false"
      content-style="padding: 20px"
    >
      <div class="login-title-root">
        <p class="login-title">欢迎</p>
        <p class="login-desc">使用本软件前需要授权</p>
      </div>

      <div>
        <n-form
          :show-label="true"
          :show-require-mark="true"
          :show-feedback="false"
          :model="loginValue"
          ref="formRef"
          :rules="formRules"
        >
          <n-form-item class="bottom-gap" label="邮箱地址" path="email">
            <n-input
              clearable
              placeholder="输入您的邮箱地址"
              v-model:value="loginValue.email"
            ></n-input>
          </n-form-item>

          <n-form-item class="bottom-gap" label="密码" path="password">
            <n-input
              clearable
              placeholder="输入您的密码"
              type="password"
              v-model:value="loginValue.password"
            ></n-input>
          </n-form-item>

          <n-form-item class="bottom-gap" label="授权密钥" path="key_content">
            <n-input
              clearable
              placeholder="输入有效的激活密钥"
              v-model:value="loginValue.key_content"
            ></n-input>
          </n-form-item>
        </n-form>


        <n-button class="submit-btn" type="info" secondary @click="submitClick">确认</n-button>
      </div>


    </n-card>



  </div>


</div>
</template>

<style scoped lang="less">
.root {
  height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  background: linear-gradient(to right, rgb(190, 147, 197), rgb(123, 198, 204));
  .root-inner {
    width: 320px;
    height: 400px;
    transition: all ease-in-out 200ms;
  }
  .root-inner:hover {
    transform: translateX(0) translateY(-3px);
  }

}

.login-title-root {
  margin-bottom: 20px;
  .login-title {
    font-size: 1.2rem;
    font-weight: bold;
  }
  .login-desc {
    opacity: 0.8;
  }
}

.bottom-gap {
  margin-bottom: 10px;
}

.submit-btn {
  width: 100%;
  margin-top: 30px;
}
</style>
