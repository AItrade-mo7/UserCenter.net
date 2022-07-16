<script setup lang="ts">
import { NewSocket } from '@/api/CoinAI/CoinAIWss';
import { onMounted, onUnmounted } from 'vue';
const props = defineProps({
  config: Object,
});

let wssData = $ref({});
let socketObj: any;

onMounted(() => {
  if (props.config.AppEnv.CoinServeID) {
    socketObj = NewSocket({
      Host: props.config.AppEnv.CoinServeID,
      MessageEvent(res) {
        if (res.Response.Code == 1) {
          wssData = res.Response.Data;
        }
      },
    });
  }
});

onUnmounted(() => {
  socketObj.close();
});
</script>

<template>
  <div class="InfoOk" v-if="wssData.DataSource">
    {{ wssData }}
  </div>
  <div v-else>正在连接系统。。。</div>
</template>

<style lang="less" scoped></style>
