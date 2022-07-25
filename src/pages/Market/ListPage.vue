<script setup lang="ts">
import { h, onMounted } from 'vue';
import { GetTickerList } from '@/api/CoinMarket';
import type { TickerParam } from '@/api/CoinMarket';
import { defineAsyncComponent } from 'vue';
const PageTitle = defineAsyncComponent(() => import('@/lib/PageTitle.vue'));
const PriceView = defineAsyncComponent(() => import('./lib/PriceView.vue'));

const CoinSort: TickerParam['SortType'] = $ref('Amount');

let CoinTickerList = $ref([]);

const GetCoinTickerList = () => {
  GetTickerList({
    SortType: CoinSort,
  }).then((res) => {
    if (res.Code > 0) {
      CoinTickerList = res.Data;
    }
  });
};

onMounted(() => {
  GetCoinTickerList();
});

const columns = [
  {
    title: 'Coin',
    key: 'CcyName',
  },
  {
    title: 'Amount',
    key: 'Amount',
  },
  {
    title: 'Price',
    key: 'Price',
    render(row) {
      return h(PriceView, {
        data: row,
      });
    },
  },
];
</script>

<template>
  <PageTitle> Market </PageTitle>
  <div class="ListWrapper">
    <div class="TableWrapper">
      <n-data-table :columns="columns" :data="CoinTickerList" />
    </div>
  </div>
</template>

<style lang="less" scoped>
@import '@/config/constant.less';
.TableWrapper {
}
</style>
