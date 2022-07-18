import { ajax_json } from '@/utils/http';

export const Buy = () => {
  return ajax_json({
    url: '/CoinAI/Order/Buy',
    data: null,
    method: 'post',
  });
};

export const Sell = () => {
  return ajax_json({
    url: '/CoinAI/Order/Sell',
    data: null,
    method: 'post',
  });
};

export const Close = () => {
  return ajax_json({
    url: '/CoinAI/Order/Close',
    data: null,
    method: 'post',
  });
};
