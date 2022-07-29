<script setup lang="ts">
import { onMounted } from 'vue';
import { GetTickerAnalyse } from '@/api/CoinMarket/index';

let TickerAnalyseData = $ref({});

const GetData = () => {
  GetTickerAnalyse().then((res) => {
    if (res.Code > 0) {
      TickerAnalyseData = res.Data;
    }
  });
};
onMounted(() => {
  GetData();
});
</script>

<template>
  <div class="TickerAnalyse">
    {{ TickerAnalyseData }}
    <n-space class="data-wrapper">
      <div class="block">
        <span class="label">连接状态</span>
        <span class="value green">xxxx</span>
      </div>
      <div class="block">
        <span class="label">系统时间</span>
        <span class="value"><n-time :time="TickerAnalyseData.Ts" /></span>
      </div>
    </n-space>
  </div>
</template>

<style lang="less" scoped></style>
