<script setup lang="ts">
import { onMounted, defineAsyncComponent } from 'vue';
import { GetAnalyHistory } from '@/api/CoinMarket';
const PageTitle = defineAsyncComponent(() => import('@/lib/PageTitle.vue'));

let HistoryList = $ref([]);

const GetHistoryList = () => {
  GetAnalyHistory().then((res) => {
    if (res.Code > 0) {
      HistoryList = res.Data;
      // console.log(JSON.stringify(HistoryList[HistoryList.length - 1]));
    }
  });
};

onMounted(() => {
  GetHistoryList();
});

const WholeDirFormat = (n: any) => {
  var ReturnObj = {
    text: '空仓观望',
    class: 'gray',
  };

  var Type = n - 0;

  switch (Type) {
    case 1:
      ReturnObj.text = '看涨';
      ReturnObj.class = 'green';
      break;
    case 2:
      ReturnObj.text = '震荡上涨';
      ReturnObj.class = 'green';
      break;
    case -1:
      ReturnObj.text = '看跌';
      ReturnObj.class = 'red';
      break;
    case -2:
      ReturnObj.text = '震荡下跌';
      ReturnObj.class = 'red';
      break;
    default:
      ReturnObj.text = '空仓观望';
      ReturnObj.class = 'gray';
      break;
  }

  return ReturnObj;
};
</script>

<template>
  <PageTitle> AnalyHistory </PageTitle>
  <div class="PageWrapper AnalyHistory">
    <div>最近72小时程序预测结果</div>
    <div v-for="item in HistoryList" class="DataBox" :class="WholeDirFormat(item.WholeDir).class">
      <n-space>
        <div>时间: <n-time :time="item.TimeUnix" /></div>
        <div>算法结果: {{ WholeDirFormat(item.WholeDir).text }}</div>
      </n-space>
    </div>
  </div>
</template>

<style lang="less" scoped>
@import '@/config/constant.less';

.DataBox {
  margin: 10px 0;
  border-width: 1;
  border-style: solid;
  padding: 2px 6px;
  &.green {
    border-color: @greenColor;
    color: @greenColor;
  }
  &.red {
    border-color: @redColor;
    color: @redColor;
  }
}
</style>
