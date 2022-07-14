import { ajax_json } from '@/utils/http';

export interface HunterServer {
  Host: string;
}

interface HunterNetParam {
  ServerInfo: HunterServer;
}

export const HunterPing = (data: HunterNetParam) => {
  return ajax_json({
    url: '/hunter_net/ping',
    data: null,
    method: 'get',
    HunterNet: data.ServerInfo,
  });
};

export const GetHunterConfig = (data: HunterNetParam): Promise<any> => {
  return ajax_json({
    url: '/hunter_net/config',
    data: null,
    method: 'get',
    HunterNet: data.ServerInfo,
  });
};
