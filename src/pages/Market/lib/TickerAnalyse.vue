<script setup lang="ts">
import { onMounted } from 'vue';
import { GetTickerAnalyse } from '@/api/CoinMarket/index';

let Data = $ref({});

const GetData = () => {
  GetTickerAnalyse().then((res) => {
    if (res.Code > 0) {
      Data = res.Data;
    }
  });
};
onMounted(() => {
  GetData();
});

const UPIndex = () => {
  if (Data.UPIndex - 50 > 0) {
    return 'green';
  }
  if (Data.UPIndex - 50 < 0) {
    return 'red';
  }
  return '';
};

const UDAvg = () => {
  if (Data.UDAvg - 0 > 0) {
    return 'green';
  }
  if (Data.UDAvg - 0 < 0) {
    return 'red';
  }
  return '';
};

const DirIndex = () => {
  const Return = {
    style: '',
    text: '震荡',
    value: Data.DirIndex,
  };

  if (Data.DirIndex - 0 > 0) {
    Return.style = 'green';
    Return.text = '上涨';
  }
  if (Data.DirIndex - 0 < 0) {
    Return.style = 'red';
    Return.text = '下跌';
  }

  return Return;
};

const CoinTicker = (label: string) => {
  const Coin = Data[label];
  const Return = {
    CcyName: Coin.CcyName,
    U_R24: Coin.U_R24,
    style: '',
    VolIdx: Coin.VolIdx,
  };

  if (Return.U_R24 - 0 > 0) {
    Return.style = 'green';
  }
  if (Return.U_R24 - 0 < 0) {
    Return.style = 'red';
  }

  return Return;
};
</script>

<template>
  <div class="TickerAnalyse" v-if="Data.MaxUP">
    <n-space class="data-wrapper">
      <div class="block">
        <span class="label">上涨指数</span>
        <span class="value" :class="UPIndex()">{{ Data.UPIndex }}%</span>
      </div>
      <div class="block">
        <span class="label">综合涨幅均值</span>
        <span class="value" :class="UDAvg()">{{ Data.UDAvg }}%</span>
      </div>
      <div class="block">
        <span class="label">市场整体情况</span>
        <span class="value" :class="DirIndex().style">{{ DirIndex().text }}</span>
      </div>

      <div class="block">
        <span class="label">最惨币</span>
        <span class="value" :class="CoinTicker('MaxDown').style">
          {{ CoinTicker('MaxDown').CcyName }}
          {{ CoinTicker('MaxDown').U_R24 }}%
        </span>
      </div>

      <div class="block">
        <span class="label">最牛币</span>
        <span class="value" :class="CoinTicker('MaxUP').style">
          {{ CoinTicker('MaxUP').CcyName }}
          {{ CoinTicker('MaxUP').U_R24 }}%
        </span>
      </div>

      <div class="block">
        <span class="label">市场时间</span>
        <span class="value"><n-time :time="Data.Ts" /></span>
      </div>
    </n-space>
    <div>持仓建议：{{ Data.Suggest }}</div>
  </div>
</template>

<style lang="less" scoped>
@import '@/config/constant.less';

.TickerAnalyse {
  background-color: antiquewhite;
  padding: 6px;
}

.value {
  color: #333;
}

.green {
  color: @greenColor;
}
.red {
  color: @redColor;
}

.block {
  font-size: 14px;
  .label {
    color: #666;
    font-size: 12px;
    &::after {
      content: ' : ';
    }
  }
}
</style>
