<script setup lang="ts">
import { defineAsyncComponent } from 'vue';
const PageTitle = defineAsyncComponent(() => import('@/lib/PageTitle.vue'));
</script>

<template>
  <PageTitle>StockServe</PageTitle>
  <div class="PageWrapper">尚在开发中。。。</div>
</template>

<style lang="less" scoped>
@import '@/config/constant.less';
</style>
