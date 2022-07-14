import { ajax_json } from '@/utils/http';
import { Md5 } from '@/utils/tools';

export interface HunterServer {
  Host: string;
}

interface HunterNetParam {
  ServerInfo: HunterServer;
}

interface SysParam {
  Password: string;
  Code: string;
  ServerInfo: HunterNetParam;
}

export const ReStart = (param: SysParam) => {
  const data = {
    ...param,
    Password: Md5(param.Password),
    Code: Md5(param.Code),
  };

  return ajax_json({
    url: '/hunter_net/sys/restart',
    data,
    method: 'post',
    HunterNet: data.ServerInfo,
  });
};

export const Remove = (param: SysParam): Promise<any> => {
  const data = {
    ...param,
    Password: Md5(param.Password),
    Code: Md5(param.Code),
  };

  return ajax_json({
    url: '/hunter_net/sys/remove',
    data,
    method: 'post',
    HunterNet: data.ServerInfo,
  });
};
