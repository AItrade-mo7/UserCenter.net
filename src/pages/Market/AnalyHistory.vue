<script setup lang="ts">
import { onMounted, defineAsyncComponent } from 'vue';
import { GetAnalyHistory } from '@/api/CoinMarket';
const PageTitle = defineAsyncComponent(() => import('@/lib/PageTitle.vue'));

let HistoryList = $ref([]);
let Current = $ref(1);
let Total = $ref(0);
let Size = $ref(300);

const GetHistoryList = (page) => {
  Current = page;
  GetAnalyHistory({
    Size: Size,
    Current: Current - 1,
    Sort: {
      TimeUnix: -1,
    },
  }).then((res) => {
    if (res.Code > 0) {
      HistoryList = res.Data.List;
      Total = res.Data.Total;
      Current = res.Data.Current + 1;
      Size = res.Data.Size;
    }
  });
};

onMounted(() => {
  GetHistoryList({});
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
    <div>最近72小时程序大盘预测结果</div>
    <div>
      {{ Current }}
      <n-pagination
        v-model:page="Current"
        size="small"
        :item-count="Total"
        :page-size="Size"
        :on-update:page="GetHistoryList"
      />
      <div v-for="item in HistoryList" class="DataBox" :class="WholeDirFormat(item.WholeDir).class">
        <n-space>
          <div>时间: <n-time :time="item.TimeUnix" /></div>
          <div>算法结果: {{ WholeDirFormat(item.WholeDir).text }}</div>
        </n-space>
      </div>
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
  display: inline-block;
  width: 324px;
  margin-right: 12px;
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
