<script setup lang="ts">
import { VolumeFormat } from '@/utils/filters';
const props = defineProps({
  Single: Array<any>,
});

console.log(props.Single);

const CountUR = (ur: any) => {
  if (ur -0 > 0) {
    return 'green';
  }
  if (ur - 0 < 0) {
    return 'red';
  }
  return "";
};
</script>

<template>
  <div class="TickerAnaly" v-if="props.Single.length">
    <n-space class="data-wrapper" v-for="item in props.Single">
      <div class="block">
        <span class="label">时间切片</span>
        <span class="value hour">{{ item.DiffHour }} hour </span>
      </div>
      <div class="block">
        <span class="label">总成交量</span>
        <span class="value vol">{{ VolumeFormat(item.Volume) }} </span>
      </div>
      <div class="block">
        <span class="label">小时平均成交量</span>
        <span class="value vol">{{ VolumeFormat(item.VolumeAvg) }} </span>
      </div>
      <div class="block">
        <span class="label">涨幅</span>
        <span class="value ur" :class="CountUR(item.RosePer)">{{ item.RosePer }}% </span>
      </div>
      <div class="block">
        <span class="label">最高振幅</span>
        <span class="value ur">{{ item.HLPerMax }}%</span>
      </div>
      <div class="block">
        <span class="label">平均振幅</span>
        <span class="value ur">{{ item.HLPerAvg }}</span>
      </div>
      <div class="block">
        <span class="label">最高价</span>
        <span class="value">{{ item.H }}</span>
      </div>
      <div class="block">
        <span class="label">最低价</span>
        <span class="value">{{ item.L }}</span>
      </div>
    </n-space>
  </div>
</template>

<style lang="less" scoped>
@import '@/config/constant.less';

.data-wrapper {
  background-color: antiquewhite;
  padding: 6px;
  margin-bottom: 6px;
}

.value {
  color: #333;
  display: inline-block;
  max-width: 100%;
  word-break: break-all;
  &.hour {
    width: 56px;
  }
  &.vol {
    width: 72px;
  }
  &.ur {
    width: 66px;
  }
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
