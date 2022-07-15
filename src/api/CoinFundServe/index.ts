import { ajax_json } from '@/utils/http';

interface AIFundNetParam {
  CoinServeID: string;
}

export const CoinFundPing = (data: AIFundNetParam) => {
  return ajax_json({
    url: '/CoinFundServe/ping',
    data,
    method: 'get',
  });
};

export const GetCoinFundConfig = (data: AIFundNetParam): Promise<any> => {
  return ajax_json({
    url: '/CoinFundServe/config',
    data,
    method: 'get',
  });
};
