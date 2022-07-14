import { ajax_json } from '@/utils/http';
import { Md5 } from '@/utils/tools';

export const GetServerList = () => {
  return ajax_json({
    url: '/api/private/server_list',
    data: null,
    method: 'get',
  });
};

interface AddAIFundServerParam {
  OkxKeyID: string;
  Host: string;
  Port: string;
  Note: string;
  Password: string;
}
export const CreateServer = (param: AddAIFundServerParam) => {
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

export const GetDeployShell = (data: { AIFundServerID: string; Password: string }) => {
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

interface DelServerParam {
  AIFundServerID: string;
  Password: string;
}

export const DelServer = (param: DelServerParam) => {
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
