<script setup lang="ts">
import { cloneDeep, setToken } from '@/utils/tools';
import { useRouter } from 'vue-router';
import { Logo } from '@/config/constant';
import { login } from '@/api/Account';
import { verifyConfig } from '@/utils/verify';
import { UserInfoStore, TopBarStore } from '@/store';

import { defineAsyncComponent } from 'vue';
const XIcon = defineAsyncComponent(() => import('@/lib/XIcon.vue'));
const PageTitle = defineAsyncComponent(() => import('@/lib/PageTitle.vue'));

const $router = useRouter();

let SubmitStatus: boolean = $ref(false);
const formValue = $ref({
  Email: '',
  Password: '',
});

const Submit = async () => {
  const v1 = verifyConfig('Email', formValue.Email);
  const v2 = verifyConfig('Password', formValue.Password);
  if (v1 || v2) {
    window.$message.warning(v1 || v2);
    return;
  }

  SubmitStatus = true;
  const res = await login({
    ...cloneDeep(formValue),
  });

  if (res.Code > 0) {
    const InfoRes = await setToken(res.Data.Token);
    window.$message.success(res.Msg);
    UserInfoStore.update(InfoRes.Data);
    TopBarStore.open();
    $router.replace('/CoinServe');
  }
  SubmitStatus = false;
};
</script>

<template>
  <PageTitle>Login</PageTitle>
  <div class="PageWrapper">
    <div className="Login__logo-box">
      <img className="Login__logo" :src="Logo" alt="" />
    </div>
    <h2 className="Login__title">
      <div>登录</div>
    </h2>

    <n-form ref="loginForm" :model="formValue" size="small" class="myForm">
      <n-form-item class="myForm__item">
        <n-input
          name="Email"
          v-model:value="formValue.Email"
          :inputProps="{ autocomplete: 'password' }"
          placeholder="请输入邮箱地址"
        >
          <template #prefix> <XIcon name="MailOutlined" /> </template>
        </n-input>
      </n-form-item>
      <n-form-item class="myForm__item">
        <n-input
          v-model:value="formValue.Password"
          type="password"
          name="Password"
          show-password-on="click"
          placeholder="请输入密码"
          :inputProps="{ autocomplete: 'password' }"
        ></n-input>
      </n-form-item>

      <div class="change_password">
        <RouterLink to="/change_password"> 忘记密码? </RouterLink>
      </div>

      <n-form-item class="myForm__item">
        <n-button class="Submit" :disabled="SubmitStatus" type="primary" @click="Submit"> 登录 </n-button>
      </n-form-item>
    </n-form>
  </div>
</template>

<style lang="less" scoped>
@import '@/config/constant.less';
.Login__logo-box {
  padding-bottom: 36px;
}

.Login__logo {
  display: block;
  border-radius: 100px;
  overflow: hidden;
  width: 100px;
  height: 100px;
  margin: 0 auto;
}

.Login__title {
  margin: 0;
  text-align: center;
}

.Login__title-str {
  display: block;
}
.Login__forget {
  text-align: center;
}
.change_password {
  text-align: right;
  color: @mainColor;
}
</style>
