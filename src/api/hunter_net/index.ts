import { ajax_json } from '@/utils/http';

export interface AIFundServer {
  Host: string;
}

interface AIFundNetParam {
  ServerInfo: AIFundServer;
}

export const AIFundPing = (data: AIFundNetParam) => {
  return ajax_json({
    url: '/AIFund_net/ping',
    data: null,
    method: 'get',
    AIFundNet: data.ServerInfo,
  });
};

export const GetAIFundConfig = (data: AIFundNetParam): Promise<any> => {
  return ajax_json({
    url: '/AIFund_net/config',
    data: null,
    method: 'get',
    AIFundNet: data.ServerInfo,
  });
};
