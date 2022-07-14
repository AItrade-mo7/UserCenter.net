<script setup lang="ts">
import { useRouter } from 'vue-router';
import { defineAsyncComponent } from 'vue';
const XIcon = defineAsyncComponent(() => import('@/lib/XIcon.vue'));

const $router = useRouter();

const props = defineProps({
  Back: {
    type: Boolean,
    default() {
      return true;
    },
  },
});

const back = () => {
  $router.go(-1);
};
</script>

<template>
  <div class="PageTitle">
    <n-button v-if="props.Back" size="tiny" strong secondary circle type="primary" @click="back">
      <template #icon>
        <XIcon name="ArrowLeftOutlined" />
      </template>
    </n-button>
    <div class="PageTitle__text">
      <slot></slot>
    </div>
    <div class="PageTitle__after">
      <slot name="after"></slot>
    </div>
  </div>
</template>

<style lang="less" scoped>
@import '@/config/constant.less';

.PageTitle {
  display: flex;
  align-items: center;
  justify-content: flex-start;

  .n-button {
    margin-right: 4px;
  }
}

.PageTitle__text {
  color: rgba(0, 0, 0, 0.7);
  display: flex;
  align-items: center;
  justify-content: flex-start;
  font-size: 14px;
  font-weight: 900;
}
.PageTitle__after {
  display: flex;
  align-items: center;
  justify-content: flex-start;
  margin-left: 4px;
}
</style>
