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

export const GetTickerAnalyse = (): Promise<any> => {
  return ajax_json({
    url: '/CoinMarket/public/TickerAnalyse',
    data: null,
    method: 'post',
  });
};
