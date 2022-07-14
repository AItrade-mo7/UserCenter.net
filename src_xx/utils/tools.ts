import store from 'store';
import copy from 'copy-to-clipboard';
import _ from 'lodash';
import Qs from 'qs';

export const mStorage = store;

export const getUrlParams = () => {
  const url = window.location.href;
  const index = url.indexOf('?');
  if (index === -1) {
    return {};
  }
  const params = url.slice(index + 1);
  const paramObj = Qs.parse(params);
  return paramObj;
};

export const isPwa = (): boolean => {
  return ['fullscreen', 'standalone', 'minimal-ui'].some(
    (displayMode) => window.matchMedia(`(display-mode: ${displayMode})`).matches,
  );
};

export const cloneDeep = (data: any): any => {
  return JSON.parse(JSON.stringify(data));
};

export const $lcg = (obj: any, path: string, defaultValue: any = ''): any => {
  const data = _.get(obj, path, defaultValue);
  return cloneDeep(data);
};

export const getToken = (): string => {
  const token: string = mStorage.get('token');
  if (token) {
    return token;
  } else {
    return '';
  }
};

export const removeToken = () => {
  mStorage.clearAll();
};

import { getUserInfo } from '@/api/Account';
import type { resDataType } from './utils.d';
import { UserInfoStore } from '@/store';

export const setToken = async (tokenStr: string) => {
  mStorage.set('token', tokenStr);
  return new Promise<resDataType>((resolve, reject) => {
    const localToken = getToken();
    if (localToken) {
      getUserInfo().then((res) => {
        if (res.Code > 0 && res.Data.UserID) {
          UserInfoStore.update(res.Data);
          resolve(res);
        } else {
          removeToken();
          reject(res);
        }
      });
    }
  });
};

export const CopyText = (text: string) => {
  copy(text);
};

export const Md5 = (message: string): string => {
  return window.mo7Md5(message);
};

export const Encrypt = (message: string): string => {
  return window.mo7Encrypt(message);
};

export const UpLoadFileEncrypt = () => {
  const stamp = new Date().getTime();
  const time = stamp / 1000 / 30;
  const timeStamp = parseInt(`${time}`);
  const EncStr = `mo7${timeStamp}`; // mo7 + 分钟时间戳
  return window.mo7Md5(EncStr); // 然后融合MD5
};

export const Shield = (str: string): string => {
  if (!str) {
    return `****`;
  }

  let last = str.substring(str.length - 3);

  const pre = str.substring(0, 3);

  if (str.indexOf('@') > -1) {
    last = str.split('@')[1];
  }

  return `${pre}***${last}`;
};

/* 
apikey = "a5f97f33-4958-4b54-820d-43d7b1e6f189"
secretkey = "AAFE8D2F84D88E50F1709F95C93B34D8"
IP = "123.123.123.123"
备注名 = "备注名"
权限 = "只读/提现/交易"
*/
export const ParseOkxKey = (str: string) => {
  const str_arr = str.split('\r\n');

  const str_obj = {
    ApiKey: '',
    SecretKey: '',
    IP: '',
    Name: '',
    Note: '',
  };
  for (const iterator of str_arr) {
    const valStr = iterator.replace(/\s*/g, '');
    const str_arr = valStr.split('=');
    if (str_arr.length < 2) {
      break;
    }
    const Key = str_arr[0];
    const Val_arr = str_arr[1].match(/"(.+)"/gi);

    if (!Val_arr) {
      break;
    }
    if (Val_arr.length < 1) {
      break;
    }
    const Val_str = Val_arr[0];
    const Val = Val_str.replace(/"/g, '');

    switch (Key) {
      case 'apikey':
        str_obj.ApiKey = Val;
        break;
      case 'secretkey':
        str_obj.SecretKey = Val;
        break;
      case 'IP':
        str_obj.IP = Val;
        break;
      case '备注名':
        str_obj.Name = Val;
        break;
      case '权限':
        str_obj.Note = Val;
        break;
      default:
        break;
    }
  }

  return str_obj;
};
