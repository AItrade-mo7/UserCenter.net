import { ajax_json } from '@/utils/http';
import { Md5 } from '@/utils/tools';

export const GetOkxKeyList = () => {
  return ajax_json({
    url: '/api/private/okx_list',
    data: null,
    method: 'get',
  });
};

interface AddOkxKeyParam {
  ApiKey: string;
  SecretKey: string;
  Passphrase: string;
  IP: string;
  Name: string;
  Note: string;
  Password: string;
}
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

interface DelOkxkeyParam {
  OkxKeyID: string;
  Password: string;
}

export const DelOkxkey = (param: DelOkxkeyParam) => {
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
