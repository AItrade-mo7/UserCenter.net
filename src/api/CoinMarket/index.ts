import { ajax_json } from '@/utils/http';

export interface TickerParam {
  SortType: 'U_R24' | 'Amount';
}

export const GetTickerList = (data: TickerParam) => {
  return ajax_json({
    url: '/CoinMarket/public/Tickers',
    data,
    method: 'post',
  });
};

export interface InstParam {
  InstType: 'SPOT' | 'SWAP';
}

export const GetInstList = (data: InstParam): Promise<any> => {
  return ajax_json({
    url: '/CoinMarket/public/Inst',
    data,
    method: 'post',
  });
};

export interface GetAnalyHistoryParam {
  Size: number;
  Current: number;
  Sort: {
    [key: string]: 1 | -1;
  };
}

export const GetAnalyHistory = (data: GetAnalyHistoryParam): Promise<any> => {
  return ajax_json({
    url: '/CoinMarket/public/GetAnalyHistory',
    data: data,
    method: 'post',
  });
};
