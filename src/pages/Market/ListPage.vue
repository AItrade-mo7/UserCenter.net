<script setup lang="ts">
import { h, onMounted, onUnmounted } from 'vue';
import { GetTickerList } from '@/api/CoinMarket';
import type { TickerParam } from '@/api/CoinMarket';
import { DateFormat } from '@/utils/filters';
import { defineAsyncComponent } from 'vue';
const PageTitle = defineAsyncComponent(() => import('@/lib/PageTitle.vue'));
const PriceView = defineAsyncComponent(() => import('./lib/PriceView.vue'));
const VolumeView = defineAsyncComponent(() => import('./lib/VolumeView.vue'));
const TickerAnalyse = defineAsyncComponent(() => import('./lib/TickerAnalyse.vue'));

const CoinSort: TickerParam['SortType'] = $ref('Amount');

let CoinTickerList = $ref([]);
let AnalyseData = $ref({});

const GetCoinTickerList = () => {
  GetTickerList({
    SortType: CoinSort,
  }).then((res) => {
    if (res.Code > 0) {
      CoinTickerList = res.Data.List;
      AnalyseData = res.Data.Analyse;
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
    title: '#',
    width: 34,
    render: (_, index) => {
      return `${index + 1}`;
    },
  },
  {
    title: 'Coin',
    width: 68,
    fixed: 'left',
    key: 'CcyName',
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
</script>

<template>
  <PageTitle> Market </PageTitle>
  <div class="ListWrapper">
    <div v-if="CoinTickerList.length" class="Describe">
      OKX、Binance 综合交易量排名前 {{ CoinTickerList.length }} 的币种。 <br />
      列表数据更新时间 {{ DateFormat(CoinTickerList[0].Ts) }}
    </div>
    <div class="TableWrapper">
      <n-data-table size="small" striped :columns="columns" :data="CoinTickerList" />
    </div>
    <TickerAnalyse :Analyse="AnalyseData" />
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
</style>
