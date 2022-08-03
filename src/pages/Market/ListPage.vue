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

const CoinSort: TickerParam['SortType'] = $ref('Amount');

let CoinTickerList = $ref([]);
let AnalyWhole = $ref({});

const GetCoinTickerList = () => {
  GetTickerList({
    SortType: CoinSort,
  }).then((res) => {
    if (res.Code > 0) {
      CoinTickerList = res.Data.List;
      AnalyWhole = res.Data.AnalyWhole;
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
    key: '12313',
    expandable: (rowData) => {
      return true;
    },
    renderExpand: (rowData) => {
      return `is a good guy.`;
    },
  },
  {
    title: '#',
    key: '1',
    width: 34,
    render: (_, index) => {
      return `${index + 1}`;
    },
  },
  {
    title: 'Coin',
    key: '2',
    width: 68,
    fixed: 'left',
    align: 'left',
  },
  {
    title: 'OKX',
    key: '3',
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
    key: '4',
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
    key: '5',
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
    key: '6',
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
  if (AnalyWhole.MaxUP.InstID == rowData.InstID) {
    return 'MaxUP';
  }

  if (AnalyWhole.MaxDown.InstID == rowData.InstID) {
    return 'MaxDown';
  }
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
    <div v-if="CoinTickerList.length" class="Describe">
      OKX、Binance 综合交易量排名前 {{ CoinTickerList.length }} 的币种。 <br />
      列表数据更新时间 {{ DateFormat(CoinTickerList[0].Ts) }}
    </div>
    <div class="TableWrapper">
      <n-data-table
        :expanded-row-keys="RowOpenKey"
        :on-update:expanded-row-keys="RowOpen"
        :row-class-name="RowClassName"
        :row-key="RowKey"
        size="small"
        striped
        bordered
        :columns="columns"
        :data="CoinTickerList"
      />
    </div>
    <TickerAnalyWhole :Analy="AnalyWhole" />
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
