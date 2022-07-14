<script setup lang="ts">
import { UserInfoStore } from '@/store';
import { Logo } from '@/config/constant';
import { EditProfile, logout } from '@/api/Account';
import { cloneDeep, UpLoadFileEncrypt, $lcg, getToken, setToken } from '@/utils/tools';
import AuthModal from '@/lib/AuthModal';

import { defineAsyncComponent } from 'vue';
const SendCode = defineAsyncComponent(() => import('@/lib/SendCode.vue'));
const XIcon = defineAsyncComponent(() => import('@/lib/XIcon.vue'));
const PageTitle = defineAsyncComponent(() => import('@/lib/PageTitle.vue'));

let SubmitStatus: boolean = $ref(false);
const formValue = $ref({
  OldEmailCode: '',
  NewEmail: '',
  NewEmailCode: '',
  Avatar: '',
  NickName: '',
  SecurityCode: '',
  Password: '',
});

// 数据回填
let UserInfo: any = {};
const SetUserInfo = () => {
  if (!UserInfo.NickName) {
    UserInfo = cloneDeep(UserInfoStore.value);
    formValue.Avatar = UserInfo.Avatar;
    formValue.NickName = UserInfo.NickName;
    formValue.SecurityCode = UserInfo.SecurityCode;
  }
  return '';
};

// 发送表单
const SendForm = async (isEmail: boolean) => {
  const res = await EditProfile({
    ...cloneDeep(formValue),
  });
  SubmitStatus = false;
  if (res.Code > 0) {
    if (isEmail) {
      window.$notification.success({
        meta: res.Data,
        content: res.Msg,
        onClose() {
          logout();
        },
      });
    } else {
      window.$message.success(res.Msg);
      setToken(getToken());
    }
  }
};

// 提交按钮
const Submit = () => {
  SubmitStatus = true;

  const isEmail = formValue.NewEmail.length > 4 && formValue.NewEmail != UserInfoStore.value.Email;
  const isSecurityCode = formValue.SecurityCode != UserInfoStore.value.SecurityCode;

  if (isEmail || isSecurityCode) {
    AuthModal({
      EmailAction: '修改资料',
      IsPassword: true,
      async OkBack(param) {
        formValue.OldEmailCode = param.Code;
        formValue.Password = param.Password;
        return SendForm(isEmail);
      },
    });
    return;
  }

  SendForm(isEmail);
};

const loseUpLoad = (): any => {
  window.$message.error('头像上传失败');
};
const winUpLoad = (res): any => {
  if (res.file.percentage == 100) {
    const response = $lcg(res, 'event.target.response', {});
    const resJson = JSON.parse(response);

    if (resJson.Code > 0) {
      window.$message.success('上传成功');
      formValue.Avatar = $lcg(resJson, 'Data.Url', '');
    } else {
      window.$message.error('上传失败');
    }
  }
};
</script>

<template>
  {{ SetUserInfo() }}
  <PageTitle>Edit Profile</PageTitle>
  <div class="PageWrapper" v-if="UserInfoStore.value.Email">
    <div className="User__info">
      <n-avatar round :size="70" :src="formValue.Avatar" :fallback-src="Logo" />

      <n-upload
        :show-file-list="false"
        :on-error="loseUpLoad"
        :on-finish="winUpLoad"
        action="//file.mo7.cc/api/public/upload"
        :headers="{
          'File-Path': 'hunter/' + UserInfoStore.value.UserID,
          'Auth-File': UpLoadFileEncrypt(),
        }"
        class="Edit_avatar"
      >
        <n-button size="tiny" type="primary"> 修改头像 </n-button>
      </n-upload>
      <div className="User__Email">{{ UserInfoStore.value.Email }}</div>
    </div>
    <n-form ref="loginForm" :model="formValue" size="small" class="myForm">
      <n-form-item class="myForm__item" label="昵称:">
        <n-input
          v-model:value="formValue.NickName"
          name="NickName"
          placeholder="修改昵称"
          :inputProps="{ autocomplete: 'password' }"
        >
          <template #prefix> <XIcon name="SkinOutlined" /> </template>
        </n-input>
      </n-form-item>

      <n-form-item class="myForm__item" label="防伪标识符:">
        <n-input
          v-model:value="formValue.SecurityCode"
          name="SecurityCode"
          placeholder="修改防伪标识符"
          :inputProps="{ autocomplete: 'password' }"
        >
          <template #prefix> <XIcon name="SafetyOutlined" /> </template>
        </n-input>
      </n-form-item>

      <n-form-item class="myForm__item" label="修改登录邮箱?">
        <n-input
          name="NewEmail"
          v-model:value="formValue.NewEmail"
          :inputProps="{ autocomplete: 'password' }"
          placeholder="请输入新邮箱地址"
        >
          <template #prefix> <XIcon name="MailOutlined" /> </template>
        </n-input>
      </n-form-item>

      <n-form-item class="myForm__item" v-if="formValue.NewEmail">
        <n-input-group>
          <n-input
            name="NewEmailCode"
            v-model:value="formValue.NewEmailCode"
            :inputProps="{ autocomplete: 'password' }"
            placeholder="新邮箱验证码"
            :maxlength="6"
          />
          <SendCode :Email="formValue.NewEmail" Action="修改资料" />
        </n-input-group>
      </n-form-item>

      <n-form-item class="myForm__item">
        <n-button class="Submit" :disabled="SubmitStatus" type="primary" @click="Submit"> 提交 </n-button>
      </n-form-item>
    </n-form>
  </div>
</template>

<style lang="less" scoped>
.PageWrapper {
  text-align: center;
}

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

.Edit_avatar {
  margin-bottom: 12px;
  .n-upload-file-list {
    display: none !important;
  }
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
