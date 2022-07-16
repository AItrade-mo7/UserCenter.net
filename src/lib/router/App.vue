<script setup lang="ts">
import { RouterView } from 'vue-router';
import { NConfigProvider, zhCN, dateZhCN } from 'naive-ui';
import type { GlobalThemeOverrides } from 'naive-ui';
import { Ping } from '@/api/Ping';
import { setToken, removeToken } from '@/utils/tools';
import { PingDataStore } from '@/store';

import { defineAsyncComponent } from 'vue';
const LoadingView = defineAsyncComponent(() => import('./LoadingView.vue'));
const TopBar = defineAsyncComponent(() => import('@/lib/TopBar.vue'));

const fetchPing = async () => {
  const res = await Ping();
  if (res.Code > 0) {
    const Token = res.Data.Token;
    PingDataStore.update(res.Data);
    if (Token) {
      await setToken(Token);
    }
  } else {
    removeToken();
  }
};
fetchPing();

const themeOverrides: GlobalThemeOverrides = {
  common: {
    primaryColor: '#f0b90b',
  },
};
</script>

<template>
  <NConfigProvider :theme-overrides="themeOverrides" :locale="zhCN" :date-locale="dateZhCN">
    <n-message-provider>
      <n-notification-provider>
        <n-dialog-provider>
          <LoadingView />
          <TopBar />
          <RouterView />
        </n-dialog-provider>
      </n-notification-provider>
    </n-message-provider>
  </NConfigProvider>
</template>
