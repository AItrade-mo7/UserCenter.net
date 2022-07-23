<script setup lang="ts">
import { defineAsyncComponent } from 'vue';
const PageTitle = defineAsyncComponent(() => import('@/lib/PageTitle.vue'));
</script>

<template>
  <PageTitle> CoinList </PageTitle>
  <div class="PageWrapper ListWrapper">
    CoinList <br />
    CoinList <br />
    CoinList <br />
    CoinList <br />
    CoinList <br />
  </div>
</template>

<style lang="less" scoped>
@import '@/config/constant.less';
</style>
