<script setup lang="ts">
import { h, onMounted, onUnmounted } from 'vue';
import { GetTickerList } from '@/api/CoinMarket';
import type { TickerParam } from '@/api/CoinMarket';
import { DateFormat } from '@/utils/filters';
import { defineAsyncComponent } from 'vue';
const PageTitle = defineAsyncComponent(() => import('@/lib/PageTitle.vue'));
const PriceView = defineAsyncComponent(() => import('./lib/PriceView.vue'));
const VolumeView = defineAsyncComponent(() => import('./lib/VolumeView.vue'));
const TickerAnalyWhole = defineAsyncComponent(() => import('./lib/TickerAnalyWhole.vue'));
const TickerAnalySingle = defineAsyncComponent(() => import('./lib/TickerAnalySingle.vue'));

const CoinSort: TickerParam['SortType'] = $ref('Amount');

let CoinTickerList = $ref([]);
let AnalyWhole = $ref([]);
let AnalySingle = $ref({});

const GetCoinTickerList = () => {
  GetTickerList({
    SortType: CoinSort,
  }).then((res) => {
    if (res.Code > 0) {
      CoinTickerList = res.Data.List;
      AnalyWhole = res.Data.AnalyWhole;
      AnalySingle = res.Data.AnalySingle;
    }
  });
};

let timer: any = null;
onMounted(() => {
  GetCoinTickerList();

  clearInterval(timer);
  timer = setInterval(() => {
    GetCoinTickerList();
  }, 180000);
});

onUnmounted(() => {
  clearInterval(timer);
});

const columns: any[] = [
  {
    type: 'expand',
    expandable: () => {
      return true;
    },
    renderExpand: (row) => {
      const Single = AnalySingle[row.InstID];
      return h(TickerAnalySingle, {
        Single,
      });
    },
  },
  {
    title: '#',
    width: 34,
    render: (_, index) => {
      return `${index + 1}`;
    },
  },
  {
    title: 'Coin',
    key: 'CcyName',
    width: 68,
    fixed: 'left',
    align: 'left',
  },
  {
    title: 'OKX',
    width: 86,
    className: 'OKX',
    align: 'right',
    render(row) {
      return h(VolumeView, {
        Data: row,
        Volume: row.OKXVol24H,
        Bourse: 'OKX',
      });
    },
  },
  {
    title: 'Binance',
    width: 86,
    className: 'Binance',
    align: 'right',
    render(row) {
      return h(VolumeView, {
        Data: row,
        Volume: row.BinanceVol24H,
        Bourse: 'Binance',
      });
    },
  },
  {
    title: 'Volume',
    width: 104,
    className: 'Volume',
    align: 'right',
    render(row) {
      return h(VolumeView, {
        Data: row,
        Volume: row.Volume,
        Bourse: 'Volume',
      });
    },
  },
  {
    title: '24h',
    width: 100,
    align: 'right',
    render(row) {
      return h(PriceView, {
        data: row,
      });
    },
  },
];

const RowClassName = (rowData) => {
  // if (AnalyWhole.MaxUP.InstID == rowData.InstID) {
  //   return 'MaxUP';
  // }
  // if (AnalyWhole.MaxDown.InstID == rowData.InstID) {
  //   return 'MaxDown';
  // }
};

let RowOpenKey = $ref([]);

const RowOpen = (keys) => {
  RowOpenKey = keys;
};

const RowKey = (rowData) => {
  return rowData.CcyName;
};
</script>

<template>
  <PageTitle> Market </PageTitle>
  <div class="ListWrapper">
    <div v-if="CoinTickerList.length" class="Describe">OKX、Binance 综合交易量排名前 N 的币种。 <br /></div>
    <div class="TableWrapper">
      <n-data-table
        :expanded-row-keys="RowOpenKey"
        :on-update:expanded-row-keys="RowOpen"
        :xx-row-class-name="RowClassName"
        :row-key="RowKey"
        size="small"
        striped
        bordered
        :columns="columns"
        :data="CoinTickerList"
      />
    </div>
    <div v-for="item in AnalyWhole">
      <TickerAnalyWhole :Analy="item" />
    </div>
  </div>
</template>

<style lang="less">
@import '@/config/constant.less';

.Describe {
  font-size: 16px;
  margin-top: 16px;
  margin-bottom: 12px;
}

.TableWrapper {
  .OKX {
    color: #999;
  }
  .Binance {
    color: #f0b90b;
  }
  .Volume {
    color: #000;
  }
}

.n-data-table-tbody .n-data-table-tr {
  &.MaxUP {
    td {
      border-top: 2px solid @greenColor;
      border-bottom: 2px solid @greenColor;
      &:first-child {
        border-left: 2px solid @greenColor;
      }
      &:last-child {
        border-right: 2px solid @greenColor;
      }
    }
  }
  &.MaxDown {
    td {
      border-top: 2px solid @redColor;
      border-bottom: 2px solid @redColor;
      &:first-child {
        border-left: 2px solid @redColor;
      }
      &:last-child {
        border-right: 2px solid @redColor;
      }
    }
  }
}
</style>
