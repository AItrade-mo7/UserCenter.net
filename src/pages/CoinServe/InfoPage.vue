<script setup lang="ts">
import { mStorage } from '@/utils/tools';
import { useRouter } from 'vue-router';
import { GetAIFundConfig } from '@/api/CoinFundServe';
import { $lcg } from '@/utils/tools';
import { defineAsyncComponent } from 'vue';
const PageTitle = defineAsyncComponent(() => import('@/lib/PageTitle.vue'));
const XIcon = defineAsyncComponent(() => import('@/lib/XIcon.vue'));
const InfoOk = defineAsyncComponent(() => import('./lib/InfoOk.vue'));
const InfoNot = defineAsyncComponent(() => import('./lib/InfoNot.vue'));
const SysManage = defineAsyncComponent(() => import('./lib/SysManage.vue'));

const $router = useRouter();
const FundServeHost = mStorage.get('FundServeHost');
let AIFund_config = $ref({});

const GetConfig = () => {
  GetAIFundConfig({
    ServerInfo: {
      Host: FundServeHost,
    },
  }).then((res) => {
    if (res.Code > 0) {
      AIFund_config = res.Data;
    }
  });
};

if (FundServeHost.length < 6) {
  window.$message.warning('缺少 FundServeHost');
  $router.replace('/');
} else {
  // 开始
  GetConfig();
}

// 控制栏
let drawerStatus = $ref(false);
const OpenSet = () => {
  drawerStatus = true;
};
</script>

<template>
  <PageTitle>
    {{ FundServeHost }}
    <template #after v-if="AIFund_config.AppInfo">
      <n-badge
        class="AIFundServer__dotNet"
        :dot="$lcg(AIFund_config, 'AppInfo.version', '') != $lcg(AIFund_config, 'GithubInfo.version', '')"
      >
        <n-button size="tiny" quaternary @click="OpenSet">
          <template #icon>
            <XIcon spin name="SettingOutlined" />
          </template>
        </n-button>
      </n-badge>
    </template>
  </PageTitle>

  <n-drawer v-model:show="drawerStatus" placement="top">
    <n-drawer-content class="AIFundServer__drawer-content">
      <SysManage v-if="drawerStatus" :config="AIFund_config" />
    </n-drawer-content>
  </n-drawer>

  <div class="PageWrapper">
    <InfoOk v-if="AIFund_config.AppInfo" :config="AIFund_config" />
    <InfoNot v-if="!AIFund_config.AppInfo" />
  </div>
</template>

<style lang="less">
@import '@/config/constant.less';

.n-badge.n-badge--dot.AIFundServer__dotNet {
  position: relative;
  .n-badge-sup {
    position: absolute;
    top: 0;
    right: 0;
    left: auto;
    height: 5px;
    width: 5px;
    min-width: 5px;
  }
}
.n-drawer .AIFundServer__drawer-content .n-drawer-body-content-wrapper {
  padding: 16px;
}
</style>
