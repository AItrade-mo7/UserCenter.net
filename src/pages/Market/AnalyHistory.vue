<script setup lang="ts">
import { onMounted, defineAsyncComponent } from 'vue';
import { GetAnalyHistory } from '@/api/CoinMarket';
import { cloneDeep } from '@/utils/tools';
const PageTitle = defineAsyncComponent(() => import('@/lib/PageTitle.vue'));
const ListPage = defineAsyncComponent(() => import('./ListPage.vue'));

let HistoryList = $ref([]);
let Current = $ref(0);
let Total = $ref(0);
let Size = $ref(300);

const GetHistoryList = (page: number) => {
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
  GetHistoryList(1);
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

// 详情展示

let DrawerStatus = $ref(false);
let AnalyInfo = $ref({});
const CheckItemFunc = (item) => {
  AnalyInfo = cloneDeep(item);
  showDrawer();
};

const closeDrawer = () => {
  DrawerStatus = false;
};
const showDrawer = () => {
  DrawerStatus = true;
};
</script>

<template>
  <PageTitle> AnalyHistory </PageTitle>
  <div class="PageWrapper AnalyHistory">
    <div>最近72小时程序大盘预测结果</div>
    <n-pagination
      v-model:page="Current"
      size="small"
      :item-count="Total"
      :page-size="Size"
      :on-update:page="GetHistoryList"
    />
    <div>
      <div v-for="item in HistoryList" class="DataBox" :class="WholeDirFormat(item.WholeDir).class">
        <n-space>
          <div>时间: <n-time :time="item.TimeUnix" /></div>
          <div>算法结果: {{ WholeDirFormat(item.WholeDir).text }}</div>
        </n-space>
        <n-button class="CheckBtn" size="small" @click="CheckItemFunc(item)">查看</n-button>
      </div>
    </div>

    <n-drawer
      display-directive="show"
      :auto-focus="false"
      height="80%"
      placement="top"
      :show="DrawerStatus"
      :on-mask-click="closeDrawer"
    >
      <n-drawer-content class="TopBarDrawer">
        <ListPage :AnalyInfo="AnalyInfo"></ListPage>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<style lang="less" scoped>
@import '@/config/constant.less';

.DataBox {
  margin: 10px 0;
  border-width: 1;
  border-style: solid;
  padding: 2px 6px;
  margin-right: 12px;
  display: flex;
  align-items: center;
  justify-content: flex-start;
  .CheckBtn {
    display: none;
    margin-left: 16px;
  }
  &:hover .CheckBtn {
    display: block;
  }
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
