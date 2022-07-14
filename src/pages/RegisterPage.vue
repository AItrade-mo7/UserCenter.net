<script setup lang="ts">
import { cloneDeep } from '@/utils/tools';
import { useRouter } from 'vue-router';
import { Logo } from '@/config/constant';
import { Register } from '@/api/Account';
import { verifyConfig } from '@/utils/verify';

import { defineAsyncComponent } from 'vue';
const XIcon = defineAsyncComponent(() => import('@/lib/XIcon.vue'));
const SendCode = defineAsyncComponent(() => import('@/lib/SendCode.vue'));
const PageTitle = defineAsyncComponent(() => import('@/lib/PageTitle.vue'));

const $router = useRouter();

let SubmitStatus: boolean = $ref(false);

const formValue = $ref({
  Email: '',
  Code: '',
});

const Submit = async () => {
  const v1 = verifyConfig('Email', formValue.Email);
  const v2 = verifyConfig('Code', formValue.Code);
  if (v1 || v2) {
    window.$message.warning(v1 || v2);
    return;
  }

  SubmitStatus = true;
  const res = await Register({
    ...cloneDeep(formValue),
  });
  SubmitStatus = false;
  if (res.Code > 0) {
    window.$notification.success({
      meta: res.Data,
      content: res.Msg,
      onClose() {
        $router.replace('/Login');
      },
    });
  }
};
</script>

<template>
  <PageTitle>Register</PageTitle>
  <div class="PageWrapper">
    <div className="logo-box">
      <img className="logo" :src="Logo" alt="" />
    </div>
    <h2 className="title">
      <div>注册</div>
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
        <n-input-group>
          <n-input
            name="Code"
            v-model:value="formValue.Code"
            :inputProps="{ autocomplete: 'password' }"
            placeholder="请输入6位验证码"
            :maxlength="6"
          />
          <SendCode :Email="formValue.Email" Action="注册" />
        </n-input-group>
      </n-form-item>

      <n-form-item class="myForm__item">
        <n-button class="Submit" :disabled="SubmitStatus" type="primary" @click="Submit"> 注册 </n-button>
      </n-form-item>
    </n-form>
  </div>
</template>

<style lang="less" scoped>
.logo-box {
  padding-bottom: 36px;
}

.logo {
  display: block;
  border-radius: 100px;
  overflow: hidden;
  width: 100px;
  height: 100px;
  margin: 0 auto;
}

.title {
  margin: 0;
  text-align: center;
}

.title-str {
  display: block;
}
.forget {
  text-align: center;
}
</style>
