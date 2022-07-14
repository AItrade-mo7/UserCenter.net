<template>
  <n-input
    type="password"
    name="Password"
    size="small"
    style="text-align: center; margin-bottom: 12px"
    show-password-on="click"
    placeholder="请输入密码"
    :inputProps="{ autocomplete: 'password', id: 'AuthModal_Password' }"
  ></n-input>

  <n-input
    name="Code"
    style="text-align: center"
    size="small"
    :inputProps="{ autocomplete: 'password', id: 'AuthModal_Code' }"
    placeholder="6位验证码"
    :maxlength="6"
  />
</template>
