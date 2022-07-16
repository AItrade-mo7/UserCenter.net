<script setup lang="ts">
import { RouterLink } from 'vue-router';
import { cloneDeep } from '@/utils/tools';
import { CreateServer } from '@/api/CoinServe';
import AuthModal from '@/lib/AuthModal';
import { useRouter } from 'vue-router';
import { GetOkxKeyList } from '@/api/OkxKey';
import { defineAsyncComponent } from 'vue';
const PageTitle = defineAsyncComponent(() => import('@/lib/PageTitle.vue'));
const XIcon = defineAsyncComponent(() => import('@/lib/XIcon.vue'));

const $router = useRouter();

let options = $ref([]);
let SelectHost = $ref('');

const GetList = async () => {
  const res = await GetOkxKeyList();
  if (res.Code > 0) {
    options = res.Data || [];
  }
};
GetList();

const OkxKeyChange = (event) => {
  for (const item of options) {
    if (item.OkxKeyID == event) {
      SelectHost = item.IP;
    }
  }
};

let SubmitStatus: boolean = $ref(false);
const formValue = $ref({
  OkxKeyID: '',
  Port: '',
  Note: '',
  Password: '',
});

const SendForm = async () => {
  const res = await CreateServer({
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
</script>

<template>
  <PageTitle>
    注册一个 CoinAI
    <template #after>
      <RouterLink to="/About/AIServe" className="CreateOkxKey__about">
        <n-button circle quaternary size="tiny" class="About__btn">
          <template #icon> <XIcon name="QuestionCircleOutlined" /> </template>
        </n-button>
      </RouterLink>
    </template>
  </PageTitle>

  <div class="PageWrapper">
    <n-form ref="loginForm" :model="formValue" size="small" class="myForm">
      <n-form-item class="myForm__item" label="选择一个密钥">
        <n-select
          label-field="Name"
          value-field="OkxKeyID"
          @update:value="OkxKeyChange"
          placeholder="请选择一个密钥"
          v-model:value="formValue.OkxKeyID"
          size="small"
          :options="options"
        />
      </n-form-item>

      <n-alert type="warning" class="Alert" v-if="!options.length">
        发现您还没有 okx 秘钥，无法创建服务，点击按钮，新增一个 Key
        <RouterLink to="/OkxKey/Add">
          <n-button size="tiny" type="primary"> 新增一个Key </n-button>
        </RouterLink>
      </n-alert>

      <n-form-item class="myForm__item" label="服务运行的IP地址与端口">
        <n-input-group>
          <n-input
            name="Host"
            v-model:value="SelectHost"
            :inputProps="{ autocomplete: 'password' }"
            placeholder="IP地址"
            disabled
          ></n-input>
          <n-input
            name="Port"
            class="Port__input"
            v-model:value="formValue.Port"
            :inputProps="{ autocomplete: 'password' }"
            placeholder="端口"
          >
            <template #prefix>:</template>
          </n-input>
        </n-input-group>
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

.Port__input {
  width: 116px;
}

.Alert {
  margin-bottom: 12px;
}
</style>
