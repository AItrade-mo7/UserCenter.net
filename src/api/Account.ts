import { ajax_json } from '@/utils/http';
import { Md5, removeToken } from '@/utils/tools';
import type { LoginParam, RegisterParam, SendCodeParam, ChangePasswordParam, EditProfileParam } from './api.d';

export const login = (data: LoginParam) => {
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

export const Register = (data: RegisterParam) => {
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

export const fetchSendCode = (data: SendCodeParam) => {
  return ajax_json({
    url: '/api/public/send_code',
    data,
    method: 'post',
  });
};

export const ChangePassword = (data: ChangePasswordParam) => {
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

export const EditProfile = (data: EditProfileParam) => {
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
