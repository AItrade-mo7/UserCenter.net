import { ajax_json } from '@/utils/http';
import { Md5, removeToken } from '@/utils/tools';
import type { loginType } from './api.d';

export const login = (data: loginType) => {
  const param = {
    ...data,
    Password: Md5(data.Password),
  };

  return ajax_json({
    url: '/api/public/login',
    data: param,
    method: 'post',
  });
};

export const logout = () => {
  removeToken();
  window.location.replace('/Login');
};

export const getUserInfo = () => {
  return ajax_json({
    url: '/api/private/get_user_info',
    data: null,
    method: 'get',
  });
};

interface registerData {
  Email: string;
  Code: string;
}
export const Register = (data: registerData) => {
  const param = {
    ...data,
    Code: Md5(data.Code),
  };

  return ajax_json({
    url: '/api/public/register',
    data: param,
    method: 'post',
  });
};

interface sendCodeData {
  Email: string;
  Action: string;
}
export const fetchSendCode = (data: sendCodeData) => {
  return ajax_json({
    url: '/api/public/send_code',
    data,
    method: 'post',
  });
};

interface editPasswordData {
  Email: string;
  Code: string;
  Password: string;
  AgainPassword: string;
}
export const ChangePassword = (data: editPasswordData) => {
  const param = {
    ...data,
    Code: Md5(data.Code),
    Password: Md5(data.Password),
    AgainPassword: Md5(data.AgainPassword),
  };

  return ajax_json({
    url: '/api/public/change_password',
    data: param,
    method: 'post',
  });
};

interface editProfileParam {
  OldEmailCode: string;
  NewEmail: string;
  NewEmailCode: string;
  Avatar: string;
  NickName: string;
  SecurityCode: string;
  Password: string;
}
export const EditProfile = (data: editProfileParam) => {
  const param = {
    ...data,
    OldEmailCode: Md5(data.OldEmailCode),
    NewEmailCode: Md5(data.NewEmailCode),
    Password: Md5(data.Password),
  };

  return ajax_json({
    url: '/api/private/edit_profile',
    data: param,
    method: 'post',
  });
};
