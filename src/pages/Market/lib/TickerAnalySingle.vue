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
        <span class="value">{{ item.DiffHour }} 小时 </span>
      </div>
      <div class="block">
        <span class="label">成交量</span>
        <span class="value">{{ VolumeFormat(item.Volume) }} </span>
      </div>
      <div class="block">
        <span class="label">涨幅</span>
        <span class="value" :class="CountUR(item.RosePer)">{{ item.RosePer }}% </span>
      </div>
    </n-space>
  </div>
</template>

<style lang="less" scoped>
@import '@/config/constant.less';

.TickerAnaly {
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
