import { ajax_json } from '@/utils/http';
import { Md5 } from '@/utils/tools';
import type { AddCoinServerParam, DelCoinServerParam } from './api.d';

export const GetServerList = () => {
  return ajax_json({
    url: '/api/private/server_list',
    data: null,
    method: 'get',
  });
};

export const CreateServer = (param: AddCoinServerParam) => {
  const data = {
    ...param,
    Password: Md5(param.Password),
  };

  return ajax_json({
    url: '/api/private/add_server',
    data,
    method: 'post',
  });
};

export const GetDeployShell = (data: { CoinServeID: string; Password: string }) => {
  const param = {
    ...data,
  };
  if (data.Password) {
    param.Password = Md5(data.Password);
  }

  return ajax_json({
    url: '/api/private/get_deploy_shell',
    data: param,
    method: 'post',
  });
};

export const DelServer = (param: DelCoinServerParam) => {
  const data = {
    ...param,
    Password: Md5(param.Password),
  };
  return ajax_json({
    url: '/api/private/del_server',
    data,
    method: 'post',
  });
};
