<template>
  <n-input
    type="password"
    name="Password"
    style="text-align: center"
    show-password-on="click"
    size="small"
    placeholder="请输入密码"
    :inputProps="{ autocomplete: 'password', id: 'AuthModal_Password' }"
  ></n-input>
</template>
