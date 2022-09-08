<script setup lang="ts">
import { h, onMounted, onUnmounted } from 'vue';
import { GetAnalyHistory } from '@/api/CoinMarket';

const GetHistoryList = () => {
  GetAnalyHistory().then((res) => {
    if (res.Code > 0) {
      console.log(res);
    }
  });
};

onMounted(() => {
  GetHistoryList();
});
</script>

<template>
  <PageTitle> AnalyHistory </PageTitle>
  <div class="PageWrapper AnalyHistory"></div>
</template>

<style lang="less" scoped></style>
