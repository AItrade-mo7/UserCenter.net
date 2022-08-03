<script setup lang="ts">
const props = defineProps({
  data: Object,
});

const CountUR = (ur: any) => {
  if (ur - 0 > 0) {
    return 'green';
  }
  if (ur - 0 < 0) {
    return 'red';
  }
  return '';
};
const CountURNum = (ur: any) => {
  if (ur - 0 > 0) {
    return `+${ur}%`;
  }
  if (ur - 0 < 0) {
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
