<script setup lang="ts">
import { Logo } from '@/config/constant';
import { defineAsyncComponent } from 'vue';
const XIcon = defineAsyncComponent(() => import('@/lib/XIcon.vue'));
const PageTitle = defineAsyncComponent(() => import('@/lib/PageTitle.vue'));
</script>

<template>
  <PageTitle :Back="false">Home</PageTitle>
  <div class="PageWrapper">
    <img class="logo" :src="Logo" alt="" srcset="" />
    <div class="appName"><span class="aiText">AI</span> <span class="fundText">Fund</span></div>
    <div>
      基于<span class="fundText">大数据</span>和<span class="fundText">深度学习</span> , 打造属于您个人的 AI 基金服务。
    </div>

    <h3>开始：</h3>
    <RouterLink to="/CoinServe">
      <n-button type="primary">
        <template #icon>
          <XIcon name="OkxLogo" />
        </template>
        我是币圈用户
      </n-button>
    </RouterLink>

    <RouterLink to="/CoinServe">
      <n-button type="primary">
        <template #icon>
          <XIcon name="OkxLogo" />
        </template>
        我是股市用户
      </n-button>
    </RouterLink>
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
