import { ajax_json } from '@/utils/http';

interface TickerParam {
  SortType: 'U_R24' | 'Amount';
}

export const GetTickerList = (data: TickerParam) => {
  return ajax_json({
    url: '/CoinMarket/public/Tickers',
    data,
    method: 'post',
  });
};

interface InstParam {
  InstType: 'SPOT' | 'SWAP';
}

export const GetInstList = (data: InstParam): Promise<any> => {
  return ajax_json({
    url: '/CoinMarket/public/Inst',
    data,
    method: 'post',
  });
};
