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
  <div class="TickerAnaly" v-if="props.Analy.MaxUP">
    <n-space class="data-wrapper">
      <div class="block">
        <span class="label">上涨指数</span>
        <span class="value" :class="UPIndex()">{{ props.Analy.UPIndex }}%</span>
      </div>
      <div class="block">
        <span class="label">综合涨幅均值</span>
        <span class="value" :class="UDAvg()">{{ props.Analy.UDAvg }}%</span>
      </div>
      <div class="block">
        <span class="label">市场整体情况</span>
        <span class="value" :class="DirIndex().style">{{ DirIndex().text }}</span>
      </div>

      <div class="block">
        <span class="label">最惨币</span>
        <span class="value" :class="CountUR(props.Analy.MaxDown.RosePer)">
          {{ props.Analy.MaxDown.CcyName }} {{ props.Analy.MaxDown.RosePer }}%
        </span>
      </div>

      <div class="block">
        <span class="label">最牛币</span>
        <span class="value" :class="CountUR(props.Analy.MaxUP.RosePer)">
          {{ props.Analy.MaxUP.CcyName }} {{ props.Analy.MaxUP.RosePer }}%
        </span>
      </div>

      <div class="block">
        <span class="label">数据时间</span>
        <span class="value">
          【<n-time :time="props.Analy.StartTimeUnix" />
          至
          <n-time :time="props.Analy.StartTimeUnix" />】
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
