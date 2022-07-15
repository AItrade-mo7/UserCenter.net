import axios from 'axios';
import { res_dispose } from './res_dispose';
import { getToken, Encrypt } from './tools';
import { LoadingStore } from '@/store';

const service = axios.create();

const origin = window.location.origin;

const axios_baseURL = null;
if (origin.includes('localhost')) {
  // axios_baseURL = '//localhost:9004/';
} else if (origin.includes('xxxx')) {
  // baseUrl = '';
} else {
  // baseUrl = '';
}

(() => {
  // 需要初始化
  service.defaults.baseURL = axios_baseURL; // 默认请求的 baseUrl
  service.defaults.timeout = 20000; // 超时 20 秒
  // 请求拦截
  service.interceptors.request.use(
    (config) => {
      LoadingStore.open(`正在请求: ${config.url}`);
      return config;
    },
    (error) => {
      LoadingStore.close();
      // window.$message.error(error.response.status);
      return Promise.reject(error);
    },
  );

  // 响应拦截
  service.interceptors.response.use(
    (response: any) => {
      LoadingStore.close();
      return res_dispose(response);
    },
    (error) => {
      LoadingStore.close();
      // window.$message.error(error.response.status);
      return Promise.reject(error);
    },
  );
})();

import type { axiosParam, resDataType } from './utils.d';
const ajax_json = (param: axiosParam): Promise<resDataType> => {
  const config: axiosParam = {
    headers: {
      Token: getToken(),
      'Auth-Encrypt': Encrypt(param.url + window.navigator.userAgent),
    },
    ...param,
  };

  if (param?.data?.CoinServeID) {
    config.headers['Coin-Serve-ID'] = param.data.CoinServeID;
  }

  // 请求参数转换
  if (config.method === 'get' || !config.method) {
    config.params = config.data;
    delete config.data;
  }

  return service(config) as any;
};

export { ajax_json, axios };
