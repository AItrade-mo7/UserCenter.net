<template>
  <n-input
    name="Code"
    style="text-align: center"
    size="small"
    :inputProps="{ autocomplete: 'password', id: 'AuthModal_Code' }"
    placeholder="6位验证码"
    :maxlength="6"
  />
</template>
