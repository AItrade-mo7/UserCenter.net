import { ajax_json } from '@/utils/http';
import { Md5 } from '@/utils/tools';

interface SysParam {
  Password: string;
  Code: string;
  CoinServeID: string;
}

export const ReStart = (param: SysParam) => {
  const data = {
    ...param,
    Password: Md5(param.Password),
    Code: Md5(param.Code),
  };

  return ajax_json({
    url: '/CoinFundServe/sys/restart',
    data,
    method: 'post',
    CoinServeID: data.CoinServeID,
  });
};

export const Remove = (param: SysParam): Promise<any> => {
  const data = {
    ...param,
    Password: Md5(param.Password),
    Code: Md5(param.Code),
  };

  return ajax_json({
    url: '/CoinFundServe/sys/remove',
    data,
    method: 'post',
    CoinServeID: data.CoinServeID,
  });
};
