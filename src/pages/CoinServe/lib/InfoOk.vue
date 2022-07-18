<script setup lang="ts">
import { NewSocket } from '@/api/CoinAI/CoinAIWss';
import { onMounted, onUnmounted } from 'vue';
import type { WssDataType } from '@/api/CoinAI/WssType.d';
import type { resDataType } from '@/utils/utils.d';
import { defineAsyncComponent } from 'vue';
const TradeBtn = defineAsyncComponent(() => import('./TradeBtn.vue'));

const props = defineProps({
  config: Object,
});

let WssData: WssDataType = $ref({});
let ResData: resDataType = $ref({});
let socketObj: any;

onMounted(() => {
  if (props.config.AppEnv.CoinServeID) {
    socketObj = NewSocket({
      Host: props.config.AppEnv.CoinServeID,
      MessageEvent(res) {
        ResData = res.Response;
        if (res.Response.Code == 1) {
          WssData = res.Response.Data;
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
  <div>{{ WssData }}</div>
  <div class="InfoOk" v-if="WssData.DataSource">
    <div class="title">系统状态</div>
    <n-space class="data-wrapper">
      <div class="block">
        <span class="label">连接状态</span>
        <span class="value green">{{ ResData.Msg }}</span>
      </div>
      <div class="block">
        <span class="label">系统时间</span>
        <span class="value"><n-time :time="WssData.SysTime" /></span>
      </div>
    </n-space>
  </div>
  <TradeBtn />
  <div v-if="!WssData.DataSource">
    正在连接到 <span class="lineHight">{{ props.config.AppEnv.CoinServeID }}</span> .....
  </div>
</template>

<style lang="less" scoped>
@import '@/config/constant.less';
.lineHight {
  color: @mainColor;
}

.title {
  font-weight: 700;
  font-size: 18px;
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

.data-wrapper {
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
}
</style>
