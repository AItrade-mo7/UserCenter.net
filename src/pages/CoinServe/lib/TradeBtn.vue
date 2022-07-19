<script setup lang="ts">
import { Buy, Sell, Close } from '@/api/CoinAI/Trade';

const props = defineProps({
  WssData: Object,
});

const BuyUp = () => {
  Buy({
    CoinServeID: props.WssData.CoinServeID,
  });
};

const SellDown = () => {
  Sell({
    CoinServeID: props.WssData.CoinServeID,
  });
};

const CloseAll = () => {
  Close({
    CoinServeID: props.WssData.CoinServeID,
  });
};

const BuySPOT = () => {
  Buy({
    CoinServeID: props.WssData.CoinServeID,
  });
};
</script>

<template>
  <div class="TradeBtn">
    <n-button type="primary" color="#18a058" @click="BuyUp">开多</n-button>
    <n-button type="primary" color="#d03050" @click="SellDown">开空</n-button>
    <n-button type="primary" @click="CloseAll">市价全平</n-button>
    <n-button type="primary" color="#0c7a43" @click="BuySPOT">现货买入</n-button>
  </div>
</template>

<style lang="less" scoped>
.TradeBtn {
  display: block;
}
</style>
