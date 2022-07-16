<script setup lang="ts">
import { defineAsyncComponent } from 'vue';
const PageTitle = defineAsyncComponent(() => import('@/lib/PageTitle.vue'));
</script>

<template>
  <PageTitle>StockServe</PageTitle>
  <div class="PageWrapper">
    股市交易相关功能尚未完成。
    <br />
    该项目目前为无偿服务
    <br />
    请开发者喝杯茶:
    <img class="wx_proceeds" src="//file.mo7.cc/static/img/wx_proceeds.jpg" alt="" />
    <br />
    <a class="lineHeight" href="https://github.com/mo7cc">开发者的 Github 主页</a>
    <br />
    <a class="lineHeight" href="https://mo7.cc/">主页</a>
  </div>
</template>

<style lang="less" scoped>
@import '@/config/constant.less';
.wx_proceeds {
  width: 200px;
  display: block;
}
.lineHeight {
  color: @mainColor;
}
</style>
