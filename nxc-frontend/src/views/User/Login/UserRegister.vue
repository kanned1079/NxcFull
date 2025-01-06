<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { computed, onMounted, ref } from "vue";
import { useRouter } from "vue-router";
import useThemeStore from "@/stores/useThemeStore";
import { useMessage } from "naive-ui";
import type { FormInst, FormItemInst, FormItemRule, FormRules } from 'naive-ui'
import { ChevronBackOutline as backIcon } from "@vicons/ionicons5";
import useAppInfosStore from "@/stores/useAppInfosStore";
import { handleSendVerifyCode, handleVerifyCode, handleFinalRegister } from "@/api/user/register";

const { t } = useI18n();
const appInfosStore = useAppInfosStore();
const themeStore = useThemeStore();
const message = useMessage();
const router = useRouter();

const formRef = ref<FormInst | null>(null);

let verifyCodePassed = ref<boolean>(false);
let animated = ref<{ leftAnimated: boolean; rightAnimated: boolean }>({
  leftAnimated: false,
  rightAnimated: false,
});
let bgColor = computed(() =>
    themeStore.enableDarkMode
        ? { backgroundColor: "rgba(40, 41, 41, 1)" }
        : { backgroundColor: "#fff" }
);
let coverBgColor = computed(() =>
    themeStore.enableDarkMode
        ? { backgroundColor: "rgba(40, 40, 40, 0.2)" }
        : { backgroundColor: "rgba(255, 255, 255, 0.0)" }
);
let placeholderBgColor = computed(() =>
    themeStore.enableDarkMode
        ? {}
        : { backgroundColor: "#f2fafd", height: "45px" }
);
let textLeftColor = computed(() => (themeStore.enableDarkMode ? { color: "#fff" } : {}));

let agreementChecked = ref<boolean>(false); // 同意按钮是否被选中
let enableRegister = ref<boolean>(true); // 由检查逻辑控制
let regBtnEnabled = computed(() => agreementChecked.value && enableRegister.value);
let clickedCount = ref<number>(0); // 注册按钮点击次数
let enableSendCode = ref<boolean>(true);
let waitSendMail = ref<number>(0);

let formValue = ref({
  user: {
    email: "",
    verify_code: "",
    password: "",
    password_confirmation: "",
    invite_user_id: "",
  },
});

// 表单验证规则
let rules: FormRules = {
  user: {
    email: {
      required: true,
      message: "邮箱不能为空",
      type: "email",
      trigger: "blur",
      validator: async (rule: FormItemRule, value: string) => {
        if (!value) {
          return new Error("请输入邮箱");
        }

        // 调用 validateEmail 函数进行验证
        const isValid = validateEmail(value);
        if (!isValid) {
          return new Error("邮箱格式不正确");
        }

        return true; // 验证通过
      },
    },
    verify_code: {
      required: appInfosStore.registerPageConfig.is_email_verify,
      message: "验证码不能为空",
      trigger: "blur",
      validator: async (rule: FormItemRule, value: string) => {
        if (!appInfosStore.registerPageConfig.is_email_verify) {
          return true; // 不需要验证码时直接通过
        }
        if (!value) {
          return new Error("请输入验证码");
        }
        const codeRegex = /^\d{6}$/;
        if (!codeRegex.test(value)) {
          return new Error("验证码格式不正确");
        }

        return true; // 验证通过
      },
    },
    password: {
      required: true,
      message: "密码不能为空",
      trigger: "blur",
      validator: async (rule: FormItemRule, value: string) => {
        if (!value) {
          return new Error("请输入密码");
        }
        // 校验密码长度
        if (value.length < 6) {
          return new Error("密码长度不能少于6位");
        }
        return true;
      },
    },
    password_confirmation: {
      required: true,
      message: "请确认密码",
      trigger: "blur",
      validator: async (rule: FormItemRule, value: string) => {
        if (!value) {
          return new Error("请输入确认密码");
        }
        // 校验与密码是否一致
        if (value !== formValue.value.user.password) {
          return new Error("两次密码输入不一致");
        }
        return true;
      },
    },
    invite_user_id: {
      required: false,
    },
  },
};

// 邮箱格式验证
let validateEmail = (email: string): boolean => {
  const emailRegex = /^[0-9a-zA-Z_.-]+@[0-9a-zA-Z_.-]+(\.[a-zA-Z]+){1,2}$/;
  if (!emailRegex.test(email)) {
    message.error("邮箱格式不正确");
    return false;
  }
  const [localPart, domain] = email.split("@");
  const isGmail = domain.toLowerCase() === "gmail.com";
  const isGooglemail = domain.toLowerCase() === "googlemail.com";

  if ((isGmail || isGooglemail) && appInfosStore.registerPageConfig.email_gmail_limit_enable) {
    if (localPart.includes("+")) {
      return false;
    }
    const normalizedLocalPart = localPart.replace(/\./g, "");
    if (normalizedLocalPart !== localPart) {
      return false;
    }
    if (localPart !== localPart.toLowerCase()) {
      return false;
    }
    if (isGooglemail) {
      return false;
    }
  }
  return true;
};

// 发送邮箱验证码
let sendVerifyCode = async () => {
  if (formValue.value.user.email === "" || !validateEmail(formValue.value.user.email.trim())) {
    message.error(t("userRegister.invalidEmailFormat"));
    return;
  } else {
    let data = await handleSendVerifyCode(formValue.value.user.email.trim());
    if (data.code === 200) {
      enableSendCode.value = false;
      startWait(60);
      message.success(t("userRegister.sendSuccess"));
    } else {
      message.error(data.error.toString() || "err on send verify code");
    }
  }
};

// 验证验证码
let verifyCode = async () => {
  let data = await handleVerifyCode(formValue.value.user.email.trim(), formValue.value.user.verify_code.trim());
  if (data && data.code === 200 && data.passed) {
    verifyCodePassed.value = data.passed as boolean;
    console.log("验证码正确");
  } else {
    message.error("验证码错误");
  }
};

// 处理注册
let handleRegister = async () => {
  let data = await handleFinalRegister(
      formValue.value.user.email.trim(),
      formValue.value.user.password.trim(),
      formValue.value.user.invite_user_id
  );
  if (data.code === 200 && data.is_registered) {
    message.success(t("userRegister.regSuccess"));
    setTimeout(async () => {
      await router.push({ path: "/login" });
    }, 1000);
  } else if (data.code === 409) {
    message.error("该邮箱已经被注册");
  } else {
    message.error(t("userRegister.regFailure") + data.msg);
  }
};

// 按钮点击次数限制
let handleValidateClick = async () => {
  clickedCount.value += 1;
  if (clickedCount.value === 5) {
    enableRegister.value = false;
    setTimeout(() => {
      enableRegister.value = true;
      clickedCount.value = 0;
    }, 60000);
    return;
  }

  if (!formRef.value) {
    message.error("Form reference is not defined");
    return;
  }

  try {
    // 验证表单
    await formRef.value.validate();

    // 验证邮箱
    if (appInfosStore.registerPageConfig.is_email_verify && !verifyCodePassed.value) {
      message.error(t("userRegister.verifyCodeErr"));
      return;
    }

    // 执行注册逻辑
    await handleRegister();
  } catch (errors) {
    message.error("表单验证未通过:" + errors);
  }
};

// 开始等待验证码倒计时
let startWait = (sec: number) => {
  let countdown = sec;
  let intervalId = setInterval(() => {
    if (countdown >= 0) {
      waitSendMail.value = countdown;
      countdown--;
    } else {
      clearInterval(intervalId);
      enableSendCode.value = true;
    }
  }, 1000);
};

onMounted(async () => {
  const queryParams = new URLSearchParams(window.location.search);
  const code = queryParams.get("code");
  code ? (formValue.value.user.invite_user_id = code) : null;

  setTimeout(() => (animated.value.leftAnimated = true), 100);
  setTimeout(() => (animated.value.rightAnimated = true), 150);
  console.log("邮箱验证", appInfosStore.registerPageConfig);
});
</script>

<script lang="ts">
export default {
  name: 'UserRegister',
}
</script>

<template>
  <div class="root">
    <div class="l-content" :style="bgColor">
      <div class="l-content-top" :style="coverBgColor"></div>
      <div class="l-content-color1"></div>
      <div class="l-content-color2"></div>
      <div class="l-content-color3"></div>
      <div class="l-content-color4"></div>
      <transition name="slide-register-fade">
        <div v-if="animated.leftAnimated" class="l-container">
          <p :style="textLeftColor" class="title">{{ appInfosStore.appCommonConfig.app_name }}</p>
          <p :style="textLeftColor" class="sub">{{ appInfosStore.appCommonConfig.app_description }}</p>
        </div>
      </transition>

    </div>
    <div class="r-content" :style="bgColor">
      <div class="r-content-color-cover"></div>
      <div class="r-content-color1"></div>
      <div class="r-content-color2"></div>
      <div class="r-content-color3"></div>
      <transition name="slide-register-fade">
        <n-card v-if="animated.rightAnimated" class="login-card" :embedded="false" :bordered="false">
          <n-button class="back" text :bordered="false" @click="router.replace({path: '/welcome'})">
            <n-icon style="margin-right: 5px" size="20">
              <backIcon/>
            </n-icon>
            {{ t('userRegister.backHomePage') }}
          </n-button>
          <p class="login-title">{{ t('userRegister.newAccount') }}</p>

          <n-form
              ref="formRef"
              :model="formValue"
              :rules="rules"
          >
            <!-- 邮箱输入框 -->
            <n-form-item
                path="user.email"
                :show-feedback="false"
            >
              <n-input
                  v-model:value="formValue.user.email"
                  :placeholder="t('userRegister.email')"
                  size="large"
                  :bordered="false"
                  :style="placeholderBgColor"
              />
            </n-form-item>

            <!-- 验证码输入框 -->
            <n-form-item
                path="user.verify_code"
                :show-feedback="false"
                v-if="appInfosStore.registerPageConfig.is_email_verify"
            >
              <n-input
                  v-model:value.number="formValue.user.verify_code"
                  :placeholder="t('userRegister.verifyCode')"
                  size="large"
                  class="input-general"
                  :bordered="false"
                  :style="placeholderBgColor"
              />
              <n-button
                  :disabled="!enableSendCode"
                  secondary
                  type="primary"
                  @click="sendVerifyCode"
                  style="margin-left: 12px;"
                  size="large"
              >
                {{ t('userRegister.sendVerifyCode') }}
                {{ waitSendMail !== 0 ? ` (${waitSendMail})` : '' }}
              </n-button>
            </n-form-item>

            <!-- 密码输入框 -->
            <n-form-item
                path="user.password"
                :show-feedback="false"
            >
              <n-input
                  type="password"
                  v-model:value="formValue.user.password"
                  :placeholder="t('userRegister.pwd')"
                  size="large"
                  :bordered="false"
                  :style="placeholderBgColor"
              />
            </n-form-item>

            <!-- 确认密码输入框 -->
            <n-form-item
                path="user.password_confirmation"
                :show-feedback="false"
            >
              <n-input
                  type="password"
                  v-model:value="formValue.user.password_confirmation"
                  :placeholder="t('userRegister.pwdAgain')"
                  size="large"
                  class="input-general"
                  :bordered="false"
                  :style="placeholderBgColor"
              />
            </n-form-item>

            <!-- 邀请码输入框 -->
            <n-form-item
                path="user.invite_user_id"
                :show-feedback="false"
                v-if="appInfosStore.registerPageConfig.is_invite_force"
            >
              <n-input
                  v-model:value.number="formValue.user.invite_user_id"
                  :placeholder="t('userRegister.inviteCode')"
                  size="large"
                  class="input-general"
                  :bordered="false"
                  :style="placeholderBgColor"
              />
            </n-form-item>

            <!-- 用户协议 -->
            <n-form-item>
              <n-checkbox v-model:checked="agreementChecked"></n-checkbox>
              <p style="font-weight: bold; opacity: 0.8; margin-left: 8px">{{ t('userRegister.agreement') }}</p>
              <n-button
                  text
                  type="primary"
                  style="font-weight: bold; opacity: 0.9; margin-left: 3px"
              >
                {{ t('userRegister.terminalUserAgreement') }}
              </n-button>
            </n-form-item>

            <!-- 注册按钮 -->
            <n-form-item>
              <n-button
                  :bordered="false"
                  type="primary"
                  class="login-btn"
                  size="large"
                  @click="handleValidateClick"
                  :disabled="!regBtnEnabled"
              >
                {{ t('userRegister.reg') }}
              </n-button>
            </n-form-item>
          </n-form>
        </n-card>
      </transition>
    </div>
  </div>
</template>

<style scoped lang="less">
.slide-register-fade-enter-active {
  transition: all 500ms ease;
}

.slide-register-fade-leave-active {
  transition: all 500ms ease;
}

.slide-register-fade-enter-from,
.slide-register-fade-leave-to {
  transform: translateY(150px);
  opacity: 0;
}

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
    top: 35%;
    left: 15%;
    //background-color: #4cae4c;
    width: 70%;
    transition: ease 300ms;


    .title {
      font-size: 2.5rem;
      font-weight: 700;
      opacity: 0.8;
      //color: #205387;
      //color: white;
    }

    .sub {
      //color: #136685;
      margin-top: 20px;
      font-size: 1.2rem;
      opacity: 0.5;

      //color: white;
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
    //background-color: rgba(255, 255, 255, 0.0);
    backdrop-filter: blur(90px);
    -webkit-backdrop-filter: blur(100px);
    z-index: 1000;
  }

  .l-content-color1 {
    width: 500px;
    height: 500px;
    border-radius: 50%;
    background-color: rgba(225, 107, 140, 0.2);
    position: absolute;
    //top: calc(70% - 400px / 2);
    //left: calc(30% - 400px / 2);
    //bottom: calc(100% - 100px);
    left: -250px;
    bottom: -250px;
    z-index: 1;
  }

  .l-content-color2 {
    width: 150px;
    height: 150px;
    border-radius: 50%;
    //background-color: rgba(84, 234, 255, 0);
    background: linear-gradient(to right, rgb(129, 129, 221), rgb(142, 177, 255), rgb(131, 255, 244));
    position: absolute;
    top: calc(30% - 150px / 2);
    left: calc(30% - 150px / 2);
    z-index: 2;
  }

  .l-content-color3 {
    width: 200px;
    height: 200px;
    border-radius: 50%;
    //background-color: rgba(220, 135, 227, 0);
    background: linear-gradient(to right, rgb(142, 255, 22), rgb(100, 179, 244));
    position: absolute;
    top: calc(60% - 400px / 2);
    left: calc(50% - 400px / 2);
    z-index: 3;
  }

  .l-content-color4 {
    width: 200px;
    height: 200px;
    border-radius: 50%;
    //background-color: rgba(220, 135, 227, 0);
    background: linear-gradient(to right, rgb(217, 167, 199), rgb(255, 252, 220));
    position: absolute;
    top: calc(50% - 200px / 2);
    left: calc(50% - 200px / 2);
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
    background-color: rgba(218, 144, 144, 0.0);
    //background-color: #66afe9;
    width: 75%;

    .back {
      //background-color: #31739f;
      font-size: 1rem;
      font-weight: 600;
      margin-bottom: 20px;
      opacity: 0.7;
    }


    .login-title {
      font-size: 2rem;
      font-weight: 500;
      margin-bottom: 10px;
      opacity: 0.9;
    }

    .login-btn {
      margin-top: 50px;
      width: 100%;
      height: 45px;
      color: white;
      background: linear-gradient(to right, rgba(167, 112, 239, 0.5), rgba(207, 139, 243, 0.9), rgba(253, 185, 155, 0.5));
      box-shadow: rgba(139, 129, 195, 0.5) 3px 3px 200px 0;
      margin-bottom: 20px;
      transition: ease 300ms;
    }

    .login-btn:hover {
      transform: translateX(0) translateY(-3px);
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