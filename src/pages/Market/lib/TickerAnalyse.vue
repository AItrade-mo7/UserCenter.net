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
  </div>
</template>

<style lang="less" scoped></style>
