<script setup lang="ts">
import { onMounted } from 'vue';
import { GetAnalyHistory } from '@/api/CoinMarket';

let HistoryList = $ref([]);

const GetHistoryList = () => {
  GetAnalyHistory().then((res) => {
    if (res.Code > 0) {
      HistoryList = res.Data;
    }
  });
};

onMounted(() => {
  GetHistoryList();
});
</script>

<template>
  <PageTitle> AnalyHistory </PageTitle>
  <div class="PageWrapper AnalyHistory">
    <div v-for="item in HistoryList" class="DataBox">
      <n-time :time="item.TimeUnix" />
      {{ item.WholeDir }}
    </div>
  </div>
</template>

<style lang="less" scoped>
.DataBox {
  margin: 20px;
  border: 1px solid red;
}
</style>
