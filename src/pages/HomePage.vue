<script setup lang="ts">
import { TopBarStore } from '@/store';
import { Logo } from '@/config/constant';
import { defineAsyncComponent } from 'vue';

const XIcon = defineAsyncComponent(() => import('@/lib/XIcon.vue'));
const PageTitle = defineAsyncComponent(() => import('@/lib/PageTitle.vue'));
</script>

<template>
  <PageTitle :Back="false">Home</PageTitle>
  <div class="PageWrapper">
    <img class="logo" :src="Logo" alt="" srcset="" />
    <div class="appName"><span class="aiText">AI</span><span class="fundText">Trade</span></div>
    <div>
      基于<span class="fundText">大数据</span>和<span class="fundText">深度学习</span> , 打造属于您个人的 AI 交易服务。
    </div>
    <br />

    <br />
    <div>
      <RouterLink to="/CoinServe">
        <n-button type="primary">
          <template #icon>
            <XIcon name="OkxLogo" />
          </template>
          数字货币 CoinAI ->
        </n-button>
      </RouterLink>
    </div>
    <br />
    <div>
      <RouterLink to="/StockServe">
        <n-button type="primary" ghost color="#f00">
          <template #icon>
            <XIcon name="ZTSLogo" />
          </template>
          股票交易 StockAI ->
        </n-button>
      </RouterLink>
    </div>
    <br />
    <div>
      <RouterLink to="/Market">
        <n-button strong secondary type="error">
          <template #icon>
            <XIcon name="BarChartOutlined" />
          </template>
          查看行情
        </n-button>
      </RouterLink>
      <n-button type="info" @click="TopBarStore.open()">
        <template #icon>
          <XIcon name="MenuFoldOutlined" />
        </template>
        开始
      </n-button>
    </div>
  </div>
</template>

<style lang="less" scoped>
@import '@/config/constant.less';

.PageWrapper {
  text-align: center;
}

.logo {
  width: 30%;
}
.appName {
  font-size: 36px;
}
.aiText {
  color: @blueColor;
}

.fundText {
  color: @mainColor;
}
</style>
