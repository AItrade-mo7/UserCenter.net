<script setup lang="ts">
import { defineAsyncComponent } from 'vue';
const PageTitle = defineAsyncComponent(() => import('@/lib/PageTitle.vue'));
</script>

<template>
  <PageTitle>StockServe</PageTitle>
  <div class="PageWrapper">
    股市交易相关功能尚未完成。
    <br />
    该程序当前为免费服务
    <br />
    请开发者喝杯茶:
    <img class="wx_proceeds" src="//file.mo7.cc/static/img/wx_proceeds.jpg" alt="" />
    <br />
    <br />
    <a class="lineHeight" href="https://github.com/mo7cc" target="_blank">开发者的 Github 主页</a>
    <br />
    <br />
    <a class="lineHeight" href="https://mo7.cc/" target="_blank">开发者 个人 主页</a>
    <br />
    <br />
    <a class="lineHeight" href="https://github.com/AITrade-mo7" target="_blank">项目仓库(目前为闭源状态)</a>
    <br />
    <br />
    <a class="lineHeight" href="https://meichangliang.feishu.cn/docs/doccnWBXTCgLrw3Ez40PuVw8FYd" target="_blank">
      程序架构图
    </a>
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
