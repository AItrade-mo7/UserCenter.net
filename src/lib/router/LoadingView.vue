<script setup lang="ts">
import { LoadingStore } from '@/store';
import { useMessage, useNotification, useDialog } from 'naive-ui';
window.$message = useMessage();
window.$notification = useNotification();
window.$dialog = useDialog();
</script>

<template>
  <Transition>
    <div v-if="LoadingStore.status" class="Loading">
      <div class="Loading__icon"></div>
      <div class="info">{{ LoadingStore.info }}</div>
    </div>
  </Transition>
</template>

<style lang="less" scoped>
@import '@/config/constant.less';
.Loading {
  position: fixed;
  left: 0;
  top: 0;
  width: 100vw;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: rgba(3, 3, 3, 0.4);
  z-index: 3000;
  flex-direction: column;
}

.info {
  margin-top: 8px;
  font-size: 14px;
  color: @mainColor;
  max-width: 90%;
  word-break: break-all;
  opacity: 0.7;
}

.Loading__icon {
  display: block;
  width: 36px;
  height: 36px;
  animation: AppLoader 1s infinite linear;
  background-color: rgba(240, 185, 11, 0.7);
}
@keyframes AppLoader {
  0% {
    transform: rotate(0deg);
    border-radius: 50%;
  }

  50% {
    transform: rotate(90deg);
    border-radius: 0%;
  }

  100% {
    transform: rotate(180deg);
    border-radius: 50%;
  }
}
</style>
