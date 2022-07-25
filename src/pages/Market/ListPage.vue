<script setup lang="ts">
import { GetTickerList } from '@/api/CoinMarket';
import type { TickerParam } from '@/api/CoinMarket';
import type { DataTableColumns } from 'naive-ui';
import { defineAsyncComponent } from 'vue';
const PageTitle = defineAsyncComponent(() => import('@/lib/PageTitle.vue'));

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
GetCoinTickerList();

const columns = [
  {
    title: 'Coin',
    key: 'CcyName',
  },
  {
    title: 'Price',
    key: 'Last',
  },
  {
    title: 'Amount',
    key: 'Amount',
  },
  {
    title: 'U_R24',
    key: 'U_R24',
  },
  {
    title: 'Amount',
    key: 'Amount',
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
