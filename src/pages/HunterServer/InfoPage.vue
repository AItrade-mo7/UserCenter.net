<script setup lang="ts">
import { mStorage } from '@/utils/tools';
import { useRouter } from 'vue-router';
import { GetHunterConfig } from '@/api/hunter_net';
import { $lcg } from '@/utils/tools';
import { defineAsyncComponent } from 'vue';
const PageTitle = defineAsyncComponent(() => import('@/lib/PageTitle.vue'));
const XIcon = defineAsyncComponent(() => import('@/lib/XIcon.vue'));
const InfoOk = defineAsyncComponent(() => import('./lib/InfoOk.vue'));
const InfoNot = defineAsyncComponent(() => import('./lib/InfoNot.vue'));
const SysManage = defineAsyncComponent(() => import('./lib/SysManage.vue'));

const $router = useRouter();
const hunter_host = mStorage.get('hunter_host');
let hunter_config = $ref({});

const GetConfig = () => {
  GetHunterConfig({
    ServerInfo: {
      Host: hunter_host,
    },
  }).then((res) => {
    if (res.Code > 0) {
      hunter_config = res.Data;
    }
  });
};

if (hunter_host.length < 6) {
  window.$message.warning('缺少 hunter_host');
  $router.replace('/hunter_serve');
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
    {{ hunter_host }}
    <template #after v-if="hunter_config.AppInfo">
      <n-badge
        class="HunterServer__dotNet"
        :dot="$lcg(hunter_config, 'AppInfo.version', '') != $lcg(hunter_config, 'GithubInfo.version', '')"
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
    <n-drawer-content class="HunterServer__drawer-content">
      <SysManage v-if="drawerStatus" :config="hunter_config" />
    </n-drawer-content>
  </n-drawer>

  <div class="PageWrapper">
    <InfoOk v-if="hunter_config.AppInfo" :config="hunter_config" />
    <InfoNot v-if="!hunter_config.AppInfo" />
  </div>
</template>

<style lang="less">
@import '@/config/constant.less';

.n-badge.n-badge--dot.HunterServer__dotNet {
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
.n-drawer .HunterServer__drawer-content .n-drawer-body-content-wrapper {
  padding: 16px;
}
</style>
