export const verifyConfig = (label: string, val: string, val2?: string): string => {
  const passwordReg = /^[a-zA-Z0-9]{6,16}$/g;
  const emailReg = /^[a-z0-9]+([._\\-]*[a-z0-9])*@([a-z0-9]+[-a-z0-9]*[a-z0-9]+.){1,63}[a-z0-9]+$/;
  const nickNameReg = /^[a-zA-Z0-9_\u4e00-\u9fa5]{1,12}$/g;

  if (label === 'Email') {
    if (!emailReg.test(val)) {
      return '请输入有效的邮箱地址';
    }
  }

  if (label === 'NickName') {
    if (!nickNameReg.test(val)) {
      return '1-12位汉字、数字、字母、下划线';
    }
  }

  if (label === 'NewEmail') {
    if (!emailReg.test(val)) {
      return '请输入有效的邮箱地址';
    }
  }

  if (label === 'OldEmailCode') {
    if (val.length < 6) {
      return '请输入有效的验证码';
    }
  }

  if (label === 'NewEmailCode') {
    if (val.length < 6) {
      return '请输入有效的验证码';
    }
  }

  if (label === 'Code') {
    if (val.length < 6) {
      return '请输入6位验证码';
    }
  }

  if (label === 'AgainPassword') {
    if (val !== val2) {
      return '两次输入的密码不一致';
    }
  }

  if (label.indexOf('Password') > -1) {
    if (!passwordReg.test(val)) {
      return '请输入6-16位字母或数字';
    }
  }

  return null;
};

export const AITradeVerifyConfig = (label: string, val: string): string => {
  const nickNameReg = /^[a-zA-Z0-9_\u4e00-\u9fa5]{1,12}$/g;

  if (label === 'Name') {
    if (!nickNameReg.test(val)) {
      return '1-12位汉字、数字、字母、下划线';
    }
  }

  if (label === 'ApiKey') {
    if (val.length < 30) {
      return '请输入有效的 ApiKey ';
    }
  }
  if (label === 'SecretKey') {
    if (val.length < 30) {
      return '请输入有效的 SecretKey ';
    }
  }
  if (label === 'Passphrase') {
    if (val.length < 2) {
      return '请输入 Passphrase ';
    }
  }

  return null;
};
