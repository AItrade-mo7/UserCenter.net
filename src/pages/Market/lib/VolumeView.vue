<script setup lang="ts">
import { defineProps } from 'vue';
import { VolumeFormat } from '@/utils/filters';

const props = defineProps({
  Volume: String,
  Bourse: String, // 'Binance' || 'OKX' || 'ALL'
});
</script>

<template>
  <div class="VolumeView">
    <div :class="props.Bourse">{{ VolumeFormat(props.Volume) }}</div>
  </div>
</template>

<style lang="less" scoped>
.VolumeView {
  font-size: 14px;
}
</style>
