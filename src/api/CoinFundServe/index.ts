import { ajax_json } from '@/utils/http';

export interface AIFundServer {
  Host: string;
}

interface AIFundNetParam {
  ServerInfo: AIFundServer;
}

export const AIFundPing = (data: AIFundNetParam) => {
  return ajax_json({
    url: '/CoinFundServe/ping',
    data: null,
    method: 'get',
    AIFundNet: data.ServerInfo,
  });
};

export const GetAIFundConfig = (data: AIFundNetParam): Promise<any> => {
  return ajax_json({
    url: '/CoinFundServe/config',
    data: null,
    method: 'get',
    AIFundNet: data.ServerInfo,
  });
};
