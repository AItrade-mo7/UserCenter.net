<script setup lang="ts">
import { verifyConfig } from '@/utils/verify';
import { fetchSendCode } from '@/api/Account';

let num = $ref(0);

const props = defineProps({
  Email: String,
  Action: String,
});

let timer: any = null;

const FilterNum = (): string => {
  if (num === 0) {
    SubmitStatus = false;
    return '发送验证码';
  } else {
    SubmitStatus = true;
    return `重新发送${num}`;
  }
};

let SubmitStatus: boolean = $ref(false);
const handleClick = async () => {
  clearInterval(timer);

  const verifyResult = verifyConfig('Email', props.Email);
  if (verifyResult) {
    window.$message.error(verifyResult);
    return;
  }

  // 开始倒计时
  num = 60;
  timer = setInterval(() => {
    num--;
  }, 1000);

  const res = await fetchSendCode({
    Email: props.Email,
    Action: props.Action,
  }).catch(() => {
    SubmitStatus = false;
    num = 0;
    clearInterval(timer);
  });

  if (!res) {
    return;
  }

  if (res.Code > 0) {
    window.$message.success(res.Msg);
    if (num === 0) {
      num = 60;
      clearInterval(timer);
    }
  } else {
    num = 0;
    clearInterval(timer);
  }
};
</script>

<template>
  <n-button type="primary" :Email="props.Email" :disabled="SubmitStatus" @click="handleClick" ghost>{{
    FilterNum()
  }}</n-button>
</template>
