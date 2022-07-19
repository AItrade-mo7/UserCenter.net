import { ajax_json } from '@/utils/http';

interface TradeParam {
  CoinServeID: string;
}

export const Buy = (data: TradeParam) => {
  return ajax_json({
    url: '/CoinAI/Order/Buy',
    data,
    method: 'post',
  });
};

export const Sell = (data: TradeParam) => {
  return ajax_json({
    url: '/CoinAI/Order/Sell',
    data,
    method: 'post',
  });
};

export const Close = (data: TradeParam) => {
  return ajax_json({
    url: '/CoinAI/Order/Close',
    data,
    method: 'post',
  });
};
