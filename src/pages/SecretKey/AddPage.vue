<script setup lang="ts">
import { RouterLink } from 'vue-router';
import { cloneDeep, ParseOkxKey } from '@/utils/tools';
import { CreateOkxKey } from '@/api/okxKey';
import AuthModal from '@/lib/AuthModal';
import { useRouter } from 'vue-router';
import { defineAsyncComponent } from 'vue';
const PageTitle = defineAsyncComponent(() => import('@/lib/PageTitle.vue'));
const XIcon = defineAsyncComponent(() => import('@/lib/XIcon.vue'));

const $router = useRouter();

let SubmitStatus: boolean = $ref(false);
let formValue = $ref({
  ApiKey: '',
  SecretKey: '',
  Passphrase: '',
  IP: '',
  Name: '',
  Note: '',
  Password: '',
});

const SendForm = async () => {
  const res = await CreateOkxKey({
    ...cloneDeep(formValue),
  });
  if (res.Code > 0) {
    window.$message.success(res.Msg);
    $router.go(-1);
  }
};

const Submit = async () => {
  SubmitStatus = true;
  AuthModal({
    IsPassword: true,
    async OkBack(param) {
      SubmitStatus = false;
      formValue.Password = param.Password;
      return SendForm();
    },
  });
  SubmitStatus = false;
};

const readPlate = async () => {
  const text = await navigator.clipboard.readText();
  const obj = ParseOkxKey(text);
  if (obj.ApiKey.length > 10) {
    window.$message.success('解析成功!');
    formValue = {
      ...formValue,
      ...obj,
    };
  } else {
    window.$message.warning('解析失败!');
  }
};
</script>

<template>
  <PageTitle>
    新增 okx 密钥组
    <template #after>
      <RouterLink to="/about/okxkey" className="CreateOkxKey__about">
        <n-button circle quaternary size="tiny" class="About__btn">
          <template #icon> <XIcon name="QuestionCircleOutlined" /> </template>
        </n-button>
      </RouterLink>
    </template>
  </PageTitle>

  <div class="PageWrapper">
    <n-form ref="loginForm" :model="formValue" size="small" class="myForm">
      <n-form-item class="myForm__item" label="Name">
        <n-input
          name="Name"
          v-model:value="formValue.Name"
          :inputProps="{ autocomplete: 'password' }"
          placeholder="给密钥组起一个名字"
        >
          <template #prefix> <XIcon name="SkinOutlined" /> </template>
        </n-input>
      </n-form-item>

      <n-form-item class="myForm__item" label="ApiKey">
        <n-input
          name="ApiKey"
          v-model:value="formValue.ApiKey"
          :inputProps="{ autocomplete: 'password' }"
          placeholder="在 okx 申请的 [Api Key]"
        >
          <template #prefix> <XIcon name="OkxLogo" /> </template>
        </n-input>
      </n-form-item>

      <n-form-item class="myForm__item" label="SecretKey">
        <n-input
          name="SecretKey"
          v-model:value="formValue.SecretKey"
          :inputProps="{ autocomplete: 'password' }"
          placeholder="在 okx 申请的 [密钥]"
        >
          <template #prefix> <XIcon name="OkxLogo" /> </template>
        </n-input>
      </n-form-item>

      <n-form-item class="myForm__item" label="Passphrase">
        <n-input
          name="Passphrase"
          v-model:value="formValue.Passphrase"
          :inputProps="{ autocomplete: 'password' }"
          placeholder="申请密钥时设置的密码"
        >
          <template #prefix> <XIcon name="OkxLogo" /> </template>
        </n-input>
      </n-form-item>

      <n-form-item class="myForm__item" label="IP">
        <n-input
          name="IP"
          v-model:value="formValue.IP"
          :inputProps="{ autocomplete: 'password' }"
          placeholder="申请密钥时绑定的IP地址"
        >
          <template #prefix> <XIcon name="OkxLogo" /> </template>
        </n-input>
      </n-form-item>

      <n-form-item class="myForm__item" label="Note">
        <n-input
          name="Note"
          type="textarea"
          class="my-textarea"
          v-model:value="formValue.Note"
          :inputProps="{ autocomplete: 'password' }"
          placeholder="给个备注"
          :autosize="{
            minRows: 2,
          }"
        >
        </n-input>
      </n-form-item>

      <n-form-item class="myForm__item">
        <n-button class="Submit" :disabled="SubmitStatus" type="primary" @click="Submit"> 提交 </n-button>
      </n-form-item>
      <div class="page__footer">
        <n-button size="tiny" tertiary type="success" @click="readPlate"> 读取剪切板 </n-button>
      </div>
    </n-form>
  </div>
</template>

<style lang="less" scoped>
@import '@/config/constant.less';
.Create__title {
  font-size: 16px;
  margin: 0;
  .About__btn {
    color: @mainColor;
  }
}
.CreateOkxKey__about {
  display: flex;
  align-items: center;
  justify-content: flex-start;
  .n-button {
    color: @mainColor;
  }
}

.page__footer {
  display: flex;
  align-items: center;
  justify-content: flex-end;
}
</style>
