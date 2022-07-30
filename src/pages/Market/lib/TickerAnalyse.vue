<script setup lang="ts">
import { cloneDeep } from '@/utils/tools';

const props = defineProps({
  Analyse: Object,
});

const Data = cloneDeep(props.Analyse);

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
        <span class="label">数据时间</span>
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
