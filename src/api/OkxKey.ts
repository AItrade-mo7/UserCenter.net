import { ajax_json } from '@/utils/http';
import { Md5 } from '@/utils/tools';
import type { AddOkxKeyParam, DelOkxKeyParam } from './api.d';

export const GetOkxKeyList = () => {
  return ajax_json({
    url: '/api/private/okx_list',
    data: null,
    method: 'get',
  });
};

export const CreateOkxKey = (param: AddOkxKeyParam) => {
  const data = {
    ...param,
    Password: Md5(param.Password),
  };

  return ajax_json({
    url: '/api/private/add_okx_key',
    data,
    method: 'post',
  });
};

export const DelOkxkey = (param: DelOkxKeyParam) => {
  const data = {
    ...param,
    Password: Md5(param.Password),
  };
  return ajax_json({
    url: '/api/private/del_okxkey',
    data,
    method: 'post',
  });
};
