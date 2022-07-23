<script setup lang="ts">
import { RouterView } from 'vue-router';
import { defineAsyncComponent } from 'vue';
const PageTitle = defineAsyncComponent(() => import('@/lib/PageTitle.vue'));
</script>

<template>
  <PageTitle></PageTitle>
  <div class="PageWrapper">
    <RouterView />
  </div>
</template>
<style lang="less" scoped>
.PageWrapper {
  padding-top: 36px;
}
</style>
