<script setup lang="ts">
const props = defineProps({
  data: Object,
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
const CountURNum = (ur: string) => {
  const Num = parseInt(ur, 10);
  if (Num > 0) {
    return `+${ur}%`;
  }
  if (Num < 0) {
    return `${ur}%`;
  }
  return `${ur}%`;
};
</script>

<template>
  <div class="Price" v-if="props.data && props.data.CcyName">
    <div class="ur" :class="CountUR(props.data.U_R24)">{{ CountURNum(props.data.U_R24) }}</div>
    <div class="last">
      {{ props.data.Last }}
    </div>
  </div>
</template>

<style lang="less" scoped>
.Price {
  text-align: right;
  font-size: 14px;
  .ur {
    color: #333;
  }
  .red {
    color: red;
  }
  .green {
    color: green;
  }
  .last {
    color: #999;
  }
}
</style>
