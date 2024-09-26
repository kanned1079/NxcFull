<script setup lang="ts">
import {useI18n} from "vue-i18n";
import {ref, onMounted} from "vue";

const {t} = useI18n()

import {LogoApple as appleIcon, LogoGithub as githubIcon, LogoGoogle as googleIcon,} from "@vicons/ionicons5"
import instance from "@/axios";

let formValue = ref<{email: string, password: string}>({
  email: '',
  password: '',
})

let loginBtn = () => {
  console.log(formValue.value)
}

let loginViaGithub = async () => {
  try {
    let  {data} = await instance.get('https://api.github.com/octocat')
    console.log(data)
  }catch (error) {
    console.log(error)
  }
}

onMounted(() => {

})

</script>

<script lang="ts">
export default {
  name: 'WhyChooseUsA',
}
</script>


<template>
  <div class="root">
    <div class="l-content">
      <div class="l-content-top"></div>
      <div class="l-content-color1"></div>
      <div class="l-content-color2"></div>
      <div class="l-content-color3"></div>
      <div class="l-content-color4"></div>
      <div class="l-container">
        <p class="title">Nxc Networks</p>
        <p class="sub">穿过县境上长长的隧道，便是雪国。夜空下，大地一片莹白，火车在信号所前停下来。 </p>
      </div>
    </div>
    <div class="r-content">
      <div class="r-content-color-cover"></div>
      <div class="r-content-color1"></div>
      <div class="r-content-color2"></div>
      <div class="r-content-color3"></div>
      <n-card class="login-card" :embedded="false" :bordered="false">
        <p class="login-title">登入以继续</p>
        <n-input :bordered="false" type="text" class="username" placeholder="邮箱" v-model:value="formValue.email"></n-input>
        <n-input :bordered="false" type="password" class="password" placeholder="密码" v-model:value="formValue.password"></n-input>
        <div class="no-account">
          <p class="no-account-prefix">还没有账号？</p>
          <n-button class="no-account-btn" text>立即注册</n-button>
        </div>
        <n-button :bordered="false" type="primary" class="login-btn" @click="loginBtn">登入</n-button>
        <n-divider>
          <p style="opacity: 0.5;  font-weight: 400; font-size: 0.8rem">或使用其他登录方式</p>
        </n-divider>
        <n-button :bordered="false" class="other-login-method github" type="primary" @click="loginViaGithub ">
          <n-icon style="margin-right: 10px;" size="25">
            <githubIcon/>
          </n-icon>
          Sign in with Github
        </n-button>
        <n-button :bordered="false" class="other-login-method apple" type="primary">
          <n-icon style="margin-right: 10px;" size="25">
            <appleIcon/>
          </n-icon>
          Sign in with Apple &nbsp;
        </n-button>
        <n-button :bordered="false" class="other-login-method google" type="primary">
          <n-icon style="margin-right: 10px;" size="25">
            <googleIcon/>
          </n-icon>
          Sign in with Google
        </n-button>
      </n-card>
    </div>
  </div>
</template>

<style scoped lang="less">
.root {
  height: 100vh;
  display: flex;
  position: relative;
}

.l-content {
  flex: 1;
  position: relative;
  overflow: hidden;
  //background-color: #292929;

  .l-container {
    position: absolute;
    z-index: 1001;
    top: 40%;
    left: 15%;
    //background-color: #4cae4c;
    width: 70%;
    transition: ease 300ms;


    .title {
      font-size: 2.5rem;
      font-weight: 700;
      opacity: 0.8;
      color: #205387;
    }

    .sub {
      color: #136685;
      margin-top: 20px;
      font-size: 1.2rem;
      opacity: 0.5;
    }
  }

  .l-container:hover {
    transform: translateX(0) translateY(-10px);
  }

  .l-content-top {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    backdrop-filter: blur(100px);
    -webkit-backdrop-filter: blur(100px);
    z-index: 1000;
  }

  .l-content-color1 {
    width: 500px;
    height: 500px;
    border-radius: 50%;
    background-color: rgba(225, 107, 140, 0.5);
    position: absolute;
    //top: calc(70% - 400px / 2);
    //left: calc(30% - 400px / 2);
    //bottom: calc(100% - 100px);
    left: -250px;
    bottom: -250px;
    z-index: 1;
  }

  .l-content-color2 {
    width: 300px;
    height: 300px;
    border-radius: 50%;
    //background-color: rgba(84, 234, 255, 0);
    background: linear-gradient(to right, rgb(127, 127, 213), rgb(134, 168, 231), rgb(145, 234, 228));
    position: absolute;
    top: calc(40% - 400px / 2);
    left: calc(45% - 400px / 2);
    z-index: 2;
  }

  .l-content-color3 {
    width: 300px;
    height: 300px;
    border-radius: 50%;
    //background-color: rgba(220, 135, 227, 0);
    background: linear-gradient(to right, rgb(194, 229, 156), rgb(100, 179, 244));
    position: absolute;
    top: calc(50% - 400px / 2);
    left: calc(60% - 400px / 2);
    z-index: 3;
  }

  .l-content-color4 {
    width: 300px;
    height: 300px;
    border-radius: 50%;
    //background-color: rgba(220, 135, 227, 0);
    background: linear-gradient(to right, rgb(217, 167, 199), rgb(255, 252, 220));
    position: absolute;
    top: calc(10% - 400px / 2);
    left: calc(80% - 400px / 2);
    z-index: 3;
  }
}

@media (max-width: 900px) {
  .l-content {
    display: flex;
    flex: 0;
  }
}

.r-content {
  flex: 1;
  //background-color: #c8d9eb;
  //position: relative;
  display: flex;
  flex-direction: row;

  overflow: hidden;
  justify-content: center;
  align-items: center;

  .login-card {
    z-index: 2000;
    display: flex;
    //background-color: rgba(218, 144, 144, 0.5);
    //background-color: #66afe9;
    width: 75%;

    .login-title {
      font-size: 2rem;
      font-weight: 500;
      margin-bottom: 10px;
      opacity: 0.9;
    }

    .username {
      margin-top: 20px;
      height: 45px;
      border-radius: 10px;
      align-items: center;
      background-color: #f2fafd;
    }

    .password {
      margin-top: 30px;
      height: 45px;
      border-radius: 10px;
      align-items: center;
      background-color: #f2fafd;
    }

    .no-account {
      margin-top: 10px;
      display: flex;
      flex-direction: row;
      justify-content: right;

      .no-account-prefix {
        font-size: 1rem;
      }

      .no-account-btn {
        opacity: 0.5;
        font-size: 1rem;
      }
    }

    .login-btn {
      margin-top: 50px;
      width: 100%;
      height: 45px;
      color: white;
      background: linear-gradient(to right, rgba(167, 112, 239, 0.5), rgba(207, 139, 243, 0.9), rgba(253, 185, 155, 0.5));
      box-shadow: rgba(139, 129, 195, 0.5) 3px 3px 200px 0;
      margin-bottom: 20px;
    }

    .other-login-method {
      width: 100%;
      height: 45px;
      margin-bottom: 20px;
      transition: ease 200ms;
    }

    .other-login-method:hover {
      transform: translateX(0) translateY(-3px);
    }

    .github {
      color: white;
      background-color: #3a4754;
      box-shadow: #3a4754 1px 1px 5px 0;
    }

    .google {
      color: white;
      background-color: #3382f0;
      box-shadow: #3382f0 1px 1px 5px 0;
    }

    .apple {
      color: white;
      background-color: #000;
      box-shadow: #000 1px 1px 5px 0;
    }
  }

  @media (max-width: 900px) {
    .login-card {
      width: 70%;
    }
  }
  @media (max-width: 768px) {
    .login-card {
      width: 90%;
    }
  }
}

@media (max-width: 900px) {
  .r-content {
    //background: linear-gradient(to right, rgb(217, 167, 199), rgb(255, 252, 220));

  }
}
</style>