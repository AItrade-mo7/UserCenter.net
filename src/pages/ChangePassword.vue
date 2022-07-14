<script setup lang="ts">
import { UserInfoStore } from '@/store';
import { Logo } from '@/config/constant';
import { ChangePassword } from '@/api/Account';
// import { verifyConfig } from '@/utils/verify';
import { cloneDeep } from '@/utils/tools';
import { useRouter } from 'vue-router';
import { defineAsyncComponent } from 'vue';
const SendCode = defineAsyncComponent(() => import('@/lib/SendCode.vue'));
const XIcon = defineAsyncComponent(() => import('@/lib/XIcon.vue'));
const PageTitle = defineAsyncComponent(() => import('@/lib/PageTitle.vue'));

const $router = useRouter();

let SubmitStatus: boolean = $ref(false);

const formValue = $ref({
  Email: UserInfoStore.value.Email,
  Code: '',
  Password: '',
  AgainPassword: '',
});

const Submit = async () => {
  SubmitStatus = true;
  const res = await ChangePassword({
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
  <PageTitle>Change Password</PageTitle>
  <div class="PageWrapper">
    <div className="User__info">
      <n-avatar round :size="70" :src="UserInfoStore.value.Avatar" :fallback-src="Logo" />
      <div className="User__NickName">{{ UserInfoStore.value.NickName }}</div>
      <div className="User__Email">{{ UserInfoStore.value.Email }}</div>
    </div>
    <n-form :model="formValue" size="small" class="myForm">
      <n-form-item class="myForm__item" v-if="!UserInfoStore.value.Email">
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
        <n-input
          v-model:value="formValue.Password"
          type="password"
          name="Password"
          show-password-on="click"
          placeholder="请输入密码"
          :inputProps="{ autocomplete: 'password' }"
        ></n-input>
      </n-form-item>

      <n-form-item class="myForm__item">
        <n-input
          v-model:value="formValue.AgainPassword"
          type="password"
          name="AgainPassword"
          show-password-on="click"
          placeholder="再次输入密码"
          :inputProps="{ autocomplete: 'password' }"
        ></n-input>
      </n-form-item>

      <n-form-item class="myForm__item">
        <n-button class="Submit" :disabled="SubmitStatus" type="primary" @click="Submit"> 注册 </n-button>
      </n-form-item>
    </n-form>
  </div>
</template>

<style lang="less" scoped>
.PageWrapper {
  text-align: center;
}

.User__info {
  margin-bottom: 24px;
  .n-avatar {
    margin-bottom: 6px;
    animation: Rotate infinite 36s linear;
  }
  .User__NickName {
    font-size: 16px;
    color: #333;
    margin-bottom: 4px;
  }
  .User__Email {
    font-size: 12px;
    color: #999;
  }
}
</style>
