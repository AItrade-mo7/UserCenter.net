<script setup lang="ts">
import AuthModal from '@/lib/AuthModal';
import { GetDeployShell } from '@/api/CoinServe';
import { defineAsyncComponent } from 'vue';
const ShellAbout = defineAsyncComponent(() => import('./ShellAbout.vue'));

const props = defineProps({
  CoinServeID: String,
});

let Url = $ref('http://file.mo7.cc/AITrade/install/i-9010_1.sh');
const deployFunc = () => {
  const CoinServeID = props?.CoinServeID;

  AuthModal({
    IsPassword: true,
    async OkBack(param) {
      return GetDeployShell({
        CoinServeID,
        Password: param.Password,
      }).then((res) => {
        Url = res.Data.Src;
      });
    },
  });
};
</script>

<template>
  <div class="InfoNot">
    <div v-if="!Url">
      <div className="ServerInfo_hint">服务状态检查失败, 您可能需要 :</div>
      <div class="InfoNot__btn">
        <n-button type="primary" @click="deployFunc"> 获取部署脚本 </n-button>
      </div>
    </div>

    <div v-if="Url && props?.CoinServeID">
      <ShellAbout :Src="Url" :CoinServeID="props?.CoinServeID"></ShellAbout>
    </div>
  </div>
</template>

<style lang="less" scoped>
.InfoNot__btn {
  margin: 0 auto;
  text-align: center;
  margin-top: 24px;
  max-width: 450px;
  .n-button {
    width: 100%;
  }
}
</style>
