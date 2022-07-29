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

const CountUR = (ur: string) => {
  const Num = parseInt(ur, 10);
  if (Num > 0) {
    return 'green';
  }
  if (Num < 0) {
    return 'red';
  }
  return '';
};
</script>

<template>
  <div class="TickerAnalyse" v-if="Data.MaxUP">
    <n-space class="data-wrapper">
      <div class="block">
        <span class="label">上涨指数</span>
        <span class="value">{{ Data.UPIndex }}</span>
      </div>
      <div class="block">
        <span class="label">综合涨幅均值</span>
        <span class="value">{{ Data.UDAvg }}%</span>
      </div>
      <div class="block">
        <span class="label">上涨趋势</span>
        <span class="value">{{ Data.UPLe }}</span>
      </div>
      <div class="block">
        <span class="label">上涨强度</span>
        <span class="value">{{ Data.UDLe }}</span>
      </div>
      <div class="block">
        <span class="label">当前市场情况</span>
        <span class="value">{{ Data.DirIndex }}</span>
      </div>

      <div class="block">
        <span class="label">最牛币</span>
        <span class="value">{{ Data.MaxUP.CcyName }} {{ Data.MaxUP.U_R24 }} </span>
      </div>

      <div class="block">
        <span class="label">最惨币</span>
        <span class="value">{{ Data.MaxDown.CcyName }} {{ Data.MaxDown.U_R24 }} </span>
      </div>

      <div class="block">
        <span class="label">市场时间</span>
        <span class="value"><n-time :time="Data.Ts" /></span>
      </div>
    </n-space>
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
