import { ajax_json } from '@/utils/http';
import { Md5 } from '@/utils/tools';

export interface AIFundServer {
  Host: string;
}

interface AIFundNetParam {
  ServerInfo: AIFundServer;
}

interface SysParam {
  Password: string;
  Code: string;
  ServerInfo: AIFundNetParam;
}

export const ReStart = (param: SysParam) => {
  const data = {
    ...param,
    Password: Md5(param.Password),
    Code: Md5(param.Code),
  };

  return ajax_json({
    url: '/CoinServe/sys/restart',
    data,
    method: 'post',
    AIFundNet: data.ServerInfo,
  });
};

export const Remove = (param: SysParam): Promise<any> => {
  const data = {
    ...param,
    Password: Md5(param.Password),
    Code: Md5(param.Code),
  };

  return ajax_json({
    url: '/CoinServe/sys/remove',
    data,
    method: 'post',
    AIFundNet: data.ServerInfo,
  });
};
