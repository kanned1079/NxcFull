<script setup lang="ts">
import {useI18n} from "vue-i18n";
import VueHcaptcha from '@hcaptcha/vue3-hcaptcha';
import {computed, h, onBeforeMount, onMounted, onUnmounted, ref} from "vue";
import {useRouter} from "vue-router";
import useThemeStore from "@/stores/useThemeStore";
import useAppInfosStore from "@/stores/useAppInfosStore";
import type {FormInst, FormItemRule, FormRules} from 'naive-ui'
import {useMessage, useNotification} from "naive-ui";
import {ChevronBackOutline as backIcon, LogInOutline as loginIcon} from "@vicons/ionicons5";
import {handleFinalRegister, handleSendVerifyCode, handleVerifyCode} from "@/api/user/register";
import {handleFetchRegisterConfig} from "@/api/common/env"

const {t} = useI18n();
const appInfosStore = useAppInfosStore();
const themeStore = useThemeStore();
const message = useMessage();
const notification = useNotification();
const router = useRouter();

let hCaptchaVerified = ref<boolean>(false);
let hCaptchaExpired = ref<boolean>(false);
let hCaptchaToken = ref<string>('')
let hCaptchaEKey = ref<string>('')
let hCaptchaErr = ref<string>('')

const onVerify = (tokenStr?: string, eKey?: string) => {
  hCaptchaVerified.value = true // 通过认证
  hCaptchaToken.value = tokenStr || ''
  hCaptchaEKey.value = eKey || ''
  message.success(t('userRegister.hCaptcha.passed'))
  console.log(hCaptchaToken.value, hCaptchaEKey.value)
}

const onExpire = () => {
  hCaptchaVerified.value = true // 没有通过认证
  hCaptchaExpired.value = true  // 标记验证过期
  message.warning(t('userRegister.hCaptcha.expired'))
}

const onChallengeExpire = () => {
  hCaptchaVerified.value = true // 没有通过认证
  hCaptchaExpired.value = true  // 标记验证过期
  message.warning(t('userRegister.hCaptcha.challengeExpired'))
}

const onError = (err: any) => {
  hCaptchaVerified.value = true // 没有通过认证
  hCaptchaExpired.value = true  // 标记验证过期
  hCaptchaErr.value = `Error: ${err}`
  message.error(t('userRegister.hCaptcha.err') + hCaptchaErr.value)
}

const calculateContentStyle = () => {
  return window.innerWidth > 768
      ? "display: flex; justify-content: center; flex-direction: column; align-items: center; height: 100%"
      : "display: flex; justify-content: center; flex-direction: column; align-items: center;";
}

const formContentStyle = ref<string>('')
const reCalculateContentStyle = () => {
  formContentStyle.value = calculateContentStyle()
}

const registerFormRef = ref<FormInst | null>(null);
let verifyCodePassed = ref<boolean>(false);
let animated = ref<{ leftAnimated: boolean; rightAnimated: boolean }>({
  leftAnimated: false,
  rightAnimated: false,
});
let bgColor = computed(() =>
    themeStore.enableDarkMode
        ? {backgroundColor: "rgba(40, 41, 41, 1)"}
        : {backgroundColor: "#fff"}
);
let coverBgColor = computed(() =>
    themeStore.enableDarkMode
        ? {backgroundColor: "rgba(40, 40, 40, 0.2)"}
        : {backgroundColor: "rgba(255, 255, 255, 0.0)"}
);
let placeholderBgColor = computed(() =>
    themeStore.enableDarkMode
        ? {}
        // : {backgroundColor: "#f2fafd",}
        : {}
);
let textLeftColor = computed(() => (themeStore.enableDarkMode ? {color: "#fff"} : {}));

let agreementChecked = ref<boolean>(false); // 同意按钮是否被选中
let enableRegister = ref<boolean>(false); // 由检查逻辑控制
// let regBtnEnabled = computed(() => agreementChecked.value && enableRegister.value && !appInfosStore.registerPageConfig.stop_register);
let regBtnEnabled = computed(() =>
    agreementChecked.value &&
    enableRegister.value &&
    !appInfosStore.registerPageConfig.stop_register &&
    (!appInfosStore.registerPageConfig.recaptcha_enable || hCaptchaVerified.value)
);
// console.log(appInfosStore.registerPageConfig.recaptcha_enable, hCaptchaVerified.value)
let clickedCount = ref<number>(0); // 注册按钮点击次数
let enableSendCode = ref<boolean>(true);
let waitSendMail = ref<number>(0);

let autoCompleteEmailSuffixOptions = computed(() => {
  return ['@outlook.com', '@gmail.com', '@hotmail.com', '@yahoo.com', '@msn.com', '@163.com'].map((suffix) => {
    const prefix = formValue.value.user.email.split('@')[0]
    return {
      label: prefix + suffix,
      value: prefix + suffix
    }
  })
})
let getShow = (value: string) => value.endsWith('@');

let formValue = ref({
  user: {
    email: "",
    verify_code: "",
    password: "",
    password_confirmation: "",
    invite_user_id: "",
  },
});

let failReasons = ref<string[]>([]);

// 邮箱格式验证
let validateEmail = (email: string): boolean => {
  const emailRegex = /^[0-9a-zA-Z_.-]+@[0-9a-zA-Z_.-]+(\.[a-zA-Z]+){1,2}$/;

  // 检查邮箱格式是否符合正则
  if (!emailRegex.test(email)) {
    failReasons.value.push(t('userRegister.form.emailFormatErr'));
    return false;
  }

  const [localPart, domain] = email.split("@");
  const isGmail = domain.toLowerCase() === "gmail.com";
  const isGooglemail = domain.toLowerCase() === "googlemail.com";

  // 如果是 Gmail 或 Googlemail，且开启了 Gmail 限制
  if ((isGmail || isGooglemail) && appInfosStore.registerPageConfig.email_gmail_limit_enable) {

    // 不允许使用 Gmail 的多别名功能（含 "+"）
    if (localPart.includes("+")) {
      failReasons.value.push(t('userRegister.form.gmailLimitErr'));
      return false;
    }
    // 如果本地部分包含 .，但去掉 . 后本地部分与原本不一致
    const normalizedLocalPart = localPart.replace(/\./g, "");
    if (normalizedLocalPart !== localPart) {
      failReasons.value.push(t('userRegister.form.gmailDotNotAllowed'));
      return false;
    }
    // 本地部分必须是小写
    if (localPart !== localPart.toLowerCase()) {
      failReasons.value.push(t('userRegister.form.gmailPartLowerForced'));
      return false;
    }
    // 如果是 googlemail.com，则直接不允许
    if (isGooglemail) {
      failReasons.value.push(t('userRegister.form.googlemailNotAllowed'));
      return false;
    }
  }

  // 如果所有检查都通过，返回 true
  return true;
};

// 表单验证规则
let rules: FormRules = {
  user: {
    email: {
      required: true,
      trigger: 'blur',
      validator: () => validateEmail(formValue.value.user.email) ? true : new Error('Form failed due to invalid email address')
    },
    verify_code: {
      required: appInfosStore.registerPageConfig.email_verify,
      message: "Verification code cannot be empty",
      trigger: ['blur', 'input'],
      validator: (rule: FormItemRule, value: string) => {
        if (!appInfosStore.registerPageConfig.email_verify) {
          return true; // Skip verification if not needed
        }
        if (!value) {
          failReasons.value.push(t('userRegister.form.verifyCodeRequire')); // 繁體中文
          return new Error('Verification code is required');
        }
        const codeRegex = /^\d{6}$/;
        if (!codeRegex.test(value)) {
          failReasons.value.push(t('userRegister.form.verifyCodeFormatErr')); // 繁體中文
          return new Error('Invalid verification code format');
        }
        return true;
      },
    },
    password: {
      required: true,
      trigger: 'blur',
      validator: (rule: FormItemRule, value: string) => {
        // Check if password is empty
        if (!value) {
          failReasons.value.push(t('userRegister.form.passwordRequire')); // 繁體中文
          return new Error('Password cannot be empty');
        }

        // Check password length
        if (value.length < 10) {
          failReasons.value.push(t('userRegister.form.passwordLengthRequire')); // 繁體中文
          return new Error("Password length is less than 12 characters");
        }

        // Check if password contains at least 3 character types
        const hasUpperCase = /[A-Z]/.test(value);  // Uppercase letter
        const hasLowerCase = /[a-z]/.test(value);  // Lowercase letter
        const hasNumber = /\d/.test(value);        // Number
        const hasSpecialChar = /[!@#$%^&*(),.?":{}|<>]/.test(value);  // Special character

        const typeCount = [hasUpperCase, hasLowerCase, hasNumber, hasSpecialChar].filter(Boolean).length;

        // Check if password meets at least 3 types
        if (typeCount < 3) {
          failReasons.value.push(t('userRegister.form.passwordComplexRequire')); // 繁體中文
          return new Error("Password must contain at least three types of characters: uppercase, lowercase, digits, or special characters");
        }

        return true;
      }
    },
    password_confirmation: [
      {
        required: true,
        message: "Please confirm your password",
        trigger: 'blur',
        validator(rule: FormItemRule, value: string) {
          if (value === '') {
            failReasons.value.push(t('userRegister.form.passwordAgainRequire'))
            return new Error("Password confirmation require");
          }
          let isValid = formValue.value.user.password === value && value.trim() !== ''
          if (!isValid) {
            failReasons.value.push(t('userRegister.form.passwordAgainNotMatch'))
            return new Error("Password confirmation does not match");
          } else return true
        },
      }
    ],
    invite_user_id: {
      required: appInfosStore.registerPageConfig.invite_require,
      trigger: 'blur',
      validator(rule: FormItemRule, value: string) {
        if (!appInfosStore.registerPageConfig.invite_require || value.trim() !== '') {
          return true;
        }
        failReasons.value.push(t('userRegister.form.inviteCodeRequire'));
        return new Error('You must be invited');
      }
    }
  },
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
let callVerifyCode = async () => {
  let data = await handleVerifyCode(formValue.value.user.email.trim(), formValue.value.user.verify_code.trim());
  if (data && data.code === 200) {
    verifyCodePassed.value = data.passed as boolean;
  } else {
    message.error(t('userRegister.verifyCodeExpireErr'));
    verifyCodePassed.value = false;
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
      await router.push({path: "/login"});
    }, 1000);
  } else if (data.code === 409) {
    message.error(t('userRegister.thisMailAlreadyExist'));
  } else {
    message.error(t("userRegister.regFailure") + data.msg);
  }
};

// 按钮点击次数限制

let showFailReasons = () => {
  let _failReasons = failReasons.value
  notification.create({
    title: computed(() => t("userRegister.form.checkForm")).value, // 通知标题
    duration: 10000,
    description: () => {
      // 使用 div 来渲染多个错误信息
      return h('div', [
        _failReasons.map((reason: string, index: number) => {
          return h('p', {key: index}, `${index + 1} ${reason}`); // 渲染每个错误信息
        }),
      ]);
    },
    // meta: new Date().toLocaleString(), // 可以加入时间戳或其他额外信息
    onClose: () => {
      // 关闭通知时的回调
      console.log('Fail reasons notification closed');
    },
  });
};

let registerClick = async () => {
  failReasons.value = []
  await registerFormRef.value?.validate((errors) => {
    if (errors) return showFailReasons()
  })

  appInfosStore.registerPageConfig.email_verify && await callVerifyCode();
  if (appInfosStore.registerPageConfig.email_verify && !verifyCodePassed.value) {
    message.error(t("userRegister.verifyCodeErr"));
    return;
  }
  await handleRegister();

}

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
}

let turn2UserAgreement = () => window.open(appInfosStore.registerPageConfig.tos_url, "_blank")

onBeforeMount(async () => {
  let isFetchSuccess = await handleFetchRegisterConfig('en')
  if (isFetchSuccess) {
    enableRegister.value = true
  } else {
    message.error(t('userRegister.pageConfigFetchFailure'))
    enableRegister.value = false
  }
})

onMounted(async () => {
  const queryParams = new URLSearchParams(window.location.search);
  const code = queryParams.get("code");
  code ? (formValue.value.user.invite_user_id = code) : null;

  formContentStyle.value = calculateContentStyle()
  window.addEventListener('resize', reCalculateContentStyle)

  setTimeout(() => (animated.value.leftAnimated = true), 100);
  setTimeout(() => (animated.value.rightAnimated = true), 150);
});

onUnmounted(() => window.removeEventListener('resize', reCalculateContentStyle))
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
        <n-scrollbar
            style="height: 100vh"
            :content-style="formContentStyle"
        >
          <n-card v-if="animated.rightAnimated" class="login-card" :embedded="false" :bordered="false">
            <n-button class="back" text :bordered="false" @click="router.replace({path: '/welcome'})"
                      icon-placement="left">
              <template #icon>
                <n-icon style="margin-right: 5px" size="20">
                  <backIcon/>
                </n-icon>
              </template>
              {{ t('userRegister.backHomePage') }}
            </n-button>
            <p class="login-title">{{ t('userRegister.newAccount') }}</p>

            <n-alert
                type="warning"
                :title="t('userRegister.stopRegisterTitle')"
                v-if="appInfosStore.registerPageConfig.stop_register"
                style="margin-bottom: 50px;"
            >
              <n-p style="font-size: 1rem">
                {{ t('userRegister.stopRegisterHint') }}
              </n-p>
              <n-p style="opacity: 0.6; font-weight: bold">by {{ appInfosStore.appCommonConfig.app_name }}</n-p>
            </n-alert>

            <n-form
                v-if="!appInfosStore.registerPageConfig.stop_register"
                ref="registerFormRef"
                :model="formValue"
                :rules="rules"
                :show-label="false"
                :show-feedback="false"
                size="medium"
                autofocus
                :disabled="appInfosStore.registerPageConfig.stop_register"
            >
              <!-- 邮箱输入框 -->
              <n-form-item path="user.email" class="btn-bottom-gap">
                <!--              <n-input-->
                <!--                  v-model:value="formValue.user.email"-->
                <!--                  :placeholder="t('userRegister.email')"-->
                <!--                  :bordered="true"-->
                <!--                  :style="placeholderBgColor"-->
                <!--              />-->
                <n-auto-complete
                    clearable
                    :append="false"
                    :get-show="getShow"
                    v-model:value="formValue.user.email"
                    :placeholder="t('userRegister.email')"
                    :options="autoCompleteEmailSuffixOptions"
                    :input-props="{autocomplete: 'disabled'}"
                    :style="placeholderBgColor"
                ></n-auto-complete>
              </n-form-item>

              <!-- 验证码输入框 -->
              <n-form-item
                  path="user.verify_code"
                  :show-feedback="false"
                  v-if="appInfosStore.registerPageConfig.email_verify"
                  class="btn-bottom-gap"
              >
                <n-input
                    v-model:value.number="formValue.user.verify_code"
                    :placeholder="t('userRegister.verifyCode')"
                    :bordered="true"
                    :style="placeholderBgColor"
                />
                <n-button
                    :disabled="!enableSendCode"
                    secondary
                    type="primary"
                    @click="sendVerifyCode"
                    style="margin-left: 12px;"
                >
                  {{ t('userRegister.sendVerifyCode') }}
                  {{ waitSendMail !== 0 ? ` (${waitSendMail})` : '' }}
                </n-button>
              </n-form-item>

              <!-- 密码输入框 -->
              <n-popover :show-arrow="false" trigger="hover" placement="top">
                <template #trigger>
                  <n-tag
                      :bordered="false"
                      :checked="false"
                      style="font-size: 0.9rem !important; opacity: 0.8; background-color: rgba(0,0,0,0.0)"
                  >
                    {{ t('userRegister.passwordComplexRequirePart1') }}
                    <n-button
                        type="primary"
                        text
                        style="text-decoration: underline"
                    >
                      {{ t('userRegister.passwordComplexRequirePart2') }}
                    </n-button>
                  </n-tag>
                </template>
                {{ t('userRegister.passwordComplexHint1') }}<br>
                {{ t('userRegister.passwordComplexHint2') }}
              </n-popover>
              <n-form-item
                  path="user.password"
                  class="btn-bottom-gap"
              >
                <n-input
                    type="password"
                    showPasswordOn="click"
                    v-model:value="formValue.user.password"
                    :placeholder="t('userRegister.pwd')"
                    :bordered="true"
                    :style="placeholderBgColor"
                />
              </n-form-item>

              <!-- 确认密码输入框 -->
              <n-form-item
                  path="user.password_confirmation"
                  class="btn-bottom-gap"
              >
                <n-input
                    type="password"
                    showPasswordOn="click"
                    v-model:value="formValue.user.password_confirmation"
                    :placeholder="t('userRegister.pwdAgain')"
                    :bordered="true"
                    :style="placeholderBgColor"
                />
              </n-form-item>

              <!-- 邀请码输入框 -->
              <n-form-item
                  path="user.invite_user_id"
                  class="btn-bottom-gap"
              >
                <n-input
                    clearable
                    v-model:value.number="formValue.user.invite_user_id"
                    :placeholder="t('userRegister.inviteCode')"
                    :bordered="true"
                    :style="placeholderBgColor"
                />
              </n-form-item>

              <!-- 用户协议 -->
              <n-form-item
                  class="btn-bottom-gap"
              >
                <n-checkbox v-model:checked="agreementChecked"></n-checkbox>
                <p style="font-weight: bold; opacity: 0.8; margin-left: 8px">{{ t('userRegister.agreement') }}</p>
                <n-button
                    text
                    type="primary"
                    style="font-weight: bold; opacity: 0.9; margin-left: 3px"
                    @click="turn2UserAgreement"
                >
                  {{ t('userRegister.terminalUserAgreement') }}
                </n-button>
              </n-form-item>

              <div v-if="appInfosStore.registerPageConfig.recaptcha_enable">
                <vue-hcaptcha
                    :sitekey="appInfosStore.registerPageConfig.recaptcha_site_key || ''"
                    :theme="themeStore.enableDarkMode?'dark':null"
                    @verify="onVerify"
                    @expired="onExpire"
                    @challenge-expired="onChallengeExpire"
                    @error="onError"
                >
                </vue-hcaptcha>
              </div>


              <!-- 注册按钮 -->
              <n-form-item>
                <n-button
                    :bordered="false"
                    type="primary"
                    class="login-btn"
                    size="large"
                    @click="registerClick"
                    :disabled="!regBtnEnabled"
                >
                  {{ t('userRegister.reg') }}
                  <template #icon>
                    <n-icon>
                      <loginIcon/>
                    </n-icon>
                  </template>
                </n-button>
              </n-form-item>
            </n-form>

<!--            <div class="bottom-hint-root">-->
<!--              <div v-if="appInfosStore.registerPageConfig.stop_register" style="height: 50px"></div>-->
<!--              <n-divider style="width: 98%; margin: 0 !important;">-->
<!--                <p style="opacity: 0.2">·</p>-->
<!--              </n-divider>-->
<!--              &lt;!&ndash;            <div style="margin: 50px 0 5px 0">&ndash;&gt;-->
<!--              &lt;!&ndash;              <img style="width: 45px" :src="appInfosStore.appCommonConfig.logo_url" alt="icon">&ndash;&gt;-->
<!--              &lt;!&ndash;            </div>&ndash;&gt;-->
<!--              <div style="font-size: 1rem; opacity: 0.8; margin-top: 20px">{{ appInfosStore.appCommonConfig.app_name }}-->
<!--                {{ t('userRegister.allRightsReserved') }}-->
<!--              </div>-->
<!--              <div style="opacity: 0.4; margin-top: 10px"> {{ t('userRegister.securityAndLaws') }}</div>-->
<!--            </div>-->
          </n-card>
        </n-scrollbar>
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
    width: 65%;

    .btn-bottom-gap {
      margin-bottom: 30px;
    }

    .pwd-complex-hint {
      margin: 0;
    }

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
      margin-bottom: 40px;
      opacity: 0.9;
    }

    .login-btn {
      margin-top: 30px;
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

    .bottom-hint-root {
      width: 100%;
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      text-align: center;
      margin: 30px 0;
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

@media (min-width: 768px) {
  .sc {
    height: 100vh;
  }
}

</style>