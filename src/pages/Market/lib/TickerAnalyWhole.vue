<script setup lang="ts">
const props = defineProps({
  Analy: Object,
});

const UPIndex = () => {
  if (props.Analy.UPIndex - 50 > 0) {
    return 'green';
  }
  if (props.Analy.UPIndex - 50 < 0) {
    return 'red';
  }
  return '';
};

const UDAvg = () => {
  if (props.Analy.UDAvg - 0 > 0) {
    return 'green';
  }
  if (props.Analy.UDAvg - 0 < 0) {
    return 'red';
  }
  return '';
};

const DirIndex = () => {
  const Return = {
    style: '',
    text: '震荡',
    value: props.Analy.DirIndex,
  };

  if (props.Analy.DirIndex - 0 > 0) {
    Return.style = 'green';
    Return.text = '上涨';
  }
  if (props.Analy.DirIndex - 0 < 0) {
    Return.style = 'red';
    Return.text = '下跌';
  }

  return Return;
};

const CountUR = (ur: any) => {
  if (ur - 0 > 0) {
    return 'green';
  }
  if (ur - 0 < 0) {
    return 'red';
  }
};
</script>

<template>
  <div class="TickerAnaly" v-if="props.Analy.MaxUP">
    <n-space class="data-wrapper">
      <div class="block">
        <span class="label">时间切片</span>
        <span class="value hour">{{ props.Analy.DiffHour }}hour</span>
      </div>
      <div class="block">
        <span class="label">上涨指数</span>
        <span class="value ur" :class="UPIndex()">{{ props.Analy.UPIndex }}%</span>
      </div>
      <div class="block">
        <span class="label">综合涨幅均值</span>
        <span class="value ur" :class="UDAvg()">{{ props.Analy.UDAvg }}%</span>
      </div>
      <div class="block">
        <span class="label">市场整体情况</span>
        <span class="value upIndex" :class="DirIndex().style">{{ DirIndex().text }}</span>
      </div>

      <div class="block">
        <span class="label">最惨币</span>
        <span class="value CoinUR" :class="CountUR(props.Analy.MaxDown.RosePer)">
          {{ props.Analy.MaxDown.CcyName }} {{ props.Analy.MaxDown.RosePer }}%
        </span>
      </div>

      <div class="block">
        <span class="label">最牛币</span>
        <span class="value CoinUR" :class="CountUR(props.Analy.MaxUP.RosePer)">
          {{ props.Analy.MaxUP.CcyName }} {{ props.Analy.MaxUP.RosePer }}%
        </span>
      </div>

      <div class="block">
        <span class="label">数据时间</span>
        <span class="value timeRange">
          【<n-time :time="props.Analy.StartTimeUnix" />
          至
          <n-time :time="props.Analy.EndTimeUnix" />】
        </span>
      </div>
    </n-space>
  </div>
</template>

<style lang="less" scoped>
@import '@/config/constant.less';

.TickerAnaly {
  background-color: antiquewhite;
  padding: 6px;
  margin-bottom: 6px;
}

.value {
  color: #333;
  display: inline-block;
  word-break: break-all;
  max-width: 100%;
  &.hour {
    width: 56px;
  }
  &.upIndex {
    width: 40px;
  }
  &.ur {
    width: 60px;
  }
  &.CoinUR {
    width: 112px;
  }
  &.timeRange {
    width: 328px;
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
